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
	"testing"

	tmrand "github.com/cometbft/cometbft/libs/rand"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	"github.com/zigbee-alliance/distributed-compliance-ledger/testutil/sample"
	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/validator"
)

func TestMsgDeletePkiRevocationDistributionPoint_ValidateBasic(t *testing.T) {
	negativeTests := []struct {
		name string
		msg  MsgDeletePkiRevocationDistributionPoint
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeletePkiRevocationDistributionPoint{
				Signer: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "empty vid",
			msg: MsgDeletePkiRevocationDistributionPoint{
				Signer: sample.AccAddress(),
			},
			err: validator.ErrFieldLowerBoundViolated,
		},
		{
			name: "vid < 1",
			msg: MsgDeletePkiRevocationDistributionPoint{
				Signer:             sample.AccAddress(),
				Vid:                0,
				Label:              "label",
				IssuerSubjectKeyID: testconstants.SubjectKeyIDWithoutColons,
			},
			err: validator.ErrFieldLowerBoundViolated,
		},
		{
			name: "vid > 65535",
			msg: MsgDeletePkiRevocationDistributionPoint{
				Signer:             sample.AccAddress(),
				Vid:                65536,
				Label:              "label",
				IssuerSubjectKeyID: testconstants.SubjectKeyIDWithoutColons,
			},
			err: validator.ErrFieldUpperBoundViolated,
		},
		{
			name: "label empty",
			msg: MsgDeletePkiRevocationDistributionPoint{
				Signer: sample.AccAddress(),
				Vid:    testconstants.Vid,
			},
			err: validator.ErrRequiredFieldMissing,
		},
		{
			name: "issuerSubjectKeyID empty",
			msg: MsgDeletePkiRevocationDistributionPoint{
				Signer: sample.AccAddress(),
				Vid:    testconstants.Vid,
				Label:  "label",
			},
			err: validator.ErrRequiredFieldMissing,
		},
		{
			name: "wrong IssuerSubjectKeyID format (not [0-9A-F])",
			msg: MsgDeletePkiRevocationDistributionPoint{
				Signer:             sample.AccAddress(),
				Vid:                testconstants.Vid,
				Label:              "label",
				IssuerSubjectKeyID: "QWERTY",
			},
			err: pkitypes.ErrWrongSubjectKeyIDFormat,
		},
		{
			name: "wrong IssuerSubjectKeyID format (not even number of symbols)",
			msg: MsgDeletePkiRevocationDistributionPoint{
				Signer:             sample.AccAddress(),
				Vid:                testconstants.Vid,
				Label:              "label",
				IssuerSubjectKeyID: "123",
			},
			err: pkitypes.ErrWrongSubjectKeyIDFormat,
		},
		{
			name: "wrong IssuerSubjectKeyID format (not even number of symbols)",
			msg: MsgDeletePkiRevocationDistributionPoint{
				Signer:             sample.AccAddress(),
				Vid:                testconstants.Vid,
				Label:              "label",
				IssuerSubjectKeyID: "123",
			},
			err: pkitypes.ErrWrongSubjectKeyIDFormat,
		},
		{
			name: "issuerSubjectKeyID > 64",
			msg: MsgDeletePkiRevocationDistributionPoint{
				Signer:             sample.AccAddress(),
				Vid:                testconstants.PAACertWithNumericVidVid,
				Label:              "label",
				IssuerSubjectKeyID: tmrand.Str(65),
			},
			err: validator.ErrFieldMaxLengthExceeded,
		},
		{
			name: "label > 64",
			msg: MsgDeletePkiRevocationDistributionPoint{
				Signer:             sample.AccAddress(),
				Vid:                testconstants.PAACertWithNumericVidVid,
				Label:              tmrand.Str(65),
				IssuerSubjectKeyID: testconstants.SubjectKeyIDWithoutColons,
			},
			err: validator.ErrFieldMaxLengthExceeded,
		},
	}

	positiveTests := []struct {
		name string
		msg  MsgDeletePkiRevocationDistributionPoint
	}{
		{
			name: "example msg",
			msg: MsgDeletePkiRevocationDistributionPoint{
				Signer:             sample.AccAddress(),
				Vid:                testconstants.PAACertWithNumericVidVid,
				Label:              "label",
				IssuerSubjectKeyID: testconstants.SubjectKeyIDWithoutColons,
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
