/*
 * Copyright (C) 2021 The Zion Authors
 * This file is part of The Zion library.
 *
 * The Zion is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The Zion is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public License
 * along with The Zion.  If not, see <http://www.gnu.org/licenses/>.
 */

package node_manager

import (
	"fmt"
	"strings"

	"github.com/devfans/zion-sdk/contracts/native/utils"
	"github.com/devfans/zion-sdk/contracts/native/go_abi/node_manager_abi"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

const contractName = "node manager"

func InitABI() {
	ab, err := abi.JSON(strings.NewReader(node_manager_abi.INodeManagerABI))
	if err != nil {
		panic(fmt.Sprintf("failed to load abi json string: [%v]", err))
	}
	ABI = &ab
}

var (
	ABI  *abi.ABI
	this = utils.NodeManagerContractAddress
)

type MethodContractNameInput struct{}

func (m *MethodContractNameInput) Decode(payload []byte) error { return nil }

type MethodContractNameOutput struct {
	Name string
}

type MethodProposeInput struct {
	StartHeight uint64
	Peers       *Peers
}

type MethodProposeOutput struct {
	Success bool
}


type MethodVoteInput struct {
	EpochID   uint64
	EpochHash common.Hash
}

type MethodVoteOutput struct {
	Success bool
}


// useless input
type MethodEpochInput struct{}

func (m *MethodEpochInput) Decode(payload []byte) error { return nil }

type MethodEpochOutput struct {
	Epoch *EpochInfo
}

// useless input
type MethodGetChangingEpochInput struct{}

func (m *MethodGetChangingEpochInput) Decode(payload []byte) error { return nil }

type MethodGetEpochByIDInput struct {
	EpochID uint64
}


type MethodProofInput struct {
	EpochID uint64
}


type MethodProofOutput struct {
	Hash common.Hash
}
