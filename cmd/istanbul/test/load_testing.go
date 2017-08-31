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
	"errors"
	"fmt"
	"log"
	"os/exec"

	"github.com/getamis/istanbul-tools/cmd/istanbul/test/charts"
	"github.com/getamis/istanbul-tools/cmd/utils"
	"github.com/urfave/cli"
)

func runLoadTesting(ctx *cli.Context) error {
	cleanDeployments()

	numOfValidators := int(ctx.Int64(validatorSizeFlag.Name))
	if numOfValidators <= 0 {
		return errors.New("number of validator should be larger than zero")
	}

	_, _, addrs := utils.GenerateKeys(numOfValidators)

	gasLimit := ctx.Uint64(gasLimitFlag.Name)

	charts := []ChartInstaller{
		charts.NewGenesisChart(addrs, gasLimit),
	}

	for _, chart := range charts {
		if err := chart.Install(false); err != nil {
			return err
		}
	}

	return nil
}

func cleanDeployments() {
	cmd := "helm list | tail -n +2 | awk '{print $1}' | xargs -I {} helm delete --purge {}"
	output, err := exec.Command(
		"bash", "-c", cmd).Output()
	if err != nil {
		log.Fatalf("Failed to execute command: %s\n", cmd)
	}
	fmt.Println(string(output))
}
