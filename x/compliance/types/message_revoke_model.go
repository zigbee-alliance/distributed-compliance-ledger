package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRevokeModel = "revoke_model"

var _ sdk.Msg = &MsgRevokeModel{}

func NewMsgRevokeModel(signer string, vid int32, pid int32, softwareVersion uint32, softwareVersionString string, revocationDate string, certificationType string, reason string) *MsgRevokeModel {
	return &MsgRevokeModel{
		Signer:                signer,
		Vid:                   vid,
		Pid:                   pid,
		SoftwareVersion:       softwareVersion,
		SoftwareVersionString: softwareVersionString,
		RevocationDate:        revocationDate,
		CertificationType:     certificationType,
		Reason:                reason,
	}
}

func (msg *MsgRevokeModel) Route() string {
	return RouterKey
}

func (msg *MsgRevokeModel) Type() string {
	return TypeMsgRevokeModel
}

func (msg *MsgRevokeModel) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{signer}
}

func (msg *MsgRevokeModel) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRevokeModel) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signer address (%s)", err)
	}
	return nil
}
