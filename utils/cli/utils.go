package cli

import (
	"io/ioutil"
	"os"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	NotFoundOutput = "\"Not Found\""
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
