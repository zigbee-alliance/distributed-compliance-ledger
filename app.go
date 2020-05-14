package app

//nolint:goimports
import (
	"encoding/json"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliance"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/pki"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/validator"
	authutils "github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/params"
	"os"

	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/auth"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliancetest"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/genaccounts"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/genutil"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/modelinfo"
	abci "github.com/tendermint/tendermint/abci/types"
	cmn "github.com/tendermint/tendermint/libs/common"
	"github.com/tendermint/tendermint/libs/log"
	tmtypes "github.com/tendermint/tendermint/types"
	dbm "github.com/tendermint/tm-db"

	bam "github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/version"
)

const appName = "zb-ledger"

var (
	// default home directories for the application CLI
	DefaultCLIHome = os.ExpandEnv("$HOME/.zblcli")

	// DefaultNodeHome sets the folder where the applcation data and configuration will be stored
	DefaultNodeHome = os.ExpandEnv("$HOME/.zbld")
)

var (
	// NewBasicManager is in charge of setting up basic module elemnets
	ModuleBasics = module.NewBasicManager(
		genaccounts.AppModuleBasic{},
		auth.AppModuleBasic{},
		validator.AppModuleBasic{},
		genutil.AppModuleBasic{},
		modelinfo.AppModuleBasic{},
		compliance.AppModuleBasic{},
		compliancetest.AppModuleBasic{},
		pki.AppModuleBasic{},
	)
)

// MakeCodec generates the necessary codecs for Amino.
func MakeCodec() *codec.Codec {
	var cdc = codec.New()

	ModuleBasics.RegisterCodec(cdc)
	sdk.RegisterCodec(cdc)
	codec.RegisterCrypto(cdc)

	return cdc
}

type zbLedgerApp struct {
	*bam.BaseApp
	cdc *codec.Codec

	// keys to access the substores
	keys  map[string]*sdk.KVStoreKey
	tkeys map[string]*sdk.TransientStoreKey

	// Keepers
	authKeeper           auth.Keeper
	validatorKeeper      validator.Keeper
	modelinfoKeeper      modelinfo.Keeper
	pkiKeeper            pki.Keeper
	complianceKeeper     compliance.Keeper
	compliancetestKeeper compliancetest.Keeper

	// Module Manager
	mm *module.Manager
}

// NewZbLedgerApp is a constructor function for zbLedgerApp.
func NewZbLedgerApp(logger log.Logger, db dbm.DB, baseAppOptions ...func(*bam.BaseApp)) *zbLedgerApp {
	// First define the top level codec that will be shared by the different modules
	cdc := MakeCodec()

	// BaseApp handles interactions with Tendermint through the ABCI protocol
	bApp := bam.NewBaseApp(appName, logger, db, authutils.DefaultTxDecoder(cdc), baseAppOptions...)

	bApp.SetAppVersion(version.Version)

	keys := sdk.NewKVStoreKeys(bam.MainStoreKey, auth.StoreKey, validator.StoreKey,
		modelinfo.StoreKey, compliance.StoreKey, compliancetest.StoreKey, pki.StoreKey)

	tkeys := sdk.NewTransientStoreKeys(params.TStoreKey)

	// Here you initialize your application with the store keys it requires.
	var app = &zbLedgerApp{
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

func InitModuleManager(app *zbLedgerApp) {
	app.mm = module.NewManager(
		genaccounts.NewAppModule(app.authKeeper),
		genutil.NewAppModule(app.authKeeper, app.validatorKeeper, app.BaseApp.DeliverTx),
		auth.NewAppModule(app.authKeeper),
		validator.NewAppModule(app.validatorKeeper, app.authKeeper),
		modelinfo.NewAppModule(app.modelinfoKeeper, app.authKeeper),
		compliance.NewAppModule(app.complianceKeeper, app.modelinfoKeeper, app.compliancetestKeeper, app.authKeeper),
		compliancetest.NewAppModule(app.compliancetestKeeper, app.authKeeper, app.modelinfoKeeper),
		pki.NewAppModule(app.pkiKeeper, app.authKeeper),
	)

	app.mm.SetOrderBeginBlockers(validator.ModuleName)
	app.mm.SetOrderEndBlockers(validator.ModuleName)

	app.mm.SetOrderInitGenesis(
		genaccounts.ModuleName,
		auth.ModuleName,
		validator.ModuleName,
		modelinfo.ModuleName,
		compliance.ModuleName,
		compliancetest.ModuleName,
		pki.ModuleName,
		genutil.ModuleName,
	)

	// register all module routes and module queriers
	app.mm.RegisterRoutes(app.Router(), app.QueryRouter())
}

func InitKeepers(app *zbLedgerApp, keys map[string]*sdk.KVStoreKey) {
	// The Validator keeper
	app.validatorKeeper = MakeValidatorKeeper(keys, app)

	// The ModelinfoKeeper keeper
	app.modelinfoKeeper = MakeModelinfoKeeper(keys, app)

	// The ComplianceKeeper keeper
	app.complianceKeeper = MakeComplianceKeeper(keys, app)

	// The CompliancetestKeeper keeper
	app.compliancetestKeeper = MakeCompliancetestKeeper(keys, app)

	// The PKI keeper
	app.pkiKeeper = MakePkiKeeper(keys, app)

	// The AuthKeeper keeper
	app.authKeeper = MakeAuthKeeper(keys, app)
}

func MakeAuthKeeper(keys map[string]*sdk.KVStoreKey, app *zbLedgerApp) auth.Keeper {
	return auth.NewKeeper(
		keys[auth.StoreKey],
		app.cdc,
	)
}

func MakeModelinfoKeeper(keys map[string]*sdk.KVStoreKey, app *zbLedgerApp) modelinfo.Keeper {
	return modelinfo.NewKeeper(
		keys[modelinfo.StoreKey],
		app.cdc,
	)
}

func MakeComplianceKeeper(keys map[string]*sdk.KVStoreKey, app *zbLedgerApp) compliance.Keeper {
	return compliance.NewKeeper(
		keys[compliance.StoreKey],
		app.cdc,
	)
}

func MakeCompliancetestKeeper(keys map[string]*sdk.KVStoreKey, app *zbLedgerApp) compliancetest.Keeper {
	return compliancetest.NewKeeper(
		keys[compliancetest.StoreKey],
		app.cdc,
	)
}

func MakePkiKeeper(keys map[string]*sdk.KVStoreKey, app *zbLedgerApp) pki.Keeper {
	return pki.NewKeeper(
		keys[pki.StoreKey],
		app.cdc,
	)
}

func MakeValidatorKeeper(keys map[string]*sdk.KVStoreKey, app *zbLedgerApp) validator.Keeper {
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

func (app *zbLedgerApp) InitChainer(ctx sdk.Context, req abci.RequestInitChain) abci.ResponseInitChain {
	var genesisState GenesisState

	err := app.cdc.UnmarshalJSON(req.AppStateBytes, &genesisState)
	if err != nil {
		panic(err)
	}

	return app.mm.InitGenesis(ctx, genesisState)
}

func (app *zbLedgerApp) BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock) abci.ResponseBeginBlock {
	return app.mm.BeginBlock(ctx, req)
}
func (app *zbLedgerApp) EndBlocker(ctx sdk.Context, req abci.RequestEndBlock) abci.ResponseEndBlock {
	return app.mm.EndBlock(ctx, req)
}
func (app *zbLedgerApp) LoadHeight(height int64) error {
	return app.LoadVersion(height, app.keys[bam.MainStoreKey])
}

//_________________________________________________________

func (app *zbLedgerApp) ExportAppStateAndValidators(forZeroHeight bool, jailWhiteList []string,
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
