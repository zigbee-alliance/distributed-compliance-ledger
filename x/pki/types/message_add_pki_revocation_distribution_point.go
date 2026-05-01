package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/validator"
)

const TypeMsgAddPkiRevocationDistributionPoint = "add_pki_revocation_distribution_point"

var _ sdk.Msg = &MsgAddPkiRevocationDistributionPoint{}

func NewMsgAddPkiRevocationDistributionPoint(signer string, vid int32, pid int32, isPAA bool, label string,
	crlSignerCertificate string, crlSignerDelegator string, issuerSubjectKeyID string, dataURL string, dataFileSize uint64, dataDigest string,
	dataDigestType uint32, revocationType uint32, schemaVersion uint32) *MsgAddPkiRevocationDistributionPoint {
	return &MsgAddPkiRevocationDistributionPoint{
		Signer:               signer,
		Vid:                  vid,
		Pid:                  pid,
		IsPAA:                isPAA,
		Label:                label,
		CrlSignerCertificate: crlSignerCertificate,
		CrlSignerDelegator:   crlSignerDelegator,
		IssuerSubjectKeyID:   issuerSubjectKeyID,
		DataURL:              dataURL,
		DataFileSize:         dataFileSize,
		DataDigest:           dataDigest,
		DataDigestType:       dataDigestType,
		RevocationType:       revocationType,
		SchemaVersion:        schemaVersion,
	}
}

func (msg *MsgAddPkiRevocationDistributionPoint) Route() string {
	return pkitypes.RouterKey
}

func (msg *MsgAddPkiRevocationDistributionPoint) Type() string {
	return TypeMsgAddPkiRevocationDistributionPoint
}

func (msg *MsgAddPkiRevocationDistributionPoint) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{signer}
}

func (msg *MsgAddPkiRevocationDistributionPoint) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)

	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddPkiRevocationDistributionPoint) verifyFields() error {
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

	if !isDataDigestInTypes {
		return pkitypes.NewErrInvalidDataDigestType(msg.DataDigestType, allowedDataDigestTypes[:])
	}

	isRevocationInTypes := false
	for _, revocationType := range allowedRevocationTypes {
		if revocationType == msg.RevocationType {
			isRevocationInTypes = true

			break
		}
	}

	if !isRevocationInTypes {
		return pkitypes.NewErrInvalidRevocationType(msg.RevocationType, allowedRevocationTypes[:])
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

	if msg.RevocationType == CRLRevocationType && (msg.DataFileSize != 0 || msg.DataDigest != "" || msg.DataDigestType != 0) {
		return pkitypes.NewErrDataFieldPresented(CRLRevocationType)
	}

	match := VerifyRevocationPointIssuerSubjectKeyIDFormat(msg.IssuerSubjectKeyID)

	if !match {
		return pkitypes.NewErrWrongIssuerSubjectKeyIDFormat()
	}

	return nil
}

func (msg *MsgAddPkiRevocationDistributionPoint) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return pkitypes.NewErrInvalidAddress(err)
	}

	err = validator.Validate(msg)
	if err != nil {
		return err
	}

	if err = msg.verifyFields(); err != nil {
		return err
	}

	return nil
}
