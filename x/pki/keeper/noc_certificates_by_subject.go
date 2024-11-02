package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

// SetNocCertificatesBySubject set a specific nocCertificatesBySubject in the store from its index
func (k Keeper) SetNocCertificatesBySubject(ctx sdk.Context, nocCertificatesBySubject types.NocCertificatesBySubject) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.NocCertificatesBySubjectKeyPrefix))
	b := k.cdc.MustMarshal(&nocCertificatesBySubject)
	store.Set(types.NocCertificatesBySubjectKey(
		nocCertificatesBySubject.Subject,
	), b)
}

// Add a NOC certificate to the list of NOC certificates: subject->subjectKeyId index.
func (k Keeper) AddNocCertificateBySubject(ctx sdk.Context, nocCertificate types.Certificate) {
	nocCertificatesBySubject, _ := k.GetNocCertificatesBySubject(ctx, nocCertificate.Subject)

	// Check if cert is already there
	for _, existingID := range nocCertificatesBySubject.SubjectKeyIds {
		if existingID == nocCertificate.SubjectKeyId {
			return
		}
	}

	nocCertificatesBySubject.Subject = nocCertificate.Subject
	nocCertificatesBySubject.SubjectKeyIds = append(nocCertificatesBySubject.SubjectKeyIds, nocCertificate.SubjectKeyId)

	k.SetNocCertificatesBySubject(ctx, nocCertificatesBySubject)
}

// GetNocCertificatesBySubject returns a nocCertificatesBySubject from its index
func (k Keeper) GetNocCertificatesBySubject(
	ctx sdk.Context,
	subject string,

) (val types.NocCertificatesBySubject, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.NocCertificatesBySubjectKeyPrefix))

	b := store.Get(types.NocCertificatesBySubjectKey(
		subject,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)

	return val, true
}

// GetAllNocCertificatesBySubject returns all nocCertificatesBySubject
func (k Keeper) GetAllNocCertificatesBySubject(ctx sdk.Context) (list []types.NocCertificatesBySubject) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.NocCertificatesBySubjectKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.NocCertificatesBySubject
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// Remove revoked root certificate from the list.
func (k Keeper) RemoveNocCertificateBySubject(ctx sdk.Context, subject string, subjectKeyID string) {
	nocCertificatesBySubject, _ := k.GetNocCertificatesBySubject(ctx, subject)

	certIDIndex := -1
	for i, existingIdentifier := range nocCertificatesBySubject.SubjectKeyIds {
		if existingIdentifier == subjectKeyID {
			certIDIndex = i

			break
		}
	}
	if certIDIndex == -1 {
		return
	}

	nocCertificatesBySubject.SubjectKeyIds = append(nocCertificatesBySubject.SubjectKeyIds[:certIDIndex], nocCertificatesBySubject.SubjectKeyIds[certIDIndex+1:]...)

	if len(nocCertificatesBySubject.SubjectKeyIds) > 0 {
		k.SetNocCertificatesBySubject(ctx, nocCertificatesBySubject)
	} else {
		k.RemoveNocCertificatesBySubject(ctx, subject)
	}
}

// RemoveNocCertificatesBySubject removes a nocCertificatesBySubject from the store
func (k Keeper) RemoveNocCertificatesBySubject(
	ctx sdk.Context,
	subject string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.NocCertificatesBySubjectKeyPrefix))
	store.Delete(types.NocCertificatesBySubjectKey(
		subject,
	))
}
