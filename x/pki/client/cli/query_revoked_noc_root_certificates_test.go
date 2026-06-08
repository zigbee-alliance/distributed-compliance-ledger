package cli_test

/* TODO issue #197
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
	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/client/cli"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

// Prevent strconv unused error.
var _ = strconv.IntSize

func networkWithRevokedNocRootCertificatesObjects(t *testing.T, n int) (*network.Network, []types.RevokedNocRootCertificates) {
	t.Helper()
	cfg := network.DefaultConfig()
	state := types.GenesisState{}
	require.NoError(t, cfg.Codec.UnmarshalJSON(cfg.GenesisState[pkitypes.ModuleName], &state))

	for i := 0; i < n; i++ {
		revokedNocRootCertificates := types.RevokedNocRootCertificates{
			Subject:      strconv.Itoa(i),
			SubjectKeyID: strconv.Itoa(i),
		}
		nullify.Fill(&revokedNocRootCertificates)
		state.RevokedNocRootCertificatesList = append(state.RevokedNocRootCertificatesList, revokedNocRootCertificates)
	}
	buf, err := cfg.Codec.MarshalJSON(&state)
	require.NoError(t, err)
	cfg.GenesisState[pkitypes.ModuleName] = buf

	return network.New(t, cfg), state.RevokedNocRootCertificatesList
}

func TestShowRevokedNocRootCertificates(t *testing.T) {
	net, objs := networkWithRevokedNocRootCertificatesObjects(t, 2)

	ctx := net.Validators[0].ClientCtx
	common := []string{
		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	}
	for _, tc := range []struct {
		desc           string
		idSubject      string
		idSubjectKeyID string

		args []string
		err  error
		obj  types.RevokedNocRootCertificates
	}{
		{
			desc:           "found",
			idSubject:      objs[0].Subject,
			idSubjectKeyID: objs[0].SubjectKeyID,

			args: common,
			obj:  objs[0],
		},
		{
			desc:           "not found",
			idSubject:      strconv.Itoa(100000),
			idSubjectKeyID: strconv.Itoa(100000),

			args: common,
			err:  status.Error(codes.InvalidArgument, "not found"),
		},
	} {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			args := []string{
				tc.idSubject,
				tc.idSubjectKeyID,
			}
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdShowRevokedNocRootCertificates(), args)
			if tc.err != nil {
				stat, ok := status.FromError(tc.err)
				require.True(t, ok)
				require.ErrorIs(t, stat.Err(), tc.err)
			} else {
				require.NoError(t, err)
				var resp types.QueryGetRevokedNocRootCertificatesResponse
				require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.NotNil(t, resp.RevokedNocRootCertificates)
				require.Equal(t,
					nullify.Fill(&tc.obj),
					nullify.Fill(&resp.RevokedNocRootCertificates),
				)
			}
		})
	}
}

func TestListRevokedNocRootCertificates(t *testing.T) {
	net, objs := networkWithRevokedNocRootCertificatesObjects(t, 5)

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
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListRevokedNocRootCertificates(), args)
			require.NoError(t, err)
			var resp types.QueryAllRevokedNocRootCertificatesResponse
			require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.LessOrEqual(t, len(resp.RevokedNocRootCertificates), step)
			require.Subset(t,
				nullify.Fill(objs),
				nullify.Fill(resp.RevokedNocRootCertificates),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(objs); i += step {
			args := request(next, 0, uint64(step), false)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListRevokedNocRootCertificates(), args)
			require.NoError(t, err)
			var resp types.QueryAllRevokedNocRootCertificatesResponse
			require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.LessOrEqual(t, len(resp.RevokedNocRootCertificates), step)
			require.Subset(t,
				nullify.Fill(objs),
				nullify.Fill(resp.RevokedNocRootCertificates),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		args := request(nil, 0, uint64(len(objs)), true)
		out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListRevokedNocRootCertificates(), args)
		require.NoError(t, err)
		var resp types.QueryAllRevokedNocRootCertificatesResponse
		require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
		require.NoError(t, err)
		require.Equal(t, len(objs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(objs),
			nullify.Fill(resp.RevokedNocRootCertificates),
		)
	})
}
*/
