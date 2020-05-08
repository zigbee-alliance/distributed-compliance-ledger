package keeper

import (
	abci "github.com/tendermint/tendermint/abci/types"

	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/validator/internal/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// query endpoints supported by the validator Querier
const (
	QueryValidators = "validators"
	QueryValidator  = "validator"
)

// creates a querier for validator module
func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {
		switch path[0] {
		case QueryValidators:
			return queryValidators(ctx, k)
		case QueryValidator:
			return queryValidator(ctx, path[1:], k)
		default:
			return nil, sdk.ErrUnknownRequest("unknown pki query endpoint")
		}
	}
}

func queryValidators(ctx sdk.Context, k Keeper) ([]byte, sdk.Error) {
	validators, total := k.GetAllValidators(ctx)

	result := types.LisValidatorItems{
		Total: total,
		Items: validators,
	}

	res := codec.MustMarshalJSONIndent(types.ModuleCdc, result)
	return res, nil
}

func queryValidator(ctx sdk.Context, path []string, k Keeper) ([]byte, sdk.Error) {
	validatorAddr, err := sdk.ValAddressFromBech32(path[0])
	if err != nil {
		return nil, sdk.ErrUnknownRequest(err.Error())
	}

	if !k.IsValidatorPresent(ctx, validatorAddr) {
		return nil, types.ErrValidatorDoesNotExist(validatorAddr)
	}

	validator := k.GetValidator(ctx, validatorAddr)

	res := codec.MustMarshalJSONIndent(types.ModuleCdc, validator)

	return res, nil
}
