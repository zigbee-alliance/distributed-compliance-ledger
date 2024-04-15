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

func TestMsgRevokeNocX509RootCert_ValidateBasic(t *testing.T) {
	negativeTests := []struct {
		name string
		msg  MsgRevokeNocX509RootCert
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgRevokeNocX509RootCert{
				Signer: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "empty subject",
			msg: MsgRevokeNocX509RootCert{
				Signer:       sample.AccAddress(),
				Subject:      "",
				SubjectKeyId: testconstants.RootSubjectKeyID,
			},
			err: validator.ErrRequiredFieldMissing,
		},
		{
			name: "empty SubjectKeyId",
			msg: MsgRevokeNocX509RootCert{
				Signer:       sample.AccAddress(),
				Subject:      testconstants.RootSubject,
				SubjectKeyId: "",
			},
			err: validator.ErrRequiredFieldMissing,
		},
		{
			name: "subject len > 1024 (1 KB)",
			msg: MsgRevokeNocX509RootCert{
				Signer:       sample.AccAddress(),
				Subject:      testconstants.RootSubject + tmrand.Str(1025-len(testconstants.RootSubject)),
				SubjectKeyId: testconstants.RootSubjectKeyID,
			},
			err: validator.ErrFieldMaxLengthExceeded,
		},
		{
			name: "subject key id len > 256",
			msg: MsgRevokeNocX509RootCert{
				Signer:       sample.AccAddress(),
				Subject:      testconstants.RootSubject,
				SubjectKeyId: testconstants.RootSubjectKeyID + tmrand.Str(257-len(testconstants.RootSubjectKeyID)),
			},
			err: validator.ErrFieldMaxLengthExceeded,
		},
		{
			name: "info len > 4096",
			msg: MsgRevokeNocX509RootCert{
				Signer:       sample.AccAddress(),
				Subject:      testconstants.RootSubject,
				SubjectKeyId: testconstants.RootSubjectKeyID,
				Info:         tmrand.Str(4097),
			},
			err: validator.ErrFieldMaxLengthExceeded,
		},
		{
			name: "schemaVersion > 65535",
			msg: MsgRevokeNocX509RootCert{
				Signer:        sample.AccAddress(),
				Subject:       testconstants.RootSubject,
				SubjectKeyId:  testconstants.RootSubjectKeyID,
				SerialNumber:  testconstants.RootSerialNumber,
				Info:          testconstants.Info,
				Time:          12345,
				SchemaVersion: 65536,
			},
			err: validator.ErrFieldUpperBoundViolated,
		},
	}
	positiveTests := []struct {
		name string
		msg  MsgRevokeNocX509RootCert
	}{
		{
			name: "valid revoke x509cert msg",
			msg: MsgRevokeNocX509RootCert{
				Signer:       sample.AccAddress(),
				Subject:      testconstants.RootSubject,
				SubjectKeyId: testconstants.RootSubjectKeyID,
				SerialNumber: testconstants.RootSerialNumber,
				Info:         testconstants.Info,
				Time:         12345,
			},
		},
		{
			name: "valid revoke x509cert msg with revokeChild true flag",
			msg: MsgRevokeNocX509RootCert{
				Signer:       sample.AccAddress(),
				Subject:      testconstants.RootSubject,
				SubjectKeyId: testconstants.RootSubjectKeyID,
				SerialNumber: testconstants.RootSerialNumber,
				Info:         testconstants.Info,
				Time:         12345,
				RevokeChild:  true,
			},
		},
		{
			name: "info field is 4096 characters long",
			msg: MsgRevokeNocX509RootCert{
				Signer:       sample.AccAddress(),
				Subject:      testconstants.RootSubject,
				SubjectKeyId: testconstants.RootSubjectKeyID,
				Info:         tmrand.Str(4096),
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
