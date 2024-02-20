package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

// SetNocRootCertificates set a specific nocRootCertificates in the store from its index.
func (k Keeper) SetNocRootCertificates(ctx sdk.Context, nocRootCertificates types.NocRootCertificates) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.NocRootCertificatesKeyPrefix))
	b := k.cdc.MustMarshal(&nocRootCertificates)
	store.Set(types.NocRootCertificatesKey(
		nocRootCertificates.Vid,
	), b)
}

// GetNocRootCertificates returns a nocRootCertificates from its index.
func (k Keeper) GetNocRootCertificates(
	ctx sdk.Context,
	vid int32,

) (val types.NocRootCertificates, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.NocRootCertificatesKeyPrefix))

	b := store.Get(types.NocRootCertificatesKey(
		vid,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)

	return val, true
}

// Add a NOC root certificate to the list of NOC root certificates for the VID map.
func (k Keeper) AddNocRootCertificate(ctx sdk.Context, nocCertificate types.Certificate) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.NocRootCertificatesKeyPrefix))

	nocCertificatesBytes := store.Get(types.NocRootCertificatesKey(nocCertificate.Vid))
	var nocCertificates types.NocRootCertificates

	if nocCertificatesBytes == nil {
		nocCertificates = types.NocRootCertificates{
			Vid:   nocCertificate.Vid,
			Certs: []*types.Certificate{},
		}
	} else {
		k.cdc.MustUnmarshal(nocCertificatesBytes, &nocCertificates)
	}

	nocCertificates.Certs = append(nocCertificates.Certs, &nocCertificate)

	b := k.cdc.MustMarshal(&nocCertificates)
	store.Set(types.NocRootCertificatesKey(nocCertificate.Vid), b)
}

// RemoveNocRootCertificates removes a nocRootCertificates from the store.
func (k Keeper) RemoveNocRootCertificates(
	ctx sdk.Context,
	vid int32,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.NocRootCertificatesKeyPrefix))
	store.Delete(types.NocRootCertificatesKey(
		vid,
	))
}

// GetAllNocRootCertificates returns all nocRootCertificates.
func (k Keeper) GetAllNocRootCertificates(ctx sdk.Context) (list []types.NocRootCertificates) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.NocRootCertificatesKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.NocRootCertificates
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
