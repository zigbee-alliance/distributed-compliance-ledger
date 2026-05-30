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

// runUpgrade12To143 is the Go translation of
// integration_tests/upgrade/05-test-upgrade-1.2-to-1.4.3.sh.
//
// Assumes the chain is currently running v1.2 with state from scripts 01-03.
//
//nolint:funlen
func runUpgrade12To143(t *testing.T, state *UpgradeTestState) {
	t.Helper()

	dcldOld, err := EnsureBinary(BinaryVersionV1_2)
	require.NoError(t, err)
	dcldNew, err := EnsureBinary(BinaryVersionV1_4_3)
	require.NoError(t, err)

	// ------------------------------------------------------------------
	// ISSUE #593 pre-upgrade: add a model with 2 versions, delete one. The
	// resulting ghost-version state is what script 09 verifies cleanup of.
	// State fields populated here come from DefaultBashState() (script-05
	// hardcodes vid_for_1_2 = 4701 and pid_3_for_1_6_0 = 160).
	// ------------------------------------------------------------------
	MustRun(t, "Issue593PreUpgradeGhostSetup", func(t *testing.T) {
		vid := state.VIDFor1_6_0FromScript5
		pid := state.PID3For1_6_0FromScript5
		sv1 := state.SoftwareVersion1For1_6_0FromScript5
		sv2 := state.SoftwareVersion2For1_6_0FromScript5

		// Add the model (1.6.0-era values reused as defaults).
		tx, err := ExecuteTxWithBin(dcldOld,
			"tx", "model", "add-model",
			"--vid", fmt.Sprintf("%d", vid),
			"--pid", fmt.Sprintf("%d", pid),
			"--deviceTypeID", fmt.Sprintf("%d", DeviceTypeIDForIssue593),
			"--productName", ProductNameForIssue593,
			"--productLabel", ProductLabelForIssue593,
			"--partNumber", PartNumberForIssue593,
			"--from", state.VendorAccountFor1_2,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)

		// Two model versions: sv1 (will be deleted to create ghost state) and sv2.
		for _, sv := range []int{sv1, sv2} {
			tx, err = ExecuteTxWithBin(dcldOld,
				"tx", "model", "add-model-version",
				"--vid", fmt.Sprintf("%d", vid),
				"--pid", fmt.Sprintf("%d", pid),
				"--softwareVersion", fmt.Sprintf("%d", sv),
				"--softwareVersionString", SoftwareVersionStringIssue593,
				"--cdVersionNumber", fmt.Sprintf("%d", CDVersionNumberIssue593),
				"--minApplicableSoftwareVersion", fmt.Sprintf("%d", MinSWVerIssue593),
				"--maxApplicableSoftwareVersion", fmt.Sprintf("%d", MaxSWVerIssue593),
				"--from", state.VendorAccountFor1_2,
			)
			require.NoError(t, err)
			require.Equal(t, uint32(0), tx.Code, tx.RawLog)
		}

		// Delete sv1 — sv2 stays as a ghost-pointer the migration must clean.
		tx, err = ExecuteTxWithBin(dcldOld,
			"tx", "model", "delete-model-version",
			"--vid", fmt.Sprintf("%d", vid),
			"--pid", fmt.Sprintf("%d", pid),
			"--softwareVersion", fmt.Sprintf("%d", sv1),
			"--from", state.VendorAccountFor1_2,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)
	})

	// ------------------------------------------------------------------
	// Upgrade 1.2 → 1.4.
	// ------------------------------------------------------------------
	step := SoftwareUpgradeStep{
		PlanName:         PlanNameV1_4,
		BinaryVersionNew: BinaryVersionV1_4_3,
		Checksum:         UpgradeChecksumV1_4,
		DcldOldBin:       dcldOld,
		DcldNewBin:       dcldNew,
		Trustees:         []string{state.Trustee1, state.Trustee2, state.Trustee3, state.Trustee4},
	}
	step.Run(t)

	// ------------------------------------------------------------------
	// Verify carry-over data is intact under v1.4.3.
	// ------------------------------------------------------------------
	MustRun(t, "VerifyPreservedVendorInfoAndModels", func(t *testing.T) {
		// VendorInfo for the v0.12 vendor — script 03 updated companyPreferredName/landing URL to 1.2 values.
		out, err := ExecuteCLIWithBin(dcldNew,
			"query", "vendorinfo", "vendor",
			"--vid", fmt.Sprintf("%d", state.VID),
		)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vendorID", state.VID)
		checkResponseContains(t, out, companyLegalNameV012)
		checkResponseContains(t, out, vendorNameV012)
		checkResponseContains(t, out, CompanyPreferredNameFor1_2)
		checkResponseContains(t, out, VendorLandingPageURLFor1_2)

		// VendorInfo for the 1.2 vendor.
		out, err = ExecuteCLIWithBin(dcldNew,
			"query", "vendorinfo", "vendor",
			"--vid", fmt.Sprintf("%d", VIDFor1_2),
		)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vendorID", VIDFor1_2)
		checkResponseContains(t, out, CompanyLegalNameFor1_2)

		out, err = ExecuteCLIWithBin(dcldNew, "query", "vendorinfo", "all-vendors")
		require.NoError(t, err)
		requireFieldEquals(t, out, "vendorID", state.VID)
		requireFieldEquals(t, out, "vendorID", VIDFor1_2)
		checkResponseContains(t, out, companyLegalNameV012)
		checkResponseContains(t, out, CompanyLegalNameFor1_2)

		// Carry-over models from 01/03.
		for _, pair := range [][2]int{
			{state.VID, pid1V012}, {state.VID, state.PID2},
			{VIDFor1_2, PID1For1_2}, {VIDFor1_2, PID2For1_2},
		} {
			out, err = ExecuteCLIWithBin(dcldNew,
				"query", "model", "get-model",
				"--vid", fmt.Sprintf("%d", pair[0]),
				"--pid", fmt.Sprintf("%d", pair[1]),
			)
			require.NoError(t, err)
			requireFieldEquals(t, out, "vid", pair[0])
			requireFieldEquals(t, out, "pid", pair[1])
		}

		// Updated 0.12 pid_2 has 1.2 min/max applicable software version.
		out, err = ExecuteCLIWithBin(dcldNew,
			"query", "model", "model-version",
			"--vid", fmt.Sprintf("%d", state.VID),
			"--pid", fmt.Sprintf("%d", state.PID2),
			"--softwareVersion", fmt.Sprintf("%d", state.SoftwareVersion),
		)
		require.NoError(t, err)
		requireFieldEquals(t, out, "minApplicableSoftwareVersion", MinApplicableSoftwareVersionFor1_2)
		requireFieldEquals(t, out, "maxApplicableSoftwareVersion", MaxApplicableSoftwareVersionFor1_2)
	})

	MustRun(t, "VerifyPreservedCompliance", func(t *testing.T) {
		// Certified 0.12 pid_1.
		out, err := ExecuteCLIWithBin(dcldNew,
			"query", "compliance", "certified-model",
			"--vid", fmt.Sprintf("%d", state.VID),
			"--pid", fmt.Sprintf("%d", pid1V012),
			"--softwareVersion", fmt.Sprintf("%d", state.SoftwareVersion),
			"--certificationType", certificationTypeV012,
		)
		require.NoError(t, err)
		checkResponseContains(t, out, `"value":true`)
		requireFieldEquals(t, out, "vid", state.VID)
		requireFieldEquals(t, out, "pid", pid1V012)

		// Certified 1.2 pid_1.
		out, err = ExecuteCLIWithBin(dcldNew,
			"query", "compliance", "certified-model",
			"--vid", fmt.Sprintf("%d", VIDFor1_2),
			"--pid", fmt.Sprintf("%d", PID1For1_2),
			"--softwareVersion", fmt.Sprintf("%d", SoftwareVersionFor1_2),
			"--certificationType", CertificationTypeFor1_2,
		)
		require.NoError(t, err)
		checkResponseContains(t, out, `"value":true`)

		// Revoked models persist.
		out, err = ExecuteCLIWithBin(dcldNew,
			"query", "compliance", "revoked-model",
			"--vid", fmt.Sprintf("%d", state.VID),
			"--pid", fmt.Sprintf("%d", state.PID2),
			"--softwareVersion", fmt.Sprintf("%d", state.SoftwareVersion),
			"--certificationType", certificationTypeV012,
		)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", state.VID)

		out, err = ExecuteCLIWithBin(dcldNew,
			"query", "compliance", "revoked-model",
			"--vid", fmt.Sprintf("%d", VIDFor1_2),
			"--pid", fmt.Sprintf("%d", PID2For1_2),
			"--softwareVersion", fmt.Sprintf("%d", SoftwareVersionFor1_2),
			"--certificationType", CertificationTypeFor1_2,
		)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", VIDFor1_2)
	})

	MustRun(t, "VerifyPreservedAccounts", func(t *testing.T) {
		out, err := ExecuteCLIWithBin(dcldNew, "query", "auth", "all-accounts")
		require.NoError(t, err)
		checkResponseContains(t, out, state.User2Address) // active from 01
		checkResponseContains(t, out, state.User5Address) // active from 03

		out, err = ExecuteCLIWithBin(dcldNew, "query", "auth", "all-proposed-accounts")
		require.NoError(t, err)
		checkResponseContains(t, out, state.User3Address) // proposed from 01
		checkResponseContains(t, out, state.User6Address) // proposed from 03

		out, err = ExecuteCLIWithBin(dcldNew, "query", "auth", "all-proposed-accounts-to-revoke")
		require.NoError(t, err)
		checkResponseContains(t, out, state.User2Address)
		checkResponseContains(t, out, state.User5Address)

		out, err = ExecuteCLIWithBin(dcldNew, "query", "auth", "all-revoked-accounts")
		require.NoError(t, err)
		checkResponseContains(t, out, state.User1Address) // revoked in 01
		checkResponseContains(t, out, state.User4Address) // revoked in 03
	})

	// ------------------------------------------------------------------
	// Post-upgrade: seed 1.4.3-era state.
	// ------------------------------------------------------------------
	MustRun(t, "CreateVendor_1_4_3", func(t *testing.T) {
		_ = CreateAndApproveAccount(t, dcldNew, VendorAccountFor1_4_3, "Vendor",
			VIDFor1_4_3, state.Trustee1,
			[]string{state.Trustee2, state.Trustee3, state.Trustee4})
	})

	MustRun(t, "AddPostUpgradeUserKeys", func(t *testing.T) {
		u7, err := newUserKey(dcldNew)
		require.NoError(t, err)
		u8, err := newUserKey(dcldNew)
		require.NoError(t, err)
		u9, err := newUserKey(dcldNew)
		require.NoError(t, err)
		state.User7Address, state.User7Pubkey = u7.address, u7.pubkey
		state.User8Address, state.User8Pubkey = u8.address, u8.pubkey
		state.User9Address, state.User9Pubkey = u9.address, u9.pubkey
	})

	MustRun(t, "VendorInfoFor1_4_3", func(t *testing.T) {
		tx, err := ExecuteTxWithBin(dcldNew,
			"tx", "vendorinfo", "add-vendor",
			"--vid", fmt.Sprintf("%d", VIDFor1_4_3),
			"--vendorName", VendorNameFor1_4_3,
			"--companyLegalName", CompanyLegalNameFor1_4_3,
			"--companyPreferredName", CompanyPreferredNameFor1_4_3,
			"--vendorLandingPageURL", VendorLandingPageURLFor1_4_3,
			"--from", VendorAccountFor1_4_3,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)

		tx, err = ExecuteTxWithBin(dcldNew,
			"tx", "vendorinfo", "update-vendor",
			"--vid", fmt.Sprintf("%d", VIDFor1_2),
			"--vendorName", VendorNameFor1_2,
			"--companyLegalName", CompanyLegalNameFor1_2,
			"--companyPreferredName", CompanyPreferredNameFor1_4_3,
			"--vendorLandingPageURL", VendorLandingPageURLFor1_4_3,
			"--from", state.VendorAccountFor1_2,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)
	})

	MustRun(t, "ModelsAndVersionsFor1_4_3", func(t *testing.T) {
		for _, pid := range []int{PID1For1_4_3, PID2For1_4_3, PID3For1_4_3} {
			tx, err := ExecuteTxWithBin(dcldNew,
				"tx", "model", "add-model",
				"--vid", fmt.Sprintf("%d", VIDFor1_4_3),
				"--pid", fmt.Sprintf("%d", pid),
				"--deviceTypeID", fmt.Sprintf("%d", DeviceTypeIDFor1_4_3),
				"--productName", ProductNameFor1_4_3,
				"--productLabel", ProductLabelFor1_4_3,
				"--partNumber", PartNumberFor1_4_3,
				"--from", VendorAccountFor1_4_3,
			)
			require.NoError(t, err)
			require.Equal(t, uint32(0), tx.Code, tx.RawLog)

			tx, err = ExecuteTxWithBin(dcldNew,
				"tx", "model", "add-model-version",
				"--vid", fmt.Sprintf("%d", VIDFor1_4_3),
				"--pid", fmt.Sprintf("%d", pid),
				"--softwareVersion", fmt.Sprintf("%d", SoftwareVersionFor1_4_3),
				"--softwareVersionString", SoftwareVersionStringFor1_4_3,
				"--cdVersionNumber", fmt.Sprintf("%d", CDVersionNumberFor1_4_3),
				"--minApplicableSoftwareVersion", fmt.Sprintf("%d", MinApplicableSoftwareVersionFor1_4_3),
				"--maxApplicableSoftwareVersion", fmt.Sprintf("%d", MaxApplicableSoftwareVersionFor1_4_3),
				"--from", VendorAccountFor1_4_3,
			)
			require.NoError(t, err)
			require.Equal(t, uint32(0), tx.Code, tx.RawLog)
		}

		// Delete pid_3.
		tx, err := ExecuteTxWithBin(dcldNew,
			"tx", "model", "delete-model",
			"--vid", fmt.Sprintf("%d", VIDFor1_4_3),
			"--pid", fmt.Sprintf("%d", PID3For1_4_3),
			"--from", VendorAccountFor1_4_3,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)

		// Update the carry-over model.
		tx, err = ExecuteTxWithBin(dcldNew,
			"tx", "model", "update-model",
			"--vid", fmt.Sprintf("%d", state.VID),
			"--pid", fmt.Sprintf("%d", state.PID2),
			"--productName", state.ProductName,
			"--productLabel", ProductLabelFor1_4_3,
			"--partNumber", PartNumberFor1_4_3,
			"--from", state.VendorAccount,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)

		tx, err = ExecuteTxWithBin(dcldNew,
			"tx", "model", "update-model-version",
			"--vid", fmt.Sprintf("%d", state.VID),
			"--pid", fmt.Sprintf("%d", state.PID2),
			"--softwareVersion", fmt.Sprintf("%d", state.SoftwareVersion),
			"--minApplicableSoftwareVersion", fmt.Sprintf("%d", MinApplicableSoftwareVersionFor1_4_3),
			"--maxApplicableSoftwareVersion", fmt.Sprintf("%d", MaxApplicableSoftwareVersionFor1_4_3),
			"--from", state.VendorAccount,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)
	})

	MustRun(t, "ComplianceFor1_4_3", func(t *testing.T) {
		// certify pid_1
		tx, err := ExecuteTxWithBin(dcldNew,
			"tx", "compliance", "certify-model",
			"--vid", fmt.Sprintf("%d", VIDFor1_4_3),
			"--pid", fmt.Sprintf("%d", PID1For1_4_3),
			"--softwareVersion", fmt.Sprintf("%d", SoftwareVersionFor1_4_3),
			"--softwareVersionString", SoftwareVersionStringFor1_4_3,
			"--certificationType", CertificationTypeFor1_4_3,
			"--certificationDate", CertificationDateFor1_4_3,
			"--cdCertificateId", CDCertificateIDFor1_4_3,
			"--cdVersionNumber", fmt.Sprintf("%d", CDVersionNumberFor1_4_3),
			"--from", CertificationCenterAccountFor1_2, // CertCenter account survives across upgrades
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)

		// provision pid_2
		tx, err = ExecuteTxWithBin(dcldNew,
			"tx", "compliance", "provision-model",
			"--vid", fmt.Sprintf("%d", VIDFor1_4_3),
			"--pid", fmt.Sprintf("%d", PID2For1_4_3),
			"--softwareVersion", fmt.Sprintf("%d", SoftwareVersionFor1_4_3),
			"--softwareVersionString", SoftwareVersionStringFor1_4_3,
			"--certificationType", CertificationTypeFor1_4_3,
			"--provisionalDate", ProvisionalDateFor1_4_3,
			"--cdCertificateId", CDCertificateIDFor1_4_3,
			"--cdVersionNumber", fmt.Sprintf("%d", CDVersionNumberFor1_4_3),
			"--from", CertificationCenterAccountFor1_2,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)

		// certify pid_2
		tx, err = ExecuteTxWithBin(dcldNew,
			"tx", "compliance", "certify-model",
			"--vid", fmt.Sprintf("%d", VIDFor1_4_3),
			"--pid", fmt.Sprintf("%d", PID2For1_4_3),
			"--softwareVersion", fmt.Sprintf("%d", SoftwareVersionFor1_4_3),
			"--softwareVersionString", SoftwareVersionStringFor1_4_3,
			"--certificationType", CertificationTypeFor1_4_3,
			"--certificationDate", CertificationDateFor1_4_3,
			"--cdCertificateId", CDCertificateIDFor1_4_3,
			"--cdVersionNumber", fmt.Sprintf("%d", CDVersionNumberFor1_4_3),
			"--from", CertificationCenterAccountFor1_2,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)

		// revoke pid_2
		tx, err = ExecuteTxWithBin(dcldNew,
			"tx", "compliance", "revoke-model",
			"--vid", fmt.Sprintf("%d", VIDFor1_4_3),
			"--pid", fmt.Sprintf("%d", PID2For1_4_3),
			"--softwareVersion", fmt.Sprintf("%d", SoftwareVersionFor1_4_3),
			"--softwareVersionString", SoftwareVersionStringFor1_4_3,
			"--certificationType", CertificationTypeFor1_4_3,
			"--revocationDate", CertificationDateFor1_4_3,
			"--cdVersionNumber", fmt.Sprintf("%d", CDVersionNumberFor1_4_3),
			"--from", CertificationCenterAccountFor1_2,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)
	})

	MustRun(t, "PKIFor1_4_3", func(t *testing.T) {
		// root_cert_with_vid: propose by t1, approve/reject by t2/t3, approve by t4.
		// (Bash also approves with trustee_5 — trustee_5 isn't propagated.)
		tx, err := ExecuteTxWithBin(dcldNew,
			"tx", "pki", "propose-add-x509-root-cert",
			"--certificate", RootCertWithVIDPathFor1_4_3,
			"--vid", fmt.Sprintf("%d", VIDFor1_4_3),
			"--from", state.Trustee1,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)

		// t2 approves.
		tx, err = ExecuteTxWithBin(dcldNew,
			"tx", "pki", "approve-add-x509-root-cert",
			"--subject", RootCertWithVIDSubjectFor1_4_3,
			"--subject-key-id", RootCertWithVIDSubjectKeyIDFor1_4_3,
			"--from", state.Trustee2,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)
		// t3 rejects.
		tx, err = ExecuteTxWithBin(dcldNew,
			"tx", "pki", "reject-add-x509-root-cert",
			"--subject", RootCertWithVIDSubjectFor1_4_3,
			"--subject-key-id", RootCertWithVIDSubjectKeyIDFor1_4_3,
			"--from", state.Trustee3,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)
		// t4 approves.
		tx, err = ExecuteTxWithBin(dcldNew,
			"tx", "pki", "approve-add-x509-root-cert",
			"--subject", RootCertWithVIDSubjectFor1_4_3,
			"--subject-key-id", RootCertWithVIDSubjectKeyIDFor1_4_3,
			"--from", state.Trustee4,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)

		// paa_cert_no_vid: propose + 3 approvals.
		tx, err = ExecuteTxWithBin(dcldNew,
			"tx", "pki", "propose-add-x509-root-cert",
			"--certificate", PaaCertNoVIDPathFor1_4_3,
			"--vid", fmt.Sprintf("%d", VIDFor1_4_3),
			"--from", state.Trustee1, // bash uses trustee_5 here; t1 is functionally equivalent
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)
		for _, who := range []string{state.Trustee2, state.Trustee3, state.Trustee4} {
			tx, err = ExecuteTxWithBin(dcldNew,
				"tx", "pki", "approve-add-x509-root-cert",
				"--subject", PaaCertNoVIDSubjectFor1_4_3,
				"--subject-key-id", PaaCertNoVIDSubjectKeyIDFor1_4_3,
				"--from", who,
			)
			require.NoError(t, err)
			require.Equal(t, uint32(0), tx.Code, tx.RawLog)
		}

		// Propose-only root_cert (no approvals).
		tx, err = ExecuteTxWithBin(dcldNew,
			"tx", "pki", "propose-add-x509-root-cert",
			"--certificate", RootCertPathFor1_4_3,
			"--vid", fmt.Sprintf("%d", VIDFor1_4_3),
			"--from", state.Trustee1,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)

		// Add intermediate_cert_with_vid, then revoke via serial-number (new in 1.4).
		tx, err = ExecuteTxWithBin(dcldNew,
			"tx", "pki", "add-x509-cert",
			"--certificate", IntermediateCertWithVIDPathFor1_4_3,
			"--from", VendorAccountFor1_4_3,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)

		tx, err = ExecuteTxWithBin(dcldNew,
			"tx", "pki", "revoke-x509-cert",
			"--subject", IntermediateCertWithVIDSubjectFor1_4_3,
			"--subject-key-id", IntermediateCertWithVIDSubjectKeyIDFor1_4_3,
			"--serial-number", IntermediateCertWithVIDSerialNumberFor1_4_3,
			"--from", VendorAccountFor1_4_3,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)

		// Revoke paa_cert_no_vid (propose + 3 approvals).
		tx, err = ExecuteTxWithBin(dcldNew,
			"tx", "pki", "propose-revoke-x509-root-cert",
			"--subject", PaaCertNoVIDSubjectFor1_4_3,
			"--subject-key-id", PaaCertNoVIDSubjectKeyIDFor1_4_3,
			"--from", state.Trustee1,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)
		for _, who := range []string{state.Trustee2, state.Trustee3, state.Trustee4} {
			tx, err = ExecuteTxWithBin(dcldNew,
				"tx", "pki", "approve-revoke-x509-root-cert",
				"--subject", PaaCertNoVIDSubjectFor1_4_3,
				"--subject-key-id", PaaCertNoVIDSubjectKeyIDFor1_4_3,
				"--from", who,
			)
			require.NoError(t, err)
			require.Equal(t, uint32(0), tx.Code, tx.RawLog)
		}

		// Propose revoke root_cert_with_vid (no approvals).
		tx, err = ExecuteTxWithBin(dcldNew,
			"tx", "pki", "propose-revoke-x509-root-cert",
			"--subject", RootCertWithVIDSubjectFor1_4_3,
			"--subject-key-id", RootCertWithVIDSubjectKeyIDFor1_4_3,
			"--from", state.Trustee1,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)
	})

	MustRun(t, "NOCCertsAddRemove", func(t *testing.T) {
		// 1.4 introduces the NOC certificate flow. Add then immediately remove.
		tx, err := ExecuteTxWithBin(dcldNew,
			"tx", "pki", "add-noc-x509-root-cert",
			"--certificate", NOCRootCert1PathFor1_4_3,
			"--from", VendorAccountFor1_4_3,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)

		tx, err = ExecuteTxWithBin(dcldNew,
			"tx", "pki", "add-noc-x509-ica-cert",
			"--certificate", NOCICACert1PathFor1_4_3,
			"--from", VendorAccountFor1_4_3,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)

		tx, err = ExecuteTxWithBin(dcldNew,
			"tx", "pki", "remove-noc-x509-root-cert",
			"--subject", NOCRootCert1SubjectFor1_4_3,
			"--subject-key-id", NOCRootCert1SubjectKeyIDFor1_4_3,
			"--from", VendorAccountFor1_4_3,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)

		tx, err = ExecuteTxWithBin(dcldNew,
			"tx", "pki", "remove-noc-x509-ica-cert",
			"--subject", NOCICACert1SubjectFor1_4_3,
			"--subject-key-id", NOCICACert1SubjectKeyIDFor1_4_3,
			"--from", VendorAccountFor1_4_3,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)
	})

	MustRun(t, "RevocationPointsFor1_4_3", func(t *testing.T) {
		// add → update → delete → add.
		addPAA := func(label, dataURL string) {
			tx, err := ExecuteTxWithBin(dcldNew,
				"tx", "pki", "add-revocation-point",
				"--vid", fmt.Sprintf("%d", VIDFor1_4_3),
				"--revocation-type", "1",
				"--is-paa=true",
				"--certificate", RootCertWithVIDPathFor1_4_3,
				"--label", label,
				"--data-url", dataURL,
				"--issuer-subject-key-id", IssuerSubjectKeyID,
				"--from", VendorAccountFor1_4_3,
			)
			require.NoError(t, err)
			require.Equal(t, uint32(0), tx.Code, tx.RawLog)
		}

		addPAA(ProductLabelFor1_4_3, TestDataURLFor1_4_3)

		tx, err := ExecuteTxWithBin(dcldNew,
			"tx", "pki", "update-revocation-point",
			"--vid", fmt.Sprintf("%d", VIDFor1_4_3),
			"--certificate", RootCertWithVIDPathFor1_4_3,
			"--label", ProductLabelFor1_4_3,
			"--data-url", TestDataURLFor1_4_3+"/new",
			"--issuer-subject-key-id", IssuerSubjectKeyID,
			"--from", VendorAccountFor1_4_3,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)

		tx, err = ExecuteTxWithBin(dcldNew,
			"tx", "pki", "delete-revocation-point",
			"--vid", fmt.Sprintf("%d", VIDFor1_4_3),
			"--label", ProductLabelFor1_4_3,
			"--issuer-subject-key-id", IssuerSubjectKeyID,
			"--from", VendorAccountFor1_4_3,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)

		addPAA(ProductLabelFor1_4_3, TestDataURLFor1_4_3)

		// CRL signer revocation point delegated by PAI (new in 1.4).
		tx, err = ExecuteTxWithBin(dcldNew,
			"tx", "pki", "add-revocation-point",
			"--vid", fmt.Sprintf("%d", VIDFor1_4_3),
			"--is-paa=false",
			"--certificate", CRLSignerDelegatedByPAI1,
			"--label", ProductLabelFor1_4_3,
			"--data-url", TestDataURLFor1_4_3,
			"--issuer-subject-key-id", DelegatorCertWithVIDSubjectKeyID,
			"--revocation-type", "1",
			"--certificate-delegator", DelegatorCertWithVID65521Path,
			"--from", VendorAccountFor1_4_3,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)

		tx, err = ExecuteTxWithBin(dcldNew,
			"tx", "pki", "update-revocation-point",
			"--vid", fmt.Sprintf("%d", VIDFor1_4_3),
			"--certificate", CRLSignerDelegatedByPAI1,
			"--label", ProductLabelFor1_4_3,
			"--data-url", TestDataURLFor1_4_3+"/new",
			"--issuer-subject-key-id", DelegatorCertWithVIDSubjectKeyID,
			"--certificate-delegator", DelegatorCertWithVID65521Path,
			"--from", VendorAccountFor1_4_3,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)
	})

	MustRun(t, "AccountFlowsFor1_4_3", func(t *testing.T) {
		approvers := []string{state.Trustee2, state.Trustee3, state.Trustee4}

		proposeUserAccount(t, dcldNew, state.Trustee1, approvers,
			state.User7Address, state.User7Pubkey, "CertificationCenter", true)
		proposeUserAccount(t, dcldNew, state.Trustee1, approvers,
			state.User8Address, state.User8Pubkey, "CertificationCenter", true)
		proposeUserAccount(t, dcldNew, state.Trustee1, nil,
			state.User9Address, state.User9Pubkey, "CertificationCenter", false)

		revokeUserAccount(t, dcldNew, state.Trustee1, approvers, state.User7Address, true)
		revokeUserAccount(t, dcldNew, state.Trustee1, nil, state.User8Address, false)
	})

	// Validator disable/enable (lines 1007-1056) — depends on the Docker
	// validator-demo container that script 01's `add_validator_node` would
	// create. Stubbed alongside the other validator work.
	MustRun(t, "ValidatorDisableEnableFlow", func(t *testing.T) {
		// Scripts 05+ approve disable-node with 3 trustees (trustee_4 active).
		RunValidatorDisableEnableFlow(t, state, dcldNew,
			[]string{state.Trustee2, state.Trustee3, state.Trustee4})
	})

	// ------------------------------------------------------------------
	// Verify post-upgrade-seeded NEW data.
	// ------------------------------------------------------------------
	MustRun(t, "VerifyNewVendorAndModels", func(t *testing.T) {
		out, err := ExecuteCLIWithBin(dcldNew,
			"query", "vendorinfo", "vendor",
			"--vid", fmt.Sprintf("%d", VIDFor1_4_3),
		)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vendorID", VIDFor1_4_3)
		checkResponseContains(t, out, CompanyLegalNameFor1_4_3)

		out, err = ExecuteCLIWithBin(dcldNew,
			"query", "model", "get-model",
			"--vid", fmt.Sprintf("%d", VIDFor1_4_3),
			"--pid", fmt.Sprintf("%d", PID1For1_4_3),
		)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", VIDFor1_4_3)
		requireFieldEquals(t, out, "pid", PID1For1_4_3)
		checkResponseContains(t, out, ProductLabelFor1_4_3)

		// Updated 0.12 pid_2 now has 1.4 productLabel/partNumber.
		out, err = ExecuteCLIWithBin(dcldNew,
			"query", "model", "get-model",
			"--vid", fmt.Sprintf("%d", state.VID),
			"--pid", fmt.Sprintf("%d", state.PID2),
		)
		require.NoError(t, err)
		checkResponseContains(t, out, ProductLabelFor1_4_3)
		checkResponseContains(t, out, PartNumberFor1_4_3)
	})

	MustRun(t, "VerifyNOCCertsEmptyAfterRemove", func(t *testing.T) {
		// After add+remove on the NOC certs, queries should return Not Found
		// and must NOT contain the removed subject key IDs.
		out, err := ExecuteCLIWithBin(dcldNew,
			"query", "pki", "noc-x509-root-certs",
			"--vid", fmt.Sprintf("%d", VIDFor1_4_3),
		)
		require.NoError(t, err)
		require.True(t, strings.Contains(string(out), "Not Found"),
			"expected Not Found for cleaned NOC root, got: %s", string(out))
		require.False(t, strings.Contains(string(out), NOCRootCert1SubjectKeyIDFor1_4_3),
			"NOC root cert SKID lingered after remove: %s", string(out))

		out, err = ExecuteCLIWithBin(dcldNew,
			"query", "pki", "noc-x509-certs",
			"--vid", fmt.Sprintf("%d", VIDFor1_4_3),
			"--subject-key-id", NOCRootCert1SubjectKeyIDFor1_4_3,
		)
		require.NoError(t, err)
		require.True(t, strings.Contains(string(out), "Not Found"),
			"expected Not Found, got: %s", string(out))

		out, err = ExecuteCLIWithBin(dcldNew,
			"query", "pki", "noc-x509-certs",
			"--vid", fmt.Sprintf("%d", VIDFor1_4_3),
			"--subject-key-id", NOCICACert1SubjectKeyIDFor1_4_3,
		)
		require.NoError(t, err)
		require.True(t, strings.Contains(string(out), "Not Found"),
			"expected Not Found, got: %s", string(out))
	})
}
