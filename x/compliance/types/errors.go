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
	ErrInvalidPFCCertificationRoute   = sdkerrors.Register(ModuleName, 309, "invalid PFC certification route")
	ErrComplianceInfoDoesNotExist     = sdkerrors.Register(ModuleName, 310, "compliance info not found")
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

func NewErrComplianceInfoDoesNotExist(vid interface{}, pid interface{}, sv interface{}, certificationType interface{}) error {
	return sdkerrors.Wrapf(
		ErrComplianceInfoDoesNotExist,
		"Compliance info with vid=%v, pid=%v, softwareVersion=%v, certificationType=%v does not exist on the ledger",
		vid, pid, sv, certificationType,
	)
}

func NewErrModelVersionStringDoesNotMatch(vid interface{}, pid interface{},
	softwareVersion interface{}, softwareVersionString interface{},
) error {
	return sdkerrors.Wrapf(
		ErrModelVersionStringDoesNotMatch,
		"Model with vid=%v, pid=%v, softwareVersion=%v present on the ledger does not have"+
			" matching softwareVersionString=%v",
		vid, pid, softwareVersion, softwareVersionString,
	)
}

func NewErrModelVersionCDVersionNumberDoesNotMatch(vid interface{}, pid interface{},
	softwareVersion interface{}, cDVersionNumber interface{},
) error {
	return sdkerrors.Wrapf(
		ErrModelVersionStringDoesNotMatch,
		"Model with vid=%v, pid=%v, softwareVersion=%v present on the ledger does not have"+
			" matching CDVersionNumber=%v",
		vid, pid, softwareVersion, cDVersionNumber,
	)
}

func NewErrModelVersionSoftwareVersionStringDoesNotMatch(vid interface{}, pid interface{},
	softwareVersion interface{}, softwareVersionString interface{},
) error {
	return sdkerrors.Wrapf(
		ErrModelVersionStringDoesNotMatch,
		"Model with vid=%v, pid=%v, softwareVersion=%v present on the ledger does not have"+
			" matching SoftwareVersionString=%s",
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

func NewErrInvalidPFCCertificationRoute(certRoute interface{}, certList interface{}) error {
	return sdkerrors.Wrapf(ErrInvalidPFCCertificationRoute,
		"Invalid PFCCertificationRoute: \"%s\". Supported types: [%s]",
		certRoute, certList,
	)
}
