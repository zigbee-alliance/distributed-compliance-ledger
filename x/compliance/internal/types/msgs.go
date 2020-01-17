package types

import sdk "github.com/cosmos/cosmos-sdk/types"

const RouterKey = ModuleName

type MsgAddModelInfo struct {
	ID     string         `json:"id"`
	Family string         `json:"family"`
	Cert   string         `json:"cert"`
	Owner  sdk.AccAddress `json:"owner"`
	Signer sdk.AccAddress `json:"signer"`
}

func NewMsgAddModelInfo(id string, family string, cert string, owner sdk.AccAddress,
	signer sdk.AccAddress) MsgAddModelInfo {
	return MsgAddModelInfo{ID: id, Family: family, Cert: cert, Owner: owner, Signer: signer}
}

func (m MsgAddModelInfo) Route() string {
	return RouterKey
}

func (m MsgAddModelInfo) Type() string {
	return "add_model_info"
}

func (m MsgAddModelInfo) ValidateBasic() sdk.Error {
	if m.Owner.Empty() {
		return sdk.ErrInvalidAddress(m.Owner.String())
	}

	if m.Signer.Empty() {
		return sdk.ErrInvalidAddress(m.Signer.String())
	}

	if len(m.ID) == 0 || len(m.Family) == 0 || len(m.Cert) == 0 {
		return sdk.ErrUnknownRequest("ID, Family and Cert cannot be empty")
	}

	return nil
}

func (m MsgAddModelInfo) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(m))
}

func (m MsgAddModelInfo) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Signer}
}

type MsgUpdateModelInfo struct {
	ID        string         `json:"id"`
	NewFamily string         `json:"new_family"`
	NewCert   string         `json:"new_cert"`
	NewOwner  sdk.AccAddress `json:"owner"`
	Signer    sdk.AccAddress `json:"signer"`
}

func NewMsgUpdateModelInfo(id string, newFamily string, newCert string, newOwner sdk.AccAddress,
	signer sdk.AccAddress) MsgUpdateModelInfo {
	return MsgUpdateModelInfo{ID: id, NewFamily: newFamily, NewCert: newCert, NewOwner: newOwner, Signer: signer}
}

func (m MsgUpdateModelInfo) Route() string {
	return RouterKey
}

func (m MsgUpdateModelInfo) Type() string {
	return "update_model_info"
}

func (m MsgUpdateModelInfo) ValidateBasic() sdk.Error {
	if m.NewOwner.Empty() {
		return sdk.ErrInvalidAddress(m.NewOwner.String())
	}

	if m.Signer.Empty() {
		return sdk.ErrInvalidAddress(m.Signer.String())
	}

	if len(m.ID) == 0 || len(m.NewFamily) == 0 || len(m.NewCert) == 0 {
		return sdk.ErrUnknownRequest("ID, NewFamily and NewCert cannot be empty")
	}

	return nil
}

func (m MsgUpdateModelInfo) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(m))
}

func (m MsgUpdateModelInfo) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Signer}
}

type MsgDeleteModelInfo struct {
	ID     string         `json:"id"`
	Signer sdk.AccAddress `json:"signer"`
}

func NewMsgDeleteModelInfo(id string, signer sdk.AccAddress) MsgDeleteModelInfo {
	return MsgDeleteModelInfo{ID: id, Signer: signer}
}

func (m MsgDeleteModelInfo) Route() string {
	return RouterKey
}

func (m MsgDeleteModelInfo) Type() string {
	return "delete_model_info"
}

func (m MsgDeleteModelInfo) ValidateBasic() sdk.Error {
	if m.Signer.Empty() {
		return sdk.ErrInvalidAddress(m.Signer.String())
	}

	if len(m.ID) == 0 {
		return sdk.ErrUnknownRequest("ID cannot be empty")
	}

	return nil
}

func (m MsgDeleteModelInfo) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(m))
}

func (m MsgDeleteModelInfo) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Signer}
}
