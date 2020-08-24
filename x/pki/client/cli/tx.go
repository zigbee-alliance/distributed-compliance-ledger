package cli

import (
	"git.dsr-corporation.com/zb-ledger/zb-ledger/utils/cli"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/pki/internal/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func GetTxCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	complianceTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "PKI transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	complianceTxCmd.AddCommand(cli.SignedCommands(client.PostCommands(
		GetCmdProposeAddX509RootCertificate(cdc),
		GetCmdApproveAddX509RootCertificate(cdc),
		GetCmdAddX509Certificate(cdc),
		GetCmdRevokeX509Certificate(cdc),
	)...)...)

	return complianceTxCmd
}

func GetCmdProposeAddX509RootCertificate(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "propose-add-x509-root-cert",
		Short: "Proposes a new self-signed root certificate",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := cli.NewCLIContext().WithCodec(cdc)

			cert, err := cliCtx.ReadFromFile(viper.GetString(FlagCertificate))
			if err != nil {
				return err
			}

			msg := types.NewMsgProposeAddX509RootCert(cert, cliCtx.FromAddress())

			return cliCtx.HandleWriteMessage(msg)
		},
	}

	cmd.Flags().StringP(FlagCertificate, FlagCertificateShortcut, "",
		"PEM encoded certificate (string or path to file containing data)")

	_ = cmd.MarkFlagRequired(FlagCertificate)

	return cmd
}

func GetCmdApproveAddX509RootCertificate(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "approve-add-x509-root-cert",
		Short: "Approves the proposed root certificate correspondent to combination of subject and subject-key-id",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := cli.NewCLIContext().WithCodec(cdc)

			subject := viper.GetString(FlagSubject)
			subjectKeyID := viper.GetString(FlagSubjectKeyID)

			msg := types.NewMsgApproveAddX509RootCert(subject, subjectKeyID, cliCtx.FromAddress())

			return cliCtx.HandleWriteMessage(msg)
		},
	}

	cmd.Flags().StringP(FlagSubject, FlagSubjectShortcut, "", "Certificate's subject")
	cmd.Flags().StringP(FlagSubjectKeyID, FlagSubjectKeyIDShortcut, "",
		"Certificate's subject key id (hex)")

	_ = cmd.MarkFlagRequired(FlagSubject)
	_ = cmd.MarkFlagRequired(FlagSubjectKeyID)

	return cmd
}

func GetCmdAddX509Certificate(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use: "add-x509-cert",
		Short: "Adds an intermediate or leaf X509 certificate signed by a chain " +
			"of certificates which must be already present on the ledger",
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := cli.NewCLIContext().WithCodec(cdc)

			cert, err := cliCtx.ReadFromFile(viper.GetString(FlagCertificate))
			if err != nil {
				return err
			}

			msg := types.NewMsgAddX509Cert(cert, cliCtx.FromAddress())

			return cliCtx.HandleWriteMessage(msg)
		},
	}

	cmd.Flags().StringP(FlagCertificate, FlagCertificateShortcut, "",
		"PEM encoded certificate (string or path to file containing data)")

	_ = cmd.MarkFlagRequired(FlagCertificate)

	return cmd
}

func GetCmdRevokeX509Certificate(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use: "revoke-x509-cert",
		Short: "Revokes the given X509 certificate (either intermediate or leaf)." +
			"All the certificates in the subtree signed by the revoked certificate will be revoked as well.",
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := cli.NewCLIContext().WithCodec(cdc)

			subject := viper.GetString(FlagSubject)
			subjectKeyID := viper.GetString(FlagSubjectKeyID)

			msg := types.NewMsgRevokeX509Cert(subject, subjectKeyID, cliCtx.FromAddress())

			return cliCtx.HandleWriteMessage(msg)
		},
	}

	cmd.Flags().StringP(FlagSubject, FlagSubjectShortcut, "", "Certificate's subject")
	cmd.Flags().StringP(FlagSubjectKeyID, FlagSubjectKeyIDShortcut, "",
		"Certificate's subject key id (hex)")

	_ = cmd.MarkFlagRequired(FlagSubject)
	_ = cmd.MarkFlagRequired(FlagSubjectKeyID)

	return cmd
}
