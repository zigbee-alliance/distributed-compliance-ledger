package pki

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	cliputils "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/cli/utils"
)

const (
	nocRevChildCert2CopyPath         = "../../constants/noc_cert_2_copy"
	nocRevChildCert2CopySerialNumber = "252687488758567844896720928536709119387931444024"

	// nocLeafCert2* are repurposed to point at vvsc_leaf_cert_2 because the
	// strict §6.5.12 ICA handler no longer accepts NOC end-entity profiles via
	// add-noc-x509-ica-cert. The cascade-revoke section keeps the same
	// variable names so its assertions still read naturally.
	nocLeafCert2Path         = "../../constants/vvsc_leaf_cert_2"
	nocLeafCert2Subject      = "MIGYMQswCQYDVQQGEwJVWjETMBEGA1UECAwKU29tZSBTdGF0ZTETMBEGA1UEBwwKU29tZSBTdGF0ZTEYMBYGA1UECgwPRXhhbXBsZSBDb21wYW55MRkwFwYDVQQLDBBUZXN0aW5nIERpdmlzaW9uMRQwEgYDVQQDDAtWVlNDLUxlYWYtMjEUMBIGCisGAQQBgqJ8AgEMBDAwMDE="
	nocLeafCert2SubjectKeyID = "8D:F6:2A:9C:24:D0:92:36:83:32:38:47:35:3A:0B:E9:19:CD:90:B3"
	nocLeafCert2SerialNumber = "5068329979159654450"

	// VvscRootCert1Copy reuses VvscRootCert1's key (same Subject + SKID), but
	// with a different serial number — used to re-establish an active VVSC root
	// after the original VvscRootCert1 has been revoked (the UniqueCertificate
	// record keyed by Issuer+SerialNumber survives revocation, so re-adding the
	// same serial fails with "certificate already exists").
	vvscRootCert1CopyPath         = "../../constants/vvsc_root_cert_1_copy"
	vvscRootCert1CopySerialNumber = "5068329979260121137"

	// Second VVSC intermediate (chained under VvscRoot1) — pre-seeded so
	// VvscLeafCert2's AuthorityKeyID resolves through verifyVVSCCertificate.
	vvscIcaCert2Path         = "../../constants/vvsc_ica_cert_2"
	vvscIcaCert2Subject      = "MIGXMQswCQYDVQQGEwJVWjETMBEGA1UECAwKU29tZSBTdGF0ZTETMBEGA1UEBwwKU29tZSBTdGF0ZTEYMBYGA1UECgwPRXhhbXBsZSBDb21wYW55MRkwFwYDVQQLDBBUZXN0aW5nIERpdmlzaW9uMRMwEQYDVQQDDApWVlNDLUlDQS0yMRQwEgYKKwYBBAGConwCAQwEMDAwMQ=="
	vvscIcaCert2SubjectKeyID = "ED:8C:5B:36:E7:3C:E4:54:09:A2:59:D4:E8:0A:D6:6C:99:C6:A2:CC"
	vvscIcaCert2SerialNumber = "5068329979109130546"
)

// noc_root_cert_1/copy and noc_cert_1/copy were added and revoked by TestPKINocCerts,
// so this test removes them from the revoked pool and re-adds them before revoking again.
func TestPKINocRevocationWithRevokingChild(t *testing.T) {
	vendorAccount := fmt.Sprintf("vendor_account_%d", nocVid)
	cliputils.CreateVendorAccount(t, vendorAccount, nocVid)

	t.Run("RevokeNocRootCertWithChildFlag", func(t *testing.T) {
		// noc_root_cert_1 and noc_root_cert_1_copy are in the revoked pool from TestPKINocCerts.
		// Remove them so they can be re-added.
		txResult, err := RemoveNocRootCert(nocRootCert1Subject, nocRootCert1SubjectKeyID, vendorAccount)
		cliputils.RequireTxOK(t, txResult, err)

		// noc_cert_1 and noc_cert_1_copy are in the revoked pool from TestPKINocCerts.
		// Remove them so they can be re-added.
		txResult, err = RemoveNocCert(nocCert1Subject, nocCert1SubjectKeyID, vendorAccount)
		cliputils.RequireTxOK(t, txResult, err)

		// Re-add root certs
		txResult, err = AddNocRootCert(nocRootCert1Path, vendorAccount)
		cliputils.RequireTxOK(t, txResult, err)

		txResult, err = AddNocRootCert(nocRootCert1CopyPath, vendorAccount)
		cliputils.RequireTxOK(t, txResult, err)

		// Re-add ICA cert (noc_cert_1)
		txResult, err = AddNocX509IcaCert(nocCert1Path, vendorAccount)
		cliputils.RequireTxOK(t, txResult, err)

		// VvscLeafCert1 is already active on-chain from TestPKINocCerts (chained
		// under VvscRoot1 → VvscIca1, was never revoked). Verify it exists.
		icas, err := GetNocX509IcaCerts(nocVid)
		require.NoError(t, err)
		require.NotNil(t, icas)
		require.True(t, containsCertSubjectSerial(icas.Certs, vvscLeafCert1Subject, vvscLeafCert1SerialNumber))

		// Verify root certs exist.
		allRoots, err := GetAllNocRootCerts()
		require.NoError(t, err)
		require.True(t, containsNocRootCertSerial(allRoots, nocRootCert1SerialNumber))
		require.True(t, containsNocRootCertSerial(allRoots, nocRootCert1CopySerialNumber))

		// Revoke the OperationalPKI root NOC certificate with revoke-child=true.
		// Cascade hits NocCert1 only — the VVSC chain is structurally disjoint (Matter §6.5.12).
		txResult, err = RevokeNocRootCert(nocRootCert1Subject, nocRootCert1SubjectKeyID, vendorAccount, RevokeNocCertOpts{RevokeChild: true})
		cliputils.RequireTxOK(t, txResult, err)

		// Also revoke the VVSC root with revoke-child=true so the cascade picks
		// up VvscIca1 + the VVSC leaf.
		txResult, err = RevokeNocRootCert(vvscRootCert1Subject, vvscRootCert1SubjectKeyID, vendorAccount, RevokeNocCertOpts{RevokeChild: true})
		cliputils.RequireTxOK(t, txResult, err)

		// Both NOC root certs + VVSC root should be revoked.
		revokedRoots, err := GetAllRevokedNocRootCerts()
		require.NoError(t, err)
		require.True(t, containsRevokedNocRootCertSerial(revokedRoots, nocRootCert1SerialNumber))
		require.True(t, containsRevokedNocRootCertSerial(revokedRoots, nocRootCert1CopySerialNumber))
		require.True(t, containsRevokedNocRootCertSerial(revokedRoots, vvscRootCert1SerialNumber))

		// Revoked ICA list now contains both the NOC ICA and the VVSC ICA + leaf.
		revokedIcas, err := GetAllRevokedNocX509IcaCerts()
		require.NoError(t, err)
		require.True(t, containsRevokedNocIcaCertSubject(revokedIcas, nocCert1Subject))
		require.True(t, containsRevokedNocIcaCertSubject(revokedIcas, vvscIcaCert1Subject))
		require.True(t, containsRevokedNocIcaCertSubject(revokedIcas, vvscLeafCert1Subject))

		// Root cert 1 (both serials) should no longer be in the approved root list for VID.
		// (Root cert 2 from TestPKINocCerts may still be approved.)
		roots, err := GetNocRootCerts(nocVid)
		require.NoError(t, err)
		if roots != nil {
			require.False(t, containsCertSerial(roots.Certs, nocRootCert1SerialNumber))
			require.False(t, containsCertSerial(roots.Certs, nocRootCert1CopySerialNumber))
			require.False(t, containsCertSerial(roots.Certs, vvscRootCert1SerialNumber))
		}

		// cert1 should be revoked after revoking root cert 1 with child flag.
		// (noc_cert_2 — child of root cert 2 from TestPKINocCerts — may still be approved.)
		icas, err = GetNocX509IcaCerts(nocVid)
		require.NoError(t, err)
		if icas != nil {
			require.False(t, containsCertSerial(icas.Certs, nocCert1SerialNumber))
			require.False(t, containsCertSerial(icas.Certs, vvscLeafCert1SerialNumber))
		}

		// NOC certs must not appear in the DA (all-x509-certs) list.
		allDa, err := GetAllX509Certs()
		require.NoError(t, err)
		require.False(t, containsApprovedCertSerial(allDa, nocRootCert1SerialNumber))
		require.False(t, containsApprovedCertSerial(allDa, nocCert1SerialNumber))
		require.False(t, containsApprovedCertSerial(allDa, vvscLeafCert1SerialNumber))

		// The single-record revoked NOC root query returns the revoked root.
		revokedRoot, err := GetRevokedNocRootCert(nocRootCert1Subject, nocRootCert1SubjectKeyID)
		require.NoError(t, err)
		require.NotNil(t, revokedRoot)
		require.Equal(t, nocRootCert1Subject, revokedRoot.Subject)

		// Namespace separation: a revoked NOC root must not leak into the DA
		// revoked-root list (all-revoked-x509-root-certs).
		daRevokedRoots, err := GetAllRevokedX509RootCerts()
		require.NoError(t, err)
		if daRevokedRoots != nil {
			for _, id := range daRevokedRoots.Certs {
				require.NotEqual(t, nocRootCert1Subject, id.Subject)
				require.NotEqual(t, vvscRootCert1Subject, id.Subject)
			}
		}
	})

	t.Run("RevokeNocIcaCertWithChildFlag", func(t *testing.T) {
		// noc_root_cert_2 and noc_cert_2 were already added by TestPKINocCerts — verify they are on-chain.
		roots, err := GetNocRootCerts(nocVid)
		require.NoError(t, err)
		require.NotNil(t, roots)
		require.True(t, containsCertSerial(roots.Certs, nocRootCert2SerialNumber))

		icas, err := GetNocX509IcaCerts(nocVid)
		require.NoError(t, err)
		require.NotNil(t, icas)
		require.True(t, containsCertSerial(icas.Certs, nocCert2SerialNumber))

		// Add cert2copy (OperationalPKI ICA copy) — not yet on-chain.
		txResult, err := AddNocX509IcaCert(nocRevChildCert2CopyPath, vendorAccount)
		cliputils.RequireTxOK(t, txResult, err)

		// Re-establish an active VVSC root via VvscRootCert1Copy. Revocation
		// soft-deletes the cert into the revoked list but the
		// (Issuer, SerialNumber) UniqueCertificate record survives — re-adding
		// the same PEM would fail with ErrCertificateAlreadyExists. The Copy
		// shares VvscRoot1's key (same Subject + SubjectKeyID) so VvscIca2's
		// AuthorityKeyID still resolves to a present VIDSignerPKI entry during
		// verifyVVSCCertificate's chain walk.
		txResult, err = AddNocRootCert(vvscRootCert1CopyPath, vendorAccount, AddNocCertOpts{IsVidVerificationSigner: true})
		cliputils.RequireTxOK(t, txResult, err)

		// Pre-seed the second VVSC intermediate (VvscIca2 under VvscRoot1) so
		// the leaf-2 chain VvscRoot1 → VvscIca2 → vvsc_leaf_cert_2 resolves
		// through verifyVVSCCertificate.
		txResult, err = AddNocX509IcaCert(vvscIcaCert2Path, vendorAccount, AddNocCertOpts{IsVidVerificationSigner: true})
		cliputils.RequireTxOK(t, txResult, err)

		// Add VVSC leaf certificate 2 (replaces the legacy NocLeafCert2 — the
		// nocLeafCert2* constants now hold VvscLeafCert2 values).
		txResult, err = AddNocX509IcaCert(nocLeafCert2Path, vendorAccount, AddNocCertOpts{IsVidVerificationSigner: true})
		cliputils.RequireTxOK(t, txResult, err)

		// Revoke the OperationalPKI ICA with revoke-child=true.
		// Cascade hits NocCert2 / NocCert2Copy but, per Matter §6.5.12, does not reach the VVSC leaf.
		txResult, err = RevokeNocX509IcaCert(nocCert2Subject, nocCert2SubjectKeyID, vendorAccount, RevokeNocCertOpts{RevokeChild: true})
		cliputils.RequireTxOK(t, txResult, err)

		// Also revoke VvscIca2 with revoke-child=true so the cascade picks up the VVSC leaf 2.
		txResult, err = RevokeNocX509IcaCert(vvscIcaCert2Subject, vvscIcaCert2SubjectKeyID, vendorAccount, RevokeNocCertOpts{RevokeChild: true})
		cliputils.RequireTxOK(t, txResult, err)

		// Revoked ICA list contains NocCert2 (+ copy), VvscIca2, and VvscLeaf2.
		revokedIcas, err := GetAllRevokedNocX509IcaCerts()
		require.NoError(t, err)
		require.True(t, containsRevokedNocIcaCertSubject(revokedIcas, nocCert2Subject))
		require.True(t, containsRevokedNocIcaCertSubject(revokedIcas, vvscIcaCert2Subject))
		require.True(t, containsRevokedNocIcaCertSubject(revokedIcas, nocLeafCert2Subject))
		require.True(t, containsRevokedNocIcaCertSerial(revokedIcas, nocCert2SerialNumber))
		require.True(t, containsRevokedNocIcaCertSerial(revokedIcas, nocRevChildCert2CopySerialNumber))
		require.True(t, containsRevokedNocIcaCertSerial(revokedIcas, nocLeafCert2SerialNumber))

		// Root should not be in revoked ICA list.
		require.False(t, containsRevokedNocIcaCertSubject(revokedIcas, nocRootCert2Subject))

		// NOC certs by VID should not contain ICA/leaf.
		icas, err = GetNocX509IcaCerts(nocVid)
		require.NoError(t, err)
		if icas != nil {
			for _, c := range icas.Certs {
				require.NotEqual(t, nocCert2Subject, c.Subject)
				require.NotEqual(t, nocLeafCert2Subject, c.Subject)
			}
		}

		// All NOC certs should not contain revoked ICA/leaf but should still have root.
		all, err := GetAllNocX509Certs()
		require.NoError(t, err)
		require.NotNil(t, all)
		require.True(t, containsCertSubjectSerial(all.Certs, nocRootCert2Subject, nocRootCert2SerialNumber))
		for _, c := range all.Certs {
			require.NotEqual(t, nocCert2Subject, c.Subject)
			require.NotEqual(t, nocLeafCert2Subject, c.Subject)
		}
	})
}
