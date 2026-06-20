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
		rootCert, err := GetX509Cert(removeX509RootCertSubject, removeX509RootCertSubjectKeyID)
		require.NoError(t, err)
		require.NotNil(t, rootCert)
		require.Equal(t, removeX509RootCertSubject, rootCert.Subject)

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

		all, err := GetAllX509Certs()
		require.NoError(t, err)
		require.True(t, containsApprovedCertSerial(all, removeX509RootCert1SerialNum))
		require.True(t, containsApprovedCertSerial(all, removeX509IntermCert1SerialNum))
		require.True(t, containsApprovedCertSerial(all, removeX509IntermCert2SerialNum))
		require.True(t, containsApprovedCertSerial(all, removeX509LeafCertSerialNum))
	})

	t.Run("RevokeAndRemoveIntermCert", func(t *testing.T) {
		// Revoke first intermediate cert
		txResult, err := RevokeX509Cert(removeX509IntermCertSubject, removeX509IntermCertSubjectKeyID, vendorAccount65521, RevokeNocCertOpts{SerialNumber: removeX509IntermCert1SerialNum})
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Try to remove with invalid serial
		txResult, err = RemoveX509Cert(removeX509IntermCertSubject, removeX509IntermCertSubjectKeyID, vendorAccount65521, RevokeNocCertOpts{SerialNumber: "invalid"})
		require.NoError(t, err)
		require.Equal(t, uint32(404), txResult.Code)

		// Try to remove when sender is not Vendor
		txResult, err = RemoveX509Cert(removeX509IntermCertSubject, removeX509IntermCertSubjectKeyID, jack, RevokeNocCertOpts{SerialNumber: removeX509IntermCert1SerialNum})
		require.NoError(t, err)
		require.Equal(t, uint32(4), txResult.Code)

		// Try to remove with different VID vendor
		txResult, err = RemoveX509Cert(removeX509IntermCertSubject, removeX509IntermCertSubjectKeyID, vendorAccount65522, RevokeNocCertOpts{SerialNumber: removeX509IntermCert1SerialNum})
		require.NoError(t, err)
		require.Equal(t, uint32(4), txResult.Code)

		// Remove revoked intermediate cert by serial
		txResult, err = RemoveX509Cert(removeX509IntermCertSubject, removeX509IntermCertSubjectKeyID, vendorAccount65521, RevokeNocCertOpts{SerialNumber: removeX509IntermCert1SerialNum})
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Only second intermediate cert should remain
		cert, err := GetX509Cert(removeX509IntermCertSubject, removeX509IntermCertSubjectKeyID)
		require.NoError(t, err)
		require.NotNil(t, cert)
		require.True(t, containsCertSerial(cert.Certs, removeX509IntermCert2SerialNum))
		require.False(t, containsCertSerial(cert.Certs, removeX509IntermCert1SerialNum))

		// Remove remaining intermediate cert by subject+subjectKeyID
		txResult, err = RemoveX509Cert(removeX509IntermCertSubject, removeX509IntermCertSubjectKeyID, vendorAccount65521)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		cert, err = GetX509Cert(removeX509IntermCertSubject, removeX509IntermCertSubjectKeyID)
		require.NoError(t, err)
		require.Nil(t, cert)

		// This test's intermediate certs should not be in the revoked list.
		// (TestPKIDemo may have left other certs revoked, so we cannot assert "[]".)
		revoked, err := GetAllRevokedX509Certs()
		require.NoError(t, err)
		for _, r := range revoked {
			require.NotEqual(t, removeX509IntermCertSubject, r.Subject)
		}

		// Only root and leaf should remain
		all, err := GetAllX509Certs()
		require.NoError(t, err)
		require.True(t, containsApprovedCertSerial(all, removeX509RootCert1SerialNum))
		require.True(t, containsApprovedCertSerial(all, removeX509LeafCertSerialNum))
		require.False(t, containsApprovedCertSerial(all, removeX509IntermCert1SerialNum))
		require.False(t, containsApprovedCertSerial(all, removeX509IntermCert2SerialNum))
	})

	t.Run("RemoveLeafCert", func(t *testing.T) {
		txResult, err := RemoveX509Cert(removeX509LeafCertSubject, removeX509LeafCertSubjectKeyID, vendorAccount65521)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		cert, err := GetX509Cert(removeX509LeafCertSubject, removeX509LeafCertSubjectKeyID)
		require.NoError(t, err)
		require.Nil(t, cert)

		// Only root should remain
		all, err := GetAllX509Certs()
		require.NoError(t, err)
		require.True(t, containsApprovedCertSerial(all, removeX509RootCert1SerialNum))
		require.False(t, containsApprovedCertSerial(all, removeX509LeafCertSerialNum))
	})
}
