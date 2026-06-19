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

	removeX509RootCertSubject      = "MIGCMQswCQYDVQQGEwJVUzERMA8GA1UECBMITmV3IFlvcmsxETAPBgNVBAcTCE5ldyBZb3JrMRgwFgYDVQQKEw9FeGFtcGxlIENvbXBhbnkxGTAXBgNVBAsTEFRlc3RpbmcgRGl2aXNpb24xGDAWBgNVBAMTD3d3dy5leGFtcGxlLmNvbQ=="
	removeX509RootCertSubjectKeyID = "C1:48:66:ED:6F:23:D8:28:1A:D9:37:7C:58:AC:3F:DA:04:C1:41:E8"
	removeX509RootCert1SerialNum   = "1"

	removeX509IntermCertSubject      = "MEUxCzAJBgNVBAYTAkFVMRMwEQYDVQQIEwpTb21lLVN0YXRlMSEwHwYDVQQKExhJbnRlcm5ldCBXaWRnaXRzIFB0eSBMdGQ="
	removeX509IntermCertSubjectKeyID = "A1:E0:92:89:FA:18:82:12:14:9D:B8:AE:19:43:BE:44:31:6B:F1:F5"
	removeX509IntermCert1Path        = "../../constants/intermediate_with_same_subject_and_skid_1"
	removeX509IntermCert2Path        = "../../constants/intermediate_with_same_subject_and_skid_2"
	removeX509IntermCert1SerialNum   = "3"
	removeX509IntermCert2SerialNum   = "4"

	removeX509LeafCertSubject      = "MEUxCzAJBgNVBAYTAkFVMRMwEQYDVQQIEwpTb21lLVN0YXRlMSEwHwYDVQQKExhJbnRlcm5ldCBXaWRnaXRzIFB0eSBMdGQ="
	removeX509LeafCertSubjectKeyID = "90:81:84:C7:EC:B8:81:14:66:61:2F:82:BB:E9:51:67:F2:4D:99:A3"
	removeX509LeafCertPath         = "../../constants/leaf_with_same_subject_and_skid"
	removeX509LeafCertSerialNum    = "5"
)

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
