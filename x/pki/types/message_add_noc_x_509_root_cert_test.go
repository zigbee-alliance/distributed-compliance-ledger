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

func TestMsgAddNocX509RootCert_ValidateBasic(t *testing.T) {
	negativeTests := []struct {
		name string
		msg  MsgAddNocX509RootCert
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgAddNocX509RootCert{
				Signer: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "empty certificate",
			msg: MsgAddNocX509RootCert{
				Signer: sample.AccAddress(),
				Cert:   "",
			},
			err: validator.ErrRequiredFieldMissing,
		},
		{
			name: "cert size > 20480 (20KB)",
			msg: MsgAddNocX509RootCert{
				Signer: sample.AccAddress(),
				Cert:   tmrand.Str(20490),
			},
			err: validator.ErrFieldMaxLengthExceeded,
		},
		{
			name: "certSchemaVersion != 0",
			msg: MsgAddNocX509RootCert{
				Signer:            sample.AccAddress(),
				Cert:              testconstants.NocRootCert1,
				CertSchemaVersion: 5,
			},
			err: validator.ErrFieldEqualBoundViolated,
		},
	}

	positiveTests := []struct {
		name string
		msg  MsgAddNocX509RootCert
	}{
		{
			name: "valid add NOC root cert msg",
			msg: MsgAddNocX509RootCert{
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
