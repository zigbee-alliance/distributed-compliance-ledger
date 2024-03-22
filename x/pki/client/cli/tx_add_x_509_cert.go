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

func CmdAddX509Cert() *cobra.Command {
	var (
		certSchemaVersion uint32
		schemaVersion     uint32
	)

	cmd := &cobra.Command{
		Use: "add-x509-cert",
		Short: "Adds an intermediate or leaf certificate signed by a chain " +
			"of certificates which must be already present on the ledger",
		Args: cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			cert, err := cli.ReadFromFile(viper.GetString(FlagCertificate))
			if err != nil {
				return err
			}

			msg := types.NewMsgAddX509Cert(
				clientCtx.GetFromAddress().String(),
				cert,
				certSchemaVersion,
				schemaVersion,
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
	cmd.Flags().Uint32Var(&certSchemaVersion, FlagCertificateSchemaVersion, 0, "Schema version of certificate")
	cmd.Flags().Uint32Var(&schemaVersion, common.FlagSchemaVersion, 0, "Schema version")

	cli.AddTxFlagsToCmd(cmd)

	_ = cmd.MarkFlagRequired(flags.FlagFrom)
	_ = cmd.MarkFlagRequired(FlagCertificate)

	return cmd
}
