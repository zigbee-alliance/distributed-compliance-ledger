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

func TestMsgRevokeNocX509Cert_ValidateBasic(t *testing.T) {
	negativeTests := []struct {
		name string
		msg  MsgRevokeNocX509Cert
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgRevokeNocX509Cert{
				Signer: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "empty subject",
			msg: MsgRevokeNocX509Cert{
				Signer:       sample.AccAddress(),
				Subject:      "",
				SubjectKeyId: testconstants.NocCert1SubjectKeyID,
			},
			err: validator.ErrRequiredFieldMissing,
		},
		{
			name: "empty SubjectKeyId",
			msg: MsgRevokeNocX509Cert{
				Signer:       sample.AccAddress(),
				Subject:      testconstants.NocCert1Subject,
				SubjectKeyId: "",
			},
			err: validator.ErrRequiredFieldMissing,
		},
		{
			name: "subject len > 1024 (1 KB)",
			msg: MsgRevokeNocX509Cert{
				Signer:       sample.AccAddress(),
				Subject:      testconstants.NocCert1Subject + tmrand.Str(1025-len(testconstants.NocCert1Subject)),
				SubjectKeyId: testconstants.NocCert1SubjectKeyID,
			},
			err: validator.ErrFieldMaxLengthExceeded,
		},
		{
			name: "subject key id len > 256",
			msg: MsgRevokeNocX509Cert{
				Signer:       sample.AccAddress(),
				Subject:      testconstants.NocCert1Subject,
				SubjectKeyId: testconstants.NocCert1SubjectKeyID + tmrand.Str(257-len(testconstants.NocCert1SubjectKeyID)),
			},
			err: validator.ErrFieldMaxLengthExceeded,
		},
		{
			name: "info len > 4096",
			msg: MsgRevokeNocX509Cert{
				Signer:       sample.AccAddress(),
				Subject:      testconstants.NocCert1Subject,
				SubjectKeyId: testconstants.NocCert1SubjectKeyID,
				Info:         tmrand.Str(4097),
			},
			err: validator.ErrFieldMaxLengthExceeded,
		},
	}
	positiveTests := []struct {
		name string
		msg  MsgRevokeNocX509Cert
	}{
		{
			name: "valid revoke x509cert msg",
			msg: MsgRevokeNocX509Cert{
				Signer:       sample.AccAddress(),
				Subject:      testconstants.NocCert1Subject,
				SubjectKeyId: testconstants.NocCert1SubjectKeyID,
				SerialNumber: testconstants.NocCert1SerialNumber,
				Info:         testconstants.Info,
				Time:         12345,
			},
		},
		{
			name: "valid revoke x509cert msg with revokeChild true flag",
			msg: MsgRevokeNocX509Cert{
				Signer:       sample.AccAddress(),
				Subject:      testconstants.NocCert1Subject,
				SubjectKeyId: testconstants.NocCert1SubjectKeyID,
				SerialNumber: testconstants.NocCert1SerialNumber,
				Info:         testconstants.Info,
				Time:         12345,
				RevokeChild:  true,
			},
		},
		{
			name: "info field is 4096 characters long",
			msg: MsgRevokeNocX509Cert{
				Signer:       sample.AccAddress(),
				Subject:      testconstants.NocCert1Subject,
				SubjectKeyId: testconstants.NocCert1SubjectKeyID,
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
