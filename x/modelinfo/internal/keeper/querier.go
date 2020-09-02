package keeper

import (
	"fmt"

	"git.dsr-corporation.com/zb-ledger/zb-ledger/utils/conversions"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/utils/pagination"
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
			return queryVendorModels(ctx, path[1:], keeper)
		default:
			return nil, sdk.ErrUnknownRequest("unknown modelinfo query endpoint")
		}
	}
}

func queryModel(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) (res []byte, err sdk.Error) {
	vid, err := conversions.ParseVID(path[0])
	if err != nil {
		return nil, err
	}

	pid, err := conversions.ParsePID(path[1])
	if err != nil {
		return nil, err
	}

	if !keeper.IsModelInfoPresent(ctx, vid, pid) {
		return nil, types.ErrModelInfoDoesNotExist(vid, pid)
	}

	modelInfo := keeper.GetModelInfo(ctx, vid, pid)

	res = codec.MustMarshalJSONIndent(keeper.cdc, modelInfo)

	return res, nil
}

func queryAllModels(ctx sdk.Context, req abci.RequestQuery, keeper Keeper) (res []byte, err sdk.Error) {
	var params pagination.PaginationParams
	if err := keeper.cdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, sdk.ErrUnknownRequest(fmt.Sprintf("failed to parse request params: %s", err))
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

	res = codec.MustMarshalJSONIndent(keeper.cdc, result)

	return res, nil
}

func queryVendors(ctx sdk.Context, req abci.RequestQuery, keeper Keeper) (res []byte, err sdk.Error) {
	var params pagination.PaginationParams
	if err := keeper.cdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, sdk.ErrUnknownRequest(fmt.Sprintf("failed to parse request params: %s", err))
	}

	result := types.ListVendorItems{
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

	res = codec.MustMarshalJSONIndent(keeper.cdc, result)

	return res, nil
}

func queryVendorModels(ctx sdk.Context, path []string, keeper Keeper) (res []byte, err sdk.Error) {
	vid, err := conversions.ParseVID(path[0])
	if err != nil {
		return nil, err
	}

	if !keeper.IsVendorProductsPresent(ctx, vid) {
		return nil, types.ErrVendorProductsDoNotExist(vid)
	}

	vendorProducts := keeper.GetVendorProducts(ctx, vid)

	res = codec.MustMarshalJSONIndent(keeper.cdc, vendorProducts)

	return res, nil
}
