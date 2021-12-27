package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAddTestingResult = "add_testing_result"

var _ sdk.Msg = &MsgAddTestingResult{}

func NewMsgAddTestingResult(signer string, vid int32, pid int32, softwareVersion uint32, softwareVersionString string, testResult string, testDate string) *MsgAddTestingResult {
	return &MsgAddTestingResult{
		Signer:                signer,
		Vid:                   vid,
		Pid:                   pid,
		SoftwareVersion:       softwareVersion,
		SoftwareVersionString: softwareVersionString,
		TestResult:            testResult,
		TestDate:              testDate,
	}
}

func (msg *MsgAddTestingResult) Route() string {
	return RouterKey
}

func (msg *MsgAddTestingResult) Type() string {
	return TypeMsgAddTestingResult
}

func (msg *MsgAddTestingResult) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{signer}
}

func (msg *MsgAddTestingResult) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddTestingResult) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signer address (%s)", err)
	}
	return nil
}
