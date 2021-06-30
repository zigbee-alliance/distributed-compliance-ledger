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

package modelinfo

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/modelinfo/internal/types"
)

type GenesisState struct {
	ModelInfoRecords []ModelInfo `json:"model_info_records"`
}

func NewGenesisState() GenesisState {
	return GenesisState{ModelInfoRecords: []ModelInfo{}}
}

//nolint:cognit
func ValidateGenesis(data GenesisState) error {
	for _, record := range data.ModelInfoRecords {
		if record.Model.VID == 0 {
			return sdk.ErrUnknownRequest(fmt.Sprintf("Invalid ModelInfo: Invalid VID. Value: %v", record))
		}

		if record.Model.PID == 0 {
			return sdk.ErrUnknownRequest(fmt.Sprintf("Invalid ModelInfo: Invalid PID. Value: %v", record))
		}

		if record.Model.ProductName == "" {
			return sdk.ErrUnknownRequest(fmt.Sprintf("Invalid ModelInfo: Missed ProductName. Value: %v", record))
		}

		if record.Model.Description == "" {
			return sdk.ErrUnknownRequest(fmt.Sprintf("Invalid ModelInfo: Missed Description. Value: %v", record))
		}

		if record.Model.SKU == "" {
			return sdk.ErrUnknownRequest(fmt.Sprintf("Invalid ModelInfo: Missed SKU. Value: %v", record))
		}

		if record.Model.HardwareVersion == 0 {
			return sdk.ErrUnknownRequest(fmt.Sprintf("Invalid ModelInfo: Missed HardwareVersion. Value: %v", record))
		}

		if record.Model.HardwareVersionString == "" {
			return sdk.ErrUnknownRequest(fmt.Sprintf("Invalid ModelInfo: Missed HardwareVersionString. Value: %v", record))
		}

		if record.Model.SoftwareVersion == 0 {
			return sdk.ErrUnknownRequest(fmt.Sprintf("Invalid ModelInfo: Missed SoftwareVersion. Value: %v", record))
		}

		if record.Model.SoftwareVersionString == "" {
			return sdk.ErrUnknownRequest(fmt.Sprintf("Invalid ModelInfo: Missed SoftwareVersionString. Value: %v", record))
		}

		if record.Owner.Empty() {
			return sdk.ErrUnknownRequest(fmt.Sprintf("Invalid ModelInfo: Missed Owner. Value: %v", record))
		}

		if record.Model.OtaURL != "" || record.Model.OtaChecksum != "" || record.Model.OtaChecksumType != "" {
			if record.Model.OtaURL == "" || record.Model.OtaChecksum == "" || record.Model.OtaChecksumType == "" {
				return sdk.ErrUnknownRequest(fmt.Sprintf("Invalid ModelInfo: The fields OtaURL, OtaChecksum and "+
					"OtaChecksumType must be either specified together, or not specified together. Value: %v", record))
			}
		}
	}

	return nil
}

func DefaultGenesisState() GenesisState {
	return NewGenesisState()
}

func InitGenesis(ctx sdk.Context, keeper Keeper, data GenesisState) []abci.ValidatorUpdate {
	for _, record := range data.ModelInfoRecords {
		keeper.SetModelInfo(ctx, record)
	}

	return []abci.ValidatorUpdate{}
}

func ExportGenesis(ctx sdk.Context, k Keeper) GenesisState {
	var records []ModelInfo

	k.IterateModelInfos(ctx, func(modelInfo types.ModelInfo) (stop bool) {
		records = append(records, modelInfo)

		return false
	})

	return GenesisState{ModelInfoRecords: records}
}
