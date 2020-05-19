//nolint:testpackage
package keeper

//nolint:goimports
import (
	"git.dsr-corporation.com/zb-ledger/zb-ledger/integration_tests/constants"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/utils/pagination"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/auth/internal/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	"testing"
)

func TestQuerier_QueryAccount(t *testing.T) {
	setup := Setup()

	// store account
	account := types.NewAccount(testconstants.Address1, testconstants.PubKey1, types.AccountRoles{types.Trustee})
	setup.Keeper.SetAccount(setup.Ctx, account)

	// query account
	result, _ := setup.Querier(
		setup.Ctx,
		[]string{QueryAccount},
		abci.RequestQuery{Data: queryAccountParas(setup, account.Address)},
	)

	var receivedAccount types.Account
	_ = setup.Cdc.UnmarshalJSON(result, &receivedAccount)

	// check
	require.Equal(t, account.Address, receivedAccount.Address)
	require.Equal(t, account.PubKey, receivedAccount.PubKey)
	require.Equal(t, account.Roles, receivedAccount.Roles)
}

func TestQuerier_QueryAccount_ForNotFound(t *testing.T) {
	setup := Setup()

	// query proposed certificate
	_, err := setup.Querier(
		setup.Ctx,
		[]string{QueryAccount},
		abci.RequestQuery{Data: queryAccountParas(setup, testconstants.Address1)},
	)

	// check
	require.NotNil(t, err)
	require.Equal(t, types.CodeAccountDoesNotExist, err.Code())
}

func TestQuerier_QueryAllAccounts(t *testing.T) {
	setup := Setup()

	// store active account
	account1 := types.NewAccount(testconstants.Address1, testconstants.PubKey1, types.AccountRoles{types.Trustee})
	setup.Keeper.SetAccount(setup.Ctx, account1)

	// store active account
	account2 := types.NewAccount(testconstants.Address2, testconstants.PubKey2, types.AccountRoles{types.Vendor})
	setup.Keeper.SetAccount(setup.Ctx, account2)

	// store second proposed account
	proposedAccount := types.NewPendingAccount(
		testconstants.Address3,
		testconstants.PubKey3,
		types.AccountRoles{types.Vendor},
		testconstants.Address1,
	)
	setup.Keeper.SetProposedAccount(setup.Ctx, proposedAccount)

	// query proposed account
	result, _ := setup.Querier(
		setup.Ctx,
		[]string{QueryAllAccounts},
		abci.RequestQuery{Data: queryListEmptyQueryParams(setup)},
	)

	var listAccounts types.ListAccountItems
	_ = setup.Cdc.UnmarshalJSON(result, &listAccounts)

	// check
	require.Equal(t, 2, len(listAccounts.Items))
	require.Equal(t, account1, listAccounts.Items[0])
	require.Equal(t, account2, listAccounts.Items[1])
}

func TestQuerier_QueryAllProposedAccounts(t *testing.T) {
	setup := Setup()

	// store active account
	account := types.NewAccount(testconstants.Address1, testconstants.PubKey1, types.AccountRoles{types.Trustee})
	setup.Keeper.SetAccount(setup.Ctx, account)

	// store proposed account
	account1 := types.NewPendingAccount(
		testconstants.Address2,
		testconstants.PubKey2,
		types.AccountRoles{types.Trustee},
		testconstants.Address1,
	)
	setup.Keeper.SetProposedAccount(setup.Ctx, account1)

	// store second proposed account
	account2 := types.NewPendingAccount(
		testconstants.Address3,
		testconstants.PubKey3,
		types.AccountRoles{types.Vendor},
		testconstants.Address1,
	)
	setup.Keeper.SetProposedAccount(setup.Ctx, account2)

	// query proposed account
	result, _ := setup.Querier(
		setup.Ctx,
		[]string{QueryAllProposedAccounts},
		abci.RequestQuery{Data: queryListEmptyQueryParams(setup)},
	)

	var listProposedAccounts types.ListProposedAccountItems
	_ = setup.Cdc.UnmarshalJSON(result, &listProposedAccounts)

	// check
	require.Equal(t, 2, len(listProposedAccounts.Items))
	require.Equal(t, account1, listProposedAccounts.Items[0])
	require.Equal(t, account2, listProposedAccounts.Items[1])
}

func queryAccountParas(setup TestSetup, address sdk.AccAddress) []byte {
	params := types.NewQueryAccountParams(address)
	res, _ := setup.Cdc.MarshalJSON(params)

	return res
}

func queryListEmptyQueryParams(setup TestSetup) []byte {
	paginationParams := pagination.NewPaginationParams(0, 0)
	res, _ := setup.Cdc.MarshalJSON(paginationParams)

	return res
}
