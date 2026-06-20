package pki

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	cliputils "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/cli/utils"
	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/utils"
)

// After TestPKINocRevocationWithRevokingChild, noc_root_cert_1/copy, noc_cert_1,
// vvsc_root_cert_1, vvsc_ica_cert_1, and vvsc_leaf_cert_1 are in the revoked pool.
// (VvscRoot1Copy was added in the prior test and is still active.)
// This test removes them and re-adds them before running serial revocation.
func TestPKINocRevocationWithSerialNumber(t *testing.T) {
	vendorAccount := fmt.Sprintf("vendor_account_%d", nocVid)
	cliputils.CreateVendorAccount(t, vendorAccount, nocVid)

	t.Run("RevokeNocRootCertBySerial", func(t *testing.T) {
		// Remove revoked NOC root certs so they can be re-added.
		txResult, err := RemoveNocRootCert(nocRootCert1Subject, nocRootCert1SubjectKeyID, vendorAccount)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Remove revoked NOC ICA so it can be re-added.
		txResult, err = RemoveNocCert(nocCert1Subject, nocCert1SubjectKeyID, vendorAccount)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Remove revoked VVSC leaf + VVSC ICA so they can be re-added.
		// VVSC root was also revoked but VvscRoot1Copy (active, same SKID) still
		// satisfies AuthorityKeyID resolution for the chain walk, so we don't
		// remove/re-add the VVSC root here — re-adding only ICA1 + Leaf1 is enough.
		txResult, err = RemoveNocCert(vvscLeafCert1Subject, vvscLeafCert1SubjectKeyID, vendorAccount)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		txResult, err = RemoveNocCert(vvscIcaCert1Subject, vvscIcaCert1SubjectKeyID, vendorAccount)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Re-add NOC root certs
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

		// Re-add NOC ICA cert
		txResult, err = AddNocX509IcaCert(nocCert1Path, vendorAccount)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Re-add VVSC ICA1 + Leaf1 (chained under VvscRoot1Copy which is still active).
		txResult, err = AddNocX509IcaCert(vvscIcaCert1Path, vendorAccount, AddNocCertOpts{IsVidVerificationSigner: true})
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		txResult, err = AddNocX509IcaCert(vvscLeafCert1Path, vendorAccount, AddNocCertOpts{IsVidVerificationSigner: true})
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Try to revoke with invalid serial number
		txResult, err = RevokeNocRootCert(nocRootCert1Subject, nocRootCert1SubjectKeyID, vendorAccount, RevokeNocCertOpts{SerialNumber: "invalid"})
		require.NoError(t, err)
		require.Equal(t, uint32(404), txResult.Code)

		// Revoke only first root cert by serial number (child certs should NOT be revoked)
		txResult, err = RevokeNocRootCert(nocRootCert1Subject, nocRootCert1SubjectKeyID, vendorAccount, RevokeNocCertOpts{SerialNumber: nocRootCert1SerialNumber})
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Only first root cert should be in revoked list
		out, err := QueryAllRevokedNocRootCerts()
		require.NoError(t, err)
		require.Contains(t, string(out), nocRootCert1SerialNumber)
		require.NotContains(t, string(out), nocRootCert1CopySerialNumber)
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, nocCert1Subject))
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, vvscLeafCert1Subject))

		// Second root cert should still be active
		out, err = QueryNocRootCerts(nocVid)
		require.NoError(t, err)
		require.Contains(t, string(out), nocRootCert1CopySerialNumber)
		require.NotContains(t, string(out), nocRootCert1SerialNumber)

		// NOC ICA + VVSC leaf should still be active (VVSC chain is structurally
		// disjoint from the NOC root revocation — Matter §6.5.12).
		out, err = QueryNocX509IcaCerts(nocVid)
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, nocCert1Subject))
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, vvscLeafCert1Subject))

		// Revoke second root cert with revoke-child=true. Cascade hits NocCert1
		// (under the NOC chain) — the VVSC chain is structurally disjoint.
		txResult, err = RevokeNocRootCert(nocRootCert1Subject, nocRootCert1SubjectKeyID, vendorAccount, RevokeNocCertOpts{SerialNumber: nocRootCert1CopySerialNumber, RevokeChild: true})
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Also revoke the VVSC root with revoke-child=true. The VVSC chain
		// VvscRoot1 → VvscIca1 → VvscLeaf1 is structurally disjoint from the
		// OperationalPKI cascade (Matter §6.5.12 / §6.4.10) — without an
		// explicit VVSC root revocation the leaf would remain active and the
		// revoked-ICA assertion below would fail.
		txResult, err = RevokeNocRootCert(vvscRootCert1Subject, vvscRootCert1SubjectKeyID, vendorAccount, RevokeNocCertOpts{RevokeChild: true})
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Both NOC root certs should be revoked
		out, err = QueryAllRevokedNocRootCerts()
		require.NoError(t, err)
		require.Contains(t, string(out), nocRootCert1SerialNumber)
		require.Contains(t, string(out), nocRootCert1CopySerialNumber)

		// Revoked ICA list now contains the cascaded NOC ICA + VVSC ICA + VVSC leaf.
		out, err = QueryAllRevokedNocX509IcaCerts()
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, nocCert1Subject))
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, vvscIcaCert1Subject))
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, vvscLeafCert1Subject))
		require.Contains(t, string(out), nocCert1SerialNumber)
		require.Contains(t, string(out), vvscLeafCert1SerialNumber)

		// NOC certs for VID1 (nocVid) should not have root_cert_1 active any more.
		// root_cert_2 (different chain) may still be active — that is expected.
		out, err = QueryAllRevokedNocRootCerts()
		require.NoError(t, err)
		require.Contains(t, string(out), nocRootCert1SerialNumber)
		require.Contains(t, string(out), nocRootCert1CopySerialNumber)

		out, err = QueryNocRootCerts(nocVid)
		require.NoError(t, err)
		require.NotContains(t, string(out), nocRootCert1Subject)
	})

	t.Run("RevokeNocIcaCertBySerial", func(t *testing.T) {
		// noc_root_cert_2 is active from TestPKINocRevocationWithRevokingChild (not revoked there).
		// noc_cert_2, noc_cert_2_copy, vvsc_ica_cert_2, and vvsc_leaf_cert_2 (held in
		// nocLeafCert2*) were revoked by RevokeNocIcaCertWithChildFlag. Remove them
		// so they can be re-added.
		txResult, err := RemoveNocCert(nocCert2Subject, nocCert2SubjectKeyID, vendorAccount)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		txResult, err = RemoveNocCert(nocLeafCert2Subject, nocLeafCert2SubjectKeyID, vendorAccount)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		txResult, err = RemoveNocCert(vvscIcaCert2Subject, vvscIcaCert2SubjectKeyID, vendorAccount)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Re-add OperationalPKI ICA certs (noc_root_cert_2 is already active — no need to re-add)
		txResult, err = AddNocX509IcaCert(nocCert2Path, vendorAccount)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		txResult, err = AddNocX509IcaCert(nocRevChildCert2CopyPath, vendorAccount)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Subtest 1's cascade (revoke vvsc_root_cert_1 with --revoke-child at line 130)
		// soft-deleted VvscRoot1Copy into the revoked root pool, so no active VVSC
		// trust anchor remains at (VvscRoot1 Subject, SKID). Remove both revoked
		// entries (clears their UniqueCertificate records) and re-add VvscRoot1 so
		// verifyVVSCCertificate's chain walk for VvscIca2 below resolves.
		txResult, err = RemoveNocRootCert(vvscRootCert1Subject, vvscRootCert1SubjectKeyID, vendorAccount)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		txResult, err = AddNocRootCert(vvscRootCert1Path, vendorAccount, AddNocCertOpts{IsVidVerificationSigner: true})
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Re-add VVSC ICA2 + VVSC leaf 2 (chained under the freshly re-added VvscRoot1).
		txResult, err = AddNocX509IcaCert(vvscIcaCert2Path, vendorAccount, AddNocCertOpts{IsVidVerificationSigner: true})
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		txResult, err = AddNocX509IcaCert(nocLeafCert2Path, vendorAccount, AddNocCertOpts{IsVidVerificationSigner: true})
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Try to revoke with invalid serial number
		txResult, err = RevokeNocX509IcaCert(nocCert2Subject, nocCert2SubjectKeyID, vendorAccount, RevokeNocCertOpts{SerialNumber: "invalid"})
		require.NoError(t, err)
		require.Equal(t, uint32(404), txResult.Code)

		// Revoke only first ICA cert by serial (child should not be revoked)
		txResult, err = RevokeNocX509IcaCert(nocCert2Subject, nocCert2SubjectKeyID, vendorAccount, RevokeNocCertOpts{SerialNumber: nocCert2SerialNumber})
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		out, err := QueryAllRevokedNocX509IcaCerts()
		require.NoError(t, err)
		require.Contains(t, string(out), nocCert2SerialNumber)
		require.NotContains(t, string(out), nocRevChildCert2CopySerialNumber)
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, nocLeafCert2Subject))

		// Second ICA cert should still be active
		out, err = QueryNocCert("--subject-key-id", nocCert2SubjectKeyID)
		require.NoError(t, err)
		require.Contains(t, string(out), nocRevChildCert2CopySerialNumber)
		require.NotContains(t, string(out), nocCert2SerialNumber)

		// Revoke second NOC ICA cert with revoke-child=true. Cascade is contained
		// to the NOC chain — the VVSC leaf is structurally disjoint (Matter §6.5.12).
		txResult, err = RevokeNocX509IcaCert(nocCert2Subject, nocCert2SubjectKeyID, vendorAccount, RevokeNocCertOpts{SerialNumber: nocRevChildCert2CopySerialNumber, RevokeChild: true})
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Also revoke VvscIca2 with revoke-child=true so the cascade picks up VVSC leaf 2.
		txResult, err = RevokeNocX509IcaCert(vvscIcaCert2Subject, vvscIcaCert2SubjectKeyID, vendorAccount, RevokeNocCertOpts{RevokeChild: true})
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// All ICAs (NOC + VVSC) and the VVSC leaf should be revoked.
		out, err = QueryAllRevokedNocX509IcaCerts()
		require.NoError(t, err)
		require.Contains(t, string(out), nocCert2SerialNumber)
		require.Contains(t, string(out), nocRevChildCert2CopySerialNumber)
		require.Contains(t, string(out), vvscIcaCert2SerialNumber)
		require.Contains(t, string(out), nocLeafCert2SerialNumber)

		// Only root cert should remain in the active NOC list (for nocVid).
		out, err = QueryAllNocX509Certs()
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, nocRootCert2Subject))
		require.NotContains(t, string(out), nocCert2Subject)
		require.NotContains(t, string(out), nocLeafCert2Subject)
	})
}
