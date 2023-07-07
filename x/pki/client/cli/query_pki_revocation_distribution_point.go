package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/cli"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

func CmdListPkiRevocationDistributionPoint() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "all-revocation-points",
		Short: "Gets all PKI revocation distribution points",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllPkiRevocationDistributionPointRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.PkiRevocationDistributionPointAll(context.Background(), params)
			if cli.IsKeyNotFoundRPCError(err) {
				return clientCtx.PrintString(cli.LightClientProxyForListQueries)
			}
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowPkiRevocationDistributionPoint() *cobra.Command {
	var (
		vid                int32
		label              string
		issuerSubjectKeyID string
	)

	cmd := &cobra.Command{
		Use:   "revocation-point",
		Short: "Gets a PKI revocation distribution point",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			var res types.PkiRevocationDistributionPoint

			return cli.QueryWithProof(
				clientCtx,
				pkitypes.StoreKey,
				types.PkiRevocationDistributionPointKeyPrefix,
				types.PkiRevocationDistributionPointKey(vid, label, issuerSubjectKeyID),
				&res,
			)
		},
	}

	cmd.Flags().Int32Var(&vid, FlagVid, 0, "Vendor ID (positive non-zero)")
	cmd.Flags().StringVarP(&label, FlagLabel, FlagLabelShortcut, "", "A label to disambiguate multiple revocation information partitions of a particular issuer.")
	cmd.Flags().StringVar(&issuerSubjectKeyID, FlagIssuerSubjectKeyID, "", "Uniquely identifies the PAA or PAI for which this revocation distribution point is provided. Must consist of even number of uppercase hexadecimal characters ([0-9A-F]), with no whitespace and no non-hexadecimal characters., e.g: 5A880E6C3653D07FB08971A3F473790930E62BDB")
	flags.AddQueryFlagsToCmd(cmd)

	_ = cmd.MarkFlagRequired(FlagVid)
	_ = cmd.MarkFlagRequired(FlagLabel)
	_ = cmd.MarkFlagRequired(FlagIssuerSubjectKeyID)

	return cmd
}
