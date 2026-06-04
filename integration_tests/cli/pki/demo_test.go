package pki

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	cliputils "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/cli/utils"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
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
	intermediateCertSubjectText  = "O=intermediate-ca,ST=some-state,C=AU"

	leafCertSubject      = "MDExCzAJBgNVBAYTAkFVMRMwEQYDVQQIDApzb21lLXN0YXRlMQ0wCwYDVQQKDARsZWFm"
	leafCertSubjectKeyID = "30:F4:65:75:14:20:B2:AF:3D:14:71:17:AC:49:90:93:3E:24:A0:1F"
	leafCertSerialNumber = "143290473708569835418599774898811724528308722063"
	leafCertSubjectText  = "O=leaf,ST=some-state,C=AU"

	googleCertPath         = "../../constants/google_root_cert"
	googleCertSubject      = "MEsxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZHb29nbGUxFTATBgNVBAMMDE1hdHRlciBQQUEgMTEUMBIGCisGAQQBgqJ8AgEMBDYwMDY="
	googleCertSubjectKeyID = "B0:00:56:81:B8:88:62:89:62:80:E1:21:18:A1:A8:BE:09:DE:93:21"
	googleCertSerialNumber = "1"
	googleCertSubjectText  = "CN=Matter PAA 1,O=Google,C=US,vid=0x6006"
	googleCertVid          = 24582

	testCertPath         = "../../constants/test_root_cert"
	testCertSubject      = "MDAxGDAWBgNVBAMMD01hdHRlciBUZXN0IFBBQTEUMBIGCisGAQQBgqJ8AgEMBDEyNUQ="
	testCertSubjectKeyID = "E2:90:8D:36:9C:3C:A3:C1:13:BB:09:E2:4D:C1:CC:C5:A6:66:91:D4"
	testCertSerialNumber = "1647312298631"
	testCertSubjectText  = "CN=Matter Test PAA,vid=0x125D"
	testCertVid          = 4701

	pkiDemoVid = 1
)

func TestPKIDemo(t *testing.T) {
	jack := testconstants.JackAccount
	alice := testconstants.AliceAccount
	bob := testconstants.BobAccount

	vendorAccount := fmt.Sprintf("vendor_account_%d", pkiDemoVid)
	cliputils.CreateVendorAccount(t, vendorAccount, pkiDemoVid)

	vendorAccount65522 := "vendor_account_65522"
	cliputils.CreateVendorAccount(t, vendorAccount65522, 65522)

	userAccount := cliputils.CreateAccount(t, "CertificationCenter")

	// ── Section 1: Query All Empty ──────────────────────────────────────────────

	t.Run("QueryAllEmpty", func(t *testing.T) {
		out, err := QueryX509Cert(rootCertSubject, rootCertSubjectKeyID)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, rootCertSubject))
		require.NotContains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, rootCertSubjectKeyID))
		require.NotContains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, rootCertSerialNumber))

		out, err = QueryAllX509Certs()
		require.NoError(t, err)
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, rootCertSubject))
		require.NotContains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, rootCertSubjectKeyID))

		out, err = QueryProposedX509RootCert(rootCertSubject, rootCertSubjectKeyID)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, rootCertSubject))

		out, err = QueryAllProposedX509RootCerts()
		require.NoError(t, err)
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, rootCertSubject))

		out, err = QueryRevokedX509Cert(rootCertSubject, rootCertSubjectKeyID)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, rootCertSubject))

		out, err = QueryAllRevokedX509Certs()
		require.NoError(t, err)
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, rootCertSubject))

		out, err = QueryX509CertsBySubject(rootCertSubject)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")
		require.NotContains(t, string(out), fmt.Sprintf(`"%s"`, rootCertSubject))
		require.NotContains(t, string(out), fmt.Sprintf(`"%s"`, rootCertSubjectKeyID))

		out, err = QueryX509CertBySKID(rootCertSubjectKeyID)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")
		require.NotContains(t, string(out), fmt.Sprintf(`"%s"`, rootCertSubject))

		out, err = QueryAllX509RootCerts()
		require.NoError(t, err)
		require.NotContains(t, string(out), fmt.Sprintf(`"%s"`, rootCertSubject))

		out, err = QueryAllRevokedX509RootCerts()
		require.NoError(t, err)
		require.NotContains(t, string(out), fmt.Sprintf(`"%s"`, rootCertSubject))

		out, err = QueryProposedRevokedX509RootCert(rootCertSubject, rootCertSubjectKeyID)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")
		require.NotContains(t, string(out), fmt.Sprintf(`"%s"`, rootCertSubject))

		out, err = QueryAllProposedRevokedX509RootCerts()
		require.NoError(t, err)
		require.NotContains(t, string(out), fmt.Sprintf(`"%s"`, rootCertSubject))

		out, err = QueryChildX509Certs(rootCertSubject, rootCertSubjectKeyID)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")
		require.NotContains(t, string(out), fmt.Sprintf(`"%s"`, rootCertSubject))
	})

	// ── Section 2: Propose Root Cert ───────────────────────────────────────────

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

		// All proposed root certs — contains schemaVersion
		out, err := QueryAllProposedX509RootCerts()
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, rootCertSubject))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, rootCertSubjectKeyID))
		require.Contains(t, string(out), `"schemaVersion":0`)

		// Proposed cert by subject+skid
		out, err = QueryProposedX509RootCert(rootCertSubject, rootCertSubjectKeyID)
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, rootCertSubject))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, rootCertSubjectKeyID))
		require.Contains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, rootCertSerialNumber))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectAsText":"%s"`, rootCertSubjectText))
		require.Contains(t, string(out), fmt.Sprintf(`"vid":%d`, pkiDemoVid))

		// Approved cert must still be absent
		out, err = QueryX509Cert(rootCertSubject, rootCertSubjectKeyID)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, rootCertSubject))

		// Approved list must not contain this cert yet
		out, err = QueryAllX509Certs()
		require.NoError(t, err)
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, rootCertSubject))
		require.NotContains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, rootCertSerialNumber))

		// Revoked list must be empty of this cert
		out, err = QueryAllRevokedX509Certs()
		require.NoError(t, err)
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, rootCertSubject))

		// Approved root certs must not contain this cert yet
		out, err = QueryAllX509RootCerts()
		require.NoError(t, err)
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, rootCertSubject))

		// Revoked root certs must not contain this cert
		out, err = QueryAllRevokedX509RootCerts()
		require.NoError(t, err)
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, rootCertSubject))

		// Subject query must be empty
		out, err = QueryX509CertsBySubject(rootCertSubject)
		require.NoError(t, err)
		require.NotContains(t, string(out), fmt.Sprintf(`"%s"`, rootCertSubject))
		require.NotContains(t, string(out), fmt.Sprintf(`"%s"`, rootCertSubjectKeyID))
	})

	// ── Section 3: Approve Root Cert ───────────────────────────────────────────

	t.Run("ApproveRootCert_Trustee", func(t *testing.T) {
		txResult, err := ApproveAddX509RootCert(rootCertSubject, rootCertSubjectKeyID, alice)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Approved cert by subject+skid
		out, err := QueryX509Cert(rootCertSubject, rootCertSubjectKeyID)
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, rootCertSubject))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, rootCertSubjectKeyID))
		require.Contains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, rootCertSerialNumber))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectAsText":"%s"`, rootCertSubjectText))
		require.Contains(t, string(out), `"isRoot":true`)

		// Approved cert by skid only
		out, err = QueryX509CertBySKID(rootCertSubjectKeyID)
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, rootCertSubjectKeyID))
		require.Contains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, rootCertSerialNumber))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectAsText":"%s"`, rootCertSubjectText))

		// Proposed cert must now be gone
		out, err = QueryProposedX509RootCert(rootCertSubject, rootCertSubjectKeyID)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")

		out, err = QueryAllProposedX509RootCerts()
		require.NoError(t, err)
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, rootCertSubject))
		require.NotContains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, rootCertSerialNumber))

		// Approved list contains this cert
		out, err = QueryAllX509Certs()
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, rootCertSubject))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, rootCertSubjectKeyID))

		// Approved root certs contains this cert
		out, err = QueryAllX509RootCerts()
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, rootCertSubject))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, rootCertSubjectKeyID))

		// Revoked certs still empty
		out, err = QueryAllRevokedX509Certs()
		require.NoError(t, err)
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, rootCertSubject))
	})

	// ── Section 4: Add Intermediate Cert ───────────────────────────────────────

	t.Run("AddIntermediateCert", func(t *testing.T) {
		txResult, err := AddX509Cert(intermediateCertPath, vendorAccount, "--schemaVersion", "0")
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Intermediate cert by subject+skid
		out, err := QueryX509Cert(intermediateCertSubject, intermediateCertSubjectKeyID)
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, intermediateCertSubject))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, intermediateCertSubjectKeyID))
		require.Contains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, intermediateCertSerialNumber))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectAsText":"%s"`, intermediateCertSubjectText))
		require.Contains(t, string(out), `"schemaVersion":0`)
		require.Contains(t, string(out), `"approvals":[]`)
		require.Contains(t, string(out), `"isRoot":false`)

		// Intermediate cert by skid only
		out, err = QueryX509CertBySKID(intermediateCertSubjectKeyID)
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, intermediateCertSubjectKeyID))
		require.Contains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, intermediateCertSerialNumber))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectAsText":"%s"`, intermediateCertSubjectText))
		require.Contains(t, string(out), `"approvals":[]`)

		// All proposed root certs must not contain either cert
		out, err = QueryAllProposedX509RootCerts()
		require.NoError(t, err)
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, rootCertSubject))
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, intermediateCertSubject))

		// All approved certs — root and intermediate
		out, err = QueryAllX509Certs()
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, rootCertSubject))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, rootCertSubjectKeyID))
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, intermediateCertSubject))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, intermediateCertSubjectKeyID))

		// All approved root certs — root only, intermediate not root
		out, err = QueryAllX509RootCerts()
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, rootCertSubject))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, rootCertSubjectKeyID))
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, intermediateCertSubject))
		require.NotContains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, intermediateCertSubjectKeyID))
	})

	// ── Section 5: Add Leaf Cert ────────────────────────────────────────────────

	t.Run("AddLeafCert", func(t *testing.T) {
		txResult, err := AddX509Cert(leafCertPath, vendorAccount)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Leaf cert by subject+skid
		out, err := QueryX509Cert(leafCertSubject, leafCertSubjectKeyID)
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, leafCertSubject))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, leafCertSubjectKeyID))
		require.Contains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, leafCertSerialNumber))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectAsText":"%s"`, leafCertSubjectText))
		require.Contains(t, string(out), `"schemaVersion":0`)
		require.Contains(t, string(out), `"approvals":[]`)

		// Leaf cert by skid only
		out, err = QueryX509CertBySKID(leafCertSubjectKeyID)
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, leafCertSubjectKeyID))
		require.Contains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, leafCertSerialNumber))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectAsText":"%s"`, leafCertSubjectText))
		require.Contains(t, string(out), `"approvals":[]`)

		// All approved certs — all three
		out, err = QueryAllX509Certs()
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, rootCertSubject))
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, intermediateCertSubject))
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, leafCertSubject))

		// All approved root certs — root only
		out, err = QueryAllX509RootCerts()
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, rootCertSubject))
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, intermediateCertSubject))
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, leafCertSubject))

		// Subject queries
		out, err = QueryX509CertsBySubject(rootCertSubject)
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"%s"`, rootCertSubject))
		require.Contains(t, string(out), fmt.Sprintf(`"%s"`, rootCertSubjectKeyID))
		require.NotContains(t, string(out), fmt.Sprintf(`"%s"`, intermediateCertSubject))
		require.NotContains(t, string(out), fmt.Sprintf(`"%s"`, leafCertSubject))

		out, err = QueryX509CertsBySubject(leafCertSubject)
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"%s"`, leafCertSubject))
		require.Contains(t, string(out), fmt.Sprintf(`"%s"`, leafCertSubjectKeyID))
		require.NotContains(t, string(out), fmt.Sprintf(`"%s"`, rootCertSubject))
		require.NotContains(t, string(out), fmt.Sprintf(`"%s"`, intermediateCertSubject))

		out, err = QueryX509CertsBySubject(intermediateCertSubject)
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"%s"`, intermediateCertSubject))
		require.Contains(t, string(out), fmt.Sprintf(`"%s"`, intermediateCertSubjectKeyID))
		require.NotContains(t, string(out), fmt.Sprintf(`"%s"`, rootCertSubject))
		require.NotContains(t, string(out), fmt.Sprintf(`"%s"`, leafCertSubject))

		// No proposed-to-revoke entries
		out, err = QueryAllProposedRevokedX509RootCerts()
		require.NoError(t, err)
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, rootCertSubject))
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, intermediateCertSubject))
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, leafCertSubject))

		// No revoked certs yet
		out, err = QueryAllRevokedX509Certs()
		require.NoError(t, err)
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, rootCertSubject))
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, intermediateCertSubject))
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, leafCertSubject))

		// Child certs of root — only intermediate, not leaf
		out, err = QueryChildX509Certs(rootCertSubject, rootCertSubjectKeyID)
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, intermediateCertSubject))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, intermediateCertSubjectKeyID))
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, leafCertSubject))

		// Child certs of intermediate — only leaf
		out, err = QueryChildX509Certs(intermediateCertSubject, intermediateCertSubjectKeyID)
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, leafCertSubject))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, leafCertSubjectKeyID))
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, rootCertSubject))

		// Child certs of leaf — none
		out, err = QueryChildX509Certs(leafCertSubject, leafCertSubjectKeyID)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, rootCertSubject))
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, intermediateCertSubject))
	})

	// ── Section 6: Revoke Intermediate Cert ────────────────────────────────────

	t.Run("RevokeIntermediateCert_Unauthorized_Fails", func(t *testing.T) {
		// Non-vendor account cannot revoke
		txResult, err := RevokeX509Cert(intermediateCertSubject, intermediateCertSubjectKeyID, userAccount)
		require.NoError(t, err)
		require.Equal(t, uint32(4), txResult.Code)
		_, _ = utils.AwaitTxConfirmation(txResult.TxHash)

		// Vendor with different VID cannot revoke
		txResult, err = RevokeX509Cert(intermediateCertSubject, intermediateCertSubjectKeyID, vendorAccount65522)
		require.NoError(t, err)
		require.Equal(t, uint32(4), txResult.Code)
		_, _ = utils.AwaitTxConfirmation(txResult.TxHash)
	})

	t.Run("RevokeIntermediateCert", func(t *testing.T) {
		// Revoke intermediate without --revoke-child: leaf must survive
		txResult, err := RevokeX509Cert(intermediateCertSubject, intermediateCertSubjectKeyID, vendorAccount)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// No proposed-to-revoke entries (intermediate is not a root cert)
		out, err := QueryAllProposedRevokedX509RootCerts()
		require.NoError(t, err)
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, rootCertSubject))
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, intermediateCertSubject))
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, leafCertSubject))

		// All revoked — intermediate present, leaf NOT present, root NOT present
		out, err = QueryAllRevokedX509Certs()
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, intermediateCertSubject))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, intermediateCertSubjectKeyID))
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, leafCertSubject))
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, rootCertSubject))

		// All revoked root certs — none of the three
		out, err = QueryAllRevokedX509RootCerts()
		require.NoError(t, err)
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, rootCertSubject))
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, intermediateCertSubject))
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, leafCertSubject))

		// Revoked intermediate cert by subject+skid
		out, err = QueryRevokedX509Cert(intermediateCertSubject, intermediateCertSubjectKeyID)
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, intermediateCertSubject))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, intermediateCertSubjectKeyID))
		require.Contains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, intermediateCertSerialNumber))

		// Leaf cert is NOT in revoked (it was not revoked)
		out, err = QueryRevokedX509Cert(leafCertSubject, leafCertSubjectKeyID)
		require.NoError(t, err)
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, leafCertSubject))
		require.NotContains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, leafCertSubjectKeyID))

		// All approved — root and leaf present, intermediate gone
		out, err = QueryAllX509Certs()
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, rootCertSubject))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, rootCertSubjectKeyID))
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, leafCertSubject))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, leafCertSubjectKeyID))
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, intermediateCertSubject))
		require.NotContains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, intermediateCertSubjectKeyID))

		// Approved root certs — root only
		out, err = QueryAllX509RootCerts()
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, rootCertSubject))
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, intermediateCertSubject))
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, leafCertSubject))

		// Subject query — leaf still in approved
		out, err = QueryX509CertsBySubject(leafCertSubject)
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"%s"`, leafCertSubject))
		require.Contains(t, string(out), fmt.Sprintf(`"%s"`, leafCertSubjectKeyID))

		// Subject query — intermediate gone from approved
		out, err = QueryX509CertsBySubject(intermediateCertSubject)
		require.NoError(t, err)
		require.NotContains(t, string(out), fmt.Sprintf(`"%s"`, intermediateCertSubject))
		require.NotContains(t, string(out), fmt.Sprintf(`"%s"`, intermediateCertSubjectKeyID))

		// Intermediate approved cert — gone
		out, err = QueryX509Cert(intermediateCertSubject, intermediateCertSubjectKeyID)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, intermediateCertSubject))

		// Leaf approved cert — still present
		out, err = QueryX509Cert(leafCertSubject, leafCertSubjectKeyID)
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, leafCertSubject))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, leafCertSubjectKeyID))
		require.Contains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, leafCertSerialNumber))
	})

	// ── Section 7: Propose Revoke Root Cert ────────────────────────────────────

	t.Run("ProposeRevokeRootCert", func(t *testing.T) {
		txResult, err := ProposeRevokeX509RootCert(rootCertSubject, rootCertSubjectKeyID, jack)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Proposed-to-revoke contains root
		out, err := QueryProposedRevokedX509RootCert(rootCertSubject, rootCertSubjectKeyID)
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"%s"`, rootCertSubject))
		require.Contains(t, string(out), fmt.Sprintf(`"%s"`, rootCertSubjectKeyID))

		out, err = QueryAllProposedRevokedX509RootCerts()
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, rootCertSubject))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, rootCertSubjectKeyID))
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, intermediateCertSubject))
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, leafCertSubject))

		// All revoked — intermediate still there, root and leaf absent
		out, err = QueryAllRevokedX509Certs()
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, intermediateCertSubject))
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, leafCertSubject))
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, rootCertSubject))

		// All revoked root certs — root not yet revoked
		out, err = QueryAllRevokedX509RootCerts()
		require.NoError(t, err)
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, rootCertSubject))

		// Root cert still approved
		out, err = QueryAllX509Certs()
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, rootCertSubject))
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, leafCertSubject))
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, intermediateCertSubject))

		out, err = QueryAllX509RootCerts()
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, rootCertSubject))
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, intermediateCertSubject))
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, leafCertSubject))

		// Subject query — root still in approved
		out, err = QueryX509CertsBySubject(rootCertSubject)
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"%s"`, rootCertSubject))
		require.Contains(t, string(out), fmt.Sprintf(`"%s"`, rootCertSubjectKeyID))
		require.NotContains(t, string(out), fmt.Sprintf(`"%s"`, leafCertSubject))
		require.NotContains(t, string(out), fmt.Sprintf(`"%s"`, intermediateCertSubject))
	})

	// ── Section 8: Approve Revoke Root Cert ────────────────────────────────────

	t.Run("ApproveRevokeRootCert", func(t *testing.T) {
		txResult, err := ApproveRevokeX509RootCert(rootCertSubject, rootCertSubjectKeyID, alice)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Proposed-to-revoke list empty
		out, err := QueryAllProposedRevokedX509RootCerts()
		require.NoError(t, err)
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, rootCertSubject))
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, intermediateCertSubject))
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, leafCertSubject))

		// All revoked — root and intermediate present, leaf absent
		out, err = QueryAllRevokedX509Certs()
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, rootCertSubject))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, rootCertSubjectKeyID))
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, intermediateCertSubject))
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, leafCertSubject))

		// All revoked root certs — root present, others absent
		out, err = QueryAllRevokedX509RootCerts()
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, rootCertSubject))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, rootCertSubjectKeyID))
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, intermediateCertSubject))
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, leafCertSubject))

		// Revoked root cert by subject+skid
		out, err = QueryRevokedX509Cert(rootCertSubject, rootCertSubjectKeyID)
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, rootCertSubject))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, rootCertSubjectKeyID))
		require.Contains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, rootCertSerialNumber))
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, intermediateCertSubject))
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, leafCertSubject))

		// All approved — root and intermediate gone, leaf still present
		out, err = QueryAllX509Certs()
		require.NoError(t, err)
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, rootCertSubject))
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, intermediateCertSubject))
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, leafCertSubject))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, leafCertSubjectKeyID))

		// All approved root certs — empty
		out, err = QueryAllX509RootCerts()
		require.NoError(t, err)
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, rootCertSubject))
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, intermediateCertSubject))
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, leafCertSubject))

		// Intermediate approved — gone
		out, err = QueryX509Cert(intermediateCertSubject, intermediateCertSubjectKeyID)
		require.NoError(t, err)
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, intermediateCertSubject))

		// Leaf approved — still present
		out, err = QueryX509Cert(leafCertSubject, leafCertSubjectKeyID)
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, leafCertSubject))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, leafCertSubjectKeyID))
		require.Contains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, leafCertSerialNumber))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectAsText":"%s"`, leafCertSubjectText))

		// Root approved cert — gone
		out, err = QueryX509Cert(rootCertSubject, rootCertSubjectKeyID)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, rootCertSubject))

		// Subject query — root gone from approved
		out, err = QueryX509CertsBySubject(rootCertSubject)
		require.NoError(t, err)
		require.NotContains(t, string(out), fmt.Sprintf(`"%s"`, rootCertSubject))
		require.NotContains(t, string(out), fmt.Sprintf(`"%s"`, rootCertSubjectKeyID))
		require.NotContains(t, string(out), fmt.Sprintf(`"%s"`, intermediateCertSubject))
		require.NotContains(t, string(out), fmt.Sprintf(`"%s"`, leafCertSubject))
	})

	// ── Section 9: Google Cert Query All Empty ──────────────────────────────────

	t.Run("GoogleCert_QueryAllEmpty", func(t *testing.T) {
		out, err := QueryX509Cert(googleCertSubject, googleCertSubjectKeyID)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, googleCertSubject))
		require.NotContains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, googleCertSerialNumber))

		out, err = QueryAllX509Certs()
		require.NoError(t, err)
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, googleCertSubject))

		out, err = QueryProposedX509RootCert(googleCertSubject, googleCertSubjectKeyID)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, googleCertSubject))

		out, err = QueryAllProposedX509RootCerts()
		require.NoError(t, err)
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, googleCertSubject))

		out, err = QueryRevokedX509Cert(googleCertSubject, googleCertSubjectKeyID)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")
		require.NotContains(t, string(out), fmt.Sprintf(`"%s"`, googleCertSubject))

		out, err = QueryAllRevokedX509Certs()
		require.NoError(t, err)
		require.NotContains(t, string(out), fmt.Sprintf(`"%s"`, googleCertSubject))

		out, err = QueryX509CertsBySubject(googleCertSubject)
		require.NoError(t, err)
		require.NotContains(t, string(out), fmt.Sprintf(`"%s"`, googleCertSubject))

		out, err = QueryAllX509RootCerts()
		require.NoError(t, err)
		require.NotContains(t, string(out), fmt.Sprintf(`"%s"`, googleCertSubject))

		out, err = QueryAllRevokedX509RootCerts()
		require.NoError(t, err)
		require.NotContains(t, string(out), fmt.Sprintf(`"%s"`, googleCertSubject))

		out, err = QueryProposedRevokedX509RootCert(googleCertSubject, googleCertSubjectKeyID)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")
		require.NotContains(t, string(out), fmt.Sprintf(`"%s"`, googleCertSubject))

		out, err = QueryAllProposedRevokedX509RootCerts()
		require.NoError(t, err)
		require.NotContains(t, string(out), fmt.Sprintf(`"%s"`, googleCertSubject))

		out, err = QueryChildX509Certs(googleCertSubject, googleCertSubjectKeyID)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")
		require.NotContains(t, string(out), fmt.Sprintf(`"%s"`, googleCertSubject))
	})

	// ── Section 10: Propose Google Root Cert ───────────────────────────────────

	t.Run("ProposeGoogleRootCert", func(t *testing.T) {
		// Non-trustee fails
		txResult, err := ProposeAddX509RootCert(googleCertPath, userAccount,
			"--vid", fmt.Sprintf("%d", googleCertVid),
		)
		require.NoError(t, err)
		require.NotEqual(t, uint32(0), txResult.Code)
		_, _ = utils.AwaitTxConfirmation(txResult.TxHash)

		// Trustee proposes
		txResult, err = ProposeAddX509RootCert(googleCertPath, jack,
			"--vid", fmt.Sprintf("%d", googleCertVid),
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Proposed cert present
		out, err := QueryProposedX509RootCert(googleCertSubject, googleCertSubjectKeyID)
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, googleCertSubject))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, googleCertSubjectKeyID))
		require.Contains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, googleCertSerialNumber))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectAsText":"%s"`, googleCertSubjectText))
		require.Contains(t, string(out), fmt.Sprintf(`"vid":%d`, googleCertVid))
		require.Contains(t, string(out), `"schemaVersion":0`)

		// Approved cert must be absent
		out, err = QueryAllX509Certs()
		require.NoError(t, err)
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, googleCertSubject))

		out, err = QueryX509Cert(googleCertSubject, googleCertSubjectKeyID)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, googleCertSubject))

		out, err = QueryAllRevokedX509Certs()
		require.NoError(t, err)
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, googleCertSubject))

		out, err = QueryAllX509RootCerts()
		require.NoError(t, err)
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, googleCertSubject))

		out, err = QueryAllRevokedX509RootCerts()
		require.NoError(t, err)
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, googleCertSubject))

		out, err = QueryX509CertsBySubject(googleCertSubject)
		require.NoError(t, err)
		require.NotContains(t, string(out), fmt.Sprintf(`"%s"`, googleCertSubject))
		require.NotContains(t, string(out), fmt.Sprintf(`"%s"`, googleCertSubjectKeyID))
	})

	// ── Section 11: Approve Google Root Cert ───────────────────────────────────

	t.Run("ApproveGoogleRootCert", func(t *testing.T) {
		// Still proposed, not yet approved
		out, err := QueryProposedX509RootCert(googleCertSubject, googleCertSubjectKeyID)
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, googleCertSubject))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, googleCertSubjectKeyID))
		require.Contains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, googleCertSerialNumber))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectAsText":"%s"`, googleCertSubjectText))
		require.Contains(t, string(out), fmt.Sprintf(`"vid":%d`, googleCertVid))

		out, err = QueryAllX509Certs()
		require.NoError(t, err)
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, googleCertSubject))

		// Alice approves
		txResult, err := ApproveAddX509RootCert(googleCertSubject, googleCertSubjectKeyID, alice)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Now approved
		out, err = QueryX509Cert(googleCertSubject, googleCertSubjectKeyID)
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, googleCertSubject))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, googleCertSubjectKeyID))
		require.Contains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, googleCertSerialNumber))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectAsText":"%s"`, googleCertSubjectText))

		out, err = QueryX509CertBySKID(googleCertSubjectKeyID)
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, googleCertSubjectKeyID))
		require.Contains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, googleCertSerialNumber))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectAsText":"%s"`, googleCertSubjectText))

		out, err = QueryAllProposedX509RootCerts()
		require.NoError(t, err)
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, googleCertSubject))

		out, err = QueryAllX509Certs()
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, googleCertSubject))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, googleCertSubjectKeyID))

		out, err = QueryAllX509RootCerts()
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, googleCertSubject))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, googleCertSubjectKeyID))

		out, err = QueryAllRevokedX509Certs()
		require.NoError(t, err)
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, googleCertSubject))
	})

	// ── Section 12: Propose Revoke Google Root Cert ─────────────────────────────

	t.Run("ProposeRevokeGoogleRootCert", func(t *testing.T) {
		txResult, err := ProposeRevokeX509RootCert(googleCertSubject, googleCertSubjectKeyID, jack)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		out, err := QueryProposedRevokedX509RootCert(googleCertSubject, googleCertSubjectKeyID)
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"%s"`, googleCertSubject))
		require.Contains(t, string(out), fmt.Sprintf(`"%s"`, googleCertSubjectKeyID))

		out, err = QueryAllProposedRevokedX509RootCerts()
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, googleCertSubject))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, googleCertSubjectKeyID))

		out, err = QueryAllRevokedX509Certs()
		require.NoError(t, err)
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, googleCertSubject))

		out, err = QueryAllRevokedX509RootCerts()
		require.NoError(t, err)
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, googleCertSubject))

		out, err = QueryAllX509Certs()
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, googleCertSubject))

		out, err = QueryAllX509RootCerts()
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, googleCertSubject))

		out, err = QueryX509CertsBySubject(googleCertSubject)
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"%s"`, googleCertSubject))
		require.Contains(t, string(out), fmt.Sprintf(`"%s"`, googleCertSubjectKeyID))
	})

	// ── Section 13: Approve Revoke Google Root Cert ─────────────────────────────

	t.Run("ApproveRevokeGoogleRootCert", func(t *testing.T) {
		txResult, err := ApproveRevokeX509RootCert(googleCertSubject, googleCertSubjectKeyID, alice)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		out, err := QueryAllProposedRevokedX509RootCerts()
		require.NoError(t, err)
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, googleCertSubject))

		out, err = QueryAllRevokedX509Certs()
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, googleCertSubject))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, googleCertSubjectKeyID))

		out, err = QueryAllRevokedX509RootCerts()
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, googleCertSubject))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, googleCertSubjectKeyID))

		out, err = QueryRevokedX509Cert(googleCertSubject, googleCertSubjectKeyID)
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, googleCertSubject))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, googleCertSubjectKeyID))
		require.Contains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, googleCertSerialNumber))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectAsText":"%s"`, googleCertSubjectText))

		out, err = QueryAllX509Certs()
		require.NoError(t, err)
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, googleCertSubject))

		out, err = QueryAllX509RootCerts()
		require.NoError(t, err)
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, googleCertSubject))

		out, err = QueryX509Cert(googleCertSubject, googleCertSubjectKeyID)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, googleCertSubject))

		out, err = QueryX509CertsBySubject(googleCertSubject)
		require.NoError(t, err)
		require.NotContains(t, string(out), fmt.Sprintf(`"%s"`, googleCertSubject))
		require.NotContains(t, string(out), fmt.Sprintf(`"%s"`, googleCertSubjectKeyID))
	})

	// ── Section 14: Propose and Reject Test Cert (single trustee) ──────────────

	t.Run("ProposeAndRejectTestCert_SingleTrustee", func(t *testing.T) {
		// Jack proposes
		txResult, err := ProposeAddX509RootCert(testCertPath, jack,
			"--vid", fmt.Sprintf("%d", testCertVid),
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Jack rejects (same trustee who proposed)
		txResult, err = RejectAddX509RootCert(testCertSubject, testCertSubjectKeyID, jack)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Cert gone from proposed
		out, err := QueryProposedX509RootCert(testCertSubject, testCertSubjectKeyID)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")

		// Cert not in rejected (single-trustee reject doesn't reach quorum alone)
		out, err = QueryRejectedX509RootCert(testCertSubject, testCertSubjectKeyID)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")

		// Cert not in approved
		out, err = QueryX509Cert(testCertSubject, testCertSubjectKeyID)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")

		out, err = QueryX509CertBySKID(testCertSubjectKeyID)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")
	})

	// ── Section 15: Propose Test Root Cert ─────────────────────────────────────

	t.Run("ProposeTestRootCert", func(t *testing.T) {
		// Non-trustee fails
		txResult, err := ProposeAddX509RootCert(testCertPath, userAccount,
			"--vid", fmt.Sprintf("%d", testCertVid),
		)
		require.NoError(t, err)
		require.NotEqual(t, uint32(0), txResult.Code)
		_, _ = utils.AwaitTxConfirmation(txResult.TxHash)

		// Trustee proposes
		txResult, err = ProposeAddX509RootCert(testCertPath, jack,
			"--vid", fmt.Sprintf("%d", testCertVid),
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Proposed cert present
		out, err := QueryProposedX509RootCert(testCertSubject, testCertSubjectKeyID)
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, testCertSubject))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, testCertSubjectKeyID))
		require.Contains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, testCertSerialNumber))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectAsText":"%s"`, testCertSubjectText))
		require.Contains(t, string(out), fmt.Sprintf(`"vid":%d`, testCertVid))
	})

	// ── Section 16: Reject Test Root Cert (multi-trustee scenario) ─────────────

	t.Run("RejectTestRootCert_MultiTrustee", func(t *testing.T) {
		// Add a 4th trustee to raise approval/rejection thresholds
		newTrustee1 := cliputils.CreateAccount(t, "Trustee")
		newTrustee1AddrOut, err := utils.ExecuteCLI("keys", "show", newTrustee1, "-a", "--keyring-backend", "test")
		require.NoError(t, err)
		newTrustee1Addr := strings.TrimSpace(string(newTrustee1AddrOut))

		// Bob approves the test cert (2 approvals with 4 trustees — not enough, need 3)
		txResult, err := ApproveAddX509RootCert(testCertSubject, testCertSubjectKeyID, bob)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Jack rejects (removes Jack's approval, adds rejection: 1 approval Bob, 1 rejection Jack)
		txResult, err = RejectAddX509RootCert(testCertSubject, testCertSubjectKeyID, jack)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Jack can approve even after rejecting (switches back: 2 approvals Jack+Bob)
		txResult, err = ApproveAddX509RootCert(testCertSubject, testCertSubjectKeyID, jack)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Jack rejects again (1 approval Bob, 1 rejection Jack)
		txResult, err = RejectAddX509RootCert(testCertSubject, testCertSubjectKeyID, jack)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Jack cannot reject twice in a row
		txResult, err = RejectAddX509RootCert(testCertSubject, testCertSubjectKeyID, jack)
		require.NoError(t, err)
		require.NotEqual(t, uint32(0), txResult.Code)
		_, _ = utils.AwaitTxConfirmation(txResult.TxHash)

		// Cert still in proposed — not yet at rejection quorum
		out, err := QueryProposedX509RootCert(testCertSubject, testCertSubjectKeyID)
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, testCertSubject))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, testCertSubjectKeyID))
		require.Contains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, testCertSerialNumber))

		// Not yet in all-rejected list
		out, err = QueryAllRejectedX509RootCerts()
		require.NoError(t, err)
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, testCertSubject))

		// Alice rejects — now 2 rejections with 4 trustees = rejection quorum reached
		txResult, err = RejectAddX509RootCert(testCertSubject, testCertSubjectKeyID, alice)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Revoke new_trustee1 (back to 3 trustees)
		txResult, err = utils.ExecuteTx("tx", "auth", "propose-revoke-account",
			"--address", newTrustee1Addr,
			"--from", alice,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code, txResult.RawLog)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		txResult, err = utils.ExecuteTx("tx", "auth", "approve-revoke-account",
			"--address", newTrustee1Addr,
			"--from", bob,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code, txResult.RawLog)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		txResult, err = utils.ExecuteTx("tx", "auth", "approve-revoke-account",
			"--address", newTrustee1Addr,
			"--from", jack,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code, txResult.RawLog)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Cert is now rejected
		out, err = QueryRejectedX509RootCert(testCertSubject, testCertSubjectKeyID)
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, testCertSubject))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, testCertSubjectKeyID))
		require.Contains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, testCertSerialNumber))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectAsText":"%s"`, testCertSubjectText))
		require.Contains(t, string(out), `"schemaVersion":0`)

		// No longer in proposed
		out, err = QueryAllProposedX509RootCerts()
		require.NoError(t, err)
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, testCertSubject))

		// Not in approved root certs
		out, err = QueryAllX509RootCerts()
		require.NoError(t, err)
		require.NotContains(t, string(out), fmt.Sprintf(`"subject":"%s"`, testCertSubject))

		// In all-rejected list
		out, err = QueryAllRejectedX509RootCerts()
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, testCertSubject))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, testCertSubjectKeyID))
	})

	// ── Section 17: Propose Test Root Cert Again ────────────────────────────────

	t.Run("ProposeTestRootCertAgain", func(t *testing.T) {
		// Non-trustee fails
		txResult, err := ProposeAddX509RootCert(testCertPath, userAccount,
			"--vid", fmt.Sprintf("%d", testCertVid),
		)
		require.NoError(t, err)
		require.NotEqual(t, uint32(0), txResult.Code)
		_, _ = utils.AwaitTxConfirmation(txResult.TxHash)

		// Trustee proposes again
		txResult, err = ProposeAddX509RootCert(testCertPath, jack,
			"--vid", fmt.Sprintf("%d", testCertVid),
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Proposed cert present again
		out, err := QueryProposedX509RootCert(testCertSubject, testCertSubjectKeyID)
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, testCertSubject))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, testCertSubjectKeyID))
		require.Contains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, testCertSerialNumber))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectAsText":"%s"`, testCertSubjectText))
		require.Contains(t, string(out), fmt.Sprintf(`"vid":%d`, testCertVid))

		// No longer in rejected (moving from rejected back to proposed clears it)
		out, err = QueryRejectedX509RootCert(testCertSubject, testCertSubjectKeyID)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")
	})

	// ── Section 18: Approve Test Root Cert ─────────────────────────────────────

	t.Run("ApproveTestRootCert", func(t *testing.T) {
		// Alice approves — with 3 trustees (jack+alice+bob), 2 approvals = quorum
		txResult, err := ApproveAddX509RootCert(testCertSubject, testCertSubjectKeyID, alice)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Approved by subject+skid
		out, err := QueryX509Cert(testCertSubject, testCertSubjectKeyID)
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subject":"%s"`, testCertSubject))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, testCertSubjectKeyID))
		require.Contains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, testCertSerialNumber))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectAsText":"%s"`, testCertSubjectText))
		require.Contains(t, string(out), fmt.Sprintf(`"vid":%d`, testCertVid))

		// Approved by skid only
		out, err = QueryX509CertBySKID(testCertSubjectKeyID)
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"subjectKeyId":"%s"`, testCertSubjectKeyID))
		require.Contains(t, string(out), fmt.Sprintf(`"serialNumber":"%s"`, testCertSerialNumber))
		require.Contains(t, string(out), fmt.Sprintf(`"subjectAsText":"%s"`, testCertSubjectText))
		require.Contains(t, string(out), fmt.Sprintf(`"vid":%d`, testCertVid))
	})

	_, _ = bob, alice
}
