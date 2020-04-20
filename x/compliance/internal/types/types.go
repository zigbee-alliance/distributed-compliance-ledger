package types

import (
	"encoding/json"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"time"
)

type ComplianceState string

const (
	Certified ComplianceState = "certified"
	Revoked   ComplianceState = "revoked"
)

type CertificationType string

const (
	ZbCertificationType    CertificationType = "zb"
	EmptyCertificationType CertificationType = ""
)

type ComplianceInfo struct {
	VID               int16                   `json:"vid"`
	PID               int16                   `json:"pid"`
	State             ComplianceState         `json:"state"`
	Date              time.Time               `json:"date"` // rfc3339 encoded date
	CertificationType CertificationType       `json:"certification_type,omitempty"`
	Reason            string                  `json:"reason,omitempty"`
	Owner             sdk.AccAddress          `json:"owner"`
	History           []ComplianceHistoryItem `json:"history,omitempty"`
}

func NewCertifiedComplianceInfo(vid int16, pid int16, certificationType CertificationType, date time.Time, reason string, owner sdk.AccAddress) ComplianceInfo {
	if certificationType == EmptyCertificationType { // `zb` certification_type is only supported now
		certificationType = ZbCertificationType
	}

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

func NewRevokedComplianceInfo(vid int16, pid int16, certificationType CertificationType, date time.Time, reason string, owner sdk.AccAddress) ComplianceInfo {
	if certificationType == EmptyCertificationType { // `zb` certification_type is only supported now
		certificationType = ZbCertificationType
	}

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
