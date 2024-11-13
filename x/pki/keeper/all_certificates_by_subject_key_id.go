package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

// SetAllCertificatesBySubjectKeyID set a specific AllCertificatesBySubjectKeyId in the store from its index.
func (k Keeper) SetAllCertificatesBySubjectKeyID(ctx sdk.Context, allCertificatesBySubjectKeyID types.AllCertificatesBySubjectKeyId) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.AllCertificatesBySubjectKeyIDKeyPrefix))
	b := k.cdc.MustMarshal(&allCertificatesBySubjectKeyID)
	store.Set(types.AllCertificatesBySubjectKeyIDKey(
		allCertificatesBySubjectKeyID.SubjectKeyId,
	), b)
}

// Add an All certificate to the list of All certificates with the subjectKeyId map.
func (k Keeper) AddAllCertificateBySubjectKeyID(ctx sdk.Context, certificate types.Certificate) {
	k.addAllCertificates(ctx, certificate.SubjectKeyId, []*types.Certificate{&certificate})
}

// Add an All certificates list to All certificates with the subjectKeyId map.
func (k Keeper) AddAllCertificatesBySubjectKeyID(ctx sdk.Context, allCertificate types.AllCertificates) {
	k.addAllCertificates(ctx, allCertificate.SubjectKeyId, allCertificate.Certs)
}

func (k Keeper) addAllCertificates(ctx sdk.Context, subjectKeyID string, certs []*types.Certificate) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.AllCertificatesBySubjectKeyIDKeyPrefix))

	AllCertificatesBytes := store.Get(types.AllCertificatesBySubjectKey(
		subjectKeyID,
	))
	var AllCertificates types.AllCertificatesBySubjectKeyId

	if AllCertificatesBytes == nil {
		AllCertificates = types.AllCertificatesBySubjectKeyId{
			SubjectKeyId: subjectKeyID,
			Certs:        []*types.Certificate{},
		}
	} else {
		k.cdc.MustUnmarshal(AllCertificatesBytes, &AllCertificates)
	}

	AllCertificates.Certs = append(AllCertificates.Certs, certs...)

	k.SetAllCertificatesBySubjectKeyID(ctx, AllCertificates)
}

// GetAllCertificatesBySubjectKeyID returns a AllCertificatesBySubjectKeyId from its index.
func (k Keeper) GetAllCertificatesBySubjectKeyID(
	ctx sdk.Context,
	subjectKeyID string,

) (val types.AllCertificatesBySubjectKeyId, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.AllCertificatesBySubjectKeyIDKeyPrefix))

	b := store.Get(types.AllCertificatesBySubjectKeyIDKey(
		subjectKeyID,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)

	return val, true
}

// RemoveAllCertificatesBySubjectKeyID removes a AllCertificatesBySubjectKeyId from the store.
func (k Keeper) RemoveAllCertificatesBySubjectKeyID(
	ctx sdk.Context,
	subject string,
	subjectKeyID string,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.AllCertificatesBySubjectKeyIDKeyPrefix))
	certs, found := k.GetAllCertificatesBySubjectKeyID(ctx, subjectKeyID)
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
		store.Delete(types.AllCertificatesBySubjectKeyIDKey(
			subjectKeyID,
		))
	} else {
		k.SetAllCertificatesBySubjectKeyID(ctx, certs)
	}
}

func (k Keeper) RemoveAllCertificatesBySubjectKeyIDBySerialNumber(ctx sdk.Context, subject, subjectKeyID, serialNumber string) {
	k._removeAllCertificatesFromSubjectKeyIDState(ctx, subjectKeyID, func(cert *types.Certificate) bool {
		return cert.Subject == subject && cert.SubjectKeyId == subjectKeyID && cert.SerialNumber == serialNumber
	})
}

// GetAllAllCertificatesBySubjectKeyID returns all AllCertificatesBySubjectKeyId.
func (k Keeper) GetAllAllCertificatesBySubjectKeyID(ctx sdk.Context) (list []types.AllCertificatesBySubjectKeyId) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.AllCertificatesBySubjectKeyIDKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.AllCertificatesBySubjectKeyId
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

func (k Keeper) _removeAllCertificatesFromSubjectKeyIDState(ctx sdk.Context, subjectKeyID string, filter func(cert *types.Certificate) bool) {
	certs, found := k.GetAllCertificatesBySubjectKeyID(ctx, subjectKeyID)
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
		store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.AllCertificatesBySubjectKeyIDKeyPrefix))
		store.Delete(types.AllCertificatesBySubjectKeyIDKey(
			subjectKeyID,
		))
	} else if numCertsBefore > len(certs.Certs) { // Update state only if any certificate is removed
		k.SetAllCertificatesBySubjectKeyID(ctx, certs)
	}
}
