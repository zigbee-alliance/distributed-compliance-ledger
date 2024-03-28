package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

// SetChildCertificates set a specific childCertificates in the store from its index.
func (k Keeper) SetChildCertificates(ctx sdk.Context, childCertificates types.ChildCertificates) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.ChildCertificatesKeyPrefix))
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
	authorityKeyID string,
) (val types.ChildCertificates, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.ChildCertificatesKeyPrefix))

	b := store.Get(types.ChildCertificatesKey(
		issuer,
		authorityKeyID,
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
	authorityKeyID string,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.ChildCertificatesKeyPrefix))
	store.Delete(types.ChildCertificatesKey(
		issuer,
		authorityKeyID,
	))
}

// GetAllChildCertificates returns all childCertificates.
func (k Keeper) GetAllChildCertificates(ctx sdk.Context) (list []types.ChildCertificates) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.ChildCertificatesKeyPrefix))
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
func (k Keeper) AddChildCertificate(ctx sdk.Context, issuer string, authorityKeyID string, certID types.CertificateIdentifier) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.ChildCertificatesKeyPrefix))

	childCertificatesBytes := store.Get(types.ChildCertificatesKey(
		issuer,
		authorityKeyID,
	))

	var childCertificates types.ChildCertificates
	if childCertificatesBytes == nil {
		childCertificates = types.ChildCertificates{
			Issuer:         issuer,
			AuthorityKeyId: authorityKeyID,
			CertIds:        []*types.CertificateIdentifier{},
		}
	} else {
		k.cdc.MustUnmarshal(childCertificatesBytes, &childCertificates)
	}

	for _, existingCertID := range childCertificates.CertIds {
		if *existingCertID == certID {
			return
		}
	}

	childCertificates.CertIds = append(childCertificates.CertIds, &certID)

	b := k.cdc.MustMarshal(&childCertificates)
	store.Set(types.ChildCertificatesKey(
		issuer,
		authorityKeyID,
	), b)
}

func (k msgServer) RevokeChildCertificates(ctx sdk.Context, issuer string, authorityKeyID string, schemaVersion uint32) {
	// Get issuer's ChildCertificates record
	childCertificates, _ := k.GetChildCertificates(ctx, issuer, authorityKeyID)

	// For each child certificate subject/subjectKeyID combination
	for _, certIdentifier := range childCertificates.CertIds {
		// Revoke certificates with this subject/subjectKeyID combination
		certificates, _ := k.GetApprovedCertificates(ctx, certIdentifier.Subject, certIdentifier.SubjectKeyId)
		k.AddRevokedCertificates(ctx, certificates, schemaVersion)
		// FIXME: Below two lines is not in the context of RevokeChildCertificates method. In future current implementation must be refactored
		if len(certificates.Certs) > 0 {
			// If cert is NOC then remove it from NOC certificates list
			k.RemoveNocCertificate(ctx, certIdentifier.Subject, certIdentifier.SubjectKeyId, certificates.Certs[0].Vid)
		}
		k.RemoveApprovedCertificates(ctx, certIdentifier.Subject, certIdentifier.SubjectKeyId)

		// remove from subject -> subject key ID map
		k.RemoveApprovedCertificateBySubject(ctx, certIdentifier.Subject, certIdentifier.SubjectKeyId)

		// remove from subject key ID -> certificates map
		k.RemoveApprovedCertificatesBySubjectKeyID(ctx, certIdentifier.Subject, certIdentifier.SubjectKeyId)

		// Process child certificates recursively
		k.RevokeChildCertificates(ctx, certIdentifier.Subject, certIdentifier.SubjectKeyId, schemaVersion)
	}

	// Delete entire ChildCertificates record of issuer
	k.RemoveChildCertificates(ctx, issuer, authorityKeyID)
}

func (k msgServer) RemoveChildCertificate(ctx sdk.Context, issuer string, authorityKeyID string,
	certIdentifier types.CertificateIdentifier,
) {
	childCertificates, _ := k.GetChildCertificates(ctx, issuer, authorityKeyID)

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

	childCertificates.CertIds = append(childCertificates.CertIds[:certIDIndex], childCertificates.CertIds[certIDIndex+1:]...)

	if len(childCertificates.CertIds) > 0 {
		k.SetChildCertificates(ctx, childCertificates)
	} else {
		k.RemoveChildCertificates(ctx, issuer, authorityKeyID)
	}
}
