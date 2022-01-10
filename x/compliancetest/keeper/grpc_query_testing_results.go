package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliancetest/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) TestingResultsAll(c context.Context, req *types.QueryAllTestingResultsRequest) (*types.QueryAllTestingResultsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var testingResultss []types.TestingResults
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	testingResultsStore := prefix.NewStore(store, types.KeyPrefix(types.TestingResultsKeyPrefix))

	pageRes, err := query.Paginate(testingResultsStore, req.Pagination, func(key []byte, value []byte) error {
		var testingResults types.TestingResults
		if err := k.cdc.Unmarshal(value, &testingResults); err != nil {
			return err
		}

		testingResultss = append(testingResultss, testingResults)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllTestingResultsResponse{TestingResults: testingResultss, Pagination: pageRes}, nil
}

func (k Keeper) TestingResults(c context.Context, req *types.QueryGetTestingResultsRequest) (*types.QueryGetTestingResultsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetTestingResults(
		ctx,
		req.Vid,
		req.Pid,
		req.SoftwareVersion,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetTestingResultsResponse{TestingResults: &val}, nil
}
