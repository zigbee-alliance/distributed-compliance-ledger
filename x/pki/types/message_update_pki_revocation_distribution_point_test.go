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

func TestMsgUpdatePkiRevocationDistributionPoint_ValidateBasic(t *testing.T) {
	negativeTests := []struct {
		name string
		msg  MsgUpdatePkiRevocationDistributionPoint
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdatePkiRevocationDistributionPoint{
				Signer: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "empty vid",
			msg: MsgUpdatePkiRevocationDistributionPoint{
				Signer: sample.AccAddress(),
			},
			err: validator.ErrFieldLowerBoundViolated,
		},
		{
			name: "vid < 1",
			msg: MsgUpdatePkiRevocationDistributionPoint{
				Signer: sample.AccAddress(),
				Vid:    0,
			},
			err: validator.ErrFieldLowerBoundViolated,
		},
		{
			name: "vid > 65535",
			msg: MsgUpdatePkiRevocationDistributionPoint{
				Signer: sample.AccAddress(),
				Vid:    65536,
			},
			err: validator.ErrFieldUpperBoundViolated,
		},
		{
			name: "label empty",
			msg: MsgUpdatePkiRevocationDistributionPoint{
				Signer: sample.AccAddress(),
				Vid:    1,
			},
			err: validator.ErrRequiredFieldMissing,
		},
		{
			name: "issuerSubjectKeyId empty",
			msg: MsgUpdatePkiRevocationDistributionPoint{
				Signer:               sample.AccAddress(),
				Vid:                  1,
				CrlSignerCertificate: testconstants.RootCertPem,
				Label:                "label",
			},
			err: validator.ErrRequiredFieldMissing,
		},
		{
			name: fmt.Sprintf("dataDigestType is not one of %v", allowedDataDigestTypes),
			msg: MsgUpdatePkiRevocationDistributionPoint{
				Signer:               sample.AccAddress(),
				Vid:                  1,
				CrlSignerCertificate: testconstants.RootCertPem,
				Label:                "label",
				DataUrl:              testconstants.DataURL,
				IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
				DataDigestType:       3,
			},
			err: pkitypes.ErrInvalidDataDigestType,
		},
		{
			name: "dataUrl starts not with http or https",
			msg: MsgUpdatePkiRevocationDistributionPoint{
				Signer:               sample.AccAddress(),
				Vid:                  65521,
				CrlSignerCertificate: testconstants.RootCertWithPidVidInSubject,
				Label:                "label",
				DataUrl:              "ftp://" + testconstants.URLWithoutProtocol,
				IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
			},
			err: pkitypes.ErrInvalidDataUrlFormat,
		},
		{
			name: "dataUrl without protocol",
			msg: MsgUpdatePkiRevocationDistributionPoint{
				Signer:               sample.AccAddress(),
				Vid:                  65521,
				CrlSignerCertificate: testconstants.RootCertWithPidVidInSubject,
				Label:                "label",
				DataUrl:              testconstants.URLWithoutProtocol,
				IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
			},
			err: validator.ErrFieldNotValid,
		},
		{
			name: "dataDigest presented, DataFileSize not presented",
			msg: MsgUpdatePkiRevocationDistributionPoint{
				Signer:               sample.AccAddress(),
				Vid:                  1,
				CrlSignerCertificate: testconstants.RootCertPem,
				Label:                "label",
				DataUrl:              testconstants.DataURL,
				IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
				DataDigest:           testconstants.DataDigest,
			},
			err: pkitypes.ErrEmptyDataFileSize,
		},
		{
			name: "dataDigestType presented, DataDigest not presented",
			msg: MsgUpdatePkiRevocationDistributionPoint{
				Signer:               sample.AccAddress(),
				Vid:                  1,
				CrlSignerCertificate: testconstants.RootCertPem,
				Label:                "label",
				DataUrl:              testconstants.DataURL,
				IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
				DataDigestType:       1,
			},
			err: pkitypes.ErrNotEmptyDataDigestType,
		},
		{
			name: "dataDigest presented, DataDigestType not presented",
			msg: MsgUpdatePkiRevocationDistributionPoint{
				Signer:               sample.AccAddress(),
				Vid:                  1,
				CrlSignerCertificate: testconstants.RootCertPem,
				Label:                "label",
				DataUrl:              testconstants.DataURL,
				IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
				DataDigest:           testconstants.DataDigest,
				DataFileSize:         123,
			},
			err: pkitypes.ErrEmptyDataDigestType,
		},
		{
			name: "wrong IssuerSubjectKeyId format (not [0-9A-F])",
			msg: MsgUpdatePkiRevocationDistributionPoint{
				Signer:               sample.AccAddress(),
				Vid:                  1,
				CrlSignerCertificate: testconstants.RootCertPem,
				Label:                "label",
				DataUrl:              testconstants.DataURL,
				IssuerSubjectKeyID:   "QWERTY",
			},
			err: pkitypes.ErrWrongSubjectKeyIDFormat,
		},
		{
			name: "wrong IssuerSubjectKeyId format (not even number of symbols)",
			msg: MsgUpdatePkiRevocationDistributionPoint{
				Signer:               sample.AccAddress(),
				Vid:                  1,
				CrlSignerCertificate: testconstants.RootCertPem,
				Label:                "label",
				DataUrl:              testconstants.DataURL,
				IssuerSubjectKeyID:   "123",
			},
			err: pkitypes.ErrWrongSubjectKeyIDFormat,
		},
		{
			name: "wrong IssuerSubjectKeyId format (not even number of symbols)",
			msg: MsgUpdatePkiRevocationDistributionPoint{
				Signer:               sample.AccAddress(),
				Vid:                  1,
				CrlSignerCertificate: testconstants.RootCertPem,
				Label:                "label",
				DataUrl:              testconstants.DataURL,
				IssuerSubjectKeyID:   "123",
			},
			err: pkitypes.ErrWrongSubjectKeyIDFormat,
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
			},
		},
		{
			name: "maximum msg",
			msg: MsgUpdatePkiRevocationDistributionPoint{
				Signer:               sample.AccAddress(),
				Label:                "label",
				Vid:                  1,
				IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
				CrlSignerCertificate: testconstants.RootCertWithPidVidInSubject,
				DataUrl:              testconstants.DataURL,
				DataDigest:           testconstants.DataDigest,
				DataDigestType:       1,
				DataFileSize:         123,
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
