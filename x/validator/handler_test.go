package validator

import (
	constants "git.dsr-corporation.com/zb-ledger/zb-ledger/integration_tests/constants"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/authz"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/validator/internal/keeper"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/validator/internal/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	"testing"
)

func TestHandler_CreateValidator(t *testing.T) {
	setup := Setup()

	// create validator
	msgCreateValidator := types.NewMsgCreateValidator(constants.ValAddress1, constants.ConsensusPubKey1, types.Description{Name: constants.Name})
	result := setup.Handler(setup.Ctx, msgCreateValidator)
	require.Equal(t, sdk.CodeOK, result.Code)

	events := result.Events.ToABCIEvents()
	require.Equal(t, 2, len(events))
	require.Equal(t, types.EventTypeCreateValidator, events[0].Type)
	require.Equal(t, sdk.EventTypeMessage, events[1].Type)

	// check corresponding records are created
	require.True(t, setup.ValidatorKeeper.IsValidatorPresent(setup.Ctx, msgCreateValidator.ValidatorAddress))
	require.True(t, setup.ValidatorKeeper.IsValidatorByConsAddrPresent(setup.Ctx, sdk.ConsAddress(msgCreateValidator.GetPubKey().Address())))

	// this record will be added in the end block handler
	require.False(t, setup.ValidatorKeeper.IsLastValidatorPowerPresent(setup.Ctx, msgCreateValidator.ValidatorAddress))

	// query validator
	validator, _ := queryValidator(setup, msgCreateValidator.ValidatorAddress)
	require.Equal(t, msgCreateValidator.ValidatorAddress, validator.OperatorAddress)
	require.Equal(t, msgCreateValidator.PubKey, validator.ConsensusPubKey)
	require.Equal(t, msgCreateValidator.Description, validator.Description)
}

func TestHandler_CreateValidator_ByNotNodeAdmin(t *testing.T) {
	setup := Setup()

	msgCreateValidator := types.NewMsgCreateValidator(constants.ValAddress2, constants.ConsensusPubKey1, types.Description{Name: constants.Name})

	for _, role := range []authz.AccountRole{authz.Administrator, authz.TestHouse, authz.ZBCertificationCenter, authz.Vendor, authz.Trustee} {
		// assign role
		setup.AuthzKeeper.AssignRole(setup.Ctx, sdk.AccAddress(constants.ValAddress2), role)

		// try to create validator
		result := setup.Handler(setup.Ctx, msgCreateValidator)
		require.Equal(t, sdk.CodeUnauthorized, result.Code)
	}
}

func TestHandler_CreateValidator_Twice(t *testing.T) {
	setup := Setup()

	// create validator
	msgCreateValidator := types.NewMsgCreateValidator(constants.ValAddress1, constants.ConsensusPubKey1, types.Description{Name: constants.Name})
	result := setup.Handler(setup.Ctx, msgCreateValidator)
	require.Equal(t, sdk.CodeOK, result.Code)

	// create validator
	result = setup.Handler(setup.Ctx, msgCreateValidator)
	require.Equal(t, types.CodeValidatorOperatorAddressExist, result.Code)
}

func queryValidator(setup TestSetup, address sdk.ValAddress) (*types.Validator, sdk.Error) {
	// query validator
	result, err := setup.Querier(
		setup.Ctx,
		[]string{keeper.QueryValidator, address.String()},
		abci.RequestQuery{},
	)

	if err != nil {
		return nil, err
	}

	var validator types.Validator
	setup.Cdc.MustUnmarshalJSON(result, &validator)

	return &validator, nil
}
