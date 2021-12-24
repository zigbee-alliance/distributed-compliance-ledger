package cli_test

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
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/client/cli"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func networkWithChildCertificatesObjects(t *testing.T, n int) (*network.Network, []types.ChildCertificates) {
	t.Helper()
	cfg := network.DefaultConfig()
	state := types.GenesisState{}
	require.NoError(t, cfg.Codec.UnmarshalJSON(cfg.GenesisState[types.ModuleName], &state))

	for i := 0; i < n; i++ {
		childCertificates := types.ChildCertificates{
			Issuer:         strconv.Itoa(i),
			AuthorityKeyId: strconv.Itoa(i),
		}
		nullify.Fill(&childCertificates)
		state.ChildCertificatesList = append(state.ChildCertificatesList, childCertificates)
	}
	buf, err := cfg.Codec.MarshalJSON(&state)
	require.NoError(t, err)
	cfg.GenesisState[types.ModuleName] = buf
	return network.New(t, cfg), state.ChildCertificatesList
}

func TestShowChildCertificates(t *testing.T) {
	net, objs := networkWithChildCertificatesObjects(t, 2)

	ctx := net.Validators[0].ClientCtx
	common := []string{
		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	}
	for _, tc := range []struct {
		desc             string
		idIssuer         string
		idAuthorityKeyId string

		args []string
		err  error
		obj  types.ChildCertificates
	}{
		{
			desc:             "found",
			idIssuer:         objs[0].Issuer,
			idAuthorityKeyId: objs[0].AuthorityKeyId,

			args: common,
			obj:  objs[0],
		},
		{
			desc:             "not found",
			idIssuer:         strconv.Itoa(100000),
			idAuthorityKeyId: strconv.Itoa(100000),

			args: common,
			err:  status.Error(codes.InvalidArgument, "not found"),
		},
	} {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			args := []string{
				tc.idIssuer,
				tc.idAuthorityKeyId,
			}
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdShowChildCertificates(), args)
			if tc.err != nil {
				stat, ok := status.FromError(tc.err)
				require.True(t, ok)
				require.ErrorIs(t, stat.Err(), tc.err)
			} else {
				require.NoError(t, err)
				var resp types.QueryGetChildCertificatesResponse
				require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.NotNil(t, resp.ChildCertificates)
				require.Equal(t,
					nullify.Fill(&tc.obj),
					nullify.Fill(&resp.ChildCertificates),
				)
			}
		})
	}
}
