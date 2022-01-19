package cli

import (
	"errors"
	"io/ioutil"
	"os"
	"strings"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	rpctypes "github.com/tendermint/tendermint/rpc/jsonrpc/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	NotFoundOutput            = "\"Not Found\"\n"
	LightClientForListQueries = "\"List queries don't work with a Light Client Proxy. Please connect to a Full Node client you trust if you need to use list queries.\"\n"
)

func ReadFromFile(target string) (string, error) {
	if _, err := os.Stat(target); err == nil { // check whether it is a path
		bytes, err := ioutil.ReadFile(target)
		if err != nil {
			return "", err
		}

		return string(bytes), nil
	} else { // else return as is
		return target, nil
	}
}

func IsNotFound(err error) bool {
	if err == nil {
		return false
	}
	s, ok := status.FromError(err)
	if !ok {
		return false
	}
	return s.Code() == codes.NotFound
}
func IsKeyNotFoundRpcError(err error) bool {
	if err == nil {
		return false
	}
	var rpcerror *rpctypes.RPCError
	if !errors.As(err, &rpcerror) {
		return false
	}
	return strings.Contains(rpcerror.Message, "Internal error") && strings.Contains(rpcerror.Data, "empty key")
}

func AddTxFlagsToCmd(cmd *cobra.Command) {
	flags.AddTxFlagsToCmd(cmd)

	// TODO there might be a better way how to filter that
	hiddenFlags := []string{
		flags.FlagFees,
		flags.FlagFeeAccount,
		flags.FlagGasPrices,
		flags.FlagGasAdjustment,
		flags.FlagGas,
		flags.FlagFeeAccount,
		flags.FlagDryRun, // TODO that flag might be actually useful but relates to gas
	}
	for _, f := range hiddenFlags {
		_ = cmd.Flags().MarkHidden(f)
	}
}
