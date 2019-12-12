package types

import (
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type ModelInfo struct {
	ID     string         `json:"id"`
	Family string         `json:"family"`
	Cert   string         `json:"cert"`
	Owner  sdk.AccAddress `json:"owner"`
}

func NewModelInfo(id string, family string, cert string, owner sdk.AccAddress) ModelInfo {
	return ModelInfo{ID: id, Family: family, Cert: cert, Owner: owner}
}

func (d ModelInfo) String() string {
	bytes, err := json.Marshal(d)
	if err != nil {
		panic(err)
	}

	return string(bytes)
}
