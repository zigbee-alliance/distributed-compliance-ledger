package genutil

import (
	"encoding/json"

	"github.com/gorilla/mux"
	"github.com/spf13/cobra"

	abci "github.com/tendermint/tendermint/abci/types"

	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/genutil/types"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
)

var (
	_ module.AppModuleGenesis = AppModule{}
	_ module.AppModuleBasic   = AppModuleBasic{}
)

// app module basics object.
type AppModuleBasic struct{}

// module name.
func (AppModuleBasic) Name() string {
	return ModuleName
}

// register module codec.
func (AppModuleBasic) RegisterCodec(cdc *codec.Codec) {}

// default genesis state.
func (AppModuleBasic) DefaultGenesis() json.RawMessage {
	return ModuleCdc.MustMarshalJSON(GenesisState{})
}

// module validate genesis.
func (AppModuleBasic) ValidateGenesis(bz json.RawMessage) error {
	var data GenesisState
	err := ModuleCdc.UnmarshalJSON(bz, &data)

	if err != nil {
		return err
	}

	return types.ValidateGenesis(data)
}

// register rest routes.
func (AppModuleBasic) RegisterRESTRoutes(_ context.CLIContext, _ *mux.Router) {}

// get the root tx command of this module.
func (AppModuleBasic) GetTxCmd(_ *codec.Codec) *cobra.Command { return nil }

// get the root query command of this module.
func (AppModuleBasic) GetQueryCmd(_ *codec.Codec) *cobra.Command { return nil }

//___________________________
// app module.
type AppModule struct {
	AppModuleBasic
	authKeeper      types.AuthKeeper
	validatorKeeper types.ValidatorKeeper
	deliverTx       deliverTxfn
}

// NewAppModule creates a new AppModule object.
func NewAppModule(authKeeper types.AuthKeeper,
	validatorKeeper types.ValidatorKeeper, deliverTx deliverTxfn) module.AppModule {
	return module.NewGenesisOnlyAppModule(AppModule{
		AppModuleBasic:  AppModuleBasic{},
		authKeeper:      authKeeper,
		validatorKeeper: validatorKeeper,
		deliverTx:       deliverTx,
	})
}

// module init-genesis.
func (am AppModule) InitGenesis(ctx sdk.Context, data json.RawMessage) []abci.ValidatorUpdate {
	var genesisState GenesisState

	ModuleCdc.MustUnmarshalJSON(data, &genesisState)

	return InitGenesis(ctx, ModuleCdc, am.authKeeper, am.validatorKeeper, am.deliverTx, genesisState)
}

// module export genesis.
func (am AppModule) ExportGenesis(sdk.Context) json.RawMessage {
	return nil
}
