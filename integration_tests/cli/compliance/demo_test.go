package compliance

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
	cliputils "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/cli/utils"
	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/utils"
)

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
	cdCertID := fmt.Sprintf("cert-%014d", rand.Intn(1<<30))
	cdVersionNumber := 1

	t.Run("CertifyUnknownModel_Succeeds", func(t *testing.T) {
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
		require.Contains(t, string(out), fmt.Sprintf(`"softwareVersionString":"%s"`, svs))
		require.Contains(t, string(out), fmt.Sprintf(`"certificationType":"%s"`, zigbeeCertType))
		require.Contains(t, string(out), fmt.Sprintf(`"date":"%s"`, certificationDate))
		require.Contains(t, string(out), fmt.Sprintf(`"cDCertificateId":"%s"`, cdCertID))

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
		require.NotContains(t, string(out), fmt.Sprintf(`"pid":%d`, pid))
		require.NotContains(t, string(out), fmt.Sprintf(`"vid":%d`, vid))

		out, err = QueryDeviceSoftwareCompliance(cdCertID)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")
		require.NotContains(t, string(out), fmt.Sprintf(`"pid":%d`, pid))
		require.NotContains(t, string(out), fmt.Sprintf(`"vid":%d`, vid))

		out, err = QueryCertifiedModel(vid, pid, sv, zigbeeCertType)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")
		require.NotContains(t, string(out), fmt.Sprintf(`"pid":%d`, pid))
		require.NotContains(t, string(out), fmt.Sprintf(`"vid":%d`, vid))

		out, err = QueryRevokedModel(vid, pid, sv, zigbeeCertType)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")
		require.NotContains(t, string(out), fmt.Sprintf(`"pid":%d`, pid))
		require.NotContains(t, string(out), fmt.Sprintf(`"vid":%d`, vid))

		out, err = QueryProvisionalModel(vid, pid, sv, zigbeeCertType)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")
		require.NotContains(t, string(out), fmt.Sprintf(`"pid":%d`, pid))
		require.NotContains(t, string(out), fmt.Sprintf(`"vid":%d`, vid))

		// QueryAll — this vid/pid must not appear yet
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
		// Certified model checks
		out, err := QueryCertifiedModel(vid, pid, sv, matterCertType)
		require.NoError(t, err)
		require.Contains(t, string(out), `"value":true`)
		require.Contains(t, string(out), fmt.Sprintf(`"pid":%d`, pid))
		require.Contains(t, string(out), fmt.Sprintf(`"vid":%d`, vid))

		out, err = QueryCertifiedModel(vid, pid, sv, zigbeeCertType)
		require.NoError(t, err)
		require.Contains(t, string(out), `"value":true`)
		require.Contains(t, string(out), fmt.Sprintf(`"pid":%d`, pid))
		require.Contains(t, string(out), fmt.Sprintf(`"vid":%d`, vid))

		// After certifying, revoked/provisional records exist with value=false
		out, err = QueryRevokedModel(vid, pid, sv, matterCertType)
		require.NoError(t, err)
		require.Contains(t, string(out), `"value":false`)
		require.Contains(t, string(out), fmt.Sprintf(`"pid":%d`, pid))
		require.Contains(t, string(out), fmt.Sprintf(`"vid":%d`, vid))

		out, err = QueryRevokedModel(vid, pid, sv, zigbeeCertType)
		require.NoError(t, err)
		require.Contains(t, string(out), `"value":false`)

		out, err = QueryProvisionalModel(vid, pid, sv, matterCertType)
		require.NoError(t, err)
		require.Contains(t, string(out), `"value":false`)
		require.Contains(t, string(out), fmt.Sprintf(`"pid":%d`, pid))
		require.Contains(t, string(out), fmt.Sprintf(`"vid":%d`, vid))

		out, err = QueryProvisionalModel(vid, pid, sv, zigbeeCertType)
		require.NoError(t, err)
		require.Contains(t, string(out), `"value":false`)

		// Compliance info — zigbee
		out, err = QueryComplianceInfo(vid, pid, sv, zigbeeCertType)
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"vid":%d`, vid))
		require.Contains(t, string(out), fmt.Sprintf(`"pid":%d`, pid))
		require.Contains(t, string(out), `"softwareVersionCertificationStatus":2`)
		require.Contains(t, string(out), fmt.Sprintf(`"cDCertificateId":"%s"`, cdCertID))
		require.Contains(t, string(out), fmt.Sprintf(`"date":"%s"`, certificationDate))
		require.Contains(t, string(out), fmt.Sprintf(`"certificationType":"%s"`, zigbeeCertType))
		require.Contains(t, string(out), `"schemaVersion":0`)

		// Compliance info — matter
		out, err = QueryComplianceInfo(vid, pid, sv, matterCertType)
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"vid":%d`, vid))
		require.Contains(t, string(out), fmt.Sprintf(`"pid":%d`, pid))
		require.Contains(t, string(out), `"softwareVersionCertificationStatus":2`)
		require.Contains(t, string(out), fmt.Sprintf(`"cDCertificateId":"%s"`, cdCertID))
		require.Contains(t, string(out), fmt.Sprintf(`"date":"%s"`, certificationDate))
		require.Contains(t, string(out), fmt.Sprintf(`"certificationType":"%s"`, matterCertType))
		require.Contains(t, string(out), `"schemaVersion":0`)

		// Device software compliance
		out, err = QueryDeviceSoftwareCompliance(cdCertID)
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"vid":%d`, vid))
		require.Contains(t, string(out), fmt.Sprintf(`"pid":%d`, pid))
		require.Contains(t, string(out), `"softwareVersionCertificationStatus":2`)
		require.Contains(t, string(out), fmt.Sprintf(`"cDCertificateId":"%s"`, cdCertID))
		require.Contains(t, string(out), fmt.Sprintf(`"date":"%s"`, certificationDate))
		require.Contains(t, string(out), fmt.Sprintf(`"certificationType":"%s"`, zigbeeCertType))
		require.Contains(t, string(out), fmt.Sprintf(`"certificationType":"%s"`, matterCertType))

		// All certified models — both cert types
		out, err = QueryAllCertifiedModels()
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"vid":%d`, vid))
		require.Contains(t, string(out), fmt.Sprintf(`"pid":%d`, pid))
		require.Contains(t, string(out), fmt.Sprintf(`"certificationType":"%s"`, zigbeeCertType))
		require.Contains(t, string(out), fmt.Sprintf(`"certificationType":"%s"`, matterCertType))

		out, err = QueryAllRevokedModels()
		require.NoError(t, err)
		require.NotContains(t, string(out), fmt.Sprintf(`"vid":%d,"pid":%d`, vid, pid))

		out, err = QueryAllProvisionalModels()
		require.NoError(t, err)
		require.NotContains(t, string(out), fmt.Sprintf(`"vid":%d,"pid":%d`, vid, pid))

		// All compliance info — both cert types and date
		out, err = QueryAllComplianceInfo()
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"vid":%d`, vid))
		require.Contains(t, string(out), fmt.Sprintf(`"pid":%d`, pid))
		require.Contains(t, string(out), fmt.Sprintf(`"certificationType":"%s"`, zigbeeCertType))
		require.Contains(t, string(out), fmt.Sprintf(`"certificationType":"%s"`, matterCertType))
		require.Contains(t, string(out), fmt.Sprintf(`"date":"%s"`, certificationDate))

		// All device software compliance
		out, err = QueryAllDeviceSoftwareCompliance()
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"vid":%d`, vid))
		require.Contains(t, string(out), fmt.Sprintf(`"pid":%d`, pid))
		require.Contains(t, string(out), fmt.Sprintf(`"certificationType":"%s"`, zigbeeCertType))
		require.Contains(t, string(out), fmt.Sprintf(`"certificationType":"%s"`, matterCertType))
		require.Contains(t, string(out), fmt.Sprintf(`"date":"%s"`, certificationDate))
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
		require.Contains(t, string(out), fmt.Sprintf(`"vid":%d`, vid))
		require.Contains(t, string(out), fmt.Sprintf(`"pid":%d`, pid))
		require.Contains(t, string(out), `"softwareVersionCertificationStatus":3`)
		require.Contains(t, string(out), fmt.Sprintf(`"cDCertificateId":"%s"`, cdCertID))
		require.Contains(t, string(out), fmt.Sprintf(`"date":"%s"`, revocationDate))
		require.Contains(t, string(out), fmt.Sprintf(`"reason":"%s"`, revocationReason))
		require.Contains(t, string(out), fmt.Sprintf(`"certificationType":"%s"`, zigbeeCertType))
		require.Contains(t, string(out), "history")

		// Device software compliance — matter still active, zigbee gone
		out, err = QueryDeviceSoftwareCompliance(cdCertID)
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"vid":%d`, vid))
		require.Contains(t, string(out), fmt.Sprintf(`"pid":%d`, pid))
		require.Contains(t, string(out), `"softwareVersionCertificationStatus":2`)
		require.Contains(t, string(out), fmt.Sprintf(`"cDCertificateId":"%s"`, cdCertID))
		require.Contains(t, string(out), fmt.Sprintf(`"date":"%s"`, certificationDate))
		require.Contains(t, string(out), fmt.Sprintf(`"certificationType":"%s"`, matterCertType))
		require.NotContains(t, string(out), fmt.Sprintf(`"certificationType":"%s"`, zigbeeCertType))

		out, err = QueryRevokedModel(vid, pid, sv, zigbeeCertType)
		require.NoError(t, err)
		require.Contains(t, string(out), `"value":true`)

		out, err = QueryCertifiedModel(vid, pid, sv, zigbeeCertType)
		require.NoError(t, err)
		require.Contains(t, string(out), `"value":false`)

		out, err = QueryProvisionalModel(vid, pid, sv, zigbeeCertType)
		require.NoError(t, err)
		require.Contains(t, string(out), `"value":false`)

		// All revoked — zigbee present
		out, err = QueryAllRevokedModels()
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"vid":%d`, vid))
		require.Contains(t, string(out), fmt.Sprintf(`"pid":%d`, pid))
		require.Contains(t, string(out), fmt.Sprintf(`"certificationType":"%s"`, zigbeeCertType))

		// All certified — matter still present, this test's zigbee entry gone
		out, err = QueryAllCertifiedModels()
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"vid":%d`, vid))
		require.Contains(t, string(out), fmt.Sprintf(`"pid":%d`, pid))
		require.Contains(t, string(out), fmt.Sprintf(`"certificationType":"%s"`, matterCertType))
		require.NotContains(t, string(out), fmt.Sprintf(`"vid":%d,"pid":%d,"softwareVersion":%d,"certificationType":"%s"`, vid, pid, sv, zigbeeCertType))
	})

	reCertDate := "2020-03-03T00:00:00Z"

	t.Run("ReCertifyAfterRevocation_Success", func(t *testing.T) {
		txResult, err := CertifyModel(vid, pid, sv, svs, zigbeeCertType, reCertDate, cdCertID, zbAccount,
			"--cdVersionNumber", fmt.Sprintf("%d", cdVersionNumber),
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Compliance info — back to certified with new date
		out, err := QueryComplianceInfo(vid, pid, sv, zigbeeCertType)
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"vid":%d`, vid))
		require.Contains(t, string(out), fmt.Sprintf(`"pid":%d`, pid))
		require.Contains(t, string(out), `"softwareVersionCertificationStatus":2`)
		require.Contains(t, string(out), fmt.Sprintf(`"cDCertificateId":"%s"`, cdCertID))
		require.Contains(t, string(out), fmt.Sprintf(`"date":"%s"`, reCertDate))
		require.Contains(t, string(out), fmt.Sprintf(`"certificationType":"%s"`, zigbeeCertType))

		// Device software compliance — both cert types again
		out, err = QueryDeviceSoftwareCompliance(cdCertID)
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"vid":%d`, vid))
		require.Contains(t, string(out), fmt.Sprintf(`"pid":%d`, pid))
		require.Contains(t, string(out), `"softwareVersionCertificationStatus":2`)
		require.Contains(t, string(out), fmt.Sprintf(`"cDCertificateId":"%s"`, cdCertID))
		require.Contains(t, string(out), fmt.Sprintf(`"certificationType":"%s"`, zigbeeCertType))
		require.Contains(t, string(out), fmt.Sprintf(`"certificationType":"%s"`, matterCertType))

		out, err = QueryCertifiedModel(vid, pid, sv, zigbeeCertType)
		require.NoError(t, err)
		require.Contains(t, string(out), `"value":true`)

		out, err = QueryRevokedModel(vid, pid, sv, zigbeeCertType)
		require.NoError(t, err)
		require.Contains(t, string(out), `"value":false`)

		// All compliance info — status=2
		out, err = QueryAllComplianceInfo()
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"vid":%d`, vid))
		require.Contains(t, string(out), fmt.Sprintf(`"pid":%d`, pid))
		require.Contains(t, string(out), `"softwareVersionCertificationStatus":2`)

		// All device software compliance — status=2
		out, err = QueryAllDeviceSoftwareCompliance()
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"vid":%d`, vid))
		require.Contains(t, string(out), fmt.Sprintf(`"pid":%d`, pid))
		require.Contains(t, string(out), `"softwareVersionCertificationStatus":2`)

		// All revoked — no longer contains this zigbee entry
		out, err = QueryAllRevokedModels()
		require.NoError(t, err)
		require.NotContains(t, string(out), fmt.Sprintf(`"vid":%d,"pid":%d`, vid, pid))

		// All certified — both cert types
		out, err = QueryAllCertifiedModels()
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"vid":%d`, vid))
		require.Contains(t, string(out), fmt.Sprintf(`"pid":%d`, pid))
		require.Contains(t, string(out), fmt.Sprintf(`"certificationType":"%s"`, matterCertType))
		require.Contains(t, string(out), fmt.Sprintf(`"certificationType":"%s"`, zigbeeCertType))
	})

	// Second pid/sv pair for optional-fields and update/delete tests.
	pid2 := rand.Intn(65534) + 1
	sv2 := rand.Intn(65534) + 1
	svs2 := fmt.Sprintf("%d", rand.Intn(65534)+1)

	t.Run("CreateModelAndVersion2", func(t *testing.T) {
		txResult, err := utils.ExecuteTx("tx", "model", "add-model",
			"--vid", fmt.Sprintf("%d", vid),
			"--pid", fmt.Sprintf("%d", pid2),
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
			"--pid", fmt.Sprintf("%d", pid2),
			"--softwareVersion", fmt.Sprintf("%d", sv2),
			"--softwareVersionString", svs2,
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

	t.Run("CertifyWithOptionalFields", func(t *testing.T) {
		txResult, err := CertifyModel(vid, pid2, sv2, svs2, zigbeeCertType, reCertDate, cdCertID, zbAccount,
			"--cdVersionNumber", fmt.Sprintf("%d", cdVersionNumber),
			"--programTypeVersion", "1.0",
			"--familyId", "someFID",
			"--supportedClusters", "someClusters",
			"--compliantPlatformUsed", "WIFI",
			"--compliantPlatformVersion", "V1",
			"--OSVersion", "someV",
			"--certificationRoute", "Full",
			"--programType", "pType",
			"--transport", "someTransport",
			"--parentChild", "parent",
			"--certificationIDOfSoftwareComponent", "someIDOfSoftwareComponent",
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Certified model
		out, err := QueryCertifiedModel(vid, pid2, sv2, zigbeeCertType)
		require.NoError(t, err)
		require.Contains(t, string(out), `"value":true`)
		require.Contains(t, string(out), fmt.Sprintf(`"pid":%d`, pid2))
		require.Contains(t, string(out), fmt.Sprintf(`"vid":%d`, vid))

		// Compliance info — all optional fields present
		out, err = QueryComplianceInfo(vid, pid2, sv2, zigbeeCertType)
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"vid":%d`, vid))
		require.Contains(t, string(out), fmt.Sprintf(`"pid":%d`, pid2))
		require.Contains(t, string(out), `"softwareVersionCertificationStatus":2`)
		require.Contains(t, string(out), fmt.Sprintf(`"date":"%s"`, reCertDate))
		require.Contains(t, string(out), fmt.Sprintf(`"certificationType":"%s"`, zigbeeCertType))
		require.Contains(t, string(out), fmt.Sprintf(`"cDCertificateId":"%s"`, cdCertID))
		require.Contains(t, string(out), `"programTypeVersion":"1.0"`)
		require.Contains(t, string(out), `"familyId":"someFID"`)
		require.Contains(t, string(out), `"supportedClusters":"someClusters"`)
		require.Contains(t, string(out), `"compliantPlatformUsed":"WIFI"`)
		require.Contains(t, string(out), `"compliantPlatformVersion":"V1"`)
		require.Contains(t, string(out), `"OSVersion":"someV"`)
		require.Contains(t, string(out), `"certificationRoute":"Full"`)
		require.Contains(t, string(out), `"programType":"pType"`)
		require.Contains(t, string(out), `"transport":"someTransport"`)
		require.Contains(t, string(out), `"parentChild":"parent"`)
		require.Contains(t, string(out), `"certificationIdOfSoftwareComponent":"someIDOfSoftwareComponent"`)

		// Device software compliance — all optional fields present
		out, err = QueryDeviceSoftwareCompliance(cdCertID)
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"vid":%d`, vid))
		require.Contains(t, string(out), fmt.Sprintf(`"pid":%d`, pid2))
		require.Contains(t, string(out), `"softwareVersionCertificationStatus":2`)
		require.Contains(t, string(out), fmt.Sprintf(`"date":"%s"`, reCertDate))
		require.Contains(t, string(out), fmt.Sprintf(`"certificationType":"%s"`, zigbeeCertType))
		require.Contains(t, string(out), fmt.Sprintf(`"cDCertificateId":"%s"`, cdCertID))
		require.Contains(t, string(out), `"programTypeVersion":"1.0"`)
		require.Contains(t, string(out), `"familyId":"someFID"`)
		require.Contains(t, string(out), `"supportedClusters":"someClusters"`)
		require.Contains(t, string(out), `"compliantPlatformUsed":"WIFI"`)
		require.Contains(t, string(out), `"compliantPlatformVersion":"V1"`)
		require.Contains(t, string(out), `"OSVersion":"someV"`)
		require.Contains(t, string(out), `"certificationRoute":"Full"`)
		require.Contains(t, string(out), `"programType":"pType"`)
		require.Contains(t, string(out), `"transport":"someTransport"`)
		require.Contains(t, string(out), `"parentChild":"parent"`)
		require.Contains(t, string(out), `"certificationIdOfSoftwareComponent":"someIDOfSoftwareComponent"`)
	})

	t.Run("UpdateComplianceInfo_ByCertCenter", func(t *testing.T) {
		txResult, err := UpdateComplianceInfo(vid, pid2, sv2, zigbeeCertType, zbAccount,
			"--reason", "new_reason",
			"--programType", "new_program_type",
			"--parentChild", "child",
			"--transport", "new_transport",
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Verify updated fields changed, non-updated optional fields unchanged
		out, err := QueryComplianceInfo(vid, pid2, sv2, zigbeeCertType)
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"vid":%d`, vid))
		require.Contains(t, string(out), fmt.Sprintf(`"pid":%d`, pid2))
		require.Contains(t, string(out), `"softwareVersionCertificationStatus":2`)
		require.Contains(t, string(out), fmt.Sprintf(`"date":"%s"`, reCertDate))
		require.Contains(t, string(out), fmt.Sprintf(`"certificationType":"%s"`, zigbeeCertType))
		require.Contains(t, string(out), fmt.Sprintf(`"cDCertificateId":"%s"`, cdCertID))
		require.Contains(t, string(out), `"programTypeVersion":"1.0"`)
		require.Contains(t, string(out), `"familyId":"someFID"`)
		require.Contains(t, string(out), `"supportedClusters":"someClusters"`)
		require.Contains(t, string(out), `"compliantPlatformUsed":"WIFI"`)
		require.Contains(t, string(out), `"compliantPlatformVersion":"V1"`)
		require.Contains(t, string(out), `"OSVersion":"someV"`)
		require.Contains(t, string(out), `"certificationRoute":"Full"`)
		require.Contains(t, string(out), `"programType":"new_program_type"`)
		require.Contains(t, string(out), `"transport":"new_transport"`)
		require.Contains(t, string(out), `"parentChild":"child"`)
		require.Contains(t, string(out), `"reason":"new_reason"`)
		require.Contains(t, string(out), `"certificationIdOfSoftwareComponent":"someIDOfSoftwareComponent"`)
	})

	t.Run("UpdateComplianceInfo_ByVendor_Fails", func(t *testing.T) {
		txResult, err := UpdateComplianceInfo(vid, pid2, sv2, zigbeeCertType, vendorAccount,
			"--reason", "by_vendor_reason",
			"--programType", "by_vendor_program_type",
			"--parentChild", "parent",
			"--transport", "by_vendor_transport",
		)
		require.NoError(t, err)
		require.NotEqual(t, uint32(0), txResult.Code)
		require.Contains(t, txResult.RawLog, "unauthorized")

		// Fields must be unchanged after vendor update attempt
		out, err := QueryComplianceInfo(vid, pid2, sv2, zigbeeCertType)
		require.NoError(t, err)
		require.Contains(t, string(out), `"programType":"new_program_type"`)
		require.Contains(t, string(out), `"transport":"new_transport"`)
		require.Contains(t, string(out), `"parentChild":"child"`)
		require.Contains(t, string(out), `"reason":"new_reason"`)
	})

	t.Run("UpdateComplianceInfo_NoFields_Unchanged", func(t *testing.T) {
		txResult, err := UpdateComplianceInfo(vid, pid2, sv2, zigbeeCertType, zbAccount)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Fields must remain as set by the first update
		out, err := QueryComplianceInfo(vid, pid2, sv2, zigbeeCertType)
		require.NoError(t, err)
		require.Contains(t, string(out), `"programType":"new_program_type"`)
		require.Contains(t, string(out), `"transport":"new_transport"`)
		require.Contains(t, string(out), `"parentChild":"child"`)
		require.Contains(t, string(out), `"reason":"new_reason"`)
		require.Contains(t, string(out), `"certificationIdOfSoftwareComponent":"someIDOfSoftwareComponent"`)
	})

	updCdCertID := fmt.Sprintf("ucrt-%014d", rand.Intn(1<<30))

	t.Run("UpdateComplianceInfo_AllFields", func(t *testing.T) {
		txResult, err := UpdateComplianceInfo(vid, pid2, sv2, zigbeeCertType, zbAccount,
			"--cdVersionNumber", "1",
			"--certificationDate", "2022-01-01T00:00:01Z",
			"--reason", "brand_new_reason",
			"--cdCertificateId", updCdCertID,
			"--certificationRoute", "brand_new_route",
			"--programType", "brand_new_program_type",
			"--programTypeVersion", "brand_new_program_type_version",
			"--compliantPlatformUsed", "brand_new_compliant_platform_used",
			"--compliantPlatformVersion", "brand_new_compliant_platform_version",
			"--transport", "brand_new_transport",
			"--familyId", "brand_new_family_ID",
			"--supportedClusters", "brand_new_clusters",
			"--OSVersion", "brand_new_os_version",
			"--parentChild", "parent",
			"--certificationIDOfSoftwareComponent", "brand_new_component",
			"--schemaVersion", "0",
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// All fields updated
		out, err := QueryComplianceInfo(vid, pid2, sv2, zigbeeCertType)
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"vid":%d`, vid))
		require.Contains(t, string(out), fmt.Sprintf(`"pid":%d`, pid2))
		require.Contains(t, string(out), fmt.Sprintf(`"softwareVersion":%d`, sv2))
		require.Contains(t, string(out), fmt.Sprintf(`"certificationType":"%s"`, zigbeeCertType))
		require.Contains(t, string(out), `"cDVersionNumber":1`)
		require.Contains(t, string(out), `"date":"2022-01-01T00:00:01Z"`)
		require.Contains(t, string(out), `"reason":"brand_new_reason"`)
		require.Contains(t, string(out), fmt.Sprintf(`"cDCertificateId":"%s"`, updCdCertID))
		require.Contains(t, string(out), `"certificationRoute":"brand_new_route"`)
		require.Contains(t, string(out), `"programType":"brand_new_program_type"`)
		require.Contains(t, string(out), `"programTypeVersion":"brand_new_program_type_version"`)
		require.Contains(t, string(out), `"compliantPlatformUsed":"brand_new_compliant_platform_used"`)
		require.Contains(t, string(out), `"compliantPlatformVersion":"brand_new_compliant_platform_version"`)
		require.Contains(t, string(out), `"transport":"brand_new_transport"`)
		require.Contains(t, string(out), `"familyId":"brand_new_family_ID"`)
		require.Contains(t, string(out), `"supportedClusters":"brand_new_clusters"`)
		require.Contains(t, string(out), `"OSVersion":"brand_new_os_version"`)
		require.Contains(t, string(out), `"parentChild":"parent"`)
		require.Contains(t, string(out), `"certificationIdOfSoftwareComponent":"brand_new_component"`)
		require.Contains(t, string(out), `"schemaVersion":0`)

		// Device software compliance accessible under the updated cdCertificateId
		out, err = QueryDeviceSoftwareCompliance(updCdCertID)
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"vid":%d`, vid))
		require.Contains(t, string(out), fmt.Sprintf(`"pid":%d`, pid2))
	})

	t.Run("DeleteComplianceInfo_AndVerify", func(t *testing.T) {
		txResult, err := DeleteComplianceInfo(vid, pid2, sv2, zigbeeCertType, zbAccount)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Compliance info — Not Found
		out, err := QueryComplianceInfo(vid, pid2, sv2, zigbeeCertType)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")

		// Device software compliance — entry for pid2 no longer present
		out, err = QueryDeviceSoftwareCompliance(cdCertID)
		require.NoError(t, err)
		require.NotContains(t, string(out), fmt.Sprintf(`"pid":%d`, pid2))

		// Certified model — Not Found
		out, err = QueryCertifiedModel(vid, pid2, sv2, zigbeeCertType)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")
	})
}
