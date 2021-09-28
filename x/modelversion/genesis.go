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

package modelversion

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/modelversion/internal/types"
)

type GenesisState struct {
	ModelVersionRecords []ModelVersion `json:"model_version_records"`
}

func NewGenesisState() GenesisState {
	return GenesisState{ModelVersionRecords: []ModelVersion{}}
}

//nolint:cognit
func ValidateGenesis(data GenesisState) error {
	for _, record := range data.ModelVersionRecords {
		if record.VID == 0 {
			return sdk.ErrUnknownRequest(fmt.Sprintf("Invalid Model: Invalid VID. Value: %v", record))
		}

		if record.PID == 0 {
			return sdk.ErrUnknownRequest(fmt.Sprintf("Invalid Model: Invalid PID. Value: %v", record))
		}

		if record.SoftwareVersion == 0 {
			return sdk.ErrUnknownRequest(fmt.Sprintf("Invalid Model: Missed SoftwareVersion. Value: %v", record))
		}

		if record.OtaURL != "" || record.OtaChecksum != "" || record.OtaChecksumType != 0 {
			if record.OtaURL == "" || record.OtaChecksum == "" || record.OtaChecksumType == 0 {
				return sdk.ErrUnknownRequest(fmt.Sprintf("Invalid Model: The fields OtaURL, OtaChecksum and "+
					"OtaChecksumType must be either specified together, or not specified together. Value: %v", record))
			}
		}
		if record.OtaURL != "" || record.OtaFileSize != 0 || record.OtaChecksum != "" || record.OtaChecksumType != 0 {
			if record.OtaURL == "" || record.OtaFileSize == 0 || record.OtaChecksum != "" || record.OtaChecksumType == 0 {
				return sdk.ErrUnknownRequest(fmt.Sprintf("Invalid ModelVersion: the fields OtaURL, OtaFileSize, OtaChecksum and "+
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
	for _, record := range data.ModelVersionRecords {
		keeper.SetModelVersion(ctx, record)
	}

	return []abci.ValidatorUpdate{}
}

func ExportGenesis(ctx sdk.Context, k Keeper) GenesisState {
	var records []ModelVersion

	k.IterateModelVersions(ctx, func(modelVersionInfo types.ModelVersion) (stop bool) {
		records = append(records, modelVersionInfo)

		return false
	})

	return GenesisState{ModelVersionRecords: records}
}
