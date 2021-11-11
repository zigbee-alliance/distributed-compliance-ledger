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

package vendorinfo

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/validator"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/vendorinfo/internal/types"
)

type GenesisState struct {
	VendorInfoRecords []VendorInfo `json:"vendor_info_records"`
}

func NewGenesisState() GenesisState {
	return GenesisState{VendorInfoRecords: []VendorInfo{}}
}

//nolint:cognit
func ValidateGenesis(data GenesisState) error {
	for _, record := range data.VendorInfoRecords {
		err := validator.ValidateAdd(record)
		if err != nil {
			return err
		}
	}
	return nil
}

func DefaultGenesisState() GenesisState {
	return NewGenesisState()
}

func InitGenesis(ctx sdk.Context, keeper Keeper, data GenesisState) []abci.ValidatorUpdate {
	for _, record := range data.VendorInfoRecords {
		keeper.SetVendorInfo(ctx, record)
	}

	return []abci.ValidatorUpdate{}
}

func ExportGenesis(ctx sdk.Context, k Keeper) GenesisState {
	var records []VendorInfo

	k.IterateVendorInfos(ctx, func(model types.VendorInfo) (stop bool) {
		records = append(records, model)

		return false
	})

	return GenesisState{VendorInfoRecords: records}
}
