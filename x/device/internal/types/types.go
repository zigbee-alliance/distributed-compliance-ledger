package types

import sdk "github.com/cosmos/cosmos-sdk/types"

type Device struct {
	Family     string         `json:"family"`
	Cert       string         `json:"cert"`
	Owner      sdk.AccAddress `json:"owner"`
	Compliance []Compliance   `json:"compliance"`
}

type Compliance struct {
	Test      string           `json:"test"`
	Tester    sdk.AccAddress   `json:"tester"`
	Approvers []sdk.AccAddress `json:"approvers"`
}
