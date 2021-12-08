package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateNewVendorInfo{}

func NewMsgCreateNewVendorInfo(
	creator string,
	index string,
	vendorInfo *VendorInfo,

) *MsgCreateNewVendorInfo {
	return &MsgCreateNewVendorInfo{
		Creator:    creator,
		Index:      index,
		VendorInfo: vendorInfo,
	}
}

func (msg *MsgCreateNewVendorInfo) Route() string {
	return RouterKey
}

func (msg *MsgCreateNewVendorInfo) Type() string {
	return "CreateNewVendorInfo"
}

func (msg *MsgCreateNewVendorInfo) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateNewVendorInfo) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateNewVendorInfo) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateNewVendorInfo{}

func NewMsgUpdateNewVendorInfo(
	creator string,
	index string,
	vendorInfo *VendorInfo,

) *MsgUpdateNewVendorInfo {
	return &MsgUpdateNewVendorInfo{
		Creator:    creator,
		Index:      index,
		VendorInfo: vendorInfo,
	}
}

func (msg *MsgUpdateNewVendorInfo) Route() string {
	return RouterKey
}

func (msg *MsgUpdateNewVendorInfo) Type() string {
	return "UpdateNewVendorInfo"
}

func (msg *MsgUpdateNewVendorInfo) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateNewVendorInfo) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateNewVendorInfo) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteNewVendorInfo{}

func NewMsgDeleteNewVendorInfo(
	creator string,
	index string,

) *MsgDeleteNewVendorInfo {
	return &MsgDeleteNewVendorInfo{
		Creator: creator,
		Index:   index,
	}
}
func (msg *MsgDeleteNewVendorInfo) Route() string {
	return RouterKey
}

func (msg *MsgDeleteNewVendorInfo) Type() string {
	return "DeleteNewVendorInfo"
}

func (msg *MsgDeleteNewVendorInfo) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteNewVendorInfo) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteNewVendorInfo) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
