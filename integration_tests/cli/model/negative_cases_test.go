package model

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
	cliputils "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/cli/utils"
	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/utils"
)

// TestModelNegativeCases translates model-negative-cases.sh.
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
		txResult, err := utils.ExecuteTx("tx", "model", "add-model",
			"--vid", fmt.Sprintf("%d", vid),
			"--pid", fmt.Sprintf("%d", pid),
			"--deviceTypeID", "1",
			"--productName", "TestProduct",
			"--productLabel", "TestingProductLabel",
			"--partNumber", "1",
			"--commissioningCustomFlow", "0",
			"--enhancedSetupFlowOptions", "0",
			"--from", certificationHouse,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(4), txResult.Code)
		_, _ = utils.AwaitTxConfirmation(txResult.TxHash)
	})

	t.Run("AddModel_VendorNonAssociatedPID_Fails", func(t *testing.T) {
		txResult, err := utils.ExecuteTx("tx", "model", "add-model",
			"--vid", fmt.Sprintf("%d", vidWithPids),
			"--pid", "101",
			"--deviceTypeID", "1",
			"--productName", "TestProduct",
			"--productLabel", "TestingProductLabel",
			"--partNumber", "1",
			"--commissioningCustomFlow", "0",
			"--enhancedSetupFlowOptions", "0",
			"--from", vendorAccountWithPids,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(4), txResult.Code)
		_, _ = utils.AwaitTxConfirmation(txResult.TxHash)
	})

	t.Run("AddModel_WrongVendorID_Fails", func(t *testing.T) {
		vid1 := rand.Intn(65534) + 1
		txResult, err := utils.ExecuteTx("tx", "model", "add-model",
			"--vid", fmt.Sprintf("%d", vid1),
			"--pid", fmt.Sprintf("%d", pid),
			"--deviceTypeID", "1",
			"--productName", "TestProduct",
			"--productLabel", "TestingProductLabel",
			"--partNumber", "1",
			"--commissioningCustomFlow", "0",
			"--enhancedSetupFlowOptions", "0",
			"--from", vendorAccount,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(4), txResult.Code)
		_, _ = utils.AwaitTxConfirmation(txResult.TxHash)
	})

	t.Run("AddModelTwice_Fails", func(t *testing.T) {
		// First add succeeds
		txResult, err := utils.ExecuteTx("tx", "model", "add-model",
			"--vid", fmt.Sprintf("%d", vid),
			"--pid", fmt.Sprintf("%d", pid),
			"--deviceTypeID", "1",
			"--productName", "TestProduct",
			"--productLabel", "TestingProductLabel",
			"--partNumber", "1",
			"--commissioningCustomFlow", "0",
			"--enhancedSetupFlowOptions", "0",
			"--from", vendorAccount,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Second add fails with 501
		txResult, err = utils.ExecuteTx("tx", "model", "add-model",
			"--vid", fmt.Sprintf("%d", vid),
			"--pid", fmt.Sprintf("%d", pid),
			"--deviceTypeID", "1",
			"--productName", "TestProduct",
			"--productLabel", "TestingProductLabel",
			"--partNumber", "1",
			"--commissioningCustomFlow", "0",
			"--enhancedSetupFlowOptions", "0",
			"--from", vendorAccount,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(501), txResult.Code)
		_, _ = utils.AwaitTxConfirmation(txResult.TxHash)
	})

	sv := rand.Intn(65534) + 1

	t.Run("AddModelVersion_ThenCertify_ThenDeleteCertifiedModel_Fails", func(t *testing.T) {
		txResult, err := utils.ExecuteTx("tx", "model", "add-model-version",
			"--cdVersionNumber", "1",
			"--maxApplicableSoftwareVersion", "10",
			"--minApplicableSoftwareVersion", "1",
			"--vid", fmt.Sprintf("%d", vid),
			"--pid", fmt.Sprintf("%d", pid),
			"--softwareVersion", fmt.Sprintf("%d", sv),
			"--softwareVersionString", softwareVersionString,
			"--from", vendorAccount,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		certificationDate := "2020-01-01T00:00:01Z"
		txResult, err = utils.ExecuteTx("tx", "compliance", "certify-model",
			"--vid", fmt.Sprintf("%d", vid),
			"--pid", fmt.Sprintf("%d", pid),
			"--softwareVersion", fmt.Sprintf("%d", sv),
			"--cdVersionNumber", "1",
			"--certificationType", "zigbee",
			"--certificationDate", certificationDate,
			"--softwareVersionString", softwareVersionString,
			"--cdCertificateId", "1230000000000000000",
			"--from", zbAccount,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Delete certified model — should fail with code 525
		txResult, err = utils.ExecuteTx("tx", "model", "delete-model",
			"--vid", fmt.Sprintf("%d", vid),
			"--pid", fmt.Sprintf("%d", pid),
			"--from", vendorAccount,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(525), txResult.Code)
		_, _ = utils.AwaitTxConfirmation(txResult.TxHash)
	})

	t.Run("AddModel_UnknownAccount_Fails", func(t *testing.T) {
		out, err := utils.ExecuteCLI("tx", "model", "add-model",
			"--vid", fmt.Sprintf("%d", vid),
			"--pid", fmt.Sprintf("%d", pid),
			"--deviceTypeID", "1",
			"--productName", "TestProduct",
			"--productLabel", "TestingProductLabel",
			"--partNumber", "1",
			"--commissioningCustomFlow", "0",
			"--enhancedSetupFlowOptions", "0",
			"--from", "Unknown",
			"--yes", "-o", "json", "--keyring-backend", "test",
		)
		combined := string(out)
		if err != nil {
			combined += err.Error()
		}
		require.Contains(t, combined, "key not found")
	})

	t.Run("AddModel_InvalidVidPid", func(t *testing.T) {
		for _, inv := range []string{"-1", "0", "65536", "string"} {
			out, err := utils.ExecuteCLI("tx", "model", "add-model",
				"--vid", inv,
				"--pid", fmt.Sprintf("%d", pid),
				"--deviceTypeID", "1",
				"--productName", "TestProduct",
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
			// Expect some error related to vid validation
			hasErr := len(combined) > 0 && (containsAny(combined, "Vid must not be", "invalid syntax", "invalid argument"))
			require.True(t, hasErr, "expected error for vid=%s, got: %s", inv, combined)
		}
	})

	t.Run("AddModel_EmptyProductName_Fails", func(t *testing.T) {
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
		out, err := utils.ExecuteCLI("tx", "model", "add-model",
			"--vid", fmt.Sprintf("%d", vid),
			"--pid", fmt.Sprintf("%d", pid),
			"--deviceTypeID", "1",
			"--productName", "TestProduct",
			"--productLabel", "TestingProductLabel",
			"--partNumber", "1",
			"--commissioningCustomFlow", "0",
			"--enhancedSetupFlowOptions", "0",
			"--from", "",
			"--yes", "-o", "json", "--keyring-backend", "test",
		)
		combined := string(out)
		if err != nil {
			combined += err.Error()
		}
		require.Contains(t, combined, "invalid creator address")
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

func containsAny(s string, subs ...string) bool {
	for _, sub := range subs {
		if len(sub) > 0 {
			idx := 0
			for i := 0; i <= len(s)-len(sub); i++ {
				if s[i:i+len(sub)] == sub {
					idx = 1

					break
				}
			}
			if idx > 0 {
				return true
			}
		}
	}

	return false
}
