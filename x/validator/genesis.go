package validator

import (
	"fmt"

	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/validator/internal/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
	tmtypes "github.com/tendermint/tendermint/types"
)

type GenesisState struct {
	Validators     []Validator                           `json:"validators"`
	LastValidators []types.LastValidatorPower            `json:"last_validators"`
	SigningInfos   map[string]types.ValidatorSigningInfo `json:"signing_infos"`
	MissedBlocks   map[string][]MissedBlock              `json:"missed_blocks"`
}

type MissedBlock struct {
	Index  int64 `json:"index"`
	Missed bool  `json:"missed"`
}

func DefaultGenesisState() GenesisState {
	return GenesisState{
		Validators:     []Validator{},
		LastValidators: []types.LastValidatorPower{},
		SigningInfos:   make(map[string]types.ValidatorSigningInfo),
		MissedBlocks:   make(map[string][]MissedBlock),
	}
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
		keeper.SetValidatorOwner(ctx, validator.Owner, validator.Address)
	}

	for _, address := range data.LastValidators {
		keeper.SetLastValidatorPower(ctx, address)
	}

	for _, info := range data.SigningInfos {
		keeper.SetValidatorSigningInfo(ctx, info)
	}

	for addr, array := range data.MissedBlocks {
		address, err := sdk.ConsAddressFromBech32(addr)
		if err != nil {
			panic(err)
		}

		for _, missed := range array {
			keeper.SetValidatorMissedBlockBitArray(ctx, address, missed.Index, missed.Missed)
		}
	}

	res = keeper.ApplyAndReturnValidatorSetUpdates(ctx)

	return res
}

// ExportGenesis returns a GenesisState for a given context and keeper.
func ExportGenesis(ctx sdk.Context, keeper Keeper) GenesisState {
	validators := keeper.GetAllValidators(ctx)
	lastValidators := keeper.GetLastValidatorPowers(ctx)

	signingInfos := make(map[string]types.ValidatorSigningInfo)
	missedBlocks := make(map[string][]MissedBlock)

	keeper.IterateValidatorSigningInfos(ctx, func(info types.ValidatorSigningInfo) (stop bool) {
		bechAddr := info.Address.String()
		signingInfos[bechAddr] = info
		var localMissedBlocks []MissedBlock

		keeper.IterateValidatorMissedBlockBitArray(ctx, info.Address, func(index int64, missed bool) (stop bool) {
			localMissedBlocks = append(localMissedBlocks, MissedBlock{index, missed})

			return false
		})
		missedBlocks[bechAddr] = localMissedBlocks

		return false
	})

	return GenesisState{
		Validators:     validators,
		LastValidators: lastValidators,
		SigningInfos:   signingInfos,
		MissedBlocks:   missedBlocks,
	}
}

// ValidateGenesis validates the provided validator genesis state to ensure the
// expected invariants holds.
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
			return sdk.ErrUnknownRequest(fmt.Sprintf(
				"duplicate validator in genesis state: name %v, address %v", val.GetName(), val.Address))
		}

		addrMap[strKey] = true
	}

	return
}

// WriteValidators returns a slice of genesis validators.
func WriteValidators(ctx sdk.Context, keeper Keeper) (vals []tmtypes.GenesisValidator) {
	keeper.IterateLastValidators(ctx, func(validator types.Validator) (stop bool) {
		vals = append(vals, tmtypes.GenesisValidator{
			PubKey: validator.GetConsPubKey(),
			Power:  validator.GetPower(),
			Name:   validator.GetName(),
		})

		return false
	})

	return
}
