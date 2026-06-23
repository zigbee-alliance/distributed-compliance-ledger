package model

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/cli/compliance"
	cliputils "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/cli/utils"
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
		cliputils.RequireTxFailCode(t, txResult, err, 4)
	})

	t.Run("AddModel_VendorNonAssociatedPID_Fails", func(t *testing.T) {
		txResult, err := AddModel(AddModelOpts{VID: vidWithPids, PID: 101, From: vendorAccountWithPids})
		cliputils.RequireTxFailCode(t, txResult, err, 4)
	})

	t.Run("AddModel_WrongVendorID_Fails", func(t *testing.T) {
		vid1 := rand.Intn(65534) + 1
		txResult, err := AddModel(AddModelOpts{VID: vid1, PID: pid, From: vendorAccount})
		cliputils.RequireTxFailCode(t, txResult, err, 4)
	})

	t.Run("AddModelTwice_Fails", func(t *testing.T) {
		// First add succeeds
		txResult, err := AddModel(AddModelOpts{
			VID:  vid,
			PID:  pid,
			From: vendorAccount,
		})
		cliputils.RequireTxOK(t, txResult, err)

		// Second add fails with code 501 (model already exists).
		txResult, err = AddModel(AddModelOpts{VID: vid, PID: pid, From: vendorAccount})
		cliputils.RequireTxFailCode(t, txResult, err, 501)
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
		cliputils.RequireTxOK(t, txResult, err)

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
		cliputils.RequireTxOK(t, txResult, err)

		// Delete certified model — should fail with code 525.
		txResult, err = DeleteModel(vid, pid, vendorAccount)
		cliputils.RequireTxFailCode(t, txResult, err, 525)
	})

	t.Run("AddModel_UnknownAccount_Fails", func(t *testing.T) {
		// AddModel routes through ExecuteTx; an unknown --from is rejected at the
		// CLI keyring layer before broadcast, so the failure surfaces as a Go err.
		_, err := AddModel(AddModelOpts{VID: vid, PID: pid, From: "Unknown"})
		require.Error(t, err)
		require.Contains(t, err.Error(), "key not found")
	})

	t.Run("AddModel_InvalidVidPid", func(t *testing.T) {
		// Cover both the VID and PID sides with the exact range/parse messages.
		// VIDHex/PIDHex send a non-numeric token to trigger the CLI parse error.
		cases := []struct {
			label string
			opts  AddModelOpts
			want  string
		}{
			{"vid<1", AddModelOpts{VID: -1, PID: pid, From: vendorAccount}, "Vid must not be less than 1"},
			{"vid>65535", AddModelOpts{VID: 65536, PID: pid, From: vendorAccount}, "Vid must not be greater than 65535"},
			{"vid-nonnumeric", AddModelOpts{VIDHex: "string", PID: pid, From: vendorAccount}, "invalid syntax"},
			{"pid<1", AddModelOpts{VID: vid, PID: -1, From: vendorAccount}, "Pid must not be less than 1"},
			{"pid>65535", AddModelOpts{VID: vid, PID: 65536, From: vendorAccount}, "Pid must not be greater than 65535"},
			{"pid-nonnumeric", AddModelOpts{VID: vid, PIDHex: "string", From: vendorAccount}, "invalid syntax"},
		}
		for _, tc := range cases {
			tc := tc
			txResult, err := AddModel(tc.opts)
			require.Contains(t, cliputils.TxFailureText(txResult, err), tc.want, "case %s", tc.label)
		}
	})

	t.Run("AddModel_EmptyProductName_Fails", func(t *testing.T) {
		txResult, err := AddModel(AddModelOpts{
			VID: vid, PID: pid, From: vendorAccount,
			ProductName: EmptyField,
		})
		cliputils.RequireTxFailContains(t, txResult, err, "ProductName is a required field")
	})

	t.Run("AddModel_EmptyFrom_Fails", func(t *testing.T) {
		// AddModel forwards From="" verbatim; the CLI rejects it before broadcast.
		_, err := AddModel(AddModelOpts{VID: vid, PID: pid, From: ""})
		require.Error(t, err)
		require.Contains(t, err.Error(), "invalid creator address")
	})

	t.Run("AddModel_EmptyProductLabel_Fails", func(t *testing.T) {
		txResult, err := AddModel(AddModelOpts{
			VID: vid, PID: pid, From: vendorAccount,
			ProductLabel: EmptyField,
		})
		cliputils.RequireTxFailContains(t, txResult, err, "ProductLabel is a required field")
	})

	t.Run("AddModel_EmptyPartNumber_Fails", func(t *testing.T) {
		txResult, err := AddModel(AddModelOpts{
			VID: vid, PID: pid, From: vendorAccount,
			PartNumber: EmptyField,
		})
		cliputils.RequireTxFailContains(t, txResult, err, "PartNumber is a required field")
	})

	t.Run("AddModel_DiscoveryBitmaskTooHigh_Fails", func(t *testing.T) {
		// 31 is one above the allowed max of 30. Reaches ValidateBasic.
		txResult, err := AddModel(AddModelOpts{
			VID: vid, PID: rand.Intn(65534) + 1, From: vendorAccount,
			DiscoveryCapabilitiesBitmask: 31,
		})
		cliputils.RequireTxFailContains(t, txResult, err, "DiscoveryCapabilitiesBitmask must not be greater than 30")
	})

	t.Run("AddModel_NoFromFlag_Fails", func(t *testing.T) {
		// Omit --from entirely (distinct from the empty-string case above): the CLI
		// must reject the command for a missing required flag before broadcast.
		txResult, err := AddModel(AddModelOpts{VID: vid, PID: pid, OmitFrom: true})
		cliputils.RequireTxFailContains(t, txResult, err, `required flag(s) "from" not set`)
	})

	t.Run("AddModelVersion_OtaNegatives", func(t *testing.T) {
		// A model is needed before adding versions.
		mvPid := rand.Intn(65534) + 1
		mvSv := rand.Intn(65534) + 1
		mvSvs := fmt.Sprintf("%d", rand.Intn(65534)+1)
		txResult, err := AddModel(AddModelOpts{VID: vid, PID: mvPid, From: vendorAccount})
		cliputils.RequireTxOK(t, txResult, err)

		// otaChecksumType outside the IANA allow-list is rejected at the handler.
		txResult, err = AddModelVersion(AddModelVersionOpts{
			VID: vid, PID: mvPid, SoftwareVersion: mvSv, SoftwareVersionString: mvSvs,
			OtaURL: "https://ota.url.com", OtaFileSize: 123,
			OtaChecksum:     "MjFiZmYxN2YyMTRlMGJiMGMwNzhlNzIzOGIxZWE1ODk=",
			OtaChecksumType: 2,
			From:            vendorAccount,
		})
		cliputils.RequireTxFailContains(t, txResult, err, "OtaChecksumType 2 is not supported")

		// otaChecksum longer than 88 chars is rejected by ValidateBasic.
		txResult, err = AddModelVersion(AddModelVersionOpts{
			VID: vid, PID: mvPid, SoftwareVersion: mvSv, SoftwareVersionString: mvSvs,
			OtaURL: "https://ota.url.com", OtaFileSize: 123,
			OtaChecksum:     strings.Repeat("a", 89),
			OtaChecksumType: 1,
			From:            vendorAccount,
		})
		cliputils.RequireTxFailContains(t, txResult, err, "maximum length for OtaChecksum allowed is 88")
	})

	t.Run("AddModel_EmptyProductLabel_Fails", func(t *testing.T) {
		out, err := utils.ExecuteCLI("tx", "model", "add-model",
			"--vid", fmt.Sprintf("%d", vid),
			"--pid", fmt.Sprintf("%d", pid),
			"--deviceTypeID", "1",
			"--productName", "TestProduct",
			"--productLabel", "",
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
		require.Contains(t, combined, "ProductLabel is a required field")
	})

	t.Run("AddModel_EmptyPartNumber_Fails", func(t *testing.T) {
		out, err := utils.ExecuteCLI("tx", "model", "add-model",
			"--vid", fmt.Sprintf("%d", vid),
			"--pid", fmt.Sprintf("%d", pid),
			"--deviceTypeID", "1",
			"--productName", "TestProduct",
			"--productLabel", "TestingProductLabel",
			"--partNumber", "",
			"--commissioningCustomFlow", "0",
			"--enhancedSetupFlowOptions", "0",
			"--from", vendorAccount,
			"--yes", "-o", "json", "--keyring-backend", "test",
		)
		combined := string(out)
		if err != nil {
			combined += err.Error()
		}
		require.Contains(t, combined, "PartNumber is a required field")
	})

	t.Run("AddModel_DiscoveryCapabilitiesBitmaskAboveRange_Fails", func(t *testing.T) {
		out, err := utils.ExecuteCLI("tx", "model", "add-model",
			"--vid", fmt.Sprintf("%d", vid),
			"--pid", fmt.Sprintf("%d", pid),
			"--deviceTypeID", "1",
			"--productName", "TestProduct",
			"--productLabel", "TestingProductLabel",
			"--partNumber", "1",
			"--commissioningCustomFlow", "0",
			"--enhancedSetupFlowOptions", "0",
			"--discoveryCapabilitiesBitmask", "31",
			"--from", vendorAccount,
			"--yes", "-o", "json", "--keyring-backend", "test",
		)
		combined := string(out)
		if err != nil {
			combined += err.Error()
		}
		require.Contains(t, combined, "DiscoveryCapabilitiesBitmask must not be greater than 30")
	})

	// AddModelVersion negative cases for v1.5.2+ ota validations.
	// Use a fresh pid so we can add a model first.
	t.Run("AddModelVersion_OTANegativeCases", func(t *testing.T) {
		mvPid := rand.Intn(65534) + 1
		mvSv := rand.Intn(65534) + 1
		mvSvs := fmt.Sprintf("%d", rand.Intn(65534)+1)

		txResult, err := utils.ExecuteTx("tx", "model", "add-model",
			"--vid", fmt.Sprintf("%d", vid),
			"--pid", fmt.Sprintf("%d", mvPid),
			"--deviceTypeID", "1",
			"--productName", "TestProduct",
			"--productLabel", "TestingProductLabel",
			"--partNumber", "1",
			"--commissioningCustomFlow", "0",
			"--enhancedSetupFlowOptions", "0",
			"--from", vendorAccount,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code, "add-model failed: %s", txResult.RawLog)

		t.Run("OtaChecksumType_NotAllowed_Fails", func(t *testing.T) {
			// otaChecksumType=2 is outside the IANA-allowed subset {1,7,8,10,11,12}.
			// This is enforced at the handler, so the tx is broadcast and rejected on-chain.
			res, txErr := utils.ExecuteTx("tx", "model", "add-model-version",
				"--vid", fmt.Sprintf("%d", vid),
				"--pid", fmt.Sprintf("%d", mvPid),
				"--softwareVersion", fmt.Sprintf("%d", mvSv),
				"--softwareVersionString", mvSvs,
				"--cdVersionNumber", "1",
				"--minApplicableSoftwareVersion", "1",
				"--maxApplicableSoftwareVersion", "10",
				"--otaURL", "https://ota.url.com",
				"--otaFileSize", "123",
				"--otaChecksum", "MjFiZmYxN2YyMTRlMGJiMGMwNzhlNzIzOGIxZWE1ODk=",
				"--otaChecksumType", "2",
				"--from", vendorAccount,
			)
			require.NoError(t, txErr)
			require.NotEqual(t, uint32(0), res.Code)
			require.Contains(t, res.RawLog, "OtaChecksumType 2 is not supported")
		})

		t.Run("OtaChecksum_TooLong_Fails", func(t *testing.T) {
			// 89 'A's exceeds the 88-char validate-basic cap, so it's rejected client-side.
			longChecksum := make([]byte, 89)
			for i := range longChecksum {
				longChecksum[i] = 'A'
			}
			out, cliErr := utils.ExecuteCLI("tx", "model", "add-model-version",
				"--vid", fmt.Sprintf("%d", vid),
				"--pid", fmt.Sprintf("%d", mvPid),
				"--softwareVersion", fmt.Sprintf("%d", mvSv),
				"--softwareVersionString", mvSvs,
				"--cdVersionNumber", "1",
				"--minApplicableSoftwareVersion", "1",
				"--maxApplicableSoftwareVersion", "10",
				"--otaURL", "https://ota.url.com",
				"--otaFileSize", "123",
				"--otaChecksum", string(longChecksum),
				"--otaChecksumType", "1",
				"--from", vendorAccount,
				"--yes", "-o", "json", "--keyring-backend", "test",
			)
			combined := string(out)
			if cliErr != nil {
				combined += cliErr.Error()
			}
			require.Contains(t, combined, "maximum length for OtaChecksum allowed is 88")
		})
	})
}
