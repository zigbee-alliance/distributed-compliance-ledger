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
		BinaryVersionNew: BinaryVersionV1_5_1,
		Checksum:         UpgradeChecksumV1_5_1,
		DcldOldBin:       dcldOld,
		DcldNewBin:       dcldNew,
		Trustees:         []string{state.Trustee1, state.Trustee2, state.Trustee3, state.Trustee4},
	}
	step.Run(t)

	// ------------------------------------------------------------------
	// Verify carry-over data is intact under v1.5.1.
	// ------------------------------------------------------------------
	MustRun(t, "VerifyPreservedAcrossFourEras", func(t *testing.T) {
		t.Helper()
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

	MustRun(t, "VerifyPreservedAccounts", func(t *testing.T) {
		t.Helper()
		out, err := ExecuteCLIWithBin(dcldNew, "query", "auth", "all-accounts")
		require.NoError(t, err)
		// Active accounts across all prior scripts.
		for _, addr := range []string{
			state.User2Address, state.User5Address, state.User8Address, state.User11Address,
		} {
			checkResponseContains(t, out, addr)
		}

		out, err = ExecuteCLIWithBin(dcldNew, "query", "auth", "all-proposed-accounts")
		require.NoError(t, err)
		for _, addr := range []string{
			state.User3Address, state.User6Address, state.User9Address, state.User12Address,
		} {
			checkResponseContains(t, out, addr)
		}

		out, err = ExecuteCLIWithBin(dcldNew, "query", "auth", "all-proposed-accounts-to-revoke")
		require.NoError(t, err)
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

		// Single-record account variants (bash 07 lines 364-411 style).
		for _, addr := range []string{
			state.User2Address, state.User5Address, state.User8Address, state.User11Address,
		} {
			out, err = ExecuteCLIWithBin(dcldNew, "query", "auth", "account", "--address", addr)
			require.NoError(t, err)
			checkResponseContains(t, out, addr)
		}
		for _, addr := range []string{
			state.User3Address, state.User6Address, state.User9Address, state.User12Address,
		} {
			out, err = ExecuteCLIWithBin(dcldNew, "query", "auth", "proposed-account", "--address", addr)
			require.NoError(t, err)
			checkResponseContains(t, out, addr)
		}
		for _, addr := range []string{
			state.User2Address, state.User5Address, state.User8Address, state.User11Address,
		} {
			out, err = ExecuteCLIWithBin(dcldNew, "query", "auth", "proposed-account-to-revoke", "--address", addr)
			require.NoError(t, err)
			checkResponseContains(t, out, addr)
		}
		for _, addr := range []string{
			state.User1Address, state.User4Address, state.User7Address, state.User10Address,
		} {
			out, err = ExecuteCLIWithBin(dcldNew, "query", "auth", "revoked-account", "--address", addr)
			require.NoError(t, err)
			checkResponseContains(t, out, addr)
		}
	})

	// Bulk readback from bash 07. Adds gap-fill compliance/model/pki listings
	// + remaining single-record forms, spanning the four pre-1.5.1 eras.
	MustRun(t, "VerifyPreservedListings_1_5_1", func(t *testing.T) {
		t.Helper()
		out, err := ExecuteCLIWithBin(dcldNew, "query", "vendorinfo", "all-vendors")
		require.NoError(t, err)
		requireFieldEquals(t, out, "vendorID", state.VID)
		requireFieldEquals(t, out, "vendorID", VIDFor1_2)
		requireFieldEquals(t, out, "vendorID", VIDFor1_4_3)
		requireFieldEquals(t, out, "vendorID", VIDFor1_4_4)

		// Model bulk listings.
		out, err = ExecuteCLIWithBin(dcldNew, "query", "model", "all-models")
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", VIDFor1_4_4)

		for _, vid := range []int{state.VID, VIDFor1_2, VIDFor1_4_3, VIDFor1_4_4} {
			_, err = ExecuteCLIWithBin(dcldNew,
				"query", "model", "vendor-models",
				"--vid", fmt.Sprintf("%d", vid),
			)
			require.NoError(t, err)
		}
		_, err = ExecuteCLIWithBin(dcldNew,
			"query", "model", "all-model-versions",
			"--vid", fmt.Sprintf("%d", VIDFor1_4_4),
			"--pid", fmt.Sprintf("%d", PID1For1_4_4),
		)
		require.NoError(t, err)

		// Compliance single-record forms.
		out, err = ExecuteCLIWithBin(dcldNew,
			"query", "compliance", "certified-model",
			"--vid", fmt.Sprintf("%d", VIDFor1_4_4),
			"--pid", fmt.Sprintf("%d", PID1For1_4_4),
			"--softwareVersion", fmt.Sprintf("%d", SoftwareVersionFor1_4_4),
			"--certificationType", CertificationTypeFor1_4_4,
		)
		require.NoError(t, err)
		checkResponseContains(t, out, `"value":true`)

		_, err = ExecuteCLIWithBin(dcldNew,
			"query", "compliance", "revoked-model",
			"--vid", fmt.Sprintf("%d", VIDFor1_4_4),
			"--pid", fmt.Sprintf("%d", PID2For1_4_4),
			"--softwareVersion", fmt.Sprintf("%d", SoftwareVersionFor1_4_4),
			"--certificationType", CertificationTypeFor1_4_4,
		)
		require.NoError(t, err)

		_, err = ExecuteCLIWithBin(dcldNew,
			"query", "compliance", "provisional-model",
			"--vid", fmt.Sprintf("%d", state.VID),
			"--pid", fmt.Sprintf("%d", pid3V012),
			"--softwareVersion", fmt.Sprintf("%d", state.SoftwareVersion),
			"--certificationType", certificationTypeV012,
		)
		require.NoError(t, err)

		_, err = ExecuteCLIWithBin(dcldNew,
			"query", "compliance", "compliance-info",
			"--vid", fmt.Sprintf("%d", VIDFor1_4_4),
			"--pid", fmt.Sprintf("%d", PID1For1_4_4),
			"--softwareVersion", fmt.Sprintf("%d", SoftwareVersionFor1_4_4),
			"--certificationType", CertificationTypeFor1_4_4,
		)
		require.NoError(t, err)

		for _, cdID := range []string{
			cdCertificateIDV012, CDCertificateIDFor1_2, CDCertificateIDFor1_4_3, CDCertificateIDFor1_4_4,
		} {
			out, err = ExecuteCLIWithBin(dcldNew,
				"query", "compliance", "device-software-compliance",
				"--cdCertificateId", cdID,
			)
			require.NoError(t, err)
			checkResponseContains(t, out, cdID)
		}

		// Compliance all-* listings.
		_, err = ExecuteCLIWithBin(dcldNew, "query", "compliance", "all-certified-models")
		require.NoError(t, err)
		_, err = ExecuteCLIWithBin(dcldNew, "query", "compliance", "all-provisional-models")
		require.NoError(t, err)
		_, err = ExecuteCLIWithBin(dcldNew, "query", "compliance", "all-revoked-models")
		require.NoError(t, err)
		_, err = ExecuteCLIWithBin(dcldNew, "query", "compliance", "all-compliance-info")
		require.NoError(t, err)
		_, err = ExecuteCLIWithBin(dcldNew, "query", "compliance", "all-device-software-compliance")
		require.NoError(t, err)

		// PKI single-record forms.
		for _, c := range []struct{ subj, kid string }{
			{RootCertWithVIDSubjectFor1_4_3, RootCertWithVIDSubjectKeyIDFor1_4_3},
			{TestRootCertSubjectFor1_2, TestRootCertSubjectKeyIDFor1_2},
			{testRootCertSubject, testRootCertSubjectKeyID},
		} {
			out, err = ExecuteCLIWithBin(dcldNew,
				"query", "pki", "cert",
				"--subject", c.subj, "--subject-key-id", c.kid,
			)
			require.NoError(t, err)
			checkResponseContains(t, out, c.subj)

			out, err = ExecuteCLIWithBin(dcldNew,
				"query", "pki", "x509-cert",
				"--subject", c.subj, "--subject-key-id", c.kid,
			)
			require.NoError(t, err)
			checkResponseContains(t, out, c.subj)

			_, _ = ExecuteCLIWithBin(dcldNew,
				"query", "pki", "noc-x509-cert",
				"--subject", c.subj, "--subject-key-id", c.kid,
			)
		}

		// Revoked + revocation points.
		_, _ = ExecuteCLIWithBin(dcldNew,
			"query", "pki", "revoked-x509-cert",
			"--subject", IntermediateCertSubjectFor1_2,
			"--subject-key-id", IntermediateCertSubjectKeyIDFor1_2,
		)

		_, _ = ExecuteCLIWithBin(dcldNew,
			"query", "pki", "revoked-noc-x509-root-cert",
			"--subject", NOCRootCert1SubjectFor1_4_3,
			"--subject-key-id", NOCRootCert1SubjectKeyIDFor1_4_3,
		)

		out, err = ExecuteCLIWithBin(dcldNew,
			"query", "pki", "revocation-point",
			"--vid", fmt.Sprintf("%d", VIDFor1_2),
			"--label", ProductLabelFor1_2,
			"--issuer-subject-key-id", IssuerSubjectKeyID,
		)
		require.NoError(t, err)
		checkResponseContains(t, out, IssuerSubjectKeyID)

		_, err = ExecuteCLIWithBin(dcldNew,
			"query", "pki", "revocation-points",
			"--issuer-subject-key-id", IssuerSubjectKeyID,
		)
		require.NoError(t, err)

		// PKI all-* listings.
		_, err = ExecuteCLIWithBin(dcldNew, "query", "pki", "all-certs")
		require.NoError(t, err)
		_, err = ExecuteCLIWithBin(dcldNew, "query", "pki", "all-x509-certs")
		require.NoError(t, err)
		_, err = ExecuteCLIWithBin(dcldNew, "query", "pki", "all-revoked-x509-certs")
		require.NoError(t, err)
		_, err = ExecuteCLIWithBin(dcldNew, "query", "pki", "all-revoked-x509-root-certs")
		require.NoError(t, err)
		_, err = ExecuteCLIWithBin(dcldNew, "query", "pki", "all-noc-x509-certs")
		require.NoError(t, err)
		_, err = ExecuteCLIWithBin(dcldNew, "query", "pki", "all-revoked-noc-x509-root-certs")
		require.NoError(t, err)
		_, err = ExecuteCLIWithBin(dcldNew, "query", "pki", "all-revoked-noc-x509-ica-certs")
		require.NoError(t, err)
		_, err = ExecuteCLIWithBin(dcldNew, "query", "pki", "all-revocation-points")
		require.NoError(t, err)

		// Subject-based listings.
		for _, c := range []struct{ subj string }{
			{RootCertWithVIDSubjectFor1_4_3},
			{TestRootCertSubjectFor1_2},
			{testRootCertSubject},
		} {
			_, _ = ExecuteCLIWithBin(dcldNew,
				"query", "pki", "all-subject-certs", "--subject", c.subj,
			)
			_, _ = ExecuteCLIWithBin(dcldNew,
				"query", "pki", "all-subject-x509-certs", "--subject", c.subj,
			)
			_, _ = ExecuteCLIWithBin(dcldNew,
				"query", "pki", "all-noc-subject-x509-certs", "--subject", c.subj,
			)
		}

		// Validator (host-side).
		if state.ValidatorAddress != "" {
			out, err = ExecuteCLIWithBin(dcldNew, "query", "validator", "all-nodes")
			require.NoError(t, err)
			checkResponseContains(t, out, state.ValidatorAddress)
		}
	})

	// ------------------------------------------------------------------
	// Post-upgrade: seed 1.5.1-era state.
	// ------------------------------------------------------------------
	MustRun(t, "CreateVendor_1_5_1", func(t *testing.T) {
		t.Helper()
		_ = CreateAndApproveAccount(t, dcldNew, VendorAccountFor1_5_1, "Vendor",
			state.VIDFor1_5_1, state.Trustee1,
			[]string{state.Trustee2, state.Trustee3, state.Trustee4})
	})

	MustRun(t, "AddPostUpgradeUserKeys", func(t *testing.T) {
		t.Helper()
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

	MustRun(t, "VendorInfoFor1_5_1", func(t *testing.T) {
		t.Helper()
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

	MustRun(t, "ModelsAndVersionsFor1_5_1", func(t *testing.T) {
		t.Helper()
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

	MustRun(t, "ComplianceFor1_5_1", func(t *testing.T) {
		t.Helper()
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

	MustRun(t, "AccountFlowsFor1_5_1", func(t *testing.T) {
		t.Helper()
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

	MustRun(t, "ValidatorDisableEnableFlow", func(t *testing.T) {
		t.Helper()
		RunValidatorDisableEnableFlow(t, state, dcldNew,
			[]string{state.Trustee2, state.Trustee3, state.Trustee4})
	})

	// ------------------------------------------------------------------
	// Verify post-upgrade-seeded NEW 1.5.1 data. The Phase 1 subtests
	// (08/09) rely on this state being present.
	// ------------------------------------------------------------------
	MustRun(t, "VerifyNew_1_5_1_Data", func(t *testing.T) {
		t.Helper()
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
