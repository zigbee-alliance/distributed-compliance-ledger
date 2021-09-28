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
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/vendorinfo/internal/types"
)

const (
	QueryVendor     = "vendor"
	QueryAllVendors = "all_vendors"
)

func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {
		switch path[0] {
		case QueryVendor:
			return queryVendor(ctx, path[1:], req, keeper)
		case QueryAllVendors:
			return queryAllVendors(ctx, req, keeper)

		default:
			return nil, sdk.ErrUnknownRequest("unknown vendorinfo query endpoint")
		}
	}
}

func queryVendor(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) (res []byte, err sdk.Error) {
	vid, err := conversions.ParseVID(path[0])
	if err != nil {
		return nil, err
	}

	if !keeper.IsVendorInfoPresent(ctx, vid) {
		return nil, types.ErrVendorInfoDoesNotExist(vid)
	}

	model := keeper.GetVendorInfo(ctx, vid)

	res = codec.MustMarshalJSONIndent(keeper.cdc, model)

	return res, nil
}

func queryAllVendors(ctx sdk.Context, req abci.RequestQuery, keeper Keeper) (res []byte, err sdk.Error) {
	var params pagination.PaginationParams
	if err := keeper.cdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, sdk.ErrUnknownRequest(fmt.Sprintf("failed to parse request params: %s", err))
	}

	result := types.ListVendors{
		Total:   keeper.CountTotalVendorInfos(ctx),
		Vendors: []types.VendorInfo{},
	}

	skipped := 0

	keeper.IterateVendorInfos(ctx, func(vendorInfo types.VendorInfo) (stop bool) {
		if skipped < params.Skip {
			skipped++

			return false
		}
		if len(result.Vendors) < params.Take || params.Take == 0 {

			result.Vendors = append(result.Vendors, vendorInfo)

			return false
		}

		return true
	})

	res = codec.MustMarshalJSONIndent(keeper.cdc, result)

	return res, nil
}
