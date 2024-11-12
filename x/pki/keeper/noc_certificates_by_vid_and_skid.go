package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

// SetNocCertificatesByVidAndSkid set a specific nocCertificatesByVidAndSkid in the store from its index.
func (k Keeper) SetNocCertificatesByVidAndSkid(ctx sdk.Context, nocCertificatesByVidAndSkid types.NocCertificatesByVidAndSkid) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.NocCertificatesByVidAndSkidKeyPrefix))
	b := k.cdc.MustMarshal(&nocCertificatesByVidAndSkid)
	store.Set(types.NocCertificatesByVidAndSkidKey(
		nocCertificatesByVidAndSkid.Vid,
		nocCertificatesByVidAndSkid.SubjectKeyId,
	), b)
}

// GetNocCertificatesByVidAndSkid returns a nocCertificatesByVidAndSkid from its index.
func (k Keeper) GetNocCertificatesByVidAndSkid(
	ctx sdk.Context,
	vid int32,
	subjectKeyID string,

) (val types.NocCertificatesByVidAndSkid, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.NocCertificatesByVidAndSkidKeyPrefix))

	b := store.Get(types.NocCertificatesByVidAndSkidKey(
		vid,
		subjectKeyID,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)

	return val, true
}

// Add a NOC certificate to the list of NOC certificates for the VID map.
func (k Keeper) AddNocCertificateByVidAndSkid(ctx sdk.Context, nocCertificate types.Certificate) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.NocCertificatesByVidAndSkidKeyPrefix))

	nocCertificatesByVidAndSkidKeyBytes := store.Get(types.NocCertificatesByVidAndSkidKey(nocCertificate.Vid, nocCertificate.SubjectKeyId))
	var nocCertificatesByVidAndSkid types.NocCertificatesByVidAndSkid

	if nocCertificatesByVidAndSkidKeyBytes == nil {
		nocCertificatesByVidAndSkid = types.NocCertificatesByVidAndSkid{
			Vid:          nocCertificate.Vid,
			SubjectKeyId: nocCertificate.SubjectKeyId,
			Certs:        []*types.Certificate{},
			Tq:           1,
		}
	} else {
		k.cdc.MustUnmarshal(nocCertificatesByVidAndSkidKeyBytes, &nocCertificatesByVidAndSkid)
	}

	nocCertificatesByVidAndSkid.Certs = append(nocCertificatesByVidAndSkid.Certs, &nocCertificate)

	b := k.cdc.MustMarshal(&nocCertificatesByVidAndSkid)
	store.Set(types.NocCertificatesByVidAndSkidKey(nocCertificate.Vid, nocCertificate.SubjectKeyId), b)
}

// RemoveNocCertificatesByVidAndSkid removes a nocCertificatesByVidAndSkid from the store.
func (k Keeper) RemoveNocCertificatesByVidAndSkid(
	ctx sdk.Context,
	vid int32,
	subjectKeyID string,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.NocCertificatesByVidAndSkidKeyPrefix))
	store.Delete(types.NocCertificatesByVidAndSkidKey(
		vid,
		subjectKeyID,
	))
}

// RemoveNocCertificateByVidSkidSubjectAndSerialNumber removes certificate with specified subject from the list.
func (k Keeper) RemoveNocCertificateByVidSubjectAndSkid(
	ctx sdk.Context,
	vid int32,
	subject string,
	subjectKeyID string,
) {
	k._filterAndSetNocCertificateByVidAndSkid(
		ctx,
		vid,
		subjectKeyID,
		func(cert *types.Certificate) bool {
			return cert.Subject != subject
		},
	)
}

// RemoveNocCertificatesByVidAndSkidBySerialNumber removes certificate with specified subject and serial number from the list.
func (k Keeper) RemoveNocCertificatesByVidAndSkidBySerialNumber(
	ctx sdk.Context,
	vid int32,
	subject string,
	subjectKeyID string,
	serialNumber string,
) {
	k._filterAndSetNocCertificateByVidAndSkid(
		ctx,
		vid,
		subjectKeyID,
		func(cert *types.Certificate) bool {
			return !(cert.Subject == subject && cert.SubjectKeyId == subjectKeyID && cert.SerialNumber == serialNumber)
		},
	)
}

// RemoveNocCertificateByVidSkidSubjectAndSerialNumber removes certificate with specified subject and serial number from the list.
func (k Keeper) _filterAndSetNocCertificateByVidAndSkid(
	ctx sdk.Context,
	vid int32,
	subjectKeyID string,
	predicate CertificatePredicate,
) {
	nocCertificates, _ := k.GetNocCertificatesByVidAndSkid(ctx, vid, subjectKeyID)
	filteredCertificates := filterCertificates(&nocCertificates.Certs, predicate)

	if len(filteredCertificates) > 0 {
		nocCertificates.Certs = filteredCertificates
		k.SetNocCertificatesByVidAndSkid(ctx, nocCertificates)
	} else {
		k.RemoveNocCertificatesByVidAndSkid(ctx, vid, subjectKeyID)
	}
}

// GetAllNocCertificatesByVidAndSkid returns all nocCertificatesByVidAndSkid.
func (k Keeper) GetAllNocCertificatesByVidAndSkid(ctx sdk.Context) (list []types.NocCertificatesByVidAndSkid) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.NocCertificatesByVidAndSkidKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.NocCertificatesByVidAndSkid
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
