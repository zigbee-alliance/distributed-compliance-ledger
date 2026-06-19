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
		out, err := QueryModelVersionHex(vidHex, pidHex, sv)
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"vid":%d`, vid))
		require.Contains(t, string(out), fmt.Sprintf(`"pid":%d`, pid))
		require.Contains(t, string(out), fmt.Sprintf(`"softwareVersion":%d`, sv))
		require.Contains(t, string(out), `"softwareVersionString":"1"`)
		require.Contains(t, string(out), `"cdVersionNumber":1`)
		require.Contains(t, string(out), `"softwareVersionValid":true`)
		require.Contains(t, string(out), `"minApplicableSoftwareVersion":1`)
		require.Contains(t, string(out), `"maxApplicableSoftwareVersion":10`)
	})

	t.Run("QueryAllModelVersions_WithHexVID", func(t *testing.T) {
		out, err := QueryAllModelVersionsHex(vidHex, pidHex)
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"vid":%d`, vid))
		require.Contains(t, string(out), fmt.Sprintf(`"pid":%d`, pid))
		require.Contains(t, string(out), "softwareVersions")
		require.Contains(t, string(out), fmt.Sprintf("%d", sv))
	})

	t.Run("QueryNonExistentModelVersion_WithHexVID", func(t *testing.T) {
		out, err := QueryModelVersionHex(vidHex, pidHex, 123456)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")
	})
}
