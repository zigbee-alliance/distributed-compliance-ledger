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
	cliputils "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/cli/utils"
)

// runRollback12 is the Go translation of
// integration_tests/upgrade/04-test-upgrade-1.2-rollback.sh.
//
// Same shape as runRollback012, but submitted against a v1.2 chain with
// "wrong_plan_name_2" — verifying the chain doesn't accidentally upgrade to
// v1.4.3 just because a plan with that target binary was approved.
//
//nolint:funlen
func runRollback12(t *testing.T, state *UpgradeTestState) {
	t.Helper()

	dcld, err := EnsureBinary(BinaryVersionV1_2)
	require.NoError(t, err)

	// ------------------------------------------------------------------
	// Wrong-plan-name upgrade attempt.
	// ------------------------------------------------------------------
	MustRun(t, "WrongPlanName2Upgrade", func(t *testing.T) {
		currentHeight, err := cliputils.GetHeight()
		require.NoError(t, err)
		planHeight := currentHeight + 20

		upgradeInfo := UpgradeInfoForVersion(BinaryVersionV1_4_3, WrongPlanChecksumV143)

		tx, err := ProposeUpgrade(dcld, WrongPlanName2, planHeight, upgradeInfo, state.Trustee1)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)

		for _, who := range []string{state.Trustee2, state.Trustee3, state.Trustee4} {
			tx, err = ApproveUpgrade(dcld, WrongPlanName2, who)
			require.NoError(t, err)
			require.Equal(t, uint32(0), tx.Code, "approve %s: %s", who, tx.RawLog)
		}

		cliputils.WaitForHeight(t, planHeight+1, 300)

		out, _ := QueryUpgradePlan(dcld)
		require.True(t, strings.Contains(string(out), "no upgrade scheduled"),
			"expected 'no upgrade scheduled', got: %s", string(out))

		out, err = QueryAppliedPlan(dcld, WrongPlanName2)
		if err == nil {
			require.False(t, strings.Contains(string(out), `"height"`),
				"upgrade unexpectedly applied: %s", string(out))
		}
	})

	// ------------------------------------------------------------------
	// Verify carry-over from scripts 01/02/03 is intact. Mirrors the
	// post-rollback readback in 04-test-upgrade-1.2-rollback.sh lines 83-454.
	// ------------------------------------------------------------------
	MustRun(t, "VerifyPreservedAfterRollback1_2", func(t *testing.T) {
		// ----- VendorInfo -----
		out, err := ExecuteCLIWithBin(dcld,
			"query", "vendorinfo", "vendor",
			"--vid", fmt.Sprintf("%d", state.VID),
		)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vendorID", state.VID)
		checkResponseContains(t, out, companyLegalNameV012)
		checkResponseContains(t, out, vendorNameV012)
		checkResponseContains(t, out, CompanyPreferredNameFor1_2)
		checkResponseContains(t, out, VendorLandingPageURLFor1_2)

		out, err = ExecuteCLIWithBin(dcld,
			"query", "vendorinfo", "vendor",
			"--vid", fmt.Sprintf("%d", VIDFor1_2),
		)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vendorID", VIDFor1_2)
		checkResponseContains(t, out, CompanyLegalNameFor1_2)

		out, err = ExecuteCLIWithBin(dcld, "query", "vendorinfo", "all-vendors")
		require.NoError(t, err)
		requireFieldEquals(t, out, "vendorID", state.VID)
		requireFieldEquals(t, out, "vendorID", VIDFor1_2)
		checkResponseContains(t, out, companyLegalNameV012)
		checkResponseContains(t, out, CompanyLegalNameFor1_2)

		// ----- Model: 0.12-era pid_1 + pid_2 -----
		out, err = ExecuteCLIWithBin(dcld,
			"query", "model", "get-model",
			"--vid", fmt.Sprintf("%d", state.VID),
			"--pid", fmt.Sprintf("%d", pid1V012),
		)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", state.VID)
		requireFieldEquals(t, out, "pid", pid1V012)
		checkResponseContains(t, out, state.ProductLabel)

		out, err = ExecuteCLIWithBin(dcld,
			"query", "model", "get-model",
			"--vid", fmt.Sprintf("%d", state.VID),
			"--pid", fmt.Sprintf("%d", state.PID2),
		)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", state.VID)
		requireFieldEquals(t, out, "pid", state.PID2)
		checkResponseContains(t, out, ProductLabelFor1_2)
		checkResponseContains(t, out, PartNumberFor1_2)

		// 1.2-era pid_1 + pid_2.
		for _, pid := range []int{PID1For1_2, PID2For1_2} {
			out, err = ExecuteCLIWithBin(dcld,
				"query", "model", "get-model",
				"--vid", fmt.Sprintf("%d", VIDFor1_2),
				"--pid", fmt.Sprintf("%d", pid),
			)
			require.NoError(t, err)
			requireFieldEquals(t, out, "vid", VIDFor1_2)
			requireFieldEquals(t, out, "pid", pid)
			checkResponseContains(t, out, ProductLabelFor1_2)
		}

		out, err = ExecuteCLIWithBin(dcld, "query", "model", "all-models")
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", state.VID)
		requireFieldEquals(t, out, "pid", pid1V012)
		requireFieldEquals(t, out, "pid", state.PID2)
		requireFieldEquals(t, out, "vid", VIDFor1_2)
		requireFieldEquals(t, out, "pid", PID1For1_2)
		requireFieldEquals(t, out, "pid", PID2For1_2)

		out, err = ExecuteCLIWithBin(dcld,
			"query", "model", "vendor-models",
			"--vid", fmt.Sprintf("%d", state.VID),
		)
		require.NoError(t, err)
		requireFieldEquals(t, out, "pid", pid1V012)
		requireFieldEquals(t, out, "pid", state.PID2)

		out, err = ExecuteCLIWithBin(dcld,
			"query", "model", "vendor-models",
			"--vid", fmt.Sprintf("%d", VIDFor1_2),
		)
		require.NoError(t, err)
		requireFieldEquals(t, out, "pid", PID1For1_2)
		requireFieldEquals(t, out, "pid", PID2For1_2)

		// Model versions.
		for _, pid := range []int{pid1V012, state.PID2} {
			out, err = ExecuteCLIWithBin(dcld,
				"query", "model", "model-version",
				"--vid", fmt.Sprintf("%d", state.VID),
				"--pid", fmt.Sprintf("%d", pid),
				"--softwareVersion", fmt.Sprintf("%d", state.SoftwareVersion),
			)
			require.NoError(t, err)
			requireFieldEquals(t, out, "vid", state.VID)
			requireFieldEquals(t, out, "pid", pid)
			requireFieldEquals(t, out, "softwareVersion", state.SoftwareVersion)
		}
		for _, pid := range []int{PID1For1_2, PID2For1_2} {
			out, err = ExecuteCLIWithBin(dcld,
				"query", "model", "model-version",
				"--vid", fmt.Sprintf("%d", VIDFor1_2),
				"--pid", fmt.Sprintf("%d", pid),
				"--softwareVersion", fmt.Sprintf("%d", SoftwareVersionFor1_2),
			)
			require.NoError(t, err)
			requireFieldEquals(t, out, "vid", VIDFor1_2)
			requireFieldEquals(t, out, "pid", pid)
			requireFieldEquals(t, out, "softwareVersion", SoftwareVersionFor1_2)
		}

		// ----- Compliance -----
		out, err = ExecuteCLIWithBin(dcld,
			"query", "compliance", "certified-model",
			"--vid", fmt.Sprintf("%d", state.VID),
			"--pid", fmt.Sprintf("%d", pid1V012),
			"--softwareVersion", fmt.Sprintf("%d", state.SoftwareVersion),
			"--certificationType", certificationTypeV012,
		)
		require.NoError(t, err)
		checkResponseContains(t, out, `"value":true`)

		out, err = ExecuteCLIWithBin(dcld,
			"query", "compliance", "certified-model",
			"--vid", fmt.Sprintf("%d", VIDFor1_2),
			"--pid", fmt.Sprintf("%d", PID1For1_2),
			"--softwareVersion", fmt.Sprintf("%d", SoftwareVersionFor1_2),
			"--certificationType", CertificationTypeFor1_2,
		)
		require.NoError(t, err)
		checkResponseContains(t, out, `"value":true`)

		out, err = ExecuteCLIWithBin(dcld,
			"query", "compliance", "revoked-model",
			"--vid", fmt.Sprintf("%d", state.VID),
			"--pid", fmt.Sprintf("%d", state.PID2),
			"--softwareVersion", fmt.Sprintf("%d", state.SoftwareVersion),
			"--certificationType", certificationTypeV012,
		)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", state.VID)
		requireFieldEquals(t, out, "pid", state.PID2)

		out, err = ExecuteCLIWithBin(dcld,
			"query", "compliance", "revoked-model",
			"--vid", fmt.Sprintf("%d", VIDFor1_2),
			"--pid", fmt.Sprintf("%d", PID2For1_2),
			"--softwareVersion", fmt.Sprintf("%d", SoftwareVersionFor1_2),
			"--certificationType", CertificationTypeFor1_2,
		)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", VIDFor1_2)
		requireFieldEquals(t, out, "pid", PID2For1_2)

		out, err = ExecuteCLIWithBin(dcld,
			"query", "compliance", "provisional-model",
			"--vid", fmt.Sprintf("%d", state.VID),
			"--pid", fmt.Sprintf("%d", pid3V012),
			"--softwareVersion", fmt.Sprintf("%d", state.SoftwareVersion),
			"--certificationType", certificationTypeV012,
		)
		require.NoError(t, err)
		checkResponseContains(t, out, `"value":true`)

		for _, pid := range []int{pid1V012, state.PID2} {
			out, err = ExecuteCLIWithBin(dcld,
				"query", "compliance", "compliance-info",
				"--vid", fmt.Sprintf("%d", state.VID),
				"--pid", fmt.Sprintf("%d", pid),
				"--softwareVersion", fmt.Sprintf("%d", state.SoftwareVersion),
				"--certificationType", certificationTypeV012,
			)
			require.NoError(t, err)
			requireFieldEquals(t, out, "vid", state.VID)
			requireFieldEquals(t, out, "pid", pid)
		}
		for _, pid := range []int{PID1For1_2, PID2For1_2} {
			out, err = ExecuteCLIWithBin(dcld,
				"query", "compliance", "compliance-info",
				"--vid", fmt.Sprintf("%d", VIDFor1_2),
				"--pid", fmt.Sprintf("%d", pid),
				"--softwareVersion", fmt.Sprintf("%d", SoftwareVersionFor1_2),
				"--certificationType", CertificationTypeFor1_2,
			)
			require.NoError(t, err)
			requireFieldEquals(t, out, "vid", VIDFor1_2)
			requireFieldEquals(t, out, "pid", pid)
		}

		out, err = ExecuteCLIWithBin(dcld,
			"query", "compliance", "device-software-compliance",
			"--cdCertificateId", cdCertificateIDV012,
		)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", state.VID)
		requireFieldEquals(t, out, "pid", pid1V012)

		out, err = ExecuteCLIWithBin(dcld,
			"query", "compliance", "device-software-compliance",
			"--cdCertificateId", CDCertificateIDFor1_2,
		)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", VIDFor1_2)
		requireFieldEquals(t, out, "pid", PID1For1_2)

		// Compliance all-* listings.
		out, err = ExecuteCLIWithBin(dcld, "query", "compliance", "all-certified-models")
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", state.VID)
		requireFieldEquals(t, out, "pid", pid1V012)
		requireFieldEquals(t, out, "vid", VIDFor1_2)
		requireFieldEquals(t, out, "pid", PID1For1_2)

		out, err = ExecuteCLIWithBin(dcld, "query", "compliance", "all-provisional-models")
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", state.VID)
		requireFieldEquals(t, out, "pid", pid3V012)

		out, err = ExecuteCLIWithBin(dcld, "query", "compliance", "all-revoked-models")
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", state.VID)
		requireFieldEquals(t, out, "pid", state.PID2)
		requireFieldEquals(t, out, "vid", VIDFor1_2)
		requireFieldEquals(t, out, "pid", PID2For1_2)

		out, err = ExecuteCLIWithBin(dcld, "query", "compliance", "all-compliance-info")
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", state.VID)
		requireFieldEquals(t, out, "pid", pid1V012)
		requireFieldEquals(t, out, "pid", state.PID2)
		requireFieldEquals(t, out, "vid", VIDFor1_2)
		requireFieldEquals(t, out, "pid", PID1For1_2)
		requireFieldEquals(t, out, "pid", PID2For1_2)

		out, err = ExecuteCLIWithBin(dcld, "query", "compliance", "all-device-software-compliance")
		require.NoError(t, err)
		checkResponseContains(t, out, cdCertificateIDV012)
		checkResponseContains(t, out, CDCertificateIDFor1_2)

		// ----- PKI single-record + listings -----
		// 1.2-era root_cert: test_root_cert assigned vid in script 03's assign-vid.
		out, err = ExecuteCLIWithBin(dcld,
			"query", "pki", "x509-cert",
			"--subject", TestRootCertSubjectFor1_2,
			"--subject-key-id", TestRootCertSubjectKeyIDFor1_2,
		)
		require.NoError(t, err)
		checkResponseContains(t, out, TestRootCertSubjectFor1_2)
		checkResponseContains(t, out, TestRootCertSubjectKeyIDFor1_2)

		// 0.12-era test_root_cert.
		out, err = ExecuteCLIWithBin(dcld,
			"query", "pki", "x509-cert",
			"--subject", testRootCertSubject,
			"--subject-key-id", testRootCertSubjectKeyID,
		)
		require.NoError(t, err)
		checkResponseContains(t, out, testRootCertSubject)
		checkResponseContains(t, out, testRootCertSubjectKeyID)

		out, err = ExecuteCLIWithBin(dcld,
			"query", "pki", "all-subject-x509-certs",
			"--subject", TestRootCertSubjectFor1_2,
		)
		require.NoError(t, err)
		checkResponseContains(t, out, TestRootCertSubjectFor1_2)
		checkResponseContains(t, out, TestRootCertSubjectKeyIDFor1_2)

		out, err = ExecuteCLIWithBin(dcld,
			"query", "pki", "all-subject-x509-certs",
			"--subject", testRootCertSubject,
		)
		require.NoError(t, err)
		checkResponseContains(t, out, testRootCertSubject)
		checkResponseContains(t, out, testRootCertSubjectKeyID)

		out, err = ExecuteCLIWithBin(dcld,
			"query", "pki", "proposed-x509-root-cert",
			"--subject", GoogleRootCertSubjectFor1_2,
			"--subject-key-id", GoogleRootCertSubjectKeyIDFor1_2,
		)
		require.NoError(t, err)
		checkResponseContains(t, out, GoogleRootCertSubjectFor1_2)
		checkResponseContains(t, out, GoogleRootCertSubjectKeyIDFor1_2)

		out, err = ExecuteCLIWithBin(dcld,
			"query", "pki", "proposed-x509-root-cert",
			"--subject", googleRootCertSubject,
			"--subject-key-id", googleRootCertSubjectKeyID,
		)
		require.NoError(t, err)
		checkResponseContains(t, out, googleRootCertSubject)
		checkResponseContains(t, out, googleRootCertSubjectKeyID)

		// Revoked intermediate certs (v0.12 + v1.2).
		out, err = ExecuteCLIWithBin(dcld,
			"query", "pki", "revoked-x509-cert",
			"--subject", IntermediateCertSubjectFor1_2,
			"--subject-key-id", IntermediateCertSubjectKeyIDFor1_2,
		)
		require.NoError(t, err)
		checkResponseContains(t, out, IntermediateCertSubjectFor1_2)
		checkResponseContains(t, out, IntermediateCertSubjectKeyIDFor1_2)

		out, err = ExecuteCLIWithBin(dcld,
			"query", "pki", "revoked-x509-cert",
			"--subject", intermediateCertSubject,
			"--subject-key-id", intermediateCertSubjectKeyID,
		)
		require.NoError(t, err)
		checkResponseContains(t, out, intermediateCertSubject)
		checkResponseContains(t, out, intermediateCertSubjectKeyID)

		// Proposed-to-revoke (both eras).
		out, err = ExecuteCLIWithBin(dcld,
			"query", "pki", "proposed-x509-root-cert-to-revoke",
			"--subject", TestRootCertSubjectFor1_2,
			"--subject-key-id", TestRootCertSubjectKeyIDFor1_2,
		)
		require.NoError(t, err)
		checkResponseContains(t, out, TestRootCertSubjectFor1_2)
		checkResponseContains(t, out, TestRootCertSubjectKeyIDFor1_2)

		out, err = ExecuteCLIWithBin(dcld,
			"query", "pki", "proposed-x509-root-cert-to-revoke",
			"--subject", testRootCertSubject,
			"--subject-key-id", testRootCertSubjectKeyID,
		)
		require.NoError(t, err)
		checkResponseContains(t, out, testRootCertSubject)
		checkResponseContains(t, out, testRootCertSubjectKeyID)

		// Revocation point (single + listing by issuer + all).
		out, err = ExecuteCLIWithBin(dcld,
			"query", "pki", "revocation-point",
			"--vid", fmt.Sprintf("%d", VIDFor1_2),
			"--label", ProductLabelFor1_2,
			"--issuer-subject-key-id", IssuerSubjectKeyID,
		)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", VIDFor1_2)
		checkResponseContains(t, out, IssuerSubjectKeyID)
		checkResponseContains(t, out, ProductLabelFor1_2)
		checkResponseContains(t, out, TestDataURL)

		out, err = ExecuteCLIWithBin(dcld,
			"query", "pki", "revocation-points",
			"--issuer-subject-key-id", IssuerSubjectKeyID,
		)
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", VIDFor1_2)
		checkResponseContains(t, out, IssuerSubjectKeyID)
		checkResponseContains(t, out, ProductLabelFor1_2)

		out, err = ExecuteCLIWithBin(dcld, "query", "pki", "all-revocation-points")
		require.NoError(t, err)
		requireFieldEquals(t, out, "vid", VIDFor1_2)
		checkResponseContains(t, out, IssuerSubjectKeyID)

		// All-* PKI listings.
		out, err = ExecuteCLIWithBin(dcld, "query", "pki", "all-proposed-x509-root-certs")
		require.NoError(t, err)
		checkResponseContains(t, out, GoogleRootCertSubjectFor1_2)
		checkResponseContains(t, out, GoogleRootCertSubjectKeyIDFor1_2)
		checkResponseContains(t, out, googleRootCertSubject)
		checkResponseContains(t, out, googleRootCertSubjectKeyID)

		out, err = ExecuteCLIWithBin(dcld, "query", "pki", "all-revoked-x509-root-certs")
		require.NoError(t, err)
		checkResponseContains(t, out, RootCertSubjectFor1_2)
		checkResponseContains(t, out, RootCertSubjectKeyIDFor1_2)
		checkResponseContains(t, out, rootCertSubject)
		checkResponseContains(t, out, rootCertSubjectKeyID)

		out, err = ExecuteCLIWithBin(dcld, "query", "pki", "all-proposed-x509-root-certs-to-revoke")
		require.NoError(t, err)
		checkResponseContains(t, out, TestRootCertSubjectFor1_2)
		checkResponseContains(t, out, TestRootCertSubjectKeyIDFor1_2)
		checkResponseContains(t, out, testRootCertSubject)
		checkResponseContains(t, out, testRootCertSubjectKeyID)

		out, err = ExecuteCLIWithBin(dcld, "query", "pki", "all-x509-certs")
		require.NoError(t, err)
		checkResponseContains(t, out, TestRootCertSubjectFor1_2)
		checkResponseContains(t, out, TestRootCertSubjectKeyIDFor1_2)
		checkResponseContains(t, out, testRootCertSubject)
		checkResponseContains(t, out, testRootCertSubjectKeyID)

		// ----- Auth: full single-record + listing coverage -----
		out, err = ExecuteCLIWithBin(dcld, "query", "auth", "all-accounts")
		require.NoError(t, err)
		checkResponseContains(t, out, state.User5Address)
		checkResponseContains(t, out, state.User2Address)

		for _, addr := range []string{state.User5Address, state.User2Address} {
			out, err = ExecuteCLIWithBin(dcld,
				"query", "auth", "account", "--address", addr,
			)
			require.NoError(t, err)
			checkResponseContains(t, out, addr)
		}

		out, err = ExecuteCLIWithBin(dcld, "query", "auth", "all-proposed-accounts")
		require.NoError(t, err)
		checkResponseContains(t, out, state.User6Address)
		checkResponseContains(t, out, state.User3Address)

		for _, addr := range []string{state.User6Address, state.User3Address} {
			out, err = ExecuteCLIWithBin(dcld,
				"query", "auth", "proposed-account", "--address", addr,
			)
			require.NoError(t, err)
			checkResponseContains(t, out, addr)
		}

		out, err = ExecuteCLIWithBin(dcld, "query", "auth", "all-proposed-accounts-to-revoke")
		require.NoError(t, err)
		checkResponseContains(t, out, state.User5Address)
		checkResponseContains(t, out, state.User2Address)

		for _, addr := range []string{state.User5Address, state.User2Address} {
			out, err = ExecuteCLIWithBin(dcld,
				"query", "auth", "proposed-account-to-revoke", "--address", addr,
			)
			require.NoError(t, err)
			checkResponseContains(t, out, addr)
		}

		out, err = ExecuteCLIWithBin(dcld, "query", "auth", "all-revoked-accounts")
		require.NoError(t, err)
		checkResponseContains(t, out, state.User4Address)
		checkResponseContains(t, out, state.User1Address)

		for _, addr := range []string{state.User4Address, state.User1Address} {
			out, err = ExecuteCLIWithBin(dcld,
				"query", "auth", "revoked-account", "--address", addr,
			)
			require.NoError(t, err)
			checkResponseContains(t, out, addr)
		}

		// ----- Validator (host-side) -----
		if state.ValidatorAddress != "" {
			out, err = ExecuteCLIWithBin(dcld, "query", "validator", "all-nodes")
			require.NoError(t, err)
			checkResponseContains(t, out, state.ValidatorAddress)
		}
	})

	// ------------------------------------------------------------------
	// Seed _for_1_2_r2 state (still on v1.2).
	// ------------------------------------------------------------------
	MustRun(t, "CreateR2Accounts", func(t *testing.T) {
		approvers := []string{state.Trustee2, state.Trustee3, state.Trustee4}
		_ = CreateAndApproveAccount(t, dcld, VendorAccountFor1_2R2, "Vendor",
			VIDFor1_2R2, state.Trustee1, approvers)
	})

	MustRun(t, "AddR2UserKeys", func(t *testing.T) {
		u4, err := newUserKey(dcld)
		require.NoError(t, err)
		u5, err := newUserKey(dcld)
		require.NoError(t, err)
		u6, err := newUserKey(dcld)
		require.NoError(t, err)
		state.User4Address, state.User4Pubkey = u4.address, u4.pubkey
		state.User5Address, state.User5Pubkey = u5.address, u5.pubkey
		state.User6Address, state.User6Pubkey = u6.address, u6.pubkey
	})

	MustRun(t, "VendorInfoForR2", func(t *testing.T) {
		tx, err := ExecuteTxWithBin(dcld,
			"tx", "vendorinfo", "add-vendor",
			"--vid", fmt.Sprintf("%d", VIDFor1_2R2),
			"--vendorName", VendorNameFor1_2R2,
			"--companyLegalName", CompanyLegalNameFor1_2R2,
			"--companyPreferredName", CompanyPreferredNameFor1_2R2,
			"--vendorLandingPageURL", VendorLandingPageURLFor1_2R2,
			"--from", VendorAccountFor1_2R2,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)
	})

	MustRun(t, "ModelsForR2", func(t *testing.T) {
		for _, pid := range []int{PID1For1_2R2, PID2For1_2R2, PID3For1_2R2} {
			tx, err := ExecuteTxWithBin(dcld,
				"tx", "model", "add-model",
				"--vid", fmt.Sprintf("%d", VIDFor1_2R2),
				"--pid", fmt.Sprintf("%d", pid),
				"--deviceTypeID", fmt.Sprintf("%d", DeviceTypeIDFor1_2R2),
				"--productName", ProductNameFor1_2R2,
				"--productLabel", ProductLabelFor1_2R2,
				"--partNumber", PartNumberFor1_2R2,
				"--from", VendorAccountFor1_2R2,
			)
			require.NoError(t, err)
			require.Equal(t, uint32(0), tx.Code, tx.RawLog)

			tx, err = ExecuteTxWithBin(dcld,
				"tx", "model", "add-model-version",
				"--vid", fmt.Sprintf("%d", VIDFor1_2R2),
				"--pid", fmt.Sprintf("%d", pid),
				"--softwareVersion", fmt.Sprintf("%d", SoftwareVersionFor1_2R2),
				"--softwareVersionString", SoftwareVersionStringFor1_2R2,
				"--cdVersionNumber", fmt.Sprintf("%d", CDVersionNumberFor1_2R2),
				"--minApplicableSoftwareVersion", fmt.Sprintf("%d", MinApplicableSoftwareVersionFor1_2R2),
				"--maxApplicableSoftwareVersion", fmt.Sprintf("%d", MaxApplicableSoftwareVersionFor1_2R2),
				"--from", VendorAccountFor1_2R2,
			)
			require.NoError(t, err)
			require.Equal(t, uint32(0), tx.Code, tx.RawLog)
		}

		// Delete pid_3.
		tx, err := ExecuteTxWithBin(dcld,
			"tx", "model", "delete-model",
			"--vid", fmt.Sprintf("%d", VIDFor1_2R2),
			"--pid", fmt.Sprintf("%d", PID3For1_2R2),
			"--from", VendorAccountFor1_2R2,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)
	})

	MustRun(t, "ComplianceForR2", func(t *testing.T) {
		// certify pid_1.
		tx, err := ExecuteTxWithBin(dcld,
			"tx", "compliance", "certify-model",
			"--vid", fmt.Sprintf("%d", VIDFor1_2R2),
			"--pid", fmt.Sprintf("%d", PID1For1_2R2),
			"--softwareVersion", fmt.Sprintf("%d", SoftwareVersionFor1_2R2),
			"--softwareVersionString", SoftwareVersionStringFor1_2R2,
			"--certificationType", CertificationTypeFor1_2R2,
			"--certificationDate", CertificationDateFor1_2R2,
			"--cdCertificateId", CDCertificateIDFor1_2R2,
			"--cdVersionNumber", fmt.Sprintf("%d", CDVersionNumberFor1_2R2),
			"--from", CertificationCenterAccountFor1_2,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), tx.Code, tx.RawLog)

		// provision pid_2, certify pid_2, revoke pid_2.
		// revoke-model does not accept --cdCertificateId, so it is appended
		// only for the provision/certify actions.
		for _, action := range []struct {
			cmd, dateFlag, dateVal string
		}{
			{"provision-model", "--provisionalDate", ProvisionalDateFor1_2R2},
			{"certify-model", "--certificationDate", CertificationDateFor1_2R2},
			{"revoke-model", "--revocationDate", CertificationDateFor1_2R2},
		} {
			args := []string{
				"tx", "compliance", action.cmd,
				"--vid", fmt.Sprintf("%d", VIDFor1_2R2),
				"--pid", fmt.Sprintf("%d", PID2For1_2R2),
				"--softwareVersion", fmt.Sprintf("%d", SoftwareVersionFor1_2R2),
				"--softwareVersionString", SoftwareVersionStringFor1_2R2,
				"--certificationType", CertificationTypeFor1_2R2,
				action.dateFlag, action.dateVal,
				"--cdVersionNumber", fmt.Sprintf("%d", CDVersionNumberFor1_2R2),
				"--from", CertificationCenterAccountFor1_2,
			}
			if action.cmd != "revoke-model" {
				args = append(args, "--cdCertificateId", CDCertificateIDFor1_2R2)
			}
			tx, err = ExecuteTxWithBin(dcld, args...)
			require.NoError(t, err)
			require.Equal(t, uint32(0), tx.Code, "%s pid_2: %s", action.cmd, tx.RawLog)
		}
	})

	MustRun(t, "AccountFlowsForR2", func(t *testing.T) {
		approvers := []string{state.Trustee2, state.Trustee3, state.Trustee4}

		proposeUserAccount(t, dcld, state.Trustee1, approvers,
			state.User4Address, state.User4Pubkey, "CertificationCenter", true)
		proposeUserAccount(t, dcld, state.Trustee1, approvers,
			state.User5Address, state.User5Pubkey, "CertificationCenter", true)
		proposeUserAccount(t, dcld, state.Trustee1, nil,
			state.User6Address, state.User6Pubkey, "CertificationCenter", false)

		revokeUserAccount(t, dcld, state.Trustee1, approvers, state.User4Address, true)
		revokeUserAccount(t, dcld, state.Trustee1, nil, state.User5Address, false)
	})

	MustRun(t, "ValidatorDisableEnableFlow", func(t *testing.T) {
		RunValidatorDisableEnableFlow(t, state, dcld,
			[]string{state.Trustee2, state.Trustee3, state.Trustee4})
	})
}
