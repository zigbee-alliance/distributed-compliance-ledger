package cli_test

/* TODO issue #197
import (
	"fmt"
	"strconv"
	"testing"

	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/stretchr/testify/require"
	tmcli "github.com/cometbft/cometbft/libs/cli"
	"github.com/zigbee-alliance/distributed-compliance-ledger/testutil/network"
	"github.com/zigbee-alliance/distributed-compliance-ledger/testutil/nullify"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/client/cli"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Prevent strconv unused error.
var _ = strconv.IntSize

func networkWithApprovedCertificatesBySubjectObjects(t *testing.T, n int) (*network.Network, []types.ApprovedCertificatesBySubject) {
	t.Helper()
	cfg := network.DefaultConfig()
	state := types.GenesisState{}
	require.NoError(t, cfg.Codec.UnmarshalJSON(cfg.GenesisState[types.ModuleName], &state))

	for i := 0; i < n; i++ {
		approvedCertificatesBySubject := types.ApprovedCertificatesBySubject{
			Subject: strconv.Itoa(i),
		}
		nullify.Fill(&approvedCertificatesBySubject)
		state.ApprovedCertificatesBySubjectList = append(state.ApprovedCertificatesBySubjectList, approvedCertificatesBySubject)
	}
	buf, err := cfg.Codec.MarshalJSON(&state)
	require.NoError(t, err)
	cfg.GenesisState[types.ModuleName] = buf
	return network.New(t, cfg), state.ApprovedCertificatesBySubjectList
}

func TestShowApprovedCertificatesBySubject(t *testing.T) {
	net, objs := networkWithApprovedCertificatesBySubjectObjects(t, 2)

	ctx := net.Validators[0].ClientCtx
	common := []string{
		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	}
	for _, tc := range []struct {
		desc      string
		idSubject string

		args []string
		err  error
		obj  types.ApprovedCertificatesBySubject
	}{
		{
			desc:      "found",
			idSubject: objs[0].Subject,

			args: common,
			obj:  objs[0],
		},
		{
			desc:      "not found",
			idSubject: strconv.Itoa(100000),

			args: common,
			err:  status.Error(codes.NotFound, "not found"),
		},
	} {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			args := []string{
				tc.idSubject,
			}
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdShowApprovedCertificatesBySubject(), args)
			if tc.err != nil {
				stat, ok := status.FromError(tc.err)
				require.True(t, ok)
				require.ErrorIs(t, stat.Err(), tc.err)
			} else {
				require.NoError(t, err)
				var resp types.QueryGetApprovedCertificatesBySubjectResponse
				require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.NotNil(t, resp.ApprovedCertificatesBySubject)
				require.Equal(t,
					nullify.Fill(&tc.obj),
					nullify.Fill(&resp.ApprovedCertificatesBySubject),
				)
			}
		})
	}
}
*/
