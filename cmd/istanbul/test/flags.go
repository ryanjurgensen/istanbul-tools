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

package test

import (
	"github.com/urfave/cli"

	"github.com/getamis/istanbul-tools/genesis"
)

var (
	chartFlag = cli.StringFlag{
		Name:  "chart",
		Usage: "Chart to deploy",
	}

	validatorSizeFlag = cli.Int64Flag{
		Name:   "validator_size",
		Usage:  "The number of validators",
		EnvVar: "VALIDATOR_SIZE",
		Value:  4,
	}

	gasLimitFlag = cli.Uint64Flag{
		Name:   "gaslimit",
		Usage:  "Gas limit per block",
		EnvVar: "GAS_LIMIT",
		Value:  genesis.InitGasLimit,
	}
)
