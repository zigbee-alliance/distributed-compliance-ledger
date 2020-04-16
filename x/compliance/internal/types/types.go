package types

import (
	"encoding/json"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"time"
)

const ZbCertificationType = "zb"

type CertifiedModel struct {
	VID               int16          `json:"vid"`
	PID               int16          `json:"pid"`
	CertificationDate time.Time      `json:"certification_date"` // rfc3339 encoded date
	CertificationType string         `json:"certification_type"`
	Owner             sdk.AccAddress `json:"owner"`
}

func NewCertifiedModel(vid int16, pid int16, certificationDate time.Time, certificationType string, owner sdk.AccAddress) CertifiedModel {
	return CertifiedModel{
		VID:               vid,
		PID:               pid,
		CertificationDate: certificationDate,
		CertificationType: ZbCertificationType, // zb certification type is only supported now
		Owner:             owner,
	}
}

func (d CertifiedModel) String() string {
	bytes, err := json.Marshal(d)
	if err != nil {
		panic(err)
	}

	return string(bytes)
}
