package helpers

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func Tx(module, command, from string, txArgs ...string) (sdk.TxResponse, error) {
	args := []string{"tx", module, command}

	// TXN arguments
	args = append(args, txArgs...)

	// Sender account
	args = append(args, "--from", from)

	// Broadcast
	args = append(args, "--yes")

	output, err := Command(args...)
	if err != nil {
		return sdk.TxResponse{}, err
	}

	var resp sdk.TxResponse
	err = Codec.UnmarshalJSON(output, &resp)
	if err != nil {
		return sdk.TxResponse{}, err
	}

	return resp, nil
}

func AwaitTxConfirmation(hash string) (string, error) {
	var (
		result []byte
		err    error
	)
	for i := 1; i <= 20; i++ {
		result, err = Command("query", "tx", hash)
		if err == nil {
			return string(result), nil
		} else {
			time.Sleep(2 * time.Second)
		}
	}

	return "", err
}
