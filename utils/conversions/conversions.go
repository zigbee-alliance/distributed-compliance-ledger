// Copyright 2020 DSR Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package conversions

import (
	"fmt"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func ParseUInt8FromString(str string) (uint8, sdk.Error) {
	val, err := strconv.ParseUint(str, 10, 8)
	if err != nil {
		return 0, sdk.ErrUnknownRequest(fmt.Sprintf("Parsing Error: \"%v\" must be 16 bit unsigned integer", str))
	}

	return uint8(val), nil
}

func ParseUInt16FromString(str string) (uint16, sdk.Error) {
	val, err := strconv.ParseUint(str, 10, 16)
	if err != nil {
		return 0, sdk.ErrUnknownRequest(fmt.Sprintf("Parsing Error: \"%v\" must be 16 bit unsigned integer", str))
	}

	return uint16(val), nil
}

func ParseUInt32FromString(str string) (uint32, sdk.Error) {
	val, err := strconv.ParseUint(str, 10, 32)
	if err != nil {
		return 0, sdk.ErrUnknownRequest(fmt.Sprintf("Parsing Error: \"%v\" must be 16 bit unsigned integer", str))
	}

	return uint32(val), nil
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
