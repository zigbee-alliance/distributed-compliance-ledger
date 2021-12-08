package cli

import (
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/model/types"
)

func CmdCreateModelVersions() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-model-versions [vid] [pid] [software-versions]",
		Short: "Create a new ModelVersions",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexVid, err := cast.ToInt32E(args[0])
			if err != nil {
				return err
			}
			indexPid, err := cast.ToInt32E(args[1])
			if err != nil {
				return err
			}

			// Get value arguments
			argCastSoftwareVersions := strings.Split(args[2], ",")
			argSoftwareVersions := make([]uint64, len(argCastSoftwareVersions))
			for i, arg := range argCastSoftwareVersions {
				value, err := cast.ToUint64E(arg)
				if err != nil {
					return err
				}
				argSoftwareVersions[i] = value
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateModelVersions(
				clientCtx.GetFromAddress().String(),
				indexVid,
				indexPid,
				argSoftwareVersions,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdateModelVersions() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-model-versions [vid] [pid] [software-versions]",
		Short: "Update a ModelVersions",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexVid, err := cast.ToInt32E(args[0])
			if err != nil {
				return err
			}
			indexPid, err := cast.ToInt32E(args[1])
			if err != nil {
				return err
			}

			// Get value arguments
			argCastSoftwareVersions := strings.Split(args[2], ",")
			argSoftwareVersions := make([]uint64, len(argCastSoftwareVersions))
			for i, arg := range argCastSoftwareVersions {
				value, err := cast.ToUint64E(arg)
				if err != nil {
					return err
				}
				argSoftwareVersions[i] = value
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateModelVersions(
				clientCtx.GetFromAddress().String(),
				indexVid,
				indexPid,
				argSoftwareVersions,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeleteModelVersions() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-model-versions [vid] [pid]",
		Short: "Delete a ModelVersions",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			indexVid, err := cast.ToInt32E(args[0])
			if err != nil {
				return err
			}
			indexPid, err := cast.ToInt32E(args[1])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteModelVersions(
				clientCtx.GetFromAddress().String(),
				indexVid,
				indexPid,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
