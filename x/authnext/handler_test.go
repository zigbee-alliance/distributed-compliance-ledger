package authnext

import (
	"git.dsr-corporation.com/zb-ledger/zb-ledger/integration_tests/constants"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/authnext/internal/keeper"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/authnext/internal/types"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/authz"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	"testing"
)

func TestHandler_CreateAccount(t *testing.T) {
	setup := Setup()
	setup.AuthzKeeper.AssignRole(setup.Ctx, test_constants.Address1, authz.Trustee)

	// add new account
	msgAddAccount := types.NewMsgAddAccount(test_constants.Address2, test_constants.PubKey, test_constants.Address1)
	result := setup.Handler(setup.Ctx, msgAddAccount)
	require.Equal(t, sdk.CodeOK, result.Code)

	// query account
	receivedAccount := queryAccount(setup, msgAddAccount.Address)

	// check
	require.Equal(t, receivedAccount.Address, msgAddAccount.Address)
	require.Equal(t, receivedAccount.PubKey, msgAddAccount.PublicKey)
	require.Nil(t, receivedAccount.Roles)
}

func TestHandler_ByNonTrustee(t *testing.T) {
	setup := Setup()

	for _, role := range []authz.AccountRole{authz.Administrator, authz.Vendor, authz.TestHouse} {
		setup.AuthzKeeper.AssignRole(setup.Ctx, test_constants.Address1, role)

		msgAddAccount := types.NewMsgAddAccount(test_constants.Address2, test_constants.PubKey, test_constants.Address1)

		// add new account
		result := setup.Handler(setup.Ctx, msgAddAccount)
		require.Equal(t, sdk.CodeUnauthorized, result.Code)
	}
}

func TestHandler_CreateAccount_Twice(t *testing.T) {
	setup := Setup()

	setup.AuthzKeeper.AssignRole(setup.Ctx, test_constants.Address1, authz.Trustee)

	// add new account
	msgAddAccount := types.NewMsgAddAccount(test_constants.Address2, test_constants.PubKey, test_constants.Address1)
	result := setup.Handler(setup.Ctx, msgAddAccount)
	require.Equal(t, sdk.CodeOK, result.Code)

	// add same account second time
	result = setup.Handler(setup.Ctx, msgAddAccount)
	require.Equal(t, sdk.CodeInvalidAddress, result.Code)

}

func queryAccount(setup TestSetup, account sdk.AccAddress) types.AccountHeader {
	result, _ := setup.Querier(
		setup.Ctx,
		[]string{keeper.QueryAccount, account.String()},
		abci.RequestQuery{},
	)

	var accountHeader types.AccountHeader
	_ = setup.Cdc.UnmarshalJSON(result, &accountHeader)
	return accountHeader
}
