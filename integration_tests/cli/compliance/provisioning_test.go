package compliance

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
	cliputils "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/cli/utils"
	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/utils"
	compliancetypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/compliance/types"
)

// TestComplianceProvisioning exercises the provisional → certified / revoked
// state machine for both the zigbee and matter certification types, the
// provision error codes (305 already-provisional, 303 already-certified,
// 304 already-revoked), the single-record and all-* queries at each stage, and
// provision/certify with the full set of optional fields.
//
// NOTE on assertions: the original CLI checks verified state with substring
// matching over the JSON output, so e.g. they asserted
// "softwareVersionCertificationStatus": 1 after a certify — that "1" actually
// matches the provisional record preserved in the `history` array, not the
// (now 2) top-level status. The typed assertions below assert the correct
// top-level value (CodeCertified) and check `history` only after a genuine
// state transition, where it is populated.
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
	certificationDate := "2021-02-02T02:20:19Z"
	certificationReason := "some reason 2"
	revocationDate := "2021-02-02T02:20:20Z"
	revocationReasonZb := "some reason 11"
	revocationReasonMatter := "some reason 22"
	cdCertID := fmt.Sprintf("cert-%014d", rand.Intn(1<<30))

	// findDSCEntry returns the ComplianceInfo for the given pid within a
	// device-software-compliance record, or nil.
	findDSCEntry := func(dsc *compliancetypes.DeviceSoftwareCompliance, pid int32) *compliancetypes.ComplianceInfo {
		if dsc == nil {
			return nil
		}
		for _, ci := range dsc.ComplianceInfo {
			if ci != nil && ci.Pid == pid {
				return ci
			}
		}

		return nil
	}

	// provisionalCertTypes collects the certification types present in a
	// provisional-model list for the given (vid, pid).
	provisionalCertTypes := func(list []compliancetypes.ProvisionalModel, vid, pid int32) map[string]bool {
		m := map[string]bool{}
		for i := range list {
			if list[i].Vid == vid && list[i].Pid == pid {
				m[list[i].CertificationType] = true
			}
		}

		return m
	}

	// ── Provision an unknown model (no model record yet), verify, delete ──

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
		cliputils.RequireTxOK(t, txResult, err)

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

		// Delete compliance info before creating the model.
		txResult, err = DeleteComplianceInfo(vid, pid, sv, certTypeZb, zbAccount)
		require.NoError(t, err)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)
	})

	t.Run("CreateModelAndVersion", func(t *testing.T) {
		cliputils.CreateModelAndVersion(t, vid, pid, sv, svs, vendorAccount)
	})

	// ── Provision the (now-existing) model for ZB then Matter ──

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
		cliputils.RequireTxOK(t, txResult, err)
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
		cliputils.RequireTxOK(t, txResult, err)
	})

	// ── Re-provision is rejected whether by a different or the same account ──

	t.Run("ReProvision_DifferentAccount_Fails", func(t *testing.T) {
		// A different CertificationCenter account cannot re-provision an already
		// provisional model — the state check (305) precedes any ownership check.
		txResult, err := ProvisionModel(ProvisionModelOpts{
			VID: vid, PID: pid,
			SoftwareVersion:       sv,
			SoftwareVersionString: svs,
			CertificationType:     certTypeZb,
			ProvisionalDate:       provisionDate,
			CDCertificateID:       cdCertID,
			Reason:                provisionReason,
			From:                  secondZbAccount,
		})
		require.NoError(t, err)
		require.Equal(t, uint32(305), txResult.Code)
		require.Contains(t, txResult.RawLog, "already in provisional state")
	})

	t.Run("ReProvision_SameAccount_Fails", func(t *testing.T) {
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

	// ── While provisional: query single records and verify absent states ──

	t.Run("QueryProvisionalAndAbsentStates", func(t *testing.T) {
		for _, ct := range []string{certTypeZb, certTypeMatter} {
			provisional, err := GetProvisionalModel(ComplianceQueryOpts{VID: vid, PID: pid, SoftwareVersion: sv, CertificationType: ct})
			require.NoError(t, err)
			require.NotNil(t, provisional, "provisional-model missing for certType %s", ct)
			require.True(t, provisional.Value)
			require.Equal(t, int32(vid), provisional.Vid)
			require.Equal(t, int32(pid), provisional.Pid)
			require.Equal(t, ct, provisional.CertificationType)

			// Not certified or revoked yet — those records do not exist.
			certified, err := GetCertifiedModel(ComplianceQueryOpts{VID: vid, PID: pid, SoftwareVersion: sv, CertificationType: ct})
			require.NoError(t, err)
			require.Nil(t, certified, "certified-model must be absent for certType %s", ct)

			revoked, err := GetRevokedModel(ComplianceQueryOpts{VID: vid, PID: pid, SoftwareVersion: sv, CertificationType: ct})
			require.NoError(t, err)
			require.Nil(t, revoked, "revoked-model must be absent for certType %s", ct)
		}
	})

	t.Run("QueryComplianceInfoWhileProvisional", func(t *testing.T) {
		for _, ct := range []string{certTypeZb, certTypeMatter} {
			info, err := GetComplianceInfo(ComplianceQueryOpts{VID: vid, PID: pid, SoftwareVersion: sv, CertificationType: ct})
			require.NoError(t, err)
			require.NotNil(t, info)
			require.Equal(t, int32(vid), info.Vid)
			require.Equal(t, int32(pid), info.Pid)
			require.Equal(t, compliancetypes.CodeProvisional, info.SoftwareVersionCertificationStatus)
			require.Equal(t, cdCertID, info.CDCertificateId)
			require.Equal(t, provisionDate, info.Date)
			require.Equal(t, provisionReason, info.Reason)
			require.Equal(t, ct, info.CertificationType)
		}
	})

	t.Run("QueryAllListsWhileProvisional", func(t *testing.T) {
		// all-compliance-info contains our (vid,pid) for both certTypes.
		allInfo, err := GetAllComplianceInfo()
		require.NoError(t, err)
		require.True(t, containsComplianceInfo(allInfo, int32(vid), int32(pid)))

		// No device-software-compliance, revoked or certified record yet.
		allDSC, err := GetAllDeviceSoftwareCompliance()
		require.NoError(t, err)
		require.False(t, containsDeviceSoftwareCompliance(allDSC, cdCertID))

		allRevoked, err := GetAllRevokedModels()
		require.NoError(t, err)
		require.False(t, containsRevokedModel(allRevoked, int32(vid), int32(pid)))

		allCertified, err := GetAllCertifiedModels()
		require.NoError(t, err)
		require.False(t, containsCertifiedModel(allCertified, int32(vid), int32(pid)))

		// all-provisional-models contains both certTypes for our (vid,pid).
		allProvisional, err := GetAllProvisionalModels()
		require.NoError(t, err)
		cts := provisionalCertTypes(allProvisional, int32(vid), int32(pid))
		require.True(t, cts[certTypeZb], "all-provisional-models missing zigbee entry")
		require.True(t, cts[certTypeMatter], "all-provisional-models missing matter entry")
	})

	// ── Second model (pid2): certify for matter, then provision is rejected ──

	pid2 := rand.Intn(65534) + 1
	sv2 := rand.Intn(65534) + 1
	svs2 := fmt.Sprintf("%d", rand.Intn(65534)+1)

	t.Run("Pid2_CertifyMatter_ThenProvisionFails", func(t *testing.T) {
		cliputils.CreateModelAndVersion(t, vid, pid2, sv2, svs2, vendorAccount)

		txResult, err := CertifyModel(CertifyModelOpts{
			VID: vid, PID: pid2,
			SoftwareVersion:       sv2,
			SoftwareVersionString: svs2,
			CertificationType:     certTypeMatter,
			CertificationDate:     certificationDate,
			CDCertificateID:       cdCertID,
			From:                  zbAccount,
			Reason:                certificationReason,
		})
		cliputils.RequireTxOK(t, txResult, err)

		// device-software-compliance for cdCertID now contains pid2.
		dsc, err := GetDeviceSoftwareCompliance(cdCertID)
		require.NoError(t, err)
		require.NotNil(t, dsc)
		entry := findDSCEntry(dsc, int32(pid2))
		require.NotNil(t, entry, "device-software-compliance missing pid2 entry")
		require.Equal(t, int32(vid), entry.Vid)
		require.Equal(t, uint32(sv2), entry.SoftwareVersion)
		require.Equal(t, svs2, entry.SoftwareVersionString)
		require.Equal(t, certTypeMatter, entry.CertificationType)
		require.Equal(t, certificationDate, entry.Date)
		require.Equal(t, certificationReason, entry.Reason)

		// Provisioning an already-certified model is rejected with 303.
		txResult, err = ProvisionModel(ProvisionModelOpts{
			VID: vid, PID: pid2,
			SoftwareVersion:       sv2,
			SoftwareVersionString: svs2,
			CertificationType:     certTypeMatter,
			ProvisionalDate:       provisionDate,
			CDCertificateID:       cdCertID,
			Reason:                provisionReason,
			From:                  zbAccount,
		})
		require.NoError(t, err)
		require.Equal(t, uint32(303), txResult.Code)
		require.Contains(t, txResult.RawLog, "already certified on the ledger")
	})

	// ── Third model (pid3): revoke uncertified, then provision is rejected ──

	pid3 := rand.Intn(65534) + 1
	sv3 := rand.Intn(65534) + 1
	svs3 := fmt.Sprintf("%d", rand.Intn(65534)+1)

	t.Run("Pid3_RevokeUncertified_ThenProvisionFails", func(t *testing.T) {
		cliputils.CreateModelAndVersion(t, vid, pid3, sv3, svs3, vendorAccount)

		txResult, err := RevokeModel(RevokeModelOpts{
			VID: vid, PID: pid3,
			SoftwareVersion:       sv3,
			SoftwareVersionString: svs3,
			CertificationType:     certTypeZb,
			RevocationDate:        revocationDate,
			Reason:                revocationReasonZb,
			From:                  zbAccount,
		})
		cliputils.RequireTxOK(t, txResult, err)

		// Provisioning an already-revoked model is rejected with 304.
		txResult, err = ProvisionModel(ProvisionModelOpts{
			VID: vid, PID: pid3,
			SoftwareVersion:       sv3,
			SoftwareVersionString: svs3,
			CertificationType:     certTypeZb,
			ProvisionalDate:       provisionDate,
			CDCertificateID:       cdCertID,
			Reason:                provisionReason,
			From:                  zbAccount,
		})
		require.NoError(t, err)
		require.Equal(t, uint32(304), txResult.Code)
		require.Contains(t, txResult.RawLog, "already revoked on the ledger")
	})

	// ── Certify the provisioned model (ZB) and revoke it (Matter) ──

	t.Run("CertifyProvisionedModel_ZB", func(t *testing.T) {
		txResult, err := CertifyModel(CertifyModelOpts{
			VID: vid, PID: pid,
			SoftwareVersion:       sv,
			SoftwareVersionString: svs,
			CertificationType:     certTypeZb,
			CertificationDate:     certificationDate,
			CDCertificateID:       cdCertID,
			Reason:                certificationReason,
			From:                  zbAccount,
		})
		cliputils.RequireTxOK(t, txResult, err)
	})

	t.Run("RevokeProvisionedModel_Matter", func(t *testing.T) {
		txResult, err := RevokeModel(RevokeModelOpts{
			VID: vid, PID: pid,
			SoftwareVersion:       sv,
			SoftwareVersionString: svs,
			CertificationType:     certTypeMatter,
			RevocationDate:        revocationDate,
			Reason:                revocationReasonMatter,
			From:                  zbAccount,
		})
		cliputils.RequireTxOK(t, txResult, err)
	})

	// ── After certify(ZB)+revoke(Matter): verify the per-record states ──

	t.Run("QueryStatesAfterCertifyRevoke", func(t *testing.T) {
		// provisional is now false for both certTypes.
		for _, ct := range []string{certTypeZb, certTypeMatter} {
			provisional, err := GetProvisionalModel(ComplianceQueryOpts{VID: vid, PID: pid, SoftwareVersion: sv, CertificationType: ct})
			require.NoError(t, err)
			require.NotNil(t, provisional)
			require.False(t, provisional.Value, "provisional should be false for certType %s", ct)
		}

		// certified: true for ZB, false for matter.
		certifiedZb, err := GetCertifiedModel(ComplianceQueryOpts{VID: vid, PID: pid, SoftwareVersion: sv, CertificationType: certTypeZb})
		require.NoError(t, err)
		require.NotNil(t, certifiedZb)
		require.True(t, certifiedZb.Value)

		certifiedMatter, err := GetCertifiedModel(ComplianceQueryOpts{VID: vid, PID: pid, SoftwareVersion: sv, CertificationType: certTypeMatter})
		require.NoError(t, err)
		require.NotNil(t, certifiedMatter)
		require.False(t, certifiedMatter.Value)

		// revoked: false for ZB, true for matter.
		revokedZb, err := GetRevokedModel(ComplianceQueryOpts{VID: vid, PID: pid, SoftwareVersion: sv, CertificationType: certTypeZb})
		require.NoError(t, err)
		require.NotNil(t, revokedZb)
		require.False(t, revokedZb.Value)

		revokedMatter, err := GetRevokedModel(ComplianceQueryOpts{VID: vid, PID: pid, SoftwareVersion: sv, CertificationType: certTypeMatter})
		require.NoError(t, err)
		require.NotNil(t, revokedMatter)
		require.True(t, revokedMatter.Value)
	})

	t.Run("QueryComplianceInfoAfterCertifyRevoke", func(t *testing.T) {
		// ZB is certified (status 2) with the certification date/reason; history
		// is populated by the provisional → certified transition.
		infoZb, err := GetComplianceInfo(ComplianceQueryOpts{VID: vid, PID: pid, SoftwareVersion: sv, CertificationType: certTypeZb})
		require.NoError(t, err)
		require.NotNil(t, infoZb)
		require.Equal(t, compliancetypes.CodeCertified, infoZb.SoftwareVersionCertificationStatus)
		require.Equal(t, cdCertID, infoZb.CDCertificateId)
		require.Equal(t, certificationDate, infoZb.Date)
		require.Equal(t, certificationReason, infoZb.Reason)
		require.Equal(t, certTypeZb, infoZb.CertificationType)
		require.NotEmpty(t, infoZb.History)

		// Matter is revoked (status 3) with the revocation date/reason.
		infoMatter, err := GetComplianceInfo(ComplianceQueryOpts{VID: vid, PID: pid, SoftwareVersion: sv, CertificationType: certTypeMatter})
		require.NoError(t, err)
		require.NotNil(t, infoMatter)
		require.Equal(t, compliancetypes.CodeRevoked, infoMatter.SoftwareVersionCertificationStatus)
		require.Equal(t, revocationDate, infoMatter.Date)
		require.Equal(t, revocationReasonMatter, infoMatter.Reason)
		require.Equal(t, certTypeMatter, infoMatter.CertificationType)
		require.NotEmpty(t, infoMatter.History)
	})

	t.Run("QueryAllListsAfterCertifyRevoke", func(t *testing.T) {
		// all-compliance-info has our (vid,pid) for both the certified (status 2)
		// and revoked (status 3) records.
		allInfo, err := GetAllComplianceInfo()
		require.NoError(t, err)
		require.True(t, hasComplianceInfoStatus(allInfo, int32(vid), int32(pid), compliancetypes.CodeCertified))
		require.True(t, hasComplianceInfoStatus(allInfo, int32(vid), int32(pid), compliancetypes.CodeRevoked))

		// all-device-software-compliance has the cdCertID with a pid2 entry.
		allDSC, err := GetAllDeviceSoftwareCompliance()
		require.NoError(t, err)
		require.True(t, containsDeviceSoftwareCompliance(allDSC, cdCertID))

		// all-revoked-models contains (vid,pid) for matter.
		allRevoked, err := GetAllRevokedModels()
		require.NoError(t, err)
		require.True(t, containsRevokedModel(allRevoked, int32(vid), int32(pid)))

		// all-certified-models contains (vid,pid) for zigbee.
		allCertified, err := GetAllCertifiedModels()
		require.NoError(t, err)
		require.True(t, hasCertifiedModelCertType(allCertified, int32(vid), int32(pid), certTypeZb))

		// (vid,pid) is no longer provisional for either certType.
		allProvisional, err := GetAllProvisionalModels()
		require.NoError(t, err)
		require.True(t, hasProvisionalWithValue(allProvisional, int32(vid), int32(pid), false) ||
			!containsProvisionalModel(allProvisional, int32(vid), int32(pid)))
	})

	// ── Provision a fresh model with the full set of optional fields ──

	pidOpt := rand.Intn(65534) + 1
	svOpt := rand.Intn(65534) + 1
	svsOpt := fmt.Sprintf("%d", rand.Intn(65534)+1)

	t.Run("ProvisionWithAllOptionalFields", func(t *testing.T) {
		cliputils.CreateModelAndVersion(t, vid, pidOpt, svOpt, svsOpt, vendorAccount)

		txResult, err := ProvisionModel(ProvisionModelOpts{
			VID: vid, PID: pidOpt,
			SoftwareVersion:       svOpt,
			SoftwareVersionString: svsOpt,
			CertificationType:     certTypeZb,
			ProvisionalDate:       provisionDate,
			CDCertificateID:       cdCertID,
			Reason:                provisionReason,
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
				CertificationIDOfSoftwareComponent: "someIDOfSoftwareComponent1",
			},
		})
		cliputils.RequireTxOK(t, txResult, err)

		provisional, err := GetProvisionalModel(ComplianceQueryOpts{VID: vid, PID: pidOpt, SoftwareVersion: svOpt, CertificationType: certTypeZb})
		require.NoError(t, err)
		require.NotNil(t, provisional)
		require.True(t, provisional.Value)

		info, err := GetComplianceInfo(ComplianceQueryOpts{VID: vid, PID: pidOpt, SoftwareVersion: svOpt, CertificationType: certTypeZb})
		require.NoError(t, err)
		require.NotNil(t, info)
		require.Equal(t, compliancetypes.CodeProvisional, info.SoftwareVersionCertificationStatus)
		require.Equal(t, provisionReason, info.Reason)
		require.Equal(t, provisionDate, info.Date)
		require.Equal(t, certTypeZb, info.CertificationType)
		require.Equal(t, uint32(1), info.SpecificationVersion)
		require.Equal(t, "1.0", info.ProgramTypeVersion)
		require.Equal(t, cdCertID, info.CDCertificateId)
		require.Equal(t, "FAM123456abc", info.FamilyId)
		require.Equal(t, "0x0003,0x0004", info.SupportedClusters)
		require.Equal(t, "WIFI", info.CompliantPlatformUsed)
		require.Equal(t, "V1", info.CompliantPlatformVersion)
		require.Equal(t, "someV", info.OSVersion)
		require.Equal(t, "fullTested", info.CertificationRoute)
		require.Equal(t, "endProduct", info.ProgramType)
		require.Equal(t, "wi-fi", info.Transport)
		require.Equal(t, "parent", info.ParentChild)
		require.Equal(t, "someIDOfSoftwareComponent1", info.CertificationIdOfSoftwareComponent)
	})

	// ── Certify the same model with SOME optional fields: certify updates the
	// supplied fields and leaves the others (OSVersion, route, transport,
	// parentChild) as provisioned. ──

	t.Run("CertifyWithSomeOptionalFields", func(t *testing.T) {
		txResult, err := CertifyModel(CertifyModelOpts{
			VID: vid, PID: pidOpt,
			SoftwareVersion:       svOpt,
			SoftwareVersionString: svsOpt,
			CertificationType:     certTypeZb,
			CertificationDate:     certificationDate,
			CDCertificateID:       cdCertID,
			From:                  zbAccount,
			Optional: OptionalFields{
				ProgramType:                        "softwareComponent",
				ProgramTypeVersion:                 "2.0",
				FamilyID:                           "FAM54321cba",
				SupportedClusters:                  "0x0006,0x0008",
				CompliantPlatformUsed:              "ETHERNET",
				CompliantPlatformVersion:           "V2",
				CertificationIDOfSoftwareComponent: "someIDOfSoftwareComponent2",
			},
		})
		cliputils.RequireTxOK(t, txResult, err)

		certified, err := GetCertifiedModel(ComplianceQueryOpts{VID: vid, PID: pidOpt, SoftwareVersion: svOpt, CertificationType: certTypeZb})
		require.NoError(t, err)
		require.NotNil(t, certified)
		require.True(t, certified.Value)

		// Compliance info: top-level status is now certified (a "1" would only
		// match the provisional record preserved in history). Updated fields
		// reflect the certify; OSVersion/route/transport/parentChild persist.
		info, err := GetComplianceInfo(ComplianceQueryOpts{VID: vid, PID: pidOpt, SoftwareVersion: svOpt, CertificationType: certTypeZb})
		require.NoError(t, err)
		require.NotNil(t, info)
		require.Equal(t, compliancetypes.CodeCertified, info.SoftwareVersionCertificationStatus)
		require.Equal(t, certificationDate, info.Date)
		require.Equal(t, certTypeZb, info.CertificationType)
		require.Equal(t, uint32(1), info.SpecificationVersion)
		require.Equal(t, "2.0", info.ProgramTypeVersion)
		require.Equal(t, cdCertID, info.CDCertificateId)
		require.Equal(t, "FAM54321cba", info.FamilyId)
		require.Equal(t, "0x0006,0x0008", info.SupportedClusters)
		require.Equal(t, "ETHERNET", info.CompliantPlatformUsed)
		require.Equal(t, "V2", info.CompliantPlatformVersion)
		require.Equal(t, "someV", info.OSVersion)
		require.Equal(t, "fullTested", info.CertificationRoute)
		require.Equal(t, "softwareComponent", info.ProgramType)
		require.Equal(t, "wi-fi", info.Transport)
		require.Equal(t, "parent", info.ParentChild)
		require.Equal(t, "someIDOfSoftwareComponent2", info.CertificationIdOfSoftwareComponent)

		// device-software-compliance for cdCertID now contains the pidOpt entry
		// with the certify-updated optional fields.
		dsc, err := GetDeviceSoftwareCompliance(cdCertID)
		require.NoError(t, err)
		require.NotNil(t, dsc)
		entry := findDSCEntry(dsc, int32(pidOpt))
		require.NotNil(t, entry, "device-software-compliance missing pidOpt entry")
		require.Equal(t, int32(vid), entry.Vid)
		require.Equal(t, certTypeZb, entry.CertificationType)
		require.Equal(t, "2.0", entry.ProgramTypeVersion)
		require.Equal(t, "FAM54321cba", entry.FamilyId)
		require.Equal(t, "0x0006,0x0008", entry.SupportedClusters)
		require.Equal(t, "ETHERNET", entry.CompliantPlatformUsed)
		require.Equal(t, "V2", entry.CompliantPlatformVersion)
		require.Equal(t, "someV", entry.OSVersion)
		require.Equal(t, "fullTested", entry.CertificationRoute)
		require.Equal(t, "softwareComponent", entry.ProgramType)
		require.Equal(t, "wi-fi", entry.Transport)
		require.Equal(t, "parent", entry.ParentChild)
		require.Equal(t, "someIDOfSoftwareComponent2", entry.CertificationIdOfSoftwareComponent)
	})
}
