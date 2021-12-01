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

type MsgAddModel struct {
	Model
	Signer sdk.AccAddress `json:"signer" validate:"required,address"`
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
	err := validator.ValidateAdd(m)
	if err != nil {
		return err
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
	Signer sdk.AccAddress `json:"signer" validate:"required,address"`
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
	err := validator.ValidateUpdate(m)
	if err != nil {
		return err
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
	VID    uint16         `json:"vid" validate:"required"`
	PID    uint16         `json:"pid" validate:"required"`
	Signer sdk.AccAddress `json:"signer" validate:"required,address"`
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
	err := validator.ValidateUpdate(m)
	if err != nil {
		return err
	}

	return nil
}

func (m MsgDeleteModel) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(m))
}

func (m MsgDeleteModel) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Signer}
}

// Model Versions

type MsgAddModelVersion struct {
	ModelVersion
	Signer sdk.AccAddress `json:"signer" validate:"required,address"`
}

func NewMsgAddModelVersion(
	modelVersion ModelVersion,
	signer sdk.AccAddress,
) MsgAddModelVersion {
	return MsgAddModelVersion{
		ModelVersion: modelVersion,
		Signer:       signer,
	}
}

func (m MsgAddModelVersion) Route() string {
	return RouterKey
}

func (m MsgAddModelVersion) Type() string {
	return "add_model_version"
}

func (m MsgAddModelVersion) ValidateBasic() sdk.Error {
	err := validator.ValidateAdd(m)
	if err != nil {
		return err
	}

	return nil
}

func (m MsgAddModelVersion) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(m))
}

func (m MsgAddModelVersion) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Signer}
}

type MsgUpdateModelVersion struct {
	ModelVersion
	Signer sdk.AccAddress `json:"signer" validate:"required,address"`
}

func NewMsgUpdateModelVersion(
	modelVersion ModelVersion,
	signer sdk.AccAddress,
) MsgUpdateModelVersion {
	return MsgUpdateModelVersion{
		ModelVersion: modelVersion,
		Signer:       signer,
	}
}

func (m MsgUpdateModelVersion) Route() string {
	return RouterKey
}

func (m MsgUpdateModelVersion) Type() string {
	return "update_model_version"
}

func (m MsgUpdateModelVersion) ValidateBasic() sdk.Error {
	err := validator.ValidateUpdate(m)
	if err != nil {
		return err
	}

	return nil
}

func (m MsgUpdateModelVersion) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(m))
}

func (m MsgUpdateModelVersion) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Signer}
}
