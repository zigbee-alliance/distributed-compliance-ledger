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

	// Re-affirm sync mode — v1.4+ binaries no longer accept `block`.
	_ = SetBroadcastMode(dcldNew, "sync")

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
		out, err := QueryGetModel(dcldNew, state.VIDFor1_5_1, state.PID1For1_5_1)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", state.VIDFor1_5_1)
		requireFieldEquals(t, out, "pid", state.PID1For1_5_1)
		checkResponseContains(t, out, state.ProductLabelFor1_5_1)
		requireFieldEquals(t, out, "commissioningModeSecondaryStepsHint",
			state.CommissioningModeSecondaryStepsHintFor1_5_1)

		out, err = QueryGetModel(dcldNew, state.VIDFor1_5_1, state.PID2For1_5_1)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", state.VIDFor1_5_1)
		requireFieldEquals(t, out, "pid", state.PID2For1_5_1)
		checkResponseContains(t, out, state.ProductLabelFor1_5_1)
		// Migration default for this field is 4.
		requireFieldEquals(t, out, "commissioningModeSecondaryStepsHint", 4)
	})

	MustRun(t, "VerifyUpdatedModelFromScript1", func(t *testing.T) {
		t.Helper()
		out, err := QueryGetModel(dcldNew, state.VID, state.PID2)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", state.VID)
		requireFieldEquals(t, out, "pid", state.PID2)
		checkResponseContains(t, out, state.ProductLabelFor1_5_1)
		checkResponseContains(t, out, state.PartNumberFor1_5_1)
	})

	MustRun(t, "VerifyModelVersionPreserved", func(t *testing.T) {
		t.Helper()
		out, err := QueryModelVersion(dcldNew, state.VID, state.PID2, state.SoftwareVersion)
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
		out, err := QueryAllModels(dcldNew)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", state.VIDFor1_5_1)
		requireFieldEquals(t, out, "pid", state.PID1For1_5_1)
		requireFieldEquals(t, out, "pid", state.PID2For1_5_1)

		out, err = QueryAllModelVersions(dcldNew, state.VIDFor1_5_1, state.PID1For1_5_1)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", state.VIDFor1_5_1)
		requireFieldEquals(t, out, "pid", state.PID1For1_5_1)

		out, err = QueryVendorModels(dcldNew, state.VIDFor1_5_1)
		require.NoError(t, err)
		requireFieldEquals(t, out, "pid", state.PID1For1_5_1)
		requireFieldEquals(t, out, "pid", state.PID2For1_5_1)

		out, err = QueryModelVersion(dcldNew, state.VIDFor1_5_1, state.PID1For1_5_1, state.SoftwareVersionFor1_5_1)
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
		txResult, err := AddModel(dcldNew, AddModelArgs{
			VID: VIDFor1_5_2, PID: PID1For1_5_2, DeviceTypeID: DeviceTypeIDFor1_5_2,
			ProductName: ProductNameFor1_5_2, ProductLabel: ProductLabelFor1_5_2, PartNumber: PartNumberFor1_5_2,
			IcdUserActiveModeTriggerHint:        ICDUserActiveModeTriggerHintFor1_5_2,
			IcdUserActiveModeTriggerInstruction: ICDUserActiveModeTriggerInstructionFor1_5_2,
			FactoryResetStepsHint:               FactoryResetStepsHintFor1_5_2,
			FactoryResetStepsInstruction:        FactoryResetStepsInstructionFor1_5_2,
			CommissioningModeSecondaryStepsHint: CommissioningModeSecondaryStepsHintFor1_5_2,
			From:                                VendorAccountFor1_5_2,
		})
		requireTxSuccess(t, txResult, err)

		// Add model-version for pid_1 (with specificationVersion).
		txResult, err = AddModelVersion(dcldNew, AddModelVersionArgs{
			VID: VIDFor1_5_2, PID: PID1For1_5_2,
			SoftwareVersion: SoftwareVersionFor1_5_2, SoftwareVersionString: SoftwareVersionStringFor1_5_2,
			CDVersionNumber:              CDVersionNumberFor1_5_2,
			MinApplicableSoftwareVersion: MinApplicableSoftwareVersionFor1_5_2,
			MaxApplicableSoftwareVersion: MaxApplicableSoftwareVersionFor1_5_2,
			SpecificationVersion:         SpecificationVersionFor1_5_2,
			From:                         VendorAccountFor1_5_2,
		})
		requireTxSuccess(t, txResult, err)

		// Add model pid_2 (no ICD fields, no factory reset).
		txResult, err = AddModel(dcldNew, AddModelArgs{
			VID: VIDFor1_5_2, PID: PID2For1_5_2, DeviceTypeID: DeviceTypeIDFor1_5_2,
			ProductName: ProductNameFor1_5_2, ProductLabel: ProductLabelFor1_5_2, PartNumber: PartNumberFor1_5_2,
			From: VendorAccountFor1_5_2,
		})
		requireTxSuccess(t, txResult, err)

		txResult, err = AddModelVersion(dcldNew, AddModelVersionArgs{
			VID: VIDFor1_5_2, PID: PID2For1_5_2,
			SoftwareVersion: SoftwareVersionFor1_5_2, SoftwareVersionString: SoftwareVersionStringFor1_5_2,
			CDVersionNumber:              CDVersionNumberFor1_5_2,
			MinApplicableSoftwareVersion: MinApplicableSoftwareVersionFor1_5_2,
			MaxApplicableSoftwareVersion: MaxApplicableSoftwareVersionFor1_5_2,
			From:                         VendorAccountFor1_5_2,
		})
		requireTxSuccess(t, txResult, err)

		// Add + immediately delete model for pid_3.
		txResult, err = AddModel(dcldNew, AddModelArgs{
			VID: VIDFor1_5_2, PID: PID3For1_5_2, DeviceTypeID: DeviceTypeIDFor1_5_2,
			ProductName: ProductNameFor1_5_2, ProductLabel: ProductLabelFor1_5_2, PartNumber: PartNumberFor1_5_2,
			From: VendorAccountFor1_5_2,
		})
		requireTxSuccess(t, txResult, err)

		txResult, err = AddModelVersion(dcldNew, AddModelVersionArgs{
			VID: VIDFor1_5_2, PID: PID3For1_5_2,
			SoftwareVersion: SoftwareVersionFor1_5_2, SoftwareVersionString: SoftwareVersionStringFor1_5_2,
			CDVersionNumber:              CDVersionNumberFor1_5_2,
			MinApplicableSoftwareVersion: MinApplicableSoftwareVersionFor1_5_2,
			MaxApplicableSoftwareVersion: MaxApplicableSoftwareVersionFor1_5_2,
			From:                         VendorAccountFor1_5_2,
		})
		requireTxSuccess(t, txResult, err)

		txResult, err = DeleteModel(dcldNew, VIDFor1_5_2, PID3For1_5_2, VendorAccountFor1_5_2)
		requireTxSuccess(t, txResult, err)

		// Update the carry-over model from script 1.
		txResult, err = UpdateModel(dcldNew, UpdateModelArgs{
			VID: state.VID, PID: state.PID2,
			ProductName: state.ProductName, ProductLabel: ProductLabelFor1_5_2, PartNumber: PartNumberFor1_5_2,
			From: state.VendorAccount,
		})
		requireTxSuccess(t, txResult, err)

		txResult, err = UpdateModelVersion(dcldNew, UpdateModelVersionArgs{
			VID: state.VID, PID: state.PID2, SoftwareVersion: state.SoftwareVersion,
			MinApplicableSoftwareVersion: MinApplicableSoftwareVersionFor1_5_2,
			MaxApplicableSoftwareVersion: MaxApplicableSoftwareVersionFor1_5_2,
			From:                         state.VendorAccount,
		})
		requireTxSuccess(t, txResult, err)
	})

	// Seed compliance state for pid_1 — the v1.6.0 upgrade phase reads this
	// back to confirm pre-#730 records survive the schema-v1 bump intact.
	MustRun(t, "ComplianceFor1_5_2", func(t *testing.T) {
		t.Helper()
		tx, err := CertifyModel(dcldNew, CertifyModelArgs{
			VID: VIDFor1_5_2, PID: PID1For1_5_2,
			SoftwareVersion: SoftwareVersionFor1_5_2, SoftwareVersionString: SoftwareVersionStringFor1_5_2,
			CertificationType: CertificationTypeFor1_5_2, CertificationDate: CertificationDateFor1_5_2,
			CDCertificateID: CDCertificateIDFor1_5_2, CDVersionNumber: CDVersionNumberFor1_5_2,
			From: CertificationCenterAccountFor1_2,
		})
		requireTxSuccess(t, tx, err)
	})

	MustRun(t, "VerifyNewModels_1_5_2", func(t *testing.T) {
		t.Helper()
		out, err := QueryGetModel(dcldNew, VIDFor1_5_2, PID1For1_5_2)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", VIDFor1_5_2)
		requireFieldEquals(t, out, "pid", PID1For1_5_2)
		checkResponseContains(t, out, ProductLabelFor1_5_2)
		requireFieldEquals(t, out, "icdUserActiveModeTriggerHint", ICDUserActiveModeTriggerHintFor1_5_2)
		checkResponseContains(t, out, ICDUserActiveModeTriggerInstructionFor1_5_2)
		requireFieldEquals(t, out, "factoryResetStepsHint", FactoryResetStepsHintFor1_5_2)
		checkResponseContains(t, out, FactoryResetStepsInstructionFor1_5_2)
		requireFieldEquals(t, out, "commissioningModeSecondaryStepsHint", CommissioningModeSecondaryStepsHintFor1_5_2)

		out, err = QueryGetModel(dcldNew, VIDFor1_5_2, PID2For1_5_2)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", VIDFor1_5_2)
		requireFieldEquals(t, out, "pid", PID2For1_5_2)
		checkResponseContains(t, out, ProductLabelFor1_5_2)
		// Migration default for this field is 4.
		requireFieldEquals(t, out, "commissioningModeSecondaryStepsHint", 4)

		// Updated carry-over model.
		out, err = QueryGetModel(dcldNew, state.VID, state.PID2)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", state.VID)
		requireFieldEquals(t, out, "pid", state.PID2)
		checkResponseContains(t, out, ProductLabelFor1_5_2)
		checkResponseContains(t, out, PartNumberFor1_5_2)

		// Model version specificationVersion check.
		out, err = QueryModelVersion(dcldNew, VIDFor1_5_2, PID1For1_5_2, SoftwareVersionFor1_5_2)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", VIDFor1_5_2)
		requireFieldEquals(t, out, "pid", PID1For1_5_2)
		requireFieldEquals(t, out, "softwareVersion", SoftwareVersionFor1_5_2)
		requireFieldEquals(t, out, "specificationVersion", SpecificationVersionFor1_5_2)
	})
}
