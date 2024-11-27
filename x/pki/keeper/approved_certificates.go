package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

// SetApprovedCertificates set a specific approvedCertificates in the store from its index.
func (k Keeper) SetApprovedCertificates(ctx sdk.Context, approvedCertificates types.ApprovedCertificates) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.ApprovedCertificatesKeyPrefix))
	b := k.cdc.MustMarshal(&approvedCertificates)
	store.Set(types.ApprovedCertificatesKey(
		approvedCertificates.Subject,
		approvedCertificates.SubjectKeyId,
	), b)
}

// GetApprovedCertificates returns a approvedCertificates from its index.
func (k Keeper) GetApprovedCertificates(
	ctx sdk.Context,
	subject string,
	subjectKeyID string,
) (val types.ApprovedCertificates, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.ApprovedCertificatesKeyPrefix))

	b := store.Get(types.ApprovedCertificatesKey(
		subject,
		subjectKeyID,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)

	return val, true
}

// RemoveApprovedCertificates removes a approvedCertificates from the store.
func (k Keeper) RemoveApprovedCertificates(
	ctx sdk.Context,
	subject string,
	subjectKeyID string,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.ApprovedCertificatesKeyPrefix))
	store.Delete(types.ApprovedCertificatesKey(
		subject,
		subjectKeyID,
	))
}

func (k Keeper) RemoveApprovedCertificatesBySerialNumber(ctx sdk.Context, subject, subjectKeyID, serialNumber string) {
	k._removeApprovedCertificatesBySerialNumber(ctx, subject, subjectKeyID, func(cert *types.Certificate) bool {
		return cert.Subject == subject && cert.SubjectKeyId == subjectKeyID && cert.SerialNumber == serialNumber
	})
}

func (k Keeper) _removeApprovedCertificatesBySerialNumber(ctx sdk.Context, subject string, subjectKeyID string, filter func(cert *types.Certificate) bool) {
	certs, found := k.GetApprovedCertificates(ctx, subject, subjectKeyID)
	if !found {
		return
	}

	numCertsBefore := len(certs.Certs)
	for i := 0; i < len(certs.Certs); {
		cert := certs.Certs[i]
		if filter(cert) {
			certs.Certs = append(certs.Certs[:i], certs.Certs[i+1:]...)
		} else {
			i++
		}
	}

	if len(certs.Certs) == 0 {
		store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.ApprovedCertificatesKeyPrefix))
		store.Delete(types.AllCertificatesKey(
			subject,
			subjectKeyID,
		))
	} else if numCertsBefore > len(certs.Certs) { // Update state only if any certificate is removed
		k.SetApprovedCertificates(ctx, certs)
	}
}

// GetAllApprovedCertificates returns all approvedCertificates.
func (k Keeper) GetAllApprovedCertificates(ctx sdk.Context) (list []types.ApprovedCertificates) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.ApprovedCertificatesKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ApprovedCertificates
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// Add an approved certificate to the list of approved certificates for the subject/subjectKeyId map.
func (k Keeper) AddApprovedCertificate(ctx sdk.Context, approvedCertificate types.Certificate) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.ApprovedCertificatesKeyPrefix))

	approvedCertificatesBytes := store.Get(types.ApprovedCertificatesKey(
		approvedCertificate.Subject,
		approvedCertificate.SubjectKeyId,
	))
	var approvedCertificates types.ApprovedCertificates

	if approvedCertificatesBytes == nil {
		approvedCertificates = types.ApprovedCertificates{
			Subject:      approvedCertificate.Subject,
			SubjectKeyId: approvedCertificate.SubjectKeyId,
			Certs:        []*types.Certificate{},
		}
	} else {
		k.cdc.MustUnmarshal(approvedCertificatesBytes, &approvedCertificates)
	}

	approvedCertificates.Certs = append(approvedCertificates.Certs, &approvedCertificate)

	b := k.cdc.MustMarshal(&approvedCertificates)
	store.Set(types.ApprovedCertificatesKey(
		approvedCertificates.Subject,
		approvedCertificates.SubjectKeyId,
	), b)
}

// IsApprovedCertificatesPresent Check if the Approved Certificate is present in the store.
func (k Keeper) IsApprovedCertificatesPresent(
	ctx sdk.Context,
	subject string,
	subjectKeyID string,
) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.ApprovedCertificatesKeyPrefix))

	return store.Has(types.ApprovedCertificatesKey(
		subject,
		subjectKeyID,
	))
}
