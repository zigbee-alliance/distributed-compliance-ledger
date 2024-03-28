package types

import (
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/validator"
)

const TypeMsgUpdatePkiRevocationDistributionPoint = "update_pki_revocation_distribution_point"

var _ sdk.Msg = &MsgUpdatePkiRevocationDistributionPoint{}

func NewMsgUpdatePkiRevocationDistributionPoint(signer string, vid int32, label string, crlSignerCertificate string,
	crlSignerDelegator string, issuerSubjectKeyID string, dataURL string, dataFileSize uint64, dataDigest string,
	dataDigestType uint32, schemaVersion uint32) *MsgUpdatePkiRevocationDistributionPoint {
	return &MsgUpdatePkiRevocationDistributionPoint{
		Signer:               signer,
		Vid:                  vid,
		Label:                label,
		CrlSignerCertificate: crlSignerCertificate,
		CrlSignerDelegator:   crlSignerDelegator,
		IssuerSubjectKeyID:   issuerSubjectKeyID,
		DataURL:              dataURL,
		DataFileSize:         dataFileSize,
		DataDigest:           dataDigest,
		DataDigestType:       dataDigestType,
		SchemaVersion:        schemaVersion,
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
		return pkitypes.NewErrInvalidAddress(err)
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

	if msg.DataURL != "" && !strings.HasPrefix(msg.DataURL, "https://") && !strings.HasPrefix(msg.DataURL, "http://") {
		return pkitypes.NewErrInvalidDataURLSchema()
	}

	if !isDataDigestInTypes {
		return pkitypes.NewErrInvalidDataDigestType(msg.DataDigestType, allowedDataDigestTypes[:])
	}

	if msg.DataFileSize == 0 && msg.DataDigest != "" {
		return pkitypes.NewErrNonEmptyDataDigest()
	}

	if msg.DataFileSize != 0 && msg.DataDigest == "" {
		return pkitypes.NewErrEmptyDataDigest()
	}

	if msg.DataDigest == "" && msg.DataDigestType != 0 {
		return pkitypes.NewErrNotEmptyDataDigestType()
	}

	if msg.DataDigest != "" && msg.DataDigestType == 0 {
		return pkitypes.NewErrEmptyDataDigestType()
	}

	match := VerifyRevocationPointIssuerSubjectKeyIDFormat(msg.IssuerSubjectKeyID)

	if !match {
		return pkitypes.NewErrWrongSubjectKeyIDFormat()
	}

	return nil
}
