package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"

	// "github.com/cosmos/cosmos-sdk/crypto/keyring".
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/server"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	dclauthtypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/dclgenutil"
	dclgenutiltypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclgenutil/types"
)

const (
	FlagAddress = "address"
	FlagPubKey  = "pubkey"
	FlagRoles   = "roles"
	FlagVID     = "vid"
)

// AddGenesisAccountCmd returns add-genesis-account cobra Command.
func AddGenesisAccountCmd(defaultNodeHome string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-genesis-account ",
		Short: "Add a genesis account to genesis.json",
		Long: `Add a genesis account to genesis.json. The provided account must specify
the account address or key name. If a key name is given,
the address will be looked up in the local Keybase.
`,
		Args: cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			//nolint:staticcheck
			depCdc := clientCtx.JSONCodec
			cdc := depCdc.(codec.Codec)

			serverCtx := server.GetServerContextFromCmd(cmd)
			config := serverCtx.Config

			config.SetRoot(clientCtx.HomeDir)

			addr, err := sdk.AccAddressFromBech32(viper.GetString(FlagAddress))
			// if err != nil {
			//	return err
			//}
			// TODO migration of keyring was not released yet in cosmos (in v.0.44.4)
			if err != nil {
				inBuf := bufio.NewReader(cmd.InOrStdin())
				keyringBackend, err := cmd.Flags().GetString(flags.FlagKeyringBackend)
				if err != nil {
					return err
				}

				// attempt to lookup address from Keybase if no address was provided
				kb, err := keyring.New(sdk.KeyringServiceName(), keyringBackend, clientCtx.HomeDir, inBuf)
				if err != nil {
					return err
				}

				info, err := kb.Key(args[0])
				if err != nil {
					return fmt.Errorf("failed to get address from Keybase: %w", err)
				}

				addr = info.GetAddress()
			}

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

			var vendorID int32 = 0
			if viper.GetString(FlagVID) != "" {
				vendorID, err = cast.ToInt32E(viper.GetString(FlagVID))
				if err != nil {
					return err
				}
			}

			// FIXME issue 99 VendorID
			genAccount = dclauthtypes.NewAccount(ba, roles, []*dclauthtypes.Grant{}, vendorID)

			if err := genAccount.Validate(); err != nil {
				return fmt.Errorf("failed to validate new genesis account: %w", err)
			}

			genFile := config.GenesisFile()
			appState, genDoc, err := dclgenutiltypes.GenesisStateFromGenFile(genFile)
			if err != nil {
				return fmt.Errorf("failed to unmarshal genesis state: %w", err)
			}

			authGenState := dclauthtypes.GetGenesisStateFromAppState(cdc, appState)
			accs := authGenState.AccountList

			// Add the new account to the set of genesis accounts and sanitize the
			// accounts afterwards.
			accs = append(accs, *genAccount)

			authGenState.AccountList = accs

			authGenStateBz, err := cdc.MarshalJSON(authGenState)
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
	cmd.Flags().String(FlagVID, "", "Vendor ID associated with this account. Required only for Vendor Roles")

	cmd.Flags().String(flags.FlagHome, defaultNodeHome, "The application home directory")
	cmd.Flags().String(flags.FlagKeyringBackend, flags.DefaultKeyringBackend, "Select keyring's backend (os|file|kwallet|pass|test)")

	/*
		cmd.Flags().String(flags.FlagKeyringBackend, flags.DefaultKeyringBackend, "Select keyring's backend (os|file|kwallet|pass|test)")
		cmd.Flags().String(flags.FlagHome, defaultNodeHome, "The application home directory")
		cmd.Flags().String(flagVestingAmt, "", "amount of coins for vesting accounts")
		cmd.Flags().Int64(flagVestingStart, 0, "schedule start time (unix epoch) for vesting accounts")
		flags.AddQueryFlagsToCmd(cmd)


	*/

	_ = cmd.MarkFlagRequired(FlagAddress)
	_ = cmd.MarkFlagRequired(FlagPubKey)

	return cmd
}
