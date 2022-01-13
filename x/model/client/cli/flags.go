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

package cli

const (
	// Flags for Model.
	FlagVid                                        = "vid"
	FlagPid                                        = "pid"
	FlagDeviceTypeId                               = "deviceTypeID"
	FlagProductName                                = "productName"
	FlagProductNameShortcut                        = "n"
	FlagProductLabel                               = "productLabel"
	FlagProductLabelShortcut                       = "d"
	FlagPartNumber                                 = "partNumber"
	FlagCommissioningCustomFlow                    = "commissioningCustomFlow"
	FlagCommissioningCustomFlowUrl                 = "commissioningCustomFlowURL"
	FlagCommissioningModeInitialStepsHint          = "commissioningModeInitialStepsHint"
	FlagCommissioningModeInitialStepsInstruction   = "commissioningModeInitialStepsInstruction"
	FlagCommissioningModeSecondaryStepsHint        = "commissioningModeSecondaryStepsHint"
	FlagCommissioningModeSecondaryStepsInstruction = "commissioningModeSecondaryStepsInstruction"
	FlagUserManualUrl                              = "userManualURL"
	FlagSupportUrl                                 = "supportURL"
	FlagProductUrl                                 = "productURL"

	// Flags for ModelVersion.
	FlagSoftwareVersion              = "softwareVersion"
	FlagSoftwareVersionShortcut      = "v"
	FlagSoftwareVersionString        = "softwareVersionString"
	FlagCdVersionNumber              = "cdVersionNumber"
	FlagFirmwareDigests              = "firmwareDigests"
	FlagSoftwareVersionValid         = "softwareVersionValid"
	FlagOtaUrl                       = "otaURL"
	FlagOtaFileSize                  = "otaFileSize"
	FlagOtaChecksum                  = "otaChecksum"
	FlagOtaChecksumType              = "otaChecksumType"
	FlagMinApplicableSoftwareVersion = "minApplicableSoftwareVersion"
	FlagMaxApplicableSoftwareVersion = "maxApplicableSoftwareVersion"
	FlagReleaseNotesUrl              = "releaseNotesURL"
)
