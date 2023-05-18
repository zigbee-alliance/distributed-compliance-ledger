package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/validator"
)

const TypeMsgUpdatePkiRevocationDistributionPoint = "update_pki_revocation_distribution_point"

var _ sdk.Msg = &MsgUpdatePkiRevocationDistributionPoint{}

func NewMsgUpdatePkiRevocationDistributionPoint(signer string, vid int32, label string, crlSignerCertificate string, issuerSubjectKeyID string, dataUrl string, dataFileSize uint64, dataDigest string, dataDigestType uint32) *MsgUpdatePkiRevocationDistributionPoint {
	return &MsgUpdatePkiRevocationDistributionPoint{
		Signer:               signer,
		Vid:                  vid,
		Label:                label,
		CrlSignerCertificate: crlSignerCertificate,
		IssuerSubjectKeyID:   issuerSubjectKeyID,
		DataUrl:              dataUrl,
		DataFileSize:         dataFileSize,
		DataDigest:           dataDigest,
		DataDigestType:       dataDigestType,
	}
}

func (msg *MsgUpdatePkiRevocationDistributionPoint) Route() string {
	return pkitypes.RouterKey
}

func (msg *MsgUpdatePkiRevocationDistributionPoint) Type() string {
	return TypeMsgUpdatePkiRevocationDistributionPoint
}

func (msg *MsgUpdatePkiRevocationDistributionPoint) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{signer}
}

func (msg *MsgUpdatePkiRevocationDistributionPoint) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdatePkiRevocationDistributionPoint) ValidateBasic() error {
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
