package dclauth

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/utils"
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
		out, err := QueryAccountRaw(userAddr)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")

		out, err = QueryProposedAccountToRevoke(userAddr)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")

		out, err = QueryProposedAccount(userAddr)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")

		out, err = QueryRevokedAccount(userAddr)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")
	})

	t.Run("InitialListsEmpty", func(t *testing.T) {
		out, err := QueryAllProposedAccounts()
		require.NoError(t, err)
		require.Contains(t, string(out), "[]")

		out, err = QueryAllProposedAccountsToRevoke()
		require.NoError(t, err)
		require.Contains(t, string(out), "[]")

		out, err = QueryAllRevokedAccounts()
		require.NoError(t, err)
		require.Contains(t, string(out), "[]")
	})

	t.Run("JackProposes", func(t *testing.T) {
		txResult, err := utils.ExecuteTx("tx", "auth", "propose-add-account",
			"--info", "Jack is proposing this account",
			"--address", userAddr,
			"--pubkey", userPubkey,
			"--roles", "NodeAdmin",
			"--from", jack,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Account not yet active
		out, err := QueryAllAccountsRaw()
		require.NoError(t, err)
		require.NotContains(t, string(out), userAddr)

		out, err = QueryAccountRaw(userAddr)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")

		// Now in proposed list
		out, err = QueryProposedAccount(userAddr)
		require.NoError(t, err)
		require.Contains(t, string(out), userAddr)
		require.Contains(t, string(out), jackAddr)
		require.Contains(t, string(out), "Jack is proposing this account")
		require.NotContains(t, string(out), aliceAddr)

		out, err = QueryAllProposedAccounts()
		require.NoError(t, err)
		require.Contains(t, string(out), userAddr)
	})

	t.Run("AliceApproves", func(t *testing.T) {
		txResult, err := utils.ExecuteTx("tx", "auth", "approve-add-account",
			"--address", userAddr,
			"--info", "Alice is approving this account",
			"--from", alice,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Alice cannot reject after approving
		txBad, err := utils.ExecuteTx("tx", "auth", "reject-add-account",
			"--address", userAddr,
			"--info", "Alice is rejecting this account",
			"--from", alice,
		)
		// Either error or non-zero code
		if err == nil {
			require.NotEqual(t, uint32(0), txBad.Code)
		}

		// Account is now active
		out, err := QueryAllAccountsRaw()
		require.NoError(t, err)
		require.Contains(t, string(out), userAddr)

		out, err = QueryAccountRaw(userAddr)
		require.NoError(t, err)
		require.Contains(t, string(out), userAddr)
		require.Contains(t, string(out), jackAddr)
		require.Contains(t, string(out), aliceAddr)
		require.Contains(t, string(out), "Alice is approving this account")
		require.Contains(t, string(out), "Jack is proposing this account")

		// No longer in proposed list
		out, err = QueryAllProposedAccounts()
		require.NoError(t, err)
		require.Contains(t, string(out), "[]")

		out, err = QueryProposedAccount(userAddr)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")
	})

	t.Run("AliceProposeRevoke", func(t *testing.T) {
		txResult, err := utils.ExecuteTx("tx", "auth", "propose-revoke-account",
			"--address", userAddr,
			"--info", "Alice proposes to revoke account",
			"--from", alice,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Still in active accounts (not enough approvals to revoke)
		out, err := QueryAllAccountsRaw()
		require.NoError(t, err)
		require.Contains(t, string(out), userAddr)

		// In proposed-to-revoke list
		out, err = QueryAllProposedAccountsToRevoke()
		require.NoError(t, err)
		require.Contains(t, string(out), userAddr)

		out, err = QueryProposedAccountToRevoke(userAddr)
		require.NoError(t, err)
		require.Contains(t, string(out), userAddr)
		require.Contains(t, string(out), aliceAddr)
		require.Contains(t, string(out), "Alice proposes to revoke account")

		// Not yet revoked
		out, err = QueryAllRevokedAccounts()
		require.NoError(t, err)
		require.NotContains(t, string(out), userAddr)
	})

	t.Run("BobApprovesRevoke", func(t *testing.T) {
		txResult, err := utils.ExecuteTx("tx", "auth", "approve-revoke-account",
			"--address", userAddr,
			"--from", bob,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Now revoked
		out, err := QueryAllRevokedAccounts()
		require.NoError(t, err)
		require.Contains(t, string(out), userAddr)
		require.Contains(t, string(out), "TrusteeVoting")

		// No longer in active accounts
		out, err = QueryAllAccountsRaw()
		require.NoError(t, err)
		require.NotContains(t, string(out), userAddr)

		out, err = QueryAllProposedAccountsToRevoke()
		require.NoError(t, err)
		require.Contains(t, string(out), "[]")

		out, err = QueryProposedAccountToRevoke(userAddr)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")

		out, err = QueryRevokedAccount(userAddr)
		require.NoError(t, err)
		require.Contains(t, string(out), userAddr)
		require.Contains(t, string(out), "TrusteeVoting")
	})

	t.Run("ReAddAfterRevoke_ProposeApprove", func(t *testing.T) {
		// Jack proposes again
		txResult, err := utils.ExecuteTx("tx", "auth", "propose-add-account",
			"--info", "Jack is proposing this account",
			"--address", userAddr,
			"--pubkey", userPubkey,
			"--roles", "NodeAdmin",
			"--from", jack,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Still revoked in revoked list
		out, err := QueryAllRevokedAccounts()
		require.NoError(t, err)
		require.Contains(t, string(out), userAddr)

		// Not yet active
		out, err = QueryAllAccountsRaw()
		require.NoError(t, err)
		require.NotContains(t, string(out), userAddr)

		out, err = QueryProposedAccount(userAddr)
		require.NoError(t, err)
		require.Contains(t, string(out), userAddr)

		// Alice approves
		txResult, err = utils.ExecuteTx("tx", "auth", "approve-add-account",
			"--address", userAddr,
			"--info", "Alice is approving this account",
			"--from", alice,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// No longer in revoked list
		out, err = QueryAllRevokedAccounts()
		require.NoError(t, err)
		require.NotContains(t, string(out), userAddr)

		// Active again
		out, err = QueryAllAccountsRaw()
		require.NoError(t, err)
		require.Contains(t, string(out), userAddr)
	})

	t.Run("RejectScenario", func(t *testing.T) {
		// Propose-revoke again then approve-revoke
		txResult, err := utils.ExecuteTx("tx", "auth", "propose-revoke-account",
			"--address", userAddr,
			"--info", "Alice proposes to revoke account",
			"--from", alice,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		txResult, err = utils.ExecuteTx("tx", "auth", "approve-revoke-account",
			"--address", userAddr,
			"--from", bob,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Jack proposes again for rejection test
		txResult, err = utils.ExecuteTx("tx", "auth", "propose-add-account",
			"--info", "Jack is proposing this account",
			"--address", userAddr,
			"--pubkey", userPubkey,
			"--roles", "NodeAdmin",
			"--from", jack,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Alice rejects
		txResult, err = utils.ExecuteTx("tx", "auth", "reject-add-account",
			"--address", userAddr,
			"--info", "Alice is rejecting this account",
			"--from", alice,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Still in proposed (alice rejected but not enough rejections)
		out, err := QueryAllProposedAccounts()
		require.NoError(t, err)
		require.Contains(t, string(out), userAddr)

		out, err = QueryAllAccountsRaw()
		require.NoError(t, err)
		require.NotContains(t, string(out), userAddr)

		out, err = QueryAllRejectedAccounts()
		require.NoError(t, err)
		require.Contains(t, string(out), "[]")

		// Bob rejects
		txResult, err = utils.ExecuteTx("tx", "auth", "reject-add-account",
			"--address", userAddr,
			"--info", "Bob is rejecting this account",
			"--from", bob,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Bob cannot reject again
		txBad, errBad := utils.ExecuteTx("tx", "auth", "reject-add-account",
			"--address", userAddr,
			"--info", "Bob is rejecting this account",
			"--from", bob,
		)
		if errBad == nil {
			require.NotEqual(t, uint32(0), txBad.Code)
		}

		// Now in rejected list (enough rejections)
		out, err = QueryAllRejectedAccounts()
		require.NoError(t, err)
		require.Contains(t, string(out), userAddr)

		out, err = QueryAllProposedAccounts()
		require.NoError(t, err)
		require.NotContains(t, string(out), userAddr)

		out, err = QueryAllAccountsRaw()
		require.NoError(t, err)
		require.NotContains(t, string(out), userAddr)

		out, err = QueryRejectedAccount(userAddr)
		require.NoError(t, err)
		require.Contains(t, string(out), userAddr)
		require.Contains(t, string(out), jackAddr)
		require.Contains(t, string(out), "Jack is proposing this account")
		require.Contains(t, string(out), aliceAddr)
		require.Contains(t, string(out), "Alice is rejecting this account")
		require.Contains(t, string(out), bobAddr)
		require.Contains(t, string(out), "Bob is rejecting this account")
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
		txResult, err := utils.ExecuteTx("tx", "auth", "propose-add-account",
			"--info", "Jack is proposing this account",
			"--address", userAddr,
			"--pubkey", userPubkey,
			"--roles", "NodeAdmin",
			"--from", jack,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Jack rejects his own proposal
		txResult, err = utils.ExecuteTx("tx", "auth", "reject-add-account",
			"--address", userAddr,
			"--info", "Jack is rejecting this account",
			"--from", jack,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Not in proposed (jack removed his approval+proposal)
		out, err := QueryProposedAccount(userAddr)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")

		// Not in rejected (single rejection doesn't reach quorum with 3 trustees)
		out, err = QueryRejectedAccount(userAddr)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")

		// Not in approved
		out, err = QueryAccountRaw(userAddr)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")
	})
}

// TestAuthDemoDynamicTrusteeCount is intentionally skipped: it requires exactly
// 3 initial trustees and leaves extra trustees on chain if it fails mid-way,
// corrupting subsequent tests. Re-enable once a robust cleanup mechanism exists.
func TestAuthDemoDynamicTrusteeCount(t *testing.T) {
	t.Skip("skipped: requires clean chain with exactly 3 trustees; see REWRITE_PLAN.md")
	jack := testconstants.JackAccount
	alice := testconstants.AliceAccount
	bob := testconstants.BobAccount

	vid := rand.Intn(65534) + 1

	// ── Create two new trustees → 5 trustees total ─────────────────────────

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

	t.Run("AddTwoNewTrustees", func(t *testing.T) {
		// Jack proposes + Alice approves new_trustee1 → 4 trustees
		txResult, err := utils.ExecuteTx("tx", "auth", "propose-add-account",
			"--address", newTrustee1Addr,
			"--pubkey", newTrustee1Pubkey,
			"--roles", "Trustee",
			"--from", jack,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		txResult, err = utils.ExecuteTx("tx", "auth", "approve-add-account",
			"--address", newTrustee1Addr,
			"--from", alice,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Jack proposes new_trustee2, then can reject and re-approve
		txResult, err = utils.ExecuteTx("tx", "auth", "propose-add-account",
			"--address", newTrustee2Addr,
			"--pubkey", newTrustee2Pubkey,
			"--roles", "Trustee",
			"--from", jack,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Jack can reject even after proposing
		txResult, err = utils.ExecuteTx("tx", "auth", "reject-add-account",
			"--address", newTrustee2Addr,
			"--info", "Jack is rejecting this account",
			"--from", jack,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Jack re-approves
		txResult, err = utils.ExecuteTx("tx", "auth", "approve-add-account",
			"--address", newTrustee2Addr,
			"--info", "Jack re-approving",
			"--from", jack,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Alice approves → 2nd approval with 4 trustees = quorum (ceil(4*2/3)=3, need 3... hmm)
		// Actually with 4 trustees: ceil(4*2/3)=ceil(2.67)=3 approvals needed.
		// Jack proposed (1) + Alice (2) + Bob (3) → active.
		txResult, err = utils.ExecuteTx("tx", "auth", "approve-add-account",
			"--address", newTrustee2Addr,
			"--from", alice,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		txResult, err = utils.ExecuteTx("tx", "auth", "approve-add-account",
			"--address", newTrustee2Addr,
			"--from", bob,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Both new trustees are now active
		out, err := QueryAccountRaw(newTrustee1Addr)
		require.NoError(t, err)
		require.Contains(t, string(out), newTrustee1Addr)

		out, err = QueryAccountRaw(newTrustee2Addr)
		require.NoError(t, err)
		require.Contains(t, string(out), newTrustee2Addr)
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
		txResult, err := utils.ExecuteTx("tx", "auth", "propose-add-account",
			"--info", "Jack is proposing this account",
			"--address", vendorAddr,
			"--pubkey", vendorPubkey,
			"--roles", "Vendor",
			"--vid", fmt.Sprintf("%d", vid),
			"--from", jack,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// With 5 trustees, vendor needs ceil(5/3)=2 approvals, so Jack's proposal alone is not enough
		out, err := QueryAccountRaw(vendorAddr)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")

		out, err = QueryAllProposedAccounts()
		require.NoError(t, err)
		require.Contains(t, string(out), vendorAddr)

		// Alice approves → 2 approvals = quorum → account active
		txResult, err = utils.ExecuteTx("tx", "auth", "approve-add-account",
			"--address", vendorAddr,
			"--info", "Alice is approving this account",
			"--from", alice,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		out, err = QueryAccountRaw(vendorAddr)
		require.NoError(t, err)
		require.Contains(t, string(out), vendorAddr)

		out, err = QueryProposedAccount(vendorAddr)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")
	})

	// ── Revoke vendor: with 5 trustees needs ceil(10/3)=4 approvals ────────

	t.Run("RevokeVendorWith5Trustees", func(t *testing.T) {
		// Alice proposes revocation (1)
		txResult, err := utils.ExecuteTx("tx", "auth", "propose-revoke-account",
			"--address", vendorAddr,
			"--info", "Alice proposes to revoke",
			"--from", alice,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Bob approves (2) — still not enough, need 4 with 5 trustees
		txResult, err = utils.ExecuteTx("tx", "auth", "approve-revoke-account",
			"--address", vendorAddr,
			"--from", bob,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Account still active (need 4 approvals)
		out, err := QueryAccountRaw(vendorAddr)
		require.NoError(t, err)
		require.Contains(t, string(out), vendorAddr)

		// Revoke new_trustee1 → 4 trustees total
		// With 4 trustees: revocation needs ceil(8/3)=3 approvals → we have alice+bob already
		txResult, err = utils.ExecuteTx("tx", "auth", "propose-revoke-account",
			"--address", newTrustee1Addr,
			"--from", alice,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		txResult, err = utils.ExecuteTx("tx", "auth", "approve-revoke-account",
			"--address", newTrustee1Addr,
			"--from", bob,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		txResult, err = utils.ExecuteTx("tx", "auth", "approve-revoke-account",
			"--address", newTrustee1Addr,
			"--from", jack,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		txResult, err = utils.ExecuteTx("tx", "auth", "approve-revoke-account",
			"--address", newTrustee1Addr,
			"--from", newTrustee1Name,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// new_trustee1 is revoked → 4 trustees remain
		out, err = QueryRevokedAccount(newTrustee1Addr)
		require.NoError(t, err)
		require.Contains(t, string(out), newTrustee1Addr)

		// Now approve vendor revocation — with 4 trustees need ceil(8/3)=3 approvals
		// alice(1) + bob(2) + jack(3) = 3 → quorum
		txResult, err = utils.ExecuteTx("tx", "auth", "approve-revoke-account",
			"--address", vendorAddr,
			"--from", jack,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Vendor is now revoked
		out, err = QueryRevokedAccount(vendorAddr)
		require.NoError(t, err)
		require.Contains(t, string(out), vendorAddr)
		require.Contains(t, string(out), "TrusteeVoting")
	})

	// ── Reject scenario with dynamic trustee count ──────────────────────────
	// With 4 trustees (jack, alice, bob, new_trustee2), rejection needs ceil(8/3)=3

	t.Run("RejectWithDynamicTrusteeCount", func(t *testing.T) {
		// Jack re-proposes the revoked vendor
		txResult, err := utils.ExecuteTx("tx", "auth", "propose-add-account",
			"--info", "Jack is proposing this account",
			"--address", vendorAddr,
			"--pubkey", vendorPubkey,
			"--roles", "Vendor,NodeAdmin",
			"--vid", fmt.Sprintf("%d", vid),
			"--from", jack,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Bob rejects (1 rejection — not enough with 4 trustees, need 3)
		txResult, err = utils.ExecuteTx("tx", "auth", "reject-add-account",
			"--address", vendorAddr,
			"--info", "Bob is rejecting this account",
			"--from", bob,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Still in proposed
		out, err := QueryAllProposedAccounts()
		require.NoError(t, err)
		require.Contains(t, string(out), vendorAddr)

		// Revoke new_trustee2 → 3 trustees (jack, alice, bob)
		txResult, err = utils.ExecuteTx("tx", "auth", "propose-revoke-account",
			"--address", newTrustee2Addr,
			"--from", alice,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		txResult, err = utils.ExecuteTx("tx", "auth", "approve-revoke-account",
			"--address", newTrustee2Addr,
			"--from", bob,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		txResult, err = utils.ExecuteTx("tx", "auth", "approve-revoke-account",
			"--address", newTrustee2Addr,
			"--from", jack,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// new_trustee2 is revoked → 3 trustees
		out, err = QueryRevokedAccount(newTrustee2Addr)
		require.NoError(t, err)
		require.Contains(t, string(out), newTrustee2Addr)

		// Alice rejects → with 3 trustees need ceil(6/3)=2 rejections → bob+alice=2 → quorum
		txResult, err = utils.ExecuteTx("tx", "auth", "reject-add-account",
			"--address", vendorAddr,
			"--info", "Alice is rejecting this account",
			"--from", alice,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// Account is now in rejected list
		out, err = QueryAllRejectedAccounts()
		require.NoError(t, err)
		require.Contains(t, string(out), vendorAddr)

		out, err = QueryRejectedAccount(vendorAddr)
		require.NoError(t, err)
		require.Contains(t, string(out), vendorAddr)

		out, err = QueryAllProposedAccounts()
		require.NoError(t, err)
		require.NotContains(t, string(out), vendorAddr)
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
		txResult, err := utils.ExecuteTx("tx", "auth", "propose-add-account",
			"--info", "Jack is proposing this account",
			"--address", userAddr,
			"--pubkey", userPubkey,
			"--roles", "Vendor",
			"--vid", fmt.Sprintf("%d", vid),
			"--from", jack,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// With Vendor role, 1 approval is sufficient so account is already active
		out, err := QueryAllAccountsRaw()
		require.NoError(t, err)
		require.Contains(t, string(out), userAddr)

		out, err = QueryAccountRaw(userAddr)
		require.NoError(t, err)
		require.Contains(t, string(out), userAddr)
		require.Contains(t, string(out), jackAddr)
		require.Contains(t, string(out), "Jack is proposing this account")

		out, err = QueryProposedAccount(userAddr)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")

		out, err = QueryAllProposedAccounts()
		require.NoError(t, err)
		require.Contains(t, string(out), "[]")

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

		txResult, err := utils.ExecuteTx("tx", "auth", "propose-add-account",
			"--info", "Jack is proposing this account",
			"--address", userAddr,
			"--pubkey", userPubkey,
			"--roles", "Vendor",
			"--vid", fmt.Sprintf("%d", vidWithPids),
			"--pid_ranges", pidRanges,
			"--from", jack,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		out, err := QueryAllAccountsRaw()
		require.NoError(t, err)
		require.Contains(t, string(out), userAddr)

		out, err = QueryAccountRaw(userAddr)
		require.NoError(t, err)
		require.Contains(t, string(out), userAddr)
		require.Contains(t, string(out), jackAddr)

		out, err = QueryProposedAccount(userAddr)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")

		out, err = QueryAllProposedAccounts()
		require.NoError(t, err)
		require.Contains(t, string(out), "[]")
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
		// Expect error about invalid PID range
		combined := string(out)
		if err != nil {
			combined = combined + err.Error()
		}
		require.Contains(t, combined, "invalid PID Range is provided")

		out2, _ := QueryProposedAccount(userAddr)
		require.Contains(t, string(out2), "Not Found")

		out3, _ := QueryAllProposedAccounts()
		require.Contains(t, string(out3), "[]")

		out4, _ := QueryAllAccountsRaw()
		require.NotContains(t, string(out4), userAddr)
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

		txResult, err := utils.ExecuteTx("tx", "auth", "propose-add-account",
			"--info", "Jack is proposing this account",
			"--address", userAddr,
			"--pubkey", userPubkey,
			"--roles", "Vendor,NodeAdmin",
			"--vid", fmt.Sprintf("%d", vid),
			"--from", jack,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		// NodeAdmin requires 2/3 approval so not yet active
		out, err := QueryAllAccountsRaw()
		require.NoError(t, err)
		require.NotContains(t, string(out), userAddr)

		out, err = QueryAccountRaw(userAddr)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")

		out, err = QueryAllProposedAccounts()
		require.NoError(t, err)
		require.Contains(t, string(out), userAddr)

		// Alice approves — now has 2 approvals, should become active
		txResult, err = utils.ExecuteTx("tx", "auth", "approve-add-account",
			"--address", userAddr,
			"--info", "Alice is approving this account",
			"--from", alice,
		)
		require.NoError(t, err)
		require.Equal(t, uint32(0), txResult.Code)
		_, err = utils.AwaitTxConfirmation(txResult.TxHash)
		require.NoError(t, err)

		out, err = QueryAllAccountsRaw()
		require.NoError(t, err)
		require.Contains(t, string(out), userAddr)

		out, err = QueryAccountRaw(userAddr)
		require.NoError(t, err)
		require.Contains(t, string(out), userAddr)
		require.Contains(t, string(out), jackAddr)
		require.Contains(t, string(out), aliceAddr)
		require.Contains(t, string(out), "Alice is approving this account")

		out, err = QueryAllProposedAccounts()
		require.NoError(t, err)
		require.NotContains(t, string(out), userAddr)

		out, err = QueryProposedAccount(userAddr)
		require.NoError(t, err)
		require.Contains(t, string(out), "Not Found")
	})
}
