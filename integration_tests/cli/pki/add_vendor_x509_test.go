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
	addVendorVid = 65522

	// Use google_root_cert_gsr4 + intermediate_cert_gsr4 to avoid conflicting with root_cert used by TestPKIDemo.
	addVendorRootCertPath           = "../../constants/google_root_cert_gsr4"
	addVendorRootCertSubject        = "MFAxJDAiBgNVBAsTG0dsb2JhbFNpZ24gRUNDIFJvb3QgQ0EgLSBSNDETMBEGA1UEChMKR2xvYmFsU2lnbjETMBEGA1UEAxMKR2xvYmFsU2lnbg=="
	addVendorRootCertSubjectKeyID   = "54:B0:7B:AD:45:B8:E2:40:7F:FB:0A:6E:FB:BE:33:C9:3C:A3:84:D5"
	addVendorIntermCertPath         = "../../constants/intermediate_cert_gsr4"
	addVendorIntermCertSubject      = "MEYxCzAJBgNVBAYTAlVTMSIwIAYDVQQKExlHb29nbGUgVHJ1c3QgU2VydmljZXMgTExDMRMwEQYDVQQDEwpHVFMgQ0EgMkQ0"
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
