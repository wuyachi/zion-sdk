package side_chain_manager

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type RippleExtraInfo struct {
	Operator      common.Address
	Sequence      uint64
	Quorum        uint64
	SignerNum     uint64
	Pks           [][]byte
	ReserveAmount *big.Int
}
