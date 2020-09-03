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
	Certified ComplianceState = "certified"
	Revoked   ComplianceState = "revoked"
)

type CertificationType string

const (
	ZbCertificationType CertificationType = "zb"
)

/*
	Compliance info stored into KVStore
*/
type ComplianceInfo struct {
	VID               uint16                  `json:"vid"`
	PID               uint16                  `json:"pid"`
	State             ComplianceState         `json:"state"`
	Date              time.Time               `json:"date"` // rfc3339 encoded date
	CertificationType CertificationType       `json:"certification_type"`
	Reason            string                  `json:"reason,omitempty"`
	Owner             sdk.AccAddress          `json:"owner"`
	History           []ComplianceHistoryItem `json:"history,omitempty"`
}

func NewCertifiedComplianceInfo(vid uint16, pid uint16, certificationType CertificationType,
	date time.Time, reason string, owner sdk.AccAddress) ComplianceInfo {
	return ComplianceInfo{
		VID:               vid,
		PID:               pid,
		State:             Certified,
		Date:              date,
		CertificationType: certificationType,
		Reason:            reason,
		Owner:             owner,
		History:           []ComplianceHistoryItem{},
	}
}

func NewRevokedComplianceInfo(vid uint16, pid uint16, certificationType CertificationType,
	date time.Time, reason string, owner sdk.AccAddress) ComplianceInfo {
	return ComplianceInfo{
		VID:               vid,
		PID:               pid,
		State:             Revoked,
		Date:              date,
		CertificationType: certificationType,
		Reason:            reason,
		Owner:             owner,
		History:           []ComplianceHistoryItem{},
	}
}

func (d *ComplianceInfo) UpdateComplianceInfo(date time.Time, reason string) {
	// Toggle state
	var state ComplianceState
	if d.State == Certified {
		state = Revoked
	} else {
		state = Certified
	}

	d.History = append(d.History, NewComplianceHistoryItem(d.State, d.Date, d.Reason))
	d.State = state
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
	State  ComplianceState `json:"state"`
	Date   time.Time       `json:"date"` // rfc3339 encoded date
	Reason string          `json:"reason,omitempty"`
}

func NewComplianceHistoryItem(state ComplianceState, date time.Time, reason string) ComplianceHistoryItem {
	return ComplianceHistoryItem{
		State:  state,
		Date:   date,
		Reason: reason,
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
	CertificationType CertificationType `json:"certification_type"`
}
