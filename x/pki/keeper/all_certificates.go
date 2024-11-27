package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/x509"
)

// SetAllCertificates set a specific certificates in the store from its index
func (k Keeper) SetAllCertificates(ctx sdk.Context, certificates types.AllCertificates) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.AllCertificatesKeyPrefix))
	b := k.cdc.MustMarshal(&certificates)
	store.Set(types.AllCertificatesKey(
		certificates.Subject,
		certificates.SubjectKeyId,
	), b)
}

// AddAllCertificate add a certificate to the list of all certificates for the subject/subjectKeyId map.
func (k Keeper) AddAllCertificate(ctx sdk.Context, certificate types.Certificate) {
	k._addAllCertificates(ctx, certificate.Subject, certificate.SubjectKeyId, certificate.SchemaVersion, []*types.Certificate{&certificate})
}

// AddAllCertificates add list of certificates in the store from its index
func (k Keeper) AddAllCertificates(ctx sdk.Context, subject string, subjectKeyID string, schemaVersion uint32, certs []*types.Certificate) {
	k._addAllCertificates(ctx, subject, subjectKeyID, schemaVersion, certs)
}

func (k Keeper) _addAllCertificates(ctx sdk.Context, subject string, subjectKeyID string, schemaVersion uint32, certs []*types.Certificate) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.AllCertificatesKeyPrefix))

	certificatesBytes := store.Get(types.AllCertificatesKey(
		subject,
		subjectKeyID,
	))
	var certificates types.AllCertificates

	if certificatesBytes == nil {
		certificates = types.AllCertificates{
			Subject:       subject,
			SubjectKeyId:  subjectKeyID,
			Certs:         []*types.Certificate{},
			SchemaVersion: schemaVersion,
		}
	} else {
		k.cdc.MustUnmarshal(certificatesBytes, &certificates)
	}

	certificates.Certs = append(certificates.Certs, certs...)

	k.SetAllCertificates(ctx, certificates)
}

// GetAllCertificates returns a certificates from its index
func (k Keeper) GetAllCertificates(
	ctx sdk.Context,
	subject string,
	subjectKeyID string,

) (val types.AllCertificates, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.AllCertificatesKeyPrefix))

	b := store.Get(types.AllCertificatesKey(
		subject,
		subjectKeyID,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)

	return val, true
}

// GetAllAllCertificates returns all certificates
func (k Keeper) GetAllAllCertificates(ctx sdk.Context) (list []types.AllCertificates) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.AllCertificatesKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.AllCertificates
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// RemoveAllCertificates removes a certificates from the store
func (k Keeper) RemoveAllCertificates(
	ctx sdk.Context,
	subject string,
	subjectKeyID string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.AllCertificatesKeyPrefix))
	store.Delete(types.AllCertificatesKey(
		subject,
		subjectKeyID,
	))
}

func (k Keeper) RemoveAllCertificatesBySerialNumber(ctx sdk.Context, subject, subjectKeyID, serialNumber string) {
	k._removeAllCertificatesBySerialNumber(ctx, subject, subjectKeyID, func(cert *types.Certificate) bool {
		return cert.Subject == subject && cert.SubjectKeyId == subjectKeyID && cert.SerialNumber == serialNumber
	})
}

func (k Keeper) _removeAllCertificatesBySerialNumber(ctx sdk.Context, subject string, subjectKeyID string, filter func(cert *types.Certificate) bool) {
	certs, found := k.GetAllCertificates(ctx, subject, subjectKeyID)
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
		store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.AllCertificatesKeyPrefix))
		store.Delete(types.AllCertificatesKey(
			subject,
			subjectKeyID,
		))
	} else if numCertsBefore > len(certs.Certs) { // Update state only if any certificate is removed
		k.SetAllCertificates(ctx, certs)
	}
}

// Tries to build a valid certificate chain for the given certificate.
// Returns the RootSubject/RootSubjectKeyID combination or an error in case no valid certificate chain can be built.
func (k Keeper) verifyCertificate(ctx sdk.Context,
	x509Certificate *x509.Certificate,
) (*x509.Certificate, error) {
	//nolint:nestif
	if x509Certificate.IsSelfSigned() {
		// in this system a certificate is self-signed if and only if it is a root certificate
		if err := x509Certificate.Verify(x509Certificate, ctx.BlockTime()); err == nil {
			return x509Certificate, nil
		}
	} else {
		parentCertificates, found := k.GetAllCertificates(ctx, x509Certificate.Issuer, x509Certificate.AuthorityKeyID)
		if !found {
			return nil, pkitypes.NewErrRootCertificateDoesNotExist(x509Certificate.Issuer, x509Certificate.AuthorityKeyID)
		}

		for _, cert := range parentCertificates.Certs {
			parentX509Certificate, err := x509.DecodeX509Certificate(cert.PemCert)
			if err != nil {
				continue
			}

			// verify certificate against parent
			if err := x509Certificate.Verify(parentX509Certificate, ctx.BlockTime()); err != nil {
				continue
			}

			// verify parent certificate
			if rootCertificate, err := k.verifyCertificate(ctx, parentX509Certificate); err == nil {
				return rootCertificate, nil
			}
		}
	}

	return nil, pkitypes.NewErrInvalidCertificate(
		fmt.Sprintf("Certificate verification failed for certificate with subject=%v and subjectKeyID=%v",
			x509Certificate.Subject, x509Certificate.SubjectKeyID))
}

// IsAllCertificatePresent Check if the All Certificate is present in the store.
func (k Keeper) IsAllCertificatePresent(
	ctx sdk.Context,
	subject string,
	subjectKeyID string,
) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.AllCertificatesKeyPrefix))

	return store.Has(types.AllCertificatesKey(
		subject,
		subjectKeyID,
	))
}
