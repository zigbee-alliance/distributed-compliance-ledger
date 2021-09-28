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

package app

import (
	"encoding/json"
	"os"

	bam "github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/version"
	authutils "github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/params"
	abci "github.com/tendermint/tendermint/abci/types"
	cmn "github.com/tendermint/tendermint/libs/common"
	"github.com/tendermint/tendermint/libs/log"
	tmtypes "github.com/tendermint/tendermint/types"
	dbm "github.com/tendermint/tm-db"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/auth"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliance"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliancetest"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/genutil"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/model"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/modelversion"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/validator"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/vendorinfo"
)

const appName = "dc-ledger"

var (
	// default home directories for the application CLI.
	DefaultCLIHome = os.ExpandEnv("$HOME/.dclcli")

	// DefaultNodeHome sets the folder where the applcation data and configuration will be stored.
	DefaultNodeHome = os.ExpandEnv("$HOME/.dcld")
)

// ModuleBasics is in charge of setting up basic module elemnets.
var ModuleBasics = module.NewBasicManager(
	auth.AppModuleBasic{},
	validator.AppModuleBasic{},
	genutil.AppModuleBasic{},
	model.AppModuleBasic{},
	modelversion.AppModuleBasic{},
	compliance.AppModuleBasic{},
	compliancetest.AppModuleBasic{},
	pki.AppModuleBasic{},
	vendorinfo.AppModuleBasic{},
)

// MakeCodec generates the necessary codecs for Amino.
func MakeCodec() *codec.Codec {
	cdc := codec.New()

	ModuleBasics.RegisterCodec(cdc)
	sdk.RegisterCodec(cdc)
	codec.RegisterCrypto(cdc)

	return cdc
}

type dcLedgerApp struct {
	*bam.BaseApp
	cdc *codec.Codec

	// keys to access the substores
	keys  map[string]*sdk.KVStoreKey
	tkeys map[string]*sdk.TransientStoreKey

	// Keepers
	authKeeper           auth.Keeper
	validatorKeeper      validator.Keeper
	modelKeeper          model.Keeper
	modelversionKeeper   modelversion.Keeper
	pkiKeeper            pki.Keeper
	complianceKeeper     compliance.Keeper
	compliancetestKeeper compliancetest.Keeper
	vendorinfoKeeper     vendorinfo.Keeper

	// Module Manager
	mm *module.Manager
}

// NewDcLedgerApp is a constructor function for dcLedgerApp.
func NewDcLedgerApp(logger log.Logger, db dbm.DB, baseAppOptions ...func(*bam.BaseApp)) *dcLedgerApp {
	// First define the top level codec that will be shared by the different modules
	cdc := MakeCodec()

	// BaseApp handles interactions with Tendermint through the ABCI protocol
	bApp := bam.NewBaseApp(appName, logger, db, authutils.DefaultTxDecoder(cdc), baseAppOptions...)

	bApp.SetAppVersion(version.Version)

	keys := sdk.NewKVStoreKeys(bam.MainStoreKey, auth.StoreKey, validator.StoreKey,
		model.StoreKey, modelversion.StoreKey, compliance.StoreKey, compliancetest.StoreKey, pki.StoreKey, vendorinfo.StoreKey)

	tkeys := sdk.NewTransientStoreKeys(params.TStoreKey)

	// Here you initialize your application with the store keys it requires.
	app := &dcLedgerApp{
		BaseApp: bApp,
		cdc:     cdc,
		keys:    keys,
		tkeys:   tkeys,
	}

	InitKeepers(app, keys)

	InitModuleManager(app)

	// The initChainer handles translating the genesis.json file into initial state for the network.
	app.SetInitChainer(app.InitChainer)
	app.SetBeginBlocker(app.BeginBlocker)
	app.SetEndBlocker(app.EndBlocker)

	// The AnteHandler handles signature verification and transaction pre-processing.
	app.SetAnteHandler(
		auth.NewAnteHandler(
			app.authKeeper,
			auth.DefaultSigVerificationGasConsumer,
		),
	)

	// initialize stores
	app.MountKVStores(keys)
	app.MountTransientStores(tkeys)

	err := app.LoadLatestVersion(app.keys[bam.MainStoreKey])
	if err != nil {
		cmn.Exit(err.Error())
	}

	return app
}

func InitModuleManager(app *dcLedgerApp) {
	app.mm = module.NewManager(
		genutil.NewAppModule(app.authKeeper, app.validatorKeeper, app.BaseApp.DeliverTx),
		auth.NewAppModule(app.authKeeper),
		validator.NewAppModule(app.validatorKeeper, app.authKeeper),
		model.NewAppModule(app.modelKeeper, app.authKeeper),
		modelversion.NewAppModule(app.modelversionKeeper, app.authKeeper, app.modelKeeper),
		compliance.NewAppModule(app.complianceKeeper, app.modelversionKeeper, app.compliancetestKeeper, app.authKeeper),
		compliancetest.NewAppModule(app.compliancetestKeeper, app.authKeeper, app.modelversionKeeper),
		pki.NewAppModule(app.pkiKeeper, app.authKeeper),
		vendorinfo.NewAppModule(app.vendorinfoKeeper, app.authKeeper),
	)

	app.mm.SetOrderBeginBlockers(validator.ModuleName)
	app.mm.SetOrderEndBlockers(validator.ModuleName)

	app.mm.SetOrderInitGenesis(
		auth.ModuleName,
		validator.ModuleName,
		model.ModuleName,
		modelversion.ModuleName,
		compliance.ModuleName,
		compliancetest.ModuleName,
		pki.ModuleName,
		vendorinfo.ModuleName,
		genutil.ModuleName,
	)

	// register all module routes and module queriers
	app.mm.RegisterRoutes(app.Router(), app.QueryRouter())
}

func InitKeepers(app *dcLedgerApp, keys map[string]*sdk.KVStoreKey) {
	// The Validator keeper
	app.validatorKeeper = MakeValidatorKeeper(keys, app)

	// The ModelKeeper keeper
	app.modelKeeper = MakeModelKeeper(keys, app)

	// The ModelversionKeeper keeper
	app.modelversionKeeper = MakeModelversionKeeper(keys, app)

	// The ComplianceKeeper keeper
	app.complianceKeeper = MakeComplianceKeeper(keys, app)

	// The CompliancetestKeeper keeper
	app.compliancetestKeeper = MakeCompliancetestKeeper(keys, app)

	// The PKI keeper
	app.pkiKeeper = MakePkiKeeper(keys, app)

	// The AuthKeeper keeper
	app.authKeeper = MakeAuthKeeper(keys, app)

	// The Vendor keeper
	app.vendorinfoKeeper = MakeVendorInfoKeeper(keys, app)
}

func MakeAuthKeeper(keys map[string]*sdk.KVStoreKey, app *dcLedgerApp) auth.Keeper {
	return auth.NewKeeper(
		keys[auth.StoreKey],
		app.cdc,
	)
}

func MakeModelKeeper(keys map[string]*sdk.KVStoreKey, app *dcLedgerApp) model.Keeper {
	return model.NewKeeper(
		keys[model.StoreKey],
		app.cdc,
	)
}

func MakeModelversionKeeper(keys map[string]*sdk.KVStoreKey, app *dcLedgerApp) modelversion.Keeper {
	return modelversion.NewKeeper(
		keys[modelversion.StoreKey],
		app.cdc,
	)
}

func MakeVendorInfoKeeper(keys map[string]*sdk.KVStoreKey, app *dcLedgerApp) vendorinfo.Keeper {
	return vendorinfo.NewKeeper(
		keys[vendorinfo.StoreKey],
		app.cdc,
	)
}

func MakeComplianceKeeper(keys map[string]*sdk.KVStoreKey, app *dcLedgerApp) compliance.Keeper {
	return compliance.NewKeeper(
		keys[compliance.StoreKey],
		app.cdc,
	)
}

func MakeCompliancetestKeeper(keys map[string]*sdk.KVStoreKey, app *dcLedgerApp) compliancetest.Keeper {
	return compliancetest.NewKeeper(
		keys[compliancetest.StoreKey],
		app.cdc,
	)
}

func MakePkiKeeper(keys map[string]*sdk.KVStoreKey, app *dcLedgerApp) pki.Keeper {
	return pki.NewKeeper(
		keys[pki.StoreKey],
		app.cdc,
	)
}

func MakeValidatorKeeper(keys map[string]*sdk.KVStoreKey, app *dcLedgerApp) validator.Keeper {
	return validator.NewKeeper(
		keys[validator.StoreKey],
		app.cdc,
	)
}

// GenesisState represents chain state at the start of the chain. Any initial state (account balances) are stored here.
type GenesisState map[string]json.RawMessage

func NewDefaultGenesisState() GenesisState {
	return ModuleBasics.DefaultGenesis()
}

func (app *dcLedgerApp) InitChainer(ctx sdk.Context, req abci.RequestInitChain) abci.ResponseInitChain {
	var genesisState GenesisState

	err := app.cdc.UnmarshalJSON(req.AppStateBytes, &genesisState)
	if err != nil {
		panic(err)
	}

	return app.mm.InitGenesis(ctx, genesisState)
}

func (app *dcLedgerApp) BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock) abci.ResponseBeginBlock {
	return app.mm.BeginBlock(ctx, req)
}

func (app *dcLedgerApp) EndBlocker(ctx sdk.Context, req abci.RequestEndBlock) abci.ResponseEndBlock {
	return app.mm.EndBlock(ctx, req)
}

func (app *dcLedgerApp) LoadHeight(height int64) error {
	return app.LoadVersion(height, app.keys[bam.MainStoreKey])
}

//_________________________________________________________

func (app *dcLedgerApp) ExportAppStateAndValidators(forZeroHeight bool, jailWhiteList []string,
) (appState json.RawMessage, validators []tmtypes.GenesisValidator, err error) {
	// as if they could withdraw from the start of the next block
	ctx := app.NewContext(true, abci.Header{Height: app.LastBlockHeight()})

	genState := app.mm.ExportGenesis(ctx)

	appState, err = codec.MarshalJSONIndent(app.cdc, genState)
	if err != nil {
		return nil, nil, err
	}

	validators = validator.WriteValidators(ctx, app.validatorKeeper)

	return appState, validators, nil
}
