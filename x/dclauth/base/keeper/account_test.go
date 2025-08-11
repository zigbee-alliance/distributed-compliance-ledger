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

func putAccount(t *testing.T, cdc codec.BinaryCodec, storeKey storetypes.StoreKey, ctx sdk.Context, addr sdk.AccAddress) dclauthtypes.Account {
	t.Helper()

	ba := authtypes.NewBaseAccountWithAddress(addr)
	acc := dclauthtypes.NewAccount(ba, nil, nil, nil, 0, nil)

	store := prefix.NewStore(ctx.KVStore(storeKey), dclauthtypes.KeyPrefix(dclauthtypes.AccountKeyPrefix))
	bz := cdc.MustMarshal(acc)
	store.Set(dclauthtypes.AccountKey(addr), bz)

	return *acc
}

func TestAccount_GetAccountO(t *testing.T) {
	addr_put := sdk.AccAddress("acc_1____________")

	cases := []struct {
		name     string
		addr_get string
		found    bool
	}{
		{
			name:     "account_found",
			addr_get: "acc_1____________",
			found:    true,
		},
		{
			name:     "account_not_found",
			addr_get: "acc_2____________",
			found:    false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			k, cdc, ctx, storeKey := setupKeeperWithAuthInterfaces(t)
			addr_get := sdk.AccAddress(tc.addr_get)

			acc_put := putAccount(t, cdc, storeKey, ctx, addr_put)
			acc_get, found := k.GetAccountO(ctx, addr_get)

			require.Equal(t, found, tc.found)

			if tc.found {
				require.Equal(t, acc_get, acc_put)
			}
		})
	}
}

func TestAccount_GetAccount(t *testing.T) {
	addr_put := sdk.AccAddress("acc_1____________")

	cases := []struct {
		name     string
		addr_get string
		found    bool
	}{
		{
			name:     "account_found",
			addr_get: "acc_1____________",
			found:    true,
		},
		{
			name:     "account_not_found",
			addr_get: "acc_2____________",
			found:    false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			k, cdc, ctx, storeKey := setupKeeperWithAuthInterfaces(t)

			addr_get := sdk.AccAddress(tc.addr_get)

			acc_put := putAccount(t, cdc, storeKey, ctx, addr_put)
			get_acc := k.GetAccount(ctx, addr_get)

			if tc.found {
				require.Equal(t, get_acc, acc_put.BaseAccount)
			} else {
				require.Nil(t, get_acc)
			}
		})
	}
}

func TestAccount_GetAllAccount(t *testing.T) {
	cases := []struct {
		name      string
		addrs_put []string
	}{
		{
			name:      "accounts_exists",
			addrs_put: []string{"acc_1____________", "acc_2____________", "acc_3____________"},
		},
		{
			name:      "accounts_do_not_exist",
			addrs_put: []string{},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			k, cdc, ctx, storeKey := setupKeeperWithAuthInterfaces(t)
			var accs_put []dclauthtypes.Account

			for _, tc_addr_put := range tc.addrs_put {
				addr_put := sdk.AccAddress(tc_addr_put)
				acc_put := putAccount(t, cdc, storeKey, ctx, addr_put)
				accs_put = append(accs_put, acc_put)
			}

			accs_get := k.GetAllAccount(ctx)

			require.Equal(t, accs_put, accs_get)
		})
	}
}

func TestAccount_IterateAccounts(t *testing.T) {
	cases := []struct {
		name      string
		addrs_put []string
		stop      bool
	}{
		{
			name:      "accounts_iterate_with_stop",
			addrs_put: []string{"acc_1____________", "acc_2____________", "acc_3____________"},
			stop:      true,
		},
		{
			name:      "accounts_iterate_without_stop",
			addrs_put: []string{"acc_1____________", "acc_2____________", "acc_3____________"},
			stop:      false,
		},
		{
			name:      "accounts_not_iterate",
			addrs_put: []string{},
			stop:      false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			k, cdc, ctx, storeKey := setupKeeperWithAuthInterfaces(t)
			var callback func(acc dclauthtypes.Account) (stop bool)

			if len(tc.addrs_put) > 0 {
				var accs_put []dclauthtypes.Account

				for _, tc_addr_put := range tc.addrs_put {
					addr_put := sdk.AccAddress(tc_addr_put)
					acc_put := putAccount(t, cdc, storeKey, ctx, addr_put)
					accs_put = append(accs_put, acc_put)
				}

				callback = func(acc dclauthtypes.Account) (stop bool) {
					require.Contains(t, accs_put, acc)

					return tc.stop
				}
			} else {
				callback = func(acc dclauthtypes.Account) (stop bool) {
					require.FailNow(t, "this function should not be called")

					return tc.stop
				}
			}

			k.IterateAccounts(ctx, callback)
		})
	}
}
