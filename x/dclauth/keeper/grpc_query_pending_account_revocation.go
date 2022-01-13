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
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) PendingAccountRevocationAll(c context.Context, req *types.QueryAllPendingAccountRevocationRequest) (*types.QueryAllPendingAccountRevocationResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var pendingAccountRevocations []types.PendingAccountRevocation
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	pendingAccountRevocationStore := prefix.NewStore(store, types.KeyPrefix(types.PendingAccountRevocationKeyPrefix))

	pageRes, err := query.Paginate(pendingAccountRevocationStore, req.Pagination, func(key []byte, value []byte) error {
		var pendingAccountRevocation types.PendingAccountRevocation
		if err := k.cdc.Unmarshal(value, &pendingAccountRevocation); err != nil {
			return err
		}

		pendingAccountRevocations = append(pendingAccountRevocations, pendingAccountRevocation)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllPendingAccountRevocationResponse{PendingAccountRevocation: pendingAccountRevocations, Pagination: pageRes}, nil
}

func (k Keeper) PendingAccountRevocation(c context.Context, req *types.QueryGetPendingAccountRevocationRequest) (*types.QueryGetPendingAccountRevocationResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	addr, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, err
	}

	val, found := k.GetPendingAccountRevocation(
		ctx,
		addr,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetPendingAccountRevocationResponse{PendingAccountRevocation: val}, nil
}
