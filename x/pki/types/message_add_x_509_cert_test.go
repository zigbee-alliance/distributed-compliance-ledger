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

func TestMsgAddX509Cert_ValidateBasic(t *testing.T) {
	negativeTests := []struct {
		name string
		msg  MsgAddX509Cert
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgAddX509Cert{
				Signer: "invalid_address",
				Cert:   testconstants.RootCertPem,
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "empty certificate",
			msg: MsgAddX509Cert{
				Signer: sample.AccAddress(),
				Cert:   "",
			},
			err: validator.ErrRequiredFieldMissing,
		},
		{
			name: "cert len > 10485760 (10 MB)",
			msg: MsgAddX509Cert{
				Signer: sample.AccAddress(),
				Cert:   testconstants.RootCertPem + tmrand.Str(10485761-len(testconstants.RootCertPem)),
			},
			err: validator.ErrFieldMaxLengthExceeded,
		},
	}

	positiveTests := []struct {
		name string
		msg  MsgAddX509Cert
	}{
		{
			name: "valid propose add x509cert msg",
			msg: MsgAddX509Cert{
				Signer: sample.AccAddress(),
				Cert:   testconstants.RootCertPem,
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
