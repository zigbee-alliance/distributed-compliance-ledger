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
	"time"

	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRejectAddAccount = "reject_add_account"

var _ sdk.Msg = &MsgRejectAddAccount{}

func NewMsgRejectAddAccount(signer sdk.AccAddress, address sdk.AccAddress, info string) *MsgRejectAddAccount {
	return &MsgRejectAddAccount{
		Signer:  signer.String(),
		Address: address.String(),
		Info:    info,
		Time:    time.Now().Unix(),
	}
}

func (msg *MsgRejectAddAccount) Route() string {
	return RouterKey
}

func (msg *MsgRejectAddAccount) Type() string {
	return TypeMsgRejectAddAccount
}

func (msg *MsgRejectAddAccount) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{signer}
}

func (msg *MsgRejectAddAccount) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)

	return sdk.MustSortJSON(bz)
}

func (msg *MsgRejectAddAccount) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signer address (%s)", err)
	}

	return nil
}
