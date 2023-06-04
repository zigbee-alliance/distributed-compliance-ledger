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

func CmdAddPkiRevocationDistributionPoint() *cobra.Command {
	var (
		vid                  int32
		pid                  int32
		isPAA                bool
		label                string
		crlSignerCertificate string
		issuerSubjectKeyID   string
		dataUrl              string
		dataFileSize         uint64
		dataDigest           string
		dataDigestType       uint32
		revocationType       uint32
	)

	cmd := &cobra.Command{
		Use:   "add-revocation-point",
		Short: "Publishes a PKI Revocation distribution endpoint owned by the Vendor",
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

			msg := types.NewMsgAddPkiRevocationDistributionPoint(
				clientCtx.GetFromAddress().String(),
				vid,
				pid,
				isPAA,
				label,
				cert,
				issuerSubjectKeyID,
				dataUrl,
				dataFileSize,
				dataDigest,
				dataDigestType,
				revocationType,
			)

			// validate basic will be called in GenerateOrBroadcastTxCLI
			err = tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
			if cli.IsWriteInsteadReadRPCError(err) {
				return clientCtx.PrintString(cli.LightClientProxyForWriteRequests)
			}

			return err
		},
	}

	cmd.Flags().Int32Var(&vid, FlagVid, 0,
		"Vendor ID (positive non-zero). Must be the same as Vendor account's VID and vid field in the VID-scoped CRLSignerCertificate")
	cmd.Flags().Int32Var(&pid, FlagPid, 0,
		"Product ID (positive non-zero). Must be empty if IsPAA is true. Must be equal to a pid field in CRLSignerCertificate")
	cmd.Flags().BoolVar(&isPAA, FlagIsPAA, true, "True if the revocation information distribution point relates to a PAA (Root certificate)")
	cmd.Flags().StringVarP(&label, FlagLabel, FlagLabelShortcut, "", " A label to disambiguate multiple revocation information partitions of a particular issuer")
	cmd.Flags().StringVarP(&crlSignerCertificate, FlagCertificate, FlagCertificateShortcut, "", "The issuer certificate whose revocation information is provided in the distribution point entry, encoded in X.509v3 PEM format. The corresponding CLI parameter can contain either a PEM string or a path to a file containing the data")
	cmd.Flags().StringVar(&issuerSubjectKeyID, FlagIssuerSubjectKeyID, "", "Uniquely identifies the PAA or PAI for which this revocation distribution point is provided. Must consist of even number of uppercase hexadecimal characters ([0-9A-F]), with no whitespace and no non-hexadecimal characters., e.g: 5A880E6C3653D07FB08971A3F473790930E62BDB")
	cmd.Flags().StringVar(&dataUrl, FlagDataUrl, "", "The URL where to obtain the information in the format indicated by the RevocationType field. Must start with either http or https")
	cmd.Flags().Uint64Var(&dataFileSize, FlagDataFileSize, 0, "Total size in bytes of the file found at the DataUrl. Must be omitted if RevocationType is 1")
	cmd.Flags().StringVar(&dataDigest, FlagDataDigest, "", "Digest of the entire contents of the associated file downloaded from the DataUrl. Must be omitted if RevocationType is 1. Must be provided if and only if the DataFileSize field is present")
	cmd.Flags().Uint32Var(&dataDigestType, FlagDataDigestType, 0, "The type of digest used in the DataDigest field from the list of [1, 7, 8, 10, 11, 12] (IANA Named Information Hash Algorithm Registry). Must be provided if and only if the DataDigest field is present") //TODO: will give error if omitted
	cmd.Flags().Uint32Var(&revocationType, FlagRevocationType, 1, "The type of file found at the DataUrl for this entry. Supported types: 1 - RFC5280 Certificate Revocation List (CRL)")
	flags.AddTxFlagsToCmd(cmd)

	_ = cmd.MarkFlagRequired(FlagVid)
	_ = cmd.MarkFlagRequired(FlagIsPAA)
	_ = cmd.MarkFlagRequired(FlagLabel)
	_ = cmd.MarkFlagRequired(FlagCertificate)
	_ = cmd.MarkFlagRequired(FlagIssuerSubjectKeyID)
	_ = cmd.MarkFlagRequired(FlagDataUrl)
	_ = cmd.MarkFlagRequired(FlagRevocationType)
	_ = cmd.MarkFlagRequired(flags.FlagFrom)

	return cmd
}
