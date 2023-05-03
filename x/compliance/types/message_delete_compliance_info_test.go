package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	"github.com/zigbee-alliance/distributed-compliance-ledger/testutil/sample"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/validator"
)

func TestMsgDeleteComplianceInfo_ValidateBasic(t *testing.T) {
	negativeTests := []struct {
		name string
		msg  MsgDeleteComplianceInfo
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteComplianceInfo{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "vid is 0",
			msg: MsgDeleteComplianceInfo{
				Creator:           sample.AccAddress(),
				Vid:               0,
				Pid:               1,
				CertificationType: testconstants.CertificationType,
				SoftwareVersion:   testconstants.SoftwareVersion,
			},
			err: validator.ErrFieldLowerBoundViolated,
		},
		{
			name: "vid < 0",
			msg: MsgDeleteComplianceInfo{
				Creator:           sample.AccAddress(),
				Vid:               -1,
				Pid:               1,
				CertificationType: testconstants.CertificationType,
				SoftwareVersion:   testconstants.SoftwareVersion,
			},
			err: validator.ErrFieldLowerBoundViolated,
		},
		{
			name: "vid > 65535",
			msg: MsgDeleteComplianceInfo{
				Creator:           sample.AccAddress(),
				Vid:               65536,
				Pid:               1,
				CertificationType: testconstants.CertificationType,
				SoftwareVersion:   testconstants.SoftwareVersion,
			},
			err: validator.ErrFieldUpperBoundViolated,
		},
		{
			name: "pid is 0",
			msg: MsgDeleteComplianceInfo{
				Creator:           sample.AccAddress(),
				Pid:               0,
				Vid:               1,
				CertificationType: testconstants.CertificationType,
				SoftwareVersion:   testconstants.SoftwareVersion,
			},
			err: validator.ErrFieldLowerBoundViolated,
		},
		{
			name: "pid < 0",
			msg: MsgDeleteComplianceInfo{
				Creator:           sample.AccAddress(),
				Pid:               -1,
				Vid:               1,
				CertificationType: testconstants.CertificationType,
				SoftwareVersion:   testconstants.SoftwareVersion,
			},
			err: validator.ErrFieldLowerBoundViolated,
		},
		{
			name: "pid > 65535",
			msg: MsgDeleteComplianceInfo{
				Creator:           sample.AccAddress(),
				Pid:               65536,
				Vid:               1,
				CertificationType: testconstants.CertificationType,
				SoftwareVersion:   testconstants.SoftwareVersion,
			},
			err: validator.ErrFieldUpperBoundViolated,
		},
		{
			name: "invalid certification type",
			msg: MsgDeleteComplianceInfo{
				Creator:           sample.AccAddress(),
				Pid:               1,
				Vid:               1,
				CertificationType: "invalid certification type",
			},
			err: ErrInvalidCertificationType,
		},
		{
			name: "certification type not set",
			msg: MsgDeleteComplianceInfo{
				Creator:           sample.AccAddress(),
				Pid:               1,
				Vid:               1,
				CertificationType: "",
			},
			err: validator.ErrRequiredFieldMissing,
		},
	}

	positiveTests := []struct {
		name string
		msg  MsgDeleteComplianceInfo
		err  error
	}{
		{
			name: "valid provision model msg",
			msg: MsgDeleteComplianceInfo{
				Creator:           sample.AccAddress(),
				Pid:               1,
				Vid:               1,
				CertificationType: testconstants.CertificationType,
				SoftwareVersion:   testconstants.SoftwareVersion,
			},
		},
		{
			name: "software version = 0",
			msg: MsgDeleteComplianceInfo{
				Creator:           sample.AccAddress(),
				Pid:               1,
				Vid:               1,
				CertificationType: testconstants.CertificationType,
				SoftwareVersion:   0,
			},
		},
		{
			name: "cd version number = 0",
			msg: MsgDeleteComplianceInfo{
				Creator:           sample.AccAddress(),
				Pid:               1,
				Vid:               1,
				CertificationType: testconstants.CertificationType,
				SoftwareVersion:   1,
			},
		},
		{
			name: "certification type is zigbee",
			msg: MsgDeleteComplianceInfo{
				Creator:           sample.AccAddress(),
				Pid:               1,
				Vid:               1,
				CertificationType: "zigbee",
				SoftwareVersion:   testconstants.SoftwareVersion,
			},
		},
		{
			name: "certification type is matter",
			msg: MsgDeleteComplianceInfo{
				Creator:           sample.AccAddress(),
				Pid:               1,
				Vid:               1,
				CertificationType: "matter",
				SoftwareVersion:   testconstants.SoftwareVersion,
			},
		},
		{
			name: "minimal pid, vid values",
			msg: MsgDeleteComplianceInfo{
				Creator:           sample.AccAddress(),
				Pid:               1,
				Vid:               1,
				CertificationType: testconstants.CertificationType,
				SoftwareVersion:   0,
			},
		},
		{
			name: "max pid, vid values",
			msg: MsgDeleteComplianceInfo{
				Creator:           sample.AccAddress(),
				Pid:               65535,
				Vid:               65535,
				CertificationType: testconstants.CertificationType,
				SoftwareVersion:   0,
			},
		},
	}

	for _, tt := range negativeTests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)

				return
			}
			require.NoError(t, err)
		})
	}

	for _, tt := range positiveTests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)

				return
			}
			require.NoError(t, err)
		})
	}
}
