package dclauth

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/utils"
	dclauthtypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
)

func TestAuthDemoNodeAdmin(t *testing.T) {
	jack := testconstants.JackAccount
	alice := testconstants.AliceAccount
	bob := testconstants.BobAccount

	// Generate new user key
	name := fmt.Sprintf("user%d", rand.Intn(99999))
	err := AddKey(name)
	require.NoError(t, err)

	userAddr, err := GetAddress(name)
	require.NoError(t, err)
	userPubkey, err := GetPubkey(name)
	require.NoError(t, err)

	jackAddr, err := GetAddress(jack)
	require.NoError(t, err)
	aliceAddr, err := GetAddress(alice)
	require.NoError(t, err)
	bobAddr, err := GetAddress(bob)
	require.NoError(t, err)

	t.Run("InitialState_NotFound", func(t *testing.T) {
		acc, err := GetAccount(userAddr)
		require.NoError(t, err)
		require.Nil(t, acc)

		propRev, err := GetProposedAccountToRevoke(userAddr)
		require.NoError(t, err)
		require.Nil(t, propRev)

		prop, err := GetProposedAccount(userAddr)
		require.NoError(t, err)
		require.Nil(t, prop)

		revoked, err := GetRevokedAccount(userAddr)
		require.NoError(t, err)
		require.Nil(t, revoked)
	})

	t.Run("InitialListsEmpty", func(t *testing.T) {
		allProposed, err := GetAllProposedAccounts()
		require.NoError(t, err)
		require.False(t, containsPendingAccountAddress(allProposed, userAddr))

		allProposedRev, err := GetAllProposedAccountsToRevoke()
		require.NoError(t, err)
		require.False(t, containsPendingAccountRevocationAddress(allProposedRev, userAddr))

		allRevoked, err := GetAllRevokedAccounts()
		require.NoError(t, err)
		require.False(t, containsRevokedAccountAddress(allRevoked, userAddr))
	})

	t.Run("JackProposes", func(t *testing.T) {
		txResult, err := ProposeAccount(userAddr, userPubkey, "NodeAdmin", jack, ProposeAccountOpts{Info: "Jack is proposing this account"})
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Account not yet active.
		all, err := GetAllAccounts()
		require.NoError(t, err)
		require.False(t, containsAccountAddress(all, userAddr))

		acc, err := GetAccount(userAddr)
		require.NoError(t, err)
		require.Nil(t, acc)

		// Now in proposed list.
		prop, err := GetProposedAccount(userAddr)
		require.NoError(t, err)
		require.NotNil(t, prop)
		require.NotNil(t, prop.Account)
		require.Equal(t, userAddr, prop.Account.Address)
		require.Len(t, prop.Account.Approvals, 1)
		require.Equal(t, jackAddr, prop.Account.Approvals[0].Address)
		require.Equal(t, "Jack is proposing this account", prop.Account.Approvals[0].Info)

		allProposed, err := GetAllProposedAccounts()
		require.NoError(t, err)
		require.True(t, containsPendingAccountAddress(allProposed, userAddr))
	})

	t.Run("AliceApproves", func(t *testing.T) {
		txResult, err := ApproveAccount(userAddr, alice, AccountActionOpts{Info: "Alice is approving this account"})
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Alice cannot reject after approving
		txBad, err := RejectAccount(userAddr, alice, AccountActionOpts{Info: "Alice is rejecting this account"})
		// Either error or non-zero code
		if err == nil {
			require.NotEqual(t, uint32(0), txBad.Code)
		}

		// Account is now active.
		all, err := GetAllAccounts()
		require.NoError(t, err)
		require.True(t, containsAccountAddress(all, userAddr))

		acc, err := GetAccount(userAddr)
		require.NoError(t, err)
		require.NotNil(t, acc)
		require.Len(t, acc.Approvals, 2)
		approvers := []string{acc.Approvals[0].Address, acc.Approvals[1].Address}
		require.Contains(t, approvers, jackAddr)
		require.Contains(t, approvers, aliceAddr)
		infos := []string{acc.Approvals[0].Info, acc.Approvals[1].Info}
		require.Contains(t, infos, "Jack is proposing this account")
		require.Contains(t, infos, "Alice is approving this account")

		// No longer in proposed list.
		allProposed, err := GetAllProposedAccounts()
		require.NoError(t, err)
		require.False(t, containsPendingAccountAddress(allProposed, userAddr))

		prop, err := GetProposedAccount(userAddr)
		require.NoError(t, err)
		require.Nil(t, prop)
	})

	t.Run("AliceProposeRevoke", func(t *testing.T) {
		txResult, err := ProposeRevokeAccount(userAddr, alice, AccountActionOpts{Info: "Alice proposes to revoke account"})
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Still in active accounts (not enough approvals to revoke).
		all, err := GetAllAccounts()
		require.NoError(t, err)
		require.True(t, containsAccountAddress(all, userAddr))

		// In proposed-to-revoke list.
		allProposedRev, err := GetAllProposedAccountsToRevoke()
		require.NoError(t, err)
		require.True(t, containsPendingAccountRevocationAddress(allProposedRev, userAddr))

		propRev, err := GetProposedAccountToRevoke(userAddr)
		require.NoError(t, err)
		require.NotNil(t, propRev)
		require.Equal(t, userAddr, propRev.Address)
		require.Len(t, propRev.Approvals, 1)
		require.Equal(t, aliceAddr, propRev.Approvals[0].Address)
		require.Equal(t, "Alice proposes to revoke account", propRev.Approvals[0].Info)

		// Not yet revoked.
		allRevoked, err := GetAllRevokedAccounts()
		require.NoError(t, err)
		require.False(t, containsRevokedAccountAddress(allRevoked, userAddr))
	})

	t.Run("BobApprovesRevoke", func(t *testing.T) {
		txResult, err := ApproveRevokeAccount(userAddr, bob)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Now revoked.
		allRevoked, err := GetAllRevokedAccounts()
		require.NoError(t, err)
		require.True(t, containsRevokedAccountAddress(allRevoked, userAddr))

		// No longer in active accounts.
		all, err := GetAllAccounts()
		require.NoError(t, err)
		require.False(t, containsAccountAddress(all, userAddr))

		allProposedRev, err := GetAllProposedAccountsToRevoke()
		require.NoError(t, err)
		require.False(t, containsPendingAccountRevocationAddress(allProposedRev, userAddr))

		propRev, err := GetProposedAccountToRevoke(userAddr)
		require.NoError(t, err)
		require.Nil(t, propRev)

		revoked, err := GetRevokedAccount(userAddr)
		require.NoError(t, err)
		require.NotNil(t, revoked)
		require.NotNil(t, revoked.Account)
		require.Equal(t, userAddr, revoked.Account.Address)
		require.Equal(t, dclauthtypes.RevokedAccount_TrusteeVoting, revoked.Reason)
	})

	t.Run("ReAddAfterRevoke_ProposeApprove", func(t *testing.T) {
		// Jack proposes again.
		txResult, err := ProposeAccount(userAddr, userPubkey, "NodeAdmin", jack, ProposeAccountOpts{Info: "Jack is proposing this account"})
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Still revoked in revoked list.
		allRevoked, err := GetAllRevokedAccounts()
		require.NoError(t, err)
		require.True(t, containsRevokedAccountAddress(allRevoked, userAddr))

		// Not yet active.
		all, err := GetAllAccounts()
		require.NoError(t, err)
		require.False(t, containsAccountAddress(all, userAddr))

		prop, err := GetProposedAccount(userAddr)
		require.NoError(t, err)
		require.NotNil(t, prop)
		require.Equal(t, userAddr, prop.Account.Address)

		// Alice approves.
		txResult, err = ApproveAccount(userAddr, alice, AccountActionOpts{Info: "Alice is approving this account"})
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// No longer in revoked list.
		allRevoked, err = GetAllRevokedAccounts()
		require.NoError(t, err)
		require.False(t, containsRevokedAccountAddress(allRevoked, userAddr))

		// Active again.
		all, err = GetAllAccounts()
		require.NoError(t, err)
		require.True(t, containsAccountAddress(all, userAddr))
	})

	t.Run("RejectScenario", func(t *testing.T) {
		// Propose-revoke again then approve-revoke
		txResult, err := ProposeRevokeAccount(userAddr, alice, AccountActionOpts{Info: "Alice proposes to revoke account"})
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		txResult, err = ApproveRevokeAccount(userAddr, bob)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Jack proposes again for rejection test
		txResult, err = ProposeAccount(userAddr, userPubkey, "NodeAdmin", jack, ProposeAccountOpts{Info: "Jack is proposing this account"})
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Alice rejects
		txResult, err = RejectAccount(userAddr, alice, AccountActionOpts{Info: "Alice is rejecting this account"})
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Still in proposed (alice rejected but not enough rejections).
		allProposed, err := GetAllProposedAccounts()
		require.NoError(t, err)
		require.True(t, containsPendingAccountAddress(allProposed, userAddr))

		all, err := GetAllAccounts()
		require.NoError(t, err)
		require.False(t, containsAccountAddress(all, userAddr))

		allRejected, err := GetAllRejectedAccounts()
		require.NoError(t, err)
		require.False(t, containsRejectedAccountAddress(allRejected, userAddr))

		// Bob rejects.
		txResult, err = RejectAccount(userAddr, bob, AccountActionOpts{Info: "Bob is rejecting this account"})
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Bob cannot reject again.
		txBad, errBad := RejectAccount(userAddr, bob, AccountActionOpts{Info: "Bob is rejecting this account"})
		if errBad == nil {
			require.NotEqual(t, uint32(0), txBad.Code)
		}

		// Now in rejected list (enough rejections).
		allRejected, err = GetAllRejectedAccounts()
		require.NoError(t, err)
		require.True(t, containsRejectedAccountAddress(allRejected, userAddr))

		allProposed, err = GetAllProposedAccounts()
		require.NoError(t, err)
		require.False(t, containsPendingAccountAddress(allProposed, userAddr))

		all, err = GetAllAccounts()
		require.NoError(t, err)
		require.False(t, containsAccountAddress(all, userAddr))

		rejected, err := GetRejectedAccount(userAddr)
		require.NoError(t, err)
		require.NotNil(t, rejected)
		require.NotNil(t, rejected.Account)
		require.Equal(t, userAddr, rejected.Account.Address)
		approvers := []string{}
		infos := []string{}
		for _, a := range rejected.Account.Approvals {
			approvers = append(approvers, a.Address)
			infos = append(infos, a.Info)
		}
		require.Contains(t, approvers, jackAddr)
		require.Contains(t, infos, "Jack is proposing this account")
		var rejectors []string
		var rejectInfos []string
		for _, r := range rejected.Account.Rejects {
			rejectors = append(rejectors, r.Address)
			rejectInfos = append(rejectInfos, r.Info)
		}
		require.Contains(t, rejectors, aliceAddr)
		require.Contains(t, rejectInfos, "Alice is rejecting this account")
		require.Contains(t, rejectors, bobAddr)
		require.Contains(t, rejectInfos, "Bob is rejecting this account")
	})

	t.Run("RejectedAccountCannotAddModel", func(t *testing.T) {
		// The rejected user's key is in the keyring, but the account was never
		// created on-chain, so signing a model tx fails ("key not found") — a
		// rejected account cannot transact (auth-demo.sh:627-631).
		mvid := rand.Intn(65534) + 1
		mpid := rand.Intn(65534) + 1
		txResult, err := utils.ExecuteTx("tx", "model", "add-model",
			"--vid", fmt.Sprintf("%d", mvid),
			"--pid", fmt.Sprintf("%d", mpid),
			"--productName", "Device #2",
			"--productLabel", "Device Description",
			"--commissioningCustomFlow", "0",
			"--deviceTypeID", "12",
			"--partNumber", "12",
			"--enhancedSetupFlowOptions", "0",
			"--from", name,
		)
		combined := ""
		if err != nil {
			combined += err.Error()
		}
		if txResult != nil {
			combined += txResult.RawLog
		}
		require.Contains(t, combined, "key not found")
	})

	// Unused variables referenced to avoid compiler errors
	_ = strings.TrimSpace
}

// TestAuthDemoJackRejectOwnProposal tests that a single trustee can propose and then
// self-reject their own proposal. With only 1 rejection (Jack), the rejection quorum
// is not reached with 3 trustees (need 2), so the account ends up nowhere.
func TestAuthDemoJackRejectOwnProposal(t *testing.T) {
	jack := testconstants.JackAccount

	name := fmt.Sprintf("user%d", rand.Intn(99999))
	err := AddKey(name)
	require.NoError(t, err)

	userAddr, err := GetAddress(name)
	require.NoError(t, err)
	userPubkey, err := GetPubkey(name)
	require.NoError(t, err)

	t.Run("ProposeAndSelfReject", func(t *testing.T) {
		txResult, err := ProposeAccount(userAddr, userPubkey, "NodeAdmin", jack, ProposeAccountOpts{Info: "Jack is proposing this account"})
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Jack rejects his own proposal
		txResult, err = RejectAccount(userAddr, jack, AccountActionOpts{Info: "Jack is rejecting this account"})
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Not in proposed (jack removed his approval+proposal).
		prop, err := GetProposedAccount(userAddr)
		require.NoError(t, err)
		require.Nil(t, prop)

		// Not in rejected (single rejection doesn't reach quorum with 3 trustees).
		rejected, err := GetRejectedAccount(userAddr)
		require.NoError(t, err)
		require.Nil(t, rejected)

		// Not in approved.
		acc, err := GetAccount(userAddr)
		require.NoError(t, err)
		require.Nil(t, acc)
	})
}

// revokeTrusteeIfPresent restores the genesis trustee set by revoking the
// trustee at addr when it is still an active account. It is tolerant of partial
// state (e.g. a test that failed mid-way): it proposes a revocation when none is
// pending and then collects approvals from the candidate voters until the
// account is gone. Votes the chain rejects (voter already voted, voter itself
// revoked, quorum already met) are skipped, and the loop stops as soon as the
// account is no longer active. A leftover *pending* trustee proposal is ignored
// on purpose — only active accounts count toward the trustee quorum.
func revokeTrusteeIfPresent(t *testing.T, addr string, voters ...string) {
	t.Helper()

	acc, err := GetAccount(addr)
	if err != nil || acc == nil {
		return // not active — nothing to revoke
	}

	proposed, _ := GetProposedAccountToRevoke(addr)
	for _, voter := range voters {
		if rev, _ := GetRevokedAccount(addr); rev != nil {
			return // revocation reached quorum
		}
		if cur, _ := GetAccount(addr); cur == nil {
			return
		}

		var txResult *utils.TxResult
		if proposed == nil {
			txResult, err = ProposeRevokeAccount(addr, voter)
		} else {
			txResult, err = ApproveRevokeAccount(addr, voter)
		}
		if err != nil || txResult == nil || txResult.Code != 0 {
			continue // voter cannot act right now — skip it
		}
		_, _ = utils.AwaitTxConfirmation(txResult.TxHash)
		if proposed == nil {
			proposed, _ = GetProposedAccountToRevoke(addr)
		}
	}
}

// TestAuthDemoDynamicTrusteeCount exercises the add / revoke / reject approval
// quorums as the trustee set grows (3 → 5) and shrinks back to 3. The genesis
// chain has exactly 3 trustees (jack, alice, bob); this test temporarily adds
// extra Trustee accounts, so it MUST leave the chain with exactly 3 trustees
// again — otherwise the quorum math in every other dclauth test breaks.
//
// It was previously skipped because a mid-way failure left extra trustees on
// chain, corrupting subsequent tests. The t.Cleanup hook below now revokes any
// trustee this test added, regardless of where the test stopped, restoring the
// genesis 3-trustee set even on failure — so the test is safe to run.
func TestAuthDemoDynamicTrusteeCount(t *testing.T) {
	jack := testconstants.JackAccount
	alice := testconstants.AliceAccount
	bob := testconstants.BobAccount

	vid := rand.Intn(65534) + 1

	// ── Provision the trustee keys up front so cleanup can reference them ──

	newTrustee1Name := fmt.Sprintf("trustee5a%d", rand.Intn(99999))
	err := AddKey(newTrustee1Name)
	require.NoError(t, err)
	newTrustee1Addr, err := GetAddress(newTrustee1Name)
	require.NoError(t, err)
	newTrustee1Pubkey, err := GetPubkey(newTrustee1Name)
	require.NoError(t, err)

	newTrustee2Name := fmt.Sprintf("trustee5b%d", rand.Intn(99999))
	err = AddKey(newTrustee2Name)
	require.NoError(t, err)
	newTrustee2Addr, err := GetAddress(newTrustee2Name)
	require.NoError(t, err)
	newTrustee2Pubkey, err := GetPubkey(newTrustee2Name)
	require.NoError(t, err)

	newTrustee3Name := fmt.Sprintf("trustee4c%d", rand.Intn(99999))
	err = AddKey(newTrustee3Name)
	require.NoError(t, err)
	newTrustee3Addr, err := GetAddress(newTrustee3Name)
	require.NoError(t, err)
	newTrustee3Pubkey, err := GetPubkey(newTrustee3Name)
	require.NoError(t, err)

	// Restore the genesis 3-trustee set no matter where the test stops. Any of
	// the base or added trustees may cast the revocation votes; the helper
	// skips voters the chain rejects and stops once each account is gone.
	t.Cleanup(func() {
		voters := []string{jack, alice, bob, newTrustee1Name, newTrustee2Name, newTrustee3Name}
		revokeTrusteeIfPresent(t, newTrustee1Addr, voters...)
		revokeTrusteeIfPresent(t, newTrustee2Addr, voters...)
		revokeTrusteeIfPresent(t, newTrustee3Addr, voters...)
	})

	t.Run("AddTwoNewTrustees", func(t *testing.T) {
		// Jack proposes + Alice approves new_trustee1 → 4 trustees
		txResult, err := ProposeAccount(newTrustee1Addr, newTrustee1Pubkey, "Trustee", jack)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		txResult, err = ApproveAccount(newTrustee1Addr, alice)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Jack proposes new_trustee2, then can reject and re-approve
		txResult, err = ProposeAccount(newTrustee2Addr, newTrustee2Pubkey, "Trustee", jack)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Jack can reject even after proposing
		txResult, err = RejectAccount(newTrustee2Addr, jack, AccountActionOpts{Info: "Jack is rejecting this account"})
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Jack re-approves
		txResult, err = ApproveAccount(newTrustee2Addr, jack, AccountActionOpts{Info: "Jack re-approving"})
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Alice approves → 2nd approval with 4 trustees = quorum (ceil(4*2/3)=3, need 3... hmm)
		// Actually with 4 trustees: ceil(4*2/3)=ceil(2.67)=3 approvals needed.
		// Jack proposed (1) + Alice (2) + Bob (3) → active.
		txResult, err = ApproveAccount(newTrustee2Addr, alice)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		txResult, err = ApproveAccount(newTrustee2Addr, bob)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Both new trustees are now active.
		acc1, err := GetAccount(newTrustee1Addr)
		require.NoError(t, err)
		require.NotNil(t, acc1)
		require.Equal(t, newTrustee1Addr, acc1.Address)

		acc2, err := GetAccount(newTrustee2Addr)
		require.NoError(t, err)
		require.NotNil(t, acc2)
		require.Equal(t, newTrustee2Addr, acc2.Address)
	})

	// ── With 5 trustees: Vendor needs 2 approvals (ceil(5/3)=2) ───────────

	vendorName := fmt.Sprintf("vendor5t%d", rand.Intn(99999))
	err = AddKey(vendorName)
	require.NoError(t, err)
	vendorAddr, err := GetAddress(vendorName)
	require.NoError(t, err)
	vendorPubkey, err := GetPubkey(vendorName)
	require.NoError(t, err)

	t.Run("VendorWith5TrusteesNeeds2Approvals", func(t *testing.T) {
		// Jack proposes (1 approval)
		txResult, err := ProposeAccount(vendorAddr, vendorPubkey, "Vendor", jack, ProposeAccountOpts{Info: "Jack is proposing this account", VID: vid})
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// With 5 trustees, vendor needs ceil(5/3)=2 approvals, so Jack's proposal alone is not enough.
		acc, err := GetAccount(vendorAddr)
		require.NoError(t, err)
		require.Nil(t, acc)

		allProposed, err := GetAllProposedAccounts()
		require.NoError(t, err)
		require.True(t, containsPendingAccountAddress(allProposed, vendorAddr))

		// Alice approves → 2 approvals = quorum → account active.
		txResult, err = ApproveAccount(vendorAddr, alice, AccountActionOpts{Info: "Alice is approving this account"})
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		acc, err = GetAccount(vendorAddr)
		require.NoError(t, err)
		require.NotNil(t, acc)
		require.Equal(t, vendorAddr, acc.Address)

		prop, err := GetProposedAccount(vendorAddr)
		require.NoError(t, err)
		require.Nil(t, prop)
	})

	// ── Revoke vendor: with 5 trustees needs ceil(10/3)=4 approvals ────────

	t.Run("RevokeVendorWith5Trustees", func(t *testing.T) {
		// Alice proposes revocation (1)
		txResult, err := ProposeRevokeAccount(vendorAddr, alice, AccountActionOpts{Info: "Alice proposes to revoke"})
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Bob approves (2) — still not enough, need 4 with 5 trustees
		txResult, err = ApproveRevokeAccount(vendorAddr, bob)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Account still active (need 4 approvals).
		acc, err := GetAccount(vendorAddr)
		require.NoError(t, err)
		require.NotNil(t, acc)
		require.Equal(t, vendorAddr, acc.Address)

		// Revoke new_trustee1 → 4 trustees total
		// With 4 trustees: revocation needs ceil(8/3)=3 approvals → we have alice+bob already
		txResult, err = ProposeRevokeAccount(newTrustee1Addr, alice)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		txResult, err = ApproveRevokeAccount(newTrustee1Addr, bob)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		txResult, err = ApproveRevokeAccount(newTrustee1Addr, jack)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		txResult, err = ApproveRevokeAccount(newTrustee1Addr, newTrustee1Name)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// new_trustee1 is revoked → 4 trustees remain.
		revT1, err := GetRevokedAccount(newTrustee1Addr)
		require.NoError(t, err)
		require.NotNil(t, revT1)
		require.Equal(t, newTrustee1Addr, revT1.Account.Address)

		// Now approve vendor revocation — with 4 trustees need ceil(8/3)=3 approvals.
		// alice(1) + bob(2) + jack(3) = 3 → quorum.
		txResult, err = ApproveRevokeAccount(vendorAddr, jack)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Vendor is now revoked.
		revVendor, err := GetRevokedAccount(vendorAddr)
		require.NoError(t, err)
		require.NotNil(t, revVendor)
		require.Equal(t, vendorAddr, revVendor.Account.Address)
		require.Equal(t, dclauthtypes.RevokedAccount_TrusteeVoting, revVendor.Reason)
	})

	// ── Reject scenario with dynamic trustee count ──────────────────────────
	// With 4 trustees (jack, alice, bob, new_trustee2), rejection needs ceil(8/3)=3

	t.Run("RejectWithDynamicTrusteeCount", func(t *testing.T) {
		// Jack re-proposes the revoked vendor
		txResult, err := ProposeAccount(vendorAddr, vendorPubkey, "Vendor,NodeAdmin", jack, ProposeAccountOpts{Info: "Jack is proposing this account", VID: vid})
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Bob rejects (1 rejection — not enough with 4 trustees, need 3)
		txResult, err = RejectAccount(vendorAddr, bob, AccountActionOpts{Info: "Bob is rejecting this account"})
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Still in proposed.
		allProposed, err := GetAllProposedAccounts()
		require.NoError(t, err)
		require.True(t, containsPendingAccountAddress(allProposed, vendorAddr))

		// Revoke new_trustee2 → 3 trustees (jack, alice, bob).
		txResult, err = ProposeRevokeAccount(newTrustee2Addr, alice)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		txResult, err = ApproveRevokeAccount(newTrustee2Addr, bob)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		txResult, err = ApproveRevokeAccount(newTrustee2Addr, jack)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// new_trustee2 is revoked → 3 trustees.
		revT2, err := GetRevokedAccount(newTrustee2Addr)
		require.NoError(t, err)
		require.NotNil(t, revT2)
		require.Equal(t, newTrustee2Addr, revT2.Account.Address)

		// Alice rejects → with 3 trustees need ceil(6/3)=2 rejections → bob+alice=2 → quorum.
		txResult, err = RejectAccount(vendorAddr, alice, AccountActionOpts{Info: "Alice is rejecting this account"})
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Account is now in rejected list.
		allRejected, err := GetAllRejectedAccounts()
		require.NoError(t, err)
		require.True(t, containsRejectedAccountAddress(allRejected, vendorAddr))

		rejected, err := GetRejectedAccount(vendorAddr)
		require.NoError(t, err)
		require.NotNil(t, rejected)
		require.Equal(t, vendorAddr, rejected.Account.Address)

		allProposed, err = GetAllProposedAccounts()
		require.NoError(t, err)
		require.False(t, containsPendingAccountAddress(allProposed, vendorAddr))
	})

	// ── Critical scenario: approvals collected under a larger trustee set ──
	// must NOT auto-activate an account once the set shrinks; the account stays
	// pending until the next approval re-evaluates the (now-met) quorum. This
	// mirrors the documented edge case in the original auth-demo.sh.
	t.Run("PendingApprovalsSurviveTrusteeRevocation", func(t *testing.T) {
		// Add a 4th trustee: with 3 trustees a Trustee needs ceil(2/3*3)=2
		// approvals → jack proposes + alice approves → active → 4 trustees.
		txResult, err := ProposeAccount(newTrustee3Addr, newTrustee3Pubkey, "Trustee", jack, ProposeAccountOpts{Info: "Jack is proposing this account"})
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		txResult, err = ApproveAccount(newTrustee3Addr, alice, AccountActionOpts{Info: "Alice is approving this account"})
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		acc, err := GetAccount(newTrustee3Addr)
		require.NoError(t, err)
		require.NotNil(t, acc)

		// Re-propose the (rejected) vendor as Vendor,NodeAdmin. The NodeAdmin
		// role uses the 2/3 quorum → with 4 trustees needs ceil(2/3*4)=3.
		txResult, err = ProposeAccount(vendorAddr, vendorPubkey, "Vendor,NodeAdmin", jack, ProposeAccountOpts{Info: "Jack is proposing this account", VID: vid})
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Bob approves → 2 approvals, still 1 short with 4 trustees.
		txResult, err = ApproveAccount(vendorAddr, bob, AccountActionOpts{Info: "Bob is approving this account"})
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		acc, err = GetAccount(vendorAddr)
		require.NoError(t, err)
		require.Nil(t, acc) // still pending

		allProposed, err := GetAllProposedAccounts()
		require.NoError(t, err)
		require.True(t, containsPendingAccountAddress(allProposed, vendorAddr))

		// Revoke the 4th trustee → back to 3 trustees. Revocation needs
		// ceil(2/3*4)=3 approvals: alice proposes + bob + jack.
		txResult, err = ProposeRevokeAccount(newTrustee3Addr, alice, AccountActionOpts{Info: "Alice proposes to revoke account"})
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		txResult, err = ApproveRevokeAccount(newTrustee3Addr, bob)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		txResult, err = ApproveRevokeAccount(newTrustee3Addr, jack)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		revT3, err := GetRevokedAccount(newTrustee3Addr)
		require.NoError(t, err)
		require.NotNil(t, revT3)

		// Even though the vendor's 2 approvals now meet the 3-trustee quorum
		// (ceil(2/3*3)=2), the revocation does NOT retroactively activate it —
		// it stays pending until another approval re-evaluates the threshold.
		acc, err = GetAccount(vendorAddr)
		require.NoError(t, err)
		require.Nil(t, acc)

		allProposed, err = GetAllProposedAccounts()
		require.NoError(t, err)
		require.True(t, containsPendingAccountAddress(allProposed, vendorAddr))

		// Alice approves → re-evaluates quorum → account becomes active.
		txResult, err = ApproveAccount(vendorAddr, alice, AccountActionOpts{Info: "Alice is approving this account"})
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		acc, err = GetAccount(vendorAddr)
		require.NoError(t, err)
		require.NotNil(t, acc)
		require.Equal(t, vendorAddr, acc.Address)

		prop, err := GetProposedAccount(vendorAddr)
		require.NoError(t, err)
		require.Nil(t, prop)
	})

	// ── The now-active Vendor (VID=vid) can add/update/query a model for its own
	// VID, but adding a model for a different VID is rejected (auth-demo.sh:1336-1362).
	t.Run("VendorModelLifecycle", func(t *testing.T) {
		mpid := rand.Intn(65534) + 1
		productName := "Device #1"

		// Add a model for the vendor's own VID → success.
		txResult, err := utils.ExecuteTx("tx", "model", "add-model",
			"--vid", fmt.Sprintf("%d", vid),
			"--pid", fmt.Sprintf("%d", mpid),
			"--productName", productName,
			"--productLabel", "Device Description",
			"--commissioningCustomFlow", "0",
			"--deviceTypeID", "12",
			"--partNumber", "12",
			"--enhancedSetupFlowOptions", "0",
			"--from", vendorName,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Add a model for a different VID → rejected: the vendor is not associated
		// with that VID. (vid ≤ 65534, so wrongVid ≤ 65535 stays a valid VID and
		// reaches the vendor-VID permission check.)
		wrongVid := vid + 1
		txResult, err = utils.ExecuteTx("tx", "model", "add-model",
			"--vid", fmt.Sprintf("%d", wrongVid),
			"--pid", fmt.Sprintf("%d", mpid),
			"--productName", productName,
			"--productLabel", "Device Description",
			"--commissioningCustomFlow", "0",
			"--deviceTypeID", "12",
			"--partNumber", "12",
			"--enhancedSetupFlowOptions", "0",
			"--from", vendorName,
		)
		combined := ""
		if err != nil {
			combined += err.Error()
		}
		if txResult != nil {
			combined += txResult.RawLog
		}
		require.Contains(t, combined,
			fmt.Sprintf("transaction should be signed by a vendor account containing the vendorID %d", wrongVid))

		// Update the model → success.
		txResult, err = utils.ExecuteTx("tx", "model", "update-model",
			"--vid", fmt.Sprintf("%d", vid),
			"--pid", fmt.Sprintf("%d", mpid),
			"--productName", productName,
			"--productLabel", "Device Description",
			"--partNumber", "12",
			"--enhancedSetupFlowOptions", "2",
			"--from", vendorName,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Query the model and confirm it is present with the expected fields.
		out, err := utils.ExecuteCLI("query", "model", "get-model",
			"--vid", fmt.Sprintf("%d", vid),
			"--pid", fmt.Sprintf("%d", mpid),
			"-o", "json",
		)
		require.NoError(t, err)
		require.False(t, utils.IsNotFound(out))
		require.Contains(t, string(out), productName)
		require.Contains(t, string(out), fmt.Sprintf("%d", mpid))
	})
}

func TestAuthDemoVendorAccount(t *testing.T) {
	jack := testconstants.JackAccount
	alice := testconstants.AliceAccount

	vid := rand.Intn(65535)

	t.Run("VendorApprovedByOneApproval", func(t *testing.T) {
		name := fmt.Sprintf("vendor%d", rand.Intn(99999))
		err := AddKey(name)
		require.NoError(t, err)

		userAddr, err := GetAddress(name)
		require.NoError(t, err)
		userPubkey, err := GetPubkey(name)
		require.NoError(t, err)

		jackAddr, err := GetAddress(jack)
		require.NoError(t, err)
		aliceAddr, err := GetAddress(alice)
		require.NoError(t, err)

		// Jack proposes Vendor — only needs 1/3 trustee approvals, so jack's proposal is enough
		txResult, err := ProposeAccount(userAddr, userPubkey, "Vendor", jack, ProposeAccountOpts{Info: "Jack is proposing this account", VID: vid})
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// With Vendor role, 1 approval is sufficient so account is already active.
		all, err := GetAllAccounts()
		require.NoError(t, err)
		require.True(t, containsAccountAddress(all, userAddr))

		acc, err := GetAccount(userAddr)
		require.NoError(t, err)
		require.NotNil(t, acc)
		require.Equal(t, userAddr, acc.Address)
		require.Len(t, acc.Approvals, 1)
		require.Equal(t, jackAddr, acc.Approvals[0].Address)
		require.Equal(t, "Jack is proposing this account", acc.Approvals[0].Info)

		prop, err := GetProposedAccount(userAddr)
		require.NoError(t, err)
		require.Nil(t, prop)

		allProposed, err := GetAllProposedAccounts()
		require.NoError(t, err)
		require.False(t, containsPendingAccountAddress(allProposed, userAddr))

		_ = aliceAddr
	})

	t.Run("VendorWithPidRanges_Success", func(t *testing.T) {
		pid := rand.Intn(65535)
		pidRanges := fmt.Sprintf("%d-%d", pid, pid)
		vidWithPids := vid + 1

		name := fmt.Sprintf("vendor%d", rand.Intn(99999))
		err := AddKey(name)
		require.NoError(t, err)

		userAddr, err := GetAddress(name)
		require.NoError(t, err)
		userPubkey, err := GetPubkey(name)
		require.NoError(t, err)

		jackAddr, err := GetAddress(jack)
		require.NoError(t, err)

		txResult, err := ProposeAccount(userAddr, userPubkey, "Vendor", jack, ProposeAccountOpts{Info: "Jack is proposing this account", VID: vidWithPids, PidRanges: pidRanges})
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		all, err := GetAllAccounts()
		require.NoError(t, err)
		require.True(t, containsAccountAddress(all, userAddr))

		acc, err := GetAccount(userAddr)
		require.NoError(t, err)
		require.NotNil(t, acc)
		require.Equal(t, userAddr, acc.Address)
		require.Len(t, acc.Approvals, 1)
		require.Equal(t, jackAddr, acc.Approvals[0].Address)

		prop, err := GetProposedAccount(userAddr)
		require.NoError(t, err)
		require.Nil(t, prop)

		allProposed, err := GetAllProposedAccounts()
		require.NoError(t, err)
		require.False(t, containsPendingAccountAddress(allProposed, userAddr))
	})

	t.Run("VendorWithInvalidPidRanges_Fails", func(t *testing.T) {
		invalidPidRanges := "100-101,1-200"

		name := fmt.Sprintf("vendor%d", rand.Intn(99999))
		err := AddKey(name)
		require.NoError(t, err)

		userAddr, err := GetAddress(name)
		require.NoError(t, err)
		userPubkey, err := GetPubkey(name)
		require.NoError(t, err)

		out, err := utils.ExecuteCLI("tx", "auth", "propose-add-account",
			"--info", "Jack is proposing this account",
			"--address", userAddr,
			"--pubkey", userPubkey,
			"--roles", "Vendor",
			"--vid", fmt.Sprintf("%d", vid),
			"--pid_ranges", invalidPidRanges,
			"--from", jack,
			"--yes", "-o", "json", "--keyring-backend", "test",
		)
		// Expect error about invalid PID range.
		combined := string(out)
		if err != nil {
			combined += err.Error()
		}
		require.Contains(t, combined, "invalid PID Range is provided")

		prop, _ := GetProposedAccount(userAddr)
		require.Nil(t, prop)

		allProposed, _ := GetAllProposedAccounts()
		require.False(t, containsPendingAccountAddress(allProposed, userAddr))

		all, _ := GetAllAccounts()
		require.False(t, containsAccountAddress(all, userAddr))
	})

	t.Run("NodeAdminWithVendorRole_NeedsMoreApprovals", func(t *testing.T) {
		name := fmt.Sprintf("user%d", rand.Intn(99999))
		err := AddKey(name)
		require.NoError(t, err)

		userAddr, err := GetAddress(name)
		require.NoError(t, err)
		userPubkey, err := GetPubkey(name)
		require.NoError(t, err)

		jackAddr, err := GetAddress(jack)
		require.NoError(t, err)
		aliceAddr, err := GetAddress(alice)
		require.NoError(t, err)

		txResult, err := ProposeAccount(userAddr, userPubkey, "Vendor,NodeAdmin", jack, ProposeAccountOpts{Info: "Jack is proposing this account", VID: vid})
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// NodeAdmin requires 2/3 approval so not yet active.
		all, err := GetAllAccounts()
		require.NoError(t, err)
		require.False(t, containsAccountAddress(all, userAddr))

		acc, err := GetAccount(userAddr)
		require.NoError(t, err)
		require.Nil(t, acc)

		allProposed, err := GetAllProposedAccounts()
		require.NoError(t, err)
		require.True(t, containsPendingAccountAddress(allProposed, userAddr))

		// Alice approves — now has 2 approvals, should become active.
		txResult, err = ApproveAccount(userAddr, alice, AccountActionOpts{Info: "Alice is approving this account"})
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		all, err = GetAllAccounts()
		require.NoError(t, err)
		require.True(t, containsAccountAddress(all, userAddr))

		acc, err = GetAccount(userAddr)
		require.NoError(t, err)
		require.NotNil(t, acc)
		require.Equal(t, userAddr, acc.Address)
		require.Len(t, acc.Approvals, 2)
		approvers := []string{acc.Approvals[0].Address, acc.Approvals[1].Address}
		require.Contains(t, approvers, jackAddr)
		require.Contains(t, approvers, aliceAddr)

		allProposed, err = GetAllProposedAccounts()
		require.NoError(t, err)
		require.False(t, containsPendingAccountAddress(allProposed, userAddr))

		prop, err := GetProposedAccount(userAddr)
		require.NoError(t, err)
		require.Nil(t, prop)
	})
}
