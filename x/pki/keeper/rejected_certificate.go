package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

// SetRejectedCertificate set a specific rejectedCertificate in the store from its index
func (k Keeper) SetRejectedCertificate(ctx sdk.Context, rejectedCertificate types.RejectedCertificate) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RejectedCertificateKeyPrefix))
	b := k.cdc.MustMarshal(&rejectedCertificate)
	store.Set(types.RejectedCertificateKey(
		rejectedCertificate.Subject,
		rejectedCertificate.SubjectKeyId,
	), b)
}

// GetRejectedCertificate returns a rejectedCertificate from its index
func (k Keeper) GetRejectedCertificate(
	ctx sdk.Context,
	subject string,
	subjectKeyId string,

) (val types.RejectedCertificate, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RejectedCertificateKeyPrefix))

	b := store.Get(types.RejectedCertificateKey(
		subject,
		subjectKeyId,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveRejectedCertificate removes a rejectedCertificate from the store
func (k Keeper) RemoveRejectedCertificate(
	ctx sdk.Context,
	subject string,
	subjectKeyId string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RejectedCertificateKeyPrefix))
	store.Delete(types.RejectedCertificateKey(
		subject,
		subjectKeyId,
	))
}

// GetAllRejectedCertificate returns all rejectedCertificate
func (k Keeper) GetAllRejectedCertificate(ctx sdk.Context) (list []types.RejectedCertificate) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RejectedCertificateKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.RejectedCertificate
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
