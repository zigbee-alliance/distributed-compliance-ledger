// Copyright 2020 DSR Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package compliancetest

import (
	"fmt"

	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliancetest/internal/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

type GenesisState struct {
	TestingResultRecords []TestingResults `json:"testing_result_records"`
}

func NewGenesisState() GenesisState {
	return GenesisState{TestingResultRecords: []TestingResults{}}
}

func ValidateGenesis(data GenesisState) error {
	for _, record := range data.TestingResultRecords {
		if record.VID == 0 {
			return sdk.ErrUnknownRequest(fmt.Sprintf("Invalid TestingResultRecord: value: %v. "+
				"Error: Invalid VID", record.VID))
		}

		if record.PID == 0 {
			return sdk.ErrUnknownRequest(fmt.Sprintf("Invalid TestingResultRecord: value: %v. "+
				"Error: Invalid PID", record.PID))
		}

		if len(record.Results) == 0 {
			return sdk.ErrUnknownRequest(fmt.Sprintf("Invalid TestingResultRecord: value: %s. "+
				"Error: Missing TestResult", record.Results))
		}
	}

	return nil
}

func DefaultGenesisState() GenesisState {
	return NewGenesisState()
}

func InitGenesis(ctx sdk.Context, keeper Keeper, data GenesisState) []abci.ValidatorUpdate {
	for _, record := range data.TestingResultRecords {
		keeper.SetTestingResults(ctx, record)
	}

	return []abci.ValidatorUpdate{}
}

func ExportGenesis(ctx sdk.Context, k Keeper) GenesisState {
	var records []TestingResults

	k.IterateTestingResults(ctx, func(testingResult types.TestingResults) (stop bool) {
		records = append(records, testingResult)

		return false
	})

	return GenesisState{TestingResultRecords: records}
}
