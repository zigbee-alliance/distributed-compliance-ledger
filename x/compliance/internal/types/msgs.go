package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"time"
)

const RouterKey = ModuleName

type MsgCertifyModel struct {
	VID               int16          `json:"vid"`
	PID               int16          `json:"pid"`
	CertificationDate time.Time      `json:"certification_date"` // rfc3339 encoded date
	CertificationType string         `json:"certification_type,omitempty"`
	Signer            sdk.AccAddress `json:"signer"`
}

func NewMsgCertifyModel(vid int16, pid int16, certificationDate time.Time, certificationType string, signer sdk.AccAddress) MsgCertifyModel {
	return MsgCertifyModel{
		VID:               vid,
		PID:               pid,
		CertificationDate: certificationDate,
		CertificationType: certificationType,
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
		return sdk.ErrInvalidAddress(m.Signer.String())
	}

	if m.VID == 0 {
		return sdk.ErrUnknownRequest("Invalid VID: it must be not zero 16-bit integer")
	}
	if m.PID == 0 {
		return sdk.ErrUnknownRequest("Invalid PID: it must be not zero 16-bit integer")
	}

	if m.CertificationDate.IsZero() {
		return sdk.ErrUnknownRequest("Invalid CertificationDate: it cannot be empty")
	}

	if m.CertificationType != "" && m.CertificationType != ZbCertificationType {
		return sdk.ErrUnknownRequest(
			fmt.Sprintf("Invalid CertificationType: \"%s\". Supported pagination: [%s]", m.CertificationType, ZbCertificationType))
	}

	return nil
}

func (m MsgCertifyModel) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(m))
}

func (m MsgCertifyModel) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Signer}
}
