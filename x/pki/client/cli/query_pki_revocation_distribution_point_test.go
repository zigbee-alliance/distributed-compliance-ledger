package cli_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/cosmos/cosmos-sdk/client/flags"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/stretchr/testify/require"
	tmcli "github.com/tendermint/tendermint/libs/cli"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/zigbee-alliance/distributed-compliance-ledger/testutil/network"
	"github.com/zigbee-alliance/distributed-compliance-ledger/testutil/nullify"
	dclpkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/client/cli"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func networkWithPkiRevocationDistributionPointObjects(t *testing.T, n int) (*network.Network, []types.PkiRevocationDistributionPoint) {
	t.Helper()
	cfg := network.DefaultConfig()
	state := types.GenesisState{}
	require.NoError(t, cfg.Codec.UnmarshalJSON(cfg.GenesisState[dclpkitypes.ModuleName], &state))

	for i := 0; i < n; i++ {
		pKIRevocationDistributionPoint := types.PkiRevocationDistributionPoint{
			Vid:                uint64(i),
			Label:              strconv.Itoa(i),
			IssuerSubjectKeyID: strconv.Itoa(i),
		}
		nullify.Fill(&pKIRevocationDistributionPoint)
		state.PkiRevocationDistributionPointList = append(state.PkiRevocationDistributionPointList, pKIRevocationDistributionPoint)
	}
	buf, err := cfg.Codec.MarshalJSON(&state)
	require.NoError(t, err)
	cfg.GenesisState[dclpkitypes.ModuleName] = buf
	return network.New(t, cfg), state.PkiRevocationDistributionPointList
}

func TestShowPkiRevocationDistributionPoint(t *testing.T) {
	net, objs := networkWithPkiRevocationDistributionPointObjects(t, 2)

	ctx := net.Validators[0].ClientCtx
	common := []string{
		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	}
	for _, tc := range []struct {
		desc                 string
		idVid                uint64
		idLabel              string
		idIssuerSubjectKeyID string

		args []string
		err  error
		obj  types.PkiRevocationDistributionPoint
	}{
		{
			desc:                 "found",
			idVid:                objs[0].Vid,
			idLabel:              objs[0].Label,
			idIssuerSubjectKeyID: objs[0].IssuerSubjectKeyID,

			args: common,
			obj:  objs[0],
		},
		{
			desc:                 "not found",
			idVid:                100000,
			idLabel:              strconv.Itoa(100000),
			idIssuerSubjectKeyID: strconv.Itoa(100000),

			args: common,
			err:  status.Error(codes.InvalidArgument, "not found"),
		},
	} {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			args := []string{
				strconv.Itoa(int(tc.idVid)),
				tc.idLabel,
				tc.idIssuerSubjectKeyID,
			}
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdShowPkiRevocationDistributionPoint(), args)
			if tc.err != nil {
				stat, ok := status.FromError(tc.err)
				require.True(t, ok)
				require.ErrorIs(t, stat.Err(), tc.err)
			} else {
				require.NoError(t, err)
				var resp types.QueryGetPkiRevocationDistributionPointResponse
				require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.NotNil(t, resp.PkiRevocationDistributionPoint)
				require.Equal(t,
					nullify.Fill(&tc.obj),
					nullify.Fill(&resp.PkiRevocationDistributionPoint),
				)
			}
		})
	}
}

func TestListPkiRevocationDistributionPoint(t *testing.T) {
	net, objs := networkWithPkiRevocationDistributionPointObjects(t, 5)

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
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListPkiRevocationDistributionPoint(), args)
			require.NoError(t, err)
			var resp types.QueryAllPkiRevocationDistributionPointResponse
			require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.LessOrEqual(t, len(resp.PkiRevocationDistributionPoint), step)
			require.Subset(t,
				nullify.Fill(objs),
				nullify.Fill(resp.PkiRevocationDistributionPoint),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(objs); i += step {
			args := request(next, 0, uint64(step), false)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListPkiRevocationDistributionPoint(), args)
			require.NoError(t, err)
			var resp types.QueryAllPkiRevocationDistributionPointResponse
			require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.LessOrEqual(t, len(resp.PkiRevocationDistributionPoint), step)
			require.Subset(t,
				nullify.Fill(objs),
				nullify.Fill(resp.PkiRevocationDistributionPoint),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		args := request(nil, 0, uint64(len(objs)), true)
		out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListPkiRevocationDistributionPoint(), args)
		require.NoError(t, err)
		var resp types.QueryAllPkiRevocationDistributionPointResponse
		require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
		require.NoError(t, err)
		require.Equal(t, len(objs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(objs),
			nullify.Fill(resp.PkiRevocationDistributionPoint),
		)
	})
}
