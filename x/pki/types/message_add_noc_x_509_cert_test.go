package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	tmrand "github.com/tendermint/tendermint/libs/rand"

	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	"github.com/zigbee-alliance/distributed-compliance-ledger/testutil/sample"
	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/validator"
)

func TestMsgAddNocX509Cert_ValidateBasic(t *testing.T) {
	negativeTests := []struct {
		name string
		msg  MsgAddNocX509Cert
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgAddNocX509Cert{
				Signer: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "empty certificate",
			msg: MsgAddNocX509Cert{
				Signer: sample.AccAddress(),
				Cert:   "",
			},
			err: validator.ErrRequiredFieldMissing,
		},
		{
			name: "invalid certificate",
			msg: MsgAddNocX509Cert{
				Signer: sample.AccAddress(),
				Cert:   testconstants.StubCertPem,
			},
			err: pkitypes.ErrInvalidCertificate,
		},
		{
			name: "cert len > 10485760",
			msg: MsgAddNocX509Cert{
				Signer: sample.AccAddress(),
				Cert:   tmrand.Str(10485761),
			},
			err: validator.ErrFieldMaxLengthExceeded,
		},
	}
	positiveTests := []struct {
		name string
		msg  MsgAddNocX509Cert
	}{
		{
			name: "valid add NOC cert msg",
			msg: MsgAddNocX509Cert{
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
