package contracts

import (
	//"math/big"

	"github.com/AthanorLabs/go-relayer/common"
	//ethcommon "github.com/ethereum/go-ethereum/common"
)

func (r *IMinimalForwarderForwardRequest) FromSubmitTransactionRequest(
	req *common.SubmitTransactionRequest,
) {
	r.From = req.From
	r.To = req.To
	r.Value = req.Value
	r.Gas = req.Gas
	r.Nonce = req.Nonce
	r.Data = req.Data
}

func NewIMinimalForwarderForwardRequest() common.ForwardRequest {
	return &IMinimalForwarderForwardRequest{}
}
