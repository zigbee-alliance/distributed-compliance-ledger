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

// x/compliance module sentinel errors.
var (
	ErrComplianceInfoAlreadyExist     = sdkerrors.Register(ModuleName, 301, "compliance info already exist")
	ErrInconsistentDates              = sdkerrors.Register(ModuleName, 302, "inconsistent dates")
	ErrAlreadyCertified               = sdkerrors.Register(ModuleName, 303, "model already certified")
	ErrAlreadyRevoked                 = sdkerrors.Register(ModuleName, 304, "model already revoked")
	ErrAlreadyProvisional             = sdkerrors.Register(ModuleName, 305, "model already in provisional state")
	ErrModelVersionStringDoesNotMatch = sdkerrors.Register(ModuleName, 306, "model version does not match")
	ErrInvalidTestDateFormat          = sdkerrors.Register(ModuleName, 307, "test date must be in RFC3339 format")
	ErrInvalidCertificationType       = sdkerrors.Register(ModuleName, 308, "invalid certification type")
)

func NewErrInconsistentDates(err interface{}) error {
	return sdkerrors.Wrapf(
		ErrInconsistentDates,
		"%v",
		err,
	)
}

func NewErrAlreadyCertified(vid interface{}, pid interface{}, sv interface{}, certificationType interface{}) error {
	return sdkerrors.Wrapf(
		ErrAlreadyCertified,
		"Model with vid=%v, pid=%v, softwareVersion=%v, certificationType=%v already certified on the ledger",
		vid, pid, sv, certificationType,
	)
}

func NewErrAlreadyRevoked(vid interface{}, pid interface{}, sv interface{}, certificationType interface{}) error {
	return sdkerrors.Wrapf(
		ErrAlreadyRevoked,
		"Model with vid=%v, pid=%v, softwareVersion=%v, certificationType=%v already revoked on the ledger",
		vid, pid, sv, certificationType,
	)
}

func NewErrAlreadyProvisional(vid interface{}, pid interface{}, sv interface{}, certificationType interface{}) error {
	return sdkerrors.Wrapf(
		ErrAlreadyProvisional,
		"Model with vid=%v, pid=%v, softwareVersion=%v, certificationType=%v is already in provisional state on the ledger",
		vid, pid, sv, certificationType,
	)
}

func NewErrComplianceInfoAlreadyExist(vid interface{}, pid interface{}, sv interface{}, certificationType interface{}) error {
	return sdkerrors.Wrapf(
		ErrAlreadyRevoked,
		"Model with vid=%v, pid=%v, softwareVersion=%v, certificationType=%v already has compliance info on the ledger",
		vid, pid, sv, certificationType,
	)
}

func NewErrModelVersionStringDoesNotMatch(vid interface{}, pid interface{},
	softwareVersion interface{}, softwareVersionString interface{}) error {
	return sdkerrors.Wrapf(
		ErrModelVersionStringDoesNotMatch,
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

func NewErrInvalidCertificationType(certType interface{}, certList interface{}) error {
	return sdkerrors.Wrapf(ErrInvalidCertificationType,
		"Invalid CertificationType: \"%s\". Supported types: [%s]",
		certType, certList,
	)
}
