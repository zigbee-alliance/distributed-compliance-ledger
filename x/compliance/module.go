package compliance

import (
	"encoding/json"

	"github.com/askolesov/zb-ledger/x/compliance/client/cli"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"

	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
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

// Register rest routes
func (AppModuleBasic) RegisterRESTRoutes(ctx context.CLIContext, rtr *mux.Router) {}

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
	keeper Keeper
}

func NewAppModule(keeper Keeper) AppModule {
	return AppModule{AppModuleBasic: AppModuleBasic{}, keeper: keeper}
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

func (a AppModule) RegisterInvariants(types.InvariantRegistry) {}

func (a AppModule) Route() string {
	return RouterKey
}

func (a AppModule) NewHandler() types.Handler {
	return NewHandler(a.keeper)
}

func (a AppModule) QuerierRoute() string {
	return RouterKey
}

func (a AppModule) NewQuerierHandler() types.Querier {
	return NewQuerier(a.keeper)
}

func (a AppModule) BeginBlock(types.Context, abci.RequestBeginBlock) {}

func (a AppModule) EndBlock(types.Context, abci.RequestEndBlock) []abci.ValidatorUpdate {
	return []abci.ValidatorUpdate{}
}
