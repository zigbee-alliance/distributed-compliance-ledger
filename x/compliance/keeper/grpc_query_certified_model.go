// Copyright 2022 DSR Corporation
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
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliance/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) CertifiedModelAll(c context.Context, req *types.QueryAllCertifiedModelRequest) (*types.QueryAllCertifiedModelResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var certifiedModels []types.CertifiedModel
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	certifiedModelStore := prefix.NewStore(store, types.KeyPrefix(types.CertifiedModelKeyPrefix))

	pageRes, err := query.FilteredPaginate(
		certifiedModelStore, req.Pagination,
		func(key []byte, value []byte, accumulate bool) (bool, error) {
			var certifiedModel types.CertifiedModel
			if err := k.cdc.Unmarshal(value, &certifiedModel); err != nil {
				return false, err
			}
			if !certifiedModel.Value {
				return false, nil
			}
			if accumulate {
				certifiedModels = append(certifiedModels, certifiedModel)
			}
			return true, nil
		},
	)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllCertifiedModelResponse{CertifiedModel: certifiedModels, Pagination: pageRes}, nil
}

func (k Keeper) CertifiedModel(c context.Context, req *types.QueryGetCertifiedModelRequest) (*types.QueryGetCertifiedModelResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetCertifiedModel(
		ctx,
		req.Vid,
		req.Pid,
		req.SoftwareVersion,
		req.CertificationType,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetCertifiedModelResponse{CertifiedModel: val}, nil
}
