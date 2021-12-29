package cli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
)

var _ = strconv.Itoa(0)

func CmdProposeAddAccount() *cobra.Command {
	cmd := &cobra.Command{
		// TODO issue 99: move from required flags to positional arguments
		//		(to keep the same direction as in cosmos sdk https://github.com/cosmos/cosmos-sdk/issues/1956)
		Use:   "propose-add-account",
		Short: "Broadcast message ProposeAddAccount",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argAddress, err := sdk.AccAddressFromBech32(viper.GetString(FlagAddress))
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			var argPubKey cryptotypes.PubKey
			if err := clientCtx.Codec.UnmarshalInterfaceJSON(
				[]byte(viper.GetString(FlagPubKey)),
				&argPubKey,
			); err != nil {
				return err
			}

			var argRoles types.AccountRoles
			if rolesStr := viper.GetString(FlagRoles); len(rolesStr) > 0 {
				for _, role := range strings.Split(rolesStr, ",") {
					argRoles = append(argRoles, types.AccountRole(role))
				}
			}

			var argVendorID uint64
			if viper.GetString(FlagVID) != "" {
				argVendorID, err = cast.ToUint64E(viper.GetString(FlagVID))
				if err != nil {
					return err
				}
			}

			msg, err := types.NewMsgProposeAddAccount(
				clientCtx.GetFromAddress(),
				argAddress,
				argPubKey,
				argRoles,
				argVendorID,
			)
			if err != nil {
				return err
			}

			/* it is done inside next step
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			*/
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(FlagAddress, "", "Bench32 encoded account address")
	cmd.Flags().String(FlagPubKey, "", "The validator's Protobuf JSON encoded public key")
	cmd.Flags().String(FlagRoles, "",
		fmt.Sprintf("The list of roles, comma-separated, assigning to the account (supported roles: %v)",
			types.Roles))
	cmd.Flags().String(FlagVID, "", "Vendor ID associated with this account. Required only for Vendor Roles")

	flags.AddTxFlagsToCmd(cmd)

	_ = cmd.MarkFlagRequired(flags.FlagFrom) // XXX issue 99: was absent in legacy code ???
	_ = cmd.MarkFlagRequired(FlagAddress)
	_ = cmd.MarkFlagRequired(FlagPubKey)

	return cmd
}
