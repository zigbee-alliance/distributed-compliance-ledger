package keeper

import (
	"fmt"

	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliance/internal/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

const (
	QueryModelInfo        = "model_info"
	QueryModelInfoHeaders = "model_info_headers"
)

func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {
		switch path[0] {
		case QueryModelInfo:
			return queryModelInfo(ctx, path[1:], req, keeper)
		case QueryModelInfoHeaders:
			return queryModelInfoHeaders(ctx, req, keeper)
		default:
			return nil, sdk.ErrUnknownRequest("unknown compliance query endpoint")
		}
	}
}

func queryModelInfo(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) ([]byte, sdk.Error) {
	vid, err := types.ParseVID(path[0])
	if err != nil {
		return nil, sdk.ErrInternal(fmt.Sprintf("failed to parse params: %s", err))
	}

	pid, err := types.ParsePID(path[1])
	if err != nil {
		return nil, sdk.ErrInternal(fmt.Sprintf("failed to parse params: %s", err))
	}

	if !keeper.IsModelInfoPresent(ctx, vid, pid) {
		return nil, types.ErrModelInfoDoesNotExist()
	}

	modelInfo := keeper.GetModelInfo(ctx, vid, pid)

	res, err := codec.MarshalJSONIndent(keeper.cdc, modelInfo)
	if err != nil {
		panic("could not marshal result to JSON")
	}

	return res, nil
}

func queryModelInfoHeaders(ctx sdk.Context, req abci.RequestQuery, keeper Keeper) ([]byte, sdk.Error) {
	var params types.QueryModelInfoHeadersParams
	if err := keeper.cdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, sdk.ErrInternal(fmt.Sprintf("failed to parse params: %s", err))
	}

	result := types.QueryModelInfoHeadersResult{
		Total: keeper.CountTotal(ctx),
		Items: []types.ModelInfoHeader{},
	}

	skipped := 0

	keeper.IterateModelInfos(ctx, func(modelInfo types.ModelInfo) (stop bool) {
		if skipped < params.Skip {
			skipped++
			return false
		}

		if len(result.Items) < params.Take || params.Take == 0 {
			header := types.ModelInfoHeader{
				VID:   modelInfo.VID,
				PID:   modelInfo.PID,
				Name:  modelInfo.Name,
				Owner: modelInfo.Owner,
				SKU:   modelInfo.SKU,
			}

			result.Items = append(result.Items, header)
			return false
		}

		return true
	})

	res, err := codec.MarshalJSONIndent(keeper.cdc, result)
	if err != nil {
		panic("could not marshal result to JSON")
	}

	return res, nil
}
