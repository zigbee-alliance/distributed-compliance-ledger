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
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/internal/types"
)

type Keeper struct {
	// Unexposed key to access store from sdk.Context
	storeKey sdk.StoreKey

	// The wire codec for binary encoding/decoding
	cdc *codec.Codec
}

func NewKeeper(storeKey sdk.StoreKey, cdc *codec.Codec) Keeper {
	return Keeper{storeKey: storeKey, cdc: cdc}
}

/*
	Approved Certificate (root or non-root)
*/

// Gets the entire Approved Certificates record associated with a Subject/SubjectKeyID combination.
func (k Keeper) GetApprovedCertificates(ctx sdk.Context, subject string, subjectKeyID string) types.Certificates {
	if !k.IsApprovedCertificatesPresent(ctx, subject, subjectKeyID) {
		return types.NewCertificates([]types.Certificate{})
	}

	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetApprovedCertificateKey(subject, subjectKeyID))

	var cert types.Certificates

	k.cdc.MustUnmarshalBinaryBare(bz, &cert)

	return cert
}

// Sets the entire Approved Certificates record for a Subject/SubjectKeyID combination.
func (k Keeper) SetApprovedCertificates(ctx sdk.Context, subject string, subjectKeyID string,
	certificates types.Certificates) {
	if len(certificates.Items) == 0 {
		panic("Cannot set approved Certificates record with no items")
	}

	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetApprovedCertificateKey(subject, subjectKeyID), k.cdc.MustMarshalBinaryBare(certificates))
}

// Add the Certificate to the Approved Certificates record with the corresponding Subject/SubjectKeyID combination.
func (k Keeper) AddApprovedCertificate(ctx sdk.Context, certificate types.Certificate) {
	certificates := k.GetApprovedCertificates(ctx, certificate.Subject, certificate.SubjectKeyID)
	certificates.Items = append(certificates.Items, certificate)
	k.SetApprovedCertificates(ctx, certificate.Subject, certificate.SubjectKeyID, certificates)
}

// Check if the Approved Certificates record associated with a Subject/SubjectKeyID combination
// is present in the store or not.
func (k Keeper) IsApprovedCertificatesPresent(ctx sdk.Context, subject string, subjectKeyID string) bool {
	store := ctx.KVStore(k.storeKey)

	return store.Has(types.GetApprovedCertificateKey(subject, subjectKeyID))
}

// Iterate over all Approved Certificates.
func (k Keeper) IterateApprovedCertificatesRecords(ctx sdk.Context, prefix string,
	process func(info types.Certificates) (stop bool)) {
	store := ctx.KVStore(k.storeKey)

	iter := sdk.KVStorePrefixIterator(store, append(types.ApprovedCertificatePrefix, []byte(prefix)...))
	defer iter.Close()

	for {
		if !iter.Valid() {
			return
		}

		val := iter.Value()

		var certificates types.Certificates

		k.cdc.MustUnmarshalBinaryBare(val, &certificates)

		if process(certificates) {
			return
		}

		iter.Next()
	}
}

// Deletes the entire Approved Certificates record associated with a Subject/SubjectKeyID combination.
func (k Keeper) DeleteApprovedCertificates(ctx sdk.Context, subject string, subjectKeyID string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetApprovedCertificateKey(subject, subjectKeyID))
}

/*
	Proposed Root Certificate
*/

// Gets the entire Proposed Certificate record associated with a Subject/SubjectKeyID combination.
func (k Keeper) GetProposedCertificate(ctx sdk.Context,
	subject string, subjectKeyID string) types.ProposedCertificate {
	if !k.IsProposedCertificatePresent(ctx, subject, subjectKeyID) {
		panic("Proposed Certificate does not exist")
	}

	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetProposedCertificateKey(subject, subjectKeyID))

	var cert types.ProposedCertificate

	k.cdc.MustUnmarshalBinaryBare(bz, &cert)

	return cert
}

// Sets the entire Proposed Certificate record for a Subject/SubjectKeyID combination.
func (k Keeper) SetProposedCertificate(ctx sdk.Context, certificate types.ProposedCertificate) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetProposedCertificateKey(
		certificate.Subject, certificate.SubjectKeyID), k.cdc.MustMarshalBinaryBare(certificate))
}

// Check if the Proposed Certificate record associated with a
// Subject/SubjectKeyID combination is present in the store or not.
func (k Keeper) IsProposedCertificatePresent(ctx sdk.Context, subject string, subjectKeyID string) bool {
	store := ctx.KVStore(k.storeKey)

	return store.Has(types.GetProposedCertificateKey(subject, subjectKeyID))
}

// Iterate over all Proposed Certificates.
func (k Keeper) IterateProposedCertificates(ctx sdk.Context,
	process func(info types.ProposedCertificate) (stop bool)) {
	store := ctx.KVStore(k.storeKey)

	iter := sdk.KVStorePrefixIterator(store, types.ProposedCertificatePrefix)
	defer iter.Close()

	for {
		if !iter.Valid() {
			return
		}

		val := iter.Value()

		var proposedCertificate types.ProposedCertificate

		k.cdc.MustUnmarshalBinaryBare(val, &proposedCertificate)

		if process(proposedCertificate) {
			return
		}

		iter.Next()
	}
}

// Deletes the Proposed Certificate from the store.
func (k Keeper) DeleteProposedCertificate(ctx sdk.Context, subject string, subjectKeyID string) {
	if !k.IsProposedCertificatePresent(ctx, subject, subjectKeyID) {
		panic("Proposed Certificate does not exist")
	}

	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetProposedCertificateKey(subject, subjectKeyID))
}

/*
	Record containing the list of direct child certificates. Referenced by an Issuer/AuthorityKeyID combination.
*/

// Gets the Child Certificates for a record associated with a combination Issuer/AuthorityKeyID.
func (k Keeper) GetChildCertificates(ctx sdk.Context, issuer string, authorityKeyID string) types.ChildCertificates {
	if !k.IsChildCertificatesPresent(ctx, issuer, authorityKeyID) {
		return types.NewChildCertificates(issuer, authorityKeyID)
	}

	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetChildCertificatesKey(issuer, authorityKeyID))

	var childCertificates types.ChildCertificates

	k.cdc.MustUnmarshalBinaryBare(bz, &childCertificates)

	return childCertificates
}

// Sets the Child Certificates record for the combination Issuer/AuthorityKeyID.
func (k Keeper) SetChildCertificates(ctx sdk.Context, childCertificates types.ChildCertificates) {
	if len(childCertificates.CertIdentifiers) == 0 {
		panic("Cannot set ChildCertificates record with no CertIdentifiers")
	}

	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetChildCertificatesKey(childCertificates.Issuer,
		childCertificates.AuthorityKeyID), k.cdc.MustMarshalBinaryBare(childCertificates))
}

// Checks if the the list of Child Certificates for a combination Issuer/AuthorityKeyID is present in the store or not.
func (k Keeper) IsChildCertificatesPresent(ctx sdk.Context, issuer string, authorityKeyID string) bool {
	store := ctx.KVStore(k.storeKey)

	return store.Has(types.GetChildCertificatesKey(issuer, authorityKeyID))
}

// Iterate over all Child Certificates records.
func (k Keeper) IterateChildCertificatesRecords(ctx sdk.Context,
	process func(info types.ChildCertificates) (stop bool)) {
	store := ctx.KVStore(k.storeKey)

	iter := sdk.KVStorePrefixIterator(store, types.ChildCertificatesPrefix)
	defer iter.Close()

	for {
		if !iter.Valid() {
			return
		}

		val := iter.Value()

		var childCertificatesRecord types.ChildCertificates

		k.cdc.MustUnmarshalBinaryBare(val, &childCertificatesRecord)

		if process(childCertificatesRecord) {
			return
		}

		iter.Next()
	}
}

// Deletes the Child Certificates record.
func (k Keeper) DeleteChildCertificates(ctx sdk.Context, issuer string, authorityKeyID string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetChildCertificatesKey(issuer, authorityKeyID))
}

/*
	Unique certificate key (Issuer/SerialNumber combination).
	Keeper allows to register these keys to track their uniqueness.
*/

// Registers the unique certificate key (Issuer/SerialNumber combination) to indicate that it is busy.
func (k Keeper) SetUniqueCertificateKey(ctx sdk.Context, issuer string, serialNumber string) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetUniqueCertificateKey(issuer, serialNumber), k.cdc.MustMarshalBinaryBare(true))
}

// Check if the unique certificate key (Issuer/SerialNumber combination) is busy.
func (k Keeper) IsUniqueCertificateKeyPresent(ctx sdk.Context, issuer string, serialNumber string) bool {
	store := ctx.KVStore(k.storeKey)

	return store.Has(types.GetUniqueCertificateKey(issuer, serialNumber))
}

// Deletes the unique certificate key (Issuer/SerialNumber combination).
func (k Keeper) DeleteUniqueCertificateKey(ctx sdk.Context, issuer string, serialNumber string) {
	if !k.IsUniqueCertificateKeyPresent(ctx, issuer, serialNumber) {
		panic("Unique Certificate Key does not exist")
	}

	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetUniqueCertificateKey(issuer, serialNumber))
}

/*
	Proposed Root Certificate Revocation
*/

// Gets the Proposed Certificate Revocation record associated with a Subject/SubjectKeyID combination.
func (k Keeper) GetProposedCertificateRevocation(ctx sdk.Context,
	subject string, subjectKeyID string) types.ProposedCertificateRevocation {
	if !k.IsProposedCertificateRevocationPresent(ctx, subject, subjectKeyID) {
		panic("Proposed Certificate Revocation does not exist")
	}

	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetProposedCertificateRevocationKey(subject, subjectKeyID))

	var revocation types.ProposedCertificateRevocation

	k.cdc.MustUnmarshalBinaryBare(bz, &revocation)

	return revocation
}

// Sets the Proposed Certificate Revocation record for a Subject/SubjectKeyID combination.
func (k Keeper) SetProposedCertificateRevocation(ctx sdk.Context, revocation types.ProposedCertificateRevocation) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetProposedCertificateRevocationKey(
		revocation.Subject, revocation.SubjectKeyID), k.cdc.MustMarshalBinaryBare(revocation))
}

// Check if the Proposed Certificate Revocation record associated with a
// Subject/SubjectKeyID combination is present in the store or not.
func (k Keeper) IsProposedCertificateRevocationPresent(ctx sdk.Context, subject string, subjectKeyID string) bool {
	store := ctx.KVStore(k.storeKey)

	return store.Has(types.GetProposedCertificateRevocationKey(subject, subjectKeyID))
}

// Iterate over all Proposed Certificate Revocations.
func (k Keeper) IterateProposedCertificateRevocations(ctx sdk.Context,
	process func(info types.ProposedCertificateRevocation) (stop bool)) {
	store := ctx.KVStore(k.storeKey)

	iter := sdk.KVStorePrefixIterator(store, types.ProposedCertificateRevocationPrefix)
	defer iter.Close()

	for {
		if !iter.Valid() {
			return
		}

		val := iter.Value()

		var revocation types.ProposedCertificateRevocation

		k.cdc.MustUnmarshalBinaryBare(val, &revocation)

		if process(revocation) {
			return
		}

		iter.Next()
	}
}

/*
	Revoked Certificate (root or non-root)
*/

// Gets the entire Revoked Certificates record associated with a Subject/SubjectKeyID combination.
func (k Keeper) GetRevokedCertificates(ctx sdk.Context, subject string, subjectKeyID string) types.Certificates {
	if !k.IsRevokedCertificatesPresent(ctx, subject, subjectKeyID) {
		return types.NewCertificates([]types.Certificate{})
	}

	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetRevokedCertificateKey(subject, subjectKeyID))

	var cert types.Certificates

	k.cdc.MustUnmarshalBinaryBare(bz, &cert)

	return cert
}

// Sets the entire Revoked Certificates record for a Subject/SubjectKeyID combination.
func (k Keeper) SetRevokedCertificates(ctx sdk.Context, subject string, subjectKeyID string,
	certificates types.Certificates) {
	if len(certificates.Items) == 0 {
		panic("Cannot set revoked Certificates record with no items")
	}

	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetRevokedCertificateKey(subject, subjectKeyID), k.cdc.MustMarshalBinaryBare(certificates))
}

// Add certificates to the Revoked Certificates record for a Subject/SubjectKeyID combination.
func (k Keeper) AddRevokedCertificates(ctx sdk.Context, subject string, subjectKeyID string,
	certificates types.Certificates) {
	revokedCertificates := k.GetRevokedCertificates(ctx, subject, subjectKeyID)
	revokedCertificates.Items = append(revokedCertificates.Items, certificates.Items...)
	k.SetRevokedCertificates(ctx, subject, subjectKeyID, revokedCertificates)
}

// Check if the Revoked Certificates record associated with a Subject/SubjectKeyID combination
// is present in the store or not.
func (k Keeper) IsRevokedCertificatesPresent(ctx sdk.Context, subject string, subjectKeyID string) bool {
	store := ctx.KVStore(k.storeKey)

	return store.Has(types.GetRevokedCertificateKey(subject, subjectKeyID))
}

func (k Keeper) IterateRevokedCertificatesRecords(ctx sdk.Context, prefix string,
	process func(info types.Certificates) (stop bool)) {
	store := ctx.KVStore(k.storeKey)

	iter := sdk.KVStorePrefixIterator(store, append(types.RevokedCertificatePrefix, []byte(prefix)...))
	defer iter.Close()

	for {
		if !iter.Valid() {
			return
		}

		val := iter.Value()

		var certificates types.Certificates

		k.cdc.MustUnmarshalBinaryBare(val, &certificates)

		if process(certificates) {
			return
		}

		iter.Next()
	}
}

// Deletes the entire Revoked Certificates record associated with a Subject/SubjectKeyID combination.
func (k Keeper) DeleteRevokedCertificates(ctx sdk.Context, subject string, subjectKeyID string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetRevokedCertificateKey(subject, subjectKeyID))
}
