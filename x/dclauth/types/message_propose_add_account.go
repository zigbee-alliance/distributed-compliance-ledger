package types

import (
	"bytes"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/validator"
)

const TypeMsgProposeAddAccount = "propose_add_account"

var _ sdk.Msg = &MsgProposeAddAccount{}

func NewMsgProposeAddAccount(
	signer sdk.AccAddress,
	address sdk.AccAddress,
	pubKey cryptotypes.PubKey, //nolint:interfacer
	roles AccountRoles,
	// roles []string,
	vendorID int32,
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

func (msg *MsgProposeAddAccount) HasRole(targetRole AccountRole) bool {
	for _, role := range msg.Roles {
		if role == targetRole {
			return true
		}
	}

	return false
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
	if msg.Address == "" {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid Account Address: it cannot be empty")
	}

	accAddr, err := sdk.AccAddressFromBech32(msg.Address)
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

	if msg.PubKey == nil {
		return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "Invalid PublicKey: it cannot be empty")
	}

	pk, err2 := msg.PubKey.GetCachedValue().(cryptotypes.PubKey)
	if !err2 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "expecting cryptotypes.PubKey for PubKey, got %T", err)
	}

	if !bytes.Equal(pk.Address().Bytes(), accAddr.Bytes()) {
		return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "account address and pubkey address do not match")
	}

	if len(msg.Roles) == 0 {
		return ErrMissingRoles()
	}

	for _, role := range msg.Roles {
		if err := role.Validate(); err != nil {
			return err
		}
	}

	// can not create Vendor with vid=0 (reserved)
	if msg.HasRole(Vendor) && msg.VendorID <= 0 {
		return ErrMissingVendorIDForVendorAccount()
	}

	err = validator.Validate(msg)
	if err != nil {
		return err
	}

	return nil
}

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces.
func (msg MsgProposeAddAccount) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	var pubKey cryptotypes.PubKey
	return unpacker.UnpackAny(msg.PubKey, &pubKey)
}
