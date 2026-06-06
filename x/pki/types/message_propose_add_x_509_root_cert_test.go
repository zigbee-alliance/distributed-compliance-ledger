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
			name: "cert size > 20480 (20KB)",
			msg: MsgProposeAddX509RootCert{
				Signer: sample.AccAddress(),
				Cert:   tmrand.Str(20490),
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
			name: "certSchemaVersion != 0",
			msg: MsgProposeAddX509RootCert{
				Signer:            sample.AccAddress(),
				Cert:              testconstants.RootCertPem,
				Info:              testconstants.Info,
				Time:              12345,
				Vid:               testconstants.Vid,
				CertSchemaVersion: 5,
			},
			err: validator.ErrFieldEqualBoundViolated,
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
