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
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type ComplianceState string

const (
	DevTest     ComplianceState = "dev-test"
	Provisional ComplianceState = "provisional"
	Certified   ComplianceState = "certified"
	Revoked     ComplianceState = "revoked"
)

type SoftwareVersionCertificationStatus uint8

const (
	CodeDevTest     SoftwareVersionCertificationStatus = 0
	CodeProvisional SoftwareVersionCertificationStatus = 1
	CodeCertified   SoftwareVersionCertificationStatus = 2
	CodeRevoked     SoftwareVersionCertificationStatus = 3
)

type CertificationType string

const (
	ZbCertificationType  CertificationType = "zb"
	CSACertificationType CertificationType = "csa"
)

/*
	Compliance info stored into KVStore
*/
type ComplianceInfo struct {
	VID                                uint16                             `json:"vid"`
	PID                                uint16                             `json:"pid"`
	SoftwareVersion                    uint32                             `json:"softwareVersion"`
	SoftwareVersionString              string                             `json:"softwareVersionString,omitempty"`
	CDVersionNumber                    uint32                             `json:"CDVersionNumber,omitempty"`
	SoftwareVersionCertificationStatus SoftwareVersionCertificationStatus `json:"softwareVersionCertificationStatus"`
	Date                               time.Time                          `json:"date"` // rfc3339 encoded date
	CertificationType                  CertificationType                  `json:"certification_type"`
	Reason                             string                             `json:"reason,omitempty"`
	Owner                              sdk.AccAddress                     `json:"owner"`
	History                            []ComplianceHistoryItem            `json:"history,omitempty"`
}

func NewCertifiedComplianceInfo(vid uint16, pid uint16,
	softwareVersion uint32, softwareVersionString string,
	certificationType CertificationType,
	date time.Time, reason string, owner sdk.AccAddress) ComplianceInfo {
	return ComplianceInfo{
		VID:                                vid,
		PID:                                pid,
		SoftwareVersion:                    softwareVersion,
		SoftwareVersionString:              softwareVersionString,
		SoftwareVersionCertificationStatus: CodeCertified,
		Date:                               date,
		CertificationType:                  certificationType,
		Reason:                             reason,
		Owner:                              owner,
		History:                            []ComplianceHistoryItem{},
	}
}

func NewRevokedComplianceInfo(vid uint16, pid uint16,
	softwareVersion uint32, softwareVersionString string,
	certificationType CertificationType,
	date time.Time, reason string, owner sdk.AccAddress) ComplianceInfo {
	return ComplianceInfo{
		VID:                                vid,
		PID:                                pid,
		SoftwareVersion:                    softwareVersion,
		SoftwareVersionString:              softwareVersionString,
		SoftwareVersionCertificationStatus: CodeRevoked,
		Date:                               date,
		CertificationType:                  certificationType,
		Reason:                             reason,
		Owner:                              owner,
		History:                            []ComplianceHistoryItem{},
	}
}

func (d *ComplianceInfo) UpdateComplianceInfo(date time.Time, reason string) {
	// Toggle state
	var svCertificationStatus SoftwareVersionCertificationStatus
	if d.SoftwareVersionCertificationStatus == CodeCertified {
		svCertificationStatus = CodeRevoked
	} else {
		svCertificationStatus = CodeCertified
	}

	d.History = append(d.History, NewComplianceHistoryItem(d.SoftwareVersionCertificationStatus, d.Date, d.Reason))
	d.SoftwareVersionCertificationStatus = svCertificationStatus
	d.Date = date
	d.Reason = reason
}

func (d ComplianceInfo) String() string {
	bytes, err := json.Marshal(d)
	if err != nil {
		panic(err)
	}

	return string(bytes)
}

/*
	Compliance info state changes
*/
type ComplianceHistoryItem struct {
	SoftwareVersionCertificationStatus SoftwareVersionCertificationStatus `json:"softwareVersionCertificationStatus"`
	Date                               time.Time                          `json:"date"` // rfc3339 encoded date
	Reason                             string                             `json:"reason,omitempty"`
}

func NewComplianceHistoryItem(svCertificationStatus SoftwareVersionCertificationStatus, date time.Time, reason string) ComplianceHistoryItem {
	return ComplianceHistoryItem{
		SoftwareVersionCertificationStatus: svCertificationStatus,
		Date:                               date,
		Reason:                             reason,
	}
}

func (d ComplianceHistoryItem) String() string {
	bytes, err := json.Marshal(d)
	if err != nil {
		panic(err)
	}

	return string(bytes)
}

type ComplianceInfoKey struct {
	VID               uint16            `json:"vid"`
	PID               uint16            `json:"pid"`
	SoftwareVersion   uint32            `json:"softwareVersion"`
	CertificationType CertificationType `json:"certification_type"`
}
