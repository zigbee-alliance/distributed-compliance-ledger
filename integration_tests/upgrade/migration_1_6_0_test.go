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
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// runUpgrade152To160 runs the v1.5.2 → v1.6.0 cosmovisor upgrade and
// verifies Issue #593 ghost-cleanup: a model version deleted on v1.5.2
// becomes properly inaccessible after the v1.6.0 migration, and the
// underlying model can now be deleted cleanly.
//
// Assumes the chain is at v1.5.2 with state from phases 01-08.
//
//nolint:funlen
func runUpgrade152To160(t *testing.T, state *UpgradeTestState) {
	t.Helper()

	dcldOld, err := EnsureBinary("1.5.2")
	require.NoError(t, err, "fetch dcld v1.5.2")
	dcldNew, err := EnsureBinary(BinaryVersionV1_6_0)
	require.NoError(t, err, "fetch dcld %s", BinaryVersionV1_6_0)

	// ------------------------------------------------------------------
	// ISSUE #593: pre-upgrade, delete one model-version and verify the
	// "ghost" semantics on the old binary.
	// ------------------------------------------------------------------
	MustRun(t, "Issue593_PreUpgradeDelete", func(t *testing.T) {
		t.Helper()
		tx, err := DeleteModelVersion(dcldOld, state.VIDFor1_6_0FromScript5, state.PID3For1_6_0FromScript5, state.SoftwareVersion2For1_6_0FromScript5, state.VendorAccountFor1_2)
		requireTxSuccess(t, tx, err)

		out, err := QueryAllModelVersions(dcldOld, state.VIDFor1_6_0FromScript5, state.PID3For1_6_0FromScript5)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", state.VIDFor1_6_0FromScript5)
		requireFieldEquals(t, out, "pid", state.PID3For1_6_0FromScript5)
		require.True(t, strings.Contains(string(out),
			fmt.Sprintf("%d", state.SoftwareVersion1For1_6_0FromScript5)),
			"expected sv1 still present, got: %s", string(out))
		require.False(t, strings.Contains(string(out),
			fmt.Sprintf("%d", state.SoftwareVersion2For1_6_0FromScript5)),
			"expected sv2 absent post-delete, got: %s", string(out))
	})

	// Re-affirm sync mode — v1.4+ binaries no longer accept `block`.
	_ = SetBroadcastMode(dcldNew, "sync")

	step := SoftwareUpgradeStep{
		PlanName:         PlanNameV1_6_0,
		BinaryVersionNew: BinaryVersionV1_6_0,
		Checksum:         UpgradeChecksumV1_6_0,
		DcldOldBin:       dcldOld,
		DcldNewBin:       dcldNew,
		Trustees: []string{
			state.Trustee1, state.Trustee2, state.Trustee3, state.Trustee4,
		},
	}
	step.Run(t)

	// ------------------------------------------------------------------
	// ISSUE #593: post-upgrade ghost cleanup verification.
	// ------------------------------------------------------------------
	MustRun(t, "Issue593_GhostModelVersionsRemoved", func(t *testing.T) {
		t.Helper()
		out, err := QueryAllModelVersions(dcldNew, state.VIDFor1_6_0FromScript5, state.PID3For1_6_0FromScript5)
		require.NoError(t, err)
		require.True(t, strings.Contains(string(out), "Not Found"),
			"expected 'Not Found' for ghost-cleaned versions, got: %s", string(out))

		// And we can now delete the model itself.
		tx, err := DeleteModel(dcldNew, state.VIDFor1_6_0FromScript5, state.PID3For1_6_0FromScript5, state.VendorAccountFor1_2)
		requireTxSuccess(t, tx, err)
	})

	// ------------------------------------------------------------------
	// After upgrade to v1.6.0, the compliance record certified in 1.5.2
	// must remain queryable. The schema-v1 bump (#730) only tightens
	// write-path constraints — pre-existing stored records keep their values.
	// ------------------------------------------------------------------
	MustRun(t, "ComplianceCarryoverFrom_1_5_2", func(t *testing.T) {
		t.Helper()
		out, err := QueryComplianceInfo(dcldNew, VIDFor1_5_2, PID1For1_5_2, SoftwareVersionFor1_5_2, CertificationTypeFor1_5_2)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", VIDFor1_5_2)
		requireFieldEquals(t, out, "pid", PID1For1_5_2)
		requireFieldEquals(t, out, "softwareVersion", SoftwareVersionFor1_5_2)
		checkResponseContains(t, out, CertificationTypeFor1_5_2)
		checkResponseContains(t, out, CDCertificateIDFor1_5_2)

		out, err = QueryCertifiedModel(dcldNew, VIDFor1_5_2, PID1For1_5_2, SoftwareVersionFor1_5_2, CertificationTypeFor1_5_2)
		require.NoError(t, err)
		requireFieldEquals(t, out, "value", true)
		requireFieldEquals(t, out, "vid", VIDFor1_5_2)
		requireFieldEquals(t, out, "pid", PID1For1_5_2)

		out, err = QueryDeviceSoftwareCompliance(dcldNew, CDCertificateIDFor1_5_2)
		require.NoError(t, err)
		checkResponseContains(t, out, CDCertificateIDFor1_5_2)
		requireFieldEquals(t, out, "vid", VIDFor1_5_2)
		requireFieldEquals(t, out, "pid", PID1For1_5_2)
	})

	// ------------------------------------------------------------------
	// Verify carry-over data from the v1.5.2 era is intact.
	// ------------------------------------------------------------------
	MustRun(t, "VerifyPreserved_1_5_2_Models", func(t *testing.T) {
		t.Helper()
		out, err := QueryGetModel(dcldNew, VIDFor1_5_2, PID1For1_5_2)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", VIDFor1_5_2)
		requireFieldEquals(t, out, "pid", PID1For1_5_2)
		checkResponseContains(t, out, ProductLabelFor1_5_2)

		out, err = QueryGetModel(dcldNew, VIDFor1_5_2, PID2For1_5_2)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", VIDFor1_5_2)
		requireFieldEquals(t, out, "pid", PID2For1_5_2)
		checkResponseContains(t, out, ProductLabelFor1_5_2)

		out, err = QueryGetModel(dcldNew, state.VID, state.PID2)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", state.VID)
		requireFieldEquals(t, out, "pid", state.PID2)
		checkResponseContains(t, out, ProductLabelFor1_5_2)
		checkResponseContains(t, out, PartNumberFor1_5_2)

		out, err = QueryModelVersion(dcldNew, state.VID, state.PID2, state.SoftwareVersion)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", state.VID)
		requireFieldEquals(t, out, "pid", state.PID2)
		requireFieldEquals(t, out, "minApplicableSoftwareVersion", MinApplicableSoftwareVersionFor1_5_2)
		requireFieldEquals(t, out, "maxApplicableSoftwareVersion", MaxApplicableSoftwareVersionFor1_5_2)

		// Bulk model listings (all-models + vendor-models).
		out, err = QueryAllModels(dcldNew)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", VIDFor1_5_2)
		requireFieldEquals(t, out, "pid", PID1For1_5_2)
		requireFieldEquals(t, out, "pid", PID2For1_5_2)

		out, err = QueryVendorModels(dcldNew, VIDFor1_5_2)
		require.NoError(t, err)
		requireFieldEquals(t, out, "pid", PID1For1_5_2)
		requireFieldEquals(t, out, "pid", PID2For1_5_2)
	})

	// ------------------------------------------------------------------
	// Post-upgrade: create a v1.6.0-era vendor and models.
	// ------------------------------------------------------------------
	MustRun(t, "CreateVendor_1_6_0", func(t *testing.T) {
		t.Helper()
		_ = CreateAndApproveAccount(t, dcldNew, VendorAccountFor1_6_0, "Vendor",
			VIDFor1_6_0, state.Trustee1,
			[]string{state.Trustee2, state.Trustee3, state.Trustee4})
	})

	MustRun(t, "AddModelsAndVersions_1_6_0", func(t *testing.T) {
		t.Helper()
		tx, err := AddModel(dcldNew, AddModelArgs{VID: VIDFor1_6_0, PID: PID1For1_6_0, DeviceTypeID: DeviceTypeIDFor1_6_0, ProductName: ProductNameFor1_6_0, ProductLabel: ProductLabelFor1_6_0, PartNumber: PartNumberFor1_6_0, IcdUserActiveModeTriggerHint: ICDUserActiveModeTriggerHintFor1_6_0, IcdUserActiveModeTriggerInstruction: ICDUserActiveModeTriggerInstructionFor1_6_0, FactoryResetStepsHint: FactoryResetStepsHintFor1_6_0, FactoryResetStepsInstruction: FactoryResetStepsInstructionFor1_6_0, CommissioningCustomFlow: intPtr(CommissioningCustomFlow), From: VendorAccountFor1_6_0})
		requireTxSuccess(t, tx, err)

		tx, err = AddModelVersion(dcldNew, AddModelVersionArgs{VID: VIDFor1_6_0, PID: PID1For1_6_0, SoftwareVersion: SoftwareVersionFor1_6_0, SoftwareVersionString: SoftwareVersionStringFor1_6_0, CDVersionNumber: CDVersionNumberFor1_6_0, MinApplicableSoftwareVersion: MinApplicableSoftwareVersionFor1_6_0, MaxApplicableSoftwareVersion: MaxApplicableSoftwareVersionFor1_6_0, SpecificationVersion: SpecificationVersionFor1_6_0, From: VendorAccountFor1_6_0})
		requireTxSuccess(t, tx, err)

		tx, err = AddModel(dcldNew, AddModelArgs{VID: VIDFor1_6_0, PID: PID2For1_6_0, DeviceTypeID: DeviceTypeIDFor1_6_0, ProductName: ProductNameFor1_6_0, ProductLabel: ProductLabelFor1_6_0, PartNumber: PartNumberFor1_6_0, CommissioningCustomFlow: intPtr(CommissioningCustomFlow), From: VendorAccountFor1_6_0})
		requireTxSuccess(t, tx, err)

		tx, err = AddModelVersion(dcldNew, AddModelVersionArgs{VID: VIDFor1_6_0, PID: PID2For1_6_0, SoftwareVersion: SoftwareVersionFor1_6_0, SoftwareVersionString: SoftwareVersionStringFor1_6_0, CDVersionNumber: CDVersionNumberFor1_6_0, MinApplicableSoftwareVersion: MinApplicableSoftwareVersionFor1_6_0, MaxApplicableSoftwareVersion: MaxApplicableSoftwareVersionFor1_6_0, From: VendorAccountFor1_6_0})
		requireTxSuccess(t, tx, err)

		// Update a v1.5.2-era model through the v1.5.2 vendor account.
		tx, err = UpdateModel(dcldNew, UpdateModelArgs{VID: VIDFor1_5_2, PID: PID2For1_5_2, ProductName: ProductNameFor1_6_0, ProductLabel: ProductLabelFor1_6_0, PartNumber: PartNumberFor1_6_0, From: VendorAccountFor1_5_2})
		requireTxSuccess(t, tx, err)
	})

	// Add a model with discoveryCapabilitiesBitmask=20 — allowed range
	// widened from 0-14 to 0-30 in v1.6.0.
	MustRun(t, "AddModelWithWidenedBitmask_1_6_0", func(t *testing.T) {
		t.Helper()
		tx, err := AddModel(dcldNew, AddModelArgs{VID: VIDFor1_6_0, PID: PIDWidenedBitmaskFor1_6_0, DeviceTypeID: DeviceTypeIDFor1_6_0, ProductName: ProductNameFor1_6_0, ProductLabel: ProductLabelFor1_6_0, PartNumber: PartNumberFor1_6_0, CommissioningCustomFlow: intPtr(CommissioningCustomFlowFor1_6), DiscoveryCapabilitiesBitmask: DiscoveryCapabilitiesBitmask, From: VendorAccountFor1_6_0})
		requireTxSuccess(t, tx, err)
	})

	// Compliance writes after upgrade — schemaVersion=1 (default per #730),
	// specificationVersion now required on every write. cDCertificateId is
	// reused across schema-v0 (carry-over) and schema-v1 records so the
	// device-software-compliance index covers both eras.
	MustRun(t, "ComplianceWrites_SchemaV1_1_6_0", func(t *testing.T) {
		t.Helper()
		tx, err := CertifyModel(dcldNew, CertifyModelArgs{VID: VIDFor1_6_0, PID: PID1For1_6_0, SoftwareVersion: SoftwareVersionFor1_6_0, SoftwareVersionString: SoftwareVersionStringFor1_6_0, CDVersionNumber: CDVersionNumberFor1_6_0, CertificationType: CertificationTypeFor1_6_0, CertificationDate: CertificationDateFor1_6_0, SpecificationVersion: SpecificationVersionFor1_6_0, CDCertificateID: CDCertificateIDFor1_5_2, SchemaVersion: "1", From: CertificationCenterAccountFor1_2})
		requireTxSuccess(t, tx, err)

		tx, err = ProvisionModel(dcldNew, ProvisionModelArgs{VID: VIDFor1_6_0, PID: PID2For1_6_0, SoftwareVersion: SoftwareVersionFor1_6_0, SoftwareVersionString: SoftwareVersionStringFor1_6_0, CDVersionNumber: CDVersionNumberFor1_6_0, CertificationType: CertificationTypeFor1_6_0, ProvisionalDate: ProvisionalDateFor1_6_0, SpecificationVersion: SpecificationVersionFor1_6_0, CDCertificateID: CDCertificateIDFor1_5_2, SchemaVersion: "1", From: CertificationCenterAccountFor1_2})
		requireTxSuccess(t, tx, err)
	})

	MustRun(t, "VerifyNewModels_1_6_0", func(t *testing.T) {
		t.Helper()
		out, err := QueryGetModel(dcldNew, VIDFor1_6_0, PID1For1_6_0)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", VIDFor1_6_0)
		requireFieldEquals(t, out, "pid", PID1For1_6_0)
		checkResponseContains(t, out, ProductLabelFor1_6_0)

		out, err = QueryGetModel(dcldNew, VIDFor1_6_0, PID2For1_6_0)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", VIDFor1_6_0)
		requireFieldEquals(t, out, "pid", PID2For1_6_0)
		checkResponseContains(t, out, ProductLabelFor1_6_0)

		// Verify the 1.5.2 model now has 1.6.0 fields after the update.
		out, err = QueryGetModel(dcldNew, VIDFor1_5_2, PID2For1_5_2)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", VIDFor1_5_2)
		requireFieldEquals(t, out, "pid", PID2For1_5_2)
		checkResponseContains(t, out, ProductLabelFor1_6_0)
		checkResponseContains(t, out, PartNumberFor1_6_0)

		// Model versions / specificationVersion.
		out, err = QueryModelVersion(dcldNew, VIDFor1_6_0, PID1For1_6_0, SoftwareVersionFor1_6_0)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", VIDFor1_6_0)
		requireFieldEquals(t, out, "pid", PID1For1_6_0)
		requireFieldEquals(t, out, "softwareVersion", SoftwareVersionFor1_6_0)
		requireFieldEquals(t, out, "specificationVersion", SpecificationVersionFor1_6_0)
	})

	// Verify the widened-bitmask model landed with its requested value.
	MustRun(t, "VerifyWidenedBitmaskModel_1_6_0", func(t *testing.T) {
		t.Helper()
		out, err := QueryGetModel(dcldNew, VIDFor1_6_0, PIDWidenedBitmaskFor1_6_0)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", VIDFor1_6_0)
		requireFieldEquals(t, out, "pid", PIDWidenedBitmaskFor1_6_0)
		requireFieldEquals(t, out, "discoveryCapabilitiesBitmask", DiscoveryCapabilitiesBitmask)
	})

	// Verify the schema-v1 compliance writes landed with specificationVersion
	// now persisted on ComplianceInfo.
	MustRun(t, "VerifyComplianceWrites_SchemaV1_1_6_0", func(t *testing.T) {
		t.Helper()
		out, err := QueryCertifiedModel(dcldNew, VIDFor1_6_0, PID1For1_6_0, SoftwareVersionFor1_6_0, CertificationTypeFor1_6_0)
		require.NoError(t, err)
		requireFieldEquals(t, out, "value", true)
		requireFieldEquals(t, out, "vid", VIDFor1_6_0)

		out, err = QueryComplianceInfo(dcldNew, VIDFor1_6_0, PID1For1_6_0, SoftwareVersionFor1_6_0, CertificationTypeFor1_6_0)
		require.NoError(t, err)
		requireFieldEquals(t, out, "schemaVersion", 1)
		requireFieldEquals(t, out, "specificationVersion", SpecificationVersionFor1_6_0)

		out, err = QueryProvisionalModel(dcldNew, VIDFor1_6_0, PID2For1_6_0, SoftwareVersionFor1_6_0, CertificationTypeFor1_6_0)
		require.NoError(t, err)
		requireFieldEquals(t, out, "value", true)
		requireFieldEquals(t, out, "vid", VIDFor1_6_0)
	})
}
