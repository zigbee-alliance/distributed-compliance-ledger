package pki

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	cliputils "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/cli/utils"
	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/utils"
)

const (
	leafCertWithVid65522Path         = "../../constants/leaf_cert_with_vid_65522"
	leafCertWithVid65522Subject      = "MDExCzAJBgNVBAYTAkFVMRMwEQYDVQQIDApzb21lLXN0YXRlMQ0wCwYDVQQKDARsZWFm"
	leafCertWithVid65522SubjectKeyID = "30:F4:65:75:14:20:B2:AF:3D:14:71:17:AC:49:90:93:3E:24:A0:1F"
	intermCertWithVid1Path           = "../../constants/intermediate_cert_with_vid_1"
	intermCertWithVid1Subject        = "MDwxCzAJBgNVBAYTAkFVMRMwEQYDVQQIDApzb21lLXN0YXRlMRgwFgYDVQQKDA9pbnRlcm1lZGlhdGUtY2E="
	intermCertWithVid1SubjectKeyID   = "4E:3B:73:F4:70:4D:C2:98:0D:DB:C8:5A:5F:02:3B:BF:86:25:56:2B"
	addVendorVid                     = 65522

	// Use google_root_cert_gsr4 + intermediate_cert_gsr4 to avoid conflicting with root_cert used by TestPKIDemo.
	addVendorRootCertPath         = "../../constants/google_root_cert_gsr4"
	addVendorRootCertSubject      = "MFAxJDAiBgNVBAsTG0dsb2JhbFNpZ24gRUNDIFJvb3QgQ0EgLSBSNDETMBEGA1UEChMKR2xvYmFsU2lnbjETMBEGA1UEAxMKR2xvYmFsU2lnbg=="
	addVendorRootCertSubjectKeyID = "54:B0:7B:AD:45:B8:E2:40:7F:FB:0A:6E:FB:BE:33:C9:3C:A3:84:D5"
	addVendorIntermCertPath       = "../../constants/intermediate_cert_gsr4"
	addVendorIntermCertSubject    = "MEYxCzAJBgNVBAYTAlVTMSIwIAYDVQQKExlHb29nbGUgVHJ1c3QgU2VydmljZXMgTExDMRMwEQYDVQQDEwpHVFMgQ0EgMkQ0"
	addVendorIntermCertSubjectKeyID = "A8:88:D9:8A:39:AC:65:D5:82:4B:37:A8:95:6C:65:43:CD:44:01:E0"
)

// TestPKIAddVendorX509Certificates translates pki-add-vendor-x509-certificates.sh.
func TestPKIAddVendorX509Certificates(t *testing.T) {
	jack := testconstants.JackAccount
	alice := testconstants.AliceAccount

	vendorAccount := fmt.Sprintf("vendor_account_%d", addVendorVid)
	cliputils.CreateVendorAccount(t, vendorAccount, addVendorVid)

	vendorAccount1 := fmt.Sprintf("vendor_account_%d", pkiDemoVid)
	cliputils.CreateVendorAccount(t, vendorAccount1, pkiDemoVid)

	t.Run("ProposeAndApproveRootCert", func(t *testing.T) {
		txResult, err := ProposeAddX509RootCert(addVendorRootCertPath, jack,
			"--vid", fmt.Sprintf("%d", pkiDemoVid),
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		txResult, err = ApproveAddX509RootCert(addVendorRootCertSubject, addVendorRootCertSubjectKeyID, alice)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)
	})

	t.Run("AddIntermediateCert_WithVendorVid", func(t *testing.T) {
		txResult, err := AddX509Cert(addVendorIntermCertPath, vendorAccount1)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		out, err := QueryX509Cert(addVendorIntermCertSubject, addVendorIntermCertSubjectKeyID)
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, addVendorIntermCertSubject))
	})

	t.Run("AddLeafCert_WrongVendor_Fails", func(t *testing.T) {
		// Adding leaf cert signed by intermediate belonging to vid=1 from a vendor with vid=65522
		txResult, err := AddX509Cert(leafCertPath, vendorAccount)
		// This should fail — vendor VID mismatch
		if err == nil && txResult != nil {
			require.NotEqual(t, uint32(0), txResult.Code)
		}
	})
}
