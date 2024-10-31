package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

// SetNocCertificatesBySubjectKeyId set a specific nocCertificatesBySubjectKeyId in the store from its index
func (k Keeper) SetNocCertificatesBySubjectKeyId(ctx sdk.Context, nocCertificatesBySubjectKeyId types.NocCertificatesBySubjectKeyId) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.NocCertificatesBySubjectKeyIdKeyPrefix))
	b := k.cdc.MustMarshal(&nocCertificatesBySubjectKeyId)
	store.Set(types.NocCertificatesBySubjectKeyIdKey(
		nocCertificatesBySubjectKeyId.SubjectKeyId,
	), b)
}

// Add a noc certificate to the list of noc certificates with the subjectKeyId map.
func (k Keeper) AddNocCertificateBySubjectKeyID(ctx sdk.Context, certificate types.Certificate) {
	k._addNocCertificates(ctx, certificate.SubjectKeyId, []*types.Certificate{&certificate})
}

// Add a noc certificates list to noc certificates with the subjectKeyId map.
func (k Keeper) AddNocCertificatesBySubjectKeyID(ctx sdk.Context, nocCertificate types.NocCertificates) {
	k._addNocCertificates(ctx, nocCertificate.SubjectKeyId, nocCertificate.Certs)
}

// GetNocCertificatesBySubjectKeyId returns a nocCertificatesBySubjectKeyId from its index
func (k Keeper) GetNocCertificatesBySubjectKeyId(
	ctx sdk.Context,
	subjectKeyId string,

) (val types.NocCertificatesBySubjectKeyId, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.NocCertificatesBySubjectKeyIdKeyPrefix))

	b := store.Get(types.NocCertificatesBySubjectKeyIdKey(
		subjectKeyId,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// GetAllNocCertificatesBySubjectKeyId returns all nocCertificatesBySubjectKeyId
func (k Keeper) GetAllNocCertificatesBySubjectKeyId(ctx sdk.Context) (list []types.NocCertificatesBySubjectKeyId) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.NocCertificatesBySubjectKeyIdKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.NocCertificatesBySubjectKeyId
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// RemoveNocCertificatesBySubjectKeyID removes a nocCertificatesBySubjectKeyId from the store.
func (k Keeper) RemoveNocCertificatesBySubjectKeyID(
	ctx sdk.Context,
	subject string,
	subjectKeyID string,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.NocCertificatesBySubjectKeyIdKeyPrefix))
	certs, found := k.GetNocCertificatesBySubjectKeyId(ctx, subjectKeyID)
	if !found {
		return
	}

	for i := 0; i < len(certs.Certs); {
		if certs.Certs[i].Subject == subject {
			certs.Certs = append(certs.Certs[:i], certs.Certs[i+1:]...)
		} else {
			i++
		}
	}

	if len(certs.Certs) == 0 {
		store.Delete(types.NocCertificatesBySubjectKeyIdKey(
			subjectKeyID,
		))
	} else {
		k.SetNocCertificatesBySubjectKeyId(ctx, certs)
	}
}

// RemoveNocCertificatesBySubjectKeyId removes a nocCertificatesBySubjectKeyId from the store
func (k Keeper) RemoveNocCertificatesBySubjectKeyId(
	ctx sdk.Context,
	subjectKeyId string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.NocCertificatesBySubjectKeyIdKeyPrefix))
	store.Delete(types.NocCertificatesBySubjectKeyIdKey(
		subjectKeyId,
	))
}

func (k Keeper) RemoveNocCertificatesBySubjectKeyIdBySerialNumber(ctx sdk.Context, subject, subjectKeyID, serialNumber string) {
	k._removeNocCertificatesBySubjectKeyIDBySerialNumber(ctx, subjectKeyID, func(cert *types.Certificate) bool {
		return cert.Subject == subject && cert.SubjectKeyId == subjectKeyID && cert.SerialNumber == serialNumber
	})
}

func (k Keeper) _addNocCertificates(ctx sdk.Context, subjectKeyID string, certs []*types.Certificate) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.NocCertificatesBySubjectKeyIdKeyPrefix))

	nocCertificatesBytes := store.Get(types.NocCertificatesBySubjectKey(
		subjectKeyID,
	))
	var nocCertificates types.NocCertificatesBySubjectKeyId

	if nocCertificatesBytes == nil {
		nocCertificates = types.NocCertificatesBySubjectKeyId{
			SubjectKeyId: subjectKeyID,
			Certs:        []*types.Certificate{},
		}
	} else {
		k.cdc.MustUnmarshal(nocCertificatesBytes, &nocCertificates)
	}

	nocCertificates.Certs = append(nocCertificates.Certs, certs...)

	k.SetNocCertificatesBySubjectKeyId(ctx, nocCertificates)
}

func (k Keeper) _removeNocCertificatesBySubjectKeyIDBySerialNumber(ctx sdk.Context, subjectKeyID string, filter func(cert *types.Certificate) bool) {
	certs, found := k.GetNocCertificatesBySubjectKeyId(ctx, subjectKeyID)
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
		store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.NocCertificatesBySubjectKeyIdKeyPrefix))
		store.Delete(types.NocCertificatesBySubjectKeyIdKey(
			subjectKeyID,
		))
	} else if numCertsBefore > len(certs.Certs) { // Update state only if any certificate is removed
		k.SetNocCertificatesBySubjectKeyId(ctx, certs)
	}
}
