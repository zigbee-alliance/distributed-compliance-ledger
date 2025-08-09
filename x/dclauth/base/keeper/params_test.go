package keeper

import (
	"testing"

	tmdb "github.com/cometbft/cometbft-db"
	"github.com/cometbft/cometbft/libs/log"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/stretchr/testify/require"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	dclauthtypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
)

// Sets up Keeper with a params Subspace configured for auth module.
func setupKeeperWithParams(t *testing.T) (*Keeper, sdk.Context) {
	t.Helper()

	// Base keeper stores
	storeKey := sdk.NewKVStoreKey(dclauthtypes.StoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(dclauthtypes.MemStoreKey)

	// Params stores
	paramsKey := sdk.NewKVStoreKey(paramstypes.StoreKey)
	tparamsKey := sdk.NewTransientStoreKey(paramstypes.TStoreKey)

	db := tmdb.NewMemDB()
	cms := store.NewCommitMultiStore(db)
	cms.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	cms.MountStoreWithDB(memStoreKey, storetypes.StoreTypeMemory, nil)
	cms.MountStoreWithDB(paramsKey, storetypes.StoreTypeIAVL, db)
	cms.MountStoreWithDB(tparamsKey, storetypes.StoreTypeTransient, nil)
	require.NoError(t, cms.LoadLatestVersion())

	// Codecs
	reg := codectypes.NewInterfaceRegistry()
	appCodec := codec.NewProtoCodec(reg)
	legacyAmino := codec.NewLegacyAmino()

	// Params keeper and subspace for auth with its key table
	pk := paramskeeper.NewKeeper(appCodec, legacyAmino, paramsKey, tparamsKey)
	authParams := authtypes.DefaultGenesisState().Params
	authKeyTable := paramtypes.NewKeyTable().RegisterParamSet(&authParams)
	subspace := pk.Subspace(authtypes.ModuleName).WithKeyTable(authKeyTable)

	// Base keeper and context
	k := NewKeeper(appCodec, storeKey, memStoreKey)
	k.paramSubspace = subspace

	ctx := sdk.NewContext(cms, tmproto.Header{}, false, log.NewNopLogger())

	return k, ctx
}

func TestParams_SetAndGet_RoundTrip(t *testing.T) {
	k, ctx := setupKeeperWithParams(t)

	expected := authtypes.Params{
		MaxMemoCharacters:      512,
		TxSigLimit:             12,
		TxSizeCostPerByte:      7,
		SigVerifyCostED25519:   2000,
		SigVerifyCostSecp256k1: 1000,
	}

	k.SetParams(ctx, expected)
	got := k.GetParams(ctx)

	require.Equal(t, expected, got)
}

func TestParams_GRPC_Success(t *testing.T) {
	k, ctx := setupKeeperWithParams(t)

	expected := authtypes.Params{
		MaxMemoCharacters:      1024,
		TxSigLimit:             10,
		TxSizeCostPerByte:      5,
		SigVerifyCostED25519:   1900,
		SigVerifyCostSecp256k1: 900,
	}
	k.SetParams(ctx, expected)

	resp, err := k.Params(sdk.WrapSDKContext(ctx), &authtypes.QueryParamsRequest{})
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, expected, resp.Params)
}
