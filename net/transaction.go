package net

import (
	"context"
	"fmt"
	"time"

	libp2pnetwork "github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/protocol"

	net "github.com/athanorlabs/go-p2p-net"
)

const (
	transactionID      = "/transaction/0"
	transactionTimeout = time.Second * 5
)

func (h *Host) handleTransactionStream(stream libp2pnetwork.Stream) {
	msg, err := readStreamMessage(stream, maxMessageSize)
	if err != nil {
		log.Debugf("error reading TransactionRequest: %w", err)
	}

	req, ok := msg.(*TransactionRequest)
	if !ok {
		return
	}

	resp, err := h.handleTransaction(req.SubmitTransactionRequest)
	if err != nil {
		log.Debugf("failed to handle transaction rrequest: %w", err)
	}

	msgResp := &TransactionResponse{
		SubmitTransactionResponse: resp,
	}

	if err := net.WriteStreamMessage(stream, msgResp, stream.Conn().RemotePeer()); err != nil {
		log.Warnf("failed to send TransactionResponse message to peer: %s", err)
	}

	_ = stream.Close()
}

func (h *Host) SubmitTransaction(who peer.ID, msg *TransactionRequest) (*TransactionResponse, error) {
	ctx, cancel := context.WithTimeout(h.ctx, transactionTimeout)
	defer cancel()

	if err := h.h.Connect(ctx, peer.AddrInfo{ID: who}); err != nil {
		return nil, err
	}

	stream, err := h.h.NewStream(ctx, who, protocol.ID(transactionID))
	if err != nil {
		return nil, fmt.Errorf("failed to open stream with peer: %w", err)
	}

	log.Debugf("opened query stream: %s", stream.Conn())
	defer func() {
		_ = stream.Close()
	}()

	return submitTransaction(msg, stream)
}

func submitTransaction(msg *TransactionRequest, stream libp2pnetwork.Stream) (*TransactionResponse, error) {
	err := net.WriteStreamMessage(stream, msg, stream.Conn().RemotePeer())
	if err != nil {
		return nil, err
	}

	return receiveTransactionResponse(stream)
}

func receiveTransactionResponse(stream libp2pnetwork.Stream) (*TransactionResponse, error) {
	msg, err := readStreamMessage(stream, maxMessageSize)
	if err != nil {
		return nil, fmt.Errorf("error reading TransactionResponse: %w", err)
	}

	resp, ok := msg.(*TransactionResponse)
	if !ok {
		return nil, fmt.Errorf("expected message with type %d but received %d",
			TransactionResponseType,
			msg.Type())
	}

	return resp, nil
}
