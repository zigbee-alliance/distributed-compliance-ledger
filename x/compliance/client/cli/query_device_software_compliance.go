package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/cli"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliance/types"
)

func CmdListDeviceSoftwareCompliance() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "all-device-software-compliance",
		Short: "Query the list of all device software compliances",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllDeviceSoftwareComplianceRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.DeviceSoftwareComplianceAll(context.Background(), params)
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

func CmdShowDeviceSoftwareCompliance() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "device-software-compliance",
		Short: "Query device software compliance for Model (identified by the `cdCertificateId`)",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			argCDCertificateID := viper.GetString(FlagCDCertificationID)

			var res types.DeviceSoftwareCompliance

			return cli.QueryWithProof(
				clientCtx,
				types.StoreKey,
				types.DeviceSoftwareComplianceKeyPrefix,
				types.DeviceSoftwareComplianceKey(argCDCertificateID),
				&res,
			)
		},
	}

	cmd.Flags().String(FlagCDCertificationID, "", "CD Certification ID of the certification")
	flags.AddQueryFlagsToCmd(cmd)

	_ = cmd.MarkFlagRequired(FlagCDCertificationID)

	return cmd
}
