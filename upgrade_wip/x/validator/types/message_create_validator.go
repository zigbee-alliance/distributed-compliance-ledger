package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateValidator{}

func NewMsgCreateValidator(signer string, address string, pubKey string, description *Description) *MsgCreateValidator {
	return &MsgCreateValidator{
		Signer:      signer,
		Address:     address,
		PubKey:      pubKey,
		Description: description,
	}
}

func (msg *MsgCreateValidator) Route() string {
	return RouterKey
}

func (msg *MsgCreateValidator) Type() string {
	return EventTypeCreateValidator
}

func (msg *MsgCreateValidator) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{signer}
}

func (msg *MsgCreateValidator) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateValidator) ValidateBasic() error {
	accAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signer address (%s)", err)
	}

	if accAddr.Empty() {
		return sdk.ErrInvalidAddress("Invalid Signer: it cannot be empty")
	}

	valAddr, err := sdk.ConsAddressFromBech32(msg.Address)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid validator address (%s)", err)
	}

	if valAddr.Empty() {
		return sdk.ErrUnknownRequest("Invalid Validator Address: it cannot be empty")
	}

	if msg.Pubkey == nil {
		return ErrEmptyValidatorPubKey
	}

	if msg.Description == (Description{}) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "empty description")
	}

	if err := msg.Description.Validate(); err != nil {
		return err
	}

	// TODO issue 99
	/*
		// OLD CHECKS

		pubkey, err := sdk.GetConsPubKeyBech32(m.PubKey)
		if err != nil {
			return sdk.ErrUnknownRequest(fmt.Sprintf("Invalid Validator Public Key: %v", err))
		}

		if !valAddr.Equals(sdk.ConsAddress(pubkey.Address())) {
			return sdk.ErrUnknownRequest("Validator Pubkey does not match to Validator Address")
		}

	*/

	return nil
}
