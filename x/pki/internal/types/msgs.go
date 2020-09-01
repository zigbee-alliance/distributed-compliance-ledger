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
		return sdk.ErrInvalidAddress("Invalid Signer: it cannot be empty")
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
	SubjectKeyID string         `json:"subject_key_id"`
	Signer       sdk.AccAddress `json:"signer"`
}

func NewMsgApproveAddX509RootCert(subject string, subjectKeyID string,
	signer sdk.AccAddress) MsgApproveAddX509RootCert {
	return MsgApproveAddX509RootCert{
		Subject:      subject,
		SubjectKeyID: subjectKeyID,
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
		return sdk.ErrInvalidAddress("Invalid Signer: it cannot be empty")
	}

	if len(m.Subject) == 0 {
		return sdk.ErrUnknownRequest("Invalid Subject: it cannot be empty")
	}

	if len(m.SubjectKeyID) == 0 {
		return sdk.ErrUnknownRequest("Invalid SubjectKeyID: it cannot be empty")
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
		return sdk.ErrInvalidAddress("Invalid Signer: it cannot be empty")
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

/*
	PROPOSE_REVOKE_X509_ROOT_CERT
*/

type MsgProposeRevokeX509RootCert struct {
	Subject      string         `json:"subject"`
	SubjectKeyID string         `json:"subject_key_id"`
	Signer       sdk.AccAddress `json:"signer"`
}

func NewMsgProposeRevokeX509RootCert(subject string, subjectKeyID string,
	signer sdk.AccAddress) MsgProposeRevokeX509RootCert {
	return MsgProposeRevokeX509RootCert{
		Subject:      subject,
		SubjectKeyID: subjectKeyID,
		Signer:       signer,
	}
}

func (m MsgProposeRevokeX509RootCert) Route() string {
	return RouterKey
}

func (m MsgProposeRevokeX509RootCert) Type() string {
	return "propose_revoke_x509_root_cert"
}

func (m MsgProposeRevokeX509RootCert) ValidateBasic() sdk.Error {
	if m.Signer.Empty() {
		return sdk.ErrInvalidAddress("Invalid Signer: it cannot be empty")
	}

	if len(m.Subject) == 0 {
		return sdk.ErrUnknownRequest("Invalid Subject: it cannot be empty")
	}

	if len(m.SubjectKeyID) == 0 {
		return sdk.ErrUnknownRequest("Invalid SubjectKeyID: it cannot be empty")
	}

	return nil
}

func (m MsgProposeRevokeX509RootCert) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(m))
}

func (m MsgProposeRevokeX509RootCert) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Signer}
}

/*
	APPROVE_REVOKE_X509_ROOT_CERT
*/

type MsgApproveRevokeX509RootCert struct {
	Subject      string         `json:"subject"`
	SubjectKeyID string         `json:"subject_key_id"`
	Signer       sdk.AccAddress `json:"signer"`
}

func NewMsgApproveRevokeX509RootCert(subject string, subjectKeyID string,
	signer sdk.AccAddress) MsgApproveRevokeX509RootCert {
	return MsgApproveRevokeX509RootCert{
		Subject:      subject,
		SubjectKeyID: subjectKeyID,
		Signer:       signer,
	}
}

func (m MsgApproveRevokeX509RootCert) Route() string {
	return RouterKey
}

func (m MsgApproveRevokeX509RootCert) Type() string {
	return "approve_revoke_x509_root_cert"
}

func (m MsgApproveRevokeX509RootCert) ValidateBasic() sdk.Error {
	if m.Signer.Empty() {
		return sdk.ErrInvalidAddress("Invalid Signer: it cannot be empty")
	}

	if len(m.Subject) == 0 {
		return sdk.ErrUnknownRequest("Invalid Subject: it cannot be empty")
	}

	if len(m.SubjectKeyID) == 0 {
		return sdk.ErrUnknownRequest("Invalid SubjectKeyID: it cannot be empty")
	}

	return nil
}

func (m MsgApproveRevokeX509RootCert) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(m))
}

func (m MsgApproveRevokeX509RootCert) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Signer}
}

/*
	REVOKE_X509_CERT
*/

type MsgRevokeX509Cert struct {
	Subject      string         `json:"subject"`
	SubjectKeyID string         `json:"subject_key_id"`
	Signer       sdk.AccAddress `json:"signer"`
}

func NewMsgRevokeX509Cert(subject string, subjectKeyID string, signer sdk.AccAddress) MsgRevokeX509Cert {
	return MsgRevokeX509Cert{
		Subject:      subject,
		SubjectKeyID: subjectKeyID,
		Signer:       signer,
	}
}

func (m MsgRevokeX509Cert) Route() string {
	return RouterKey
}

func (m MsgRevokeX509Cert) Type() string {
	return "revoke_x509_cert"
}

func (m MsgRevokeX509Cert) ValidateBasic() sdk.Error {
	if m.Signer.Empty() {
		return sdk.ErrInvalidAddress("Invalid Signer: it cannot be empty")
	}

	if len(m.Subject) == 0 {
		return sdk.ErrUnknownRequest("Invalid Subject: it cannot be empty")
	}

	if len(m.SubjectKeyID) == 0 {
		return sdk.ErrUnknownRequest("Invalid SubjectKeyID: it cannot be empty")
	}

	return nil
}

func (m MsgRevokeX509Cert) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(m))
}

func (m MsgRevokeX509Cert) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Signer}
}
