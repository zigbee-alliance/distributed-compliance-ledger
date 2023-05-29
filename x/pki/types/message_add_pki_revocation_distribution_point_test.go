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
	true_ := true
	false_ := false
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
				IsPAA:  &true_,
			},
			err: validator.ErrRequiredFieldMissing,
		},
		{
			name: "crl signer certificate empty",
			msg: MsgAddPkiRevocationDistributionPoint{
				Signer: sample.AccAddress(),
				Vid:    1,
				IsPAA:  &true_,
				Label:  "label",
			},
			err: validator.ErrRequiredFieldMissing,
		},
		{
			name: "issuerSubjectKeyId empty",
			msg: MsgAddPkiRevocationDistributionPoint{
				Signer:               sample.AccAddress(),
				Vid:                  1,
				IsPAA:                &true_,
				CrlSignerCertificate: testconstants.RootCertPem,
				Label:                "label",
			},
			err: validator.ErrRequiredFieldMissing,
		},
		{
			name: "dataUrl empty",
			msg: MsgAddPkiRevocationDistributionPoint{
				Signer:               sample.AccAddress(),
				Vid:                  1,
				IsPAA:                &true_,
				CrlSignerCertificate: testconstants.RootCertPem,
				Label:                "label",
				IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
			},
			err: validator.ErrRequiredFieldMissing,
		},
		{
			name: "revocationType empty",
			msg: MsgAddPkiRevocationDistributionPoint{
				Signer:               sample.AccAddress(),
				Vid:                  1,
				IsPAA:                &true_,
				CrlSignerCertificate: testconstants.RootCertPem,
				Label:                "label",
				DataUrl:              testconstants.DataURL,
				IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
			},
			err: validator.ErrRequiredFieldMissing,
		},
		{
			name: fmt.Sprintf("dataDigestType is not one of %v", allowedDataDigestTypes),
			msg: MsgAddPkiRevocationDistributionPoint{
				Signer:               sample.AccAddress(),
				Vid:                  1,
				IsPAA:                &true_,
				CrlSignerCertificate: testconstants.RootCertPem,
				Label:                "label",
				DataUrl:              testconstants.DataURL,
				IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
				RevocationType:       1,
				DataDigestType:       3,
			},
			err: pkitypes.ErrInvalidDataDigestType,
		},
		{
			name: fmt.Sprintf("revocationType is not one of %v", allowedRevocationType),
			msg: MsgAddPkiRevocationDistributionPoint{
				Signer:               sample.AccAddress(),
				Vid:                  1,
				IsPAA:                &true_,
				CrlSignerCertificate: testconstants.RootCertPem,
				Label:                "label",
				DataUrl:              testconstants.DataURL,
				IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
				RevocationType:       2,
			},
			err: pkitypes.ErrInvalidRevocationType,
		},
		{
			name: "pid not empty when IsPAA is true",
			msg: MsgAddPkiRevocationDistributionPoint{
				Signer:               sample.AccAddress(),
				Vid:                  1,
				IsPAA:                &true_,
				CrlSignerCertificate: testconstants.RootCertPem,
				Label:                "label",
				DataUrl:              testconstants.DataURL,
				IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
				RevocationType:       1,
				Pid:                  1,
			},
			err: pkitypes.ErrNotEmptyPid,
		},
		{
			name: "dataUrl starts not with http or https",
			msg: MsgAddPkiRevocationDistributionPoint{
				Signer:               sample.AccAddress(),
				Vid:                  1,
				IsPAA:                &true_,
				CrlSignerCertificate: testconstants.RootCertPem,
				Label:                "label",
				DataUrl:              testconstants.URLWithoutProtocol,
				IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
				RevocationType:       1,
			},
			err: validator.ErrFieldNotValid,
		},
		{
			name: "dataDigest presented, DataFileSize not presented",
			msg: MsgAddPkiRevocationDistributionPoint{
				Signer:               sample.AccAddress(),
				Vid:                  1,
				IsPAA:                &true_,
				CrlSignerCertificate: testconstants.RootCertPem,
				Label:                "label",
				DataUrl:              testconstants.DataURL,
				IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
				DataDigest:           testconstants.DataDigest,
				RevocationType:       1,
			},
			err: pkitypes.ErrEmptyDataFileSize,
		},
		{
			name: "dataDigestType presented, DataDigest not presented",
			msg: MsgAddPkiRevocationDistributionPoint{
				Signer:               sample.AccAddress(),
				Vid:                  1,
				IsPAA:                &true_,
				CrlSignerCertificate: testconstants.RootCertPem,
				Label:                "label",
				DataUrl:              testconstants.DataURL,
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
				Vid:                  1,
				IsPAA:                &true_,
				CrlSignerCertificate: testconstants.RootCertPem,
				Label:                "label",
				DataUrl:              testconstants.DataURL,
				IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
				DataDigest:           testconstants.DataDigest,
				DataFileSize:         123,
				RevocationType:       1,
			},
			err: pkitypes.ErrEmptyDataDigestType,
		},
		{
			name: "wrong IssuerSubjectKeyId format (not [0-9A-F])",
			msg: MsgAddPkiRevocationDistributionPoint{
				Signer:               sample.AccAddress(),
				Vid:                  1,
				IsPAA:                &true_,
				CrlSignerCertificate: testconstants.RootCertPem,
				Label:                "label",
				DataUrl:              testconstants.DataURL,
				IssuerSubjectKeyID:   "QWERTY",
				RevocationType:       1,
			},
			err: pkitypes.ErrWrongSubjectKeyIDFormat,
		},
		{
			name: "wrong IssuerSubjectKeyId format (not even number of symbols)",
			msg: MsgAddPkiRevocationDistributionPoint{
				Signer:               sample.AccAddress(),
				Vid:                  1,
				IsPAA:                &true_,
				CrlSignerCertificate: testconstants.RootCertPem,
				Label:                "label",
				DataUrl:              testconstants.DataURL,
				IssuerSubjectKeyID:   "123",
				RevocationType:       1,
			},
			err: pkitypes.ErrWrongSubjectKeyIDFormat,
		},
		{
			name: "wrong IssuerSubjectKeyId format (not even number of symbols)",
			msg: MsgAddPkiRevocationDistributionPoint{
				Signer:               sample.AccAddress(),
				Vid:                  1,
				IsPAA:                &true_,
				CrlSignerCertificate: testconstants.RootCertPem,
				Label:                "label",
				DataUrl:              testconstants.DataURL,
				IssuerSubjectKeyID:   "123",
				RevocationType:       1,
			},
			err: pkitypes.ErrWrongSubjectKeyIDFormat,
		},
		{
			name: "data fields present when revocationType is 1",
			msg: MsgAddPkiRevocationDistributionPoint{
				Signer:               sample.AccAddress(),
				Vid:                  1,
				IsPAA:                &true_,
				CrlSignerCertificate: testconstants.RootCertPem,
				Label:                "label",
				DataUrl:              testconstants.DataURL,
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
				Vid:                  1,
				IsPAA:                &true_,
				CrlSignerCertificate: testconstants.RootCertPem,
				Label:                "label",
				DataUrl:              testconstants.DataURL,
				IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
				RevocationType:       1,
				Pid:                  1,
			},
			err: pkitypes.ErrNotEmptyPid,
		},
		{
			name: "pid not encoded in CRL signer certificate",
			msg: MsgAddPkiRevocationDistributionPoint{
				Signer:               sample.AccAddress(),
				Vid:                  1,
				IsPAA:                &false_,
				CrlSignerCertificate: testconstants.RootCertPem,
				Label:                "label",
				DataUrl:              testconstants.DataURL,
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
				Vid:                  1,
				IsPAA:                &false_,
				CrlSignerCertificate: testconstants.RootCertPem,
				Label:                "label",
				DataUrl:              testconstants.DataURL,
				IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
				RevocationType:       1,
			},
			err: pkitypes.ErrNonRootCertificateSelfSigned,
		},
		{
			name: "IsPAA true, certificate is non-root",
			msg: MsgAddPkiRevocationDistributionPoint{
				Signer:               sample.AccAddress(),
				Vid:                  1,
				IsPAA:                &true_,
				CrlSignerCertificate: testconstants.IntermediateCertPem,
				Label:                "label",
				DataUrl:              testconstants.DataURL,
				IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
				RevocationType:       1,
			},
			err: pkitypes.ErrRootCertificateIsNotSelfSigned,
		},
		{
			name: "PAA is true, CRL signer certificate contains vid != msg.vid",
			msg: MsgAddPkiRevocationDistributionPoint{
				Signer:               sample.AccAddress(),
				Vid:                  1,
				IsPAA:                &true_,
				CrlSignerCertificate: testconstants.RootCertWithPidVidInSubject,
				Label:                "label",
				DataUrl:              testconstants.DataURL,
				IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
				RevocationType:       1,
			},
			err: pkitypes.ErrCRLSignerCertificateVidNotEqualMsgVid,
		},
		{
			name: "PAA is false, cert does not contain vid",
			msg: MsgAddPkiRevocationDistributionPoint{
				Signer:               sample.AccAddress(),
				Vid:                  1,
				IsPAA:                &false_,
				CrlSignerCertificate: testconstants.IntermediateCertPem,
				Label:                "label",
				DataUrl:              testconstants.DataURL,
				IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
				RevocationType:       1,
			},
			err: pkitypes.ErrVidNotFound,
		},
		{
			name: "PAA is false, cert pid provided, msg pid does not provided",
			msg: MsgAddPkiRevocationDistributionPoint{
				Signer:               sample.AccAddress(),
				Vid:                  1,
				IsPAA:                &false_,
				CrlSignerCertificate: testconstants.NonRootCertWithPidVidInSubject,
				Label:                "label",
				DataUrl:              testconstants.DataURL,
				IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
				RevocationType:       1,
			},
			err: pkitypes.ErrPidNotFound,
		},
		{
			name: "PAA is false, cert pid != msg.pid",
			msg: MsgAddPkiRevocationDistributionPoint{
				Signer:               sample.AccAddress(),
				Vid:                  1,
				IsPAA:                &false_,
				CrlSignerCertificate: testconstants.NonRootCertWithPidVidInSubject,
				Label:                "label",
				DataUrl:              testconstants.DataURL,
				Pid:                  1,
				IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
				RevocationType:       1,
			},
			err: pkitypes.ErrCRLSignerCertificatePidNotEqualMsgPid,
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
				Vid:                  65521,
				IsPAA:                &true_,
				CrlSignerCertificate: testconstants.RootCertWithPidVidInSubject,
				Label:                "label",
				DataUrl:              testconstants.DataURL,
				IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
				RevocationType:       1,
			},
		},
		{
			name: "minimal msg isPAA false",
			msg: MsgAddPkiRevocationDistributionPoint{
				Signer:               sample.AccAddress(),
				Vid:                  65521,
				IsPAA:                &false_,
				Pid:                  32769,
				CrlSignerCertificate: testconstants.NonRootCertWithPidVidInSubject,
				Label:                "label",
				DataUrl:              testconstants.DataURL,
				IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
				RevocationType:       1,
			},
		},
		{
			name: "vid == cert.vid",
			msg: MsgAddPkiRevocationDistributionPoint{
				Signer:               sample.AccAddress(),
				Vid:                  65521,
				IsPAA:                &true_,
				CrlSignerCertificate: testconstants.RootCertWithPidVidInSubject,
				Label:                "label",
				DataUrl:              testconstants.DataURL,
				IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
				RevocationType:       1,
			},
		},
		{
			name: "vid == cert.vid, pid == cert.pid",
			msg: MsgAddPkiRevocationDistributionPoint{
				Signer:               sample.AccAddress(),
				Vid:                  65521,
				IsPAA:                &false_,
				CrlSignerCertificate: testconstants.NonRootCertWithPidVidInSubject,
				Label:                "label",
				DataUrl:              testconstants.DataURL,
				IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
				Pid:                  32769,
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
