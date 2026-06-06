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
	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/validator"
)

const TypeMsgRevokeNocX509IcaCert = "revoke_noc_x_509_ica_cert"

var _ sdk.Msg = &MsgRevokeNocX509IcaCert{}

func NewMsgRevokeNocX509IcaCert(signer, subject, subjectKeyID, serialNumber, info string, revokeChild bool) *MsgRevokeNocX509IcaCert {
	return &MsgRevokeNocX509IcaCert{
		Signer:       signer,
		Subject:      subject,
		SubjectKeyId: subjectKeyID,
		SerialNumber: serialNumber,
		Info:         info,
		Time:         time.Now().Unix(),
		RevokeChild:  revokeChild,
	}
}

func (msg *MsgRevokeNocX509IcaCert) Route() string {
	return pkitypes.RouterKey
}

func (msg *MsgRevokeNocX509IcaCert) Type() string {
	return TypeMsgRevokeNocX509IcaCert
}

func (msg *MsgRevokeNocX509IcaCert) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{signer}
}

func (msg *MsgRevokeNocX509IcaCert) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)

	return sdk.MustSortJSON(bz)
}

func (msg *MsgRevokeNocX509IcaCert) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signer address (%s)", err)
	}

	err = validator.Validate(msg)
	if err != nil {
		return err
	}

	return nil
}
