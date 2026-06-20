package pki

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/cli/dclauth"
	cliputils "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/cli/utils"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/utils"
	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

const (
	rootCertPath         = "../../constants/root_cert"
	intermediateCertPath = "../../constants/intermediate_cert"
	leafCertPath         = "../../constants/leaf_cert"

	rootCertSubject      = "MDQxCzAJBgNVBAYTAkFVMRMwEQYDVQQIEwpzb21lLXN0YXRlMRAwDgYDVQQKEwdyb290LWNh"
	rootCertSubjectKeyID = "DF:4E:AF:B0:8C:9C:37:78:1A:E7:53:12:CA:E4:78:6B:48:1E:AF:B0"
	rootCertSerialNumber = "81311506302208030248766861785118937702312370677"
	rootCertSubjectText  = "O=root-ca,ST=some-state,C=AU"

	intermediateCertSubject      = "MDwxCzAJBgNVBAYTAkFVMRMwEQYDVQQIEwpzb21lLXN0YXRlMRgwFgYDVQQKEw9pbnRlcm1lZGlhdGUtY2E="
	intermediateCertSubjectKeyID = "1B:73:2A:91:34:46:8A:90:2A:87:19:91:E4:BD:8F:69:3A:F9:04:77"
	intermediateCertSerialNumber = "486736128900935106101503663840421220667833341899"
	intermediateCertSubjectText  = "O=intermediate-ca,ST=some-state,C=AU"

	leafCertSubject      = "MDExCzAJBgNVBAYTAkFVMRMwEQYDVQQIEwpzb21lLXN0YXRlMQ0wCwYDVQQKEwRsZWFm"
	leafCertSubjectKeyID = "2A:31:8D:39:6E:50:DA:96:DF:95:C5:98:83:68:F0:58:B2:15:B3:3A"
	leafCertSerialNumber = "409691117370409054634487600348183880852961428328"
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

	t.Run("QueryAllEmpty", pkiDemoQueryAllEmpty)
	t.Run("ProposeRootCert_NotTrustee_Fails", func(t *testing.T) {
		pkiDemoProposeRootCertNotTrustee(t, userAccount)
	})
	t.Run("ProposeRootCert_Trustee_Succeeds", func(t *testing.T) {
		pkiDemoProposeRootCertTrustee(t, jack)
	})
	t.Run("ApproveRootCert_Trustee", func(t *testing.T) {
		pkiDemoApproveRootCert(t, alice)
	})
	t.Run("AddIntermediateCert", func(t *testing.T) {
		pkiDemoAddIntermediateCert(t, vendorAccount)
	})
	t.Run("AddLeafCert", func(t *testing.T) {
		pkiDemoAddLeafCert(t, vendorAccount)
	})
	t.Run("RevokeIntermediateCert_Unauthorized_Fails", func(t *testing.T) {
		pkiDemoRevokeIntermediateCertUnauthorized(t, userAccount, vendorAccount65522)
	})
	t.Run("RevokeIntermediateCert", func(t *testing.T) {
		pkiDemoRevokeIntermediateCert(t, vendorAccount)
	})
	t.Run("ProposeRevokeRootCert", func(t *testing.T) {
		pkiDemoProposeRevokeRootCert(t, jack)
	})
	t.Run("ApproveRevokeRootCert", func(t *testing.T) {
		pkiDemoApproveRevokeRootCert(t, alice)
	})
	t.Run("GoogleCert_QueryAllEmpty", pkiDemoGoogleCertQueryAllEmpty)
	t.Run("ProposeGoogleRootCert", func(t *testing.T) {
		pkiDemoProposeGoogleRootCert(t, jack, userAccount)
	})
	t.Run("ApproveGoogleRootCert", func(t *testing.T) {
		pkiDemoApproveGoogleRootCert(t, alice)
	})
	t.Run("ProposeRevokeGoogleRootCert", func(t *testing.T) {
		pkiDemoProposeRevokeGoogleRootCert(t, jack)
	})
	t.Run("ApproveRevokeGoogleRootCert", func(t *testing.T) {
		pkiDemoApproveRevokeGoogleRootCert(t, alice)
	})
	t.Run("ProposeAndRejectTestCert_SingleTrustee", func(t *testing.T) {
		pkiDemoProposeAndRejectTestCertSingleTrustee(t, jack)
	})
	t.Run("ProposeTestRootCert", func(t *testing.T) {
		pkiDemoProposeTestRootCert(t, jack, userAccount)
	})
	t.Run("RejectTestRootCert_MultiTrustee", func(t *testing.T) {
		pkiDemoRejectTestRootCertMultiTrustee(t, jack, alice, bob)
	})
	t.Run("ProposeTestRootCertAgain", func(t *testing.T) {
		pkiDemoProposeTestRootCertAgain(t, jack, userAccount)
	})
	t.Run("ApproveTestRootCert", func(t *testing.T) {
		pkiDemoApproveTestRootCert(t, alice)
	})
}

// ── Section 1: Query All Empty ──────────────────────────────────────────────

func pkiDemoQueryAllEmpty(t *testing.T) {
	t.Helper()
	cert, err := GetX509Cert(rootCertSubject, rootCertSubjectKeyID)
	require.NoError(t, err)
	require.Nil(t, cert)

	all, err := GetAllX509Certs()
	require.NoError(t, err)
	require.False(t, containsApprovedCertSerial(all, rootCertSerialNumber))

	proposed, err := GetProposedX509RootCert(rootCertSubject, rootCertSubjectKeyID)
	require.NoError(t, err)
	require.Nil(t, proposed)

	allProposed, err := GetAllProposedX509RootCerts()
	require.NoError(t, err)
	for _, p := range allProposed {
		require.NotEqual(t, rootCertSubject, p.Subject)
	}

	revoked, err := GetRevokedX509Cert(rootCertSubject, rootCertSubjectKeyID)
	require.NoError(t, err)
	require.Nil(t, revoked)

	allRevoked, err := GetAllRevokedX509Certs()
	require.NoError(t, err)
	require.False(t, containsRevokedCertSerial(allRevoked, rootCertSerialNumber))

	bySubject, err := GetX509CertsBySubject(rootCertSubject)
	require.NoError(t, err)
	require.Nil(t, bySubject)

	bySkid, err := GetX509CertBySKID(rootCertSubjectKeyID)
	require.NoError(t, err)
	require.Nil(t, bySkid)

	allRoots, err := GetAllX509RootCerts()
	require.NoError(t, err)
	if allRoots != nil {
		for _, id := range allRoots.Certs {
			require.NotEqual(t, rootCertSubject, id.Subject)
		}
	}

	allRevokedRoots, err := GetAllRevokedX509RootCerts()
	require.NoError(t, err)
	if allRevokedRoots != nil {
		for _, id := range allRevokedRoots.Certs {
			require.NotEqual(t, rootCertSubject, id.Subject)
		}
	}

	proposedRev, err := GetProposedRevokedX509RootCert(rootCertSubject, rootCertSubjectKeyID)
	require.NoError(t, err)
	require.Nil(t, proposedRev)

	allProposedRev, err := GetAllProposedRevokedX509RootCerts()
	require.NoError(t, err)
	for _, p := range allProposedRev {
		require.NotEqual(t, rootCertSubject, p.Subject)
	}

	children, err := GetChildX509Certs(rootCertSubject, rootCertSubjectKeyID)
	require.NoError(t, err)
	require.Nil(t, children)
}

// ── Section 2: Propose Root Cert ───────────────────────────────────────────

func pkiDemoProposeRootCertNotTrustee(t *testing.T, userAccount string) {
	t.Helper()
	txResult, err := ProposeAddX509RootCert(rootCertPath, userAccount, X509ProposeOpts{VID: pkiDemoVid})
	require.NoError(t, err)
	require.NotEqual(t, uint32(0), txResult.Code)
	_, _ = utils.AwaitTxConfirmation(txResult.TxHash)
}

func pkiDemoProposeRootCertTrustee(t *testing.T, jack string) {
	t.Helper()
	txResult, err := ProposeAddX509RootCert(rootCertPath, jack, X509ProposeOpts{
		VID:           pkiDemoVid,
		SchemaVersion: "0",
	})
	require.NoError(t, err)
	require.Equal(t, uint32(0), txResult.Code)
	_, err = utils.AwaitTxConfirmation(txResult.TxHash)
	require.NoError(t, err)

	// All proposed root certs — contains our cert.
	allProposed, err := GetAllProposedX509RootCerts()
	require.NoError(t, err)
	var proposed *pkitypes.ProposedCertificate
	for i := range allProposed {
		if allProposed[i].Subject == rootCertSubject {
			proposed = &allProposed[i]

			break
		}
	}
	require.NotNil(t, proposed)
	require.Equal(t, rootCertSubjectKeyID, proposed.SubjectKeyId)

	// Proposed cert by subject+skid.
	proposed2, err := GetProposedX509RootCert(rootCertSubject, rootCertSubjectKeyID)
	require.NoError(t, err)
	require.NotNil(t, proposed2)
	require.Equal(t, rootCertSubject, proposed2.Subject)
	require.Equal(t, rootCertSubjectKeyID, proposed2.SubjectKeyId)
	require.Equal(t, rootCertSerialNumber, proposed2.SerialNumber)
	require.Equal(t, rootCertSubjectText, proposed2.SubjectAsText)
	require.Equal(t, int32(pkiDemoVid), proposed2.Vid)

	// Approved cert must still be absent.
	cert, err := GetX509Cert(rootCertSubject, rootCertSubjectKeyID)
	require.NoError(t, err)
	require.Nil(t, cert)

	// Approved list must not contain this cert yet.
	all, err := GetAllX509Certs()
	require.NoError(t, err)
	require.False(t, containsApprovedCertSerial(all, rootCertSerialNumber))

	// Revoked list must not contain this cert.
	allRevoked, err := GetAllRevokedX509Certs()
	require.NoError(t, err)
	require.False(t, containsRevokedCertSerial(allRevoked, rootCertSerialNumber))

	// Approved root certs must not contain this cert.
	allRoots, err := GetAllX509RootCerts()
	require.NoError(t, err)
	if allRoots != nil {
		for _, id := range allRoots.Certs {
			require.NotEqual(t, rootCertSubject, id.Subject)
		}
	}

	// Revoked root certs must not contain this cert.
	allRevokedRoots, err := GetAllRevokedX509RootCerts()
	require.NoError(t, err)
	if allRevokedRoots != nil {
		for _, id := range allRevokedRoots.Certs {
			require.NotEqual(t, rootCertSubject, id.Subject)
		}
	}

	// Subject query must be empty.
	bySubject, err := GetX509CertsBySubject(rootCertSubject)
	require.NoError(t, err)
	require.Nil(t, bySubject)
}

// ── Section 3: Approve Root Cert ───────────────────────────────────────────

func pkiDemoApproveRootCert(t *testing.T, alice string) {
	t.Helper()
	txResult, err := ApproveAddX509RootCert(rootCertSubject, rootCertSubjectKeyID, alice)
	require.NoError(t, err)
	require.Equal(t, uint32(0), txResult.Code)
	_, err = utils.AwaitTxConfirmation(txResult.TxHash)
	require.NoError(t, err)

	// Approved cert by subject+skid.
	cert, err := GetX509Cert(rootCertSubject, rootCertSubjectKeyID)
	require.NoError(t, err)
	require.NotNil(t, cert)
	require.Equal(t, rootCertSubject, cert.Subject)
	require.Equal(t, rootCertSubjectKeyID, cert.SubjectKeyId)
	c := findCertBySerial(cert.Certs, rootCertSerialNumber)
	require.NotNil(t, c)
	require.Equal(t, rootCertSubjectText, c.SubjectAsText)
	require.True(t, c.IsRoot)

	// Approved cert by skid only.
	bySkid, err := GetX509CertBySKID(rootCertSubjectKeyID)
	require.NoError(t, err)
	require.NotNil(t, bySkid)
	require.Equal(t, rootCertSubjectKeyID, bySkid.SubjectKeyId)
	c2 := findCertBySerial(bySkid.Certs, rootCertSerialNumber)
	require.NotNil(t, c2)
	require.Equal(t, rootCertSubjectText, c2.SubjectAsText)

	// Proposed cert must now be gone.
	proposed, err := GetProposedX509RootCert(rootCertSubject, rootCertSubjectKeyID)
	require.NoError(t, err)
	require.Nil(t, proposed)

	allProposed, err := GetAllProposedX509RootCerts()
	require.NoError(t, err)
	for _, p := range allProposed {
		require.NotEqual(t, rootCertSubject, p.Subject)
	}

	// Approved list contains this cert.
	all, err := GetAllX509Certs()
	require.NoError(t, err)
	require.True(t, containsApprovedCertSerial(all, rootCertSerialNumber))

	// Approved root certs contain this cert.
	allRoots, err := GetAllX509RootCerts()
	require.NoError(t, err)
	require.NotNil(t, allRoots)
	var foundRoot bool
	for _, id := range allRoots.Certs {
		if id.Subject == rootCertSubject && id.SubjectKeyId == rootCertSubjectKeyID {
			foundRoot = true

			break
		}
	}
	require.True(t, foundRoot)

	// Revoked certs still empty for this cert.
	allRevoked, err := GetAllRevokedX509Certs()
	require.NoError(t, err)
	require.False(t, containsRevokedCertSerial(allRevoked, rootCertSerialNumber))
}

// ── Section 4: Add Intermediate Cert ───────────────────────────────────────

func pkiDemoAddIntermediateCert(t *testing.T, vendorAccount string) {
	t.Helper()
	txResult, err := AddX509Cert(intermediateCertPath, vendorAccount, X509ProposeOpts{SchemaVersion: "0"})
	require.NoError(t, err)
	require.Equal(t, uint32(0), txResult.Code)
	_, err = utils.AwaitTxConfirmation(txResult.TxHash)
	require.NoError(t, err)

	// Intermediate cert by subject+skid.
	cert, err := GetX509Cert(intermediateCertSubject, intermediateCertSubjectKeyID)
	require.NoError(t, err)
	require.NotNil(t, cert)
	require.Equal(t, intermediateCertSubject, cert.Subject)
	require.Equal(t, intermediateCertSubjectKeyID, cert.SubjectKeyId)
	c := findCertBySerial(cert.Certs, intermediateCertSerialNumber)
	require.NotNil(t, c)
	require.Equal(t, intermediateCertSubjectText, c.SubjectAsText)
	require.False(t, c.IsRoot)

	// Intermediate cert by skid only.
	bySkid, err := GetX509CertBySKID(intermediateCertSubjectKeyID)
	require.NoError(t, err)
	require.NotNil(t, bySkid)
	c2 := findCertBySerial(bySkid.Certs, intermediateCertSerialNumber)
	require.NotNil(t, c2)
	require.Equal(t, intermediateCertSubjectText, c2.SubjectAsText)

	// All proposed root certs must not contain either cert.
	allProposed, err := GetAllProposedX509RootCerts()
	require.NoError(t, err)
	for _, p := range allProposed {
		require.NotEqual(t, rootCertSubject, p.Subject)
		require.NotEqual(t, intermediateCertSubject, p.Subject)
	}

	// All approved certs — root and intermediate.
	all, err := GetAllX509Certs()
	require.NoError(t, err)
	require.True(t, containsApprovedCertSerial(all, rootCertSerialNumber))
	require.True(t, containsApprovedCertSerial(all, intermediateCertSerialNumber))

	// All approved root certs — root only, intermediate not root.
	allRoots, err := GetAllX509RootCerts()
	require.NoError(t, err)
	require.NotNil(t, allRoots)
	foundRoot, foundInterm := false, false
	for _, id := range allRoots.Certs {
		if id.Subject == rootCertSubject && id.SubjectKeyId == rootCertSubjectKeyID {
			foundRoot = true
		}
		if id.Subject == intermediateCertSubject {
			foundInterm = true
		}
	}
	require.True(t, foundRoot)
	require.False(t, foundInterm)
}

// ── Section 5: Add Leaf Cert ────────────────────────────────────────────────

func pkiDemoAddLeafCert(t *testing.T, vendorAccount string) {
	t.Helper()
	txResult, err := AddX509Cert(leafCertPath, vendorAccount)
	require.NoError(t, err)
	require.Equal(t, uint32(0), txResult.Code)
	_, err = utils.AwaitTxConfirmation(txResult.TxHash)
	require.NoError(t, err)

	// Leaf cert by subject+skid.
	cert, err := GetX509Cert(leafCertSubject, leafCertSubjectKeyID)
	require.NoError(t, err)
	require.NotNil(t, cert)
	require.Equal(t, leafCertSubject, cert.Subject)
	c := findCertBySerial(cert.Certs, leafCertSerialNumber)
	require.NotNil(t, c)
	require.Equal(t, leafCertSubjectText, c.SubjectAsText)

	// Leaf cert by skid only.
	bySkid, err := GetX509CertBySKID(leafCertSubjectKeyID)
	require.NoError(t, err)
	require.NotNil(t, bySkid)
	require.NotNil(t, findCertBySerial(bySkid.Certs, leafCertSerialNumber))

	// All approved certs — all three.
	all, err := GetAllX509Certs()
	require.NoError(t, err)
	require.True(t, containsApprovedCertSerial(all, rootCertSerialNumber))
	require.True(t, containsApprovedCertSerial(all, intermediateCertSerialNumber))
	require.True(t, containsApprovedCertSerial(all, leafCertSerialNumber))

	// All approved root certs — root only.
	allRoots, err := GetAllX509RootCerts()
	require.NoError(t, err)
	require.NotNil(t, allRoots)
	foundRoot, foundInterm, foundLeaf := false, false, false
	for _, id := range allRoots.Certs {
		switch id.Subject {
		case rootCertSubject:
			foundRoot = true
		case intermediateCertSubject:
			foundInterm = true
		case leafCertSubject:
			foundLeaf = true
		}
	}
	require.True(t, foundRoot)
	require.False(t, foundInterm)
	require.False(t, foundLeaf)

	// Subject queries.
	rootBy, err := GetX509CertsBySubject(rootCertSubject)
	require.NoError(t, err)
	require.NotNil(t, rootBy)
	require.Equal(t, rootCertSubject, rootBy.Subject)
	require.Contains(t, rootBy.SubjectKeyIds, rootCertSubjectKeyID)

	leafBy, err := GetX509CertsBySubject(leafCertSubject)
	require.NoError(t, err)
	require.NotNil(t, leafBy)
	require.Equal(t, leafCertSubject, leafBy.Subject)
	require.Contains(t, leafBy.SubjectKeyIds, leafCertSubjectKeyID)

	intermBy, err := GetX509CertsBySubject(intermediateCertSubject)
	require.NoError(t, err)
	require.NotNil(t, intermBy)
	require.Equal(t, intermediateCertSubject, intermBy.Subject)
	require.Contains(t, intermBy.SubjectKeyIds, intermediateCertSubjectKeyID)

	// No proposed-to-revoke entries.
	allProposedRev, err := GetAllProposedRevokedX509RootCerts()
	require.NoError(t, err)
	for _, p := range allProposedRev {
		require.NotEqual(t, rootCertSubject, p.Subject)
		require.NotEqual(t, intermediateCertSubject, p.Subject)
		require.NotEqual(t, leafCertSubject, p.Subject)
	}

	// No revoked certs yet.
	allRevoked, err := GetAllRevokedX509Certs()
	require.NoError(t, err)
	require.False(t, containsRevokedCertSerial(allRevoked, rootCertSerialNumber))
	require.False(t, containsRevokedCertSerial(allRevoked, intermediateCertSerialNumber))
	require.False(t, containsRevokedCertSerial(allRevoked, leafCertSerialNumber))

	// Child certs of root — only intermediate.
	rootChildren, err := GetChildX509Certs(rootCertSubject, rootCertSubjectKeyID)
	require.NoError(t, err)
	require.NotNil(t, rootChildren)
	foundIntermInRootChildren, foundLeafInRootChildren := false, false
	for _, cid := range rootChildren.CertIds {
		if cid.Subject == intermediateCertSubject {
			foundIntermInRootChildren = true
		}
		if cid.Subject == leafCertSubject {
			foundLeafInRootChildren = true
		}
	}
	require.True(t, foundIntermInRootChildren)
	require.False(t, foundLeafInRootChildren)

	// Child certs of intermediate — only leaf.
	intermChildren, err := GetChildX509Certs(intermediateCertSubject, intermediateCertSubjectKeyID)
	require.NoError(t, err)
	require.NotNil(t, intermChildren)
	foundLeafInIntermChildren := false
	for _, cid := range intermChildren.CertIds {
		if cid.Subject == leafCertSubject && cid.SubjectKeyId == leafCertSubjectKeyID {
			foundLeafInIntermChildren = true
		}
	}
	require.True(t, foundLeafInIntermChildren)

	// Child certs of leaf — none.
	leafChildren, err := GetChildX509Certs(leafCertSubject, leafCertSubjectKeyID)
	require.NoError(t, err)
	require.Nil(t, leafChildren)
}

// ── Section 6: Revoke Intermediate Cert ────────────────────────────────────

func pkiDemoRevokeIntermediateCertUnauthorized(t *testing.T, userAccount, vendorAccount65522 string) {
	t.Helper()
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
}

func pkiDemoRevokeIntermediateCert(t *testing.T, vendorAccount string) {
	t.Helper()
	// Revoke intermediate without --revoke-child: leaf must survive
	txResult, err := RevokeX509Cert(intermediateCertSubject, intermediateCertSubjectKeyID, vendorAccount)
	require.NoError(t, err)
	require.Equal(t, uint32(0), txResult.Code)
	_, err = utils.AwaitTxConfirmation(txResult.TxHash)
	require.NoError(t, err)

	// No proposed-to-revoke entries (intermediate is not a root cert).
	allProposedRev, err := GetAllProposedRevokedX509RootCerts()
	require.NoError(t, err)
	for _, p := range allProposedRev {
		require.NotEqual(t, rootCertSubject, p.Subject)
		require.NotEqual(t, intermediateCertSubject, p.Subject)
		require.NotEqual(t, leafCertSubject, p.Subject)
	}

	// All revoked — intermediate present, leaf NOT present, root NOT present.
	allRevoked, err := GetAllRevokedX509Certs()
	require.NoError(t, err)
	require.True(t, containsRevokedCertSerial(allRevoked, intermediateCertSerialNumber))
	require.False(t, containsRevokedCertSerial(allRevoked, leafCertSerialNumber))
	require.False(t, containsRevokedCertSerial(allRevoked, rootCertSerialNumber))

	// All revoked root certs — none of the three.
	allRevokedRoots, err := GetAllRevokedX509RootCerts()
	require.NoError(t, err)
	if allRevokedRoots != nil {
		for _, id := range allRevokedRoots.Certs {
			require.NotEqual(t, rootCertSubject, id.Subject)
			require.NotEqual(t, intermediateCertSubject, id.Subject)
			require.NotEqual(t, leafCertSubject, id.Subject)
		}
	}

	// Revoked intermediate cert by subject+skid.
	revokedInterm, err := GetRevokedX509Cert(intermediateCertSubject, intermediateCertSubjectKeyID)
	require.NoError(t, err)
	require.NotNil(t, revokedInterm)
	require.Equal(t, intermediateCertSubject, revokedInterm.Subject)
	require.Equal(t, intermediateCertSubjectKeyID, revokedInterm.SubjectKeyId)
	require.True(t, containsCertSerial(revokedInterm.Certs, intermediateCertSerialNumber))

	// Leaf cert is NOT in revoked.
	revokedLeaf, err := GetRevokedX509Cert(leafCertSubject, leafCertSubjectKeyID)
	require.NoError(t, err)
	require.Nil(t, revokedLeaf)

	// All approved — root and leaf present, intermediate gone.
	all, err := GetAllX509Certs()
	require.NoError(t, err)
	require.True(t, containsApprovedCertSerial(all, rootCertSerialNumber))
	require.True(t, containsApprovedCertSerial(all, leafCertSerialNumber))
	require.False(t, containsApprovedCertSerial(all, intermediateCertSerialNumber))

	// Approved root certs — root only.
	allRoots, err := GetAllX509RootCerts()
	require.NoError(t, err)
	require.NotNil(t, allRoots)
	hasRoot, hasInterm, hasLeaf := false, false, false
	for _, id := range allRoots.Certs {
		switch id.Subject {
		case rootCertSubject:
			hasRoot = true
		case intermediateCertSubject:
			hasInterm = true
		case leafCertSubject:
			hasLeaf = true
		}
	}
	require.True(t, hasRoot)
	require.False(t, hasInterm)
	require.False(t, hasLeaf)

	// Subject query — leaf still in approved.
	leafBy, err := GetX509CertsBySubject(leafCertSubject)
	require.NoError(t, err)
	require.NotNil(t, leafBy)
	require.Contains(t, leafBy.SubjectKeyIds, leafCertSubjectKeyID)

	// Subject query — intermediate gone from approved.
	intermBy, err := GetX509CertsBySubject(intermediateCertSubject)
	require.NoError(t, err)
	if intermBy != nil {
		require.NotContains(t, intermBy.SubjectKeyIds, intermediateCertSubjectKeyID)
	}

	// Intermediate approved cert — gone.
	intermCert, err := GetX509Cert(intermediateCertSubject, intermediateCertSubjectKeyID)
	require.NoError(t, err)
	require.Nil(t, intermCert)

	// Leaf approved cert — still present.
	leafCert, err := GetX509Cert(leafCertSubject, leafCertSubjectKeyID)
	require.NoError(t, err)
	require.NotNil(t, leafCert)
	require.True(t, containsCertSubjectSerial(leafCert.Certs, leafCertSubject, leafCertSerialNumber))
}

// ── Section 7: Propose Revoke Root Cert ────────────────────────────────────

func pkiDemoProposeRevokeRootCert(t *testing.T, jack string) {
	t.Helper()
	txResult, err := ProposeRevokeX509RootCert(rootCertSubject, rootCertSubjectKeyID, jack)
	require.NoError(t, err)
	require.Equal(t, uint32(0), txResult.Code)
	_, err = utils.AwaitTxConfirmation(txResult.TxHash)
	require.NoError(t, err)

	// Proposed-to-revoke contains root.
	proposedRev, err := GetProposedRevokedX509RootCert(rootCertSubject, rootCertSubjectKeyID)
	require.NoError(t, err)
	require.NotNil(t, proposedRev)
	require.Equal(t, rootCertSubject, proposedRev.Subject)
	require.Equal(t, rootCertSubjectKeyID, proposedRev.SubjectKeyId)

	allProposedRev, err := GetAllProposedRevokedX509RootCerts()
	require.NoError(t, err)
	foundRoot := false
	for _, p := range allProposedRev {
		if p.Subject == rootCertSubject {
			foundRoot = true
		}
		require.NotEqual(t, intermediateCertSubject, p.Subject)
		require.NotEqual(t, leafCertSubject, p.Subject)
	}
	require.True(t, foundRoot)

	// All revoked — intermediate still there, root and leaf absent.
	allRevoked, err := GetAllRevokedX509Certs()
	require.NoError(t, err)
	require.True(t, containsRevokedCertSerial(allRevoked, intermediateCertSerialNumber))
	require.False(t, containsRevokedCertSerial(allRevoked, leafCertSerialNumber))
	require.False(t, containsRevokedCertSerial(allRevoked, rootCertSerialNumber))

	// All revoked root certs — root not yet revoked.
	allRevokedRoots, err := GetAllRevokedX509RootCerts()
	require.NoError(t, err)
	if allRevokedRoots != nil {
		for _, id := range allRevokedRoots.Certs {
			require.NotEqual(t, rootCertSubject, id.Subject)
		}
	}

	// Root cert still approved.
	all, err := GetAllX509Certs()
	require.NoError(t, err)
	require.True(t, containsApprovedCertSerial(all, rootCertSerialNumber))
	require.True(t, containsApprovedCertSerial(all, leafCertSerialNumber))
	require.False(t, containsApprovedCertSerial(all, intermediateCertSerialNumber))

	allRoots, err := GetAllX509RootCerts()
	require.NoError(t, err)
	require.NotNil(t, allRoots)
	hasRoot, hasInterm, hasLeaf := false, false, false
	for _, id := range allRoots.Certs {
		switch id.Subject {
		case rootCertSubject:
			hasRoot = true
		case intermediateCertSubject:
			hasInterm = true
		case leafCertSubject:
			hasLeaf = true
		}
	}
	require.True(t, hasRoot)
	require.False(t, hasInterm)
	require.False(t, hasLeaf)

	// Subject query — root still in approved.
	rootBy, err := GetX509CertsBySubject(rootCertSubject)
	require.NoError(t, err)
	require.NotNil(t, rootBy)
	require.Contains(t, rootBy.SubjectKeyIds, rootCertSubjectKeyID)
}

// ── Section 8: Approve Revoke Root Cert ────────────────────────────────────

func pkiDemoApproveRevokeRootCert(t *testing.T, alice string) {
	t.Helper()
	txResult, err := ApproveRevokeX509RootCert(rootCertSubject, rootCertSubjectKeyID, alice)
	require.NoError(t, err)
	require.Equal(t, uint32(0), txResult.Code)
	_, err = utils.AwaitTxConfirmation(txResult.TxHash)
	require.NoError(t, err)

	// Proposed-to-revoke list empty.
	allProposedRev, err := GetAllProposedRevokedX509RootCerts()
	require.NoError(t, err)
	for _, p := range allProposedRev {
		require.NotEqual(t, rootCertSubject, p.Subject)
		require.NotEqual(t, intermediateCertSubject, p.Subject)
		require.NotEqual(t, leafCertSubject, p.Subject)
	}

	// All revoked — root and intermediate present, leaf absent.
	allRevoked, err := GetAllRevokedX509Certs()
	require.NoError(t, err)
	require.True(t, containsRevokedCertSerial(allRevoked, rootCertSerialNumber))
	require.True(t, containsRevokedCertSerial(allRevoked, intermediateCertSerialNumber))
	require.False(t, containsRevokedCertSerial(allRevoked, leafCertSerialNumber))

	// All revoked root certs — root present, others absent.
	allRevokedRoots, err := GetAllRevokedX509RootCerts()
	require.NoError(t, err)
	require.NotNil(t, allRevokedRoots)
	var rootRev, intermRev, leafRev bool
	for _, id := range allRevokedRoots.Certs {
		switch id.Subject {
		case rootCertSubject:
			rootRev = true
		case intermediateCertSubject:
			intermRev = true
		case leafCertSubject:
			leafRev = true
		}
	}
	require.True(t, rootRev)
	require.False(t, intermRev)
	require.False(t, leafRev)

	// Revoked root cert by subject+skid.
	revokedRoot, err := GetRevokedX509Cert(rootCertSubject, rootCertSubjectKeyID)
	require.NoError(t, err)
	require.NotNil(t, revokedRoot)
	require.Equal(t, rootCertSubject, revokedRoot.Subject)
	require.True(t, containsCertSerial(revokedRoot.Certs, rootCertSerialNumber))

	// All approved — root and intermediate gone, leaf still present.
	all, err := GetAllX509Certs()
	require.NoError(t, err)
	require.False(t, containsApprovedCertSerial(all, rootCertSerialNumber))
	require.False(t, containsApprovedCertSerial(all, intermediateCertSerialNumber))
	require.True(t, containsApprovedCertSerial(all, leafCertSerialNumber))

	// All approved root certs — empty of these.
	allRoots, err := GetAllX509RootCerts()
	require.NoError(t, err)
	if allRoots != nil {
		for _, id := range allRoots.Certs {
			require.NotEqual(t, rootCertSubject, id.Subject)
			require.NotEqual(t, intermediateCertSubject, id.Subject)
			require.NotEqual(t, leafCertSubject, id.Subject)
		}
	}

	// Intermediate approved — gone.
	intermCert, err := GetX509Cert(intermediateCertSubject, intermediateCertSubjectKeyID)
	require.NoError(t, err)
	require.Nil(t, intermCert)

	// Leaf approved — still present.
	leafCert, err := GetX509Cert(leafCertSubject, leafCertSubjectKeyID)
	require.NoError(t, err)
	require.NotNil(t, leafCert)
	c := findCertBySerial(leafCert.Certs, leafCertSerialNumber)
	require.NotNil(t, c)
	require.Equal(t, leafCertSubjectText, c.SubjectAsText)

	// Root approved cert — gone.
	rootCert, err := GetX509Cert(rootCertSubject, rootCertSubjectKeyID)
	require.NoError(t, err)
	require.Nil(t, rootCert)

	// Subject query — root gone from approved.
	rootBy, err := GetX509CertsBySubject(rootCertSubject)
	require.NoError(t, err)
	if rootBy != nil {
		require.NotContains(t, rootBy.SubjectKeyIds, rootCertSubjectKeyID)
	}
}

// ── Section 9: Google Cert Query All Empty ──────────────────────────────────

func pkiDemoGoogleCertQueryAllEmpty(t *testing.T) {
	t.Helper()
	cert, err := GetX509Cert(googleCertSubject, googleCertSubjectKeyID)
	require.NoError(t, err)
	require.Nil(t, cert)

	all, err := GetAllX509Certs()
	require.NoError(t, err)
	require.False(t, containsApprovedCertSerial(all, googleCertSerialNumber))

	proposed, err := GetProposedX509RootCert(googleCertSubject, googleCertSubjectKeyID)
	require.NoError(t, err)
	require.Nil(t, proposed)

	allProposed, err := GetAllProposedX509RootCerts()
	require.NoError(t, err)
	for _, p := range allProposed {
		require.NotEqual(t, googleCertSubject, p.Subject)
	}

	revoked, err := GetRevokedX509Cert(googleCertSubject, googleCertSubjectKeyID)
	require.NoError(t, err)
	require.Nil(t, revoked)

	allRevoked, err := GetAllRevokedX509Certs()
	require.NoError(t, err)
	require.False(t, containsRevokedCertSerial(allRevoked, googleCertSerialNumber))

	bySubject, err := GetX509CertsBySubject(googleCertSubject)
	require.NoError(t, err)
	require.Nil(t, bySubject)

	allRoots, err := GetAllX509RootCerts()
	require.NoError(t, err)
	if allRoots != nil {
		for _, id := range allRoots.Certs {
			require.NotEqual(t, googleCertSubject, id.Subject)
		}
	}

	allRevokedRoots, err := GetAllRevokedX509RootCerts()
	require.NoError(t, err)
	if allRevokedRoots != nil {
		for _, id := range allRevokedRoots.Certs {
			require.NotEqual(t, googleCertSubject, id.Subject)
		}
	}

	proposedRev, err := GetProposedRevokedX509RootCert(googleCertSubject, googleCertSubjectKeyID)
	require.NoError(t, err)
	require.Nil(t, proposedRev)

	allProposedRev, err := GetAllProposedRevokedX509RootCerts()
	require.NoError(t, err)
	for _, p := range allProposedRev {
		require.NotEqual(t, googleCertSubject, p.Subject)
	}

	children, err := GetChildX509Certs(googleCertSubject, googleCertSubjectKeyID)
	require.NoError(t, err)
	require.Nil(t, children)
}

// ── Section 10: Propose Google Root Cert ───────────────────────────────────

func pkiDemoProposeGoogleRootCert(t *testing.T, jack, userAccount string) {
	t.Helper()
	// Non-trustee fails
	txResult, err := ProposeAddX509RootCert(googleCertPath, userAccount, X509ProposeOpts{VID: googleCertVid})
	require.NoError(t, err)
	require.NotEqual(t, uint32(0), txResult.Code)
	_, _ = utils.AwaitTxConfirmation(txResult.TxHash)

	// Trustee proposes
	txResult, err = ProposeAddX509RootCert(googleCertPath, jack, X509ProposeOpts{VID: googleCertVid})
	require.NoError(t, err)
	require.Equal(t, uint32(0), txResult.Code)
	_, err = utils.AwaitTxConfirmation(txResult.TxHash)
	require.NoError(t, err)

	// Proposed cert present.
	proposed, err := GetProposedX509RootCert(googleCertSubject, googleCertSubjectKeyID)
	require.NoError(t, err)
	require.NotNil(t, proposed)
	require.Equal(t, googleCertSubject, proposed.Subject)
	require.Equal(t, googleCertSubjectKeyID, proposed.SubjectKeyId)
	require.Equal(t, googleCertSerialNumber, proposed.SerialNumber)
	require.Equal(t, googleCertSubjectText, proposed.SubjectAsText)
	require.Equal(t, int32(googleCertVid), proposed.Vid)

	// Approved cert must be absent.
	all, err := GetAllX509Certs()
	require.NoError(t, err)
	require.False(t, containsApprovedCertSerial(all, googleCertSerialNumber))

	cert, err := GetX509Cert(googleCertSubject, googleCertSubjectKeyID)
	require.NoError(t, err)
	require.Nil(t, cert)

	allRevoked, err := GetAllRevokedX509Certs()
	require.NoError(t, err)
	require.False(t, containsRevokedCertSerial(allRevoked, googleCertSerialNumber))

	allRoots, err := GetAllX509RootCerts()
	require.NoError(t, err)
	if allRoots != nil {
		for _, id := range allRoots.Certs {
			require.NotEqual(t, googleCertSubject, id.Subject)
		}
	}

	allRevokedRoots, err := GetAllRevokedX509RootCerts()
	require.NoError(t, err)
	if allRevokedRoots != nil {
		for _, id := range allRevokedRoots.Certs {
			require.NotEqual(t, googleCertSubject, id.Subject)
		}
	}

	bySubject, err := GetX509CertsBySubject(googleCertSubject)
	require.NoError(t, err)
	require.Nil(t, bySubject)
}

// ── Section 11: Approve Google Root Cert ───────────────────────────────────

func pkiDemoApproveGoogleRootCert(t *testing.T, alice string) {
	t.Helper()
	// Still proposed, not yet approved.
	proposed, err := GetProposedX509RootCert(googleCertSubject, googleCertSubjectKeyID)
	require.NoError(t, err)
	require.NotNil(t, proposed)
	require.Equal(t, googleCertSubject, proposed.Subject)
	require.Equal(t, googleCertSerialNumber, proposed.SerialNumber)

	all, err := GetAllX509Certs()
	require.NoError(t, err)
	require.False(t, containsApprovedCertSerial(all, googleCertSerialNumber))

	// Alice approves.
	txResult, err := ApproveAddX509RootCert(googleCertSubject, googleCertSubjectKeyID, alice)
	require.NoError(t, err)
	require.Equal(t, uint32(0), txResult.Code)
	_, err = utils.AwaitTxConfirmation(txResult.TxHash)
	require.NoError(t, err)

	// Now approved.
	cert, err := GetX509Cert(googleCertSubject, googleCertSubjectKeyID)
	require.NoError(t, err)
	require.NotNil(t, cert)
	require.Equal(t, googleCertSubject, cert.Subject)
	c := findCertBySerial(cert.Certs, googleCertSerialNumber)
	require.NotNil(t, c)
	require.Equal(t, googleCertSubjectText, c.SubjectAsText)

	bySkid, err := GetX509CertBySKID(googleCertSubjectKeyID)
	require.NoError(t, err)
	require.NotNil(t, bySkid)
	c2 := findCertBySerial(bySkid.Certs, googleCertSerialNumber)
	require.NotNil(t, c2)
	require.Equal(t, googleCertSubjectText, c2.SubjectAsText)

	allProposed, err := GetAllProposedX509RootCerts()
	require.NoError(t, err)
	for _, p := range allProposed {
		require.NotEqual(t, googleCertSubject, p.Subject)
	}

	all, err = GetAllX509Certs()
	require.NoError(t, err)
	require.True(t, containsApprovedCertSerial(all, googleCertSerialNumber))

	allRoots, err := GetAllX509RootCerts()
	require.NoError(t, err)
	require.NotNil(t, allRoots)
	foundGoogle := false
	for _, id := range allRoots.Certs {
		if id.Subject == googleCertSubject && id.SubjectKeyId == googleCertSubjectKeyID {
			foundGoogle = true
		}
	}
	require.True(t, foundGoogle)

	allRevoked, err := GetAllRevokedX509Certs()
	require.NoError(t, err)
	require.False(t, containsRevokedCertSerial(allRevoked, googleCertSerialNumber))
}

// ── Section 12: Propose Revoke Google Root Cert ─────────────────────────────

func pkiDemoProposeRevokeGoogleRootCert(t *testing.T, jack string) {
	t.Helper()
	txResult, err := ProposeRevokeX509RootCert(googleCertSubject, googleCertSubjectKeyID, jack)
	require.NoError(t, err)
	require.Equal(t, uint32(0), txResult.Code)
	_, err = utils.AwaitTxConfirmation(txResult.TxHash)
	require.NoError(t, err)

	proposedRev, err := GetProposedRevokedX509RootCert(googleCertSubject, googleCertSubjectKeyID)
	require.NoError(t, err)
	require.NotNil(t, proposedRev)
	require.Equal(t, googleCertSubject, proposedRev.Subject)
	require.Equal(t, googleCertSubjectKeyID, proposedRev.SubjectKeyId)

	allProposedRev, err := GetAllProposedRevokedX509RootCerts()
	require.NoError(t, err)
	found := false
	for _, p := range allProposedRev {
		if p.Subject == googleCertSubject {
			found = true
		}
	}
	require.True(t, found)

	allRevoked, err := GetAllRevokedX509Certs()
	require.NoError(t, err)
	require.False(t, containsRevokedCertSerial(allRevoked, googleCertSerialNumber))

	allRevokedRoots, err := GetAllRevokedX509RootCerts()
	require.NoError(t, err)
	if allRevokedRoots != nil {
		for _, id := range allRevokedRoots.Certs {
			require.NotEqual(t, googleCertSubject, id.Subject)
		}
	}

	all, err := GetAllX509Certs()
	require.NoError(t, err)
	require.True(t, containsApprovedCertSerial(all, googleCertSerialNumber))

	allRoots, err := GetAllX509RootCerts()
	require.NoError(t, err)
	require.NotNil(t, allRoots)
	foundRoot := false
	for _, id := range allRoots.Certs {
		if id.Subject == googleCertSubject {
			foundRoot = true
		}
	}
	require.True(t, foundRoot)

	bySubject, err := GetX509CertsBySubject(googleCertSubject)
	require.NoError(t, err)
	require.NotNil(t, bySubject)
	require.Contains(t, bySubject.SubjectKeyIds, googleCertSubjectKeyID)
}

// ── Section 13: Approve Revoke Google Root Cert ─────────────────────────────

func pkiDemoApproveRevokeGoogleRootCert(t *testing.T, alice string) {
	t.Helper()
	txResult, err := ApproveRevokeX509RootCert(googleCertSubject, googleCertSubjectKeyID, alice)
	require.NoError(t, err)
	require.Equal(t, uint32(0), txResult.Code)
	_, err = utils.AwaitTxConfirmation(txResult.TxHash)
	require.NoError(t, err)

	allProposedRev, err := GetAllProposedRevokedX509RootCerts()
	require.NoError(t, err)
	for _, p := range allProposedRev {
		require.NotEqual(t, googleCertSubject, p.Subject)
	}

	allRevoked, err := GetAllRevokedX509Certs()
	require.NoError(t, err)
	require.True(t, containsRevokedCertSerial(allRevoked, googleCertSerialNumber))

	allRevokedRoots, err := GetAllRevokedX509RootCerts()
	require.NoError(t, err)
	require.NotNil(t, allRevokedRoots)
	foundRevokedGoogle := false
	for _, id := range allRevokedRoots.Certs {
		if id.Subject == googleCertSubject && id.SubjectKeyId == googleCertSubjectKeyID {
			foundRevokedGoogle = true
		}
	}
	require.True(t, foundRevokedGoogle)

	revoked, err := GetRevokedX509Cert(googleCertSubject, googleCertSubjectKeyID)
	require.NoError(t, err)
	require.NotNil(t, revoked)
	require.Equal(t, googleCertSubject, revoked.Subject)
	c := findCertBySerial(revoked.Certs, googleCertSerialNumber)
	require.NotNil(t, c)
	require.Equal(t, googleCertSubjectText, c.SubjectAsText)

	all, err := GetAllX509Certs()
	require.NoError(t, err)
	require.False(t, containsApprovedCertSerial(all, googleCertSerialNumber))

	allRoots, err := GetAllX509RootCerts()
	require.NoError(t, err)
	if allRoots != nil {
		for _, id := range allRoots.Certs {
			require.NotEqual(t, googleCertSubject, id.Subject)
		}
	}

	cert, err := GetX509Cert(googleCertSubject, googleCertSubjectKeyID)
	require.NoError(t, err)
	require.Nil(t, cert)

	bySubject, err := GetX509CertsBySubject(googleCertSubject)
	require.NoError(t, err)
	if bySubject != nil {
		require.NotContains(t, bySubject.SubjectKeyIds, googleCertSubjectKeyID)
	}
}

// ── Section 14: Propose and Reject Test Cert (single trustee) ──────────────

func pkiDemoProposeAndRejectTestCertSingleTrustee(t *testing.T, jack string) {
	t.Helper()
	// Jack proposes
	txResult, err := ProposeAddX509RootCert(testCertPath, jack, X509ProposeOpts{VID: testCertVid})
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

	// Cert gone from proposed.
	proposed, err := GetProposedX509RootCert(testCertSubject, testCertSubjectKeyID)
	require.NoError(t, err)
	require.Nil(t, proposed)

	// Cert not in rejected (single-trustee reject doesn't reach quorum alone).
	rejected, err := GetRejectedX509RootCert(testCertSubject, testCertSubjectKeyID)
	require.NoError(t, err)
	require.Nil(t, rejected)

	// Cert not in approved.
	cert, err := GetX509Cert(testCertSubject, testCertSubjectKeyID)
	require.NoError(t, err)
	require.Nil(t, cert)

	bySkid, err := GetX509CertBySKID(testCertSubjectKeyID)
	require.NoError(t, err)
	require.Nil(t, bySkid)
}

// ── Section 15: Propose Test Root Cert ─────────────────────────────────────

func pkiDemoProposeTestRootCert(t *testing.T, jack, userAccount string) {
	t.Helper()
	// Non-trustee fails
	txResult, err := ProposeAddX509RootCert(testCertPath, userAccount, X509ProposeOpts{VID: testCertVid})
	require.NoError(t, err)
	require.NotEqual(t, uint32(0), txResult.Code)
	_, _ = utils.AwaitTxConfirmation(txResult.TxHash)

	// Trustee proposes
	txResult, err = ProposeAddX509RootCert(testCertPath, jack, X509ProposeOpts{VID: testCertVid})
	require.NoError(t, err)
	require.Equal(t, uint32(0), txResult.Code)
	_, err = utils.AwaitTxConfirmation(txResult.TxHash)
	require.NoError(t, err)

	// Proposed cert present.
	proposed, err := GetProposedX509RootCert(testCertSubject, testCertSubjectKeyID)
	require.NoError(t, err)
	require.NotNil(t, proposed)
	require.Equal(t, testCertSubject, proposed.Subject)
	require.Equal(t, testCertSubjectKeyID, proposed.SubjectKeyId)
	require.Equal(t, testCertSerialNumber, proposed.SerialNumber)
	require.Equal(t, testCertSubjectText, proposed.SubjectAsText)
	require.Equal(t, int32(testCertVid), proposed.Vid)
}

// ── Section 16: Reject Test Root Cert (multi-trustee scenario) ─────────────

func pkiDemoRejectTestRootCertMultiTrustee(t *testing.T, jack, alice, bob string) {
	t.Helper()
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

	// Cert still in proposed — not yet at rejection quorum.
	proposed, err := GetProposedX509RootCert(testCertSubject, testCertSubjectKeyID)
	require.NoError(t, err)
	require.NotNil(t, proposed)
	require.Equal(t, testCertSubject, proposed.Subject)
	require.Equal(t, testCertSerialNumber, proposed.SerialNumber)

	// Not yet in all-rejected list.
	allRejected, err := GetAllRejectedX509RootCerts()
	require.NoError(t, err)
	for _, r := range allRejected {
		require.NotEqual(t, testCertSubject, r.Subject)
	}

	// Alice rejects — now 2 rejections with 4 trustees = rejection quorum reached
	txResult, err = RejectAddX509RootCert(testCertSubject, testCertSubjectKeyID, alice)
	require.NoError(t, err)
	require.Equal(t, uint32(0), txResult.Code)
	_, err = utils.AwaitTxConfirmation(txResult.TxHash)
	require.NoError(t, err)

	// Revoke new_trustee1 (back to 3 trustees)
	txResult, err = dclauth.ProposeRevokeAccount(newTrustee1Addr, alice)
	require.NoError(t, err)
	require.Equal(t, uint32(0), txResult.Code, txResult.RawLog)
	_, err = utils.AwaitTxConfirmation(txResult.TxHash)
	require.NoError(t, err)

	txResult, err = dclauth.ApproveRevokeAccount(newTrustee1Addr, bob)
	require.NoError(t, err)
	require.Equal(t, uint32(0), txResult.Code, txResult.RawLog)
	_, err = utils.AwaitTxConfirmation(txResult.TxHash)
	require.NoError(t, err)

	txResult, err = dclauth.ApproveRevokeAccount(newTrustee1Addr, jack)
	require.NoError(t, err)
	require.Equal(t, uint32(0), txResult.Code, txResult.RawLog)
	_, err = utils.AwaitTxConfirmation(txResult.TxHash)
	require.NoError(t, err)

	// Cert is now rejected.
	rejected, err := GetRejectedX509RootCert(testCertSubject, testCertSubjectKeyID)
	require.NoError(t, err)
	require.NotNil(t, rejected)
	require.Equal(t, testCertSubject, rejected.Subject)
	require.Equal(t, testCertSubjectKeyID, rejected.SubjectKeyId)
	rejCert := findCertBySerial(rejected.Certs, testCertSerialNumber)
	require.NotNil(t, rejCert)
	require.Equal(t, testCertSubjectText, rejCert.SubjectAsText)

	// No longer in proposed.
	allProposed, err := GetAllProposedX509RootCerts()
	require.NoError(t, err)
	for _, p := range allProposed {
		require.NotEqual(t, testCertSubject, p.Subject)
	}

	// Not in approved root certs.
	allRoots, err := GetAllX509RootCerts()
	require.NoError(t, err)
	if allRoots != nil {
		for _, id := range allRoots.Certs {
			require.NotEqual(t, testCertSubject, id.Subject)
		}
	}

	// In all-rejected list.
	allRejected, err = GetAllRejectedX509RootCerts()
	require.NoError(t, err)
	foundTest := false
	for _, r := range allRejected {
		if r.Subject == testCertSubject && r.SubjectKeyId == testCertSubjectKeyID {
			foundTest = true
		}
	}
	require.True(t, foundTest)
}

// ── Section 17: Propose Test Root Cert Again ────────────────────────────────

func pkiDemoProposeTestRootCertAgain(t *testing.T, jack, userAccount string) {
	t.Helper()
	// Non-trustee fails
	txResult, err := ProposeAddX509RootCert(testCertPath, userAccount, X509ProposeOpts{VID: testCertVid})
	require.NoError(t, err)
	require.NotEqual(t, uint32(0), txResult.Code)
	_, _ = utils.AwaitTxConfirmation(txResult.TxHash)

	// Trustee proposes again
	txResult, err = ProposeAddX509RootCert(testCertPath, jack, X509ProposeOpts{VID: testCertVid})
	require.NoError(t, err)
	require.Equal(t, uint32(0), txResult.Code)
	_, err = utils.AwaitTxConfirmation(txResult.TxHash)
	require.NoError(t, err)

	// Proposed cert present again.
	proposed, err := GetProposedX509RootCert(testCertSubject, testCertSubjectKeyID)
	require.NoError(t, err)
	require.NotNil(t, proposed)
	require.Equal(t, testCertSubject, proposed.Subject)
	require.Equal(t, testCertSubjectKeyID, proposed.SubjectKeyId)
	require.Equal(t, testCertSerialNumber, proposed.SerialNumber)
	require.Equal(t, testCertSubjectText, proposed.SubjectAsText)
	require.Equal(t, int32(testCertVid), proposed.Vid)

	// No longer in rejected (moving from rejected back to proposed clears it).
	rejected, err := GetRejectedX509RootCert(testCertSubject, testCertSubjectKeyID)
	require.NoError(t, err)
	require.Nil(t, rejected)
}

// ── Section 18: Approve Test Root Cert ─────────────────────────────────────

func pkiDemoApproveTestRootCert(t *testing.T, alice string) {
	t.Helper()
	// Alice approves — with 3 trustees (jack+alice+bob), 2 approvals = quorum
	txResult, err := ApproveAddX509RootCert(testCertSubject, testCertSubjectKeyID, alice)
	require.NoError(t, err)
	require.Equal(t, uint32(0), txResult.Code)
	_, err = utils.AwaitTxConfirmation(txResult.TxHash)
	require.NoError(t, err)

	// Approved by subject+skid.
	cert, err := GetX509Cert(testCertSubject, testCertSubjectKeyID)
	require.NoError(t, err)
	require.NotNil(t, cert)
	require.Equal(t, testCertSubject, cert.Subject)
	c := findCertBySerial(cert.Certs, testCertSerialNumber)
	require.NotNil(t, c)
	require.Equal(t, testCertSubjectText, c.SubjectAsText)
	require.Equal(t, int32(testCertVid), c.Vid)

	// Approved by skid only.
	bySkid, err := GetX509CertBySKID(testCertSubjectKeyID)
	require.NoError(t, err)
	require.NotNil(t, bySkid)
	c2 := findCertBySerial(bySkid.Certs, testCertSerialNumber)
	require.NotNil(t, c2)
	require.Equal(t, testCertSubjectText, c2.SubjectAsText)
	require.Equal(t, int32(testCertVid), c2.Vid)
}
