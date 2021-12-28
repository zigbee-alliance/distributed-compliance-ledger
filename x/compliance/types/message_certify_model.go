package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCertifyModel = "certify_model"

var _ sdk.Msg = &MsgCertifyModel{}

func NewMsgCertifyModel(signer string, vid int32, pid int32, softwareVersion uint64, softwareVersionString string, certificationDate string, certificationType string, reason string) *MsgCertifyModel {
	return &MsgCertifyModel{
		Signer:                signer,
		Vid:                   vid,
		Pid:                   pid,
		SoftwareVersion:       softwareVersion,
		SoftwareVersionString: softwareVersionString,
		CertificationDate:     certificationDate,
		CertificationType:     certificationType,
		Reason:                reason,
	}
}

func (msg *MsgCertifyModel) Route() string {
	return RouterKey
}

func (msg *MsgCertifyModel) Type() string {
	return TypeMsgCertifyModel
}

func (msg *MsgCertifyModel) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{signer}
}

func (msg *MsgCertifyModel) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCertifyModel) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signer address (%s)", err)
	}
	return nil
}
