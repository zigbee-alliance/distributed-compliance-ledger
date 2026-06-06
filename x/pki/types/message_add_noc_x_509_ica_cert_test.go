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

func TestMsgAddNocX509IcaCert_ValidateBasic(t *testing.T) {
	negativeTests := []struct {
		name string
		msg  MsgAddNocX509IcaCert
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgAddNocX509IcaCert{
				Signer: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "empty certificate",
			msg: MsgAddNocX509IcaCert{
				Signer: sample.AccAddress(),
				Cert:   "",
			},
			err: validator.ErrRequiredFieldMissing,
		},
		{
			name: "cert size > 20480 (20KB)",
			msg: MsgAddNocX509IcaCert{
				Signer: sample.AccAddress(),
				Cert:   tmrand.Str(20490),
			},
			err: validator.ErrFieldMaxLengthExceeded,
		},
		{
			name: "certSchemaVersion != 0",
			msg: MsgAddNocX509IcaCert{
				Signer:            sample.AccAddress(),
				Cert:              testconstants.NocCert1,
				CertSchemaVersion: 5,
			},
			err: validator.ErrFieldEqualBoundViolated,
		},
	}
	positiveTests := []struct {
		name string
		msg  MsgAddNocX509IcaCert
	}{
		{
			name: "valid add NOC cert msg",
			msg: MsgAddNocX509IcaCert{
				Signer: sample.AccAddress(),
				Cert:   testconstants.NocCert1,
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
