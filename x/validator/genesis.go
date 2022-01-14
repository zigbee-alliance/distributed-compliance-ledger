package validator

import (
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
	tmtypes "github.com/tendermint/tendermint/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/validator/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/validator/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) (res []abci.ValidatorUpdate) {
	// Set all the validator
	for _, elem := range genState.ValidatorList {
		k.SetValidator(ctx, elem)
		_ = k.SetValidatorByConsAddr(ctx, elem)
	}
	// Set all the lastValidatorPower
	for _, elem := range genState.LastValidatorPowerList {
		k.SetLastValidatorPower(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init

	// TODO issue 99 error processing
	res, _ = k.ApplyAndReturnValidatorSetUpdates(ctx)

	return res
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	genesis.ValidatorList = k.GetAllValidator(ctx)
	genesis.LastValidatorPowerList = k.GetAllLastValidatorPower(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}

// WriteValidators returns a slice of bonded genesis validators.
func WriteValidators(ctx sdk.Context, k keeper.Keeper) (vals []tmtypes.GenesisValidator, err error) {
	k.IterateLastValidators(ctx, func(validator types.Validator) (stop bool) {
		pk, err := validator.GetConsPubKey()
		if err != nil {
			return true
		}
		tmPk, err := cryptocodec.ToTmPubKeyInterface(pk)
		if err != nil {
			return true
		}

		vals = append(vals, tmtypes.GenesisValidator{
			Address: sdk.ConsAddress(tmPk.Address()).Bytes(),
			PubKey:  tmPk,
			Power:   int64(validator.GetPower()),
			Name:    validator.GetMoniker(),
		})

		return false
	})

	return
}
