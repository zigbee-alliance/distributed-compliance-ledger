package keeper_test

import (
	"testing"

	tmdb "github.com/cometbft/cometbft-db"
	"github.com/cometbft/cometbft/libs/log"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	"github.com/cosmos/cosmos-sdk/store"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkquery "github.com/cosmos/cosmos-sdk/types/query"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/stretchr/testify/require"

	basekeeper "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/base/keeper"
	dclauthtypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
)

// setupKeeperWithAuthInterfaces is like setupKeeper but ensures the interface
// registry has auth types registered so interface marshal/unmarshal works.
func setupKeeperWithAuthInterfaces(t *testing.T) (*basekeeper.Keeper, codec.BinaryCodec, sdk.Context, storetypes.StoreKey) {
	t.Helper()

	storeKey := sdk.NewKVStoreKey(dclauthtypes.StoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(dclauthtypes.MemStoreKey)

	db := tmdb.NewMemDB()
	stateStore := store.NewCommitMultiStore(db)
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(memStoreKey, storetypes.StoreTypeMemory, nil)
	require.NoError(t, stateStore.LoadLatestVersion())

	registry := codectypes.NewInterfaceRegistry()
	// Register crypto and auth interface implementations (e.g., BaseAccount).
	cryptocodec.RegisterInterfaces(registry)
	authtypes.RegisterInterfaces(registry)

	cdc := codec.NewProtoCodec(registry)
	k := basekeeper.NewKeeper(cdc, storeKey, memStoreKey)
	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, log.NewNopLogger())

	return k, cdc, ctx, storeKey
}

// putAccountAny stores an AccountI encoded as an Any/interface so that
// base keeper's Accounts query can decode via UnmarshalInterface.
func putAccountAny(t *testing.T, cdc codec.BinaryCodec, storeKey storetypes.StoreKey, ctx sdk.Context, addr sdk.AccAddress) {
	t.Helper()

	ba := authtypes.NewBaseAccountWithAddress(addr)
	// Encode as interface (wraps into Any under the hood for ProtoCodec).
	bz, err := cdc.MarshalInterface(ba)
	require.NoError(t, err)

	s := prefix.NewStore(ctx.KVStore(storeKey), dclauthtypes.KeyPrefix(dclauthtypes.AccountKeyPrefix))
	s.Set(dclauthtypes.AccountKey(addr), bz)
}

func putAccount(t *testing.T, k *basekeeper.Keeper, cdc codec.BinaryCodec, storeKey storetypes.StoreKey, ctx sdk.Context, addr sdk.AccAddress) dclauthtypes.Account {
	t.Helper()

	ba := authtypes.NewBaseAccountWithAddress(addr)
	acc := dclauthtypes.NewAccount(ba, nil, nil, nil, 0, nil)

	store := prefix.NewStore(ctx.KVStore(storeKey), dclauthtypes.KeyPrefix(dclauthtypes.AccountKeyPrefix))
	bz := cdc.MustMarshal(acc)
	store.Set(dclauthtypes.AccountKey(addr), bz)

	return *acc
}

func TestAccounts_InvalidRequest(t *testing.T) {
	k, _, ctx, _ := setupKeeperWithAuthInterfaces(t)

	_, err := k.Accounts(sdk.WrapSDKContext(ctx), nil)
	require.Error(t, err)
}

func TestAccounts_EmptyStore_ReturnsEmptyList(t *testing.T) {
	k, _, ctx, _ := setupKeeperWithAuthInterfaces(t)

	resp, err := k.Accounts(sdk.WrapSDKContext(ctx), &authtypes.QueryAccountsRequest{Pagination: nil})
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Empty(t, resp.Accounts)
	require.NotNil(t, resp.Pagination)
	require.Nil(t, resp.Pagination.NextKey)
	require.EqualValues(t, 0, resp.Pagination.Total)
}

func TestAccounts_ReturnsAll_AsAnys(t *testing.T) {
	k, cdc, ctx, storeKey := setupKeeperWithAuthInterfaces(t)

	addrs := []sdk.AccAddress{
		sdk.AccAddress("acc_list_1____________"),
		sdk.AccAddress("acc_list_2____________"),
		sdk.AccAddress("acc_list_3____________"),
	}
	for _, a := range addrs {
		putAccountAny(t, cdc, storeKey, ctx, a)
	}

	resp, err := k.Accounts(sdk.WrapSDKContext(ctx), &authtypes.QueryAccountsRequest{Pagination: &sdkquery.PageRequest{Limit: 100}})
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Len(t, resp.Accounts, len(addrs))

	// Validate the returned Any bytes decode to BaseAccount having one of the addresses.
	got := make(map[string]struct{})
	for _, anyVal := range resp.Accounts {
		var ba authtypes.BaseAccount
		require.NoError(t, cdc.Unmarshal(anyVal.Value, &ba))
		got[ba.Address] = struct{}{}
	}
	for _, a := range addrs {
		_, ok := got[a.String()]
		require.True(t, ok)
	}
}

func TestAccounts_Pagination(t *testing.T) {
	k, cdc, ctx, storeKey := setupKeeperWithAuthInterfaces(t)

	// Add three accounts
	addrs := []sdk.AccAddress{
		sdk.AccAddress("acc_pg_1_______________"),
		sdk.AccAddress("acc_pg_2_______________"),
		sdk.AccAddress("acc_pg_3_______________"),
	}
	for _, a := range addrs {
		putAccountAny(t, cdc, storeKey, ctx, a)
	}

	// Page 1: limit 2
	resp1, err := k.Accounts(sdk.WrapSDKContext(ctx), &authtypes.QueryAccountsRequest{
		Pagination: &sdkquery.PageRequest{Limit: 2},
	})
	require.NoError(t, err)
	require.NotNil(t, resp1)
	require.Len(t, resp1.Accounts, 2)
	require.NotNil(t, resp1.Pagination)
	require.NotEmpty(t, resp1.Pagination.NextKey)

	// Page 2: use next key
	resp2, err := k.Accounts(sdk.WrapSDKContext(ctx), &authtypes.QueryAccountsRequest{
		Pagination: &sdkquery.PageRequest{Key: resp1.Pagination.NextKey, Limit: 2},
	})
	require.NoError(t, err)
	require.NotNil(t, resp2)
	require.Len(t, resp2.Accounts, 1)
}

func TestAccount_InvalidRequest(t *testing.T) {
	k, _, ctx, _ := setupKeeperWithAuthInterfaces(t)

	_, err := k.Account(sdk.WrapSDKContext(ctx), nil)
	require.Error(t, err)
}

func TestAccount_InvalidAddress(t *testing.T) {
	k, _, ctx, _ := setupKeeperWithAuthInterfaces(t)

	_, err := k.Account(sdk.WrapSDKContext(ctx), &authtypes.QueryAccountRequest{Address: "not-a-bech32"})
	require.Error(t, err)
}

func TestAccount_NotFound(t *testing.T) {
	k, _, ctx, _ := setupKeeperWithAuthInterfaces(t)

	// Parse a valid address string; no record stored so it should be NotFound.
	addr := sdk.AccAddress("missing_acc_address____")
	_, err := k.Account(sdk.WrapSDKContext(ctx), &authtypes.QueryAccountRequest{Address: addr.String()})
	require.Error(t, err)
}

func TestAccount_Found(t *testing.T) {
	// Reuse helper from account_test.go to store a DCL account (so GetAccount works).
	k, cdc, ctx, storeKey := setupKeeperWithAuthInterfaces(t)
	addr := sdk.AccAddress("found_acc_address______")
	_ = putAccount(t, k, cdc, storeKey, ctx, addr) // from account_test.go

	resp, err := k.Account(sdk.WrapSDKContext(ctx), &authtypes.QueryAccountRequest{Address: addr.String()})
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, resp.Account)

	// Unpack the Any to BaseAccount to verify address
	var ba authtypes.BaseAccount
	require.NoError(t, cdc.Unmarshal(resp.Account.Value, &ba))
	require.Equal(t, addr.String(), ba.Address)
}

func TestParams_InvalidRequest(t *testing.T) {
	k, _, ctx, _ := setupKeeperWithAuthInterfaces(t)

	_, err := k.Params(sdk.WrapSDKContext(ctx), nil)
	require.Error(t, err)
}
