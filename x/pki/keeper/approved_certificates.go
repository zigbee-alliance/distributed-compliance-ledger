package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/x509"
)

// SetApprovedCertificates set a specific approvedCertificates in the store from its index
func (k Keeper) SetApprovedCertificates(ctx sdk.Context, approvedCertificates types.ApprovedCertificates) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ApprovedCertificatesKeyPrefix))
	b := k.cdc.MustMarshal(&approvedCertificates)
	store.Set(types.ApprovedCertificatesKey(
		approvedCertificates.Subject,
		approvedCertificates.SubjectKeyId,
	), b)
}

// GetApprovedCertificates returns a approvedCertificates from its index
func (k Keeper) GetApprovedCertificates(
	ctx sdk.Context,
	subject string,
	subjectKeyId string,

) (val types.ApprovedCertificates, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ApprovedCertificatesKeyPrefix))

	b := store.Get(types.ApprovedCertificatesKey(
		subject,
		subjectKeyId,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveApprovedCertificates removes a approvedCertificates from the store
func (k Keeper) RemoveApprovedCertificates(
	ctx sdk.Context,
	subject string,
	subjectKeyId string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ApprovedCertificatesKeyPrefix))
	store.Delete(types.ApprovedCertificatesKey(
		subject,
		subjectKeyId,
	))
}

// GetAllApprovedCertificates returns all approvedCertificates
func (k Keeper) GetAllApprovedCertificates(ctx sdk.Context) (list []types.ApprovedCertificates) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ApprovedCertificatesKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ApprovedCertificates
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// SetApprovedCertificates set a specific approvedCertificates in the store from its index
func (k Keeper) AddApprovedCertificate(ctx sdk.Context, approvedCertificate types.Certificate) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ApprovedCertificatesKeyPrefix))

	approvedCertificatesBytes := store.Get(types.ApprovedCertificatesKey(
		approvedCertificate.Subject,
		approvedCertificate.SubjectKeyId,
	))
	var approvedCertificates types.ApprovedCertificates

	if approvedCertificatesBytes == nil {
		approvedCertificates = types.ApprovedCertificates{
			Subject:      approvedCertificate.Subject,
			SubjectKeyId: approvedCertificate.SubjectKeyId,
			Certs:        []*types.Certificate{},
		}
	} else {
		k.cdc.MustUnmarshal(approvedCertificatesBytes, &approvedCertificates)
	}

	approvedCertificates.Certs = append(approvedCertificates.Certs, &approvedCertificate)

	b := k.cdc.MustMarshal(&approvedCertificates)
	store.Set(types.ApprovedCertificatesKey(
		approvedCertificates.Subject,
		approvedCertificates.SubjectKeyId,
	), b)
}

// Tries to build a valid certificate chain for the given certificate.
// Returns the RootSubject/RootSubjectKeyID combination or an error in case no valid certificate chain can be built.
func (k Keeper) verifyCertificate(ctx sdk.Context,
	x509Certificate *x509.X509Certificate) (string, string, error) {
	// nolint:nestif
	if x509Certificate.IsSelfSigned() {
		// in this system a certificate is self-signed if and only if it is a root certificate
		if err := x509Certificate.Verify(x509Certificate); err == nil {
			return x509Certificate.Subject, x509Certificate.SubjectKeyID, nil
		}
	} else {
		parentCertificates, found := k.GetApprovedCertificates(ctx, x509Certificate.Issuer, x509Certificate.AuthorityKeyID)
		if !found {
			return "", "", types.NewErrCodeInvalidCertificate(
				fmt.Sprintf("Certificate verification failed for certificate with subject=%v and subjectKeyID=%v",
					x509Certificate.Subject, x509Certificate.SubjectKeyID))
		}

		for _, cert := range parentCertificates.Certs {
			parentX509Certificate, err := x509.DecodeX509Certificate(cert.PemCert)
			if err != nil {
				continue
			}

			// verify certificate against parent
			if err := x509Certificate.Verify(parentX509Certificate); err != nil {
				continue
			}

			// verify parent certificate
			if subject, subjectKeyID, err := k.verifyCertificate(ctx, parentX509Certificate); err == nil {
				return subject, subjectKeyID, nil
			}
		}
	}

	return "", "", types.NewErrCodeInvalidCertificate(
		fmt.Sprintf("Certificate verification failed for certificate with subject=%v and subjectKeyID=%v",
			x509Certificate.Subject, x509Certificate.SubjectKeyID))
}
