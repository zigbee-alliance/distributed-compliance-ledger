package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliance/types"
)

func CmdListComplianceInfo() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-compliance-info",
		Short: "list all ComplianceInfo",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllComplianceInfoRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.ComplianceInfoAll(context.Background(), params)
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

func CmdShowComplianceInfo() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-compliance-info [vid] [pid] [software-version] [certification-type]",
		Short: "shows a ComplianceInfo",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argVid, err := cast.ToInt32E(args[0])
			if err != nil {
				return err
			}
			argPid, err := cast.ToInt32E(args[1])
			if err != nil {
				return err
			}
			argSoftwareVersion, err := cast.ToUint64E(args[2])
			if err != nil {
				return err
			}
			argCertificationType := args[3]

			params := &types.QueryGetComplianceInfoRequest{
				Vid:               argVid,
				Pid:               argPid,
				SoftwareVersion:   argSoftwareVersion,
				CertificationType: argCertificationType,
			}

			res, err := queryClient.ComplianceInfo(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
