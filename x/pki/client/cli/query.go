package cli

import (
	"encoding/json"
	"fmt"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/utils/cli"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/utils/pagination"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/pki/internal/keeper"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/pki/internal/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/viper"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"
)

const (
	FlagRootSubject      = "root-subject"
	FlagRootSubjectKeyId = "root-subject-key-id"
)

func GetQueryCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	complianceQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the compliancetest module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	complianceQueryCmd.AddCommand(client.GetCommands(
		GetCmdGetAllProposedX509RootCerts(storeKey, cdc),
		GetCmdProposedX509RootCert(storeKey, cdc),
		GetCmdGetAllX509RootCerts(storeKey, cdc),
		GetCmdX509Cert(storeKey, cdc),
		GetCmdGetAllX509Certs(storeKey, cdc),
		GetCmdGetAllSubjectX509Certs(storeKey, cdc),
	)...)

	return complianceQueryCmd
}

func GetCmdGetAllProposedX509RootCerts(queryRoute string, cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "all-proposed-x509-root-certs",
		Short: "Gets all proposed but not approved root certificates",
		RunE: func(cmd *cobra.Command, args []string) error {
			return getAllCertificates(cdc, fmt.Sprintf("custom/%s/all_proposed_x509_root_certs", queryRoute))
		},
	}

	cmd.Flags().Int(pagination.FlagSkip, 0, "amount of certificates to skip")
	cmd.Flags().Int(pagination.FlagTake, 0, "amount of certificates to take")

	return cmd
}

func GetCmdProposedX509RootCert(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "proposed-x509-root-cert [subject] [subject-key-id]",
		Short: "Gets a proposed but not approved root certificate with the given combination of subject and subject-key-id",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			subject := args[0]
			subjectKeyId := args[1]

			res, height, err := cliCtx.QueryStore([]byte(keeper.ProposedCertificateId(subject, subjectKeyId)), queryRoute)
			if err != nil || res == nil {
				return types.ErrProposedCertificateDoesNotExist(subject, subjectKeyId)
			}

			var proposedCertificate types.ProposedCertificate
			cdc.MustUnmarshalBinaryBare(res, &proposedCertificate)

			out, err := json.Marshal(proposedCertificate)
			if err != nil {
				return sdk.ErrInternal(fmt.Sprintf("Could not encode result: %v", err))
			}

			return cliCtx.PrintOutput(cli.NewReadResult(cdc, out, height))
		},
	}
}

func GetCmdGetAllX509RootCerts(queryRoute string, cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "all-x509-root-certs",
		Short: "Gets all approved root certificates",
		RunE: func(cmd *cobra.Command, args []string) error {
			return getAllCertificates(cdc, fmt.Sprintf("custom/%s/all_x509_root_certs", queryRoute))
		},
	}

	cmd.Flags().Int(pagination.FlagSkip, 0, "amount of certificates to skip")
	cmd.Flags().Int(pagination.FlagTake, 0, "amount of certificates to take")

	return cmd
}

func GetCmdX509Cert(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "x509-cert [subject] [subject-key-id]",
		Short: "Gets a certificates (either root, intermediate or leaf) by the given combination of subject and subject-key-id",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			subject := args[0]
			subjectKeyId := args[1]

			res, height, err := cliCtx.QueryStore([]byte(keeper.CertificateId(subject, subjectKeyId)), queryRoute)
			if err != nil || res == nil {
				return types.ErrCertificateDoesNotExist(subject, subjectKeyId)
			}

			var certificate types.Certificates
			cdc.MustUnmarshalBinaryBare(res, &certificate)

			out, err := json.Marshal(certificate)
			if err != nil {
				return sdk.ErrInternal(fmt.Sprintf("Could not encode result: %v", err))
			}

			return cliCtx.PrintOutput(cli.NewReadResult(cdc, out, height))
		},
	}
}

func GetCmdGetAllX509Certs(queryRoute string, cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "all-x509-certs",
		Short: "Gets all certificates (root, intermediate and leaf)",
		RunE: func(cmd *cobra.Command, args []string) error {
			return getAllCertificates(cdc, fmt.Sprintf("custom/%s/all_x509_certs", queryRoute))
		},
	}

	cmd.Flags().String(FlagRootSubject, "", "filter certificates by `Subject` of root certificate (only the certificates started with the given root certificate are returned)")
	cmd.Flags().String(FlagRootSubjectKeyId, "", "filter certificates by `Subject Key Id` of root certificate (only the certificates started with the given root certificate are returned)")
	cmd.Flags().Int(pagination.FlagSkip, 0, "amount of certificates to skip")
	cmd.Flags().Int(pagination.FlagTake, 0, "amount of certificates to take")

	return cmd
}

func GetCmdGetAllSubjectX509Certs(queryRoute string, cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "all-subject-x509-certs [subject]",
		Short: "Gets all certificates (root, intermediate and leaf) associated with subject",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			subject := args[0]
			return getAllCertificates(cdc, fmt.Sprintf("custom/%s/all_subject_x509_certs/%s", queryRoute, subject))
		},
	}

	cmd.Flags().String(FlagRootSubject, "", "filter certificates by `Subject` of root certificate (only the certificates started with the given root certificate are returned)")
	cmd.Flags().String(FlagRootSubjectKeyId, "", "filter certificates by `Subject Key Id` of root certificate (only the certificates started with the given root certificate are returned)")
	cmd.Flags().Int(pagination.FlagSkip, 0, "amount of certificates to skip")
	cmd.Flags().Int(pagination.FlagTake, 0, "amount of certificates to take")

	return cmd
}

func getAllCertificates(cdc *codec.Codec, route string) error {
	cliCtx := context.NewCLIContext().WithCodec(cdc)

	rootSubject := viper.GetString(FlagRootSubject)
	rootSubjectKeyId := viper.GetString(FlagRootSubjectKeyId)

	paginationParams := pagination.ParsePaginationParamsFromFlags()
	params := types.NewListCertificatesQueryParams(paginationParams, rootSubject, rootSubjectKeyId)

	res, height, err := cliCtx.QueryWithData(route, cdc.MustMarshalJSON(params))
	if err != nil {
		return sdk.ErrInternal(fmt.Sprintf("Could not query certificates: %s\n", err))
	}

	return cliCtx.PrintOutput(cli.NewReadResult(cdc, res, height))
}
