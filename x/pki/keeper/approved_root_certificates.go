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

// SetApprovedRootCertificates set approvedRootCertificates in the store.
func (k Keeper) SetApprovedRootCertificates(ctx sdk.Context, approvedRootCertificates types.ApprovedRootCertificates) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ApprovedRootCertificatesKey))
	b := k.cdc.MustMarshal(&approvedRootCertificates)
	store.Set([]byte{0}, b)
}

// GetApprovedRootCertificates returns approvedRootCertificates.
func (k Keeper) GetApprovedRootCertificates(ctx sdk.Context) (val types.ApprovedRootCertificates, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ApprovedRootCertificatesKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveApprovedRootCertificates removes approvedRootCertificates from the store.
func (k Keeper) RemoveApprovedRootCertificates(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ApprovedRootCertificatesKey))
	store.Delete([]byte{0})
}

// Add root certificate to the list.
func (k Keeper) AddApprovedRootCertificate(ctx sdk.Context, certId types.CertificateIdentifier) {
	rootCertificates, _ := k.GetApprovedRootCertificates(ctx)

	// Check if the root cert is already there
	for _, existingСertId := range rootCertificates.Certs {
		if *existingСertId == certId {
			return
		}
	}

	rootCertificates.Certs = append(rootCertificates.Certs, &certId)

	k.SetApprovedRootCertificates(ctx, rootCertificates)
}

// Remove root certificate from the list.
func (k Keeper) RemoveApprovedRootCertificate(ctx sdk.Context, certId types.CertificateIdentifier) {
	rootCertificates, _ := k.GetApprovedRootCertificates(ctx)

	certIDIndex := -1
	for i, existingIdentifier := range rootCertificates.Certs {
		if *existingIdentifier == certId {
			certIDIndex = i
			break
		}
	}
	if certIDIndex == -1 {
		return
	}

	rootCertificates.Certs =
		append(rootCertificates.Certs[:certIDIndex], rootCertificates.Certs[certIDIndex+1:]...)

	k.SetApprovedRootCertificates(ctx, rootCertificates)
}
