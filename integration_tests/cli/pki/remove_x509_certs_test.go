package pki

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	cliputils "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/cli/utils"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/utils"
)

const (
	removeX509RootCertVid = 65521
	removeX509OtherVid    = 65522

	removeX509RootCertSubject      = "MIGCMQswCQYDVQQGEwJVUzERMA8GA1UECAwITmV3IFlvcmsxETAPBgNVBAcMCE5ldyBZb3JrMRgwFgYDVQQKDA9FeGFtcGxlIENvbXBhbnkxGTAXBgNVBAsMEFRlc3RpbmcgRGl2aXNpb24xGDAWBgNVBAMMD3d3dy5leGFtcGxlLmNvbQ=="
	removeX509RootCertSubjectKeyID = "33:5E:0C:07:44:F8:B5:9C:CD:55:01:9B:6D:71:23:83:6F:D0:D4:BE"
	removeX509RootCert1Path        = "../../constants/root_with_same_subject_and_skid_1"
	removeX509RootCert1SerialNum   = "1"

	removeX509IntermCertSubject      = "MEUxCzAJBgNVBAYTAkFVMRMwEQYDVQQIDApTb21lLVN0YXRlMSEwHwYDVQQKDBhJbnRlcm5ldCBXaWRnaXRzIFB0eSBMdGQ="
	removeX509IntermCertSubjectKeyID = "2E:13:3B:44:52:2C:30:E9:EC:FB:45:FA:5D:E5:04:0A:C1:C6:E6:B9"
	removeX509IntermCert1Path        = "../../constants/intermediate_with_same_subject_and_skid_1"
	removeX509IntermCert2Path        = "../../constants/intermediate_with_same_subject_and_skid_2"
	removeX509IntermCert1SerialNum   = "3"
	removeX509IntermCert2SerialNum   = "4"

	removeX509LeafCertSubject      = "MEUxCzAJBgNVBAYTAkFVMRMwEQYDVQQIDApTb21lLVN0YXRlMSEwHwYDVQQKDBhJbnRlcm5ldCBXaWRnaXRzIFB0eSBMdGQ="
	removeX509LeafCertSubjectKeyID = "12:16:55:8E:5E:2A:DF:04:D7:E6:FE:D1:53:69:61:98:EF:17:2F:03"
	removeX509LeafCertPath         = "../../constants/leaf_with_same_subject_and_skid"
	removeX509LeafCertSerialNum    = "5"
)

// TestPKIRemoveX509Certificates translates pki-remove-x509-certificates.sh.
func TestPKIRemoveX509Certificates(t *testing.T) {
	jack := testconstants.JackAccount

	vendorAccount65521 := fmt.Sprintf("vendor_account_%d", removeX509RootCertVid)
	cliputils.CreateVendorAccount(t, vendorAccount65521, removeX509RootCertVid)

	vendorAccount65522 := fmt.Sprintf("vendor_account_%d", removeX509OtherVid)
	cliputils.CreateVendorAccount(t, vendorAccount65522, removeX509OtherVid)

	t.Run("SetupCerts", func(t *testing.T) {
		// Root cert was already proposed and approved by TestPKICombineCerts.
		// Verify it is on-chain before proceeding.
		out, err := QueryX509Cert(removeX509RootCertSubject, removeX509RootCertSubjectKeyID)
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, removeX509RootCertSubject))

		// Add intermediate certs
		txResult, err := AddX509Cert(removeX509IntermCert1Path, vendorAccount65521)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		txResult, err = AddX509Cert(removeX509IntermCert2Path, vendorAccount65521)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Add leaf cert
		txResult, err = AddX509Cert(removeX509LeafCertPath, vendorAccount65521)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		out, err = QueryAllX509Certs()
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, removeX509RootCert1SerialNum))
		require.Contains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, removeX509IntermCert1SerialNum))
		require.Contains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, removeX509IntermCert2SerialNum))
		require.Contains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, removeX509LeafCertSerialNum))
	})

	t.Run("RevokeAndRemoveIntermCert", func(t *testing.T) {
		// Revoke first intermediate cert
		txResult, err := RevokeX509Cert(removeX509IntermCertSubject, removeX509IntermCertSubjectKeyID, vendorAccount65521,
			"--serial-number", removeX509IntermCert1SerialNum,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Try to remove with invalid serial
		txResult, err = RemoveX509Cert(removeX509IntermCertSubject, removeX509IntermCertSubjectKeyID, vendorAccount65521,
			"--serial-number", "invalid",
		)
		require.NoError(t, err)
		require.Equal(t, uint32(404), txResult.Code)

		// Try to remove when sender is not Vendor
		txResult, err = RemoveX509Cert(removeX509IntermCertSubject, removeX509IntermCertSubjectKeyID, jack,
			"--serial-number", removeX509IntermCert1SerialNum,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(4), txResult.Code)

		// Try to remove with different VID vendor
		txResult, err = RemoveX509Cert(removeX509IntermCertSubject, removeX509IntermCertSubjectKeyID, vendorAccount65522,
			"--serial-number", removeX509IntermCert1SerialNum,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(4), txResult.Code)

		// Remove revoked intermediate cert by serial
		txResult, err = RemoveX509Cert(removeX509IntermCertSubject, removeX509IntermCertSubjectKeyID, vendorAccount65521,
			"--serial-number", removeX509IntermCert1SerialNum,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Only second intermediate cert should remain
		out, err := QueryX509Cert(removeX509IntermCertSubject, removeX509IntermCertSubjectKeyID)
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, removeX509IntermCert2SerialNum))
		require.NotContains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, removeX509IntermCert1SerialNum))

		// Remove remaining intermediate cert by subject+subjectKeyID
		txResult, err = RemoveX509Cert(removeX509IntermCertSubject, removeX509IntermCertSubjectKeyID, vendorAccount65521)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		out, err = QueryX509Cert(removeX509IntermCertSubject, removeX509IntermCertSubjectKeyID)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")

		// This test's intermediate certs should not be in the revoked list.
		// (TestPKIDemo may have left other certs revoked, so we cannot assert "[]".)
		out, err = QueryAllRevokedX509Certs()
		require.NoError(t, err)
		require.NotContains(t, string(out), removeX509IntermCertSubject)

		// Only root and leaf should remain
		out, err = QueryAllX509Certs()
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, removeX509RootCert1SerialNum))
		require.Contains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, removeX509LeafCertSerialNum))
		require.NotContains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, removeX509IntermCert1SerialNum))
		require.NotContains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, removeX509IntermCert2SerialNum))
	})

	t.Run("RemoveLeafCert", func(t *testing.T) {
		txResult, err := RemoveX509Cert(removeX509LeafCertSubject, removeX509LeafCertSubjectKeyID, vendorAccount65521)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		out, err := QueryX509Cert(removeX509LeafCertSubject, removeX509LeafCertSubjectKeyID)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")

		// Only root should remain
		out, err = QueryAllX509Certs()
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, removeX509RootCert1SerialNum))
		require.NotContains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, removeX509LeafCertSerialNum))
	})
}
