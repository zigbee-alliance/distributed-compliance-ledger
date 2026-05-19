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
	delegatorCertSubjectKeyID         = "0E8CE8C8B8AA50BC258556B9B19CC2C7D9C52F17"

	crlSignerDelegatedByPAI1Path = "../../constants/leaf_cert_with_vid_65521"

	// Use google_root_cert_r1 instead of root_cert to avoid conflict: TestPKIDemo revokes root_cert
	// and the unique-cert store retains the entry, permanently blocking re-addition.
	revPointsTestRootCertPath         = "../../constants/google_root_cert_r1"
	revPointsTestRootCertSubject      = "MEcxCzAJBgNVBAYTAlVTMSIwIAYDVQQKExlHb29nbGUgVHJ1c3QgU2VydmljZXMgTExDMRQwEgYDVQQDEwtHVFMgUm9vdCBSMQ=="
	revPointsTestRootCertSubjectKeyID = "E4:AF:2B:26:71:1A:2B:48:27:85:2F:52:66:2C:EF:F0:89:13:71:3E"

	testRootCertPath         = "../../constants/test_root_cert"
	testRootCertSubject      = "MDAxGDAWBgNVBAMMD01hdHRlciBUZXN0IFBBQTEUMBIGCisGAQQBgqJ8AgEMBDEyNUQ="
	testRootCertSubjectKeyID = "E2:90:8D:36:9C:3C:A3:C1:13:BB:09:E2:4D:C1:CC:C5:A6:66:91:D4"

	rootCertWithVidRevPath         = "../../constants/root_cert_with_vid"
	rootCertWithVidRevSubject      = "MIGYMQswCQYDVQQGEwJVUzERMA8GA1UECAwITmV3IFlvcmsxETAPBgNVBAcMCE5ldyBZb3JrMRgwFgYDVQQKDA9FeGFtcGxlIENvbXBhbnkxGTAXBgNVBAsMEFRlc3RpbmcgRGl2aXNpb24xGDAWBgNVBAMMD3d3dy5leGFtcGxlLmNvbTEUMBIGCisGAQQBgqJ8AgEMBEZGRjE="
	rootCertWithVidRevSubjectKeyID = "CE:A8:92:66:EA:E0:80:BD:2B:B5:68:E4:0B:07:C4:FA:2C:34:6D:31"

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
	revPointIssuerSKID       = "5A880E6C3653D07FB08971A3F473790930E62BDB"
	// SKID of google_root_cert_gsr4 (no colons) — used for non-VID-scoped PAI revocation point.
	// intermediate_cert_gsr4 (no VID) chains to this root, which is on the ledger after TestPKIAddVendorX509Certificates.
	revPointGsr4IssuerSKID = "54B07BAD45B8E2407FFB0A6EFBBE33C93CA384D5"
)

// TestPKIRevocationPoints translates pki-revocation-points.sh.
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
		out, err := QueryAllPkiRevocationDistributionPoints()
		require.NoError(t, err)
		require.Contains(t, string(out), "[]")
	})

	t.Run("QueryRevocationPointNotFound", func(t *testing.T) {
		out, err := QueryPkiRevocationDistributionPoint(revPointVid, revPointLabel, "AB")
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")

		out, err = QueryPkiRevocationDistributionPointsByIssuer("AB")
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")
	})

	t.Run("AddRevocationPointFailures", func(t *testing.T) {
		// Not by vendor
		txResult, err := AddRevocationPoint(jack,
			"--vid", fmt.Sprintf("%d", revPointVid),
			"--is-paa=true",
			"--certificate", paaCertWithNumericVidPath,
			"--label", revPointLabel,
			"--data-url", revPointDataURL,
			"--issuer-subject-key-id", revPointIssuerSKID,
			"--revocation-type", "1",
		)
		require.NoError(t, err)
		require.NotEqual(t, uint32(0), txResult.Code)

		// Sender VID not equal to field VID
		txResult, err = AddRevocationPoint(vendorAccount65522,
			"--vid", fmt.Sprintf("%d", revPointVid),
			"--is-paa=true",
			"--certificate", paaCertWithNumericVidPath,
			"--label", revPointLabel,
			"--data-url", revPointDataURL,
			"--issuer-subject-key-id", revPointIssuerSKID,
			"--revocation-type", "1",
		)
		require.NoError(t, err)
		require.NotEqual(t, uint32(0), txResult.Code)

		// Certificate does not exist on ledger
		txResult, err = AddRevocationPoint(vendorAccount,
			"--vid", fmt.Sprintf("%d", revPointVid),
			"--is-paa=true",
			"--certificate", paaCertWithNumericVidPath,
			"--label", revPointLabel,
			"--data-url", revPointDataURL,
			"--issuer-subject-key-id", revPointIssuerSKID,
			"--revocation-type", "1",
		)
		require.NoError(t, err)
		require.NotEqual(t, uint32(0), txResult.Code)
	})

	t.Run("AddCertsToLedger", func(t *testing.T) {
		// Trustees add PAA cert with numeric VID
		txResult, err := ProposeAddX509RootCert(paaCertWithNumericVidPath, jack,
			"--vid", fmt.Sprintf("%d", revPointVid),
		)
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
		txResult, err = ProposeAddX509RootCert(paaCertNoVidPath, jack,
			"--vid", fmt.Sprintf("%d", revPointVidNonScoped),
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		txResult, err = ApproveAddX509RootCert(paaCertNoVidSubject, paaCertNoVidSubjectKeyID, alice)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Add root cert
		txResult, err = ProposeAddX509RootCert(revPointsTestRootCertPath, jack,
			"--vid", fmt.Sprintf("%d", revPointVid),
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		txResult, err = ApproveAddX509RootCert(revPointsTestRootCertSubject, revPointsTestRootCertSubjectKeyID, alice)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Add test root cert
		txResult, err = ProposeAddX509RootCert(testRootCertPath, jack,
			"--vid", fmt.Sprintf("%d", revPointVidNonScoped),
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		txResult, err = ApproveAddX509RootCert(testRootCertSubject, testRootCertSubjectKeyID, alice)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Add VID-scoped root cert
		txResult, err = ProposeAddX509RootCert(rootCertWithVidRevPath, jack,
			"--vid", fmt.Sprintf("%d", revPointVid),
		)
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

	t.Run("AddRevocationPointForVidScopedPAA", func(t *testing.T) {
		// CRL signer cert PEM value not equal to stored — should fail
		txResult, err := AddRevocationPoint(vendorAccount65522,
			"--vid", fmt.Sprintf("%d", revPointVid65522),
			"--is-paa=true",
			"--certificate", paaCertWithNumericVid1Path,
			"--label", revPointLabel,
			"--data-url", revPointDataURL,
			"--issuer-subject-key-id", revPointIssuerSKID,
			"--revocation-type", "1",
		)
		require.NoError(t, err)
		require.NotEqual(t, uint32(0), txResult.Code)

		// Add for VID-scoped PAA
		txResult, err = AddRevocationPoint(vendorAccount,
			"--vid", fmt.Sprintf("%d", revPointVid),
			"--is-paa=true",
			"--certificate", paaCertWithNumericVidPath,
			"--label", revPointLabel,
			"--data-url", revPointDataURL,
			"--issuer-subject-key-id", revPointIssuerSKID,
			"--revocation-type", "1",
			"--schemaVersion", "0",
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		out, err := QueryAllPkiRevocationDistributionPoints()
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"vid":%d`, revPointVid))
		require.Contains(t, string(out), fmt.Sprintf(`"label":"%s"`, revPointLabel))

		// Cannot add same point twice (same vid, issuer, label)
		txResult, err = AddRevocationPoint(vendorAccount,
			"--vid", fmt.Sprintf("%d", revPointVid),
			"--is-paa=true",
			"--certificate", paaCertWithNumericVidPath,
			"--label", revPointLabel,
			"--data-url", revPointDataURL+"-new",
			"--issuer-subject-key-id", revPointIssuerSKID,
			"--revocation-type", "1",
		)
		require.NoError(t, err)
		require.NotEqual(t, uint32(0), txResult.Code)
	})

	t.Run("AddRevocationPointForNonVidScopedPAA", func(t *testing.T) {
		txResult, err := AddRevocationPoint(vendorAccountNonScoped,
			"--vid", fmt.Sprintf("%d", revPointVidNonScoped),
			"--is-paa=true",
			"--certificate", paaCertNoVidPath,
			"--label", revPointLabelNonScoped,
			"--data-url", revPointDataURLNonScoped,
			"--issuer-subject-key-id", revPointIssuerSKID,
			"--revocation-type", "1",
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		out, err := QueryAllPkiRevocationDistributionPoints()
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"vid":%d`, revPointVid))
		require.Contains(t, string(out), fmt.Sprintf(`"vid":%d`, revPointVidNonScoped))
	})

	t.Run("AddRevocationPointForPAI", func(t *testing.T) {
		txResult, err := AddRevocationPoint(vendorAccount65522,
			"--vid", fmt.Sprintf("%d", revPointVid65522),
			"--is-paa=false",
			"--certificate", paiCertWithNumericVidPath,
			"--label", revPointLabelPAI,
			"--data-url", revPointDataURL,
			"--issuer-subject-key-id", revPointIssuerSKID,
			"--revocation-type", "1",
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		out, err := QueryPkiRevocationDistributionPoint(revPointVid65522, revPointLabelPAI, revPointIssuerSKID)
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"vid":%d`, revPointVid65522))
		require.Contains(t, string(out), fmt.Sprintf(`"label":"%s"`, revPointLabelPAI))
	})

	t.Run("AddRevocationPointWithDelegator", func(t *testing.T) {
		// Add PAI cert to ledger
		txResult, err := AddX509Cert(delegatorCertWithVid65521Path, vendorAccount)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Add revocation point with delegator
		txResult, err = AddRevocationPoint(vendorAccount,
			"--vid", fmt.Sprintf("%d", revPointVid),
			"--is-paa=false",
			"--certificate", crlSignerDelegatedByPAI1Path,
			"--label", revPointLabelLeafDel,
			"--data-url", revPointDataURL,
			"--issuer-subject-key-id", delegatorCertSubjectKeyID,
			"--revocation-type", "1",
			"--certificate-delegator", delegatorCertWithVid65521Path,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		out, err := QueryPkiRevocationDistributionPoint(revPointVid, revPointLabelLeafDel, delegatorCertSubjectKeyID)
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"vid":%d`, revPointVid))
		require.Contains(t, string(out), fmt.Sprintf(`"label":"%s"`, revPointLabelLeafDel))
	})

	t.Run("AddRevocationPointForNonVidScopedPAI", func(t *testing.T) {
		// Use intermediate_cert_gsr4 (no VID) instead of intermediate_cert: root_cert (issuer of
		// intermediate_cert) is revoked by TestPKIDemo, so its chain cannot be verified.
		// google_root_cert_gsr4 (issuer of intermediate_cert_gsr4) is already on the ledger
		// from TestPKIAddVendorX509Certificates.
		txResult, err := AddRevocationPoint(vendorAccount,
			"--vid", fmt.Sprintf("%d", revPointVid),
			"--is-paa=false",
			"--certificate", addVendorIntermCertPath,
			"--label", revPointLabelIntermediate,
			"--data-url", revPointDataURLNonScoped,
			"--issuer-subject-key-id", revPointGsr4IssuerSKID,
			"--revocation-type", "1",
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)
	})

	t.Run("UpdateRevocationPoints", func(t *testing.T) {
		dataURLNew := revPointDataURL + "_new"

		// Update with delegator
		txResult, err := UpdateRevocationPoint(vendorAccount,
			"--vid", fmt.Sprintf("%d", revPointVid),
			"--certificate", crlSignerDelegatedByPAI1Path,
			"--label", revPointLabelLeafDel,
			"--data-url", dataURLNew,
			"--issuer-subject-key-id", delegatorCertSubjectKeyID,
			"--certificate-delegator", delegatorCertWithVid65521CopyPath,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		out, err := QueryPkiRevocationDistributionPoint(revPointVid, revPointLabelLeafDel, delegatorCertSubjectKeyID)
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"dataURL":"%s"`, dataURLNew))

		// Update non-VID-scoped PAI (uses addVendorIntermCertPath / revPointGsr4IssuerSKID)
		dataURLNonScopedNew := revPointDataURLNonScoped + "_new"
		txResult, err = UpdateRevocationPoint(vendorAccount,
			"--vid", fmt.Sprintf("%d", revPointVid),
			"--label", revPointLabelIntermediate,
			"--certificate", addVendorIntermCertPath,
			"--data-url", dataURLNonScopedNew,
			"--issuer-subject-key-id", revPointGsr4IssuerSKID,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		out, err = QueryPkiRevocationDistributionPoint(revPointVid, revPointLabelIntermediate, revPointGsr4IssuerSKID)
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"dataURL":"%s"`, dataURLNonScopedNew))

		// Update VID-scoped PAA (use revPointsTestRootCertPath which is on-ledger and not revoked)
		txResult, err = UpdateRevocationPoint(vendorAccount,
			"--vid", fmt.Sprintf("%d", revPointVid),
			"--certificate", revPointsTestRootCertPath,
			"--label", revPointLabel,
			"--data-url", revPointDataURL,
			"--issuer-subject-key-id", revPointIssuerSKID,
			"--schemaVersion", "0",
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Update failure: point not found
		txResult, err = UpdateRevocationPoint(vendorAccount65522,
			"--vid", fmt.Sprintf("%d", revPointVid65522),
			"--certificate", paiCertWithNumericVidPidPath,
			"--label", revPointLabel,
			"--data-url", revPointDataURL,
			"--issuer-subject-key-id", revPointIssuerSKID,
		)
		require.NoError(t, err)
		require.NotEqual(t, uint32(0), txResult.Code)

		// Update failure: sender not vendor
		txResult, err = UpdateRevocationPoint(jack,
			"--vid", fmt.Sprintf("%d", revPointVid),
			"--certificate", paaCertWithNumericVidPath,
			"--label", revPointLabel,
			"--data-url", revPointDataURL,
			"--issuer-subject-key-id", revPointIssuerSKID,
		)
		require.NoError(t, err)
		require.NotEqual(t, uint32(0), txResult.Code)
	})

	t.Run("DeleteRevocationPoint", func(t *testing.T) {
		txResult, err := DeleteRevocationPoint(vendorAccount,
			"--vid", fmt.Sprintf("%d", revPointVid),
			"--label", revPointLabel,
			"--issuer-subject-key-id", revPointIssuerSKID,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		out, err := QueryPkiRevocationDistributionPoint(revPointVid, revPointLabel, revPointIssuerSKID)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")
	})
}
