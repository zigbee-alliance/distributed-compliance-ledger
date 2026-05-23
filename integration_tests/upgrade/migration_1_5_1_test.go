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

// runUpgrade144To151 is the Go translation of
// integration_tests/upgrade/07-test-upgrade-1.4.4-to-1.5.1.sh.
//
// This is the final Phase 2 script — after it runs the chain is at v1.5.1 and
// Phase 1's 08/09 subtests can take over.
//
//nolint:funlen
func runUpgrade144To151(t *testing.T, state *UpgradeTestState) {
	t.Helper()

	dcldOld, err := EnsureBinary(BinaryVersionV1_4_4)
	require.NoError(t, err)
	dcldNew, err := EnsureBinary(BinaryVersionV1_5_1)
	require.NoError(t, err)

	step := SoftwareUpgradeStep{
		PlanName:         PlanNameV1_5_1,
		BinaryVersionNew: "v" + BinaryVersionV1_5_1,
		Checksum:         UpgradeChecksumV1_5_1,
		DcldOldBin:       dcldOld,
		DcldNewBin:       dcldNew,
		Trustees:         []string{state.Trustee1, state.Trustee2, state.Trustee3, state.Trustee4},
	}
	step.Run(t)

	// ------------------------------------------------------------------
	// Verify carry-over data is intact under v1.5.1.
	// ------------------------------------------------------------------
	t.Run("VerifyPreservedAcrossFourEras", func(t *testing.T) {
		for _, vid := range []int{state.VID, VIDFor1_2, VIDFor1_4_3, VIDFor1_4_4} {
			out, qerr := ExecuteCLIWithBin(dcldNew,
				"query", "vendorinfo", "vendor",
				"--vid", fmt.Sprintf("%d", vid),
			)
			require.NoError(t, qerr)
			requireFieldEquals(t, out, "vendorID", vid)
		}

		// 0.12 pid_2 now carries 1.4.4 productLabel/partNumber (set in script 06).
		out, err := ExecuteCLIWithBin(dcldNew,
			"query", "model", "get-model",
			"--vid", fmt.Sprintf("%d", state.VID),
			"--pid", fmt.Sprintf("%d", state.PID2),
		)
		require.NoError(t, err)
		checkResponseContains(t, out, ProductLabelFor1_4_4)
		checkResponseContains(t, out, PartNumberFor1_4_4)

		out, err = ExecuteCLIWithBin(dcldNew,
			"query", "model", "model-version",
			"--vid", fmt.Sprintf("%d", state.VID),
			"--pid", fmt.Sprintf("%d", state.PID2),
			"--softwareVersion", fmt.Sprintf("%d", state.SoftwareVersion),
		)
		require.NoError(t, err)
		requireFieldEquals(t, out, "minApplicableSoftwareVersion", MinApplicableSoftwareVersionFor1_4_4)
		requireFieldEquals(t, out, "maxApplicableSoftwareVersion", MaxApplicableSoftwareVersionFor1_4_4)
	})

	t.Run("VerifyPreservedAccounts", func(t *testing.T) {
		out, err := ExecuteCLIWithBin(dcldNew, "query", "auth", "all-accounts")
		require.NoError(t, err)
		// Active accounts across all prior scripts.
		for _, addr := range []string{
			state.User2Address, state.User5Address, state.User8Address, state.User11Address,
		} {
			checkResponseContains(t, out, addr)
		}

		out, err = ExecuteCLIWithBin(dcldNew, "query", "auth", "all-revoked-accounts")
		require.NoError(t, err)
		for _, addr := range []string{
			state.User1Address, state.User4Address, state.User7Address, state.User10Address,
		} {
			checkResponseContains(t, out, addr)
		}
	})

	// ------------------------------------------------------------------
	// Post-upgrade: seed 1.5.1-era state.
	// ------------------------------------------------------------------
	t.Run("CreateVendor_1_5_1", func(t *testing.T) {
		_ = CreateAndApproveAccount(t, dcldNew, VendorAccountFor1_5_1, "Vendor",
			state.VIDFor1_5_1, state.Trustee1,
			[]string{state.Trustee2, state.Trustee3, state.Trustee4})
	})

	t.Run("AddPostUpgradeUserKeys", func(t *testing.T) {
		u13, err := newUserKey(dcldNew)
		require.NoError(t, err)
		u14, err := newUserKey(dcldNew)
		require.NoError(t, err)
		u15, err := newUserKey(dcldNew)
		require.NoError(t, err)
		state.User13Address, state.User13Pubkey = u13.address, u13.pubkey
		state.User14Address, state.User14Pubkey = u14.address, u14.pubkey
		state.User15Address, state.User15Pubkey = u15.address, u15.pubkey
	})

	t.Run("VendorInfoFor1_5_1", func(t *testing.T) {
		tx, err := ExecuteTxWithBin(dcldNew,
			"tx", "vendorinfo", "add-vendor",
			"--vid", fmt.Sprintf("%d", state.VIDFor1_5_1),
			"--vendorName", VendorNameFor1_5_1,
			"--companyLegalName", CompanyLegalNameFor1_5_1,
			"--companyPreferredName", CompanyPreferredNameFor1_5_1,
			"--vendorLandingPageURL", VendorLandingPageURLFor1_5_1,
			"--from", VendorAccountFor1_5_1,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)

		tx, err = ExecuteTxWithBin(dcldNew,
			"tx", "vendorinfo", "update-vendor",
			"--vid", fmt.Sprintf("%d", VIDFor1_2),
			"--vendorName", VendorNameFor1_2,
			"--companyLegalName", CompanyLegalNameFor1_2,
			"--companyPreferredName", CompanyPreferredNameFor1_5_1,
			"--vendorLandingPageURL", VendorLandingPageURLFor1_5_1,
			"--from", state.VendorAccountFor1_2,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)
	})

	t.Run("ModelsAndVersionsFor1_5_1", func(t *testing.T) {
		// pid_1 with full 1.5-era field set (ICD, factory-reset, commissioning sec hint).
		tx, err := ExecuteTxWithBin(dcldNew,
			"tx", "model", "add-model",
			"--vid", fmt.Sprintf("%d", state.VIDFor1_5_1),
			"--pid", fmt.Sprintf("%d", state.PID1For1_5_1),
			"--deviceTypeID", fmt.Sprintf("%d", DeviceTypeIDFor1_5_1),
			"--productName", ProductNameFor1_5_1,
			"--productLabel", state.ProductLabelFor1_5_1,
			"--partNumber", state.PartNumberFor1_5_1,
			"--commissioningModeSecondaryStepsHint",
			fmt.Sprintf("%d", state.CommissioningModeSecondaryStepsHintFor1_5_1),
			"--icdUserActiveModeTriggerHint",
			fmt.Sprintf("%d", ICDUserActiveModeTriggerHintFor1_5_1),
			"--icdUserActiveModeTriggerInstruction",
			ICDUserActiveModeTriggerInstructionFor1_5_1,
			"--factoryResetStepsHint",
			fmt.Sprintf("%d", FactoryResetStepsHintFor1_5_1),
			"--factoryResetStepsInstruction",
			FactoryResetStepsInstructionFor1_5_1,
			"--from", VendorAccountFor1_5_1,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)

		// pid_1 version with specificationVersion (1.5-era).
		tx, err = ExecuteTxWithBin(dcldNew,
			"tx", "model", "add-model-version",
			"--vid", fmt.Sprintf("%d", state.VIDFor1_5_1),
			"--pid", fmt.Sprintf("%d", state.PID1For1_5_1),
			"--softwareVersion", fmt.Sprintf("%d", state.SoftwareVersionFor1_5_1),
			"--softwareVersionString", SoftwareVersionStringFor1_5_1,
			"--cdVersionNumber", fmt.Sprintf("%d", CDVersionNumberFor1_5_1),
			"--minApplicableSoftwareVersion", fmt.Sprintf("%d", state.MinApplicableSoftwareVersionFor1_5_1),
			"--maxApplicableSoftwareVersion", fmt.Sprintf("%d", state.MaxApplicableSoftwareVersionFor1_5_1),
			"--specificationVersion", fmt.Sprintf("%d", SpecificationVersionFor1_5_1),
			"--from", VendorAccountFor1_5_1,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)

		// pid_2 (no new fields).
		tx, err = ExecuteTxWithBin(dcldNew,
			"tx", "model", "add-model",
			"--vid", fmt.Sprintf("%d", state.VIDFor1_5_1),
			"--pid", fmt.Sprintf("%d", state.PID2For1_5_1),
			"--deviceTypeID", fmt.Sprintf("%d", DeviceTypeIDFor1_5_1),
			"--productName", ProductNameFor1_5_1,
			"--productLabel", state.ProductLabelFor1_5_1,
			"--partNumber", state.PartNumberFor1_5_1,
			"--from", VendorAccountFor1_5_1,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)
		tx, err = ExecuteTxWithBin(dcldNew,
			"tx", "model", "add-model-version",
			"--vid", fmt.Sprintf("%d", state.VIDFor1_5_1),
			"--pid", fmt.Sprintf("%d", state.PID2For1_5_1),
			"--softwareVersion", fmt.Sprintf("%d", state.SoftwareVersionFor1_5_1),
			"--softwareVersionString", SoftwareVersionStringFor1_5_1,
			"--cdVersionNumber", fmt.Sprintf("%d", CDVersionNumberFor1_5_1),
			"--minApplicableSoftwareVersion", fmt.Sprintf("%d", state.MinApplicableSoftwareVersionFor1_5_1),
			"--maxApplicableSoftwareVersion", fmt.Sprintf("%d", state.MaxApplicableSoftwareVersionFor1_5_1),
			"--from", VendorAccountFor1_5_1,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)

		// pid_3 add + delete.
		tx, err = ExecuteTxWithBin(dcldNew,
			"tx", "model", "add-model",
			"--vid", fmt.Sprintf("%d", state.VIDFor1_5_1),
			"--pid", fmt.Sprintf("%d", PID3For1_5_1),
			"--deviceTypeID", fmt.Sprintf("%d", DeviceTypeIDFor1_5_1),
			"--productName", ProductNameFor1_5_1,
			"--productLabel", state.ProductLabelFor1_5_1,
			"--partNumber", state.PartNumberFor1_5_1,
			"--from", VendorAccountFor1_5_1,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)
		tx, err = ExecuteTxWithBin(dcldNew,
			"tx", "model", "add-model-version",
			"--vid", fmt.Sprintf("%d", state.VIDFor1_5_1),
			"--pid", fmt.Sprintf("%d", PID3For1_5_1),
			"--softwareVersion", fmt.Sprintf("%d", state.SoftwareVersionFor1_5_1),
			"--softwareVersionString", SoftwareVersionStringFor1_5_1,
			"--cdVersionNumber", fmt.Sprintf("%d", CDVersionNumberFor1_5_1),
			"--minApplicableSoftwareVersion", fmt.Sprintf("%d", state.MinApplicableSoftwareVersionFor1_5_1),
			"--maxApplicableSoftwareVersion", fmt.Sprintf("%d", state.MaxApplicableSoftwareVersionFor1_5_1),
			"--from", VendorAccountFor1_5_1,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)
		tx, err = ExecuteTxWithBin(dcldNew,
			"tx", "model", "delete-model",
			"--vid", fmt.Sprintf("%d", state.VIDFor1_5_1),
			"--pid", fmt.Sprintf("%d", PID3For1_5_1),
			"--from", VendorAccountFor1_5_1,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)

		// Update the 0.12 pid_2 model with 1.5.1 productLabel/partNumber.
		tx, err = ExecuteTxWithBin(dcldNew,
			"tx", "model", "update-model",
			"--vid", fmt.Sprintf("%d", state.VID),
			"--pid", fmt.Sprintf("%d", state.PID2),
			"--productName", state.ProductName,
			"--productLabel", state.ProductLabelFor1_5_1,
			"--partNumber", state.PartNumberFor1_5_1,
			"--from", state.VendorAccount,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)

		tx, err = ExecuteTxWithBin(dcldNew,
			"tx", "model", "update-model-version",
			"--vid", fmt.Sprintf("%d", state.VID),
			"--pid", fmt.Sprintf("%d", state.PID2),
			"--softwareVersion", fmt.Sprintf("%d", state.SoftwareVersion),
			"--minApplicableSoftwareVersion", fmt.Sprintf("%d", state.MinApplicableSoftwareVersionFor1_5_1),
			"--maxApplicableSoftwareVersion", fmt.Sprintf("%d", state.MaxApplicableSoftwareVersionFor1_5_1),
			"--from", state.VendorAccount,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)
	})

	t.Run("ComplianceFor1_5_1", func(t *testing.T) {
		// certify pid_1
		tx, err := ExecuteTxWithBin(dcldNew,
			"tx", "compliance", "certify-model",
			"--vid", fmt.Sprintf("%d", state.VIDFor1_5_1),
			"--pid", fmt.Sprintf("%d", state.PID1For1_5_1),
			"--softwareVersion", fmt.Sprintf("%d", state.SoftwareVersionFor1_5_1),
			"--softwareVersionString", SoftwareVersionStringFor1_5_1,
			"--certificationType", CertificationTypeFor1_5_1,
			"--certificationDate", CertificationDateFor1_5_1,
			"--cdCertificateId", CDCertificateIDFor1_5_1,
			"--cdVersionNumber", fmt.Sprintf("%d", CDVersionNumberFor1_5_1),
			"--from", CertificationCenterAccountFor1_2,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)

		// provision pid_2
		tx, err = ExecuteTxWithBin(dcldNew,
			"tx", "compliance", "provision-model",
			"--vid", fmt.Sprintf("%d", state.VIDFor1_5_1),
			"--pid", fmt.Sprintf("%d", state.PID2For1_5_1),
			"--softwareVersion", fmt.Sprintf("%d", state.SoftwareVersionFor1_5_1),
			"--softwareVersionString", SoftwareVersionStringFor1_5_1,
			"--certificationType", CertificationTypeFor1_5_1,
			"--provisionalDate", ProvisionalDateFor1_5_1,
			"--cdCertificateId", CDCertificateIDFor1_5_1,
			"--cdVersionNumber", fmt.Sprintf("%d", CDVersionNumberFor1_5_1),
			"--from", CertificationCenterAccountFor1_2,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)

		// certify pid_2
		tx, err = ExecuteTxWithBin(dcldNew,
			"tx", "compliance", "certify-model",
			"--vid", fmt.Sprintf("%d", state.VIDFor1_5_1),
			"--pid", fmt.Sprintf("%d", state.PID2For1_5_1),
			"--softwareVersion", fmt.Sprintf("%d", state.SoftwareVersionFor1_5_1),
			"--softwareVersionString", SoftwareVersionStringFor1_5_1,
			"--certificationType", CertificationTypeFor1_5_1,
			"--certificationDate", CertificationDateFor1_5_1,
			"--cdCertificateId", CDCertificateIDFor1_5_1,
			"--cdVersionNumber", fmt.Sprintf("%d", CDVersionNumberFor1_5_1),
			"--from", CertificationCenterAccountFor1_2,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)

		// revoke pid_2
		tx, err = ExecuteTxWithBin(dcldNew,
			"tx", "compliance", "revoke-model",
			"--vid", fmt.Sprintf("%d", state.VIDFor1_5_1),
			"--pid", fmt.Sprintf("%d", state.PID2For1_5_1),
			"--softwareVersion", fmt.Sprintf("%d", state.SoftwareVersionFor1_5_1),
			"--softwareVersionString", SoftwareVersionStringFor1_5_1,
			"--certificationType", CertificationTypeFor1_5_1,
			"--revocationDate", CertificationDateFor1_5_1,
			"--cdVersionNumber", fmt.Sprintf("%d", CDVersionNumberFor1_5_1),
			"--from", CertificationCenterAccountFor1_2,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)
	})

	t.Run("AccountFlowsFor1_5_1", func(t *testing.T) {
		approvers := []string{state.Trustee2, state.Trustee3, state.Trustee4}

		proposeUserAccount(t, dcldNew, state.Trustee1, approvers,
			state.User13Address, state.User13Pubkey, "CertificationCenter", true)
		proposeUserAccount(t, dcldNew, state.Trustee1, approvers,
			state.User14Address, state.User14Pubkey, "CertificationCenter", true)
		proposeUserAccount(t, dcldNew, state.Trustee1, nil,
			state.User15Address, state.User15Pubkey, "CertificationCenter", false)

		revokeUserAccount(t, dcldNew, state.Trustee1, approvers, state.User13Address, true)
		revokeUserAccount(t, dcldNew, state.Trustee1, nil, state.User14Address, false)
	})

	t.Run("ValidatorDisableEnableFlow", func(t *testing.T) {
		RunValidatorDisableEnableFlow(t, state, dcldNew,
			[]string{state.Trustee2, state.Trustee3, state.Trustee4})
	})

	// ------------------------------------------------------------------
	// Verify post-upgrade-seeded NEW 1.5.1 data. The Phase 1 subtests
	// (08/09) rely on this state being present.
	// ------------------------------------------------------------------
	t.Run("VerifyNew_1_5_1_Data", func(t *testing.T) {
		out, err := ExecuteCLIWithBin(dcldNew,
			"query", "vendorinfo", "vendor",
			"--vid", fmt.Sprintf("%d", state.VIDFor1_5_1),
		)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vendorID", state.VIDFor1_5_1)
		checkResponseContains(t, out, CompanyLegalNameFor1_5_1)

		// pid_1 has full 1.5-era fields.
		out, err = ExecuteCLIWithBin(dcldNew,
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
		requireFieldEquals(t, out, "icdUserActiveModeTriggerHint",
			ICDUserActiveModeTriggerHintFor1_5_1)
		requireFieldEquals(t, out, "factoryResetStepsHint",
			FactoryResetStepsHintFor1_5_1)

		// pid_2 with defaults.
		out, err = ExecuteCLIWithBin(dcldNew,
			"query", "model", "get-model",
			"--vid", fmt.Sprintf("%d", state.VIDFor1_5_1),
			"--pid", fmt.Sprintf("%d", state.PID2For1_5_1),
		)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", state.VIDFor1_5_1)
		requireFieldEquals(t, out, "pid", state.PID2For1_5_1)

		// 0.12 pid_2 now has 1.5.1 productLabel/partNumber.
		out, err = ExecuteCLIWithBin(dcldNew,
			"query", "model", "get-model",
			"--vid", fmt.Sprintf("%d", state.VID),
			"--pid", fmt.Sprintf("%d", state.PID2),
		)
		require.NoError(t, err)
		checkResponseContains(t, out, state.ProductLabelFor1_5_1)
		checkResponseContains(t, out, state.PartNumberFor1_5_1)

		// Model version specificationVersion.
		out, err = ExecuteCLIWithBin(dcldNew,
			"query", "model", "model-version",
			"--vid", fmt.Sprintf("%d", state.VIDFor1_5_1),
			"--pid", fmt.Sprintf("%d", state.PID1For1_5_1),
			"--softwareVersion", fmt.Sprintf("%d", state.SoftwareVersionFor1_5_1),
		)
		require.NoError(t, err)
		requireFieldEquals(t, out, "specificationVersion", SpecificationVersionFor1_5_1)
	})
}
