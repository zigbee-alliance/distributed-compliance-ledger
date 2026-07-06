package keeper_test

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	dclauthtypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
)

// putInterfaceAccount stores an interface-encoded (Any-wrapped) BaseAccount,
// which is the encoding the cosmos-compatible Accounts list query decodes with
// (k.decodeAccount -> UnmarshalInterface).
func putInterfaceAccount(t *testing.T, cdc codec.BinaryCodec, storeKey storetypes.StoreKey, ctx sdk.Context, addr sdk.AccAddress) {
	t.Helper()

	ba := authtypes.NewBaseAccountWithAddress(addr)
	bz, err := cdc.MarshalInterface(ba)
	require.NoError(t, err)

	store := prefix.NewStore(ctx.KVStore(storeKey), dclauthtypes.KeyPrefix(dclauthtypes.AccountKeyPrefix))
	store.Set(dclauthtypes.AccountKey(addr), bz)
}

func TestAccountsQuery(t *testing.T) {
	k, cdc, ctx, storeKey := setupKeeperWithAuthInterfaces(t)
	wctx := sdk.WrapSDKContext(ctx)

	addrs := []string{"acc_1____________", "acc_2____________", "acc_3____________"}
	for _, a := range addrs {
		putInterfaceAccount(t, cdc, storeKey, ctx, sdk.AccAddress(a))
	}

	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := k.Accounts(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})

	t.Run("All", func(t *testing.T) {
		resp, err := k.Accounts(wctx, &authtypes.QueryAccountsRequest{
			Pagination: &query.PageRequest{CountTotal: true},
		})
		require.NoError(t, err)
		require.Len(t, resp.Accounts, len(addrs))
		require.Equal(t, uint64(len(addrs)), resp.Pagination.Total)
	})

	t.Run("Paginated", func(t *testing.T) {
		resp, err := k.Accounts(wctx, &authtypes.QueryAccountsRequest{
			Pagination: &query.PageRequest{Limit: 1, CountTotal: true},
		})
		require.NoError(t, err)
		require.Len(t, resp.Accounts, 1)
		require.Equal(t, uint64(len(addrs)), resp.Pagination.Total)
	})
}

func TestAccountQuery(t *testing.T) {
	k, cdc, ctx, storeKey := setupKeeperWithAuthInterfaces(t)
	wctx := sdk.WrapSDKContext(ctx)

	addr := sdk.AccAddress("acc_1____________")
	acc := putAccount(t, cdc, storeKey, ctx, addr)

	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := k.Account(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})

	t.Run("Found", func(t *testing.T) {
		resp, err := k.Account(wctx, &authtypes.QueryAccountRequest{Address: addr.String()})
		require.NoError(t, err)
		require.NotNil(t, resp.Account)
		require.Equal(t, acc.BaseAccount, resp.Account.GetCachedValue())
	})

	t.Run("InvalidAddress", func(t *testing.T) {
		_, err := k.Account(wctx, &authtypes.QueryAccountRequest{Address: "not-a-bech32-address"})
		require.Error(t, err)
	})

	t.Run("NotFound", func(t *testing.T) {
		missing := sdk.AccAddress("acc_missing______")
		_, err := k.Account(wctx, &authtypes.QueryAccountRequest{Address: missing.String()})
		require.Error(t, err)
		require.Equal(t, codes.NotFound, status.Code(err))
	})
}

func TestParamsQuery_InvalidRequest(t *testing.T) {
	k, _, ctx, _ := setupKeeperWithAuthInterfaces(t)
	wctx := sdk.WrapSDKContext(ctx)

	_, err := k.Params(wctx, nil)
	require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "empty request"))
}

// The remaining query methods are unimplemented passthrough stubs required to
// satisfy the cosmos auth QueryServer interface; exercise them so the wiring is
// covered and stays nil-safe.
func TestAccountQueryStubs(t *testing.T) {
	k, _, ctx, _ := setupKeeperWithAuthInterfaces(t)
	wctx := sdk.WrapSDKContext(ctx)

	_, err := k.AccountAddressByID(wctx, nil)
	require.NoError(t, err)
	_, err = k.ModuleAccounts(wctx, nil)
	require.NoError(t, err)
	_, err = k.ModuleAccountByName(wctx, nil)
	require.NoError(t, err)
	_, err = k.Bech32Prefix(wctx, nil)
	require.NoError(t, err)
	_, err = k.AddressBytesToString(wctx, nil)
	require.NoError(t, err)
	_, err = k.AddressStringToBytes(wctx, nil)
	require.NoError(t, err)
	_, err = k.AccountInfo(wctx, nil)
	require.NoError(t, err)
}
