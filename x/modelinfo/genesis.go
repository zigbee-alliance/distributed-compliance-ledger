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

func ValidateGenesis(data GenesisState) error {
	for _, record := range data.ModelInfoRecords {
		if record.VID == 0 {
			return sdk.ErrUnknownRequest(fmt.Sprintf("Invalid ModelInfo: Invalid VID. Value: %v", record))
		}

		if record.PID == 0 {
			return sdk.ErrUnknownRequest(fmt.Sprintf("Invalid ModelInfo: Invalid PID. Value: %v", record))
		}

		if record.Name == "" {
			return sdk.ErrUnknownRequest(fmt.Sprintf("Invalid ModelInfo: Missed Name. Value: %v", record))
		}

		if record.Description == "" {
			return sdk.ErrUnknownRequest(fmt.Sprintf("Invalid ModelInfo: Missed Description. Value: %v", record))
		}

		if record.SKU == "" {
			return sdk.ErrUnknownRequest(fmt.Sprintf("Invalid ModelInfo: Missed SKU. Value: %v", record))
		}

		if record.HardwareVersion == "" {
			return sdk.ErrUnknownRequest(fmt.Sprintf("Invalid ModelInfo: Missed HardwareVersion. Value: %v", record))
		}

		if record.FirmwareVersion == "" {
			return sdk.ErrUnknownRequest(fmt.Sprintf("Invalid ModelInfo: Missed FirmwareVersion. Value: %v", record))
		}

		if record.Owner.Empty() {
			return sdk.ErrUnknownRequest(fmt.Sprintf("Invalid ModelInfo: Missed Owner. Value: %v", record))
		}

		if record.OtaURL != "" || record.OtaChecksum != "" || record.OtaChecksumType != "" {
			if record.OtaURL == "" || record.OtaChecksum == "" || record.OtaChecksumType == "" {
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
