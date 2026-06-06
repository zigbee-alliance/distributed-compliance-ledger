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
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

func (k Keeper) RevokedNocRootCertificatesAll(c context.Context, req *types.QueryAllRevokedNocRootCertificatesRequest) (*types.QueryAllRevokedNocRootCertificatesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var revokedNocRootCertificatess []types.RevokedNocRootCertificates
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	revokedNocRootCertificatesStore := prefix.NewStore(store, pkitypes.KeyPrefix(types.RevokedNocRootCertificatesKeyPrefix))

	pageRes, err := query.Paginate(revokedNocRootCertificatesStore, req.Pagination, func(key []byte, value []byte) error {
		var revokedNocRootCertificates types.RevokedNocRootCertificates
		if err := k.cdc.Unmarshal(value, &revokedNocRootCertificates); err != nil {
			return err
		}

		revokedNocRootCertificatess = append(revokedNocRootCertificatess, revokedNocRootCertificates)

		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllRevokedNocRootCertificatesResponse{RevokedNocRootCertificates: revokedNocRootCertificatess, Pagination: pageRes}, nil
}

func (k Keeper) RevokedNocRootCertificates(c context.Context, req *types.QueryGetRevokedNocRootCertificatesRequest) (*types.QueryGetRevokedNocRootCertificatesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetRevokedNocRootCertificates(
		ctx,
		req.Subject,
		req.SubjectKeyId,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetRevokedNocRootCertificatesResponse{RevokedNocRootCertificates: val}, nil
}
