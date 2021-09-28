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

package compliance

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliance/internal/types"
)

type GenesisState struct {
	ComplianceInfoRecords []ComplianceInfo `json:"compliance_model_records"`
}

func NewGenesisState() GenesisState {
	return GenesisState{ComplianceInfoRecords: []ComplianceInfo{}}
}

func ValidateGenesis(data GenesisState) error {
	for _, record := range data.ComplianceInfoRecords {
		if record.VID == 0 {
			return sdk.ErrUnknownRequest(
				fmt.Sprintf("Invalid CertifiedModelRecord: value: %d. "+
					"Error: Invalid VID: it cannot be 0", record.VID))
		}

		if record.PID == 0 {
			return sdk.ErrUnknownRequest(
				fmt.Sprintf("Invalid CertifiedModelRecord: value: %d. "+
					"Error: Invalid PID: it cannot be 0", record.PID))
		}

		if record.SoftwareVersionCertificationStatus > types.CodeRevoked {
			return sdk.ErrUnknownRequest(
				fmt.Sprintf("Invalid CertifiedModelRecord: value: %d."+
					" Error: Invalid SoftwareVersionCertificationStatus: It should be either 0,1,2 or 3", record.PID))
		}

		if record.Date.IsZero() {
			return sdk.ErrUnknownRequest("Invalid Date: it cannot be empty")
		}

		if record.CertificationType != "" && record.CertificationType != types.ZbCertificationType {
			return sdk.ErrUnknownRequest(
				fmt.Sprintf("Invalid CertifiedModelRecord: value: %v."+
					" Error: Invalid CertificationType: "+
					"unknown type; supported types: [%s]", record.CertificationType, types.ZbCertificationType))
		}
	}

	return nil
}

func DefaultGenesisState() GenesisState {
	return NewGenesisState()
}

func InitGenesis(ctx sdk.Context, keeper Keeper, data GenesisState) []abci.ValidatorUpdate {
	for _, record := range data.ComplianceInfoRecords {
		keeper.SetComplianceInfo(ctx, record)
	}

	return []abci.ValidatorUpdate{}
}

func ExportGenesis(ctx sdk.Context, k Keeper) GenesisState {
	var records []ComplianceInfo

	k.IterateComplianceInfos(ctx, "", func(deviceCompliance types.ComplianceInfo) (stop bool) {
		records = append(records, deviceCompliance)

		return false
	})

	return GenesisState{ComplianceInfoRecords: records}
}
