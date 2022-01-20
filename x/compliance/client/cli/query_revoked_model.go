package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/cli"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliance/types"
)

func CmdListRevokedModel() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "all-revoked-models",
		Short: "Query the list of all revoked models",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllRevokedModelRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.RevokedModelAll(context.Background(), params)
			if cli.IsKeyNotFoundRpcError(err) {
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

func CmdShowRevokedModel() *cobra.Command {
	var (
		vid               int32
		pid               int32
		softwareVersion   uint32
		certificationType string
	)

	cmd := &cobra.Command{
		Use:   "revoked-model",
		Short: "Gets a boolean if the given Model (identified by the `vid`, `pid`, `softwareVersion` and `certification_type`) is revoked",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)
			var res types.RevokedModel
			return cli.QueryWithProof(
				clientCtx,
				types.StoreKey,
				types.RevokedModelKeyPrefix,
				types.RevokedModelKey(vid, pid, softwareVersion, certificationType),
				&res,
			)
		},
	}

	cmd.Flags().Int32Var(&vid, FlagVID, 0, "Model vendor ID")
	cmd.Flags().Int32Var(&pid, FlagPID, 0, "Model product ID")
	cmd.Flags().Uint32Var(&softwareVersion, FlagSoftwareVersion, 0, "Model software version")
	cmd.Flags().StringVarP(&certificationType, FlagCertificationType, FlagCertificationTypeShortcut, "", TextCertificationType)

	_ = cmd.MarkFlagRequired(FlagVID)
	_ = cmd.MarkFlagRequired(FlagPID)
	_ = cmd.MarkFlagRequired(FlagSoftwareVersion)
	_ = cmd.MarkFlagRequired(FlagCertificationType)

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
