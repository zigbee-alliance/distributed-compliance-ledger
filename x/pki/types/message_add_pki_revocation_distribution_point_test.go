package types

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	x509std "crypto/x509"
	"fmt"
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	"github.com/zigbee-alliance/distributed-compliance-ledger/testutil/sample"
	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/validator"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/x509"
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
				Signer:        "invalid_address",
				SchemaVersion: 0,
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "empty vid",
			msg: MsgAddPkiRevocationDistributionPoint{
				Signer:        sample.AccAddress(),
				SchemaVersion: 0,
			},
			err: validator.ErrFieldLowerBoundViolated,
		},
		{
			name: "vid < 1",
			msg: MsgAddPkiRevocationDistributionPoint{
				Signer:        sample.AccAddress(),
				Vid:           0,
				SchemaVersion: 0,
			},
			err: validator.ErrFieldLowerBoundViolated,
		},
		{
			name: "vid > 65535",
			msg: MsgAddPkiRevocationDistributionPoint{
				Signer:        sample.AccAddress(),
				Vid:           65536,
				SchemaVersion: 0,
			},
			err: validator.ErrFieldUpperBoundViolated,
		},
		{
			name: "pid < 0",
			msg: MsgAddPkiRevocationDistributionPoint{
				Signer:        sample.AccAddress(),
				Vid:           1,
				Pid:           -1,
				SchemaVersion: 0,
			},
			err: validator.ErrFieldLowerBoundViolated,
		},
		{
			name: "pid < 65535",
			msg: MsgAddPkiRevocationDistributionPoint{
				Signer:        sample.AccAddress(),
				Vid:           1,
				Pid:           65536,
				SchemaVersion: 0,
			},
			err: validator.ErrFieldUpperBoundViolated,
		},
		{
			name: "IsPAA empty",
			msg: MsgAddPkiRevocationDistributionPoint{
				Signer:        sample.AccAddress(),
				Vid:           1,
				SchemaVersion: 0,
			},
			err: validator.ErrRequiredFieldMissing,
		},
		{
			name: "label empty",
			msg: MsgAddPkiRevocationDistributionPoint{
				Signer:        sample.AccAddress(),
				Vid:           1,
				IsPAA:         true,
				SchemaVersion: 0,
			},
			err: validator.ErrRequiredFieldMissing,
		},
		{
			name: "crl signer certificate empty",
			msg: MsgAddPkiRevocationDistributionPoint{
				Signer:        sample.AccAddress(),
				Vid:           1,
				IsPAA:         true,
				Label:         "label",
				SchemaVersion: 0,
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
				SchemaVersion:        0,
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
				SchemaVersion:        0,
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
				SchemaVersion:        0,
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
				SchemaVersion:        0,
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
				SchemaVersion:        0,
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
				SchemaVersion:        0,
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
				SchemaVersion:        0,
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
				SchemaVersion:        0,
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
				SchemaVersion:        0,
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
				SchemaVersion:        0,
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
				SchemaVersion:        0,
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
				SchemaVersion:        0,
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
				SchemaVersion:        0,
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
				SchemaVersion:        0,
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
				SchemaVersion:        0,
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
				SchemaVersion:        0,
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
				SchemaVersion:        0,
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
				SchemaVersion:        0,
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
				SchemaVersion:        0,
			},
			err: pkitypes.ErrNonRootCertificateSelfSigned,
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
				SchemaVersion:        0,
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
				SchemaVersion:        0,
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
				SchemaVersion:        0,
			},
			err: pkitypes.ErrCRLSignerCertificatePidNotEqualMsgPid,
		},
		{
			name: "schemaVersion != 0",
			msg: MsgAddPkiRevocationDistributionPoint{
				Signer:               sample.AccAddress(),
				Vid:                  testconstants.PAACertWithNumericVidVid,
				IsPAA:                true,
				CrlSignerCertificate: testconstants.PAACertWithNumericVid,
				Label:                "label",
				DataURL:              testconstants.DataURL,
				IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
				RevocationType:       1,
				SchemaVersion:        5,
			},
			err: validator.ErrFieldEqualBoundViolated,
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
				Vid:                  testconstants.LeafCertWithVidVid,
				IsPAA:                true,
				CrlSignerCertificate: testconstants.LeafCertWithVid,
				Label:                "label",
				DataURL:              testconstants.DataURL,
				IssuerSubjectKeyID:   testconstants.IntermediateCertWithVid1SubjectKeyIDWithoutColumns,
				RevocationType:       1,
				SchemaVersion:        0,
			},
		},
		{
			name: "minimal msg isPAA false",
			msg: MsgAddPkiRevocationDistributionPoint{
				Signer:               sample.AccAddress(),
				Vid:                  testconstants.LeafCertWithVidVid,
				IsPAA:                false,
				CrlSignerCertificate: testconstants.LeafCertWithVid,
				Label:                "label",
				DataURL:              testconstants.DataURL,
				IssuerSubjectKeyID:   testconstants.IntermediateCertWithVid1SubjectKeyIDWithoutColumns,
				RevocationType:       1,
				SchemaVersion:        0,
			},
		},
		{
			name: "vid == cert.vid",
			msg: MsgAddPkiRevocationDistributionPoint{
				Signer:               sample.AccAddress(),
				Vid:                  testconstants.LeafCertWithVidVid,
				IsPAA:                true,
				CrlSignerCertificate: testconstants.LeafCertWithVid,
				Label:                "label",
				DataURL:              testconstants.DataURL,
				IssuerSubjectKeyID:   testconstants.IntermediateCertWithVid1SubjectKeyIDWithoutColumns,
				RevocationType:       1,
				SchemaVersion:        0,
			},
		},
		{
			name: "vid == cert.vid, pid == cert.pid",
			msg: MsgAddPkiRevocationDistributionPoint{
				Signer:               sample.AccAddress(),
				Vid:                  testconstants.LeafCertWithVidPidVid,
				IsPAA:                false,
				CrlSignerCertificate: testconstants.LeafCertWithVidPid,
				Label:                "label",
				DataURL:              testconstants.DataURL,
				IssuerSubjectKeyID:   testconstants.IntermediateCertWithVid1SubjectKeyIDWithoutColumns,
				Pid:                  testconstants.LeafCertWithVidPidPid,
				RevocationType:       1,
				SchemaVersion:        0,
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
				SchemaVersion:        0,
			},
		},
		{
			name: "PAA is true, cert non-vid scoped",
			msg: MsgAddPkiRevocationDistributionPoint{
				Signer:               sample.AccAddress(),
				Vid:                  testconstants.Vid,
				IsPAA:                true,
				CrlSignerCertificate: testconstants.LeafCertWithoutVidPid,
				Label:                "label",
				DataURL:              testconstants.DataURL,
				IssuerSubjectKeyID:   testconstants.IntermediateCertWithVid1SubjectKeyIDWithoutColumns,
				RevocationType:       1,
				SchemaVersion:        0,
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
				SchemaVersion:        0,
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
				SchemaVersion:        0,
			},
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
				SchemaVersion:        0,
			},
		},
		{
			name: "PAA is false, cert pid == 0, msg.pid == 0",
			msg: MsgAddPkiRevocationDistributionPoint{
				Signer:               sample.AccAddress(),
				Vid:                  testconstants.PAICertWithNumericVidVid,
				IsPAA:                false,
				CrlSignerCertificate: testconstants.PAICertWithVid,
				Label:                "label",
				DataURL:              testconstants.DataURL,
				Pid:                  0,
				IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
				RevocationType:       1,
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
func TestMsgAddPkiRevocationDistributionPoint_verifyCRLCertFormat(t *testing.T) {
	negativeTests := []struct {
		name string
		init func(*x509.Certificate)
		err  error
	}{
		{
			name: "empty subject-key-id",
			init: func(certificate *x509.Certificate) {
				certificate.SubjectKeyID = ""
			},
			err: pkitypes.ErrWrongSubjectKeyIDFormat,
		},
		{
			name: "version is not v3",
			init: func(certificate *x509.Certificate) {
				certificate.Certificate.Version = 2
			},
			err: pkitypes.ErrCRLSignerCertificateInvalidFormat,
		},
		{
			name: "SignatureAlgorithm is not ECDSAWithSHA256",
			init: func(certificate *x509.Certificate) {
				certificate.Certificate.SignatureAlgorithm = x509std.ECDSAWithSHA384
			},
			err: pkitypes.ErrCRLSignerCertificateInvalidFormat,
		},
		{
			name: "PublicKey is not use prime256v1 curve",
			init: func(certificate *x509.Certificate) {
				ecdsaPubKey, _ := certificate.Certificate.PublicKey.(*ecdsa.PublicKey)
				ecdsaPubKey.Curve = elliptic.P224()
				certificate.Certificate.PublicKey = ecdsaPubKey
			},
			err: pkitypes.ErrCRLSignerCertificateInvalidFormat,
		},
		{
			name: "Key Usage extension is not critical",
			init: func(certificate *x509.Certificate) {
				certificate.Certificate.Extensions[3].Critical = false
			},
			err: pkitypes.ErrCRLSignerCertificateInvalidFormat,
		},
		{
			name: "The cRLSign bits is not in the KeyUsage bitstring",
			init: func(certificate *x509.Certificate) {
				certificate.Certificate.KeyUsage = x509std.KeyUsageCertSign
			},
			err: pkitypes.ErrCRLSignerCertificateInvalidFormat,
		},
		{
			name: "Other Key Usage bits expect KeyUsageCRLSign and KeyUsageDigitalSignature is not be set",
			init: func(certificate *x509.Certificate) {
				certificate.Certificate.KeyUsage = x509std.KeyUsageCertSign | x509std.KeyUsageCRLSign | x509std.KeyUsageDigitalSignature
			},
			err: pkitypes.ErrCRLSignerCertificateInvalidFormat,
		},
	}

	for _, tt := range negativeTests {
		t.Run(tt.name, func(t *testing.T) {
			cert, err := x509.DecodeX509Certificate(testconstants.LeafCertWithVid)
			require.NoError(t, err)

			tt.init(cert)

			err = VerifyCRLSignerCertFormat(cert)
			require.Error(t, err)
			require.ErrorIs(t, err, tt.err)
		})
	}

	// Positive case
	cert, err := x509.DecodeX509Certificate(testconstants.LeafCertWithVid)
	require.NoError(t, err)

	err = VerifyCRLSignerCertFormat(cert)
	require.NoError(t, err)
}
