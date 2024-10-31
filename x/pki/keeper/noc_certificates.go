package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

// SetNocCertificates set a specific nocCertificates in the store from its index
func (k Keeper) SetNocCertificates(ctx sdk.Context, nocCertificates types.NocCertificates) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.NocCertificatesKeyPrefix))
	b := k.cdc.MustMarshal(&nocCertificates)
	store.Set(types.NocCertificatesKey(
		nocCertificates.Subject,
		nocCertificates.SubjectKeyId,
	), b)
}

// Add a NOC certificate to the list of NOC certificates.
func (k Keeper) AddNocCertificate(ctx sdk.Context, nocCertificate types.Certificate) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.NocCertificatesKeyPrefix))

	nocCertificatesKeyBytes := store.Get(types.NocCertificatesKey(nocCertificate.Subject, nocCertificate.SubjectKeyId))
	var nocCertificates types.NocCertificates

	if nocCertificatesKeyBytes == nil {
		nocCertificates = types.NocCertificates{
			Subject:       nocCertificate.Subject,
			SubjectKeyId:  nocCertificate.SubjectKeyId,
			Certs:         []*types.Certificate{},
			Tq:            1,
			SchemaVersion: nocCertificate.SchemaVersion,
		}
	} else {
		k.cdc.MustUnmarshal(nocCertificatesKeyBytes, &nocCertificates)
	}

	nocCertificates.Certs = append(nocCertificates.Certs, &nocCertificate)

	b := k.cdc.MustMarshal(&nocCertificates)
	store.Set(types.NocCertificatesKey(nocCertificate.Subject, nocCertificate.SubjectKeyId), b)
}

// GetNocCertificates returns a nocCertificates from its index
func (k Keeper) GetNocCertificates(
	ctx sdk.Context,
	subject string,
	subjectKeyId string,

) (val types.NocCertificates, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.NocCertificatesKeyPrefix))

	b := store.Get(types.NocCertificatesKey(
		subject,
		subjectKeyId,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// GetAllNocCertificates returns all nocCertificates
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

// RemoveNocCertificates removes a nocCertificates from the store
func (k Keeper) RemoveNocCertificates(
	ctx sdk.Context,
	subject string,
	subjectKeyId string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.NocCertificatesKeyPrefix))
	store.Delete(types.NocCertificatesKey(
		subject,
		subjectKeyId,
	))
}

func (k Keeper) RemoveNocCertificatesBySerialNumber(ctx sdk.Context, subject, subjectKeyID, serialNumber string) {
	k._removeNocCertificatesBySerialNumber(ctx, subject, subjectKeyID, func(cert *types.Certificate) bool {
		return cert.Subject == subject && cert.SubjectKeyId == subjectKeyID && cert.SerialNumber == serialNumber
	})
}

func (k Keeper) _removeNocCertificatesBySerialNumber(ctx sdk.Context, subject string, subjectKeyID string, filter func(cert *types.Certificate) bool) {
	certs, found := k.GetNocCertificates(ctx, subject, subjectKeyID)
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
		store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.NocCertificatesKeyPrefix))
		store.Delete(types.AllCertificatesKey(
			subject,
			subjectKeyID,
		))
	} else if numCertsBefore > len(certs.Certs) { // Update state only if any certificate is removed
		k.SetNocCertificates(ctx, certs)
	}
}
