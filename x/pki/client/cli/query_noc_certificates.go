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

func CmdListNocCertificates() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "all-noc-x509-certs",
		Short: "Gets all NOC non-root certificates",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllNocCertificatesRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.NocCertificatesAll(context.Background(), params)
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

func CmdShowNocCertificates() *cobra.Command {
	var vid int32
	cmd := &cobra.Command{
		Use:   "noc-x509-certs",
		Short: "Gets NOC non-root certificates by VID",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			var res types.NocCertificates

			return cli.QueryWithProof(
				clientCtx,
				pkitypes.StoreKey,
				types.NocCertificatesKeyPrefix,
				types.NocCertificatesKey(vid),
				&res,
			)
		},
	}

	cmd.Flags().Int32Var(&vid, FlagVid, 0, "Vendor ID (positive non-zero)")
	flags.AddQueryFlagsToCmd(cmd)

	_ = cmd.MarkFlagRequired(FlagVid)

	return cmd
}
