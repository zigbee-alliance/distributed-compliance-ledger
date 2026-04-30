package types

import (
	fmt "fmt"
	"testing"

	tmrand "github.com/cometbft/cometbft/libs/rand"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	"github.com/zigbee-alliance/distributed-compliance-ledger/testutil/sample"
	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/validator"
)

func TestMsgUpdatePkiRevocationDistributionPoint_ValidateBasic(t *testing.T) {
	negativeTests := []struct {
		name string
		msg  MsgUpdatePkiRevocationDistributionPoint
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdatePkiRevocationDistributionPoint{
				Signer:        "invalid_address",
				SchemaVersion: 0,
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "empty vid",
			msg: MsgUpdatePkiRevocationDistributionPoint{
				Signer:        sample.AccAddress(),
				SchemaVersion: 0,
			},
			err: validator.ErrFieldLowerBoundViolated,
		},
		{
			name: "vid < 1",
			msg: MsgUpdatePkiRevocationDistributionPoint{
				Signer:        sample.AccAddress(),
				Vid:           0,
				SchemaVersion: 0,
			},
			err: validator.ErrFieldLowerBoundViolated,
		},
		{
			name: "vid > 65535",
			msg: MsgUpdatePkiRevocationDistributionPoint{
				Signer:        sample.AccAddress(),
				Vid:           65536,
				SchemaVersion: 0,
			},
			err: validator.ErrFieldUpperBoundViolated,
		},
		{
			name: "label empty",
			msg: MsgUpdatePkiRevocationDistributionPoint{
				Signer:        sample.AccAddress(),
				Vid:           testconstants.Vid,
				SchemaVersion: 0,
			},
			err: validator.ErrRequiredFieldMissing,
		},
		{
			name: "label > 64",
			msg: MsgUpdatePkiRevocationDistributionPoint{
				Signer:        sample.AccAddress(),
				Vid:           testconstants.Vid,
				SchemaVersion: 0,
				Label:         tmrand.Str(65),
			},
			err: validator.ErrFieldMaxLengthExceeded,
		},
		{
			name: "issuerSubjectKeyID empty",
			msg: MsgUpdatePkiRevocationDistributionPoint{
				Signer:               sample.AccAddress(),
				Vid:                  testconstants.Vid,
				CrlSignerCertificate: testconstants.RootCertPem,
				Label:                "label",
				SchemaVersion:        0,
			},
			err: validator.ErrRequiredFieldMissing,
		},
		{
			name: fmt.Sprintf("dataDigestType is not one of %v", allowedDataDigestTypes),
			msg: MsgUpdatePkiRevocationDistributionPoint{
				Signer:               sample.AccAddress(),
				Vid:                  testconstants.Vid,
				CrlSignerCertificate: testconstants.RootCertPem,
				Label:                "label",
				DataURL:              testconstants.DataURL,
				IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
				DataDigestType:       3,
				SchemaVersion:        0,
			},
			err: pkitypes.ErrInvalidDataDigestType,
		},
		{
			name: "dataURL starts not with http or https",
			msg: MsgUpdatePkiRevocationDistributionPoint{
				Signer:               sample.AccAddress(),
				Vid:                  testconstants.PAACertWithNumericVidVid,
				CrlSignerCertificate: testconstants.PAACertWithNumericVid,
				Label:                "label",
				DataURL:              "ftp://" + testconstants.URLWithoutProtocol,
				IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
				SchemaVersion:        0,
			},
			err: validator.ErrFieldNotValid,
		},
		{
			name: "dataURL without protocol",
			msg: MsgUpdatePkiRevocationDistributionPoint{
				Signer:               sample.AccAddress(),
				Vid:                  testconstants.PAACertWithNumericVidVid,
				CrlSignerCertificate: testconstants.PAACertWithNumericVid,
				Label:                "label",
				DataURL:              testconstants.URLWithoutProtocol,
				IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
				SchemaVersion:        0,
			},
			err: validator.ErrFieldNotValid,
		},
		{
			name: "dataURL length > 256",
			msg: MsgUpdatePkiRevocationDistributionPoint{
				Signer:               sample.AccAddress(),
				Vid:                  testconstants.PAACertWithNumericVidVid,
				CrlSignerCertificate: testconstants.PAACertWithNumericVid,
				Label:                "label",
				DataURL: func() string {
					longUrl := testconstants.DataURL
					for i := 0; i < 29; i++ {
						longUrl += "/longurl"
					}

					return longUrl
				}(),
				IssuerSubjectKeyID: testconstants.SubjectKeyIDWithoutColons,
				SchemaVersion:      0,
			},
			err: validator.ErrFieldMaxLengthExceeded,
		},
		{
			name: "dataDigest presented, DataFileSize not presented",
			msg: MsgUpdatePkiRevocationDistributionPoint{
				Signer:               sample.AccAddress(),
				Vid:                  testconstants.Vid,
				CrlSignerCertificate: testconstants.RootCertPem,
				Label:                "label",
				DataURL:              testconstants.DataURL,
				IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
				DataDigest:           testconstants.DataDigest,
				SchemaVersion:        0,
			},
			err: pkitypes.ErrEmptyDataFileSize,
		},
		{
			name: "dataDigestType presented, DataDigest not presented",
			msg: MsgUpdatePkiRevocationDistributionPoint{
				Signer:               sample.AccAddress(),
				Vid:                  testconstants.Vid,
				CrlSignerCertificate: testconstants.RootCertPem,
				Label:                "label",
				DataURL:              testconstants.DataURL,
				IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
				DataDigestType:       1,
				SchemaVersion:        0,
			},
			err: pkitypes.ErrNotEmptyDataDigestType,
		},
		{
			name: "dataDigest presented, DataDigestType not presented",
			msg: MsgUpdatePkiRevocationDistributionPoint{
				Signer:               sample.AccAddress(),
				Vid:                  testconstants.Vid,
				CrlSignerCertificate: testconstants.RootCertPem,
				Label:                "label",
				DataURL:              testconstants.DataURL,
				IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
				DataDigest:           testconstants.DataDigest,
				DataFileSize:         123,
				SchemaVersion:        0,
			},
			err: pkitypes.ErrEmptyDataDigestType,
		},
		{
			name: "dataDigest length > 128",
			msg: MsgUpdatePkiRevocationDistributionPoint{
				Signer:               sample.AccAddress(),
				Vid:                  testconstants.Vid,
				CrlSignerCertificate: testconstants.RootCertPem,
				Label:                "label",
				DataURL:              testconstants.DataURL,
				IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
				DataDigestType:       1,
				DataDigest: func() string {
					longDataDigest := testconstants.DataDigest
					for i := 0; i < 5; i++ {
						longDataDigest += testconstants.DataDigest
					}

					return longDataDigest
				}(), DataFileSize: 123,
				SchemaVersion: 0,
			},
			err: validator.ErrFieldMaxLengthExceeded,
		},
		{
			name: "wrong IssuerSubjectKeyID format (not [0-9A-F])",
			msg: MsgUpdatePkiRevocationDistributionPoint{
				Signer:               sample.AccAddress(),
				Vid:                  testconstants.Vid,
				CrlSignerCertificate: testconstants.RootCertPem,
				Label:                "label",
				DataURL:              testconstants.DataURL,
				IssuerSubjectKeyID:   "QWERTY",
				SchemaVersion:        0,
			},
			err: pkitypes.ErrWrongSubjectKeyIDFormat,
		},
		{
			name: "wrong IssuerSubjectKeyID format (not even number of symbols)",
			msg: MsgUpdatePkiRevocationDistributionPoint{
				Signer:               sample.AccAddress(),
				Vid:                  testconstants.Vid,
				CrlSignerCertificate: testconstants.RootCertPem,
				Label:                "label",
				DataURL:              testconstants.DataURL,
				IssuerSubjectKeyID:   "123",
				SchemaVersion:        0,
			},
			err: pkitypes.ErrWrongSubjectKeyIDFormat,
		},
		{
			name: "wrong IssuerSubjectKeyID format (not even number of symbols)",
			msg: MsgUpdatePkiRevocationDistributionPoint{
				Signer:               sample.AccAddress(),
				Vid:                  testconstants.Vid,
				CrlSignerCertificate: testconstants.RootCertPem,
				Label:                "label",
				DataURL:              testconstants.DataURL,
				IssuerSubjectKeyID:   "123",
				SchemaVersion:        0,
			},
			err: pkitypes.ErrWrongSubjectKeyIDFormat,
		},
		{
			name: "when CrlSignerCertificate size exceeds 2KB",
			msg: MsgUpdatePkiRevocationDistributionPoint{
				Signer:               sample.AccAddress(),
				Vid:                  testconstants.Vid,
				CrlSignerCertificate: testconstants.CertWithSizeGreater2KB,
				Label:                "label",
				DataURL:              testconstants.DataURL,
				IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
				SchemaVersion:        0,
			},
			err: validator.ErrFieldMaxLengthExceeded,
		},
		{
			name: "when CrlSignerDelegator size exceeds 2KB",
			msg: MsgUpdatePkiRevocationDistributionPoint{
				Signer:               sample.AccAddress(),
				Vid:                  testconstants.PAACertWithNumericVidVid,
				CrlSignerCertificate: testconstants.PAACertWithNumericVid,
				Label:                "label",
				DataURL:              testconstants.DataURL,
				IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
				CrlSignerDelegator:   testconstants.CertWithSizeGreater2KB,
				SchemaVersion:        0,
			},
			err: validator.ErrFieldMaxLengthExceeded,
		},
		{
			name: "issuerSubjectKeyID > 64",
			msg: MsgUpdatePkiRevocationDistributionPoint{
				Signer:               sample.AccAddress(),
				Vid:                  testconstants.PAACertWithNumericVidVid,
				CrlSignerCertificate: testconstants.PAACertWithNumericVid,
				Label:                "label",
				DataURL:              testconstants.DataURL,
				IssuerSubjectKeyID:   tmrand.Str(65),
				SchemaVersion:        0,
			},
			err: validator.ErrFieldMaxLengthExceeded,
		},
		{
			name: "schemaVersion != 0",
			msg: MsgUpdatePkiRevocationDistributionPoint{
				Signer:               sample.AccAddress(),
				Label:                "label",
				Vid:                  1,
				IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
				CrlSignerCertificate: testconstants.PAACertWithNumericVid,
				DataURL:              testconstants.DataURL,
				DataDigest:           testconstants.DataDigest,
				DataDigestType:       1,
				DataFileSize:         123,
				SchemaVersion:        5,
			},
			err: validator.ErrFieldEqualBoundViolated,
		},
	}

	positiveTests := []struct {
		name string
		msg  MsgUpdatePkiRevocationDistributionPoint
	}{
		{
			name: "minimal msg",
			msg: MsgUpdatePkiRevocationDistributionPoint{
				Signer:             sample.AccAddress(),
				Vid:                1,
				Label:              "label",
				IssuerSubjectKeyID: testconstants.SubjectKeyIDWithoutColons,
				SchemaVersion:      0,
			},
		},
		{
			name: "maximum msg",
			msg: MsgUpdatePkiRevocationDistributionPoint{
				Signer:               sample.AccAddress(),
				Label:                "label",
				Vid:                  1,
				IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
				CrlSignerCertificate: testconstants.PAACertWithNumericVid,
				DataURL:              testconstants.DataURL,
				DataDigest:           testconstants.DataDigest,
				DataDigestType:       1,
				DataFileSize:         123,
				SchemaVersion:        0,
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
