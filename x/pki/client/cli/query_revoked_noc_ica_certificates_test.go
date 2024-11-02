package cli_test

// import (
//	"fmt"
//	"strconv"
//	"testing"
//
//	"github.com/cosmos/cosmos-sdk/client/flags"
//	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
//	"github.com/stretchr/testify/require"
//	tmcli "github.com/cometbft/cometbft/libs/cli"
//	"google.golang.org/grpc/codes"
//	"google.golang.org/grpc/status"
//
//	"github.com/zigbee-alliance/distributed-compliance-ledger/testutil/network"
//	"github.com/zigbee-alliance/distributed-compliance-ledger/testutil/nullify"
//	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/client/cli"
//    "github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
//)
//
//// Prevent strconv unused error
// var _ = strconv.IntSize
//
// func networkWithRevokedNocIcaCertificatesObjects(t *testing.T, n int) (*network.Network, []types.RevokedNocIcaCertificates) {
//	t.Helper()
//	cfg := network.DefaultConfig()
//	state := types.GenesisState{}
//	for i := 0; i < n; i++ {
//	revokedNocIcaCertificates := types.RevokedNocIcaCertificates{
//			Subject: strconv.Itoa(i),
//			SubjectKeyID: strconv.Itoa(i),
//
//		}
//		nullify.Fill(&revokedNocIcaCertificates)
//		state.RevokedNocIcaCertificatesList = append(state.RevokedNocIcaCertificatesList, revokedNocIcaCertificates)
//	}
//	buf, err := cfg.Codec.MarshalJSON(&state)
//	require.NoError(t, err)
//	cfg.GenesisState[types.ModuleName] = buf
//	return network.New(t, cfg), state.RevokedNocIcaCertificatesList
//}
//
// func TestShowRevokedNocIcaCertificates(t *testing.T) {
//	net, objs := networkWithRevokedNocIcaCertificatesObjects(t, 2)
//
//	ctx := net.Validators[0].ClientCtx
//	common := []string{
//		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
//	}
//	tests := []struct {
//		desc string
//		idSubject string
//        idSubjectKeyId string
//
//		args []string
//		err  error
//		obj  types.RevokedNocIcaCertificates
//	}{
//		{
//			desc: "found",
//			idSubject: objs[0].Subject,
//            idSubjectKeyId: objs[0].SubjectKeyID,
//
//			args: common,
//			obj:  objs[0],
//		},
//		{
//			desc: "not found",
//			idSubject: strconv.Itoa(100000),
//            idSubjectKeyId: strconv.Itoa(100000),
//
//			args: common,
//			err:  status.Error(codes.NotFound, "not found"),
//		},
//	}
//	for _, tc := range tests {
//		t.Run(tc.desc, func(t *testing.T) {
//			args := []string{
//			    tc.idSubject,
//                tc.idSubjectKeyId,
//
//			}
//			args = append(args, tc.args...)
//			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdShowRevokedNocIcaCertificates(), args)
//			if tc.err != nil {
//				stat, ok := status.FromError(tc.err)
//				require.True(t, ok)
//				require.ErrorIs(t, stat.Err(), tc.err)
//			} else {
//				require.NoError(t, err)
//				var resp types.QueryGetRevokedNocIcaCertificatesResponse
//				require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
//				require.NotNil(t, resp.RevokedNocIcaCertificates)
//				require.Equal(t,
//					nullify.Fill(&tc.obj),
//					nullify.Fill(&resp.RevokedNocIcaCertificates),
//				)
//			}
//		})
//	}
//}
//
//func TestListRevokedNocIcaCertificates(t *testing.T) {
//	net, objs := networkWithRevokedNocIcaCertificatesObjects(t, 5)
//
//	ctx := net.Validators[0].ClientCtx
//	request := func(next []byte, offset, limit uint64, total bool) []string {
//		args := []string{
//			fmt.Sprintf("--%s=json", tmcli.OutputFlag),
//		}
//		if next == nil {
//			args = append(args, fmt.Sprintf("--%s=%d", flags.FlagOffset, offset))
//		} else {
//			args = append(args, fmt.Sprintf("--%s=%s", flags.FlagPageKey, next))
//		}
//		args = append(args, fmt.Sprintf("--%s=%d", flags.FlagLimit, limit))
//		if total {
//			args = append(args, fmt.Sprintf("--%s", flags.FlagCountTotal))
//		}
//		return args
//	}
//	t.Run("ByOffset", func(t *testing.T) {
//		step := 2
//		for i := 0; i < len(objs); i += step {
//			args := request(nil, uint64(i), uint64(step), false)
//			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListRevokedNocIcaCertificates(), args)
//			require.NoError(t, err)
//			var resp types.QueryAllRevokedNocIcaCertificatesResponse
//			require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
//			require.LessOrEqual(t, len(resp.RevokedNocIcaCertificates), step)
//			require.Subset(t,
//            	nullify.Fill(objs),
//            	nullify.Fill(resp.RevokedNocIcaCertificates),
//            )
//		}
//	})
//	t.Run("ByKey", func(t *testing.T) {
//		step := 2
//		var next []byte
//		for i := 0; i < len(objs); i += step {
//			args := request(next, 0, uint64(step), false)
//			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListRevokedNocIcaCertificates(), args)
//			require.NoError(t, err)
//			var resp types.QueryAllRevokedNocIcaCertificatesResponse
//			require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
//			require.LessOrEqual(t, len(resp.RevokedNocIcaCertificates), step)
//			require.Subset(t,
//            	nullify.Fill(objs),
//            	nullify.Fill(resp.RevokedNocIcaCertificates),
//            )
//			next = resp.Pagination.NextKey
//		}
//	})
//	t.Run("Total", func(t *testing.T) {
//		args := request(nil, 0, uint64(len(objs)), true)
//		out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListRevokedNocIcaCertificates(), args)
//		require.NoError(t, err)
//		var resp types.QueryAllRevokedNocIcaCertificatesResponse
//		require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
//		require.NoError(t, err)
//		require.Equal(t, len(objs), int(resp.Pagination.Total))
//		require.ElementsMatch(t,
//			nullify.Fill(objs),
//			nullify.Fill(resp.RevokedNocIcaCertificates),
//		)
//	})
//}
