package net

import (
	"context"
	"time"

	logging "github.com/ipfs/go-log"
	libp2pnetwork "github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/protocol"

	net "github.com/athanorlabs/go-p2p-net"
)

const (
	// ProtocolID is the base relayer network protocol ID.
	ProtocolID           = "/relayer/0.1"
	advertisedRelayerStr = "isrelayer"
	maxMessageSize       = 1 << 17
)

var log = logging.Logger("net")

// NetHost contains libp2p functionality used by the Host.
type NetHost interface {
	Start() error
	Stop() error

	Advertise([]string)
	Discover(provides string, searchTime time.Duration) ([]peer.ID, error)

	SetStreamHandler(string, func(libp2pnetwork.Stream))
	SetShouldAdvertiseFunc(net.ShouldAdvertiseFunc)

	Connectedness(peer.ID) libp2pnetwork.Connectedness
	Connect(context.Context, peer.AddrInfo) error
	NewStream(context.Context, peer.ID, protocol.ID) (libp2pnetwork.Stream, error)

	AddrInfo() peer.AddrInfo
	Addresses() []string
	PeerID() peer.ID
	ConnectedPeers() []string
}

type handleTransactionFunc = func(*TransactionRequest) (*TransactionResponse, error)

// Host represents a p2p node that implements the atomic swap protocol.
type Host struct {
	ctx               context.Context
	h                 NetHost
	handleTransaction handleTransactionFunc
}

// NewHost returns a new Host.
func NewHost(cfg *net.Config) (*Host, error) {
	h, err := net.NewHost(cfg)
	if err != nil {
		return nil, err
	}

	return &Host{
		ctx: cfg.Ctx,
		h:   h,
	}, nil
}

// Start starts the bootstrap and discovery process.
func (h *Host) Start() error {
	return h.h.Start()
}

func (h *Host) HandleTransactions() {
	h.h.SetStreamHandler(transactionID, h.handleTransactionStream)
	// TODO: also advertise
}

// Stop stops the host.
func (h *Host) Stop() error {
	return h.h.Stop()
}

// Advertise advertises in the DHT.
func (h *Host) Advertise() {
	h.h.Advertise([]string{advertisedRelayerStr})
}

// Discover searches the DHT for peers that advertise that they provide the given coin..
// It searches for up to `searchTime` duration of time.
func (h *Host) Discover(searchTime time.Duration) ([]peer.ID, error) {
	return h.h.Discover(advertisedRelayerStr, searchTime)
}

// AddrInfo returns the host's AddrInfo.
func (h *Host) AddrInfo() peer.AddrInfo {
	return h.h.AddrInfo()
}

// Addresses returns the list of multiaddress the host is listening on.
func (h *Host) Addresses() []string {
	return h.h.Addresses()
}

// ConnectedPeers returns the multiaddresses of our currently connected peers.
func (h *Host) ConnectedPeers() []string {
	return h.h.ConnectedPeers()
}

// PeerID returns the host's peer ID.
func (h *Host) PeerID() peer.ID {
	return h.h.AddrInfo().ID
}

func readStreamMessage(stream libp2pnetwork.Stream, maxMessageSize uint32) (Message, error) {
	msgBytes, err := net.ReadStreamMessage(stream, maxMessageSize)
	if err != nil {
		return nil, err
	}

	return decodeMessage(msgBytes)
}
