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

func TestMsgApproveRevokeX509RootCert_ValidateBasic(t *testing.T) {
	negative_tests := []struct {
		name string
		msg  MsgApproveRevokeX509RootCert
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgApproveRevokeX509RootCert{
				Signer:       "invalid_address",
				Subject:      testconstants.RootSubject,
				SubjectKeyId: testconstants.RootSubjectKeyID,
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "empty subject",
			msg: MsgApproveRevokeX509RootCert{
				Signer:       sample.AccAddress(),
				Subject:      "",
				SubjectKeyId: testconstants.RootSubjectKeyID,
			},
			err: validator.ErrRequiredFieldMissing,
		},
		{
			name: "empty SubjectKeyId",
			msg: MsgApproveRevokeX509RootCert{
				Signer:       sample.AccAddress(),
				Subject:      testconstants.RootSubject,
				SubjectKeyId: "",
			},
			err: validator.ErrRequiredFieldMissing,
		},
		{
			name: "subject len > 50",
			msg: MsgApproveRevokeX509RootCert{
				Signer:       sample.AccAddress(),
				Subject:      testconstants.RootSubject + tmrand.Str(51-len(testconstants.RootSubject)),
				SubjectKeyId: testconstants.RootSubjectKeyID,
			},
			err: validator.ErrFieldMaxLengthExceeded,
		},
		{
			name: "subject key id len > 150",
			msg: MsgApproveRevokeX509RootCert{
				Signer:       sample.AccAddress(),
				Subject:      testconstants.RootSubject,
				SubjectKeyId: testconstants.RootSubjectKeyID + tmrand.Str(151-len(testconstants.RootSubjectKeyID)),
			},
			err: validator.ErrFieldMaxLengthExceeded,
		},
	}

	positive_tests := []struct {
		name string
		msg  MsgApproveRevokeX509RootCert
	}{
		{
			name: "valid approve revoke x509cert msg",
			msg: MsgApproveRevokeX509RootCert{
				Signer:       sample.AccAddress(),
				Subject:      testconstants.RootSubject,
				SubjectKeyId: testconstants.RootSubjectKeyID,
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
