package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

// SetUniqueCertificate set a specific uniqueCertificate in the store from its index
func (k Keeper) SetUniqueCertificate(ctx sdk.Context, uniqueCertificate types.UniqueCertificate) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UniqueCertificateKeyPrefix))
	b := k.cdc.MustMarshal(&uniqueCertificate)
	store.Set(types.UniqueCertificateKey(
		uniqueCertificate.Issuer,
		uniqueCertificate.SerialNumber,
	), b)
}

// GetUniqueCertificate returns a uniqueCertificate from its index
func (k Keeper) GetUniqueCertificate(
	ctx sdk.Context,
	issuer string,
	serialNumber string,

) (val types.UniqueCertificate, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UniqueCertificateKeyPrefix))

	b := store.Get(types.UniqueCertificateKey(
		issuer,
		serialNumber,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveUniqueCertificate removes a uniqueCertificate from the store
func (k Keeper) RemoveUniqueCertificate(
	ctx sdk.Context,
	issuer string,
	serialNumber string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UniqueCertificateKeyPrefix))
	store.Delete(types.UniqueCertificateKey(
		issuer,
		serialNumber,
	))
}

// GetAllUniqueCertificate returns all uniqueCertificate
func (k Keeper) GetAllUniqueCertificate(ctx sdk.Context) (list []types.UniqueCertificate) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UniqueCertificateKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.UniqueCertificate
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// Check if the unique certificate key (Issuer/SerialNumber combination) is busy.
func (k Keeper) IsUniqueCertificatePresent(
	ctx sdk.Context,
	issuer string,
	serialNumber string,

) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UniqueCertificateKeyPrefix))
	return store.Has(types.UniqueCertificateKey(
		issuer,
		serialNumber,
	))
}
