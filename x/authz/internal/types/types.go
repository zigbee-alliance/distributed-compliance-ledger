package types

import (
	"encoding/json"
	"errors"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type AccountRole string

const (
	Administrator AccountRole = "Administrator"
	Vendor        AccountRole = "Vendor"
	TestHouse     AccountRole = "TestHouse"
)

func (lt *AccountRole) UnmarshalJSON(b []byte) error {
	accountRole := AccountRole(strings.Trim(string(b), `"`))

	switch accountRole {
	case Administrator, Vendor, TestHouse:
		*lt = accountRole
		return nil
	}

	return errors.New("invalid account role")
}

type AccountRoles struct {
	Address sdk.AccAddress `json:"address"`
	// We are not using map because it's unsupported by go-amino (Cosmos's serializer)
	Roles []AccountRole `json:"roles"`
}

func NewAccountRoles(address sdk.AccAddress, roles []AccountRole) AccountRoles {
	return AccountRoles{Address: address, Roles: roles}
}

func (d AccountRoles) String() string {
	bytes, err := json.Marshal(d)
	if err != nil {
		panic(err)
	}

	return string(bytes)
}
