package types

import (
	"bytes"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgProposeAddAccount = "propose_add_account"

var _ sdk.Msg = &MsgProposeAddAccount{}

func NewMsgProposeAddAccount(
	signer sdk.AccAddress,
	address sdk.AccAddress,
	pubKey cryptotypes.PubKey, //nolint:interfacer
	roles AccountRoles,
	// roles []string,
	vendorID uint64,
) (*MsgProposeAddAccount, error) {
	var pkAny *codectypes.Any
	if pubKey != nil {
		var err error
		if pkAny, err = codectypes.NewAnyWithValue(pubKey); err != nil {
			return nil, err
		}
	}

	return &MsgProposeAddAccount{
		Signer:   signer.String(),
		Address:  address.String(),
		PubKey:   pkAny,
		Roles:    roles,
		VendorID: vendorID,
	}, nil
}

func (msg *MsgProposeAddAccount) Route() string {
	return RouterKey
}

func (msg *MsgProposeAddAccount) Type() string {
	return TypeMsgProposeAddAccount
}

func (msg *MsgProposeAddAccount) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{signer}
}

func (msg *MsgProposeAddAccount) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgProposeAddAccount) ValidateBasic() error {
	if m.Address.Empty() {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid Account Address: it cannot be empty")
	}

	accAddr, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid Account Address: (%s)", err)
	}

	if m.Signer.Empty() {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid Signer: it cannot be empty")
	}

	_, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid Signer: (%s)", err)
	}

	if len(m.PubKey) == 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "Invalid PublicKey: it cannot be empty")
	}

	pk, err := msg.Pubkey.GetCachedValue().(cryptotypes.PubKey)
	if !err {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "expecting cryptotypes.PubKey for PubKey, got %T", err)
	}

	if !bytes.Equal(pk.Address().Bytes(), accAddr.Bytes()) {
		return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "account address and pubkey address do not match")
	}

	if err := m.Roles.Validate(); err != nil {
		return err
	}

	if m.HasRole(Vendor) && m.VendorID <= 0 {
		return ErrMissingVendorIDForVendorAccount()
	}

	return nil
}

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (msg MsgProposeAddAccount) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	var pubKey cryptotypes.PubKey
	return unpacker.UnpackAny(msg.Pubkey, &pubKey)
}
