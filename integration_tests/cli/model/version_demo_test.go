package model

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
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
		txResult, err := utils.ExecuteTx("tx", "model", "add-model",
			"--vid", fmt.Sprintf("%d", vid),
			"--pid", fmt.Sprintf("%d", pid),
			"--deviceTypeID", "1",
			"--productName", "TestProduct",
			"--productLabel", "Test Product",
			"--partNumber", "1",
			"--commissioningCustomFlow", "0",
			"--enhancedSetupFlowOptions", "0",
			"--from", vendorAccount,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)
	})

	sv := rand.Intn(65534) + 1

	t.Run("CertifyModel_BeforeVersion", func(t *testing.T) {
		// Certify unknown SV — allowed
		txResult, err := utils.ExecuteTx("tx", "compliance", "certify-model",
			"--vid", fmt.Sprintf("%d", vid),
			"--pid", fmt.Sprintf("%d", pid),
			"--softwareVersion", fmt.Sprintf("%d", sv),
			"--cdVersionNumber", "1",
			"--softwareVersionString", "1",
			"--cdCertificateId", "0000000000000000001",
			"--certificationType", "zigbee",
			"--certificationDate", "2020-01-01T00:00:01Z",
			"--from", zbAccount,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)
	})

	t.Run("AddModelVersion", func(t *testing.T) {
		txResult, err := utils.ExecuteTx("tx", "model", "add-model-version",
			"--cdVersionNumber", "1",
			"--maxApplicableSoftwareVersion", "10",
			"--minApplicableSoftwareVersion", "1",
			"--vid", fmt.Sprintf("%d", vid),
			"--pid", fmt.Sprintf("%d", pid),
			"--softwareVersion", fmt.Sprintf("%d", sv),
			"--softwareVersionString", "1",
			"--schemaVersion", "0",
			"--from", vendorAccount,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)
	})

	t.Run("QueryModelVersion", func(t *testing.T) {
		out, err := QueryModelVersion(vid, pid, sv)
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"vid":%d`, vid))
		require.Contains(t, string(out), fmt.Sprintf(`"pid":%d`, pid))
		require.Contains(t, string(out), fmt.Sprintf(`"softwareVersion":%d`, sv))
		require.Contains(t, string(out), `"softwareVersionString":"1"`)
		require.Contains(t, string(out), `"cdVersionNumber":1`)
		require.Contains(t, string(out), `"softwareVersionValid":true`)
		require.Contains(t, string(out), `"minApplicableSoftwareVersion":1`)
		require.Contains(t, string(out), `"maxApplicableSoftwareVersion":10`)
		require.Contains(t, string(out), `"schemaVersion":0`)
	})

	t.Run("QueryAllModelVersions", func(t *testing.T) {
		out, err := QueryAllModelVersions(vid, pid)
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"vid":%d`, vid))
		require.Contains(t, string(out), fmt.Sprintf(`"pid":%d`, pid))
		require.Contains(t, string(out), "softwareVersions")
		require.Contains(t, string(out), fmt.Sprintf("%d", sv))
	})

	t.Run("QueryNonExistentModelVersion", func(t *testing.T) {
		out, err := QueryModelVersion(vid, pid, 123456)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")

		vid1 := rand.Intn(65534) + 1
		pid1 := rand.Intn(65534) + 1
		out, err = QueryAllModelVersions(vid1, pid1)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")
	})

	t.Run("UpdateModelVersion", func(t *testing.T) {
		txResult, err := utils.ExecuteTx("tx", "model", "update-model-version",
			"--vid", fmt.Sprintf("%d", vid),
			"--pid", fmt.Sprintf("%d", pid),
			"--softwareVersion", fmt.Sprintf("%d", sv),
			"--minApplicableSoftwareVersion", "2",
			"--maxApplicableSoftwareVersion", "10",
			"--softwareVersionValid=false",
			"--schemaVersion", "0",
			"--from", vendorAccount,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		out, err := QueryModelVersion(vid, pid, sv)
		require.NoError(t, err)
		require.Contains(t, string(out), `"softwareVersionValid":false`)
		require.Contains(t, string(out), `"minApplicableSoftwareVersion":2`)
		require.Contains(t, string(out), `"maxApplicableSoftwareVersion":10`)
	})

	sv2 := rand.Intn(65534) + 1

	t.Run("AddSecondModelVersion", func(t *testing.T) {
		txResult, err := utils.ExecuteTx("tx", "model", "add-model-version",
			"--cdVersionNumber", "1",
			"--maxApplicableSoftwareVersion", "10",
			"--minApplicableSoftwareVersion", "1",
			"--vid", fmt.Sprintf("%d", vid),
			"--pid", fmt.Sprintf("%d", pid),
			"--softwareVersion", fmt.Sprintf("%d", sv2),
			"--softwareVersionString", "1",
			"--from", vendorAccount,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		out, err := QueryAllModelVersions(vid, pid)
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf("%d", sv))
		require.Contains(t, string(out), fmt.Sprintf("%d", sv2))
	})

	t.Run("AddModelVersionFromDifferentVendor_Fails", func(t *testing.T) {
		newVid := rand.Intn(65534) + 1
		differentVendor := fmt.Sprintf("vendor_account_%d", newVid)
		cliputils.CreateVendorAccount(t, differentVendor, newVid)

		txResult, err := utils.ExecuteTx("tx", "model", "add-model-version",
			"--cdVersionNumber", "1",
			"--maxApplicableSoftwareVersion", "10",
			"--minApplicableSoftwareVersion", "1",
			"--vid", fmt.Sprintf("%d", vid),
			"--pid", fmt.Sprintf("%d", pid),
			"--softwareVersion", fmt.Sprintf("%d", sv),
			"--softwareVersionString", "1",
			"--from", differentVendor,
		)
		require.NoError(t, err)
		require.Contains(t, txResult.RawLog, fmt.Sprintf("vendorID %d", vid))
	})

	t.Run("UpdateModelVersionFromDifferentVendor_Fails", func(t *testing.T) {
		newVid := rand.Intn(65534) + 1
		differentVendor := fmt.Sprintf("vendor_account_%d", newVid)
		cliputils.CreateVendorAccount(t, differentVendor, newVid)

		txResult, err := utils.ExecuteTx("tx", "model", "update-model-version",
			"--vid", fmt.Sprintf("%d", vid),
			"--pid", fmt.Sprintf("%d", pid),
			"--softwareVersion", fmt.Sprintf("%d", sv),
			"--softwareVersionValid=false",
			"--from", differentVendor,
		)
		require.NoError(t, err)
		require.Contains(t, txResult.RawLog, fmt.Sprintf("vendorID %d", vid))
	})

	t.Run("DeleteModelVersion", func(t *testing.T) {
		// Delete compliance info first
		txResult, err := utils.ExecuteTx("tx", "compliance", "delete-compliance-info",
			"--vid", fmt.Sprintf("%d", vid),
			"--pid", fmt.Sprintf("%d", pid),
			"--softwareVersion", fmt.Sprintf("%d", sv),
			"--certificationType", "zigbee",
			"--from", zbAccount,
		)
		require.NoError(t, err)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		txResult, err = utils.ExecuteTx("tx", "model", "delete-model-version",
			"--vid", fmt.Sprintf("%d", vid),
			"--pid", fmt.Sprintf("%d", pid),
			"--softwareVersion", fmt.Sprintf("%d", sv),
			"--from", vendorAccount,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		out, err := QueryModelVersion(vid, pid, sv)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")
	})
}
