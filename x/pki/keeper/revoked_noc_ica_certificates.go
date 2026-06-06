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

// SetRevokedNocIcaCertificates set a specific revokedNocIcaCertificates in the store from its index
func (k Keeper) SetRevokedNocIcaCertificates(ctx sdk.Context, revokedNocIcaCertificates types.RevokedNocIcaCertificates) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.RevokedNocIcaCertificatesKeyPrefix))
	b := k.cdc.MustMarshal(&revokedNocIcaCertificates)
	store.Set(types.RevokedNocIcaCertificatesKey(
		revokedNocIcaCertificates.Subject,
		revokedNocIcaCertificates.SubjectKeyId,
	), b)
}

// AddRevokedNocIcaCertificates adds revoked NOC certificates to the list of revoked NOC certificates for the subject/subjectKeyId map.
func (k Keeper) AddRevokedNocIcaCertificates(ctx sdk.Context, revokedNocIcaCertificates types.RevokedNocIcaCertificates) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.RevokedNocIcaCertificatesKeyPrefix))

	revokedCertsBytes := store.Get(types.RevokedNocIcaCertificatesKey(
		revokedNocIcaCertificates.Subject,
		revokedNocIcaCertificates.SubjectKeyId,
	))
	var revokedCerts types.RevokedNocIcaCertificates

	if revokedCertsBytes == nil {
		revokedCerts = types.RevokedNocIcaCertificates{
			Subject:      revokedNocIcaCertificates.Subject,
			SubjectKeyId: revokedNocIcaCertificates.SubjectKeyId,
			Certs:        []*types.Certificate{},
		}
	} else {
		k.cdc.MustUnmarshal(revokedCertsBytes, &revokedCerts)
	}

	revokedCerts.Certs = append(revokedCerts.Certs, revokedNocIcaCertificates.Certs...)

	b := k.cdc.MustMarshal(&revokedCerts)
	store.Set(types.RevokedNocIcaCertificatesKey(
		revokedCerts.Subject,
		revokedCerts.SubjectKeyId,
	), b)
}

// GetRevokedNocIcaCertificates returns a revokedNocIcaCertificates from its index
func (k Keeper) GetRevokedNocIcaCertificates(
	ctx sdk.Context,
	subject string,
	subjectKeyID string,

) (val types.RevokedNocIcaCertificates, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.RevokedNocIcaCertificatesKeyPrefix))

	b := store.Get(types.RevokedNocIcaCertificatesKey(
		subject,
		subjectKeyID,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)

	return val, true
}

// RemoveRevokedNocIcaCertificates removes a revokedNocIcaCertificates from the store
func (k Keeper) RemoveRevokedNocIcaCertificates(
	ctx sdk.Context,
	subject string,
	subjectKeyID string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.RevokedNocIcaCertificatesKeyPrefix))
	store.Delete(types.RevokedNocIcaCertificatesKey(
		subject,
		subjectKeyID,
	))
}

// GetAllRevokedNocIcaCertificates returns all revokedNocIcaCertificates
func (k Keeper) GetAllRevokedNocIcaCertificates(ctx sdk.Context) (list []types.RevokedNocIcaCertificates) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.RevokedNocIcaCertificatesKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer func() { _ = iterator.Close() }()

	for ; iterator.Valid(); iterator.Next() {
		var val types.RevokedNocIcaCertificates
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// IsRevokedNocIcaCertificatePresent Check if the Revoked Noc ICA Certificate record associated with a Subject/SubjectKeyID combination is present in the store.
func (k Keeper) IsRevokedNocIcaCertificatePresent(
	ctx sdk.Context,
	subject string,
	subjectKeyID string,
) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.RevokedNocIcaCertificatesKeyPrefix))

	return store.Has(types.RevokedNocIcaCertificatesKey(
		subject,
		subjectKeyID,
	))
}
