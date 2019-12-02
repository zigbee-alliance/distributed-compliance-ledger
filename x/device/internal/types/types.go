package types

import (
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Device struct {
	Family     string                `json:"family"`
	Cert       string                `json:"cert"`
	Owner      sdk.AccAddress        `json:"owner"`
	Compliance map[string]Compliance `json:"compliance"`
}

func NewDevice() Device {
	return Device{
		Compliance: map[string]Compliance{},
	}
}

func (d Device) String() string {
	bytes, err := json.Marshal(d)
	if err != nil {
		panic(err)
	}

	return string(bytes)
}

type Compliance struct {
	Description string           `json:"description"`
	TestedBy    sdk.AccAddress   `json:"tested_by"`
	ApprovedBy  []sdk.AccAddress `json:"approved_by"`
}
