package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateModelVersions = "create_model_versions"
	TypeMsgUpdateModelVersions = "update_model_versions"
	TypeMsgDeleteModelVersions = "delete_model_versions"
)

var _ sdk.Msg = &MsgCreateModelVersions{}

func NewMsgCreateModelVersions(
	creator string,
	vid int32,
	pid int32,
	softwareVersions []uint64,

) *MsgCreateModelVersions {
	return &MsgCreateModelVersions{
		Creator:          creator,
		Vid:              vid,
		Pid:              pid,
		SoftwareVersions: softwareVersions,
	}
}

func (msg *MsgCreateModelVersions) Route() string {
	return RouterKey
}

func (msg *MsgCreateModelVersions) Type() string {
	return TypeMsgCreateModelVersions
}

func (msg *MsgCreateModelVersions) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateModelVersions) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateModelVersions) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if msg.Vid < 0 || msg.Vid > 65535 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Vid must be in range from 0 to 65535")
	}

	if msg.Pid < 0 || msg.Pid > 65535 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Pid must be in range from 0 to 65535")
	}

	for _, softwareVersion := range msg.SoftwareVersions {
		if softwareVersion > 4294967295 {
			return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Each element of SoftwareVersions must not be greater than 4294967295")
		}
	}

	return nil
}

var _ sdk.Msg = &MsgUpdateModelVersions{}

func NewMsgUpdateModelVersions(
	creator string,
	vid int32,
	pid int32,
	softwareVersions []uint64,

) *MsgUpdateModelVersions {
	return &MsgUpdateModelVersions{
		Creator:          creator,
		Vid:              vid,
		Pid:              pid,
		SoftwareVersions: softwareVersions,
	}
}

func (msg *MsgUpdateModelVersions) Route() string {
	return RouterKey
}

func (msg *MsgUpdateModelVersions) Type() string {
	return TypeMsgUpdateModelVersions
}

func (msg *MsgUpdateModelVersions) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateModelVersions) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateModelVersions) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if msg.Vid < 0 || msg.Vid > 65535 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Vid must be in range from 0 to 65535")
	}

	if msg.Pid < 0 || msg.Pid > 65535 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Pid must be in range from 0 to 65535")
	}

	for _, softwareVersion := range msg.SoftwareVersions {
		if softwareVersion > 4294967295 {
			return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Each element of SoftwareVersions must not be greater than 4294967295")
		}
	}

	return nil
}

var _ sdk.Msg = &MsgDeleteModelVersions{}

func NewMsgDeleteModelVersions(
	creator string,
	vid int32,
	pid int32,

) *MsgDeleteModelVersions {
	return &MsgDeleteModelVersions{
		Creator: creator,
		Vid:     vid,
		Pid:     pid,
	}
}
func (msg *MsgDeleteModelVersions) Route() string {
	return RouterKey
}

func (msg *MsgDeleteModelVersions) Type() string {
	return TypeMsgDeleteModelVersions
}

func (msg *MsgDeleteModelVersions) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteModelVersions) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteModelVersions) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if msg.Vid < 0 || msg.Vid > 65535 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Vid must be in range from 0 to 65535")
	}

	if msg.Pid < 0 || msg.Pid > 65535 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Pid must be in range from 0 to 65535")
	}

	return nil
}
