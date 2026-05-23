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

// TestModelValidationCases translates model-validation-cases.sh.
func TestModelValidationCases(t *testing.T) {
	vid1 := rand.Intn(65534) + 1
	vendorAccount1 := fmt.Sprintf("vendor_account_%d", vid1)
	cliputils.CreateVendorAccount(t, vendorAccount1, vid1)

	pid1 := rand.Intn(65534) + 1
	pid2 := rand.Intn(65534) + 1
	pid3 := rand.Intn(65534) + 1

	// --- AddModel ---

	t.Run("AddModel_MinimumFields", func(t *testing.T) {
		txResult, err := utils.ExecuteTx("tx", "model", "add-model",
			"--vid", fmt.Sprintf("%d", vid1),
			"--pid", fmt.Sprintf("%d", pid1),
			"--deviceTypeID", "1",
			"--productName", "TestProduct",
			"--productLabel", "Test Product",
			"--partNumber", "1",
			"--enhancedSetupFlowOptions", "0",
			"--commissioningCustomFlow", "0",
			"--from", vendorAccount1,
		)
		requireTxOK(t, txResult, err)

		out, err := QueryModel(vid1, pid1)
		require.NoError(t, err)
		require.Contains(t, string(out), fmt.Sprintf(`"vid":%d`, vid1))
		require.Contains(t, string(out), fmt.Sprintf(`"pid":%d`, pid1))
		require.Contains(t, string(out), `"productName":"TestProduct"`)
		require.Contains(t, string(out), `"partNumber":"1"`)
		require.Contains(t, string(out), `"commissioningCustomFlow":0`)
	})

	t.Run("AddModel_AllFields", func(t *testing.T) {
		txResult, err := utils.ExecuteTx("tx", "model", "add-model",
			"--vid", fmt.Sprintf("%d", vid1),
			"--pid", fmt.Sprintf("%d", pid2),
			"--deviceTypeID", "2",
			"--productName", "Test Product with All Fields",
			"--productLabel", "Test Product with All fields",
			"--partNumber", "23.456",
			"--commissioningCustomFlow", "1",
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
			"--enhancedSetupFlowOptions", "0",
			"--from", vendorAccount1,
		)
		requireTxOK(t, txResult, err)

		out, err := QueryModel(vid1, pid2)
		require.NoError(t, err)
		s := string(out)
		require.Contains(t, s, fmt.Sprintf(`"vid":%d`, vid1))
		require.Contains(t, s, fmt.Sprintf(`"pid":%d`, pid2))
		require.Contains(t, s, `"deviceTypeId":2`)
		require.Contains(t, s, `"productName":"Test Product with All Fields"`)
		require.Contains(t, s, `"productLabel":"Test Product with All fields"`)
		require.Contains(t, s, `"partNumber":"23.456"`)
		require.Contains(t, s, `"commissioningCustomFlow":1`)
		require.Contains(t, s, `"commissioningCustomFlowUrl":"https://customflow.url.info"`)
		require.Contains(t, s, `"commissioningModeInitialStepsHint":1`)
		require.Contains(t, s, `"commissioningModeInitialStepsInstruction":"Initial Instructions"`)
		require.Contains(t, s, `"commissioningModeSecondaryStepsHint":2`)
		require.Contains(t, s, `"commissioningModeSecondaryStepsInstruction":"Secondary Steps Instruction"`)
		require.Contains(t, s, `"icdUserActiveModeTriggerHint":4`)
		require.Contains(t, s, `"icdUserActiveModeTriggerInstruction":"ICD User Active Mode Trigger Instruction"`)
		require.Contains(t, s, `"factoryResetStepsHint":3`)
		require.Contains(t, s, `"factoryResetStepsInstruction":"Factory Reset Steps Instruction"`)
		require.Contains(t, s, `"userManualUrl":"https://usermanual.url"`)
		require.Contains(t, s, `"supportUrl":"https://support.url.info"`)
		require.Contains(t, s, `"productUrl":"https://product.url.info"`)
		require.Contains(t, s, `"lsfUrl":"https://lsf.url.info"`)
		require.Contains(t, s, `"lsfRevision":1`)
	})

	t.Run("AddModel_MandatoryAndSomeOptional_Pid3", func(t *testing.T) {
		txResult, err := utils.ExecuteTx("tx", "model", "add-model",
			"--vid", fmt.Sprintf("%d", vid1),
			"--pid", fmt.Sprintf("%d", pid3),
			"--deviceTypeID", "2",
			"--productName", "Test Product with All Fields",
			"--productLabel", "Test Product with All fields",
			"--partNumber", "23.456",
			"--commissioningCustomFlow", "1",
			"--commissioningCustomFlowURL", "https://customflow.url.info",
			"--commissioningModeInitialStepsHint", "1",
			"--commissioningModeInitialStepsInstruction", "Initial Instructions",
			"--commissioningModeSecondaryStepsHint", "2",
			"--commissioningModeSecondaryStepsInstruction", "Secondary Steps Instruction",
			"--icdUserActiveModeTriggerHint", "4",
			"--icdUserActiveModeTriggerInstruction", "ICD User Active Mode Trigger Instruction",
			"--factoryResetStepsHint", "3",
			"--factoryResetStepsInstruction", "Factory Reset Steps Instruction",
			"--enhancedSetupFlowOptions", "0",
			"--from", vendorAccount1,
		)
		requireTxOK(t, txResult, err)

		out, err := QueryModel(vid1, pid3)
		require.NoError(t, err)
		s := string(out)
		require.Contains(t, s, fmt.Sprintf(`"vid":%d`, vid1))
		require.Contains(t, s, fmt.Sprintf(`"pid":%d`, pid3))
		require.Contains(t, s, `"factoryResetStepsInstruction":"Factory Reset Steps Instruction"`)
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

		out, err := QueryModel(vid1, pid1)
		require.NoError(t, err)
		s := string(out)
		require.Contains(t, s, `"productName":"Updated Product Name"`)
		require.Contains(t, s, `"partNumber":"2"`)
		require.Contains(t, s, `"productLabel":"Updated Test Product"`)
		require.Contains(t, s, `"commissioningCustomFlow":0`)
		require.Contains(t, s, `"lsfUrl":"https://lsf.url.info?v=1"`)
		require.Contains(t, s, `"lsfRevision":1`)
	})

	t.Run("UpdateModel_SingleField_Pid1", func(t *testing.T) {
		txResult, err := UpdateModel(vid1, pid1, vendorAccount1,
			"--productLabel", "Updated Test Product V2",
			"--enhancedSetupFlowOptions", "0",
		)
		requireTxOK(t, txResult, err)

		out, err := QueryModel(vid1, pid1)
		require.NoError(t, err)
		s := string(out)
		require.Contains(t, s, `"productName":"Updated Product Name"`) // unchanged
		require.Contains(t, s, `"partNumber":"2"`)                     // unchanged
		require.Contains(t, s, `"productLabel":"Updated Test Product V2"`)
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

		out, err := QueryModel(vid1, pid1)
		require.NoError(t, err)
		s := string(out)
		require.Contains(t, s, `"productName":"Updated Product Name V3"`)
		require.Contains(t, s, `"partNumber":"V3"`)
		require.Contains(t, s, `"productLabel":"Updated Test Product V3"`)
		require.Contains(t, s, `"commissioningCustomFlowUrl":"https://updated.url.info"`)
		require.Contains(t, s, `"commissioningModeInitialStepsInstruction":"Instructions updated v3"`)
		require.Contains(t, s, `"commissioningModeSecondaryStepsInstruction":"Secondary Instructions v3"`)
		require.Contains(t, s, `"icdUserActiveModeTriggerInstruction":"ICD User Active Mode Trigger Instructions v3"`)
		require.Contains(t, s, `"factoryResetStepsInstruction":"Factory Reset Instructions v3"`)
		require.Contains(t, s, `"userManualUrl":"https://userManual.info/v3"`)
		require.Contains(t, s, `"supportUrl":"https://support.url.info/v3"`)
		require.Contains(t, s, `"productUrl":"https://product.landingpage.url"`)
		require.Contains(t, s, `"lsfUrl":"https://lsf.url.info?v=2"`)
		require.Contains(t, s, `"lsfRevision":2`)
	})

	t.Run("UpdateModel_OneFieldPreservesOthers_Pid1", func(t *testing.T) {
		txResult, err := UpdateModel(vid1, pid1, vendorAccount1,
			"--productLabel", "Updated Test Product V4",
			"--enhancedSetupFlowOptions", "0",
		)
		requireTxOK(t, txResult, err)

		out, err := QueryModel(vid1, pid1)
		require.NoError(t, err)
		s := string(out)
		require.Contains(t, s, `"productLabel":"Updated Test Product V4"`)
		require.Contains(t, s, `"productName":"Updated Product Name V3"`) // unchanged
		require.Contains(t, s, `"lsfRevision":2`)                         // unchanged
	})

	t.Run("UpdateModel_NoFields_Pid1", func(t *testing.T) {
		txResult, err := UpdateModel(vid1, pid1, vendorAccount1,
			"--enhancedSetupFlowOptions", "0",
		)
		requireTxOK(t, txResult, err)

		out, err := QueryModel(vid1, pid1)
		require.NoError(t, err)
		s := string(out)
		// All previously-set fields should remain.
		require.Contains(t, s, `"productLabel":"Updated Test Product V4"`)
		require.Contains(t, s, `"productName":"Updated Product Name V3"`)
		require.Contains(t, s, `"lsfRevision":2`)
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
		txResult, err := utils.ExecuteTx("tx", "model", "add-model-version",
			"--vid", fmt.Sprintf("%d", vid1),
			"--pid", fmt.Sprintf("%d", pid1),
			"--softwareVersion", fmt.Sprintf("%d", svBasic),
			"--softwareVersionString", "1",
			"--cdVersionNumber", "1",
			"--maxApplicableSoftwareVersion", "20",
			"--minApplicableSoftwareVersion", "10",
			"--from", vendorAccount1,
		)
		requireTxOK(t, txResult, err)

		out, err := QueryModelVersion(vid1, pid1, svBasic)
		require.NoError(t, err)
		s := string(out)
		require.Contains(t, s, fmt.Sprintf(`"vid":%d`, vid1))
		require.Contains(t, s, fmt.Sprintf(`"pid":%d`, pid1))
		require.Contains(t, s, fmt.Sprintf(`"softwareVersion":%d`, svBasic))
		require.Contains(t, s, `"softwareVersionString":"1"`)
		require.Contains(t, s, `"cdVersionNumber":1`)
		require.Contains(t, s, `"softwareVersionValid":true`)
		require.Contains(t, s, `"minApplicableSoftwareVersion":"10"`)
		require.Contains(t, s, `"maxApplicableSoftwareVersion":"20"`)
	})

	t.Run("UpdateModelVersion_OnlyValidity_Basic", func(t *testing.T) {
		txResult, err := UpdateModelVersion(vid1, pid1, svBasic, vendorAccount1,
			"--softwareVersionValid", "false",
		)
		requireTxOK(t, txResult, err)

		out, err := QueryModelVersion(vid1, pid1, svBasic)
		require.NoError(t, err)
		require.Contains(t, string(out), `"softwareVersionValid":false`)
	})

	t.Run("UpdateModelVersion_FewFields_Basic", func(t *testing.T) {
		txResult, err := UpdateModelVersion(vid1, pid1, svBasic, vendorAccount1,
			"--softwareVersionValid", "true",
			"--releaseNotesURL", "https://release.url.info",
			"--otaURL", "https://ota.url.com",
			"--otaFileSize", "123",
			"--otaChecksum", "MjFiZmYxN2YyMTRlMGJiMGMwNzhlNzIzOGIxZWE1ODk1Mzg4MjA3ZmFhNmM2NTg2YTBmNDU0MDk3YTU0ZWIzMw==",
			"--otaChecksumType", "1",
			"--minApplicableSoftwareVersion", "2",
			"--maxApplicableSoftwareVersion", "20",
		)
		requireTxOK(t, txResult, err)

		out, err := QueryModelVersion(vid1, pid1, svBasic)
		require.NoError(t, err)
		s := string(out)
		require.Contains(t, s, `"softwareVersionValid":true`)
		require.Contains(t, s, `"otaUrl":"https://ota.url.com"`)
		require.Contains(t, s, `"otaFileSize":"123"`)
		require.Contains(t, s, `"otaChecksumType":1`)
		require.Contains(t, s, `"minApplicableSoftwareVersion":"2"`)
		require.Contains(t, s, `"maxApplicableSoftwareVersion":"20"`)
		require.Contains(t, s, `"releaseNotesUrl":"https://release.url.info"`)
	})

	// --- Model Version: full create + update lifecycle ---

	svFull := rand.Intn(65534) + 1

	t.Run("AddModelVersion_AllFields_Full", func(t *testing.T) {
		txResult, err := utils.ExecuteTx("tx", "model", "add-model-version",
			"--vid", fmt.Sprintf("%d", vid1),
			"--pid", fmt.Sprintf("%d", pid1),
			"--softwareVersion", fmt.Sprintf("%d", svFull),
			"--softwareVersionString", "1.0",
			"--cdVersionNumber", "21334",
			"--firmwareInformation", "123456789012345678901234567890123456789012345678901234567890123",
			"--softwareVersionValid", "true",
			"--otaURL", "https://ota.url.info",
			"--otaFileSize", "123456789",
			"--specificationVersion", "4",
			"--otaChecksum", "MjFiZmYxN2YyMTRlMGJiMGMwNzhlNzIzOGIxZWE1ODk=",
			"--releaseNotesURL", "https://release.notes.url.info",
			"--otaChecksumType", "1",
			"--maxApplicableSoftwareVersion", "32",
			"--minApplicableSoftwareVersion", "5",
			"--from", vendorAccount1,
		)
		requireTxOK(t, txResult, err)

		out, err := QueryModelVersion(vid1, pid1, svFull)
		require.NoError(t, err)
		s := string(out)
		require.Contains(t, s, `"softwareVersionString":"1.0"`)
		require.Contains(t, s, `"cdVersionNumber":21334`)
		require.Contains(t, s, `"firmwareInformation":"123456789012345678901234567890123456789012345678901234567890123"`)
		require.Contains(t, s, `"otaUrl":"https://ota.url.info"`)
		require.Contains(t, s, `"otaFileSize":"123456789"`)
		require.Contains(t, s, `"otaChecksum":"MjFiZmYxN2YyMTRlMGJiMGMwNzhlNzIzOGIxZWE1ODk="`)
		require.Contains(t, s, `"otaChecksumType":1`)
		require.Contains(t, s, `"releaseNotesUrl":"https://release.notes.url.info"`)
		require.Contains(t, s, `"maxApplicableSoftwareVersion":"32"`)
		require.Contains(t, s, `"minApplicableSoftwareVersion":"5"`)
		require.Contains(t, s, `"specificationVersion":4`)
	})

	t.Run("UpdateModelVersion_NoChange_Full", func(t *testing.T) {
		// Update with only the required identifiers — nothing should change.
		txResult, err := UpdateModelVersion(vid1, pid1, svFull, vendorAccount1)
		requireTxOK(t, txResult, err)

		out, err := QueryModelVersion(vid1, pid1, svFull)
		require.NoError(t, err)
		s := string(out)
		require.Contains(t, s, `"cdVersionNumber":21334`)
		require.Contains(t, s, `"otaUrl":"https://ota.url.info"`)
		require.Contains(t, s, `"maxApplicableSoftwareVersion":"32"`)
	})

	t.Run("UpdateModelVersion_OnlyValidity_Full", func(t *testing.T) {
		txResult, err := UpdateModelVersion(vid1, pid1, svFull, vendorAccount1,
			"--softwareVersionValid", "false",
		)
		requireTxOK(t, txResult, err)

		out, err := QueryModelVersion(vid1, pid1, svFull)
		require.NoError(t, err)
		s := string(out)
		require.Contains(t, s, `"softwareVersionValid":false`)
		require.Contains(t, s, `"otaUrl":"https://ota.url.info"`) // unchanged
	})

	t.Run("UpdateModelVersion_AllMutable_Full", func(t *testing.T) {
		txResult, err := UpdateModelVersion(vid1, pid1, svFull, vendorAccount1,
			"--softwareVersionValid", "true",
			"--otaURL", "https://updated.ota.url.info",
			"--releaseNotesURL", "https://updated.release.notes.url.info",
			"--maxApplicableSoftwareVersion", "25",
			"--minApplicableSoftwareVersion", "15",
		)
		requireTxOK(t, txResult, err)

		out, err := QueryModelVersion(vid1, pid1, svFull)
		require.NoError(t, err)
		s := string(out)
		require.Contains(t, s, `"softwareVersionValid":true`)
		require.Contains(t, s, `"otaUrl":"https://updated.ota.url.info"`)
		require.Contains(t, s, `"releaseNotesUrl":"https://updated.release.notes.url.info"`)
		require.Contains(t, s, `"maxApplicableSoftwareVersion":"25"`)
		require.Contains(t, s, `"minApplicableSoftwareVersion":"15"`)
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
		txResult, err := utils.ExecuteTx("tx", "model", "add-model-version",
			"--vid", fmt.Sprintf("%d", vid1),
			"--pid", fmt.Sprintf("%d", pid1),
			"--softwareVersion", fmt.Sprintf("%d", svBounds),
			"--softwareVersionString", "1.0",
			"--cdVersionNumber", "21334",
			"--firmwareInformation", "123456789012345678901234567890123456789012345678901234567890123",
			"--softwareVersionValid", "true",
			"--otaURL", "https://ota.url.info",
			"--otaFileSize", "123456789",
			"--specificationVersion", "6",
			"--otaChecksum", "MjFiZmYxN2YyMTRlMGJiMGMwNzhlNzIzOGIxZWE1ODk=",
			"--otaChecksumType", "1",
			"--maxApplicableSoftwareVersion", "32",
			"--minApplicableSoftwareVersion", "5",
			"--from", vendorAccount1,
		)
		requireTxOK(t, txResult, err)

		out, err := QueryModelVersion(vid1, pid1, svBounds)
		require.NoError(t, err)
		require.Contains(t, string(out), `"specificationVersion":6`)
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
		txResult, err := utils.ExecuteTx("tx", "model", "add-model-version",
			"--vid", fmt.Sprintf("%d", vid1),
			"--pid", fmt.Sprintf("%d", pid1),
			"--softwareVersion", fmt.Sprintf("%d", svNoOta),
			"--softwareVersionString", "1",
			"--cdVersionNumber", "1",
			"--maxApplicableSoftwareVersion", "20",
			"--minApplicableSoftwareVersion", "10",
			"--from", vendorAccount1,
		)
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
		txResult, err := utils.ExecuteTx("tx", "model", "add-model-version",
			"--vid", fmt.Sprintf("%d", vid1),
			"--pid", fmt.Sprintf("%d", pid1),
			"--softwareVersion", fmt.Sprintf("%d", svNoSpecVer),
			"--softwareVersionString", "1.0",
			"--cdVersionNumber", "21334",
			"--firmwareInformation", "123456789012345678901234567890123456789012345678901234567890123",
			"--softwareVersionValid", "true",
			"--otaURL", "https://ota.url.info",
			"--otaFileSize", "123456789",
			"--otaChecksum", "MjFiZmYxN2YyMTRlMGJiMGMwNzhlNzIzOGIxZWE1ODk=",
			"--releaseNotesURL", "https://release.notes.url.info",
			"--otaChecksumType", "1",
			"--maxApplicableSoftwareVersion", "32",
			"--minApplicableSoftwareVersion", "5",
			"--from", vendorAccount1,
		)
		requireTxOK(t, txResult, err)

		out, err := QueryModelVersion(vid1, pid1, svNoSpecVer)
		require.NoError(t, err)
		require.Contains(t, string(out), `"specificationVersion":0`)
	})

	// --- update-model-version must NOT accept --specificationVersion ---

	svSpecVerImmutable := rand.Intn(65534) + 1

	t.Run("AddModelVersion_AllFields_v33", func(t *testing.T) {
		txResult, err := utils.ExecuteTx("tx", "model", "add-model-version",
			"--vid", fmt.Sprintf("%d", vid1),
			"--pid", fmt.Sprintf("%d", pid1),
			"--softwareVersion", fmt.Sprintf("%d", svSpecVerImmutable),
			"--softwareVersionString", "1.0",
			"--cdVersionNumber", "21334",
			"--firmwareInformation", "123456789012345678901234567890123456789012345678901234567890123",
			"--softwareVersionValid", "true",
			"--otaURL", "https://ota.url.info",
			"--otaFileSize", "123456789",
			"--specificationVersion", "33",
			"--otaChecksum", "MjFiZmYxN2YyMTRlMGJiMGMwNzhlNzIzOGIxZWE1ODk=",
			"--releaseNotesURL", "https://release.notes.url.info",
			"--otaChecksumType", "1",
			"--maxApplicableSoftwareVersion", "32",
			"--minApplicableSoftwareVersion", "5",
			"--from", vendorAccount1,
		)
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
		out, err := QueryModelVersion(vid1, pid1, svSpecVerImmutable)
		require.NoError(t, err)
		require.Contains(t, string(out), `"specificationVersion":33`)
	})

	// Keep the AddModel_WithSchemaVersion path for ts-client compatibility coverage.
	t.Run("AddModel_WithSchemaVersion", func(t *testing.T) {
		extraPid := rand.Intn(65534) + 1
		txResult, err := utils.ExecuteTx("tx", "model", "add-model",
			"--vid", fmt.Sprintf("%d", vid1),
			"--pid", fmt.Sprintf("%d", extraPid),
			"--deviceTypeID", "1",
			"--productName", "TestProduct",
			"--productLabel", "Test Product",
			"--partNumber", "1",
			"--enhancedSetupFlowOptions", "0",
			"--commissioningCustomFlow", "0",
			"--schemaVersion", "0",
			"--from", vendorAccount1,
		)
		requireTxOK(t, txResult, err)

		out, err := QueryModel(vid1, extraPid)
		require.NoError(t, err)
		require.Contains(t, string(out), `"schemaVersion":0`)
	})
}
