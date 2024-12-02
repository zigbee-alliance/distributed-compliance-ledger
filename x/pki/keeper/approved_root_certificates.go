package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

// SetApprovedRootCertificates set approvedRootCertificates in the store.
func (k Keeper) SetApprovedRootCertificates(ctx sdk.Context, approvedRootCertificates types.ApprovedRootCertificates) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(pkitypes.ApprovedRootCertificatesKeyPrefix))
	b := k.cdc.MustMarshal(&approvedRootCertificates)
	store.Set(pkitypes.ApprovedRootCertificatesKey, b)
}

// GetApprovedRootCertificates returns approvedRootCertificates.
func (k Keeper) GetApprovedRootCertificates(ctx sdk.Context) (val types.ApprovedRootCertificates, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(pkitypes.ApprovedRootCertificatesKeyPrefix))

	b := store.Get(pkitypes.ApprovedRootCertificatesKey)
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)

	return val, true
}

// RemoveApprovedRootCertificates removes approvedRootCertificates from the store.
func (k Keeper) RemoveApprovedRootCertificates(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(pkitypes.ApprovedRootCertificatesKeyPrefix))
	store.Delete(pkitypes.ApprovedRootCertificatesKey)
}

// Add root certificate to the list.
func (k Keeper) AddApprovedRootCertificate(ctx sdk.Context, certificate types.Certificate) {
	rootCertificates, _ := k.GetApprovedRootCertificates(ctx)

	certID := types.CertificateIdentifier{
		Subject:      certificate.Subject,
		SubjectKeyId: certificate.SubjectKeyId,
	}

	// Check if the root cert is already there
	for _, existingCertID := range rootCertificates.Certs {
		if *existingCertID == certID {
			return
		}
	}

	rootCertificates.Certs = append(rootCertificates.Certs, &certID)

	k.SetApprovedRootCertificates(ctx, rootCertificates)
}

// Remove root certificate from the list.
func (k Keeper) RemoveApprovedRootCertificate(
	ctx sdk.Context,
	subject string,
	subjectKeyID string,
) {
	certID := types.CertificateIdentifier{
		Subject:      subject,
		SubjectKeyId: subjectKeyID,
	}

	rootCertificates, _ := k.GetApprovedRootCertificates(ctx)

	certIDIndex := -1

	for i, existingIdentifier := range rootCertificates.Certs {
		if *existingIdentifier == certID {
			certIDIndex = i

			break
		}
	}
	if certIDIndex == -1 {
		return
	}

	rootCertificates.Certs = append(rootCertificates.Certs[:certIDIndex], rootCertificates.Certs[certIDIndex+1:]...)

	k.SetApprovedRootCertificates(ctx, rootCertificates)
}
