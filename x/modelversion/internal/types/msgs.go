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

type MsgAddModelVersion struct {
	ModelVersion
	Signer sdk.AccAddress `json:"signer" validate:"required"`
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
	Signer sdk.AccAddress `json:"signer"`
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
