package types

import (
	"testing"

	tmrand "github.com/cometbft/cometbft/libs/rand"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	"github.com/zigbee-alliance/distributed-compliance-ledger/testutil/sample"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/validator"
)

func TestMsgRemoveNocX509IcaCert_ValidateBasic(t *testing.T) {
	negativeTests := []struct {
		name string
		msg  MsgRemoveNocX509IcaCert
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgRemoveNocX509IcaCert{
				Signer:       "invalid_address",
				Subject:      testconstants.RootSubject,
				SubjectKeyId: testconstants.RootSubjectKeyID,
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "empty subject",
			msg: MsgRemoveNocX509IcaCert{
				Signer:       sample.AccAddress(),
				Subject:      "",
				SubjectKeyId: testconstants.RootSubjectKeyID,
			},
			err: validator.ErrRequiredFieldMissing,
		},
		{
			name: "empty SubjectKeyId",
			msg: MsgRemoveNocX509IcaCert{
				Signer:       sample.AccAddress(),
				Subject:      testconstants.RootSubject,
				SubjectKeyId: "",
			},
			err: validator.ErrRequiredFieldMissing,
		},
		{
			name: "subject len > 1024 (1 KB)",
			msg: MsgRemoveNocX509IcaCert{
				Signer:       sample.AccAddress(),
				Subject:      testconstants.RootSubject + tmrand.Str(1025-len(testconstants.RootSubject)),
				SubjectKeyId: testconstants.RootSubjectKeyID,
			},
			err: validator.ErrFieldMaxLengthExceeded,
		},
		{
			name: "subject key id len > 256",
			msg: MsgRemoveNocX509IcaCert{
				Signer:       sample.AccAddress(),
				Subject:      testconstants.RootSubject,
				SubjectKeyId: testconstants.RootSubjectKeyID + tmrand.Str(257-len(testconstants.RootSubjectKeyID)),
			},
			err: validator.ErrFieldMaxLengthExceeded,
		},
	}

	positiveTests := []struct {
		name string
		msg  MsgRemoveNocX509IcaCert
	}{
		{
			name: "valid remove nocX509IcaCert msg",
			msg: MsgRemoveNocX509IcaCert{
				Signer:       sample.AccAddress(),
				Subject:      testconstants.RootSubject,
				SubjectKeyId: testconstants.RootSubjectKeyID,
			},
		},
		{
			name: "valid remove nocX509IcaCert msg with serial number",
			msg: MsgRemoveNocX509IcaCert{
				Signer:       sample.AccAddress(),
				Subject:      testconstants.RootSubject,
				SubjectKeyId: testconstants.RootSubjectKeyID,
				SerialNumber: testconstants.RootSerialNumber,
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
