package cli

import (
	"git.dsr-corporation.com/zb-ledger/zb-ledger/utils/cli"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/pki/internal/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
)

func GetTxCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	complianceTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "PKI transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	complianceTxCmd.AddCommand(client.PostCommands(
		GetCmdProposeAddX509RootCertificate(cdc),
		GetCmdApproveAddX509RootCertificate(cdc),
		GetCmdAddX509Certificate(cdc),
	)...)

	return complianceTxCmd
}

//nolint dupl
func GetCmdProposeAddX509RootCertificate(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "propose-add-x509-root-cert [certificate-path-or-pem-string]",
		Short: "Proposes a new self-signed root certificate",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := cli.NewCLIContext().WithCodec(cdc)

			cert, err := ReadCertificate(args[0])
			if err != nil {
				return err
			}

			msg := types.NewMsgProposeAddX509RootCert(cert, cliCtx.FromAddress())

			return cliCtx.HandleWriteMessage(msg)
		},
	}
}

//nolint dupl
func GetCmdApproveAddX509RootCertificate(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "approve-add-x509-root-cert [subject] [subject-key-id]",
		Short: "Approves the proposed root certificate correspondent to combination of subject and subject-key-id",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := cli.NewCLIContext().WithCodec(cdc)

			subject := args[0]
			subjectKeyId := args[1]

			msg := types.NewMsgApproveAddX509RootCert(subject, subjectKeyId, cliCtx.FromAddress())

			return cliCtx.HandleWriteMessage(msg)
		},
	}
}

//nolint dupl
func GetCmdAddX509Certificate(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "add-x509-cert [certificate-path-or-pem-string]",
		Short: "Adds an intermediate or leaf X509 certificate signed by a chain of certificates which must be already present on the ledger",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := cli.NewCLIContext().WithCodec(cdc)

			cert, err := ReadCertificate(args[0])
			if err != nil {
				return err
			}

			msg := types.NewMsgAddX509Cert(cert, cliCtx.FromAddress())

			return cliCtx.HandleWriteMessage(msg)
		},
	}
}

func ReadCertificate(cert string) (string, error) {
	if _, err := os.Stat(cert); err == nil { // check whether it is a path
		bytes, err := ioutil.ReadFile(cert)
		if err != nil {
			return "", err
		}
		return string(bytes), nil
	} else { // else return as is
		return cert, nil
	}
}
