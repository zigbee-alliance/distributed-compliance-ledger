package cli_test

import (
	"fmt"
	"strconv"
	"testing"

	tmcli "github.com/cometbft/cometbft/libs/cli"
	"github.com/cosmos/cosmos-sdk/client/flags"
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

func networkWithModelObjects(t *testing.T, n int) (*network.Network, []types.Model) {
	t.Helper()
	cfg := network.DefaultConfig()
	state := types.GenesisState{}
	require.NoError(t, cfg.Codec.UnmarshalJSON(cfg.GenesisState[types.ModuleName], &state))

	for i := 0; i < n; i++ {
		model := types.Model{
			Vid: int32(i + 1),
			Pid: int32(i + 1),
		}
		nullify.Fill(&model)
		state.ModelList = append(state.ModelList, model)
	}
	buf, err := cfg.Codec.MarshalJSON(&state)
	require.NoError(t, err)
	cfg.GenesisState[types.ModuleName] = buf

	return network.New(t, cfg), state.ModelList
}

func TestShowModel(t *testing.T) {
	net, objs := networkWithModelObjects(t, 2)

	ctx := net.Validators[0].ClientCtx
	common := []string{
		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	}
	for _, tc := range []struct {
		desc  string
		idVid int32
		idPid int32

		common []string
		obj    *types.Model
	}{
		{
			desc:  "found",
			idVid: objs[0].Vid,
			idPid: objs[0].Pid,

			common: common,
			obj:    &objs[0],
		},
		{
			desc:  "not found",
			idVid: 100000,
			idPid: 100000,

			common: common,
			obj:    nil,
		},
	} {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			args := []string{
				fmt.Sprintf("--%s=%v", cli.FlagVid, tc.idVid),
				fmt.Sprintf("--%s=%v", cli.FlagPid, tc.idPid),
			}
			args = append(args, tc.common...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdShowModel(), args)
			require.NoError(t, err)
			if tc.obj == nil {
				require.Equal(t, cliutils.NotFoundOutput, out.String())
			} else {
				var model types.Model
				require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &model))
				require.Equal(t,
					nullify.Fill(&tc.obj),
					nullify.Fill(&model),
				)
			}
		})
	}
}

func TestListModel(t *testing.T) {
	net, objs := networkWithModelObjects(t, 5)

	ctx := net.Validators[0].ClientCtx
	request := func(next []byte, offset, limit uint64, total bool) []string {
		args := []string{
			fmt.Sprintf("--%s=json", tmcli.OutputFlag),
		}
		if next == nil {
			args = append(args, fmt.Sprintf("--%s=%d", flags.FlagOffset, offset))
		} else {
			args = append(args, fmt.Sprintf("--%s=%s", flags.FlagPageKey, next))
		}
		args = append(args, fmt.Sprintf("--%s=%d", flags.FlagLimit, limit))
		if total {
			args = append(args, fmt.Sprintf("--%s", flags.FlagCountTotal))
		}

		return args
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(objs); i += step {
			args := request(nil, uint64(i), uint64(step), false)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListModel(), args)
			require.NoError(t, err)
			var resp types.QueryAllModelResponse
			require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.LessOrEqual(t, len(resp.Model), step)
			require.Subset(t,
				nullify.Fill(objs),
				nullify.Fill(resp.Model),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(objs); i += step {
			args := request(next, 0, uint64(step), false)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListModel(), args)
			require.NoError(t, err)
			var resp types.QueryAllModelResponse
			require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.LessOrEqual(t, len(resp.Model), step)
			require.Subset(t,
				nullify.Fill(objs),
				nullify.Fill(resp.Model),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		args := request(nil, 0, uint64(len(objs)), true)
		out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListModel(), args)
		require.NoError(t, err)
		var resp types.QueryAllModelResponse
		require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
		require.NoError(t, err)
		require.Equal(t, len(objs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(objs),
			nullify.Fill(resp.Model),
		)
	})
}

// TODO: Add negative tests for absence of required parameters and
// for presence of unexpected parameters (including positional ones).
