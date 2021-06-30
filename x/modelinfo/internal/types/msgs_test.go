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
		ProductName:                              testconstants.ProductName,
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
		CommissioningCustomFlowURL:               testconstants.CommissioningCustomFlowURL,
		CommissioningModeInitialStepsHint:        testconstants.CommissioningModeInitialStepsHint,
		CommissioningModeInitialStepsInstruction: testconstants.CommissioningModeInitialStepsInstruction,
		CommissioningModeSecondaryStepsHint:      testconstants.CommissioningModeSecondaryStepsHint,
		CommissioningModeSecondaryStepsInstruction: testconstants.CommissioningModeSecondaryStepsInstruction,
		ReleaseNotesURL: testconstants.ReleaseNotesURL,
		UserManualURL:   testconstants.UserManualURL,
		SupportURL:      testconstants.SupportURL,
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
		CommissioningCustomFlowURL: testconstants.CommissioningCustomFlowURL,
		ReleaseNotesURL:            testconstants.ReleaseNotesURL,
		UserManualURL:              testconstants.UserManualURL,
		SupportURL:                 testconstants.SupportURL,
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
	model.ProductName = ""
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
	expected := `{"type":"modelinfo/AddModelInfo","value":{"Model":{"CDVersionNumber":312,"chipBlob":"Chip Blob Text","cid":12345,"commissioningCustomFlow":1,"commissioningCustomFlowURL":"https://sampleflowurl.dclmodel","commissioningModeInitialStepsHint":2,"commissioningModeInitialStepsInstruction":"commissioningModeInitialStepsInstruction details","commissioningModeSecondaryStepsHint":3,"commissioningModeSecondaryStepsInstruction":"commissioningModeSecondaryStepsInstruction steps","description":"Device Description","firmwareDigests":"Firmware Digest String","hardwareVersion":21,"hardwareVersionString":"2.1","otaBlob":"OTABlob Text","otaChecksum":"0xe3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855","otaChecksumType":"SHA-256","otaURL":"http://ota.firmware.com","pid":22,"productName":"Device Name","productURL":"https://url.producturl.dclmodel","releaseNotesURL":"https://url.releasenotes.dclmodel","revoked":false,"sku":"RCU2205A","softwareVersion":1,"softwareVersionString":"1.0","supportURL":"https://url.supporturl.dclmodel","userManualURL":"https://url.usermanual.dclmodel","vendorBlob":"Vendor Blob Text","vid":1},"signer":"cosmos1p72j8mgkf39qjzcmr283w8l8y9qv30qpj056uz"}}` //nolint:lll
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

	expected := `{"type":"modelinfo/UpdateModelInfo","value":{"Model":{"CDVersionNumber":312,"chipBlob":"Chip Blob Text","cid":12345,"commissioningCustomFlowURL":"https://sampleflowurl.dclmodel","description":"Device Description","firmwareDigests":"Firmware Digest String","hardwareVersion":0,"hardwareVersionString":"","otaBlob":"OTABlob Text","otaChecksum":"0xe3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855","otaChecksumType":"SHA-256","otaURL":"http://ota.firmware.com","pid":22,"productName":"","productURL":"https://url.producturl.dclmodel","releaseNotesURL":"https://url.releasenotes.dclmodel","revoked":false,"sku":"","softwareVersion":0,"softwareVersionString":"","supportURL":"https://url.supporturl.dclmodel","userManualURL":"https://url.usermanual.dclmodel","vendorBlob":"Vendor Blob Text","vid":1},"signer":"cosmos1p72j8mgkf39qjzcmr283w8l8y9qv30qpj056uz"}}` //nolint:lll

	require.Equal(t, expected, string(msg.GetSignBytes()))
}
