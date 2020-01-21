package authnext

import (
	"encoding/json"

	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/authz"

	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/authnext/client/rest"

	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/authnext/client/cli"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/authnext/internal/types"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
	abci "github.com/tendermint/tendermint/abci/types"
)

// type check to ensure the interface is properly implemented
var (
	_ module.AppModule      = AppModule{}
	_ module.AppModuleBasic = AppModuleBasic{}
)

// app module Basics object
type AppModuleBasic struct{}

func (a AppModuleBasic) Name() string {
	return ModuleName
}

func (a AppModuleBasic) RegisterCodec(cdc *codec.Codec) {
	RegisterCodec(cdc)
}

func (a AppModuleBasic) DefaultGenesis() json.RawMessage {
	return nil
}

func (a AppModuleBasic) ValidateGenesis(bz json.RawMessage) error {
	return nil
}

// Register rest routes
func (AppModuleBasic) RegisterRESTRoutes(ctx context.CLIContext, rtr *mux.Router) {
	rest.RegisterRoutes(ctx, rtr, RouterKey)
}

// Get the root query command of this module
func (AppModuleBasic) GetQueryCmd(cdc *codec.Codec) *cobra.Command {
	return cli.GetQueryCmd(StoreKey, cdc)
}

// Get the root tx command of this module
func (AppModuleBasic) GetTxCmd(cdc *codec.Codec) *cobra.Command {
	return cli.GetTxCmd(StoreKey, cdc)
}

type AppModule struct {
	AppModuleBasic
	accKeeper   types.AccountKeeper
	authzKeeper authz.Keeper
	cdc         *codec.Codec
}

func NewAppModule(accKeeper types.AccountKeeper, authzKeeper authz.Keeper, cdc *codec.Codec) AppModule {
	return AppModule{AppModuleBasic: AppModuleBasic{}, accKeeper: accKeeper, authzKeeper: authzKeeper, cdc: cdc}
}

func (a AppModule) InitGenesis(ctx sdk.Context, data json.RawMessage) []abci.ValidatorUpdate {
	return abci.ValidatorUpdates{}
}

func (a AppModule) ExportGenesis(ctx sdk.Context) json.RawMessage {
	return nil
}

func (a AppModule) RegisterInvariants(sdk.InvariantRegistry) {}

func (a AppModule) Route() string {
	return RouterKey
}

func (a AppModule) NewHandler() sdk.Handler {
	return NewHandler(a.accKeeper, a.cdc)
}

func (a AppModule) QuerierRoute() string {
	return RouterKey
}

func (a AppModule) NewQuerierHandler() sdk.Querier {
	return NewQuerier(a.accKeeper, a.authzKeeper, a.cdc)
}

func (a AppModule) BeginBlock(sdk.Context, abci.RequestBeginBlock) {}

func (a AppModule) EndBlock(sdk.Context, abci.RequestEndBlock) []abci.ValidatorUpdate {
	return []abci.ValidatorUpdate{}
}
