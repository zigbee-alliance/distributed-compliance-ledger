package pki

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	cliputils "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/cli/utils"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
)

// requirePEMEquals asserts that a stored PEM field equals the PEM in the given
// fixture file, ignoring surrounding whitespace (the CLI may add/strip a trailing
// newline when reading the --certificate/--certificate-delegator file).
func requirePEMEquals(t *testing.T, fixturePath, stored string) {
	t.Helper()
	want, err := os.ReadFile(fixturePath)
	require.NoError(t, err)
	require.Equal(t, strings.TrimSpace(string(want)), strings.TrimSpace(stored))
}

// ensureMainnetPAAOnLedger makes paa_cert_no_vid_mainnet (Matter PAA 2, VID 24582)
// available as an approved root cert, proposing+approving it only if it is not
// already present (other PKI tests may add/remove it on the shared chain).
// It reuses the approvalTestRootCert* constants (same fixture, defined in approval_test.go).
func ensureMainnetPAAOnLedger(t *testing.T, jack, alice string) {
	t.Helper()
	cert, err := GetX509Cert(approvalTestRootCertSubject, approvalTestRootCertSubjectKeyID)
	require.NoError(t, err)
	if cert != nil {
		return
	}

	txResult, err := ProposeAddX509RootCert(approvalTestRootCertPath, jack, X509ProposeOpts{VID: revPointVid24582})
	cliputils.RequireTxOK(t, txResult, err)

	txResult, err = ApproveAddX509RootCert(approvalTestRootCertSubject, approvalTestRootCertSubjectKeyID, alice)
	cliputils.RequireTxOK(t, txResult, err)
}

const (
	paaCertWithNumericVidPath         = "../../constants/paa_cert_numeric_vid"
	paaCertWithNumericVidSubject      = "MDAxGDAWBgNVBAMMD01hdHRlciBUZXN0IFBBQTEUMBIGCisGAQQBgqJ8AgEMBEZGRjE="
	paaCertWithNumericVidSubjectKeyID = "6A:FD:22:77:1F:51:1F:EC:BF:16:41:97:67:10:DC:DC:31:A1:71:7E"

	paaCertNoVidPath         = "../../constants/paa_cert_no_vid"
	paaCertNoVidSubject      = "MBoxGDAWBgNVBAMMD01hdHRlciBUZXN0IFBBQQ=="
	paaCertNoVidSubjectKeyID = "78:5C:E7:05:B8:6B:8F:4E:6F:C7:93:AA:60:CB:43:EA:69:68:82:D5"

	paaCertWithNumericVid1Path = "../../constants/paa_cert_numeric_vid_1"

	paiCertWithNumericVidPath    = "../../constants/pai_cert_numeric_vid"
	paiCertWithNumericVidPidPath = "../../constants/pai_cert_numeric_vid_pid"

	delegatorCertWithVid65521Path     = "../../constants/intermediate_cert_with_vid_1"
	delegatorCertWithVid65521CopyPath = "../../constants/intermediate_cert_with_vid_1_copy"
	delegatorCertSubjectKeyID         = "B07B3FF14501918FC1FAEECB9A0106C7479B5DEC"

	crlSignerDelegatedByPAI1Path = "../../constants/leaf_cert_with_vid_65521"
	crlSignerDelegatedByPAI2Path = "../../constants/leaf_cert_with_vid_65522"
	// leaf_cert_without_vid is delegated by intermediate_cert_with_vid_1 (SKID B07B…),
	// the same delegator SKID as the leaf_cert_with_vid_* certs. Used for the
	// "CRL signer delegated by PAA" (is-paa=true) add path.
	crlSignerDelegatedByPAAPath = "../../constants/leaf_cert_without_vid"

	// delegated_CRL_signer_certificate is a non-self-signed CRL signer delegated by
	// pai_cert_certificate_delegator, which chains back to paa_cert_no_vid_mainnet
	// (Matter PAA 2, VID 24582). Used for the PAI add path with --certificate-delegator.
	delegatedCRLSignerCertPath       = "../../constants/delegated_CRL_signer_certificate"
	paiCertCertificateDelegatorPath  = "../../constants/pai_cert_certificate_delegator"
	delegatedCRLSignerCertIssuerSKID = "E981D0E419765AB12F6D03A734CF003307870F0A"

	// pai_cert_vid is a VID-scoped PAI (VID FFF2=65522) chained to paa_cert_no_vid.
	paiCertVidPath = "../../constants/pai_cert_vid"

	rootCertWithVidRevPath         = "../../constants/root_cert_with_vid"
	rootCertWithVidRevSubject      = "MIGYMQswCQYDVQQGEwJVUzERMA8GA1UECBMITmV3IFlvcmsxETAPBgNVBAcTCE5ldyBZb3JrMRgwFgYDVQQKEw9FeGFtcGxlIENvbXBhbnkxGTAXBgNVBAsTEFRlc3RpbmcgRGl2aXNpb24xGDAWBgNVBAMTD3d3dy5leGFtcGxlLmNvbTEUMBIGCisGAQQBgqJ8AgETBEZGRjE="
	rootCertWithVidRevSubjectKeyID = "6B:8C:77:1E:AD:CB:A8:3C:33:9C:2F:10:27:5F:42:03:1D:0A:F4:8E"

	// intermediate_cert_with_vid_2 is VID-scoped to FFF2 (65522) but chains to
	// root_cert_with_vid (FFF1=65521) — used for the add-x509-cert 440 case.
	intermediateCertWithVid2Path = "../../constants/intermediate_cert_with_vid_2"

	revPointVid          = 65521
	revPointVid65522     = 65522
	revPointVid24582     = 24582
	revPointVidNonScoped = 4701

	revPointLabel             = "label"
	revPointLabelPAI          = "label_pai"
	revPointLabelLeaf         = "label_leaf"
	revPointLabelLeafDel      = "label_leaf_with_delegator"
	revPointLabelIntermediate = "label_intermediate"
	revPointLabelNonScoped    = "label2"

	revPointDataURL          = "https://url.data.dclmodel"
	revPointDataURLNonScoped = "https://url.data.dclmodel2"
	revPointIssuerSKID       = "DF4EAFB08C9C37781AE75312CAE4786B481EAFB0"
	// SKID of google_root_cert_gsr4 (no colons) — used for non-VID-scoped PAI revocation point.
	// intermediate_cert_gsr4 (no VID) chains to this root, which is on the ledger after TestPKIAddVendorX509Certificates.
	revPointGsr4IssuerSKID = "54B07BAD45B8E2407FFB0A6EFBBE33C93CA384D5"
)

func TestPKIRevocationPoints(t *testing.T) {
	jack := testconstants.JackAccount
	alice := testconstants.AliceAccount

	vendorAccount := fmt.Sprintf("vendor_account_%d", revPointVid)
	cliputils.CreateVendorAccount(t, vendorAccount, revPointVid)

	vendorAccount65522 := fmt.Sprintf("vendor_account_%d", revPointVid65522)
	cliputils.CreateVendorAccount(t, vendorAccount65522, revPointVid65522)

	vendorAccount24582 := fmt.Sprintf("vendor_account_%d", revPointVid24582)
	cliputils.CreateVendorAccount(t, vendorAccount24582, revPointVid24582)

	vendorAccountNonScoped := fmt.Sprintf("vendor_account_%d", revPointVidNonScoped)
	cliputils.CreateVendorAccount(t, vendorAccountNonScoped, revPointVidNonScoped)

	t.Run("QueryAllEmpty", func(t *testing.T) {
		all, err := GetAllPkiRevocationDistributionPoints()
		require.NoError(t, err)
		require.Empty(t, all)
	})

	t.Run("QueryRevocationPointNotFound", func(t *testing.T) {
		point, err := GetPkiRevocationDistributionPoint(revPointVid, revPointLabel, "AB")
		require.NoError(t, err)
		require.Nil(t, point)

		byIssuer, err := GetPkiRevocationDistributionPointsByIssuer("AB")
		require.NoError(t, err)
		require.Nil(t, byIssuer)
	})

	t.Run("AddRevocationPointFailures", func(t *testing.T) {
		// Not by vendor
		txResult, err := AddRevocationPoint(jack, RevocationPointOpts{
			VID:                revPointVid,
			IsPAA:              true,
			Certificate:        paaCertWithNumericVidPath,
			Label:              revPointLabel,
			DataURL:            revPointDataURL,
			IssuerSubjectKeyID: revPointIssuerSKID,
			RevocationType:     "1",
		})
		cliputils.RequireTxFails(t, txResult, err)

		// Sender VID not equal to field VID
		txResult, err = AddRevocationPoint(vendorAccount65522, RevocationPointOpts{
			VID:                revPointVid,
			IsPAA:              true,
			Certificate:        paaCertWithNumericVidPath,
			Label:              revPointLabel,
			DataURL:            revPointDataURL,
			IssuerSubjectKeyID: revPointIssuerSKID,
			RevocationType:     "1",
		})
		cliputils.RequireTxFails(t, txResult, err)

		// Certificate does not exist on ledger
		txResult, err = AddRevocationPoint(vendorAccount, RevocationPointOpts{
			VID:                revPointVid,
			IsPAA:              true,
			Certificate:        paaCertWithNumericVidPath,
			Label:              revPointLabel,
			DataURL:            revPointDataURL,
			IssuerSubjectKeyID: revPointIssuerSKID,
			RevocationType:     "1",
		})
		cliputils.RequireTxFails(t, txResult, err)

		// PAI revocation point added by a non-vendor (trustee) — should fail
		txResult, err = AddRevocationPoint(jack, RevocationPointOpts{
			VID:                revPointVid,
			PID:                32768,
			Certificate:        paiCertWithNumericVidPidPath,
			Label:              revPointLabel,
			DataURL:            revPointDataURL,
			IssuerSubjectKeyID: revPointIssuerSKID,
			RevocationType:     "1",
		})
		cliputils.RequireTxFails(t, txResult, err)
	})

	t.Run("AddCertsToLedger", func(t *testing.T) {
		// Trustees add PAA cert with numeric VID
		txResult, err := ProposeAddX509RootCert(paaCertWithNumericVidPath, jack, X509ProposeOpts{VID: revPointVid})
		cliputils.RequireTxOK(t, txResult, err)

		txResult, err = ApproveAddX509RootCert(paaCertWithNumericVidSubject, paaCertWithNumericVidSubjectKeyID, alice)
		cliputils.RequireTxOK(t, txResult, err)

		// Trustees add PAA no VID
		txResult, err = ProposeAddX509RootCert(paaCertNoVidPath, jack, X509ProposeOpts{VID: revPointVidNonScoped})
		cliputils.RequireTxOK(t, txResult, err)

		txResult, err = ApproveAddX509RootCert(paaCertNoVidSubject, paaCertNoVidSubjectKeyID, alice)
		cliputils.RequireTxOK(t, txResult, err)

		// Add VID-scoped root cert
		txResult, err = ProposeAddX509RootCert(rootCertWithVidRevPath, jack, X509ProposeOpts{VID: revPointVid})
		cliputils.RequireTxOK(t, txResult, err)

		txResult, err = ApproveAddX509RootCert(rootCertWithVidRevSubject, rootCertWithVidRevSubjectKeyID, alice)
		cliputils.RequireTxOK(t, txResult, err)
	})

	t.Run("AddChildVidNotEqualRootVid_Fails", func(t *testing.T) {
		// Port of pki-add-vendor-x509-certificates.sh:79-82 (code 440): adding an
		// intermediate whose VID (65522) differs from its VID-scoped root's VID
		// (root_cert_with_vid, 65521) is rejected. This lives here because this
		// test owns root_cert_with_vid on the ledger; the add fails, so no ledger
		// state changes. The childVid≠rootVid check precedes the account-VID check,
		// so 440 fires regardless of the signer's VID.
		txResult, err := AddX509Cert(intermediateCertWithVid2Path, vendorAccount65522)
		require.NoError(t, err)
		require.Equal(t, uint32(440), txResult.Code, "expected cert-vid-not-equal-root-vid (440), raw: %s", txResult.RawLog)
	})

	t.Run("AddRevocationPointForVidScopedPAA", func(t *testing.T) {
		// CRL signer cert PEM value not equal to stored — should fail
		txResult, err := AddRevocationPoint(vendorAccount65522, RevocationPointOpts{
			VID:                revPointVid65522,
			IsPAA:              true,
			Certificate:        paaCertWithNumericVid1Path,
			Label:              revPointLabel,
			DataURL:            revPointDataURL,
			IssuerSubjectKeyID: revPointIssuerSKID,
			RevocationType:     "1",
		})
		cliputils.RequireTxFails(t, txResult, err)

		// Add for VID-scoped PAA
		txResult, err = AddRevocationPoint(vendorAccount, RevocationPointOpts{
			VID:                revPointVid,
			IsPAA:              true,
			Certificate:        paaCertWithNumericVidPath,
			Label:              revPointLabel,
			DataURL:            revPointDataURL,
			IssuerSubjectKeyID: revPointIssuerSKID,
			RevocationType:     "1",
			SchemaVersion:      "0",
		})
		cliputils.RequireTxOK(t, txResult, err)

		all, err := GetAllPkiRevocationDistributionPoints()
		require.NoError(t, err)
		require.True(t, containsRevocationPointByLabel(all, int32(revPointVid), revPointLabel))

		// Cannot add same point twice (same vid, issuer, label)
		txResult, err = AddRevocationPoint(vendorAccount, RevocationPointOpts{
			VID:                revPointVid,
			IsPAA:              true,
			Certificate:        paaCertWithNumericVidPath,
			Label:              revPointLabel,
			DataURL:            revPointDataURL + "-new",
			IssuerSubjectKeyID: revPointIssuerSKID,
			RevocationType:     "1",
		})
		cliputils.RequireTxFails(t, txResult, err)

		// Cannot add same point twice (same vid, issuer, dataURL) — new label reuses
		// the already-stored dataURL, which the handler rejects as a duplicate key.
		txResult, err = AddRevocationPoint(vendorAccount, RevocationPointOpts{
			VID:                revPointVid,
			IsPAA:              true,
			Certificate:        paaCertWithNumericVidPath,
			Label:              revPointLabel + "-new",
			DataURL:            revPointDataURL,
			IssuerSubjectKeyID: revPointIssuerSKID,
			RevocationType:     "1",
		})
		cliputils.RequireTxFails(t, txResult, err)
	})

	t.Run("AddRevocationPointForNonVidScopedPAA", func(t *testing.T) {
		txResult, err := AddRevocationPoint(vendorAccountNonScoped, RevocationPointOpts{
			VID:                revPointVidNonScoped,
			IsPAA:              true,
			Certificate:        paaCertNoVidPath,
			Label:              revPointLabelNonScoped,
			DataURL:            revPointDataURLNonScoped,
			IssuerSubjectKeyID: revPointIssuerSKID,
			RevocationType:     "1",
		})
		cliputils.RequireTxOK(t, txResult, err)

		all, err := GetAllPkiRevocationDistributionPoints()
		require.NoError(t, err)
		require.True(t, containsRevocationPointByLabel(all, int32(revPointVid), revPointLabel))
		require.True(t, containsRevocationPointByLabel(all, int32(revPointVidNonScoped), revPointLabelNonScoped))
	})

	t.Run("AddRevocationPointForPAI", func(t *testing.T) {
		txResult, err := AddRevocationPoint(vendorAccount65522, RevocationPointOpts{
			VID:                revPointVid65522,
			Certificate:        paiCertWithNumericVidPath,
			Label:              revPointLabelPAI,
			DataURL:            revPointDataURL,
			IssuerSubjectKeyID: revPointIssuerSKID,
			RevocationType:     "1",
		})
		cliputils.RequireTxOK(t, txResult, err)

		point, err := GetPkiRevocationDistributionPoint(revPointVid65522, revPointLabelPAI, revPointIssuerSKID)
		require.NoError(t, err)
		require.NotNil(t, point)
		require.Equal(t, int32(revPointVid65522), point.Vid)
		require.Equal(t, revPointLabelPAI, point.Label)
	})

	t.Run("AddRevocationPointWithDelegator", func(t *testing.T) {
		// Add PAI cert to ledger
		txResult, err := AddX509Cert(delegatorCertWithVid65521Path, vendorAccount)
		cliputils.RequireTxOK(t, txResult, err)

		// Add revocation point with delegator
		txResult, err = AddRevocationPoint(vendorAccount, RevocationPointOpts{
			VID:                  revPointVid,
			Certificate:          crlSignerDelegatedByPAI1Path,
			Label:                revPointLabelLeafDel,
			DataURL:              revPointDataURL,
			IssuerSubjectKeyID:   delegatorCertSubjectKeyID,
			RevocationType:       "1",
			CertificateDelegator: delegatorCertWithVid65521Path,
		})
		cliputils.RequireTxOK(t, txResult, err)

		point, err := GetPkiRevocationDistributionPoint(revPointVid, revPointLabelLeafDel, delegatorCertSubjectKeyID)
		require.NoError(t, err)
		require.NotNil(t, point)
		require.Equal(t, int32(revPointVid), point.Vid)
		require.Equal(t, revPointLabelLeafDel, point.Label)
		// CRL signer PEM-body field assertions (the cert and its delegator).
		requirePEMEquals(t, crlSignerDelegatedByPAI1Path, point.CrlSignerCertificate)
		requirePEMEquals(t, delegatorCertWithVid65521Path, point.CrlSignerDelegator)
	})

	t.Run("AddRevocationPointForCRLSignerDelegatedByPAA", func(t *testing.T) {
		// is-paa=true path: leaf_cert_without_vid is delegated by
		// intermediate_cert_with_vid_1 (added to the ledger in the previous subtest).
		txResult, err := AddRevocationPoint(vendorAccount65522, RevocationPointOpts{
			VID:                revPointVid65522,
			IsPAA:              true,
			Certificate:        crlSignerDelegatedByPAAPath,
			Label:              revPointLabelLeaf,
			DataURL:            revPointDataURL,
			IssuerSubjectKeyID: delegatorCertSubjectKeyID,
			RevocationType:     "1",
		})
		cliputils.RequireTxOK(t, txResult, err)

		point, err := GetPkiRevocationDistributionPoint(revPointVid65522, revPointLabelLeaf, delegatorCertSubjectKeyID)
		require.NoError(t, err)
		require.NotNil(t, point)
		require.Equal(t, int32(revPointVid65522), point.Vid)
		require.Equal(t, revPointLabelLeaf, point.Label)
		require.Equal(t, delegatorCertSubjectKeyID, point.IssuerSubjectKeyID)
		requirePEMEquals(t, crlSignerDelegatedByPAAPath, point.CrlSignerCertificate)
	})

	t.Run("AddRevocationPointForPAIWithDelegator", func(t *testing.T) {
		// PAI add path with --certificate-delegator. The delegator
		// (pai_cert_certificate_delegator) chains back to paa_cert_no_vid_mainnet
		// (Matter PAA 2, VID 24582), which must be on the ledger first.
		ensureMainnetPAAOnLedger(t, jack, alice)

		txResult, err := AddRevocationPoint(vendorAccount, RevocationPointOpts{
			VID:                  revPointVid,
			Certificate:          delegatedCRLSignerCertPath,
			Label:                revPointLabelLeafDel,
			DataURL:              revPointDataURL,
			IssuerSubjectKeyID:   delegatedCRLSignerCertIssuerSKID,
			RevocationType:       "1",
			CertificateDelegator: paiCertCertificateDelegatorPath,
		})
		cliputils.RequireTxOK(t, txResult, err)

		point, err := GetPkiRevocationDistributionPoint(revPointVid, revPointLabelLeafDel, delegatedCRLSignerCertIssuerSKID)
		require.NoError(t, err)
		require.NotNil(t, point)
		require.Equal(t, int32(revPointVid), point.Vid)
		require.Equal(t, revPointLabelLeafDel, point.Label)
		require.Equal(t, delegatedCRLSignerCertIssuerSKID, point.IssuerSubjectKeyID)
		// Assert the stored CRL signer cert and its delegator PEM-body values.
		requirePEMEquals(t, delegatedCRLSignerCertPath, point.CrlSignerCertificate)
		requirePEMEquals(t, paiCertCertificateDelegatorPath, point.CrlSignerDelegator)
	})

	t.Run("AddRevocationPointForNonVidScopedPAI", func(t *testing.T) {
		// Use intermediate_cert_gsr4 (no VID) instead of intermediate_cert: root_cert (issuer of
		// intermediate_cert) is revoked by TestPKIDemo, so its chain cannot be verified.
		// google_root_cert_gsr4 (issuer of intermediate_cert_gsr4) is already on the ledger
		// from TestPKIAddVendorX509Certificates.
		txResult, err := AddRevocationPoint(vendorAccount, RevocationPointOpts{
			VID:                revPointVid,
			Certificate:        addVendorIntermCertPath,
			Label:              revPointLabelIntermediate,
			DataURL:            revPointDataURLNonScoped,
			IssuerSubjectKeyID: revPointGsr4IssuerSKID,
			RevocationType:     "1",
		})
		cliputils.RequireTxOK(t, txResult, err)
	})

	t.Run("UpdateRevocationPoints", func(t *testing.T) {
		dataURLNew := revPointDataURL + "_new"

		// Update with delegator
		txResult, err := UpdateRevocationPoint(vendorAccount, RevocationPointOpts{
			VID:                  revPointVid,
			Certificate:          crlSignerDelegatedByPAI1Path,
			Label:                revPointLabelLeafDel,
			DataURL:              dataURLNew,
			IssuerSubjectKeyID:   delegatorCertSubjectKeyID,
			CertificateDelegator: delegatorCertWithVid65521CopyPath,
		})
		cliputils.RequireTxOK(t, txResult, err)

		point, err := GetPkiRevocationDistributionPoint(revPointVid, revPointLabelLeafDel, delegatorCertSubjectKeyID)
		require.NoError(t, err)
		require.NotNil(t, point)
		require.Equal(t, dataURLNew, point.DataURL)
		requirePEMEquals(t, crlSignerDelegatedByPAI1Path, point.CrlSignerCertificate)
		requirePEMEquals(t, delegatorCertWithVid65521CopyPath, point.CrlSignerDelegator)

		// Update CRL signer delegated by PAA (label_leaf point added with IsPAA=true).
		// Swap to leaf_cert_with_vid_65522, which is also delegated by
		// intermediate_cert_with_vid_1 (on the ledger) and VID-scoped to FFF2 (65522).
		txResult, err = UpdateRevocationPoint(vendorAccount65522, RevocationPointOpts{
			VID:                revPointVid65522,
			Certificate:        crlSignerDelegatedByPAI2Path,
			Label:              revPointLabelLeaf,
			DataURL:            dataURLNew,
			IssuerSubjectKeyID: delegatorCertSubjectKeyID,
		})
		cliputils.RequireTxOK(t, txResult, err)

		point, err = GetPkiRevocationDistributionPoint(revPointVid65522, revPointLabelLeaf, delegatorCertSubjectKeyID)
		require.NoError(t, err)
		require.NotNil(t, point)
		require.Equal(t, dataURLNew, point.DataURL)
		requirePEMEquals(t, crlSignerDelegatedByPAI2Path, point.CrlSignerCertificate)

		// Update non-VID-scoped PAI (uses addVendorIntermCertPath / revPointGsr4IssuerSKID)
		dataURLNonScopedNew := revPointDataURLNonScoped + "_new"
		txResult, err = UpdateRevocationPoint(vendorAccount, RevocationPointOpts{
			VID:                revPointVid,
			Label:              revPointLabelIntermediate,
			Certificate:        addVendorIntermCertPath,
			DataURL:            dataURLNonScopedNew,
			IssuerSubjectKeyID: revPointGsr4IssuerSKID,
		})
		cliputils.RequireTxOK(t, txResult, err)

		point, err = GetPkiRevocationDistributionPoint(revPointVid, revPointLabelIntermediate, revPointGsr4IssuerSKID)
		require.NoError(t, err)
		require.NotNil(t, point)
		require.Equal(t, dataURLNonScopedNew, point.DataURL)

		// Update VID-scoped PAA (use rootCertWithVidRevPath which is on-ledger, ECDSA P-256,
		// and shares the FFF1 VID-scope so the handler's revocationPoint.Vid check passes)
		txResult, err = UpdateRevocationPoint(vendorAccount, RevocationPointOpts{
			VID:                revPointVid,
			Certificate:        rootCertWithVidRevPath,
			Label:              revPointLabel,
			DataURL:            revPointDataURL,
			IssuerSubjectKeyID: revPointIssuerSKID,
			SchemaVersion:      "0",
		})
		cliputils.RequireTxOK(t, txResult, err)

		point, err = GetPkiRevocationDistributionPoint(revPointVid, revPointLabel, revPointIssuerSKID)
		require.NoError(t, err)
		require.NotNil(t, point)
		requirePEMEquals(t, rootCertWithVidRevPath, point.CrlSignerCertificate)
		require.Equal(t, uint32(0), point.SchemaVersion)

		// Update non-VID-scoped PAA (label2 point, paa_cert_no_vid is self-signed and
		// non-VID-scoped, VID 4701 assigned at proposal). Re-supplying the same on-ledger
		// cert exercises verifyUpdatedPAA positively without depending on test_root_cert,
		// whose ledger presence is owned by TestPKIDemo.
		dataURLNonScopedPAANew := revPointDataURLNonScoped + "_paa_new"
		txResult, err = UpdateRevocationPoint(vendorAccountNonScoped, RevocationPointOpts{
			VID:                revPointVidNonScoped,
			Certificate:        paaCertNoVidPath,
			Label:              revPointLabelNonScoped,
			DataURL:            dataURLNonScopedPAANew,
			IssuerSubjectKeyID: revPointIssuerSKID,
		})
		cliputils.RequireTxOK(t, txResult, err)

		point, err = GetPkiRevocationDistributionPoint(revPointVidNonScoped, revPointLabelNonScoped, revPointIssuerSKID)
		require.NoError(t, err)
		require.NotNil(t, point)
		require.Equal(t, dataURLNonScopedPAANew, point.DataURL)
		requirePEMEquals(t, paaCertNoVidPath, point.CrlSignerCertificate)

		// Update PAI (label_pai point). Swap to pai_cert_vid (VID-scoped FFF2=65522),
		// which chains back to paa_cert_no_vid on the ledger.
		txResult, err = UpdateRevocationPoint(vendorAccount65522, RevocationPointOpts{
			VID:                revPointVid65522,
			Certificate:        paiCertVidPath,
			Label:              revPointLabelPAI,
			DataURL:            revPointDataURL,
			IssuerSubjectKeyID: revPointIssuerSKID,
		})
		cliputils.RequireTxOK(t, txResult, err)

		point, err = GetPkiRevocationDistributionPoint(revPointVid65522, revPointLabelPAI, revPointIssuerSKID)
		require.NoError(t, err)
		require.NotNil(t, point)
		requirePEMEquals(t, paiCertVidPath, point.CrlSignerCertificate)

		// Update failure: new cert is not a PAA (the VID-scoped PAA point's old cert is
		// self-signed, so verifyUpdatedPAA rejects a non-self-signed replacement).
		txResult, err = UpdateRevocationPoint(vendorAccount, RevocationPointOpts{
			VID:                revPointVid,
			Certificate:        paiCertWithNumericVidPidPath,
			Label:              revPointLabel,
			DataURL:            revPointDataURL,
			IssuerSubjectKeyID: revPointIssuerSKID,
		})
		cliputils.RequireTxFails(t, txResult, err)

		// Update failure: sender VID (account 65522) not equal to cert/msg VID (65521).
		txResult, err = UpdateRevocationPoint(vendorAccount65522, RevocationPointOpts{
			VID:                revPointVid,
			Certificate:        paaCertWithNumericVidPath,
			Label:              revPointLabel,
			DataURL:            revPointDataURL,
			IssuerSubjectKeyID: revPointIssuerSKID,
		})
		cliputils.RequireTxFails(t, txResult, err)

		// Update failure: msg VID (65522) not equal to cert VID (paa_cert_numeric_vid is FFF1=65521).
		txResult, err = UpdateRevocationPoint(vendorAccount, RevocationPointOpts{
			VID:                revPointVid65522,
			Certificate:        paaCertWithNumericVidPath,
			Label:              revPointLabel,
			DataURL:            revPointDataURL,
			IssuerSubjectKeyID: revPointIssuerSKID,
		})
		cliputils.RequireTxFails(t, txResult, err)

		// Update failure: point not found
		txResult, err = UpdateRevocationPoint(vendorAccount65522, RevocationPointOpts{
			VID:                revPointVid65522,
			Certificate:        paiCertWithNumericVidPidPath,
			Label:              revPointLabel,
			DataURL:            revPointDataURL,
			IssuerSubjectKeyID: revPointIssuerSKID,
		})
		cliputils.RequireTxFails(t, txResult, err)

		// Update failure: sender not vendor
		txResult, err = UpdateRevocationPoint(jack, RevocationPointOpts{
			VID:                revPointVid,
			Certificate:        paaCertWithNumericVidPath,
			Label:              revPointLabel,
			DataURL:            revPointDataURL,
			IssuerSubjectKeyID: revPointIssuerSKID,
		})
		cliputils.RequireTxFails(t, txResult, err)
	})

	t.Run("DeleteRevocationPoint", func(t *testing.T) {
		txResult, err := DeleteRevocationPoint(vendorAccount, RevocationPointOpts{
			VID:                revPointVid,
			Label:              revPointLabel,
			IssuerSubjectKeyID: revPointIssuerSKID,
		})
		cliputils.RequireTxOK(t, txResult, err)

		point, err := GetPkiRevocationDistributionPoint(revPointVid, revPointLabel, revPointIssuerSKID)
		require.NoError(t, err)
		require.Nil(t, point)
	})
}
