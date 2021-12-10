package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/model/types"
)

func CmdCreateModelVersion() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-model-version [vid] [pid] [software-version] [software-version-string] [cd-version-number] [firmware-digests] [software-version-valid] [ota-url] [ota-file-size] [ota-checksum] [ota-checksum-type] [min-applicable-software-version] [max-applicable-software-version] [release-notes-url]",
		Short: "Create a new ModelVersion",
		Args:  cobra.ExactArgs(14),
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
			indexSoftwareVersion, err := cast.ToUint64E(args[2])
			if err != nil {
				return err
			}

			// Get value arguments
			argSoftwareVersionString := args[3]
			argCdVersionNumber, err := cast.ToInt32E(args[4])
			if err != nil {
				return err
			}
			argFirmwareDigests := args[5]
			argSoftwareVersionValid, err := cast.ToBoolE(args[6])
			if err != nil {
				return err
			}
			argOtaUrl := args[7]
			argOtaFileSize, err := cast.ToUint64E(args[8])
			if err != nil {
				return err
			}
			argOtaChecksum := args[9]
			argOtaChecksumType, err := cast.ToInt32E(args[10])
			if err != nil {
				return err
			}
			argMinApplicableSoftwareVersion, err := cast.ToUint64E(args[11])
			if err != nil {
				return err
			}
			argMaxApplicableSoftwareVersion, err := cast.ToUint64E(args[12])
			if err != nil {
				return err
			}
			argReleaseNotesUrl := args[13]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateModelVersion(
				clientCtx.GetFromAddress().String(),
				indexVid,
				indexPid,
				indexSoftwareVersion,
				argSoftwareVersionString,
				argCdVersionNumber,
				argFirmwareDigests,
				argSoftwareVersionValid,
				argOtaUrl,
				argOtaFileSize,
				argOtaChecksum,
				argOtaChecksumType,
				argMinApplicableSoftwareVersion,
				argMaxApplicableSoftwareVersion,
				argReleaseNotesUrl,
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

func CmdUpdateModelVersion() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-model-version [vid] [pid] [software-version] [software-version-valid] [ota-url] [min-applicable-software-version] [max-applicable-software-version] [release-notes-url]",
		Short: "Update a ModelVersion",
		Args:  cobra.ExactArgs(14),
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
			indexSoftwareVersion, err := cast.ToUint64E(args[2])
			if err != nil {
				return err
			}

			// Get value arguments
			argSoftwareVersionValid, err := cast.ToBoolE(args[3])
			if err != nil {
				return err
			}
			argOtaUrl := args[4]
			argMinApplicableSoftwareVersion, err := cast.ToUint64E(args[5])
			if err != nil {
				return err
			}
			argMaxApplicableSoftwareVersion, err := cast.ToUint64E(args[6])
			if err != nil {
				return err
			}
			argReleaseNotesUrl := args[7]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateModelVersion(
				clientCtx.GetFromAddress().String(),
				indexVid,
				indexPid,
				indexSoftwareVersion,
				argSoftwareVersionValid,
				argOtaUrl,
				argMinApplicableSoftwareVersion,
				argMaxApplicableSoftwareVersion,
				argReleaseNotesUrl,
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

func CmdDeleteModelVersion() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-model-version [vid] [pid] [software-version]",
		Short: "Delete a ModelVersion",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			indexVid, err := cast.ToInt32E(args[0])
			if err != nil {
				return err
			}
			indexPid, err := cast.ToInt32E(args[1])
			if err != nil {
				return err
			}
			indexSoftwareVersion, err := cast.ToUint64E(args[2])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteModelVersion(
				clientCtx.GetFromAddress().String(),
				indexVid,
				indexPid,
				indexSoftwareVersion,
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
