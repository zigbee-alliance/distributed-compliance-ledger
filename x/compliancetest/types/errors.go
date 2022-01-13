// Copyright 2022 DSR Corporation
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

package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/compliancetest module sentinel errors.
var (
	ErrTestingResultsDoNotExist       = sdkerrors.Register(ModuleName, 201, "testing result does not exist")
	ErrModelVersionStringDoesNotMatch = sdkerrors.Register(ModuleName, 202, "model version does not match")
	ErrInvalidTestDateFormat          = sdkerrors.Register(ModuleName, 203, "test date must be in RFC3339 format")
)

func NewErrTestingResultsDoNotExist(vid interface{}, pid interface{}, softwareVersion uint32) error {
	return sdkerrors.Wrapf(ErrTestingResultsDoNotExist,
		"No testing results about the model with vid=%v pid=%v and softwareVersion=%v on the ledger",
		vid, pid, softwareVersion,
	)
}

func NewErrModelVersionStringDoesNotMatch(vid interface{}, pid interface{},
	softwareVersion interface{}, softwareVersionString interface{}) error {
	return sdkerrors.Wrapf(ErrModelVersionStringDoesNotMatch,
		"Model with vid=%v, pid=%v, softwareVersion=%v present on the ledger does not have"+
			" matching softwareVersionString=%v",
		vid, pid, softwareVersion, softwareVersionString,
	)
}

func NewErrInvalidTestDateFormat(testDate interface{}) error {
	return sdkerrors.Wrapf(ErrInvalidTestDateFormat,
		"Invalid TestDate \"%v\": it must be RFC3339 encoded date, for example 2019-10-12T07:20:50.52Z",
		testDate,
	)
}
