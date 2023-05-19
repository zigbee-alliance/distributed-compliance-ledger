package types

import (
	"fmt"
	"regexp"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/validator"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/x509"
)

const TypeMsgAddPkiRevocationDistributionPoint = "add_pki_revocation_distribution_point"

var _ sdk.Msg = &MsgAddPkiRevocationDistributionPoint{}

func NewMsgAddPkiRevocationDistributionPoint(signer string, vid int32, pid int32, isPAA bool, label string, crlSignerCertificate string, issuerSubjectKeyID string, dataUrl string, dataFileSize uint64, dataDigest string, dataDigestType uint32, revocationType uint32) *MsgAddPkiRevocationDistributionPoint {
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

	isDataDigestInTypes := false
	for _, digestType := range allowedDataDigestTypes {
		if digestType == msg.DataDigestType {
			isDataDigestInTypes = true

			break
		}
	}

	if !isDataDigestInTypes {
		return pkitypes.NewErrInvalidDataDigestType(fmt.Sprintf("invalid DataDigestType: %d. Supported types are: %v", msg.DataDigestType, allowedDataDigestTypes))
	}

	if msg.RevocationType != allowedRevocationType {
		return pkitypes.NewErrInvalidRevocationType(fmt.Sprintf("invalid RevocationType: %d. Supported types are: %d", msg.RevocationType, allowedRevocationType))
	}

	cert, err := x509.DecodeX509Certificate(msg.CrlSignerCertificate)
	if err != nil {
		return pkitypes.NewErrInvalidCertificate(err)
	}

	subjectAsMap := x509.SubjectAsTextToMap(cert.SubjectAsText)

	if msg.IsPAA {
		if msg.Pid != 0 {
			return pkitypes.NewErrNotEmptyPid("Product ID (pid) must be empty for root certificates when isPAA is true")
		}

		if !cert.IsSelfSigned() {
			return pkitypes.NewErrPAANotSelfSigned(fmt.Sprintf("CRL Signer Certificate must be self-signed if isPAA is True"))
		}

		strVid, found := subjectAsMap["vid"]
		if found {
			vid, err := strconv.ParseInt(strVid, 10, 32)
			if err != nil {
				return err
			}

			if int32(vid) != msg.Vid {
				return pkitypes.NewErrCRLSignerCertificateVidNotEqualMsgVid("CRL Signer Certificate vid must equal to message vid")
			}
		}
	} else {
		strPid, found := subjectAsMap["pid"]
		if found {
			pid, err := strconv.ParseInt(strPid, 10, 32)
			if err != nil {
				return err
			}

			if int32(pid) != msg.Pid {
				return pkitypes.NewErrCRLSignerCertificatePidNotEqualMsgPid("CRL Signer Certificate pid must equal to message pid")
			}
		} else {
			if msg.Pid != 0 {
				return pkitypes.NewErrNotEmptyPid("Product ID (pid) must be empty when it is not found in root certificate")
			}
		}

		strVid, found := subjectAsMap["vid"]
		if found {
			vid, err := strconv.ParseInt(strVid, 10, 32)
			if err != nil {
				return err
			}

			if int32(vid) != msg.Vid {
				return pkitypes.NewErrCRLSignerCertificateVidNotEqualMsgVid("CRL Signer Certificate vid must equal to message vid")
			}
		} else {
			return pkitypes.NewErrVidNotFound("vid not found in CRL Signer Certificate subject")
		}

		if cert.IsSelfSigned() {
			return pkitypes.NewErrNonPAASelfSigned(fmt.Sprintf("CRL Signer Certificate shall not be self-signed if isPAA is False"))
		}
	}

	if msg.DataFileSize == 0 && msg.DataDigest != "" {
		return pkitypes.NewErrNonEmptyDataDigest("Data Digest must be provided only if Data File Size is provided")
	}

	if msg.DataFileSize != 0 && msg.DataDigest == "" {
		return pkitypes.NewErrEmptyDataDigest("Data Digest must be provided if Data File Size is provided")
	}

	if msg.DataDigest == "" && msg.DataDigestType != 0 {
		return pkitypes.NewErrNonEmptyDataDigestType("Data Digest Type must be provided only if Data Digest is provided")
	}

	if msg.DataDigest != "" && msg.DataDigestType == 0 {
		return pkitypes.NewErrEmptyDataDigestType("Data Digest Type must be provided if Data Digest is provided")
	}

	if msg.RevocationType == allowedRevocationType && (msg.DataFileSize != 0 || msg.DataDigest != "" || msg.DataDigestType != 0) {
		return pkitypes.NewErrDataFieldPresented("Data Digest, Data File Size and Data Digest Type must be omitted for Revocation Type 1")
	}

	match, _ := regexp.MatchString("^(?:[0-9A-F]{2})+$", msg.IssuerSubjectKeyID)

	if !match {
		return pkitypes.NewErrWrongSubjectKeyIDFormat("Wrong IssuerSubjectKeyID format. It must consist of even number of uppercase hexadecimal characters ([0-9A-F]), with no whitespace and no non-hexadecimal characters")
	}

	return nil
}
