// Copyright 2020 DSR Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cli

import (
	"fmt"

	"git.dsr-corporation.com/zb-ledger/zb-ledger/utils/cli"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/utils/pagination"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/pki/internal/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func GetQueryCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	complianceQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the pki module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	complianceQueryCmd.AddCommand(client.GetCommands(
		GetCmdGetAllProposedX509RootCerts(storeKey, cdc),
		GetCmdGetProposedX509RootCert(storeKey, cdc),
		GetCmdGetAllX509RootCerts(storeKey, cdc),
		GetCmdGetX509Cert(storeKey, cdc),
		GetCmdGetX509CertChain(storeKey, cdc),
		GetCmdGetAllX509Certs(storeKey, cdc),
		GetCmdGetAllSubjectX509Certs(storeKey, cdc),
		GetCmdGetAllProposedX509RootCertsToRevoke(storeKey, cdc),
		GetCmdGetProposedX509RootCertToRevoke(storeKey, cdc),
		GetCmdGetRevokedX509Cert(storeKey, cdc),
		GetCmdGetAllRevokedX509RootCerts(storeKey, cdc),
		GetCmdGetAllRevokedX509Certs(storeKey, cdc),
	)...)

	return complianceQueryCmd
}

func GetCmdGetAllProposedX509RootCerts(queryRoute string, cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "all-proposed-x509-root-certs",
		Short: "Gets all proposed but not approved root certificates",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return performPkiQuery(cdc, fmt.Sprintf("custom/%s/all_proposed_x509_root_certs", queryRoute))
		},
	}

	cmd.Flags().Int(pagination.FlagSkip, 0, "amount of certificates to skip")
	cmd.Flags().Int(pagination.FlagTake, 0, "amount of certificates to take")

	return cmd
}

func GetCmdGetProposedX509RootCert(queryRoute string, cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "proposed-x509-root-cert",
		Short: "Gets a proposed but not approved root certificate with the given combination of subject and subject-key-id",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := cli.NewCLIContext().WithCodec(cdc)

			subject := viper.GetString(FlagSubject)
			subjectKeyID := viper.GetString(FlagSubjectKeyID)

			res, height, err := cliCtx.QueryStore(types.GetProposedCertificateKey(subject, subjectKeyID), queryRoute)
			if err != nil || res == nil {
				return types.ErrProposedCertificateDoesNotExist(subject, subjectKeyID)
			}

			var proposedCertificate types.ProposedCertificate
			cdc.MustUnmarshalBinaryBare(res, &proposedCertificate)

			return cliCtx.EncodeAndPrintWithHeight(proposedCertificate, height)
		},
	}

	cmd.Flags().StringP(FlagSubject, FlagSubjectShortcut, "", "Certificate's subject")
	cmd.Flags().StringP(FlagSubjectKeyID, FlagSubjectKeyIDShortcut, "", "Certificate's subject key id (hex)")

	_ = cmd.MarkFlagRequired(FlagSubject)
	_ = cmd.MarkFlagRequired(FlagSubjectKeyID)

	return cmd
}

func GetCmdGetAllX509RootCerts(queryRoute string, cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "all-x509-root-certs",
		Short: "Gets all approved root certificates",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return performPkiQuery(cdc, fmt.Sprintf("custom/%s/all_x509_root_certs", queryRoute))
		},
	}

	cmd.Flags().Int(pagination.FlagSkip, 0, "amount of certificates to skip")
	cmd.Flags().Int(pagination.FlagTake, 0, "amount of certificates to take")

	return cmd
}

// nolint:dupl
func GetCmdGetX509Cert(queryRoute string, cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use: "x509-cert",
		Short: "Gets certificates (either root, intermediate or leaf) " +
			"by the given combination of subject and subject-key-id",
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := cli.NewCLIContext().WithCodec(cdc)

			subject := viper.GetString(FlagSubject)
			subjectKeyID := viper.GetString(FlagSubjectKeyID)

			res, height, err := cliCtx.QueryStore(types.GetApprovedCertificateKey(subject, subjectKeyID), queryRoute)
			if err != nil || res == nil {
				return types.ErrCertificateDoesNotExist(subject, subjectKeyID)
			}

			var certificates types.Certificates
			cdc.MustUnmarshalBinaryBare(res, &certificates)

			return cliCtx.EncodeAndPrintWithHeight(certificates, height)
		},
	}

	cmd.Flags().StringP(FlagSubject, FlagSubjectShortcut, "", "Certificate's subject")
	cmd.Flags().StringP(FlagSubjectKeyID, FlagSubjectKeyIDShortcut, "", "Certificate's subject key id (hex)")

	_ = cmd.MarkFlagRequired(FlagSubject)
	_ = cmd.MarkFlagRequired(FlagSubjectKeyID)

	return cmd
}

func GetCmdGetX509CertChain(queryRoute string, cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use: "x509-cert-chain",
		Short: "Gets the complete chain for a certificate with " +
			"the given combination of subject and subject-key-id",
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := cli.NewCLIContext().WithCodec(cdc)

			subject := viper.GetString(FlagSubject)
			subjectKeyID := viper.GetString(FlagSubjectKeyID)

			chain := types.NewCertificates([]types.Certificate{})

			height, err := chainCertificates(cliCtx, queryRoute, subject, subjectKeyID, &chain)
			if err != nil {
				return types.ErrCertificateDoesNotExist(subject, subjectKeyID)
			}

			return cliCtx.EncodeAndPrintWithHeight(chain, height)
		},
	}

	cmd.Flags().StringP(FlagSubject, FlagSubjectShortcut, "", "Certificate's subject")
	cmd.Flags().StringP(FlagSubjectKeyID, FlagSubjectKeyIDShortcut, "", "Certificate's subject key id (hex)")

	_ = cmd.MarkFlagRequired(FlagSubject)
	_ = cmd.MarkFlagRequired(FlagSubjectKeyID)

	return cmd
}

func GetCmdGetAllX509Certs(queryRoute string, cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "all-x509-certs",
		Short: "Gets all certificates (root, intermediate and leaf)",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return performPkiQuery(cdc, fmt.Sprintf("custom/%s/all_x509_certs", queryRoute))
		},
	}

	cmd.Flags().StringP(FlagRootSubject, FlagRootSubjectShortcut, "",
		"filter certificates by `Subject` of root certificate "+
			"(only the certificates originated from the given root certificate are returned)")
	cmd.Flags().StringP(FlagRootSubjectKeyID, FlagRootSubjectKeyIDShortcut, "",
		"filter certificates by `Subject Key Id` of root certificate "+
			"(only the certificates originated from the given root certificate are returned)")
	cmd.Flags().Int(pagination.FlagSkip, 0, "amount of certificates to skip")
	cmd.Flags().Int(pagination.FlagTake, 0, "amount of certificates to take")

	return cmd
}

func GetCmdGetAllSubjectX509Certs(queryRoute string, cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "all-subject-x509-certs",
		Short: "Gets all certificates (root, intermediate and leaf) associated with subject",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			subject := viper.GetString(FlagSubject)

			return performPkiQuery(cdc, fmt.Sprintf("custom/%s/all_subject_x509_certs/%s", queryRoute, subject))
		},
	}

	cmd.Flags().StringP(FlagSubject, FlagSubjectShortcut, "", "Certificate's subject")
	cmd.Flags().StringP(FlagRootSubject, FlagRootSubjectShortcut, "",
		"filter certificates by `Subject` of root certificate "+
			"(only the certificates originated from the given root certificate are returned)")
	cmd.Flags().StringP(FlagRootSubjectKeyID, FlagRootSubjectKeyIDShortcut, "",
		"filter certificates by `Subject Key Id` of root certificate "+
			"(only the certificates originated from the given root certificate are returned)")
	cmd.Flags().Int(pagination.FlagSkip, 0, "amount of certificates to skip")
	cmd.Flags().Int(pagination.FlagTake, 0, "amount of certificates to take")

	_ = cmd.MarkFlagRequired(FlagSubject)

	return cmd
}

func GetCmdGetAllProposedX509RootCertsToRevoke(queryRoute string, cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "all-proposed-x509-root-certs-to-revoke",
		Short: "Gets all proposed but not approved root certificates to be revoked",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return performPkiQuery(cdc, fmt.Sprintf("custom/%s/all_proposed_x509_root_cert_revocations", queryRoute))
		},
	}

	cmd.Flags().Int(pagination.FlagSkip, 0, "amount of certificates to skip")
	cmd.Flags().Int(pagination.FlagTake, 0, "amount of certificates to take")

	return cmd
}

// nolint:dupl
func GetCmdGetProposedX509RootCertToRevoke(queryRoute string, cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use: "proposed-x509-root-cert-to-revoke",
		Short: "Gets a proposed but not approved root certificate to be revoked " +
			"with the given combination of subject and subject-key-id",
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := cli.NewCLIContext().WithCodec(cdc)

			subject := viper.GetString(FlagSubject)
			subjectKeyID := viper.GetString(FlagSubjectKeyID)

			res, height, err := cliCtx.QueryStore(types.GetProposedCertificateRevocationKey(subject, subjectKeyID), queryRoute)
			if err != nil || res == nil {
				return types.ErrProposedCertificateRevocationDoesNotExist(subject, subjectKeyID)
			}

			var revocation types.ProposedCertificateRevocation
			cdc.MustUnmarshalBinaryBare(res, &revocation)

			return cliCtx.EncodeAndPrintWithHeight(revocation, height)
		},
	}

	cmd.Flags().StringP(FlagSubject, FlagSubjectShortcut, "", "Certificate's subject")
	cmd.Flags().StringP(FlagSubjectKeyID, FlagSubjectKeyIDShortcut, "", "Certificate's subject key id (hex)")

	_ = cmd.MarkFlagRequired(FlagSubject)
	_ = cmd.MarkFlagRequired(FlagSubjectKeyID)

	return cmd
}

// nolint:dupl
func GetCmdGetRevokedX509Cert(queryRoute string, cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use: "revoked-x509-cert",
		Short: "Gets revoked certificates (either root, intermediate or leaf) " +
			"by the given combination of subject and subject-key-id",
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := cli.NewCLIContext().WithCodec(cdc)

			subject := viper.GetString(FlagSubject)
			subjectKeyID := viper.GetString(FlagSubjectKeyID)

			res, height, err := cliCtx.QueryStore(types.GetRevokedCertificateKey(subject, subjectKeyID), queryRoute)
			if err != nil || res == nil {
				return types.ErrRevokedCertificateDoesNotExist(subject, subjectKeyID)
			}

			var certificates types.Certificates
			cdc.MustUnmarshalBinaryBare(res, &certificates)

			return cliCtx.EncodeAndPrintWithHeight(certificates, height)
		},
	}

	cmd.Flags().StringP(FlagSubject, FlagSubjectShortcut, "", "Certificate's subject")
	cmd.Flags().StringP(FlagSubjectKeyID, FlagSubjectKeyIDShortcut, "", "Certificate's subject key id (hex)")

	_ = cmd.MarkFlagRequired(FlagSubject)
	_ = cmd.MarkFlagRequired(FlagSubjectKeyID)

	return cmd
}

func GetCmdGetAllRevokedX509RootCerts(queryRoute string, cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "all-revoked-x509-root-certs",
		Short: "Gets all revoked root certificates",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return performPkiQuery(cdc, fmt.Sprintf("custom/%s/all_revoked_x509_root_certs", queryRoute))
		},
	}

	cmd.Flags().Int(pagination.FlagSkip, 0, "amount of certificates to skip")
	cmd.Flags().Int(pagination.FlagTake, 0, "amount of certificates to take")

	return cmd
}

func GetCmdGetAllRevokedX509Certs(queryRoute string, cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "all-revoked-x509-certs",
		Short: "Gets all revoked certificates (root, intermediate and leaf)",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return performPkiQuery(cdc, fmt.Sprintf("custom/%s/all_revoked_x509_certs", queryRoute))
		},
	}

	cmd.Flags().StringP(FlagRootSubject, FlagRootSubjectShortcut, "",
		"filter certificates by `Subject` of root certificate "+
			"(only the certificates originated from the given root certificate are returned)")
	cmd.Flags().StringP(FlagRootSubjectKeyID, FlagRootSubjectKeyIDShortcut, "",
		"filter certificates by `Subject Key Id` of root certificate "+
			"(only the certificates originated from the given root certificate are returned)")
	cmd.Flags().Int(pagination.FlagSkip, 0, "amount of certificates to skip")
	cmd.Flags().Int(pagination.FlagTake, 0, "amount of certificates to take")

	return cmd
}

func chainCertificates(cliCtx cli.CliContext, queryRoute string,
	subject string, subjectKeyID string, chain *types.Certificates) (int64, sdk.Error) {
	res, height, err := cliCtx.QueryStore(types.GetApprovedCertificateKey(subject, subjectKeyID), queryRoute)
	if err != nil || res == nil {
		return height, types.ErrCertificateDoesNotExist(subject, subjectKeyID)
	}

	var certificates types.Certificates

	cliCtx.Codec().MustUnmarshalBinaryBare(res, &certificates)

	certificate := certificates.Items[len(certificates.Items)-1]
	chain.Items = append(chain.Items, certificate)

	if !certificate.IsRoot {
		return chainCertificates(cliCtx, queryRoute, certificate.Issuer, certificate.AuthorityKeyID, chain)
	}

	return height, nil
}

func performPkiQuery(cdc *codec.Codec, route string) error {
	cliCtx := cli.NewCLIContext().WithCodec(cdc)

	rootSubject := viper.GetString(FlagRootSubject)
	rootSubjectKeyID := viper.GetString(FlagRootSubjectKeyID)

	paginationParams := pagination.ParsePaginationParamsFromFlags()
	params := types.NewPkiQueryParams(paginationParams, rootSubject, rootSubjectKeyID)

	return cliCtx.QueryList(route, params)
}
