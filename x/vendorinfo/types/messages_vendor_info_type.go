package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateVendorInfoType{}

func NewMsgCreateVendorInfoType(
    creator string,
    index string,
    vendorID uint64,
    vendorName string,
    companyLegalName string,
    companyPrefferedName string,
    vendorLandingPageURL string,
    
) *MsgCreateVendorInfoType {
  return &MsgCreateVendorInfoType{
		Creator : creator,
		Index: index,
		VendorID: vendorID,
        VendorName: vendorName,
        CompanyLegalName: companyLegalName,
        CompanyPrefferedName: companyPrefferedName,
        VendorLandingPageURL: vendorLandingPageURL,
        
	}
}

func (msg *MsgCreateVendorInfoType) Route() string {
  return RouterKey
}

func (msg *MsgCreateVendorInfoType) Type() string {
  return "CreateVendorInfoType"
}

func (msg *MsgCreateVendorInfoType) GetSigners() []sdk.AccAddress {
  creator, err := sdk.AccAddressFromBech32(msg.Creator)
  if err != nil {
    panic(err)
  }
  return []sdk.AccAddress{creator}
}

func (msg *MsgCreateVendorInfoType) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateVendorInfoType) ValidateBasic() error {
  _, err := sdk.AccAddressFromBech32(msg.Creator)
  	if err != nil {
  		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
  	}
  return nil
}

var _ sdk.Msg = &MsgUpdateVendorInfoType{}

func NewMsgUpdateVendorInfoType(
    creator string,
    index string,
    vendorID uint64,
    vendorName string,
    companyLegalName string,
    companyPrefferedName string,
    vendorLandingPageURL string,
    
) *MsgUpdateVendorInfoType {
  return &MsgUpdateVendorInfoType{
		Creator: creator,
        Index: index,
        VendorID: vendorID,
        VendorName: vendorName,
        CompanyLegalName: companyLegalName,
        CompanyPrefferedName: companyPrefferedName,
        VendorLandingPageURL: vendorLandingPageURL,
        
	}
}

func (msg *MsgUpdateVendorInfoType) Route() string {
  return RouterKey
}

func (msg *MsgUpdateVendorInfoType) Type() string {
  return "UpdateVendorInfoType"
}

func (msg *MsgUpdateVendorInfoType) GetSigners() []sdk.AccAddress {
  creator, err := sdk.AccAddressFromBech32(msg.Creator)
  if err != nil {
    panic(err)
  }
  return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateVendorInfoType) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateVendorInfoType) ValidateBasic() error {
  _, err := sdk.AccAddressFromBech32(msg.Creator)
  if err != nil {
    return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
  }
   return nil
}

var _ sdk.Msg = &MsgDeleteVendorInfoType{}

func NewMsgDeleteVendorInfoType(
    creator string,
    index string,
    
) *MsgDeleteVendorInfoType {
  return &MsgDeleteVendorInfoType{
		Creator: creator,
		Index: index,
        
	}
}
func (msg *MsgDeleteVendorInfoType) Route() string {
  return RouterKey
}

func (msg *MsgDeleteVendorInfoType) Type() string {
  return "DeleteVendorInfoType"
}

func (msg *MsgDeleteVendorInfoType) GetSigners() []sdk.AccAddress {
  creator, err := sdk.AccAddressFromBech32(msg.Creator)
  if err != nil {
    panic(err)
  }
  return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteVendorInfoType) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteVendorInfoType) ValidateBasic() error {
  _, err := sdk.AccAddressFromBech32(msg.Creator)
  if err != nil {
    return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
  }
  return nil
}