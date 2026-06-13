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

// matterVVSCMaxPathLength caps VVSC chain depth at the Matter R1.6 §6.4.10
// step 12.a.iii bound ("The path length SHALL NOT be longer than 3"): the
// trust anchor (self-signed VVSC) plus at most two non-self-issued VVSCs
// below it.
const matterVVSCMaxPathLength = 3

// verifyVVSCCertificate is the Matter R1.6 §6.4.10 step 12.a.iii equivalent
// of verifyCertificate, scoped to the VVSC chain semantics:
//
//   - Chain walks up via (Issuer, AuthorityKeyID) like verifyCertificate.
//   - Each parent step is validated with VerifyVVSCSignature, which bypasses
//     crypto/x509's BasicConstraints cA=TRUE + KeyUsage keyCertSign checks
//     (a §6.5.12 VVSC has cA=FALSE / KU=digitalSignature, so a VVSC parent
//     would never pass stdlib CheckSignatureFrom).
//   - Only stored entries with CertificateType_VIDSignerPKI are considered as
//     parents; "Only the self-signed certificates among the set SHALL be
//     considered as trust anchors during certificate path validation" — the
//     "set" being the Operational Trust Anchors entries under the VID, all of
//     which are VIDSignerPKI by construction.
//   - Total chain length (this cert + ancestors) is bounded by 3.
func (k Keeper) verifyVVSCCertificate(ctx sdk.Context,
	x509Certificate *x509.Certificate,
	depth int,
) (*x509.Certificate, error) {
	if depth > matterVVSCMaxPathLength {
		return nil, pkitypes.NewErrInvalidCertificate(
			fmt.Sprintf("VVSC chain length exceeds Matter R1.6 §6.4.10 limit of %d", matterVVSCMaxPathLength))
	}

	if x509Certificate.IsSelfSigned() {
		if err := x509Certificate.VerifyVVSCSignature(x509Certificate, ctx.BlockTime()); err == nil {
			return x509Certificate, nil
		}

		return nil, pkitypes.NewErrInvalidCertificate(
			fmt.Sprintf("VVSC trust anchor self-signature verification failed for subject=%v subjectKeyID=%v",
				x509Certificate.Subject, x509Certificate.SubjectKeyID))
	}

	parentCertificates, found := k.GetAllCertificates(ctx, x509Certificate.Issuer, x509Certificate.AuthorityKeyID)
	if !found {
		return nil, pkitypes.NewErrRootCertificateDoesNotExist(x509Certificate.Issuer, x509Certificate.AuthorityKeyID)
	}

	for _, cert := range parentCertificates.Certs {
		if cert.CertificateType != types.CertificateType_VIDSignerPKI {
			continue
		}
		parentX509Certificate, err := x509.DecodeX509Certificate(cert.PemCert)
		if err != nil {
			continue
		}
		if err := x509Certificate.VerifyVVSCSignature(parentX509Certificate, ctx.BlockTime()); err != nil {
			continue
		}
		if rootCertificate, err := k.verifyVVSCCertificate(ctx, parentX509Certificate, depth+1); err == nil {
			return rootCertificate, nil
		}
	}

	return nil, pkitypes.NewErrInvalidCertificate(
		fmt.Sprintf("VVSC chain verification failed for subject=%v subjectKeyID=%v",
			x509Certificate.Subject, x509Certificate.SubjectKeyID))
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
