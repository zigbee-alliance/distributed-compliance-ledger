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
		txResult, err := UpdateModel(vid, pid, vendorAccount,
			"--productLabel", description,
			"--schemaVersion", "0",
			"--commissioningModeInitialStepsHint", fmt.Sprintf("%d", newCommissioningModeInitialStepsHint),
			"--commissioningModeSecondaryStepsHint", fmt.Sprintf("%d", newCommissioningModeSecondaryStepsHint),
			"--icdUserActiveModeTriggerHint", fmt.Sprintf("%d", newIcdUserActiveModeTriggerHint),
			"--enhancedSetupFlowOptions", "2",
			"--factoryResetStepsHint", fmt.Sprintf("%d", newFactoryResetStepsHint),
		)
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
		txResult, err := UpdateModel(vid, pid, vendorAccount,
			"--supportURL", supportURL,
			"--enhancedSetupFlowOptions", "0",
		)
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
}
