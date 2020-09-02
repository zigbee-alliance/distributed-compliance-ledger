package main

import (
	"encoding/json"
	"io"

	app "git.dsr-corporation.com/zb-ledger/zb-ledger"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/cmd/settings"
	genutilcli "git.dsr-corporation.com/zb-ledger/zb-ledger/x/genutil/client/cli"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/server"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/cli"
	"github.com/tendermint/tendermint/libs/log"
	tmtypes "github.com/tendermint/tendermint/types"
	dbm "github.com/tendermint/tm-db"
)

func main() {
	cobra.EnableCommandSorting = false

	cdc := app.MakeCodec()

	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount(sdk.Bech32PrefixAccAddr, sdk.Bech32PrefixAccPub)
	config.SetBech32PrefixForValidator(sdk.Bech32PrefixValAddr, sdk.Bech32PrefixValPub)
	config.SetBech32PrefixForConsensusNode(sdk.Bech32PrefixConsAddr, sdk.Bech32PrefixConsPub)
	config.Seal()

	ctx := server.NewDefaultContext()

	rootCmd := &cobra.Command{
		Use:               "zbld",
		Short:             "ZbLedger App Daemon (server)",
		PersistentPreRunE: server.PersistentPreRunEFn(ctx),
	}
	// CLI commands to initialize the chain
	rootCmd.AddCommand(
		genutilcli.InitCmd(ctx, cdc, app.ModuleBasics, app.DefaultNodeHome),
		genutilcli.CollectGenTxsCmd(ctx, cdc, app.DefaultNodeHome),
		genutilcli.GenTxCmd(
			ctx, cdc, app.ModuleBasics, app.DefaultNodeHome, app.DefaultCLIHome,
		),
		genutilcli.ValidateGenesisCmd(ctx, cdc, app.ModuleBasics),
		// AddGenesisAccountCmd allows users to add accounts to the genesis file
		genutilcli.AddGenesisAccountCmd(ctx, cdc, app.DefaultNodeHome, app.DefaultCLIHome),
	)

	server.AddCommands(ctx, cdc, rootCmd, newApp, exportAppStateAndTMValidators)

	// prepare and add flags
	executor := cli.PrepareBaseCmd(rootCmd, "NS", app.DefaultNodeHome)

	err := executor.Execute()
	if err != nil {
		panic(err)
	}
}

func newApp(logger log.Logger, db dbm.DB, traceStore io.Writer) abci.Application {
	return app.NewZbLedgerApp(logger, db, baseapp.SetPruning(settings.PruningStrategy))
}

func exportAppStateAndTMValidators(logger log.Logger, db dbm.DB, traceStore io.Writer,
	height int64, forZeroHeight bool, jailWhiteList []string) (json.RawMessage, []tmtypes.GenesisValidator, error) {
	if height != -1 {
		nsApp := app.NewZbLedgerApp(logger, db, baseapp.SetPruning(settings.PruningStrategy))

		err := nsApp.LoadHeight(height)
		if err != nil {
			return nil, nil, err
		}

		return nsApp.ExportAppStateAndValidators(forZeroHeight, jailWhiteList)
	}

	nsApp := app.NewZbLedgerApp(logger, db, baseapp.SetPruning(settings.PruningStrategy))

	return nsApp.ExportAppStateAndValidators(forZeroHeight, jailWhiteList)
}
