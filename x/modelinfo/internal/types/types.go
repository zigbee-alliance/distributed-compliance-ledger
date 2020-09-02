package types

//nolint:goimports
import (
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

/*
	Model Info stored in KVStore
*/
type ModelInfo struct {
	Name                     string         `json:"name"`
	Description              string         `json:"description"`
	SKU                      string         `json:"sku"`
	FirmwareVersion          string         `json:"firmware_version"`
	HardwareVersion          string         `json:"hardware_version"`
	Custom                   string         `json:"custom,omitempty"`
	TisOrTrpTestingCompleted bool           `json:"tis_or_trp_testing_completed"`
	VID                      uint16         `json:"vid"`
	PID                      uint16         `json:"pid"`
	CID                      uint16         `json:"cid,omitempty"`
	Owner                    sdk.AccAddress `json:"owner"`
}

func NewModelInfo(vid uint16, pid uint16, cid uint16, name string, owner sdk.AccAddress,
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

/*
	List of products for specific Vendor
*/
type VendorProducts struct {
	VID      uint16    `json:"vid"`
	Products []Product `json:"products"`
}

func NewVendorProducts(vid uint16) VendorProducts {
	return VendorProducts{
		VID:      vid,
		Products: []Product{},
	}
}

func (d *VendorProducts) AddVendorProduct(pid Product) {
	d.Products = append(d.Products, pid)
}

func (d *VendorProducts) RemoveVendorProduct(pid uint16) {
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

// Single Vendor Product.
type Product struct {
	PID   uint16         `json:"pid"`
	Name  string         `json:"name"`
	Owner sdk.AccAddress `json:"owner"`
	SKU   string         `json:"sku"`
}
