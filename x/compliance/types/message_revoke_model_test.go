package types

import (
	fmt "fmt"
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	"github.com/zigbee-alliance/distributed-compliance-ledger/testutil/sample"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/validator"
)

func TestMsgRevokeModel_ValidateBasic(t *testing.T) {
	negative_tests := []struct {
		name string
		msg  MsgRevokeModel
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgRevokeModel{
				Signer:                "invalid_address",
				SoftwareVersionString: testconstants.SoftwareVersionString,
				RevocationDate:        testconstants.CertificationDate,
				CertificationType:     testconstants.CertificationType,
				SoftwareVersion:       testconstants.SoftwareVersion,
				CDVersionNumber:       uint32(testconstants.CdVersionNumber),
				Reason:                testconstants.Reason,
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "vid is 0",
			msg: MsgRevokeModel{
				Signer:                sample.AccAddress(),
				Vid:                   0,
				Pid:                   1,
				SoftwareVersionString: testconstants.SoftwareVersionString,
				RevocationDate:        testconstants.CertificationDate,
				CertificationType:     testconstants.CertificationType,
				SoftwareVersion:       testconstants.SoftwareVersion,
				CDVersionNumber:       uint32(testconstants.CdVersionNumber),
				Reason:                testconstants.Reason,
			},
			err: validator.ErrFieldLowerBoundViolated,
		},
		{
			name: "vid < 0",
			msg: MsgRevokeModel{
				Signer:                sample.AccAddress(),
				Vid:                   -1,
				Pid:                   1,
				SoftwareVersionString: testconstants.SoftwareVersionString,
				RevocationDate:        testconstants.CertificationDate,
				CertificationType:     testconstants.CertificationType,
				SoftwareVersion:       testconstants.SoftwareVersion,
				CDVersionNumber:       uint32(testconstants.CdVersionNumber),
				Reason:                testconstants.Reason,
			},
			err: validator.ErrFieldLowerBoundViolated,
		},
		{
			name: "vid > 65535",
			msg: MsgRevokeModel{
				Signer:                sample.AccAddress(),
				Vid:                   65536,
				Pid:                   1,
				SoftwareVersionString: testconstants.SoftwareVersionString,
				RevocationDate:        testconstants.CertificationDate,
				CertificationType:     testconstants.CertificationType,
				SoftwareVersion:       testconstants.SoftwareVersion,
				CDVersionNumber:       uint32(testconstants.CdVersionNumber),
				Reason:                testconstants.Reason,
			},
			err: validator.ErrFieldUpperBoundViolated,
		},
		{
			name: "pid is 0",
			msg: MsgRevokeModel{
				Signer:                sample.AccAddress(),
				Pid:                   0,
				Vid:                   1,
				SoftwareVersionString: testconstants.SoftwareVersionString,
				RevocationDate:        testconstants.CertificationDate,
				CertificationType:     testconstants.CertificationType,
				SoftwareVersion:       testconstants.SoftwareVersion,
				CDVersionNumber:       uint32(testconstants.CdVersionNumber),
				Reason:                testconstants.Reason,
			},
			err: validator.ErrFieldLowerBoundViolated,
		},
		{
			name: "pid < 0",
			msg: MsgRevokeModel{
				Signer:                sample.AccAddress(),
				Pid:                   -1,
				Vid:                   1,
				SoftwareVersionString: testconstants.SoftwareVersionString,
				RevocationDate:        testconstants.CertificationDate,
				CertificationType:     testconstants.CertificationType,
				SoftwareVersion:       testconstants.SoftwareVersion,
				CDVersionNumber:       uint32(testconstants.CdVersionNumber),
				Reason:                testconstants.Reason,
			},
			err: validator.ErrFieldLowerBoundViolated,
		},
		{
			name: "pid > 65535",
			msg: MsgRevokeModel{
				Signer:                sample.AccAddress(),
				Pid:                   65536,
				Vid:                   1,
				SoftwareVersionString: testconstants.SoftwareVersionString,
				RevocationDate:        testconstants.CertificationDate,
				CertificationType:     testconstants.CertificationType,
				SoftwareVersion:       testconstants.SoftwareVersion,
				CDVersionNumber:       uint32(testconstants.CdVersionNumber),
				Reason:                testconstants.Reason,
			},
			err: validator.ErrFieldUpperBoundViolated,
		},
		{
			name: "cd version number > 65535",
			msg: MsgRevokeModel{
				Signer:                sample.AccAddress(),
				CDVersionNumber:       65536,
				Pid:                   1,
				Vid:                   1,
				SoftwareVersionString: testconstants.SoftwareVersionString,
				RevocationDate:        testconstants.CertificationDate,
				CertificationType:     testconstants.CertificationType,
				SoftwareVersion:       testconstants.SoftwareVersion,
				Reason:                testconstants.Reason,
			},
			err: validator.ErrFieldUpperBoundViolated,
		},
		{
			name: "revocation date not set",
			msg: MsgRevokeModel{
				Signer:                sample.AccAddress(),
				Pid:                   1,
				Vid:                   1,
				SoftwareVersionString: testconstants.SoftwareVersionString,
				RevocationDate:        "",
				CertificationType:     testconstants.CertificationType,
				CDVersionNumber:       uint32(testconstants.CdVersionNumber),
				Reason:                testconstants.Reason,
			},
			err: validator.ErrRequiredFieldMissing,
		},
		{
			name: "certification type not set",
			msg: MsgRevokeModel{
				Signer:                sample.AccAddress(),
				Pid:                   1,
				Vid:                   1,
				SoftwareVersionString: testconstants.SoftwareVersionString,
				RevocationDate:        testconstants.CertificationDate,
				CertificationType:     "",
				CDVersionNumber:       uint32(testconstants.CdVersionNumber),
				Reason:                testconstants.Reason,
			},
			err: validator.ErrRequiredFieldMissing,
		},
		{
			name: "revocation date is not RFC3339",
			msg: MsgRevokeModel{
				Signer:                sample.AccAddress(),
				Pid:                   1,
				Vid:                   1,
				SoftwareVersionString: testconstants.SoftwareVersionString,
				RevocationDate:        "2020-01-01",
				CertificationType:     testconstants.TestResult,
				CDVersionNumber:       uint32(testconstants.CdVersionNumber),
				Reason:                testconstants.Reason,
			},
			err: ErrInvalidTestDateFormat,
		},
		{
			name: "invalid certification type",
			msg: MsgRevokeModel{
				Signer:                sample.AccAddress(),
				Pid:                   1,
				Vid:                   1,
				SoftwareVersionString: testconstants.SoftwareVersionString,
				RevocationDate:        testconstants.CertificationDate,
				CertificationType:     "invalid certification type",
				CDVersionNumber:       uint32(testconstants.CdVersionNumber),
				Reason:                testconstants.Reason,
			},
			err: ErrInvalidCertificationType,
		},
		{
			name: "software version string len > 64",
			msg: MsgRevokeModel{
				Signer:                sample.AccAddress(),
				Pid:                   1,
				Vid:                   1,
				SoftwareVersionString: fmt.Sprintf("1.%063d", 0),
				RevocationDate:        testconstants.CertificationDate,
				CertificationType:     testconstants.CertificationType,
				CDVersionNumber:       uint32(testconstants.CdVersionNumber),
				Reason:                testconstants.Reason,
			},
			err: validator.ErrFieldMaxLengthExceeded,
		},
	}

	positive_tests := []struct {
		name string
		msg  MsgRevokeModel
	}{
		{
			name: "valid revoke model msg",
			msg: MsgRevokeModel{
				Signer:                sample.AccAddress(),
				SoftwareVersionString: testconstants.SoftwareVersionString,
				Pid:                   1,
				Vid:                   1,
				RevocationDate:        testconstants.CertificationDate,
				CertificationType:     testconstants.CertificationType,
				SoftwareVersion:       testconstants.SoftwareVersion,
				CDVersionNumber:       uint32(testconstants.CdVersionNumber),
				Reason:                testconstants.Reason,
			},
		},
		{
			name: "software version = 0",
			msg: MsgRevokeModel{
				Signer:                sample.AccAddress(),
				SoftwareVersionString: testconstants.SoftwareVersionString,
				Pid:                   1,
				Vid:                   1,
				RevocationDate:        testconstants.CertificationDate,
				CertificationType:     testconstants.CertificationType,
				SoftwareVersion:       0,
				CDVersionNumber:       uint32(testconstants.CdVersionNumber),
				Reason:                testconstants.Reason,
			},
		},
		{
			name: "cd version number = 0",
			msg: MsgRevokeModel{
				Signer:                sample.AccAddress(),
				SoftwareVersionString: testconstants.SoftwareVersionString,
				Pid:                   1,
				Vid:                   1,
				RevocationDate:        testconstants.CertificationDate,
				CertificationType:     testconstants.CertificationType,
				SoftwareVersion:       1,
				CDVersionNumber:       0,
				Reason:                testconstants.Reason,
			},
		},
		{
			name: "certification type is zigbee",
			msg: MsgRevokeModel{
				Signer:                sample.AccAddress(),
				SoftwareVersionString: testconstants.SoftwareVersionString,
				Pid:                   1,
				Vid:                   1,
				RevocationDate:        testconstants.CertificationDate,
				CertificationType:     "zigbee",
				SoftwareVersion:       testconstants.SoftwareVersion,
				CDVersionNumber:       uint32(testconstants.CdVersionNumber),
				Reason:                testconstants.Reason,
			},
		},
		{
			name: "certification type is matter",
			msg: MsgRevokeModel{
				Signer:                sample.AccAddress(),
				SoftwareVersionString: testconstants.SoftwareVersionString,
				Pid:                   1,
				Vid:                   1,
				RevocationDate:        testconstants.CertificationDate,
				CertificationType:     "matter",
				SoftwareVersion:       testconstants.SoftwareVersion,
				CDVersionNumber:       uint32(testconstants.CdVersionNumber),
				Reason:                testconstants.Reason,
			},
		},
		{
			name: "minimal pid, vid values",
			msg: MsgRevokeModel{
				Signer:                sample.AccAddress(),
				SoftwareVersionString: "1",
				Pid:                   1,
				Vid:                   1,
				RevocationDate:        testconstants.CertificationDate,
				CertificationType:     testconstants.CertificationType,
				SoftwareVersion:       0,
				CDVersionNumber:       uint32(testconstants.CdVersionNumber),
			},
		},
		{
			name: "max pid, vid values",
			msg: MsgRevokeModel{
				Signer:                sample.AccAddress(),
				SoftwareVersionString: "1",
				Pid:                   65535,
				Vid:                   65535,
				RevocationDate:        testconstants.CertificationDate,
				CertificationType:     testconstants.CertificationType,
				SoftwareVersion:       0,
				CDVersionNumber:       uint32(testconstants.CdVersionNumber),
			},
		},
	}

	for _, tt := range negative_tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			require.Error(t, err)
			require.ErrorIs(t, err, tt.err)
		})
	}

	for _, tt := range positive_tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			require.NoError(t, err)
		})
	}
}
