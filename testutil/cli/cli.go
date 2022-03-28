package cli

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/client"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/require"
)

func ExecTestCLITxCmd(t *testing.T, clientCtx client.Context, cmd *cobra.Command, args []string) (*sdk.TxResponse, error) {
	t.Helper()

	out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, args)
	require.NoError(t, err)

	var resp sdk.TxResponse
	err = clientCtx.Codec.UnmarshalJSON(out.Bytes(), &resp)
	require.NoError(t, err)

	if resp.Code != 0 {
		err = sdkerrors.ABCIError(resp.Codespace, resp.Code, resp.RawLog)

		return nil, err
	}

	return &resp, nil
}
