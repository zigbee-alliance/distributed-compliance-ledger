package conversions

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"strconv"
)

func ParseUInt16FromString(str string) (uint16, sdk.Error) {
	val, err := strconv.ParseUint(str, 10, 16)
	if err != nil || val < 0 {
		return 0, sdk.ErrUnknownRequest(fmt.Sprintf("Parsing Error: \"%v\" must be 16 bit unsigned integer", str))
	}
	return uint16(val), nil
}

func ParseVID(str string) (uint16, sdk.Error) {
	res, err := ParseUInt16FromString(str)
	if err != nil {
		return 0, sdk.ErrUnknownRequest(fmt.Sprintf("Invalid VID: %v", err.Data()))
	}
	if res == 0 {
		return 0, sdk.ErrUnknownRequest("Invalid VID: it must be non zero 16-bit unsigned integer")
	}
	return res, nil
}

func ParsePID(str string) (uint16, sdk.Error) {
	res, err := ParseUInt16FromString(str)
	if err != nil {
		return 0, sdk.ErrUnknownRequest(fmt.Sprintf("Invalid PID: %v", err.Data()))
	}
	if res == 0 {
		return 0, sdk.ErrUnknownRequest("Invalid PID: it must be non zero 16-bit unsigned integer")
	}
	return res, nil
}

func ParseCID(str string) (uint16, sdk.Error) {
	res, err := ParseUInt16FromString(str)
	if err != nil {
		return 0, sdk.ErrUnknownRequest(fmt.Sprintf("Invalid CID: %v", err.Data()))
	}
	return res, nil
}
