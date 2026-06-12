// Copyright 2020 DSR Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package upgrade

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

// runUpgrade151To152 runs the v1.5.1 → v1.5.2 cosmovisor upgrade, then
// seeds 1.5.2-era state: a new vendor account (VID=65519, no vendorinfo
// record), models exercising the new ICD / factory-reset /
// specificationVersion fields, and a pid_3 add+delete to seed ghost state.
//
// Assumes the chain is currently running v1.5.1 with state seeded by
// phases 01-07.
//
//nolint:funlen
func runUpgrade151To152(t *testing.T, state *UpgradeTestState) {
	t.Helper()

	dcldOld, err := EnsureBinary("1.5.1")
	require.NoError(t, err, "fetch dcld v1.5.1")
	dcldNew, err := EnsureBinary("1.5.2")
	require.NoError(t, err, "fetch dcld v1.5.2")

	// `config broadcast-mode sync` — v1.4+ binaries no longer accept `block`.
	_, _ = ExecuteCLIWithBin(dcldNew, "config", "broadcast-mode", "sync")

	step := SoftwareUpgradeStep{
		PlanName:         PlanNameV1_5_2,
		BinaryVersionNew: BinaryVersionV1_5_2,
		Checksum:         UpgradeChecksumV1_5_2,
		DcldOldBin:       dcldOld,
		DcldNewBin:       dcldNew,
		Trustees: []string{
			state.Trustee1, state.Trustee2, state.Trustee3, state.Trustee4,
		},
	}
	step.Run(t)

	// ------------------------------------------------------------------
	// Verify carry-over data is intact under the new binary.
	// ------------------------------------------------------------------

	MustRun(t, "VerifyPreservedModels", func(t *testing.T) {
		t.Helper()
		out, err := ExecuteCLIWithBin(dcldNew,
			"query", "model", "get-model",
			"--vid", fmt.Sprintf("%d", state.VIDFor1_5_1),
			"--pid", fmt.Sprintf("%d", state.PID1For1_5_1),
		)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", state.VIDFor1_5_1)
		requireFieldEquals(t, out, "pid", state.PID1For1_5_1)
		checkResponseContains(t, out, state.ProductLabelFor1_5_1)
		requireFieldEquals(t, out, "commissioningModeSecondaryStepsHint",
			state.CommissioningModeSecondaryStepsHintFor1_5_1)

		out, err = ExecuteCLIWithBin(dcldNew,
			"query", "model", "get-model",
			"--vid", fmt.Sprintf("%d", state.VIDFor1_5_1),
			"--pid", fmt.Sprintf("%d", state.PID2For1_5_1),
		)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", state.VIDFor1_5_1)
		requireFieldEquals(t, out, "pid", state.PID2For1_5_1)
		checkResponseContains(t, out, state.ProductLabelFor1_5_1)
		// Migration default for this field is 4.
		requireFieldEquals(t, out, "commissioningModeSecondaryStepsHint", 4)
	})

	MustRun(t, "VerifyUpdatedModelFromScript1", func(t *testing.T) {
		t.Helper()
		out, err := ExecuteCLIWithBin(dcldNew,
			"query", "model", "get-model",
			"--vid", fmt.Sprintf("%d", state.VID),
			"--pid", fmt.Sprintf("%d", state.PID2),
		)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", state.VID)
		requireFieldEquals(t, out, "pid", state.PID2)
		checkResponseContains(t, out, state.ProductLabelFor1_5_1)
		checkResponseContains(t, out, state.PartNumberFor1_5_1)
	})

	MustRun(t, "VerifyModelVersionPreserved", func(t *testing.T) {
		t.Helper()
		out, err := ExecuteCLIWithBin(dcldNew,
			"query", "model", "model-version",
			"--vid", fmt.Sprintf("%d", state.VID),
			"--pid", fmt.Sprintf("%d", state.PID2),
			"--softwareVersion", fmt.Sprintf("%d", state.SoftwareVersion),
		)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", state.VID)
		requireFieldEquals(t, out, "pid", state.PID2)
		requireFieldEquals(t, out, "minApplicableSoftwareVersion",
			state.MinApplicableSoftwareVersionFor1_5_1)
		requireFieldEquals(t, out, "maxApplicableSoftwareVersion",
			state.MaxApplicableSoftwareVersionFor1_5_1)
	})

	MustRun(t, "VerifyAllModelsListings", func(t *testing.T) {
		t.Helper()
		out, err := ExecuteCLIWithBin(dcldNew, "query", "model", "all-models")
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", state.VIDFor1_5_1)
		requireFieldEquals(t, out, "pid", state.PID1For1_5_1)
		requireFieldEquals(t, out, "pid", state.PID2For1_5_1)

		out, err = ExecuteCLIWithBin(dcldNew,
			"query", "model", "all-model-versions",
			"--vid", fmt.Sprintf("%d", state.VIDFor1_5_1),
			"--pid", fmt.Sprintf("%d", state.PID1For1_5_1),
		)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", state.VIDFor1_5_1)
		requireFieldEquals(t, out, "pid", state.PID1For1_5_1)

		out, err = ExecuteCLIWithBin(dcldNew,
			"query", "model", "vendor-models",
			"--vid", fmt.Sprintf("%d", state.VIDFor1_5_1),
		)
		require.NoError(t, err)
		requireFieldEquals(t, out, "pid", state.PID1For1_5_1)
		requireFieldEquals(t, out, "pid", state.PID2For1_5_1)

		out, err = ExecuteCLIWithBin(dcldNew,
			"query", "model", "model-version",
			"--vid", fmt.Sprintf("%d", state.VIDFor1_5_1),
			"--pid", fmt.Sprintf("%d", state.PID1For1_5_1),
			"--softwareVersion", fmt.Sprintf("%d", state.SoftwareVersionFor1_5_1),
		)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", state.VIDFor1_5_1)
		requireFieldEquals(t, out, "pid", state.PID1For1_5_1)
		requireFieldEquals(t, out, "softwareVersion", state.SoftwareVersionFor1_5_1)
	})

	// ------------------------------------------------------------------
	// Post-upgrade: create a new vendor + several models against v1.5.2.
	// ------------------------------------------------------------------

	MustRun(t, "CreateVendor_1_5_2", func(t *testing.T) {
		t.Helper()
		_ = CreateAndApproveAccount(t, dcldNew, VendorAccountFor1_5_2, "Vendor",
			VIDFor1_5_2, state.Trustee1,
			[]string{state.Trustee2, state.Trustee3, state.Trustee4})
	})

	MustRun(t, "AddModelsAndVersions_1_5_2", func(t *testing.T) {
		t.Helper()
		// Add model (pid_1) with all new ICD/factory-reset fields.
		txResult, err := ExecuteTxWithBin(dcldNew,
			"tx", "model", "add-model",
			"--vid", fmt.Sprintf("%d", VIDFor1_5_2),
			"--pid", fmt.Sprintf("%d", PID1For1_5_2),
			"--deviceTypeID", fmt.Sprintf("%d", DeviceTypeIDFor1_5_2),
			"--productName", ProductNameFor1_5_2,
			"--productLabel", ProductLabelFor1_5_2,
			"--partNumber", PartNumberFor1_5_2,
			"--icdUserActiveModeTriggerHint", fmt.Sprintf("%d", ICDUserActiveModeTriggerHintFor1_5_2),
			"--icdUserActiveModeTriggerInstruction", ICDUserActiveModeTriggerInstructionFor1_5_2,
			"--factoryResetStepsHint", fmt.Sprintf("%d", FactoryResetStepsHintFor1_5_2),
			"--factoryResetStepsInstruction", FactoryResetStepsInstructionFor1_5_2,
			"--commissioningModeSecondaryStepsHint", fmt.Sprintf("%d", CommissioningModeSecondaryStepsHintFor1_5_2),
			"--from", VendorAccountFor1_5_2,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code, txResult.RawLog)

		// Add model-version for pid_1 (with specificationVersion).
		txResult, err = ExecuteTxWithBin(dcldNew,
			"tx", "model", "add-model-version",
			"--vid", fmt.Sprintf("%d", VIDFor1_5_2),
			"--pid", fmt.Sprintf("%d", PID1For1_5_2),
			"--softwareVersion", fmt.Sprintf("%d", SoftwareVersionFor1_5_2),
			"--softwareVersionString", SoftwareVersionStringFor1_5_2,
			"--cdVersionNumber", fmt.Sprintf("%d", CDVersionNumberFor1_5_2),
			"--minApplicableSoftwareVersion", fmt.Sprintf("%d", MinApplicableSoftwareVersionFor1_5_2),
			"--maxApplicableSoftwareVersion", fmt.Sprintf("%d", MaxApplicableSoftwareVersionFor1_5_2),
			"--specificationVersion", fmt.Sprintf("%d", SpecificationVersionFor1_5_2),
			"--from", VendorAccountFor1_5_2,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code, txResult.RawLog)

		// Add model pid_2 (no ICD fields, no factory reset).
		txResult, err = ExecuteTxWithBin(dcldNew,
			"tx", "model", "add-model",
			"--vid", fmt.Sprintf("%d", VIDFor1_5_2),
			"--pid", fmt.Sprintf("%d", PID2For1_5_2),
			"--deviceTypeID", fmt.Sprintf("%d", DeviceTypeIDFor1_5_2),
			"--productName", ProductNameFor1_5_2,
			"--productLabel", ProductLabelFor1_5_2,
			"--partNumber", PartNumberFor1_5_2,
			"--from", VendorAccountFor1_5_2,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code, txResult.RawLog)

		txResult, err = ExecuteTxWithBin(dcldNew,
			"tx", "model", "add-model-version",
			"--vid", fmt.Sprintf("%d", VIDFor1_5_2),
			"--pid", fmt.Sprintf("%d", PID2For1_5_2),
			"--softwareVersion", fmt.Sprintf("%d", SoftwareVersionFor1_5_2),
			"--softwareVersionString", SoftwareVersionStringFor1_5_2,
			"--cdVersionNumber", fmt.Sprintf("%d", CDVersionNumberFor1_5_2),
			"--minApplicableSoftwareVersion", fmt.Sprintf("%d", MinApplicableSoftwareVersionFor1_5_2),
			"--maxApplicableSoftwareVersion", fmt.Sprintf("%d", MaxApplicableSoftwareVersionFor1_5_2),
			"--from", VendorAccountFor1_5_2,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code, txResult.RawLog)

		// Add + immediately delete model for pid_3.
		txResult, err = ExecuteTxWithBin(dcldNew,
			"tx", "model", "add-model",
			"--vid", fmt.Sprintf("%d", VIDFor1_5_2),
			"--pid", fmt.Sprintf("%d", PID3For1_5_2),
			"--deviceTypeID", fmt.Sprintf("%d", DeviceTypeIDFor1_5_2),
			"--productName", ProductNameFor1_5_2,
			"--productLabel", ProductLabelFor1_5_2,
			"--partNumber", PartNumberFor1_5_2,
			"--from", VendorAccountFor1_5_2,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code, txResult.RawLog)

		txResult, err = ExecuteTxWithBin(dcldNew,
			"tx", "model", "add-model-version",
			"--vid", fmt.Sprintf("%d", VIDFor1_5_2),
			"--pid", fmt.Sprintf("%d", PID3For1_5_2),
			"--softwareVersion", fmt.Sprintf("%d", SoftwareVersionFor1_5_2),
			"--softwareVersionString", SoftwareVersionStringFor1_5_2,
			"--cdVersionNumber", fmt.Sprintf("%d", CDVersionNumberFor1_5_2),
			"--minApplicableSoftwareVersion", fmt.Sprintf("%d", MinApplicableSoftwareVersionFor1_5_2),
			"--maxApplicableSoftwareVersion", fmt.Sprintf("%d", MaxApplicableSoftwareVersionFor1_5_2),
			"--from", VendorAccountFor1_5_2,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code, txResult.RawLog)

		txResult, err = ExecuteTxWithBin(dcldNew,
			"tx", "model", "delete-model",
			"--vid", fmt.Sprintf("%d", VIDFor1_5_2),
			"--pid", fmt.Sprintf("%d", PID3For1_5_2),
			"--from", VendorAccountFor1_5_2,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code, txResult.RawLog)

		// Update the carry-over model from script 1.
		txResult, err = ExecuteTxWithBin(dcldNew,
			"tx", "model", "update-model",
			"--vid", fmt.Sprintf("%d", state.VID),
			"--pid", fmt.Sprintf("%d", state.PID2),
			"--productName", state.ProductName,
			"--productLabel", ProductLabelFor1_5_2,
			"--partNumber", PartNumberFor1_5_2,
			"--from", state.VendorAccount,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code, txResult.RawLog)

		txResult, err = ExecuteTxWithBin(dcldNew,
			"tx", "model", "update-model-version",
			"--vid", fmt.Sprintf("%d", state.VID),
			"--pid", fmt.Sprintf("%d", state.PID2),
			"--softwareVersion", fmt.Sprintf("%d", state.SoftwareVersion),
			"--minApplicableSoftwareVersion", fmt.Sprintf("%d", MinApplicableSoftwareVersionFor1_5_2),
			"--maxApplicableSoftwareVersion", fmt.Sprintf("%d", MaxApplicableSoftwareVersionFor1_5_2),
			"--from", state.VendorAccount,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code, txResult.RawLog)
	})

	// Seed compliance state for pid_1 — the v1.6.0 upgrade phase reads this
	// back to confirm pre-#730 records survive the schema-v1 bump intact.
	MustRun(t, "ComplianceFor1_5_2", func(t *testing.T) {
		t.Helper()
		tx, err := ExecuteTxWithBin(dcldNew,
			"tx", "compliance", "certify-model",
			"--vid", fmt.Sprintf("%d", VIDFor1_5_2),
			"--pid", fmt.Sprintf("%d", PID1For1_5_2),
			"--softwareVersion", fmt.Sprintf("%d", SoftwareVersionFor1_5_2),
			"--softwareVersionString", SoftwareVersionStringFor1_5_2,
			"--certificationType", CertificationTypeFor1_5_2,
			"--certificationDate", CertificationDateFor1_5_2,
			"--cdCertificateId", CDCertificateIDFor1_5_2,
			"--cdVersionNumber", fmt.Sprintf("%d", CDVersionNumberFor1_5_2),
			"--from", CertificationCenterAccountFor1_2,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)
	})

	MustRun(t, "VerifyNewModels_1_5_2", func(t *testing.T) {
		t.Helper()
		out, err := ExecuteCLIWithBin(dcldNew,
			"query", "model", "get-model",
			"--vid", fmt.Sprintf("%d", VIDFor1_5_2),
			"--pid", fmt.Sprintf("%d", PID1For1_5_2),
		)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", VIDFor1_5_2)
		requireFieldEquals(t, out, "pid", PID1For1_5_2)
		checkResponseContains(t, out, ProductLabelFor1_5_2)
		requireFieldEquals(t, out, "icdUserActiveModeTriggerHint", ICDUserActiveModeTriggerHintFor1_5_2)
		checkResponseContains(t, out, ICDUserActiveModeTriggerInstructionFor1_5_2)
		requireFieldEquals(t, out, "factoryResetStepsHint", FactoryResetStepsHintFor1_5_2)
		checkResponseContains(t, out, FactoryResetStepsInstructionFor1_5_2)
		requireFieldEquals(t, out, "commissioningModeSecondaryStepsHint", CommissioningModeSecondaryStepsHintFor1_5_2)

		out, err = ExecuteCLIWithBin(dcldNew,
			"query", "model", "get-model",
			"--vid", fmt.Sprintf("%d", VIDFor1_5_2),
			"--pid", fmt.Sprintf("%d", PID2For1_5_2),
		)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", VIDFor1_5_2)
		requireFieldEquals(t, out, "pid", PID2For1_5_2)
		checkResponseContains(t, out, ProductLabelFor1_5_2)
		// Migration default for this field is 4.
		requireFieldEquals(t, out, "commissioningModeSecondaryStepsHint", 4)

		// Updated carry-over model.
		out, err = ExecuteCLIWithBin(dcldNew,
			"query", "model", "get-model",
			"--vid", fmt.Sprintf("%d", state.VID),
			"--pid", fmt.Sprintf("%d", state.PID2),
		)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", state.VID)
		requireFieldEquals(t, out, "pid", state.PID2)
		checkResponseContains(t, out, ProductLabelFor1_5_2)
		checkResponseContains(t, out, PartNumberFor1_5_2)

		// Model version specificationVersion check.
		out, err = ExecuteCLIWithBin(dcldNew,
			"query", "model", "model-version",
			"--vid", fmt.Sprintf("%d", VIDFor1_5_2),
			"--pid", fmt.Sprintf("%d", PID1For1_5_2),
			"--softwareVersion", fmt.Sprintf("%d", SoftwareVersionFor1_5_2),
		)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", VIDFor1_5_2)
		requireFieldEquals(t, out, "pid", PID1For1_5_2)
		requireFieldEquals(t, out, "softwareVersion", SoftwareVersionFor1_5_2)
		requireFieldEquals(t, out, "specificationVersion", SpecificationVersionFor1_5_2)
	})
}
