package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	"github.com/zigbee-alliance/distributed-compliance-ledger/testutil/sample"
	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/validator"
)

func TestMsgDeletePkiRevocationDistributionPoint_ValidateBasic(t *testing.T) {
	negativeTests := []struct {
		name string
		msg  MsgDeletePkiRevocationDistributionPoint
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeletePkiRevocationDistributionPoint{
				Signer: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "empty vid",
			msg: MsgDeletePkiRevocationDistributionPoint{
				Signer: sample.AccAddress(),
			},
			err: validator.ErrFieldLowerBoundViolated,
		},
		{
			name: "vid < 1",
			msg: MsgDeletePkiRevocationDistributionPoint{
				Signer:             sample.AccAddress(),
				Vid:                0,
				Label:              "label",
				IssuerSubjectKeyID: testconstants.SubjectKeyIDWithoutColons,
			},
			err: validator.ErrFieldLowerBoundViolated,
		},
		{
			name: "vid > 65535",
			msg: MsgDeletePkiRevocationDistributionPoint{
				Signer:             sample.AccAddress(),
				Vid:                65536,
				Label:              "label",
				IssuerSubjectKeyID: testconstants.SubjectKeyIDWithoutColons,
			},
			err: validator.ErrFieldUpperBoundViolated,
		},
		{
			name: "label empty",
			msg: MsgDeletePkiRevocationDistributionPoint{
				Signer: sample.AccAddress(),
				Vid:    1,
			},
			err: validator.ErrRequiredFieldMissing,
		},
		{
			name: "issuerSubjectKeyId empty",
			msg: MsgDeletePkiRevocationDistributionPoint{
				Signer: sample.AccAddress(),
				Vid:    1,
				Label:  "label",
			},
			err: validator.ErrRequiredFieldMissing,
		},
		{
			name: "wrong IssuerSubjectKeyId format (not [0-9A-F])",
			msg: MsgDeletePkiRevocationDistributionPoint{
				Signer:             sample.AccAddress(),
				Vid:                1,
				Label:              "label",
				IssuerSubjectKeyID: "QWERTY",
			},
			err: pkitypes.ErrWrongSubjectKeyIDFormat,
		},
		{
			name: "wrong IssuerSubjectKeyId format (not even number of symbols)",
			msg: MsgDeletePkiRevocationDistributionPoint{
				Signer:             sample.AccAddress(),
				Vid:                1,
				Label:              "label",
				IssuerSubjectKeyID: "123",
			},
			err: pkitypes.ErrWrongSubjectKeyIDFormat,
		},
		{
			name: "wrong IssuerSubjectKeyId format (not even number of symbols)",
			msg: MsgDeletePkiRevocationDistributionPoint{
				Signer:             sample.AccAddress(),
				Vid:                1,
				Label:              "label",
				IssuerSubjectKeyID: "123",
			},
			err: pkitypes.ErrWrongSubjectKeyIDFormat,
		},
	}

	positiveTests := []struct {
		name string
		msg  MsgAddPkiRevocationDistributionPoint
	}{
		// {
		// 	name: "valid approve add x509cert msg",
		// },
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
