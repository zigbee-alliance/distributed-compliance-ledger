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

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	"github.com/zigbee-alliance/distributed-compliance-ledger/testutil/sample"
)

func TestMsgApproveDisableValidator_ValidateBasic(t *testing.T) {
	negativeTests := []struct {
		name string
		msg  MsgApproveDisableValidator
		err  error
	}{
		{
			name: "invalid creator address",
			msg: MsgApproveDisableValidator{
				Creator: "invalid_address",
				Address: testconstants.ValidatorAddress1,
				Info:    testconstants.Info,
				Time:    testconstants.Time,
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "omitted creator address",
			msg: MsgApproveDisableValidator{
				Creator: "",
				Address: testconstants.ValidatorAddress1,
				Info:    testconstants.Info,
				Time:    testconstants.Time,
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "invalid validator address",
			msg: MsgApproveDisableValidator{
				Creator: testconstants.Address1.String(),
				Address: "invalid_address",
				Info:    testconstants.Info,
				Time:    testconstants.Time,
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "omitted validator address",
			msg: MsgApproveDisableValidator{
				Creator: testconstants.Address1.String(),
				Address: "",
				Info:    testconstants.Info,
				Time:    testconstants.Time,
			},
			err: sdkerrors.ErrInvalidAddress,
		},
	}

	positiveTests := []struct {
		name string
		msg  MsgApproveDisableValidator
	}{
		{
			name: "valid ApproveDisableValidator message",
			msg: MsgApproveDisableValidator{
				Creator: sample.AccAddress(),
				Address: testconstants.ValidatorAddress1,
				Info:    testconstants.Info,
				Time:    testconstants.Time,
			},
		},
	}
	for _, tt := range positiveTests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			require.NoError(t, err)
		})
	}

	for _, tt := range negativeTests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			require.Error(t, err)
			require.ErrorIs(t, err, tt.err)
		})
	}
}
