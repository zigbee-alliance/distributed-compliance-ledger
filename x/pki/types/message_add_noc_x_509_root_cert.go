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
	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/validator"
)

const TypeMsgAddNocX509RootCert = "add_noc_x_509_root_cert"

var _ sdk.Msg = &MsgAddNocX509RootCert{}

func NewMsgAddNocX509RootCert(signer string, cert string, certSchemaVersion uint32, isVidVerificationSigner bool) *MsgAddNocX509RootCert {
	return &MsgAddNocX509RootCert{
		Signer:                  signer,
		Cert:                    cert,
		CertSchemaVersion:       certSchemaVersion,
		IsVidVerificationSigner: isVidVerificationSigner,
	}
}

func (msg *MsgAddNocX509RootCert) Route() string {
	return pkitypes.RouterKey
}

func (msg *MsgAddNocX509RootCert) Type() string {
	return TypeMsgAddNocX509RootCert
}

func (msg *MsgAddNocX509RootCert) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{signer}
}

func (msg *MsgAddNocX509RootCert) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)

	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddNocX509RootCert) ValidateBasic() error {
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
