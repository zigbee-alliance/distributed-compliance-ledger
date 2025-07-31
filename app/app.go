package app

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"

	reflectionv1 "cosmossdk.io/api/cosmos/reflection/v1"
	dbm "github.com/cometbft/cometbft-db"
	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cometbft/cometbft/libs/log"
	tmos "github.com/cometbft/cometbft/libs/os"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	nodeservice "github.com/cosmos/cosmos-sdk/client/grpc/node"
	"github.com/cosmos/cosmos-sdk/client/grpc/tmservice"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	runtimeservices "github.com/cosmos/cosmos-sdk/runtime/services"
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/server/api"
	"github.com/cosmos/cosmos-sdk/server/config"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/version"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	consensusparamkeeper "github.com/cosmos/cosmos-sdk/x/consensus/keeper"
	"github.com/cosmos/cosmos-sdk/x/params"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/cosmos/cosmos-sdk/x/upgrade"
	upgradekeeper "github.com/cosmos/cosmos-sdk/x/upgrade/keeper"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"

	appparams "github.com/zigbee-alliance/distributed-compliance-ledger/app/params"
	dclpkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	compliancemodule "github.com/zigbee-alliance/distributed-compliance-ledger/x/compliance"
	compliancemodulekeeper "github.com/zigbee-alliance/distributed-compliance-ledger/x/compliance/keeper"
	dclcompliancetypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/compliance/types"
	dclauthmodule "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/ante"
	baseauthmodulekeeper "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/base/keeper"
	dclauthmodulekeeper "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/keeper"
	dclauthmoduletypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
	dclgenutilmodule "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclgenutil"
	dclgenutilmoduletypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclgenutil/types"
	dclupgrademodule "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclupgrade"
	dclupgrademodulekeeper "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclupgrade/keeper"
	dclupgrademoduletypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclupgrade/types"
	modelmodule "github.com/zigbee-alliance/distributed-compliance-ledger/x/model"
	modelmodulekeeper "github.com/zigbee-alliance/distributed-compliance-ledger/x/model/keeper"
	dclmodeltypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/model/types"
	pkimodule "github.com/zigbee-alliance/distributed-compliance-ledger/x/pki"
	pkimodulekeeper "github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/keeper"
	validatormodule "github.com/zigbee-alliance/distributed-compliance-ledger/x/validator"
	validatormodulekeeper "github.com/zigbee-alliance/distributed-compliance-ledger/x/validator/keeper"
	validatormoduletypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/validator/types"
	vendorinfomodule "github.com/zigbee-alliance/distributed-compliance-ledger/x/vendorinfo"
	vendorinfomodulekeeper "github.com/zigbee-alliance/distributed-compliance-ledger/x/vendorinfo/keeper"
	vendorinfomoduletypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/vendorinfo/types"
)

const (
	AccountAddressPrefix = "cosmos"
	Name                 = "dcl"
)

// this line is used by starport scaffolding # stargate/wasm/app/enabledProposals
/*
func getGovProposalHandlers() []govclient.ProposalHandler {
	var govProposalHandlers []govclient.ProposalHandler
	// this line is used by starport scaffolding # stargate/app/govProposalHandlers

	govProposalHandlers = append(govProposalHandlers,
		paramsclient.ProposalHandler,
		distrclient.ProposalHandler,
		upgradeclient.ProposalHandler,
		upgradeclient.CancelProposalHandler,
		// this line is used by starport scaffolding # stargate/app/govProposalHandler
	)

	return govProposalHandlers
}
*/

var (
	// DefaultNodeHome default home directories for the application daemon.
	DefaultNodeHome string

	// ModuleBasics defines the module BasicManager is in charge of setting up basic,
	// non-dependant module elements, such as codec registration
	// and genesis verification.
	ModuleBasics = module.NewBasicManager(
		/*
			auth.AppModuleBasic{},
			genutil.AppModuleBasic{},
			bank.AppModuleBasic{},
			capability.AppModuleBasic{},
			staking.AppModuleBasic{},
			mint.AppModuleBasic{},
			distr.AppModuleBasic{},
			gov.NewAppModuleBasic(getGovProposalHandlers()...),
		*/
		params.AppModuleBasic{},
		/*
			crisis.AppModuleBasic{},
			slashing.AppModuleBasic{},
			feegrantmodule.AppModuleBasic{},
			ibc.AppModuleBasic{},
		*/
		upgrade.AppModuleBasic{},
		/*
			evidence.AppModuleBasic{},
			transfer.AppModuleBasic{},
			vesting.AppModuleBasic{},
		*/
		dclauthmodule.AppModuleBasic{},
		validatormodule.AppModuleBasic{},
		dclgenutilmodule.AppModuleBasic{},
		dclupgrademodule.AppModuleBasic{},
		pkimodule.AppModuleBasic{},
		vendorinfomodule.AppModuleBasic{},
		modelmodule.AppModuleBasic{},
		compliancemodule.AppModuleBasic{},
		// this line is used by starport scaffolding # stargate/app/moduleBasic
	)

	// module account permissions.
	/* maccPerms = map[string][]string{
			authtypes.FeeCollectorName:     nil,
			distrtypes.ModuleName:          nil,
			minttypes.ModuleName:           {authtypes.Minter},
			stakingtypes.BondedPoolName:    {authtypes.Burner, authtypes.Staking},
			stakingtypes.NotBondedPoolName: {authtypes.Burner, authtypes.Staking},
			govtypes.ModuleName:            {authtypes.Burner},
			ibctransfertypes.ModuleName:    {authtypes.Minter, authtypes.Burner},
			// this line is used by starport scaffolding # stargate/app/maccPerms
	}
	*/
)

var (
	_ runtime.AppI            = (*App)(nil)
	_ servertypes.Application = (*App)(nil)
)

func init() {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	DefaultNodeHome = filepath.Join(userHomeDir, "."+Name)
}

// App extends an ABCI application, but with most of its parameters exported.
// They are exported for convenience in creating helper functions, as object
// capabilities aren't needed for testing.
type App struct {
	*baseapp.BaseApp

	cdc               *codec.LegacyAmino
	appCodec          codec.Codec
	interfaceRegistry types.InterfaceRegistry
	txConfig          client.TxConfig

	invCheckPeriod uint

	// keys to access the substores
	keys    map[string]*storetypes.KVStoreKey
	tkeys   map[string]*storetypes.TransientStoreKey
	memKeys map[string]*storetypes.MemoryStoreKey

	// keepers
	/*
		AccountKeeper    authkeeper.AccountKeeper
		BankKeeper       bankkeeper.Keeper
		CapabilityKeeper *capabilitykeeper.Keeper
		StakingKeeper    stakingkeeper.Keeper
		SlashingKeeper   slashingkeeper.Keeper
		MintKeeper       mintkeeper.Keeper
		DistrKeeper      distrkeeper.Keeper
		GovKeeper        govkeeper.Keeper
		CrisisKeeper     crisiskeeper.Keeper
	*/
	UpgradeKeeper *upgradekeeper.Keeper
	ParamsKeeper  paramskeeper.Keeper
	/*
		IBCKeeper        *ibckeeper.Keeper // IBC Keeper must be a pointer in the app, so we can SetRouter on it correctly
		EvidenceKeeper   evidencekeeper.Keeper
		TransferKeeper   ibctransferkeeper.Keeper
		FeeGrantKeeper   feegrantkeeper.Keeper

		// make scoped keepers public for test purposes
		ScopedIBCKeeper      capabilitykeeper.ScopedKeeper
		ScopedTransferKeeper capabilitykeeper.ScopedKeeper
	*/

	DclauthKeeper dclauthmodulekeeper.Keeper

	BaseauthKeeper baseauthmodulekeeper.Keeper

	ValidatorKeeper validatormodulekeeper.Keeper

	DclupgradeKeeper dclupgrademodulekeeper.Keeper

	PkiKeeper pkimodulekeeper.Keeper

	VendorinfoKeeper vendorinfomodulekeeper.Keeper

	ModelKeeper modelmodulekeeper.Keeper

	ComplianceKeeper compliancemodulekeeper.Keeper

	ConsensusParamsKeeper consensusparamkeeper.Keeper
	// this line is used by starport scaffolding # stargate/app/keeperDeclaration

	// the module manager
	mm *module.Manager

	// sm is the simulation manager
	sm           *module.SimulationManager
	configurator module.Configurator
}

func (app *App) RegisterNodeService(clientCtx client.Context) {
	nodeservice.RegisterNodeService(clientCtx, app.GRPCQueryRouter())
}

func (app *App) SimulationManager() *module.SimulationManager {
	return app.sm
}

// New returns a reference to an initialized Gaia.
func New(
	logger log.Logger,
	db dbm.DB,
	traceStore io.Writer,
	loadLatest bool,
	skipUpgradeHeights map[int64]bool,
	homePath string,
	invCheckPeriod uint,
	encodingConfig appparams.EncodingConfig,
	baseAppOptions ...func(*baseapp.BaseApp),
) *App {
	appCodec := encodingConfig.Codec
	cdc := encodingConfig.Amino
	interfaceRegistry := encodingConfig.InterfaceRegistry
	txConfig := encodingConfig.TxConfig

	bApp := baseapp.NewBaseApp(Name, logger, db, encodingConfig.TxConfig.TxDecoder(), baseAppOptions...)
	bApp.SetCommitMultiStoreTracer(traceStore)
	bApp.SetVersion(version.Version)
	bApp.SetInterfaceRegistry(interfaceRegistry)
	bApp.SetTxEncoder(txConfig.TxEncoder())

	keys := sdk.NewKVStoreKeys(
		paramstypes.StoreKey,
		upgradetypes.StoreKey,
		/*
			capabilitytypes.StoreKey,
			authtypes.StoreKey, banktypes.StoreKey, stakingtypes.StoreKey,
			minttypes.StoreKey, distrtypes.StoreKey, slashingtypes.StoreKey,
			govtypes.StoreKey, ibchost.StoreKey, feegrant.StoreKey,
			evidencetypes.StoreKey, ibctransfertypes.StoreKey, capabilitytypes.StoreKey,
		*/
		dclauthmoduletypes.StoreKey,
		validatormoduletypes.StoreKey,
		dclupgrademoduletypes.StoreKey,
		dclpkitypes.StoreKey,
		vendorinfomoduletypes.StoreKey,
		dclmodeltypes.StoreKey,
		dclcompliancetypes.StoreKey,
		// this line is used by starport scaffolding # stargate/app/storeKey
	)
	tkeys := sdk.NewTransientStoreKeys(paramstypes.TStoreKey)
	memKeys := sdk.NewMemoryStoreKeys( /*capabilitytypes.MemStoreKey*/ )

	app := &App{
		BaseApp:           bApp,
		cdc:               cdc,
		appCodec:          appCodec,
		interfaceRegistry: interfaceRegistry,
		txConfig:          txConfig,
		invCheckPeriod:    invCheckPeriod,
		keys:              keys,
		tkeys:             tkeys,
		memKeys:           memKeys,
	}

	app.ParamsKeeper = initParamsKeeper(appCodec, cdc, keys[paramstypes.StoreKey], tkeys[paramstypes.TStoreKey])

	// set the BaseApp's parameter store
	app.ConsensusParamsKeeper = consensusparamkeeper.NewKeeper(appCodec, keys[upgradetypes.StoreKey], authtypes.NewModuleAddress(authtypes.ModuleName).String())
	bApp.SetParamStore(&app.ConsensusParamsKeeper) // add capability keeper and ScopeToModule for ibc module
	// app.CapabilityKeeper = capabilitykeeper.NewKeeper(appCodec, keys[capabilitytypes.StoreKey], memKeys[capabilitytypes.MemStoreKey])

	/*
		// grant capabilities for the ibc and ibc-transfer modules
		scopedIBCKeeper := app.CapabilityKeeper.ScopeToModule(ibchost.ModuleName)
		scopedTransferKeeper := app.CapabilityKeeper.ScopeToModule(ibctransfertypes.ModuleName)
	*/
	// this line is used by starport scaffolding # stargate/app/scopedKeeper

	// add keepers
	/*
		app.AccountKeeper = authkeeper.NewAccountKeeper(
			appCodec, keys[authtypes.StoreKey], app.GetSubspace(authtypes.ModuleName), authtypes.ProtoBaseAccount, maccPerms,
		)
		app.BankKeeper = bankkeeper.NewBaseKeeper(
			appCodec, keys[banktypes.StoreKey], app.AccountKeeper, app.GetSubspace(banktypes.ModuleName), app.ModuleAccountAddrs(),
		)
		stakingKeeper := stakingkeeper.NewKeeper(
			appCodec, keys[stakingtypes.StoreKey], app.AccountKeeper, app.BankKeeper, app.GetSubspace(stakingtypes.ModuleName),
		)
		app.MintKeeper = mintkeeper.NewKeeper(
			appCodec, keys[minttypes.StoreKey], app.GetSubspace(minttypes.ModuleName), &stakingKeeper,
			app.AccountKeeper, app.BankKeeper, authtypes.FeeCollectorName,
		)
		app.DistrKeeper = distrkeeper.NewKeeper(
			appCodec, keys[distrtypes.StoreKey], app.GetSubspace(distrtypes.ModuleName), app.AccountKeeper, app.BankKeeper,
			&stakingKeeper, authtypes.FeeCollectorName, app.ModuleAccountAddrs(),
		)
		app.SlashingKeeper = slashingkeeper.NewKeeper(
			appCodec, keys[slashingtypes.StoreKey], &stakingKeeper, app.GetSubspace(slashingtypes.ModuleName),
		)
		app.CrisisKeeper = crisiskeeper.NewKeeper(
			app.GetSubspace(crisistypes.ModuleName), invCheckPeriod, app.BankKeeper, authtypes.FeeCollectorName,
		)

		app.FeeGrantKeeper = feegrantkeeper.NewKeeper(appCodec, keys[feegrant.StoreKey], app.AccountKeeper)
	*/
	app.UpgradeKeeper = upgradekeeper.NewKeeper(skipUpgradeHeights, keys[upgradetypes.StoreKey], appCodec, homePath, app.BaseApp, authtypes.NewModuleAddress(upgradetypes.ModuleName).String())
	/*
		// register the staking hooks
		// NOTE: stakingKeeper above is passed by reference, so that it will contain these hooks
		app.StakingKeeper = *stakingKeeper.SetHooks(
			stakingtypes.NewMultiStakingHooks(app.DistrKeeper.Hooks(), app.SlashingKeeper.Hooks()),
		)

		// ... other modules keepers

		// Create IBC Keeper
		app.IBCKeeper = ibckeeper.NewKeeper(
			appCodec, keys[ibchost.StoreKey], app.GetSubspace(ibchost.ModuleName), app.StakingKeeper, app.UpgradeKeeper, scopedIBCKeeper,
		)

		// register the proposal types
		govRouter := govtypes.NewRouter()
		govRouter.AddRoute(govtypes.RouterKey, govtypes.ProposalHandler).
			AddRoute(paramproposal.RouterKey, params.NewParamChangeProposalHandler(app.ParamsKeeper)).
			AddRoute(distrtypes.RouterKey, distr.NewCommunityPoolSpendProposalHandler(app.DistrKeeper)).
			AddRoute(upgradetypes.RouterKey, upgrade.NewSoftwareUpgradeProposalHandler(app.UpgradeKeeper)).
			AddRoute(ibchost.RouterKey, ibcclient.NewClientProposalHandler(app.IBCKeeper.ClientKeeper))

		// Create Transfer Keepers
		app.TransferKeeper = ibctransferkeeper.NewKeeper(
			appCodec, keys[ibctransfertypes.StoreKey], app.GetSubspace(ibctransfertypes.ModuleName),
			app.IBCKeeper.ChannelKeeper, &app.IBCKeeper.PortKeeper,
			app.AccountKeeper, app.BankKeeper, scopedTransferKeeper,
		)
		transferModule := transfer.NewAppModule(app.TransferKeeper)

		// Create evidence Keeper for to register the IBC light client misbehaviour evidence route
		evidenceKeeper := evidencekeeper.NewKeeper(
			appCodec, keys[evidencetypes.StoreKey], &app.StakingKeeper, app.SlashingKeeper,
		)
		// If evidence needs to be handled for the app, set routes in router here and seal
		app.EvidenceKeeper = *evidenceKeeper

		app.GovKeeper = govkeeper.NewKeeper(
			appCodec, keys[govtypes.StoreKey], app.GetSubspace(govtypes.ModuleName), app.AccountKeeper, app.BankKeeper,
			&stakingKeeper, govRouter,
		)
	*/

	app.DclauthKeeper = *dclauthmodulekeeper.NewKeeper(
		appCodec,
		keys[dclauthmoduletypes.StoreKey],
		keys[dclauthmoduletypes.MemStoreKey],
	)

	app.BaseauthKeeper = *baseauthmodulekeeper.NewKeeper(
		appCodec,
		keys[dclauthmoduletypes.StoreKey],
		keys[dclauthmoduletypes.MemStoreKey],
	)

	dclauthModule := dclauthmodule.NewAppModule(appCodec, app.DclauthKeeper, app.BaseauthKeeper)

	app.ValidatorKeeper = *validatormodulekeeper.NewKeeper(
		appCodec,
		keys[validatormoduletypes.StoreKey],
		keys[validatormoduletypes.MemStoreKey],

		app.DclauthKeeper,
	)
	validatorModule := validatormodule.NewAppModule(appCodec, app.ValidatorKeeper)

	dclgenutilModule := dclgenutilmodule.NewAppModule(
		app.DclauthKeeper, app.ValidatorKeeper, app.BaseApp.DeliverTx,
		encodingConfig.TxConfig,
	)

	app.DclupgradeKeeper = *dclupgrademodulekeeper.NewKeeper(
		appCodec,
		keys[dclupgrademoduletypes.StoreKey],
		keys[dclupgrademoduletypes.MemStoreKey],

		app.DclauthKeeper,
		app.UpgradeKeeper,
	)
	dclupgradeModule := dclupgrademodule.NewAppModule(appCodec, app.DclupgradeKeeper)

	app.PkiKeeper = *pkimodulekeeper.NewKeeper(
		appCodec,
		keys[dclpkitypes.StoreKey],
		keys[dclpkitypes.MemStoreKey],

		app.DclauthKeeper,
	)
	pkiModule := pkimodule.NewAppModule(appCodec, app.PkiKeeper)

	app.VendorinfoKeeper = *vendorinfomodulekeeper.NewKeeper(
		appCodec,
		keys[vendorinfomoduletypes.StoreKey],
		keys[vendorinfomoduletypes.MemStoreKey],

		app.DclauthKeeper,
	)
	vendorinfoModule := vendorinfomodule.NewAppModule(appCodec, app.VendorinfoKeeper)

	app.ModelKeeper = *modelmodulekeeper.NewKeeper(
		appCodec,
		keys[dclmodeltypes.StoreKey],
		keys[dclmodeltypes.MemStoreKey],

		app.DclauthKeeper,
		app.ComplianceKeeper,
	)

	app.ComplianceKeeper = *compliancemodulekeeper.NewKeeper(
		appCodec,
		keys[dclcompliancetypes.StoreKey],
		keys[dclcompliancetypes.MemStoreKey],
		app.DclauthKeeper,
		app.ModelKeeper,
	)

	app.ModelKeeper.SetComplianceKeeper(app.ComplianceKeeper)

	modelModule := modelmodule.NewAppModule(appCodec, app.ModelKeeper)
	complianceModule := compliancemodule.NewAppModule(appCodec, app.ComplianceKeeper)

	// this line is used by starport scaffolding # stargate/app/keeperDefinition

	/*
		// Create static IBC router, add transfer route, then set and seal it
		ibcRouter := ibcporttypes.NewRouter()
		ibcRouter.AddRoute(ibctransfertypes.ModuleName, transferModule)
		// this line is used by starport scaffolding # ibc/app/router
		app.IBCKeeper.SetRouter(ibcRouter)
	*/

	/****  Module Options ****/

	// NOTE: we may consider parsing `appOpts` inside module constructors. For the moment
	// we prefer to be more strict in what arguments the modules expect.
	/*
		var skipGenesisInvariants = cast.ToBool(appOpts.Get(crisis.FlagSkipGenesisInvariants))
	*/

	// NOTE: Any module instantiated in the module manager that is later modified
	// must be passed by reference here.

	app.mm = module.NewManager(
		/*
			genutil.NewAppModule(
				app.AccountKeeper, app.StakingKeeper, app.BaseApp.DeliverTx,
				encodingConfig.TxConfig,
			),
			auth.NewAppModule(appCodec, app.AccountKeeper, nil),
			vesting.NewAppModule(app.AccountKeeper, app.BankKeeper),
			bank.NewAppModule(appCodec, app.BankKeeper, app.AccountKeeper),
			capability.NewAppModule(appCodec, *app.CapabilityKeeper),
			feegrantmodule.NewAppModule(appCodec, app.AccountKeeper, app.BankKeeper, app.FeeGrantKeeper, app.interfaceRegistry),
			crisis.NewAppModule(&app.CrisisKeeper, skipGenesisInvariants),
			gov.NewAppModule(appCodec, app.GovKeeper, app.AccountKeeper, app.BankKeeper),
			mint.NewAppModule(appCodec, app.MintKeeper, app.AccountKeeper),
			slashing.NewAppModule(appCodec, app.SlashingKeeper, app.AccountKeeper, app.BankKeeper, app.StakingKeeper),
			distr.NewAppModule(appCodec, app.DistrKeeper, app.AccountKeeper, app.BankKeeper, app.StakingKeeper),
			staking.NewAppModule(appCodec, app.StakingKeeper, app.AccountKeeper, app.BankKeeper),
			evidence.NewAppModule(app.EvidenceKeeper),
			ibc.NewAppModule(app.IBCKeeper),
			transferModule,
		*/
		params.NewAppModule(app.ParamsKeeper),
		upgrade.NewAppModule(app.UpgradeKeeper),
		dclauthModule,
		validatorModule,
		dclgenutilModule,
		dclupgradeModule,
		pkiModule,
		vendorinfoModule,
		modelModule,
		complianceModule,
		// this line is used by starport scaffolding # stargate/app/appModule
	)

	// During begin block slashing happens after distr.BeginBlocker so that
	// there is nothing left over in the validator fee pool, so as to keep the
	// CanWithdrawInvariant invariant.
	// NOTE: staking module is required if HistoricalEntries param > 0
	/*
		app.mm.SetOrderBeginBlockers(
			upgradetypes.ModuleName, capabilitytypes.ModuleName, minttypes.ModuleName, distrtypes.ModuleName, slashingtypes.ModuleName,
			evidencetypes.ModuleName, stakingtypes.ModuleName, ibchost.ModuleName,
			feegrant.ModuleName,
		)

		app.mm.SetOrderEndBlockers(crisistypes.ModuleName, govtypes.ModuleName, stakingtypes.ModuleName)
	*/

	app.mm.SetOrderBeginBlockers(
		// TODO [issue 99] verify the order
		dclauthmoduletypes.ModuleName,
		validatormoduletypes.ModuleName,
		dclgenutilmoduletypes.ModuleName,
		dclupgrademoduletypes.ModuleName,
		upgradetypes.ModuleName,
		dclpkitypes.ModuleName,
		dclauthmoduletypes.ModuleName,
		dclmodeltypes.ModuleName,
		dclcompliancetypes.ModuleName,
		vendorinfomoduletypes.ModuleName,
		paramstypes.ModuleName,
	)

	app.mm.SetOrderEndBlockers(
		dclauthmoduletypes.ModuleName,
		validatormoduletypes.ModuleName,
		dclgenutilmoduletypes.ModuleName,
		dclupgrademoduletypes.ModuleName,
		upgradetypes.ModuleName,
		dclpkitypes.ModuleName,
		dclauthmoduletypes.ModuleName,
		dclmodeltypes.ModuleName,
		dclcompliancetypes.ModuleName,
		vendorinfomoduletypes.ModuleName,
		paramstypes.ModuleName,
	)

	// NOTE: The genutils module must occur after staking so that pools are
	// properly initialized with tokens from genesis accounts.
	// NOTE: Capability module must occur first so that it can initialize any capabilities
	// so that other modules that want to create or claim capabilities afterwards in InitChain
	// can do so safely.
	app.mm.SetOrderInitGenesis(
		// TODO [issue 99] verify the order
		/*
			capabilitytypes.ModuleName,
			authtypes.ModuleName,
			banktypes.ModuleName,
			distrtypes.ModuleName,
			stakingtypes.ModuleName,
			slashingtypes.ModuleName,
			govtypes.ModuleName,
			minttypes.ModuleName,
			crisistypes.ModuleName,
			ibchost.ModuleName,
			genutiltypes.ModuleName,
			evidencetypes.ModuleName,
			ibctransfertypes.ModuleName,
		*/
		dclauthmoduletypes.ModuleName,
		validatormoduletypes.ModuleName,
		dclgenutilmoduletypes.ModuleName,
		dclupgrademoduletypes.ModuleName,
		dclpkitypes.ModuleName,
		vendorinfomoduletypes.ModuleName,
		dclmodeltypes.ModuleName,
		dclcompliancetypes.ModuleName,
		paramstypes.ModuleName,
		upgradetypes.ModuleName,
		dclupgrademoduletypes.ModuleName,
		// this line is used by starport scaffolding # stargate/app/initGenesis
	)

	// app.mm.RegisterInvariants(&app.CrisisKeeper)
	// app.mm.RegisterRoutes(app.Router(), app.QueryRouter(), encodingConfig.Amino)

	app.configurator = module.NewConfigurator(app.appCodec, app.MsgServiceRouter(), app.GRPCQueryRouter())
	app.mm.RegisterServices(app.configurator)

	// autocliv1.RegisterQueryServer(app.GRPCQueryRouter(), runtimeservices.NewAutoCLIQueryService(app.mm.Modules))
	reflectionSvc, err := runtimeservices.NewReflectionService()
	if err != nil {
		panic(err)
	}
	reflectionv1.RegisterReflectionServiceServer(app.GRPCQueryRouter(), reflectionSvc)

	// create the simulation manager and define the order of the modules for deterministic simulations
	overrideModules := map[string]module.AppModuleSimulation{
		dclauthmoduletypes.ModuleName: dclauthModule}
	app.sm = module.NewSimulationManagerFromAppModules(app.mm.Modules, overrideModules)
	app.sm.RegisterStoreDecoders()

	// initialize stores
	app.MountKVStores(keys)
	app.MountTransientStores(tkeys)
	app.MountMemoryStores(memKeys)

	// initialize BaseApp
	app.SetInitChainer(app.InitChainer)
	app.SetBeginBlocker(app.BeginBlocker)

	anteHandler, err := ante.NewAnteHandler(
		ante.HandlerOptions{
			AccountKeeper:   app.DclauthKeeper,
			SignModeHandler: encodingConfig.TxConfig.SignModeHandler(),
			SigGasConsumer:  ante.DefaultSigVerificationGasConsumer,
		},
	)
	if err != nil {
		panic(err)
	}

	app.SetAnteHandler(anteHandler)
	app.SetEndBlocker(app.EndBlocker)

	if loadLatest {
		if err := app.LoadLatestVersion(); err != nil {
			tmos.Exit(err.Error())
		}
	}

	/*
		app.ScopedIBCKeeper = scopedIBCKeeper
		app.ScopedTransferKeeper = scopedTransferKeeper
	*/
	// this line is used by starport scaffolding # stargate/app/beforeInitReturn
	app.UpgradeKeeper.SetUpgradeHandler(
		"v0.10.0",
		func(ctx sdk.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
			return make(module.VersionMap), nil
		},
	)

	app.UpgradeKeeper.SetUpgradeHandler(
		"v0.11.0",
		func(ctx sdk.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
			return make(module.VersionMap), nil
		},
	)

	app.UpgradeKeeper.SetUpgradeHandler(
		"v0.12.0",
		func(ctx sdk.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
			return make(module.VersionMap), nil
		},
	)

	app.UpgradeKeeper.SetUpgradeHandler(
		"v0.13.0-pre",
		func(ctx sdk.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
			return make(module.VersionMap), nil
		},
	)

	app.UpgradeKeeper.SetUpgradeHandler(
		"v1.2",
		func(ctx sdk.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
			return make(module.VersionMap), nil
		},
	)

	app.UpgradeKeeper.SetUpgradeHandler(
		"v1.4",
		func(ctx sdk.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
			return app.mm.RunMigrations(ctx, app.configurator, fromVM)
		},
	)

	app.UpgradeKeeper.SetUpgradeHandler(
		"v1.4.4",
		func(ctx sdk.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
			return app.mm.RunMigrations(ctx, app.configurator, fromVM)
		},
	)

	app.UpgradeKeeper.SetUpgradeHandler(
		"v1.5",
		func(ctx sdk.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
			return make(module.VersionMap), nil
		},
	)

	return app
}

// Name returns the name of the App.
func (app *App) Name() string { return app.BaseApp.Name() }

// BeginBlocker application updates every begin block.
func (app *App) BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock) abci.ResponseBeginBlock {
	return app.mm.BeginBlock(ctx, req)
}

// EndBlocker application updates every end block.
func (app *App) EndBlocker(ctx sdk.Context, req abci.RequestEndBlock) abci.ResponseEndBlock {
	return app.mm.EndBlock(ctx, req)
}

// InitChainer application update at chain initialization.
func (app *App) InitChainer(ctx sdk.Context, req abci.RequestInitChain) abci.ResponseInitChain {
	var genesisState GenesisState
	if err := json.Unmarshal(req.AppStateBytes, &genesisState); err != nil {
		panic(err)
	}
	app.UpgradeKeeper.SetModuleVersionMap(ctx, app.mm.GetVersionMap())

	return app.mm.InitGenesis(ctx, app.appCodec, genesisState)
}

// LoadHeight loads a particular height.
func (app *App) LoadHeight(height int64) error {
	return app.LoadVersion(height)
}

// ModuleAccountAddrs returns all the app's module account addresses.
func (app *App) ModuleAccountAddrs() map[string]bool {
	modAccAddrs := make(map[string]bool)
	/*
		for acc := range maccPerms {
			modAccAddrs[authtypes.NewModuleAddress(acc).String()] = true
		}
	*/

	return modAccAddrs
}

// LegacyAmino returns SimApp's amino codec.
//
// NOTE: This is solely to be used for testing purposes as it may be desirable
// for modules to register their own custom testing types.
func (app *App) LegacyAmino() *codec.LegacyAmino {
	return app.cdc
}

// AppCodec returns Gaia's app codec.
//
// NOTE: This is solely to be used for testing purposes as it may be desirable
// for modules to register their own custom testing types.
func (app *App) AppCodec() codec.Codec {
	return app.appCodec
}

// InterfaceRegistry returns Gaia's InterfaceRegistry.
func (app *App) InterfaceRegistry() types.InterfaceRegistry {
	return app.interfaceRegistry
}

// GetKey returns the KVStoreKey for the provided store key.
//
// NOTE: This is solely to be used for testing purposes.
func (app *App) GetKey(storeKey string) *storetypes.KVStoreKey {
	return app.keys[storeKey]
}

// GetTKey returns the TransientStoreKey for the provided store key.
//
// NOTE: This is solely to be used for testing purposes.
func (app *App) GetTKey(storeKey string) *storetypes.TransientStoreKey {
	return app.tkeys[storeKey]
}

// GetMemKey returns the MemStoreKey for the provided mem key.
//
// NOTE: This is solely used for testing purposes.
func (app *App) GetMemKey(storeKey string) *storetypes.MemoryStoreKey {
	return app.memKeys[storeKey]
}

// GetSubspace returns a param subspace for a given module name.
//
// NOTE: This is solely to be used for testing purposes.
func (app *App) GetSubspace(moduleName string) paramstypes.Subspace {
	subspace, _ := app.ParamsKeeper.GetSubspace(moduleName)

	return subspace
}

// RegisterAPIRoutes registers all application module routes with the provided
// API server.
func (app *App) RegisterAPIRoutes(apiSvr *api.Server, apiConfig config.APIConfig) {
	clientCtx := apiSvr.ClientCtx
	// Register new tx routes from grpc-gateway.
	authtx.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)

	// Register new tendermint queries routes from grpc-gateway.
	tmservice.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)

	// Register legacy and grpc-gateway routes for all modules.
	ModuleBasics.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)

	// register swagger API from root so that other applications can override easily
	if err := server.RegisterSwaggerAPI(apiSvr.ClientCtx, apiSvr.Router, apiConfig.Swagger); err != nil {
		panic(err)
	}
}

// RegisterTxService implements the Application.RegisterTxService method.
func (app *App) RegisterTxService(clientCtx client.Context) {
	authtx.RegisterTxService(app.BaseApp.GRPCQueryRouter(), clientCtx, app.BaseApp.Simulate, app.interfaceRegistry)
}

// RegisterTendermintService implements the Application.RegisterTendermintService method.
func (app *App) RegisterTendermintService(clientCtx client.Context) {
	tmservice.RegisterTendermintService(
		clientCtx,
		app.BaseApp.GRPCQueryRouter(),
		app.interfaceRegistry,
		app.Query,
	)
}

// TxConfig returns App's TxConfig.
func (app *App) TxConfig() client.TxConfig {
	return app.txConfig
}

// Configurator get app configurator
func (app *App) Configurator() module.Configurator {
	return app.configurator
}

// ModuleManager returns the app ModuleManager
func (app *App) ModuleManager() *module.Manager {
	return app.mm
}

// initParamsKeeper init params keeper and its subspaces.
func initParamsKeeper(appCodec codec.BinaryCodec, legacyAmino *codec.LegacyAmino, key, tkey storetypes.StoreKey) paramskeeper.Keeper {
	paramsKeeper := paramskeeper.NewKeeper(appCodec, legacyAmino, key, tkey)

	paramsKeeper.Subspace(upgradetypes.ModuleName)
	/*
		paramsKeeper.Subspace(authtypes.ModuleName)
		paramsKeeper.Subspace(banktypes.ModuleName)
		paramsKeeper.Subspace(stakingtypes.ModuleName)
		paramsKeeper.Subspace(minttypes.ModuleName)
		paramsKeeper.Subspace(distrtypes.ModuleName)
		paramsKeeper.Subspace(slashingtypes.ModuleName)
		paramsKeeper.Subspace(govtypes.ModuleName).WithKeyTable(govtypes.ParamKeyTable())
		paramsKeeper.Subspace(crisistypes.ModuleName)
		paramsKeeper.Subspace(ibctransfertypes.ModuleName)
		paramsKeeper.Subspace(ibchost.ModuleName)
	*/
	paramsKeeper.Subspace(dclauthmoduletypes.ModuleName)
	paramsKeeper.Subspace(validatormoduletypes.ModuleName)
	paramsKeeper.Subspace(dclgenutilmoduletypes.ModuleName)
	paramsKeeper.Subspace(dclupgrademoduletypes.ModuleName)
	paramsKeeper.Subspace(dclpkitypes.ModuleName)
	paramsKeeper.Subspace(vendorinfomoduletypes.ModuleName)
	paramsKeeper.Subspace(dclmodeltypes.ModuleName)
	paramsKeeper.Subspace(dclcompliancetypes.ModuleName)
	// this line is used by starport scaffolding # stargate/app/paramSubspace

	return paramsKeeper
}
