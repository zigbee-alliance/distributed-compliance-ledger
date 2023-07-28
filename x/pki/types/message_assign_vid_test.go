package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	tmrand "github.com/tendermint/tendermint/libs/rand"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	"github.com/zigbee-alliance/distributed-compliance-ledger/testutil/sample"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/validator"
)

func TestMsgAssignVid_ValidateBasic(t *testing.T) {
	negativeTests := []struct {
		name string
		msg  MsgAssignVid
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgAssignVid{
				Signer: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "empty subject",
			msg: MsgAssignVid{
				Signer:       sample.AccAddress(),
				Subject:      "",
				SubjectKeyId: testconstants.PAACertWithNumericVidSubjectKeyID,
				Vid:          testconstants.PAACertWithNumericVidVid,
			},
			err: validator.ErrRequiredFieldMissing,
		},
		{
			name: "empty SubjectKeyId",
			msg: MsgAssignVid{
				Signer:       sample.AccAddress(),
				Subject:      testconstants.PAACertWithNumericVidSubject,
				SubjectKeyId: "",
				Vid:          testconstants.PAACertWithNumericVidVid,
			},
			err: validator.ErrRequiredFieldMissing,
		},
		{
			name: "subject len > 1024 (1 KB)",
			msg: MsgAssignVid{
				Signer:       sample.AccAddress(),
				Subject:      testconstants.PAACertWithNumericVidSubject + tmrand.Str(1025-len(testconstants.PAACertWithNumericVidSubjectKeyID)),
				SubjectKeyId: testconstants.PAACertWithNumericVidSubjectKeyID,
				Vid:          testconstants.PAACertWithNumericVidVid,
			},
			err: validator.ErrFieldMaxLengthExceeded,
		},
		{
			name: "subject key id len > 256",
			msg: MsgAssignVid{
				Signer:       sample.AccAddress(),
				Subject:      testconstants.PAACertWithNumericVidSubject,
				SubjectKeyId: testconstants.PAACertWithNumericVidSubjectKeyID + tmrand.Str(257-len(testconstants.PAACertWithNumericVidSubjectKeyID)),
				Vid:          testconstants.PAACertWithNumericVidVid,
			},
			err: validator.ErrFieldMaxLengthExceeded,
		},
		{
			name: "empty vid",
			msg: MsgAssignVid{
				Signer:       sample.AccAddress(),
				Subject:      testconstants.PAACertWithNumericVidSubject,
				SubjectKeyId: testconstants.PAACertWithNumericVidSubjectKeyID,
			},
			err: validator.ErrFieldLowerBoundViolated,
		},
		{
			name: "vid < 1",
			msg: MsgAssignVid{
				Signer:       sample.AccAddress(),
				Subject:      testconstants.PAACertWithNumericVidSubject,
				SubjectKeyId: testconstants.PAACertWithNumericVidSubjectKeyID,
				Vid:          0,
			},
			err: validator.ErrFieldLowerBoundViolated,
		},
		{
			name: "vid > 65535",
			msg: MsgAssignVid{
				Signer:       sample.AccAddress(),
				Subject:      testconstants.PAACertWithNumericVidSubject,
				SubjectKeyId: testconstants.PAACertWithNumericVidSubjectKeyID,
				Vid:          65536,
			},
			err: validator.ErrFieldUpperBoundViolated,
		},
	}

	positiveTests := []struct {
		name string
		msg  MsgAssignVid
	}{
		{
			name: "valid assign vid msg",
			msg: MsgAssignVid{
				Signer:       sample.AccAddress(),
				Subject:      testconstants.PAACertWithNumericVidSubject,
				SubjectKeyId: testconstants.PAACertWithNumericVidSubjectKeyID,
				Vid:          testconstants.PAACertWithNumericVidVid,
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
