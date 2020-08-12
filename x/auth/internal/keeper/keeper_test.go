//nolint:testpackage
package keeper

//nolint:goimports
import (
	"testing"

	testconstants "git.dsr-corporation.com/zb-ledger/zb-ledger/integration_tests/constants"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/auth/internal/types"
	"github.com/stretchr/testify/require"
)

func TestKeeper_AccountGetSet(t *testing.T) {
	setup := Setup()

	// check if account present
	require.False(t, setup.Keeper.IsAccountPresent(setup.Ctx, testconstants.Address1))

	// no account before its created
	require.Panics(t, func() {
		setup.Keeper.GetAccount(setup.Ctx, testconstants.Address1)
	})

	// store account
	account := types.NewAccount(testconstants.Address1, testconstants.PubKey1, types.AccountRoles{types.Trustee})
	setup.Keeper.SetAccount(setup.Ctx, account)

	// check if account present
	require.True(t, setup.Keeper.IsAccountPresent(setup.Ctx, testconstants.Address1))

	// get account
	receivedAccount := setup.Keeper.GetAccount(setup.Ctx, testconstants.Address1)
	require.Equal(t, account.Address, receivedAccount.Address)
	require.Equal(t, account.PubKey, receivedAccount.PubKey)
	require.Equal(t, account.Roles, receivedAccount.Roles)
	require.Equal(t, account.Sequence, receivedAccount.Sequence)
	require.Equal(t, account.AccountNumber, receivedAccount.AccountNumber)

	// check if account has role
	require.True(t, setup.Keeper.HasRole(setup.Ctx, testconstants.Address1, types.Trustee))

	// get all accounts
	accounts := setup.Keeper.GetAllAccounts(setup.Ctx)
	require.Equal(t, 1, len(accounts))
	require.Equal(t, account.Address, accounts[0].Address)

	// count accounts with role
	require.Equal(t, 1, setup.Keeper.CountAccountsWithRole(setup.Ctx, types.Trustee))
	require.Equal(t, 0, setup.Keeper.CountAccountsWithRole(setup.Ctx, types.Vendor))

	// delete account
	setup.Keeper.DeleteAccount(setup.Ctx, testconstants.Address1)
	require.False(t, setup.Keeper.IsAccountPresent(setup.Ctx, testconstants.Address1))
	require.Panics(t, func() {
		setup.Keeper.GetAccount(setup.Ctx, testconstants.Address1)
	})
}

func TestKeeper_PendingAccountGetSet(t *testing.T) {
	setup := Setup()

	// check if pending account present
	require.False(t, setup.Keeper.IsPendingAccountPresent(setup.Ctx, testconstants.Address1))

	// no pending account before its created
	require.Panics(t, func() {
		setup.Keeper.GetPendingAccount(setup.Ctx, testconstants.Address1)
	})

	// store pending account
	pendAcc := types.NewPendingAccount(
		testconstants.Address1,
		testconstants.PubKey1,
		types.AccountRoles{types.Trustee},
		testconstants.Address2,
	)

	setup.Keeper.SetPendingAccount(setup.Ctx, pendAcc)

	// check if pending account present
	require.True(t, setup.Keeper.IsPendingAccountPresent(setup.Ctx, testconstants.Address1))

	// get pending account
	receivedPendAcc := setup.Keeper.GetPendingAccount(setup.Ctx, testconstants.Address1)
	require.Equal(t, pendAcc.Address, receivedPendAcc.Address)
	require.Equal(t, pendAcc.PubKey, receivedPendAcc.PubKey)
	require.Equal(t, pendAcc.Roles, receivedPendAcc.Roles)
	require.Equal(t, pendAcc.Approvals, receivedPendAcc.Approvals)

	// delete pending account
	setup.Keeper.DeletePendingAccount(setup.Ctx, testconstants.Address1)
	require.False(t, setup.Keeper.IsPendingAccountPresent(setup.Ctx, testconstants.Address1))
	require.Panics(t, func() {
		setup.Keeper.GetPendingAccount(setup.Ctx, testconstants.Address1)
	})
}

func TestKeeper_PendingAccountRevocationGetSet(t *testing.T) {
	setup := Setup()

	// check if pending account revocation present
	require.False(t, setup.Keeper.IsPendingAccountRevocationPresent(setup.Ctx, testconstants.Address1))

	// no pending account revocation before its created
	require.Panics(t, func() {
		setup.Keeper.GetPendingAccountRevocation(setup.Ctx, testconstants.Address1)
	})

	// store pending account revocation
	revocation := types.NewPendingAccountRevocation(
		testconstants.Address1,
		testconstants.Address2,
	)

	setup.Keeper.SetPendingAccountRevocation(setup.Ctx, revocation)

	// check if pending account revocation present
	require.True(t, setup.Keeper.IsPendingAccountRevocationPresent(setup.Ctx, testconstants.Address1))

	// get pending account revocation
	receivedRevocation := setup.Keeper.GetPendingAccountRevocation(setup.Ctx, testconstants.Address1)
	require.Equal(t, revocation.Address, receivedRevocation.Address)
	require.Equal(t, revocation.Approvals, receivedRevocation.Approvals)

	// delete pending account revocation
	setup.Keeper.DeletePendingAccountRevocation(setup.Ctx, testconstants.Address1)
	require.False(t, setup.Keeper.IsPendingAccountRevocationPresent(setup.Ctx, testconstants.Address1))
	require.Panics(t, func() {
		setup.Keeper.GetPendingAccountRevocation(setup.Ctx, testconstants.Address1)
	})
}

func TestKeeper_AccountNumber(t *testing.T) {
	setup := Setup()

	for i := uint64(0); i < 5; i++ {
		require.Equal(t, i, setup.Keeper.GetNextAccountNumber(setup.Ctx))
	}
}
