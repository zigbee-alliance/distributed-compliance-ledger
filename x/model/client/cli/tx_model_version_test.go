package cli_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/stretchr/testify/require"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	testcli "github.com/zigbee-alliance/distributed-compliance-ledger/testutil/cli"
	"github.com/zigbee-alliance/distributed-compliance-ledger/testutil/network"
	"github.com/zigbee-alliance/distributed-compliance-ledger/testutil/nullify"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/model/client/cli"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/model/types"
)

// Prevent strconv unused error.
var _ = strconv.IntSize

func networkWithPreconditions(t *testing.T) *network.Network {
	t.Helper()
	cfg := network.DefaultConfig()
	state := types.GenesisState{}
	require.NoError(t, cfg.Codec.UnmarshalJSON(cfg.GenesisState[types.ModuleName], &state))

	model := types.Model{
		Vid: testconstants.Vid,
		Pid: testconstants.Pid,
	}
	nullify.Fill(&model)
	state.ModelList = append(state.ModelList, model)

	buf, err := cfg.Codec.MarshalJSON(&state)
	require.NoError(t, err)
	cfg.GenesisState[types.ModuleName] = buf
	return network.New(t, cfg)
}

func TestCreateModelVersion(t *testing.T) {
	net := networkWithPreconditions(t)
	val := net.Validators[0]
	ctx := val.ClientCtx

	fields := []string{
		fmt.Sprintf("--%s=%v", cli.FlagSoftwareVersionString, testconstants.SoftwareVersionString),
		fmt.Sprintf("--%s=%v", cli.FlagCdVersionNumber, testconstants.CdVersionNumber),
		fmt.Sprintf("--%s=%v", cli.FlagFirmwareDigests, testconstants.FirmwareDigests),
		fmt.Sprintf("--%s=%v", cli.FlagSoftwareVersionValid, testconstants.SoftwareVersionValid),
		fmt.Sprintf("--%s=%v", cli.FlagOtaUrl, testconstants.OtaUrl),
		fmt.Sprintf("--%s=%v", cli.FlagOtaFileSize, testconstants.OtaFileSize),
		fmt.Sprintf("--%s=%v", cli.FlagOtaChecksum, testconstants.OtaChecksum),
		fmt.Sprintf("--%s=%v", cli.FlagOtaChecksumType, testconstants.OtaChecksumType),
		fmt.Sprintf("--%s=%v", cli.FlagMinApplicableSoftwareVersion, testconstants.MinApplicableSoftwareVersion),
		fmt.Sprintf("--%s=%v", cli.FlagMaxApplicableSoftwareVersion, testconstants.MaxApplicableSoftwareVersion),
		fmt.Sprintf("--%s=%v", cli.FlagReleaseNotesUrl, testconstants.ReleaseNotesUrl),
	}
	common := []string{
		fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
	}

	for _, tc := range []struct {
		desc              string
		idVid             int32
		idPid             int32
		idSoftwareVersion uint32

		err error
	}{
		{
			desc:              "valid",
			idVid:             testconstants.Vid,
			idPid:             testconstants.Pid,
			idSoftwareVersion: testconstants.SoftwareVersion,
		},
	} {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			args := []string{
				fmt.Sprintf("--%s=%v", cli.FlagVid, tc.idVid),
				fmt.Sprintf("--%s=%v", cli.FlagPid, tc.idPid),
				fmt.Sprintf("--%s=%v", cli.FlagSoftwareVersion, tc.idSoftwareVersion),
			}
			args = append(args, fields...)
			args = append(args, common...)
			_, err := testcli.ExecTestCLICmd(t, ctx, cli.CmdCreateModelVersion(), args)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestUpdateModelVersion(t *testing.T) {
	net := networkWithPreconditions(t)
	val := net.Validators[0]
	ctx := val.ClientCtx

	common := []string{
		fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
	}

	args := []string{
		fmt.Sprintf("--%s=%v", cli.FlagVid, testconstants.Vid),
		fmt.Sprintf("--%s=%v", cli.FlagPid, testconstants.Pid),
		fmt.Sprintf("--%s=%v", cli.FlagSoftwareVersion, testconstants.SoftwareVersion),
		fmt.Sprintf("--%s=%v", cli.FlagSoftwareVersionString, testconstants.SoftwareVersionString),
		fmt.Sprintf("--%s=%v", cli.FlagCdVersionNumber, testconstants.CdVersionNumber),
		fmt.Sprintf("--%s=%v", cli.FlagFirmwareDigests, testconstants.FirmwareDigests),
		fmt.Sprintf("--%s=%v", cli.FlagSoftwareVersionValid, testconstants.SoftwareVersionValid),
		fmt.Sprintf("--%s=%v", cli.FlagOtaUrl, testconstants.OtaUrl),
		fmt.Sprintf("--%s=%v", cli.FlagOtaFileSize, testconstants.OtaFileSize),
		fmt.Sprintf("--%s=%v", cli.FlagOtaChecksum, testconstants.OtaChecksum),
		fmt.Sprintf("--%s=%v", cli.FlagOtaChecksumType, testconstants.OtaChecksumType),
		fmt.Sprintf("--%s=%v", cli.FlagMinApplicableSoftwareVersion, testconstants.MinApplicableSoftwareVersion),
		fmt.Sprintf("--%s=%v", cli.FlagMaxApplicableSoftwareVersion, testconstants.MaxApplicableSoftwareVersion),
		fmt.Sprintf("--%s=%v", cli.FlagReleaseNotesUrl, testconstants.ReleaseNotesUrl),
	}
	args = append(args, common...)
	_, err := testcli.ExecTestCLICmd(t, ctx, cli.CmdCreateModelVersion(), args)
	require.NoError(t, err)

	fields := []string{
		fmt.Sprintf("--%s=%v", cli.FlagSoftwareVersionValid, !testconstants.SoftwareVersionValid),
		fmt.Sprintf("--%s=%v", cli.FlagOtaUrl, testconstants.OtaUrl+"/updated"),
		fmt.Sprintf("--%s=%v", cli.FlagMinApplicableSoftwareVersion, testconstants.MinApplicableSoftwareVersion+1),
		fmt.Sprintf("--%s=%v", cli.FlagMaxApplicableSoftwareVersion, testconstants.MaxApplicableSoftwareVersion+1),
		fmt.Sprintf("--%s=%v", cli.FlagReleaseNotesUrl, testconstants.ReleaseNotesUrl+"/updated"),
	}

	for _, tc := range []struct {
		desc              string
		idVid             int32
		idPid             int32
		idSoftwareVersion uint32

		err error
	}{
		{
			desc:              "valid",
			idVid:             testconstants.Vid,
			idPid:             testconstants.Pid,
			idSoftwareVersion: testconstants.SoftwareVersion,
		},
		{
			desc:              "model version does not exist",
			idVid:             testconstants.Vid,
			idPid:             testconstants.Pid + 1,
			idSoftwareVersion: testconstants.SoftwareVersion + 1,

			err: types.ErrModelVersionDoesNotExist,
		},
	} {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			args := []string{
				fmt.Sprintf("--%s=%v", cli.FlagVid, tc.idVid),
				fmt.Sprintf("--%s=%v", cli.FlagPid, tc.idPid),
				fmt.Sprintf("--%s=%v", cli.FlagSoftwareVersion, tc.idSoftwareVersion),
			}
			args = append(args, fields...)
			args = append(args, common...)
			_, err := testcli.ExecTestCLICmd(t, ctx, cli.CmdUpdateModelVersion(), args)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
