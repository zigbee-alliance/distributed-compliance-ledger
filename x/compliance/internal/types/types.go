package types

import (
	"encoding/json"
	"strconv"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type ModelInfo struct {
	VID                      int16          `json:"vid"`
	PID                      int16          `json:"pid"`
	CID                      int16          `json:"cid,omitempty"`
	Name                     string         `json:"name"`
	Owner                    sdk.AccAddress `json:"owner"`
	Description              string         `json:"description"`
	SKU                      string         `json:"sku"`
	FirmwareVersion          string         `json:"firmware_version"`
	HardwareVersion          string         `json:"hardware_version"`
	Custom                   string         `json:"custom"`
	CertificateID            string         `json:"certificate_id,omitempty"`
	CertifiedDate            time.Time      `json:"certified_date,omitempty"` // rfc3339 data format
	TisOrTrpTestingCompleted bool           `json:"tis_or_trp_testing_completed"`
}

func NewModelInfo(vid int16, pid int16, cid int16, name string, owner sdk.AccAddress,
	description string, sku string, firmwareVersion string, hardwareVersion string, custom string, certificateID string,
	certifiedDate time.Time, tisOrTrpTestingCompleted bool) ModelInfo {
	return ModelInfo{
		VID:                      vid,
		PID:                      pid,
		CID:                      cid,
		Name:                     name,
		Owner:                    owner,
		Description:              description,
		SKU:                      sku,
		FirmwareVersion:          firmwareVersion,
		HardwareVersion:          hardwareVersion,
		Custom:                   custom,
		CertificateID:            certificateID,
		CertifiedDate:            certifiedDate,
		TisOrTrpTestingCompleted: tisOrTrpTestingCompleted,
	}
}

func (d ModelInfo) String() string {
	bytes, err := json.Marshal(d)
	if err != nil {
		panic(err)
	}

	return string(bytes)
}

func parseInt16FromString(str string) (int16, error) {
	val_, err := strconv.ParseInt(str, 10, 16)
	if err != nil {
		return 0, err
	}
	vid := int16(val_)
	return vid, nil
}

func ParseVID(str string) (int16, error) {
	return parseInt16FromString(str)
}

func ParsePID(str string) (int16, error) {
	return parseInt16FromString(str)
}

func ParseCID(str string) (int16, error) {
	return parseInt16FromString(str)
}
