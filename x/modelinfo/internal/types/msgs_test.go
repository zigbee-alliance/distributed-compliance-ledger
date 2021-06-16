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
	msg := NewMsgAddModelInfo(testconstants.VID, testconstants.PID, testconstants.CID, testconstants.Name,
		testconstants.Description, testconstants.SKU, testconstants.SoftwareVersion, testconstants.SoftwareVersionString,
		testconstants.HardwareVersion, testconstants.HardwareVersionString, testconstants.CDVersionNumber,
		testconstants.FirmwareDigests, testconstants.Revoked, testconstants.OtaURL,
		testconstants.OtaChecksum, testconstants.OtaChecksumType, testconstants.OtaBlob,
		testconstants.CommissioningCustomFlow, testconstants.CommissioningCustomFlowUrl, testconstants.CommissioningModeInitialStepsHint,
		testconstants.CommissioningModeInitialStepsInstruction, testconstants.CommissioningModeSecondaryStepsHint,
		testconstants.CommissioningModeSecondaryStepsInstruction, testconstants.ReleaseNotesUrl, testconstants.UserManualUrl,
		testconstants.SupportUrl, testconstants.ProductURL, testconstants.ChipBlob, testconstants.VendorBlob,
		testconstants.Signer)

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
		{true, NewMsgAddModelInfo(testconstants.VID, testconstants.PID, testconstants.CID, testconstants.Name,
			testconstants.Description, testconstants.SKU, testconstants.SoftwareVersion, testconstants.SoftwareVersionString,
			testconstants.HardwareVersion, testconstants.HardwareVersionString, testconstants.CDVersionNumber,
			testconstants.FirmwareDigests, testconstants.Revoked, testconstants.OtaURL,
			testconstants.OtaChecksum, testconstants.OtaChecksumType, testconstants.OtaBlob,
			testconstants.CommissioningCustomFlow, testconstants.CommissioningCustomFlowUrl, testconstants.CommissioningModeInitialStepsHint,
			testconstants.CommissioningModeInitialStepsInstruction, testconstants.CommissioningModeSecondaryStepsHint,
			testconstants.CommissioningModeSecondaryStepsInstruction, testconstants.ReleaseNotesUrl, testconstants.UserManualUrl,
			testconstants.SupportUrl, testconstants.ProductURL, testconstants.ChipBlob, testconstants.VendorBlob,
			testconstants.Signer)},
		// Invalid VID
		{false, NewMsgAddModelInfo(0, testconstants.PID, testconstants.CID, testconstants.Name,
			testconstants.Description, testconstants.SKU, testconstants.SoftwareVersion, testconstants.SoftwareVersionString,
			testconstants.HardwareVersion, testconstants.HardwareVersionString, testconstants.CDVersionNumber,
			testconstants.FirmwareDigests, testconstants.Revoked, testconstants.OtaURL,
			testconstants.OtaChecksum, testconstants.OtaChecksumType, testconstants.OtaBlob,
			testconstants.CommissioningCustomFlow, testconstants.CommissioningCustomFlowUrl, testconstants.CommissioningModeInitialStepsHint,
			testconstants.CommissioningModeInitialStepsInstruction, testconstants.CommissioningModeSecondaryStepsHint,
			testconstants.CommissioningModeSecondaryStepsInstruction, testconstants.ReleaseNotesUrl, testconstants.UserManualUrl,
			testconstants.SupportUrl, testconstants.ProductURL, testconstants.ChipBlob, testconstants.VendorBlob,
			testconstants.Signer)},

		// Invalid PID
		{false, NewMsgAddModelInfo(testconstants.VID, 0, testconstants.CID, testconstants.Name,
			testconstants.Description, testconstants.SKU, testconstants.SoftwareVersion, testconstants.SoftwareVersionString,
			testconstants.HardwareVersion, testconstants.HardwareVersionString, testconstants.CDVersionNumber,
			testconstants.FirmwareDigests, testconstants.Revoked, testconstants.OtaURL,
			testconstants.OtaChecksum, testconstants.OtaChecksumType, testconstants.OtaBlob,
			testconstants.CommissioningCustomFlow, testconstants.CommissioningCustomFlowUrl, testconstants.CommissioningModeInitialStepsHint,
			testconstants.CommissioningModeInitialStepsInstruction, testconstants.CommissioningModeSecondaryStepsHint,
			testconstants.CommissioningModeSecondaryStepsInstruction, testconstants.ReleaseNotesUrl, testconstants.UserManualUrl,
			testconstants.SupportUrl, testconstants.ProductURL, testconstants.ChipBlob, testconstants.VendorBlob,
			testconstants.Signer)},

		// CID is optional
		{true, NewMsgAddModelInfo(testconstants.VID, testconstants.PID, 0, testconstants.Name,
			testconstants.Description, testconstants.SKU, testconstants.SoftwareVersion, testconstants.SoftwareVersionString,
			testconstants.HardwareVersion, testconstants.HardwareVersionString, testconstants.CDVersionNumber,
			testconstants.FirmwareDigests, testconstants.Revoked, testconstants.OtaURL,
			testconstants.OtaChecksum, testconstants.OtaChecksumType, testconstants.OtaBlob,
			testconstants.CommissioningCustomFlow, testconstants.CommissioningCustomFlowUrl, testconstants.CommissioningModeInitialStepsHint,
			testconstants.CommissioningModeInitialStepsInstruction, testconstants.CommissioningModeSecondaryStepsHint,
			testconstants.CommissioningModeSecondaryStepsInstruction, testconstants.ReleaseNotesUrl, testconstants.UserManualUrl,
			testconstants.SupportUrl, testconstants.ProductURL, testconstants.ChipBlob, testconstants.VendorBlob,
			testconstants.Signer)},

		// Name is mandatory
		{false, NewMsgAddModelInfo(testconstants.VID, testconstants.PID, testconstants.CID, "",
			testconstants.Description, testconstants.SKU, testconstants.SoftwareVersion, testconstants.SoftwareVersionString,
			testconstants.HardwareVersion, testconstants.HardwareVersionString, testconstants.CDVersionNumber,
			testconstants.FirmwareDigests, testconstants.Revoked, testconstants.OtaURL,
			testconstants.OtaChecksum, testconstants.OtaChecksumType, testconstants.OtaBlob,
			testconstants.CommissioningCustomFlow, testconstants.CommissioningCustomFlowUrl, testconstants.CommissioningModeInitialStepsHint,
			testconstants.CommissioningModeInitialStepsInstruction, testconstants.CommissioningModeSecondaryStepsHint,
			testconstants.CommissioningModeSecondaryStepsInstruction, testconstants.ReleaseNotesUrl, testconstants.UserManualUrl,
			testconstants.SupportUrl, testconstants.ProductURL, testconstants.ChipBlob, testconstants.VendorBlob,
			testconstants.Signer)},

		// Description is mandatory
		{false, NewMsgAddModelInfo(testconstants.VID, testconstants.PID, testconstants.CID, testconstants.Name,
			"", testconstants.SKU, testconstants.SoftwareVersion, testconstants.SoftwareVersionString,
			testconstants.HardwareVersion, testconstants.HardwareVersionString, testconstants.CDVersionNumber,
			testconstants.FirmwareDigests, testconstants.Revoked, testconstants.OtaURL,
			testconstants.OtaChecksum, testconstants.OtaChecksumType, testconstants.OtaBlob,
			testconstants.CommissioningCustomFlow, testconstants.CommissioningCustomFlowUrl, testconstants.CommissioningModeInitialStepsHint,
			testconstants.CommissioningModeInitialStepsInstruction, testconstants.CommissioningModeSecondaryStepsHint,
			testconstants.CommissioningModeSecondaryStepsInstruction, testconstants.ReleaseNotesUrl, testconstants.UserManualUrl,
			testconstants.SupportUrl, testconstants.ProductURL, testconstants.ChipBlob, testconstants.VendorBlob,
			testconstants.Signer)},

		// SKU is mandatory
		{false, NewMsgAddModelInfo(testconstants.VID, testconstants.PID, testconstants.CID, testconstants.Name,
			testconstants.Description, "", testconstants.SoftwareVersion, testconstants.SoftwareVersionString,
			testconstants.HardwareVersion, testconstants.HardwareVersionString, testconstants.CDVersionNumber,
			testconstants.FirmwareDigests, testconstants.Revoked, testconstants.OtaURL,
			testconstants.OtaChecksum, testconstants.OtaChecksumType, testconstants.OtaBlob,
			testconstants.CommissioningCustomFlow, testconstants.CommissioningCustomFlowUrl, testconstants.CommissioningModeInitialStepsHint,
			testconstants.CommissioningModeInitialStepsInstruction, testconstants.CommissioningModeSecondaryStepsHint,
			testconstants.CommissioningModeSecondaryStepsInstruction, testconstants.ReleaseNotesUrl, testconstants.UserManualUrl,
			testconstants.SupportUrl, testconstants.ProductURL, testconstants.ChipBlob, testconstants.VendorBlob,
			testconstants.Signer)},

		// SoftwareVersion is mandatory
		{false, NewMsgAddModelInfo(testconstants.VID, testconstants.PID, testconstants.CID, testconstants.Name,
			testconstants.Description, testconstants.SKU, 0, testconstants.SoftwareVersionString,
			testconstants.HardwareVersion, testconstants.HardwareVersionString, testconstants.CDVersionNumber,
			testconstants.FirmwareDigests, testconstants.Revoked, testconstants.OtaURL,
			testconstants.OtaChecksum, testconstants.OtaChecksumType, testconstants.OtaBlob,
			testconstants.CommissioningCustomFlow, testconstants.CommissioningCustomFlowUrl, testconstants.CommissioningModeInitialStepsHint,
			testconstants.CommissioningModeInitialStepsInstruction, testconstants.CommissioningModeSecondaryStepsHint,
			testconstants.CommissioningModeSecondaryStepsInstruction, testconstants.ReleaseNotesUrl, testconstants.UserManualUrl,
			testconstants.SupportUrl, testconstants.ProductURL, testconstants.ChipBlob, testconstants.VendorBlob,
			testconstants.Signer)},

		// SoftwareVersionString is mandatory
		{false, NewMsgAddModelInfo(testconstants.VID, testconstants.PID, testconstants.CID, testconstants.Name,
			testconstants.Description, testconstants.SKU, testconstants.SoftwareVersion, "",
			testconstants.HardwareVersion, testconstants.HardwareVersionString, testconstants.CDVersionNumber,
			testconstants.FirmwareDigests, testconstants.Revoked, testconstants.OtaURL,
			testconstants.OtaChecksum, testconstants.OtaChecksumType, testconstants.OtaBlob,
			testconstants.CommissioningCustomFlow, testconstants.CommissioningCustomFlowUrl, testconstants.CommissioningModeInitialStepsHint,
			testconstants.CommissioningModeInitialStepsInstruction, testconstants.CommissioningModeSecondaryStepsHint,
			testconstants.CommissioningModeSecondaryStepsInstruction, testconstants.ReleaseNotesUrl, testconstants.UserManualUrl,
			testconstants.SupportUrl, testconstants.ProductURL, testconstants.ChipBlob, testconstants.VendorBlob,
			testconstants.Signer)},

		// HardwareVersion is mandatory
		{false, NewMsgAddModelInfo(testconstants.VID, testconstants.PID, testconstants.CID, testconstants.Name,
			testconstants.Description, testconstants.SKU, testconstants.SoftwareVersion, testconstants.SoftwareVersionString,
			0, testconstants.HardwareVersionString, testconstants.CDVersionNumber,
			testconstants.FirmwareDigests, testconstants.Revoked, testconstants.OtaURL,
			testconstants.OtaChecksum, testconstants.OtaChecksumType, testconstants.OtaBlob,
			testconstants.CommissioningCustomFlow, testconstants.CommissioningCustomFlowUrl, testconstants.CommissioningModeInitialStepsHint,
			testconstants.CommissioningModeInitialStepsInstruction, testconstants.CommissioningModeSecondaryStepsHint,
			testconstants.CommissioningModeSecondaryStepsInstruction, testconstants.ReleaseNotesUrl, testconstants.UserManualUrl,
			testconstants.SupportUrl, testconstants.ProductURL, testconstants.ChipBlob, testconstants.VendorBlob,
			testconstants.Signer)},

		// HardwareVersionString is mandatory
		{false, NewMsgAddModelInfo(testconstants.VID, testconstants.PID, testconstants.CID, testconstants.Name,
			testconstants.Description, testconstants.SKU, testconstants.SoftwareVersion, testconstants.SoftwareVersionString,
			testconstants.HardwareVersion, "", testconstants.CDVersionNumber,
			testconstants.FirmwareDigests, testconstants.Revoked, testconstants.OtaURL,
			testconstants.OtaChecksum, testconstants.OtaChecksumType, testconstants.OtaBlob,
			testconstants.CommissioningCustomFlow, testconstants.CommissioningCustomFlowUrl, testconstants.CommissioningModeInitialStepsHint,
			testconstants.CommissioningModeInitialStepsInstruction, testconstants.CommissioningModeSecondaryStepsHint,
			testconstants.CommissioningModeSecondaryStepsInstruction, testconstants.ReleaseNotesUrl, testconstants.UserManualUrl,
			testconstants.SupportUrl, testconstants.ProductURL, testconstants.ChipBlob, testconstants.VendorBlob,
			testconstants.Signer)},

		// Ota Combination checks - Missing checksum
		{false, NewMsgAddModelInfo(testconstants.VID, testconstants.PID, testconstants.CID, testconstants.Name,
			testconstants.Description, testconstants.SKU, testconstants.SoftwareVersion, testconstants.SoftwareVersionString,
			testconstants.HardwareVersion, testconstants.HardwareVersionString, testconstants.CDVersionNumber,
			testconstants.FirmwareDigests, testconstants.Revoked, testconstants.OtaURL,
			"", testconstants.OtaChecksumType, testconstants.OtaBlob,
			testconstants.CommissioningCustomFlow, testconstants.CommissioningCustomFlowUrl, testconstants.CommissioningModeInitialStepsHint,
			testconstants.CommissioningModeInitialStepsInstruction, testconstants.CommissioningModeSecondaryStepsHint,
			testconstants.CommissioningModeSecondaryStepsInstruction, testconstants.ReleaseNotesUrl, testconstants.UserManualUrl,
			testconstants.SupportUrl, testconstants.ProductURL, testconstants.ChipBlob, testconstants.VendorBlob,
			testconstants.Signer)},

		// Ota Combination checks - Missing url and checksumType
		{false, NewMsgAddModelInfo(testconstants.VID, testconstants.PID, testconstants.CID, testconstants.Name,
			testconstants.Description, testconstants.SKU, testconstants.SoftwareVersion, testconstants.SoftwareVersionString,
			testconstants.HardwareVersion, testconstants.HardwareVersionString, testconstants.CDVersionNumber,
			testconstants.FirmwareDigests, testconstants.Revoked, "",
			testconstants.OtaChecksum, "", testconstants.OtaBlob,
			testconstants.CommissioningCustomFlow, testconstants.CommissioningCustomFlowUrl, testconstants.CommissioningModeInitialStepsHint,
			testconstants.CommissioningModeInitialStepsInstruction, testconstants.CommissioningModeSecondaryStepsHint,
			testconstants.CommissioningModeSecondaryStepsInstruction, testconstants.ReleaseNotesUrl, testconstants.UserManualUrl,
			testconstants.SupportUrl, testconstants.ProductURL, testconstants.ChipBlob, testconstants.VendorBlob,
			testconstants.Signer)},

		// Ota Combination checks - Missing url and checksum
		{false, NewMsgAddModelInfo(testconstants.VID, testconstants.PID, testconstants.CID, testconstants.Name,
			testconstants.Description, testconstants.SKU, testconstants.SoftwareVersion, testconstants.SoftwareVersionString,
			testconstants.HardwareVersion, testconstants.HardwareVersionString, testconstants.CDVersionNumber,
			testconstants.FirmwareDigests, testconstants.Revoked, "",
			"", testconstants.OtaChecksumType, testconstants.OtaBlob,
			testconstants.CommissioningCustomFlow, testconstants.CommissioningCustomFlowUrl, testconstants.CommissioningModeInitialStepsHint,
			testconstants.CommissioningModeInitialStepsInstruction, testconstants.CommissioningModeSecondaryStepsHint,
			testconstants.CommissioningModeSecondaryStepsInstruction, testconstants.ReleaseNotesUrl, testconstants.UserManualUrl,
			testconstants.SupportUrl, testconstants.ProductURL, testconstants.ChipBlob, testconstants.VendorBlob,
			testconstants.Signer)},

		// Ota Combination checks - Missing checksumType
		{false, NewMsgAddModelInfo(testconstants.VID, testconstants.PID, testconstants.CID, testconstants.Name,
			testconstants.Description, testconstants.SKU, testconstants.SoftwareVersion, testconstants.SoftwareVersionString,
			testconstants.HardwareVersion, testconstants.HardwareVersionString, testconstants.CDVersionNumber,
			testconstants.FirmwareDigests, testconstants.Revoked, testconstants.OtaURL,
			testconstants.OtaChecksum, "", testconstants.OtaBlob,
			testconstants.CommissioningCustomFlow, testconstants.CommissioningCustomFlowUrl, testconstants.CommissioningModeInitialStepsHint,
			testconstants.CommissioningModeInitialStepsInstruction, testconstants.CommissioningModeSecondaryStepsHint,
			testconstants.CommissioningModeSecondaryStepsInstruction, testconstants.ReleaseNotesUrl, testconstants.UserManualUrl,
			testconstants.SupportUrl, testconstants.ProductURL, testconstants.ChipBlob, testconstants.VendorBlob,
			testconstants.Signer)},

		// Ota Combination checks - Missing all three
		{true, NewMsgAddModelInfo(testconstants.VID, testconstants.PID, testconstants.CID, testconstants.Name,
			testconstants.Description, testconstants.SKU, testconstants.SoftwareVersion, testconstants.SoftwareVersionString,
			testconstants.HardwareVersion, testconstants.HardwareVersionString, testconstants.CDVersionNumber,
			testconstants.FirmwareDigests, testconstants.Revoked, "",
			"", "", testconstants.OtaBlob,
			testconstants.CommissioningCustomFlow, testconstants.CommissioningCustomFlowUrl, testconstants.CommissioningModeInitialStepsHint,
			testconstants.CommissioningModeInitialStepsInstruction, testconstants.CommissioningModeSecondaryStepsHint,
			testconstants.CommissioningModeSecondaryStepsInstruction, testconstants.ReleaseNotesUrl, testconstants.UserManualUrl,
			testconstants.SupportUrl, testconstants.ProductURL, testconstants.ChipBlob, testconstants.VendorBlob,
			testconstants.Signer)},

		// Missing non mandatory
		{true, NewMsgAddModelInfo(testconstants.VID, testconstants.PID, testconstants.CID, testconstants.Name,
			testconstants.Description, testconstants.SKU, testconstants.SoftwareVersion, testconstants.SoftwareVersionString,
			testconstants.HardwareVersion, testconstants.HardwareVersionString, testconstants.CDVersionNumber,
			testconstants.FirmwareDigests, testconstants.Revoked, "",
			"", "", "",
			testconstants.CommissioningCustomFlow, testconstants.CommissioningCustomFlowUrl, testconstants.CommissioningModeInitialStepsHint,
			testconstants.CommissioningModeInitialStepsInstruction, testconstants.CommissioningModeSecondaryStepsHint,
			"", testconstants.ReleaseNotesUrl, testconstants.UserManualUrl,
			testconstants.SupportUrl, testconstants.ProductURL, testconstants.ChipBlob, testconstants.VendorBlob,
			testconstants.Signer)},

		{true, NewMsgAddModelInfo(testconstants.VID, testconstants.PID, testconstants.CID, testconstants.Name,
			testconstants.Description, testconstants.SKU, testconstants.SoftwareVersion, testconstants.SoftwareVersionString,
			testconstants.HardwareVersion, testconstants.HardwareVersionString, testconstants.CDVersionNumber,
			testconstants.FirmwareDigests, testconstants.Revoked, "",
			"", "", "",
			testconstants.CommissioningCustomFlow, "", testconstants.CommissioningModeInitialStepsHint,
			testconstants.CommissioningModeInitialStepsInstruction, testconstants.CommissioningModeSecondaryStepsHint,
			"", testconstants.ReleaseNotesUrl, "",
			"", "", "", "",
			testconstants.Signer)},

		// Signer is nil
		{false, NewMsgAddModelInfo(testconstants.VID, testconstants.PID, testconstants.CID, testconstants.Name,
			testconstants.Description, testconstants.SKU, testconstants.SoftwareVersion, testconstants.SoftwareVersionString,
			testconstants.HardwareVersion, testconstants.HardwareVersionString, testconstants.CDVersionNumber,
			testconstants.FirmwareDigests, testconstants.Revoked, testconstants.OtaURL,
			testconstants.OtaChecksum, testconstants.OtaChecksumType, testconstants.OtaBlob,
			testconstants.CommissioningCustomFlow, testconstants.CommissioningCustomFlowUrl, testconstants.CommissioningModeInitialStepsHint,
			testconstants.CommissioningModeInitialStepsInstruction, testconstants.CommissioningModeSecondaryStepsHint,
			testconstants.CommissioningModeSecondaryStepsInstruction, testconstants.ReleaseNotesUrl, testconstants.UserManualUrl,
			testconstants.SupportUrl, testconstants.ProductURL, testconstants.ChipBlob, testconstants.VendorBlob,
			nil)},
		// Signer is nil
		{false, NewMsgAddModelInfo(testconstants.VID, testconstants.PID, testconstants.CID, testconstants.Name,
			testconstants.Description, testconstants.SKU, testconstants.SoftwareVersion, testconstants.SoftwareVersionString,
			testconstants.HardwareVersion, testconstants.HardwareVersionString, testconstants.CDVersionNumber,
			testconstants.FirmwareDigests, testconstants.Revoked, testconstants.OtaURL,
			testconstants.OtaChecksum, testconstants.OtaChecksumType, testconstants.OtaBlob,
			testconstants.CommissioningCustomFlow, testconstants.CommissioningCustomFlowUrl, testconstants.CommissioningModeInitialStepsHint,
			testconstants.CommissioningModeInitialStepsInstruction, testconstants.CommissioningModeSecondaryStepsHint,
			testconstants.CommissioningModeSecondaryStepsInstruction, testconstants.ReleaseNotesUrl, testconstants.UserManualUrl,
			testconstants.SupportUrl, testconstants.ProductURL, testconstants.ChipBlob, testconstants.VendorBlob,
			[]byte{})},
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
		testconstants.Description, testconstants.SKU, testconstants.SoftwareVersion, testconstants.SoftwareVersionString,
		testconstants.HardwareVersion, testconstants.HardwareVersionString, testconstants.CDVersionNumber,
		testconstants.FirmwareDigests, testconstants.Revoked, testconstants.OtaURL,
		testconstants.OtaChecksum, testconstants.OtaChecksumType, testconstants.OtaBlob,
		testconstants.CommissioningCustomFlow, testconstants.CommissioningCustomFlowUrl, testconstants.CommissioningModeInitialStepsHint,
		testconstants.CommissioningModeInitialStepsInstruction, testconstants.CommissioningModeSecondaryStepsHint,
		testconstants.CommissioningModeSecondaryStepsInstruction, testconstants.ReleaseNotesUrl, testconstants.UserManualUrl,
		testconstants.SupportUrl, testconstants.ProductURL, testconstants.ChipBlob, testconstants.VendorBlob,
		testconstants.Signer)

	expected := "{\"type\":\"modelinfo/AddModelInfo\",\"value\":{\"cd_version_number\":312,\"chip-blob\":\"Chip Blob Text\",\"cid\":12345,\"commisioning_mode_initial_steps_hint\":2,\"commisioning_mode_initial_steps_instruction\":\"commissioningModeInitialStepsInstruction details\",\"commisioning_mode_secondary_steps_hint\":3,\"commisioning_mode_secondary_steps_instruction\":\"commissioningModeSecondaryStepsInstruction steps\",\"commission_custom_flow\":1,\"commission_custom_flow_url\":\"https://sampleflowurl.dclmodel\",\"description\":\"Device Description\",\"firmware_digests\":\"Firmware Digest String\",\"hardware_version\":21,\"hardware_version_string\":\"2.1\",\"name\":\"Device Name\",\"ota_blob\":\"OTABlob Text\",\"ota_checksum\":\"0xe3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855\",\"ota_checksum_type\":\"SHA-256\",\"ota_url\":\"http://ota.firmware.com\",\"pid\":22,\"product-url\":\"https://url.producturl.dclmodel\",\"release_notes_url\":\"https://url.releasenotes.dclmodel\",\"revoked\":false,\"signer\":\"cosmos1p72j8mgkf39qjzcmr283w8l8y9qv30qpj056uz\",\"sku\":\"RCU2205A\",\"software_version\":1,\"software_version_string\":\"1.0\",\"support-url\":\"https://url.supporturl.dclmodel\",\"user-manual-url\":\"https://url.usermanual.dclmodel\",\"vendor-blob\":\"Vendor Blob Text\",\"vid\":1}}"

	require.Equal(t, expected, string(msg.GetSignBytes()))

}

func TestNewMsgUpdateModelInfo(t *testing.T) {
	msg := NewMsgUpdateModelInfo(testconstants.VID, testconstants.PID, testconstants.CID,
		testconstants.Description, testconstants.CDVersionNumber, testconstants.Revoked, testconstants.OtaURL,
		testconstants.OtaChecksum, testconstants.OtaChecksumType, testconstants.OtaBlob, testconstants.CommissioningCustomFlowUrl,
		testconstants.ReleaseNotesUrl, testconstants.UserManualUrl, testconstants.SupportUrl, testconstants.ProductURL,
		testconstants.ChipBlob, testconstants.VendorBlob, testconstants.Signer)

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
			testconstants.Description, testconstants.CDVersionNumber, false, testconstants.OtaURL,
			testconstants.OtaChecksum, testconstants.OtaChecksumType,
			testconstants.OtaBlob, testconstants.CommissioningCustomFlowUrl,
			testconstants.ReleaseNotesUrl, testconstants.UserManualUrl, testconstants.SupportUrl,
			testconstants.ProductURL, testconstants.ChipBlob, testconstants.VendorBlob, testconstants.Signer)},
		{false, NewMsgUpdateModelInfo(
			testconstants.VID, testconstants.PID, testconstants.CID,
			testconstants.Description, testconstants.CDVersionNumber, false, testconstants.OtaURL,
			testconstants.OtaChecksum, testconstants.OtaChecksumType,
			testconstants.OtaBlob, testconstants.CommissioningCustomFlowUrl,
			testconstants.ReleaseNotesUrl, testconstants.UserManualUrl, testconstants.SupportUrl,
			testconstants.ProductURL, testconstants.ChipBlob, testconstants.VendorBlob, nil)},
		{false, NewMsgUpdateModelInfo(
			testconstants.VID, testconstants.PID, testconstants.CID,
			testconstants.Description, testconstants.CDVersionNumber, false, testconstants.OtaURL,
			testconstants.OtaChecksum, testconstants.OtaChecksumType,
			testconstants.OtaBlob, testconstants.CommissioningCustomFlowUrl,
			testconstants.ReleaseNotesUrl, testconstants.UserManualUrl, testconstants.SupportUrl,
			testconstants.ProductURL, testconstants.ChipBlob, testconstants.VendorBlob, []byte{})},
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
		testconstants.Description, testconstants.CDVersionNumber, testconstants.Revoked, testconstants.OtaURL,
		testconstants.OtaChecksum, testconstants.OtaChecksumType, testconstants.OtaBlob, testconstants.CommissioningCustomFlowUrl,
		testconstants.ReleaseNotesUrl, testconstants.UserManualUrl, testconstants.SupportUrl, testconstants.ProductURL,
		testconstants.ChipBlob, testconstants.VendorBlob, testconstants.Signer)

	expected := "{\"type\":\"modelinfo/UpdateModelInfo\",\"value\":{\"cd_version_number\":312,\"chip-blob\":\"Chip Blob Text\",\"cid\":12345,\"commission_custom_flow_url\":\"https://sampleflowurl.dclmodel\",\"description\":\"Device Description\",\"ota_blob\":\"OTABlob Text\",\"ota_checksum\":\"0xe3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855\",\"ota_checksum_type\":\"SHA-256\",\"ota_url\":\"http://ota.firmware.com\",\"pid\":22,\"product-url\":\"https://url.producturl.dclmodel\",\"release_notes_url\":\"https://url.releasenotes.dclmodel\",\"revoked\":false,\"signer\":\"cosmos1p72j8mgkf39qjzcmr283w8l8y9qv30qpj056uz\",\"support-url\":\"https://url.supporturl.dclmodel\",\"user-manual-url\":\"https://url.usermanual.dclmodel\",\"vendor-blob\":\"Vendor Blob Text\",\"vid\":1}}"

	require.Equal(t, expected, string(msg.GetSignBytes()))
}
