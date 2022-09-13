package contracts

import (
	"math/big"

	ethcommon "github.com/ethereum/go-ethereum/common"
)

func NewIMinimalForwarderForwardRequest(
	from ethcommon.Address,
	to ethcommon.Address,
	value *big.Int,
	gas *big.Int,
	nonce *big.Int,
	data []byte,
	_ *big.Int,
) IMinimalForwarderForwardRequest {
	return IMinimalForwarderForwardRequest{
		From:  from,
		To:    to,
		Value: value,
		Gas:   gas,
		Nonce: nonce,
		Data:  data,
	}
}
