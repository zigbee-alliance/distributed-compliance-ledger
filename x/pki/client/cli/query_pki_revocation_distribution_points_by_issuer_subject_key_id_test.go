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

func networkWithPkiRevocationDistributionPointsByIssuerSubjectKeyIDObjects(t *testing.T, n int) (*network.Network, []types.PkiRevocationDistributionPointsByIssuerSubjectKeyID) {
	t.Helper()
	cfg := network.DefaultConfig()
	state := types.GenesisState{}
	require.NoError(t, cfg.Codec.UnmarshalJSON(cfg.GenesisState[pkitypes.ModuleName], &state))

	for i := 0; i < n; i++ {
		pkiRevocationDistributionPointsByIssuerSubjectKeyID := types.PkiRevocationDistributionPointsByIssuerSubjectKeyID{
			IssuerSubjectKeyID: strconv.Itoa(i),
		}
		nullify.Fill(&pkiRevocationDistributionPointsByIssuerSubjectKeyID)
		state.PkiRevocationDistributionPointsByIssuerSubjectKeyIDList = append(state.PkiRevocationDistributionPointsByIssuerSubjectKeyIDList, pkiRevocationDistributionPointsByIssuerSubjectKeyID)
	}
	buf, err := cfg.Codec.MarshalJSON(&state)
	require.NoError(t, err)
	cfg.GenesisState[pkitypes.ModuleName] = buf
	return network.New(t, cfg), state.PkiRevocationDistributionPointsByIssuerSubjectKeyIDList
}

func TestShowPkiRevocationDistributionPointsByIssuerSubjectKeyID(t *testing.T) {
	net, objs := networkWithPkiRevocationDistributionPointsByIssuerSubjectKeyIDObjects(t, 2)

	ctx := net.Validators[0].ClientCtx
	common := []string{
		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	}
	for _, tc := range []struct {
		desc                 string
		idIssuerSubjectKeyID string

		args []string
		err  error
		obj  types.PkiRevocationDistributionPointsByIssuerSubjectKeyID
	}{
		{
			desc:                 "found",
			idIssuerSubjectKeyID: objs[0].IssuerSubjectKeyID,

			args: common,
			obj:  objs[0],
		},
		{
			desc:                 "not found",
			idIssuerSubjectKeyID: strconv.Itoa(100000),

			args: common,
			err:  status.Error(codes.InvalidArgument, "not found"),
		},
	} {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			args := []string{
				tc.idIssuerSubjectKeyID,
			}
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdShowPkiRevocationDistributionPointsByIssuerSubjectKeyID(), args)
			if tc.err != nil {
				stat, ok := status.FromError(tc.err)
				require.True(t, ok)
				require.ErrorIs(t, stat.Err(), tc.err)
			} else {
				require.NoError(t, err)
				var resp types.QueryGetPkiRevocationDistributionPointsByIssuerSubjectKeyIDResponse
				require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.NotNil(t, resp.PkiRevocationDistributionPointsByIssuerSubjectKeyID)
				require.Equal(t,
					nullify.Fill(&tc.obj),
					nullify.Fill(&resp.PkiRevocationDistributionPointsByIssuerSubjectKeyID),
				)
			}
		})
	}
}
*/
