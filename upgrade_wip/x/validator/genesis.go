package validator

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/validator/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/validator/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the validator
	for _, elem := range genState.ValidatorList {
		k.SetValidator(ctx, elem)
	}
	// Set all the lastValidatorPower
	for _, elem := range genState.LastValidatorPowerList {
		k.SetLastValidatorPower(ctx, elem)
	}
	// Set all the validatorSigningInfo
	for _, elem := range genState.ValidatorSigningInfoList {
		k.SetValidatorSigningInfo(ctx, elem)
	}
	// Set all the validatorMissedBlockBitArray
	for _, elem := range genState.ValidatorMissedBlockBitArrayList {
		k.SetValidatorMissedBlockBitArray(ctx, elem)
	}
	// Set all the validatorOwner
	for _, elem := range genState.ValidatorOwnerList {
		k.SetValidatorOwner(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	genesis.ValidatorList = k.GetAllValidator(ctx)
	genesis.LastValidatorPowerList = k.GetAllLastValidatorPower(ctx)
	genesis.ValidatorSigningInfoList = k.GetAllValidatorSigningInfo(ctx)
	genesis.ValidatorMissedBlockBitArrayList = k.GetAllValidatorMissedBlockBitArray(ctx)
	genesis.ValidatorOwnerList = k.GetAllValidatorOwner(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
