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

func TestMsgProposeAddX509RootCert_ValidateBasic(t *testing.T) {
	negativeTests := []struct {
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
				Vid:    testconstants.Vid,
			},
			err: validator.ErrFieldMaxLengthExceeded,
		},
		{
			name: "info len > 4096",
			msg: MsgProposeAddX509RootCert{
				Signer: sample.AccAddress(),
				Cert:   testconstants.RootCertPem,
				Info:   tmrand.Str(4097),
				Vid:    testconstants.Vid,
			},
			err: validator.ErrFieldMaxLengthExceeded,
		},

		{
			name: "VID is required",
			msg: MsgProposeAddX509RootCert{
				Signer: sample.AccAddress(),
				Cert:   testconstants.RootCertPem,
				Info:   testconstants.Info,
				Time:   12345,
			},
			err: validator.ErrFieldLowerBoundViolated,
		},
		{
			name: "invalid VID",
			msg: MsgProposeAddX509RootCert{
				Signer: sample.AccAddress(),
				Cert:   testconstants.PAACertWithNumericVid,
				Info:   testconstants.Info,
				Time:   12345,
				Vid:    testconstants.Vid + 5,
			},
			err: pkitypes.ErrCertificateVidNotEqualMsgVid,
		},
		{
			name: "schemaVersion > 65535",
			msg: MsgProposeAddX509RootCert{
				Signer:            sample.AccAddress(),
				Cert:              testconstants.RootCertPem,
				Info:              testconstants.Info,
				Time:              12345,
				Vid:               testconstants.Vid,
				CertSchemaVersion: testconstants.CertSchemaVersion,
				SchemaVersion:     65536,
			},
			err: validator.ErrFieldUpperBoundViolated,
		},
		{
			name: "certSchemaVersion > 65535",
			msg: MsgProposeAddX509RootCert{
				Signer:            sample.AccAddress(),
				Cert:              testconstants.RootCertPem,
				Info:              testconstants.Info,
				Time:              12345,
				Vid:               testconstants.Vid,
				CertSchemaVersion: 65536,
				SchemaVersion:     testconstants.SchemaVersion,
			},
			err: validator.ErrFieldUpperBoundViolated,
		},
	}

	positiveTests := []struct {
		name string
		msg  MsgProposeAddX509RootCert
	}{
		{
			name: "valid propose add x509cert msg",
			msg: MsgProposeAddX509RootCert{
				Signer: sample.AccAddress(),
				Cert:   testconstants.PAACertWithNumericVid,
				Info:   testconstants.Info,
				Time:   12345,
				Vid:    testconstants.GoogleVid,
			},
		},
		{
			name: "info field length = 4096",
			msg: MsgProposeAddX509RootCert{
				Signer: sample.AccAddress(),
				Cert:   testconstants.RootCertPem,
				Info:   tmrand.Str(4096),
				Vid:    testconstants.Vid,
			},
		},
		{
			name: "info field length is empty",
			msg: MsgProposeAddX509RootCert{
				Signer: sample.AccAddress(),
				Cert:   testconstants.RootCertPem,
				Info:   "",
				Vid:    testconstants.Vid,
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
