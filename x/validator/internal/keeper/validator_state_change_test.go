package keeper

import (
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/validator/internal/types"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	"testing"
	"time"
)

func TestValidatorStateChange_ApplyAndReturnValidatorSetUpdates_ForEmpty(t *testing.T) {
	setup := Setup()
	updates := setup.ValidatorKeeper.ApplyAndReturnValidatorSetUpdates(setup.Ctx)
	require.Equal(t, 0, len(updates))
}

func TestValidatorStateChange_ApplyAndReturnValidatorSetUpdates_ForAddedNewValidator(t *testing.T) {
	setup := Setup()

	// create validator
	validator := DefaultValidator()
	setup.ValidatorKeeper.SetValidator(setup.Ctx, validator)

	// ensure last validator power is zero
	require.Equal(t, types.ZeroPower, setup.ValidatorKeeper.GetLastValidatorPower(setup.Ctx, validator.OperatorAddress).Power)

	// check for updates
	updates := setup.ValidatorKeeper.ApplyAndReturnValidatorSetUpdates(setup.Ctx)
	require.Equal(t, 1, len(updates))
	require.Equal(t, validator.ABCIValidatorUpdate(), updates[0])

	// ensure last validator record is set
	require.Equal(t, types.Power, setup.ValidatorKeeper.GetLastValidatorPower(setup.Ctx, validator.OperatorAddress).Power)

	// check for updates second time
	updates = setup.ValidatorKeeper.ApplyAndReturnValidatorSetUpdates(setup.Ctx)
	require.Equal(t, 0, len(updates))
}

func TestValidatorStateChange_ApplyAndReturnValidatorSetUpdates_TwoValidators(t *testing.T) {
	setup := Setup()

	// add 2 validators
	validator1, validator2 := StoreTwoValidators(setup)

	// check for updates
	updates := setup.ValidatorKeeper.ApplyAndReturnValidatorSetUpdates(setup.Ctx)
	require.Equal(t, 2, len(updates))
	require.Equal(t, validator1.ABCIValidatorUpdate(), updates[0])
	require.Equal(t, validator2.ABCIValidatorUpdate(), updates[1])

	// ensure last validator record is set
	require.Equal(t, types.Power, setup.ValidatorKeeper.GetLastValidatorPower(setup.Ctx, validator1.OperatorAddress).Power)
	require.Equal(t, types.Power, setup.ValidatorKeeper.GetLastValidatorPower(setup.Ctx, validator2.OperatorAddress).Power)

	// check for updates second time
	updates = setup.ValidatorKeeper.ApplyAndReturnValidatorSetUpdates(setup.Ctx)
	require.Equal(t, 0, len(updates))
}

func TestValidatorStateChange_HandleDoubleSign(t *testing.T) {
	setup := Setup()

	timestamp := time.Now().UTC()

	// create validator
	validator := DefaultValidator()
	setup.ValidatorKeeper.SetValidator(setup.Ctx, validator)
	setup.ValidatorKeeper.SetValidatorByConsAddr(setup.Ctx, validator)

	// check it is not slashed
	receivedValidator := setup.ValidatorKeeper.GetValidator(setup.Ctx, validator.OperatorAddress)
	require.False(t, receivedValidator.Jailed)
	require.Equal(t, types.Power, receivedValidator.Power)

	// imitate double sign for validator
	setup.Ctx = setup.Ctx.WithBlockHeader(abci.Header{
		Time: timestamp.Add(time.Second * time.Duration(5)),
	})
	setup.ValidatorKeeper.HandleDoubleSign(setup.Ctx, validator.GetConsPubKey().Address(), 5, timestamp, types.Power)

	// check validator is slashed
	receivedValidator = setup.ValidatorKeeper.GetValidator(setup.Ctx, validator.OperatorAddress)
	require.True(t, receivedValidator.Jailed)
	require.Equal(t, types.ZeroPower, receivedValidator.Power)

	events := setup.Ctx.EventManager().Events().ToABCIEvents()
	require.Equal(t, 1, len(events))
	require.Equal(t, slashing.EventTypeSlash, events[0].Type)
}

func TestValidatorStateChange_HandleDoubleSign_ForOutdated(t *testing.T) {
	setup := Setup()

	timestamp := time.Now().UTC()

	// create validator
	validator := DefaultValidator()
	setup.ValidatorKeeper.SetValidator(setup.Ctx, validator)
	setup.ValidatorKeeper.SetValidatorByConsAddr(setup.Ctx, validator)

	// check it is not slashed
	receivedValidator := setup.ValidatorKeeper.GetValidator(setup.Ctx, validator.OperatorAddress)
	require.False(t, receivedValidator.Jailed)
	require.Equal(t, types.Power, receivedValidator.Power)

	// imitate double sign for validator with outdated timestamp
	setup.Ctx = setup.Ctx.WithBlockHeader(abci.Header{
		Time: timestamp.Add(types.MaxEvidenceAge + 2*time.Second),
	})
	setup.ValidatorKeeper.HandleDoubleSign(setup.Ctx, validator.GetConsPubKey().Address(), 5, timestamp, types.Power)

	// check validator is not slashed
	receivedValidator = setup.ValidatorKeeper.GetValidator(setup.Ctx, validator.OperatorAddress)
	require.False(t, receivedValidator.Jailed)
	require.Equal(t, types.Power, receivedValidator.Power)
}

func TestValidatorStateChange_ApplyAndReturnValidatorSetUpdates_ForJailedValidator(t *testing.T) {
	setup := Setup()

	timestamp := time.Now().UTC()

	// create validator
	validator := DefaultValidator()
	setup.ValidatorKeeper.SetValidator(setup.Ctx, validator)
	setup.ValidatorKeeper.SetValidatorByConsAddr(setup.Ctx, validator)
	setup.ValidatorKeeper.SetLastValidatorPower(setup.Ctx, types.NewLastValidatorPower(validator.OperatorAddress))

	// imitate double sign for validator
	setup.Ctx = setup.Ctx.WithBlockHeader(abci.Header{
		Time: timestamp.Add(time.Second * time.Duration(5)),
	})
	setup.ValidatorKeeper.HandleDoubleSign(setup.Ctx, validator.GetConsPubKey().Address(), 5, timestamp, types.Power)

	// check for updates
	updates := setup.ValidatorKeeper.ApplyAndReturnValidatorSetUpdates(setup.Ctx)
	require.Equal(t, 1, len(updates))
	require.Equal(t, validator.ABCIValidatorUpdateZero(), updates[0])

	// ensure last validator record is zeroed
	require.Equal(t, types.ZeroPower, setup.ValidatorKeeper.GetLastValidatorPower(setup.Ctx, validator.OperatorAddress).Power)

	// check for updates second time
	updates = setup.ValidatorKeeper.ApplyAndReturnValidatorSetUpdates(setup.Ctx)
	require.Equal(t, 0, len(updates))
}

func TestValidatorStateChange_HandleValidatorSignature_ForSignedBlock(t *testing.T) {
	setup := Setup()

	// create validator
	validator := createValidator(setup)

	for i := int64(1); i <= 10; i++ {
		// handle signed block
		setup.ValidatorKeeper.HandleValidatorSignature(setup.Ctx, validator.GetConsPubKey().Address(), types.Power, true)

		// fetch signing info
		signInfo := setup.ValidatorKeeper.GetValidatorSigningInfo(setup.Ctx, validator.GetConsAddress())
		require.Equal(t, i, signInfo.IndexOffset)
		require.Equal(t, int64(0), signInfo.MissedBlocksCounter)

		// check validator is not slashed
		receivedValidator := setup.ValidatorKeeper.GetValidator(setup.Ctx, validator.OperatorAddress)
		require.False(t, receivedValidator.Jailed)
		require.Equal(t, types.Power, receivedValidator.Power)
	}
}

func TestValidatorStateChange_HandleValidatorSignature_ForMissedBlock(t *testing.T) {
	setup := Setup()

	// create validator
	validator := createValidator(setup)

	for i := int64(1); i <= 10; i++ {
		// handle not signed block
		setup.ValidatorKeeper.HandleValidatorSignature(setup.Ctx, validator.GetConsPubKey().Address(), types.Power, false)

		// fetch signing info
		signInfo := setup.ValidatorKeeper.GetValidatorSigningInfo(setup.Ctx, validator.GetConsAddress())
		require.Equal(t, i, signInfo.IndexOffset)
		require.Equal(t, i, signInfo.MissedBlocksCounter)

		// check validator is not slashed
		receivedValidator := setup.ValidatorKeeper.GetValidator(setup.Ctx, validator.OperatorAddress)
		require.False(t, receivedValidator.Jailed)
		require.Equal(t, types.Power, receivedValidator.Power)
	}
}

func TestValidatorStateChange_HandleValidatorSignature_ForMissedToManyBlocks_ForNotPassedWindow(t *testing.T) {
	setup := Setup()

	// create validator
	validator := createValidator(setup)

	// handle MinSignedPerWindow not signed blocks for n = SignedBlocksWindow Height starts from i
	for i := int64(1); i < types.SignedBlocksWindow; i++ {
		setup.Ctx = setup.Ctx.WithBlockHeight(i)

		// handle signed block
		setup.ValidatorKeeper.HandleValidatorSignature(setup.Ctx, validator.GetConsPubKey().Address(), types.Power, false)

		// fetch signing info
		signInfo := setup.ValidatorKeeper.GetValidatorSigningInfo(setup.Ctx, validator.GetConsAddress())
		require.Equal(t, i, signInfo.IndexOffset)
		require.Equal(t, i, signInfo.MissedBlocksCounter)
	}
}

func TestValidatorStateChange_HandleValidatorSignature_ForMissedToManyBlocks_PassedWindow(t *testing.T) {
	setup := Setup()

	// create validator
	validator := createValidator(setup)

	// handle MinSignedPerWindow not signed blocks for n = MinSignedPerWindow Height starts from SignedBlocksWindow + i
	for i := int64(1); i <= types.MinSignedPerWindow; i++ {
		// set height as SignedBlocksWindow + i to imitate that node passed window
		setup.Ctx = setup.Ctx.WithBlockHeight(types.SignedBlocksWindow + i)

		// handle signed block
		setup.ValidatorKeeper.HandleValidatorSignature(setup.Ctx, validator.GetConsPubKey().Address(), types.Power, false)

		// fetch signing info
		signInfo := setup.ValidatorKeeper.GetValidatorSigningInfo(setup.Ctx, validator.GetConsAddress())
		require.Equal(t, i, signInfo.IndexOffset)
		require.Equal(t, i, signInfo.MissedBlocksCounter)
	}

	// check validator is not slashed yet
	receivedValidator := setup.ValidatorKeeper.GetValidator(setup.Ctx, validator.OperatorAddress)
	require.False(t, receivedValidator.Jailed)
	require.Equal(t, types.Power, receivedValidator.Power)

	// handle one more not signed block
	setup.ValidatorKeeper.HandleValidatorSignature(setup.Ctx, validator.GetConsPubKey().Address(), types.Power, false)

	// check validator is slashed
	receivedValidator = setup.ValidatorKeeper.GetValidator(setup.Ctx, validator.OperatorAddress)
	require.True(t, receivedValidator.Jailed)
	require.Equal(t, types.ZeroPower, receivedValidator.Power)

	// fetch signing info. it must be clear
	signInfo := setup.ValidatorKeeper.GetValidatorSigningInfo(setup.Ctx, validator.GetConsAddress())
	require.Equal(t, int64(0), signInfo.IndexOffset)
	require.Equal(t, int64(0), signInfo.MissedBlocksCounter)

	// check events
	events := setup.Ctx.EventManager().Events().ToABCIEvents()
	require.Equal(t, types.MinSignedPerWindow+2, int64(len(events)))

	for i := int64(0); i < types.MinSignedPerWindow+1; i++ {
		require.Equal(t, slashing.EventTypeLiveness, events[0].Type)
	}
	require.Equal(t, slashing.EventTypeSlash, events[types.MinSignedPerWindow+1].Type)
}

func createValidator(setup TestSetup) types.Validator {
	// create validator
	validator := DefaultValidator()
	setup.ValidatorKeeper.SetValidator(setup.Ctx, validator)
	setup.ValidatorKeeper.SetValidatorByConsAddr(setup.Ctx, validator)
	setup.ValidatorKeeper.SetValidatorSigningInfo(setup.Ctx, types.NewValidatorSigningInfo(validator.GetConsAddress(), 0))
	return validator
}
