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

type MsgAddModel struct {
	Model
	Signer sdk.AccAddress `json:"signer"`
}

func NewMsgAddModel(
	model Model,
	signer sdk.AccAddress,
) MsgAddModel {
	return MsgAddModel{
		Model:  model,
		Signer: signer,
	}
}

func (m MsgAddModel) Route() string {
	return RouterKey
}

func (m MsgAddModel) Type() string {
	return "add_model_info"
}

func (m MsgAddModel) ValidateBasic() sdk.Error {
	if m.Signer.Empty() {
		return sdk.ErrInvalidAddress("Invalid Signer: it cannot be empty")
	}

	if m.VID == 0 {
		return sdk.ErrUnknownRequest("Invalid VID: it must be non-zero 16-bit unsigned integer")
	}

	if m.PID == 0 {
		return sdk.ErrUnknownRequest("Invalid PID: it must be non-zero 16-bit unsigned integer")
	}

	if len(m.ProductName) == 0 {
		return sdk.ErrUnknownRequest("Invalid ProductName: it cannot be empty")
	}

	if len(m.ProductLabel) == 0 {
		return sdk.ErrUnknownRequest("Invalid ProductLabel: it cannot be empty")
	}

	if len(m.PartNumber) == 0 {
		return sdk.ErrUnknownRequest("Invalid PartNumber: it cannot be empty")
	}
	return nil
}

func (m MsgAddModel) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(m))
}

func (m MsgAddModel) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Signer}
}

type MsgUpdateModel struct {
	Model
	Signer sdk.AccAddress `json:"signer"`
}

func NewMsgUpdateModel(
	model Model,
	signer sdk.AccAddress,
) MsgUpdateModel {
	return MsgUpdateModel{
		Model:  model,
		Signer: signer,
	}
}

func (m MsgUpdateModel) Route() string {
	return RouterKey
}

func (m MsgUpdateModel) Type() string {
	return "update_model_info"
}

func (m MsgUpdateModel) ValidateBasic() sdk.Error {
	if m.Signer.Empty() {
		return sdk.ErrInvalidAddress("Invalid Signer: it cannot be empty")
	}

	if m.VID == 0 {
		return sdk.ErrUnknownRequest("Invalid VID: it must be non-zero 16-bit unsigned integer")
	}

	if m.PID == 0 {
		return sdk.ErrUnknownRequest("Invalid PID: it must be non-zero 16-bit unsigned integer")
	}

	return nil
}

func (m MsgUpdateModel) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(m))
}

func (m MsgUpdateModel) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Signer}
}

type MsgDeleteModel struct {
	VID    uint16         `json:"vid"`
	PID    uint16         `json:"pid"`
	Signer sdk.AccAddress `json:"signer"`
}

func NewMsgDeleteModel(vid uint16, pid uint16, signer sdk.AccAddress) MsgDeleteModel {
	return MsgDeleteModel{
		VID:    vid,
		PID:    pid,
		Signer: signer,
	}
}

func (m MsgDeleteModel) Route() string {
	return RouterKey
}

func (m MsgDeleteModel) Type() string {
	return "delete_model_info"
}

func (m MsgDeleteModel) ValidateBasic() sdk.Error {
	if m.Signer.Empty() {
		return sdk.ErrInvalidAddress("Invalid Signer: it cannot be empty")
	}

	if m.VID == 0 {
		return sdk.ErrUnknownRequest("Invalid VID: it must be non-zero 16-bit unsigned integer")
	}

	if m.PID == 0 {
		return sdk.ErrUnknownRequest("Invalid PID: it must be non-zero 16-bit unsigned integer")
	}

	return nil
}

func (m MsgDeleteModel) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(m))
}

func (m MsgDeleteModel) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Signer}
}
