package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

// SetAllCertificatesBySubject set a specific allCertificatesBySubject in the store from its index
func (k Keeper) SetAllCertificatesBySubject(ctx sdk.Context, allCertificatesBySubject types.AllCertificatesBySubject) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.AllCertificatesBySubjectKeyPrefix))
	b := k.cdc.MustMarshal(&allCertificatesBySubject)
	store.Set(types.AllCertificatesBySubjectKey(
		allCertificatesBySubject.Subject,
	), b)
}

// Add AllCertificates to a subject->subjectKeyId index.
func (k Keeper) AddAllCertificateBySubject(ctx sdk.Context, subject string, subjectKeyID string) {
	AllCertificatesBySubject, _ := k.GetAllCertificatesBySubject(ctx, subject)

	// Check if cert is already there
	for _, existingID := range AllCertificatesBySubject.SubjectKeyIds {
		if existingID == subjectKeyID {
			return
		}
	}

	AllCertificatesBySubject.Subject = subject
	AllCertificatesBySubject.SubjectKeyIds = append(AllCertificatesBySubject.SubjectKeyIds, subjectKeyID)

	k.SetAllCertificatesBySubject(ctx, AllCertificatesBySubject)
}

// AddAllCertificates add list of certificates in the store from its index
func (k Keeper) AddAllCertificatesBySubject(ctx sdk.Context, subject string, schemaVersion uint32, subjectKeyIds []string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.AllCertificatesBySubjectKeyPrefix))

	certificatesBytes := store.Get(types.AllCertificatesBySubjectKey(
		subject,
	))
	var certificates types.AllCertificatesBySubject

	if certificatesBytes == nil {
		certificates = types.AllCertificatesBySubject{
			Subject:       subject,
			SubjectKeyIds: []string{},
			SchemaVersion: schemaVersion,
		}
	} else {
		k.cdc.MustUnmarshal(certificatesBytes, &certificates)
	}

	certificates.SubjectKeyIds = append(certificates.SubjectKeyIds, subjectKeyIds...)

	k.SetAllCertificatesBySubject(ctx, certificates)
}

// GetAllCertificatesBySubject returns a allCertificatesBySubject from its index
func (k Keeper) GetAllCertificatesBySubject(
	ctx sdk.Context,
	subject string,

) (val types.AllCertificatesBySubject, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.AllCertificatesBySubjectKeyPrefix))

	b := store.Get(types.AllCertificatesBySubjectKey(
		subject,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)

	return val, true
}

// RemoveAllCertificateBySubject removes a allCertificatesBySubject from the store
func (k Keeper) RemoveAllCertificateBySubject(ctx sdk.Context, subject string, subjectKeyID string) {
	AllCertificatesBySubject, _ := k.GetAllCertificatesBySubject(ctx, subject)

	certIDIndex := -1
	for i, existingIdentifier := range AllCertificatesBySubject.SubjectKeyIds {
		if existingIdentifier == subjectKeyID {
			certIDIndex = i

			break
		}
	}
	if certIDIndex == -1 {
		return
	}

	AllCertificatesBySubject.SubjectKeyIds = append(AllCertificatesBySubject.SubjectKeyIds[:certIDIndex], AllCertificatesBySubject.SubjectKeyIds[certIDIndex+1:]...)

	if len(AllCertificatesBySubject.SubjectKeyIds) > 0 {
		k.SetAllCertificatesBySubject(ctx, AllCertificatesBySubject)
	} else {
		k.RemoveAllCertificatesBySubject(ctx, subject)
	}
}

// RemoveAllCertificatesBySubject removes a AllCertificatesBySubject from the store.
func (k Keeper) RemoveAllCertificatesBySubject(
	ctx sdk.Context,
	subject string,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.AllCertificatesBySubjectKeyPrefix))
	store.Delete(types.AllCertificatesBySubjectKey(
		subject,
	))
}

// GetAllAllCertificatesBySubject returns all allCertificatesBySubject
func (k Keeper) GetAllAllCertificatesBySubject(ctx sdk.Context) (list []types.AllCertificatesBySubject) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.AllCertificatesBySubjectKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.AllCertificatesBySubject
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// IsCertificatesBySubjectPresent Check if the Certificate By Subject is present in the store.
func (k Keeper) IsCertificatesBySubjectPresent(
	ctx sdk.Context,
	subject string,
) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.AllCertificatesBySubjectKeyPrefix))

	return store.Has(types.AllCertificatesBySubjectKey(
		subject,
	))
}
