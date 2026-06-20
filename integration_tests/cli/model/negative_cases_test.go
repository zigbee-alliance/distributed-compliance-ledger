package model

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/cli/compliance"
	cliputils "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/cli/utils"
	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/utils"
)

func TestModelNegativeCases(t *testing.T) {
	certificationHouse := cliputils.CreateAccount(t, "CertificationCenter")

	vid := rand.Intn(65534) + 1
	pid := rand.Intn(65534) + 1
	softwareVersionString := fmt.Sprintf("%d", rand.Intn(65534)+1)
	vendorAccount := fmt.Sprintf("vendor_account_%d", vid)
	cliputils.CreateVendorAccount(t, vendorAccount, vid)

	zbAccount := cliputils.CreateAccount(t, "CertificationCenter")

	vidWithPids := vid + 1
	pidRanges := "1-100"
	vendorAccountWithPids := fmt.Sprintf("vendor_account_%d", vidWithPids)
	cliputils.CreateVendorAccount(t, vendorAccountWithPids, vidWithPids, pidRanges)

	t.Run("AddModel_NotVendor_Fails", func(t *testing.T) {
		txResult, err := AddModel(AddModelOpts{VID: vid, PID: pid, From: certificationHouse})
		require.NoError(t, err)
		require.Equal(t, uint32(4), txResult.Code)
		_, _ = utils.AwaitTxConfirmation(txResult.TxHash)
	})

	t.Run("AddModel_VendorNonAssociatedPID_Fails", func(t *testing.T) {
		txResult, err := AddModel(AddModelOpts{VID: vidWithPids, PID: 101, From: vendorAccountWithPids})
		require.NoError(t, err)
		require.Equal(t, uint32(4), txResult.Code)
		_, _ = utils.AwaitTxConfirmation(txResult.TxHash)
	})

	t.Run("AddModel_WrongVendorID_Fails", func(t *testing.T) {
		vid1 := rand.Intn(65534) + 1
		txResult, err := AddModel(AddModelOpts{VID: vid1, PID: pid, From: vendorAccount})
		require.NoError(t, err)
		require.Equal(t, uint32(4), txResult.Code)
		_, _ = utils.AwaitTxConfirmation(txResult.TxHash)
	})

	t.Run("AddModelTwice_Fails", func(t *testing.T) {
		// First add succeeds
		txResult, err := AddModel(AddModelOpts{
			VID:  vid,
			PID:  pid,
			From: vendorAccount,
		})
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Second add fails with code 501 (model already exists).
		txResult, err = AddModel(AddModelOpts{VID: vid, PID: pid, From: vendorAccount})
		require.NoError(t, err)
		require.Equal(t, uint32(501), txResult.Code)
		_, _ = utils.AwaitTxConfirmation(txResult.TxHash)
	})

	sv := rand.Intn(65534) + 1

	t.Run("AddModelVersion_ThenCertify_ThenDeleteCertifiedModel_Fails", func(t *testing.T) {
		txResult, err := AddModelVersion(AddModelVersionOpts{
			VID:                   vid,
			PID:                   pid,
			SoftwareVersion:       sv,
			SoftwareVersionString: softwareVersionString,
			From:                  vendorAccount,
		})
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		certificationDate := "2020-01-01T00:00:01Z"
		txResult, err = compliance.CertifyModel(compliance.CertifyModelOpts{
			VID: vid, PID: pid,
			SoftwareVersion:       sv,
			SoftwareVersionString: softwareVersionString,
			CertificationType:     "zigbee",
			CertificationDate:     certificationDate,
			CDCertificateID:       "1230000000000000000",
			From:                  zbAccount,
		})
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Delete certified model — should fail with code 525.
		txResult, err = DeleteModel(vid, pid, vendorAccount)
		require.NoError(t, err)
		require.Equal(t, uint32(525), txResult.Code)
		_, _ = utils.AwaitTxConfirmation(txResult.TxHash)
	})

	t.Run("AddModel_UnknownAccount_Fails", func(t *testing.T) {
		// AddModel routes through ExecuteTx; an unknown --from is rejected at the
		// CLI keyring layer before broadcast, so the failure surfaces as a Go err.
		_, err := AddModel(AddModelOpts{VID: vid, PID: pid, From: "Unknown"})
		require.Error(t, err)
		require.Contains(t, err.Error(), "key not found")
	})

	type vidCase struct {
		label string
		opts  AddModelOpts
	}
	t.Run("AddModel_InvalidVidPid", func(t *testing.T) {
		cases := []vidCase{
			{"-1", AddModelOpts{VID: -1, PID: pid, From: vendorAccount}},
			{"0", AddModelOpts{VID: 0, PID: pid, From: vendorAccount}},
			{"65536", AddModelOpts{VID: 65536, PID: pid, From: vendorAccount}},
			// VIDHex bypasses int formatting and lets us send a non-numeric token.
			{"string", AddModelOpts{VIDHex: "string", PID: pid, From: vendorAccount}},
		}
		for _, tc := range cases {
			tc := tc
			txResult, err := AddModel(tc.opts)
			combined := ""
			if err != nil {
				combined = err.Error()
			}
			if txResult != nil {
				combined += txResult.RawLog
			}
			hasErr := combined != "" && (strings.Contains(combined, "Vid must not be") ||
				strings.Contains(combined, "invalid syntax") ||
				strings.Contains(combined, "invalid argument"))
			require.True(t, hasErr, "expected error for vid=%s, got: %s", tc.label, combined)
		}
	})

	t.Run("AddModel_EmptyProductName_Fails", func(t *testing.T) {
		// AddModel substitutes its "TestProduct" default for an empty ProductName,
		// so we can't drive this case through the typed helper. Send the raw flags.
		out, err := utils.ExecuteCLI("tx", "model", "add-model",
			"--vid", fmt.Sprintf("%d", vid),
			"--pid", fmt.Sprintf("%d", pid),
			"--deviceTypeID", "1",
			"--productName", "",
			"--productLabel", "TestingProductLabel",
			"--partNumber", "1",
			"--commissioningCustomFlow", "0",
			"--enhancedSetupFlowOptions", "0",
			"--from", vendorAccount,
			"--yes", "-o", "json", "--keyring-backend", "test",
		)
		combined := string(out)
		if err != nil {
			combined += err.Error()
		}
		require.Contains(t, combined, "ProductName is a required field")
	})

	t.Run("AddModel_EmptyFrom_Fails", func(t *testing.T) {
		// AddModel forwards From="" verbatim; the CLI rejects it before broadcast.
		_, err := AddModel(AddModelOpts{VID: vid, PID: pid, From: ""})
		require.Error(t, err)
		require.Contains(t, err.Error(), "invalid creator address")
	})
}
