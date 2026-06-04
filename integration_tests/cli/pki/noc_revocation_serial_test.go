package pki

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	cliputils "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/cli/utils"
	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/utils"
)

// After TestPKINocRevocationWithRevokingChild, noc_root_cert_1/copy, noc_cert_1, and noc_leaf_cert_1
// are in the revoked pool. This test removes them and re-adds them before running serial revocation.
func TestPKINocRevocationWithSerialNumber(t *testing.T) {
	vendorAccount := fmt.Sprintf("vendor_account_%d", nocVid)
	cliputils.CreateVendorAccount(t, vendorAccount, nocVid)

	t.Run("RevokeNocRootCertBySerial", func(t *testing.T) {
		// Remove revoked root certs from TestPKINocRevocationWithRevokingChild so they can be re-added.
		txResult, err := RemoveNocRootCert(nocRootCert1Subject, nocRootCert1SubjectKeyID, vendorAccount)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Remove revoked ICA and leaf certs so they can be re-added.
		txResult, err = RemoveNocCert(nocCert1Subject, nocCert1SubjectKeyID, vendorAccount)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		txResult, err = RemoveNocCert(nocLeafCert1Subject, nocLeafCert1SubjectKeyID, vendorAccount)
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

		// Re-add ICA and leaf certs
		txResult, err = AddNocX509IcaCert(nocCert1Path, vendorAccount)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		txResult, err = AddNocX509IcaCert(nocLeafCert1Path, vendorAccount)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Try to revoke with invalid serial number
		txResult, err = RevokeNocRootCert(nocRootCert1Subject, nocRootCert1SubjectKeyID, vendorAccount,
			"--serial-number", "invalid",
		)
		require.NoError(t, err)
		require.Equal(t, uint32(404), txResult.Code)

		// Revoke only first root cert by serial number (child certs should NOT be revoked)
		txResult, err = RevokeNocRootCert(nocRootCert1Subject, nocRootCert1SubjectKeyID, vendorAccount,
			"--serial-number", nocRootCert1SerialNumber,
		)
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
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, nocLeafCert1Subject))

		// Second root cert should still be active
		out, err = QueryNocRootCerts(nocVid)
		require.NoError(t, err)
		require.Contains(t, string(out), nocRootCert1CopySerialNumber)
		require.NotContains(t, string(out), nocRootCert1SerialNumber)

		// ICA and leaf should still be active
		out, err = QueryNocX509IcaCerts(nocVid)
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, nocCert1Subject))
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, nocLeafCert1Subject))

		// Revoke second root cert with revoke-child=true
		txResult, err = RevokeNocRootCert(nocRootCert1Subject, nocRootCert1SubjectKeyID, vendorAccount,
			"--serial-number", nocRootCert1CopySerialNumber,
			"--revoke-child=true",
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Both root certs should be revoked
		out, err = QueryAllRevokedNocRootCerts()
		require.NoError(t, err)
		require.Contains(t, string(out), nocRootCert1SerialNumber)
		require.Contains(t, string(out), nocRootCert1CopySerialNumber)

		// ICA and leaf should also be revoked
		out, err = QueryAllRevokedNocX509IcaCerts()
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, nocCert1Subject))
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, nocLeafCert1Subject))

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
		// noc_cert_2, noc_cert_2_copy, and noc_leaf_cert_2 were revoked by RevokeNocIcaCertWithChildFlag.
		// Remove them so they can be re-added.
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

		// Re-add ICA certs (noc_root_cert_2 is already active — no need to re-add)
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

		txResult, err = AddNocX509IcaCert(nocLeafCert2Path, vendorAccount)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Try to revoke with invalid serial number
		txResult, err = RevokeNocX509IcaCert(nocCert2Subject, nocCert2SubjectKeyID, vendorAccount,
			"--serial-number", "invalid",
		)
		require.NoError(t, err)
		require.Equal(t, uint32(404), txResult.Code)

		// Revoke only first ICA cert by serial (child should not be revoked)
		txResult, err = RevokeNocX509IcaCert(nocCert2Subject, nocCert2SubjectKeyID, vendorAccount,
			"--serial-number", nocCert2SerialNumber,
		)
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

		// Revoke second ICA cert with revoke-child=true
		txResult, err = RevokeNocX509IcaCert(nocCert2Subject, nocCert2SubjectKeyID, vendorAccount,
			"--serial-number", nocRevChildCert2CopySerialNumber,
			"--revoke-child=true",
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Both ICA certs and leaf should be revoked
		out, err = QueryAllRevokedNocX509IcaCerts()
		require.NoError(t, err)
		require.Contains(t, string(out), nocCert2SerialNumber)
		require.Contains(t, string(out), nocRevChildCert2CopySerialNumber)
		require.Contains(t, string(out), nocLeafCert2SerialNumber)

		// Only root cert should remain
		out, err = QueryAllNocX509Certs()
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, nocRootCert2Subject))
		require.NotContains(t, string(out), nocCert2Subject)
		require.NotContains(t, string(out), nocLeafCert2Subject)
	})
}
