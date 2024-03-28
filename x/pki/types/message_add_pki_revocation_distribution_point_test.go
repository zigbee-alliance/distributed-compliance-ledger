package types

import (
	fmt "fmt"
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	"github.com/zigbee-alliance/distributed-compliance-ledger/testutil/sample"
	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/validator"
)

func TestMsgAddPkiRevocationDistributionPoint_ValidateBasic(t *testing.T) {
	negativeTests := []struct {
		name string
		msg  MsgAddPkiRevocationDistributionPoint
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgAddPkiRevocationDistributionPoint{
				Signer: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "empty vid",
			msg: MsgAddPkiRevocationDistributionPoint{
				Signer: sample.AccAddress(),
			},
			err: validator.ErrFieldLowerBoundViolated,
		},
		{
			name: "vid < 1",
			msg: MsgAddPkiRevocationDistributionPoint{
				Signer: sample.AccAddress(),
				Vid:    0,
			},
			err: validator.ErrFieldLowerBoundViolated,
		},
		{
			name: "vid > 65535",
			msg: MsgAddPkiRevocationDistributionPoint{
				Signer: sample.AccAddress(),
				Vid:    65536,
			},
			err: validator.ErrFieldUpperBoundViolated,
		},
		{
			name: "pid < 0",
			msg: MsgAddPkiRevocationDistributionPoint{
				Signer: sample.AccAddress(),
				Vid:    1,
				Pid:    -1,
			},
			err: validator.ErrFieldLowerBoundViolated,
		},
		{
			name: "pid < 65535",
			msg: MsgAddPkiRevocationDistributionPoint{
				Signer: sample.AccAddress(),
				Vid:    1,
				Pid:    65536,
			},
			err: validator.ErrFieldUpperBoundViolated,
		},
		{
			name: "IsPAA empty",
			msg: MsgAddPkiRevocationDistributionPoint{
				Signer: sample.AccAddress(),
				Vid:    1,
			},
			err: validator.ErrRequiredFieldMissing,
		},
		{
			name: "label empty",
			msg: MsgAddPkiRevocationDistributionPoint{
				Signer: sample.AccAddress(),
				Vid:    1,
				IsPAA:  true,
			},
			err: validator.ErrRequiredFieldMissing,
		},
		{
			name: "crl signer certificate empty",
			msg: MsgAddPkiRevocationDistributionPoint{
				Signer: sample.AccAddress(),
				Vid:    1,
				IsPAA:  true,
				Label:  "label",
			},
			err: validator.ErrRequiredFieldMissing,
		},
		{
			name: "issuerSubjectKeyID empty",
			msg: MsgAddPkiRevocationDistributionPoint{
				Signer:               sample.AccAddress(),
				Vid:                  testconstants.PAACertWithNumericVidVid,
				IsPAA:                true,
				CrlSignerCertificate: testconstants.PAACertWithNumericVid,
				Label:                "label",
			},
			err: validator.ErrRequiredFieldMissing,
		},
		{
			name: "dataURL empty",
			msg: MsgAddPkiRevocationDistributionPoint{
				Signer:               sample.AccAddress(),
				Vid:                  testconstants.PAACertWithNumericVidVid,
				IsPAA:                true,
				CrlSignerCertificate: testconstants.PAACertWithNumericVid,
				Label:                "label",
				IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
			},
			err: validator.ErrRequiredFieldMissing,
		},
		{
			name: "revocationType empty",
			msg: MsgAddPkiRevocationDistributionPoint{
				Signer:               sample.AccAddress(),
				Vid:                  testconstants.PAACertWithNumericVidVid,
				IsPAA:                true,
				CrlSignerCertificate: testconstants.PAACertWithNumericVid,
				Label:                "label",
				DataURL:              testconstants.DataURL,
				IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
			},
			err: validator.ErrRequiredFieldMissing,
		},
		{
			name: fmt.Sprintf("dataDigestType is not one of %v", allowedDataDigestTypes),
			msg: MsgAddPkiRevocationDistributionPoint{
				Signer:               sample.AccAddress(),
				Vid:                  testconstants.PAACertWithNumericVidVid,
				IsPAA:                true,
				CrlSignerCertificate: testconstants.PAACertWithNumericVid,
				Label:                "label",
				DataURL:              testconstants.DataURL,
				IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
				RevocationType:       1,
				DataDigestType:       3,
			},
			err: pkitypes.ErrInvalidDataDigestType,
		},
		{
			name: fmt.Sprintf("revocationType is not one of %v", allowedRevocationTypes),
			msg: MsgAddPkiRevocationDistributionPoint{
				Signer:               sample.AccAddress(),
				Vid:                  testconstants.PAACertWithNumericVidVid,
				IsPAA:                true,
				CrlSignerCertificate: testconstants.PAACertWithNumericVid,
				Label:                "label",
				DataURL:              testconstants.DataURL,
				IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
				RevocationType:       2,
			},
			err: pkitypes.ErrInvalidRevocationType,
		},
		{
			name: "pid not empty when IsPAA is true",
			msg: MsgAddPkiRevocationDistributionPoint{
				Signer:               sample.AccAddress(),
				Vid:                  testconstants.PAACertWithNumericVidVid,
				IsPAA:                true,
				CrlSignerCertificate: testconstants.PAACertWithNumericVid,
				Label:                "label",
				DataURL:              testconstants.DataURL,
				IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
				RevocationType:       1,
				Pid:                  1,
			},
			err: pkitypes.ErrNotEmptyPid,
		},
		{
			name: "dataURL starts not with http or https",
			msg: MsgAddPkiRevocationDistributionPoint{
				Signer:               sample.AccAddress(),
				Vid:                  testconstants.PAACertWithNumericVidVid,
				IsPAA:                true,
				CrlSignerCertificate: testconstants.PAACertWithNumericVid,
				Label:                "label",
				DataURL:              "ftp://" + testconstants.URLWithoutProtocol,
				IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
				RevocationType:       1,
			},
			err: pkitypes.ErrInvalidDataURLFormat,
		},
		{
			name: "dataURL without protocol",
			msg: MsgAddPkiRevocationDistributionPoint{
				Signer:               sample.AccAddress(),
				Vid:                  testconstants.PAACertWithNumericVidVid,
				IsPAA:                true,
				CrlSignerCertificate: testconstants.PAACertWithNumericVid,
				Label:                "label",
				DataURL:              testconstants.URLWithoutProtocol,
				IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
				RevocationType:       1,
			},
			err: validator.ErrFieldNotValid,
		},
		{
			name: "dataDigest presented, DataFileSize not presented",
			msg: MsgAddPkiRevocationDistributionPoint{
				Signer:               sample.AccAddress(),
				Vid:                  testconstants.PAACertWithNumericVidVid,
				IsPAA:                true,
				CrlSignerCertificate: testconstants.PAACertWithNumericVid,
				Label:                "label",
				DataURL:              testconstants.DataURL,
				IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
				DataDigest:           testconstants.DataDigest,
				RevocationType:       1,
			},
			err: pkitypes.ErrEmptyDataFileSize,
		},
		{
			name: "DataFileSize presented, dataDigest not presented",
			msg: MsgAddPkiRevocationDistributionPoint{
				Signer:               sample.AccAddress(),
				Vid:                  testconstants.PAACertWithNumericVidVid,
				IsPAA:                true,
				CrlSignerCertificate: testconstants.PAACertWithNumericVid,
				Label:                "label",
				DataURL:              testconstants.DataURL,
				IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
				DataFileSize:         123,
				RevocationType:       1,
			},
			err: pkitypes.ErrEmptyDataDigest,
		},
		{
			name: "dataDigestType presented, DataDigest not presented",
			msg: MsgAddPkiRevocationDistributionPoint{
				Signer:               sample.AccAddress(),
				Vid:                  testconstants.PAACertWithNumericVidVid,
				IsPAA:                true,
				CrlSignerCertificate: testconstants.PAACertWithNumericVid,
				Label:                "label",
				DataURL:              testconstants.DataURL,
				IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
				DataDigestType:       1,
				RevocationType:       1,
			},
			err: pkitypes.ErrNotEmptyDataDigestType,
		},
		{
			name: "dataDigest presented, DataDigestType not presented",
			msg: MsgAddPkiRevocationDistributionPoint{
				Signer:               sample.AccAddress(),
				Vid:                  testconstants.PAACertWithNumericVidVid,
				IsPAA:                true,
				CrlSignerCertificate: testconstants.PAACertWithNumericVid,
				Label:                "label",
				DataURL:              testconstants.DataURL,
				IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
				DataDigest:           testconstants.DataDigest,
				DataFileSize:         123,
				RevocationType:       1,
			},
			err: pkitypes.ErrEmptyDataDigestType,
		},
		{
			name: "wrong IssuerSubjectKeyID format (not [0-9A-F])",
			msg: MsgAddPkiRevocationDistributionPoint{
				Signer:               sample.AccAddress(),
				Vid:                  testconstants.PAACertWithNumericVidVid,
				IsPAA:                true,
				CrlSignerCertificate: testconstants.PAACertWithNumericVid,
				Label:                "label",
				DataURL:              testconstants.DataURL,
				IssuerSubjectKeyID:   "QWERTY",
				RevocationType:       1,
			},
			err: pkitypes.ErrWrongSubjectKeyIDFormat,
		},
		{
			name: "wrong IssuerSubjectKeyID format (not even number of symbols)",
			msg: MsgAddPkiRevocationDistributionPoint{
				Signer:               sample.AccAddress(),
				Vid:                  testconstants.PAACertWithNumericVidVid,
				IsPAA:                true,
				CrlSignerCertificate: testconstants.PAACertWithNumericVid,
				Label:                "label",
				DataURL:              testconstants.DataURL,
				IssuerSubjectKeyID:   "123",
				RevocationType:       1,
			},
			err: pkitypes.ErrWrongSubjectKeyIDFormat,
		},
		{
			name: "wrong IssuerSubjectKeyID format (with colons)",
			msg: MsgAddPkiRevocationDistributionPoint{
				Signer:               sample.AccAddress(),
				Vid:                  testconstants.PAACertWithNumericVidVid,
				IsPAA:                true,
				CrlSignerCertificate: testconstants.PAACertWithNumericVid,
				Label:                "label",
				DataURL:              testconstants.DataURL,
				IssuerSubjectKeyID:   "12:AA:BB",
				RevocationType:       1,
			},
			err: pkitypes.ErrWrongSubjectKeyIDFormat,
		},
		{
			name: "data fields present when revocationType is 1",
			msg: MsgAddPkiRevocationDistributionPoint{
				Signer:               sample.AccAddress(),
				Vid:                  testconstants.PAACertWithNumericVidVid,
				IsPAA:                true,
				CrlSignerCertificate: testconstants.PAACertWithNumericVid,
				Label:                "label",
				DataURL:              testconstants.DataURL,
				IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
				DataFileSize:         123,
				DataDigest:           testconstants.DataDigest,
				DataDigestType:       1,
				RevocationType:       1,
			},
			err: pkitypes.ErrDataFieldPresented,
		},
		{
			name: "pid provided when PAA is true",
			msg: MsgAddPkiRevocationDistributionPoint{
				Signer:               sample.AccAddress(),
				Vid:                  testconstants.Vid,
				IsPAA:                true,
				CrlSignerCertificate: testconstants.RootCertPem,
				Label:                "label",
				DataURL:              testconstants.DataURL,
				IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
				RevocationType:       1,
				Pid:                  testconstants.Pid,
			},
			err: pkitypes.ErrNotEmptyPid,
		},
		{
			name: "pid not encoded in CRL signer certificate",
			msg: MsgAddPkiRevocationDistributionPoint{
				Signer:               sample.AccAddress(),
				Vid:                  testconstants.PAICertWithVidVid,
				IsPAA:                false,
				CrlSignerCertificate: testconstants.PAICertWithVid,
				Label:                "label",
				DataURL:              testconstants.DataURL,
				IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
				RevocationType:       1,
				Pid:                  1,
			},
			err: pkitypes.ErrNotEmptyPid,
		},
		{
			name: "IsPAA false, certificate is root",
			msg: MsgAddPkiRevocationDistributionPoint{
				Signer:               sample.AccAddress(),
				Vid:                  testconstants.Vid,
				IsPAA:                false,
				CrlSignerCertificate: testconstants.RootCertPem,
				Label:                "label",
				DataURL:              testconstants.DataURL,
				IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
				RevocationType:       1,
			},
			err: pkitypes.ErrNonRootCertificateSelfSigned,
		},
		{
			name: "IsPAA true, certificate is non-root",
			msg: MsgAddPkiRevocationDistributionPoint{
				Signer:               sample.AccAddress(),
				Vid:                  testconstants.Vid,
				IsPAA:                true,
				CrlSignerCertificate: testconstants.IntermediateCertPem,
				Label:                "label",
				DataURL:              testconstants.DataURL,
				IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
				RevocationType:       1,
			},
			err: pkitypes.ErrRootCertificateIsNotSelfSigned,
		},
		{
			name: "PAA is true, CRL signer certificate contains vid != msg.vid",
			msg: MsgAddPkiRevocationDistributionPoint{
				Signer:               sample.AccAddress(),
				Vid:                  testconstants.Vid,
				IsPAA:                true,
				CrlSignerCertificate: testconstants.PAACertWithNumericVid,
				Label:                "label",
				DataURL:              testconstants.DataURL,
				IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
				RevocationType:       1,
			},
			err: pkitypes.ErrCRLSignerCertificateVidNotEqualMsgVid,
		},
		{
			name: "PAA is false, cert does not contain vid",
			msg: MsgAddPkiRevocationDistributionPoint{
				Signer:               sample.AccAddress(),
				Vid:                  testconstants.Vid,
				IsPAA:                false,
				CrlSignerCertificate: testconstants.IntermediateCertPem,
				Label:                "label",
				DataURL:              testconstants.DataURL,
				IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
				RevocationType:       1,
			},
			err: pkitypes.ErrCRLSignerCertificateVidNotEqualMsgVid,
		},
		{
			name: "PAA is false, cert pid provided, msg pid does not provided",
			msg: MsgAddPkiRevocationDistributionPoint{
				Signer:               sample.AccAddress(),
				Vid:                  testconstants.PAICertWithNumericPidVidVid,
				IsPAA:                false,
				CrlSignerCertificate: testconstants.PAICertWithNumericPidVid,
				Label:                "label",
				DataURL:              testconstants.DataURL,
				IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
				RevocationType:       1,
			},
			err: pkitypes.ErrPidNotFound,
		},
		{
			name: "PAA is false, cert pid != msg.pid",
			msg: MsgAddPkiRevocationDistributionPoint{
				Signer:               sample.AccAddress(),
				Vid:                  testconstants.PAICertWithNumericPidVidVid,
				IsPAA:                false,
				CrlSignerCertificate: testconstants.PAICertWithNumericPidVid,
				Label:                "label",
				DataURL:              testconstants.DataURL,
				Pid:                  1,
				IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
				RevocationType:       1,
			},
			err: pkitypes.ErrCRLSignerCertificatePidNotEqualMsgPid,
		},
		{
			name: "schemaVersion > 65535",
			msg: MsgAddPkiRevocationDistributionPoint{
				Signer:               sample.AccAddress(),
				Vid:                  testconstants.PAACertWithNumericVidVid,
				IsPAA:                true,
				CrlSignerCertificate: testconstants.PAACertWithNumericVid,
				Label:                "label",
				DataURL:              testconstants.DataURL,
				IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
				RevocationType:       1,
				SchemaVersion:        65536,
			},
			err: validator.ErrFieldUpperBoundViolated,
		},
	}

	positiveTests := []struct {
		name string
		msg  MsgAddPkiRevocationDistributionPoint
	}{
		{
			name: "minimal msg isPAA true",
			msg: MsgAddPkiRevocationDistributionPoint{
				Signer:               sample.AccAddress(),
				Vid:                  testconstants.PAACertWithNumericVidVid,
				IsPAA:                true,
				CrlSignerCertificate: testconstants.PAACertWithNumericVid,
				Label:                "label",
				DataURL:              testconstants.DataURL,
				IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
				RevocationType:       1,
			},
		},
		{
			name: "minimal msg isPAA false",
			msg: MsgAddPkiRevocationDistributionPoint{
				Signer:               sample.AccAddress(),
				Vid:                  testconstants.PAICertWithNumericPidVidVid,
				IsPAA:                false,
				Pid:                  testconstants.PAICertWithNumericPidVidPid,
				CrlSignerCertificate: testconstants.PAICertWithNumericPidVid,
				Label:                "label",
				DataURL:              testconstants.DataURL,
				IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
				RevocationType:       1,
			},
		},
		{
			name: "vid == cert.vid",
			msg: MsgAddPkiRevocationDistributionPoint{
				Signer:               sample.AccAddress(),
				Vid:                  testconstants.PAACertWithNumericVidVid,
				IsPAA:                true,
				CrlSignerCertificate: testconstants.PAACertWithNumericVid,
				Label:                "label",
				DataURL:              testconstants.DataURL,
				IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
				RevocationType:       1,
			},
		},
		{
			name: "vid == cert.vid, pid == cert.pid",
			msg: MsgAddPkiRevocationDistributionPoint{
				Signer:               sample.AccAddress(),
				Vid:                  testconstants.PAICertWithNumericPidVidVid,
				IsPAA:                false,
				CrlSignerCertificate: testconstants.PAICertWithNumericPidVid,
				Label:                "label",
				DataURL:              testconstants.DataURL,
				IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
				Pid:                  testconstants.PAICertWithNumericPidVidPid,
				RevocationType:       1,
			},
		},
		{
			name: "numeric MVid",
			msg: MsgAddPkiRevocationDistributionPoint{
				Signer:               sample.AccAddress(),
				Vid:                  testconstants.PAACertWithNumericVidVid,
				IsPAA:                true,
				CrlSignerCertificate: testconstants.PAACertWithNumericVid,
				Label:                "label",
				DataURL:              testconstants.DataURL,
				IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
				RevocationType:       1,
			},
		},
		{
			name: "PAA is true, cert non-vid scoped",
			msg: MsgAddPkiRevocationDistributionPoint{
				Signer:               sample.AccAddress(),
				Vid:                  testconstants.Vid,
				IsPAA:                true,
				CrlSignerCertificate: testconstants.RootCertPem,
				Label:                "label",
				DataURL:              testconstants.DataURL,
				IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
				RevocationType:       1,
			},
		},
		{
			name: "vid, pid encoded in certificate's subject as MVid, MPid",
			msg: MsgAddPkiRevocationDistributionPoint{
				Signer:               sample.AccAddress(),
				Vid:                  testconstants.PAICertWithPidVidVid,
				IsPAA:                false,
				CrlSignerCertificate: testconstants.PAICertWithPidVid,
				Label:                "label",
				DataURL:              testconstants.DataURL,
				Pid:                  testconstants.PAICertWithPidVidPid,
				IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
				RevocationType:       1,
			},
		},
		{
			name: "vid, pid encoded in certificate's subject as OID values",
			msg: MsgAddPkiRevocationDistributionPoint{
				Signer:               sample.AccAddress(),
				Vid:                  testconstants.PAICertWithNumericPidVidVid,
				IsPAA:                false,
				CrlSignerCertificate: testconstants.PAICertWithNumericPidVid,
				Label:                "label",
				DataURL:              testconstants.DataURL,
				IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
				Pid:                  testconstants.PAICertWithNumericPidVidPid,
				RevocationType:       1,
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
