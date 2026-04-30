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
	rootCertPath         = "../../constants/root_cert"
	intermediateCertPath = "../../constants/intermediate_cert"
	leafCertPath         = "../../constants/leaf_cert"

	rootCertSubject      = "MDQxCzAJBgNVBAYTAkFVMRMwEQYDVQQIDApzb21lLXN0YXRlMRAwDgYDVQQKDAdyb290LWNh"
	rootCertSubjectKeyID = "5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB"
	rootCertSerialNumber = "442314047376310867378175982234956458728610743315"
	rootCertSubjectText  = "O=root-ca,ST=some-state,C=AU"

	intermediateCertSubject      = "MDwxCzAJBgNVBAYTAkFVMRMwEQYDVQQIDApzb21lLXN0YXRlMRgwFgYDVQQKDA9pbnRlcm1lZGlhdGUtY2E="
	intermediateCertSubjectKeyID = "4E:3B:73:F4:70:4D:C2:98:0D:DB:C8:5A:5F:02:3B:BF:86:25:56:2B"
	intermediateCertSerialNumber = "169917617234879872371588777545667947720450185023"

	leafCertSubject      = "MDExCzAJBgNVBAYTAkFVMRMwEQYDVQQIDApzb21lLXN0YXRlMQ0wCwYDVQQKDARsZWFm"
	leafCertSubjectKeyID = "30:F4:65:75:14:20:B2:AF:3D:14:71:17:AC:49:90:93:3E:24:A0:1F"
	leafCertSerialNumber = "143290473708569835418599774898811724528308722063"

	pkiDemoVid = 1
)

// TestPKIDemo translates pki-demo.sh.
func TestPKIDemo(t *testing.T) {
	jack := testconstants.JackAccount
	alice := testconstants.AliceAccount
	bob := testconstants.BobAccount

	vendorAccount := fmt.Sprintf("vendor_account_%d", pkiDemoVid)
	cliputils.CreateVendorAccount(t, vendorAccount, pkiDemoVid)

	vendorAccount65522 := "vendor_account_65522"
	cliputils.CreateVendorAccount(t, vendorAccount65522, 65522)

	userAccount := cliputils.CreateAccount(t, "CertificationCenter")

	t.Run("QueryAllEmpty", func(t *testing.T) {
		// Verify this test's specific root cert is not yet present.
		// (Other tests running earlier on the shared chain may have added other certs.)
		out, err := QueryX509Cert(rootCertSubject, rootCertSubjectKeyID)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")

		out, err = QueryProposedX509RootCert(rootCertSubject, rootCertSubjectKeyID)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")

		out, err = QueryRevokedX509Cert(rootCertSubject, rootCertSubjectKeyID)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")

		out, err = QueryProposedRevokedX509RootCert(rootCertSubject, rootCertSubjectKeyID)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")
	})

	t.Run("ProposeRootCert_NotTrustee_Fails", func(t *testing.T) {
		txResult, err := ProposeAddX509RootCert(rootCertPath, userAccount,
			"--vid", fmt.Sprintf("%d", pkiDemoVid),
		)
		require.NoError(t, err)
		require.NotEqual(t, uint32(0), txResult.Code)
		_, _ = utils.AwaitTxConfirmation(txResult.TxHash)
	})

	t.Run("ProposeRootCert_Trustee_Succeeds", func(t *testing.T) {
		txResult, err := ProposeAddX509RootCert(rootCertPath, jack,
			"--vid", fmt.Sprintf("%d", pkiDemoVid),
			"--schemaVersion", "0",
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		out, err := QueryAllProposedX509RootCerts()
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, rootCertSubject))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, rootCertSubjectKeyID))

		out, err = QueryProposedX509RootCert(rootCertSubject, rootCertSubjectKeyID)
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, rootCertSubject))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, rootCertSubjectKeyID))
		require.Contains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, rootCertSerialNumber))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectAsText":"%s"`, rootCertSubjectText))
		require.Contains(t, string(out), fmt.Sprintf(`"vid":%d`, pkiDemoVid))

		out, err = QueryX509Cert(rootCertSubject, rootCertSubjectKeyID)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")
	})

	t.Run("ApproveRootCert_Trustee", func(t *testing.T) {
		txResult, err := ApproveAddX509RootCert(rootCertSubject, rootCertSubjectKeyID, alice)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Root cert now approved
		out, err := QueryX509Cert(rootCertSubject, rootCertSubjectKeyID)
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, rootCertSubject))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, rootCertSubjectKeyID))
		require.Contains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, rootCertSerialNumber))
		require.Contains(t, string(out), `"isRoot":true`)

		out, err = QueryProposedX509RootCert(rootCertSubject, rootCertSubjectKeyID)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")
	})

	t.Run("AddIntermediateCert", func(t *testing.T) {
		txResult, err := AddX509Cert(intermediateCertPath, vendorAccount)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		out, err := QueryX509Cert(intermediateCertSubject, intermediateCertSubjectKeyID)
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, intermediateCertSubject))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, intermediateCertSubjectKeyID))
		require.Contains(t, string(out), `"isRoot":false`)
	})

	t.Run("AddLeafCert", func(t *testing.T) {
		txResult, err := AddX509Cert(leafCertPath, vendorAccount)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		out, err := QueryX509Cert(leafCertSubject, leafCertSubjectKeyID)
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, leafCertSubject))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, leafCertSubjectKeyID))
	})

	t.Run("QueryAllApprovedCerts", func(t *testing.T) {
		out, err := QueryAllX509Certs()
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, rootCertSubject))
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, intermediateCertSubject))
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, leafCertSubject))
	})

	t.Run("RevokeLeafCert", func(t *testing.T) {
		txResult, err := RevokeX509Cert(leafCertSubject, leafCertSubjectKeyID, vendorAccount)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		out, err := QueryX509Cert(leafCertSubject, leafCertSubjectKeyID)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")

		out, err = QueryRevokedX509Cert(leafCertSubject, leafCertSubjectKeyID)
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, leafCertSubject))
	})

	t.Run("RevokeIntermediateCert", func(t *testing.T) {
		// Intermediate cert is a non-root cert — use revoke-x509-cert (not propose-revoke-x509-root-cert)
		txResult, err := RevokeX509Cert(intermediateCertSubject, intermediateCertSubjectKeyID, vendorAccount)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		out, err := QueryX509Cert(intermediateCertSubject, intermediateCertSubjectKeyID)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")

		out, err = QueryRevokedX509Cert(intermediateCertSubject, intermediateCertSubjectKeyID)
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, intermediateCertSubject))
	})

	t.Run("ProposeRevokeRootCert", func(t *testing.T) {
		// With 3 trustees, quorum=2: jack proposes + alice approves = cert is revoked.
		txResult, err := ProposeRevokeX509RootCert(rootCertSubject, rootCertSubjectKeyID, jack)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		out, err := QueryProposedRevokedX509RootCert(rootCertSubject, rootCertSubjectKeyID)
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, rootCertSubject))

		txResult, err = ApproveRevokeX509RootCert(rootCertSubject, rootCertSubjectKeyID, alice)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		out, err = QueryX509Cert(rootCertSubject, rootCertSubjectKeyID)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")

		out, err = QueryRevokedX509Cert(rootCertSubject, rootCertSubjectKeyID)
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, rootCertSubject))
	})

	_, _ = bob, alice
}
