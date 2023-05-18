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

func NewMsgAddPkiRevocationDistributionPoint(signer string, vid int32, pid int32, isPAA bool, label string, crlSignerCertificate string, issuerSubjectKeyID string, dataUrl string, dataFileSize uint64, dataDigest string, dataDigestType uint32, revocationType uint64) *MsgAddPkiRevocationDistributionPoint {
	return &MsgAddPkiRevocationDistributionPoint{
		Signer:               signer,
		Vid:                  vid,
		Pid:                  pid,
		IsPAA:                isPAA,
		Label:                label,
		CrlSignerCertificate: crlSignerCertificate,
		IssuerSubjectKeyID:   issuerSubjectKeyID,
		DataUrl:              dataUrl,
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

func (msg *MsgAddPkiRevocationDistributionPoint) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signer address (%s)", err)
	}

	err = validator.Validate(msg)
	if err != nil {
		return err
	}

	allowedDataDigestTypes := [6]uint32{1, 7, 8, 10, 11, 12}
	allowedRevocationType := uint64(1)

	isDataDigestInTypes := false
	for _, digestType := range allowedDataDigestTypes {
		if digestType == msg.DataDigestType {
			isDataDigestInTypes = true

			break
		}
	}

	if !isDataDigestInTypes {
		return pkitypes.NewErrInvalidDataDigestType(fmt.Sprintf("invalid DataDigestType: %d", msg.DataDigestType))
	}

	if msg.RevocationType != allowedRevocationType {
		return pkitypes.NewErrInvalidRevocationType(fmt.Sprintf("invalid RevocationType: %d", msg.RevocationType))
	}

	cert, err := x509.DecodeX509Certificate(msg.CrlSignerCertificate)
	if err != nil {
		return err
	}

	if msg.IsPAA {
		if msg.Pid != 0 {
			return pkitypes.NewErrNotEmptyPid("Product ID (pid) must be empty for root certificates when isPAA is true")
		}

		if !cert.IsSelfSigned() {
			return pkitypes.NewErrPAANotSelfSigned(fmt.Sprintf("CRL Signer Certificate must be self-signed if isPAA is True")
		}
	} else {
		if !strings.Contains(cert.SubjectAsText, string(msg.Pid)) {
			return pkitypes.NewErrCRLSignerCertificateDoesNotContainPid(fmt.Sprintf("CRLSignerCertificate with subject: %s, subjectKeyID does not contain pid: %d", cert.SubjectAsText, cert.SubjectKeyID, msg.Pid))
		}

		if cert.IsSelfSigned() {
			return pkitypes.NewErrNonPAASelfSigned(fmt.Sprintf("non-root certificate with subject: %s, subjectKeyID: %s is self-signed", cert.SubjectAsText, cert.SubjectKeyID))
		}
	}

	if msg.DataFileSize == 0 && msg.DataDigest != "" {
		return pkitypes.NewErrEmptyDataFileSize(fmt.Sprintf("msgAddRevocationDistributionPoint with CRLSignerCertificate: %s has empty DataFileSize when DataDigest is not empty", msg.CrlSignerCertificate))
	}

	if msg.DataFileSize != 0 && msg.DataDigest == "" {
		return pkitypes.NewErrEmptyDataDigest(fmt.Sprintf("msgAddRevocationDistributionPoint with CRLSignerCertificate: %s has empty DataDigest when DataFileSize is not empty", msg.CrlSignerCertificate))
	}

	if msg.DataDigest == "" && msg.DataDigestType != 0 {
                return pkitypes.NewErrNonEmptyDataDigestType("Data Digest Type must be provided only if Data Digest is provided")
	}

	if msg.DataDigest != "" && msg.DataDigestType == 0 {
		return pkitypes.NewErrEmptyDataDigestType("Data Digest Type must be provided if Data Digest is provided")
	}

	if msg.RevocationType == 1 && (msg.DataFileSize != 0 || msg.DataDigest != "" || msg.DataDigestType != 0) {
		return pkitypes.NewErrDataFieldPresented(fmt.Sprintf("msgAddRevocationDistributionPoint with CRLSignerCertificate: %s has one or more non-empty DataFields when RevocationType is 1", msg.CrlSignerCertificate))
	}

	if msg.IssuerSubjectKeyID != cert.SubjectKeyID {
		return pkitypes.NewErrIssuerSubjectKeyIDNotEqualsCertSubjectKeyID(fmt.Sprintf("msgAddRevocationDistributionPoint with CRLSignerCertificate: %s has IssuerSubjectKeyID: %s which is not equal to certificate SubjectKeyID: %s", msg.CrlSignerCertificate, msg.IssuerSubjectKeyID, cert.SubjectKeyID))
	}

	return nil
}
