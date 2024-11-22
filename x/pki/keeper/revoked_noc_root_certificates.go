package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

// SetRevokedNocRootCertificates set a specific revokedNocRootCertificates in the store from its index.
func (k Keeper) SetRevokedNocRootCertificates(ctx sdk.Context, revokedNocRootCertificates types.RevokedNocRootCertificates) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.RevokedNocRootCertificatesKeyPrefix))
	b := k.cdc.MustMarshal(&revokedNocRootCertificates)
	store.Set(types.RevokedNocRootCertificatesKey(
		revokedNocRootCertificates.Subject,
		revokedNocRootCertificates.SubjectKeyId,
	), b)
}

// AddRevokedNocRootCertificates adds revoked NOC certificates to the list of revoked NOC certificates for the subject/subjectKeyId map.
func (k Keeper) AddRevokedNocRootCertificates(ctx sdk.Context, revokedNocRootCertificates types.RevokedNocRootCertificates) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.RevokedNocRootCertificatesKeyPrefix))

	revokedCertsBytes := store.Get(types.RevokedNocRootCertificatesKey(
		revokedNocRootCertificates.Subject,
		revokedNocRootCertificates.SubjectKeyId,
	))
	var revokedCerts types.RevokedNocRootCertificates

	if revokedCertsBytes == nil {
		revokedCerts = types.RevokedNocRootCertificates{
			Subject:      revokedNocRootCertificates.Subject,
			SubjectKeyId: revokedNocRootCertificates.SubjectKeyId,
			Certs:        []*types.Certificate{},
		}
	} else {
		k.cdc.MustUnmarshal(revokedCertsBytes, &revokedCerts)
	}

	revokedCerts.Certs = append(revokedCerts.Certs, revokedNocRootCertificates.Certs...)

	b := k.cdc.MustMarshal(&revokedCerts)
	store.Set(types.RevokedNocRootCertificatesKey(
		revokedCerts.Subject,
		revokedCerts.SubjectKeyId,
	), b)
}

// GetRevokedNocRootCertificates returns a revokedNocRootCertificates from its index.
func (k Keeper) GetRevokedNocRootCertificates(
	ctx sdk.Context,
	subject string,
	subjectKeyID string,

) (val types.RevokedNocRootCertificates, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.RevokedNocRootCertificatesKeyPrefix))

	b := store.Get(types.RevokedNocRootCertificatesKey(
		subject,
		subjectKeyID,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)

	return val, true
}

// RemoveRevokedNocRootCertificates removes a revokedNocRootCertificates from the store.
func (k Keeper) RemoveRevokedNocRootCertificates(
	ctx sdk.Context,
	subject string,
	subjectKeyID string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.RevokedNocRootCertificatesKeyPrefix))
	store.Delete(types.RevokedNocRootCertificatesKey(
		subject,
		subjectKeyID,
	))
}

// GetAllRevokedNocRootCertificates returns all revokedNocRootCertificates.
func (k Keeper) GetAllRevokedNocRootCertificates(ctx sdk.Context) (list []types.RevokedNocRootCertificates) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.RevokedNocRootCertificatesKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.RevokedNocRootCertificates
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// IsRevokedNocRootCertificatePresent Check if the Revoked Noc Root Certificate record associated with a Subject/SubjectKeyID combination is present in the store.
func (k Keeper) IsRevokedNocRootCertificatePresent(
	ctx sdk.Context,
	subject string,
	subjectKeyID string,
) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.RevokedNocRootCertificatesKeyPrefix))

	return store.Has(types.RevokedNocRootCertificatesKey(
		subject,
		subjectKeyID,
	))
}
