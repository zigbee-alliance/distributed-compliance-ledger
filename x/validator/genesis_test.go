package validator_test

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	testkeeper "github.com/zigbee-alliance/distributed-compliance-ledger/testutil/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/testutil/sample"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/validator"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/validator/types"
)

func TestGenesis(t *testing.T) {
	setup := testkeeper.Setup(t)

	pk1 := ed25519.GenPrivKey().PubKey()
	pk2 := ed25519.GenPrivKey().PubKey()
	v1, err := types.NewValidator(sdk.ValAddress(pk1.Address()), pk1, types.NewDescription("v1", "", "", ""))
	require.NoError(t, err)
	v2, err := types.NewValidator(sdk.ValAddress(pk2.Address()), pk2, types.NewDescription("v2", "", "", ""))
	require.NoError(t, err)

	genesisState := types.GenesisState{
		ValidatorList: []types.Validator{v1, v2},
		LastValidatorPowerList: []types.LastValidatorPower{
			types.NewLastValidatorPower(v1.GetOwner()),
			types.NewLastValidatorPower(v2.GetOwner()),
		},
		ProposedDisableValidatorList: []types.ProposedDisableValidator{
			{Address: "a"},
			{Address: "b"},
		},
		DisabledValidatorList: []types.DisabledValidator{
			{Address: "c"},
			{Address: "d"},
		},
		RejectedValidatorList: []types.RejectedDisableValidator{
			{Address: sample.ValAddress()},
		},
	}

	validator.InitGenesis(setup.Ctx, setup.ValidatorKeeper, genesisState)

	got := validator.ExportGenesis(setup.Ctx, setup.ValidatorKeeper)
	require.NotNil(t, got)
	require.Len(t, got.ValidatorList, 2)
	require.Len(t, got.LastValidatorPowerList, 2)
	require.Len(t, got.ProposedDisableValidatorList, 2)
	require.Len(t, got.DisabledValidatorList, 2)
	require.Len(t, got.RejectedValidatorList, 1)

	// WriteValidators iterates the active set; covered regardless of result.
	_, err = validator.WriteValidators(setup.Ctx, setup.ValidatorKeeper)
	require.NoError(t, err)
}
