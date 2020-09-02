//nolint:testpackage
package validator

//nolint:goimports
import (
	"testing"

	constants "git.dsr-corporation.com/zb-ledger/zb-ledger/integration_tests/constants"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/auth"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/validator/internal/keeper"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/validator/internal/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
)

func TestHandler_CreateValidator(t *testing.T) {
	setup := Setup()

	// create validator
	msgCreateValidator := types.NewMsgCreateValidator(constants.ValidatorAddress1, constants.ValidatorPubKey1,
		types.Description{Name: constants.Name}, constants.Address1)
	result := setup.Handler(setup.Ctx, msgCreateValidator)
	require.Equal(t, sdk.CodeOK, result.Code)

	events := result.Events.ToABCIEvents()
	require.Equal(t, 2, len(events))
	require.Equal(t, types.EventTypeCreateValidator, events[0].Type)
	require.Equal(t, sdk.EventTypeMessage, events[1].Type)

	// check corresponding records are created
	require.True(t, setup.ValidatorKeeper.IsValidatorPresent(setup.Ctx, msgCreateValidator.Address))

	// this record will be added in the end block handler
	require.False(t, setup.ValidatorKeeper.IsLastValidatorPowerPresent(setup.Ctx, msgCreateValidator.Address))

	// query validator
	validator, _ := queryValidator(setup, msgCreateValidator.Address)
	require.Equal(t, msgCreateValidator.Address, validator.Address)
	require.Equal(t, msgCreateValidator.PubKey, validator.PubKey)
	require.Equal(t, msgCreateValidator.Description, validator.Description)
}

func TestHandler_CreateValidator_ByNotNodeAdmin(t *testing.T) {
	setup := Setup()

	msgCreateValidator := types.NewMsgCreateValidator(constants.ValidatorAddress1, constants.ValidatorPubKey1,
		types.Description{Name: constants.Name}, constants.Address1)

	for _, role := range []auth.AccountRole{auth.TestHouse, auth.ZBCertificationCenter, auth.Vendor, auth.Trustee} {
		// create signer account
		account := auth.NewAccount(constants.Address1, constants.PubKey1, auth.AccountRoles{role})
		setup.authKeeper.SetAccount(setup.Ctx, account)

		// try to create validator
		result := setup.Handler(setup.Ctx, msgCreateValidator)
		require.Equal(t, sdk.CodeUnauthorized, result.Code)
	}
}

func TestHandler_CreateValidator_TwiceForSameValidatorAddress(t *testing.T) {
	setup := Setup()

	// create validator
	msgCreateValidator := types.NewMsgCreateValidator(constants.ValidatorAddress1, constants.ValidatorPubKey1,
		types.Description{Name: constants.Name}, constants.Address1)
	result := setup.Handler(setup.Ctx, msgCreateValidator)
	require.Equal(t, sdk.CodeOK, result.Code)

	// create validator
	account := auth.NewAccount(constants.Address2, constants.PubKey2, auth.AccountRoles{auth.NodeAdmin})
	setup.authKeeper.SetAccount(setup.Ctx, account)

	msgCreateValidator = types.NewMsgCreateValidator(constants.ValidatorAddress1, constants.ValidatorPubKey1,
		types.Description{Name: constants.Name}, constants.Address2)
	result = setup.Handler(setup.Ctx, msgCreateValidator)
	require.Equal(t, types.CodeValidatorAlreadyExist, result.Code)
}

func TestHandler_CreateValidator_TwiceForSameValidatorOwner(t *testing.T) {
	setup := Setup()

	// create validator
	msgCreateValidator := types.NewMsgCreateValidator(constants.ValidatorAddress1, constants.ValidatorPubKey1,
		types.Description{Name: constants.Name}, constants.Address1)
	result := setup.Handler(setup.Ctx, msgCreateValidator)
	require.Equal(t, sdk.CodeOK, result.Code)

	// create validator with different address
	msgCreateValidator2 := types.NewMsgCreateValidator(constants.ValidatorAddress2, constants.ValidatorPubKey2,
		types.Description{Name: constants.Name}, constants.Address1)
	result = setup.Handler(setup.Ctx, msgCreateValidator2)
	require.Equal(t, types.CodeAccountAlreadyHasNode, result.Code)
}

func queryValidator(setup TestSetup, address sdk.ConsAddress) (*types.Validator, sdk.Error) {
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
