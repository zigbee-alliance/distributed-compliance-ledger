package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

// SetRevokedNocIcaCertificates set a specific revokedNocIcaCertificates in the store from its index
func (k Keeper) SetRevokedNocIcaCertificates(ctx sdk.Context, revokedNocIcaCertificates types.RevokedNocIcaCertificates) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.RevokedNocIcaCertificatesKeyPrefix))
	b := k.cdc.MustMarshal(&revokedNocIcaCertificates)
	store.Set(types.RevokedNocIcaCertificatesKey(
		revokedNocIcaCertificates.Subject,
		revokedNocIcaCertificates.SubjectKeyId,
	), b)
}

// AddRevokedNocIcaCertificates adds revoked NOC certificates to the list of revoked NOC certificates for the subject/subjectKeyId map.
func (k Keeper) AddRevokedNocIcaCertificates(ctx sdk.Context, revokedNocIcaCertificates types.RevokedNocIcaCertificates) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.RevokedNocIcaCertificatesKeyPrefix))

	revokedCertsBytes := store.Get(types.RevokedNocIcaCertificatesKey(
		revokedNocIcaCertificates.Subject,
		revokedNocIcaCertificates.SubjectKeyId,
	))
	var revokedCerts types.RevokedNocIcaCertificates

	if revokedCertsBytes == nil {
		revokedCerts = types.RevokedNocIcaCertificates{
			Subject:      revokedNocIcaCertificates.Subject,
			SubjectKeyId: revokedNocIcaCertificates.SubjectKeyId,
			Certs:        []*types.Certificate{},
		}
	} else {
		k.cdc.MustUnmarshal(revokedCertsBytes, &revokedCerts)
	}

	revokedCerts.Certs = append(revokedCerts.Certs, revokedNocIcaCertificates.Certs...)

	b := k.cdc.MustMarshal(&revokedCerts)
	store.Set(types.RevokedNocIcaCertificatesKey(
		revokedCerts.Subject,
		revokedCerts.SubjectKeyId,
	), b)
}

// GetRevokedNocIcaCertificates returns a revokedNocIcaCertificates from its index
func (k Keeper) GetRevokedNocIcaCertificates(
	ctx sdk.Context,
	subject string,
	subjectKeyId string,

) (val types.RevokedNocIcaCertificates, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.RevokedNocIcaCertificatesKeyPrefix))

	b := store.Get(types.RevokedNocIcaCertificatesKey(
		subject,
		subjectKeyId,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)

	return val, true
}

// RemoveRevokedNocIcaCertificates removes a revokedNocIcaCertificates from the store
func (k Keeper) RemoveRevokedNocIcaCertificates(
	ctx sdk.Context,
	subject string,
	subjectKeyID string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.RevokedNocIcaCertificatesKeyPrefix))
	store.Delete(types.RevokedNocIcaCertificatesKey(
		subject,
		subjectKeyID,
	))
}

// GetAllRevokedNocIcaCertificates returns all revokedNocIcaCertificates
func (k Keeper) GetAllRevokedNocIcaCertificates(ctx sdk.Context) (list []types.RevokedNocIcaCertificates) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.RevokedNocIcaCertificatesKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.RevokedNocIcaCertificates
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
