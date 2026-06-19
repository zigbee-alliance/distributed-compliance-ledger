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
		txResult, err := UpdateModelVersion(vid, pid, sv, vendorAccount,
			"--minApplicableSoftwareVersion", "2",
			"--maxApplicableSoftwareVersion", "10",
			"--softwareVersionValid=false",
			"--schemaVersion", "0",
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

		out, err := QueryAllModelVersions(vid, pid)
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf("%d", sv))
		require.Contains(t, string(out), fmt.Sprintf("%d", sv2))
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

		txResult, err := UpdateModelVersion(vid, pid, sv, differentVendor,
			"--softwareVersionValid=false",
		)
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

		out, err := QueryModelVersion(vid, pid, sv)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")
	})
}
