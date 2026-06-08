package keeper_test

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	x509std "crypto/x509"
	"testing"

	"github.com/stretchr/testify/require"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	"github.com/zigbee-alliance/distributed-compliance-ledger/testutil/sample"
	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/x509"
)

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
			err: pkitypes.ErrCRLSignerCertificateInvalidVersion,
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

			err = keeper.VerifyCRLSignerCertFormat(cert)
			require.Error(t, err)
			require.ErrorIs(t, err, tt.err)
		})
	}

	// Positive case
	cert, err := x509.DecodeX509Certificate(testconstants.LeafCertWithVid)
	require.NoError(t, err)

	err = keeper.VerifyCRLSignerCertFormat(cert)
	require.NoError(t, err)
}

func TestMsgAddPkiRevocationDistributionPoint_VerifyCrlSignerCertificate(t *testing.T) {
	negativeTests := []struct {
		name string
		msg  types.MsgAddPkiRevocationDistributionPoint
		err  error
	}{
		{
			name: "pid not empty when IsPAA is true",
			msg: types.MsgAddPkiRevocationDistributionPoint{
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
			name: "pid provided when PAA is true",
			msg: types.MsgAddPkiRevocationDistributionPoint{
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
			msg: types.MsgAddPkiRevocationDistributionPoint{
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
			msg: types.MsgAddPkiRevocationDistributionPoint{
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
			msg: types.MsgAddPkiRevocationDistributionPoint{
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
			msg: types.MsgAddPkiRevocationDistributionPoint{
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
			msg: types.MsgAddPkiRevocationDistributionPoint{
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
	}

	positiveTests := []struct {
		name string
		msg  types.MsgAddPkiRevocationDistributionPoint
	}{
		{
			name: "minimal msg isPAA true",
			msg: types.MsgAddPkiRevocationDistributionPoint{
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
			msg: types.MsgAddPkiRevocationDistributionPoint{
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
			msg: types.MsgAddPkiRevocationDistributionPoint{
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
			msg: types.MsgAddPkiRevocationDistributionPoint{
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
			msg: types.MsgAddPkiRevocationDistributionPoint{
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
			msg: types.MsgAddPkiRevocationDistributionPoint{
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
			msg: types.MsgAddPkiRevocationDistributionPoint{
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
			msg: types.MsgAddPkiRevocationDistributionPoint{
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
			msg: types.MsgAddPkiRevocationDistributionPoint{
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
			msg: types.MsgAddPkiRevocationDistributionPoint{
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
			cert, err := x509.DecodeX509Certificate(tt.msg.CrlSignerCertificate)
			require.NoError(t, err)

			err = keeper.VerifyCrlSignerCertificate(cert, &tt.msg)
			require.Error(t, err)
			require.ErrorIs(t, err, tt.err)
		})
	}

	for _, tt := range positiveTests {
		t.Run(tt.name, func(t *testing.T) {
			cert, err := x509.DecodeX509Certificate(tt.msg.CrlSignerCertificate)
			require.NoError(t, err)

			err = keeper.VerifyCrlSignerCertificate(cert, &tt.msg)
			require.NoError(t, err)
		})
	}
}
