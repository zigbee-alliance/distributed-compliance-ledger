package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

// SetNocRootCertificatesByVidAndSkid set a specific nocRootCertificatesByVidAndSkid in the store from its index
func (k Keeper) SetNocRootCertificatesByVidAndSkid(ctx sdk.Context, nocRootCertificatesByVidAndSkid types.NocRootCertificatesByVidAndSkid) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.NocRootCertificatesByVidAndSkidKeyPrefix))
	b := k.cdc.MustMarshal(&nocRootCertificatesByVidAndSkid)
	store.Set(types.NocRootCertificatesByVidAndSkidKey(
		nocRootCertificatesByVidAndSkid.Vid,
		nocRootCertificatesByVidAndSkid.SubjectKeyId,
	), b)
}

// GetNocRootCertificatesByVidAndSkid returns a nocRootCertificatesByVidAndSkid from its index
func (k Keeper) GetNocRootCertificatesByVidAndSkid(
	ctx sdk.Context,
	vid int32,
	subjectKeyID string,

) (val types.NocRootCertificatesByVidAndSkid, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.NocRootCertificatesByVidAndSkidKeyPrefix))

	b := store.Get(types.NocRootCertificatesByVidAndSkidKey(
		vid,
		subjectKeyID,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// Add a NOC root certificate to the list of NOC root certificates for the VID map.
func (k Keeper) AddNocRootCertificatesByVidAndSkid(ctx sdk.Context, nocCertificate types.Certificate) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.NocRootCertificatesKeyPrefix))

	nocRootCertificatesByVidAndSkidKeyBytes := store.Get(types.NocRootCertificatesByVidAndSkidKey(nocCertificate.Vid, nocCertificate.SubjectKeyId))
	var nocRootCertificatesByVidAndSkid types.NocRootCertificatesByVidAndSkid

	if nocRootCertificatesByVidAndSkidKeyBytes == nil {
		nocRootCertificatesByVidAndSkid = types.NocRootCertificatesByVidAndSkid{
			Vid:          nocCertificate.Vid,
			SubjectKeyId: nocCertificate.SubjectKeyId,
			Certs:        []*types.Certificate{},
			Tq:           1,
		}
	} else {
		k.cdc.MustUnmarshal(nocRootCertificatesByVidAndSkidKeyBytes, &nocRootCertificatesByVidAndSkid)
	}

	nocRootCertificatesByVidAndSkid.Certs = append(nocRootCertificatesByVidAndSkid.Certs, &nocCertificate)

	b := k.cdc.MustMarshal(&nocRootCertificatesByVidAndSkid)
	store.Set(types.NocRootCertificatesKey(nocCertificate.Vid), b)
}

// RemoveNocRootCertificatesByVidAndSkid removes a nocRootCertificatesByVidAndSkid from the store
func (k Keeper) RemoveNocRootCertificatesByVidAndSkid(
	ctx sdk.Context,
	vid int32,
	subjectKeyID string,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.NocRootCertificatesByVidAndSkidKeyPrefix))
	store.Delete(types.NocRootCertificatesByVidAndSkidKey(
		vid,
		subjectKeyID,
	))
}

// GetAllNocRootCertificatesByVidAndSkid returns all nocRootCertificatesByVidAndSkid
func (k Keeper) GetAllNocRootCertificatesByVidAndSkid(ctx sdk.Context) (list []types.NocRootCertificatesByVidAndSkid) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.NocRootCertificatesByVidAndSkidKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.NocRootCertificatesByVidAndSkid
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
