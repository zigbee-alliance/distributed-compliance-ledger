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
	"encoding/json"
)

//nolint:maligned
type ModelVersion struct {
	VID                          uint16 `json:"vid" validate:"required"`
	PID                          uint16 `json:"pid" validate:"required"`
	SoftwareVersion              uint32 `json:"softwareVersion" validate:"required"`
	SoftwareVersionString        string `json:"softwareVersionString,omitempty" validate:"requiredForAdd,max=64"`
	CDVersionNumber              uint16 `json:"CDVersionNumber,omitempty" validate:"requiredForAdd"`
	FirmwareDigests              string `json:"firmwareDigests,omitempty" validate:"max=512"`
	SoftwareVersionValid         bool   `json:"softwareVersionValid"`
	OtaURL                       string `json:"otaURL,omitempty" validate:"omitempty,url,max=256"`
	OtaFileSize                  uint64 `json:"otaFileSize,omitempty" validate:"required_with_all=OtaURL"`
	OtaChecksum                  string `json:"otaChecksum,omitempty" validate:"required_with_all=OtaURL,max=64"`
	OtaChecksumType              uint16 `json:"otaChecksumType,omitempty" validate:"required_with_all=OtaURL"`
	MinApplicableSoftwareVersion uint32 `json:"minApplicableSoftwareVersion,omitempty" validate:"requiredForAdd"`
	MaxApplicableSoftwareVersion uint32 `json:"maxApplicableSoftwareVersion,omitempty" validate:"requiredForAdd"`
	ReleaseNotesURL              string `json:"releaseNotesURL,omitempty" validate:"max=256"`
}

func (d ModelVersion) String() string {
	bytes, err := json.Marshal(d)
	if err != nil {
		panic(err)
	}

	return string(bytes)
}

type ModelVersions struct {
	VID              uint16   `json:"vid"`
	PID              uint16   `json:"pid"`
	SoftwareVersions []uint32 `json:"softwareVersions"`
}
