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
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/cli"
	commontypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/common/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
)

var _ = strconv.Itoa(0)

func CmdProposeAddAccount() *cobra.Command {
	cmd := &cobra.Command{
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

			var argVendorID int32
			if viper.GetString(FlagVID) != "" {
				argVendorID, err = cast.ToInt32E(viper.GetString(FlagVID))
				if err != nil {
					return err
				}
			}

			var pidRanges []*commontypes.Uint16Range
			if pidStrRanges := viper.GetString(FlagPIDs); len(pidStrRanges) > 0 {
				pidRanges, err = getPidRanges(pidStrRanges)
				if err != nil {
					return err
				}
			}

			argInfo := viper.GetString(FlagInfo)

			msg, err := types.NewMsgProposeAddAccount(
				clientCtx.GetFromAddress(),
				argAddress,
				argPubKey,
				argRoles,
				argVendorID,
				pidRanges,
				argInfo,
			)
			if err != nil {
				return err
			}

			// validate basic will be called in GenerateOrBroadcastTxCLI
			err = tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
			if cli.IsWriteInsteadReadRPCError(err) {
				return clientCtx.PrintString(cli.LightClientProxyForWriteRequests)
			}

			return err
		},
	}

	cmd.Flags().String(FlagAddress, "", "Bech32 encoded account address")
	cmd.Flags().String(FlagPubKey, "", "The account's Protobuf JSON encoded public key")
	cmd.Flags().String(FlagRoles, "",
		fmt.Sprintf("The list of roles, comma-separated, assigning to the account (supported roles: %v)",
			types.Roles))
	cmd.Flags().String(FlagVID, "", "Vendor ID associated with this account (positive non-zero uint16). Required only for Vendor Roles.")
	cmd.Flags().String(FlagPIDs, "", "Optional list of Product ID ranges (inclusive numbers split by \"-\") associated with this account (for example: 1-100,200-65535, positive non-zero uint16)")
	cmd.Flags().String(FlagInfo, "", FlagInfoUsage)

	cli.AddTxFlagsToCmd(cmd)

	_ = cmd.MarkFlagRequired(flags.FlagFrom)
	_ = cmd.MarkFlagRequired(FlagAddress)
	_ = cmd.MarkFlagRequired(FlagPubKey)

	return cmd
}

func getPidRanges(pidStrRanges string) ([]*commontypes.Uint16Range, error) {
	var pidRanges = make([]*commontypes.Uint16Range, 0)
	var lastMax int32
	for _, pidStrRange := range strings.Split(pidStrRanges, ",") {
		pidRange := strings.Split(pidStrRange, "-")
		if len(pidRange) != 2 {
			return nil, fmt.Errorf("failed to parse PID Range")
		}
		min, err := cast.ToInt32E(pidRange[0])
		if err != nil {
			return nil, fmt.Errorf("failed to parse PID Range: %w", err)
		}
		max, err := cast.ToInt32E(pidRange[1])
		if err != nil {
			return nil, fmt.Errorf("failed to parse PID Range: %w", err)
		}
		if min > max || max <= 0 || min <= 0 {
			return nil, fmt.Errorf("invalid PID Range is provided: min=%d, max=%d", min, max)
		}
		if max <= lastMax || min <= lastMax {
			return nil, fmt.Errorf("invalid PID Range is provided: {%d-%d}, ranges are overlapped, range items must be provided in increased order", min, max)
		}
		pid := commontypes.Uint16Range{Min: min, Max: max}
		pidRanges = append(pidRanges, &pid)
		lastMax = max
	}

	return pidRanges, nil
}
