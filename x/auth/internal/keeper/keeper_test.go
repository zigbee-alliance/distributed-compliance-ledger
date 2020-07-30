//nolint:testpackage
package keeper

//nolint:goimports
import (
	"git.dsr-corporation.com/zb-ledger/zb-ledger/integration_tests/constants"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/auth/internal/types"
	"github.com/stretchr/testify/require"
	"testing"
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
}

func TestKeeper_ProposedAccountGetSet(t *testing.T) {
	setup := Setup()

	// check if pending account present
	require.False(t, setup.Keeper.IsProposedAccountPresent(setup.Ctx, testconstants.Address1))

	// no account before its created
	require.Panics(t, func() {
		setup.Keeper.GetProposedAccount(setup.Ctx, testconstants.Address1)
	})

	// store proposed account
	account := types.NewPendingAccount(
		testconstants.Address1,
		testconstants.PubKey1,
		types.AccountRoles{types.Trustee},
		testconstants.Address2,
	)

	setup.Keeper.SetProposedAccount(setup.Ctx, account)

	// check if account present
	require.True(t, setup.Keeper.IsProposedAccountPresent(setup.Ctx, testconstants.Address1))

	// get account
	receivedAccount := setup.Keeper.GetProposedAccount(setup.Ctx, testconstants.Address1)
	require.Equal(t, account.Address, receivedAccount.Address)
	require.Equal(t, account.PubKey, receivedAccount.PubKey)
	require.Equal(t, account.Roles, receivedAccount.Roles)
	require.Equal(t, account.Approvals, receivedAccount.Approvals)

	// delete account
	setup.Keeper.DeleteProposedAccount(setup.Ctx, testconstants.Address1)
	require.False(t, setup.Keeper.IsProposedAccountPresent(setup.Ctx, testconstants.Address1))
	require.Panics(t, func() {
		setup.Keeper.GetProposedAccount(setup.Ctx, testconstants.Address1)
	})
}

func TestKeeper_AccountNumber(t *testing.T) {
	setup := Setup()

	for i := uint64(0); i < 5; i++ {
		require.Equal(t, i, setup.Keeper.GetNextAccountNumber(setup.Ctx))
	}
}
