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

func TestMsgProposeAddX509RootCert_ValidateBasic(t *testing.T) {
	negative_tests := []struct {
		name string
		msg  MsgProposeAddX509RootCert
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgProposeAddX509RootCert{
				Signer: "invalid_address",
				Cert:   testconstants.RootCertPem,
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "empty certificate",
			msg: MsgProposeAddX509RootCert{
				Signer: sample.AccAddress(),
				Cert:   "",
			},
			err: validator.ErrRequiredFieldMissing,
		},
		{
			name: "cert len > 10485760 (10 MB)",
			msg: MsgProposeAddX509RootCert{
				Signer: sample.AccAddress(),
				Cert:   testconstants.RootCertPem + tmrand.Str(10485761-len(testconstants.RootCertPem)),
			},
			err: validator.ErrFieldMaxLengthExceeded,
		},
	}

	positive_tests := []struct {
		name string
		msg  MsgProposeAddX509RootCert
	}{
		{
			name: "valid propose add x509cert msg",
			msg: MsgProposeAddX509RootCert{
				Signer: sample.AccAddress(),
				Cert:   testconstants.RootCertPem,
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
