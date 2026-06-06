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

package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliance/types"
)

// SetDeviceSoftwareCompliance set a specific deviceSoftwareCompliance in the store from its index.
func (k Keeper) SetDeviceSoftwareCompliance(ctx sdk.Context, deviceSoftwareCompliance types.DeviceSoftwareCompliance) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DeviceSoftwareComplianceKeyPrefix))
	b := k.cdc.MustMarshal(&deviceSoftwareCompliance)
	store.Set(types.DeviceSoftwareComplianceKey(
		deviceSoftwareCompliance.CDCertificateId,
	), b)
}

// GetDeviceSoftwareCompliance returns a deviceSoftwareCompliance from its index.
func (k Keeper) GetDeviceSoftwareCompliance(
	ctx sdk.Context,
	cDCertificateID string,
) (val types.DeviceSoftwareCompliance, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DeviceSoftwareComplianceKeyPrefix))

	b := store.Get(types.DeviceSoftwareComplianceKey(
		cDCertificateID,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)

	return val, true
}

// RemoveDeviceSoftwareCompliance removes a deviceSoftwareCompliance from the store.
func (k Keeper) RemoveDeviceSoftwareCompliance(
	ctx sdk.Context,
	cDCertificateID string,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DeviceSoftwareComplianceKeyPrefix))
	store.Delete(types.DeviceSoftwareComplianceKey(
		cDCertificateID,
	))
}

// GetAllDeviceSoftwareCompliance returns all deviceSoftwareCompliance.
func (k Keeper) GetAllDeviceSoftwareCompliance(ctx sdk.Context) (list []types.DeviceSoftwareCompliance) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DeviceSoftwareComplianceKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer func() { _ = iterator.Close() }()

	for ; iterator.Valid(); iterator.Next() {
		var val types.DeviceSoftwareCompliance
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
