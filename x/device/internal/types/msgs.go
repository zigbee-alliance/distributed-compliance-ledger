package types

import sdk "github.com/cosmos/cosmos-sdk/types"

const RouterName = ModuleName

type MsgAddDevice struct {
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
	return RouterName
}

func (m MsgAddDevice) Type() string {
	return "add_device"
}

func (m MsgAddDevice) ValidateBasic() sdk.Error {
	if m.Owner.Empty() {
		return sdk.ErrInvalidAddress(m.Owner.String())
	}

	if len(m.Family) == 0 || len(m.Cert) == 0 {
		return sdk.ErrUnknownRequest("Family and/or Cert cannot be empty")
	}

	return nil
}

func (m MsgAddDevice) GetSignBytes() []byte {
	panic("implement me")
}

func (m MsgAddDevice) GetSigners() []sdk.AccAddress {
	panic("implement me")
}
