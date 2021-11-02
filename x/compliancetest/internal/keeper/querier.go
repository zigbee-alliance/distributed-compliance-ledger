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
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliancetest/internal/types"
)

const (
	QueryTestingResult = "testresult"
)

func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {
		switch path[0] {
		case QueryTestingResult:
			return queryTestingResult(ctx, path[1:], keeper)
		default:
			return nil, sdk.ErrUnknownRequest("unknown compliancetest query endpoint")
		}
	}
}

func queryTestingResult(ctx sdk.Context, path []string, keeper Keeper) (res []byte, err sdk.Error) {
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

	if !keeper.IsTestingResultsPresents(ctx, vid, pid, softwareVersion) {
		return nil, types.ErrTestingResultDoesNotExist(vid, pid, softwareVersion)
	}

	testingResult := keeper.GetTestingResults(ctx, vid, pid, softwareVersion)

	res = codec.MustMarshalJSONIndent(keeper.cdc, testingResult)

	return res, nil
}
