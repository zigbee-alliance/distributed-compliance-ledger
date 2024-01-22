package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

// SetApprovedCertificatesBySubjectKeyID set a specific approvedCertificatesBySubjectKeyId in the store from its index.
func (k Keeper) SetApprovedCertificatesBySubjectKeyID(ctx sdk.Context, approvedCertificatesBySubjectKeyID types.ApprovedCertificatesBySubjectKeyId) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.ApprovedCertificatesBySubjectKeyIDKeyPrefix))
	b := k.cdc.MustMarshal(&approvedCertificatesBySubjectKeyID)
	store.Set(types.ApprovedCertificatesBySubjectKeyIDKey(
		approvedCertificatesBySubjectKeyID.SubjectKeyId,
	), b)
}

// Add an approved certificate to the list of approved certificates with the subjectKeyId map.
func (k Keeper) AddApprovedCertificateBySubjectKeyID(ctx sdk.Context, certificate types.Certificate) {
	k.addApprovedCertificates(ctx, certificate.SubjectKeyId, []*types.Certificate{&certificate})
}

// Add an approved certificates list to approved certificates with the subjectKeyId map.
func (k Keeper) AddApprovedCertificatesBySubjectKeyID(ctx sdk.Context, approvedCertificate types.ApprovedCertificates) {
	k.addApprovedCertificates(ctx, approvedCertificate.SubjectKeyId, approvedCertificate.Certs)
}

func (k Keeper) addApprovedCertificates(ctx sdk.Context, subjectKeyID string, certs []*types.Certificate) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.ApprovedCertificatesBySubjectKeyIDKeyPrefix))

	approvedCertificatesBytes := store.Get(types.ApprovedCertificatesBySubjectKey(
		subjectKeyID,
	))
	var approvedCertificates types.ApprovedCertificatesBySubjectKeyId

	if approvedCertificatesBytes == nil {
		approvedCertificates = types.ApprovedCertificatesBySubjectKeyId{
			SubjectKeyId: subjectKeyID,
			Certs:        []*types.Certificate{},
		}
	} else {
		k.cdc.MustUnmarshal(approvedCertificatesBytes, &approvedCertificates)
	}

	approvedCertificates.Certs = append(approvedCertificates.Certs, certs...)

	k.SetApprovedCertificatesBySubjectKeyID(ctx, approvedCertificates)
}

// GetApprovedCertificatesBySubjectKeyID returns a approvedCertificatesBySubjectKeyId from its index.
func (k Keeper) GetApprovedCertificatesBySubjectKeyID(
	ctx sdk.Context,
	subjectKeyID string,

) (val types.ApprovedCertificatesBySubjectKeyId, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.ApprovedCertificatesBySubjectKeyIDKeyPrefix))

	b := store.Get(types.ApprovedCertificatesBySubjectKeyIDKey(
		subjectKeyID,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)

	return val, true
}

// RemoveApprovedCertificatesBySubjectKeyID removes a approvedCertificatesBySubjectKeyId from the store.
func (k Keeper) RemoveApprovedCertificatesBySubjectKeyID(
	ctx sdk.Context,
	subject string,
	subjectKeyID string,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.ApprovedCertificatesBySubjectKeyIDKeyPrefix))
	certs, found := k.GetApprovedCertificatesBySubjectKeyID(ctx, subjectKeyID)
	if !found {
		return
	}

	var remainedCerts = certs.Certs
	for i, cert := range certs.Certs {
		if cert.Subject == subject {
			if i+1 != len(certs.Certs) {
				//nolint:gocritic
				remainedCerts = append(certs.Certs[:i], certs.Certs[i+1:]...)
			} else {
				remainedCerts = remainedCerts[:i]
			}
		}
	}

	if len(remainedCerts) == 0 {
		store.Delete(types.ApprovedCertificatesBySubjectKeyIDKey(
			subjectKeyID,
		))
	} else {
		certs.Certs = remainedCerts
		k.SetApprovedCertificatesBySubjectKeyID(ctx, certs)
	}
}

// GetAllApprovedCertificatesBySubjectKeyID returns all approvedCertificatesBySubjectKeyId.
func (k Keeper) GetAllApprovedCertificatesBySubjectKeyID(ctx sdk.Context) (list []types.ApprovedCertificatesBySubjectKeyId) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.ApprovedCertificatesBySubjectKeyIDKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ApprovedCertificatesBySubjectKeyId
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
