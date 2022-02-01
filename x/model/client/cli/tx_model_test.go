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
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/model/client/cli"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/model/types"
)

// Prevent strconv unused error.
var _ = strconv.IntSize

func TestCreateModel(t *testing.T) {
	net := network.New(t)
	val := net.Validators[0]
	ctx := val.ClientCtx

	fields := []string{
		fmt.Sprintf("--%s=%v", cli.FlagDeviceTypeId, testconstants.DeviceTypeId),
		fmt.Sprintf("--%s=%v", cli.FlagProductName, testconstants.ProductName),
		fmt.Sprintf("--%s=%v", cli.FlagProductLabel, testconstants.ProductLabel),
		fmt.Sprintf("--%s=%v", cli.FlagPartNumber, testconstants.PartNumber),
		fmt.Sprintf("--%s=%v", cli.FlagCommissioningCustomFlow, testconstants.CommissioningCustomFlow),
		fmt.Sprintf("--%s=%v", cli.FlagCommissioningCustomFlowUrl, testconstants.CommissioningCustomFlowUrl),
		fmt.Sprintf("--%s=%v", cli.FlagCommissioningModeInitialStepsHint, testconstants.CommissioningModeInitialStepsHint),
		fmt.Sprintf("--%s=%v", cli.FlagCommissioningModeInitialStepsInstruction, testconstants.CommissioningModeInitialStepsInstruction),
		fmt.Sprintf("--%s=%v", cli.FlagCommissioningModeSecondaryStepsHint, testconstants.CommissioningModeSecondaryStepsHint),
		fmt.Sprintf("--%s=%v", cli.FlagCommissioningModeSecondaryStepsInstruction, testconstants.CommissioningModeSecondaryStepsInstruction),
		fmt.Sprintf("--%s=%v", cli.FlagUserManualUrl, testconstants.UserManualUrl),
		fmt.Sprintf("--%s=%v", cli.FlagSupportUrl, testconstants.SupportUrl),
		fmt.Sprintf("--%s=%v", cli.FlagProductUrl, testconstants.ProductUrl),
		fmt.Sprintf("--%s=%v", cli.FlagLsfUrl, testconstants.LsfUrl),
		fmt.Sprintf("--%s=%v", cli.FlagLsfRevision, testconstants.LsfRevision),
	}
	common := []string{
		fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
		fmt.Sprintf("--%s", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
	}

	for _, tc := range []struct {
		desc  string
		idVid int32
		idPid int32

		err error
	}{
		{
			desc:  "valid",
			idVid: testconstants.Vid,
			idPid: testconstants.Pid,
		},
	} {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			args := []string{
				fmt.Sprintf("--%s=%v", cli.FlagVid, tc.idVid),
				fmt.Sprintf("--%s=%v", cli.FlagPid, tc.idPid),
			}
			args = append(args, fields...)
			args = append(args, common...)
			_, err := testcli.ExecTestCLITxCmd(t, ctx, cli.CmdCreateModel(), args)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestUpdateModel(t *testing.T) {
	net := network.New(t)
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
		fmt.Sprintf("--%s=%v", cli.FlagDeviceTypeId, testconstants.DeviceTypeId),
		fmt.Sprintf("--%s=%v", cli.FlagProductName, testconstants.ProductName),
		fmt.Sprintf("--%s=%v", cli.FlagProductLabel, testconstants.ProductLabel),
		fmt.Sprintf("--%s=%v", cli.FlagPartNumber, testconstants.PartNumber),
		fmt.Sprintf("--%s=%v", cli.FlagCommissioningCustomFlow, testconstants.CommissioningCustomFlow),
		fmt.Sprintf("--%s=%v", cli.FlagCommissioningCustomFlowUrl, testconstants.CommissioningCustomFlowUrl),
		fmt.Sprintf("--%s=%v", cli.FlagCommissioningModeInitialStepsHint, testconstants.CommissioningModeInitialStepsHint),
		fmt.Sprintf("--%s=%v", cli.FlagCommissioningModeInitialStepsInstruction, testconstants.CommissioningModeInitialStepsInstruction),
		fmt.Sprintf("--%s=%v", cli.FlagCommissioningModeSecondaryStepsHint, testconstants.CommissioningModeSecondaryStepsHint),
		fmt.Sprintf("--%s=%v", cli.FlagCommissioningModeSecondaryStepsInstruction, testconstants.CommissioningModeSecondaryStepsInstruction),
		fmt.Sprintf("--%s=%v", cli.FlagUserManualUrl, testconstants.UserManualUrl),
		fmt.Sprintf("--%s=%v", cli.FlagSupportUrl, testconstants.SupportUrl),
		fmt.Sprintf("--%s=%v", cli.FlagProductUrl, testconstants.ProductUrl),
		fmt.Sprintf("--%s=%v", cli.FlagLsfUrl, testconstants.LsfUrl),
		fmt.Sprintf("--%s=%v", cli.FlagLsfRevision, testconstants.LsfRevision),
	}
	args = append(args, common...)
	_, err := testcli.ExecTestCLITxCmd(t, ctx, cli.CmdCreateModel(), args)
	require.NoError(t, err)

	fields := []string{
		fmt.Sprintf("--%s=%v", cli.FlagProductName, testconstants.ProductName) + "-updated",
		fmt.Sprintf("--%s=%v", cli.FlagProductLabel, testconstants.ProductLabel) + "-updated",
		fmt.Sprintf("--%s=%v", cli.FlagPartNumber, testconstants.PartNumber) + "-updated",
		fmt.Sprintf("--%s=%v", cli.FlagCommissioningCustomFlowUrl, testconstants.CommissioningCustomFlowUrl) + "/updated",
		fmt.Sprintf("--%s=%v", cli.FlagCommissioningModeInitialStepsInstruction, testconstants.CommissioningModeInitialStepsInstruction) + "-updated",
		fmt.Sprintf("--%s=%v", cli.FlagCommissioningModeSecondaryStepsInstruction, testconstants.CommissioningModeSecondaryStepsInstruction) + "-updated",
		fmt.Sprintf("--%s=%v", cli.FlagUserManualUrl, testconstants.UserManualUrl) + "/updated",
		fmt.Sprintf("--%s=%v", cli.FlagSupportUrl, testconstants.SupportUrl) + "/updated",
		fmt.Sprintf("--%s=%v", cli.FlagProductUrl, testconstants.ProductUrl) + "/updated",
		fmt.Sprintf("--%s=%v", cli.FlagLsfUrl, testconstants.LsfUrl) + "/updated",
		fmt.Sprintf("--%s=%v", cli.FlagLsfRevision, testconstants.LsfRevision+1),
	}

	for _, tc := range []struct {
		desc  string
		idVid int32
		idPid int32

		err error
	}{
		{
			desc:  "valid",
			idVid: testconstants.Vid,
			idPid: testconstants.Pid,
		},
		{
			desc:  "model does not exist",
			idVid: testconstants.Vid,
			idPid: testconstants.Pid + 1,

			err: types.ErrModelDoesNotExist,
		},
	} {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			args := []string{
				fmt.Sprintf("--%s=%v", cli.FlagVid, tc.idVid),
				fmt.Sprintf("--%s=%v", cli.FlagPid, tc.idPid),
			}
			args = append(args, fields...)
			args = append(args, common...)
			_, err := testcli.ExecTestCLITxCmd(t, ctx, cli.CmdUpdateModel(), args)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestDeleteModel(t *testing.T) {
	net := network.New(t)

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
		fmt.Sprintf("--%s=%v", cli.FlagDeviceTypeId, testconstants.DeviceTypeId),
		fmt.Sprintf("--%s=%v", cli.FlagProductName, testconstants.ProductName),
		fmt.Sprintf("--%s=%v", cli.FlagProductLabel, testconstants.ProductLabel),
		fmt.Sprintf("--%s=%v", cli.FlagPartNumber, testconstants.PartNumber),
		fmt.Sprintf("--%s=%v", cli.FlagCommissioningCustomFlow, testconstants.CommissioningCustomFlow),
		fmt.Sprintf("--%s=%v", cli.FlagCommissioningCustomFlowUrl, testconstants.CommissioningCustomFlowUrl),
		fmt.Sprintf("--%s=%v", cli.FlagCommissioningModeInitialStepsHint, testconstants.CommissioningModeInitialStepsHint),
		fmt.Sprintf("--%s=%v", cli.FlagCommissioningModeInitialStepsInstruction, testconstants.CommissioningModeInitialStepsInstruction),
		fmt.Sprintf("--%s=%v", cli.FlagCommissioningModeSecondaryStepsHint, testconstants.CommissioningModeSecondaryStepsHint),
		fmt.Sprintf("--%s=%v", cli.FlagCommissioningModeSecondaryStepsInstruction, testconstants.CommissioningModeSecondaryStepsInstruction),
		fmt.Sprintf("--%s=%v", cli.FlagUserManualUrl, testconstants.UserManualUrl),
		fmt.Sprintf("--%s=%v", cli.FlagSupportUrl, testconstants.SupportUrl),
		fmt.Sprintf("--%s=%v", cli.FlagProductUrl, testconstants.ProductUrl),
		fmt.Sprintf("--%s=%v", cli.FlagLsfUrl, testconstants.LsfUrl),
		fmt.Sprintf("--%s=%v", cli.FlagLsfRevision, testconstants.LsfRevision),
	}
	args = append(args, common...)
	_, err := testcli.ExecTestCLITxCmd(t, ctx, cli.CmdCreateModel(), args)
	require.NoError(t, err)

	for _, tc := range []struct {
		desc  string
		idVid int32
		idPid int32

		err error
	}{
		{
			desc:  "valid",
			idVid: testconstants.Vid,
			idPid: testconstants.Pid,
		},
		{
			desc:  "model does not exist",
			idVid: testconstants.Vid,
			idPid: testconstants.Pid + 1,

			err: types.ErrModelDoesNotExist,
		},
	} {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			args := []string{
				fmt.Sprintf("--%s=%v", cli.FlagVid, tc.idVid),
				fmt.Sprintf("--%s=%v", cli.FlagPid, tc.idPid),
			}
			args = append(args, common...)
			_, err := testcli.ExecTestCLITxCmd(t, ctx, cli.CmdDeleteModel(), args)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

// TODO: Add negative tests for absence of required parameters and
// for presence of unexpected parameters (including positional ones).
