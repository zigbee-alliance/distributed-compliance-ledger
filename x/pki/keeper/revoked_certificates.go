package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

// SetRevokedCertificates set a specific revokedCertificates in the store from its index.
func (k Keeper) SetRevokedCertificates(ctx sdk.Context, revokedCertificates types.RevokedCertificates) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.RevokedCertificatesKeyPrefix))
	b := k.cdc.MustMarshal(&revokedCertificates)
	store.Set(types.RevokedCertificatesKey(
		revokedCertificates.Subject,
		revokedCertificates.SubjectKeyId,
	), b)
}

// GetRevokedCertificates returns a revokedCertificates from its index.
func (k Keeper) GetRevokedCertificates(
	ctx sdk.Context,
	subject string,
	subjectKeyID string,
) (val types.RevokedCertificates, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.RevokedCertificatesKeyPrefix))

	b := store.Get(types.RevokedCertificatesKey(
		subject,
		subjectKeyID,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)

	return val, true
}

// RemoveRevokedCertificates removes a revokedCertificates from the store.
func (k Keeper) RemoveRevokedCertificates(
	ctx sdk.Context,
	subject string,
	subjectKeyID string,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.RevokedCertificatesKeyPrefix))
	store.Delete(types.RevokedCertificatesKey(
		subject,
		subjectKeyID,
	))
}

// GetAllRevokedCertificates returns all revokedCertificates.
func (k Keeper) GetAllRevokedCertificates(ctx sdk.Context) (list []types.RevokedCertificates) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.RevokedCertificatesKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.RevokedCertificates
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// AddRevokedCertificates Add revoked certificates to the list of revoked certificates for the subject/subjectKeyId map.
func (k Keeper) AddRevokedCertificates(ctx sdk.Context, approvedCertificates types.RevokedCertificates) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.RevokedCertificatesKeyPrefix))

	revokedCertificatesBytes := store.Get(types.RevokedCertificatesKey(
		approvedCertificates.Subject,
		approvedCertificates.SubjectKeyId,
	))
	var revokedCertificates types.RevokedCertificates

	if revokedCertificatesBytes == nil {
		revokedCertificates = types.RevokedCertificates{
			Subject:      approvedCertificates.Subject,
			SubjectKeyId: approvedCertificates.SubjectKeyId,
			Certs:        []*types.Certificate{},
		}
	} else {
		k.cdc.MustUnmarshal(revokedCertificatesBytes, &revokedCertificates)
	}

	revokedCertificates.Certs = append(revokedCertificates.Certs, approvedCertificates.Certs...)

	b := k.cdc.MustMarshal(&revokedCertificates)
	store.Set(types.RevokedCertificatesKey(
		revokedCertificates.Subject,
		revokedCertificates.SubjectKeyId,
	), b)
}

func (k msgServer) removeOrUpdateRevokedX509Cert(
	ctx sdk.Context,
	subject string,
	subjectKeyID string,
	certificates *types.RevokedCertificates) {
	if len(certificates.Certs) == 0 {
		k.RemoveRevokedCertificates(ctx, subject, subjectKeyID)
	} else {
		k.SetRevokedCertificates(
			ctx,
			*certificates,
		)
	}
}

// IsRevokedCertificatePresent Check if the Revoked Certificate is present in the store.
func (k Keeper) IsRevokedCertificatePresent(
	ctx sdk.Context,
	subject string,
	subjectKeyID string,
) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.RevokedCertificatesKeyPrefix))

	return store.Has(types.RevokedCertificatesKey(
		subject,
		subjectKeyID,
	))
}
