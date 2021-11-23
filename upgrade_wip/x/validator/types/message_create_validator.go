package types

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	_ sdk.Msg                            = &MsgCreateValidator{}
	_ codectypes.UnpackInterfacesMessage = (*MsgCreateValidator)(nil)
)

func NewMsgCreateValidator(
	signer sdk.ValAddress,
	pubKey cryptotypes.PubKey, //nolint:interfacer
	description *Description,
) (*MsgCreateValidator, error) {
	var pkAny *codectypes.Any
	if pubKey != nil {
		var err error
		if pkAny, err = codectypes.NewAnyWithValue(pubKey); err != nil {
			return nil, err
		}
	}

	return &MsgCreateValidator{
		Signer:      signer.String(),
		PubKey:      pkAny,
		Description: description,
	}, nil
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
		return sdkerrors.Wrap(sdk.ErrInvalidAddress, "Invalid Signer: it cannot be empty")
	}

	if valAddr.Empty() {
		return sdkerrors.Wrap(sdk.ErrUnknownRequest, "Invalid Validator Address: it cannot be empty")
	}

	if msg.Pubkey == nil {
		return sdkerrors.Wrap(ErrInvalidPubKey, "Invalid Validator PubKey: it cannot be empty")
	}

	_, err := msg.Pubkey.GetCachedValue().(cryptotypes.PubKey)
	if !err {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "expecting cryptotypes.PubKey for PubKey, got %T", err)
	}

	if msg.Description == (Description{}) {
		return sdkerrors.Wrap(sdkerrors.Wrap, sdkerrors.ErrInvalidRequest, "empty description")
	}

	if err := msg.Description.Validate(); err != nil {
		return err
	}

	return nil
}

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (msg MsgCreateValidator) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	var pubKey cryptotypes.PubKey
	return unpacker.UnpackAny(msg.Pubkey, &pubKey)
}
