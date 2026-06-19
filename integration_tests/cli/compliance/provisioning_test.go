package compliance

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
	cliputils "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/cli/utils"
	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/utils"
)

func TestComplianceProvisioning(t *testing.T) {
	vid := rand.Intn(65534) + 1
	vendorAccount := fmt.Sprintf("vendor_account_%d", vid)
	cliputils.CreateVendorAccount(t, vendorAccount, vid)

	zbAccount := cliputils.CreateAccount(t, "CertificationCenter")
	secondZbAccount := cliputils.CreateAccount(t, "CertificationCenter")

	pid := rand.Intn(65534) + 1
	sv := rand.Intn(65534) + 1
	svs := fmt.Sprintf("%d", rand.Intn(65534)+1)
	certTypeZb := "zigbee"
	certTypeMatter := "matter"
	provisionDate := "2020-02-02T02:20:20Z"
	provisionReason := "some reason"
	cdCertID := fmt.Sprintf("cert-%014d", rand.Intn(1<<30))

	t.Run("ProvisionUnknownModel_Succeeds", func(t *testing.T) {
		txResult, err := ProvisionModel(ProvisionModelOpts{
			VID: vid, PID: pid,
			SoftwareVersion:       sv,
			SoftwareVersionString: svs,
			CertificationType:     certTypeZb,
			ProvisionalDate:       provisionDate,
			CDCertificateID:       cdCertID,
			From:                  zbAccount,
		})
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		out, err := QueryComplianceInfo(ComplianceQueryOpts{VID: vid, PID: pid, SoftwareVersion: sv, CertificationType: certTypeZb})
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"vid":%d`, vid))
		require.Contains(t, string(out), fmt.Sprintf(`"pid":%d`, pid))
		require.Contains(t, string(out), fmt.Sprintf(`"softwareVersion":%d`, sv))
		require.Contains(t, string(out), fmt.Sprintf(`"softwareVersionString":"%s"`, svs))
		require.Contains(t, string(out), fmt.Sprintf(`"certificationType":"%s"`, certTypeZb))
		require.Contains(t, string(out), `"specificationVersion":1`)
		require.Contains(t, string(out), fmt.Sprintf(`"date":"%s"`, provisionDate))
		require.Contains(t, string(out), fmt.Sprintf(`"cDCertificateId":"%s"`, cdCertID))

		// Delete compliance info before creating model
		txResult, err = DeleteComplianceInfo(vid, pid, sv, certTypeZb, zbAccount)
		require.NoError(t, err)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)
	})

	t.Run("CreateModelAndVersion", func(t *testing.T) {
		cliputils.CreateModelAndVersion(t, vid, pid, sv, svs, vendorAccount)
	})

	t.Run("ProvisionZB_Success", func(t *testing.T) {
		txResult, err := ProvisionModel(ProvisionModelOpts{
			VID: vid, PID: pid,
			SoftwareVersion:       sv,
			SoftwareVersionString: svs,
			CertificationType:     certTypeZb,
			ProvisionalDate:       provisionDate,
			CDCertificateID:       cdCertID,
			Reason:                provisionReason,
			From:                  zbAccount,
			Extra:                 []string{"--schemaVersion", "0"},
		})
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)
	})

	t.Run("ProvisionMatter_Success", func(t *testing.T) {
		txResult, err := ProvisionModel(ProvisionModelOpts{
			VID: vid, PID: pid,
			SoftwareVersion:       sv,
			SoftwareVersionString: svs,
			CertificationType:     certTypeMatter,
			ProvisionalDate:       provisionDate,
			CDCertificateID:       cdCertID,
			Reason:                provisionReason,
			From:                  zbAccount,
		})
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)
	})

	t.Run("ReProvision_SameState_Fails", func(t *testing.T) {
		txResult, err := ProvisionModel(ProvisionModelOpts{
			VID: vid, PID: pid,
			SoftwareVersion:       sv,
			SoftwareVersionString: svs,
			CertificationType:     certTypeZb,
			ProvisionalDate:       provisionDate,
			CDCertificateID:       cdCertID,
			Reason:                provisionReason,
			From:                  zbAccount,
		})
		require.NoError(t, err)
		require.Equal(t, uint32(305), txResult.Code)
		require.Contains(t, txResult.RawLog, "already in provisional state")
	})

	t.Run("QueryProvisional", func(t *testing.T) {
		out, err := QueryProvisionalModel(ComplianceQueryOpts{VID: vid, PID: pid, SoftwareVersion: sv, CertificationType: certTypeZb})
		require.NoError(t, err)
		require.Contains(t, string(out), `"value":true`)

		out, err = QueryProvisionalModel(ComplianceQueryOpts{VID: vid, PID: pid, SoftwareVersion: sv, CertificationType: certTypeMatter})
		require.NoError(t, err)
		require.Contains(t, string(out), `"value":true`)

		// Never certified or revoked — records don't exist yet, so "Not Found"
		out, err = QueryCertifiedModel(ComplianceQueryOpts{VID: vid, PID: pid, SoftwareVersion: sv, CertificationType: certTypeZb})
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")

		out, err = QueryRevokedModel(ComplianceQueryOpts{VID: vid, PID: pid, SoftwareVersion: sv, CertificationType: certTypeZb})
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")
	})

	t.Run("CertifyAfterProvisioning_Success", func(t *testing.T) {
		certificationDate := "2020-02-02T02:20:21Z"
		txResult, err := CertifyModel(CertifyModelOpts{
			VID: vid, PID: pid,
			SoftwareVersion:       sv,
			SoftwareVersionString: svs,
			CertificationType:     certTypeZb,
			CertificationDate:     certificationDate,
			CDCertificateID:       cdCertID,
			From:                  zbAccount,
		})
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		out, err := QueryCertifiedModel(ComplianceQueryOpts{VID: vid, PID: pid, SoftwareVersion: sv, CertificationType: certTypeZb})
		require.NoError(t, err)
		require.Contains(t, string(out), `"value":true`)

		out, err = QueryProvisionalModel(ComplianceQueryOpts{VID: vid, PID: pid, SoftwareVersion: sv, CertificationType: certTypeZb})
		require.NoError(t, err)
		require.Contains(t, string(out), `"value":false`)
	})

	_ = secondZbAccount
}
