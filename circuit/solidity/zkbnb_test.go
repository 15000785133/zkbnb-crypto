/*
 * Copyright © 2022 ZkBNB Protocol
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package solidity

import (
	"bufio"
	"encoding/gob"
	"flag"
	"fmt"
	"github.com/consensys/gnark/constraint/lazy"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"

	"github.com/bnb-chain/zkbnb-crypto/circuit"
)

var optionalBlockSizes = flag.String("blocksizes", "1,10", "block size that will be used for proof generation and verification")

func TestCompileCircuit(t *testing.T) {
	differentBlockSizes := optionalBlockSizesInt()
	gasAssetIds := []int64{0, 1}
	gasAccountIndex := int64(1)
	for i := 0; i < len(differentBlockSizes); i++ {
		var blockConstraints circuit.BlockConstraints
		blockConstraints.TxsCount = differentBlockSizes[i]
		blockConstraints.Txs = make([]circuit.TxConstraints, blockConstraints.TxsCount)
		for i := 0; i < blockConstraints.TxsCount; i++ {
			blockConstraints.Txs[i] = circuit.GetZeroTxConstraint()
		}
		blockConstraints.GasAssetIds = gasAssetIds
		blockConstraints.GasAccountIndex = gasAccountIndex
		blockConstraints.Gas = circuit.GetZeroGasConstraints(gasAssetIds)
		oR1cs, err := frontend.Compile(ecc.BN254.ScalarField(), r1cs.NewBuilder, &blockConstraints, frontend.IgnoreUnconstrainedInputs())
		if err != nil {
			panic(err)
		}
		fmt.Printf("Number of constraints: %d\n", oR1cs.GetNbConstraints())
	}
}

func TestExportSol(t *testing.T) {
	exportSol(optionalBlockSizesInt())
}

func TestExportSolSmall(t *testing.T) {
	differentBlockSizes := []int{32}
	exportSol(differentBlockSizes)
}

func exportSol(differentBlockSizes []int) {
	gasAssetIds := []int64{0, 1}
	gasAccountIndex := int64(1)
	gob.Register(lazy.GeneralLazyInputs{})
	sessionName := "zkbnb"
	for i := 0; i < len(differentBlockSizes); i++ {
		var blockConstraints circuit.BlockConstraints
		blockConstraints.TxsCount = differentBlockSizes[i]
		blockConstraints.Txs = make([]circuit.TxConstraints, blockConstraints.TxsCount)
		for i := 0; i < blockConstraints.TxsCount; i++ {
			blockConstraints.Txs[i] = circuit.GetZeroTxConstraint()
		}
		blockConstraints.GasAssetIds = gasAssetIds
		blockConstraints.GasAccountIndex = gasAccountIndex
		blockConstraints.Gas = circuit.GetZeroGasConstraints(gasAssetIds)

		oR1cs, err := frontend.Compile(ecc.BN254.ScalarField(), r1cs.NewBuilder, &blockConstraints, frontend.IgnoreUnconstrainedInputs())
		fmt.Printf("Constraints num=%v\n", oR1cs.GetNbConstraints())
		if err != nil {
			panic(err)
		}
		// pk, vk, err := groth16.Setup(oR1cs)
		f, err := os.Create("zkbnb" + fmt.Sprint(differentBlockSizes[i]) + ".dump")
		writer := bufio.NewWriter(f)
		enc := gob.NewEncoder(writer)
		err = enc.Encode(oR1cs)
		if err != nil {
			panic(err)
		}
		err = f.Close()
		if err != nil {
			panic(err)
		}
		pk, vk, err := groth16.Setup(oR1cs)
		if err != nil {
			panic(err)
		}
		err = groth16.SplitDumpPK(pk, sessionName+fmt.Sprint(differentBlockSizes[i]))
		if err != nil {
			f, err := os.Create("ZkBNBVerifier" + fmt.Sprint(differentBlockSizes[i]) + ".sol")
			if err != nil {
				panic(err)
			}
			err = vk.ExportSolidity(f)
			if err != nil {
				panic(err)
			}
		}
	}
}

func optionalBlockSizesInt() []int {
	blockSizesStr := strings.Split(*optionalBlockSizes, ",")
	blockSizesInt := make([]int, len(blockSizesStr))
	for i := range blockSizesStr {
		v, err := strconv.Atoi(blockSizesStr[i])
		if err != nil {
			panic(err)
		}
		blockSizesInt[i] = v
	}
	return blockSizesInt
}
