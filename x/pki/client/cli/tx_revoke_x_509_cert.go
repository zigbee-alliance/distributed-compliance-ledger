package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/cli"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

var _ = strconv.Itoa(0)

func CmdRevokeX509Cert() *cobra.Command {
	cmd := &cobra.Command{
		Use: "revoke-x509-cert",
		Short: "Revokes the given intermediate or leaf certificate. " +
			"If revoke-child flag is set to true then all the certificates in the subtree signed by the revoked " +
			"certificate will be revoked as well.",
		Args: cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			subject := viper.GetString(FlagSubject)
			subjectKeyID := viper.GetString(FlagSubjectKeyID)
			serialNumber := viper.GetString(FlagSerialNumber)
			revokeChild := viper.GetBool(FlagRevokeChild)
			infoArg := viper.GetString(FlagInfo)

			msg := types.NewMsgRevokeX509Cert(
				clientCtx.GetFromAddress().String(),
				subject,
				subjectKeyID,
				serialNumber,
				revokeChild,
				infoArg,
			)
			// validate basic will be called in GenerateOrBroadcastTxCLI
			err = tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
			if cli.IsWriteInsteadReadRPCError(err) {
				return clientCtx.PrintString(cli.LightClientProxyForWriteRequests)
			}

			return err
		},
	}

	cmd.Flags().StringP(FlagSubject, FlagSubjectShortcut, "", "Certificate's subject")
	cmd.Flags().StringP(FlagSubjectKeyID, FlagSubjectKeyIDShortcut, "", "Certificate's subject key id (hex)")
	cmd.Flags().StringP(FlagSerialNumber, FlagSerialNumberShortcut, "", "Certificate's serial number")
	cmd.Flags().StringP(FlagRevokeChild, FlagRevokeChildShortcut, "", "If flag is true then all the certificates in the subtree will be revoked as well - default is false")
	cmd.Flags().String(FlagInfo, "", FlagInfoUsage)
	cli.AddTxFlagsToCmd(cmd)

	_ = cmd.MarkFlagRequired(FlagSubject)
	_ = cmd.MarkFlagRequired(FlagSubjectKeyID)
	_ = cmd.MarkFlagRequired(flags.FlagFrom)

	return cmd
}
