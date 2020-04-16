package conversions

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"strconv"
)

func ParseInt16FromString(str string) (int16, sdk.Error) {
	val, err := strconv.ParseInt(str, 10, 16)
	if err != nil {
		return 0, sdk.ErrInternal(fmt.Sprintf("Parsing Error: %v must be 16 bit integer", str))
	}
	return int16(val), nil
}

func ParseVID(str string) (int16, sdk.Error) {
	res, err := ParseInt16FromString(str)
	if err != nil {
		return 0, sdk.ErrInternal(fmt.Sprintf("Invalid VID: %v", err))
	}
	return res, nil
}

func ParsePID(str string) (int16, sdk.Error) {
	res, err := ParseInt16FromString(str)
	if err != nil {
		return 0, sdk.ErrInternal(fmt.Sprintf("Invalid PID: %v", err))
	}
	return res, nil
}

func ParseCID(str string) (int16, sdk.Error) {
	res, err := ParseInt16FromString(str)
	if err != nil {
		return 0, sdk.ErrInternal(fmt.Sprintf("Invalid CID: %v", err))
	}
	return res, nil
}
