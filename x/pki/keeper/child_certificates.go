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

// SetChildCertificates set a specific childCertificates in the store from its index.
func (k Keeper) SetChildCertificates(ctx sdk.Context, childCertificates types.ChildCertificates) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChildCertificatesKeyPrefix))
	b := k.cdc.MustMarshal(&childCertificates)
	store.Set(types.ChildCertificatesKey(
		childCertificates.Issuer,
		childCertificates.AuthorityKeyId,
	), b)
}

// GetChildCertificates returns a childCertificates from its index.
func (k Keeper) GetChildCertificates(
	ctx sdk.Context,
	issuer string,
	authorityKeyId string,

) (val types.ChildCertificates, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChildCertificatesKeyPrefix))

	b := store.Get(types.ChildCertificatesKey(
		issuer,
		authorityKeyId,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveChildCertificates removes a childCertificates from the store.
func (k Keeper) RemoveChildCertificates(
	ctx sdk.Context,
	issuer string,
	authorityKeyId string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChildCertificatesKeyPrefix))
	store.Delete(types.ChildCertificatesKey(
		issuer,
		authorityKeyId,
	))
}

// GetAllChildCertificates returns all childCertificates.
func (k Keeper) GetAllChildCertificates(ctx sdk.Context) (list []types.ChildCertificates) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChildCertificatesKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ChildCertificates
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// Add a child certificate to the list of child certificate IDs for the issuer/authorityKeyId map.
func (k Keeper) AddChildCertificate(ctx sdk.Context, issuer string, authorityKeyId string, certId types.CertificateIdentifier) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChildCertificatesKeyPrefix))

	childCertificatesBytes := store.Get(types.ChildCertificatesKey(
		issuer,
		authorityKeyId,
	))

	var childCertificates types.ChildCertificates
	if childCertificatesBytes == nil {
		childCertificates = types.ChildCertificates{
			Issuer:         issuer,
			AuthorityKeyId: authorityKeyId,
			CertIds:        []*types.CertificateIdentifier{},
		}
	} else {
		k.cdc.MustUnmarshal(childCertificatesBytes, &childCertificates)
	}

	for _, existingCertId := range childCertificates.CertIds {
		if *existingCertId == certId {
			return
		}
	}

	childCertificates.CertIds = append(childCertificates.CertIds, &certId)

	b := k.cdc.MustMarshal(&childCertificates)
	store.Set(types.ChildCertificatesKey(
		issuer,
		authorityKeyId,
	), b)
}

func (k msgServer) RevokeChildCertificates(ctx sdk.Context, issuer string, authorityKeyId string) {
	// Get issuer's ChildCertificates record
	childCertificates, _ := k.GetChildCertificates(ctx, issuer, authorityKeyId)

	// For each child certificate subject/subjectKeyID combination
	for _, certIdentifier := range childCertificates.CertIds {
		// Revoke certificates with this subject/subjectKeyID combination
		certificates, _ := k.GetApprovedCertificates(ctx, certIdentifier.Subject, certIdentifier.SubjectKeyId)
		k.AddRevokedCertificates(ctx, certificates)
		k.RemoveApprovedCertificates(ctx, certIdentifier.Subject, certIdentifier.SubjectKeyId)

		// remove from subject -> subject key ID map
		k.RemoveApprovedCertificateBySubject(ctx, certIdentifier.Subject, certIdentifier.SubjectKeyId)

		// Process child certificates recursively
		k.RevokeChildCertificates(ctx, certIdentifier.Subject, certIdentifier.SubjectKeyId)
	}

	// Delete entire ChildCertificates record of issuer
	k.RemoveChildCertificates(ctx, issuer, authorityKeyId)
}

func (k msgServer) RemoveChildCertificate(ctx sdk.Context, issuer string, authorityKeyId string,
	certIdentifier types.CertificateIdentifier) {
	childCertificates, _ := k.GetChildCertificates(ctx, issuer, authorityKeyId)

	certIDIndex := -1
	for i, existingIdentifier := range childCertificates.CertIds {
		if *existingIdentifier == certIdentifier {
			certIDIndex = i
			break
		}
	}

	if certIDIndex == -1 {
		return
	}

	childCertificates.CertIds =
		append(childCertificates.CertIds[:certIDIndex], childCertificates.CertIds[certIDIndex+1:]...)

	if len(childCertificates.CertIds) > 0 {
		k.SetChildCertificates(ctx, childCertificates)
	} else {
		k.RemoveChildCertificates(ctx, issuer, authorityKeyId)
	}
}
