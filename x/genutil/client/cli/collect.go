package cli

import (
	"encoding/json"
	"path/filepath"

	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/genutil"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tendermint/tendermint/libs/cli"
	tmtypes "github.com/tendermint/tendermint/types"
)

const flagGenTxDir = "gentx-dir"

// CollectGenTxsCmd - return the cobra command to collect genesis transactions.
func CollectGenTxsCmd(ctx *server.Context, cdc *codec.Codec, defaultNodeHome string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "collect-gentxs",
		Short: "Collect genesis txs and output a genesis.json file",
		RunE: func(_ *cobra.Command, _ []string) error {
			config := ctx.Config
			config.SetRoot(viper.GetString(cli.HomeFlag))

			nodeID, valPubKey, err := genutil.InitializeNodeValidatorFiles(config)
			if err != nil {
				return err
			}

			name := viper.GetString(client.FlagName)

			genDoc, err := tmtypes.GenesisDocFromFile(config.GenesisFile())
			if err != nil {
				return err
			}

			genTxsDir := viper.GetString(flagGenTxDir)
			if genTxsDir == "" {
				genTxsDir = filepath.Join(config.RootDir, "config", "gentx")
			}

			toPrint := newPrintInfo(config.Moniker, genDoc.ChainID, nodeID, genTxsDir, json.RawMessage(""))
			initCfg := genutil.NewInitConfig(genDoc.ChainID, genTxsDir, name, nodeID, valPubKey)

			appMessage, err := genutil.GenAppStateFromConfig(cdc, config, initCfg, *genDoc)
			if err != nil {
				return err
			}

			toPrint.AppMessage = appMessage

			// print out some key information
			return displayInfo(cdc, toPrint)
		},
	}

	cmd.Flags().String(cli.HomeFlag, defaultNodeHome, "node's home directory")
	cmd.Flags().String(flagGenTxDir, "",
		"override default \"gentx\" directory from which collect and execute "+
			"genesis transactions; default [--home]/config/gentx/")

	return cmd
}
