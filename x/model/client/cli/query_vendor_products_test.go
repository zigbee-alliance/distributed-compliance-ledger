package cli_test

import (
	"fmt"
	"strconv"
	"testing"

	tmcli "github.com/cometbft/cometbft/libs/cli"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/stretchr/testify/require"
	"github.com/zigbee-alliance/distributed-compliance-ledger/testutil/network"
	"github.com/zigbee-alliance/distributed-compliance-ledger/testutil/nullify"
	cliutils "github.com/zigbee-alliance/distributed-compliance-ledger/utils/cli"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/model/client/cli"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/model/types"
)

// Prevent strconv unused error.
var _ = strconv.IntSize

func networkWithVendorProductsObjects(t *testing.T, n int) (*network.Network, []types.VendorProducts) {
	t.Helper()
	cfg := network.DefaultConfig()
	state := types.GenesisState{}
	require.NoError(t, cfg.Codec.UnmarshalJSON(cfg.GenesisState[types.ModuleName], &state))

	for i := 0; i < n; i++ {
		vendorProducts := types.VendorProducts{
			Vid: int32(i + 1),
		}
		nullify.Fill(&vendorProducts)
		state.VendorProductsList = append(state.VendorProductsList, vendorProducts)
	}
	buf, err := cfg.Codec.MarshalJSON(&state)
	require.NoError(t, err)
	cfg.GenesisState[types.ModuleName] = buf

	return network.New(t, cfg), state.VendorProductsList
}

func TestShowVendorProducts(t *testing.T) {
	net, objs := networkWithVendorProductsObjects(t, 2)

	ctx := net.Validators[0].ClientCtx
	common := []string{
		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	}
	for _, tc := range []struct {
		desc  string
		idVid int32

		common []string
		obj    *types.VendorProducts
	}{
		{
			desc:  "found",
			idVid: objs[0].Vid,

			common: common,
			obj:    &objs[0],
		},
		{
			desc:  "not found",
			idVid: 100000,

			common: common,
			obj:    nil,
		},
	} {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			args := []string{
				fmt.Sprintf("--%s=%v", cli.FlagVid, tc.idVid),
			}
			args = append(args, tc.common...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdShowVendorProducts(), args)
			if tc.obj == nil {
				require.Equal(t, cliutils.NotFoundOutput, out.String())
			} else {
				require.NoError(t, err)
				var vendorProducts types.VendorProducts
				require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &vendorProducts))
				require.Equal(t,
					nullify.Fill(&tc.obj),
					nullify.Fill(&vendorProducts),
				)
			}
		})
	}
}

// TODO: Add negative tests for absence of required parameters and
// for presence of unexpected parameters (including positional ones).
