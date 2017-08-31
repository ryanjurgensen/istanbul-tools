// Copyright 2017 AMIS Technologies
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package genesis

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/big"
	"path/filepath"
	"time"

	"github.com/ethereum/go-ethereum/consensus/istanbul"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
)

const (
	FileName       = "genesis.json"
	InitGasLimit   = 4700000
	InitDifficulty = 1
)

func New(options ...Option) *core.Genesis {
	genesis := &core.Genesis{
		Timestamp:  uint64(time.Now().Unix()),
		GasLimit:   InitGasLimit,
		Difficulty: big.NewInt(InitDifficulty),
		Alloc:      make(core.GenesisAlloc),
		Config: &params.ChainConfig{
			HomesteadBlock: big.NewInt(1),
			EIP150Block:    big.NewInt(2),
			EIP155Block:    big.NewInt(3),
			EIP158Block:    big.NewInt(3),
			Istanbul: &params.IstanbulConfig{
				ProposerPolicy: uint64(istanbul.DefaultConfig.ProposerPolicy),
				Epoch:          istanbul.DefaultConfig.Epoch,
			},
		},
		Mixhash: types.IstanbulDigest,
	}

	for _, opt := range options {
		opt(genesis)
	}

	return genesis
}

func NewFile(options ...Option) string {
	dir, err := generateRandomDir()
	if err != nil {
		log.Fatalf("Failed to create random directory, err: %v", err)
	}
	genesis := New(options...)
	if err := Save(dir, genesis); err != nil {
		log.Fatalf("Failed to save genesis to '%s', err: %v", dir, err)
	}

	return filepath.Join(dir, FileName)
}

func NewFileAt(dir string, options ...Option) string {
	genesis := New(options...)
	if err := Save(dir, genesis); err != nil {
		log.Fatalf("Failed to save genesis to '%s', err: %v", dir, err)
	}

	return filepath.Join(dir, FileName)
}

func Save(dataDir string, genesis *core.Genesis) error {
	filePath := filepath.Join(dataDir, FileName)

	raw, err := json.Marshal(genesis)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filePath, raw, 0600)
}
