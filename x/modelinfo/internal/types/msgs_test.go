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

func getTestModel() Model {
	return Model{
		VID:                                      testconstants.VID,
		PID:                                      testconstants.PID,
		CID:                                      testconstants.CID,
		Name:                                     testconstants.Name,
		Description:                              testconstants.Description,
		SKU:                                      testconstants.SKU,
		SoftwareVersion:                          testconstants.SoftwareVersion,
		SoftwareVersionString:                    testconstants.SoftwareVersionString,
		HardwareVersion:                          testconstants.HardwareVersion,
		HardwareVersionString:                    testconstants.HardwareVersionString,
		CDVersionNumber:                          testconstants.CDVersionNumber,
		FirmwareDigests:                          testconstants.FirmwareDigests,
		Revoked:                                  testconstants.Revoked,
		OtaURL:                                   testconstants.OtaURL,
		OtaChecksum:                              testconstants.OtaChecksum,
		OtaChecksumType:                          testconstants.OtaChecksumType,
		OtaBlob:                                  testconstants.OtaBlob,
		CommissioningCustomFlow:                  testconstants.CommissioningCustomFlow,
		CommissioningCustomFlowUrl:               testconstants.CommissioningCustomFlowUrl,
		CommissioningModeInitialStepsHint:        testconstants.CommissioningModeInitialStepsHint,
		CommissioningModeInitialStepsInstruction: testconstants.CommissioningModeInitialStepsInstruction,
		CommissioningModeSecondaryStepsHint:      testconstants.CommissioningModeSecondaryStepsHint,
		CommissioningModeSecondaryStepsInstruction: testconstants.CommissioningModeSecondaryStepsInstruction,
		ReleaseNotesUrl: testconstants.ReleaseNotesUrl,
		UserManualUrl:   testconstants.UserManualUrl,
		SupportUrl:      testconstants.SupportUrl,
		ProductURL:      testconstants.ProductURL,
		ChipBlob:        testconstants.ChipBlob,
		VendorBlob:      testconstants.VendorBlob,
	}
}

func getTestModelForUpdate() Model {

	return Model{
		VID:                        testconstants.VID,
		PID:                        testconstants.PID,
		CID:                        testconstants.CID,
		Description:                testconstants.Description,
		CDVersionNumber:            testconstants.CDVersionNumber,
		FirmwareDigests:            testconstants.FirmwareDigests,
		Revoked:                    testconstants.Revoked,
		OtaURL:                     testconstants.OtaURL,
		OtaChecksum:                testconstants.OtaChecksum,
		OtaChecksumType:            testconstants.OtaChecksumType,
		OtaBlob:                    testconstants.OtaBlob,
		CommissioningCustomFlowUrl: testconstants.CommissioningCustomFlowUrl,
		ReleaseNotesUrl:            testconstants.ReleaseNotesUrl,
		UserManualUrl:              testconstants.UserManualUrl,
		SupportUrl:                 testconstants.SupportUrl,
		ProductURL:                 testconstants.ProductURL,
		ChipBlob:                   testconstants.ChipBlob,
		VendorBlob:                 testconstants.VendorBlob,
	}
}

func TestNewMsgAddModelInfo(t *testing.T) {
	model := getTestModel()
	msg := NewMsgAddModelInfo(model, testconstants.Signer)

	require.Equal(t, msg.Route(), RouterKey)
	require.Equal(t, msg.Type(), "add_model_info")
	require.Equal(t, msg.GetSigners(), []sdk.AccAddress{testconstants.Signer})
}

//nolint:funlen
func TestMsgAddModelInfoValidation(t *testing.T) {

	require.Nil(t, NewMsgAddModelInfo(getTestModel(), testconstants.Signer).ValidateBasic())

	// Invalid VID
	model := getTestModel()
	model.VID = 0
	require.NotNil(t, NewMsgAddModelInfo(model, testconstants.Signer).ValidateBasic())

	// Invalid PID
	model = getTestModel()
	model.PID = 0
	require.NotNil(t, NewMsgAddModelInfo(model, testconstants.Signer).ValidateBasic())

	// CID is optional
	model = getTestModel()
	model.CID = 0
	require.Nil(t, NewMsgAddModelInfo(model, testconstants.Signer).ValidateBasic())

	// Name is mandatory
	model = getTestModel()
	model.Name = ""
	require.NotNil(t, NewMsgAddModelInfo(model, testconstants.Signer).ValidateBasic())

	// Description is mandatory
	model = getTestModel()
	model.Description = ""
	require.NotNil(t, NewMsgAddModelInfo(model, testconstants.Signer).ValidateBasic())

	// SKU is mandatory
	model = getTestModel()
	model.SKU = ""
	require.NotNil(t, NewMsgAddModelInfo(model, testconstants.Signer).ValidateBasic())

	// SoftwareVersion is mandatory
	model = getTestModel()
	model.SoftwareVersion = 0
	require.NotNil(t, NewMsgAddModelInfo(model, testconstants.Signer).ValidateBasic())

	// SoftwareVersionString is mandatory
	model = getTestModel()
	model.SoftwareVersionString = ""
	require.NotNil(t, NewMsgAddModelInfo(model, testconstants.Signer).ValidateBasic())

	// HardwareVersion is mandatory
	model = getTestModel()
	model.HardwareVersion = 0
	require.NotNil(t, NewMsgAddModelInfo(model, testconstants.Signer).ValidateBasic())

	// HardwareVersionString is mandatory
	model = getTestModel()
	model.HardwareVersionString = ""
	require.NotNil(t, NewMsgAddModelInfo(model, testconstants.Signer).ValidateBasic())

	// Ota Combination checks - Missing checksum
	model = getTestModel()
	model.OtaChecksum = ""
	require.NotNil(t, NewMsgAddModelInfo(model, testconstants.Signer).ValidateBasic())

	// Ota Combination checks - Missing url and checksumType
	model = getTestModel()
	model.OtaURL = ""
	model.OtaChecksum = ""
	require.NotNil(t, NewMsgAddModelInfo(model, testconstants.Signer).ValidateBasic())

	// Ota Combination checks - Missing checksumType
	model = getTestModel()
	model.OtaChecksumType = ""
	require.NotNil(t, NewMsgAddModelInfo(model, testconstants.Signer).ValidateBasic())

	// Ota Combination checks - Missing all three
	model = getTestModel()
	model.OtaChecksum = ""
	model.OtaChecksumType = ""
	model.OtaURL = ""
	require.Nil(t, NewMsgAddModelInfo(model, testconstants.Signer).ValidateBasic())

	// Missing non mandatory
	model = getTestModel()
	model.CommissioningModeInitialStepsHint = 0
	model.CommissioningCustomFlow = 0
	model.CommissioningModeInitialStepsHint = 0
	model.CommissioningModeInitialStepsInstruction = ""
	require.Nil(t, NewMsgAddModelInfo(model, testconstants.Signer).ValidateBasic())

	// Signer is nil
	model = getTestModel()
	require.NotNil(t, NewMsgAddModelInfo(model, nil).ValidateBasic())
	// Signer is nil
	require.NotNil(t, NewMsgAddModelInfo(model, []byte{}).ValidateBasic())

}

func TestMsgAddModelInfoGetSignBytes(t *testing.T) {
	msg := NewMsgAddModelInfo(getTestModel(), testconstants.Signer)

	expected := "{\"type\":\"modelinfo/AddModelInfo\",\"value\":{\"Model\":{\"cd_version_number\":312,\"chip_blob\":\"Chip Blob Text\",\"cid\":12345,\"commisioning_mode_initial_steps_hint\":2,\"commisioning_mode_initial_steps_instruction\":\"commissioningModeInitialStepsInstruction details\",\"commisioning_mode_secondary_steps_hint\":3,\"commisioning_mode_secondary_steps_instruction\":\"commissioningModeSecondaryStepsInstruction steps\",\"commission_custom_flow\":1,\"commission_custom_flow_url\":\"https://sampleflowurl.dclmodel\",\"description\":\"Device Description\",\"firmware_digests\":\"Firmware Digest String\",\"hardware_version\":21,\"hardware_version_string\":\"2.1\",\"name\":\"Device Name\",\"ota_blob\":\"OTABlob Text\",\"ota_checksum\":\"0xe3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855\",\"ota_checksum_type\":\"SHA-256\",\"ota_url\":\"http://ota.firmware.com\",\"pid\":22,\"product_url\":\"https://url.producturl.dclmodel\",\"release_notes_url\":\"https://url.releasenotes.dclmodel\",\"revoked\":false,\"sku\":\"RCU2205A\",\"software_version\":1,\"software_version_string\":\"1.0\",\"support_url\":\"https://url.supporturl.dclmodel\",\"user_manual_url\":\"https://url.usermanual.dclmodel\",\"vendor_blob\":\"Vendor Blob Text\",\"vid\":1},\"signer\":\"cosmos1p72j8mgkf39qjzcmr283w8l8y9qv30qpj056uz\"}}"

	require.Equal(t, expected, string(msg.GetSignBytes()))

}

func TestNewMsgUpdateModelInfo(t *testing.T) {
	msg := NewMsgUpdateModelInfo(getTestModelForUpdate(), testconstants.Signer)

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
			getTestModelForUpdate(), testconstants.Signer)},
		{false, NewMsgUpdateModelInfo(
			getTestModelForUpdate(), nil)},
		{false, NewMsgUpdateModelInfo(
			getTestModelForUpdate(), []byte{})},
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
	msg := NewMsgUpdateModelInfo(getTestModelForUpdate(), testconstants.Signer)

	expected := `{"type":"modelinfo/UpdateModelInfo","value":{"Model":{"cd_version_number":312,"chip_blob":"Chip Blob Text","cid":12345,"commission_custom_flow_url":"https://sampleflowurl.dclmodel","description":"Device Description","firmware_digests":"Firmware Digest String","hardware_version":0,"hardware_version_string":"","name":"","ota_blob":"OTABlob Text","ota_checksum":"0xe3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855","ota_checksum_type":"SHA-256","ota_url":"http://ota.firmware.com","pid":22,"product_url":"https://url.producturl.dclmodel","release_notes_url":"https://url.releasenotes.dclmodel","revoked":false,"sku":"","software_version":0,"software_version_string":"","support_url":"https://url.supporturl.dclmodel","user_manual_url":"https://url.usermanual.dclmodel","vendor_blob":"Vendor Blob Text","vid":1},"signer":"cosmos1p72j8mgkf39qjzcmr283w8l8y9qv30qpj056uz"}}`

	require.Equal(t, expected, string(msg.GetSignBytes()))
}
