package cli

import (
    "context"
	
     "github.com/spf13/cast"
    "github.com/spf13/cobra"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
    "github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

func CmdListPKIRevocationDistributionPoint() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-pki-revocation-distribution-point",
		Short: "list all PKIRevocationDistributionPoint",
		RunE: func(cmd *cobra.Command, args []string) error {
            clientCtx := client.GetClientContextFromCmd(cmd)

            pageReq, err := client.ReadPageRequest(cmd.Flags())
            if err != nil {
                return err
            }

            queryClient := types.NewQueryClient(clientCtx)

            params := &types.QueryAllPKIRevocationDistributionPointRequest{
                Pagination: pageReq,
            }

            res, err := queryClient.PKIRevocationDistributionPointAll(context.Background(), params)
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

func CmdShowPKIRevocationDistributionPoint() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-pki-revocation-distribution-point [vid] [label] [issuer-subject-key-id]",
		Short: "shows a PKIRevocationDistributionPoint",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
            clientCtx := client.GetClientContextFromCmd(cmd)

            queryClient := types.NewQueryClient(clientCtx)

             argVid, err := cast.ToUint64E(args[0])
            		if err != nil {
                		return err
            		}
             argLabel := args[1]
             argIssuerSubjectKeyID := args[2]
            
            params := &types.QueryGetPKIRevocationDistributionPointRequest{
                Vid: argVid,
                Label: argLabel,
                IssuerSubjectKeyID: argIssuerSubjectKeyID,
                
            }

            res, err := queryClient.PKIRevocationDistributionPoint(context.Background(), params)
            if err != nil {
                return err
            }

            return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

    return cmd
}
