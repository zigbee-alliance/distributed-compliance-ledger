// Copyright 2020 DSR Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package types

import (
	"fmt"
	"testing"

	tmrand "github.com/cometbft/cometbft/libs/rand"
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
			name: "label > 64",
			msg: MsgAddPkiRevocationDistributionPoint{
				Signer:        sample.AccAddress(),
				Vid:           1,
				IsPAA:         true,
				SchemaVersion: 0,
				Label:         tmrand.Str(65),
			},
			err: validator.ErrFieldMaxLengthExceeded,
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
			name: "dataURL with invalid protocol (ftp)",
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
			err: validator.ErrFieldNotValid,
		},
		{
			name: "dataURL without http or https protocol",
			msg: MsgAddPkiRevocationDistributionPoint{
				Signer:               sample.AccAddress(),
				Vid:                  testconstants.PAACertWithNumericVidVid,
				IsPAA:                true,
				CrlSignerCertificate: testconstants.PAACertWithNumericVid,
				Label:                "label",
				DataURL:              testconstants.URLStartsWithW3,
				IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
				RevocationType:       1,
				SchemaVersion:        0,
			},
			err: validator.ErrFieldNotValid,
		},
		{
			name: "dataURL length > 256",
			msg: MsgAddPkiRevocationDistributionPoint{
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
				RevocationType:     1,
				SchemaVersion:      0,
			},
			err: validator.ErrFieldMaxLengthExceeded,
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
			name: "dataDigest length > 128",
			msg: MsgAddPkiRevocationDistributionPoint{
				Signer:               sample.AccAddress(),
				Vid:                  testconstants.PAACertWithNumericVidVid,
				IsPAA:                true,
				CrlSignerCertificate: testconstants.PAACertWithNumericVid,
				Label:                "label",
				DataURL:              testconstants.DataURL,
				IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
				DataDigest: func() string {
					longDataDigest := testconstants.DataDigest
					for i := 0; i < 5; i++ {
						longDataDigest += testconstants.DataDigest
					}

					return longDataDigest
				}(),
				DataDigestType: 1,
				DataFileSize:   123,
				RevocationType: 1,
				SchemaVersion:  0,
			},
			err: validator.ErrFieldMaxLengthExceeded,
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
			name: "when CrlSignerCertificate size exceeds 2KB",
			msg: MsgAddPkiRevocationDistributionPoint{
				Signer:               sample.AccAddress(),
				Vid:                  testconstants.PAACertWithNumericVidVid,
				IsPAA:                true,
				CrlSignerCertificate: testconstants.CertWithSizeGreater2KB,
				Label:                "label",
				DataURL:              testconstants.DataURL,
				IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
				RevocationType:       1,
				SchemaVersion:        0,
			},
			err: validator.ErrFieldMaxLengthExceeded,
		},
		{
			name: "when CrlSignerDelegator size exceeds 2KB",
			msg: MsgAddPkiRevocationDistributionPoint{
				Signer:               sample.AccAddress(),
				Vid:                  testconstants.PAACertWithNumericVidVid,
				IsPAA:                true,
				CrlSignerCertificate: testconstants.PAACertWithNumericVid,
				Label:                "label",
				DataURL:              testconstants.DataURL,
				IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
				CrlSignerDelegator:   testconstants.CertWithSizeGreater2KB,
				RevocationType:       1,
				SchemaVersion:        0,
			},
			err: validator.ErrFieldMaxLengthExceeded,
		},
		{
			name: "issuerSubjectKeyID > 64",
			msg: MsgAddPkiRevocationDistributionPoint{
				Signer:               sample.AccAddress(),
				Vid:                  testconstants.PAACertWithNumericVidVid,
				IsPAA:                true,
				CrlSignerCertificate: testconstants.PAACertWithNumericVid,
				Label:                "label",
				DataURL:              testconstants.DataURL,
				IssuerSubjectKeyID:   tmrand.Str(65),
				RevocationType:       1,
				SchemaVersion:        0,
			},
			err: validator.ErrFieldMaxLengthExceeded,
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
