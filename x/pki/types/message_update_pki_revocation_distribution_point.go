package types

import (
	fmt "fmt"
	"strings"

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

	isDataDigestInTypes := true
	if msg.DataDigestType != 0 {
		isDataDigestInTypes = false
		for _, digestType := range allowedDataDigestTypes {
			if digestType == msg.DataDigestType {
				isDataDigestInTypes = true

				break
			}
		}
	}

	if msg.DataUrl != "" && !strings.HasPrefix(msg.DataUrl, "https://") && !strings.HasPrefix(msg.DataUrl, "http://") {
		return pkitypes.NewErrInvalidDataURLFormat("Data Url must start with https:// or http://")
	}

	if !isDataDigestInTypes {
		return pkitypes.NewErrInvalidDataDigestType(fmt.Sprintf("invalid DataDigestType: %d. Supported types are: %v", msg.DataDigestType, allowedDataDigestTypes))
	}

	if msg.DataFileSize == 0 && msg.DataDigest != "" {
		return pkitypes.NewErrNonEmptyDataDigest("Data Digest must be provided only if Data File Size is provided")
	}

	if msg.DataFileSize != 0 && msg.DataDigest == "" {
		return pkitypes.NewErrEmptyDataDigest("Data Digest must be provided if Data File Size is provided")
	}

	if msg.DataDigest == "" && msg.DataDigestType != 0 {
		return pkitypes.NewErrNotEmptyDataDigestType("Data Digest Type must be provided only if Data Digest is provided")
	}

	if msg.DataDigest != "" && msg.DataDigestType == 0 {
		return pkitypes.NewErrEmptyDataDigestType("Data Digest Type must be provided if Data Digest is provided")
	}

	match := VerifyRevocationPointIssuerSubjectKeyIDFormat(msg.IssuerSubjectKeyID)

	if !match {
		return pkitypes.NewErrWrongSubjectKeyIDFormat("Wrong IssuerSubjectKeyID format. It must consist of even number of uppercase hexadecimal characters ([0-9A-F]), with no whitespace and no non-hexadecimal characters")
	}

	return nil
}
