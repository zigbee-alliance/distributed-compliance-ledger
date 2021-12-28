package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/compliance module sentinel errors
var (
	ErrComplianceInfoDoesNotExist     = sdkerrors.Register(ModuleName, 301, "compliance info does not exist")
	ErrInconsistentDates              = sdkerrors.Register(ModuleName, 302, "inconsistent dates")
	ErrAlreadyCertified               = sdkerrors.Register(ModuleName, 303, "model already certified")
	ErrModelDoesNotExist              = sdkerrors.Register(ModuleName, 304, "model does not exist")
	ErrModelVersionStringDoesNotMatch = sdkerrors.Register(ModuleName, 305, "model version does not match")
	ErrInvalidTestDateFormat          = sdkerrors.Register(ModuleName, 306, "test date must be in RFC3339 format")
	ErrInvalidCertificationType       = sdkerrors.Register(ModuleName, 307, "invalid certification type")
)

func NewErrComplianceInfoDoesNotExist(vid interface{}, pid interface{},
	softwareVersion interface{}, certificationType interface{}) error {
	return sdkerrors.Wrapf(
		ErrComplianceInfoDoesNotExist,
		"No certification information about the model with vid=%v, pid=%v softwareVersion=%v "+
			"certification_type=%v on the ledger. This means that the model is either not certified yet or "+
			"certified by default (off-ledger).",
		vid, pid, softwareVersion, certificationType,
	)
}

func NewErrInconsistentDates(err interface{}) error {
	return sdkerrors.Wrapf(
		ErrInconsistentDates,
		"%v",
		err,
	)
}

func NewErrAlreadyCertifyed(vid interface{}, pid interface{}) error {
	return sdkerrors.Wrapf(
		ErrAlreadyCertified,
		"Model with vid=%v, pid=%v already certified on the ledger",
		vid, pid,
	)
}

func NewErrModelDoesNotExist(vid interface{}, pid interface{}) error {
	return sdkerrors.Wrapf(
		ErrModelDoesNotExist,
		"Model with vid=%v, pid=%v does not exist on the ledger",
		vid, pid,
	)
}

func NewErrModelVersionStringDoesNotMatch(vid interface{}, pid interface{},
	softwareVersion interface{}, softwareVersionString interface{}) error {
	return sdkerrors.Wrapf(
		ErrModelVersionStringDoesNotMatch,
		"Model with vid=%v, pid=%v, softwareVersion=%v present on the ledger does not have "+
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
