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
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliance/internal/types"
)

type Keeper struct {
	// Unexposed key to access store from sdk.Context.
	storeKey sdk.StoreKey

	// The wire codec for binary encoding/decoding.
	cdc *codec.Codec
}

func NewKeeper(storeKey sdk.StoreKey, cdc *codec.Codec) Keeper {
	return Keeper{storeKey: storeKey, cdc: cdc}
}

// Gets the entire ComplianceInfo struct for a ComplianceInfoID.
func (k Keeper) GetComplianceInfo(ctx sdk.Context, certificationType types.CertificationType,
	vid uint16, pid uint16, softwareVersion uint32) types.ComplianceInfo {
	if !k.IsComplianceInfoPresent(ctx, certificationType, vid, pid, softwareVersion) {
		panic("ComplianceInfo does not exist")
	}

	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetComplianceInfoKey(certificationType, vid, pid, softwareVersion))

	var device types.ComplianceInfo

	k.cdc.MustUnmarshalBinaryBare(bz, &device)

	return device
}

// Sets the entire ComplianceInfo metadata struct for a ComplianceInfoID.
func (k Keeper) SetComplianceInfo(ctx sdk.Context, model types.ComplianceInfo) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetComplianceInfoKey(
		model.CertificationType, model.VID, model.PID, model.SoftwareVersion), k.cdc.MustMarshalBinaryBare(model))
}

// Iterate over all ComplianceInfos.
func (k Keeper) IterateComplianceInfos(ctx sdk.Context, certificationType types.CertificationType,
	process func(info types.ComplianceInfo) (stop bool)) {
	store := ctx.KVStore(k.storeKey)

	iter := sdk.KVStorePrefixIterator(store, types.GetCertificationPrefix(certificationType))
	defer iter.Close()

	for {
		if !iter.Valid() {
			return
		}

		val := iter.Value()

		var certifiedModel types.ComplianceInfo

		k.cdc.MustUnmarshalBinaryBare(val, &certifiedModel)

		if process(certifiedModel) {
			return
		}

		iter.Next()
	}
}

func (k Keeper) CountTotalComplianceInfo(ctx sdk.Context, certificationType types.CertificationType) int {
	return k.countTotal(ctx, types.GetCertificationPrefix(certificationType))
}

// Check if the ComplianceInfo is present in the store or not.
func (k Keeper) IsComplianceInfoPresent(ctx sdk.Context,
	certificationType types.CertificationType, vid uint16, pid uint16, softwareVersion uint32) bool {
	return k.isRecordPresent(ctx, types.GetComplianceInfoKey(certificationType, vid, pid, softwareVersion))
}

// Check if the record is present in the store or not.
func (k Keeper) isRecordPresent(ctx sdk.Context, id []byte) bool {
	store := ctx.KVStore(k.storeKey)

	return store.Has(id)
}

//nolint:godox
//  TODO: Is iteration the only way to calculate the total number of elements?
//  It looks like that in a non-pagination case we iterate twice:
// to get the total number of elements and to get the real content.
func (k Keeper) countTotal(ctx sdk.Context, prefix []byte) int {
	store := ctx.KVStore(k.storeKey)
	res := 0

	iter := sdk.KVStorePrefixIterator(store, prefix)
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		res++
	}

	return res
}
