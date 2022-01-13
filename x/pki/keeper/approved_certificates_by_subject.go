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

// SetApprovedCertificatesBySubject set a specific approvedCertificatesBySubject in the store from its index.
func (k Keeper) SetApprovedCertificatesBySubject(ctx sdk.Context, approvedCertificatesBySubject types.ApprovedCertificatesBySubject) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ApprovedCertificatesBySubjectKeyPrefix))
	b := k.cdc.MustMarshal(&approvedCertificatesBySubject)
	store.Set(types.ApprovedCertificatesBySubjectKey(
		approvedCertificatesBySubject.Subject,
	), b)
}

// GetApprovedCertificatesBySubject returns a approvedCertificatesBySubject from its index.
func (k Keeper) GetApprovedCertificatesBySubject(
	ctx sdk.Context,
	subject string,

) (val types.ApprovedCertificatesBySubject, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ApprovedCertificatesBySubjectKeyPrefix))

	b := store.Get(types.ApprovedCertificatesBySubjectKey(
		subject,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveApprovedCertificatesBySubject removes a approvedCertificatesBySubject from the store.
func (k Keeper) RemoveApprovedCertificatesBySubject(
	ctx sdk.Context,
	subject string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ApprovedCertificatesBySubjectKeyPrefix))
	store.Delete(types.ApprovedCertificatesBySubjectKey(
		subject,
	))
}

// GetAllApprovedCertificatesBySubject returns all approvedCertificatesBySubject.
func (k Keeper) GetAllApprovedCertificatesBySubject(ctx sdk.Context) (list []types.ApprovedCertificatesBySubject) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ApprovedCertificatesBySubjectKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ApprovedCertificatesBySubject
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// Add ApprovedCertificates to a subject->subjectKeyId index.
func (k Keeper) AddApprovedCertificateBySubject(ctx sdk.Context, subject string, subjectKeyId string) {
	approvedCertificatesBySubject, _ := k.GetApprovedCertificatesBySubject(ctx, subject)

	// Check if cert is already there
	for _, existingId := range approvedCertificatesBySubject.SubjectKeyIds {
		if existingId == subjectKeyId {
			return
		}
	}

	approvedCertificatesBySubject.Subject = subject
	approvedCertificatesBySubject.SubjectKeyIds = append(approvedCertificatesBySubject.SubjectKeyIds, subjectKeyId)

	k.SetApprovedCertificatesBySubject(ctx, approvedCertificatesBySubject)
}

// Remove revoked root certificate from the list.
func (k Keeper) RemoveApprovedCertificateBySubject(ctx sdk.Context, subject string, subjectKeyId string) {
	approvedCertificatesBySubject, _ := k.GetApprovedCertificatesBySubject(ctx, subject)

	certIDIndex := -1
	for i, existingIdentifier := range approvedCertificatesBySubject.SubjectKeyIds {
		if existingIdentifier == subjectKeyId {
			certIDIndex = i
			break
		}
	}
	if certIDIndex == -1 {
		return
	}

	approvedCertificatesBySubject.SubjectKeyIds =
		append(approvedCertificatesBySubject.SubjectKeyIds[:certIDIndex], approvedCertificatesBySubject.SubjectKeyIds[certIDIndex+1:]...)

	if len(approvedCertificatesBySubject.SubjectKeyIds) > 0 {
		k.SetApprovedCertificatesBySubject(ctx, approvedCertificatesBySubject)
	} else {
		k.RemoveApprovedCertificatesBySubject(ctx, subject)
	}
}
