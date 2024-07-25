package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

// SetNocCertificatesByVidAndSkid set a specific nocRootCertificatesByVidAndSkid in the store from its index.
func (k Keeper) SetNocCertificatesByVidAndSkid(ctx sdk.Context, nocRootCertificatesByVidAndSkid types.NocCertificatesByVidAndSkid) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.NocCertificatesByVidAndSkidKeyPrefix))
	b := k.cdc.MustMarshal(&nocRootCertificatesByVidAndSkid)
	store.Set(types.NocCertificatesByVidAndSkidKey(
		nocRootCertificatesByVidAndSkid.Vid,
		nocRootCertificatesByVidAndSkid.SubjectKeyId,
	), b)
}

// GetNocCertificatesByVidAndSkid returns a nocRootCertificatesByVidAndSkid from its index.
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

	nocRootCertificatesByVidAndSkidKeyBytes := store.Get(types.NocCertificatesByVidAndSkidKey(nocCertificate.Vid, nocCertificate.SubjectKeyId))
	var nocRootCertificatesByVidAndSkid types.NocCertificatesByVidAndSkid

	if nocRootCertificatesByVidAndSkidKeyBytes == nil {
		nocRootCertificatesByVidAndSkid = types.NocCertificatesByVidAndSkid{
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
	store.Set(types.NocCertificatesByVidAndSkidKey(nocCertificate.Vid, nocCertificate.SubjectKeyId), b)
}

// RemoveNocCertificatesByVidAndSkid removes a nocRootCertificatesByVidAndSkid from the store.
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

// RemoveNocRootCertificateByVidSkidSubjectAndSerialNumber removes root certificate with specified subject from the list.
func (k Keeper) RemoveNocRootCertificateByVidSubjectAndSkid(
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

// RemoveNocRootCertificateByVidSkidSubjectAndSerialNumber removes root certificate with specified subject and serial number from the list.
func (k Keeper) RemoveNocRootCertificateByVidSubjectSkidAndSerialNumber(
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
			return !(cert.Subject == subject && cert.SerialNumber == serialNumber)
		},
	)
}

// RemoveNocRootCertificateByVidSkidSubjectAndSerialNumber removes root certificate with specified subject and serial number from the list.
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

// GetAllNocCertificatesByVidAndSkid returns all nocRootCertificatesByVidAndSkid.
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
