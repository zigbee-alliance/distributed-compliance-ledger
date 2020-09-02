package cli

import (
	"fmt"
	"strings"

	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/auth"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/genutil"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/genutil/types"
	"github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/server"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tendermint/tendermint/libs/cli"
)

const (
	FlagAddress = "address"
	FlagPubKey  = "pubkey"
	FlagRoles   = "roles"
)

// AddGenesisAccountCmd returns add-genesis-account cobra Command.
//nolint:funlen
func AddGenesisAccountCmd(ctx *server.Context, cdc *codec.Codec,
	defaultNodeHome, defaultClientHome string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-genesis-account",
		Short: "Add genesis account to genesis.json",
		Args:  cobra.ExactArgs(0),
		RunE: func(_ *cobra.Command, args []string) error {
			config := ctx.Config
			config.SetRoot(viper.GetString(cli.HomeFlag))

			addr, err := sdk.AccAddressFromBech32(viper.GetString(FlagAddress))
			if err != nil {
				kb, err := keys.NewKeyBaseFromDir(viper.GetString(flagClientHome))
				if err != nil {
					return err
				}

				info, err := kb.Get(args[0])
				if err != nil {
					return err
				}

				addr = info.GetAddress()
			}

			pubkey, err := sdk.GetAccPubKeyBech32(viper.GetString(FlagPubKey))
			if err != nil {
				return err
			}

			var roles auth.AccountRoles
			if rolesStr := viper.GetString(FlagRoles); len(rolesStr) > 0 {
				for _, role := range strings.Split(rolesStr, ",") {
					roles = append(roles, auth.AccountRole(role))
				}
			}

			account := auth.NewAccount(addr, pubkey, roles)
			if err := account.Validate(); err != nil {
				return err
			}

			// retrieve the app state
			genFile := config.GenesisFile()
			appState, genDoc, err := genutil.GenesisStateFromGenFile(cdc, genFile)
			if err != nil {
				return err
			}

			// add genesis account to the app state
			var genesisState types.GenesisState

			cdc.MustUnmarshalJSON(appState[genutil.ModuleName], &genesisState)

			if genesisState.Accounts.Contains(addr) {
				return sdk.ErrUnknownRequest(fmt.Sprintf("cannot add account at existing address %v", addr))
			}

			genesisState.Accounts = append(genesisState.Accounts, account)

			genesisStateBz := cdc.MustMarshalJSON(genesisState)
			appState[genutil.ModuleName] = genesisStateBz

			appStateJSON, err := cdc.MarshalJSON(appState)
			if err != nil {
				return err
			}

			// export app state
			genDoc.AppState = appStateJSON

			return genutil.ExportGenesisFile(genDoc, genFile)
		},
	}

	cmd.Flags().String(FlagAddress, "", "Bench32 encoded account address")
	cmd.Flags().String(FlagPubKey, "", "Bench32 encoded account public key")
	cmd.Flags().String(FlagRoles, "",
		fmt.Sprintf("The list of roles (split by comma) to assign to account (supported roles: %v)", auth.Roles))
	cmd.Flags().String(cli.HomeFlag, defaultNodeHome, "node's home directory")
	cmd.Flags().String(flagClientHome, defaultClientHome, "client's home directory")

	_ = cmd.MarkFlagRequired(FlagAddress)
	_ = cmd.MarkFlagRequired(FlagPubKey)

	return cmd
}
