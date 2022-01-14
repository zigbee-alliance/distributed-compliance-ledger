package keeper_test

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	evidencetypes "github.com/cosmos/cosmos-sdk/x/evidence/types"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	testkeeper "github.com/zigbee-alliance/distributed-compliance-ledger/testutil/keeper"
	dclauthkeeper "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/validator/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/validator/types"
)

type TestSetup struct {
	Ctx             sdk.Context
	ValidatorKeeper keeper.Keeper
	DclauthKeeper   dclauthkeeper.Keeper
	Validator       types.Validator
}

func DefaultValidator() types.Validator {
	v, _ := types.NewValidator(
		sdk.ValAddress(testconstants.Address1),
		testconstants.PubKey1,
		types.Description{Moniker: testconstants.ProductName},
	)
	return v
}

func Setup(t *testing.T) TestSetup {
	dclauthK, _ := testkeeper.DclauthKeeper(t)
	k, ctx := testkeeper.ValidatorKeeper(t, dclauthK)

	// create validator
	validator := DefaultValidator()

	k.SetValidator(ctx, validator)
	k.SetValidatorByConsAddr(ctx, validator)

	// check it is not slashed
	receivedValidator, _ := k.GetValidator(ctx, validator.GetOwner())
	require.False(t, receivedValidator.Jailed)
	require.Equal(t, types.Power, receivedValidator.Power)

	setup := TestSetup{
		Ctx:             ctx,
		ValidatorKeeper: *k,
		DclauthKeeper:   *dclauthK,
		Validator:       validator,
	}

	return setup
}

func TestValidatorStateChange_HandleJailUnjail(t *testing.T) {
	setup := Setup(t)

	// Jail/Slash
	setup.ValidatorKeeper.Jail(setup.Ctx, setup.Validator, "some reason")

	// check validator is slashed
	receivedValidator, _ := setup.ValidatorKeeper.GetValidator(setup.Ctx, setup.Validator.GetOwner())
	require.True(t, receivedValidator.Jailed)
	require.Equal(t, types.ZeroPower, receivedValidator.Power)
	require.Equal(t, "some reason", receivedValidator.JailedReason)

	// Unjail/unslash
	setup.ValidatorKeeper.Unjail(setup.Ctx, setup.Validator)

	// check validator is not slashed
	receivedValidator, _ = setup.ValidatorKeeper.GetValidator(setup.Ctx, setup.Validator.GetOwner())
	require.False(t, receivedValidator.Jailed)
	require.Equal(t, types.Power, receivedValidator.Power)
	require.Equal(t, "", receivedValidator.JailedReason)
}

func TestValidatorStateChange_HandleDoubleSign(t *testing.T) {
	setup := Setup(t)

	timestamp := time.Now().UTC()

	// imitate double sign for validator
	setup.Ctx = setup.Ctx.WithBlockHeader(tmproto.Header{
		Time: timestamp.Add(time.Second * time.Duration(5)),
	})
	validatorConsAddr, _ := setup.Validator.GetConsAddr()
	evidence := evidencetypes.Equivocation{
		Height:           1,
		Time:             timestamp,
		Power:            int64(types.Power),
		ConsensusAddress: validatorConsAddr.String(),
	}

	setup.ValidatorKeeper.HandleDoubleSign(setup.Ctx, &evidence)

	// check validator is slashed
	receivedValidator, _ := setup.ValidatorKeeper.GetValidator(setup.Ctx, setup.Validator.GetOwner())
	require.True(t, receivedValidator.Jailed)
	require.Equal(t, types.ZeroPower, receivedValidator.Power)

	events := setup.Ctx.EventManager().Events().ToABCIEvents()
	require.Equal(t, 1, len(events))
	require.Equal(t, slashingtypes.EventTypeSlash, events[0].Type)
}

func TestValidatorStateChange_HandleDoubleSign_ForOutdated(t *testing.T) {
	setup := Setup(t)

	timestamp := time.Now().UTC()
	initialHeight := int64(1)

	// imitate double sign for validator with outdated timestamp AND block
	maxEvidenceAge := time.Duration(1000)
	maxNumBlocks := int64(20)
	cp := abci.ConsensusParams{
		Evidence: &tmproto.EvidenceParams{
			MaxAgeDuration:  maxEvidenceAge,
			MaxAgeNumBlocks: maxNumBlocks,
		},
	}
	setup.Ctx = setup.Ctx.WithConsensusParams(&cp)
	setup.Ctx = setup.Ctx.WithBlockHeader(tmproto.Header{
		Time:   timestamp.Add(maxEvidenceAge + 2*time.Second),
		Height: maxNumBlocks + initialHeight + 1,
	})

	validatorConsAddr, _ := setup.Validator.GetConsAddr()
	evidence := evidencetypes.Equivocation{
		Height:           initialHeight,
		Time:             timestamp,
		Power:            int64(types.Power),
		ConsensusAddress: validatorConsAddr.String(),
	}

	setup.ValidatorKeeper.HandleDoubleSign(setup.Ctx, &evidence)

	// check validator is not slashed
	receivedValidator, _ := setup.ValidatorKeeper.GetValidator(setup.Ctx, setup.Validator.GetOwner())
	require.False(t, receivedValidator.Jailed)
	require.Equal(t, types.Power, receivedValidator.Power)
}

func TestValidatorStateChange_HandleDoubleSign_ForNotOutdatedBlock(t *testing.T) {
	setup := Setup(t)

	timestamp := time.Now().UTC()
	initialHeight := int64(1)

	// imitate double sign for validator with outdated timestamp and not block
	maxEvidenceAge := time.Duration(1000)
	cp := abci.ConsensusParams{
		Evidence: &tmproto.EvidenceParams{
			MaxAgeDuration: maxEvidenceAge,
		},
	}
	setup.Ctx = setup.Ctx.WithConsensusParams(&cp)
	setup.Ctx = setup.Ctx.WithBlockHeader(tmproto.Header{
		Time: timestamp.Add(maxEvidenceAge + 2*time.Second),
	})

	validatorConsAddr, _ := setup.Validator.GetConsAddr()
	evidence := evidencetypes.Equivocation{
		Height:           initialHeight,
		Time:             timestamp,
		Power:            int64(types.Power),
		ConsensusAddress: validatorConsAddr.String(),
	}

	setup.ValidatorKeeper.HandleDoubleSign(setup.Ctx, &evidence)

	// check validator is slashed
	receivedValidator, _ := setup.ValidatorKeeper.GetValidator(setup.Ctx, setup.Validator.GetOwner())
	require.True(t, receivedValidator.Jailed)
	require.Equal(t, types.ZeroPower, receivedValidator.Power)
}

func TestValidatorStateChange_HandleDoubleSign_ForNotOutdatedAge(t *testing.T) {
	setup := Setup(t)

	timestamp := time.Now().UTC()
	initialHeight := int64(1)

	// imitate double sign for validator with outdated block and not timestamp
	maxNumBlocks := int64(20)
	cp := abci.ConsensusParams{
		Evidence: &tmproto.EvidenceParams{
			MaxAgeNumBlocks: maxNumBlocks,
		},
	}
	setup.Ctx = setup.Ctx.WithConsensusParams(&cp)
	setup.Ctx = setup.Ctx.WithBlockHeader(tmproto.Header{
		Height: maxNumBlocks + initialHeight + 1,
	})

	validatorConsAddr, _ := setup.Validator.GetConsAddr()
	evidence := evidencetypes.Equivocation{
		Height:           initialHeight,
		Time:             timestamp,
		Power:            int64(types.Power),
		ConsensusAddress: validatorConsAddr.String(),
	}

	setup.ValidatorKeeper.HandleDoubleSign(setup.Ctx, &evidence)

	// check validator is slashed
	receivedValidator, _ := setup.ValidatorKeeper.GetValidator(setup.Ctx, setup.Validator.GetOwner())
	require.True(t, receivedValidator.Jailed)
	require.Equal(t, types.ZeroPower, receivedValidator.Power)
}
