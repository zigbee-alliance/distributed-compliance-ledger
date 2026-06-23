package model

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/cli/compliance"
	cliputils "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/cli/utils"
	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/utils"
)

func TestModelVersionDemo(t *testing.T) {
	vid := rand.Intn(65534) + 1
	pid := rand.Intn(65534) + 1
	vendorAccount := fmt.Sprintf("vendor_account_%d", vid)
	cliputils.CreateVendorAccount(t, vendorAccount, vid)

	zbAccount := cliputils.CreateAccount(t, "CertificationCenter")

	t.Run("AddModel", func(t *testing.T) {
		txResult, err := AddModel(AddModelOpts{
			VID:          vid,
			PID:          pid,
			ProductLabel: "Test Product",
			From:         vendorAccount,
		})
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)
	})

	sv := rand.Intn(65534) + 1

	t.Run("CertifyModel_BeforeVersion", func(t *testing.T) {
		// Certify unknown SV — allowed
		txResult, err := compliance.CertifyModel(compliance.CertifyModelOpts{
			VID: vid, PID: pid,
			SoftwareVersion:       sv,
			SoftwareVersionString: "1",
			CertificationType:     "zigbee",
			CertificationDate:     "2020-01-01T00:00:01Z",
			CDCertificateID:       "0000000000000000001",
			From:                  zbAccount,
		})
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)
	})

	t.Run("AddModelVersion", func(t *testing.T) {
		txResult, err := AddModelVersion(AddModelVersionOpts{
			VID:                   vid,
			PID:                   pid,
			SoftwareVersion:       sv,
			SoftwareVersionString: "1",
			SchemaVersion:         "0",
			From:                  vendorAccount,
		})
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)
	})

	t.Run("QueryModelVersion", func(t *testing.T) {
		mv, err := GetModelVersion(vid, pid, sv)
		require.NoError(t, err)
		require.NotNil(t, mv)
		require.Equal(t, int32(vid), mv.Vid)
		require.Equal(t, int32(pid), mv.Pid)
		require.Equal(t, uint32(sv), mv.SoftwareVersion)
		require.Equal(t, "1", mv.SoftwareVersionString)
		require.Equal(t, int32(1), mv.CdVersionNumber)
		require.True(t, mv.SoftwareVersionValid)
		require.Equal(t, uint32(1), mv.MinApplicableSoftwareVersion)
		require.Equal(t, uint32(10), mv.MaxApplicableSoftwareVersion)
		require.Equal(t, uint32(0), mv.SchemaVersion)
	})

	t.Run("QueryAllModelVersions", func(t *testing.T) {
		mvs, err := GetAllModelVersions(vid, pid)
		require.NoError(t, err)
		require.NotNil(t, mvs)
		require.Equal(t, int32(vid), mvs.Vid)
		require.Equal(t, int32(pid), mvs.Pid)
		require.Contains(t, mvs.SoftwareVersions, uint32(sv))
	})

	t.Run("QueryNonExistentModelVersion", func(t *testing.T) {
		mv, err := GetModelVersion(vid, pid, 123456)
		require.NoError(t, err)
		require.Nil(t, mv)

		vid1 := rand.Intn(65534) + 1
		pid1 := rand.Intn(65534) + 1
		mvs, err := GetAllModelVersions(vid1, pid1)
		require.NoError(t, err)
		require.Nil(t, mvs)
	})

	t.Run("UpdateModelVersion", func(t *testing.T) {
		txResult, err := UpdateModelVersion(UpdateModelVersionOpts{
			VID: vid, PID: pid, SoftwareVersion: sv, From: vendorAccount,
			MinApplicableSoftwareVersion: 2,
			MaxApplicableSoftwareVersion: 10,
			SoftwareVersionValid:         boolPtr(false),
			SchemaVersion:                "0",
		})
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		mv, err := GetModelVersion(vid, pid, sv)
		require.NoError(t, err)
		require.NotNil(t, mv)
		require.False(t, mv.SoftwareVersionValid)
		require.Equal(t, uint32(2), mv.MinApplicableSoftwareVersion)
		require.Equal(t, uint32(10), mv.MaxApplicableSoftwareVersion)
	})

	sv2 := rand.Intn(65534) + 1

	t.Run("AddSecondModelVersion", func(t *testing.T) {
		txResult, err := AddModelVersion(AddModelVersionOpts{
			VID:                   vid,
			PID:                   pid,
			SoftwareVersion:       sv2,
			SoftwareVersionString: "1",
			From:                  vendorAccount,
		})
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		mvs, err := GetAllModelVersions(vid, pid)
		require.NoError(t, err)
		require.NotNil(t, mvs)
		require.Contains(t, mvs.SoftwareVersions, uint32(sv))
		require.Contains(t, mvs.SoftwareVersions, uint32(sv2))
	})

	t.Run("AddModelVersionFromDifferentVendor_Fails", func(t *testing.T) {
		newVid := rand.Intn(65534) + 1
		differentVendor := fmt.Sprintf("vendor_account_%d", newVid)
		cliputils.CreateVendorAccount(t, differentVendor, newVid)

		txResult, err := AddModelVersion(AddModelVersionOpts{
			VID:                          vid,
			PID:                          pid,
			SoftwareVersion:              sv,
			SoftwareVersionString:        "1",
			CDVersionNumber:              1,
			MinApplicableSoftwareVersion: 1,
			MaxApplicableSoftwareVersion: 10,
			From:                         differentVendor,
		})
		require.NoError(t, err)
		require.Contains(t, txResult.RawLog, fmt.Sprintf("vendorID %d", vid))
	})

	t.Run("UpdateModelVersionFromDifferentVendor_Fails", func(t *testing.T) {
		newVid := rand.Intn(65534) + 1
		differentVendor := fmt.Sprintf("vendor_account_%d", newVid)
		cliputils.CreateVendorAccount(t, differentVendor, newVid)

		txResult, err := UpdateModelVersion(UpdateModelVersionOpts{
			VID: vid, PID: pid, SoftwareVersion: sv, From: differentVendor,
			SoftwareVersionValid: boolPtr(false),
		})
		require.NoError(t, err)
		require.Contains(t, txResult.RawLog, fmt.Sprintf("vendorID %d", vid))
	})

	t.Run("DeleteModelVersion", func(t *testing.T) {
		// Delete compliance info first
		txResult, err := compliance.DeleteComplianceInfo(vid, pid, sv, "zigbee", zbAccount)
		require.NoError(t, err)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		txResult, err = DeleteModelVersion(vid, pid, sv, vendorAccount)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		mv, err := GetModelVersion(vid, pid, sv)
		require.NoError(t, err)
		require.Nil(t, mv)
	})

	t.Run("VendorAdminModelVersionLifecycle", func(t *testing.T) {
		// A VendorAdmin account can add/update/delete model versions for any
		// vendor (modelversion-demo.sh:185-223).
		vendorAdmin := cliputils.CreateAccount(t, "VendorAdmin")
		vid3 := rand.Intn(65534) + 1
		pid3 := rand.Intn(65534) + 1
		sv3 := rand.Intn(65534) + 1

		txResult, err := AddModel(AddModelOpts{VID: vid3, PID: pid3, ProductLabel: "Test Product", From: vendorAdmin})
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		txResult, err = AddModelVersion(AddModelVersionOpts{
			VID: vid3, PID: pid3, SoftwareVersion: sv3, SoftwareVersionString: "1", From: vendorAdmin,
		})
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		mv, err := GetModelVersion(vid3, pid3, sv3)
		require.NoError(t, err)
		require.NotNil(t, mv)
		require.Equal(t, int32(vid3), mv.Vid)
		require.Equal(t, uint32(sv3), mv.SoftwareVersion)
		require.True(t, mv.SoftwareVersionValid)

		// VendorAdmin invalidates the version.
		txResult, err = UpdateModelVersion(UpdateModelVersionOpts{
			VID: vid3, PID: pid3, SoftwareVersion: sv3, From: vendorAdmin,
			SoftwareVersionValid: boolPtr(false),
		})
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		mv, err = GetModelVersion(vid3, pid3, sv3)
		require.NoError(t, err)
		require.NotNil(t, mv)
		require.False(t, mv.SoftwareVersionValid)

		// VendorAdmin deletes the version (never certified, so no compliance info).
		txResult, err = DeleteModelVersion(vid3, pid3, sv3, vendorAdmin)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		mv, err = GetModelVersion(vid3, pid3, sv3)
		require.NoError(t, err)
		require.Nil(t, mv)
	})
}
