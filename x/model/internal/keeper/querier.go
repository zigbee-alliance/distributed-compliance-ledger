// Copyright 2020 DSR Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/conversions"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/pagination"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/model/internal/types"
)

const (
	QueryModel        = "model"
	QueryAllModels    = "all_models"
	QueryVendors      = "vendors"
	QueryVendorModels = "vendor_models"

	QueryModelVersion     = "modelVersion"
	QueryAllModelVersions = "allModelVersions"
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
		case QueryModelVersion:
			return queryModelVersion(ctx, path[1:], req, keeper)
		case QueryAllModelVersions:
			return queryModelVersions(ctx, path[1:], req, keeper)
		default:
			return nil, sdk.ErrUnknownRequest("unknown model query endpoint")
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

	if !keeper.IsModelPresent(ctx, vid, pid) {
		return nil, types.ErrModelDoesNotExist(vid, pid)
	}

	model := keeper.GetModel(ctx, vid, pid)

	res = codec.MustMarshalJSONIndent(keeper.cdc, model)

	return res, nil
}

func queryAllModels(ctx sdk.Context, req abci.RequestQuery, keeper Keeper) (res []byte, err sdk.Error) {
	var params pagination.PaginationParams
	if err := keeper.cdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, sdk.ErrUnknownRequest(fmt.Sprintf("failed to parse request params: %s", err))
	}

	result := types.ListModelItems{
		Total: keeper.CountTotalModels(ctx),
		Items: []types.ModelItem{},
	}

	skipped := 0

	keeper.IterateModels(ctx, func(model types.Model) (stop bool) {
		if skipped < params.Skip {
			skipped++

			return false
		}
		if len(result.Items) < params.Take || params.Take == 0 {
			item := types.ModelItem{
				VID:        model.VID,
				PID:        model.PID,
				Name:       model.ProductName,
				PartNumber: model.PartNumber,
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

func queryModelVersion(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) (res []byte, err sdk.Error) {
	vid, err := conversions.ParseVID(path[0])
	if err != nil {
		return nil, err
	}

	pid, err := conversions.ParsePID(path[1])
	if err != nil {
		return nil, err
	}

	softwareVersion, err := conversions.ParseUInt32FromString("softwareVersion", path[2])
	if err != nil {
		return nil, err
	}

	if !keeper.IsModelVersionPresent(ctx, vid, pid, softwareVersion) {
		return nil, types.ErrModelVersionDoesNotExist(vid, pid, softwareVersion)
	}

	modelVersion := keeper.GetModelVersion(ctx, vid, pid, softwareVersion)

	res = codec.MustMarshalJSONIndent(keeper.cdc, modelVersion)

	return res, nil
}

func queryModelVersions(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) (res []byte, err sdk.Error) {
	vid, err := conversions.ParseVID(path[0])
	if err != nil {
		return nil, err
	}

	pid, err := conversions.ParsePID(path[1])
	if err != nil {
		return nil, err
	}

	if !keeper.IsModelPresent(ctx, vid, pid) {
		return nil, types.ErrNoModelVersionsExist(vid, pid)
	}

	modelVersions := keeper.GetModelVersions(ctx, vid, pid)

	res = codec.MustMarshalJSONIndent(keeper.cdc, modelVersions)

	return res, nil
}
