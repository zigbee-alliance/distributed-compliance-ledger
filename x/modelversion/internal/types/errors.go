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
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	Codespace sdk.CodespaceType = ModuleName

	CodeSoftwareVersionStringInvalid sdk.CodeType = 511
	CodeFirmwareDigestsInvalid       sdk.CodeType = 512
	CodeCDVersionNumberInvalid       sdk.CodeType = 513
	CodeErrOtaURLInvalid             sdk.CodeType = 514
	CodeErrOtaMissingInformation     sdk.CodeType = 515
	CodeReleaseNotesUrlInvalid       sdk.CodeType = 516
	CodeModelVersionDoesNotExist     sdk.CodeType = 517
	CodeNoModelVersionExist          sdk.CodeType = 518
	CodeModelVersionAlreadyExists    sdk.CodeType = 519
	CodeOtaURLCannotBeSet            sdk.CodeType = 520
	CodeModelDoesNotExist            sdk.CodeType = 521
)

func ErrSoftwareVersionStringInvalid(softwareVersion interface{}) sdk.Error {
	return sdk.NewError(Codespace, CodeSoftwareVersionStringInvalid,
		fmt.Sprintf("SoftwareVersionString %v is invalid. It should be greater then 1 and less then 64 character long", softwareVersion))
}

func ErrFirmwareDigestsInvalid(firmwareDigests interface{}) sdk.Error {
	return sdk.NewError(Codespace, CodeSoftwareVersionStringInvalid,
		fmt.Sprintf("firmwareDigests %v is of invalid length. Maximum length should be less then 512", firmwareDigests))
}

func ErrCDVersionNumberInvalid(cdVersionNumber interface{}) sdk.Error {
	return sdk.NewError(Codespace, CodeSoftwareVersionStringInvalid,
		fmt.Sprintf("CDVersionNumber %v is invalid. It should be a 16 bit unsigned integer", cdVersionNumber))
}

func ErrOtaURLInvalid(otaURL interface{}) sdk.Error {
	return sdk.NewError(Codespace, CodeSoftwareVersionStringInvalid,
		fmt.Sprintf("OtaURL %v is invalid. Maximum length should be less then 256", otaURL))
}

func ErrMissingOtaInformation() sdk.Error {
	return sdk.NewError(Codespace, CodeErrOtaMissingInformation, "OtaFileSize, OtaChecksum and OtaChecksumType are required if OtaUrl is provided")
}

func ErrReleaseNotesURLInvalid(releaseNotesURL interface{}) sdk.Error {
	return sdk.NewError(Codespace, CodeReleaseNotesUrlInvalid,
		fmt.Sprintf("ReleaseNotesURLInvalid %v is invalid. Maximum length should be less then 256", releaseNotesURL))
}

func ErrModelVersionDoesNotExist(vid interface{}, pid interface{}, softwareVersion interface{}) sdk.Error {
	return sdk.NewError(Codespace, CodeModelVersionDoesNotExist,
		fmt.Sprintf("No model version associated with vid=%v, pid=%v and softwareVersion=%v exist on the ledger", vid, pid, softwareVersion))
}

func ErrNoModelVersionsExist(vid interface{}, pid interface{}) sdk.Error {
	return sdk.NewError(Codespace, CodeNoModelVersionExist,
		fmt.Sprintf("No versions associated with vid=%v and pid=%v exist on the ledger", vid, pid))
}

func ErrModelVersionAlreadyExists(vid interface{}, pid interface{}, softwareVersion interface{}) sdk.Error {
	return sdk.NewError(Codespace, CodeModelVersionAlreadyExists,
		fmt.Sprintf("Model Version already exists on ledger with vid=%v pid=%v and softwareVersion=%v exist on the ledger", vid, pid, softwareVersion))
}

func ErrOtaURLCannotBeSet(vid interface{}, pid interface{}, softwareVersion interface{}) sdk.Error {
	return sdk.NewError(Codespace, CodeOtaURLCannotBeSet,
		fmt.Sprintf("OTA URL cannot be set for model version associated with vid=%v, pid=%v "+
			"and softwareVersion=%v because OTA was not set for this model info initially", vid, pid, softwareVersion))
}

func ErrModelDoesNotExist(vid interface{}, pid interface{}) sdk.Error {
	return sdk.NewError(Codespace, CodeModelDoesNotExist,
		fmt.Sprintf("No model associated with vid=%v and pid=%v exist on the ledger", vid, pid))
}
