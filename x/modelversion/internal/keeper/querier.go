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
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/conversions"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/modelversion/internal/types"
)

const (
	QueryModelVersion     = "modelVersion"
	QueryAllModelVersions = "allModelVersions"
)

func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {
		switch path[0] {
		case QueryModelVersion:
			return queryModelVersion(ctx, path[1:], req, keeper)
		case QueryAllModelVersions:
			return queryModelVersions(ctx, path[1:], req, keeper)
		default:
			return nil, sdk.ErrUnknownRequest("unknown model query endpoint")
		}
	}
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
