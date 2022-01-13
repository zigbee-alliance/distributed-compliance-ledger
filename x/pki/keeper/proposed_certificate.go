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

// SetProposedCertificate set a specific proposedCertificate in the store from its index.
func (k Keeper) SetProposedCertificate(ctx sdk.Context, proposedCertificate types.ProposedCertificate) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProposedCertificateKeyPrefix))
	b := k.cdc.MustMarshal(&proposedCertificate)
	store.Set(types.ProposedCertificateKey(
		proposedCertificate.Subject,
		proposedCertificate.SubjectKeyId,
	), b)
}

// GetProposedCertificate returns a proposedCertificate from its index.
func (k Keeper) GetProposedCertificate(
	ctx sdk.Context,
	subject string,
	subjectKeyId string,

) (val types.ProposedCertificate, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProposedCertificateKeyPrefix))

	b := store.Get(types.ProposedCertificateKey(
		subject,
		subjectKeyId,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveProposedCertificate removes a proposedCertificate from the store.
func (k Keeper) RemoveProposedCertificate(
	ctx sdk.Context,
	subject string,
	subjectKeyId string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProposedCertificateKeyPrefix))
	store.Delete(types.ProposedCertificateKey(
		subject,
		subjectKeyId,
	))
}

// GetAllProposedCertificate returns all proposedCertificate.
func (k Keeper) GetAllProposedCertificate(ctx sdk.Context) (list []types.ProposedCertificate) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProposedCertificateKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ProposedCertificate
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// Check if the Proposed Certificate record associated with a
// Subject/SubjectKeyID combination is present in the store.
func (k Keeper) IsProposedCertificatePresent(
	ctx sdk.Context,
	subject string,
	subjectKeyId string,

) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProposedCertificateKeyPrefix))
	return store.Has(types.ProposedCertificateKey(
		subject,
		subjectKeyId,
	))
}
