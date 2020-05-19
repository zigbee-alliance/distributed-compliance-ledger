//nolint:testpackage
package auth

// nolint:goimports
import (
	testconstants "git.dsr-corporation.com/zb-ledger/zb-ledger/integration_tests/constants"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/auth/internal/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/crypto"
	"testing"
)

func TestHandler_CreateAccount_OneApprovalIsNeeded(t *testing.T) {
	setup := Setup()

	countTrustees := 2

	for i := 0; i < countTrustees; i++ {
		// store trustee
		trustee := storeTrustee(setup)

		// ensure 1 trustee approval is needed
		require.Equal(t, 1, CountAccountApprovals(setup.Ctx, setup.Keeper))

		// propose account
		result, address, pubkey := proposeAccount(setup, trustee)
		require.Equal(t, sdk.CodeOK, result.Code)

		// ensure active account created
		account := setup.Keeper.GetAccount(setup.Ctx, address)
		require.Equal(t, address, account.Address)
		require.Equal(t, pubkey, account.PubKey)

		// ensure no pending account created
		require.False(t, setup.Keeper.IsProposedAccountPresent(setup.Ctx, address))
	}
}

func TestHandler_CreateAccount_TwoApprovalsAreNeeded(t *testing.T) {
	setup := Setup()

	// store 3 trustees
	trustee1 := storeTrustee(setup)
	trustee2 := storeTrustee(setup)
	_ = storeTrustee(setup)

	// ensure 2 trustee approvals are needed
	require.Equal(t, 2, CountAccountApprovals(setup.Ctx, setup.Keeper))

	// trustee1 propose account
	result, address, pubkey := proposeAccount(setup, trustee1)
	require.Equal(t, sdk.CodeOK, result.Code)

	// ensure proposed account created
	proposedAccount := setup.Keeper.GetProposedAccount(setup.Ctx, address)
	require.Equal(t, address, proposedAccount.Address)
	require.Equal(t, []sdk.AccAddress{trustee1}, proposedAccount.Approvals)

	// ensure no active account created
	require.False(t, setup.Keeper.IsAccountPresent(setup.Ctx, address))

	// trustee2 approves account
	approveAddAccount := types.NewMsgApproveAddAccount(address, trustee2)
	result = setup.Handler(setup.Ctx, approveAddAccount)
	require.Equal(t, sdk.CodeOK, result.Code)

	// active account must be created
	account := setup.Keeper.GetAccount(setup.Ctx, address)
	require.Equal(t, address, account.Address)
	require.Equal(t, pubkey, account.PubKey)

	// ensure no pending account created
	require.False(t, setup.Keeper.IsProposedAccountPresent(setup.Ctx, address))
}

func TestHandler_CreateAccount_ThreeApprovalsAreNeeded(t *testing.T) {
	setup := Setup()

	// store 3 trustees
	trustee1 := storeTrustee(setup)
	trustee2 := storeTrustee(setup)
	trustee3 := storeTrustee(setup)
	_ = storeTrustee(setup)

	// ensure 3 trustee approvals are needed
	require.Equal(t, 3, CountAccountApprovals(setup.Ctx, setup.Keeper))

	// trustee1 propose account
	result, address, pubkey := proposeAccount(setup, trustee1)
	require.Equal(t, sdk.CodeOK, result.Code)

	// ensure proposed account created
	proposedAccount := setup.Keeper.GetProposedAccount(setup.Ctx, address)
	require.Equal(t, address, proposedAccount.Address)
	require.Equal(t, []sdk.AccAddress{trustee1}, proposedAccount.Approvals)

	// ensure no active account created
	require.False(t, setup.Keeper.IsAccountPresent(setup.Ctx, address))

	// trustee2 approves account
	approveAddAccount := types.NewMsgApproveAddAccount(address, trustee2)
	result = setup.Handler(setup.Ctx, approveAddAccount)
	require.Equal(t, sdk.CodeOK, result.Code)

	// ensure proposed account created
	proposedAccount = setup.Keeper.GetProposedAccount(setup.Ctx, address)
	require.Equal(t, address, proposedAccount.Address)
	require.Equal(t, []sdk.AccAddress{trustee1, trustee2}, proposedAccount.Approvals)

	// ensure no active account created
	require.False(t, setup.Keeper.IsAccountPresent(setup.Ctx, address))

	// trustee3 approves account
	approveAddAccount = types.NewMsgApproveAddAccount(address, trustee3)
	result = setup.Handler(setup.Ctx, approveAddAccount)
	require.Equal(t, sdk.CodeOK, result.Code)

	// active account must be created
	account := setup.Keeper.GetAccount(setup.Ctx, address)
	require.Equal(t, address, account.Address)
	require.Equal(t, pubkey, account.PubKey)

	// ensure no pending account created
	require.False(t, setup.Keeper.IsProposedAccountPresent(setup.Ctx, address))
}

func TestHandler_ProposeAccount_ByNotTrustee(t *testing.T) {
	setup := Setup()

	for _, role := range []AccountRole{Vendor, TestHouse, ZBCertificationCenter, NodeAdmin} {
		// store signer account
		signer, _ := storeAccount(setup, role)

		// propose new account
		result, _, _ := proposeAccount(setup, signer)
		require.Equal(t, sdk.CodeUnauthorized, result.Code)
	}
}

func TestHandler_ProposeAccount_ForExistingProposedAccount(t *testing.T) {
	setup := Setup()

	// store trustee
	trustee := storeTrustee(setup)

	// propose account
	result, address, pubkey := proposeAccount(setup, trustee)
	require.Equal(t, sdk.CodeOK, result.Code)

	// propose same account second time
	proposeAddAccount := types.NewMsgProposeAddAccount(
		address,
		sdk.MustBech32ifyAccPub(pubkey),
		types.AccountRoles{types.Vendor},
		trustee,
	)
	result = setup.Handler(setup.Ctx, proposeAddAccount)
	require.Equal(t, types.CodeAccountAlreadyExist, result.Code)
}

func TestHandler_ProposeAccount_ForExistingApprovedAccount(t *testing.T) {
	setup := Setup()

	// store trustee
	trustee := storeTrustee(setup)

	// store active account
	address, pubkey := storeAccount(setup, types.Vendor)

	// propose existing account
	proposeAddAccount := types.NewMsgProposeAddAccount(
		address,
		sdk.MustBech32ifyAccPub(pubkey),
		types.AccountRoles{types.Vendor},
		trustee,
	)
	result := setup.Handler(setup.Ctx, proposeAddAccount)
	require.Equal(t, types.CodeAccountAlreadyExist, result.Code)
}

func TestHandler_ApproveAccount_ByNotTrustee(t *testing.T) {
	setup := Setup()

	// store trustee
	trustee := storeTrustee(setup)
	_ = storeTrustee(setup)
	_ = storeTrustee(setup)

	// ensure 2 trustee approvals are needed
	require.Equal(t, 2, CountAccountApprovals(setup.Ctx, setup.Keeper))

	// propose account
	result, address, _ := proposeAccount(setup, trustee)
	require.Equal(t, sdk.CodeOK, result.Code)

	for _, role := range []AccountRole{Vendor, TestHouse, ZBCertificationCenter, NodeAdmin} {
		// store signer account
		signer, _ := storeAccount(setup, role)

		// try to approve account
		approveAddAccount := types.NewMsgApproveAddAccount(address, signer)
		result := setup.Handler(setup.Ctx, approveAddAccount)
		require.Equal(t, sdk.CodeUnauthorized, result.Code)
	}
}

func TestHandler_ApproveAccount_ForUnknownAccount(t *testing.T) {
	setup := Setup()

	// store trustee
	trustee := storeTrustee(setup)

	// approve unknown account
	approveAddAccount := types.NewMsgApproveAddAccount(testconstants.Address1, trustee)
	result := setup.Handler(setup.Ctx, approveAddAccount)
	require.Equal(t, types.CodePendingAccountDoesNotExist, result.Code)
}

func TestHandler_ApproveAccount_ForDuplicateApproval(t *testing.T) {
	setup := Setup()

	// store 3 trustees
	trustee1 := storeTrustee(setup)
	_ = storeTrustee(setup)
	_ = storeTrustee(setup)

	// ensure 2 trustee approvals are needed
	require.Equal(t, 2, CountAccountApprovals(setup.Ctx, setup.Keeper))

	// proposed account
	result, address, _ := proposeAccount(setup, trustee1)
	require.Equal(t, sdk.CodeOK, result.Code)

	// the same trustee tries to approve account
	approveAddAccount := types.NewMsgApproveAddAccount(address, trustee1)
	result = setup.Handler(setup.Ctx, approveAddAccount)
	require.Equal(t, sdk.CodeUnauthorized, result.Code)
}

func storeTrustee(setup TestSetup) sdk.AccAddress {
	address, _ := storeAccount(setup, types.Trustee)
	return address
}

func storeAccount(setup TestSetup, role types.AccountRole) (sdk.AccAddress, crypto.PubKey) {
	address, pubkey, _ := testconstants.TestAddress()
	account := types.NewAccount(address, pubkey, types.AccountRoles{role})
	setup.Keeper.AssignNumberAndStoreAccount(setup.Ctx, account)
	return address, pubkey
}

func proposeAccount(setup TestSetup, signer sdk.AccAddress) (sdk.Result, sdk.AccAddress, crypto.PubKey) {
	address, pubkey, pubkeyStr := testconstants.TestAddress()
	proposeAddAccount := types.NewMsgProposeAddAccount(
		address,
		pubkeyStr,
		types.AccountRoles{types.Vendor},
		signer,
	)
	result := setup.Handler(setup.Ctx, proposeAddAccount)
	return result, address, pubkey
}
