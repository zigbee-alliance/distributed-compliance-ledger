package cli_test

// import (
// 	"fmt"
// 	"strconv"
// 	"testing"

// 	"github.com/cosmos/cosmos-sdk/client/flags"
// 	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
// 	"github.com/stretchr/testify/require"
// 	tmcli "github.com/tendermint/tendermint/libs/cli"
// 	"google.golang.org/grpc/codes"
// 	"google.golang.org/grpc/status"

// 	"github.com/zigbee-alliance/distributed-compliance-ledger/testutil/network"
// 	"github.com/zigbee-alliance/distributed-compliance-ledger/testutil/nullify"
// 	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/client/cli"
// 	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
// )

// // Prevent strconv unused error
// var _ = strconv.IntSize

// func networkWithNocRootCertificatesObjects(t *testing.T, n int) (*network.Network, []types.NocRootCertificates) {
// 	t.Helper()
// 	cfg := network.DefaultConfig()
// 	state := types.GenesisState{}
// 	require.NoError(t, cfg.Codec.UnmarshalJSON(cfg.GenesisState[types.ModuleName], &state))

// 	for i := 0; i < n; i++ {
// 		nocRootCertificates := types.NocRootCertificates{
// 			Vid: uint64(i),
// 		}
// 		nullify.Fill(&nocRootCertificates)
// 		state.NocRootCertificatesList = append(state.NocRootCertificatesList, nocRootCertificates)
// 	}
// 	buf, err := cfg.Codec.MarshalJSON(&state)
// 	require.NoError(t, err)
// 	cfg.GenesisState[types.ModuleName] = buf
// 	return network.New(t, cfg), state.NocRootCertificatesList
// }

// func TestShowNocRootCertificates(t *testing.T) {
// 	net, objs := networkWithNocRootCertificatesObjects(t, 2)

// 	ctx := net.Validators[0].ClientCtx
// 	common := []string{
// 		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
// 	}
// 	for _, tc := range []struct {
// 		desc  string
// 		idVid uint64

// 		args []string
// 		err  error
// 		obj  types.NocRootCertificates
// 	}{
// 		{
// 			desc:  "found",
// 			idVid: objs[0].Vid,

// 			args: common,
// 			obj:  objs[0],
// 		},
// 		{
// 			desc:  "not found",
// 			idVid: 100000,

// 			args: common,
// 			err:  status.Error(codes.InvalidArgument, "not found"),
// 		},
// 	} {
// 		tc := tc
// 		t.Run(tc.desc, func(t *testing.T) {
// 			args := []string{
// 				strconv.Itoa(int(tc.idVid)),
// 			}
// 			args = append(args, tc.args...)
// 			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdShowNocRootCertificates(), args)
// 			if tc.err != nil {
// 				stat, ok := status.FromError(tc.err)
// 				require.True(t, ok)
// 				require.ErrorIs(t, stat.Err(), tc.err)
// 			} else {
// 				require.NoError(t, err)
// 				var resp types.QueryGetNocRootCertificatesResponse
// 				require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
// 				require.NotNil(t, resp.NocRootCertificates)
// 				require.Equal(t,
// 					nullify.Fill(&tc.obj),
// 					nullify.Fill(&resp.NocRootCertificates),
// 				)
// 			}
// 		})
// 	}
// }

// func TestListNocRootCertificates(t *testing.T) {
// 	net, objs := networkWithNocRootCertificatesObjects(t, 5)

// 	ctx := net.Validators[0].ClientCtx
// 	request := func(next []byte, offset, limit uint64, total bool) []string {
// 		args := []string{
// 			fmt.Sprintf("--%s=json", tmcli.OutputFlag),
// 		}
// 		if next == nil {
// 			args = append(args, fmt.Sprintf("--%s=%d", flags.FlagOffset, offset))
// 		} else {
// 			args = append(args, fmt.Sprintf("--%s=%s", flags.FlagPageKey, next))
// 		}
// 		args = append(args, fmt.Sprintf("--%s=%d", flags.FlagLimit, limit))
// 		if total {
// 			args = append(args, fmt.Sprintf("--%s", flags.FlagCountTotal))
// 		}
// 		return args
// 	}
// 	t.Run("ByOffset", func(t *testing.T) {
// 		step := 2
// 		for i := 0; i < len(objs); i += step {
// 			args := request(nil, uint64(i), uint64(step), false)
// 			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListNocRootCertificates(), args)
// 			require.NoError(t, err)
// 			var resp types.QueryAllNocRootCertificatesResponse
// 			require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
// 			require.LessOrEqual(t, len(resp.NocRootCertificates), step)
// 			require.Subset(t,
// 				nullify.Fill(objs),
// 				nullify.Fill(resp.NocRootCertificates),
// 			)
// 		}
// 	})
// 	t.Run("ByKey", func(t *testing.T) {
// 		step := 2
// 		var next []byte
// 		for i := 0; i < len(objs); i += step {
// 			args := request(next, 0, uint64(step), false)
// 			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListNocRootCertificates(), args)
// 			require.NoError(t, err)
// 			var resp types.QueryAllNocRootCertificatesResponse
// 			require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
// 			require.LessOrEqual(t, len(resp.NocRootCertificates), step)
// 			require.Subset(t,
// 				nullify.Fill(objs),
// 				nullify.Fill(resp.NocRootCertificates),
// 			)
// 			next = resp.Pagination.NextKey
// 		}
// 	})
// 	t.Run("Total", func(t *testing.T) {
// 		args := request(nil, 0, uint64(len(objs)), true)
// 		out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListNocRootCertificates(), args)
// 		require.NoError(t, err)
// 		var resp types.QueryAllNocRootCertificatesResponse
// 		require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
// 		require.NoError(t, err)
// 		require.Equal(t, len(objs), int(resp.Pagination.Total))
// 		require.ElementsMatch(t,
// 			nullify.Fill(objs),
// 			nullify.Fill(resp.NocRootCertificates),
// 		)
// 	})
// }
