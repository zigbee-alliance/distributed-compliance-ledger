package compliance

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
	cliputils "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/cli/utils"
	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/utils"
)

// TestComplianceRevocation translates compliance-revocation.sh.
func TestComplianceRevocation(t *testing.T) {
	vid := rand.Intn(65534) + 1
	vendorAccount := fmt.Sprintf("vendor_account_%d", vid)
	cliputils.CreateVendorAccount(t, vendorAccount, vid)

	zbAccount := cliputils.CreateAccount(t, "CertificationCenter")
	secondZbAccount := cliputils.CreateAccount(t, "CertificationCenter")

	certType := "zigbee"
	certTypeMatter := "matter"

	pid := rand.Intn(65534) + 1
	sv := rand.Intn(65534) + 1
	svs := fmt.Sprintf("%d", rand.Intn(65534)+1)
	cdCertID := fmt.Sprintf("cert-%014d", rand.Intn(1<<30))

	t.Run("CreateModelAndVersion", func(t *testing.T) {
		cliputils.CreateModelAndVersion(t, vid, pid, sv, svs, vendorAccount)
	})

	revocationDate := "2020-02-02T02:20:20Z"
	revocationReason := "some reason"

	t.Run("RevokeUncertifiedModel_ZB_Success", func(t *testing.T) {
		txResult, err := RevokeModel(vid, pid, sv, svs, certType, revocationDate, zbAccount,
			"--reason", revocationReason,
			"--cdVersionNumber", "1",
			"--schemaVersion", "0",
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)
	})

	t.Run("ReRevoke_DifferentAccount_Fails", func(t *testing.T) {
		txResult, err := RevokeModel(vid, pid, sv, svs, certType, revocationDate, secondZbAccount,
			"--cdVersionNumber", "1",
		)
		require.NoError(t, err)
		require.Equal(t, uint32(304), txResult.Code)
		require.Contains(t, txResult.RawLog, "already revoked on the ledger")
	})

	t.Run("ReRevoke_SameAccount_Fails", func(t *testing.T) {
		txResult, err := RevokeModel(vid, pid, sv, svs, certType, revocationDate, zbAccount,
			"--cdVersionNumber", "1",
		)
		require.NoError(t, err)
		require.Equal(t, uint32(304), txResult.Code)
		require.Contains(t, txResult.RawLog, "already revoked on the ledger")
	})

	t.Run("RevokeUncertifiedModel_Matter_Success", func(t *testing.T) {
		txResult, err := RevokeModel(vid, pid, sv, svs, certTypeMatter, revocationDate, zbAccount,
			"--reason", revocationReason,
			"--cdVersionNumber", "1",
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)
	})

	t.Run("QueryAfterRevocation", func(t *testing.T) {
		out, err := QueryCertifiedModel(vid, pid, sv, certType)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")

		out, err = QueryCertifiedModel(vid, pid, sv, certTypeMatter)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")

		out, err = QueryRevokedModel(vid, pid, sv, certType)
		require.NoError(t, err)
		require.Contains(t, string(out), `"value":true`)
		require.Contains(t, string(out), fmt.Sprintf(`"pid":%d`, pid))
		require.Contains(t, string(out), fmt.Sprintf(`"vid":%d`, vid))

		out, err = QueryRevokedModel(vid, pid, sv, certTypeMatter)
		require.NoError(t, err)
		require.Contains(t, string(out), `"value":true`)

		// Never provisioned — record doesn't exist, so "Not Found"
		out, err = QueryProvisionalModel(vid, pid, sv, certType)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")

		out, err = QueryComplianceInfo(vid, pid, sv, certType)
		require.NoError(t, err)
		require.Contains(t, string(out), `"softwareVersionCertificationStatus":3`)
		require.Contains(t, string(out), fmt.Sprintf(`"date":"%s"`, revocationDate))
		require.Contains(t, string(out), fmt.Sprintf(`"reason":"%s"`, revocationReason))

		out, err = QueryAllRevokedModels()
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"vid":%d`, vid))
		require.Contains(t, string(out), fmt.Sprintf(`"pid":%d`, pid))

		out, err = QueryAllCertifiedModels()
		require.NoError(t, err)
		require.NotContains(t, string(out), fmt.Sprintf(`"vid":%d`, vid))
	})

	t.Run("CertifyAfterRevoke_Success", func(t *testing.T) {
		certificationDate := "2020-02-02T02:20:21Z"
		txResult, err := CertifyModel(vid, pid, sv, svs, certType, certificationDate, cdCertID, zbAccount,
			"--cdVersionNumber", "1",
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		out, err := QueryCertifiedModel(vid, pid, sv, certType)
		require.NoError(t, err)
		require.Contains(t, string(out), `"value":true`)

		out, err = QueryRevokedModel(vid, pid, sv, certType)
		require.NoError(t, err)
		require.Contains(t, string(out), `"value":false`)
	})
}

// TestComplianceSchemaV1Negative covers the schema-v1 (#730) and Matter-spec (#713)
// constraints appended to compliance-revocation.sh in commit 147904e9.
func TestComplianceSchemaV1Negative(t *testing.T) {
	vid := rand.Intn(65534) + 1
	vendorAccount := fmt.Sprintf("vendor_account_%d", vid)
	cliputils.CreateVendorAccount(t, vendorAccount, vid)

	zbAccount := cliputils.CreateAccount(t, "CertificationCenter")

	pid := rand.Intn(65534) + 1
	sv := rand.Intn(65534) + 1
	svs := fmt.Sprintf("%d", rand.Intn(65534)+1)
	cliputils.CreateModelAndVersion(t, vid, pid, sv, svs, vendorAccount)

	certType := "zigbee"
	certDate := "2020-02-02T02:20:20Z"
	provisionDate := "2020-02-02T02:20:20Z"
	revocationDate := "2020-02-02T02:20:20Z"
	cdCertID := "12345678910abcdefgh" // exactly 19 chars
	specVersion := "1"

	t.Run("CertifyModel_SchemaVersion0_Fails", func(t *testing.T) {
		out := runTxRaw(t,
			"tx", "compliance", "certify-model",
			"--vid", fmt.Sprintf("%d", vid),
			"--pid", fmt.Sprintf("%d", pid),
			"--softwareVersion", fmt.Sprintf("%d", sv),
			"--softwareVersionString", svs,
			"--certificationType", certType,
			"--certificationDate", certDate,
			"--cdCertificateId", cdCertID,
			"--cdVersionNumber", "1",
			"--specificationVersion", specVersion,
			"--schemaVersion", "0",
			"--from", zbAccount,
		)
		require.Contains(t, out, "SchemaVersion must be equal 1")
	})

	t.Run("CertifyModel_ShortCdCertificateId_Fails", func(t *testing.T) {
		shortID := "1234567890abcdefgh" // 18 chars
		out := runTxRaw(t,
			"tx", "compliance", "certify-model",
			"--vid", fmt.Sprintf("%d", vid),
			"--pid", fmt.Sprintf("%d", pid),
			"--softwareVersion", fmt.Sprintf("%d", sv),
			"--softwareVersionString", svs,
			"--certificationType", certType,
			"--certificationDate", certDate,
			"--cdCertificateId", shortID,
			"--cdVersionNumber", "1",
			"--specificationVersion", specVersion,
			"--schemaVersion", "1",
			"--from", zbAccount,
		)
		require.Contains(t, out, "minimum length for CDCertificateId allowed is 19")
	})

	t.Run("CertifyModel_LongCdCertificateId_Fails", func(t *testing.T) {
		longID := "12345678910abcdefghX" // 20 chars
		out := runTxRaw(t,
			"tx", "compliance", "certify-model",
			"--vid", fmt.Sprintf("%d", vid),
			"--pid", fmt.Sprintf("%d", pid),
			"--softwareVersion", fmt.Sprintf("%d", sv),
			"--softwareVersionString", svs,
			"--certificationType", certType,
			"--certificationDate", certDate,
			"--cdCertificateId", longID,
			"--cdVersionNumber", "1",
			"--specificationVersion", specVersion,
			"--schemaVersion", "1",
			"--from", zbAccount,
		)
		require.Contains(t, out, "maximum length for CDCertificateId allowed is 19")
	})

	t.Run("ProvisionModel_SchemaVersion0_Fails", func(t *testing.T) {
		out := runTxRaw(t,
			"tx", "compliance", "provision-model",
			"--vid", fmt.Sprintf("%d", vid),
			"--pid", fmt.Sprintf("%d", pid),
			"--softwareVersion", fmt.Sprintf("%d", sv),
			"--softwareVersionString", svs,
			"--certificationType", certType,
			"--provisionalDate", provisionDate,
			"--cdCertificateId", cdCertID,
			"--cdVersionNumber", "1",
			"--specificationVersion", specVersion,
			"--schemaVersion", "0",
			"--from", zbAccount,
		)
		require.Contains(t, out, "SchemaVersion must be equal 1")
	})

	t.Run("RevokeModel_SchemaVersion0_Fails", func(t *testing.T) {
		out := runTxRaw(t,
			"tx", "compliance", "revoke-model",
			"--vid", fmt.Sprintf("%d", vid),
			"--pid", fmt.Sprintf("%d", pid),
			"--softwareVersion", fmt.Sprintf("%d", sv),
			"--softwareVersionString", svs,
			"--certificationType", certType,
			"--revocationDate", revocationDate,
			"--cdVersionNumber", "1",
			"--schemaVersion", "0",
			"--from", zbAccount,
		)
		require.Contains(t, out, "SchemaVersion must be equal 1")
	})

	t.Run("CertifyModel_LongCertificationType_Fails", func(t *testing.T) {
		longCertType := "this_certification_type_is_way_too_long"
		out := runTxRaw(t,
			"tx", "compliance", "certify-model",
			"--vid", fmt.Sprintf("%d", vid),
			"--pid", fmt.Sprintf("%d", pid),
			"--softwareVersion", fmt.Sprintf("%d", sv),
			"--softwareVersionString", svs,
			"--certificationType", longCertType,
			"--certificationDate", certDate,
			"--cdCertificateId", cdCertID,
			"--cdVersionNumber", "1",
			"--specificationVersion", specVersion,
			"--schemaVersion", "1",
			"--from", zbAccount,
		)
		require.Contains(t, out, "maximum length for CertificationType allowed is 20")
	})

	t.Run("CertifyModel_SpecificationVersionZero_Fails", func(t *testing.T) {
		out := runTxRaw(t,
			"tx", "compliance", "certify-model",
			"--vid", fmt.Sprintf("%d", vid),
			"--pid", fmt.Sprintf("%d", pid),
			"--softwareVersion", fmt.Sprintf("%d", sv),
			"--softwareVersionString", svs,
			"--certificationType", certType,
			"--certificationDate", certDate,
			"--cdCertificateId", cdCertID,
			"--cdVersionNumber", "1",
			"--specificationVersion", "0",
			"--schemaVersion", "1",
			"--from", zbAccount,
		)
		require.Contains(t, out, "SpecificationVersion is a required field")
	})
}

// runTxRaw runs a tx via ExecuteCLI (bypassing the JSON-parsing ExecuteTx) so the
// raw validate-basic error message is returned, mirroring the bash `check_response_and_report ... raw` flow.
func runTxRaw(t *testing.T, args ...string) string {
	t.Helper()
	args = append(args, "--yes", "-o", "json", "--keyring-backend", "test")
	out, err := utils.ExecuteCLI(args...)
	combined := string(out)
	if err != nil {
		combined += err.Error()
	}

	return combined
}
