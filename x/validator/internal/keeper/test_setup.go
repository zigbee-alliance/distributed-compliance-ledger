package keeper

import (
	"git.dsr-corporation.com/zb-ledger/zb-ledger/integration_tests/constants"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/validator/internal/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"
)

type TestSetup struct {
	Cdc             *codec.Codec
	Ctx             sdk.Context
	ValidatorKeeper Keeper
	Querier         sdk.Querier
}

func Setup() TestSetup {
	// Init Codec
	cdc := codec.New()
	sdk.RegisterCodec(cdc)

	// Init KVSore
	db := dbm.NewMemDB()
	dbStore := store.NewCommitMultiStore(db)
	validatorKey := sdk.NewKVStoreKey(types.StoreKey)
	dbStore.MountStoreWithDB(validatorKey, sdk.StoreTypeIAVL, db)
	_ = dbStore.LoadLatestVersion()

	// Init Keepers
	validatorKeeper := NewKeeper(validatorKey, cdc)

	// Init Querier
	querier := NewQuerier(validatorKeeper)

	// Create context
	ctx := sdk.NewContext(dbStore, abci.Header{ChainID: test_constants.ChainId}, false, log.NewNopLogger())

	setup := TestSetup{
		Cdc:             cdc,
		Ctx:             ctx,
		ValidatorKeeper: validatorKeeper,
		Querier:         querier,
	}
	return setup
}

func DefaultValidator() types.Validator {
	return types.NewValidator(
		test_constants.ValidatorAddress1,
		test_constants.ValidatorPubKey1,
		types.Description{Name: test_constants.Name},
		test_constants.Owner,
	)
}

func DefaultValidatorPower() types.LastValidatorPower {
	return types.NewLastValidatorPower(test_constants.ValidatorAddress1)
}

func StoreTwoValidators(setup TestSetup) (types.Validator, types.Validator) {
	validator1 := types.NewValidator(test_constants.ValidatorAddress1, test_constants.ValidatorPubKey1,
		types.Description{Name: "Validator 1"}, test_constants.Address1)
	setup.ValidatorKeeper.SetValidator(setup.Ctx, validator1)

	validator2 := types.NewValidator(test_constants.ValidatorAddress2, test_constants.ValidatorPubKey2,
		types.Description{Name: "Validator 2"}, test_constants.Address2)
	setup.ValidatorKeeper.SetValidator(setup.Ctx, validator2)

	return validator1, validator2
}
