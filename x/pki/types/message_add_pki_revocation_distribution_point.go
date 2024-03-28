package types

import (
	x509std "crypto/x509"
	"encoding/asn1"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"

	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/validator"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/x509"
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

func (msg *MsgAddPkiRevocationDistributionPoint) verifyPAA(cert *x509.Certificate) error {
	if msg.Pid != 0 {
		return pkitypes.NewErrNotEmptyPidForRootCertificate()
	}

	pid, _ := x509.GetPidFromSubject(cert.SubjectAsText)
	if pid != 0 {
		return pkitypes.NewErrNotEmptyPidForNonRootCertificate()
	}

	// verify VID
	vid, err := x509.GetVidFromSubject(cert.SubjectAsText)
	if err != nil {
		return pkitypes.NewErrInvalidVidFormat(err)
	}

	if vid > 0 && vid != msg.Vid {
		return pkitypes.NewErrCRLSignerCertificateVidNotEqualMsgVid(vid, msg.Vid)
	}

	if !cert.IsSelfSigned() {
		err = VerifyCRLSignerCertFormat(cert)
		if err != nil {
			return err
		}
	}

	return nil
}

func (msg *MsgAddPkiRevocationDistributionPoint) verifyPAI(cert *x509.Certificate) error {
	if cert.IsSelfSigned() {
		return pkitypes.NewErrNonRootCertificateSelfSigned()
	}

	// verify VID
	vid, err := x509.GetVidFromSubject(cert.SubjectAsText)
	if err != nil {
		return pkitypes.NewErrInvalidVidFormat(err)
	}

	if vid != msg.Vid {
		return pkitypes.NewErrCRLSignerCertificateVidNotEqualMsgVid(vid, msg.Vid)
	}

	// verify PID
	pid, err := x509.GetPidFromSubject(cert.SubjectAsText)
	if err != nil {
		return pkitypes.NewErrInvalidPidFormat(err)
	}
	if pid == 0 && msg.Pid != 0 {
		return pkitypes.NewErrNotEmptyPidForNonRootCertificate()
	}
	if pid != 0 && msg.Pid == 0 {
		return pkitypes.NewErrPidNotFoundInMessage(pid)
	}
	if pid != msg.Pid {
		return pkitypes.NewErrCRLSignerCertificatePidNotEqualMsgPid(pid, msg.Pid)
	}

	if msg.CrlSignerDelegator != "" {
		err = VerifyCRLSignerCertFormat(cert)
		if err != nil {
			return err
		}
	}

	return nil
}

func VerifyCRLSignerCertFormat(certificate *x509.Certificate) error {
	if certificate.SubjectKeyID == "" {
		return pkitypes.NewErrWrongSubjectKeyIDFormat()
	}

	cert := certificate.Certificate
	if cert.Version != 3 {
		return pkitypes.NewErrCRLSignerCertificateInvalidFormat()
	}

	if cert.SignatureAlgorithm != x509std.ECDSAWithSHA256 {
		return pkitypes.NewErrCRLSignerCertificateInvalidFormat()
	}

	if cert.PublicKeyAlgorithm != x509std.ECDSA {
		return pkitypes.NewErrCRLSignerCertificateInvalidFormat()
	}
	// Basic Constraint extension should be marked critical and have the cA field set to false
	if !cert.BasicConstraintsValid || cert.IsCA {
		return pkitypes.NewErrCRLSignerCertificateInvalidFormat()
	}

	// Key Usage extension should be marked critical
	oidKeyUsage := asn1.ObjectIdentifier{2, 5, 29, 15}
	isCritical := false
	for _, ext := range cert.Extensions {
		if ext.Id.Equal(oidKeyUsage) {
			isCritical = ext.Critical

			break
		}
	}
	if !isCritical {
		return pkitypes.NewErrCRLSignerCertificateInvalidFormat()
	}

	if cert.KeyUsage&x509std.KeyUsageCRLSign == 0 {
		return pkitypes.NewErrCRLSignerCertificateInvalidFormat()
	}
	if cert.KeyUsage&^(x509std.KeyUsageCRLSign|x509std.KeyUsageDigitalSignature) != 0 {
		return pkitypes.NewErrCRLSignerCertificateInvalidFormat()
	}

	return nil
}

func (msg *MsgAddPkiRevocationDistributionPoint) verifySignerCertificate() error {
	cert, err := x509.DecodeX509Certificate(msg.CrlSignerCertificate)
	if err != nil {
		return pkitypes.NewErrInvalidCertificate(err)
	}

	if msg.IsPAA {
		err = msg.verifyPAA(cert)
	} else {
		err = msg.verifyPAI(cert)
	}

	return err
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

	if !strings.HasPrefix(msg.DataURL, "https://") && !strings.HasPrefix(msg.DataURL, "http://") {
		return pkitypes.NewErrInvalidDataURLSchema()
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
		return pkitypes.NewErrWrongSubjectKeyIDFormat()
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

	if err = msg.verifySignerCertificate(); err != nil {
		return err
	}

	return nil
}
