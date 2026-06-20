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

		info, err := GetComplianceInfo(ComplianceQueryOpts{VID: vid, PID: pid, SoftwareVersion: sv, CertificationType: certTypeZb})
		require.NoError(t, err)
		require.NotNil(t, info)
		require.Equal(t, int32(vid), info.Vid)
		require.Equal(t, int32(pid), info.Pid)
		require.Equal(t, uint32(sv), info.SoftwareVersion)
		require.Equal(t, svs, info.SoftwareVersionString)
		require.Equal(t, certTypeZb, info.CertificationType)
		require.Equal(t, uint32(1), info.SpecificationVersion)
		require.Equal(t, provisionDate, info.Date)
		require.Equal(t, cdCertID, info.CDCertificateId)

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
		provisional, err := GetProvisionalModel(ComplianceQueryOpts{VID: vid, PID: pid, SoftwareVersion: sv, CertificationType: certTypeZb})
		require.NoError(t, err)
		require.NotNil(t, provisional)
		require.True(t, provisional.Value)

		provisional, err = GetProvisionalModel(ComplianceQueryOpts{VID: vid, PID: pid, SoftwareVersion: sv, CertificationType: certTypeMatter})
		require.NoError(t, err)
		require.NotNil(t, provisional)
		require.True(t, provisional.Value)

		// Never certified or revoked — records don't exist yet.
		certified, err := GetCertifiedModel(ComplianceQueryOpts{VID: vid, PID: pid, SoftwareVersion: sv, CertificationType: certTypeZb})
		require.NoError(t, err)
		require.Nil(t, certified)

		revoked, err := GetRevokedModel(ComplianceQueryOpts{VID: vid, PID: pid, SoftwareVersion: sv, CertificationType: certTypeZb})
		require.NoError(t, err)
		require.Nil(t, revoked)
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

		certified, err := GetCertifiedModel(ComplianceQueryOpts{VID: vid, PID: pid, SoftwareVersion: sv, CertificationType: certTypeZb})
		require.NoError(t, err)
		require.NotNil(t, certified)
		require.True(t, certified.Value)

		provisional, err := GetProvisionalModel(ComplianceQueryOpts{VID: vid, PID: pid, SoftwareVersion: sv, CertificationType: certTypeZb})
		require.NoError(t, err)
		require.NotNil(t, provisional)
		require.False(t, provisional.Value)
	})

	_ = secondZbAccount
}
