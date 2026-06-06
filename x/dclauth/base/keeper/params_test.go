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
	authKeyTable := paramstypes.NewKeyTable().RegisterParamSet(&authParams)
	subspace := pk.Subspace(authtypes.ModuleName).WithKeyTable(authKeyTable)

	// Base keeper and context
	keeper := NewKeeper(appCodec, storeKey, memStoreKey)
	keeper.paramSubspace = subspace

	ctx := sdk.NewContext(cms, tmproto.Header{}, false, log.NewNopLogger())

	return keeper, ctx
}

func TestParams_SetAndGetParams(t *testing.T) {
	keeper, ctx := setupKeeperWithParams(t)

	expected := authtypes.Params{
		MaxMemoCharacters:      512,
		TxSigLimit:             12,
		TxSizeCostPerByte:      7,
		SigVerifyCostED25519:   2000,
		SigVerifyCostSecp256k1: 1000,
	}

	keeper.SetParams(ctx, expected)
	got := keeper.GetParams(ctx)

	require.Equal(t, expected, got)
}
