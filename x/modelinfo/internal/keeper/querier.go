package keeper

import (
	"fmt"

	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/modelinfo/internal/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

const (
	QueryModel        = "model"
	QueryAllModels    = "all_models"
	QueryVendors      = "vendors"
	QueryVendorModels = "vendor_models"
)

func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {
		switch path[0] {
		case QueryModel:
			return queryModel(ctx, path[1:], req, keeper)
		case QueryAllModels:
			return queryAllModels(ctx, req, keeper)
		case QueryVendors:
			return queryVendors(ctx, req, keeper)
		case QueryVendorModels:
			return queryVendorModels(ctx, path[1:], req, keeper)
		default:
			return nil, sdk.ErrUnknownRequest("unknown modelinfo query endpoint")
		}
	}
}

func queryModel(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) ([]byte, sdk.Error) {
	vid, err := types.ParseVID(path[0])
	if err != nil {
		return nil, sdk.ErrInternal(fmt.Sprintf("failed to parse vid: %s", err))
	}

	pid, err := types.ParsePID(path[1])
	if err != nil {
		return nil, sdk.ErrInternal(fmt.Sprintf("failed to parse pid: %s", err))
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

func queryAllModels(ctx sdk.Context, req abci.RequestQuery, keeper Keeper) ([]byte, sdk.Error) {
	var params types.PaginationParams
	if err := keeper.cdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, sdk.ErrInternal(fmt.Sprintf("failed to parse params: %s", err))
	}

	result := types.LisModelInfoItems{
		Total: keeper.CountTotalModelInfos(ctx),
		Items: []types.ModelInfoItem{},
	}

	skipped := 0

	keeper.IterateModelInfos(ctx, func(modelInfo types.ModelInfo) (stop bool) {
		if skipped < params.Skip {
			skipped++
			return false
		}
		if len(result.Items) < params.Take || params.Take == 0 {
			item := types.ModelInfoItem{
				VID:   modelInfo.VID,
				PID:   modelInfo.PID,
				Name:  modelInfo.Name,
				Owner: modelInfo.Owner,
				SKU:   modelInfo.SKU,
			}

			result.Items = append(result.Items, item)

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

func queryVendors(ctx sdk.Context, req abci.RequestQuery, keeper Keeper) ([]byte, sdk.Error) {
	var params types.PaginationParams
	if err := keeper.cdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, sdk.ErrInternal(fmt.Sprintf("failed to parse params: %s", err))
	}

	result := types.LisVendorItems{
		Total: keeper.CountTotalVendorProducts(ctx),
		Items: []types.VendorItem{},
	}

	skipped := 0

	keeper.IterateVendorProducts(ctx, func(vendorProducts types.VendorProducts) (stop bool) {
		if skipped < params.Skip {
			skipped++
			return false
		}

		if len(result.Items) < params.Take || params.Take == 0 {
			item := types.VendorItem{
				VID: vendorProducts.VID,
			}

			result.Items = append(result.Items, item)
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

func queryVendorModels(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) ([]byte, sdk.Error) {
	vid, err := types.ParseVID(path[0])
	if err != nil {
		return nil, sdk.ErrInternal(fmt.Sprintf("failed to parse vid: %s", err))
	}

	if !keeper.IsVendorProductsPresent(ctx, vid) {
		return nil, types.ErrVendorProductsDoNotExist()
	}

	var params types.PaginationParams
	if err := keeper.cdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, sdk.ErrInternal(fmt.Sprintf("failed to parse params: %s", err))
	}

	vendorProducts := keeper.GetVendorProducts(ctx, vid)
	total := len(vendorProducts.PIDs)

	result := types.LisModelInfoItems{
		Total: total,
		Items: []types.ModelInfoItem{},
	}

	count := params.Take
	if count <= 0 {
		count = total
	}

	for i := params.Skip; i < count+params.Skip; i++ {
		modelInfo := keeper.GetModelInfo(ctx, vid, vendorProducts.PIDs[i])

		header := types.ModelInfoItem{
			VID:   modelInfo.VID,
			PID:   modelInfo.PID,
			Name:  modelInfo.Name,
			Owner: modelInfo.Owner,
			SKU:   modelInfo.SKU,
		}

		result.Items = append(result.Items, header)
	}

	res, err := codec.MarshalJSONIndent(keeper.cdc, result)
	if err != nil {
		panic("could not marshal result to JSON")
	}

	return res, nil
}
