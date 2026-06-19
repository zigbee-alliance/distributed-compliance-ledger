package pki

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	cliputils "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/cli/utils"
	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/utils"
)

const (
	nocRevChildCert2CopyPath         = "../../constants/noc_cert_2_copy"
	nocRevChildCert2CopySerialNumber = "252687488758567844896720928536709119387931444024"

	// nocLeafCert2* are repurposed to point at vvsc_leaf_cert_2 because the
	// strict §6.5.12 ICA handler no longer accepts NOC end-entity profiles via
	// add-noc-x509-ica-cert. The cascade-revoke section keeps the same
	// variable names so its assertions still read naturally.
	nocLeafCert2Path         = "../../constants/vvsc_leaf_cert_2"
	nocLeafCert2Subject      = "MIGYMQswCQYDVQQGEwJVWjETMBEGA1UECAwKU29tZSBTdGF0ZTETMBEGA1UEBwwKU29tZSBTdGF0ZTEYMBYGA1UEChMPRXhhbXBsZSBDb21wYW55MRkwFwYDVQQLDBBUZXN0aW5nIERpdmlzaW9uMRQwEgYDVQQDDAtWVlNDLUxlYWYtMjEUMBIGCisGAQQBgqJ8AgEMBDAwMDE="
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
	vvscIcaCert2Subject      = "MIGXMQswCQYDVQQGEwJVWjETMBEGA1UECAwKU29tZSBTdGF0ZTETMBEGA1UEBwwKU29tZSBTdGF0ZTEYMBYGA1UEChMPRXhhbXBsZSBDb21wYW55MRkwFwYDVQQLDBBUZXN0aW5nIERpdmlzaW9uMRMwEQYDVQQDDApWVlNDLUlDQS0yMRQwEgYKKwYBBAGConwCAQwEMDAwMQ=="
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
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// noc_cert_1 and noc_cert_1_copy are in the revoked pool from TestPKINocCerts.
		// Remove them so they can be re-added.
		txResult, err = RemoveNocCert(nocCert1Subject, nocCert1SubjectKeyID, vendorAccount)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Re-add root certs
		txResult, err = AddNocRootCert(nocRootCert1Path, vendorAccount)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		txResult, err = AddNocRootCert(nocRootCert1CopyPath, vendorAccount)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Re-add ICA cert (noc_cert_1)
		txResult, err = AddNocX509IcaCert(nocCert1Path, vendorAccount)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// VvscLeafCert1 is already active on-chain from TestPKINocCerts (chained
		// under VvscRoot1 → VvscIca1, was never revoked). Verify it exists.
		out, err := QueryNocX509IcaCerts(nocVid)
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, vvscLeafCert1Subject))

		// Verify root certs exist
		out, err = QueryAllNocRootCerts()
		require.NoError(t, err)
		require.Contains(t, string(out), nocRootCert1SerialNumber)
		require.Contains(t, string(out), nocRootCert1CopySerialNumber)

		// Revoke the OperationalPKI root NOC certificate with revoke-child=true.
		// Cascade hits NocCert1 only — the VVSC chain is structurally disjoint (Matter §6.5.12).
		txResult, err = RevokeNocRootCert(nocRootCert1Subject, nocRootCert1SubjectKeyID, vendorAccount,
			"--revoke-child=true",
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Also revoke the VVSC root with revoke-child=true so the cascade picks
		// up VvscIca1 + the VVSC leaf.
		txResult, err = RevokeNocRootCert(vvscRootCert1Subject, vvscRootCert1SubjectKeyID, vendorAccount,
			"--revoke-child=true",
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Both NOC root certs should be revoked
		out, err = QueryAllRevokedNocRootCerts()
		require.NoError(t, err)
		require.Contains(t, string(out), nocRootCert1SerialNumber)
		require.Contains(t, string(out), nocRootCert1CopySerialNumber)
		require.Contains(t, string(out), vvscRootCert1SerialNumber)

		// Revoked ICA list now contains both the NOC ICA and the VVSC ICA + leaf.
		out, err = QueryAllRevokedNocX509IcaCerts()
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, nocCert1Subject))
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, vvscIcaCert1Subject))
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, vvscLeafCert1Subject))

		// Root cert 1 (both serials) should no longer be in the approved root list for VID.
		// (Root cert 2 from TestPKINocCerts may still be approved — don't assert "Not Found".)
		out, err = QueryNocRootCerts(nocVid)
		require.NoError(t, err)
		require.NotContains(t, string(out), nocRootCert1SerialNumber)
		require.NotContains(t, string(out), nocRootCert1CopySerialNumber)
		require.NotContains(t, string(out), vvscRootCert1SerialNumber)

		// cert1 should be revoked after revoking root cert 1 with child flag.
		// noc_cert_2 (child of root cert 2) from TestPKINocCerts is still approved — don't assert "Not Found".
		out, err = QueryNocX509IcaCerts(nocVid)
		require.NoError(t, err)
		require.NotContains(t, string(out), nocCert1SerialNumber)
		require.NotContains(t, string(out), vvscLeafCert1SerialNumber)

		// NOC certs must not appear in the DA (all-x509-certs) list.
		out, err = QueryAllX509Certs()
		require.NoError(t, err)
		require.NotContains(t, string(out), nocRootCert1Subject)
		require.NotContains(t, string(out), nocCert1Subject)
		require.NotContains(t, string(out), vvscLeafCert1Subject)
	})

	t.Run("RevokeNocIcaCertWithChildFlag", func(t *testing.T) {
		// noc_root_cert_2 and noc_cert_2 were already added by TestPKINocCerts — verify they are on-chain.
		out, err := QueryNocRootCerts(nocVid)
		require.NoError(t, err)
		require.Contains(t, string(out), nocRootCert2SerialNumber)

		out, err = QueryNocX509IcaCerts(nocVid)
		require.NoError(t, err)
		require.Contains(t, string(out), nocCert2SerialNumber)

		// Add cert2copy (OperationalPKI ICA copy) — not yet on-chain.
		txResult, err := AddNocX509IcaCert(nocRevChildCert2CopyPath, vendorAccount)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Re-establish an active VVSC root via VvscRootCert1Copy. Revocation
		// soft-deletes the cert into the revoked list but the
		// (Issuer, SerialNumber) UniqueCertificate record survives — re-adding
		// the same PEM would fail with ErrCertificateAlreadyExists. The Copy
		// shares VvscRoot1's key (same Subject + SubjectKeyID) so VvscIca2's
		// AuthorityKeyID still resolves to a present VIDSignerPKI entry during
		// verifyVVSCCertificate's chain walk.
		txResult, err = AddNocRootCert(vvscRootCert1CopyPath, vendorAccount, "--is-vid-verification-signer=true")
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Pre-seed the second VVSC intermediate (VvscIca2 under VvscRoot1) so
		// the leaf-2 chain VvscRoot1 → VvscIca2 → vvsc_leaf_cert_2 resolves
		// through verifyVVSCCertificate.
		txResult, err = AddNocX509IcaCert(vvscIcaCert2Path, vendorAccount, "--is-vid-verification-signer=true")
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Add VVSC leaf certificate 2 (replaces the legacy NocLeafCert2 — the
		// nocLeafCert2* constants now hold VvscLeafCert2 values).
		txResult, err = AddNocX509IcaCert(nocLeafCert2Path, vendorAccount, "--is-vid-verification-signer=true")
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Revoke the OperationalPKI ICA with revoke-child=true.
		// Cascade hits NocCert2 / NocCert2Copy but, per Matter §6.5.12, does not reach the VVSC leaf.
		txResult, err = RevokeNocX509IcaCert(nocCert2Subject, nocCert2SubjectKeyID, vendorAccount,
			"--revoke-child=true",
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Also revoke VvscIca2 with revoke-child=true so the cascade picks up the VVSC leaf 2.
		txResult, err = RevokeNocX509IcaCert(vvscIcaCert2Subject, vvscIcaCert2SubjectKeyID, vendorAccount,
			"--revoke-child=true",
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Revoked ICA list contains NocCert2 (+ copy), VvscIca2, and VvscLeaf2.
		out, err = QueryAllRevokedNocX509IcaCerts()
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, nocCert2Subject))
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, vvscIcaCert2Subject))
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, nocLeafCert2Subject))
		require.Contains(t, string(out), nocCert2SerialNumber)
		require.Contains(t, string(out), nocRevChildCert2CopySerialNumber)
		require.Contains(t, string(out), nocLeafCert2SerialNumber)

		// Root should not be in revoked ICA list
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s`, nocRootCert2Subject))

		// NOC certs by VID should not contain ICA/leaf
		out, err = QueryNocX509IcaCerts(nocVid)
		require.NoError(t, err)
		require.NotContains(t, string(out), nocCert2Subject)
		require.NotContains(t, string(out), nocLeafCert2Subject)

		// All NOC certs should not contain revoked ICA/leaf but should still have root
		out, err = QueryAllNocX509Certs()
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, nocRootCert2Subject))
		require.NotContains(t, string(out), nocCert2Subject)
		require.NotContains(t, string(out), nocLeafCert2Subject)
	})
}
