package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/validator"
)

var _ sdk.Msg = &MsgCreateVendorInfo{}

func NewMsgCreateVendorInfo(
	creator string,
	vendorID int32,
	vendorName string,
	companyLegalName string,
	companyPreferredName string,
	vendorLandingPageURL string,
	schemaVersion uint32,
) *MsgCreateVendorInfo {
	return &MsgCreateVendorInfo{
		Creator:              creator,
		VendorID:             vendorID,
		VendorName:           vendorName,
		CompanyLegalName:     companyLegalName,
		CompanyPreferredName: companyPreferredName,
		VendorLandingPageURL: vendorLandingPageURL,
		SchemaVersion:        schemaVersion,
	}
}

func (msg *MsgCreateVendorInfo) Route() string {
	return RouterKey
}

func (msg *MsgCreateVendorInfo) Type() string {
	return "CreateVendorInfo"
}

func (msg *MsgCreateVendorInfo) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateVendorInfo) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)

	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateVendorInfo) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	err = validator.Validate(msg)
	if err != nil {
		return err
	}

	return nil
}

var _ sdk.Msg = &MsgUpdateVendorInfo{}

func NewMsgUpdateVendorInfo(
	creator string,
	vendorID int32,
	vendorName string,
	companyLegalName string,
	companyPreferredName string,
	vendorLandingPageURL string,
	schemaVersion uint32,
) *MsgUpdateVendorInfo {
	return &MsgUpdateVendorInfo{
		Creator:              creator,
		VendorID:             vendorID,
		VendorName:           vendorName,
		CompanyLegalName:     companyLegalName,
		CompanyPreferredName: companyPreferredName,
		VendorLandingPageURL: vendorLandingPageURL,
		SchemaVersion:        schemaVersion,
	}
}

func (msg *MsgUpdateVendorInfo) Route() string {
	return RouterKey
}

func (msg *MsgUpdateVendorInfo) Type() string {
	return "UpdateVendorInfo"
}

func (msg *MsgUpdateVendorInfo) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateVendorInfo) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)

	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateVendorInfo) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	err = validator.Validate(msg)
	if err != nil {
		return err
	}

	return nil
}

// var _ sdk.Msg = &MsgDeleteVendorInfo{}

// func NewMsgDeleteVendorInfo(
// 	creator string,
// 	vendorID int32,

// ) *MsgDeleteVendorInfo {
// 	return &MsgDeleteVendorInfo{
// 		Creator:  creator,
// 		VendorID: vendorID,
// 	}
// }

// func (msg *MsgDeleteVendorInfo) Route() string {
// 	return RouterKey
// }

// func (msg *MsgDeleteVendorInfo) Type() string {
// 	return "DeleteVendorInfo"
// }

// func (msg *MsgDeleteVendorInfo) GetSigners() []sdk.AccAddress {
// 	creator, err := sdk.AccAddressFromBech32(msg.Creator)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return []sdk.AccAddress{creator}
// }

// func (msg *MsgDeleteVendorInfo) GetSignBytes() []byte {
// 	bz := ModuleCdc.MustMarshalJSON(msg)
// 	return sdk.MustSortJSON(bz)
// }

// func (msg *MsgDeleteVendorInfo) ValidateBasic() error {
// 	_, err := sdk.AccAddressFromBech32(msg.Creator)
// 	if err != nil {
// 		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
// 	}

// 	err = validator.Validate(msg)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
