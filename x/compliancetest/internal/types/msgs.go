package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const RouterKey = ModuleName

type MsgAddTestingResult struct {
	VID        int16          `json:"vid"`
	PID        int16          `json:"pid"`
	TestResult string         `json:"test_result"`
	Signer     sdk.AccAddress `json:"signer"`
}

func NewMsgAddTestingResult(vid int16, pid int16, testResult string, signer sdk.AccAddress) MsgAddTestingResult {
	return MsgAddTestingResult{
		VID:        vid,
		PID:        pid,
		TestResult: testResult,
		Signer:     signer,
	}
}

func (m MsgAddTestingResult) Route() string {
	return RouterKey
}

func (m MsgAddTestingResult) Type() string {
	return "add_testing_result"
}

func (m MsgAddTestingResult) ValidateBasic() sdk.Error {
	if m.Signer.Empty() {
		return sdk.ErrInvalidAddress(m.Signer.String())
	}

	if m.VID == 0 {
		return sdk.ErrUnknownRequest("Invalid VID: it must be 16-bit integer")
	}
	if m.PID == 0 {
		return sdk.ErrUnknownRequest("Invalid PID: it must be 16-bit integer")
	}

	if len(m.TestResult) == 0 {
		return sdk.ErrUnknownRequest("Invalid TestResult: it cannot be empty")
	}

	return nil
}

func (m MsgAddTestingResult) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(m))
}

func (m MsgAddTestingResult) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Signer}
}
