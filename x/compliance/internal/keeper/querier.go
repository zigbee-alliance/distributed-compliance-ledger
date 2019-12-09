package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/abci/types"
)

const (
	QueryModelInfo = "model_info"
)

func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req types.RequestQuery) (res []byte, err sdk.Error) {
		switch path[0] {
		case QueryModelInfo:
			return queryModelInfo(ctx, path[1:], req, keeper)
		default:
			return nil, sdk.ErrUnknownRequest("unknown compliance query endpoint")
		}
	}
}

func queryModelInfo(ctx sdk.Context, path []string, req types.RequestQuery, keeper Keeper) ([]byte, sdk.Error) {
	modelInfo := keeper.GetModelInfo(ctx, path[0])

	res, err := codec.MarshalJSONIndent(keeper.cdc, modelInfo)
	if err != nil {
		panic("could not marshal result to JSON")
	}

	return res, nil
}
