package cli_test

/*
import (
	"fmt"
	"strconv"
	"testing"

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

// Prevent strconv unused error
var _ = strconv.IntSize

func networkWithPkiRevocationDistributionPointsByIssuerSubjectKeyIdObjects(t *testing.T, n int) (*network.Network, []types.PkiRevocationDistributionPointsByIssuerSubjectKeyId) {
	t.Helper()
	cfg := network.DefaultConfig()
	state := types.GenesisState{}
	require.NoError(t, cfg.Codec.UnmarshalJSON(cfg.GenesisState[pkitypes.ModuleName], &state))

	for i := 0; i < n; i++ {
		pkiRevocationDistributionPointsByIssuerSubjectKeyId := types.PkiRevocationDistributionPointsByIssuerSubjectKeyId{
			IssuerSubjectKeyId: strconv.Itoa(i),
		}
		nullify.Fill(&pkiRevocationDistributionPointsByIssuerSubjectKeyId)
		state.PkiRevocationDistributionPointsByIssuerSubjectKeyIdList = append(state.PkiRevocationDistributionPointsByIssuerSubjectKeyIdList, pkiRevocationDistributionPointsByIssuerSubjectKeyId)
	}
	buf, err := cfg.Codec.MarshalJSON(&state)
	require.NoError(t, err)
	cfg.GenesisState[pkitypes.ModuleName] = buf
	return network.New(t, cfg), state.PkiRevocationDistributionPointsByIssuerSubjectKeyIdList
}

func TestShowPkiRevocationDistributionPointsByIssuerSubjectKeyId(t *testing.T) {
	net, objs := networkWithPkiRevocationDistributionPointsByIssuerSubjectKeyIdObjects(t, 2)

	ctx := net.Validators[0].ClientCtx
	common := []string{
		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	}
	for _, tc := range []struct {
		desc                 string
		idIssuerSubjectKeyId string

		args []string
		err  error
		obj  types.PkiRevocationDistributionPointsByIssuerSubjectKeyId
	}{
		{
			desc:                 "found",
			idIssuerSubjectKeyId: objs[0].IssuerSubjectKeyId,

			args: common,
			obj:  objs[0],
		},
		{
			desc:                 "not found",
			idIssuerSubjectKeyId: strconv.Itoa(100000),

			args: common,
			err:  status.Error(codes.InvalidArgument, "not found"),
		},
	} {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			args := []string{
				tc.idIssuerSubjectKeyId,
			}
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdShowPkiRevocationDistributionPointsByIssuerSubjectKeyId(), args)
			if tc.err != nil {
				stat, ok := status.FromError(tc.err)
				require.True(t, ok)
				require.ErrorIs(t, stat.Err(), tc.err)
			} else {
				require.NoError(t, err)
				var resp types.QueryGetPkiRevocationDistributionPointsByIssuerSubjectKeyIdResponse
				require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.NotNil(t, resp.PkiRevocationDistributionPointsByIssuerSubjectKeyId)
				require.Equal(t,
					nullify.Fill(&tc.obj),
					nullify.Fill(&resp.PkiRevocationDistributionPointsByIssuerSubjectKeyId),
				)
			}
		})
	}
}
*/
