package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/validator"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/x509"
)

const TypeMsgAddPkiRevocationDistributionPoint = "add_pki_revocation_distribution_point"

var _ sdk.Msg = &MsgAddPkiRevocationDistributionPoint{}

func NewMsgAddPkiRevocationDistributionPoint(signer string, vid int32, pid int32, isPAA bool, label string, crlSignerCertificate string, issuerSubjectKeyID string, dataURL string, dataFileSize uint64, dataDigest string, dataDigestType uint32, revocationType uint32) *MsgAddPkiRevocationDistributionPoint {
	return &MsgAddPkiRevocationDistributionPoint{
		Signer:               signer,
		Vid:                  vid,
		Pid:                  pid,
		IsPAA:                isPAA,
		Label:                label,
		CrlSignerCertificate: crlSignerCertificate,
		IssuerSubjectKeyID:   issuerSubjectKeyID,
		DataURL:              dataURL,
		DataFileSize:         dataFileSize,
		DataDigest:           dataDigest,
		DataDigestType:       dataDigestType,
		RevocationType:       revocationType,
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

func (msg *MsgAddPkiRevocationDistributionPoint) verifyVid(subjectAsText string) error {
	vid, err := x509.GetVidFromSubject(subjectAsText)

	if err != nil {
		return sdkerrors.Wrapf(pkitypes.ErrInvalidVidFormat, "Could not parse vid: %s", err)
	}

	if vid == 0 {
		return sdkerrors.Wrap(pkitypes.ErrUnsupportedOperation, "publishing a revocation point for non-VID scoped root certificates is currently not supported")
	}

	if vid != msg.Vid {
		return pkitypes.NewErrCRLSignerCertificateVidNotEqualMsgVid("CRL Signer Certificate's vid must be equal to the provided vid in the message")
	}

	return nil
}

func (msg *MsgAddPkiRevocationDistributionPoint) verifyPid(subjectAsText string) error {
	pid, err := x509.GetPidFromSubject(subjectAsText)

	if err != nil {
		return sdkerrors.Wrapf(pkitypes.ErrInvalidPidFormat, "Could not parse pid: %s", err)
	}
	if pid == 0 && msg.Pid != 0 {
		return pkitypes.NewErrNotEmptyPid("Product ID (pid) must be empty when it is not found in non-root CRL Signer Certificate")
	}

	if pid != 0 && msg.Pid == 0 {
		return pkitypes.NewErrPidNotFound("Product ID (pid) must be provided when it is found in non-root CRL Signer Certificate")
	}

	if pid != msg.Pid {
		return pkitypes.NewErrCRLSignerCertificatePidNotEqualMsgPid("CRL Signer Certificate's pid must be equal to the provided pid in the message")
	}

	return nil
}

func (msg *MsgAddPkiRevocationDistributionPoint) verifyPAA(cert *x509.Certificate) error {
	if msg.Pid != 0 {
		return pkitypes.NewErrNotEmptyPid("Product ID (pid) must be empty for root certificates when isPAA is true")
	}

	if !cert.IsSelfSigned() {
		return pkitypes.NewErrRootCertificateIsNotSelfSigned("CRL Signer Certificate must be self-signed if isPAA is True")
	}

	err := msg.verifyVid(cert.SubjectAsText)

	if err != nil {
		return err
	}

	return nil
}

func (msg *MsgAddPkiRevocationDistributionPoint) verifyPAI(cert *x509.Certificate) error {
	err := msg.verifyPid(cert.SubjectAsText)

	if err != nil {
		return err
	}

	if cert.IsSelfSigned() {
		return pkitypes.NewErrNonRootCertificateSelfSigned("CRL Signer Certificate shall not be self-signed if isPAA is False")
	}

	err = msg.verifyVid(cert.SubjectAsText)

	if err != nil {
		return err
	}

	return nil
}

func (msg *MsgAddPkiRevocationDistributionPoint) ValidateBasic() error {
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

	if !isDataDigestInTypes {
		return pkitypes.NewErrInvalidDataDigestType(fmt.Sprintf("invalid DataDigestType: %d. Supported types are: %v", msg.DataDigestType, allowedDataDigestTypes))
	}

	if msg.RevocationType != AllowedRevocationType {
		return pkitypes.NewErrInvalidRevocationType(fmt.Sprintf("invalid RevocationType: %d. Supported types are: %d", msg.RevocationType, AllowedRevocationType))
	}

	cert, err := x509.DecodeX509Certificate(msg.CrlSignerCertificate)
	if err != nil {
		return pkitypes.NewErrInvalidCertificate(err)
	}

	if msg.IsPAA {
		err = msg.verifyPAA(cert)

		if err != nil {
			return err
		}
	} else {
		err = msg.verifyPAI(cert)

		if err != nil {
			return err
		}
	}

	if !strings.HasPrefix(msg.DataURL, "https://") && !strings.HasPrefix(msg.DataURL, "http://") {
		return pkitypes.NewErrInvalidDataURLFormat("Data Url must start with https:// or http://")
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

	if msg.RevocationType == AllowedRevocationType && (msg.DataFileSize != 0 || msg.DataDigest != "" || msg.DataDigestType != 0) {
		return pkitypes.NewErrDataFieldPresented(fmt.Sprintf("Data Digest, Data File Size and Data Digest Type must be omitted for Revocation Type %d", AllowedRevocationType))
	}

	match := VerifyRevocationPointIssuerSubjectKeyIDFormat(msg.IssuerSubjectKeyID)

	if !match {
		return pkitypes.NewErrWrongSubjectKeyIDFormat("Wrong IssuerSubjectKeyID format. It must consist of even number of uppercase hexadecimal characters ([0-9A-F]), with no whitespace and no non-hexadecimal characters")
	}

	return nil
}
