package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/validator"
)

const TypeMsgProposeRevokeAccount = "propose_revoke_account"

var _ sdk.Msg = &MsgProposeRevokeAccount{}

func NewMsgProposeRevokeAccount(signer sdk.AccAddress, address sdk.AccAddress, info string) *MsgProposeRevokeAccount {
	return &MsgProposeRevokeAccount{
		Signer:  signer.String(),
		Address: address.String(),
		Info:    info,
		Time:    time.Now().Unix(),
	}
}

func (msg *MsgProposeRevokeAccount) Route() string {
	return RouterKey
}

func (msg *MsgProposeRevokeAccount) Type() string {
	return TypeMsgProposeRevokeAccount
}

func (msg *MsgProposeRevokeAccount) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{signer}
}

func (msg *MsgProposeRevokeAccount) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgProposeRevokeAccount) ValidateBasic() error {
	if msg.Address == "" {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid Account Address: it cannot be empty")
	}

	_, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid Account Address: (%s)", err)
	}

	if msg.Signer == "" {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid Signer: it cannot be empty")
	}

	_, err = sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid Signer: (%s)", err)
	}

	err = validator.Validate(msg)
	if err != nil {
		return err
	}

	return nil
}
