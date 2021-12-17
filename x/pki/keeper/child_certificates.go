package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

// SetChildCertificates set a specific childCertificates in the store from its index
func (k Keeper) SetChildCertificates(ctx sdk.Context, childCertificates types.ChildCertificates) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChildCertificatesKeyPrefix))
	b := k.cdc.MustMarshal(&childCertificates)
	store.Set(types.ChildCertificatesKey(
		childCertificates.Issuer,
		childCertificates.AuthorityKeyId,
	), b)
}

// GetChildCertificates returns a childCertificates from its index
func (k Keeper) GetChildCertificates(
	ctx sdk.Context,
	issuer string,
	authorityKeyId string,

) (val types.ChildCertificates, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChildCertificatesKeyPrefix))

	b := store.Get(types.ChildCertificatesKey(
		issuer,
		authorityKeyId,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveChildCertificates removes a childCertificates from the store
func (k Keeper) RemoveChildCertificates(
	ctx sdk.Context,
	issuer string,
	authorityKeyId string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChildCertificatesKeyPrefix))
	store.Delete(types.ChildCertificatesKey(
		issuer,
		authorityKeyId,
	))
}

// GetAllChildCertificates returns all childCertificates
func (k Keeper) GetAllChildCertificates(ctx sdk.Context) (list []types.ChildCertificates) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChildCertificatesKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ChildCertificates
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
