package model

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
	cliputils "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/cli/utils"
	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/utils"
)

func TestModelVersionDemoHex(t *testing.T) {
	vidHex := "0xA13"
	pidHex := "0xA11"
	vid := 2579
	pid := 2577

	vendorAccount := fmt.Sprintf("vendor_account_hex_%d", vid)
	cliputils.CreateVendorAccount(t, vendorAccount, vid)

	sv := rand.Intn(65534) + 1

	t.Run("AddModel", func(t *testing.T) {
		txResult, err := AddModel(AddModelOpts{
			VIDHex:       vidHex,
			PIDHex:       pidHex,
			ProductLabel: "Test Product",
			From:         vendorAccount,
		})
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)
	})

	t.Run("AddModelVersion_WithDecimalSV", func(t *testing.T) {
		txResult, err := AddModelVersion(AddModelVersionOpts{
			VID:                   vid,
			PID:                   pid,
			SoftwareVersion:       sv,
			SoftwareVersionString: "1",
			From:                  vendorAccount,
		})
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)
	})

	t.Run("QueryModelVersion_WithHexVID", func(t *testing.T) {
		mv, err := GetModelVersionHex(vidHex, pidHex, sv)
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
	})

	t.Run("QueryAllModelVersions_WithHexVID", func(t *testing.T) {
		mvs, err := GetAllModelVersionsHex(vidHex, pidHex)
		require.NoError(t, err)
		require.NotNil(t, mvs)
		require.Equal(t, int32(vid), mvs.Vid)
		require.Equal(t, int32(pid), mvs.Pid)
		require.Contains(t, mvs.SoftwareVersions, uint32(sv))
	})

	t.Run("QueryNonExistentModelVersion_WithHexVID", func(t *testing.T) {
		mv, err := GetModelVersionHex(vidHex, pidHex, 123456)
		require.NoError(t, err)
		require.Nil(t, mv)

		// all-model-versions for an unrelated hex vid/pid is also Not Found.
		mvs, err := GetAllModelVersionsHex("0xA14", "0xA15")
		require.NoError(t, err)
		require.Nil(t, mvs)
	})

	t.Run("UpdateModelVersion_WithHexVID", func(t *testing.T) {
		// The chain keys versions by the integer value, so a decimal update
		// targets the same hex-created version; read it back with the hex query.
		txResult, err := UpdateModelVersion(vid, pid, sv, vendorAccount,
			"--minApplicableSoftwareVersion", "2",
			"--maxApplicableSoftwareVersion", "10",
			"--softwareVersionValid=false",
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		mv, err := GetModelVersionHex(vidHex, pidHex, sv)
		require.NoError(t, err)
		require.NotNil(t, mv)
		require.False(t, mv.SoftwareVersionValid)
		require.Equal(t, uint32(2), mv.MinApplicableSoftwareVersion)
		require.Equal(t, uint32(10), mv.MaxApplicableSoftwareVersion)
	})

	sv2 := rand.Intn(65534) + 1

	t.Run("AddSecondModelVersion_WithHexVID", func(t *testing.T) {
		txResult, err := AddModelVersion(AddModelVersionOpts{
			VID: vid, PID: pid, SoftwareVersion: sv2, SoftwareVersionString: "1", From: vendorAccount,
		})
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		mvs, err := GetAllModelVersionsHex(vidHex, pidHex)
		require.NoError(t, err)
		require.NotNil(t, mvs)
		require.Contains(t, mvs.SoftwareVersions, uint32(sv))
		require.Contains(t, mvs.SoftwareVersions, uint32(sv2))
	})

	t.Run("AddAndUpdateModelVersionFromDifferentVendor_Fails", func(t *testing.T) {
		newVid := rand.Intn(60000) + 3000 // guaranteed != vid (2579)
		differentVendor := fmt.Sprintf("vendor_account_%d", newVid)
		cliputils.CreateVendorAccount(t, differentVendor, newVid)

		txResult, err := AddModelVersion(AddModelVersionOpts{
			VID: vid, PID: pid, SoftwareVersion: sv, SoftwareVersionString: "1",
			CDVersionNumber: 1, MinApplicableSoftwareVersion: 1, MaxApplicableSoftwareVersion: 10,
			From: differentVendor,
		})
		require.NoError(t, err)
		require.Contains(t, txResult.RawLog, fmt.Sprintf("vendorID %d", vid))

		txResult, err = UpdateModelVersion(vid, pid, sv, differentVendor, "--softwareVersionValid=false")
		require.NoError(t, err)
		require.Contains(t, txResult.RawLog, fmt.Sprintf("vendorID %d", vid))
	})
}
