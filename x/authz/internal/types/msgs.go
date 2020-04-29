package types

import sdk "github.com/cosmos/cosmos-sdk/types"

const RouterKey = ModuleName

type MsgAssignRole struct {
	Address sdk.AccAddress `json:"address"`
	Role    AccountRole    `json:"role"`
	Signer  sdk.AccAddress `json:"signer"`
}

func NewMsgAssignRole(address sdk.AccAddress, role AccountRole, signer sdk.AccAddress) MsgAssignRole {
	return MsgAssignRole{Address: address, Role: role, Signer: signer}
}

func (m MsgAssignRole) Route() string {
	return RouterKey
}

func (m MsgAssignRole) Type() string {
	return "assign_role"
}

func (m MsgAssignRole) ValidateBasic() sdk.Error {
	if m.Address.Empty() {
		return sdk.ErrInvalidAddress("Invalid Address: it cannot be empty")
	}

	if m.Signer.Empty() {
		return sdk.ErrInvalidAddress("Invalid Signer: it cannot be empty")
	}

	return nil
}

func (m MsgAssignRole) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(m))
}

func (m MsgAssignRole) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Signer}
}

type MsgRevokeRole struct {
	Address sdk.AccAddress `json:"address"`
	Role    AccountRole    `json:"role"`
	Signer  sdk.AccAddress `json:"signer"`
}

func NewMsgRevokeRole(address sdk.AccAddress, role AccountRole, signer sdk.AccAddress) MsgRevokeRole {
	return MsgRevokeRole{Address: address, Role: role, Signer: signer}
}

func (m MsgRevokeRole) Route() string {
	return RouterKey
}

func (m MsgRevokeRole) Type() string {
	return "revoke_role"
}

func (m MsgRevokeRole) ValidateBasic() sdk.Error {
	if m.Address.Empty() {
		return sdk.ErrInvalidAddress("Invalid Address: it cannot be empty")
	}

	if m.Signer.Empty() {
		return sdk.ErrInvalidAddress("Invalid Signer: it cannot be empty")
	}

	return nil
}

func (m MsgRevokeRole) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(m))
}

func (m MsgRevokeRole) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Signer}
}
