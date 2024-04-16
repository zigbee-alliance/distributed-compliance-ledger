package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

// SetNocIcaCertificates set a specific NocIcaCertificates in the store from its index.
func (k Keeper) SetNocIcaCertificates(ctx sdk.Context, nocIcaCertificates types.NocIcaCertificates) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.NocIcaCertificatesKeyPrefix))
	b := k.cdc.MustMarshal(&nocIcaCertificates)
	store.Set(types.NocIcaCertificatesKey(
		nocIcaCertificates.Vid,
	), b)
}

// GetNocIcaCertificates returns a NocIcaCertificates from its index.
func (k Keeper) GetNocIcaCertificates(
	ctx sdk.Context,
	vid int32,

) (val types.NocIcaCertificates, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.NocIcaCertificatesKeyPrefix))

	b := store.Get(types.NocIcaCertificatesKey(
		vid,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)

	return val, true
}

// AddNocIcaCertificates adds a NOC certificate to the list of NOC certificates for the VID map.
func (k Keeper) AddNocIcaCertificate(ctx sdk.Context, nocIcaCertificates types.Certificate) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.NocIcaCertificatesKeyPrefix))

	NocIcaCertificatesBytes := store.Get(types.NocIcaCertificatesKey(nocIcaCertificates.Vid))
	var NocIcaCertificates types.NocIcaCertificates

	if NocIcaCertificatesBytes == nil {
		NocIcaCertificates = types.NocIcaCertificates{
			Vid:   nocIcaCertificates.Vid,
			Certs: []*types.Certificate{},
		}
	} else {
		k.cdc.MustUnmarshal(NocIcaCertificatesBytes, &NocIcaCertificates)
	}

	NocIcaCertificates.Certs = append(NocIcaCertificates.Certs, &nocIcaCertificates)

	b := k.cdc.MustMarshal(&NocIcaCertificates)
	store.Set(types.NocIcaCertificatesKey(nocIcaCertificates.Vid), b)
}

// RemoveNocIcaCertificates removes a NocIcaCertificates from the store.
func (k Keeper) RemoveNocIcaCertificates(ctx sdk.Context, vid int32) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.NocIcaCertificatesKeyPrefix))
	store.Delete(types.NocIcaCertificatesKey(
		vid,
	))
}

func (k Keeper) RemoveNocIcaCertificate(ctx sdk.Context, subject, subjectKeyID string, vid int32) {
	k._removeNocIcaCertificates(ctx, vid, func(cert *types.Certificate) bool {
		return cert.Subject == subject && cert.SubjectKeyId == subjectKeyID
	})
}

func (k Keeper) RemoveNocIcaCertificateBySerialNumber(ctx sdk.Context, vid int32, subject, subjectKeyID, serialNumber string) {
	k._removeNocIcaCertificates(ctx, vid, func(cert *types.Certificate) bool {
		return cert.Subject == subject && cert.SubjectKeyId == subjectKeyID && cert.SerialNumber == serialNumber
	})
}

func (k Keeper) _removeNocIcaCertificates(ctx sdk.Context, vid int32, filter func(cert *types.Certificate) bool) {
	certs, found := k.GetNocIcaCertificates(ctx, vid)
	if !found {
		return
	}

	numCertsBefore := len(certs.Certs)
	for i := 0; i < len(certs.Certs); {
		cert := certs.Certs[i]
		if filter(cert) {
			certs.Certs = append(certs.Certs[:i], certs.Certs[i+1:]...)
		} else {
			i++
		}
	}

	if len(certs.Certs) == 0 {
		k.RemoveNocIcaCertificates(ctx, vid)
	} else if numCertsBefore > len(certs.Certs) { // Update state only if any certificate is removed
		k.SetNocIcaCertificates(ctx, certs)
	}
}

// GetAllNocIcaCertificates returns all NocIcaCertificates.
func (k Keeper) GetAllNocIcaCertificates(ctx sdk.Context) (list []types.NocIcaCertificates) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.NocIcaCertificatesKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.NocIcaCertificates
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
