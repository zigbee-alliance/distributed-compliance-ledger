package types

import sdk "github.com/cosmos/cosmos-sdk/types"

const RouterKey = ModuleName

type MsgAddDevice struct {
	ID     string         `json:"id"`
	Family string         `json:"family"`
	Cert   string         `json:"cert"`
	Owner  sdk.AccAddress `json:"owner"`
}

func NewMsgAddDevice(family string, cert string, owner sdk.AccAddress) MsgAddDevice {
	return MsgAddDevice{
		Family: family,
		Cert:   cert,
		Owner:  owner,
	}
}

func (m MsgAddDevice) Route() string {
	return RouterKey
}

func (m MsgAddDevice) Type() string {
	return "add_device"
}

func (m MsgAddDevice) ValidateBasic() sdk.Error {
	if m.Owner.Empty() {
		return sdk.ErrInvalidAddress(m.Owner.String())
	}

	if len(m.ID) == 0 || len(m.Family) == 0 || len(m.Cert) == 0 {
		return sdk.ErrUnknownRequest("ID, Family and Cert cannot be empty")
	}

	return nil
}

func (m MsgAddDevice) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(m))
}

func (m MsgAddDevice) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Owner}
}

type MsgAddCompliance struct {
	ID          string         `json:"id"`
	Description string         `json:"description"`
	TestedBy    sdk.AccAddress `json:"tested_by"`
}

func NewMsgAddCompliance(id string, description string, testedBy sdk.AccAddress) MsgAddCompliance {
	return MsgAddCompliance{
		ID:          id,
		Description: description,
		TestedBy:    testedBy,
	}
}

func (m MsgAddCompliance) Route() string {
	return RouterKey
}

func (m MsgAddCompliance) Type() string {
	return "add_compliance"
}

func (m MsgAddCompliance) ValidateBasic() sdk.Error {
	if m.TestedBy.Empty() {
		return sdk.ErrInvalidAddress(m.TestedBy.String())
	}

	if len(m.ID) == 0 || len(m.Description) == 0 {
		return sdk.ErrUnknownRequest("ID and/or Description cannot be empty")
	}

	return nil
}

func (m MsgAddCompliance) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(m))
}
func (m MsgAddCompliance) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.TestedBy}
}

type MsgApproveCompliance struct {
	DeviceID     string         `json:"device_id"`
	ComplianceID string         `json:"compliance_id"`
	ApprovedBy   sdk.AccAddress `json:"owner"`
}

func NewMsgApproveCompliance(deviceID string, complianceID string, approvedBy sdk.AccAddress) MsgApproveCompliance {
	return MsgApproveCompliance{
		DeviceID:     deviceID,
		ComplianceID: complianceID,
		ApprovedBy:   approvedBy,
	}
}

func (m MsgApproveCompliance) Route() string {
	return RouterKey
}

func (m MsgApproveCompliance) Type() string {
	return "approve_compliance"
}

func (m MsgApproveCompliance) ValidateBasic() sdk.Error {
	if m.ApprovedBy.Empty() {
		return sdk.ErrInvalidAddress(m.ApprovedBy.String())
	}

	if len(m.DeviceID) == 0 || len(m.ComplianceID) == 0 {
		return sdk.ErrUnknownRequest("DeviceID and/or ComplianceID cannot be empty")
	}

	return nil
}

func (m MsgApproveCompliance) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(m))
}

func (m MsgApproveCompliance) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.ApprovedBy}
}
