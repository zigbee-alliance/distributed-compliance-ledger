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

//nolint:testpackage
package types

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
)

/*
	MsgCreateValidator
*/

func TestNewMsgCreateValidator(t *testing.T) {
	msg := NewMsgCreateValidator(testconstants.ValidatorAddress1, testconstants.ValidatorPubKey1,
		Description{Name: testconstants.ProductName}, testconstants.Owner)

	require.Equal(t, msg.Route(), RouterKey)
	require.Equal(t, msg.Type(), "create_validator")
	require.Equal(t, msg.GetSigners(), []sdk.AccAddress{msg.Signer})
}

func TestValidateMsgCreateValidator(t *testing.T) {
	cases := []struct {
		valid bool
		msg   MsgCreateValidator
	}{
		{true, NewMsgCreateValidator(testconstants.ValidatorAddress1, testconstants.ValidatorPubKey1,
			Description{Name: testconstants.ProductName}, testconstants.Owner)},
		{false, NewMsgCreateValidator(nil, testconstants.ValidatorPubKey1,
			Description{Name: testconstants.ProductName}, testconstants.Owner)},
		{false, NewMsgCreateValidator(testconstants.ValidatorAddress1, "",
			Description{Name: testconstants.ProductName}, testconstants.Owner)},
		{false, NewMsgCreateValidator(testconstants.ValidatorAddress1, testconstants.ValidatorPubKey1,
			Description{}, testconstants.Owner)},
		{false, NewMsgCreateValidator(testconstants.ValidatorAddress1, testconstants.ValidatorPubKey1,
			Description{Name: testconstants.ProductName}, nil)},
	}

	for _, tc := range cases {
		err := tc.msg.ValidateBasic()

		if tc.valid {
			require.Nil(t, err)
		} else {
			require.NotNil(t, err)
		}
	}
}

func TestMsgCreateValidatorGetSignBytes(t *testing.T) {
	msg := NewMsgCreateValidator(testconstants.ValidatorAddress1, testconstants.ValidatorPubKey1,
		Description{Name: "Test"}, testconstants.Owner)

	expected := `{"type":"validator/CreateValidator","value":{` +
		`"description":{"name":"Test"},"signer":"cosmos1p72j8mgkf39qjzcmr283w8l8y9qv30qpj056uz",` +
		`"validator_address":"cosmosvalcons158uwzeqeu7zg332ztuzc5xh9k5uy3h5ttegzxd",` +
		`"validator_pubkey":"cosmosvalconspub1zcjduepqdmmjdfyvh2mrwl8p8wkwp23kh8lvjrd9u45snxqz6te6y6lwk6gqts45r3"}}`
	require.Equal(t, expected, string(msg.GetSignBytes()))
}
