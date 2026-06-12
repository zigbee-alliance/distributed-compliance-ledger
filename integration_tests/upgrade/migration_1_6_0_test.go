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
		tx, err := ExecuteTxWithBin(dcldOld,
			"tx", "model", "delete-model-version",
			"--vid", fmt.Sprintf("%d", state.VIDFor1_6_0FromScript5),
			"--pid", fmt.Sprintf("%d", state.PID3For1_6_0FromScript5),
			"--softwareVersion", fmt.Sprintf("%d", state.SoftwareVersion2For1_6_0FromScript5),
			"--from", state.VendorAccountFor1_2,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)

		out, err := ExecuteCLIWithBin(dcldOld,
			"query", "model", "all-model-versions",
			"--vid", fmt.Sprintf("%d", state.VIDFor1_6_0FromScript5),
			"--pid", fmt.Sprintf("%d", state.PID3For1_6_0FromScript5),
		)
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

	// v1.4+ binaries no longer accept --broadcast-mode block.
	_, _ = ExecuteCLIWithBin(dcldNew, "config", "broadcast-mode", "sync")

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
		out, err := ExecuteCLIWithBin(dcldNew,
			"query", "model", "all-model-versions",
			"--vid", fmt.Sprintf("%d", state.VIDFor1_6_0FromScript5),
			"--pid", fmt.Sprintf("%d", state.PID3For1_6_0FromScript5),
		)
		require.NoError(t, err)
		require.True(t, strings.Contains(string(out), "Not Found"),
			"expected 'Not Found' for ghost-cleaned versions, got: %s", string(out))

		// And we can now delete the model itself.
		tx, err := ExecuteTxWithBin(dcldNew,
			"tx", "model", "delete-model",
			"--vid", fmt.Sprintf("%d", state.VIDFor1_6_0FromScript5),
			"--pid", fmt.Sprintf("%d", state.PID3For1_6_0FromScript5),
			"--from", state.VendorAccountFor1_2,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)
	})

	// ------------------------------------------------------------------
	// After upgrade to v1.6.0, the compliance record certified in 1.5.2
	// must remain queryable. The schema-v1 bump (#730) only tightens
	// write-path constraints — pre-existing stored records keep their values.
	// ------------------------------------------------------------------
	MustRun(t, "ComplianceCarryoverFrom_1_5_2", func(t *testing.T) {
		t.Helper()
		out, err := ExecuteCLIWithBin(dcldNew,
			"query", "compliance", "compliance-info",
			"--vid", fmt.Sprintf("%d", VIDFor1_5_2),
			"--pid", fmt.Sprintf("%d", PID1For1_5_2),
			"--softwareVersion", fmt.Sprintf("%d", SoftwareVersionFor1_5_2),
			"--certificationType", CertificationTypeFor1_5_2,
		)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", VIDFor1_5_2)
		requireFieldEquals(t, out, "pid", PID1For1_5_2)
		requireFieldEquals(t, out, "softwareVersion", SoftwareVersionFor1_5_2)
		checkResponseContains(t, out, CertificationTypeFor1_5_2)
		checkResponseContains(t, out, CDCertificateIDFor1_5_2)

		out, err = ExecuteCLIWithBin(dcldNew,
			"query", "compliance", "certified-model",
			"--vid", fmt.Sprintf("%d", VIDFor1_5_2),
			"--pid", fmt.Sprintf("%d", PID1For1_5_2),
			"--softwareVersion", fmt.Sprintf("%d", SoftwareVersionFor1_5_2),
			"--certificationType", CertificationTypeFor1_5_2,
		)
		require.NoError(t, err)
		requireFieldEquals(t, out, "value", true)
		requireFieldEquals(t, out, "vid", VIDFor1_5_2)
		requireFieldEquals(t, out, "pid", PID1For1_5_2)

		out, err = ExecuteCLIWithBin(dcldNew,
			"query", "compliance", "device-software-compliance",
			"--cdCertificateId", CDCertificateIDFor1_5_2,
		)
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
		out, err := ExecuteCLIWithBin(dcldNew,
			"query", "model", "get-model",
			"--vid", fmt.Sprintf("%d", VIDFor1_5_2),
			"--pid", fmt.Sprintf("%d", PID1For1_5_2),
		)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", VIDFor1_5_2)
		requireFieldEquals(t, out, "pid", PID1For1_5_2)
		checkResponseContains(t, out, ProductLabelFor1_5_2)

		out, err = ExecuteCLIWithBin(dcldNew,
			"query", "model", "get-model",
			"--vid", fmt.Sprintf("%d", VIDFor1_5_2),
			"--pid", fmt.Sprintf("%d", PID2For1_5_2),
		)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", VIDFor1_5_2)
		requireFieldEquals(t, out, "pid", PID2For1_5_2)
		checkResponseContains(t, out, ProductLabelFor1_5_2)

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

		out, err = ExecuteCLIWithBin(dcldNew,
			"query", "model", "model-version",
			"--vid", fmt.Sprintf("%d", state.VID),
			"--pid", fmt.Sprintf("%d", state.PID2),
			"--softwareVersion", fmt.Sprintf("%d", state.SoftwareVersion),
		)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", state.VID)
		requireFieldEquals(t, out, "pid", state.PID2)
		requireFieldEquals(t, out, "minApplicableSoftwareVersion", MinApplicableSoftwareVersionFor1_5_2)
		requireFieldEquals(t, out, "maxApplicableSoftwareVersion", MaxApplicableSoftwareVersionFor1_5_2)

		// Bulk model listings (all-models + vendor-models).
		out, err = ExecuteCLIWithBin(dcldNew, "query", "model", "all-models")
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", VIDFor1_5_2)
		requireFieldEquals(t, out, "pid", PID1For1_5_2)
		requireFieldEquals(t, out, "pid", PID2For1_5_2)

		out, err = ExecuteCLIWithBin(dcldNew,
			"query", "model", "vendor-models",
			"--vid", fmt.Sprintf("%d", VIDFor1_5_2),
		)
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
		tx, err := ExecuteTxWithBin(dcldNew,
			"tx", "model", "add-model",
			"--vid", fmt.Sprintf("%d", VIDFor1_6_0),
			"--pid", fmt.Sprintf("%d", PID1For1_6_0),
			"--deviceTypeID", fmt.Sprintf("%d", DeviceTypeIDFor1_6_0),
			"--productName", ProductNameFor1_6_0,
			"--productLabel", ProductLabelFor1_6_0,
			"--partNumber", PartNumberFor1_6_0,
			"--icdUserActiveModeTriggerHint", fmt.Sprintf("%d", ICDUserActiveModeTriggerHintFor1_6_0),
			"--icdUserActiveModeTriggerInstruction", ICDUserActiveModeTriggerInstructionFor1_6_0,
			"--factoryResetStepsHint", fmt.Sprintf("%d", FactoryResetStepsHintFor1_6_0),
			"--factoryResetStepsInstruction", FactoryResetStepsInstructionFor1_6_0,
			"--commissioningModeSecondaryStepsHint", fmt.Sprintf("%d", CommissioningModeSecondaryStepsHintFor1_6_0),
			"--from", VendorAccountFor1_6_0,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)

		tx, err = ExecuteTxWithBin(dcldNew,
			"tx", "model", "add-model-version",
			"--vid", fmt.Sprintf("%d", VIDFor1_6_0),
			"--pid", fmt.Sprintf("%d", PID1For1_6_0),
			"--softwareVersion", fmt.Sprintf("%d", SoftwareVersionFor1_6_0),
			"--softwareVersionString", SoftwareVersionStringFor1_6_0,
			"--cdVersionNumber", fmt.Sprintf("%d", CDVersionNumberFor1_6_0),
			"--minApplicableSoftwareVersion", fmt.Sprintf("%d", MinApplicableSoftwareVersionFor1_6_0),
			"--maxApplicableSoftwareVersion", fmt.Sprintf("%d", MaxApplicableSoftwareVersionFor1_6_0),
			"--specificationVersion", fmt.Sprintf("%d", SpecificationVersionFor1_6_0),
			"--from", VendorAccountFor1_6_0,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)

		tx, err = ExecuteTxWithBin(dcldNew,
			"tx", "model", "add-model",
			"--vid", fmt.Sprintf("%d", VIDFor1_6_0),
			"--pid", fmt.Sprintf("%d", PID2For1_6_0),
			"--deviceTypeID", fmt.Sprintf("%d", DeviceTypeIDFor1_6_0),
			"--productName", ProductNameFor1_6_0,
			"--productLabel", ProductLabelFor1_6_0,
			"--partNumber", PartNumberFor1_6_0,
			"--from", VendorAccountFor1_6_0,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)

		tx, err = ExecuteTxWithBin(dcldNew,
			"tx", "model", "add-model-version",
			"--vid", fmt.Sprintf("%d", VIDFor1_6_0),
			"--pid", fmt.Sprintf("%d", PID2For1_6_0),
			"--softwareVersion", fmt.Sprintf("%d", SoftwareVersionFor1_6_0),
			"--softwareVersionString", SoftwareVersionStringFor1_6_0,
			"--cdVersionNumber", fmt.Sprintf("%d", CDVersionNumberFor1_6_0),
			"--minApplicableSoftwareVersion", fmt.Sprintf("%d", MinApplicableSoftwareVersionFor1_6_0),
			"--maxApplicableSoftwareVersion", fmt.Sprintf("%d", MaxApplicableSoftwareVersionFor1_6_0),
			"--from", VendorAccountFor1_6_0,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)

		// Update a v1.5.2-era model through the v1.5.2 vendor account.
		tx, err = ExecuteTxWithBin(dcldNew,
			"tx", "model", "update-model",
			"--vid", fmt.Sprintf("%d", VIDFor1_5_2),
			"--pid", fmt.Sprintf("%d", PID2For1_5_2),
			"--productName", ProductNameFor1_6_0,
			"--productLabel", ProductLabelFor1_6_0,
			"--partNumber", PartNumberFor1_6_0,
			"--from", VendorAccountFor1_5_2,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)
	})

	// Add a model with discoveryCapabilitiesBitmask=20 — allowed range
	// widened from 0-14 to 0-30 in v1.6.0.
	MustRun(t, "AddModelWithWidenedBitmask_1_6_0", func(t *testing.T) {
		t.Helper()
		tx, err := ExecuteTxWithBin(dcldNew,
			"tx", "model", "add-model",
			"--vid", fmt.Sprintf("%d", VIDFor1_6_0),
			"--pid", fmt.Sprintf("%d", PIDWidenedBitmaskFor1_6_0),
			"--deviceTypeID", fmt.Sprintf("%d", DeviceTypeIDFor1_6_0),
			"--productName", ProductNameFor1_6_0,
			"--productLabel", ProductLabelFor1_6_0,
			"--partNumber", PartNumberFor1_6_0,
			"--commissioningCustomFlow", fmt.Sprintf("%d", CommissioningCustomFlowFor1_6),
			"--discoveryCapabilitiesBitmask", fmt.Sprintf("%d", DiscoveryCapabilitiesBitmask),
			"--from", VendorAccountFor1_6_0,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)
	})

	// Compliance writes after upgrade — schemaVersion=1 (default per #730),
	// specificationVersion now required on every write. cDCertificateId is
	// reused across schema-v0 (carry-over) and schema-v1 records so the
	// device-software-compliance index covers both eras.
	MustRun(t, "ComplianceWrites_SchemaV1_1_6_0", func(t *testing.T) {
		t.Helper()
		tx, err := ExecuteTxWithBin(dcldNew,
			"tx", "compliance", "certify-model",
			"--vid", fmt.Sprintf("%d", VIDFor1_6_0),
			"--pid", fmt.Sprintf("%d", PID1For1_6_0),
			"--softwareVersion", fmt.Sprintf("%d", SoftwareVersionFor1_6_0),
			"--softwareVersionString", SoftwareVersionStringFor1_6_0,
			"--cdVersionNumber", fmt.Sprintf("%d", CDVersionNumberFor1_6_0),
			"--certificationType", CertificationTypeFor1_6_0,
			"--certificationDate", CertificationDateFor1_6_0,
			"--specificationVersion", fmt.Sprintf("%d", SpecificationVersionFor1_6_0),
			"--cdCertificateId", CDCertificateIDFor1_5_2,
			"--schemaVersion", "1",
			"--from", CertificationCenterAccountFor1_2,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)

		tx, err = ExecuteTxWithBin(dcldNew,
			"tx", "compliance", "provision-model",
			"--vid", fmt.Sprintf("%d", VIDFor1_6_0),
			"--pid", fmt.Sprintf("%d", PID2For1_6_0),
			"--softwareVersion", fmt.Sprintf("%d", SoftwareVersionFor1_6_0),
			"--softwareVersionString", SoftwareVersionStringFor1_6_0,
			"--cdVersionNumber", fmt.Sprintf("%d", CDVersionNumberFor1_6_0),
			"--certificationType", CertificationTypeFor1_6_0,
			"--provisionalDate", ProvisionalDateFor1_6_0,
			"--specificationVersion", fmt.Sprintf("%d", SpecificationVersionFor1_6_0),
			"--cdCertificateId", CDCertificateIDFor1_5_2,
			"--schemaVersion", "1",
			"--from", CertificationCenterAccountFor1_2,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)
	})

	MustRun(t, "VerifyNewModels_1_6_0", func(t *testing.T) {
		t.Helper()
		out, err := ExecuteCLIWithBin(dcldNew,
			"query", "model", "get-model",
			"--vid", fmt.Sprintf("%d", VIDFor1_6_0),
			"--pid", fmt.Sprintf("%d", PID1For1_6_0),
		)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", VIDFor1_6_0)
		requireFieldEquals(t, out, "pid", PID1For1_6_0)
		checkResponseContains(t, out, ProductLabelFor1_6_0)

		out, err = ExecuteCLIWithBin(dcldNew,
			"query", "model", "get-model",
			"--vid", fmt.Sprintf("%d", VIDFor1_6_0),
			"--pid", fmt.Sprintf("%d", PID2For1_6_0),
		)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", VIDFor1_6_0)
		requireFieldEquals(t, out, "pid", PID2For1_6_0)
		checkResponseContains(t, out, ProductLabelFor1_6_0)

		// Verify the 1.5.2 model now has 1.6.0 fields after the update.
		out, err = ExecuteCLIWithBin(dcldNew,
			"query", "model", "get-model",
			"--vid", fmt.Sprintf("%d", VIDFor1_5_2),
			"--pid", fmt.Sprintf("%d", PID2For1_5_2),
		)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", VIDFor1_5_2)
		requireFieldEquals(t, out, "pid", PID2For1_5_2)
		checkResponseContains(t, out, ProductLabelFor1_6_0)
		checkResponseContains(t, out, PartNumberFor1_6_0)

		// Model versions / specificationVersion.
		out, err = ExecuteCLIWithBin(dcldNew,
			"query", "model", "model-version",
			"--vid", fmt.Sprintf("%d", VIDFor1_6_0),
			"--pid", fmt.Sprintf("%d", PID1For1_6_0),
			"--softwareVersion", fmt.Sprintf("%d", SoftwareVersionFor1_6_0),
		)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", VIDFor1_6_0)
		requireFieldEquals(t, out, "pid", PID1For1_6_0)
		requireFieldEquals(t, out, "softwareVersion", SoftwareVersionFor1_6_0)
		requireFieldEquals(t, out, "specificationVersion", SpecificationVersionFor1_6_0)
	})

	// Verify the widened-bitmask model landed with its requested value.
	MustRun(t, "VerifyWidenedBitmaskModel_1_6_0", func(t *testing.T) {
		t.Helper()
		out, err := ExecuteCLIWithBin(dcldNew,
			"query", "model", "get-model",
			"--vid", fmt.Sprintf("%d", VIDFor1_6_0),
			"--pid", fmt.Sprintf("%d", PIDWidenedBitmaskFor1_6_0),
		)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", VIDFor1_6_0)
		requireFieldEquals(t, out, "pid", PIDWidenedBitmaskFor1_6_0)
		requireFieldEquals(t, out, "discoveryCapabilitiesBitmask", DiscoveryCapabilitiesBitmask)
	})

	// Verify the schema-v1 compliance writes landed with specificationVersion
	// now persisted on ComplianceInfo.
	MustRun(t, "VerifyComplianceWrites_SchemaV1_1_6_0", func(t *testing.T) {
		t.Helper()
		out, err := ExecuteCLIWithBin(dcldNew,
			"query", "compliance", "certified-model",
			"--vid", fmt.Sprintf("%d", VIDFor1_6_0),
			"--pid", fmt.Sprintf("%d", PID1For1_6_0),
			"--softwareVersion", fmt.Sprintf("%d", SoftwareVersionFor1_6_0),
			"--certificationType", CertificationTypeFor1_6_0,
		)
		require.NoError(t, err)
		requireFieldEquals(t, out, "value", true)
		requireFieldEquals(t, out, "vid", VIDFor1_6_0)

		out, err = ExecuteCLIWithBin(dcldNew,
			"query", "compliance", "compliance-info",
			"--vid", fmt.Sprintf("%d", VIDFor1_6_0),
			"--pid", fmt.Sprintf("%d", PID1For1_6_0),
			"--softwareVersion", fmt.Sprintf("%d", SoftwareVersionFor1_6_0),
			"--certificationType", CertificationTypeFor1_6_0,
		)
		require.NoError(t, err)
		requireFieldEquals(t, out, "schemaVersion", 1)
		requireFieldEquals(t, out, "specificationVersion", SpecificationVersionFor1_6_0)

		out, err = ExecuteCLIWithBin(dcldNew,
			"query", "compliance", "provisional-model",
			"--vid", fmt.Sprintf("%d", VIDFor1_6_0),
			"--pid", fmt.Sprintf("%d", PID2For1_6_0),
			"--softwareVersion", fmt.Sprintf("%d", SoftwareVersionFor1_6_0),
			"--certificationType", CertificationTypeFor1_6_0,
		)
		require.NoError(t, err)
		requireFieldEquals(t, out, "value", true)
		requireFieldEquals(t, out, "vid", VIDFor1_6_0)
	})
}
