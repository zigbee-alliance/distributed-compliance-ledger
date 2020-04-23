package app

import (
	"encoding/json"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliance"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/pki"
	"os"

	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/authnext"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/authz"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliancetest"
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
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	distr "github.com/cosmos/cosmos-sdk/x/distribution"
	"github.com/cosmos/cosmos-sdk/x/genaccounts"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	"github.com/cosmos/cosmos-sdk/x/staking"
	"github.com/cosmos/cosmos-sdk/x/supply"
)

const appName = "zb-ledger"

var (
	// default home directories for the application CLI
	DefaultCLIHome = os.ExpandEnv("$HOME/.zblcli")

	// DefaultNodeHome sets the folder where the applcation data and configuration will be stored
	DefaultNodeHome = os.ExpandEnv("$HOME/.zbld")

	// NewBasicManager is in charge of setting up basic module elemnets
	ModuleBasics = module.NewBasicManager(
		genaccounts.AppModuleBasic{},
		genutil.AppModuleBasic{},
		auth.AppModuleBasic{},
		bank.AppModuleBasic{},
		staking.AppModuleBasic{},
		distr.AppModuleBasic{},
		params.AppModuleBasic{},
		slashing.AppModuleBasic{},
		supply.AppModuleBasic{},

		modelinfo.AppModuleBasic{},
		compliance.AppModuleBasic{},
		compliancetest.AppModuleBasic{},
		authnext.AppModuleBasic{},
		authz.AppModuleBasic{},
		pki.AppModuleBasic{},
	)
	// account permissions
	maccPerms = map[string][]string{
		auth.FeeCollectorName:     nil,
		distr.ModuleName:          nil,
		staking.BondedPoolName:    {supply.Burner, supply.Staking},
		staking.NotBondedPoolName: {supply.Burner, supply.Staking},
	}
)

// MakeCodec generates the necessary codecs for Amino
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
	accountKeeper        auth.AccountKeeper
	bankKeeper           bank.Keeper
	stakingKeeper        staking.Keeper
	slashingKeeper       slashing.Keeper
	distrKeeper          distr.Keeper
	supplyKeeper         supply.Keeper
	paramsKeeper         params.Keeper
	modelinfoKeeper      modelinfo.Keeper
	pkiKeeper            pki.Keeper
	complianceKeeper     compliance.Keeper
	compliancetestKeeper compliancetest.Keeper
	authzKeeper          authz.Keeper

	// Module Manager
	mm *module.Manager
}

// NewZbLedgerApp is a constructor function for zbLedgerApp
func NewZbLedgerApp(logger log.Logger, db dbm.DB, baseAppOptions ...func(*bam.BaseApp)) *zbLedgerApp {
	// First define the top level codec that will be shared by the different modules
	cdc := MakeCodec()

	// BaseApp handles interactions with Tendermint through the ABCI protocol
	bApp := bam.NewBaseApp(appName, logger, db, auth.DefaultTxDecoder(cdc), baseAppOptions...)

	bApp.SetAppVersion(version.Version)

	keys := sdk.NewKVStoreKeys(bam.MainStoreKey, auth.StoreKey, staking.StoreKey,
		supply.StoreKey, distr.StoreKey, slashing.StoreKey, params.StoreKey, modelinfo.StoreKey, authz.StoreKey,
		compliance.StoreKey, compliancetest.StoreKey, pki.StoreKey)

	tkeys := sdk.NewTransientStoreKeys(staking.TStoreKey, params.TStoreKey)

	// Here you initialize your application with the store keys it requires
	var app = &zbLedgerApp{
		BaseApp: bApp,
		cdc:     cdc,
		keys:    keys,
		tkeys:   tkeys,
	}

	InitKeepers(app, keys, tkeys)

	InitModuleManager(app)

	// The initChainer handles translating the genesis.json file into initial state for the network
	app.SetInitChainer(app.InitChainer)
	app.SetBeginBlocker(app.BeginBlocker)
	app.SetEndBlocker(app.EndBlocker)

	// The AnteHandler handles signature verification and transaction pre-processing
	app.SetAnteHandler(
		auth.NewAnteHandler(
			app.accountKeeper,
			app.supplyKeeper,
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
		genaccounts.NewAppModule(app.accountKeeper),
		genutil.NewAppModule(app.accountKeeper, app.stakingKeeper, app.BaseApp.DeliverTx),
		auth.NewAppModule(app.accountKeeper),
		bank.NewAppModule(app.bankKeeper, app.accountKeeper),
		supply.NewAppModule(app.supplyKeeper, app.accountKeeper),
		distr.NewAppModule(app.distrKeeper, app.supplyKeeper),
		slashing.NewAppModule(app.slashingKeeper, app.stakingKeeper),
		staking.NewAppModule(app.stakingKeeper, app.distrKeeper, app.accountKeeper, app.supplyKeeper),

		modelinfo.NewAppModule(app.modelinfoKeeper, app.authzKeeper),
		compliance.NewAppModule(app.complianceKeeper, app.modelinfoKeeper, app.compliancetestKeeper, app.authzKeeper),
		compliancetest.NewAppModule(app.compliancetestKeeper, app.authzKeeper, app.modelinfoKeeper),
		authnext.NewAppModule(app.accountKeeper, app.authzKeeper, app.cdc),
		authz.NewAppModule(app.authzKeeper),
		pki.NewAppModule(app.pkiKeeper, app.authzKeeper),
	)

	app.mm.SetOrderBeginBlockers(distr.ModuleName, slashing.ModuleName)
	app.mm.SetOrderEndBlockers(staking.ModuleName)

	// Sets the order of Genesis - Order matters, genutil is to always come last
	// NOTE: The genutils moodule must occur after staking so that pools are
	// properly initialized with tokens from genesis accounts.
	app.mm.SetOrderInitGenesis(
		genaccounts.ModuleName,
		distr.ModuleName,
		staking.ModuleName,
		auth.ModuleName,
		bank.ModuleName,
		slashing.ModuleName,
		modelinfo.ModuleName,
		compliance.ModuleName,
		compliancetest.ModuleName,
		authnext.ModuleName,
		authz.ModuleName,
		pki.ModuleName,
		supply.ModuleName,
		genutil.ModuleName,
	)

	// register all module routes and module queriers
	app.mm.RegisterRoutes(app.Router(), app.QueryRouter())
}

func InitKeepers(app *zbLedgerApp, keys map[string]*sdk.KVStoreKey, tkeys map[string]*sdk.TransientStoreKey) {
	// The ParamsKeeper handles parameter storage for the application
	app.paramsKeeper = MakeParamKeeper(app, keys, tkeys)

	// Set specific supspaces
	authSubspace := app.paramsKeeper.Subspace(auth.DefaultParamspace)
	bankSupspace := app.paramsKeeper.Subspace(bank.DefaultParamspace)
	stakingSubspace := app.paramsKeeper.Subspace(staking.DefaultParamspace)
	distrSubspace := app.paramsKeeper.Subspace(distr.DefaultParamspace)
	slashingSubspace := app.paramsKeeper.Subspace(slashing.DefaultParamspace)

	// The AccountKeeper handles address -> account lookups
	app.accountKeeper = MakeAccountKeeper(app, keys, authSubspace)

	// The BankKeeper allows you perform sdk.Coins interactions
	app.bankKeeper = MakeBankKeeper(app, bankSupspace)

	// The SupplyKeeper collects transaction fees and renders them to the fee distribution module
	app.supplyKeeper = MakeSupplyKeeper(app, keys)

	// The staking keeper
	stakingKeeper := MakeStakingKeeper(app, keys, tkeys, stakingSubspace)

	app.distrKeeper = MakeDistrKeeper(app, keys, distrSubspace, &stakingKeeper)

	app.slashingKeeper = MakeSlashingKeeper(app, keys, &stakingKeeper, slashingSubspace)

	// register the staking hooks
	// NOTE: stakingKeeper above is passed by reference, so that it will contain these hooks
	app.stakingKeeper = *stakingKeeper.SetHooks(
		staking.NewMultiStakingHooks(
			app.distrKeeper.Hooks(),
			app.slashingKeeper.Hooks()),
	)

	// The ModelinfoKeeper keeper
	app.modelinfoKeeper = MakeModelinfoKeeper(keys, app)

	// The ComplianceKeeper keeper
	app.complianceKeeper = MakeComplianceKeeper(keys, app)

	// The CompliancetestKeeper keeper
	app.compliancetestKeeper = MakeCompliancetestKeeper(keys, app)

	// The PKI keeper
	app.pkiKeeper = MakePkiKeeper(keys, app)

	// The AuthzKeeper keeper
	app.authzKeeper = MakeAuthzKeeper(keys, app)
}

func MakeAuthzKeeper(keys map[string]*sdk.KVStoreKey, app *zbLedgerApp) authz.Keeper {
	return authz.NewKeeper(
		keys[authz.StoreKey],
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

func MakeSlashingKeeper(app *zbLedgerApp, keys map[string]*sdk.KVStoreKey,
	stakingKeeper *staking.Keeper, slashingSubspace params.Subspace) slashing.Keeper {
	return slashing.NewKeeper(
		app.cdc,
		keys[slashing.StoreKey],
		stakingKeeper,
		slashingSubspace,
		slashing.DefaultCodespace,
	)
}

func MakeDistrKeeper(app *zbLedgerApp, keys map[string]*sdk.KVStoreKey,
	distrSubspace params.Subspace, stakingKeeper *staking.Keeper) distr.Keeper {
	return distr.NewKeeper(
		app.cdc,
		keys[distr.StoreKey],
		distrSubspace,
		stakingKeeper,
		app.supplyKeeper,
		distr.DefaultCodespace,
		auth.FeeCollectorName,
		app.ModuleAccountAddrs(),
	)
}

func MakeStakingKeeper(app *zbLedgerApp, keys map[string]*sdk.KVStoreKey,
	tkeys map[string]*sdk.TransientStoreKey, stakingSubspace params.Subspace) staking.Keeper {
	return staking.NewKeeper(
		app.cdc,
		keys[staking.StoreKey],
		tkeys[staking.TStoreKey],
		app.supplyKeeper,
		stakingSubspace,
		staking.DefaultCodespace,
	)
}

func MakeSupplyKeeper(app *zbLedgerApp, keys map[string]*sdk.KVStoreKey) supply.Keeper {
	return supply.NewKeeper(
		app.cdc,
		keys[supply.StoreKey],
		app.accountKeeper,
		app.bankKeeper,
		maccPerms,
	)
}

func MakeBankKeeper(app *zbLedgerApp, bankSupspace params.Subspace) bank.Keeper {
	return bank.NewBaseKeeper(
		app.accountKeeper,
		bankSupspace,
		bank.DefaultCodespace,
		app.ModuleAccountAddrs(),
	)
}

func MakeAccountKeeper(app *zbLedgerApp,
	keys map[string]*sdk.KVStoreKey, authSubspace params.Subspace) auth.AccountKeeper {
	return auth.NewAccountKeeper(
		app.cdc,
		keys[auth.StoreKey],
		authSubspace,
		auth.ProtoBaseAccount,
	)
}

func MakeParamKeeper(app *zbLedgerApp,
	keys map[string]*sdk.KVStoreKey, tkeys map[string]*sdk.TransientStoreKey) params.Keeper {
	return params.NewKeeper(
		app.cdc, keys[params.StoreKey],
		tkeys[params.TStoreKey],
		params.DefaultCodespace,
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

// ModuleAccountAddrs returns all the app's module account addresses.
func (app *zbLedgerApp) ModuleAccountAddrs() map[string]bool {
	modAccAddrs := make(map[string]bool)
	for acc := range maccPerms {
		modAccAddrs[supply.NewModuleAddress(acc).String()] = true
	}

	return modAccAddrs
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

	validators = staking.WriteValidators(ctx, app.stakingKeeper)

	return appState, validators, nil
}
