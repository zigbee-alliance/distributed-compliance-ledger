package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const RouterKey = ModuleName

type MsgAddAccount struct {
	Address   sdk.AccAddress `json:"address"`
	PublicKey string         `json:"pub_key"`
	Signer    sdk.AccAddress `json:"signer"`
}

func NewMsgAddAccount(address sdk.AccAddress, pubKey string, signer sdk.AccAddress) MsgAddAccount {
	return MsgAddAccount{Address: address, PublicKey: pubKey, Signer: signer}
}

func (m MsgAddAccount) Route() string {
	return RouterKey
}

func (m MsgAddAccount) Type() string {
	return "add_account"
}

func (m MsgAddAccount) ValidateBasic() sdk.Error {
	if m.Address.Empty() {
		return sdk.ErrInvalidAddress(m.Address.String())
	}

	if len(m.PublicKey) == 0 {
		return sdk.ErrUnknownRequest("Invalid PublicKey: it cannot be empty")
	}

	return nil
}

func (m MsgAddAccount) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(m))
}

func (m MsgAddAccount) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Signer}
}
