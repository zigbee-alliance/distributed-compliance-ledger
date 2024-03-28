package keeper

import (
	"context"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	authtypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
)

func (k Keeper) Accounts(c context.Context, req *types.QueryAccountsRequest) (*types.QueryAccountsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	store := ctx.KVStore(k.storeKey)
	accountStore := prefix.NewStore(store, authtypes.KeyPrefix(authtypes.AccountKeyPrefix))

	var accounts []*codectypes.Any

	pageRes, err := query.Paginate(accountStore, req.Pagination, func(key []byte, value []byte) error {
		account := k.decodeAccount(value)
		anyWithValue, err := codectypes.NewAnyWithValue(account)
		if err != nil {
			return err
		}

		accounts = append(accounts, anyWithValue)

		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAccountsResponse{Accounts: accounts, Pagination: pageRes}, err
}

func (k Keeper) Account(c context.Context, req *types.QueryAccountRequest) (*types.QueryAccountResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	addr, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, err
	}

	account := k.GetAccount(ctx, addr)
	if account == nil {
		return nil, status.Errorf(codes.NotFound, "account %s not found", req.Address)
	}
	anyWithValue, err := codectypes.NewAnyWithValue(account)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &types.QueryAccountResponse{Account: anyWithValue}, nil
}

// Params returns parameters of auth module.
func (k Keeper) Params(c context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}
	ctx := sdk.UnwrapSDKContext(c)
	params := k.GetParams(ctx)

	return &types.QueryParamsResponse{Params: params}, nil
}

func (k Keeper) AccountAddressByID(_ context.Context, _ *types.QueryAccountAddressByIDRequest) (resp *types.QueryAccountAddressByIDResponse, e error) {
	return resp, nil
}

func (k Keeper) ModuleAccounts(_ context.Context, _ *types.QueryModuleAccountsRequest) (resp *types.QueryModuleAccountsResponse, e error) {
	return resp, nil
}

func (k Keeper) ModuleAccountByName(_ context.Context, _ *types.QueryModuleAccountByNameRequest) (resp *types.QueryModuleAccountByNameResponse, e error) {
	return resp, nil
}

func (k Keeper) Bech32Prefix(_ context.Context, _ *types.Bech32PrefixRequest) (resp *types.Bech32PrefixResponse, e error) {
	return resp, nil
}

func (k Keeper) AddressBytesToString(_ context.Context, _ *types.AddressBytesToStringRequest) (resp *types.AddressBytesToStringResponse, e error) {
	return resp, nil
}

func (k Keeper) AddressStringToBytes(_ context.Context, _ *types.AddressStringToBytesRequest) (resp *types.AddressStringToBytesResponse, e error) {
	return resp, nil
}

func (k Keeper) AccountInfo(_ context.Context, _ *types.QueryAccountInfoRequest) (resp *types.QueryAccountInfoResponse, e error) {
	return resp, nil
}
