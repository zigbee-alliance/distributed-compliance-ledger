package compliance

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/cli/model"
	cliputils "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/cli/utils"
	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/utils"
	compliancetypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/compliance/types"
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
		txResult, err := CertifyModel(CertifyModelOpts{
			VID: vid, PID: pid,
			SoftwareVersion:       sv,
			SoftwareVersionString: svs,
			CertificationType:     zigbeeCertType,
			CertificationDate:     certificationDate,
			CDCertificateID:       cdCertID,
			CDVersionNumber:       cdVersionNumber,
			From:                  zbAccount,
		})
		cliputils.RequireTxOK(t, txResult, err)

		dsc, err := GetDeviceSoftwareCompliance(cdCertID)
		require.NoError(t, err)
		require.NotNil(t, dsc)
		require.Equal(t, cdCertID, dsc.CDCertificateId)
		require.True(t, containsComplianceInfoCertType(dsc.ComplianceInfo, int32(vid), int32(pid), zigbeeCertType))
		for _, ci := range dsc.ComplianceInfo {
			if ci.CertificationType == zigbeeCertType {
				require.Equal(t, uint32(sv), ci.SoftwareVersion)
				require.Equal(t, svs, ci.SoftwareVersionString)
				require.Equal(t, certificationDate, ci.Date)
				require.Equal(t, cdCertID, ci.CDCertificateId)
			}
		}

		txResult, err = DeleteComplianceInfo(vid, pid, sv, zigbeeCertType, zbAccount)
		require.NoError(t, err)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)
	})

	t.Run("CreateModelAndVersion", func(t *testing.T) {
		txResult, err := model.AddModel(model.AddModelOpts{
			VID: vid, PID: pid, From: vendorAccount,
		})
		cliputils.RequireTxOK(t, txResult, err)

		txResult, err = model.AddModelVersion(model.AddModelVersionOpts{
			VID: vid, PID: pid,
			SoftwareVersion:       sv,
			SoftwareVersionString: svs,
			CDVersionNumber:       cdVersionNumber,
			From:                  vendorAccount,
		})
		cliputils.RequireTxOK(t, txResult, err)
	})

	t.Run("QueryBeforeCertification_NotFound", func(t *testing.T) {
		zbQuery := ComplianceQueryOpts{VID: vid, PID: pid, SoftwareVersion: sv, CertificationType: zigbeeCertType}

		info, err := GetComplianceInfo(zbQuery)
		require.NoError(t, err)
		require.Nil(t, info)

		dsc, err := GetDeviceSoftwareCompliance(cdCertID)
		require.NoError(t, err)
		require.Nil(t, dsc)

		certified, err := GetCertifiedModel(zbQuery)
		require.NoError(t, err)
		require.Nil(t, certified)

		revoked, err := GetRevokedModel(zbQuery)
		require.NoError(t, err)
		require.Nil(t, revoked)

		provisional, err := GetProvisionalModel(zbQuery)
		require.NoError(t, err)
		require.Nil(t, provisional)

		// QueryAll — this vid/pid must not appear yet
		allInfo, err := GetAllComplianceInfo()
		require.NoError(t, err)
		require.False(t, containsComplianceInfo(allInfo, int32(vid), int32(pid)))

		allDsc, err := GetAllDeviceSoftwareCompliance()
		require.NoError(t, err)
		require.False(t, containsDeviceSoftwareCompliance(allDsc, cdCertID))

		allCertified, err := GetAllCertifiedModels()
		require.NoError(t, err)
		require.False(t, containsCertifiedModel(allCertified, int32(vid), int32(pid)))

		allRevoked, err := GetAllRevokedModels()
		require.NoError(t, err)
		require.False(t, containsRevokedModel(allRevoked, int32(vid), int32(pid)))

		allProvisional, err := GetAllProvisionalModels()
		require.NoError(t, err)
		require.False(t, containsProvisionalModel(allProvisional, int32(vid), int32(pid)))
	})

	t.Run("CertifyWithInvalidSVS_Fails", func(t *testing.T) {
		invalidSvs := fmt.Sprintf("%d", rand.Intn(65534)+1)
		txResult, err := CertifyModel(CertifyModelOpts{
			VID: vid, PID: pid,
			SoftwareVersion:       sv,
			SoftwareVersionString: invalidSvs,
			CertificationType:     zigbeeCertType,
			CertificationDate:     certificationDate,
			CDCertificateID:       cdCertID,
			CDVersionNumber:       cdVersionNumber,
			From:                  zbAccount,
		})
		cliputils.RequireTxFailCode(t, txResult, err, 306)
	})

	t.Run("CertifyWithInvalidCDVersionNumber_Fails", func(t *testing.T) {
		// The model version was created with CDVersionNumber=1; certifying with a
		// different value must be rejected as a mismatch.
		txResult, err := CertifyModel(CertifyModelOpts{
			VID: vid, PID: pid,
			SoftwareVersion:       sv,
			SoftwareVersionString: svs,
			CertificationType:     zigbeeCertType,
			CertificationDate:     certificationDate,
			CDCertificateID:       cdCertID,
			CDVersionNumber:       2,
			From:                  zbAccount,
		})
		cliputils.RequireTxFailCode(t, txResult, err, 306)
		require.Contains(t, txResult.RawLog, "ledger does not have matching CDVersionNumber=2")
	})

	t.Run("CertifyZigbee_Success", func(t *testing.T) {
		txResult, err := CertifyModel(CertifyModelOpts{
			VID: vid, PID: pid,
			SoftwareVersion:       sv,
			SoftwareVersionString: svs,
			CertificationType:     zigbeeCertType,
			CertificationDate:     certificationDate,
			CDCertificateID:       cdCertID,
			CDVersionNumber:       cdVersionNumber,
			From:                  zbAccount,
		})
		cliputils.RequireTxOK(t, txResult, err)
	})

	t.Run("CertifyMatter_Success", func(t *testing.T) {
		txResult, err := CertifyModel(CertifyModelOpts{
			VID: vid, PID: pid,
			SoftwareVersion:       sv,
			SoftwareVersionString: svs,
			CertificationType:     matterCertType,
			CertificationDate:     certificationDate,
			CDCertificateID:       cdCertID,
			CDVersionNumber:       cdVersionNumber,
			From:                  zbAccount,
		})
		cliputils.RequireTxOK(t, txResult, err)
	})

	t.Run("ReCertify_DifferentAccount_Fails", func(t *testing.T) {
		txResult, err := CertifyModel(CertifyModelOpts{
			VID: vid, PID: pid,
			SoftwareVersion:       sv,
			SoftwareVersionString: svs,
			CertificationType:     zigbeeCertType,
			CertificationDate:     certificationDate,
			CDCertificateID:       cdCertID,
			CDVersionNumber:       cdVersionNumber,
			From:                  secondZbAccount,
		})
		cliputils.RequireTxFailCode(t, txResult, err, 303)
		require.Contains(t, txResult.RawLog, "already certified on the ledger")
	})

	t.Run("ReCertify_SameAccount_Fails", func(t *testing.T) {
		txResult, err := CertifyModel(CertifyModelOpts{
			VID: vid, PID: pid,
			SoftwareVersion:       sv,
			SoftwareVersionString: svs,
			CertificationType:     zigbeeCertType,
			CertificationDate:     certificationDate,
			CDCertificateID:       cdCertID,
			CDVersionNumber:       cdVersionNumber,
			From:                  zbAccount,
		})
		cliputils.RequireTxFailCode(t, txResult, err, 303)
		require.Contains(t, txResult.RawLog, "already certified on the ledger")
	})

	t.Run("QueryAfterCertification", func(t *testing.T) {
		// Both cert types are now certified, with revoked/provisional companion
		// records carrying value=false. Check each tuple.
		for _, ct := range []string{zigbeeCertType, matterCertType} {
			q := ComplianceQueryOpts{VID: vid, PID: pid, SoftwareVersion: sv, CertificationType: ct}

			certified, err := GetCertifiedModel(q)
			require.NoError(t, err)
			require.NotNil(t, certified)
			require.True(t, certified.Value)
			require.Equal(t, int32(vid), certified.Vid)
			require.Equal(t, int32(pid), certified.Pid)

			revoked, err := GetRevokedModel(q)
			require.NoError(t, err)
			require.NotNil(t, revoked)
			require.False(t, revoked.Value)

			provisional, err := GetProvisionalModel(q)
			require.NoError(t, err)
			require.NotNil(t, provisional)
			require.False(t, provisional.Value)

			info, err := GetComplianceInfo(q)
			require.NoError(t, err)
			require.NotNil(t, info)
			require.Equal(t, int32(vid), info.Vid)
			require.Equal(t, int32(pid), info.Pid)
			require.Equal(t, uint32(2), info.SoftwareVersionCertificationStatus)
			require.Equal(t, cdCertID, info.CDCertificateId)
			require.Equal(t, certificationDate, info.Date)
			require.Equal(t, ct, info.CertificationType)
			require.Equal(t, uint32(1), info.SchemaVersion)
		}

		// Device software compliance — both cert types present, status=2.
		dsc, err := GetDeviceSoftwareCompliance(cdCertID)
		require.NoError(t, err)
		require.NotNil(t, dsc)
		require.Equal(t, cdCertID, dsc.CDCertificateId)
		require.True(t, containsComplianceInfoCertType(dsc.ComplianceInfo, int32(vid), int32(pid), zigbeeCertType))
		require.True(t, containsComplianceInfoCertType(dsc.ComplianceInfo, int32(vid), int32(pid), matterCertType))
		for _, ci := range dsc.ComplianceInfo {
			require.Equal(t, uint32(2), ci.SoftwareVersionCertificationStatus)
			require.Equal(t, certificationDate, ci.Date)
			require.Equal(t, cdCertID, ci.CDCertificateId)
		}

		// All certified — vid/pid appears under both cert types.
		allCertified, err := GetAllCertifiedModels()
		require.NoError(t, err)
		require.True(t, hasCertifiedModelCertType(allCertified, int32(vid), int32(pid), zigbeeCertType))
		require.True(t, hasCertifiedModelCertType(allCertified, int32(vid), int32(pid), matterCertType))

		// All revoked/provisional must NOT contain our vid/pid with value=true.
		allRevoked, err := GetAllRevokedModels()
		require.NoError(t, err)
		require.False(t, hasRevokedWithValue(allRevoked, int32(vid), int32(pid), true))

		allProvisional, err := GetAllProvisionalModels()
		require.NoError(t, err)
		require.False(t, hasProvisionalWithValue(allProvisional, int32(vid), int32(pid), true))

		// All compliance info — both cert types present.
		allInfo, err := GetAllComplianceInfo()
		require.NoError(t, err)
		require.True(t, hasComplianceInfoCertType(allInfo, int32(vid), int32(pid), zigbeeCertType, certificationDate))
		require.True(t, hasComplianceInfoCertType(allInfo, int32(vid), int32(pid), matterCertType, certificationDate))

		// All device software compliance — entry present.
		allDsc, err := GetAllDeviceSoftwareCompliance()
		require.NoError(t, err)
		require.True(t, containsDeviceSoftwareCompliance(allDsc, cdCertID))
	})

	t.Run("RevokeFromPast_Fails", func(t *testing.T) {
		pastDate := "2020-01-01T00:00:00Z"
		txResult, err := RevokeModel(RevokeModelOpts{
			VID: vid, PID: pid,
			SoftwareVersion:       sv,
			SoftwareVersionString: svs,
			CertificationType:     zigbeeCertType,
			RevocationDate:        pastDate,
			Reason:                "some reason",
			CDVersionNumber:       cdVersionNumber,
			From:                  zbAccount,
		})
		cliputils.RequireTxFailCode(t, txResult, err, 302)
		require.Contains(t, txResult.RawLog, "must be after")
	})

	revocationDate := "2020-02-02T02:20:20Z"
	revocationReason := "some reason"

	t.Run("RevokeZigbee_Success", func(t *testing.T) {
		txResult, err := RevokeModel(RevokeModelOpts{
			VID: vid, PID: pid,
			SoftwareVersion:       sv,
			SoftwareVersionString: svs,
			CertificationType:     zigbeeCertType,
			RevocationDate:        revocationDate,
			Reason:                revocationReason,
			CDVersionNumber:       cdVersionNumber,
			From:                  zbAccount,
		})
		cliputils.RequireTxOK(t, txResult, err)
	})

	t.Run("QueryAfterRevocation", func(t *testing.T) {
		zbQuery := ComplianceQueryOpts{VID: vid, PID: pid, SoftwareVersion: sv, CertificationType: zigbeeCertType}

		info, err := GetComplianceInfo(zbQuery)
		require.NoError(t, err)
		require.NotNil(t, info)
		require.Equal(t, int32(vid), info.Vid)
		require.Equal(t, int32(pid), info.Pid)
		require.Equal(t, uint32(3), info.SoftwareVersionCertificationStatus)
		require.Equal(t, cdCertID, info.CDCertificateId)
		require.Equal(t, revocationDate, info.Date)
		require.Equal(t, revocationReason, info.Reason)
		require.Equal(t, zigbeeCertType, info.CertificationType)
		require.NotNil(t, info.History)

		// Device software compliance — matter still active, zigbee gone
		dsc, err := GetDeviceSoftwareCompliance(cdCertID)
		require.NoError(t, err)
		require.NotNil(t, dsc)
		require.True(t, containsComplianceInfoCertType(dsc.ComplianceInfo, int32(vid), int32(pid), matterCertType))
		require.False(t, containsComplianceInfoCertType(dsc.ComplianceInfo, int32(vid), int32(pid), zigbeeCertType))
		for _, ci := range dsc.ComplianceInfo {
			require.Equal(t, uint32(2), ci.SoftwareVersionCertificationStatus)
			require.Equal(t, cdCertID, ci.CDCertificateId)
			require.Equal(t, certificationDate, ci.Date)
		}

		revoked, err := GetRevokedModel(zbQuery)
		require.NoError(t, err)
		require.NotNil(t, revoked)
		require.True(t, revoked.Value)

		certified, err := GetCertifiedModel(zbQuery)
		require.NoError(t, err)
		require.NotNil(t, certified)
		require.False(t, certified.Value)

		provisional, err := GetProvisionalModel(zbQuery)
		require.NoError(t, err)
		require.NotNil(t, provisional)
		require.False(t, provisional.Value)

		// All revoked — zigbee present with value=true
		allRevoked, err := GetAllRevokedModels()
		require.NoError(t, err)
		require.True(t, hasRevokedWithValue(allRevoked, int32(vid), int32(pid), true))

		// All certified — matter still present (value=true), but the zigbee
		// (vid, pid, sv) entry was overwritten to value=false on revoke.
		allCertified, err := GetAllCertifiedModels()
		require.NoError(t, err)
		require.True(t, hasCertifiedModelCertType(allCertified, int32(vid), int32(pid), matterCertType))
		require.False(t, hasCertifiedTrueAt(allCertified, int32(vid), int32(pid), uint32(sv), zigbeeCertType))
	})

	reCertDate := "2020-03-03T00:00:00Z"

	t.Run("ReCertifyAfterRevocation_Success", func(t *testing.T) {
		txResult, err := CertifyModel(CertifyModelOpts{
			VID: vid, PID: pid,
			SoftwareVersion:       sv,
			SoftwareVersionString: svs,
			CertificationType:     zigbeeCertType,
			CertificationDate:     reCertDate,
			CDCertificateID:       cdCertID,
			CDVersionNumber:       cdVersionNumber,
			From:                  zbAccount,
		})
		cliputils.RequireTxOK(t, txResult, err)

		// Compliance info — back to certified with new date
		zbQuery := ComplianceQueryOpts{VID: vid, PID: pid, SoftwareVersion: sv, CertificationType: zigbeeCertType}

		info, err := GetComplianceInfo(zbQuery)
		require.NoError(t, err)
		require.NotNil(t, info)
		require.Equal(t, int32(vid), info.Vid)
		require.Equal(t, int32(pid), info.Pid)
		require.Equal(t, uint32(2), info.SoftwareVersionCertificationStatus)
		require.Equal(t, cdCertID, info.CDCertificateId)
		require.Equal(t, reCertDate, info.Date)
		require.Equal(t, zigbeeCertType, info.CertificationType)

		// Device software compliance — both cert types again
		dsc, err := GetDeviceSoftwareCompliance(cdCertID)
		require.NoError(t, err)
		require.NotNil(t, dsc)
		require.Equal(t, cdCertID, dsc.CDCertificateId)
		require.True(t, containsComplianceInfoCertType(dsc.ComplianceInfo, int32(vid), int32(pid), zigbeeCertType))
		require.True(t, containsComplianceInfoCertType(dsc.ComplianceInfo, int32(vid), int32(pid), matterCertType))
		for _, ci := range dsc.ComplianceInfo {
			require.Equal(t, uint32(2), ci.SoftwareVersionCertificationStatus)
		}

		certified, err := GetCertifiedModel(zbQuery)
		require.NoError(t, err)
		require.NotNil(t, certified)
		require.True(t, certified.Value)

		revoked, err := GetRevokedModel(zbQuery)
		require.NoError(t, err)
		require.NotNil(t, revoked)
		require.False(t, revoked.Value)

		// All compliance info — entry present with status=2.
		allInfo, err := GetAllComplianceInfo()
		require.NoError(t, err)
		require.True(t, hasComplianceInfoStatus(allInfo, int32(vid), int32(pid), 2))

		// All device software compliance — entry present.
		allDsc, err := GetAllDeviceSoftwareCompliance()
		require.NoError(t, err)
		require.True(t, containsDeviceSoftwareCompliance(allDsc, cdCertID))

		// All revoked — this (vid, pid) no longer has any value=true entry for
		// the re-certified (sv, certType). (Matter revoked entry still has
		// value=false from earlier certification.)
		allRevoked, err := GetAllRevokedModels()
		require.NoError(t, err)
		require.False(t, hasRevokedWithValue(allRevoked, int32(vid), int32(pid), true))

		// All certified — both cert types
		allCertified, err := GetAllCertifiedModels()
		require.NoError(t, err)
		require.True(t, hasCertifiedModelCertType(allCertified, int32(vid), int32(pid), zigbeeCertType))
		require.True(t, hasCertifiedModelCertType(allCertified, int32(vid), int32(pid), matterCertType))
	})

	// Second pid/sv pair for optional-fields and update/delete tests.
	pid2 := rand.Intn(65534) + 1
	sv2 := rand.Intn(65534) + 1
	svs2 := fmt.Sprintf("%d", rand.Intn(65534)+1)

	t.Run("CreateModelAndVersion2", func(t *testing.T) {
		txResult, err := model.AddModel(model.AddModelOpts{
			VID: vid, PID: pid2, From: vendorAccount,
		})
		cliputils.RequireTxOK(t, txResult, err)

		txResult, err = model.AddModelVersion(model.AddModelVersionOpts{
			VID: vid, PID: pid2,
			SoftwareVersion:       sv2,
			SoftwareVersionString: svs2,
			CDVersionNumber:       cdVersionNumber,
			From:                  vendorAccount,
		})
		cliputils.RequireTxOK(t, txResult, err)
	})

	t.Run("CertifyWithOptionalFields", func(t *testing.T) {
		txResult, err := CertifyModel(CertifyModelOpts{
			VID: vid, PID: pid2,
			SoftwareVersion:       sv2,
			SoftwareVersionString: svs2,
			CertificationType:     zigbeeCertType,
			CertificationDate:     reCertDate,
			CDCertificateID:       cdCertID,
			CDVersionNumber:       cdVersionNumber,
			From:                  zbAccount,
			Optional: OptionalFields{
				ProgramTypeVersion:                 "1.0",
				FamilyID:                           "FAM123456abc",
				SupportedClusters:                  "0x0003,0x0004",
				CompliantPlatformUsed:              "WIFI",
				CompliantPlatformVersion:           "V1",
				OSVersion:                          "someV",
				CertificationRoute:                 "fullTested",
				ProgramType:                        "endProduct",
				Transport:                          "wi-fi",
				ParentChild:                        "parent",
				CertificationIDOfSoftwareComponent: "someIDOfSoftwareComponent",
			},
		})
		cliputils.RequireTxOK(t, txResult, err)

		// Certified model
		certified, err := GetCertifiedModel(ComplianceQueryOpts{VID: vid, PID: pid2, SoftwareVersion: sv2, CertificationType: zigbeeCertType})
		require.NoError(t, err)
		require.NotNil(t, certified)
		require.True(t, certified.Value)
		require.Equal(t, int32(vid), certified.Vid)
		require.Equal(t, int32(pid2), certified.Pid)

		// Compliance info — all optional fields present
		info, err := GetComplianceInfo(ComplianceQueryOpts{VID: vid, PID: pid2, SoftwareVersion: sv2, CertificationType: zigbeeCertType})
		require.NoError(t, err)
		require.NotNil(t, info)
		require.Equal(t, int32(vid), info.Vid)
		require.Equal(t, int32(pid2), info.Pid)
		require.Equal(t, uint32(2), info.SoftwareVersionCertificationStatus)
		require.Equal(t, reCertDate, info.Date)
		require.Equal(t, zigbeeCertType, info.CertificationType)
		require.Equal(t, uint32(1), info.SpecificationVersion)
		require.Equal(t, cdCertID, info.CDCertificateId)
		require.Equal(t, "1.0", info.ProgramTypeVersion)
		require.Equal(t, "FAM123456abc", info.FamilyId)
		require.Equal(t, "0x0003,0x0004", info.SupportedClusters)
		require.Equal(t, "WIFI", info.CompliantPlatformUsed)
		require.Equal(t, "V1", info.CompliantPlatformVersion)
		require.Equal(t, "someV", info.OSVersion)
		require.Equal(t, "fullTested", info.CertificationRoute)
		require.Equal(t, "endProduct", info.ProgramType)
		require.Equal(t, "wi-fi", info.Transport)
		require.Equal(t, "parent", info.ParentChild)
		require.Equal(t, "someIDOfSoftwareComponent", info.CertificationIdOfSoftwareComponent)

		// Device software compliance — all optional fields present on the pid2 entry.
		dsc, err := GetDeviceSoftwareCompliance(cdCertID)
		require.NoError(t, err)
		require.NotNil(t, dsc)
		var pid2Info *compliancetypes.ComplianceInfo
		for _, ci := range dsc.ComplianceInfo {
			if ci.Pid == int32(pid2) && ci.CertificationType == zigbeeCertType {
				pid2Info = ci

				break
			}
		}
		require.NotNil(t, pid2Info)
		require.Equal(t, int32(vid), pid2Info.Vid)
		require.Equal(t, uint32(2), pid2Info.SoftwareVersionCertificationStatus)
		require.Equal(t, reCertDate, pid2Info.Date)
		require.Equal(t, uint32(1), pid2Info.SpecificationVersion)
		require.Equal(t, cdCertID, pid2Info.CDCertificateId)
		require.Equal(t, "1.0", pid2Info.ProgramTypeVersion)
		require.Equal(t, "FAM123456abc", pid2Info.FamilyId)
		require.Equal(t, "0x0003,0x0004", pid2Info.SupportedClusters)
		require.Equal(t, "WIFI", pid2Info.CompliantPlatformUsed)
		require.Equal(t, "V1", pid2Info.CompliantPlatformVersion)
		require.Equal(t, "someV", pid2Info.OSVersion)
		require.Equal(t, "fullTested", pid2Info.CertificationRoute)
		require.Equal(t, "endProduct", pid2Info.ProgramType)
		require.Equal(t, "wi-fi", pid2Info.Transport)
		require.Equal(t, "parent", pid2Info.ParentChild)
		require.Equal(t, "someIDOfSoftwareComponent", pid2Info.CertificationIdOfSoftwareComponent)
	})

	t.Run("UpdateComplianceInfo_ByCertCenter", func(t *testing.T) {
		txResult, err := UpdateComplianceInfo(UpdateComplianceInfoOpts{
			VID: vid, PID: pid2, SoftwareVersion: sv2, CertificationType: zigbeeCertType, From: zbAccount,
			Reason: "new_reason",
			Optional: OptionalFields{
				ProgramType: "softwareComponent",
				ParentChild: "child",
				Transport:   "ethernet",
			},
		})
		cliputils.RequireTxOK(t, txResult, err)

		// Verify updated fields changed, non-updated optional fields unchanged
		zbQuery := ComplianceQueryOpts{VID: vid, PID: pid2, SoftwareVersion: sv2, CertificationType: zigbeeCertType}

		info, err := GetComplianceInfo(zbQuery)
		require.NoError(t, err)
		require.NotNil(t, info)
		require.Equal(t, int32(vid), info.Vid)
		require.Equal(t, int32(pid2), info.Pid)
		require.Equal(t, uint32(2), info.SoftwareVersionCertificationStatus)
		require.Equal(t, reCertDate, info.Date)
		require.Equal(t, zigbeeCertType, info.CertificationType)
		require.Equal(t, uint32(1), info.SpecificationVersion)
		require.Equal(t, cdCertID, info.CDCertificateId)
		require.Equal(t, "1.0", info.ProgramTypeVersion)
		require.Equal(t, "FAM123456abc", info.FamilyId)
		require.Equal(t, "0x0003,0x0004", info.SupportedClusters)
		require.Equal(t, "WIFI", info.CompliantPlatformUsed)
		require.Equal(t, "V1", info.CompliantPlatformVersion)
		require.Equal(t, "someV", info.OSVersion)
		require.Equal(t, "fullTested", info.CertificationRoute)
		// Updated fields
		require.Equal(t, "softwareComponent", info.ProgramType)
		require.Equal(t, "ethernet", info.Transport)
		require.Equal(t, "child", info.ParentChild)
		require.Equal(t, "new_reason", info.Reason)
		require.Equal(t, "someIDOfSoftwareComponent", info.CertificationIdOfSoftwareComponent)
	})

	t.Run("UpdateComplianceInfo_ByVendor_Fails", func(t *testing.T) {
		txResult, err := UpdateComplianceInfo(UpdateComplianceInfoOpts{
			VID: vid, PID: pid2, SoftwareVersion: sv2, CertificationType: zigbeeCertType, From: vendorAccount,
			Reason: "by_vendor_reason",
			Optional: OptionalFields{
				ProgramType: "compliantPlatform",
				ParentChild: "parent",
				Transport:   "bluetooth",
			},
		})
		cliputils.RequireTxFails(t, txResult, err)
		require.Contains(t, txResult.RawLog, "unauthorized")

		// Fields must be unchanged after vendor update attempt.
		info, err := GetComplianceInfo(ComplianceQueryOpts{VID: vid, PID: pid2, SoftwareVersion: sv2, CertificationType: zigbeeCertType})
		require.NoError(t, err)
		require.NotNil(t, info)
		require.Equal(t, "softwareComponent", info.ProgramType)
		require.Equal(t, "ethernet", info.Transport)
		require.Equal(t, "child", info.ParentChild)
		require.Equal(t, "new_reason", info.Reason)
	})

	t.Run("UpdateComplianceInfo_NoFields_Unchanged", func(t *testing.T) {
		txResult, err := UpdateComplianceInfo(UpdateComplianceInfoOpts{
			VID: vid, PID: pid2, SoftwareVersion: sv2, CertificationType: zigbeeCertType, From: zbAccount,
		})
		cliputils.RequireTxOK(t, txResult, err)

		// Fields must remain as set by the first update.
		info, err := GetComplianceInfo(ComplianceQueryOpts{VID: vid, PID: pid2, SoftwareVersion: sv2, CertificationType: zigbeeCertType})
		require.NoError(t, err)
		require.NotNil(t, info)
		require.Equal(t, "softwareComponent", info.ProgramType)
		require.Equal(t, "ethernet", info.Transport)
		require.Equal(t, "child", info.ParentChild)
		require.Equal(t, "new_reason", info.Reason)
		require.Equal(t, "someIDOfSoftwareComponent", info.CertificationIdOfSoftwareComponent)
	})

	updCdCertID := fmt.Sprintf("ucrt-%014d", rand.Intn(1<<30))

	t.Run("UpdateComplianceInfo_AllFields", func(t *testing.T) {
		txResult, err := UpdateComplianceInfo(UpdateComplianceInfoOpts{
			VID: vid, PID: pid2, SoftwareVersion: sv2, CertificationType: zigbeeCertType, From: zbAccount,
			CDVersionNumber:   1,
			CertificationDate: "2022-01-01T00:00:01Z",
			Reason:            "brand_new_reason",
			CDCertificateID:   updCdCertID,
			Optional: OptionalFields{
				CertificationRoute:                 "similarity",
				ProgramType:                        "endProduct",
				ProgramTypeVersion:                 "brand_new_program_type_version",
				CompliantPlatformUsed:              "brand_new_compliant_platform_used",
				CompliantPlatformVersion:           "brand_new_compliant_platform_version",
				Transport:                          "thread,nfc",
				FamilyID:                           "FAM123456abc",
				SupportedClusters:                  "0x0006,0x0008,0x0062",
				OSVersion:                          "brand_new_os_version",
				ParentChild:                        "parent",
				CertificationIDOfSoftwareComponent: "brand_new_component",
			},
		})
		cliputils.RequireTxOK(t, txResult, err)

		// All fields updated
		info, err := GetComplianceInfo(ComplianceQueryOpts{VID: vid, PID: pid2, SoftwareVersion: sv2, CertificationType: zigbeeCertType})
		require.NoError(t, err)
		require.NotNil(t, info)
		require.Equal(t, int32(vid), info.Vid)
		require.Equal(t, int32(pid2), info.Pid)
		require.Equal(t, uint32(sv2), info.SoftwareVersion)
		require.Equal(t, zigbeeCertType, info.CertificationType)
		require.Equal(t, uint32(1), info.SpecificationVersion)
		require.Equal(t, uint32(1), info.CDVersionNumber)
		require.Equal(t, "2022-01-01T00:00:01Z", info.Date)
		require.Equal(t, "brand_new_reason", info.Reason)
		require.Equal(t, updCdCertID, info.CDCertificateId)
		require.Equal(t, "similarity", info.CertificationRoute)
		require.Equal(t, "endProduct", info.ProgramType)
		require.Equal(t, "brand_new_program_type_version", info.ProgramTypeVersion)
		require.Equal(t, "brand_new_compliant_platform_used", info.CompliantPlatformUsed)
		require.Equal(t, "brand_new_compliant_platform_version", info.CompliantPlatformVersion)
		require.Equal(t, "thread,nfc", info.Transport)
		require.Equal(t, "FAM123456abc", info.FamilyId)
		require.Equal(t, "0x0006,0x0008,0x0062", info.SupportedClusters)
		require.Equal(t, "brand_new_os_version", info.OSVersion)
		require.Equal(t, "parent", info.ParentChild)
		require.Equal(t, "brand_new_component", info.CertificationIdOfSoftwareComponent)
		require.Equal(t, uint32(1), info.SchemaVersion)

		// Device software compliance accessible under the updated cdCertificateId
		dsc, err := GetDeviceSoftwareCompliance(updCdCertID)
		require.NoError(t, err)
		require.NotNil(t, dsc)
		require.True(t, containsComplianceInfoCertType(dsc.ComplianceInfo, int32(vid), int32(pid2), zigbeeCertType))
	})

	t.Run("DeleteComplianceInfo_AndVerify", func(t *testing.T) {
		txResult, err := DeleteComplianceInfo(vid, pid2, sv2, zigbeeCertType, zbAccount)
		cliputils.RequireTxOK(t, txResult, err)

		// Compliance info — gone.
		info, err := GetComplianceInfo(ComplianceQueryOpts{VID: vid, PID: pid2, SoftwareVersion: sv2, CertificationType: zigbeeCertType})
		require.NoError(t, err)
		require.Nil(t, info)

		// Device software compliance — entry for pid2 no longer present in the
		// original cdCertID record.
		dsc, err := GetDeviceSoftwareCompliance(cdCertID)
		require.NoError(t, err)
		if dsc != nil {
			for _, ci := range dsc.ComplianceInfo {
				require.NotEqual(t, int32(pid2), ci.Pid)
			}
		}

		// Certified model — gone.
		certified, err := GetCertifiedModel(ComplianceQueryOpts{VID: vid, PID: pid2, SoftwareVersion: sv2, CertificationType: zigbeeCertType})
		require.NoError(t, err)
		require.Nil(t, certified)
	})
}
