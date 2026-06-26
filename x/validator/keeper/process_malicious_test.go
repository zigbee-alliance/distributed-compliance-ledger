package keeper_test

import (
	"testing"
	"time"

	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	evidencetypes "github.com/cosmos/cosmos-sdk/x/evidence/types"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	"github.com/stretchr/testify/require"
	testkeeper "github.com/zigbee-alliance/distributed-compliance-ledger/testutil/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/validator/types"
)

func TestProcessMalicious_HandleJailUnjail(t *testing.T) {
	setup := testkeeper.Setup(t)

	validator := testkeeper.DefaultValidator()
	setup.ValidatorKeeper.SetValidator(setup.Ctx, validator)
	_ = setup.ValidatorKeeper.SetValidatorByConsAddr(setup.Ctx, validator)

	receivedValidator, _ := setup.ValidatorKeeper.GetValidator(setup.Ctx, validator.GetOwner())
	require.False(t, receivedValidator.Jailed)
	require.Equal(t, types.Power, receivedValidator.Power)

	setup.ValidatorKeeper.Jail(setup.Ctx, validator, "some reason")

	receivedValidator, _ = setup.ValidatorKeeper.GetValidator(setup.Ctx, validator.GetOwner())
	require.True(t, receivedValidator.Jailed)
	require.Equal(t, types.ZeroPower, receivedValidator.Power)
	require.Equal(t, "some reason", receivedValidator.JailedReason)

	setup.ValidatorKeeper.Unjail(setup.Ctx, validator)

	receivedValidator, _ = setup.ValidatorKeeper.GetValidator(setup.Ctx, validator.GetOwner())
	require.False(t, receivedValidator.Jailed)
	require.Equal(t, types.Power, receivedValidator.Power)
	require.Equal(t, "", receivedValidator.JailedReason)
}

func TestProcessMalicious_HandleDoubleSign(t *testing.T) {
	setup := testkeeper.Setup(t)

	validator := testkeeper.DefaultValidator()
	setup.ValidatorKeeper.SetValidator(setup.Ctx, validator)
	_ = setup.ValidatorKeeper.SetValidatorByConsAddr(setup.Ctx, validator)

	receivedValidator, _ := setup.ValidatorKeeper.GetValidator(setup.Ctx, validator.GetOwner())
	require.False(t, receivedValidator.Jailed)
	require.Equal(t, types.Power, receivedValidator.Power)

	timestamp := time.Now().UTC()

	setup.Ctx = setup.Ctx.WithBlockHeader(tmproto.Header{
		Time: timestamp.Add(time.Second * time.Duration(5)),
	})
	validatorConsAddr, _ := validator.GetConsAddr()
	evidence := evidencetypes.Equivocation{
		Height:           1,
		Time:             timestamp,
		Power:            int64(types.Power),
		ConsensusAddress: validatorConsAddr.String(),
	}

	setup.ValidatorKeeper.HandleDoubleSign(setup.Ctx, &evidence)

	receivedValidator, _ = setup.ValidatorKeeper.GetValidator(setup.Ctx, validator.GetOwner())
	require.True(t, receivedValidator.Jailed)
	require.Equal(t, types.ZeroPower, receivedValidator.Power)

	events := setup.Ctx.EventManager().Events().ToABCIEvents()
	require.Equal(t, 1, len(events))
	require.Equal(t, slashingtypes.EventTypeSlash, events[0].Type)
}

func TestProcessMalicious_HandleDoubleSign_ForOutdated(t *testing.T) {
	setup := testkeeper.Setup(t)

	validator := testkeeper.DefaultValidator()
	setup.ValidatorKeeper.SetValidator(setup.Ctx, validator)
	_ = setup.ValidatorKeeper.SetValidatorByConsAddr(setup.Ctx, validator)

	timestamp := time.Now().UTC()
	initialHeight := int64(1)

	// imitate double sign for validator with outdated timestamp AND block
	maxEvidenceAge := time.Duration(1000)
	maxNumBlocks := int64(20)
	cp := tmproto.ConsensusParams{
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

	validatorConsAddr, _ := validator.GetConsAddr()
	evidence := evidencetypes.Equivocation{
		Height:           initialHeight,
		Time:             timestamp,
		Power:            int64(types.Power),
		ConsensusAddress: validatorConsAddr.String(),
	}

	setup.ValidatorKeeper.HandleDoubleSign(setup.Ctx, &evidence)

	receivedValidator, _ := setup.ValidatorKeeper.GetValidator(setup.Ctx, validator.GetOwner())
	require.False(t, receivedValidator.Jailed)
	require.Equal(t, types.Power, receivedValidator.Power)
}

func TestProcessMalicious_HandleDoubleSign_ForNotOutdatedBlock(t *testing.T) {
	setup := testkeeper.Setup(t)

	validator := testkeeper.DefaultValidator()
	setup.ValidatorKeeper.SetValidator(setup.Ctx, validator)
	_ = setup.ValidatorKeeper.SetValidatorByConsAddr(setup.Ctx, validator)

	timestamp := time.Now().UTC()
	initialHeight := int64(1)

	// outdated timestamp but NOT block -> still slashed (both must exceed)
	maxEvidenceAge := time.Duration(1000)
	cp := tmproto.ConsensusParams{
		Evidence: &tmproto.EvidenceParams{
			MaxAgeDuration: maxEvidenceAge,
		},
	}
	setup.Ctx = setup.Ctx.WithConsensusParams(&cp)
	setup.Ctx = setup.Ctx.WithBlockHeader(tmproto.Header{
		Time: timestamp.Add(maxEvidenceAge + 2*time.Second),
	})

	validatorConsAddr, _ := validator.GetConsAddr()
	evidence := evidencetypes.Equivocation{
		Height:           initialHeight,
		Time:             timestamp,
		Power:            int64(types.Power),
		ConsensusAddress: validatorConsAddr.String(),
	}

	setup.ValidatorKeeper.HandleDoubleSign(setup.Ctx, &evidence)

	receivedValidator, _ := setup.ValidatorKeeper.GetValidator(setup.Ctx, validator.GetOwner())
	require.True(t, receivedValidator.Jailed)
	require.Equal(t, types.ZeroPower, receivedValidator.Power)
}

func TestProcessMalicious_HandleDoubleSign_ForNotOutdatedAge(t *testing.T) {
	setup := testkeeper.Setup(t)

	validator := testkeeper.DefaultValidator()
	setup.ValidatorKeeper.SetValidator(setup.Ctx, validator)
	_ = setup.ValidatorKeeper.SetValidatorByConsAddr(setup.Ctx, validator)

	timestamp := time.Now().UTC()
	initialHeight := int64(1)

	// outdated block but NOT timestamp -> still slashed (both must exceed)
	maxNumBlocks := int64(20)
	cp := tmproto.ConsensusParams{
		Evidence: &tmproto.EvidenceParams{
			MaxAgeNumBlocks: maxNumBlocks,
		},
	}
	setup.Ctx = setup.Ctx.WithConsensusParams(&cp)
	setup.Ctx = setup.Ctx.WithBlockHeader(tmproto.Header{
		Height: maxNumBlocks + initialHeight + 1,
	})

	validatorConsAddr, _ := validator.GetConsAddr()
	evidence := evidencetypes.Equivocation{
		Height:           initialHeight,
		Time:             timestamp,
		Power:            int64(types.Power),
		ConsensusAddress: validatorConsAddr.String(),
	}

	setup.ValidatorKeeper.HandleDoubleSign(setup.Ctx, &evidence)

	receivedValidator, _ := setup.ValidatorKeeper.GetValidator(setup.Ctx, validator.GetOwner())
	require.True(t, receivedValidator.Jailed)
	require.Equal(t, types.ZeroPower, receivedValidator.Power)
}
