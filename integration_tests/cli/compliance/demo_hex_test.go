package compliance

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
	cliputils "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/cli/utils"
	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/utils"
)

// TestComplianceDemoHex translates compliance-demo-hex.sh.
func TestComplianceDemoHex(t *testing.T) {
	// Use random VID/PID expressed as hex to avoid collisions across test runs.
	vid := rand.Intn(65534) + 1
	pid := rand.Intn(65534) + 1
	vidHex := fmt.Sprintf("0x%X", vid)
	pidHex := fmt.Sprintf("0x%X", pid)

	vendorAccount := fmt.Sprintf("vendor_account_%d", vid)
	cliputils.CreateVendorAccount(t, vendorAccount, vid)

	zbAccount := cliputils.CreateAccount(t, "CertificationCenter")
	secondZbAccount := cliputils.CreateAccount(t, "CertificationCenter")

	sv := rand.Intn(65534) + 1
	svs := fmt.Sprintf("%d", rand.Intn(65534)+1)
	certificationDate := "2020-01-01T00:00:01Z"
	zigbeeCertType := "zigbee"
	matterCertType := "matter"
	cdCertID := fmt.Sprintf("cert-%014d", rand.Intn(1<<30))

	t.Run("AddModelAndVersion", func(t *testing.T) {
		txResult, err := utils.ExecuteTx("tx", "model", "add-model",
			"--vid", vidHex,
			"--pid", pidHex,
			"--deviceTypeID", "1",
			"--productName", "TestProduct",
			"--productLabel", "TestingProductLabel",
			"--partNumber", "1",
			"--commissioningCustomFlow", "0",
			"--enhancedSetupFlowOptions", "0",
			"--from", vendorAccount,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		txResult, err = utils.ExecuteTx("tx", "model", "add-model-version",
			"--vid", vidHex,
			"--pid", pidHex,
			"--softwareVersion", fmt.Sprintf("%d", sv),
			"--softwareVersionString", svs,
			"--cdVersionNumber", "1",
			"--maxApplicableSoftwareVersion", "10",
			"--minApplicableSoftwareVersion", "1",
			"--from", vendorAccount,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)
	})

	t.Run("QueryBeforeCertification_NotFound", func(t *testing.T) {
		out, err := utils.ExecuteCLI("query", "compliance", "compliance-info",
			"--vid", vidHex, "--pid", pidHex,
			"--softwareVersion", fmt.Sprintf("%d", sv),
			"--certificationType", "zigbee", "-o", "json",
		)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")
		require.NotContains(t, string(out), fmt.Sprintf(`"pid":%d`, pid))

		out, err = QueryDeviceSoftwareCompliance(cdCertID)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")

		out, err = utils.ExecuteCLI("query", "compliance", "certified-model",
			"--vid", vidHex, "--pid", pidHex,
			"--softwareVersion", fmt.Sprintf("%d", sv),
			"--certificationType", "zigbee", "-o", "json",
		)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")

		out, err = utils.ExecuteCLI("query", "compliance", "revoked-model",
			"--vid", vidHex, "--pid", pidHex,
			"--softwareVersion", fmt.Sprintf("%d", sv),
			"--certificationType", "zigbee", "-o", "json",
		)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")

		out, err = utils.ExecuteCLI("query", "compliance", "provisional-model",
			"--vid", vidHex, "--pid", pidHex,
			"--softwareVersion", fmt.Sprintf("%d", sv),
			"--certificationType", "zigbee", "-o", "json",
		)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")
	})

	t.Run("CertifyWithInvalidSVS_Fails", func(t *testing.T) {
		invalidSvs := fmt.Sprintf("%d", rand.Intn(65534)+1)
		txResult, err := utils.ExecuteTx("tx", "compliance", "certify-model",
			"--vid", vidHex, "--pid", pidHex,
			"--softwareVersion", fmt.Sprintf("%d", sv),
			"--softwareVersionString", invalidSvs,
			"--certificationType", zigbeeCertType,
			"--certificationDate", certificationDate,
			"--cdCertificateId", cdCertID,
			"--from", zbAccount,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(306), txResult.Code)
		_, _ = utils.AwaitTxConfirmation(txResult.TxHash)
	})

	t.Run("CertifyZigbee_WithHexVID_Success", func(t *testing.T) {
		txResult, err := utils.ExecuteTx("tx", "compliance", "certify-model",
			"--vid", vidHex, "--pid", pidHex,
			"--softwareVersion", fmt.Sprintf("%d", sv),
			"--softwareVersionString", svs,
			"--certificationType", zigbeeCertType,
			"--certificationDate", certificationDate,
			"--cdCertificateId", cdCertID,
			"--cdVersionNumber", "1",
			"--from", zbAccount,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)
	})

	t.Run("CertifyMatter_WithHexVID_Success", func(t *testing.T) {
		txResult, err := utils.ExecuteTx("tx", "compliance", "certify-model",
			"--vid", vidHex, "--pid", pidHex,
			"--softwareVersion", fmt.Sprintf("%d", sv),
			"--softwareVersionString", svs,
			"--certificationType", matterCertType,
			"--certificationDate", certificationDate,
			"--cdCertificateId", cdCertID,
			"--cdVersionNumber", "1",
			"--from", zbAccount,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)
	})

	t.Run("ReCertify_DifferentAccount_Fails", func(t *testing.T) {
		txResult, err := utils.ExecuteTx("tx", "compliance", "certify-model",
			"--vid", vidHex, "--pid", pidHex,
			"--softwareVersion", fmt.Sprintf("%d", sv),
			"--softwareVersionString", svs,
			"--certificationType", zigbeeCertType,
			"--certificationDate", certificationDate,
			"--cdCertificateId", cdCertID,
			"--cdVersionNumber", "1",
			"--from", secondZbAccount,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(303), txResult.Code)
		require.Contains(t, txResult.RawLog, "already certified on the ledger")
	})

	t.Run("QueryAfterCertification", func(t *testing.T) {
		out, err := utils.ExecuteCLI("query", "compliance", "certified-model",
			"--vid", vidHex, "--pid", pidHex,
			"--softwareVersion", fmt.Sprintf("%d", sv),
			"--certificationType", matterCertType, "-o", "json",
		)
		require.NoError(t, err)
		require.Contains(t, string(out), `"value":true`)
		require.Contains(t, string(out), fmt.Sprintf(`"pid":%d`, pid))
		require.Contains(t, string(out), fmt.Sprintf(`"vid":%d`, vid))

		out, err = utils.ExecuteCLI("query", "compliance", "compliance-info",
			"--vid", vidHex, "--pid", pidHex,
			"--softwareVersion", fmt.Sprintf("%d", sv),
			"--certificationType", zigbeeCertType, "-o", "json",
		)
		require.NoError(t, err)
		require.Contains(t, string(out), `"softwareVersionCertificationStatus":2`)
		require.Contains(t, string(out), fmt.Sprintf(`"date":"%s"`, certificationDate))

		out, err = QueryDeviceSoftwareCompliance(cdCertID)
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"vid":%d`, vid))
		require.Contains(t, string(out), fmt.Sprintf(`"pid":%d`, pid))
		require.Contains(t, string(out), `"softwareVersionCertificationStatus":2`)
		require.Contains(t, string(out), fmt.Sprintf(`"certificationType":"%s"`, matterCertType))
		require.Contains(t, string(out), fmt.Sprintf(`"certificationType":"%s"`, zigbeeCertType))
	})

	revocationDate := "2020-02-02T02:20:20Z"
	revocationReason := "some reason"

	t.Run("RevokeFromPast_Fails", func(t *testing.T) {
		pastDate := "2020-01-01T00:00:00Z"
		txResult, err := utils.ExecuteTx("tx", "compliance", "revoke-model",
			"--vid", vidHex, "--pid", pidHex,
			"--softwareVersion", fmt.Sprintf("%d", sv),
			"--softwareVersionString", svs,
			"--certificationType", zigbeeCertType,
			"--revocationDate", pastDate,
			"--reason", revocationReason,
			"--cdVersionNumber", "1",
			"--from", zbAccount,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(302), txResult.Code)
		require.Contains(t, txResult.RawLog, "must be after")
	})

	t.Run("RevokeZigbee_Success", func(t *testing.T) {
		txResult, err := utils.ExecuteTx("tx", "compliance", "revoke-model",
			"--vid", vidHex, "--pid", pidHex,
			"--softwareVersion", fmt.Sprintf("%d", sv),
			"--softwareVersionString", svs,
			"--certificationType", zigbeeCertType,
			"--revocationDate", revocationDate,
			"--reason", revocationReason,
			"--cdVersionNumber", "1",
			"--from", zbAccount,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)
	})

	t.Run("QueryAfterRevocation", func(t *testing.T) {
		out, err := utils.ExecuteCLI("query", "compliance", "compliance-info",
			"--vid", vidHex, "--pid", pidHex,
			"--softwareVersion", fmt.Sprintf("%d", sv),
			"--certificationType", zigbeeCertType, "-o", "json",
		)
		require.NoError(t, err)
		require.Contains(t, string(out), `"softwareVersionCertificationStatus":3`)
		require.Contains(t, string(out), fmt.Sprintf(`"date":"%s"`, revocationDate))
		require.Contains(t, string(out), fmt.Sprintf(`"reason":"%s"`, revocationReason))
		require.Contains(t, string(out), "history")

		out, err = QueryDeviceSoftwareCompliance(cdCertID)
		require.NoError(t, err)
		require.Contains(t, string(out), `"softwareVersionCertificationStatus":2`)
		require.Contains(t, string(out), fmt.Sprintf(`"certificationType":"%s"`, matterCertType))
		require.NotContains(t, string(out), fmt.Sprintf(`"certificationType":"%s"`, zigbeeCertType))

		out, err = utils.ExecuteCLI("query", "compliance", "revoked-model",
			"--vid", vidHex, "--pid", pidHex,
			"--softwareVersion", fmt.Sprintf("%d", sv),
			"--certificationType", zigbeeCertType, "-o", "json",
		)
		require.NoError(t, err)
		require.Contains(t, string(out), `"value":true`)
	})

	t.Run("ReCertifyAfterRevoke_Success", func(t *testing.T) {
		newCertDate := "2020-03-03T00:00:00Z"
		txResult, err := utils.ExecuteTx("tx", "compliance", "certify-model",
			"--vid", vidHex, "--pid", pidHex,
			"--softwareVersion", fmt.Sprintf("%d", sv),
			"--softwareVersionString", svs,
			"--certificationType", zigbeeCertType,
			"--certificationDate", newCertDate,
			"--cdCertificateId", cdCertID,
			"--cdVersionNumber", "1",
			"--from", zbAccount,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		out, err := utils.ExecuteCLI("query", "compliance", "compliance-info",
			"--vid", vidHex, "--pid", pidHex,
			"--softwareVersion", fmt.Sprintf("%d", sv),
			"--certificationType", zigbeeCertType, "-o", "json",
		)
		require.NoError(t, err)
		require.Contains(t, string(out), `"softwareVersionCertificationStatus":2`)
		require.Contains(t, string(out), fmt.Sprintf(`"date":"%s"`, newCertDate))

		out, err = utils.ExecuteCLI("query", "compliance", "certified-model",
			"--vid", vidHex, "--pid", pidHex,
			"--softwareVersion", fmt.Sprintf("%d", sv),
			"--certificationType", zigbeeCertType, "-o", "json",
		)
		require.NoError(t, err)
		require.Contains(t, string(out), `"value":true`)
	})
}
