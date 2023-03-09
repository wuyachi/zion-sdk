package side_chain_manager

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
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

type Fee struct {
	View uint64
	Fee  *big.Int
}

type UpdateFeeParam struct {
	ChainID   uint64
	ViewNum   uint64
	Fee       *big.Int
	Signature []byte
}

//Digest Digest calculate the hash of param input
func (m *UpdateFeeParam) Digest() ([]byte, error) {
	input := &UpdateFeeParam{
		ChainID: m.ChainID,
		ViewNum: m.ViewNum,
		Fee:     m.Fee,
	}
	msg, err := rlp.EncodeToBytes(input)
	if err != nil {
		return nil, fmt.Errorf("UpdateFeeParam, serialize input error: %v", err)
	}
	digest := crypto.Keccak256(msg)
	return digest, nil
}
