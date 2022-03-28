package types

import (
	fmt "fmt"
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	tmrand "github.com/tendermint/tendermint/libs/rand"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	"github.com/zigbee-alliance/distributed-compliance-ledger/testutil/sample"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/validator"
)

func TestMsgCertifyModel_ValidateBasic(t *testing.T) {
	negativeTests := []struct {
		name string
		msg  MsgCertifyModel
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCertifyModel{
				Signer:                "invalid_address",
				SoftwareVersionString: testconstants.SoftwareVersionString,
				CertificationDate:     testconstants.CertificationDate,
				CertificationType:     testconstants.CertificationType,
				SoftwareVersion:       testconstants.SoftwareVersion,
				CDVersionNumber:       uint32(testconstants.CdVersionNumber),
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "vid is 0",
			msg: MsgCertifyModel{
				Signer:                sample.AccAddress(),
				Vid:                   0,
				Pid:                   1,
				SoftwareVersionString: testconstants.SoftwareVersionString,
				CertificationDate:     testconstants.CertificationDate,
				CertificationType:     testconstants.CertificationType,
				SoftwareVersion:       testconstants.SoftwareVersion,
				CDVersionNumber:       uint32(testconstants.CdVersionNumber),
			},
			err: validator.ErrFieldLowerBoundViolated,
		},
		{
			name: "vid < 0",
			msg: MsgCertifyModel{
				Signer:                sample.AccAddress(),
				Vid:                   -1,
				Pid:                   1,
				SoftwareVersionString: testconstants.SoftwareVersionString,
				CertificationDate:     testconstants.CertificationDate,
				CertificationType:     testconstants.CertificationType,
				SoftwareVersion:       testconstants.SoftwareVersion,
				CDVersionNumber:       uint32(testconstants.CdVersionNumber),
			},
			err: validator.ErrFieldLowerBoundViolated,
		},
		{
			name: "vid > 65535",
			msg: MsgCertifyModel{
				Signer:                sample.AccAddress(),
				Vid:                   65536,
				Pid:                   1,
				SoftwareVersionString: testconstants.SoftwareVersionString,
				CertificationDate:     testconstants.CertificationDate,
				CertificationType:     testconstants.CertificationType,
				SoftwareVersion:       testconstants.SoftwareVersion,
				CDVersionNumber:       uint32(testconstants.CdVersionNumber),
			},
			err: validator.ErrFieldUpperBoundViolated,
		},
		{
			name: "pid is 0",
			msg: MsgCertifyModel{
				Signer:                sample.AccAddress(),
				Pid:                   0,
				Vid:                   1,
				SoftwareVersionString: testconstants.SoftwareVersionString,
				CertificationDate:     testconstants.CertificationDate,
				CertificationType:     testconstants.CertificationType,
				SoftwareVersion:       testconstants.SoftwareVersion,
				CDVersionNumber:       uint32(testconstants.CdVersionNumber),
			},
			err: validator.ErrFieldLowerBoundViolated,
		},
		{
			name: "pid < 0",
			msg: MsgCertifyModel{
				Signer:                sample.AccAddress(),
				Pid:                   -1,
				Vid:                   1,
				SoftwareVersionString: testconstants.SoftwareVersionString,
				CertificationDate:     testconstants.CertificationDate,
				CertificationType:     testconstants.CertificationType,
				SoftwareVersion:       testconstants.SoftwareVersion,
				CDVersionNumber:       uint32(testconstants.CdVersionNumber),
			},
			err: validator.ErrFieldLowerBoundViolated,
		},
		{
			name: "pid > 65535",
			msg: MsgCertifyModel{
				Signer:                sample.AccAddress(),
				Pid:                   65536,
				Vid:                   1,
				SoftwareVersionString: testconstants.SoftwareVersionString,
				CertificationDate:     testconstants.CertificationDate,
				CertificationType:     testconstants.CertificationType,
				SoftwareVersion:       testconstants.SoftwareVersion,
				CDVersionNumber:       uint32(testconstants.CdVersionNumber),
			},
			err: validator.ErrFieldUpperBoundViolated,
		},
		{
			name: "cd version number > 65535",
			msg: MsgCertifyModel{
				Signer:                sample.AccAddress(),
				CDVersionNumber:       65536,
				Pid:                   1,
				Vid:                   1,
				SoftwareVersionString: testconstants.SoftwareVersionString,
				CertificationDate:     testconstants.CertificationDate,
				CertificationType:     testconstants.CertificationType,
				SoftwareVersion:       testconstants.SoftwareVersion,
			},
			err: validator.ErrFieldUpperBoundViolated,
		},
		{
			name: "certification date not set",
			msg: MsgCertifyModel{
				Signer:                sample.AccAddress(),
				Pid:                   1,
				Vid:                   1,
				SoftwareVersionString: testconstants.SoftwareVersionString,
				CertificationDate:     "",
				CertificationType:     testconstants.CertificationType,
				CDVersionNumber:       uint32(testconstants.CdVersionNumber),
			},
			err: validator.ErrRequiredFieldMissing,
		},
		{
			name: "certification type not set",
			msg: MsgCertifyModel{
				Signer:                sample.AccAddress(),
				Pid:                   1,
				Vid:                   1,
				SoftwareVersionString: testconstants.SoftwareVersionString,
				CertificationDate:     testconstants.CertificationDate,
				CertificationType:     "",
				CDVersionNumber:       uint32(testconstants.CdVersionNumber),
			},
			err: validator.ErrRequiredFieldMissing,
		},
		{
			name: "certification date is not RFC3339",
			msg: MsgCertifyModel{
				Signer:                sample.AccAddress(),
				Pid:                   1,
				Vid:                   1,
				SoftwareVersionString: testconstants.SoftwareVersionString,
				CertificationDate:     "2020-01-01",
				CertificationType:     testconstants.TestResult,
				CDVersionNumber:       uint32(testconstants.CdVersionNumber),
			},
			err: ErrInvalidTestDateFormat,
		},
		{
			name: "invalid certification type",
			msg: MsgCertifyModel{
				Signer:                sample.AccAddress(),
				Pid:                   1,
				Vid:                   1,
				SoftwareVersionString: testconstants.SoftwareVersionString,
				CertificationDate:     testconstants.CertificationDate,
				CertificationType:     "invalid certification type",
				CDVersionNumber:       uint32(testconstants.CdVersionNumber),
				Reason:                testconstants.Reason,
			},
			err: ErrInvalidCertificationType,
		},
		{
			name: "software version string len > 64",
			msg: MsgCertifyModel{
				Signer:                sample.AccAddress(),
				Pid:                   1,
				Vid:                   1,
				SoftwareVersionString: fmt.Sprintf("1.%063d", 0),
				CertificationDate:     testconstants.CertificationDate,
				CertificationType:     testconstants.CertificationType,
				CDVersionNumber:       uint32(testconstants.CdVersionNumber),
				Reason:                testconstants.Reason,
			},
			err: validator.ErrFieldMaxLengthExceeded,
		},
		{
			name: "reason len > 102400 (100 KB)",
			msg: MsgCertifyModel{
				Signer:                sample.AccAddress(),
				Pid:                   1,
				Vid:                   1,
				SoftwareVersionString: testconstants.TestDate,
				CertificationDate:     testconstants.CertificationDate,
				CertificationType:     testconstants.CertificationType,
				CDVersionNumber:       uint32(testconstants.CdVersionNumber),
				Reason:                tmrand.Str(102401),
			},
			err: validator.ErrFieldMaxLengthExceeded,
		},
	}

	positiveTests := []struct {
		name string
		msg  MsgCertifyModel
	}{
		{
			name: "valid certification model msg",
			msg: MsgCertifyModel{
				Signer:                sample.AccAddress(),
				SoftwareVersionString: testconstants.SoftwareVersionString,
				Pid:                   1,
				Vid:                   1,
				CertificationDate:     testconstants.CertificationDate,
				CertificationType:     testconstants.CertificationType,
				SoftwareVersion:       testconstants.SoftwareVersion,
				CDVersionNumber:       uint32(testconstants.CdVersionNumber),
			},
		},
		{
			name: "software version = 0",
			msg: MsgCertifyModel{
				Signer:                sample.AccAddress(),
				SoftwareVersionString: testconstants.SoftwareVersionString,
				Pid:                   1,
				Vid:                   1,
				CertificationDate:     testconstants.CertificationDate,
				CertificationType:     testconstants.CertificationType,
				SoftwareVersion:       0,
				CDVersionNumber:       uint32(testconstants.CdVersionNumber),
			},
		},
		{
			name: "cd version number = 0",
			msg: MsgCertifyModel{
				Signer:                sample.AccAddress(),
				SoftwareVersionString: testconstants.SoftwareVersionString,
				Pid:                   1,
				Vid:                   1,
				CertificationDate:     testconstants.CertificationDate,
				CertificationType:     testconstants.CertificationType,
				SoftwareVersion:       1,
				CDVersionNumber:       0,
			},
		},
		{
			name: "certification type is zigbee",
			msg: MsgCertifyModel{
				Signer:                sample.AccAddress(),
				SoftwareVersionString: testconstants.SoftwareVersionString,
				Pid:                   1,
				Vid:                   1,
				CertificationDate:     testconstants.CertificationDate,
				CertificationType:     "zigbee",
				SoftwareVersion:       testconstants.SoftwareVersion,
				CDVersionNumber:       uint32(testconstants.CdVersionNumber),
			},
		},
		{
			name: "certification type is matter",
			msg: MsgCertifyModel{
				Signer:                sample.AccAddress(),
				SoftwareVersionString: testconstants.SoftwareVersionString,
				Pid:                   1,
				Vid:                   1,
				CertificationDate:     testconstants.CertificationDate,
				CertificationType:     "matter",
				SoftwareVersion:       testconstants.SoftwareVersion,
				CDVersionNumber:       uint32(testconstants.CdVersionNumber),
			},
		},
		{
			name: "minimal pid, vid values",
			msg: MsgCertifyModel{
				Signer:                sample.AccAddress(),
				SoftwareVersionString: "1",
				Pid:                   1,
				Vid:                   1,
				CertificationDate:     testconstants.CertificationDate,
				CertificationType:     testconstants.CertificationType,
				SoftwareVersion:       0,
				CDVersionNumber:       uint32(testconstants.CdVersionNumber),
			},
		},
		{
			name: "max pid, vid values",
			msg: MsgCertifyModel{
				Signer:                sample.AccAddress(),
				SoftwareVersionString: "1",
				Pid:                   65535,
				Vid:                   65535,
				CertificationDate:     testconstants.CertificationDate,
				CertificationType:     testconstants.CertificationType,
				SoftwareVersion:       0,
				CDVersionNumber:       uint32(testconstants.CdVersionNumber),
			},
		},
	}

	for _, tt := range negativeTests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			require.Error(t, err)
			require.ErrorIs(t, err, tt.err)
		})
	}

	for _, tt := range positiveTests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			require.NoError(t, err)
		})
	}
}
