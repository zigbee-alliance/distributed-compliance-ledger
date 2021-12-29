package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgApproveRevokeAccount = "approve_revoke_account"

var _ sdk.Msg = &MsgApproveRevokeAccount{}

func NewMsgApproveRevokeAccount(signer sdk.AccAddress, address sdk.AccAddress) *MsgApproveRevokeAccount {
	return &MsgApproveRevokeAccount{
		Signer:  signer.String(),
		Address: address.String(),
	}
}

func (msg *MsgApproveRevokeAccount) Route() string {
	return RouterKey
}

func (msg *MsgApproveRevokeAccount) Type() string {
	return TypeMsgApproveRevokeAccount
}

func (msg *MsgApproveRevokeAccount) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{signer}
}

func (msg *MsgApproveRevokeAccount) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgApproveRevokeAccount) ValidateBasic() error {
	if msg.Address == "" {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid Account Address: it cannot be empty")
	}

	_, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid Account Address: (%s)", err)
	}

	if msg.Signer == "" {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid Signer: it cannot be empty")
	}

	_, err = sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid Signer: (%s)", err)
	}

	err = validator.Validate(msg)
	if err != nil {
		return err
	}

	return nil
}
