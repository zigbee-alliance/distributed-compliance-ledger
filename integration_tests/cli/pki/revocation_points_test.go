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
		require.NoError(t, err)
		require.NotEqual(t, uint32(0), txResult.Code)

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
		require.NoError(t, err)
		require.NotEqual(t, uint32(0), txResult.Code)

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
		require.NoError(t, err)
		require.NotEqual(t, uint32(0), txResult.Code)
	})

	t.Run("AddCertsToLedger", func(t *testing.T) {
		// Trustees add PAA cert with numeric VID
		txResult, err := ProposeAddX509RootCert(paaCertWithNumericVidPath, jack, X509ProposeOpts{VID: revPointVid})
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		txResult, err = ApproveAddX509RootCert(paaCertWithNumericVidSubject, paaCertWithNumericVidSubjectKeyID, alice)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Trustees add PAA no VID
		txResult, err = ProposeAddX509RootCert(paaCertNoVidPath, jack, X509ProposeOpts{VID: revPointVidNonScoped})
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		txResult, err = ApproveAddX509RootCert(paaCertNoVidSubject, paaCertNoVidSubjectKeyID, alice)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Add VID-scoped root cert
		txResult, err = ProposeAddX509RootCert(rootCertWithVidRevPath, jack, X509ProposeOpts{VID: revPointVid})
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		txResult, err = ApproveAddX509RootCert(rootCertWithVidRevSubject, rootCertWithVidRevSubjectKeyID, alice)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)
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
		require.NoError(t, err)
		require.NotEqual(t, uint32(0), txResult.Code)

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
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

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
		require.NoError(t, err)
		require.NotEqual(t, uint32(0), txResult.Code)
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
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

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
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		point, err := GetPkiRevocationDistributionPoint(revPointVid65522, revPointLabelPAI, revPointIssuerSKID)
		require.NoError(t, err)
		require.NotNil(t, point)
		require.Equal(t, int32(revPointVid65522), point.Vid)
		require.Equal(t, revPointLabelPAI, point.Label)
	})

	t.Run("AddRevocationPointWithDelegator", func(t *testing.T) {
		// Add PAI cert to ledger
		txResult, err := AddX509Cert(delegatorCertWithVid65521Path, vendorAccount)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

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
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		point, err := GetPkiRevocationDistributionPoint(revPointVid, revPointLabelLeafDel, delegatorCertSubjectKeyID)
		require.NoError(t, err)
		require.NotNil(t, point)
		require.Equal(t, int32(revPointVid), point.Vid)
		require.Equal(t, revPointLabelLeafDel, point.Label)
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
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)
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
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		point, err := GetPkiRevocationDistributionPoint(revPointVid, revPointLabelLeafDel, delegatorCertSubjectKeyID)
		require.NoError(t, err)
		require.NotNil(t, point)
		require.Equal(t, dataURLNew, point.DataURL)

		// Update non-VID-scoped PAI (uses addVendorIntermCertPath / revPointGsr4IssuerSKID)
		dataURLNonScopedNew := revPointDataURLNonScoped + "_new"
		txResult, err = UpdateRevocationPoint(vendorAccount, RevocationPointOpts{
			VID:                revPointVid,
			Label:              revPointLabelIntermediate,
			Certificate:        addVendorIntermCertPath,
			DataURL:            dataURLNonScopedNew,
			IssuerSubjectKeyID: revPointGsr4IssuerSKID,
		})
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

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
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Update failure: point not found
		txResult, err = UpdateRevocationPoint(vendorAccount65522, RevocationPointOpts{
			VID:                revPointVid65522,
			Certificate:        paiCertWithNumericVidPidPath,
			Label:              revPointLabel,
			DataURL:            revPointDataURL,
			IssuerSubjectKeyID: revPointIssuerSKID,
		})
		require.NoError(t, err)
		require.NotEqual(t, uint32(0), txResult.Code)

		// Update failure: sender not vendor
		txResult, err = UpdateRevocationPoint(jack, RevocationPointOpts{
			VID:                revPointVid,
			Certificate:        paaCertWithNumericVidPath,
			Label:              revPointLabel,
			DataURL:            revPointDataURL,
			IssuerSubjectKeyID: revPointIssuerSKID,
		})
		require.NoError(t, err)
		require.NotEqual(t, uint32(0), txResult.Code)
	})

	t.Run("DeleteRevocationPoint", func(t *testing.T) {
		txResult, err := DeleteRevocationPoint(vendorAccount, RevocationPointOpts{
			VID:                revPointVid,
			Label:              revPointLabel,
			IssuerSubjectKeyID: revPointIssuerSKID,
		})
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		point, err := GetPkiRevocationDistributionPoint(revPointVid, revPointLabel, revPointIssuerSKID)
		require.NoError(t, err)
		require.Nil(t, point)
	})
}
