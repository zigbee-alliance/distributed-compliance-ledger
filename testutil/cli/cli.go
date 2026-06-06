// Copyright 2020 DSR Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cli

import (
	"testing"

	"cosmossdk.io/errors"
	"github.com/cosmos/cosmos-sdk/client"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	sdk "github.com/cosmos/cosmos-sdk/types"
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
		err = errors.ABCIError(resp.Codespace, resp.Code, resp.RawLog)

		return nil, err
	}

	return &resp, nil
}
