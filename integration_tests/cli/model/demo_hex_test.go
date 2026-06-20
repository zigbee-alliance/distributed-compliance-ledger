package model

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
	cliputils "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/cli/utils"
	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/utils"
)

func TestModelDemoHex(t *testing.T) {
	vid := rand.Intn(65534) + 1
	pid := rand.Intn(65534) + 1
	vidHex := fmt.Sprintf("0x%X", vid)
	pidHex := fmt.Sprintf("0x%X", pid)

	vendorAccount := fmt.Sprintf("vendor_account_%d", vid)
	cliputils.CreateVendorAccount(t, vendorAccount, vid)

	t.Run("QueryNonExistent", func(t *testing.T) {
		m, err := GetModelHex(vidHex, pidHex)
		require.NoError(t, err)
		require.Nil(t, m)

		vm, err := GetVendorModelsHex(vidHex)
		require.NoError(t, err)
		require.Nil(t, vm)

		all, err := GetAllModels()
		require.NoError(t, err)
		require.False(t, containsModelByPid(all, int32(vid), int32(pid)))
	})

	productLabel := "Device #1"

	t.Run("AddModel", func(t *testing.T) {
		txResult, err := AddModel(AddModelOpts{
			VIDHex: vidHex, PIDHex: pidHex,
			ProductLabel: productLabel,
			From:         vendorAccount,
		})
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)
	})

	t.Run("QueryModel", func(t *testing.T) {
		m, err := GetModelHex(vidHex, pidHex)
		require.NoError(t, err)
		require.NotNil(t, m)
		require.Equal(t, int32(vid), m.Vid)
		require.Equal(t, int32(pid), m.Pid)
		require.Equal(t, productLabel, m.ProductLabel)

		all, err := GetAllModels()
		require.NoError(t, err)
		require.True(t, containsModelByPid(all, int32(vid), int32(pid)))

		vm, err := GetVendorModelsHex(vidHex)
		require.NoError(t, err)
		require.NotNil(t, vm)
		require.True(t, containsProductByPid(vm.Products, int32(pid)))
	})

	description := "New Device Description"

	t.Run("UpdateModel", func(t *testing.T) {
		txResult, err := UpdateModelHex(vidHex, pidHex, vendorAccount,
			"--enhancedSetupFlowOptions", "2",
			"--productLabel", description,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		m, err := GetModelHex(vidHex, pidHex)
		require.NoError(t, err)
		require.NotNil(t, m)
		require.Equal(t, int32(vid), m.Vid)
		require.Equal(t, int32(pid), m.Pid)
		require.Equal(t, description, m.ProductLabel)
	})

	supportURL := "https://newsupporturl.test"

	t.Run("UpdateModelSupportURL", func(t *testing.T) {
		txResult, err := UpdateModelHex(vidHex, pidHex, vendorAccount,
			"--enhancedSetupFlowOptions", "2",
			"--supportURL", supportURL,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		m, err := GetModelHex(vidHex, pidHex)
		require.NoError(t, err)
		require.NotNil(t, m)
		require.Equal(t, supportURL, m.SupportUrl)
	})

	t.Run("DeleteModel", func(t *testing.T) {
		txResult, err := DeleteModelHex(vidHex, pidHex, vendorAccount)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		m, err := GetModelHex(vidHex, pidHex)
		require.NoError(t, err)
		require.Nil(t, m)
	})
}
