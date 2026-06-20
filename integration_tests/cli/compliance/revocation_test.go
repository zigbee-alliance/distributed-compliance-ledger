package compliance

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
	cliputils "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/cli/utils"
	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/utils"
)

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
		txResult, err := RevokeModel(RevokeModelOpts{
			VID: vid, PID: pid,
			SoftwareVersion:       sv,
			SoftwareVersionString: svs,
			CertificationType:     certType,
			RevocationDate:        revocationDate,
			Reason:                revocationReason,
			From:                  zbAccount,
		})
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)
	})

	t.Run("ReRevoke_DifferentAccount_Fails", func(t *testing.T) {
		txResult, err := RevokeModel(RevokeModelOpts{
			VID: vid, PID: pid,
			SoftwareVersion:       sv,
			SoftwareVersionString: svs,
			CertificationType:     certType,
			RevocationDate:        revocationDate,
			From:                  secondZbAccount,
		})
		require.NoError(t, err)
		require.Equal(t, uint32(304), txResult.Code)
		require.Contains(t, txResult.RawLog, "already revoked on the ledger")
	})

	t.Run("ReRevoke_SameAccount_Fails", func(t *testing.T) {
		txResult, err := RevokeModel(RevokeModelOpts{
			VID: vid, PID: pid,
			SoftwareVersion:       sv,
			SoftwareVersionString: svs,
			CertificationType:     certType,
			RevocationDate:        revocationDate,
			From:                  zbAccount,
		})
		require.NoError(t, err)
		require.Equal(t, uint32(304), txResult.Code)
		require.Contains(t, txResult.RawLog, "already revoked on the ledger")
	})

	t.Run("RevokeUncertifiedModel_Matter_Success", func(t *testing.T) {
		txResult, err := RevokeModel(RevokeModelOpts{
			VID: vid, PID: pid,
			SoftwareVersion:       sv,
			SoftwareVersionString: svs,
			CertificationType:     certTypeMatter,
			RevocationDate:        revocationDate,
			Reason:                revocationReason,
			From:                  zbAccount,
		})
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)
	})

	t.Run("QueryAfterRevocation", func(t *testing.T) {
		certified, err := GetCertifiedModel(ComplianceQueryOpts{VID: vid, PID: pid, SoftwareVersion: sv, CertificationType: certType})
		require.NoError(t, err)
		require.Nil(t, certified)

		certified, err = GetCertifiedModel(ComplianceQueryOpts{VID: vid, PID: pid, SoftwareVersion: sv, CertificationType: certTypeMatter})
		require.NoError(t, err)
		require.Nil(t, certified)

		revoked, err := GetRevokedModel(ComplianceQueryOpts{VID: vid, PID: pid, SoftwareVersion: sv, CertificationType: certType})
		require.NoError(t, err)
		require.NotNil(t, revoked)
		require.True(t, revoked.Value)
		require.Equal(t, int32(vid), revoked.Vid)
		require.Equal(t, int32(pid), revoked.Pid)

		revoked, err = GetRevokedModel(ComplianceQueryOpts{VID: vid, PID: pid, SoftwareVersion: sv, CertificationType: certTypeMatter})
		require.NoError(t, err)
		require.NotNil(t, revoked)
		require.True(t, revoked.Value)

		// Never provisioned — record doesn't exist.
		provisional, err := GetProvisionalModel(ComplianceQueryOpts{VID: vid, PID: pid, SoftwareVersion: sv, CertificationType: certType})
		require.NoError(t, err)
		require.Nil(t, provisional)

		info, err := GetComplianceInfo(ComplianceQueryOpts{VID: vid, PID: pid, SoftwareVersion: sv, CertificationType: certType})
		require.NoError(t, err)
		require.NotNil(t, info)
		require.Equal(t, uint32(3), info.SoftwareVersionCertificationStatus)
		require.Equal(t, revocationDate, info.Date)
		require.Equal(t, revocationReason, info.Reason)

		allRevoked, err := GetAllRevokedModels()
		require.NoError(t, err)
		require.True(t, containsRevokedModel(allRevoked, int32(vid), int32(pid)))

		allCertified, err := GetAllCertifiedModels()
		require.NoError(t, err)
		require.False(t, containsCertifiedModel(allCertified, int32(vid), int32(pid)))
	})

	t.Run("CertifyAfterRevoke_Success", func(t *testing.T) {
		certificationDate := "2020-02-02T02:20:21Z"
		txResult, err := CertifyModel(CertifyModelOpts{
			VID: vid, PID: pid,
			SoftwareVersion:       sv,
			SoftwareVersionString: svs,
			CertificationType:     certType,
			CertificationDate:     certificationDate,
			CDCertificateID:       cdCertID,
			From:                  zbAccount,
		})
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		certified, err := GetCertifiedModel(ComplianceQueryOpts{VID: vid, PID: pid, SoftwareVersion: sv, CertificationType: certType})
		require.NoError(t, err)
		require.NotNil(t, certified)
		require.True(t, certified.Value)

		revoked, err := GetRevokedModel(ComplianceQueryOpts{VID: vid, PID: pid, SoftwareVersion: sv, CertificationType: certType})
		require.NoError(t, err)
		require.NotNil(t, revoked)
		require.False(t, revoked.Value)
	})
}
