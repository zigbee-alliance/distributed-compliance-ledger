package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const RouterKey = ModuleName

/*
	PROPOSE_ADD_X509_ROOT_CERT
*/

type MsgProposeAddX509RootCert struct {
	Cert   string         `json:"cert"`
	Signer sdk.AccAddress `json:"signer"`
}

func NewMsgProposeAddX509RootCert(cert string, signer sdk.AccAddress) MsgProposeAddX509RootCert {
	return MsgProposeAddX509RootCert{
		Cert:   cert,
		Signer: signer,
	}
}

func (m MsgProposeAddX509RootCert) Route() string {
	return RouterKey
}

func (m MsgProposeAddX509RootCert) Type() string {
	return "propose_add_x509_root_cert"
}

func (m MsgProposeAddX509RootCert) ValidateBasic() sdk.Error {
	if m.Signer.Empty() {
		return sdk.ErrInvalidAddress(m.Signer.String())
	}

	if len(m.Cert) == 0 {
		return sdk.ErrUnknownRequest("Invalid x509Cert: it cannot be empty")
	}

	return nil
}

func (m MsgProposeAddX509RootCert) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(m))
}

func (m MsgProposeAddX509RootCert) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Signer}
}

/*
	APPROVE_ADD_X509_ROOT_CERT
*/

type MsgApproveAddX509RootCert struct {
	Subject      string         `json:"subject"`
	SubjectKeyId string         `json:"subject_key_id"`
	Signer       sdk.AccAddress `json:"signer"`
}

func NewMsgApproveAddX509RootCert(subject string, subjectKeyId string, signer sdk.AccAddress) MsgApproveAddX509RootCert {
	return MsgApproveAddX509RootCert{
		Subject:      subject,
		SubjectKeyId: subjectKeyId,
		Signer:       signer,
	}
}

func (m MsgApproveAddX509RootCert) Route() string {
	return RouterKey
}

func (m MsgApproveAddX509RootCert) Type() string {
	return "approve_add_x509_root_cert"
}

func (m MsgApproveAddX509RootCert) ValidateBasic() sdk.Error {
	if m.Signer.Empty() {
		return sdk.ErrInvalidAddress(m.Signer.String())
	}

	if len(m.Subject) == 0 {
		return sdk.ErrUnknownRequest("Invalid Subject: it cannot be empty")
	}

	if len(m.SubjectKeyId) == 0 {
		return sdk.ErrUnknownRequest("Invalid SubjectKeyId: it cannot be empty")
	}

	return nil
}

func (m MsgApproveAddX509RootCert) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(m))
}

func (m MsgApproveAddX509RootCert) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Signer}
}

/*
	ADD_X509_CERT
*/

type MsgAddX509Cert struct {
	Cert   string         `json:"cert"`
	Signer sdk.AccAddress `json:"signer"`
}

func NewMsgAddX509Cert(cert string, signer sdk.AccAddress) MsgAddX509Cert {
	return MsgAddX509Cert{
		Cert:   cert,
		Signer: signer,
	}
}

func (m MsgAddX509Cert) Route() string {
	return RouterKey
}

func (m MsgAddX509Cert) Type() string {
	return "add_x509_cert"
}

func (m MsgAddX509Cert) ValidateBasic() sdk.Error {
	if m.Signer.Empty() {
		return sdk.ErrInvalidAddress(m.Signer.String())
	}

	if len(m.Cert) == 0 {
		return sdk.ErrUnknownRequest("Invalid x509Cert: it cannot be empty")
	}

	return nil
}

func (m MsgAddX509Cert) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(m))
}

func (m MsgAddX509Cert) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Signer}
}
