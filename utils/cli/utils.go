package cli

import (
	"errors"
	"fmt"
	"os"
	"strings"

	rpctypes "github.com/cometbft/cometbft/rpc/jsonrpc/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"
)

const (
	NotFoundOutput                   = "\"Not Found\"\n"
	LightClientProxyForListQueries   = "\"List queries don't work with a Light Client Proxy. Please connect to a Full Node you trust if you need to use list queries.\"\n"
	LightClientProxyForWriteRequests = "\"Write requests don't work with a Light Client Proxy. Please connect to a Full Node (Validator or Observer) instead.\"\n"
	RequestFailed                    = "Request failed"
)

func QueryWithProof(clientCtx client.Context, storeName string, keyPrefix string, key []byte, res codec.ProtoMarshaler) error {
	key = append([]byte(keyPrefix), key...)
	resBytes, _, err := clientCtx.QueryStore(key, storeName)
	// TODO: for some reasons EOF error can be returned sometimes.
	// See https://github.com/zigbee-alliance/distributed-compliance-ledger/issues/203
	if isEOFError(err) {
		return clientCtx.PrintString(fmt.Sprintf("Request failed: %s. Please re-try.", err.Error()))
	}
	if IsEmptySubtreeRPCError(err) {
		// TODO: if no write requests has been sent for a module, an attempt to query a non-existent result will
		// cause an error: verify absence proof: could not calculate root for proof: Nonexistence proof has empty Left and Right proof: invalid proof
		// interpret as Not Found for now
		return clientCtx.PrintString(NotFoundOutput)
	}
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
	// TODO: for some reasons EOF error can be returned sometimes.
	// See https://github.com/zigbee-alliance/distributed-compliance-ledger/issues/203
	if isEOFError(err) {
		return clientCtx.PrintString(fmt.Sprintf("Request failed: %s. Please re-try.", err.Error()))
	}
	if IsEmptySubtreeRPCError(err) {
		// TODO: if no write requests has been sent for a module, an attempt to query a non-existent result will
		// cause an error: verify absence proof: could not calculate root for proof: Nonexistence proof has empty Left and Right proof: invalid proof
		// interpret as Not Found for now
		return clientCtx.PrintString(NotFoundOutput)
	}
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
		bytes, err := os.ReadFile(target)
		if err != nil {
			return "", err
		}

		return string(bytes), nil
	}

	return target, nil
}

func IsKeyNotFoundRPCError(err error) bool {
	if err == nil {
		return false
	}
	var rpcerror *rpctypes.RPCError

	if !errors.As(err, &rpcerror) {
		return false
	}

	return strings.Contains(rpcerror.Message, "Internal error") && strings.Contains(rpcerror.Data, "empty key")
}

func IsWriteInsteadReadRPCError(err error) bool {
	return isRPCError22(err) || IsKeyNotFoundRPCError(err)
}

func isRPCError22(err error) bool {
	if err == nil {
		return false
	}
	var rpcerror *rpctypes.RPCError
	if !errors.As(err, &rpcerror) {
		return false
	}

	return strings.Contains(rpcerror.Message, "Internal error") && strings.Contains(rpcerror.Data, "err response code: 22")
}

func IsEmptySubtreeRPCError(err error) bool {
	if err == nil {
		return false
	}
	var rpcerror *rpctypes.RPCError
	if !errors.As(err, &rpcerror) {
		return false
	}

	return strings.Contains(rpcerror.Message, "Internal error") && strings.Contains(rpcerror.Data, "nonexistence proof has empty Left and Right proof")
}

func isEOFError(err error) bool {
	if err == nil {
		return false
	}

	return strings.Contains(err.Error(), "EOF")
}

func AddTxFlagsToCmd(cmd *cobra.Command) {
	flags.AddTxFlagsToCmd(cmd)

	// TODO there might be a better way how to filter that
	hiddenFlags := []string{
		flags.FlagFees,
		flags.FlagFeeGranter,
		flags.FlagGasPrices,
		flags.FlagGasAdjustment,
		flags.FlagGas,
		flags.FlagFeePayer,
		flags.FlagDryRun, // TODO that flag might be actually useful but relates to gas
	}
	for _, f := range hiddenFlags {
		_ = cmd.Flags().MarkHidden(f)
	}
}
