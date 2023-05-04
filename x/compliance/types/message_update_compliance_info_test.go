package types

import (
	"strconv"
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	tmrand "github.com/tendermint/tendermint/libs/rand"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	"github.com/zigbee-alliance/distributed-compliance-ledger/testutil/sample"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/validator"
)

func TestMsgUpdateComplianceInfo_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateComplianceInfo
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateComplianceInfo{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "date is not RFC3339",
			msg: MsgUpdateComplianceInfo{
				Creator:           sample.AccAddress(),
				Pid:               1,
				Vid:               1,
				Date:              "2020-01-01",
				CertificationType: testconstants.CertificationType,
				CDCertificateId:   testconstants.CDCertificateID,
			},
			err: ErrInvalidTestDateFormat,
		},
		{
			name: "invalid certification type",
			msg: MsgUpdateComplianceInfo{
				Creator:           sample.AccAddress(),
				Pid:               1,
				Vid:               1,
				Date:              testconstants.CertificationDate,
				CertificationType: "invalid certification type",
				CDVersionNumber:   "312",
				Reason:            testconstants.Reason,
				CDCertificateId:   testconstants.CDCertificateID,
			},
			err: ErrInvalidCertificationType,
		},
		{
			name: "invalid parent child field",
			msg: MsgUpdateComplianceInfo{
				Creator:           sample.AccAddress(),
				Pid:               1,
				Vid:               1,
				Date:              testconstants.CertificationDate,
				CertificationType: testconstants.CertificationType,
				CDVersionNumber:   "312",
				Reason:            testconstants.Reason,
				CDCertificateId:   testconstants.CDCertificateID,
				ParentChild:       "invalid parent child",
			},
			err: ErrInvalidPFCCertificationRoute,
		},
		{
			name: "certificationRoute > 64",
			msg: MsgUpdateComplianceInfo{
				Creator:            sample.AccAddress(),
				Pid:                1,
				Vid:                1,
				Date:               testconstants.CertificationDate,
				CertificationType:  testconstants.CertificationType,
				CDVersionNumber:    "312",
				Reason:             testconstants.Reason,
				CertificationRoute: tmrand.Str(65),
			},
			err: validator.ErrFieldMaxLengthExceeded,
		},
		{
			name: "update compliance info message with all optional fields not set",
			msg: MsgUpdateComplianceInfo{
				Creator:           sample.AccAddress(),
				Vid:               1,
				Pid:               1,
				CertificationType: testconstants.CertificationType,
				SoftwareVersion:   testconstants.SoftwareVersion,
			},
		},
		{
			name: "update compliance info message with all optional fields set",
			msg: MsgUpdateComplianceInfo{
				Creator:                            sample.AccAddress(),
				Vid:                                1,
				Pid:                                1,
				CertificationType:                  testconstants.CertificationType,
				SoftwareVersion:                    testconstants.SoftwareVersion,
				Reason:                             testconstants.Reason,
				CDVersionNumber:                    "312",
				CDCertificateId:                    testconstants.CDCertificateID,
				Date:                               testconstants.CertificationDate,
				ParentChild:                        "parent",
				CertificationRoute:                 tmrand.Str(60),
				ProgramType:                        testconstants.ProgramType,
				ProgramTypeVersion:                 testconstants.ProgramTypeVersion,
				CompliantPlatformUsed:              testconstants.CompliantPlatformUsed,
				CompliantPlatformVersion:           testconstants.CompliantPlatformVersion,
				Transport:                          testconstants.Transport,
				FamilyId:                           testconstants.FamilyID,
				SupportedClusters:                  testconstants.SupportedClusters,
				OSVersion:                          testconstants.OSVersion,
				CertificationIdOfSoftwareComponent: testconstants.CertificationIDOfSoftwareComponent,
			},
		},
		{
			name: "valid update compliance info message",
			msg: MsgUpdateComplianceInfo{
				Creator:           sample.AccAddress(),
				Vid:               1,
				Pid:               1,
				CertificationType: testconstants.CertificationType,
				SoftwareVersion:   testconstants.SoftwareVersion,
				Reason:            testconstants.Reason,
			},
		},
		{
			name: "certification type is zigbee",
			msg: MsgUpdateComplianceInfo{
				Creator:           sample.AccAddress(),
				Pid:               1,
				Vid:               1,
				Date:              testconstants.CertificationDate,
				CertificationType: "zigbee",
				SoftwareVersion:   testconstants.SoftwareVersion,
				CDVersionNumber:   "312",
				CDCertificateId:   testconstants.CDCertificateID,
			},
		},
		{
			name: "certification type is matter",
			msg: MsgUpdateComplianceInfo{
				Creator:           sample.AccAddress(),
				Pid:               1,
				Vid:               1,
				Date:              testconstants.CertificationDate,
				CertificationType: "matter",
				SoftwareVersion:   testconstants.SoftwareVersion,
				CDVersionNumber:   "312",
				CDCertificateId:   testconstants.CDCertificateID,
			},
		},
		{
			name: "cdVersionNumber > 65535",
			msg: MsgUpdateComplianceInfo{
				Creator:           sample.AccAddress(),
				Vid:               1,
				Pid:               1,
				CertificationType: testconstants.CertificationType,
				SoftwareVersion:   testconstants.SoftwareVersion,
				Reason:            testconstants.Reason,
				CDVersionNumber:   "65536",
			},
			err: validator.ErrFieldUpperBoundViolated,
		},

		{
			name: "cdVersionNumber < 0",
			msg: MsgUpdateComplianceInfo{
				Creator:           sample.AccAddress(),
				Vid:               1,
				Pid:               1,
				CertificationType: testconstants.CertificationType,
				SoftwareVersion:   testconstants.SoftwareVersion,
				Reason:            testconstants.Reason,
				CDVersionNumber:   "-1",
			},
			err: strconv.ErrSyntax,
		},
		{
			name: "non-integer cdVersionNumber",
			msg: MsgUpdateComplianceInfo{
				Creator:           sample.AccAddress(),
				Vid:               1,
				Pid:               1,
				CertificationType: testconstants.CertificationType,
				SoftwareVersion:   testconstants.SoftwareVersion,
				Reason:            testconstants.Reason,
				CDVersionNumber:   "0.1402",
			},
			err: strconv.ErrSyntax,
		},
	}
	for _, tt := range tests {
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
