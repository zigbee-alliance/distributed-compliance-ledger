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

//nolint:goimports
import (
	"testing"

	"git.dsr-corporation.com/zb-ledger/zb-ledger/integration_tests/constants"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestNewMsgAddModelInfo(t *testing.T) {
	msg := NewMsgAddModelInfo(testconstants.VID, testconstants.PID, testconstants.CID, testconstants.Name,
		testconstants.Description, testconstants.Sku, testconstants.FirmwareVersion,
		testconstants.HardwareVersion, testconstants.Custom, testconstants.TisOrTrpTestingCompleted, testconstants.Signer)

	require.Equal(t, msg.Route(), RouterKey)
	require.Equal(t, msg.Type(), "add_model_info")
	require.Equal(t, msg.GetSigners(), []sdk.AccAddress{testconstants.Signer})
}

func TestMsgAddModelInfoValidation(t *testing.T) {
	cases := []struct {
		valid bool
		msg   MsgAddModelInfo
	}{
		{true, NewMsgAddModelInfo(
			testconstants.VID, testconstants.PID, testconstants.CID, testconstants.Name, testconstants.Description,
			testconstants.Sku, testconstants.FirmwareVersion, testconstants.HardwareVersion,
			testconstants.Custom, testconstants.TisOrTrpTestingCompleted, testconstants.Signer)},
		{false, NewMsgAddModelInfo(
			testconstants.VID, 0, testconstants.CID, testconstants.Name, testconstants.Description,
			testconstants.Sku, testconstants.FirmwareVersion, testconstants.HardwareVersion,
			testconstants.Custom, testconstants.TisOrTrpTestingCompleted, testconstants.Signer)},
		{false, NewMsgAddModelInfo(
			testconstants.VID, 0, testconstants.CID, testconstants.Name, testconstants.Description,
			testconstants.Sku, testconstants.FirmwareVersion, testconstants.HardwareVersion,
			testconstants.Custom, testconstants.TisOrTrpTestingCompleted, testconstants.Signer)},
		{true, NewMsgAddModelInfo(
			testconstants.VID, testconstants.PID, 0, testconstants.Name, testconstants.Description,
			testconstants.Sku, testconstants.FirmwareVersion, testconstants.HardwareVersion,
			testconstants.Custom, testconstants.TisOrTrpTestingCompleted, testconstants.Signer)},
		{false, NewMsgAddModelInfo(
			testconstants.VID, testconstants.PID, testconstants.CID, "", testconstants.Description,
			testconstants.Sku, testconstants.FirmwareVersion, testconstants.HardwareVersion,
			testconstants.Custom, testconstants.TisOrTrpTestingCompleted, testconstants.Signer)},
		{false, NewMsgAddModelInfo(
			testconstants.VID, testconstants.PID, testconstants.CID, testconstants.Name, "",
			testconstants.Sku, testconstants.FirmwareVersion, testconstants.HardwareVersion,
			testconstants.Custom, testconstants.TisOrTrpTestingCompleted, testconstants.Signer)},
		{false, NewMsgAddModelInfo(
			testconstants.VID, testconstants.PID, testconstants.CID, testconstants.Name, testconstants.Description,
			"", testconstants.FirmwareVersion, testconstants.HardwareVersion,
			testconstants.Custom, testconstants.TisOrTrpTestingCompleted, testconstants.Signer)},
		{false, NewMsgAddModelInfo(
			testconstants.VID, testconstants.PID, testconstants.CID, testconstants.Name, testconstants.Description,
			testconstants.Sku, "", testconstants.HardwareVersion,
			testconstants.Custom, testconstants.TisOrTrpTestingCompleted, testconstants.Signer)},
		{false, NewMsgAddModelInfo(
			testconstants.VID, testconstants.PID, testconstants.CID, testconstants.Name, testconstants.Description,
			testconstants.Sku, testconstants.FirmwareVersion, "",
			testconstants.Custom, testconstants.TisOrTrpTestingCompleted, testconstants.Signer)},
		{true, NewMsgAddModelInfo(
			testconstants.VID, testconstants.PID, testconstants.CID, testconstants.Name, testconstants.Description,
			testconstants.Sku, testconstants.FirmwareVersion, testconstants.HardwareVersion,
			"", testconstants.TisOrTrpTestingCompleted, testconstants.Signer)},
		{false, NewMsgAddModelInfo(
			testconstants.VID, testconstants.PID, testconstants.CID, testconstants.Name, testconstants.Description,
			testconstants.Sku, testconstants.FirmwareVersion, testconstants.HardwareVersion,
			testconstants.Custom, testconstants.TisOrTrpTestingCompleted, nil)},
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

func TestMsgAddModelInfoGetSignBytes(t *testing.T) {
	msg := NewMsgAddModelInfo(testconstants.VID, testconstants.PID, testconstants.CID, testconstants.Name,
		testconstants.Description, testconstants.Sku, testconstants.FirmwareVersion, testconstants.HardwareVersion,
		testconstants.Custom, testconstants.TisOrTrpTestingCompleted, testconstants.Signer)

	expected := `{"type":"modelinfo/AddModelInfo","value":{"cid":12345,"custom":"Custom data",` +
		`"description":"Device Description","firmware_version":"1.0","hardware_version":"2.0","name":"Device Name",` +
		`"pid":22,"signer":"cosmos1p72j8mgkf39qjzcmr283w8l8y9qv30qpj056uz",` +
		`"sku":"RCU2205A","tis_or_trp_testing_completed":false,"vid":1}}`

	require.Equal(t, expected, string(msg.GetSignBytes()))
}
