package types

import (
	"testing"

	tmrand "github.com/cometbft/cometbft/libs/rand"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	"github.com/zigbee-alliance/distributed-compliance-ledger/testutil/sample"
	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/validator"
)

func TestMsgAddNocX509IcaCert_ValidateBasic(t *testing.T) {
	negativeTests := []struct {
		name string
		msg  MsgAddNocX509IcaCert
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgAddNocX509IcaCert{
				Signer: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "empty certificate",
			msg: MsgAddNocX509IcaCert{
				Signer: sample.AccAddress(),
				Cert:   "",
			},
			err: validator.ErrRequiredFieldMissing,
		},
		{
			name: "invalid certificate",
			msg: MsgAddNocX509IcaCert{
				Signer: sample.AccAddress(),
				Cert:   testconstants.StubCertPem,
			},
			err: pkitypes.ErrInvalidCertificate,
		},
		{
			name: "cert len > 10485760",
			msg: MsgAddNocX509IcaCert{
				Signer: sample.AccAddress(),
				Cert:   tmrand.Str(10485761),
			},
			err: validator.ErrFieldMaxLengthExceeded,
		},
		{
			name: "certSchemaVersion != 0",
			msg: MsgAddNocX509IcaCert{
				Signer:            sample.AccAddress(),
				Cert:              testconstants.NocCert1,
				CertSchemaVersion: 5,
			},
			err: validator.ErrFieldEqualBoundViolated,
		},
	}
	positiveTests := []struct {
		name string
		msg  MsgAddNocX509IcaCert
	}{
		{
			name: "valid add NOC cert msg",
			msg: MsgAddNocX509IcaCert{
				Signer: sample.AccAddress(),
				Cert:   testconstants.NocCert1,
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
