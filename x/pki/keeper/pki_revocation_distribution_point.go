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
	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

// SetPkiRevocationDistributionPoint set a specific pKIRevocationDistributionPoint in the store from its index.
func (k Keeper) SetPkiRevocationDistributionPoint(ctx sdk.Context, pKIRevocationDistributionPoint types.PkiRevocationDistributionPoint) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.PkiRevocationDistributionPointKeyPrefix))
	b := k.cdc.MustMarshal(&pKIRevocationDistributionPoint)
	store.Set(types.PkiRevocationDistributionPointKey(
		pKIRevocationDistributionPoint.Vid,
		pKIRevocationDistributionPoint.Label,
		pKIRevocationDistributionPoint.IssuerSubjectKeyID,
	), b)
}

// GetPkiRevocationDistributionPoint returns a pKIRevocationDistributionPoint from its index.
func (k Keeper) GetPkiRevocationDistributionPoint(
	ctx sdk.Context,
	vid int32,
	label string,
	issuerSubjectKeyID string,

) (val types.PkiRevocationDistributionPoint, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.PkiRevocationDistributionPointKeyPrefix))

	b := store.Get(types.PkiRevocationDistributionPointKey(
		vid,
		label,
		issuerSubjectKeyID,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)

	return val, true
}

// RemovePkiRevocationDistributionPoint removes a pKIRevocationDistributionPoint from the store.
func (k Keeper) RemovePkiRevocationDistributionPoint(
	ctx sdk.Context,
	vid int32,
	label string,
	issuerSubjectKeyID string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.PkiRevocationDistributionPointKeyPrefix))
	store.Delete(types.PkiRevocationDistributionPointKey(
		vid,
		label,
		issuerSubjectKeyID,
	))
}

// GetAllPkiRevocationDistributionPoint returns all pKIRevocationDistributionPoint.
func (k Keeper) GetAllPkiRevocationDistributionPoint(ctx sdk.Context) (list []types.PkiRevocationDistributionPoint) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.PkiRevocationDistributionPointKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer func() { _ = iterator.Close() }()

	for ; iterator.Valid(); iterator.Next() {
		var val types.PkiRevocationDistributionPoint
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
