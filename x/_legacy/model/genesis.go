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

package model

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/model/internal/types"
)

type GenesisState struct {
	ModelRecords []Model `json:"model_records"`
}

func NewGenesisState() GenesisState {
	return GenesisState{ModelRecords: []Model{}}
}

//nolint:cognit
func ValidateGenesis(data GenesisState) error {
	for _, record := range data.ModelRecords {
		if record.VID == 0 {
			return sdk.ErrUnknownRequest(fmt.Sprintf("Invalid Model: Invalid VID. Value: %v", record))
		}

		if record.PID == 0 {
			return sdk.ErrUnknownRequest(fmt.Sprintf("Invalid Model: Invalid PID. Value: %v", record))
		}

		if record.ProductName == "" {
			return sdk.ErrUnknownRequest(fmt.Sprintf("Invalid Model: Missed ProductName. Value: %v", record))
		}

		if record.ProductLabel == "" {
			return sdk.ErrUnknownRequest(fmt.Sprintf("Invalid Model: Missed ProductLabel. Value: %v", record))
		}

		if record.PartNumber == "" {
			return sdk.ErrUnknownRequest(fmt.Sprintf("Invalid Model: Missed PartNumber. Value: %v", record))
		}

	}

	return nil
}

func DefaultGenesisState() GenesisState {
	return NewGenesisState()
}

func InitGenesis(ctx sdk.Context, keeper Keeper, data GenesisState) []abci.ValidatorUpdate {
	for _, record := range data.ModelRecords {
		keeper.SetModel(ctx, record)
	}

	return []abci.ValidatorUpdate{}
}

func ExportGenesis(ctx sdk.Context, k Keeper) GenesisState {
	var records []Model

	k.IterateModels(ctx, func(model types.Model) (stop bool) {
		records = append(records, model)

		return false
	})

	return GenesisState{ModelRecords: records}
}
