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

func TestMsgApproveAddX509RootCert_ValidateBasic(t *testing.T) {
	negative_tests := []struct {
		name string
		msg  MsgApproveAddX509RootCert
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgApproveAddX509RootCert{
				Signer:       "invalid_address",
				Subject:      testconstants.RootSubject,
				SubjectKeyId: testconstants.RootSubjectKeyID,
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "empty subject",
			msg: MsgApproveAddX509RootCert{
				Signer:       sample.AccAddress(),
				Subject:      "",
				SubjectKeyId: testconstants.RootSubjectKeyID,
			},
			err: validator.ErrRequiredFieldMissing,
		},
		{
			name: "empty SubjectKeyId",
			msg: MsgApproveAddX509RootCert{
				Signer:       sample.AccAddress(),
				Subject:      testconstants.RootSubject,
				SubjectKeyId: "",
			},
			err: validator.ErrRequiredFieldMissing,
		},
		{
			name: "subject len > 1024 (1 KB)",
			msg: MsgApproveAddX509RootCert{
				Signer:       sample.AccAddress(),
				Subject:      testconstants.RootSubject + tmrand.Str(1025-len(testconstants.RootSubject)),
				SubjectKeyId: testconstants.RootSubjectKeyID,
			},
			err: validator.ErrFieldMaxLengthExceeded,
		},
		{
			name: "subject key id len > 256",
			msg: MsgApproveAddX509RootCert{
				Signer:       sample.AccAddress(),
				Subject:      testconstants.RootSubject,
				SubjectKeyId: testconstants.RootSubjectKeyID + tmrand.Str(257-len(testconstants.RootSubjectKeyID)),
			},
			err: validator.ErrFieldMaxLengthExceeded,
		},
		{
			name: "info field length > 4096",
			msg: MsgApproveAddX509RootCert{
				Signer:       sample.AccAddress(),
				Subject:      testconstants.RootSubject,
				SubjectKeyId: testconstants.RootSubjectKeyID,
				Info:         tmrand.Str(4097),
			},
			err: validator.ErrFieldMaxLengthExceeded,
		},
	}

	positive_tests := []struct {
		name string
		msg  MsgApproveAddX509RootCert
	}{
		{
			name: "valid approve add x509cert msg",
			msg: MsgApproveAddX509RootCert{
				Signer:       sample.AccAddress(),
				Subject:      testconstants.RootSubject,
				SubjectKeyId: testconstants.RootSubjectKeyID,
			},
		},
		{
			name: "info field length = 4096",
			msg: MsgApproveAddX509RootCert{
				Signer:       sample.AccAddress(),
				Subject:      testconstants.RootSubject,
				SubjectKeyId: testconstants.RootSubjectKeyID,
				Info:         tmrand.Str(4096),
			},
		},
		{
			name: "info field is empty",
			msg: MsgApproveAddX509RootCert{
				Signer:       sample.AccAddress(),
				Subject:      testconstants.RootSubject,
				SubjectKeyId: testconstants.RootSubjectKeyID,
				Info:         "",
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
