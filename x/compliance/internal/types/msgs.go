package types

import sdk "github.com/cosmos/cosmos-sdk/types"

const RouterKey = ModuleName

type MsgAddModelInfo struct {
	ID     string         `json:"id"`
	Family string         `json:"family"`
	Cert   string         `json:"cert"`
	Owner  sdk.AccAddress `json:"owner"`
}

func NewMsgAddModelInfo(id string, family string, cert string, owner sdk.AccAddress) MsgAddModelInfo {
	return MsgAddModelInfo{ID: id, Family: family, Cert: cert, Owner: owner}
}

func (m MsgAddModelInfo) Route() string {
	return RouterKey
}

func (m MsgAddModelInfo) Type() string {
	return "add_model_info"
}

func (m MsgAddModelInfo) ValidateBasic() sdk.Error {
	if m.Owner.Empty() {
		return sdk.ErrInvalidAddress(m.Owner.String())
	}

	if len(m.ID) == 0 || len(m.Family) == 0 || len(m.Cert) == 0 {
		return sdk.ErrUnknownRequest("ID, Family and Cert cannot be empty")
	}

	return nil
}

func (m MsgAddModelInfo) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(m))
}

func (m MsgAddModelInfo) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Owner}
}
