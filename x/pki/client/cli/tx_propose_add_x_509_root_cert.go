package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/cli"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/common"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

var _ = strconv.Itoa(0)

func CmdProposeAddX509RootCert() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "propose-add-x509-root-cert",
		Short: "Proposes a new self-signed root certificate",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			cert, err := cli.ReadFromFile(viper.GetString(FlagCertificate))
			if err != nil {
				return err
			}

			vid := viper.GetInt32(FlagVid)
			info := viper.GetString(FlagInfo)
			certSchemaVersion := viper.GetUint32(common.FlagSchemaVersion)

			msg := types.NewMsgProposeAddX509RootCert(
				clientCtx.GetFromAddress().String(),
				cert,
				info,
				vid,
				certSchemaVersion,
			)
			// validate basic will be called in GenerateOrBroadcastTxCLI
			err = tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
			if cli.IsWriteInsteadReadRPCError(err) {
				return clientCtx.PrintString(cli.LightClientProxyForWriteRequests)
			}

			return err
		},
	}

	cmd.Flags().StringP(FlagCertificate, FlagCertificateShortcut, "",
		"PEM encoded certificate (string or path to file containing data)")
	cmd.Flags().String(FlagInfo, "", FlagInfoUsage)
	cmd.Flags().Int32(FlagVid, 0, "Model vendor ID (positive non-zero uint16)")
	cmd.Flags().Uint32(common.FlagSchemaVersion, 0, "Schema version of certificate")

	cli.AddTxFlagsToCmd(cmd)

	_ = cmd.MarkFlagRequired(flags.FlagFrom)
	_ = cmd.MarkFlagRequired(FlagCertificate)
	_ = cmd.MarkFlagRequired(FlagVid)

	return cmd
}
