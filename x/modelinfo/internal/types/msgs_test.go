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

func TestNewMsgAddModelInfo(t *testing.T) {
	msg := NewMsgAddModelInfo(testconstants.VID, testconstants.PID, testconstants.CID, testconstants.Version,
		testconstants.Name, testconstants.Description, testconstants.SKU, testconstants.HardwareVersion,
		testconstants.FirmwareVersion, testconstants.OtaURL, testconstants.OtaChecksum, testconstants.OtaChecksumType,
		testconstants.Custom, testconstants.TisOrTrpTestingCompleted, testconstants.Signer)

	require.Equal(t, msg.Route(), RouterKey)
	require.Equal(t, msg.Type(), "add_model_info")
	require.Equal(t, msg.GetSigners(), []sdk.AccAddress{testconstants.Signer})
}

//nolint:funlen
func TestMsgAddModelInfoValidation(t *testing.T) {
	cases := []struct {
		valid bool
		msg   MsgAddModelInfo
	}{
		{true, NewMsgAddModelInfo(
			testconstants.VID, testconstants.PID, testconstants.CID, testconstants.Version,
			testconstants.Name, testconstants.Description, testconstants.SKU, testconstants.HardwareVersion,
			testconstants.FirmwareVersion, testconstants.OtaURL, testconstants.OtaChecksum, testconstants.OtaChecksumType,
			testconstants.Custom, testconstants.TisOrTrpTestingCompleted, testconstants.Signer)},
		{false, NewMsgAddModelInfo(
			0, testconstants.PID, testconstants.CID, testconstants.Version,
			testconstants.Name, testconstants.Description, testconstants.SKU, testconstants.HardwareVersion,
			testconstants.FirmwareVersion, testconstants.OtaURL, testconstants.OtaChecksum, testconstants.OtaChecksumType,
			testconstants.Custom, testconstants.TisOrTrpTestingCompleted, testconstants.Signer)},
		{false, NewMsgAddModelInfo(
			testconstants.VID, 0, testconstants.CID, testconstants.Version,
			testconstants.Name, testconstants.Description, testconstants.SKU, testconstants.HardwareVersion,
			testconstants.FirmwareVersion, testconstants.OtaURL, testconstants.OtaChecksum, testconstants.OtaChecksumType,
			testconstants.Custom, testconstants.TisOrTrpTestingCompleted, testconstants.Signer)},
		{true, NewMsgAddModelInfo(
			testconstants.VID, testconstants.PID, 0, testconstants.Version,
			testconstants.Name, testconstants.Description, testconstants.SKU, testconstants.HardwareVersion,
			testconstants.FirmwareVersion, testconstants.OtaURL, testconstants.OtaChecksum, testconstants.OtaChecksumType,
			testconstants.Custom, testconstants.TisOrTrpTestingCompleted, testconstants.Signer)},
		{true, NewMsgAddModelInfo(
			testconstants.VID, testconstants.PID, testconstants.CID, "",
			testconstants.Name, testconstants.Description, testconstants.SKU, testconstants.HardwareVersion,
			testconstants.FirmwareVersion, testconstants.OtaURL, testconstants.OtaChecksum, testconstants.OtaChecksumType,
			testconstants.Custom, testconstants.TisOrTrpTestingCompleted, testconstants.Signer)},
		{false, NewMsgAddModelInfo(
			testconstants.VID, testconstants.PID, testconstants.CID, testconstants.Version,
			"", testconstants.Description, testconstants.SKU, testconstants.HardwareVersion,
			testconstants.FirmwareVersion, testconstants.OtaURL, testconstants.OtaChecksum, testconstants.OtaChecksumType,
			testconstants.Custom, testconstants.TisOrTrpTestingCompleted, testconstants.Signer)},
		{false, NewMsgAddModelInfo(
			testconstants.VID, testconstants.PID, testconstants.CID, testconstants.Version,
			testconstants.Name, "", testconstants.SKU, testconstants.HardwareVersion,
			testconstants.FirmwareVersion, testconstants.OtaURL, testconstants.OtaChecksum, testconstants.OtaChecksumType,
			testconstants.Custom, testconstants.TisOrTrpTestingCompleted, testconstants.Signer)},
		{false, NewMsgAddModelInfo(
			testconstants.VID, testconstants.PID, testconstants.CID, testconstants.Version,
			testconstants.Name, testconstants.Description, "", testconstants.HardwareVersion,
			testconstants.FirmwareVersion, testconstants.OtaURL, testconstants.OtaChecksum, testconstants.OtaChecksumType,
			testconstants.Custom, testconstants.TisOrTrpTestingCompleted, testconstants.Signer)},
		{false, NewMsgAddModelInfo(
			testconstants.VID, testconstants.PID, testconstants.CID, testconstants.Version,
			testconstants.Name, testconstants.Description, testconstants.SKU, "",
			testconstants.FirmwareVersion, testconstants.OtaURL, testconstants.OtaChecksum, testconstants.OtaChecksumType,
			testconstants.Custom, testconstants.TisOrTrpTestingCompleted, testconstants.Signer)},
		{false, NewMsgAddModelInfo(
			testconstants.VID, testconstants.PID, testconstants.CID, testconstants.Version,
			testconstants.Name, testconstants.Description, testconstants.SKU, testconstants.HardwareVersion,
			"", testconstants.OtaURL, testconstants.OtaChecksum, testconstants.OtaChecksumType,
			testconstants.Custom, testconstants.TisOrTrpTestingCompleted, testconstants.Signer)},
		{false, NewMsgAddModelInfo(
			testconstants.VID, testconstants.PID, testconstants.CID, testconstants.Version,
			testconstants.Name, testconstants.Description, testconstants.SKU, testconstants.HardwareVersion,
			testconstants.FirmwareVersion, "", testconstants.OtaChecksum, testconstants.OtaChecksumType,
			testconstants.Custom, testconstants.TisOrTrpTestingCompleted, testconstants.Signer)},
		{false, NewMsgAddModelInfo(
			testconstants.VID, testconstants.PID, testconstants.CID, testconstants.Version,
			testconstants.Name, testconstants.Description, testconstants.SKU, testconstants.HardwareVersion,
			testconstants.FirmwareVersion, testconstants.OtaURL, "", testconstants.OtaChecksumType,
			testconstants.Custom, testconstants.TisOrTrpTestingCompleted, testconstants.Signer)},
		{false, NewMsgAddModelInfo(
			testconstants.VID, testconstants.PID, testconstants.CID, testconstants.Version,
			testconstants.Name, testconstants.Description, testconstants.SKU, testconstants.HardwareVersion,
			testconstants.FirmwareVersion, testconstants.OtaURL, testconstants.OtaChecksum, "",
			testconstants.Custom, testconstants.TisOrTrpTestingCompleted, testconstants.Signer)},
		{false, NewMsgAddModelInfo(
			testconstants.VID, testconstants.PID, testconstants.CID, testconstants.Version,
			testconstants.Name, testconstants.Description, testconstants.SKU, testconstants.HardwareVersion,
			testconstants.FirmwareVersion, "", "", testconstants.OtaChecksumType,
			testconstants.Custom, testconstants.TisOrTrpTestingCompleted, testconstants.Signer)},
		{false, NewMsgAddModelInfo(
			testconstants.VID, testconstants.PID, testconstants.CID, testconstants.Version,
			testconstants.Name, testconstants.Description, testconstants.SKU, testconstants.HardwareVersion,
			testconstants.FirmwareVersion, "", testconstants.OtaChecksum, "",
			testconstants.Custom, testconstants.TisOrTrpTestingCompleted, testconstants.Signer)},
		{false, NewMsgAddModelInfo(
			testconstants.VID, testconstants.PID, testconstants.CID, testconstants.Version,
			testconstants.Name, testconstants.Description, testconstants.SKU, testconstants.HardwareVersion,
			testconstants.FirmwareVersion, testconstants.OtaURL, "", "",
			testconstants.Custom, testconstants.TisOrTrpTestingCompleted, testconstants.Signer)},
		{true, NewMsgAddModelInfo(
			testconstants.VID, testconstants.PID, testconstants.CID, testconstants.Version,
			testconstants.Name, testconstants.Description, testconstants.SKU, testconstants.HardwareVersion,
			testconstants.FirmwareVersion, "", "", "",
			testconstants.Custom, testconstants.TisOrTrpTestingCompleted, testconstants.Signer)},
		{true, NewMsgAddModelInfo(
			testconstants.VID, testconstants.PID, testconstants.CID, testconstants.Version,
			testconstants.Name, testconstants.Description, testconstants.SKU, testconstants.HardwareVersion,
			testconstants.FirmwareVersion, testconstants.OtaURL, testconstants.OtaChecksum, testconstants.OtaChecksumType,
			"", testconstants.TisOrTrpTestingCompleted, testconstants.Signer)},
		{false, NewMsgAddModelInfo(
			testconstants.VID, testconstants.PID, testconstants.CID, testconstants.Version,
			testconstants.Name, testconstants.Description, testconstants.SKU, testconstants.HardwareVersion,
			testconstants.FirmwareVersion, testconstants.OtaURL, testconstants.OtaChecksum, testconstants.OtaChecksumType,
			testconstants.Custom, testconstants.TisOrTrpTestingCompleted, nil)},
		{false, NewMsgAddModelInfo(
			testconstants.VID, testconstants.PID, testconstants.CID, testconstants.Version,
			testconstants.Name, testconstants.Description, testconstants.SKU, testconstants.HardwareVersion,
			testconstants.FirmwareVersion, testconstants.OtaURL, testconstants.OtaChecksum, testconstants.OtaChecksumType,
			testconstants.Custom, testconstants.TisOrTrpTestingCompleted, []byte{})},
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
	msg := NewMsgAddModelInfo(testconstants.VID, testconstants.PID, testconstants.CID, testconstants.Version,
		testconstants.Name, testconstants.Description, testconstants.SKU, testconstants.HardwareVersion,
		testconstants.FirmwareVersion, testconstants.OtaURL, testconstants.OtaChecksum, testconstants.OtaChecksumType,
		testconstants.Custom, testconstants.TisOrTrpTestingCompleted, testconstants.Signer)

	expected := `{"type":"modelinfo/AddModelInfo","value":{` +
		`"cid":12345,"custom":"Custom data","description":"Device Description",` +
		`"firmware_version":"2.0","hardware_version":"1.1","name":"Device Name",` +
		`"ota_checksum":"0xe3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",` +
		`"ota_checksum_type":"SHA-256","ota_url":"http://ota.firmware.com","pid":22,` +
		`"signer":"cosmos1p72j8mgkf39qjzcmr283w8l8y9qv30qpj056uz","sku":"RCU2205A",` +
		`"tis_or_trp_testing_completed":true,"version":"1.0","vid":1}}`

	require.Equal(t, expected, string(msg.GetSignBytes()))
}

func TestNewMsgUpdateModelInfo(t *testing.T) {
	msg := NewMsgUpdateModelInfo(testconstants.VID, testconstants.PID, testconstants.CID,
		testconstants.Description, testconstants.OtaURL, testconstants.Custom,
		testconstants.TisOrTrpTestingCompleted, testconstants.Signer)

	require.Equal(t, msg.Route(), RouterKey)
	require.Equal(t, msg.Type(), "update_model_info")
	require.Equal(t, msg.GetSigners(), []sdk.AccAddress{testconstants.Signer})
}

func TestMsgUpdateModelInfoValidation(t *testing.T) {
	cases := []struct {
		valid bool
		msg   MsgUpdateModelInfo
	}{
		{true, NewMsgUpdateModelInfo(
			testconstants.VID, testconstants.PID, testconstants.CID,
			testconstants.Description, testconstants.OtaURL, testconstants.Custom,
			testconstants.TisOrTrpTestingCompleted, testconstants.Signer)},
		{false, NewMsgUpdateModelInfo(
			0, testconstants.PID, testconstants.CID,
			testconstants.Description, testconstants.OtaURL, testconstants.Custom,
			testconstants.TisOrTrpTestingCompleted, testconstants.Signer)},
		{false, NewMsgUpdateModelInfo(
			testconstants.VID, 0, testconstants.CID,
			testconstants.Description, testconstants.OtaURL, testconstants.Custom,
			testconstants.TisOrTrpTestingCompleted, testconstants.Signer)},
		{true, NewMsgUpdateModelInfo(
			testconstants.VID, testconstants.PID, 0,
			testconstants.Description, testconstants.OtaURL, testconstants.Custom,
			testconstants.TisOrTrpTestingCompleted, testconstants.Signer)},
		{true, NewMsgUpdateModelInfo(
			testconstants.VID, testconstants.PID, testconstants.CID,
			"", testconstants.OtaURL, testconstants.Custom,
			testconstants.TisOrTrpTestingCompleted, testconstants.Signer)},
		{true, NewMsgUpdateModelInfo(
			testconstants.VID, testconstants.PID, testconstants.CID,
			testconstants.Description, "", testconstants.Custom,
			testconstants.TisOrTrpTestingCompleted, testconstants.Signer)},
		{true, NewMsgUpdateModelInfo(
			testconstants.VID, testconstants.PID, testconstants.CID,
			testconstants.Description, testconstants.OtaURL, "",
			testconstants.TisOrTrpTestingCompleted, testconstants.Signer)},
		{false, NewMsgUpdateModelInfo(
			testconstants.VID, testconstants.PID, testconstants.CID,
			testconstants.Description, testconstants.OtaURL, testconstants.Custom,
			testconstants.TisOrTrpTestingCompleted, nil)},
		{false, NewMsgUpdateModelInfo(
			testconstants.VID, testconstants.PID, testconstants.CID,
			testconstants.Description, testconstants.OtaURL, testconstants.Custom,
			testconstants.TisOrTrpTestingCompleted, []byte{})},
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

func TestMsgUpdateModelInfoGetSignBytes(t *testing.T) {
	msg := NewMsgUpdateModelInfo(testconstants.VID, testconstants.PID, testconstants.CID,
		testconstants.Description, testconstants.OtaURL, testconstants.Custom,
		testconstants.TisOrTrpTestingCompleted, testconstants.Signer)

	expected := `{"type":"modelinfo/UpdateModelInfo","value":{` +
		`"cid":12345,"custom":"Custom data","description":"Device Description",` +
		`"ota_url":"http://ota.firmware.com","pid":22,` +
		`"signer":"cosmos1p72j8mgkf39qjzcmr283w8l8y9qv30qpj056uz",` +
		`"tis_or_trp_testing_completed":true,"vid":1}}`

	require.Equal(t, expected, string(msg.GetSignBytes()))
}
