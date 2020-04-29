package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"time"
)

const RouterKey = ModuleName

type MsgCertifyModel struct {
	VID               uint16            `json:"vid"`
	PID               uint16            `json:"pid"`
	CertificationDate time.Time         `json:"certification_date"` // rfc3339 encoded date
	CertificationType CertificationType `json:"certification_type"`
	Reason            string            `json:"reason,omitempty"`
	Signer            sdk.AccAddress    `json:"signer"`
}

func NewMsgCertifyModel(vid uint16, pid uint16, certificationDate time.Time, certificationType CertificationType,
	reason string, signer sdk.AccAddress) MsgCertifyModel {
	return MsgCertifyModel{
		VID:               vid,
		PID:               pid,
		CertificationDate: certificationDate,
		CertificationType: certificationType,
		Reason:            reason,
		Signer:            signer,
	}
}

func (m MsgCertifyModel) Route() string {
	return RouterKey
}

func (m MsgCertifyModel) Type() string {
	return "certify_model"
}

func (m MsgCertifyModel) ValidateBasic() sdk.Error {
	if m.Signer.Empty() {
		return sdk.ErrInvalidAddress("Invalid Signer: it cannot be empty")
	}

	if m.VID == 0 {
		return sdk.ErrUnknownRequest("Invalid VID: it must be non zero 16-bit unsigned integer")
	}
	if m.PID == 0 {
		return sdk.ErrUnknownRequest("Invalid PID: it must be non zero 16-bit unsigned integer")
	}

	if m.CertificationDate.IsZero() {
		return sdk.ErrUnknownRequest("Invalid CertificationDate: it cannot be empty")
	}

	if m.CertificationType != ZbCertificationType {
		return sdk.ErrUnknownRequest(
			fmt.Sprintf("Invalid CertificationType: \"%s\". Supported types: [%s]", m.CertificationType, ZbCertificationType))
	}

	return nil
}

func (m MsgCertifyModel) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(m))
}

func (m MsgCertifyModel) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Signer}
}

type MsgRevokeModel struct {
	VID               uint16            `json:"vid"`
	PID               uint16            `json:"pid"`
	RevocationDate    time.Time         `json:"revocation_date"` // rfc3339 encoded date
	CertificationType CertificationType `json:"certification_type"`
	Reason            string            `json:"reason,omitempty"`
	Signer            sdk.AccAddress    `json:"signer"`
}

func NewMsgRevokeModel(vid uint16, pid uint16, revocationDate time.Time, certificationType CertificationType, revocationReason string, signer sdk.AccAddress) MsgRevokeModel {
	return MsgRevokeModel{
		VID:               vid,
		PID:               pid,
		RevocationDate:    revocationDate,
		CertificationType: certificationType,
		Reason:            revocationReason,
		Signer:            signer,
	}
}

func (m MsgRevokeModel) Route() string {
	return RouterKey
}

func (m MsgRevokeModel) Type() string {
	return "revoke_model"
}

func (m MsgRevokeModel) ValidateBasic() sdk.Error {
	if m.Signer.Empty() {
		return sdk.ErrInvalidAddress("Invalid Signer: it cannot be empty")
	}

	if m.VID == 0 {
		return sdk.ErrUnknownRequest("Invalid VID: it must be non zero 16-bit unsigned integer")
	}
	if m.PID == 0 {
		return sdk.ErrUnknownRequest("Invalid PID: it must be non zero 16-bit unsigned integer")
	}

	if m.RevocationDate.IsZero() {
		return sdk.ErrUnknownRequest("Invalid RevocationDate: it cannot be empty")
	}

	if m.CertificationType != ZbCertificationType {
		return sdk.ErrUnknownRequest(
			fmt.Sprintf("Invalid CertificationType: \"%s\". Supported types: [%s]", m.CertificationType, ZbCertificationType))
	}

	return nil
}

func (m MsgRevokeModel) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(m))
}

func (m MsgRevokeModel) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Signer}
}
