package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgDeletePkiRevocationDistributionPoint = "delete_pki_revocation_distribution_point"

var _ sdk.Msg = &MsgDeletePkiRevocationDistributionPoint{}

func NewMsgDeletePkiRevocationDistributionPoint(signer string, vid int32, label string, issuerSubjectKeyID string) *MsgDeletePkiRevocationDistributionPoint {
	return &MsgDeletePkiRevocationDistributionPoint{
		Signer:             signer,
		Vid:                vid,
		Label:              label,
		IssuerSubjectKeyID: issuerSubjectKeyID,
	}
}

func (msg *MsgDeletePkiRevocationDistributionPoint) Route() string {
	return RouterKey
}

func (msg *MsgDeletePkiRevocationDistributionPoint) Type() string {
	return TypeMsgDeletePkiRevocationDistributionPoint
}

func (msg *MsgDeletePkiRevocationDistributionPoint) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{signer}
}

func (msg *MsgDeletePkiRevocationDistributionPoint) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeletePkiRevocationDistributionPoint) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signer address (%s)", err)
	}
	return nil
}
