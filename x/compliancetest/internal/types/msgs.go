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
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const RouterKey = ModuleName

type MsgAddTestingResult struct {
	VID                   uint16         `json:"vid"`
	PID                   uint16         `json:"pid"`
	SoftwareVersion       uint32         `json:"softwareVersion"`
	SoftwareVersionString string         `json:"softwareVersionString"`
	TestResult            string         `json:"test_result"`
	TestDate              time.Time      `json:"test_date"` // rfc3339 encoded date
	Signer                sdk.AccAddress `json:"signer"`
}

func NewMsgAddTestingResult(vid uint16, pid uint16, softwareVersion uint32,
	softwareVersionString string, testResult string,
	testDate time.Time, signer sdk.AccAddress) MsgAddTestingResult {
	return MsgAddTestingResult{
		VID:                   vid,
		PID:                   pid,
		SoftwareVersion:       softwareVersion,
		SoftwareVersionString: softwareVersionString,
		TestResult:            testResult,
		TestDate:              testDate,
		Signer:                signer,
	}
}

func (m MsgAddTestingResult) Route() string {
	return RouterKey
}

func (m MsgAddTestingResult) Type() string {
	return "add_testing_result"
}

func (m MsgAddTestingResult) ValidateBasic() sdk.Error {
	if m.Signer.Empty() {
		return sdk.ErrInvalidAddress("Invalid Signer: it cannot be empty")
	}

	if m.VID == 0 {
		return sdk.ErrUnknownRequest("Invalid VID: it must be non zero 16-bit unsigned integer")
	}

	if m.PID == 0 {
		return sdk.ErrUnknownRequest("Invalid PID: it must be non zero 16-bit unsigned integer")
	}

	if m.SoftwareVersion == 0 {
		return sdk.ErrUnknownRequest("Invalid SoftwareVersion: it must be non zero 32-bit unsigned integer")
	}

	if len(m.SoftwareVersionString) == 0 {
		return sdk.ErrUnknownRequest("Invalid SoftwareVersionString: it cannot be empty")
	}

	if len(m.TestResult) == 0 {
		return sdk.ErrUnknownRequest("Invalid TestResult: it cannot be empty")
	}

	if m.TestDate.IsZero() {
		return sdk.ErrUnknownRequest("Invalid TestDate: it cannot be empty")
	}

	return nil
}

func (m MsgAddTestingResult) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(m))
}

func (m MsgAddTestingResult) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Signer}
}
