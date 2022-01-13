// Copyright 2022 DSR Corporation
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
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/validator"
)

const TypeMsgApproveRevokeX509RootCert = "approve_revoke_x_509_root_cert"

var _ sdk.Msg = &MsgApproveRevokeX509RootCert{}

func NewMsgApproveRevokeX509RootCert(signer string, subject string, subjectKeyId string) *MsgApproveRevokeX509RootCert {
	return &MsgApproveRevokeX509RootCert{
		Signer:       signer,
		Subject:      subject,
		SubjectKeyId: subjectKeyId,
	}
}

func (msg *MsgApproveRevokeX509RootCert) Route() string {
	return RouterKey
}

func (msg *MsgApproveRevokeX509RootCert) Type() string {
	return TypeMsgApproveRevokeX509RootCert
}

func (msg *MsgApproveRevokeX509RootCert) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{signer}
}

func (msg *MsgApproveRevokeX509RootCert) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgApproveRevokeX509RootCert) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signer address (%s)", err)
	}

	err = validator.Validate(msg)
	if err != nil {
		return err
	}

	return nil
}
