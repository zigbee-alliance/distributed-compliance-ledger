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

// runUpgrade143To144 runs the v1.4.3 → v1.4.4 cosmovisor upgrade, then
// seeds 1.4.4-era state: new vendor (VID=65522), DA root certs + their
// intermediates with revocation pairs, NOC certs with the new
// revoke-noc-x509-* commands, and a revocation point.
//
//nolint:funlen
func runUpgrade143To144(t *testing.T, state *UpgradeTestState) {
	t.Helper()

	dcldOld, err := EnsureBinary(BinaryVersionV1_4_3)
	require.NoError(t, err)
	dcldNew, err := EnsureBinary(BinaryVersionV1_4_4)
	require.NoError(t, err)

	step := SoftwareUpgradeStep{
		PlanName:         PlanNameV1_4_4,
		BinaryVersionNew: BinaryVersionV1_4_4,
		Checksum:         UpgradeChecksumV1_4_4,
		DcldOldBin:       dcldOld,
		DcldNewBin:       dcldNew,
		Trustees:         []string{state.Trustee1, state.Trustee2, state.Trustee3, state.Trustee4},
	}
	step.Run(t)

	// ------------------------------------------------------------------
	// Verify carry-over data is intact under v1.4.4.
	// ------------------------------------------------------------------
	MustRun(t, "VerifyPreservedAcrossThreeEras", func(t *testing.T) {
		t.Helper()
		// Spot-check the three vendor-info records.
		for _, vid := range []int{state.VID, VIDFor1_2, VIDFor1_4_3} {
			out, qerr := ExecuteCLIWithBin(dcldNew,
				"query", "vendorinfo", "vendor",
				"--vid", fmt.Sprintf("%d", vid),
			)
			require.NoError(t, qerr)
			requireFieldEquals(t, out, "vendorID", vid)
		}

		// Spot-check key models from each era.
		for _, pair := range [][2]int{
			{state.VID, pid1V012}, {state.VID, state.PID2},
			{VIDFor1_2, PID1For1_2}, {VIDFor1_2, PID2For1_2},
			{VIDFor1_4_3, PID2For1_4_3},
		} {
			out, qerr := ExecuteCLIWithBin(dcldNew,
				"query", "model", "get-model",
				"--vid", fmt.Sprintf("%d", pair[0]),
				"--pid", fmt.Sprintf("%d", pair[1]),
			)
			require.NoError(t, qerr)
			requireFieldEquals(t, out, "vid", pair[0])
			requireFieldEquals(t, out, "pid", pair[1])
		}

		// 0.12 pid_2 now has 1.4.3 productLabel + partNumber (set in script 05).
		out, err := ExecuteCLIWithBin(dcldNew,
			"query", "model", "model-version",
			"--vid", fmt.Sprintf("%d", state.VID),
			"--pid", fmt.Sprintf("%d", state.PID2),
			"--softwareVersion", fmt.Sprintf("%d", state.SoftwareVersion),
		)
		require.NoError(t, err)
		requireFieldEquals(t, out, "minApplicableSoftwareVersion", MinApplicableSoftwareVersionFor1_4_3)
		requireFieldEquals(t, out, "maxApplicableSoftwareVersion", MaxApplicableSoftwareVersionFor1_4_3)
	})

	MustRun(t, "VerifyPreservedAccounts", func(t *testing.T) {
		t.Helper()
		out, err := ExecuteCLIWithBin(dcldNew, "query", "auth", "all-accounts")
		require.NoError(t, err)
		// Active accounts from all prior scripts.
		checkResponseContains(t, out, state.User2Address)
		checkResponseContains(t, out, state.User5Address)
		checkResponseContains(t, out, state.User8Address)

		out, err = ExecuteCLIWithBin(dcldNew, "query", "auth", "all-proposed-accounts")
		require.NoError(t, err)
		checkResponseContains(t, out, state.User3Address)
		checkResponseContains(t, out, state.User6Address)
		checkResponseContains(t, out, state.User9Address)

		out, err = ExecuteCLIWithBin(dcldNew, "query", "auth", "all-proposed-accounts-to-revoke")
		require.NoError(t, err)
		checkResponseContains(t, out, state.User2Address)
		checkResponseContains(t, out, state.User5Address)
		checkResponseContains(t, out, state.User8Address)

		out, err = ExecuteCLIWithBin(dcldNew, "query", "auth", "all-revoked-accounts")
		require.NoError(t, err)
		checkResponseContains(t, out, state.User1Address)
		checkResponseContains(t, out, state.User4Address)
		checkResponseContains(t, out, state.User7Address)

		// Single-record account variants.
		for _, addr := range []string{state.User5Address, state.User2Address, state.User8Address} {
			out, err = ExecuteCLIWithBin(dcldNew, "query", "auth", "account", "--address", addr)
			require.NoError(t, err)
			checkResponseContains(t, out, addr)
		}
		for _, addr := range []string{state.User6Address, state.User3Address, state.User9Address} {
			out, err = ExecuteCLIWithBin(dcldNew, "query", "auth", "proposed-account", "--address", addr)
			require.NoError(t, err)
			checkResponseContains(t, out, addr)
		}
		for _, addr := range []string{state.User5Address, state.User2Address, state.User8Address} {
			out, err = ExecuteCLIWithBin(dcldNew, "query", "auth", "proposed-account-to-revoke", "--address", addr)
			require.NoError(t, err)
			checkResponseContains(t, out, addr)
		}
		for _, addr := range []string{state.User4Address, state.User1Address, state.User7Address} {
			out, err = ExecuteCLIWithBin(dcldNew, "query", "auth", "revoked-account", "--address", addr)
			require.NoError(t, err)
			checkResponseContains(t, out, addr)
		}
	})

	// Bulk readback — gap-fill queries covering vendorinfo/model/compliance/PKI
	// listings + global/DA/NOC cert queries introduced in 1.4.x. NOC-side
	// "Not Found" responses are run for coverage but not asserted on (treated
	// as informational).
	MustRun(t, "VerifyPreservedListings_1_4_4", func(t *testing.T) {
		t.Helper()
		// VendorInfo: all-vendors across three eras.
		out, err := ExecuteCLIWithBin(dcldNew, "query", "vendorinfo", "all-vendors")
		require.NoError(t, err)
		requireFieldEquals(t, out, "vendorID", state.VID)
		requireFieldEquals(t, out, "vendorID", VIDFor1_2)
		requireFieldEquals(t, out, "vendorID", VIDFor1_4_3)

		// Model bulk listings.
		out, err = ExecuteCLIWithBin(dcldNew, "query", "model", "all-models")
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", VIDFor1_4_3)
		requireFieldEquals(t, out, "pid", PID1For1_4_3)

		for _, vid := range []int{state.VID, VIDFor1_2, VIDFor1_4_3} {
			_, err = ExecuteCLIWithBin(dcldNew,
				"query", "model", "vendor-models",
				"--vid", fmt.Sprintf("%d", vid),
			)
			require.NoError(t, err)
		}
		_, err = ExecuteCLIWithBin(dcldNew,
			"query", "model", "all-model-versions",
			"--vid", fmt.Sprintf("%d", VIDFor1_4_3),
			"--pid", fmt.Sprintf("%d", PID1For1_4_3),
		)
		require.NoError(t, err)

		// Compliance single-record forms (gap entries).
		out, err = ExecuteCLIWithBin(dcldNew,
			"query", "compliance", "certified-model",
			"--vid", fmt.Sprintf("%d", VIDFor1_4_3),
			"--pid", fmt.Sprintf("%d", PID1For1_4_3),
			"--softwareVersion", fmt.Sprintf("%d", SoftwareVersionFor1_4_3),
			"--certificationType", CertificationTypeFor1_4_3,
		)
		require.NoError(t, err)
		checkResponseContains(t, out, `"value":true`)

		_, err = ExecuteCLIWithBin(dcldNew,
			"query", "compliance", "revoked-model",
			"--vid", fmt.Sprintf("%d", VIDFor1_4_3),
			"--pid", fmt.Sprintf("%d", PID2For1_4_3),
			"--softwareVersion", fmt.Sprintf("%d", SoftwareVersionFor1_4_3),
			"--certificationType", CertificationTypeFor1_4_3,
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
			"--vid", fmt.Sprintf("%d", VIDFor1_4_3),
			"--pid", fmt.Sprintf("%d", PID1For1_4_3),
			"--softwareVersion", fmt.Sprintf("%d", SoftwareVersionFor1_4_3),
			"--certificationType", CertificationTypeFor1_4_3,
		)
		require.NoError(t, err)

		for _, cdID := range []string{cdCertificateIDV012, CDCertificateIDFor1_2, CDCertificateIDFor1_4_3} {
			out, err = ExecuteCLIWithBin(dcldNew,
				"query", "compliance", "device-software-compliance",
				"--cdCertificateId", cdID,
			)
			require.NoError(t, err)
			checkResponseContains(t, out, cdID)
		}

		// Compliance all-* listings.
		out, err = ExecuteCLIWithBin(dcldNew, "query", "compliance", "all-certified-models")
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", VIDFor1_4_3)
		requireFieldEquals(t, out, "pid", PID1For1_4_3)

		_, err = ExecuteCLIWithBin(dcldNew, "query", "compliance", "all-provisional-models")
		require.NoError(t, err)

		out, err = ExecuteCLIWithBin(dcldNew, "query", "compliance", "all-revoked-models")
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", VIDFor1_4_3)
		requireFieldEquals(t, out, "pid", PID2For1_4_3)

		_, err = ExecuteCLIWithBin(dcldNew, "query", "compliance", "all-compliance-info")
		require.NoError(t, err)

		out, err = ExecuteCLIWithBin(dcldNew, "query", "compliance", "all-device-software-compliance")
		require.NoError(t, err)
		checkResponseContains(t, out, CDCertificateIDFor1_4_3)

		// PKI single-record forms across DA/global namespaces (cert + x509-cert).
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
			checkResponseContains(t, out, c.kid)

			out, err = ExecuteCLIWithBin(dcldNew,
				"query", "pki", "x509-cert",
				"--subject", c.subj, "--subject-key-id", c.kid,
			)
			require.NoError(t, err)
			checkResponseContains(t, out, c.subj)
			checkResponseContains(t, out, c.kid)

			_, err = ExecuteCLIWithBin(dcldNew,
				"query", "pki", "all-subject-certs",
				"--subject", c.subj,
			)
			require.NoError(t, err)

			_, err = ExecuteCLIWithBin(dcldNew,
				"query", "pki", "all-subject-x509-certs",
				"--subject", c.subj,
			)
			require.NoError(t, err)
		}

		// NOC namespace queries return Not Found for DA-side subjects —
		// confirms DA/NOC namespace separation.
		for _, c := range []struct{ subj, kid string }{
			{RootCertWithVIDSubjectFor1_4_3, RootCertWithVIDSubjectKeyIDFor1_4_3},
			{TestRootCertSubjectFor1_2, TestRootCertSubjectKeyIDFor1_2},
		} {
			_, _ = ExecuteCLIWithBin(dcldNew,
				"query", "pki", "noc-x509-cert",
				"--subject", c.subj, "--subject-key-id", c.kid,
			)
			_, _ = ExecuteCLIWithBin(dcldNew,
				"query", "pki", "all-noc-subject-x509-certs",
				"--subject", c.subj,
			)
		}

		// Proposed + revoked + propose-to-revoke (gap entries).
		for _, c := range []struct{ subj, kid string }{
			{GoogleRootCertSubjectFor1_2, GoogleRootCertSubjectKeyIDFor1_2},
			{googleRootCertSubject, googleRootCertSubjectKeyID},
		} {
			out, err = ExecuteCLIWithBin(dcldNew,
				"query", "pki", "proposed-x509-root-cert",
				"--subject", c.subj, "--subject-key-id", c.kid,
			)
			require.NoError(t, err)
			checkResponseContains(t, out, c.subj)
		}

		for _, c := range []struct{ subj, kid string }{
			{IntermediateCertWithVIDSubjectFor1_4_3, IntermediateCertWithVIDSubjectKeyIDFor1_4_3},
			{IntermediateCertSubjectFor1_2, IntermediateCertSubjectKeyIDFor1_2},
			{intermediateCertSubject, intermediateCertSubjectKeyID},
		} {
			out, err = ExecuteCLIWithBin(dcldNew,
				"query", "pki", "revoked-x509-cert",
				"--subject", c.subj, "--subject-key-id", c.kid,
			)
			require.NoError(t, err)
			checkResponseContains(t, out, c.subj)
		}

		for _, c := range []struct{ subj, kid string }{
			{RootCertWithVIDSubjectFor1_4_3, RootCertWithVIDSubjectKeyIDFor1_4_3},
			{TestRootCertSubjectFor1_2, TestRootCertSubjectKeyIDFor1_2},
			{testRootCertSubject, testRootCertSubjectKeyID},
		} {
			out, err = ExecuteCLIWithBin(dcldNew,
				"query", "pki", "proposed-x509-root-cert-to-revoke",
				"--subject", c.subj, "--subject-key-id", c.kid,
			)
			require.NoError(t, err)
			checkResponseContains(t, out, c.subj)
		}

		// Revocation points (single + by-issuer + all).
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

		out, err = ExecuteCLIWithBin(dcldNew, "query", "pki", "all-revocation-points")
		require.NoError(t, err)
		checkResponseContains(t, out, IssuerSubjectKeyID)

		// Global vs DA all-* listings.
		out, err = ExecuteCLIWithBin(dcldNew, "query", "pki", "all-certs")
		require.NoError(t, err)
		checkResponseContains(t, out, RootCertWithVIDSubjectKeyIDFor1_4_3)
		checkResponseContains(t, out, TestRootCertSubjectKeyIDFor1_2)
		checkResponseContains(t, out, testRootCertSubjectKeyID)

		out, err = ExecuteCLIWithBin(dcldNew, "query", "pki", "all-x509-certs")
		require.NoError(t, err)
		checkResponseContains(t, out, RootCertWithVIDSubjectKeyIDFor1_4_3)
		checkResponseContains(t, out, TestRootCertSubjectKeyIDFor1_2)
		checkResponseContains(t, out, testRootCertSubjectKeyID)

		out, err = ExecuteCLIWithBin(dcldNew, "query", "pki", "all-proposed-x509-root-certs")
		require.NoError(t, err)
		checkResponseContains(t, out, GoogleRootCertSubjectFor1_2)
		checkResponseContains(t, out, googleRootCertSubject)

		out, err = ExecuteCLIWithBin(dcldNew, "query", "pki", "all-revoked-x509-root-certs")
		require.NoError(t, err)
		checkResponseContains(t, out, RootCertSubjectFor1_2)
		checkResponseContains(t, out, rootCertSubject)

		out, err = ExecuteCLIWithBin(dcldNew, "query", "pki", "all-proposed-x509-root-certs-to-revoke")
		require.NoError(t, err)
		checkResponseContains(t, out, RootCertWithVIDSubjectFor1_4_3)
		checkResponseContains(t, out, TestRootCertSubjectFor1_2)
		checkResponseContains(t, out, testRootCertSubject)

		out, err = ExecuteCLIWithBin(dcldNew, "query", "pki", "all-revoked-x509-certs")
		require.NoError(t, err)
		checkResponseContains(t, out, IntermediateCertWithVIDSubjectFor1_4_3)

		// NOC-side listings (added in 1.4.3 NOC flow); since NOC certs were
		// added then removed at the script 05 tail, these should be empty —
		// run for coverage and check for absence of removed SKIDs.
		out, err = ExecuteCLIWithBin(dcldNew, "query", "pki", "all-noc-x509-certs")
		require.NoError(t, err)
		require.False(t, strings.Contains(string(out), NOCRootCert1SubjectKeyIDFor1_4_3),
			"NOC root SKID lingered: %s", string(out))

		_, err = ExecuteCLIWithBin(dcldNew,
			"query", "pki", "noc-x509-root-certs",
			"--vid", fmt.Sprintf("%d", VIDFor1_4_3),
		)
		require.NoError(t, err)

		_, err = ExecuteCLIWithBin(dcldNew, "query", "pki", "all-revoked-noc-x509-root-certs")
		require.NoError(t, err)

		_, err = ExecuteCLIWithBin(dcldNew, "query", "pki", "all-revoked-noc-x509-ica-certs")
		require.NoError(t, err)

		_, _ = ExecuteCLIWithBin(dcldNew,
			"query", "pki", "revoked-noc-x509-root-cert",
			"--subject", NOCRootCert1SubjectFor1_4_3,
			"--subject-key-id", NOCRootCert1SubjectKeyIDFor1_4_3,
		)

		// Validator (host-side).
		if state.ValidatorAddress != "" {
			out, err = ExecuteCLIWithBin(dcldNew, "query", "validator", "all-nodes")
			require.NoError(t, err)
			checkResponseContains(t, out, state.ValidatorAddress)
		}
	})

	// ------------------------------------------------------------------
	// Post-upgrade: seed 1.4.4-era state.
	// ------------------------------------------------------------------
	MustRun(t, "CreateVendor_1_4_4", func(t *testing.T) {
		t.Helper()
		_ = CreateAndApproveAccount(t, dcldNew, VendorAccountFor1_4_4, "Vendor",
			VIDFor1_4_4, state.Trustee1,
			[]string{state.Trustee2, state.Trustee3, state.Trustee4})
	})

	MustRun(t, "AddPostUpgradeUserKeys", func(t *testing.T) {
		t.Helper()
		u10, err := newUserKey(dcldNew)
		require.NoError(t, err)
		u11, err := newUserKey(dcldNew)
		require.NoError(t, err)
		u12, err := newUserKey(dcldNew)
		require.NoError(t, err)
		state.User10Address, state.User10Pubkey = u10.address, u10.pubkey
		state.User11Address, state.User11Pubkey = u11.address, u11.pubkey
		state.User12Address, state.User12Pubkey = u12.address, u12.pubkey
	})

	MustRun(t, "VendorInfoFor1_4_4", func(t *testing.T) {
		t.Helper()
		tx, err := ExecuteTxWithBin(dcldNew,
			"tx", "vendorinfo", "add-vendor",
			"--vid", fmt.Sprintf("%d", VIDFor1_4_4),
			"--vendorName", VendorNameFor1_4_4,
			"--companyLegalName", CompanyLegalNameFor1_4_4,
			"--companyPreferredName", CompanyPreferredNameFor1_4_4,
			"--vendorLandingPageURL", VendorLandingPageURLFor1_4_4,
			"--from", VendorAccountFor1_4_4,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)

		tx, err = ExecuteTxWithBin(dcldNew,
			"tx", "vendorinfo", "update-vendor",
			"--vid", fmt.Sprintf("%d", VIDFor1_2),
			"--vendorName", VendorNameFor1_2,
			"--companyLegalName", CompanyLegalNameFor1_2,
			"--companyPreferredName", CompanyPreferredNameFor1_4_4,
			"--vendorLandingPageURL", VendorLandingPageURLFor1_4_4,
			"--from", state.VendorAccountFor1_2,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)
	})

	MustRun(t, "ModelsAndVersionsFor1_4_4", func(t *testing.T) {
		t.Helper()
		for _, pid := range []int{PID1For1_4_4, PID2For1_4_4, PID3For1_4_4} {
			tx, err := ExecuteTxWithBin(dcldNew,
				"tx", "model", "add-model",
				"--vid", fmt.Sprintf("%d", VIDFor1_4_4),
				"--pid", fmt.Sprintf("%d", pid),
				"--deviceTypeID", fmt.Sprintf("%d", DeviceTypeIDFor1_4_4),
				"--productName", ProductNameFor1_4_4,
				"--productLabel", ProductLabelFor1_4_4,
				"--partNumber", PartNumberFor1_4_4,
				"--from", VendorAccountFor1_4_4,
			)
			require.NoError(t, err)
			require.Equal(t, uint32(0), tx.Code, tx.RawLog)

			tx, err = ExecuteTxWithBin(dcldNew,
				"tx", "model", "add-model-version",
				"--vid", fmt.Sprintf("%d", VIDFor1_4_4),
				"--pid", fmt.Sprintf("%d", pid),
				"--softwareVersion", fmt.Sprintf("%d", SoftwareVersionFor1_4_4),
				"--softwareVersionString", SoftwareVersionStringFor1_4_4,
				"--cdVersionNumber", fmt.Sprintf("%d", CDVersionNumberFor1_4_4),
				"--minApplicableSoftwareVersion", fmt.Sprintf("%d", MinApplicableSoftwareVersionFor1_4_4),
				"--maxApplicableSoftwareVersion", fmt.Sprintf("%d", MaxApplicableSoftwareVersionFor1_4_4),
				"--from", VendorAccountFor1_4_4,
			)
			require.NoError(t, err)
			require.Equal(t, uint32(0), tx.Code, tx.RawLog)
		}

		// Delete pid_3.
		tx, err := ExecuteTxWithBin(dcldNew,
			"tx", "model", "delete-model",
			"--vid", fmt.Sprintf("%d", VIDFor1_4_4),
			"--pid", fmt.Sprintf("%d", PID3For1_4_4),
			"--from", VendorAccountFor1_4_4,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)

		// Update carry-over 0.12 pid_2.
		tx, err = ExecuteTxWithBin(dcldNew,
			"tx", "model", "update-model",
			"--vid", fmt.Sprintf("%d", state.VID),
			"--pid", fmt.Sprintf("%d", state.PID2),
			"--productName", state.ProductName,
			"--productLabel", ProductLabelFor1_4_4,
			"--partNumber", PartNumberFor1_4_4,
			"--from", state.VendorAccount,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)

		tx, err = ExecuteTxWithBin(dcldNew,
			"tx", "model", "update-model-version",
			"--vid", fmt.Sprintf("%d", state.VID),
			"--pid", fmt.Sprintf("%d", state.PID2),
			"--softwareVersion", fmt.Sprintf("%d", state.SoftwareVersion),
			"--minApplicableSoftwareVersion", fmt.Sprintf("%d", MinApplicableSoftwareVersionFor1_4_4),
			"--maxApplicableSoftwareVersion", fmt.Sprintf("%d", MaxApplicableSoftwareVersionFor1_4_4),
			"--from", state.VendorAccount,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)
	})

	MustRun(t, "ComplianceFor1_4_4", func(t *testing.T) {
		t.Helper()
		// certify pid_1
		tx, err := ExecuteTxWithBin(dcldNew,
			"tx", "compliance", "certify-model",
			"--vid", fmt.Sprintf("%d", VIDFor1_4_4),
			"--pid", fmt.Sprintf("%d", PID1For1_4_4),
			"--softwareVersion", fmt.Sprintf("%d", SoftwareVersionFor1_4_4),
			"--softwareVersionString", SoftwareVersionStringFor1_4_4,
			"--certificationType", CertificationTypeFor1_4_4,
			"--certificationDate", CertificationDateFor1_4_4,
			"--cdCertificateId", CDCertificateIDFor1_4_4,
			"--cdVersionNumber", fmt.Sprintf("%d", CDVersionNumberFor1_4_4),
			"--from", CertificationCenterAccountFor1_2,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)

		// provision pid_2
		tx, err = ExecuteTxWithBin(dcldNew,
			"tx", "compliance", "provision-model",
			"--vid", fmt.Sprintf("%d", VIDFor1_4_4),
			"--pid", fmt.Sprintf("%d", PID2For1_4_4),
			"--softwareVersion", fmt.Sprintf("%d", SoftwareVersionFor1_4_4),
			"--softwareVersionString", SoftwareVersionStringFor1_4_4,
			"--certificationType", CertificationTypeFor1_4_4,
			"--provisionalDate", ProvisionalDateFor1_4_4,
			"--cdCertificateId", CDCertificateIDFor1_4_4,
			"--cdVersionNumber", fmt.Sprintf("%d", CDVersionNumberFor1_4_4),
			"--from", CertificationCenterAccountFor1_2,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)

		// certify pid_2
		tx, err = ExecuteTxWithBin(dcldNew,
			"tx", "compliance", "certify-model",
			"--vid", fmt.Sprintf("%d", VIDFor1_4_4),
			"--pid", fmt.Sprintf("%d", PID2For1_4_4),
			"--softwareVersion", fmt.Sprintf("%d", SoftwareVersionFor1_4_4),
			"--softwareVersionString", SoftwareVersionStringFor1_4_4,
			"--certificationType", CertificationTypeFor1_4_4,
			"--certificationDate", CertificationDateFor1_4_4,
			"--cdCertificateId", CDCertificateIDFor1_4_4,
			"--cdVersionNumber", fmt.Sprintf("%d", CDVersionNumberFor1_4_4),
			"--from", CertificationCenterAccountFor1_2,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)

		// revoke pid_2
		tx, err = ExecuteTxWithBin(dcldNew,
			"tx", "compliance", "revoke-model",
			"--vid", fmt.Sprintf("%d", VIDFor1_4_4),
			"--pid", fmt.Sprintf("%d", PID2For1_4_4),
			"--softwareVersion", fmt.Sprintf("%d", SoftwareVersionFor1_4_4),
			"--softwareVersionString", SoftwareVersionStringFor1_4_4,
			"--certificationType", CertificationTypeFor1_4_4,
			"--revocationDate", CertificationDateFor1_4_4,
			"--cdVersionNumber", fmt.Sprintf("%d", CDVersionNumberFor1_4_4),
			"--from", CertificationCenterAccountFor1_2,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)
	})

	MustRun(t, "PKIFor1_4_4_DARootCerts", func(t *testing.T) {
		t.Helper()
		tx, err := ExecuteTxWithBin(dcldNew,
			"tx", "pki", "propose-add-x509-root-cert",
			"--certificate", DARootCert1PathFor1_4_4,
			"--vid", fmt.Sprintf("%d", VIDFor1_4_4),
			"--from", state.Trustee1,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)

		tx, err = ExecuteTxWithBin(dcldNew,
			"tx", "pki", "approve-add-x509-root-cert",
			"--subject", DARootCert1SubjectFor1_4_4,
			"--subject-key-id", DARootCert1SubjectKeyIDFor1_4_4,
			"--from", state.Trustee2,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)

		tx, err = ExecuteTxWithBin(dcldNew,
			"tx", "pki", "reject-add-x509-root-cert",
			"--subject", DARootCert1SubjectFor1_4_4,
			"--subject-key-id", DARootCert1SubjectKeyIDFor1_4_4,
			"--from", state.Trustee3,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)

		tx, err = ExecuteTxWithBin(dcldNew,
			"tx", "pki", "approve-add-x509-root-cert",
			"--subject", DARootCert1SubjectFor1_4_4,
			"--subject-key-id", DARootCert1SubjectKeyIDFor1_4_4,
			"--from", state.Trustee4,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)

		tx, err = ExecuteTxWithBin(dcldNew,
			"tx", "pki", "approve-add-x509-root-cert",
			"--subject", DARootCert1SubjectFor1_4_4,
			"--subject-key-id", DARootCert1SubjectKeyIDFor1_4_4,
			"--from", state.Trustee5,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)

		// da_root_cert_2: propose by t1, approve by t1/t2/t3.
		tx, err = ExecuteTxWithBin(dcldNew,
			"tx", "pki", "propose-add-x509-root-cert",
			"--certificate", DARootCert2PathFor1_4_4,
			"--vid", fmt.Sprintf("%d", VIDFor1_4_4),
			"--from", state.Trustee1,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)

		for _, who := range []string{state.Trustee2, state.Trustee3, state.Trustee4} {
			tx, err = ExecuteTxWithBin(dcldNew,
				"tx", "pki", "approve-add-x509-root-cert",
				"--subject", DARootCert2SubjectFor1_4_4,
				"--subject-key-id", DARootCert2SubjectKeyIDFor1_4_4,
				"--from", who,
			)
			require.NoError(t, err)
			require.Equal(t, uint32(0), tx.Code, tx.RawLog)
		}

		// Add intermediates and revoke da_root_cert_1 + da_intermediate_cert_1.
		for _, certPath := range []string{
			DAIntermediateCert1PathFor1_4_4, DAIntermediateCert2PathFor1_4_4,
		} {
			tx, err = ExecuteTxWithBin(dcldNew,
				"tx", "pki", "add-x509-cert",
				"--certificate", certPath,
				"--from", VendorAccountFor1_4_4,
			)
			require.NoError(t, err)
			require.Equal(t, uint32(0), tx.Code, tx.RawLog)
		}

		// Propose-revoke + 3 approves of da_root_cert_1.
		tx, err = ExecuteTxWithBin(dcldNew,
			"tx", "pki", "propose-revoke-x509-root-cert",
			"--subject", DARootCert1SubjectFor1_4_4,
			"--subject-key-id", DARootCert1SubjectKeyIDFor1_4_4,
			"--from", state.Trustee1,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)
		for _, who := range []string{state.Trustee2, state.Trustee3, state.Trustee4} {
			tx, err = ExecuteTxWithBin(dcldNew,
				"tx", "pki", "approve-revoke-x509-root-cert",
				"--subject", DARootCert1SubjectFor1_4_4,
				"--subject-key-id", DARootCert1SubjectKeyIDFor1_4_4,
				"--from", who,
			)
			require.NoError(t, err)
			require.Equal(t, uint32(0), tx.Code, tx.RawLog)
		}

		// Propose-revoke da_root_cert_2 (no approvals).
		tx, err = ExecuteTxWithBin(dcldNew,
			"tx", "pki", "propose-revoke-x509-root-cert",
			"--subject", DARootCert2SubjectFor1_4_4,
			"--subject-key-id", DARootCert2SubjectKeyIDFor1_4_4,
			"--from", state.Trustee1,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)

		// Revoke da_intermediate_cert_1.
		tx, err = ExecuteTxWithBin(dcldNew,
			"tx", "pki", "revoke-x509-cert",
			"--subject", DAIntermediateCert1SubjectFor1_4_4,
			"--subject-key-id", DAIntermediateCert1SubjectKeyIDFor1_4_4,
			"--from", VendorAccountFor1_4_4,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)
	})

	MustRun(t, "NOCCertsAddRevoke", func(t *testing.T) {
		t.Helper()
		// 1.4.4 introduces `revoke-noc-x509-{root,ica}-cert` (vs 1.4.3's
		// `remove-noc-x509-*`). Add 2 root/ICA pairs, then revoke pair #1
		// only — pair #2 stays active.
		for _, pair := range []struct{ rootPath, icaPath string }{
			{NOCRootCert1V144PathFor1_4_4, NOCICACert1V144PathFor1_4_4},
			{NOCRootCert2V144PathFor1_4_4, NOCICACert2V144PathFor1_4_4},
		} {
			tx, err := ExecuteTxWithBin(dcldNew,
				"tx", "pki", "add-noc-x509-root-cert",
				"--certificate", pair.rootPath,
				"--from", VendorAccountFor1_4_4,
			)
			require.NoError(t, err)
			require.Equal(t, uint32(0), tx.Code, tx.RawLog)

			tx, err = ExecuteTxWithBin(dcldNew,
				"tx", "pki", "add-noc-x509-ica-cert",
				"--certificate", pair.icaPath,
				"--from", VendorAccountFor1_4_4,
			)
			require.NoError(t, err)
			require.Equal(t, uint32(0), tx.Code, tx.RawLog)
		}

		// Revoke NOC pair #1.
		tx, err := ExecuteTxWithBin(dcldNew,
			"tx", "pki", "revoke-noc-x509-root-cert",
			"--subject", NOCRootCert1V144SubjectFor1_4_4,
			"--subject-key-id", NOCRootCert1V144SubjectKeyIDFor1_4_4,
			"--from", VendorAccountFor1_4_4,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)

		tx, err = ExecuteTxWithBin(dcldNew,
			"tx", "pki", "revoke-noc-x509-ica-cert",
			"--subject", NOCICACert1V144SubjectFor1_4_4,
			"--subject-key-id", NOCICACert1V144SubjectKeyIDFor1_4_4,
			"--from", VendorAccountFor1_4_4,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)
	})

	MustRun(t, "RevocationPointsFor1_4_4", func(t *testing.T) {
		t.Helper()
		// add → update → delete → add (one active PAA revocation point at end).
		addPAA := func(dataURL string) {
			tx, err := ExecuteTxWithBin(dcldNew,
				"tx", "pki", "add-revocation-point",
				"--vid", fmt.Sprintf("%d", VIDFor1_4_4),
				"--revocation-type", "1",
				"--is-paa=true",
				"--certificate", DARootCert2PathFor1_4_4,
				"--label", ProductLabelFor1_4_4,
				"--data-url", dataURL,
				"--issuer-subject-key-id", IssuerSubjectKeyID,
				"--from", VendorAccountFor1_4_4,
			)
			require.NoError(t, err)
			require.Equal(t, uint32(0), tx.Code, tx.RawLog)
		}

		addPAA(TestDataURLFor1_4_4)

		tx, err := ExecuteTxWithBin(dcldNew,
			"tx", "pki", "update-revocation-point",
			"--vid", fmt.Sprintf("%d", VIDFor1_4_4),
			"--certificate", DARootCert2PathFor1_4_4,
			"--label", ProductLabelFor1_4_4,
			"--data-url", TestDataURLFor1_4_4+"/new",
			"--issuer-subject-key-id", IssuerSubjectKeyID,
			"--from", VendorAccountFor1_4_4,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)

		tx, err = ExecuteTxWithBin(dcldNew,
			"tx", "pki", "delete-revocation-point",
			"--vid", fmt.Sprintf("%d", VIDFor1_4_4),
			"--label", ProductLabelFor1_4_4,
			"--issuer-subject-key-id", IssuerSubjectKeyID,
			"--from", VendorAccountFor1_4_4,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)

		addPAA(TestDataURLFor1_4_4)
	})

	MustRun(t, "AccountFlowsFor1_4_4", func(t *testing.T) {
		t.Helper()
		approvers := []string{state.Trustee2, state.Trustee3, state.Trustee4}

		proposeUserAccount(t, dcldNew, state.Trustee1, approvers,
			state.User10Address, state.User10Pubkey, "CertificationCenter", true)
		proposeUserAccount(t, dcldNew, state.Trustee1, approvers,
			state.User11Address, state.User11Pubkey, "CertificationCenter", true)
		proposeUserAccount(t, dcldNew, state.Trustee1, nil,
			state.User12Address, state.User12Pubkey, "CertificationCenter", false)

		revokeUserAccount(t, dcldNew, state.Trustee1, approvers, state.User10Address, true)
		revokeUserAccount(t, dcldNew, state.Trustee1, nil, state.User11Address, false)
	})

	// Validator disable/enable (lines 1189-1238) — Docker-dependent, stubbed.
	MustRun(t, "ValidatorDisableEnableFlow", func(t *testing.T) {
		t.Helper()
		RunValidatorDisableEnableFlow(t, state, dcldNew,
			[]string{state.Trustee2, state.Trustee3, state.Trustee4})
	})

	// ------------------------------------------------------------------
	// Verify post-upgrade-seeded NEW data.
	// ------------------------------------------------------------------
	MustRun(t, "VerifyNew_1_4_4_Data", func(t *testing.T) {
		t.Helper()
		out, err := ExecuteCLIWithBin(dcldNew,
			"query", "vendorinfo", "vendor",
			"--vid", fmt.Sprintf("%d", VIDFor1_4_4),
		)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vendorID", VIDFor1_4_4)
		checkResponseContains(t, out, CompanyLegalNameFor1_4_4)

		out, err = ExecuteCLIWithBin(dcldNew,
			"query", "model", "get-model",
			"--vid", fmt.Sprintf("%d", VIDFor1_4_4),
			"--pid", fmt.Sprintf("%d", PID1For1_4_4),
		)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", VIDFor1_4_4)
		requireFieldEquals(t, out, "pid", PID1For1_4_4)
		checkResponseContains(t, out, ProductLabelFor1_4_4)

		// 0.12 pid_2 now has 1.4.4 productLabel/partNumber.
		out, err = ExecuteCLIWithBin(dcldNew,
			"query", "model", "get-model",
			"--vid", fmt.Sprintf("%d", state.VID),
			"--pid", fmt.Sprintf("%d", state.PID2),
		)
		require.NoError(t, err)
		checkResponseContains(t, out, ProductLabelFor1_4_4)
		checkResponseContains(t, out, PartNumberFor1_4_4)
	})
}
