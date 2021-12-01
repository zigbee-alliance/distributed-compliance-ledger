package cmd

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"

	// "github.com/cosmos/cosmos-sdk/crypto/keyring"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/server"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	dclauthtypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/dclgenutil"
	dclgenutiltypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclgenutil/types"
)

const (
	FlagAddress = "address"
	FlagPubKey  = "pubkey"
	FlagRoles   = "roles"
)

// AddGenesisAccountCmd returns add-genesis-account cobra Command.
func AddGenesisAccountCmd(defaultNodeHome string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-genesis-account",
		Short: "Add a genesis account to genesis.json",
		Long: `Add a genesis account to genesis.json. The provided account must specify
the account address or key name. If a key name is given,
the address will be looked up in the local Keybase.
`,
		Args: cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			serverCtx := server.GetServerContextFromCmd(cmd)
			config := serverCtx.Config

			config.SetRoot(clientCtx.HomeDir)

			addr, err := sdk.AccAddressFromBech32(viper.GetString(FlagAddress))
			if err != nil {
				return err
			}

			/* TODO migration of keyring was not released yet in cosmos (in v.0.44.4)
			var kr keyring.Keyring
			if err != nil {
				inBuf := bufio.NewReader(cmd.InOrStdin())
				keyringBackend, _ := cmd.Flags().GetString(flags.FlagKeyringBackend)

				if keyringBackend != "" && clientCtx.Keyring == nil {
					var err error
					kr, err = keyring.New(sdk.KeyringServiceName(), keyringBackend, clientCtx.HomeDir, inBuf, clientCtx.Codec)
					if err != nil {
						return err
					}
				} else {
					kr = clientCtx.Keyring
				}

				k, err := kr.Key(viper.GetString(FlagAddress))
				if err != nil {
					return fmt.Errorf("failed to get address from Keyring: %w", err)
				}

				addr, err = k.GetAddress()
				if err != nil {
					return err
				}
			}
			*/

			pkStr := viper.GetString(FlagPubKey)

			var pk cryptotypes.PubKey
			if err := clientCtx.Codec.UnmarshalInterfaceJSON([]byte(pkStr), &pk); err != nil {
				return err
			}

			var roles dclauthtypes.AccountRoles
			if rolesStr := viper.GetString(FlagRoles); len(rolesStr) > 0 {
				for _, role := range strings.Split(rolesStr, ",") {
					roles = append(roles, dclauthtypes.AccountRole(role))
				}
			}

			// create concrete account type based on input parameters
			var genAccount *dclauthtypes.Account
			// TODO issue 99: review - cosmos here works with GenesisAccount interfaces
			//	    and uses pack/unpack API to extract accounts from genesis state

			ba := authtypes.NewBaseAccount(addr, pk, 0, 0)
			// FIXME issue 99 VendorID
			genAccount = dclauthtypes.NewAccount(ba, roles, 0)

			if err := genAccount.Validate(); err != nil {
				return fmt.Errorf("failed to validate new genesis account: %w", err)
			}

			genFile := config.GenesisFile()
			appState, genDoc, err := dclgenutiltypes.GenesisStateFromGenFile(genFile)
			if err != nil {
				return fmt.Errorf("failed to unmarshal genesis state: %w", err)
			}

			authGenState := dclauthtypes.GetGenesisStateFromAppState(clientCtx.Codec, appState)
			accs := authGenState.AccountList

			// Add the new account to the set of genesis accounts and sanitize the
			// accounts afterwards.
			accs = append(accs, *genAccount)

			authGenState.AccountList = accs

			authGenStateBz, err := clientCtx.Codec.MarshalJSON(authGenState)
			if err != nil {
				return fmt.Errorf("failed to marshal auth genesis state: %w", err)
			}

			appState[dclauthtypes.ModuleName] = authGenStateBz

			appStateJSON, err := json.Marshal(appState)
			if err != nil {
				return fmt.Errorf("failed to marshal application genesis state: %w", err)
			}

			genDoc.AppState = appStateJSON
			return dclgenutil.ExportGenesisFile(genDoc, genFile)
		},
	}

	cmd.Flags().String(FlagAddress, "", "Bench32 encoded account address or key name")
	cmd.Flags().String(FlagPubKey, "", "The validator's Protobuf JSON encoded public key")
	cmd.Flags().String(FlagRoles, "",
		fmt.Sprintf("The list of roles (split by comma) to assign to account (supported roles: %v)", dclauthtypes.Roles))

	cmd.Flags().String(flags.FlagHome, defaultNodeHome, "The application home directory")
	cmd.Flags().String(flags.FlagKeyringBackend, flags.DefaultKeyringBackend, "Select keyring's backend (os|file|kwallet|pass|test)")

	_ = cmd.MarkFlagRequired(FlagAddress)
	_ = cmd.MarkFlagRequired(FlagPubKey)

	return cmd
}
