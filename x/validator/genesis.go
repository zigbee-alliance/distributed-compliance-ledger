package validator

import (
	"fmt"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/validator/internal/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
	tmtypes "github.com/tendermint/tendermint/types"
)

type GenesisState struct {
	Validators          types.Validators           `json:"validators"`
	LastValidatorPowers []types.LastValidatorPower `json:"last_validator_powers"`
	Exported            bool                       `json:"exported"`
}

func NewGenesisState(validators []Validator) GenesisState {
	return GenesisState{
		Validators: validators,
	}
}

func DefaultGenesisState() GenesisState {
	return GenesisState{}
}

func InitGenesis(ctx sdk.Context, keeper Keeper, data GenesisState) (res []abci.ValidatorUpdate) {

	// We need to pretend to be "n blocks before genesis", where "n" is the
	// validator update delay, so that e.g. slashing periods are correctly
	// initialized for the validator set e.g. with a one-block offset - the
	// first TM block is at height 1, so state updates applied from
	// genesis.json are in block 0.
	ctx = ctx.WithBlockHeight(1 - sdk.ValidatorUpdateDelay)

	for _, validator := range data.Validators {
		keeper.SetValidator(ctx, validator)
		keeper.SetValidatorByConsAddr(ctx, validator)
	}

	for _, lv := range data.LastValidatorPowers {
		keeper.SetLastValidatorPower(ctx, lv.Address, lv.Power)
	}

	res = keeper.ApplyAndReturnValidatorSetUpdates(ctx)

	return res
}

// ExportGenesis returns a GenesisState for a given context and keeper. The
// GenesisState will contain the pool, params, validators, and bonds found in
// the keeper.
func ExportGenesis(ctx sdk.Context, keeper Keeper) GenesisState {
	validators := keeper.GetAllValidators(ctx)
	lastValidatorPowers := keeper.GetLastValidatorPowers(ctx)

	return GenesisState{
		Validators:          validators,
		LastValidatorPowers: lastValidatorPowers,
		Exported:            true,
	}
}

// ValidateGenesis validates the provided validator genesis state to ensure the
// expected invariants holds. (i.e. params in correct bounds, no duplicate validators)
func ValidateGenesis(data GenesisState) error {
	err := validateGenesisStateValidators(data.Validators)
	if err != nil {
		return err
	}

	return nil
}

func validateGenesisStateValidators(validators []Validator) (err error) {
	addrMap := make(map[string]bool, len(validators))
	for i := 0; i < len(validators); i++ {
		val := validators[i]
		strKey := string(val.GetConsPubKey().Bytes())
		if _, ok := addrMap[strKey]; ok {
			return fmt.Errorf("duplicate validator in genesis state: moniker %v, address %v", val.Description.Moniker, val.GetConsAddress())
		}
		addrMap[strKey] = true
	}
	return
}

// WriteValidators returns a slice of bonded genesis validators.
func WriteValidators(ctx sdk.Context, keeper Keeper) (vals []tmtypes.GenesisValidator) {
	keeper.IterateLastValidators(ctx, func(_ int64, validator types.Validator) (stop bool) {
		vals = append(vals, tmtypes.GenesisValidator{
			PubKey: validator.GetConsPubKey(),
			Power:  validator.GetWeight(),
			Name:   validator.GetMoniker(),
		})

		return false
	})

	return
}
