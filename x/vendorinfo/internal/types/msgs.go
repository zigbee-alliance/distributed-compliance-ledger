// Copyright 2020 DSR Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const RouterKey = ModuleName

/*
	ADD_VENDOR Message
*/
type MsgAddVendorInfo struct {
	VendorInfo
	Signer sdk.AccAddress `json:"signer"`
}

func NewMsgAddVendorInfo(vendor VendorInfo, signer sdk.AccAddress) MsgAddVendorInfo {
	return MsgAddVendorInfo{
		VendorInfo: vendor,
		Signer:     signer,
	}
}

func (m MsgAddVendorInfo) Route() string {
	return RouterKey
}

func (m MsgAddVendorInfo) Type() string {
	return "add_vendor_info"
}

func (m MsgAddVendorInfo) ValidateBasic() sdk.Error {
	// TODO ADD LOGIC HERE
	/**
	if m.Address.Empty() {
		return sdk.ErrInvalidAddress("Invalid Account Address: it cannot be empty")
	}

	if len(m.PublicKey) == 0 {
		return sdk.ErrUnknownRequest("Invalid PublicKey: it cannot be empty")
	}

	if err := m.Roles.Validate(); err != nil {
		return err
	}

	if m.HasRole(Vendor) && m.VendorId <= 0 {
		return ErrMissingVendorIdForVendorAccount()
	}

	if m.Signer.Empty() {
		return sdk.ErrInvalidAddress("Invalid Signer: it cannot be empty")
	} **/

	return nil
}

func (m MsgAddVendorInfo) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(m))
}

func (m MsgAddVendorInfo) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Signer}
}

/*
	UPDATE_VENDOR Message
*/
type MsgUpdateVendorInfo struct {
	VendorInfo
	Signer sdk.AccAddress `json:"signer"`
}

func NewMsgUpdateVendorInfo(
	vendor VendorInfo,
	signer sdk.AccAddress,
) MsgUpdateVendorInfo {
	return MsgUpdateVendorInfo{
		VendorInfo: vendor,
		Signer:     signer,
	}
}

func (m MsgUpdateVendorInfo) Route() string {
	return RouterKey
}

func (m MsgUpdateVendorInfo) Type() string {
	return "update_vendor_info"
}

func (m MsgUpdateVendorInfo) ValidateBasic() sdk.Error {
	//TODO add validation
	/*
		if m.Signer.Empty() {
			return sdk.ErrInvalidAddress("Invalid Signer: it cannot be empty")
		}

		if m.VID == 0 {
			return sdk.ErrUnknownRequest("Invalid VID: it must be non-zero 16-bit unsigned integer")
		}

		if m.PID == 0 {
			return sdk.ErrUnknownRequest("Invalid PID: it must be non-zero 16-bit unsigned integer")
		}*/

	return nil
}

func (m MsgUpdateVendorInfo) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(m))
}

func (m MsgUpdateVendorInfo) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Signer}
}
