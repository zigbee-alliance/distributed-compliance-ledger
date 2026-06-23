package compliance

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/cli/model"
	cliputils "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/cli/utils"
)

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
		txResult, err := model.AddModel(model.AddModelOpts{
			VIDHex: vidHex, PIDHex: pidHex, From: vendorAccount,
		})
		cliputils.RequireTxOK(t, txResult, err)

		txResult, err = model.AddModelVersion(model.AddModelVersionOpts{
			VIDHex: vidHex, PIDHex: pidHex,
			SoftwareVersion:       sv,
			SoftwareVersionString: svs,
			From:                  vendorAccount,
		})
		cliputils.RequireTxOK(t, txResult, err)
	})

	t.Run("QueryBeforeCertification_NotFound", func(t *testing.T) {
		zbHex := ComplianceQueryOpts{
			VIDHex: vidHex, PIDHex: pidHex,
			SoftwareVersion:   sv,
			CertificationType: "zigbee",
		}

		info, err := GetComplianceInfo(zbHex)
		require.NoError(t, err)
		require.Nil(t, info)

		dsc, err := GetDeviceSoftwareCompliance(cdCertID)
		require.NoError(t, err)
		require.Nil(t, dsc)

		certified, err := GetCertifiedModel(zbHex)
		require.NoError(t, err)
		require.Nil(t, certified)

		revoked, err := GetRevokedModel(zbHex)
		require.NoError(t, err)
		require.Nil(t, revoked)

		provisional, err := GetProvisionalModel(zbHex)
		require.NoError(t, err)
		require.Nil(t, provisional)
	})

	t.Run("CertifyWithInvalidSVS_Fails", func(t *testing.T) {
		invalidSvs := fmt.Sprintf("%d", rand.Intn(65534)+1)
		txResult, err := CertifyModel(CertifyModelOpts{
			VIDHex: vidHex, PIDHex: pidHex,
			SoftwareVersion:       sv,
			SoftwareVersionString: invalidSvs,
			CertificationType:     zigbeeCertType,
			CertificationDate:     certificationDate,
			CDCertificateID:       cdCertID,
			From:                  zbAccount,
		})
		cliputils.RequireTxFailCode(t, txResult, err, 306)
	})

	t.Run("CertifyZigbee_WithHexVID_Success", func(t *testing.T) {
		txResult, err := CertifyModel(CertifyModelOpts{
			VIDHex: vidHex, PIDHex: pidHex,
			SoftwareVersion:       sv,
			SoftwareVersionString: svs,
			CertificationType:     zigbeeCertType,
			CertificationDate:     certificationDate,
			CDCertificateID:       cdCertID,
			From:                  zbAccount,
		})
		cliputils.RequireTxOK(t, txResult, err)
	})

	t.Run("CertifyMatter_WithHexVID_Success", func(t *testing.T) {
		txResult, err := CertifyModel(CertifyModelOpts{
			VIDHex: vidHex, PIDHex: pidHex,
			SoftwareVersion:       sv,
			SoftwareVersionString: svs,
			CertificationType:     matterCertType,
			CertificationDate:     certificationDate,
			CDCertificateID:       cdCertID,
			From:                  zbAccount,
		})
		cliputils.RequireTxOK(t, txResult, err)
	})

	t.Run("ReCertify_DifferentAccount_Fails", func(t *testing.T) {
		txResult, err := CertifyModel(CertifyModelOpts{
			VIDHex: vidHex, PIDHex: pidHex,
			SoftwareVersion:       sv,
			SoftwareVersionString: svs,
			CertificationType:     zigbeeCertType,
			CertificationDate:     certificationDate,
			CDCertificateID:       cdCertID,
			From:                  secondZbAccount,
		})
		require.NoError(t, err)
		require.Equal(t, uint32(303), txResult.Code)
		require.Contains(t, txResult.RawLog, "already certified on the ledger")
	})

	t.Run("ReCertify_SameAccount_Fails", func(t *testing.T) {
		txResult, err := CertifyModel(CertifyModelOpts{
			VIDHex: vidHex, PIDHex: pidHex,
			SoftwareVersion:       sv,
			SoftwareVersionString: svs,
			CertificationType:     zigbeeCertType,
			CertificationDate:     certificationDate,
			CDCertificateID:       cdCertID,
			From:                  zbAccount,
		})
		require.NoError(t, err)
		require.Equal(t, uint32(303), txResult.Code)
		require.Contains(t, txResult.RawLog, "already certified on the ledger")
	})

	t.Run("QueryAfterCertification", func(t *testing.T) {
		certified, err := GetCertifiedModel(ComplianceQueryOpts{
			VIDHex: vidHex, PIDHex: pidHex,
			SoftwareVersion:   sv,
			CertificationType: matterCertType,
		})
		require.NoError(t, err)
		require.NotNil(t, certified)
		require.True(t, certified.Value)
		require.Equal(t, int32(vid), certified.Vid)
		require.Equal(t, int32(pid), certified.Pid)

		info, err := GetComplianceInfo(ComplianceQueryOpts{
			VIDHex: vidHex, PIDHex: pidHex,
			SoftwareVersion:   sv,
			CertificationType: zigbeeCertType,
		})
		require.NoError(t, err)
		require.NotNil(t, info)
		require.Equal(t, uint32(2), info.SoftwareVersionCertificationStatus)
		require.Equal(t, certificationDate, info.Date)

		dsc, err := GetDeviceSoftwareCompliance(cdCertID)
		require.NoError(t, err)
		require.NotNil(t, dsc)
		require.Len(t, dsc.ComplianceInfo, 2)
		require.True(t, containsComplianceInfoCertType(dsc.ComplianceInfo, int32(vid), int32(pid), matterCertType))
		require.True(t, containsComplianceInfoCertType(dsc.ComplianceInfo, int32(vid), int32(pid), zigbeeCertType))
		for _, ci := range dsc.ComplianceInfo {
			require.Equal(t, uint32(2), ci.SoftwareVersionCertificationStatus)
		}

		// Companion records for BOTH cert types: certified=true, revoked=false,
		// provisional=false; compliance-info status 2.
		for _, ct := range []string{zigbeeCertType, matterCertType} {
			q := ComplianceQueryOpts{VIDHex: vidHex, PIDHex: pidHex, SoftwareVersion: sv, CertificationType: ct}

			cert, err := GetCertifiedModel(q)
			require.NoError(t, err)
			require.NotNil(t, cert, "certified-model missing for %s", ct)
			require.True(t, cert.Value)

			rev, err := GetRevokedModel(q)
			require.NoError(t, err)
			require.NotNil(t, rev, "revoked-model missing for %s", ct)
			require.False(t, rev.Value)

			prov, err := GetProvisionalModel(q)
			require.NoError(t, err)
			require.NotNil(t, prov, "provisional-model missing for %s", ct)
			require.False(t, prov.Value)

			ci, err := GetComplianceInfo(q)
			require.NoError(t, err)
			require.NotNil(t, ci)
			require.Equal(t, uint32(2), ci.SoftwareVersionCertificationStatus)
		}
	})

	revocationDate := "2020-02-02T02:20:20Z"
	revocationReason := "some reason"

	t.Run("RevokeFromPast_Fails", func(t *testing.T) {
		pastDate := "2020-01-01T00:00:00Z"
		txResult, err := RevokeModel(RevokeModelOpts{
			VIDHex: vidHex, PIDHex: pidHex,
			SoftwareVersion:       sv,
			SoftwareVersionString: svs,
			CertificationType:     zigbeeCertType,
			RevocationDate:        pastDate,
			Reason:                revocationReason,
			From:                  zbAccount,
		})
		require.NoError(t, err)
		require.Equal(t, uint32(302), txResult.Code)
		require.Contains(t, txResult.RawLog, "must be after")
	})

	t.Run("RevokeZigbee_Success", func(t *testing.T) {
		txResult, err := RevokeModel(RevokeModelOpts{
			VIDHex: vidHex, PIDHex: pidHex,
			SoftwareVersion:       sv,
			SoftwareVersionString: svs,
			CertificationType:     zigbeeCertType,
			RevocationDate:        revocationDate,
			Reason:                revocationReason,
			From:                  zbAccount,
		})
		cliputils.RequireTxOK(t, txResult, err)
	})

	t.Run("QueryAfterRevocation", func(t *testing.T) {
		zbHex := ComplianceQueryOpts{
			VIDHex: vidHex, PIDHex: pidHex,
			SoftwareVersion:   sv,
			CertificationType: zigbeeCertType,
		}

		info, err := GetComplianceInfo(zbHex)
		require.NoError(t, err)
		require.NotNil(t, info)
		require.Equal(t, uint32(3), info.SoftwareVersionCertificationStatus)
		require.Equal(t, revocationDate, info.Date)
		require.Equal(t, revocationReason, info.Reason)
		require.NotNil(t, info.History)

		dsc, err := GetDeviceSoftwareCompliance(cdCertID)
		require.NoError(t, err)
		require.NotNil(t, dsc)
		require.False(t, containsComplianceInfoCertType(dsc.ComplianceInfo, int32(vid), int32(pid), zigbeeCertType))
		require.True(t, containsComplianceInfoCertType(dsc.ComplianceInfo, int32(vid), int32(pid), matterCertType))
		for _, ci := range dsc.ComplianceInfo {
			require.Equal(t, uint32(2), ci.SoftwareVersionCertificationStatus)
		}

		revoked, err := GetRevokedModel(zbHex)
		require.NoError(t, err)
		require.NotNil(t, revoked)
		require.True(t, revoked.Value)

		// Zigbee companions flip: certified=false, provisional=false.
		certified, err := GetCertifiedModel(zbHex)
		require.NoError(t, err)
		require.NotNil(t, certified)
		require.False(t, certified.Value)

		provisional, err := GetProvisionalModel(zbHex)
		require.NoError(t, err)
		require.NotNil(t, provisional)
		require.False(t, provisional.Value)

		// Matter side is untouched — still certified (status 2).
		matterInfo, err := GetComplianceInfo(ComplianceQueryOpts{
			VIDHex: vidHex, PIDHex: pidHex, SoftwareVersion: sv, CertificationType: matterCertType,
		})
		require.NoError(t, err)
		require.NotNil(t, matterInfo)
		require.Equal(t, uint32(2), matterInfo.SoftwareVersionCertificationStatus)
	})

	t.Run("ReCertifyAfterRevoke_Success", func(t *testing.T) {
		newCertDate := "2020-03-03T00:00:00Z"
		txResult, err := CertifyModel(CertifyModelOpts{
			VIDHex: vidHex, PIDHex: pidHex,
			SoftwareVersion:       sv,
			SoftwareVersionString: svs,
			CertificationType:     zigbeeCertType,
			CertificationDate:     newCertDate,
			CDCertificateID:       cdCertID,
			From:                  zbAccount,
		})
		cliputils.RequireTxOK(t, txResult, err)

		zbHex := ComplianceQueryOpts{
			VIDHex: vidHex, PIDHex: pidHex,
			SoftwareVersion:   sv,
			CertificationType: zigbeeCertType,
		}

		info, err := GetComplianceInfo(zbHex)
		require.NoError(t, err)
		require.NotNil(t, info)
		require.Equal(t, uint32(2), info.SoftwareVersionCertificationStatus)
		require.Equal(t, newCertDate, info.Date)

		certified, err := GetCertifiedModel(zbHex)
		require.NoError(t, err)
		require.NotNil(t, certified)
		require.True(t, certified.Value)

		// Revoked flips back to false on re-certification.
		revoked, err := GetRevokedModel(zbHex)
		require.NoError(t, err)
		require.NotNil(t, revoked)
		require.False(t, revoked.Value)

		// Matter is still certified (status 2).
		matterInfo, err := GetComplianceInfo(ComplianceQueryOpts{
			VIDHex: vidHex, PIDHex: pidHex, SoftwareVersion: sv, CertificationType: matterCertType,
		})
		require.NoError(t, err)
		require.NotNil(t, matterInfo)
		require.Equal(t, uint32(2), matterInfo.SoftwareVersionCertificationStatus)
	})
}
