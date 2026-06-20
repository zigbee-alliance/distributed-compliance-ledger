package model

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	cliputils "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/cli/utils"
	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/utils"
)

// requireTxOK asserts a tx executed with on-chain code 0.
func requireTxOK(t *testing.T, txResult *utils.TxResult, err error) {
	t.Helper()
	require.NoError(t, err)
	require.Equal(t, uint32(0), txResult.Code, "tx raw_log: %s", txResult.RawLog)
	_, awaitErr := utils.AwaitTxConfirmation(txResult.TxHash)
	require.NoError(t, awaitErr)
}

// requireTxFailContains asserts a tx failed (either at CLI level or on-chain)
// and the error message contains the expected substring.
func requireTxFailContains(t *testing.T, txResult *utils.TxResult, err error, want string) {
	t.Helper()
	var msg string
	switch {
	case err != nil:
		msg = err.Error()
	case txResult == nil:
		t.Fatalf("expected failure containing %q, got nil tx and nil err", want)
	default:
		require.NotEqual(t, uint32(0), txResult.Code,
			"expected non-zero code, raw_log: %s", txResult.RawLog)
		msg = txResult.RawLog
	}
	require.True(t, strings.Contains(msg, want),
		"expected error to contain %q, got: %s", want, msg)
}

func TestModelValidationCases(t *testing.T) {
	vid1 := rand.Intn(65534) + 1
	vendorAccount1 := fmt.Sprintf("vendor_account_%d", vid1)
	cliputils.CreateVendorAccount(t, vendorAccount1, vid1)

	pid1 := rand.Intn(65534) + 1
	pid2 := rand.Intn(65534) + 1
	pid3 := rand.Intn(65534) + 1

	// --- AddModel ---

	t.Run("AddModel_MinimumFields", func(t *testing.T) {
		txResult, err := AddModel(AddModelOpts{
			VID:          vid1,
			PID:          pid1,
			ProductLabel: "Test Product",
			From:         vendorAccount1,
		})
		requireTxOK(t, txResult, err)

		m, err := GetModel(vid1, pid1)
		require.NoError(t, err)
		require.NotNil(t, m)
		require.Equal(t, int32(vid1), m.Vid)
		require.Equal(t, int32(pid1), m.Pid)
		require.Equal(t, "TestProduct", m.ProductName)
		require.Equal(t, "1", m.PartNumber)
		require.Equal(t, int32(0), m.CommissioningCustomFlow)
	})

	t.Run("AddModel_AllFields", func(t *testing.T) {
		txResult, err := AddModel(AddModelOpts{
			VID:                     vid1,
			PID:                     pid2,
			DeviceTypeID:            2,
			ProductName:             "Test Product with All Fields",
			ProductLabel:            "Test Product with All fields",
			PartNumber:              "23.456",
			CommissioningCustomFlow: 1,
			From:                    vendorAccount1,
			Extra: []string{
				"--commissioningCustomFlowURL", "https://customflow.url.info",
				"--commissioningModeInitialStepsHint", "1",
				"--commissioningModeInitialStepsInstruction", "Initial Instructions",
				"--commissioningModeSecondaryStepsHint", "2",
				"--commissioningModeSecondaryStepsInstruction", "Secondary Steps Instruction",
				"--icdUserActiveModeTriggerHint", "4",
				"--icdUserActiveModeTriggerInstruction", "ICD User Active Mode Trigger Instruction",
				"--factoryResetStepsHint", "3",
				"--factoryResetStepsInstruction", "Factory Reset Steps Instruction",
				"--userManualURL", "https://usermanual.url",
				"--productURL", "https://product.url.info",
				"--lsfURL", "https://lsf.url.info",
				"--supportURL", "https://support.url.info",
			},
		})
		requireTxOK(t, txResult, err)

		m, err := GetModel(vid1, pid2)
		require.NoError(t, err)
		require.NotNil(t, m)
		require.Equal(t, int32(vid1), m.Vid)
		require.Equal(t, int32(pid2), m.Pid)
		require.Equal(t, int32(2), m.DeviceTypeId)
		require.Equal(t, "Test Product with All Fields", m.ProductName)
		require.Equal(t, "Test Product with All fields", m.ProductLabel)
		require.Equal(t, "23.456", m.PartNumber)
		require.Equal(t, int32(1), m.CommissioningCustomFlow)
		require.Equal(t, "https://customflow.url.info", m.CommissioningCustomFlowUrl)
		require.Equal(t, uint32(1), m.CommissioningModeInitialStepsHint)
		require.Equal(t, "Initial Instructions", m.CommissioningModeInitialStepsInstruction)
		require.Equal(t, uint32(2), m.CommissioningModeSecondaryStepsHint)
		require.Equal(t, "Secondary Steps Instruction", m.CommissioningModeSecondaryStepsInstruction)
		require.Equal(t, uint32(4), m.IcdUserActiveModeTriggerHint)
		require.Equal(t, "ICD User Active Mode Trigger Instruction", m.IcdUserActiveModeTriggerInstruction)
		require.Equal(t, uint32(3), m.FactoryResetStepsHint)
		require.Equal(t, "Factory Reset Steps Instruction", m.FactoryResetStepsInstruction)
		require.Equal(t, "https://usermanual.url", m.UserManualUrl)
		require.Equal(t, "https://support.url.info", m.SupportUrl)
		require.Equal(t, "https://product.url.info", m.ProductUrl)
		require.Equal(t, "https://lsf.url.info", m.LsfUrl)
		require.Equal(t, int32(1), m.LsfRevision)
	})

	t.Run("AddModel_MandatoryAndSomeOptional_Pid3", func(t *testing.T) {
		txResult, err := AddModel(AddModelOpts{
			VID:                     vid1,
			PID:                     pid3,
			DeviceTypeID:            2,
			ProductName:             "Test Product with All Fields",
			ProductLabel:            "Test Product with All fields",
			PartNumber:              "23.456",
			CommissioningCustomFlow: 1,
			From:                    vendorAccount1,
			Extra: []string{
				"--commissioningCustomFlowURL", "https://customflow.url.info",
				"--commissioningModeInitialStepsHint", "1",
				"--commissioningModeInitialStepsInstruction", "Initial Instructions",
				"--commissioningModeSecondaryStepsHint", "2",
				"--commissioningModeSecondaryStepsInstruction", "Secondary Steps Instruction",
				"--icdUserActiveModeTriggerHint", "4",
				"--icdUserActiveModeTriggerInstruction", "ICD User Active Mode Trigger Instruction",
				"--factoryResetStepsHint", "3",
				"--factoryResetStepsInstruction", "Factory Reset Steps Instruction",
			},
		})
		requireTxOK(t, txResult, err)

		m, err := GetModel(vid1, pid3)
		require.NoError(t, err)
		require.NotNil(t, m)
		require.Equal(t, int32(vid1), m.Vid)
		require.Equal(t, int32(pid3), m.Pid)
		require.Equal(t, "Factory Reset Steps Instruction", m.FactoryResetStepsInstruction)
	})

	// --- UpdateModel ---

	t.Run("UpdateModel_MultipleFields_Pid1", func(t *testing.T) {
		txResult, err := UpdateModel(vid1, pid1, vendorAccount1,
			"--productName", "Updated Product Name",
			"--productLabel", "Updated Test Product",
			"--partNumber", "2",
			"--lsfURL", "https://lsf.url.info?v=1",
			"--lsfRevision", "1",
			"--enhancedSetupFlowOptions", "0",
		)
		requireTxOK(t, txResult, err)

		m, err := GetModel(vid1, pid1)
		require.NoError(t, err)
		require.NotNil(t, m)
		require.Equal(t, "Updated Product Name", m.ProductName)
		require.Equal(t, "2", m.PartNumber)
		require.Equal(t, "Updated Test Product", m.ProductLabel)
		require.Equal(t, int32(0), m.CommissioningCustomFlow)
		require.Equal(t, "https://lsf.url.info?v=1", m.LsfUrl)
		require.Equal(t, int32(1), m.LsfRevision)
	})

	t.Run("UpdateModel_SingleField_Pid1", func(t *testing.T) {
		txResult, err := UpdateModel(vid1, pid1, vendorAccount1,
			"--productLabel", "Updated Test Product V2",
			"--enhancedSetupFlowOptions", "0",
		)
		requireTxOK(t, txResult, err)

		m, err := GetModel(vid1, pid1)
		require.NoError(t, err)
		require.NotNil(t, m)
		require.Equal(t, "Updated Product Name", m.ProductName) // unchanged
		require.Equal(t, "2", m.PartNumber)                     // unchanged
		require.Equal(t, "Updated Test Product V2", m.ProductLabel)
	})

	t.Run("UpdateModel_AllFields_Pid1", func(t *testing.T) {
		txResult, err := UpdateModel(vid1, pid1, vendorAccount1,
			"--productName", "Updated Product Name V3",
			"--partNumber", "V3",
			"--commissioningCustomFlowURL", "https://updated.url.info",
			"--productLabel", "Updated Test Product V3",
			"--commissioningModeInitialStepsInstruction", "Instructions updated v3",
			"--commissioningModeSecondaryStepsInstruction", "Secondary Instructions v3",
			"--icdUserActiveModeTriggerInstruction", "ICD User Active Mode Trigger Instructions v3",
			"--factoryResetStepsInstruction", "Factory Reset Instructions v3",
			"--userManualURL", "https://userManual.info/v3",
			"--supportURL", "https://support.url.info/v3",
			"--productURL", "https://product.landingpage.url",
			"--lsfURL", "https://lsf.url.info?v=2",
			"--lsfRevision", "2",
			"--enhancedSetupFlowOptions", "0",
		)
		requireTxOK(t, txResult, err)

		m, err := GetModel(vid1, pid1)
		require.NoError(t, err)
		require.NotNil(t, m)
		require.Equal(t, "Updated Product Name V3", m.ProductName)
		require.Equal(t, "V3", m.PartNumber)
		require.Equal(t, "Updated Test Product V3", m.ProductLabel)
		require.Equal(t, "https://updated.url.info", m.CommissioningCustomFlowUrl)
		require.Equal(t, "Instructions updated v3", m.CommissioningModeInitialStepsInstruction)
		require.Equal(t, "Secondary Instructions v3", m.CommissioningModeSecondaryStepsInstruction)
		require.Equal(t, "ICD User Active Mode Trigger Instructions v3", m.IcdUserActiveModeTriggerInstruction)
		require.Equal(t, "Factory Reset Instructions v3", m.FactoryResetStepsInstruction)
		require.Equal(t, "https://userManual.info/v3", m.UserManualUrl)
		require.Equal(t, "https://support.url.info/v3", m.SupportUrl)
		require.Equal(t, "https://product.landingpage.url", m.ProductUrl)
		require.Equal(t, "https://lsf.url.info?v=2", m.LsfUrl)
		require.Equal(t, int32(2), m.LsfRevision)
	})

	t.Run("UpdateModel_OneFieldPreservesOthers_Pid1", func(t *testing.T) {
		txResult, err := UpdateModel(vid1, pid1, vendorAccount1,
			"--productLabel", "Updated Test Product V4",
			"--enhancedSetupFlowOptions", "0",
		)
		requireTxOK(t, txResult, err)

		m, err := GetModel(vid1, pid1)
		require.NoError(t, err)
		require.NotNil(t, m)
		require.Equal(t, "Updated Test Product V4", m.ProductLabel)
		require.Equal(t, "Updated Product Name V3", m.ProductName) // unchanged
		require.Equal(t, int32(2), m.LsfRevision)                  // unchanged
	})

	t.Run("UpdateModel_NoFields_Pid1", func(t *testing.T) {
		txResult, err := UpdateModel(vid1, pid1, vendorAccount1,
			"--enhancedSetupFlowOptions", "0",
		)
		requireTxOK(t, txResult, err)

		m, err := GetModel(vid1, pid1)
		require.NoError(t, err)
		require.NotNil(t, m)
		// All previously-set fields should remain.
		require.Equal(t, "Updated Test Product V4", m.ProductLabel)
		require.Equal(t, "Updated Product Name V3", m.ProductName)
		require.Equal(t, int32(2), m.LsfRevision)
	})

	t.Run("UpdateModel_LsfRevisionEqual_Fails_Pid1", func(t *testing.T) {
		txResult, err := UpdateModel(vid1, pid1, vendorAccount1,
			"--lsfURL", "https://lsf.url.info?v=3",
			"--lsfRevision", "2",
			"--enhancedSetupFlowOptions", "0",
		)
		requireTxFailContains(t, txResult, err, "LsfRevision should monotonically increase by 1")
	})

	t.Run("UpdateModel_LsfRevisionLess_Fails_Pid1", func(t *testing.T) {
		txResult, err := UpdateModel(vid1, pid1, vendorAccount1,
			"--lsfURL", "https://lsf.url.info?v=3",
			"--lsfRevision", "1",
			"--enhancedSetupFlowOptions", "0",
		)
		requireTxFailContains(t, txResult, err, "LsfRevision should monotonically increase by 1")
	})

	// --- Model Version: minimum fields, then increment-style updates ---

	svBasic := rand.Intn(65534) + 1

	t.Run("AddModelVersion_MinimumFields", func(t *testing.T) {
		txResult, err := AddModelVersion(AddModelVersionOpts{
			VID:                          vid1,
			PID:                          pid1,
			SoftwareVersion:              svBasic,
			SoftwareVersionString:        "1",
			MinApplicableSoftwareVersion: 10,
			MaxApplicableSoftwareVersion: 20,
			From:                         vendorAccount1,
		})
		requireTxOK(t, txResult, err)

		mv, err := GetModelVersion(vid1, pid1, svBasic)
		require.NoError(t, err)
		require.NotNil(t, mv)
		require.Equal(t, int32(vid1), mv.Vid)
		require.Equal(t, int32(pid1), mv.Pid)
		require.Equal(t, uint32(svBasic), mv.SoftwareVersion)
		require.Equal(t, "1", mv.SoftwareVersionString)
		require.Equal(t, int32(1), mv.CdVersionNumber)
		require.True(t, mv.SoftwareVersionValid)
		require.Equal(t, uint32(10), mv.MinApplicableSoftwareVersion)
		require.Equal(t, uint32(20), mv.MaxApplicableSoftwareVersion)
	})

	t.Run("UpdateModelVersion_OnlyValidity_Basic", func(t *testing.T) {
		txResult, err := UpdateModelVersion(vid1, pid1, svBasic, vendorAccount1,
			"--softwareVersionValid=false",
		)
		requireTxOK(t, txResult, err)

		mv, err := GetModelVersion(vid1, pid1, svBasic)
		require.NoError(t, err)
		require.NotNil(t, mv)
		require.False(t, mv.SoftwareVersionValid)
	})

	t.Run("UpdateModelVersion_FewFields_Basic", func(t *testing.T) {
		txResult, err := UpdateModelVersion(vid1, pid1, svBasic, vendorAccount1,
			"--softwareVersionValid=true",
			"--releaseNotesURL", "https://release.url.info",
			"--otaURL", "https://ota.url.com",
			"--otaFileSize", "123",
			"--otaChecksum", "MjFiZmYxN2YyMTRlMGJiMGMwNzhlNzIzOGIxZWE1ODk1Mzg4MjA3ZmFhNmM2NTg2YTBmNDU0MDk3YTU0ZWIzMw==",
			"--otaChecksumType", "1",
			"--minApplicableSoftwareVersion", "2",
			"--maxApplicableSoftwareVersion", "20",
		)
		requireTxOK(t, txResult, err)

		mv, err := GetModelVersion(vid1, pid1, svBasic)
		require.NoError(t, err)
		require.NotNil(t, mv)
		require.True(t, mv.SoftwareVersionValid)
		require.Equal(t, "https://ota.url.com", mv.OtaUrl)
		require.Equal(t, uint64(123), mv.OtaFileSize)
		require.Equal(t, int32(1), mv.OtaChecksumType)
		require.Equal(t, uint32(2), mv.MinApplicableSoftwareVersion)
		require.Equal(t, uint32(20), mv.MaxApplicableSoftwareVersion)
		require.Equal(t, "https://release.url.info", mv.ReleaseNotesUrl)
	})

	// --- Model Version: full create + update lifecycle ---

	svFull := rand.Intn(65534) + 1

	t.Run("AddModelVersion_AllFields_Full", func(t *testing.T) {
		txResult, err := AddModelVersion(AddModelVersionOpts{
			VID:                          vid1,
			PID:                          pid1,
			SoftwareVersion:              svFull,
			SoftwareVersionString:        "1.0",
			CDVersionNumber:              21334,
			MinApplicableSoftwareVersion: 5,
			MaxApplicableSoftwareVersion: 32,
			OtaURL:                       "https://ota.url.info",
			OtaFileSize:                  123456789,
			OtaChecksum:                  "MjFiZmYxN2YyMTRlMGJiMGMwNzhlNzIzOGIxZWE1ODk=",
			From:                         vendorAccount1,
			Extra: []string{
				"--firmwareInformation", "123456789012345678901234567890123456789012345678901234567890123",
				"--softwareVersionValid=true",
				"--specificationVersion", "4",
				"--releaseNotesURL", "https://release.notes.url.info",
				"--otaChecksumType", "1",
			},
		})
		requireTxOK(t, txResult, err)

		mv, err := GetModelVersion(vid1, pid1, svFull)
		require.NoError(t, err)
		require.NotNil(t, mv)
		require.Equal(t, "1.0", mv.SoftwareVersionString)
		require.Equal(t, int32(21334), mv.CdVersionNumber)
		require.Equal(t, "123456789012345678901234567890123456789012345678901234567890123", mv.FirmwareInformation)
		require.Equal(t, "https://ota.url.info", mv.OtaUrl)
		require.Equal(t, uint64(123456789), mv.OtaFileSize)
		require.Equal(t, "MjFiZmYxN2YyMTRlMGJiMGMwNzhlNzIzOGIxZWE1ODk=", mv.OtaChecksum)
		require.Equal(t, int32(1), mv.OtaChecksumType)
		require.Equal(t, "https://release.notes.url.info", mv.ReleaseNotesUrl)
		require.Equal(t, uint32(32), mv.MaxApplicableSoftwareVersion)
		require.Equal(t, uint32(5), mv.MinApplicableSoftwareVersion)
		require.Equal(t, uint32(4), mv.SpecificationVersion) //nolint:staticcheck // intentionally testing the deprecated field is still served
	})

	t.Run("UpdateModelVersion_NoChange_Full", func(t *testing.T) {
		// Update with only the required identifiers — nothing should change.
		txResult, err := UpdateModelVersion(vid1, pid1, svFull, vendorAccount1)
		requireTxOK(t, txResult, err)

		mv, err := GetModelVersion(vid1, pid1, svFull)
		require.NoError(t, err)
		require.NotNil(t, mv)
		require.Equal(t, int32(21334), mv.CdVersionNumber)
		require.Equal(t, "https://ota.url.info", mv.OtaUrl)
		require.Equal(t, uint32(32), mv.MaxApplicableSoftwareVersion)
	})

	t.Run("UpdateModelVersion_OnlyValidity_Full", func(t *testing.T) {
		txResult, err := UpdateModelVersion(vid1, pid1, svFull, vendorAccount1,
			"--softwareVersionValid=false",
		)
		requireTxOK(t, txResult, err)

		mv, err := GetModelVersion(vid1, pid1, svFull)
		require.NoError(t, err)
		require.NotNil(t, mv)
		require.False(t, mv.SoftwareVersionValid)
		require.Equal(t, "https://ota.url.info", mv.OtaUrl) // unchanged
	})

	t.Run("UpdateModelVersion_AllMutable_Full", func(t *testing.T) {
		txResult, err := UpdateModelVersion(vid1, pid1, svFull, vendorAccount1,
			"--softwareVersionValid=true",
			"--otaURL", "https://updated.ota.url.info",
			"--releaseNotesURL", "https://updated.release.notes.url.info",
			"--maxApplicableSoftwareVersion", "25",
			"--minApplicableSoftwareVersion", "15",
		)
		requireTxOK(t, txResult, err)

		mv, err := GetModelVersion(vid1, pid1, svFull)
		require.NoError(t, err)
		require.NotNil(t, mv)
		require.True(t, mv.SoftwareVersionValid)
		require.Equal(t, "https://updated.ota.url.info", mv.OtaUrl)
		require.Equal(t, "https://updated.release.notes.url.info", mv.ReleaseNotesUrl)
		require.Equal(t, uint32(25), mv.MaxApplicableSoftwareVersion)
		require.Equal(t, uint32(15), mv.MinApplicableSoftwareVersion)
	})

	t.Run("UpdateModelVersion_OtaFieldsWithoutUrl_Fails", func(t *testing.T) {
		// OTA already set during create — try to change otaFileSize on its own.
		txResult, err := UpdateModelVersion(vid1, pid1, svFull, vendorAccount1,
			"--otaFileSize", "12345",
		)
		requireTxFailContains(t, txResult, err,
			"OtaUrl is not provided. OtaFileSize, OtaChecksum, and OtaChecksumType fields must not be provided")
	})

	// --- Model Version: bounds validation ---

	svBounds := rand.Intn(65534) + 1

	t.Run("AddModelVersion_WithSpecVer6", func(t *testing.T) {
		txResult, err := AddModelVersion(AddModelVersionOpts{
			VID:                          vid1,
			PID:                          pid1,
			SoftwareVersion:              svBounds,
			SoftwareVersionString:        "1.0",
			CDVersionNumber:              21334,
			MinApplicableSoftwareVersion: 5,
			MaxApplicableSoftwareVersion: 32,
			OtaURL:                       "https://ota.url.info",
			OtaFileSize:                  123456789,
			OtaChecksum:                  "MjFiZmYxN2YyMTRlMGJiMGMwNzhlNzIzOGIxZWE1ODk=",
			From:                         vendorAccount1,
			Extra: []string{
				"--firmwareInformation", "123456789012345678901234567890123456789012345678901234567890123",
				"--softwareVersionValid=true",
				"--specificationVersion", "6",
				"--otaChecksumType", "1",
			},
		})
		requireTxOK(t, txResult, err)

		mv, err := GetModelVersion(vid1, pid1, svBounds)
		require.NoError(t, err)
		require.NotNil(t, mv)
		require.Equal(t, uint32(6), mv.SpecificationVersion) //nolint:staticcheck // intentionally testing the deprecated field is still served
	})

	t.Run("UpdateModelVersion_MaxLessThanMin_Fails", func(t *testing.T) {
		txResult, err := UpdateModelVersion(vid1, pid1, svBounds, vendorAccount1,
			"--maxApplicableSoftwareVersion", "3",
			"--minApplicableSoftwareVersion", "5",
		)
		requireTxFailContains(t, txResult, err,
			"MaxApplicableSoftwareVersion must not be less than MinApplicableSoftwareVersion")
	})

	t.Run("UpdateModelVersion_MinGreaterThanMax_Fails", func(t *testing.T) {
		txResult, err := UpdateModelVersion(vid1, pid1, svBounds, vendorAccount1,
			"--maxApplicableSoftwareVersion", "32",
			"--minApplicableSoftwareVersion", "33",
		)
		requireTxFailContains(t, txResult, err,
			"MaxApplicableSoftwareVersion must not be less than MinApplicableSoftwareVersion")
	})

	// --- OtaURL set on update without other OTA fields ---

	svNoOta := rand.Intn(65534) + 1

	t.Run("AddModelVersion_NoOta_ThenUpdateOtaUrlOnly_Fails", func(t *testing.T) {
		txResult, err := AddModelVersion(AddModelVersionOpts{
			VID:                          vid1,
			PID:                          pid1,
			SoftwareVersion:              svNoOta,
			SoftwareVersionString:        "1",
			MinApplicableSoftwareVersion: 10,
			MaxApplicableSoftwareVersion: 20,
			From:                         vendorAccount1,
		})
		requireTxOK(t, txResult, err)

		txResult, err = UpdateModelVersion(vid1, pid1, svNoOta, vendorAccount1,
			"--otaURL", "https://ota.url.com",
		)
		requireTxFailContains(t, txResult, err,
			"OtaFileSize, OtaChecksum and OtaChecksumType are required if OtaUrl is provided")
	})

	// --- ModelVersion without --specificationVersion: defaults to 0 ---

	svNoSpecVer := rand.Intn(65534) + 1

	t.Run("AddModelVersion_NoSpecVer_DefaultsZero", func(t *testing.T) {
		txResult, err := AddModelVersion(AddModelVersionOpts{
			VID:                          vid1,
			PID:                          pid1,
			SoftwareVersion:              svNoSpecVer,
			SoftwareVersionString:        "1.0",
			CDVersionNumber:              21334,
			MinApplicableSoftwareVersion: 5,
			MaxApplicableSoftwareVersion: 32,
			OtaURL:                       "https://ota.url.info",
			OtaFileSize:                  123456789,
			OtaChecksum:                  "MjFiZmYxN2YyMTRlMGJiMGMwNzhlNzIzOGIxZWE1ODk=",
			From:                         vendorAccount1,
			Extra: []string{
				"--firmwareInformation", "123456789012345678901234567890123456789012345678901234567890123",
				"--softwareVersionValid=true",
				"--releaseNotesURL", "https://release.notes.url.info",
				"--otaChecksumType", "1",
			},
		})
		requireTxOK(t, txResult, err)

		mv, err := GetModelVersion(vid1, pid1, svNoSpecVer)
		require.NoError(t, err)
		require.NotNil(t, mv)
		require.Equal(t, uint32(0), mv.SpecificationVersion) //nolint:staticcheck // intentionally testing the deprecated field is still served
	})

	// --- update-model-version must NOT accept --specificationVersion ---

	svSpecVerImmutable := rand.Intn(65534) + 1

	t.Run("AddModelVersion_AllFields_v33", func(t *testing.T) {
		txResult, err := AddModelVersion(AddModelVersionOpts{
			VID:                          vid1,
			PID:                          pid1,
			SoftwareVersion:              svSpecVerImmutable,
			SoftwareVersionString:        "1.0",
			CDVersionNumber:              21334,
			MinApplicableSoftwareVersion: 5,
			MaxApplicableSoftwareVersion: 32,
			OtaURL:                       "https://ota.url.info",
			OtaFileSize:                  123456789,
			OtaChecksum:                  "MjFiZmYxN2YyMTRlMGJiMGMwNzhlNzIzOGIxZWE1ODk=",
			From:                         vendorAccount1,
			Extra: []string{
				"--firmwareInformation", "123456789012345678901234567890123456789012345678901234567890123",
				"--softwareVersionValid=true",
				"--specificationVersion", "33",
				"--releaseNotesURL", "https://release.notes.url.info",
				"--otaChecksumType", "1",
			},
		})
		requireTxOK(t, txResult, err)
	})

	t.Run("UpdateModelVersion_SpecVerFlag_Rejected", func(t *testing.T) {
		// update-model-version must reject --specificationVersion at the CLI flag
		// parsing stage (the flag is not defined on the update subcommand).
		_, err := UpdateModelVersion(vid1, pid1, svSpecVerImmutable, vendorAccount1,
			"--specificationVersion", "11",
		)
		require.Error(t, err)
		require.Contains(t, err.Error(), "unknown flag: --specificationVersion")
	})

	t.Run("QueryModelVersion_SpecVerImmutable_Preserved", func(t *testing.T) {
		// After the rejected update above, specificationVersion must remain 33.
		mv, err := GetModelVersion(vid1, pid1, svSpecVerImmutable)
		require.NoError(t, err)
		require.NotNil(t, mv)
		require.Equal(t, uint32(33), mv.SpecificationVersion) //nolint:staticcheck // intentionally testing the deprecated field is still served
	})

	// Keep the AddModel_WithSchemaVersion path for ts-client compatibility coverage.
	t.Run("AddModel_WithSchemaVersion", func(t *testing.T) {
		extraPid := rand.Intn(65534) + 1
		txResult, err := AddModel(AddModelOpts{
			VID:           vid1,
			PID:           extraPid,
			ProductLabel:  "Test Product",
			SchemaVersion: "0",
			From:          vendorAccount1,
		})
		requireTxOK(t, txResult, err)

		m, err := GetModel(vid1, extraPid)
		require.NoError(t, err)
		require.NotNil(t, m)
		require.Equal(t, uint32(0), m.SchemaVersion)
	})
}
