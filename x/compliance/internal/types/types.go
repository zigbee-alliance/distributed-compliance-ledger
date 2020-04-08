package types

import (
	"encoding/json"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type ModelInfo struct {
	ID                       string         `json:"id"`
	Name                     string         `json:"name"`
	Owner                    sdk.AccAddress `json:"owner"`
	Description              string         `json:"description"`
	SKU                      string         `json:"sku"`
	FirmwareVersion          string         `json:"firmware_version"`
	HardwareVersion          string         `json:"hardware_version"`
	CertificateID            string         `json:"certificate_id,omitempty"`
	CertifiedDate            time.Time      `json:"certified_date,omitempty"`
	TisOrTrpTestingCompleted bool           `json:"tis_or_trp_testing_completed"`
}

func NewModelInfo(id string, name string, owner sdk.AccAddress,
	description string, sku string, firmwareVersion string, hardwareVersion string, certificateID string,
	certifiedDate time.Time, tisOrTrpTestingCompleted bool) ModelInfo {
	return ModelInfo{
		ID:                       id,
		Name:                     name,
		Owner:                    owner,
		Description:              description,
		SKU:                      sku,
		FirmwareVersion:          firmwareVersion,
		HardwareVersion:          hardwareVersion,
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
