package types

import (
	"encoding/json"
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type AccountRole string

const (
	Administrator         AccountRole = "Administrator"
	Vendor                AccountRole = "Vendor"
	TestHouse             AccountRole = "TestHouse"
	ZBCertificationCenter AccountRole = "ZBCertificationCenter"
	Trustee               AccountRole = "Trustee"
	NodeAdmin             AccountRole = "NodeAdmin"
)

func (lt *AccountRole) UnmarshalJSON(b []byte) error {
	accountRole := AccountRole(strings.Trim(string(b), `"`))

	switch accountRole {
	case Administrator, Vendor, TestHouse, ZBCertificationCenter, Trustee, NodeAdmin:
		*lt = accountRole
		return nil
	}

	return sdk.ErrUnknownRequest(fmt.Sprintf("invalid account role: %s", accountRole))
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
