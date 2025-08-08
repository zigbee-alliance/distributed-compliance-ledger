package keeper_test

import (
	"testing"

	tmdb "github.com/cometbft/cometbft-db"
	"github.com/cometbft/cometbft/libs/log"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/store"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/stretchr/testify/require"

	basekeeper "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/base/keeper"
	dclauthtypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
)

// test setup returns keeper, codec, and context
func setupKeeper(t *testing.T) (*basekeeper.Keeper, codec.BinaryCodec, sdk.Context, storetypes.StoreKey) {
	t.Helper()

	storeKey := sdk.NewKVStoreKey(dclauthtypes.StoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(dclauthtypes.MemStoreKey)

	db := tmdb.NewMemDB()
	stateStore := store.NewCommitMultiStore(db)
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(memStoreKey, storetypes.StoreTypeMemory, nil)
	require.NoError(t, stateStore.LoadLatestVersion())

	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)

	k := basekeeper.NewKeeper(cdc, storeKey, memStoreKey)
	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, log.NewNopLogger())

	return k, cdc, ctx, storeKey
}

func putAccount(t *testing.T, k *basekeeper.Keeper, cdc codec.BinaryCodec, storeKey storetypes.StoreKey, ctx sdk.Context, addr sdk.AccAddress) dclauthtypes.Account {
	t.Helper()

	ba := authtypes.NewBaseAccountWithAddress(addr)
	acc := dclauthtypes.NewAccount(ba, nil, nil, nil, 0, nil)

	// store under Account prefix
	store := prefix.NewStore(ctx.KVStore(storeKey), dclauthtypes.KeyPrefix(dclauthtypes.AccountKeyPrefix))
	bz := cdc.MustMarshal(acc)
	store.Set(dclauthtypes.AccountKey(addr), bz)

	return *acc
}

func TestGetAccountO_FoundAndNotFound(t *testing.T) {
	k, cdc, ctx, storeKey := setupKeeper(t)

	addr1 := sdk.AccAddress("address_one__________")
	addr2 := sdk.AccAddress("address_two__________")

	stored := putAccount(t, k, cdc, storeKey, ctx, addr1)

	// found case
	got, found := k.GetAccountO(ctx, addr1)
	require.True(t, found)
	require.Equal(t, stored.BaseAccount.Address, got.BaseAccount.Address)

	// not found case
	_, found = k.GetAccountO(ctx, addr2)
	require.False(t, found)
}

func TestGetAccount_ReturnsBaseAccountOrNil(t *testing.T) {
	k, cdc, ctx, storeKey := setupKeeper(t)
	addr := sdk.AccAddress("addr______________base")
	putAccount(t, k, cdc, storeKey, ctx, addr)

	accI := k.GetAccount(ctx, addr)
	require.NotNil(t, accI)
	require.Equal(t, addr.String(), accI.GetAddress().String())

	missing := sdk.AccAddress("missing______________")
	require.Nil(t, k.GetAccount(ctx, missing))
}

func TestGetAllAccount_ReturnsAll(t *testing.T) {
	k, cdc, ctx, storeKey := setupKeeper(t)

	addrs := []sdk.AccAddress{
		sdk.AccAddress("all_acc_1____________"),
		sdk.AccAddress("all_acc_2____________"),
		sdk.AccAddress("all_acc_3____________"),
	}
	for _, a := range addrs {
		putAccount(t, k, cdc, storeKey, ctx, a)
	}

	list := k.GetAllAccount(ctx)
	require.Len(t, list, len(addrs))

	// collect addresses from result
	gotSet := map[string]struct{}{}
	for _, a := range list {
		gotSet[a.BaseAccount.Address] = struct{}{}
	}
	for _, a := range addrs {
		_, ok := gotSet[a.String()]
		require.True(t, ok)
	}
}

func TestIterateAccounts_StopCondition(t *testing.T) {
	k, cdc, ctx, storeKey := setupKeeper(t)

	a1 := sdk.AccAddress("iter_acc_1____________")
	a2 := sdk.AccAddress("iter_acc_2____________")
	a3 := sdk.AccAddress("iter_acc_3____________")
	putAccount(t, k, cdc, storeKey, ctx, a1)
	putAccount(t, k, cdc, storeKey, ctx, a2)
	putAccount(t, k, cdc, storeKey, ctx, a3)

	count := 0
	k.IterateAccounts(ctx, func(acc dclauthtypes.Account) (stop bool) {
		count++
		return count == 2 // stop after visiting two
	})
	require.Equal(t, 2, count)
}

func TestGetModuleAddress_IsNil(t *testing.T) {
	k, _, _, _ := setupKeeper(t)
	require.Nil(t, k.GetModuleAddress("any"))
}
