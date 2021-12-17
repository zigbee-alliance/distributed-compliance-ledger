package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

// SetRevokedCertificates set a specific revokedCertificates in the store from its index
func (k Keeper) SetRevokedCertificates(ctx sdk.Context, revokedCertificates types.RevokedCertificates) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RevokedCertificatesKeyPrefix))
	b := k.cdc.MustMarshal(&revokedCertificates)
	store.Set(types.RevokedCertificatesKey(
		revokedCertificates.Subject,
		revokedCertificates.SubjectKeyId,
	), b)
}

// GetRevokedCertificates returns a revokedCertificates from its index
func (k Keeper) GetRevokedCertificates(
	ctx sdk.Context,
	subject string,
	subjectKeyId string,

) (val types.RevokedCertificates, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RevokedCertificatesKeyPrefix))

	b := store.Get(types.RevokedCertificatesKey(
		subject,
		subjectKeyId,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveRevokedCertificates removes a revokedCertificates from the store
func (k Keeper) RemoveRevokedCertificates(
	ctx sdk.Context,
	subject string,
	subjectKeyId string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RevokedCertificatesKeyPrefix))
	store.Delete(types.RevokedCertificatesKey(
		subject,
		subjectKeyId,
	))
}

// GetAllRevokedCertificates returns all revokedCertificates
func (k Keeper) GetAllRevokedCertificates(ctx sdk.Context) (list []types.RevokedCertificates) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RevokedCertificatesKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.RevokedCertificates
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
