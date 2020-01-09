package keeper

import (
	"github.com/askolesov/zb-ledger/x/compliance/internal/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

const (
	QueryModelInfo    = "model_info"
	QueryModelInfoIDs = "model_info_ids"
)

func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {
		switch path[0] {
		case QueryModelInfo:
			return queryModelInfo(ctx, path[1:], req, keeper)
		case QueryModelInfoIDs:
			return queryModelInfoIDs(ctx, req, keeper)
		default:
			return nil, sdk.ErrUnknownRequest("unknown compliance query endpoint")
		}
	}
}

func queryModelInfo(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) ([]byte, sdk.Error) {
	id := path[0]

	if !keeper.IsModelInfoPresent(ctx, id) {
		return nil, types.ErrModelInfoDoesNotExist(types.DefaultCodespace)
	}

	modelInfo := keeper.GetModelInfo(ctx, id)

	res, err := codec.MarshalJSONIndent(keeper.cdc, modelInfo)
	if err != nil {
		panic("could not marshal result to JSON")
	}

	return res, nil
}

func queryModelInfoIDs(ctx sdk.Context, req abci.RequestQuery, keeper Keeper) ([]byte, sdk.Error) {
	var namesList types.QueryModelInfoIDsResult

	iterator := keeper.GetModelInfoIDIterator(ctx)

	for ; iterator.Valid(); iterator.Next() {
		namesList = append(namesList, string(iterator.Key()))
	}

	res, err := codec.MarshalJSONIndent(keeper.cdc, namesList)
	if err != nil {
		panic("could not marshal result to JSON")
	}

	return res, nil
}
