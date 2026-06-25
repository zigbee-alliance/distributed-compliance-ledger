package compliance

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
	cliputils "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/cli/utils"
	compliancetypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/compliance/types"
)

// TestComplianceRevocation revokes an uncertified model for both zigbee and
// matter, exercises the re-revoke error path (304), verifies every per-record
// and all-* query while revoked, then re-certifies the revoked model (covering
// the 302 "must be after" guard) and re-verifies the resulting states. It
// finishes with the compliance schema-v1 (#730) ValidateBasic negative cases.
//
// NOTE on assertions: the original CLI checks verified state with substring
// matching over the JSON output. After a certify they asserted
// "softwareVersionCertificationStatus": 2, but a revoke followed by certify
// keeps the prior states in the `history` array, so the typed assertions below
// check the correct top-level value and only inspect
// `history` where it is meaningfully populated. The all-* list queries skip
// Value=false entries (see grpc_query_*_model.go), so a re-certified zigbee
// model drops out of all-revoked-models while the still-revoked matter remains.
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
	cdCertID := fmt.Sprintf("cert-%014d", rand.Intn(1<<30)) // 19 chars: "cert-" + 14 digits

	revocationDate := "2020-02-02T02:20:20Z"
	revocationReason := "some reason"
	certificationDate := "2020-02-02T02:20:21Z"
	certificationDatePast := "2020-02-02T02:20:19Z"
	certificationReason := "some reason 2"

	// findDSCEntry returns the ComplianceInfo for pid within a
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

	revokedCertTypes := func(list []compliancetypes.RevokedModel, vid, pid int32) map[string]bool {
		m := map[string]bool{}
		for i := range list {
			if list[i].Vid == vid && list[i].Pid == pid {
				m[list[i].CertificationType] = true
			}
		}

		return m
	}
	certifiedCertTypes := func(list []compliancetypes.CertifiedModel, vid, pid int32) map[string]bool {
		m := map[string]bool{}
		for i := range list {
			if list[i].Vid == vid && list[i].Pid == pid {
				m[list[i].CertificationType] = true
			}
		}

		return m
	}
	allInfoCertTypes := func(list []compliancetypes.ComplianceInfo, vid, pid int32) map[string]bool {
		m := map[string]bool{}
		for i := range list {
			if list[i].Vid == vid && list[i].Pid == pid {
				m[list[i].CertificationType] = true
			}
		}

		return m
	}

	t.Run("CreateModelAndVersion", func(t *testing.T) {
		cliputils.CreateModelAndVersion(t, vid, pid, sv, svs, vendorAccount)
	})

	// ── Revoke an uncertified model for ZB ──

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
		cliputils.RequireTxOK(t, txResult, err)
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
		cliputils.RequireTxOK(t, txResult, err)
	})

	// ── Verify all states while revoked (ZB + Matter) ──

	t.Run("QueryAfterRevocation", func(t *testing.T) {
		for _, ct := range []string{certType, certTypeMatter} {
			// Not certified.
			certified, err := GetCertifiedModel(ComplianceQueryOpts{VID: vid, PID: pid, SoftwareVersion: sv, CertificationType: ct})
			require.NoError(t, err)
			require.Nil(t, certified, "certified-model must be absent for certType %s", ct)

			// Revoked (value true).
			revoked, err := GetRevokedModel(ComplianceQueryOpts{VID: vid, PID: pid, SoftwareVersion: sv, CertificationType: ct})
			require.NoError(t, err)
			require.NotNil(t, revoked)
			require.True(t, revoked.Value)
			require.Equal(t, int32(vid), revoked.Vid)
			require.Equal(t, int32(pid), revoked.Pid)
			require.Equal(t, ct, revoked.CertificationType)

			// Compliance info: status 3 (revoked) with revocation date/reason.
			// (History stays empty here — this is a fresh revoke of an uncertified
			// model, with no prior state to push into history.)
			info, err := GetComplianceInfo(ComplianceQueryOpts{VID: vid, PID: pid, SoftwareVersion: sv, CertificationType: ct})
			require.NoError(t, err)
			require.NotNil(t, info)
			require.Equal(t, compliancetypes.CodeRevoked, info.SoftwareVersionCertificationStatus)
			require.Equal(t, revocationDate, info.Date)
			require.Equal(t, revocationReason, info.Reason)
			require.Equal(t, ct, info.CertificationType)
		}

		// Never provisioned — record does not exist.
		provisional, err := GetProvisionalModel(ComplianceQueryOpts{VID: vid, PID: pid, SoftwareVersion: sv, CertificationType: certType})
		require.NoError(t, err)
		require.Nil(t, provisional)

		// all-compliance-info has both certTypes for (vid,pid), both revoked.
		allInfo, err := GetAllComplianceInfo()
		require.NoError(t, err)
		require.True(t, hasComplianceInfoStatus(allInfo, int32(vid), int32(pid), compliancetypes.CodeRevoked))
		cts := allInfoCertTypes(allInfo, int32(vid), int32(pid))
		require.True(t, cts[certType] && cts[certTypeMatter], "all-compliance-info missing a certType: %v", cts)

		// all-revoked-models has both certTypes; all-certified and all-provisional
		// do not contain (vid,pid) yet.
		allRevoked, err := GetAllRevokedModels()
		require.NoError(t, err)
		rcts := revokedCertTypes(allRevoked, int32(vid), int32(pid))
		require.True(t, rcts[certType] && rcts[certTypeMatter], "all-revoked-models missing a certType: %v", rcts)

		allCertified, err := GetAllCertifiedModels()
		require.NoError(t, err)
		require.False(t, containsCertifiedModel(allCertified, int32(vid), int32(pid)))

		allProvisional, err := GetAllProvisionalModels()
		require.NoError(t, err)
		require.False(t, containsProvisionalModel(allProvisional, int32(vid), int32(pid)))
	})

	// ── Re-certifying a revoked model with a date before its revocation is
	// rejected (302 "must be after"). ──

	t.Run("CertifyRevokedFromPast_Fails", func(t *testing.T) {
		txResult, err := CertifyModel(CertifyModelOpts{
			VID: vid, PID: pid,
			SoftwareVersion:       sv,
			SoftwareVersionString: svs,
			CertificationType:     certType,
			CertificationDate:     certificationDatePast,
			CDCertificateID:       cdCertID,
			From:                  zbAccount,
			Reason:                certificationReason,
		})
		require.NoError(t, err)
		require.Equal(t, uint32(302), txResult.Code)
		require.Contains(t, txResult.RawLog, "must be after")
	})

	// ── Re-certify the revoked ZB model successfully, then verify states ──

	t.Run("CertifyAfterRevoke_Success", func(t *testing.T) {
		txResult, err := CertifyModel(CertifyModelOpts{
			VID: vid, PID: pid,
			SoftwareVersion:       sv,
			SoftwareVersionString: svs,
			CertificationType:     certType,
			CertificationDate:     certificationDate,
			CDCertificateID:       cdCertID,
			From:                  zbAccount,
			Reason:                certificationReason,
		})
		cliputils.RequireTxOK(t, txResult, err)

		// revoked: ZB now false, matter still true.
		revokedZb, err := GetRevokedModel(ComplianceQueryOpts{VID: vid, PID: pid, SoftwareVersion: sv, CertificationType: certType})
		require.NoError(t, err)
		require.NotNil(t, revokedZb)
		require.False(t, revokedZb.Value)

		revokedMatter, err := GetRevokedModel(ComplianceQueryOpts{VID: vid, PID: pid, SoftwareVersion: sv, CertificationType: certTypeMatter})
		require.NoError(t, err)
		require.NotNil(t, revokedMatter)
		require.True(t, revokedMatter.Value)

		// certified: ZB now true, matter absent.
		certifiedZb, err := GetCertifiedModel(ComplianceQueryOpts{VID: vid, PID: pid, SoftwareVersion: sv, CertificationType: certType})
		require.NoError(t, err)
		require.NotNil(t, certifiedZb)
		require.True(t, certifiedZb.Value)

		certifiedMatter, err := GetCertifiedModel(ComplianceQueryOpts{VID: vid, PID: pid, SoftwareVersion: sv, CertificationType: certTypeMatter})
		require.NoError(t, err)
		require.Nil(t, certifiedMatter)

		// compliance-info ZB: status 2 (certified) with certification date/reason,
		// cdCertID, and a populated history (revoked → certified transition).
		info, err := GetComplianceInfo(ComplianceQueryOpts{VID: vid, PID: pid, SoftwareVersion: sv, CertificationType: certType})
		require.NoError(t, err)
		require.NotNil(t, info)
		require.Equal(t, compliancetypes.CodeCertified, info.SoftwareVersionCertificationStatus)
		require.Equal(t, certificationDate, info.Date)
		require.Equal(t, certificationReason, info.Reason)
		require.Equal(t, certType, info.CertificationType)
		require.Equal(t, cdCertID, info.CDCertificateId)
		require.NotEmpty(t, info.History)

		// device-software-compliance for cdCertID now has the certified ZB record.
		dsc, err := GetDeviceSoftwareCompliance(cdCertID)
		require.NoError(t, err)
		require.NotNil(t, dsc)
		require.Equal(t, cdCertID, dsc.CDCertificateId)
		entry := findDSCEntry(dsc, int32(pid))
		require.NotNil(t, entry, "device-software-compliance missing pid entry")
		require.Equal(t, int32(vid), entry.Vid)
		require.Equal(t, compliancetypes.CodeCertified, entry.SoftwareVersionCertificationStatus)
		require.Equal(t, certificationDate, entry.Date)
		require.Equal(t, certificationReason, entry.Reason)
		require.Equal(t, certType, entry.CertificationType)
		require.Equal(t, cdCertID, entry.CDCertificateId)

		// all-compliance-info has both a certified (2) and revoked (3) record.
		allInfo, err := GetAllComplianceInfo()
		require.NoError(t, err)
		require.True(t, hasComplianceInfoStatus(allInfo, int32(vid), int32(pid), compliancetypes.CodeCertified))
		require.True(t, hasComplianceInfoStatus(allInfo, int32(vid), int32(pid), compliancetypes.CodeRevoked))

		// all-revoked-models now has matter only (ZB dropped — value false skipped).
		allRevoked, err := GetAllRevokedModels()
		require.NoError(t, err)
		rcts := revokedCertTypes(allRevoked, int32(vid), int32(pid))
		require.True(t, rcts[certTypeMatter], "all-revoked-models missing matter")
		require.False(t, rcts[certType], "all-revoked-models should not contain zigbee after re-certify")

		// all-certified-models now has zigbee only.
		allCertified, err := GetAllCertifiedModels()
		require.NoError(t, err)
		ccts := certifiedCertTypes(allCertified, int32(vid), int32(pid))
		require.True(t, ccts[certType], "all-certified-models missing zigbee")
		require.False(t, ccts[certTypeMatter], "all-certified-models should not contain matter")

		// provisional ZB is false; (vid,pid) absent from all-provisional-models.
		provisional, err := GetProvisionalModel(ComplianceQueryOpts{VID: vid, PID: pid, SoftwareVersion: sv, CertificationType: certType})
		require.NoError(t, err)
		require.NotNil(t, provisional)
		require.False(t, provisional.Value)

		allProvisional, err := GetAllProvisionalModels()
		require.NoError(t, err)
		require.False(t, containsProvisionalModel(allProvisional, int32(vid), int32(pid)))
	})
}

// TestComplianceSchemaV1Negative covers the schema-v1 (#730) ValidateBasic
// negative cases for certify/provision/revoke. Each fails client-side before
// broadcast (surfaced via the returned error), so the typed helpers express the
// invalid input directly: SchemaVersion "0" exercises the "must be equal 1"
// guard and SpecVersionZero the required-field guard.
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
	date := "2020-02-02T02:20:20Z"
	cdCertID := "12345678910abcdefgh" // exactly 19 chars

	t.Run("CertifyModel_SchemaVersion0_Fails", func(t *testing.T) {
		txResult, err := CertifyModel(CertifyModelOpts{
			VID: vid, PID: pid, SoftwareVersion: sv, SoftwareVersionString: svs,
			CertificationType: certType, CertificationDate: date,
			CDCertificateID: cdCertID, SchemaVersion: "0", From: zbAccount,
		})
		cliputils.RequireTxFailContains(t, txResult, err, "SchemaVersion must be equal 1")
	})

	t.Run("CertifyModel_ShortCdCertificateId_Fails", func(t *testing.T) {
		txResult, err := CertifyModel(CertifyModelOpts{
			VID: vid, PID: pid, SoftwareVersion: sv, SoftwareVersionString: svs,
			CertificationType: certType, CertificationDate: date,
			CDCertificateID: "1234567890abcdefgh", // 18 chars
			From:            zbAccount,
		})
		cliputils.RequireTxFailContains(t, txResult, err, "minimum length for CDCertificateId allowed is 19")
	})

	t.Run("CertifyModel_LongCdCertificateId_Fails", func(t *testing.T) {
		txResult, err := CertifyModel(CertifyModelOpts{
			VID: vid, PID: pid, SoftwareVersion: sv, SoftwareVersionString: svs,
			CertificationType: certType, CertificationDate: date,
			CDCertificateID: "12345678910abcdefghX", // 20 chars
			From:            zbAccount,
		})
		cliputils.RequireTxFailContains(t, txResult, err, "maximum length for CDCertificateId allowed is 19")
	})

	t.Run("ProvisionModel_SchemaVersion0_Fails", func(t *testing.T) {
		txResult, err := ProvisionModel(ProvisionModelOpts{
			VID: vid, PID: pid, SoftwareVersion: sv, SoftwareVersionString: svs,
			CertificationType: certType, ProvisionalDate: date,
			CDCertificateID: cdCertID, SchemaVersion: "0", From: zbAccount,
		})
		cliputils.RequireTxFailContains(t, txResult, err, "SchemaVersion must be equal 1")
	})

	t.Run("RevokeModel_SchemaVersion0_Fails", func(t *testing.T) {
		txResult, err := RevokeModel(RevokeModelOpts{
			VID: vid, PID: pid, SoftwareVersion: sv, SoftwareVersionString: svs,
			CertificationType: certType, RevocationDate: date,
			SchemaVersion: "0", From: zbAccount,
		})
		cliputils.RequireTxFailContains(t, txResult, err, "SchemaVersion must be equal 1")
	})

	t.Run("CertifyModel_LongCertificationType_Fails", func(t *testing.T) {
		txResult, err := CertifyModel(CertifyModelOpts{
			VID: vid, PID: pid, SoftwareVersion: sv, SoftwareVersionString: svs,
			CertificationType: "this_certification_type_is_way_too_long",
			CertificationDate: date, CDCertificateID: cdCertID, From: zbAccount,
		})
		cliputils.RequireTxFailContains(t, txResult, err, "maximum length for CertificationType allowed is 20")
	})

	t.Run("CertifyModel_SpecificationVersionZero_Fails", func(t *testing.T) {
		// SchemaVersion omitted (CLI default 1) so the schema-v1 required-field
		// guard applies; SpecVersionZero forces --specificationVersion 0.
		txResult, err := CertifyModel(CertifyModelOpts{
			VID: vid, PID: pid, SoftwareVersion: sv, SoftwareVersionString: svs,
			CertificationType: certType, CertificationDate: date,
			CDCertificateID: cdCertID, SpecificationVersion: SpecVersionZero, From: zbAccount,
		})
		cliputils.RequireTxFailContains(t, txResult, err, "SpecificationVersion is a required field")
	})
}
