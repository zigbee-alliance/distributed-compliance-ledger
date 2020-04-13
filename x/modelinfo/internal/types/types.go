package types

import (
	"encoding/json"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/utils/conversions"
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
	Custom                   string         `json:"custom,omitempty"`
	TisOrTrpTestingCompleted bool           `json:"tis_or_trp_testing_completed"`
}

func NewModelInfo(vid int16, pid int16, cid int16, name string, owner sdk.AccAddress,
	description string, sku string, firmwareVersion string, hardwareVersion string, custom string,
	tisOrTrpTestingCompleted bool) ModelInfo {
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

type VendorProducts struct {
	VID      int16     `json:"vid"`
	Products []Product `json:"products"`
}

func NewVendorProducts(vid int16) VendorProducts {
	return VendorProducts{
		VID:      vid,
		Products: []Product{},
	}
}

func (d *VendorProducts) AddVendorProduct(pid Product) {
	d.Products = append(d.Products, pid)
}

func (d *VendorProducts) RemoveVendorProduct(pid int16) {
	for i, value := range d.Products {
		if pid == value.PID {
			d.Products = append(d.Products[:i], d.Products[i+1:]...)
			return
		}
	}
}

func (d *VendorProducts) IsEmpty() bool {
	return len(d.Products) == 0
}

func (d VendorProducts) String() string {
	bytes, err := json.Marshal(d)
	if err != nil {
		panic(err)
	}

	return string(bytes)
}

type Product struct {
	PID   int16          `json:"pid"`
	Name  string         `json:"name"`
	Owner sdk.AccAddress `json:"owner"`
	SKU   string         `json:"sku"`
}

func ParseVID(str string) (int16, error) {
	return conversions.ParseInt16FromString(str)
}

func ParsePID(str string) (int16, error) {
	return conversions.ParseInt16FromString(str)
}

func ParseCID(str string) (int16, error) {
	return conversions.ParseInt16FromString(str)
}
