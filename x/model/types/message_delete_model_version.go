package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgDeleteModelVersion = "delete_model_version"

var _ sdk.Msg = &MsgDeleteModelVersion{}

func NewMsgDeleteModelVersion(signer string, vid int32, pid int32, softwareVersion int32) *MsgDeleteModelVersion {
	return &MsgDeleteModelVersion{
		Signer:          signer,
		Vid:             vid,
		Pid:             pid,
		SoftwareVersion: softwareVersion,
	}
}

func (msg *MsgDeleteModelVersion) Route() string {
	return RouterKey
}

func (msg *MsgDeleteModelVersion) Type() string {
	return TypeMsgDeleteModelVersion
}

func (msg *MsgDeleteModelVersion) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{signer}
}

func (msg *MsgDeleteModelVersion) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteModelVersion) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signer address (%s)", err)
	}
	return nil
}
