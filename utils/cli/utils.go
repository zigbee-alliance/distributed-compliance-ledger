package cli

import (
	"errors"
	"io/ioutil"
	"os"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"
	rpctypes "github.com/tendermint/tendermint/rpc/jsonrpc/types"
)

const (
	NotFoundOutput                   = "\"Not Found\"\n"
	LightClientProxyForListQueries   = "\"List queries don't work with a Light Client Proxy. Please connect to a Full Node you trust if you need to use list queries.\"\n"
	LightClientProxyForWriteRequests = "\"Write requests don't work with a Light Client Proxy. Please connect to a Full Node (Validator or Observer) instead.\"\n"
)

func QueryWithProof(clientCtx client.Context, storeName string, keyPrefix string, key []byte, res codec.ProtoMarshaler) error {
	key = append([]byte(keyPrefix), key...)
	resBytes, _, err := clientCtx.QueryStore(key, storeName)
	if err != nil {
		return err
	}
	if resBytes == nil {
		return clientCtx.PrintString(NotFoundOutput)
	}

	clientCtx.Codec.MustUnmarshal(resBytes, res)
	return clientCtx.PrintProto(res)
}

func QueryWithProofList(clientCtx client.Context, storeName string, keyPrefix string, key []byte, res codec.ProtoMarshaler) error {
	key = append([]byte(keyPrefix), key...)
	resBytes, _, err := clientCtx.QueryStore(key, storeName)
	if err != nil {
		return err
	}
	// return default (empty) list if not found
	if resBytes != nil {
		clientCtx.Codec.MustUnmarshal(resBytes, res)
	}
	return clientCtx.PrintProto(res)
}

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

func IsWriteInsteadReadRpcError(err error) bool {
	if err == nil {
		return false
	}
	var rpcerror *rpctypes.RPCError
	if !errors.As(err, &rpcerror) {
		return false
	}
	return strings.Contains(rpcerror.Message, "Internal error") && strings.Contains(rpcerror.Data, "err response code: 22")
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
