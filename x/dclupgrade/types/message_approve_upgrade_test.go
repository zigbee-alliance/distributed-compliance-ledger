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
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/validator"
)

func TestMsgApproveUpgrade_ValidateBasic(t *testing.T) {
	negativeTests := []struct {
		name string
		msg  MsgApproveUpgrade
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgApproveUpgrade{
				Creator: "invalid_address",
				Name:    testconstants.UpgradePlanName,
				Info:    testconstants.Info,
				Time:    testconstants.Time,
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "omitted address",
			msg: MsgApproveUpgrade{
				Creator: "",
				Name:    testconstants.UpgradePlanName,
				Info:    testconstants.Info,
				Time:    testconstants.Time,
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "name is not set",
			msg: MsgApproveUpgrade{
				Creator: sample.AccAddress(),
				Name:    "",
				Info:    testconstants.Info,
				Time:    testconstants.Time,
			},
			err: validator.ErrRequiredFieldMissing,
		},
	}

	positiveTests := []struct {
		name string
		msg  MsgApproveUpgrade
	}{
		{
			name: "valid MsgApproveUpgrade message",
			msg: MsgApproveUpgrade{
				Creator: sample.AccAddress(),
				Name:    testconstants.UpgradePlanName,
				Info:    testconstants.Info,
				Time:    testconstants.Time,
			},
		},
		{
			name: "info is not set",
			msg: MsgApproveUpgrade{
				Creator: sample.AccAddress(),
				Name:    testconstants.UpgradePlanName,
				Info:    "",
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
