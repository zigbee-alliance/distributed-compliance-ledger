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

import sdk "github.com/cosmos/cosmos-sdk/types"

const RouterKey = ModuleName

/*
	PROPOSE_ADD_ACCOUNT Message
*/
type MsgProposeAddAccount struct {
	Address   sdk.AccAddress `json:"address"`
	PublicKey string         `json:"pub_key"`
	Roles     AccountRoles   `json:"roles"`
	VendorId  uint16         `json:"vendorId"`
	Signer    sdk.AccAddress `json:"signer"`
}

func NewMsgProposeAddAccount(address sdk.AccAddress, pubKey string,
	roles AccountRoles, vendorId uint16, signer sdk.AccAddress) MsgProposeAddAccount {
	return MsgProposeAddAccount{
		Address:   address,
		PublicKey: pubKey,
		Roles:     roles,
		VendorId:  vendorId,
		Signer:    signer,
	}
}

func (m MsgProposeAddAccount) Route() string {
	return RouterKey
}

func (m MsgProposeAddAccount) Type() string {
	return "propose_add_account"
}

func (m MsgProposeAddAccount) ValidateBasic() sdk.Error {
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
	}

	return nil
}

func (m MsgProposeAddAccount) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(m))
}

func (m MsgProposeAddAccount) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Signer}
}

func (m MsgProposeAddAccount) HasRole(targetRole AccountRole) bool {
	for _, role := range m.Roles {
		if role == targetRole {
			return true
		}
	}
	return false
}

/*
	APPROVE_ADD_ACCOUNT Message
*/
type MsgApproveAddAccount struct {
	Address sdk.AccAddress `json:"address"`
	Signer  sdk.AccAddress `json:"signer"`
}

func NewMsgApproveAddAccount(address sdk.AccAddress, signer sdk.AccAddress) MsgApproveAddAccount {
	return MsgApproveAddAccount{
		Address: address,
		Signer:  signer,
	}
}

func (m MsgApproveAddAccount) Route() string {
	return RouterKey
}

func (m MsgApproveAddAccount) Type() string {
	return "approve_add_account"
}

func (m MsgApproveAddAccount) ValidateBasic() sdk.Error {
	if m.Address.Empty() {
		return sdk.ErrInvalidAddress("Invalid Account Address: it cannot be empty")
	}

	if m.Signer.Empty() {
		return sdk.ErrInvalidAddress("Invalid Signer: it cannot be empty")
	}

	return nil
}

func (m MsgApproveAddAccount) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(m))
}

func (m MsgApproveAddAccount) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Signer}
}

/*
	PROPOSE_REVOKE_ACCOUNT Message
*/
type MsgProposeRevokeAccount struct {
	Address sdk.AccAddress `json:"address"`
	Signer  sdk.AccAddress `json:"signer"`
}

func NewMsgProposeRevokeAccount(address sdk.AccAddress, signer sdk.AccAddress) MsgProposeRevokeAccount {
	return MsgProposeRevokeAccount{
		Address: address,
		Signer:  signer,
	}
}

func (m MsgProposeRevokeAccount) Route() string {
	return RouterKey
}

func (m MsgProposeRevokeAccount) Type() string {
	return "propose_revoke_account"
}

func (m MsgProposeRevokeAccount) ValidateBasic() sdk.Error {
	if m.Address.Empty() {
		return sdk.ErrInvalidAddress("Invalid Account Address: it cannot be empty")
	}

	if m.Signer.Empty() {
		return sdk.ErrInvalidAddress("Invalid Signer: it cannot be empty")
	}

	return nil
}

func (m MsgProposeRevokeAccount) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(m))
}

func (m MsgProposeRevokeAccount) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Signer}
}

/*
	APPROVE_REVOKE_ACCOUNT Message
*/
type MsgApproveRevokeAccount struct {
	Address sdk.AccAddress `json:"address"`
	Signer  sdk.AccAddress `json:"signer"`
}

func NewMsgApproveRevokeAccount(address sdk.AccAddress, signer sdk.AccAddress) MsgApproveRevokeAccount {
	return MsgApproveRevokeAccount{
		Address: address,
		Signer:  signer,
	}
}

func (m MsgApproveRevokeAccount) Route() string {
	return RouterKey
}

func (m MsgApproveRevokeAccount) Type() string {
	return "approve_revoke_account"
}

func (m MsgApproveRevokeAccount) ValidateBasic() sdk.Error {
	if m.Address.Empty() {
		return sdk.ErrInvalidAddress("Invalid Account Address: it cannot be empty")
	}

	if m.Signer.Empty() {
		return sdk.ErrInvalidAddress("Invalid Signer: it cannot be empty")
	}

	return nil
}

func (m MsgApproveRevokeAccount) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(m))
}

func (m MsgApproveRevokeAccount) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Signer}
}
