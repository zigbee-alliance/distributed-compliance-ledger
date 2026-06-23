package model

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
	cliputils "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/cli/utils"
	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/utils"
)

func TestModelDemo(t *testing.T) {
	vid := rand.Intn(65534) + 1
	vendorAccount := fmt.Sprintf("vendor_account_%d", vid)
	cliputils.CreateVendorAccount(t, vendorAccount, vid)

	vidWithPids := vid + 1
	pid := rand.Intn(65534) + 1
	pidRanges := fmt.Sprintf("%d-%d", pid, pid)
	vendorAccountWithPids := fmt.Sprintf("vendor_account_%d", vidWithPids)
	cliputils.CreateVendorAccount(t, vendorAccountWithPids, vidWithPids, pidRanges)

	t.Run("QueryNonExistentModel", func(t *testing.T) {
		m, err := GetModel(vid, pid)
		require.NoError(t, err)
		require.Nil(t, m)

		vm, err := GetVendorModels(vid)
		require.NoError(t, err)
		require.Nil(t, vm)

		all, err := GetAllModels()
		require.NoError(t, err)
		require.False(t, containsModelByPid(all, int32(vid), int32(pid)))
	})

	productLabel := "Device #1"
	sv := 1
	cdVersionNum := 10

	t.Run("AddModel", func(t *testing.T) {
		txResult, err := AddModel(AddModelOpts{
			VID: vid, PID: pid,
			ProductLabel:  productLabel,
			SchemaVersion: "0",
			From:          vendorAccount,
		})
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)
	})

	t.Run("AddModelWithPidRanges", func(t *testing.T) {
		txResult, err := AddModel(AddModelOpts{
			VID: vidWithPids, PID: pid,
			ProductLabel:                 productLabel,
			EnhancedSetupFlowOptions:     1,
			EnhancedSetupFlowTCUrl:       "https://example.org/file.txt",
			EnhancedSetupFlowTCRevision:  1,
			EnhancedSetupFlowTCDigest:    "MWRjNGE0NDA0MWRjYWYxMTU0NWI3NTQzZGZlOTQyZjQ3NDJmNTY4YmU2OGZlZTI3NTQ0MWIwOTJiYjYwZGVlZA==",
			EnhancedSetupFlowTCFileSize:  1024,
			MaintenanceURL:               "https://example.org",
			CommissioningFallbackURL:     "https://url.commissioningfallbackurl.dclmodel",
			DiscoveryCapabilitiesBitmask: 1,
			From:                         vendorAccountWithPids,
		})
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)
	})

	t.Run("QueryModel", func(t *testing.T) {
		m, err := GetModel(vid, pid)
		require.NoError(t, err)
		require.NotNil(t, m)
		require.Equal(t, int32(vid), m.Vid)
		require.Equal(t, int32(pid), m.Pid)
		require.Equal(t, productLabel, m.ProductLabel)
		require.Equal(t, uint32(0), m.SchemaVersion)
		require.Equal(t, int32(0), m.EnhancedSetupFlowOptions)
	})

	t.Run("QueryModelWithPidRanges", func(t *testing.T) {
		// The vidWithPids model was added with the full enhanced-setup-flow and
		// maintenance field set — verify each persisted.
		m, err := GetModel(vidWithPids, pid)
		require.NoError(t, err)
		require.NotNil(t, m)
		require.Equal(t, int32(vidWithPids), m.Vid)
		require.Equal(t, int32(pid), m.Pid)
		require.Equal(t, productLabel, m.ProductLabel)
		require.Equal(t, int32(1), m.EnhancedSetupFlowOptions)
		require.Equal(t, "https://example.org/file.txt", m.EnhancedSetupFlowTCUrl)
		require.Equal(t, int32(1), m.EnhancedSetupFlowTCRevision)
		require.Equal(t, "MWRjNGE0NDA0MWRjYWYxMTU0NWI3NTQzZGZlOTQyZjQ3NDJmNTY4YmU2OGZlZTI3NTQ0MWIwOTJiYjYwZGVlZA==", m.EnhancedSetupFlowTCDigest)
		require.Equal(t, uint32(1024), m.EnhancedSetupFlowTCFileSize)
		require.Equal(t, "https://example.org", m.MaintenanceUrl)
		require.Equal(t, "https://url.commissioningfallbackurl.dclmodel", m.CommissioningFallbackUrl)
		require.Equal(t, uint32(1), m.DiscoveryCapabilitiesBitmask)
	})

	t.Run("UpdateModelWithPidRangesFields", func(t *testing.T) {
		// Update the vidWithPids model's enhanced-setup-flow / maintenance fields
		// and confirm the new values persist.
		newTCUrl := "https://example.org/file2.txt"
		newDigest := "MWRjNGE0NDA0MWRjYWYxMTU0NWI3NTQzZGZlOTQyZjQ3NDJmNTY4YmU2OGZlZTI3NTQ0MWIwOTJiYjYwZGVlZA=="
		newMaintenanceURL := "https://example.org/maintenance"
		newFallbackURL := "https://url.newcommissioningfallbackurl.dclmodel"
		txResult, err := UpdateModel(UpdateModelOpts{
			VID: vidWithPids, PID: pid, From: vendorAccountWithPids,
			ProductLabel:                "Updated pid-ranges model",
			EnhancedSetupFlowOptions:    1,
			EnhancedSetupFlowTCUrl:      newTCUrl,
			EnhancedSetupFlowTCRevision: 2,
			EnhancedSetupFlowTCDigest:   newDigest,
			EnhancedSetupFlowTCFileSize: 2048,
			MaintenanceURL:              newMaintenanceURL,
			CommissioningFallbackURL:    newFallbackURL,
		})
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		m, err := GetModel(vidWithPids, pid)
		require.NoError(t, err)
		require.NotNil(t, m)
		require.Equal(t, int32(1), m.EnhancedSetupFlowOptions)
		require.Equal(t, newTCUrl, m.EnhancedSetupFlowTCUrl)
		require.Equal(t, int32(2), m.EnhancedSetupFlowTCRevision)
		require.Equal(t, uint32(2048), m.EnhancedSetupFlowTCFileSize)
		require.Equal(t, newMaintenanceURL, m.MaintenanceUrl)
		require.Equal(t, newFallbackURL, m.CommissioningFallbackUrl)
	})

	t.Run("AddModelVersions", func(t *testing.T) {
		svStr := fmt.Sprintf("%d", sv)
		txResult, err := AddModelVersion(AddModelVersionOpts{
			VID: vid, PID: pid,
			SoftwareVersion:              sv,
			SoftwareVersionString:        svStr,
			CDVersionNumber:              cdVersionNum,
			MaxApplicableSoftwareVersion: 15,
			From:                         vendorAccount,
		})
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		txResult, err = AddModelVersion(AddModelVersionOpts{
			VID: vidWithPids, PID: pid,
			SoftwareVersion:              sv,
			SoftwareVersionString:        svStr,
			CDVersionNumber:              cdVersionNum,
			MaxApplicableSoftwareVersion: 15,
			From:                         vendorAccountWithPids,
		})
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)
	})

	t.Run("QueryAllModelsAndVendorModels", func(t *testing.T) {
		all, err := GetAllModels()
		require.NoError(t, err)
		require.True(t, containsModelByPid(all, int32(vid), int32(pid)))

		vm, err := GetVendorModels(vid)
		require.NoError(t, err)
		require.NotNil(t, vm)
		require.True(t, containsProductByPid(vm.Products, int32(pid)))
	})

	description := "New Device Description"
	newCommissioningModeInitialStepsHint := 8
	newCommissioningModeSecondaryStepsHint := 9
	newIcdUserActiveModeTriggerHint := 7
	newFactoryResetStepsHint := 6

	t.Run("UpdateModel", func(t *testing.T) {
		txResult, err := UpdateModel(UpdateModelOpts{
			VID: vid, PID: pid, From: vendorAccount,
			ProductLabel:                        description,
			SchemaVersion:                       "0",
			CommissioningModeInitialStepsHint:   newCommissioningModeInitialStepsHint,
			CommissioningModeSecondaryStepsHint: newCommissioningModeSecondaryStepsHint,
			IcdUserActiveModeTriggerHint:        newIcdUserActiveModeTriggerHint,
			EnhancedSetupFlowOptions:            2,
			FactoryResetStepsHint:               newFactoryResetStepsHint,
		})
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)
	})

	t.Run("QueryUpdatedModel", func(t *testing.T) {
		m, err := GetModel(vid, pid)
		require.NoError(t, err)
		require.NotNil(t, m)
		require.Equal(t, int32(vid), m.Vid)
		require.Equal(t, int32(pid), m.Pid)
		require.Equal(t, description, m.ProductLabel)
		require.Equal(t, uint32(0), m.SchemaVersion)
		require.Equal(t, uint32(newCommissioningModeInitialStepsHint), m.CommissioningModeInitialStepsHint)
		require.Equal(t, uint32(newCommissioningModeSecondaryStepsHint), m.CommissioningModeSecondaryStepsHint)
		require.Equal(t, uint32(newIcdUserActiveModeTriggerHint), m.IcdUserActiveModeTriggerHint)
		require.Equal(t, uint32(newFactoryResetStepsHint), m.FactoryResetStepsHint)
		require.Equal(t, int32(2), m.EnhancedSetupFlowOptions)
	})

	supportURL := "https://newsupporturl.test"

	t.Run("UpdateModelSupportURL", func(t *testing.T) {
		txResult, err := UpdateModel(UpdateModelOpts{
			VID: vid, PID: pid, From: vendorAccount,
			SupportURL: supportURL,
		})
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		m, err := GetModel(vid, pid)
		require.NoError(t, err)
		require.NotNil(t, m)
		require.Equal(t, supportURL, m.SupportUrl)
	})

	t.Run("DeleteModels", func(t *testing.T) {
		txResult, err := DeleteModel(vid, pid, vendorAccount)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		txResult, err = DeleteModel(vidWithPids, pid, vendorAccountWithPids)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)
	})

	t.Run("QueryAfterDeletion", func(t *testing.T) {
		m, err := GetModel(vid, pid)
		require.NoError(t, err)
		require.Nil(t, m)

		m, err = GetModel(vidWithPids, pid)
		require.NoError(t, err)
		require.Nil(t, m)

		mv, err := GetModelVersion(vid, pid, sv)
		require.NoError(t, err)
		require.Nil(t, mv)

		mv, err = GetModelVersion(vidWithPids, pid, sv)
		require.NoError(t, err)
		require.Nil(t, mv)
	})

	t.Run("VendorAdminModelLifecycle", func(t *testing.T) {
		// A VendorAdmin account can add/update/delete a model for any vendor
		// (model-demo.sh:273-309).
		vendorAdmin := cliputils.CreateAccount(t, "VendorAdmin")
		vid3 := rand.Intn(65534) + 1
		pid3 := rand.Intn(65534) + 1

		txResult, err := AddModel(AddModelOpts{VID: vid3, PID: pid3, ProductLabel: "VendorAdmin Product", From: vendorAdmin})
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		m, err := GetModel(vid3, pid3)
		require.NoError(t, err)
		require.NotNil(t, m)
		require.Equal(t, "VendorAdmin Product", m.ProductLabel)

		txResult, err = UpdateModel(UpdateModelOpts{
			VID: vid3, PID: pid3, From: vendorAdmin,
			ProductLabel: "Updated by VendorAdmin",
		})
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		m, err = GetModel(vid3, pid3)
		require.NoError(t, err)
		require.NotNil(t, m)
		require.Equal(t, "Updated by VendorAdmin", m.ProductLabel)

		txResult, err = DeleteModel(vid3, pid3, vendorAdmin)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		m, err = GetModel(vid3, pid3)
		require.NoError(t, err)
		require.Nil(t, m)
	})
}
