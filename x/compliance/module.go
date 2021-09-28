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

package compliance

import (
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/auth"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliance/client/cli"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliance/client/rest"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliancetest"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/modelversion"
)

// type check to ensure the interface is properly implemented.
var (
	_ module.AppModule      = AppModule{}
	_ module.AppModuleBasic = AppModuleBasic{}
)

// app module Basics object.
type AppModuleBasic struct{}

func (a AppModuleBasic) Name() string {
	return ModuleName
}

func (a AppModuleBasic) RegisterCodec(cdc *codec.Codec) {
	RegisterCodec(cdc)
}

func (a AppModuleBasic) DefaultGenesis() json.RawMessage {
	return ModuleCdc.MustMarshalJSON(DefaultGenesisState())
}

func (a AppModuleBasic) ValidateGenesis(bz json.RawMessage) error {
	var data GenesisState

	err := ModuleCdc.UnmarshalJSON(bz, &data)
	if err != nil {
		return err
	}
	// Once json successfully marshalled, passes along to genesis.go
	return ValidateGenesis(data)
}

// Register rest routes.
func (AppModuleBasic) RegisterRESTRoutes(ctx context.CLIContext, rtr *mux.Router) {
	rest.RegisterRoutes(ctx, rtr, StoreKey)
}

// Get the root query command of this module.
func (AppModuleBasic) GetQueryCmd(cdc *codec.Codec) *cobra.Command {
	return cli.GetQueryCmd(StoreKey, cdc)
}

// Get the root tx command of this module.
func (AppModuleBasic) GetTxCmd(cdc *codec.Codec) *cobra.Command {
	return cli.GetTxCmd(StoreKey, cdc)
}

type AppModule struct {
	AppModuleBasic
	keeper               Keeper
	modelversionKeeper   modelversion.Keeper
	compliancetestKeeper compliancetest.Keeper
	authKeeper           auth.Keeper
}

func NewAppModule(keeper Keeper, modelversionKeeper modelversion.Keeper,
	compliancetestKeeper compliancetest.Keeper, authKeeper auth.Keeper) AppModule {
	return AppModule{
		AppModuleBasic:       AppModuleBasic{},
		keeper:               keeper,
		modelversionKeeper:   modelversionKeeper,
		compliancetestKeeper: compliancetestKeeper,
		authKeeper:           authKeeper,
	}
}

func (a AppModule) InitGenesis(ctx sdk.Context, data json.RawMessage) []abci.ValidatorUpdate {
	var genesisState GenesisState

	ModuleCdc.MustUnmarshalJSON(data, &genesisState)

	return InitGenesis(ctx, a.keeper, genesisState)
}

func (a AppModule) ExportGenesis(ctx sdk.Context) json.RawMessage {
	gs := ExportGenesis(ctx, a.keeper)

	return ModuleCdc.MustMarshalJSON(gs)
}

func (a AppModule) RegisterInvariants(sdk.InvariantRegistry) {}

func (a AppModule) Route() string {
	return RouterKey
}

func (a AppModule) NewHandler() sdk.Handler {
	return NewHandler(a.keeper, a.modelversionKeeper, a.compliancetestKeeper, a.authKeeper)
}

func (a AppModule) QuerierRoute() string {
	return RouterKey
}

func (a AppModule) NewQuerierHandler() sdk.Querier {
	return NewQuerier(a.keeper)
}

func (a AppModule) BeginBlock(sdk.Context, abci.RequestBeginBlock) {}

func (a AppModule) EndBlock(sdk.Context, abci.RequestEndBlock) []abci.ValidatorUpdate {
	return []abci.ValidatorUpdate{}
}
