package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto"
)

type MsgCreateValidator struct {
	Address     sdk.ConsAddress `json:"validator_address"`
	PubKey      string          `json:"validator_pubkey"`
	Description Description     `json:"description"`
	Signer      sdk.AccAddress  `json:"signer"`
}

func NewMsgCreateValidator(address sdk.ConsAddress, pubKey string,
	description Description, signer sdk.AccAddress) MsgCreateValidator {
	return MsgCreateValidator{
		Address:     address,
		PubKey:      pubKey,
		Description: description,
		Signer:      signer,
	}
}

func (m MsgCreateValidator) Route() string { return RouterKey }

func (m MsgCreateValidator) Type() string { return EventTypeCreateValidator }

func (m MsgCreateValidator) ValidateBasic() sdk.Error {
	if m.Signer.Empty() {
		return sdk.ErrInvalidAddress("Invalid Signer: it cannot be empty")
	}

	if m.Address.Empty() {
		return sdk.ErrUnknownRequest("Invalid Validator Address: it cannot be empty")
	}

	pubkey, err := sdk.GetConsPubKeyBech32(m.PubKey)
	if err != nil {
		return sdk.ErrUnknownRequest(fmt.Sprintf("Invalid Validator Public Key: %v", err))
	}

	if !m.Address.Equals(sdk.ConsAddress(pubkey.Address())) {
		return sdk.ErrUnknownRequest("Validator Pubkey does not match to Validator Address")
	}

	if len(m.Description.Name) == 0 {
		return sdk.ErrUnknownRequest("Invalid Validator Name: it cannot be empty")
	}
	if err := m.Description.Validate(); err != nil {
		return err
	}
	return nil
}

func (m MsgCreateValidator) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Signer}
}

func (m MsgCreateValidator) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(m))
}

func (m MsgCreateValidator) GetPubKey() crypto.PubKey {
	return sdk.MustGetConsPubKeyBech32(m.PubKey)
}
