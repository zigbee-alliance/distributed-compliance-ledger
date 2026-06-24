package pki

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	cliputils "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/cli/utils"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
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

// TestPKIAddVendorX509Certificates ports the add-x509-cert authorization /
// VID-scoping error matrix from pki-add-vendor-x509-certificates.sh against the
// google_root_cert_gsr4 → intermediate_cert_gsr4 chain this test owns.
//
// The gsr4 root is non-VID-scoped in its subject but is assigned msg VID=1 at
// proposal (rootCert.Vid=1); the gsr4 intermediate carries no Matter VID
// (childVid=0). That combination exercises:
//   - code 4   — a non-Vendor account cannot add an intermediate (unauthorized).
//   - code 439 — a Vendor whose VID (65522) differs from the root's VID (1).
//   - success  — the Vendor whose VID (1) matches the root's VID.
//
// The remaining shell case — code 440 (a child whose VID differs from a
// VID-scoped root's VID) — needs the VID-scoped root_cert_with_vid + the
// FFF2-scoped intermediate_cert_with_vid_2; that root is owned by
// TestPKIRevocationPoints (which runs later), so the 440 case is ported there
// (AddChildVidNotEqualRootVid_Fails) to avoid a duplicate-root conflict.
//
// Two further add-x509-cert cases from the shell could not be hosted on the
// gsr4 chain this test owns, because they require roots/fixtures that other
// tests already own on the single shared ledger (this test runs first, so
// re-proposing those roots here would make the owning test's propose fail with
// 501 "certificate already exists"):
//
//   - code 4 (a wrong-VID Vendor adds a *second* certificate that shares the
//     subject/SKID of an already-added intermediate): the handler only reaches
//     the EnsureVidMatches(existingOwner, signer) code-4 path
//     (msg_server_add_x_509_cert.go:71) when a cert with the same subject/SKID
//     but a *distinct* serial already exists. The only such pair is
//     intermediate_with_same_subject_and_skid_1/2, which chain to
//     root_with_same_subject_and_skid (owned by TestPKICombineCerts). The gsr4
//     chain has a single intermediate, so re-adding it hits the earlier
//     issuer+serial guard (code 501) instead — it cannot reach the code-4 path.
//     This add-path code-4 case is therefore not portable into this file; note
//     that remove_x509_certs_test.go:85-87 already covers the analogous
//     *remove*-path code-4 on that same fixture set.
//
//   - success: a matching-VID Vendor adds the numeric-VID PAI (FFF2=65522,
//     pai_cert_numeric_vid, serial 4428370313154203676) and all-x509-certs
//     shows it with vid=65522. That PAI is issued by "Matter Test PAA"
//     (paa_cert_no_vid), the only root it chains to, which is owned by
//     TestPKIRevocationPoints; that test also approves paa_cert_no_vid with msg
//     VID=4701, so even the parent's stored Vid differs from the shell's 65522.
//     The success/vid-field intent is exercised below on the gsr4 intermediate
//     instead (whose subject carries no Matter VID, so the stored Vid is 0).
func TestPKIAddVendorX509Certificates(t *testing.T) {
	jack := testconstants.JackAccount
	alice := testconstants.AliceAccount

	// Vendor with VID=65522 — does NOT match the gsr4 root's VID (1).
	vendorAccountWrongVid := fmt.Sprintf("vendor_account_%d", addVendorVid)
	cliputils.CreateVendorAccount(t, vendorAccountWrongVid, addVendorVid)

	// Vendor with VID=1 — matches the gsr4 root's VID.
	vendorAccountMatchingVid := fmt.Sprintf("vendor_account_%d", pkiDemoVid)
	cliputils.CreateVendorAccount(t, vendorAccountMatchingVid, pkiDemoVid)

	t.Run("ProposeAndApproveRootCert", func(t *testing.T) {
		txResult, err := ProposeAddX509RootCert(addVendorRootCertPath, jack, X509ProposeOpts{VID: pkiDemoVid})
		cliputils.RequireTxOK(t, txResult, err)

		txResult, err = ApproveAddX509RootCert(addVendorRootCertSubject, addVendorRootCertSubjectKeyID, alice)
		cliputils.RequireTxOK(t, txResult, err)
	})

	// assertIntermediateAbsent confirms the gsr4 intermediate is not on the
	// approved-cert ledger (used after each rejected add).
	assertIntermediateAbsent := func(t *testing.T) {
		t.Helper()
		cert, err := GetX509Cert(addVendorIntermCertSubject, addVendorIntermCertSubjectKeyID)
		require.NoError(t, err)
		require.Nil(t, cert, "intermediate must not be on the ledger after a rejected add")
		all, err := GetAllX509Certs()
		require.NoError(t, err)
		for i := range all {
			require.NotEqual(t, addVendorIntermCertSubject, all[i].Subject,
				"all-x509-certs must not contain the rejected intermediate")
		}
	}

	t.Run("AddIntermediate_NonVendor_Fails", func(t *testing.T) {
		// A Trustee account (no Vendor role) cannot add an intermediate cert.
		txResult, err := AddX509Cert(addVendorIntermCertPath, jack)
		require.NoError(t, err)
		require.Equal(t, uint32(4), txResult.Code, "expected unauthorized (code 4), raw: %s", txResult.RawLog)
		assertIntermediateAbsent(t)
	})

	t.Run("AddIntermediate_WrongVendorVid_Fails", func(t *testing.T) {
		// Vendor VID (65522) does not match the root's VID (1) → 439.
		txResult, err := AddX509Cert(addVendorIntermCertPath, vendorAccountWrongVid)
		require.NoError(t, err)
		require.Equal(t, uint32(439), txResult.Code, "expected vid-not-equal-account-vid (439), raw: %s", txResult.RawLog)
		assertIntermediateAbsent(t)
	})

	t.Run("AddIntermediate_MatchingVendorVid_Success", func(t *testing.T) {
		// Vendor VID (1) matches the root's VID → add succeeds.
		txResult, err := AddX509Cert(addVendorIntermCertPath, vendorAccountMatchingVid)
		cliputils.RequireTxOK(t, txResult, err)

		cert, err := GetX509Cert(addVendorIntermCertSubject, addVendorIntermCertSubjectKeyID)
		require.NoError(t, err)
		require.NotNil(t, cert)
		require.Equal(t, addVendorIntermCertSubject, cert.Subject)
		require.Equal(t, addVendorIntermCertSubjectKeyID, cert.SubjectKeyId)
		require.NotEmpty(t, cert.Certs)
		require.False(t, cert.Certs[0].IsRoot)
		// The stored Vid is the Matter VID parsed from the child's subject
		// (msg_server_add_x_509_cert.go sets it from GetVidFromSubject). The gsr4
		// intermediate carries no Matter VID, so it is 0 — this mirrors the shell's
		// "vid field is present and correct" check; the numeric-VID (65522) variant
		// is not portable here (see the file header).
		require.Equal(t, int32(0), cert.Certs[0].Vid)

		// The intermediate now appears in all-x509-certs (matched by serial).
		serial := cert.Certs[0].SerialNumber
		all, err := GetAllX509Certs()
		require.NoError(t, err)
		require.True(t, containsApprovedCertSerial(all, serial),
			"all-x509-certs should contain the approved intermediate serial %s", serial)
	})
}
