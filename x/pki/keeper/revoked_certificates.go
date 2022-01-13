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

// SetRevokedCertificates set a specific revokedCertificates in the store from its index.
func (k Keeper) SetRevokedCertificates(ctx sdk.Context, revokedCertificates types.RevokedCertificates) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RevokedCertificatesKeyPrefix))
	b := k.cdc.MustMarshal(&revokedCertificates)
	store.Set(types.RevokedCertificatesKey(
		revokedCertificates.Subject,
		revokedCertificates.SubjectKeyId,
	), b)
}

// GetRevokedCertificates returns a revokedCertificates from its index.
func (k Keeper) GetRevokedCertificates(
	ctx sdk.Context,
	subject string,
	subjectKeyId string,

) (val types.RevokedCertificates, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RevokedCertificatesKeyPrefix))

	b := store.Get(types.RevokedCertificatesKey(
		subject,
		subjectKeyId,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveRevokedCertificates removes a revokedCertificates from the store.
func (k Keeper) RemoveRevokedCertificates(
	ctx sdk.Context,
	subject string,
	subjectKeyId string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RevokedCertificatesKeyPrefix))
	store.Delete(types.RevokedCertificatesKey(
		subject,
		subjectKeyId,
	))
}

// GetAllRevokedCertificates returns all revokedCertificates.
func (k Keeper) GetAllRevokedCertificates(ctx sdk.Context) (list []types.RevokedCertificates) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RevokedCertificatesKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.RevokedCertificates
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// Add revoked certificates to the list of revoked certificates for the subject/subjectKeyId map.
func (k Keeper) AddRevokedCertificates(ctx sdk.Context, approvedCertificates types.ApprovedCertificates) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RevokedCertificatesKeyPrefix))

	revokedCertificatesBytes := store.Get(types.RevokedCertificatesKey(
		approvedCertificates.Subject,
		approvedCertificates.SubjectKeyId,
	))
	var revokedCertificates types.RevokedCertificates

	if revokedCertificatesBytes == nil {
		revokedCertificates = types.RevokedCertificates{
			Subject:      approvedCertificates.Subject,
			SubjectKeyId: approvedCertificates.SubjectKeyId,
			Certs:        []*types.Certificate{},
		}
	} else {
		k.cdc.MustUnmarshal(revokedCertificatesBytes, &revokedCertificates)
	}

	revokedCertificates.Certs = append(revokedCertificates.Certs, approvedCertificates.Certs...)

	b := k.cdc.MustMarshal(&revokedCertificates)
	store.Set(types.RevokedCertificatesKey(
		revokedCertificates.Subject,
		revokedCertificates.SubjectKeyId,
	), b)
}
