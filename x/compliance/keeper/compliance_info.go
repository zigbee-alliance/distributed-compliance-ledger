// Copyright 2022 DSR Corporation
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

package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliance/types"
)

// SetComplianceInfo set a specific complianceInfo in the store from its index
func (k Keeper) SetComplianceInfo(ctx sdk.Context, complianceInfo types.ComplianceInfo) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ComplianceInfoKeyPrefix))
	b := k.cdc.MustMarshal(&complianceInfo)
	store.Set(types.ComplianceInfoKey(
		complianceInfo.Vid,
		complianceInfo.Pid,
		complianceInfo.SoftwareVersion,
		complianceInfo.CertificationType,
	), b)
}

// GetComplianceInfo returns a complianceInfo from its index
func (k Keeper) GetComplianceInfo(
	ctx sdk.Context,
	vid int32,
	pid int32,
	softwareVersion uint32,
	certificationType string,

) (val types.ComplianceInfo, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ComplianceInfoKeyPrefix))

	b := store.Get(types.ComplianceInfoKey(
		vid,
		pid,
		softwareVersion,
		certificationType,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveComplianceInfo removes a complianceInfo from the store
func (k Keeper) RemoveComplianceInfo(
	ctx sdk.Context,
	vid int32,
	pid int32,
	softwareVersion uint32,
	certificationType string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ComplianceInfoKeyPrefix))
	store.Delete(types.ComplianceInfoKey(
		vid,
		pid,
		softwareVersion,
		certificationType,
	))
}

// GetAllComplianceInfo returns all complianceInfo
func (k Keeper) GetAllComplianceInfo(ctx sdk.Context) (list []types.ComplianceInfo) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ComplianceInfoKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ComplianceInfo
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
