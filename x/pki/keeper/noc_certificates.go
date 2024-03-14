package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

// SetNocCertificates set a specific nocCertificates in the store from its index.
func (k Keeper) SetNocCertificates(ctx sdk.Context, nocCertificates types.NocCertificates) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.NocCertificatesKeyPrefix))
	b := k.cdc.MustMarshal(&nocCertificates)
	store.Set(types.NocCertificatesKey(
		nocCertificates.Vid,
	), b)
}

// GetNocCertificates returns a nocCertificates from its index.
func (k Keeper) GetNocCertificates(
	ctx sdk.Context,
	vid int32,

) (val types.NocCertificates, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.NocCertificatesKeyPrefix))

	b := store.Get(types.NocCertificatesKey(
		vid,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)

	return val, true
}

// AddNocCertificate adds a NOC certificate to the list of NOC certificates for the VID map.
func (k Keeper) AddNocCertificate(ctx sdk.Context, nocCertificate types.Certificate) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.NocCertificatesKeyPrefix))

	nocCertificatesBytes := store.Get(types.NocCertificatesKey(nocCertificate.Vid))
	var nocCertificates types.NocCertificates

	if nocCertificatesBytes == nil {
		nocCertificates = types.NocCertificates{
			Vid:   nocCertificate.Vid,
			Certs: []*types.Certificate{},
		}
	} else {
		k.cdc.MustUnmarshal(nocCertificatesBytes, &nocCertificates)
	}

	nocCertificates.Certs = append(nocCertificates.Certs, &nocCertificate)

	b := k.cdc.MustMarshal(&nocCertificates)
	store.Set(types.NocCertificatesKey(nocCertificate.Vid), b)
}

// RemoveNocCertificates removes a nocCertificates from the store.
func (k Keeper) RemoveNocCertificates(
	ctx sdk.Context,
	vid int32,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.NocCertificatesKeyPrefix))
	store.Delete(types.NocCertificatesKey(
		vid,
	))
}

func (k Keeper) RemoveNocCertificate(ctx sdk.Context, subject, subjectKeyID string, vid int32) {
	certs, found := k.GetNocCertificates(ctx, vid)
	if !found {
		return
	}

	for i := 0; i < len(certs.Certs); {
		if certs.Certs[i].Subject == subject && certs.Certs[i].SubjectKeyId == subjectKeyID {
			certs.Certs = append(certs.Certs[:i], certs.Certs[i+1:]...)
		} else {
			i++
		}
	}

	if len(certs.Certs) == 0 {
		k.RemoveNocCertificates(ctx, vid)
	} else {
		k.SetNocCertificates(ctx, certs)
	}
}

// GetAllNocCertificates returns all nocCertificates.
func (k Keeper) GetAllNocCertificates(ctx sdk.Context) (list []types.NocCertificates) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.NocCertificatesKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.NocCertificates
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
