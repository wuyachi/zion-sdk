package common

import (
	"encoding/hex"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	EmptyHash    = common.Hash{}
	EmptyAddress = common.Address{}
)

func MappingKeyAt(position1 string, position2 string) ([]byte, error) {
	p1, err := hex.DecodeString(position1)
	if err != nil {
		return nil, err
	}

	p2, err := hex.DecodeString(position2)

	if err != nil {
		return nil, err
	}

	key := crypto.Keccak256(common.LeftPadBytes(p1, 32), common.LeftPadBytes(p2, 32))

	return key, nil
}