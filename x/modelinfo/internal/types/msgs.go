package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const RouterKey = ModuleName

type MsgAddModelInfo struct {
	VID                      int16          `json:"vid"`
	PID                      int16          `json:"pid"`
	CID                      int16          `json:"cid,omitempty"`
	Name                     string         `json:"name"`
	Description              string         `json:"description"`
	SKU                      string         `json:"sku"`
	FirmwareVersion          string         `json:"firmware_version"`
	HardwareVersion          string         `json:"hardware_version"`
	Custom                   string         `json:"custom,omitempty"`
	CertificateID            string         `json:"certificate_id,omitempty"`
	CertifiedDate            time.Time      `json:"certified_date,omitempty"`
	TisOrTrpTestingCompleted bool           `json:"tis_or_trp_testing_completed"`
	Signer                   sdk.AccAddress `json:"signer"`
}

func NewMsgAddModelInfo(vid int16, pid int16, cid int16, name string, description string, sku string,
	firmwareVersion string, hardwareVersion string, custom string, certificateID string, certifiedDate time.Time,
	tisOrTrpTestingCompleted bool, signer sdk.AccAddress) MsgAddModelInfo {
	return MsgAddModelInfo{
		VID:                      vid,
		PID:                      pid,
		CID:                      cid,
		Name:                     name,
		Description:              description,
		SKU:                      sku,
		FirmwareVersion:          firmwareVersion,
		HardwareVersion:          hardwareVersion,
		Custom:                   custom,
		CertificateID:            certificateID,
		CertifiedDate:            certifiedDate,
		TisOrTrpTestingCompleted: tisOrTrpTestingCompleted,
		Signer:                   signer,
	}
}

func (m MsgAddModelInfo) Route() string {
	return RouterKey
}

func (m MsgAddModelInfo) Type() string {
	return "add_model_info"
}

func (m MsgAddModelInfo) ValidateBasic() sdk.Error {
	if m.Signer.Empty() {
		return sdk.ErrInvalidAddress(m.Signer.String())
	}

	if m.VID == 0 ||
		m.PID == 0 ||
		len(m.Name) == 0 ||
		len(m.Description) == 0 ||
		len(m.SKU) == 0 ||
		len(m.FirmwareVersion) == 0 ||
		len(m.HardwareVersion) == 0 {
		return sdk.ErrUnknownRequest("VID, PID, Name, Description, SKU, FirmwareVersion and HardwareVersion  " +
			"cannot be empty")
	}

	return nil
}

func (m MsgAddModelInfo) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(m))
}

func (m MsgAddModelInfo) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Signer}
}

type MsgUpdateModelInfo struct {
	VID                      int16          `json:"vid"`
	PID                      int16          `json:"pid"`
	CID                      int16          `json:"cid,omitempty"`
	Name                     string         `json:"name"`
	Description              string         `json:"description"`
	SKU                      string         `json:"sku"`
	FirmwareVersion          string         `json:"firmware_version"`
	HardwareVersion          string         `json:"hardware_version"`
	Custom                   string         `json:"custom,omitempty"`
	CertificateID            string         `json:"certificate_id,omitempty"`
	CertifiedDate            time.Time      `json:"certified_date,omitempty"`
	TisOrTrpTestingCompleted bool           `json:"tis_or_trp_testing_completed"`
	Signer                   sdk.AccAddress `json:"signer"`
}

func NewMsgUpdateModelInfo(vid int16, pid int16, cid int16, name string, description string, sku string,
	firmwareVersion string, hardwareVersion string, custom string, certificateID string, certifiedDate time.Time,
	tisOrTrpTestingCompleted bool, signer sdk.AccAddress) MsgUpdateModelInfo {
	return MsgUpdateModelInfo{
		VID:                      vid,
		PID:                      pid,
		CID:                      cid,
		Name:                     name,
		Description:              description,
		SKU:                      sku,
		FirmwareVersion:          firmwareVersion,
		HardwareVersion:          hardwareVersion,
		Custom:                   custom,
		CertificateID:            certificateID,
		CertifiedDate:            certifiedDate,
		TisOrTrpTestingCompleted: tisOrTrpTestingCompleted,
		Signer:                   signer,
	}
}

func (m MsgUpdateModelInfo) Route() string {
	return RouterKey
}

func (m MsgUpdateModelInfo) Type() string {
	return "update_model_info"
}

func (m MsgUpdateModelInfo) ValidateBasic() sdk.Error {
	if m.Signer.Empty() {
		return sdk.ErrInvalidAddress(m.Signer.String())
	}

	if m.VID == 0 ||
		m.PID == 0 ||
		len(m.Name) == 0 ||
		len(m.Description) == 0 ||
		len(m.SKU) == 0 ||
		len(m.FirmwareVersion) == 0 ||
		len(m.HardwareVersion) == 0 {
		return sdk.ErrUnknownRequest("VID, PID, Name, Description, SKU, FirmwareVersion and HardwareVersion  " +
			"cannot be empty")
	}

	return nil
}

func (m MsgUpdateModelInfo) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(m))
}

func (m MsgUpdateModelInfo) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Signer}
}

type MsgDeleteModelInfo struct {
	VID    int16          `json:"vid"`
	PID    int16          `json:"pid"`
	Signer sdk.AccAddress `json:"signer"`
}

func NewMsgDeleteModelInfo(vid int16, pid int16, signer sdk.AccAddress) MsgDeleteModelInfo {
	return MsgDeleteModelInfo{
		VID:    vid,
		PID:    pid,
		Signer: signer,
	}
}

func (m MsgDeleteModelInfo) Route() string {
	return RouterKey
}

func (m MsgDeleteModelInfo) Type() string {
	return "delete_model_info"
}

func (m MsgDeleteModelInfo) ValidateBasic() sdk.Error {
	if m.Signer.Empty() {
		return sdk.ErrInvalidAddress(m.Signer.String())
	}

	if m.VID == 0 || m.PID == 0 {
		return sdk.ErrUnknownRequest("VID and PID cannot be empty")
	}

	return nil
}

func (m MsgDeleteModelInfo) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(m))
}

func (m MsgDeleteModelInfo) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Signer}
}
