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
	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) RejectedCertificateAll(c context.Context, req *types.QueryAllRejectedCertificatesRequest) (*types.QueryAllRejectedCertificatesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var rejectedCertificates []types.RejectedCertificate
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	rejectedCertificateStore := prefix.NewStore(store, pkitypes.KeyPrefix(types.RejectedCertificateKeyPrefix))

	pageRes, err := query.Paginate(rejectedCertificateStore, req.Pagination, func(key []byte, value []byte) error {
		var rejectedCertificate types.RejectedCertificate
		if err := k.cdc.Unmarshal(value, &rejectedCertificate); err != nil {
			return err
		}

		rejectedCertificates = append(rejectedCertificates, rejectedCertificate)

		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllRejectedCertificatesResponse{RejectedCertificate: rejectedCertificates, Pagination: pageRes}, nil
}

func (k Keeper) RejectedCertificate(c context.Context, req *types.QueryGetRejectedCertificatesRequest) (*types.QueryGetRejectedCertificatesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.NotFound, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetRejectedCertificate(
		ctx,
		req.Subject,
		req.SubjectKeyId,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetRejectedCertificatesResponse{RejectedCertificate: val}, nil
}
