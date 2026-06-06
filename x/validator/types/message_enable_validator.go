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
	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/validator"
)

const TypeMsgEnableValidator = "enable_validator"

var _ sdk.Msg = &MsgEnableValidator{}

func NewMsgEnableValidator(creator sdk.ValAddress) *MsgEnableValidator {
	return &MsgEnableValidator{
		Creator: creator.String(),
	}
}

func (msg *MsgEnableValidator) Route() string {
	return RouterKey
}

func (msg *MsgEnableValidator) Type() string {
	return TypeMsgEnableValidator
}

func (msg *MsgEnableValidator) GetSigners() []sdk.AccAddress {
	creator, err := sdk.ValAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{sdk.AccAddress(creator)}
}

func (msg *MsgEnableValidator) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)

	return sdk.MustSortJSON(bz)
}

func (msg *MsgEnableValidator) ValidateBasic() error {
	_, err := sdk.ValAddressFromBech32(msg.Creator)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid validator address (%s)", err)
	}

	err = validator.Validate(msg)
	if err != nil {
		return err
	}

	return nil
}
