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
		DeviceTypeID:                             testconstants.DeviceTypeID,
		ProductName:                              testconstants.ProductName,
		ProductLabel:                             testconstants.ProductLabel,
		PartNumber:                               testconstants.PartNumber,
		CommissioningCustomFlow:                  testconstants.CommissioningCustomFlow,
		CommissioningCustomFlowURL:               testconstants.CommissioningCustomFlowURL,
		CommissioningModeInitialStepsHint:        testconstants.CommissioningModeInitialStepsHint,
		CommissioningModeInitialStepsInstruction: testconstants.CommissioningModeInitialStepsInstruction,
		CommissioningModeSecondaryStepsHint:      testconstants.CommissioningModeSecondaryStepsHint,
		CommissioningModeSecondaryStepsInstruction: testconstants.CommissioningModeSecondaryStepsInstruction,
		UserManualURL: testconstants.UserManualURL,
		SupportURL:    testconstants.SupportURL,
		ProductURL:    testconstants.ProductURL,
	}
}

func getTestModelForUpdate() Model {
	return Model{
		VID:                        testconstants.VID,
		PID:                        testconstants.PID,
		DeviceTypeID:               testconstants.DeviceTypeID,
		ProductLabel:               testconstants.ProductLabel,
		CommissioningCustomFlowURL: testconstants.CommissioningCustomFlowURL,
		UserManualURL:              testconstants.UserManualURL,
		SupportURL:                 testconstants.SupportURL,
		ProductURL:                 testconstants.ProductURL,
	}
}

func TestNewMsgAddModel(t *testing.T) {
	model := getTestModel()
	msg := NewMsgAddModel(model, testconstants.Signer)

	require.Equal(t, msg.Route(), RouterKey)
	require.Equal(t, msg.Type(), "add_model_info")
	require.Equal(t, msg.GetSigners(), []sdk.AccAddress{testconstants.Signer})
}

//nolint:funlen
func TestMsgAddModelValidation(t *testing.T) {
	require.Nil(t, NewMsgAddModel(getTestModel(), testconstants.Signer).ValidateBasic())

	// Invalid VID
	model := getTestModel()
	model.VID = 0
	require.NotNil(t, NewMsgAddModel(model, testconstants.Signer).ValidateBasic())

	// Invalid PID
	model = getTestModel()
	model.PID = 0
	require.NotNil(t, NewMsgAddModel(model, testconstants.Signer).ValidateBasic())

	// DeviceTypeID is Mandator
	model = getTestModel()
	model.DeviceTypeID = 0
	require.Nil(t, NewMsgAddModel(model, testconstants.Signer).ValidateBasic())

	// Name is mandatory
	model = getTestModel()
	model.ProductName = ""
	require.NotNil(t, NewMsgAddModel(model, testconstants.Signer).ValidateBasic())

	// ProductLabel is mandatory
	model = getTestModel()
	model.ProductLabel = ""
	require.NotNil(t, NewMsgAddModel(model, testconstants.Signer).ValidateBasic())

	// PartNumber is mandatory
	model = getTestModel()
	model.PartNumber = ""
	require.NotNil(t, NewMsgAddModel(model, testconstants.Signer).ValidateBasic())

	// Missing non mandatory
	model = getTestModel()
	model.CommissioningModeInitialStepsHint = 0
	model.CommissioningCustomFlow = 0
	model.CommissioningModeInitialStepsHint = 0
	model.CommissioningModeInitialStepsInstruction = ""
	require.Nil(t, NewMsgAddModel(model, testconstants.Signer).ValidateBasic())

	// Signer is nil
	model = getTestModel()
	require.NotNil(t, NewMsgAddModel(model, nil).ValidateBasic())
	// Signer is nil
	require.NotNil(t, NewMsgAddModel(model, []byte{}).ValidateBasic())
}

func TestMsgAddModelGetSignBytes(t *testing.T) {
	msg := NewMsgAddModel(getTestModel(), testconstants.Signer)
	expected := `{"type":"model/AddModel","value":{"Model":{"commissioningCustomFlow":1,"commissioningCustomFlowURL":"https://sampleflowurl.dclmodel","commissioningModeInitialStepsHint":2,"commissioningModeInitialStepsInstruction":"commissioningModeInitialStepsInstruction details","commissioningModeSecondaryStepsHint":3,"commissioningModeSecondaryStepsInstruction":"commissioningModeSecondaryStepsInstruction steps","deviceTypeID":12345,"partNumber":"RCU2205A","pid":22,"productLabel":"Product Label and/or Product Description","productName":"Device Name","productURL":"https://url.producturl.dclmodel","supportURL":"https://url.supporturl.dclmodel","userManualURL":"https://url.usermanual.dclmodel","vid":1},"signer":"cosmos1p72j8mgkf39qjzcmr283w8l8y9qv30qpj056uz"}}` //nolint:lll
	require.Equal(t, expected, string(msg.GetSignBytes()))
}

func TestNewMsgUpdateModel(t *testing.T) {
	msg := NewMsgUpdateModel(getTestModelForUpdate(), testconstants.Signer)
	require.Equal(t, msg.Route(), RouterKey)
	require.Equal(t, msg.Type(), "update_model_info")
	require.Equal(t, msg.GetSigners(), []sdk.AccAddress{testconstants.Signer})
}

func TestMsgUpdateModelValidation(t *testing.T) {
	cases := []struct {
		valid bool
		msg   MsgUpdateModel
	}{

		{true, NewMsgUpdateModel(
			getTestModelForUpdate(), testconstants.Signer)},
		{false, NewMsgUpdateModel(
			getTestModelForUpdate(), nil)},
		{false, NewMsgUpdateModel(
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

func TestMsgUpdateModelGetSignBytes(t *testing.T) {
	msg := NewMsgUpdateModel(getTestModelForUpdate(), testconstants.Signer)

	expected := `{"type":"model/UpdateModel","value":{"Model":{"commissioningCustomFlow":0,"commissioningCustomFlowURL":"https://sampleflowurl.dclmodel","deviceTypeID":12345,"pid":22,"productLabel":"Product Label and/or Product Description","productURL":"https://url.producturl.dclmodel","supportURL":"https://url.supporturl.dclmodel","userManualURL":"https://url.usermanual.dclmodel","vid":1},"signer":"cosmos1p72j8mgkf39qjzcmr283w8l8y9qv30qpj056uz"}}` //nolint:lll

	require.Equal(t, expected, string(msg.GetSignBytes()))
}
