package pki

import (
	"testing"

	"github.com/stretchr/testify/require"
	dclauth "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/cli/dclauth"
	cliputils "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/cli/utils"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/utils"
)

const (
	// Use paa_cert_no_vid_mainnet (Matter PAA 2, VID=24582). google_root_cert is used by
	// TestPKIDemo's Google sections, so we avoid it here to keep tests state-independent.
	approvalTestRootCertPath         = "../../constants/paa_cert_no_vid_mainnet"
	approvalTestRootCertSubject      = "MEsxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZHb29nbGUxFTATBgNVBAMMDE1hdHRlciBQQUEgMjEUMBIGCisGAQQBgqJ8AgEMBDYwMDY="
	approvalTestRootCertSubjectKeyID = "7A:B9:ED:A7:6F:E9:CB:64:62:75:32:6D:D1:45:08:B8:00:F8:E1:C8"
	approvalTestVid                  = 24582
)

// This test covers approval and revocation flows requiring a 2/3-quorum with 6 trustees.
func TestPKIApproval(t *testing.T) {
	jack := testconstants.JackAccount
	alice := testconstants.AliceAccount
	bob := testconstants.BobAccount

	jackAddr, err := cliputils.GetAddress(jack)
	require.NoError(t, err)
	aliceAddr, err := cliputils.GetAddress(alice)
	require.NoError(t, err)
	bobAddr, err := cliputils.GetAddress(bob)
	require.NoError(t, err)

	userAccount := cliputils.CreateAccount(t, "CertificationCenter")

	// At genesis 3 trustees exist (Jack, Alice, Bob).
	// threshold = ceil(2/3 * N); for N=3 → 2, N=4 → 3, N=5 → 4, N=6 → 4.

	// Create 4th trustee (N=3, threshold=2): Jack proposes + Alice approves.
	fourthKey := utils.RandString()
	require.NoError(t, cliputils.AddKey(fourthKey))
	fourthAddr, err := cliputils.GetAddress(fourthKey)
	require.NoError(t, err)
	fourthPubkey, err := cliputils.GetPubkey(fourthKey)
	require.NoError(t, err)
	txResult, err := dclauth.ProposeAccount(fourthAddr, fourthPubkey, "Trustee", jack)
	require.NoError(t, err)
	require.Equal(t, uint32(0), txResult.Code, "propose 4th trustee: %s", txResult.RawLog)
	txResult, err = dclauth.ApproveAccount(fourthAddr, alice)
	require.NoError(t, err)
	require.Equal(t, uint32(0), txResult.Code, "approve 4th trustee (alice): %s", txResult.RawLog)
	fourthAcc, err := dclauth.GetAccount(fourthAddr)
	require.NoError(t, err)
	require.NotNil(t, fourthAcc, "4th trustee should be active")

	// Create 5th trustee (N=4, threshold=3): Jack + Alice + Bob.
	fifthKey := utils.RandString()
	require.NoError(t, cliputils.AddKey(fifthKey))
	fifthAddr, err := cliputils.GetAddress(fifthKey)
	require.NoError(t, err)
	fifthPubkey, err := cliputils.GetPubkey(fifthKey)
	require.NoError(t, err)
	txResult, err = dclauth.ProposeAccount(fifthAddr, fifthPubkey, "Trustee", jack)
	require.NoError(t, err)
	require.Equal(t, uint32(0), txResult.Code, "propose 5th trustee: %s", txResult.RawLog)
	for _, approver := range []string{alice, bob} {
		txResult, err = dclauth.ApproveAccount(fifthAddr, approver)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code, "approve 5th trustee (%s): %s", approver, txResult.RawLog)
	}
	fifthAcc, err := dclauth.GetAccount(fifthAddr)
	require.NoError(t, err)
	require.NotNil(t, fifthAcc, "5th trustee should be active")

	// Create 6th trustee (N=5, threshold=4): Jack + Alice + Bob + fourth.
	sixthKey := utils.RandString()
	require.NoError(t, cliputils.AddKey(sixthKey))
	sixthAddr, err := cliputils.GetAddress(sixthKey)
	require.NoError(t, err)
	sixthPubkey, err := cliputils.GetPubkey(sixthKey)
	require.NoError(t, err)
	txResult, err = dclauth.ProposeAccount(sixthAddr, sixthPubkey, "Trustee", jack)
	require.NoError(t, err)
	require.Equal(t, uint32(0), txResult.Code, "propose 6th trustee: %s", txResult.RawLog)
	for _, approver := range []string{alice, bob, fourthKey} {
		txResult, err = dclauth.ApproveAccount(sixthAddr, approver)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code, "approve 6th trustee (%s): %s", approver, txResult.RawLog)
	}
	sixthAcc, err := dclauth.GetAccount(sixthAddr)
	require.NoError(t, err)
	require.NotNil(t, sixthAcc, "6th trustee should be active")

	// Cleanup: revoke extra trustees after the test so subsequent tests aren't affected by
	// the raised quorum. Register in order 4th→5th→6th so LIFO runs 6th→5th→4th.
	t.Cleanup(func() {
		// 4th cleanup runs last (4 trustees remain, threshold=3): Jack proposes, Alice+Bob approve.
		res, _ := dclauth.ProposeRevokeAccount(fourthAddr, jack)
		if res != nil && res.Code == 0 {
			dclauth.ApproveRevokeAccount(fourthAddr, alice) //nolint:errcheck
			dclauth.ApproveRevokeAccount(fourthAddr, bob)   //nolint:errcheck
		}
	})
	t.Cleanup(func() {
		// 5th cleanup runs 2nd (5 trustees remain, threshold=4): Jack proposes, Alice+Bob+fourth approve.
		res, _ := dclauth.ProposeRevokeAccount(fifthAddr, jack)
		if res != nil && res.Code == 0 {
			dclauth.ApproveRevokeAccount(fifthAddr, alice)     //nolint:errcheck
			dclauth.ApproveRevokeAccount(fifthAddr, bob)       //nolint:errcheck
			dclauth.ApproveRevokeAccount(fifthAddr, fourthKey) //nolint:errcheck
		}
	})
	t.Cleanup(func() {
		// 6th cleanup runs first (6 trustees, threshold=4): Jack proposes, Alice+Bob+fourth approve.
		res, _ := dclauth.ProposeRevokeAccount(sixthAddr, jack)
		if res != nil && res.Code == 0 {
			dclauth.ApproveRevokeAccount(sixthAddr, alice)     //nolint:errcheck
			dclauth.ApproveRevokeAccount(sixthAddr, bob)       //nolint:errcheck
			dclauth.ApproveRevokeAccount(sixthAddr, fourthKey) //nolint:errcheck
		}
	})

	t.Run("NonTrustee_CannotProposeRootCert", func(t *testing.T) {
		txResult, err := ProposeAddX509RootCert(approvalTestRootCertPath, userAccount, X509ProposeOpts{VID: approvalTestVid})
		require.NoError(t, err)
		require.NotEqual(t, uint32(0), txResult.Code)
	})

	t.Run("ProposeAndApproveWithQuorum", func(t *testing.T) {
		// 6 trustees, threshold=4.
		// 4th proposes (=1), Jack approves (=2), Alice approves (=3), Bob approves (=4) → approved.
		txResult, err := ProposeAddX509RootCert(approvalTestRootCertPath, fourthKey, X509ProposeOpts{VID: approvalTestVid})
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)

		// After the proposal, the proposer's approval is recorded and the cert is
		// not yet in the approved store.
		proposed, err := GetProposedX509RootCert(approvalTestRootCertSubject, approvalTestRootCertSubjectKeyID)
		require.NoError(t, err)
		require.NotNil(t, proposed)
		require.True(t, grantsContain(proposed.Approvals, fourthAddr))
		cert, err := GetX509Cert(approvalTestRootCertSubject, approvalTestRootCertSubjectKeyID)
		require.NoError(t, err)
		require.Nil(t, cert, "cert must not be approved after only the proposal")

		txResult, err = ApproveAddX509RootCert(approvalTestRootCertSubject, approvalTestRootCertSubjectKeyID, jack)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)

		// Still proposed (not approved) after 2 approvals; approvals accumulate.
		proposed, err = GetProposedX509RootCert(approvalTestRootCertSubject, approvalTestRootCertSubjectKeyID)
		require.NoError(t, err)
		require.NotNil(t, proposed)
		require.Equal(t, approvalTestRootCertSubject, proposed.Subject)
		require.True(t, grantsContain(proposed.Approvals, fourthAddr))
		require.True(t, grantsContain(proposed.Approvals, jackAddr))
		cert, err = GetX509Cert(approvalTestRootCertSubject, approvalTestRootCertSubjectKeyID)
		require.NoError(t, err)
		require.Nil(t, cert, "cert must not be approved after 2/4 approvals")

		txResult, err = ApproveAddX509RootCert(approvalTestRootCertSubject, approvalTestRootCertSubjectKeyID, alice)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)

		// Still proposed (not approved) after 3 approvals.
		proposed, err = GetProposedX509RootCert(approvalTestRootCertSubject, approvalTestRootCertSubjectKeyID)
		require.NoError(t, err)
		require.NotNil(t, proposed)
		require.True(t, grantsContain(proposed.Approvals, aliceAddr))
		cert, err = GetX509Cert(approvalTestRootCertSubject, approvalTestRootCertSubjectKeyID)
		require.NoError(t, err)
		require.Nil(t, cert, "cert must not be approved after 3/4 approvals")

		txResult, err = ApproveAddX509RootCert(approvalTestRootCertSubject, approvalTestRootCertSubjectKeyID, bob)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)

		// Now approved (4 approvals = quorum); all four approver addresses are
		// recorded on the approved certificate.
		cert, err = GetX509Cert(approvalTestRootCertSubject, approvalTestRootCertSubjectKeyID)
		require.NoError(t, err)
		require.NotNil(t, cert)
		require.Equal(t, approvalTestRootCertSubject, cert.Subject)
		require.NotEmpty(t, cert.Certs)
		require.True(t, cert.Certs[0].IsRoot)
		require.True(t, grantsContain(cert.Certs[0].Approvals, fourthAddr))
		require.True(t, grantsContain(cert.Certs[0].Approvals, jackAddr))
		require.True(t, grantsContain(cert.Certs[0].Approvals, aliceAddr))
		require.True(t, grantsContain(cert.Certs[0].Approvals, bobAddr))

		proposed, err = GetProposedX509RootCert(approvalTestRootCertSubject, approvalTestRootCertSubjectKeyID)
		require.NoError(t, err)
		require.Nil(t, proposed)
	})

	t.Run("RevokeRootCertWithQuorum", func(t *testing.T) {
		// 6 trustees, threshold=4.
		// 6th proposes (=1), 5th approves (=2), 4th approves (=3), Bob approves (=4) → revoked.
		txResult, err := ProposeRevokeX509RootCert(approvalTestRootCertSubject, approvalTestRootCertSubjectKeyID, sixthKey)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)

		sixthAddr, err := cliputils.GetAddress(sixthKey)
		require.NoError(t, err)
		fifthAddr, err := cliputils.GetAddress(fifthKey)
		require.NoError(t, err)
		fourthRevAddr, err := cliputils.GetAddress(fourthKey)
		require.NoError(t, err)

		proposedRev, err := GetProposedRevokedX509RootCert(approvalTestRootCertSubject, approvalTestRootCertSubjectKeyID)
		require.NoError(t, err)
		require.NotNil(t, proposedRev)
		require.Equal(t, approvalTestRootCertSubject, proposedRev.Subject)
		require.True(t, grantsContain(proposedRev.Approvals, sixthAddr))
		// Cert is still approved (not yet revoked) with only the proposer's vote.
		cert, err := GetX509Cert(approvalTestRootCertSubject, approvalTestRootCertSubjectKeyID)
		require.NoError(t, err)
		require.NotNil(t, cert)

		txResult, err = ApproveRevokeX509RootCert(approvalTestRootCertSubject, approvalTestRootCertSubjectKeyID, fifthKey)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)

		// Still proposed (2 approvals accumulate); cert still approved.
		proposedRev, err = GetProposedRevokedX509RootCert(approvalTestRootCertSubject, approvalTestRootCertSubjectKeyID)
		require.NoError(t, err)
		require.NotNil(t, proposedRev)
		require.True(t, grantsContain(proposedRev.Approvals, sixthAddr))
		require.True(t, grantsContain(proposedRev.Approvals, fifthAddr))
		cert, err = GetX509Cert(approvalTestRootCertSubject, approvalTestRootCertSubjectKeyID)
		require.NoError(t, err)
		require.NotNil(t, cert)

		txResult, err = ApproveRevokeX509RootCert(approvalTestRootCertSubject, approvalTestRootCertSubjectKeyID, fourthKey)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)

		// Still proposed (3 approvals).
		proposedRev, err = GetProposedRevokedX509RootCert(approvalTestRootCertSubject, approvalTestRootCertSubjectKeyID)
		require.NoError(t, err)
		require.NotNil(t, proposedRev)
		require.True(t, grantsContain(proposedRev.Approvals, fourthRevAddr))

		txResult, err = ApproveRevokeX509RootCert(approvalTestRootCertSubject, approvalTestRootCertSubjectKeyID, bob)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)

		// Now revoked (4 approvals = quorum); the revoke voters are recorded.
		revoked, err := GetRevokedX509Cert(approvalTestRootCertSubject, approvalTestRootCertSubjectKeyID)
		require.NoError(t, err)
		require.NotNil(t, revoked)
		require.Equal(t, approvalTestRootCertSubject, revoked.Subject)
		require.NotEmpty(t, revoked.Certs)
		require.True(t, grantsContain(revoked.Certs[0].Approvals, sixthAddr))
		require.True(t, grantsContain(revoked.Certs[0].Approvals, fifthAddr))
		require.True(t, grantsContain(revoked.Certs[0].Approvals, fourthRevAddr))
		require.True(t, grantsContain(revoked.Certs[0].Approvals, bobAddr))

		cert, err = GetX509Cert(approvalTestRootCertSubject, approvalTestRootCertSubjectKeyID)
		require.NoError(t, err)
		require.Nil(t, cert)
	})
}
