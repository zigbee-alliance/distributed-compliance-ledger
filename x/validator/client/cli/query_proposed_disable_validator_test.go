package cli_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/cosmos/cosmos-sdk/client/flags"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/stretchr/testify/require"
	tmcli "github.com/tendermint/tendermint/libs/cli"

	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	"github.com/zigbee-alliance/distributed-compliance-ledger/testutil/network"
	"github.com/zigbee-alliance/distributed-compliance-ledger/testutil/nullify"
	cliutils "github.com/zigbee-alliance/distributed-compliance-ledger/utils/cli"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/validator/client/cli"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/validator/types"
)

// Prevent strconv unused error.
var _ = strconv.IntSize

func networkWithProposedDisableValidatorObjects(t *testing.T) (*network.Network, []types.ProposedDisableValidator) {
	t.Helper()
	cfg := network.DefaultConfig()
	state := types.GenesisState{}
	require.NoError(t, cfg.Codec.UnmarshalJSON(cfg.GenesisState[types.ModuleName], &state))

	proposedDisableValidator := types.ProposedDisableValidator{
		Address: testconstants.ValidatorAddress1,
	}
	nullify.Fill(&proposedDisableValidator)
	state.ProposedDisableValidatorList = append(state.ProposedDisableValidatorList, proposedDisableValidator)

	buf, err := cfg.Codec.MarshalJSON(&state)
	require.NoError(t, err)
	cfg.GenesisState[types.ModuleName] = buf
	return network.New(t, cfg), state.ProposedDisableValidatorList
}

func TestShowProposedDisableValidator(t *testing.T) {
	net, objs := networkWithProposedDisableValidatorObjects(t)

	ctx := net.Validators[0].ClientCtx
	common := []string{
		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	}
	for _, tc := range []struct {
		desc      string
		idAddress string

		args []string
		obj  *types.ProposedDisableValidator
	}{
		{
			desc:      "found",
			idAddress: objs[0].Address,

			args: common,
			obj:  &objs[0],
		},
		{
			desc:      "not found",
			idAddress: testconstants.Address1.String(),

			args: common,
			obj:  nil,
		},
	} {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			args := []string{
				fmt.Sprintf("--%s=%v", cli.FlagAddress, tc.idAddress),
			}
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdShowProposedDisableValidator(), args)
			require.NoError(t, err)
			if tc.obj == nil {
				require.Equal(t, cliutils.NotFoundOutput, out.String())
			} else {
				var proposedDisableValidator types.ProposedDisableValidator
				require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &proposedDisableValidator))
				require.Equal(t,
					nullify.Fill(&tc.obj),
					nullify.Fill(&proposedDisableValidator),
				)
			}
		})
	}
}

func TestListProposedDisableValidator(t *testing.T) {
	net, objs := networkWithProposedDisableValidatorObjects(t)

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
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListProposedDisableValidator(), args)
			require.NoError(t, err)
			var resp types.QueryAllProposedDisableValidatorResponse
			require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.LessOrEqual(t, len(resp.ProposedDisableValidator), step)
			require.Subset(t,
				nullify.Fill(objs),
				nullify.Fill(resp.ProposedDisableValidator),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(objs); i += step {
			args := request(next, 0, uint64(step), false)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListProposedDisableValidator(), args)
			require.NoError(t, err)
			var resp types.QueryAllProposedDisableValidatorResponse
			require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.LessOrEqual(t, len(resp.ProposedDisableValidator), step)
			require.Subset(t,
				nullify.Fill(objs),
				nullify.Fill(resp.ProposedDisableValidator),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		args := request(nil, 0, uint64(len(objs)), true)
		out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListProposedDisableValidator(), args)
		require.NoError(t, err)
		var resp types.QueryAllProposedDisableValidatorResponse
		require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
		require.NoError(t, err)
		require.Equal(t, len(objs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(objs),
			nullify.Fill(resp.ProposedDisableValidator),
		)
	})
}
