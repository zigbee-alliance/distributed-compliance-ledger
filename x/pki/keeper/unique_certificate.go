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
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

// SetUniqueCertificate set a specific uniqueCertificate in the store from its index.
func (k Keeper) SetUniqueCertificate(ctx sdk.Context, uniqueCertificate types.UniqueCertificate) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UniqueCertificateKeyPrefix))
	b := k.cdc.MustMarshal(&uniqueCertificate)
	store.Set(types.UniqueCertificateKey(
		uniqueCertificate.Issuer,
		uniqueCertificate.SerialNumber,
	), b)
}

// GetUniqueCertificate returns a uniqueCertificate from its index.
func (k Keeper) GetUniqueCertificate(
	ctx sdk.Context,
	issuer string,
	serialNumber string,

) (val types.UniqueCertificate, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UniqueCertificateKeyPrefix))

	b := store.Get(types.UniqueCertificateKey(
		issuer,
		serialNumber,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveUniqueCertificate removes a uniqueCertificate from the store.
func (k Keeper) RemoveUniqueCertificate(
	ctx sdk.Context,
	issuer string,
	serialNumber string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UniqueCertificateKeyPrefix))
	store.Delete(types.UniqueCertificateKey(
		issuer,
		serialNumber,
	))
}

// GetAllUniqueCertificate returns all uniqueCertificate.
func (k Keeper) GetAllUniqueCertificate(ctx sdk.Context) (list []types.UniqueCertificate) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UniqueCertificateKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.UniqueCertificate
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// Check if the unique certificate key (Issuer/SerialNumber combination) is busy.
func (k Keeper) IsUniqueCertificatePresent(
	ctx sdk.Context,
	issuer string,
	serialNumber string,

) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UniqueCertificateKeyPrefix))
	return store.Has(types.UniqueCertificateKey(
		issuer,
		serialNumber,
	))
}
