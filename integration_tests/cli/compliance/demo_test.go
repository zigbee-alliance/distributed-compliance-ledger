package compliance

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
	cliputils "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/cli/utils"
	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/utils"
)

// TestComplianceDemo translates compliance-demo.sh.
func TestComplianceDemo(t *testing.T) {
	vid := rand.Intn(65534) + 1
	vendorAccount := fmt.Sprintf("vendor_account_%d", vid)
	cliputils.CreateVendorAccount(t, vendorAccount, vid)

	zbAccount := cliputils.CreateAccount(t, "CertificationCenter")
	secondZbAccount := cliputils.CreateAccount(t, "CertificationCenter")

	pid := rand.Intn(65534) + 1
	sv := rand.Intn(65534) + 1
	svs := fmt.Sprintf("%d", rand.Intn(65534)+1)
	certificationDate := "2020-01-01T00:00:01Z"
	zigbeeCertType := "zigbee"
	matterCertType := "matter"
	cdCertID := fmt.Sprintf("cert-%d", rand.Intn(1<<30))
	cdVersionNumber := 1

	t.Run("CertifyUnknownModel_Succeeds", func(t *testing.T) {
		// Certify unknown model (without a model/version record) — should succeed for ZB
		txResult, err := CertifyModel(vid, pid, sv, svs, zigbeeCertType, certificationDate, cdCertID, zbAccount,
			"--cdVersionNumber", fmt.Sprintf("%d", cdVersionNumber))
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		out, err := QueryDeviceSoftwareCompliance(cdCertID)
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"vid":%d`, vid))
		require.Contains(t, string(out), fmt.Sprintf(`"pid":%d`, pid))
		require.Contains(t, string(out), fmt.Sprintf(`"softwareVersion":%d`, sv))

		// Delete the compliance info before creating model
		txResult, err = DeleteComplianceInfo(vid, pid, sv, zigbeeCertType, zbAccount)
		require.NoError(t, err)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)
	})

	t.Run("CreateModelAndVersion", func(t *testing.T) {
		txResult, err := utils.ExecuteTx("tx", "model", "add-model",
			"--vid", fmt.Sprintf("%d", vid),
			"--pid", fmt.Sprintf("%d", pid),
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
			"--vid", fmt.Sprintf("%d", vid),
			"--pid", fmt.Sprintf("%d", pid),
			"--softwareVersion", fmt.Sprintf("%d", sv),
			"--softwareVersionString", svs,
			"--cdVersionNumber", fmt.Sprintf("%d", cdVersionNumber),
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
		out, err := QueryComplianceInfo(vid, pid, sv, zigbeeCertType)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")

		out, err = QueryDeviceSoftwareCompliance(cdCertID)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")

		out, err = QueryCertifiedModel(vid, pid, sv, zigbeeCertType)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")

		out, err = QueryRevokedModel(vid, pid, sv, zigbeeCertType)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")

		out, err = QueryProvisionalModel(vid, pid, sv, zigbeeCertType)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")

		// QueryAll checks: verify this vid/pid is NOT present (chain may have other data from other test runs)
		out, err = QueryAllComplianceInfo()
		require.NoError(t, err)
		require.NotContains(t, string(out), fmt.Sprintf(`"vid":%d,"pid":%d`, vid, pid))

		out, err = QueryAllDeviceSoftwareCompliance()
		require.NoError(t, err)
		require.NotContains(t, string(out), cdCertID)

		out, err = QueryAllCertifiedModels()
		require.NoError(t, err)
		require.NotContains(t, string(out), fmt.Sprintf(`"vid":%d,"pid":%d`, vid, pid))

		out, err = QueryAllRevokedModels()
		require.NoError(t, err)
		require.NotContains(t, string(out), fmt.Sprintf(`"vid":%d,"pid":%d`, vid, pid))

		out, err = QueryAllProvisionalModels()
		require.NoError(t, err)
		require.NotContains(t, string(out), fmt.Sprintf(`"vid":%d,"pid":%d`, vid, pid))
	})

	t.Run("CertifyWithInvalidSVS_Fails", func(t *testing.T) {
		invalidSvs := fmt.Sprintf("%d", rand.Intn(65534)+1)
		txResult, err := CertifyModel(vid, pid, sv, invalidSvs, zigbeeCertType, certificationDate, cdCertID, zbAccount,
			"--cdVersionNumber", fmt.Sprintf("%d", cdVersionNumber))
		require.NoError(t, err)
		require.Equal(t, uint32(306), txResult.Code)
		_, _ = utils.AwaitTxConfirmation(txResult.TxHash)
	})

	t.Run("CertifyWithInvalidCDVersionNumber_Fails", func(t *testing.T) {
		// cdVersionNumber 0 doesn't match ledger value
		txResult, err := CertifyModel(vid, pid, sv, svs, zigbeeCertType, certificationDate, cdCertID, zbAccount)
		require.NoError(t, err)
		require.Equal(t, uint32(306), txResult.Code)
		require.Contains(t, txResult.RawLog, "ledger does not have matching CDVersionNumber=0")
		_, _ = utils.AwaitTxConfirmation(txResult.TxHash)
	})

	t.Run("CertifyZigbee_Success", func(t *testing.T) {
		txResult, err := CertifyModel(vid, pid, sv, svs, zigbeeCertType, certificationDate, cdCertID, zbAccount,
			"--cdVersionNumber", fmt.Sprintf("%d", cdVersionNumber),
			"--schemaVersion", "0",
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)
	})

	t.Run("CertifyMatter_Success", func(t *testing.T) {
		txResult, err := CertifyModel(vid, pid, sv, svs, matterCertType, certificationDate, cdCertID, zbAccount,
			"--cdVersionNumber", fmt.Sprintf("%d", cdVersionNumber),
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)
	})

	t.Run("ReCertify_DifferentAccount_Fails", func(t *testing.T) {
		txResult, err := CertifyModel(vid, pid, sv, svs, zigbeeCertType, certificationDate, cdCertID, secondZbAccount,
			"--cdVersionNumber", fmt.Sprintf("%d", cdVersionNumber),
		)
		require.NoError(t, err)
		require.Equal(t, uint32(303), txResult.Code)
		require.Contains(t, txResult.RawLog, "already certified on the ledger")
		_, _ = utils.AwaitTxConfirmation(txResult.TxHash)
	})

	t.Run("ReCertify_SameAccount_Fails", func(t *testing.T) {
		txResult, err := CertifyModel(vid, pid, sv, svs, zigbeeCertType, certificationDate, cdCertID, zbAccount,
			"--cdVersionNumber", fmt.Sprintf("%d", cdVersionNumber),
		)
		require.NoError(t, err)
		require.Equal(t, uint32(303), txResult.Code)
		require.Contains(t, txResult.RawLog, "already certified on the ledger")
		_, _ = utils.AwaitTxConfirmation(txResult.TxHash)
	})

	t.Run("QueryAfterCertification", func(t *testing.T) {
		out, err := QueryCertifiedModel(vid, pid, sv, matterCertType)
		require.NoError(t, err)
		require.Contains(t, string(out), `"value":true`)
		require.Contains(t, string(out), fmt.Sprintf(`"pid":%d`, pid))
		require.Contains(t, string(out), fmt.Sprintf(`"vid":%d`, vid))

		out, err = QueryCertifiedModel(vid, pid, sv, zigbeeCertType)
		require.NoError(t, err)
		require.Contains(t, string(out), `"value":true`)

		// After certifying, revoked-model and provisional-model records exist with value=false
		out, err = QueryRevokedModel(vid, pid, sv, matterCertType)
		require.NoError(t, err)
		require.Contains(t, string(out), `"value":false`)

		out, err = QueryRevokedModel(vid, pid, sv, zigbeeCertType)
		require.NoError(t, err)
		require.Contains(t, string(out), `"value":false`)

		out, err = QueryProvisionalModel(vid, pid, sv, matterCertType)
		require.NoError(t, err)
		require.Contains(t, string(out), `"value":false`)

		out, err = QueryProvisionalModel(vid, pid, sv, zigbeeCertType)
		require.NoError(t, err)
		require.Contains(t, string(out), `"value":false`)

		out, err = QueryComplianceInfo(vid, pid, sv, zigbeeCertType)
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"vid":%d`, vid))
		require.Contains(t, string(out), fmt.Sprintf(`"pid":%d`, pid))
		require.Contains(t, string(out), `"softwareVersionCertificationStatus":2`)
		require.Contains(t, string(out), fmt.Sprintf(`"cDCertificateId":"%s"`, cdCertID))
		require.Contains(t, string(out), fmt.Sprintf(`"date":"%s"`, certificationDate))
		require.Contains(t, string(out), fmt.Sprintf(`"certificationType":"%s"`, zigbeeCertType))
		require.Contains(t, string(out), `"schemaVersion":0`)

		out, err = QueryDeviceSoftwareCompliance(cdCertID)
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"vid":%d`, vid))
		require.Contains(t, string(out), fmt.Sprintf(`"pid":%d`, pid))
		require.Contains(t, string(out), `"softwareVersionCertificationStatus":2`)
		require.Contains(t, string(out), fmt.Sprintf(`"certificationType":"%s"`, zigbeeCertType))
		require.Contains(t, string(out), fmt.Sprintf(`"certificationType":"%s"`, matterCertType))

		out, err = QueryAllCertifiedModels()
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"vid":%d`, vid))
		require.Contains(t, string(out), fmt.Sprintf(`"pid":%d`, pid))

		out, err = QueryAllRevokedModels()
		require.NoError(t, err)
		require.NotContains(t, string(out), fmt.Sprintf(`"vid":%d,"pid":%d`, vid, pid))

		out, err = QueryAllProvisionalModels()
		require.NoError(t, err)
		require.NotContains(t, string(out), fmt.Sprintf(`"vid":%d,"pid":%d`, vid, pid))

		out, err = QueryAllComplianceInfo()
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"vid":%d`, vid))
		require.Contains(t, string(out), fmt.Sprintf(`"pid":%d`, pid))
	})

	t.Run("RevokeFromPast_Fails", func(t *testing.T) {
		pastDate := "2020-01-01T00:00:00Z"
		txResult, err := RevokeModel(vid, pid, sv, svs, zigbeeCertType, pastDate, zbAccount,
			"--reason", "some reason",
			"--cdVersionNumber", fmt.Sprintf("%d", cdVersionNumber),
		)
		require.NoError(t, err)
		require.Equal(t, uint32(302), txResult.Code)
		require.Contains(t, txResult.RawLog, "must be after")
		_, _ = utils.AwaitTxConfirmation(txResult.TxHash)
	})

	revocationDate := "2020-02-02T02:20:20Z"
	revocationReason := "some reason"

	t.Run("RevokeZigbee_Success", func(t *testing.T) {
		txResult, err := RevokeModel(vid, pid, sv, svs, zigbeeCertType, revocationDate, zbAccount,
			"--reason", revocationReason,
			"--cdVersionNumber", fmt.Sprintf("%d", cdVersionNumber),
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)
	})

	t.Run("QueryAfterRevocation", func(t *testing.T) {
		out, err := QueryComplianceInfo(vid, pid, sv, zigbeeCertType)
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

		out, err = QueryRevokedModel(vid, pid, sv, zigbeeCertType)
		require.NoError(t, err)
		require.Contains(t, string(out), `"value":true`)

		out, err = QueryCertifiedModel(vid, pid, sv, zigbeeCertType)
		require.NoError(t, err)
		require.Contains(t, string(out), `"value":false`)

		out, err = QueryAllRevokedModels()
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"vid":%d`, vid))
		require.Contains(t, string(out), fmt.Sprintf(`"pid":%d`, pid))
	})
}
