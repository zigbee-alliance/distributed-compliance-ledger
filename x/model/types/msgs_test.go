// Copyright 2022 DSR Corporation
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

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
)

func TestValidateMsgCreateModel(t *testing.T) {
	cases := []struct {
		valid bool
		msg   *MsgCreateModel
	}{

		{true, newMsgCreateModel(testconstants.Signer, 1, 1)},
		{true, newMsgCreateModel(testconstants.Signer, 65535, 65535)},

		// zero PID/VID - not OK
		{false, newMsgCreateModel(testconstants.Signer, 0, 1)},
		{false, newMsgCreateModel(testconstants.Signer, 1, 0)},

		// negative VID/PID - not OK
		{false, newMsgCreateModel(testconstants.Signer, -1, 1)},
		{false, newMsgCreateModel(testconstants.Signer, 1, -1)},

		// too large VID/PID - not OK
		{false, newMsgCreateModel(testconstants.Signer, 65535+1, 1)},
		{false, newMsgCreateModel(testconstants.Signer, 1, 65535+1)},
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

func TestValidateMsgUpdateModel(t *testing.T) {
	cases := []struct {
		valid bool
		msg   *MsgUpdateModel
	}{

		{true, newMsgUpdateModel(testconstants.Signer, 1, 1)},
		{true, newMsgUpdateModel(testconstants.Signer, 65535, 65535)},

		// zero PID/VID - not OK
		{false, newMsgUpdateModel(testconstants.Signer, 0, 1)},
		{false, newMsgUpdateModel(testconstants.Signer, 1, 0)},

		// negative VID/PID - not OK
		{false, newMsgUpdateModel(testconstants.Signer, -1, 1)},
		{false, newMsgUpdateModel(testconstants.Signer, 1, -1)},

		// too large VID/PID - not OK
		{false, newMsgUpdateModel(testconstants.Signer, 65535+1, 1)},
		{false, newMsgUpdateModel(testconstants.Signer, 1, 65535+1)},
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

func TestValidateMsgCreateModelVersion(t *testing.T) {
	cases := []struct {
		valid bool
		msg   *MsgCreateModelVersion
	}{

		{true, newMsgCreateModelVersion(testconstants.Signer, 1, 1, 1)},
		{true, newMsgCreateModelVersion(testconstants.Signer, 65535, 65535, 1)},

		// zero SV - OK
		{true, newMsgCreateModelVersion(testconstants.Signer, 1, 1, 0)},

		// zero PID/VID - not OK
		{false, newMsgCreateModelVersion(testconstants.Signer, 0, 1, 1)},
		{false, newMsgCreateModelVersion(testconstants.Signer, 1, 0, 1)},

		// negative VID/PID - not OK
		{false, newMsgCreateModelVersion(testconstants.Signer, -1, 1, 1)},
		{false, newMsgCreateModelVersion(testconstants.Signer, 1, -1, 1)},

		// too large VID/PID - not OK
		{false, newMsgCreateModelVersion(testconstants.Signer, 65535+1, 1, 1)},
		{false, newMsgCreateModelVersion(testconstants.Signer, 1, 65535+1, 1)},
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

func TestValidateMsgUpdateModelVersion(t *testing.T) {
	cases := []struct {
		valid bool
		msg   *MsgUpdateModelVersion
	}{

		{true, newMsgUpdateModelVersion(testconstants.Signer, 1, 1, 1)},
		{true, newMsgUpdateModelVersion(testconstants.Signer, 65535, 65535, 1)},

		// zero SV - OK
		{true, newMsgUpdateModelVersion(testconstants.Signer, 1, 1, 0)},

		// zero PID/VID - not OK
		{false, newMsgUpdateModelVersion(testconstants.Signer, 0, 1, 1)},
		{false, newMsgUpdateModelVersion(testconstants.Signer, 1, 0, 1)},

		// negative VID/PID - not OK
		{false, newMsgUpdateModelVersion(testconstants.Signer, -1, 1, 1)},
		{false, newMsgUpdateModelVersion(testconstants.Signer, 1, -1, 1)},

		// too large VID/PID - not OK
		{false, newMsgUpdateModelVersion(testconstants.Signer, 65535+1, 1, 1)},
		{false, newMsgUpdateModelVersion(testconstants.Signer, 1, 65535+1, 1)},
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

func newMsgCreateModel(signer sdk.AccAddress, vid int32, pid int32) *MsgCreateModel {
	return &MsgCreateModel{
		Creator:                                  signer.String(),
		Vid:                                      vid,
		Pid:                                      pid,
		DeviceTypeId:                             testconstants.DeviceTypeId,
		ProductName:                              testconstants.ProductName,
		ProductLabel:                             testconstants.ProductLabel,
		PartNumber:                               testconstants.PartNumber,
		CommissioningCustomFlow:                  testconstants.CommissioningCustomFlow,
		CommissioningCustomFlowUrl:               testconstants.CommissioningCustomFlowUrl,
		CommissioningModeInitialStepsHint:        testconstants.CommissioningModeInitialStepsHint,
		CommissioningModeInitialStepsInstruction: testconstants.CommissioningModeInitialStepsInstruction,
		CommissioningModeSecondaryStepsHint:      testconstants.CommissioningModeSecondaryStepsHint,
		CommissioningModeSecondaryStepsInstruction: testconstants.CommissioningModeSecondaryStepsInstruction,
		UserManualUrl: testconstants.UserManualUrl,
		SupportUrl:    testconstants.SupportUrl,
		ProductUrl:    testconstants.ProductUrl,
	}
}

func newMsgUpdateModel(signer sdk.AccAddress, vid int32, pid int32) *MsgUpdateModel {
	return &MsgUpdateModel{
		Creator:                                  signer.String(),
		Vid:                                      vid,
		Pid:                                      pid,
		ProductName:                              testconstants.ProductName + "-updated",
		ProductLabel:                             testconstants.ProductLabel + "-updated",
		PartNumber:                               testconstants.PartNumber + "-updated",
		CommissioningCustomFlowUrl:               testconstants.CommissioningCustomFlowUrl + "/updated",
		CommissioningModeInitialStepsInstruction: testconstants.CommissioningModeInitialStepsInstruction + "-updated",
		CommissioningModeSecondaryStepsInstruction: testconstants.CommissioningModeSecondaryStepsInstruction + "-updated",
		UserManualUrl: testconstants.UserManualUrl + "/updated",
		SupportUrl:    testconstants.SupportUrl + "/updated",
		ProductUrl:    testconstants.ProductUrl + "/updated",
	}
}

func newMsgCreateModelVersion(signer sdk.AccAddress, vid int32, pid int32, sv uint32) *MsgCreateModelVersion {
	return &MsgCreateModelVersion{
		Creator:                      signer.String(),
		Vid:                          vid,
		Pid:                          pid,
		SoftwareVersion:              sv,
		SoftwareVersionString:        testconstants.SoftwareVersionString,
		CdVersionNumber:              testconstants.CdVersionNumber,
		FirmwareDigests:              testconstants.FirmwareDigests,
		SoftwareVersionValid:         testconstants.SoftwareVersionValid,
		OtaUrl:                       testconstants.OtaUrl,
		OtaFileSize:                  testconstants.OtaFileSize,
		OtaChecksum:                  testconstants.OtaChecksum,
		OtaChecksumType:              testconstants.OtaChecksumType,
		MinApplicableSoftwareVersion: testconstants.MinApplicableSoftwareVersion,
		MaxApplicableSoftwareVersion: testconstants.MaxApplicableSoftwareVersion,
		ReleaseNotesUrl:              testconstants.ReleaseNotesUrl,
	}
}

func newMsgUpdateModelVersion(signer sdk.AccAddress, vid int32, pid int32, sv uint32) *MsgUpdateModelVersion {
	return &MsgUpdateModelVersion{
		Creator:                      signer.String(),
		Vid:                          vid,
		Pid:                          pid,
		SoftwareVersion:              sv,
		SoftwareVersionValid:         !testconstants.SoftwareVersionValid,
		OtaUrl:                       testconstants.OtaUrl + "/updated",
		MinApplicableSoftwareVersion: testconstants.MinApplicableSoftwareVersion + 1,
		MaxApplicableSoftwareVersion: testconstants.MaxApplicableSoftwareVersion + 1,
		ReleaseNotesUrl:              testconstants.ReleaseNotesUrl + "/updated",
	}
}
