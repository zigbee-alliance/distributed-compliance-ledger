package keeper

import (
	"fmt"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/pki/internal/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Keeper struct {
	// Unexposed key to access store from sdk.Context
	storeKey sdk.StoreKey

	// The wire codec for binary encoding/decoding
	cdc *codec.Codec
}

const (
	approvedCertificatePrefix = "ac"
	proposedCertificatePrefix = "pc"
	childCertificatesPrefix   = "cc"
)

func NewKeeper(storeKey sdk.StoreKey, cdc *codec.Codec) Keeper {
	return Keeper{storeKey: storeKey, cdc: cdc}
}

/*
	Approved Root or Intermediate certificate
*/

// Gets the entire Certificate record associated with a Subject/SubjectKeyId combination
func (k Keeper) GetCertificate(ctx sdk.Context, subject string, subjectKeyId string) types.Certificate {
	if !k.IsCertificatePresent(ctx, subject, subjectKeyId) {
		panic("Certificate does not exist")
	}

	store := ctx.KVStore(k.storeKey)
	bz := store.Get([]byte(CertificateId(subject, subjectKeyId)))

	var cert types.Certificate
	k.cdc.MustUnmarshalBinaryBare(bz, &cert)
	return cert
}

// Sets the entire Certificate record for a Subject/SubjectKeyId combination
func (k Keeper) SetCertificate(ctx sdk.Context, certificate types.Certificate) {
	store := ctx.KVStore(k.storeKey)
	store.Set([]byte(CertificateId(certificate.Subject, certificate.SubjectKeyId)), k.cdc.MustMarshalBinaryBare(certificate))
}

// Deletes the entire Certificate record associated with a Subject/SubjectKeyId combination
func (k Keeper) DeleteCertificate(ctx sdk.Context, certificate types.Certificate) {
	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(CertificateId(certificate.Subject, certificate.SubjectKeyId)))
}

// Check if the Certificate record associated with a Subject/SubjectKeyId combination is present in the store or not
func (k Keeper) IsCertificatePresent(ctx sdk.Context, subject string, subjectKeyId string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(CertificateId(subject, subjectKeyId)))
}

// Count total Certificates
func (k Keeper) CountTotalCertificates(ctx sdk.Context) int {
	return k.countTotal(ctx, approvedCertificatePrefix)
}

func (k Keeper) IterateCertificates(ctx sdk.Context, prefix string, process func(info types.Certificate) (stop bool)) {
	store := ctx.KVStore(k.storeKey)

	iter := sdk.KVStorePrefixIterator(store, []byte(fmt.Sprintf("%v:%v", approvedCertificatePrefix, prefix)))
	defer iter.Close()

	for {
		if !iter.Valid() {
			return
		}

		val := iter.Value()

		var certificate types.Certificate

		k.cdc.MustUnmarshalBinaryBare(val, &certificate)

		if process(certificate) {
			return
		}

		iter.Next()
	}
}

// Id builder for Certificate record
func CertificateId(subject string, subjectKeyId string) string {
	return fmt.Sprintf("%s:%v:%v", approvedCertificatePrefix, subject, subjectKeyId)
}

/*
	Proposed Root certificate
*/

// Gets the entire Proposed Certificate record associated with a Subject/SubjectKeyId combination
func (k Keeper) GetProposedCertificate(ctx sdk.Context, subject string, subjectKeyId string) types.ProposedCertificate {
	if !k.IsProposedCertificatePresent(ctx, subject, subjectKeyId) {
		panic("Proposed Certificate does not exist")
	}

	store := ctx.KVStore(k.storeKey)
	bz := store.Get([]byte(ProposedCertificateId(subject, subjectKeyId)))

	var cert types.ProposedCertificate
	k.cdc.MustUnmarshalBinaryBare(bz, &cert)
	return cert
}

// Sets the entire Proposed Certificate record for a Subject/SubjectKeyId combination
func (k Keeper) SetProposedCertificate(ctx sdk.Context, certificate types.ProposedCertificate) {
	store := ctx.KVStore(k.storeKey)
	store.Set([]byte(ProposedCertificateId(certificate.Subject, certificate.SubjectKeyId)), k.cdc.MustMarshalBinaryBare(certificate))
}

// Check if the Proposed Certificate record associated with a Subject/SubjectKeyId combination is present in the store or not
func (k Keeper) IsProposedCertificatePresent(ctx sdk.Context, subject string, subjectKeyId string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(ProposedCertificateId(subject, subjectKeyId)))
}

// Count total Proposed Certificates
func (k Keeper) CountTotalProposedCertificates(ctx sdk.Context) int {
	return k.countTotal(ctx, proposedCertificatePrefix)
}

// Iterate over all Proposed Certificates
func (k Keeper) IterateProposedCertificates(ctx sdk.Context, process func(info types.ProposedCertificate) (stop bool)) {
	store := ctx.KVStore(k.storeKey)

	iter := sdk.KVStorePrefixIterator(store, []byte(proposedCertificatePrefix))
	defer iter.Close()

	for {
		if !iter.Valid() {
			return
		}

		val := iter.Value()

		var pendingCertificate types.ProposedCertificate

		k.cdc.MustUnmarshalBinaryBare(val, &pendingCertificate)

		if process(pendingCertificate) {
			return
		}

		iter.Next()
	}
}

// Deletes the Proposed Certificate from the store
func (k Keeper) DeleteProposedCertificate(ctx sdk.Context, subject string, subjectKeyId string) {
	if !k.IsProposedCertificatePresent(ctx, subject, subjectKeyId) {
		panic("Proposed Certificate does not exist")
	}

	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(ProposedCertificateId(subject, subjectKeyId)))
}

// Id builder for Proposed Certificate
func ProposedCertificateId(subject string, subjectKeyId string) string {
	return fmt.Sprintf("%s:%v:%v", proposedCertificatePrefix, subject, subjectKeyId)
}

func (k Keeper) countTotal(ctx sdk.Context, prefix string) int {
	store := ctx.KVStore(k.storeKey)
	res := 0

	iter := sdk.KVStorePrefixIterator(store, []byte(prefix))
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		res++
	}

	return res
}

/*
	A record contains the list of direct child certificates referring to Subject/SubjectKeyId of parent certificate
*/

// Gets the Child Certificates for a record associated with a combination Subject/SubjectKeyId
func (k Keeper) GetChildCertificates(ctx sdk.Context, subject string, subjectKeyId string) types.ChildCertificates {
	if !k.IsChildCertificatesPresent(ctx, subject, subjectKeyId) {
		return types.NewChildCertificates(subject, subjectKeyId)
	}

	store := ctx.KVStore(k.storeKey)
	bz := store.Get([]byte(ChildCertificatesId(subject, subjectKeyId)))

	var childCertificates types.ChildCertificates
	k.cdc.MustUnmarshalBinaryBare(bz, &childCertificates)
	return childCertificates
}

// Sets the entire Child Certificate record associated with a combination Subject/SubjectKeyId
func (k Keeper) SetChildCertificatesList(ctx sdk.Context, childCertificates types.ChildCertificates) {
	store := ctx.KVStore(k.storeKey)
	store.Set([]byte(ChildCertificatesId(childCertificates.Subject, childCertificates.SubjectKeyId)), k.cdc.MustMarshalBinaryBare(childCertificates))
}

// Adds a Child Certificate for a record associated with a combination Subject/SubjectKeyId
func (k Keeper) AddChildCertificate(ctx sdk.Context, subject string, subjectKeyId string, childCertificate types.Certificate) {
	store := ctx.KVStore(k.storeKey)

	childIdentifier := types.NewCertificateIdentifier(childCertificate.Subject, childCertificate.SubjectKeyId)

	certificateChain := k.GetChildCertificates(ctx, subject, subjectKeyId)
	certificateChain.ChildCertificates = append(certificateChain.ChildCertificates, childIdentifier)

	store.Set([]byte(ChildCertificatesId(subject, subjectKeyId)), k.cdc.MustMarshalBinaryBare(certificateChain))
}

// Checks if the the list of Child Certificates for a combination Subject/SubjectKeyId is present in the store or not
func (k Keeper) IsChildCertificatesPresent(ctx sdk.Context, subject string, subjectKeyId string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(ChildCertificatesId(subject, subjectKeyId)))
}

// Id builder for Child Certificates
func ChildCertificatesId(subject string, subjectKeyId string) string {
	return fmt.Sprintf("%s:%v:%v", childCertificatesPrefix, subject, subjectKeyId)
}

// Iterate over all Child Certificates records
func (k Keeper) IterateChildCertificatesRecords(ctx sdk.Context, process func(info types.ChildCertificates) (stop bool)) {
	store := ctx.KVStore(k.storeKey)

	iter := sdk.KVStorePrefixIterator(store, []byte(childCertificatesPrefix))
	defer iter.Close()

	for {
		if !iter.Valid() {
			return
		}

		val := iter.Value()

		var childCertificatesList types.ChildCertificates

		k.cdc.MustUnmarshalBinaryBare(val, &childCertificatesList)

		if process(childCertificatesList) {
			return
		}

		iter.Next()
	}
}
