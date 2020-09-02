package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const RouterKey = ModuleName

type MsgAddTestingResult struct {
	VID        uint16         `json:"vid"`
	PID        uint16         `json:"pid"`
	TestResult string         `json:"test_result"`
	TestDate   time.Time      `json:"test_date"` // rfc3339 encoded date
	Signer     sdk.AccAddress `json:"signer"`
}

func NewMsgAddTestingResult(vid uint16, pid uint16, testResult string,
	testDate time.Time, signer sdk.AccAddress) MsgAddTestingResult {
	return MsgAddTestingResult{
		VID:        vid,
		PID:        pid,
		TestResult: testResult,
		TestDate:   testDate,
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
		return sdk.ErrInvalidAddress("Invalid Signer: it cannot be empty")
	}

	if m.VID == 0 {
		return sdk.ErrUnknownRequest("Invalid VID: it must be non zero 16-bit unsigned integer")
	}

	if m.PID == 0 {
		return sdk.ErrUnknownRequest("Invalid PID: it must be non zero 16-bit unsigned integer")
	}

	if len(m.TestResult) == 0 {
		return sdk.ErrUnknownRequest("Invalid TestResult: it cannot be empty")
	}

	if m.TestDate.IsZero() {
		return sdk.ErrUnknownRequest("Invalid TestDate: it cannot be empty")
	}

	return nil
}

func (m MsgAddTestingResult) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(m))
}

func (m MsgAddTestingResult) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Signer}
}
