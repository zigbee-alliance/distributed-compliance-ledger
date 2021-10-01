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
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/validator"
)

const RouterKey = ModuleName

/*
	ADD_VENDOR Message
*/
type MsgAddVendorInfo struct {
	VendorInfo
	Signer sdk.AccAddress `json:"signer" validate:"required"`
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
	return validator.ValidateAdd(m)
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
	Signer sdk.AccAddress `json:"signer" validate:"required"`
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
	return validator.ValidateUpdate(m)
}

func (m MsgUpdateVendorInfo) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(m))
}

func (m MsgUpdateVendorInfo) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Signer}
}
