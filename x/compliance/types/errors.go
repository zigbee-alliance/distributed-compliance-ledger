package types

// DONTCOVER

import (
	"cosmossdk.io/errors"
)

// x/compliance module sentinel errors.
var (
	ErrComplianceInfoAlreadyExist      = errors.Register(ModuleName, 301, "compliance info already exist")
	ErrInconsistentDates               = errors.Register(ModuleName, 302, "inconsistent dates")
	ErrAlreadyCertified                = errors.Register(ModuleName, 303, "model already certified")
	ErrAlreadyRevoked                  = errors.Register(ModuleName, 304, "model already revoked")
	ErrAlreadyProvisional              = errors.Register(ModuleName, 305, "model already in provisional state")
	ErrModelVersionStringDoesNotMatch  = errors.Register(ModuleName, 306, "model version does not match")
	ErrInvalidTestDateFormat           = errors.Register(ModuleName, 307, "test date must be in RFC3339 format")
	ErrInvalidCertificationType        = errors.Register(ModuleName, 308, "invalid certification type")
	ErrInvalidPFCCertificationRoute    = errors.Register(ModuleName, 309, "invalid PFC certification route")
	ErrComplianceInfoDoesNotExist      = errors.Register(ModuleName, 310, "compliance info not found")
	ErrInvalidUint32ForCdVersionNumber = errors.Register(ModuleName, 311, "invalid uint32 for cd version number")
	ErrInvalidCertificationRoute       = errors.Register(ModuleName, 312, "invalid certification route")
	ErrInvalidFamilyID                 = errors.Register(ModuleName, 313, "invalid familyID")
	ErrInvalidTransport                = errors.Register(ModuleName, 314, "invalid transport")
	ErrInvalidSupportedClusters        = errors.Register(ModuleName, 315, "invalid supportedClusters")
	ErrInvalidProgramType              = errors.Register(ModuleName, 316, "invalid program type")
)

func NewErrInconsistentDates(err interface{}) error {
	return errors.Wrapf(
		ErrInconsistentDates,
		"%v",
		err,
	)
}

func NewErrAlreadyCertified(vid interface{}, pid interface{}, sv interface{}, certificationType interface{}) error {
	return errors.Wrapf(
		ErrAlreadyCertified,
		"Model with vid=%v, pid=%v, softwareVersion=%v, certificationType=%v already certified on the ledger",
		vid, pid, sv, certificationType,
	)
}

func NewErrAlreadyRevoked(vid interface{}, pid interface{}, sv interface{}, certificationType interface{}) error {
	return errors.Wrapf(
		ErrAlreadyRevoked,
		"Model with vid=%v, pid=%v, softwareVersion=%v, certificationType=%v already revoked on the ledger",
		vid, pid, sv, certificationType,
	)
}

func NewErrAlreadyProvisional(vid interface{}, pid interface{}, sv interface{}, certificationType interface{}) error {
	return errors.Wrapf(
		ErrAlreadyProvisional,
		"Model with vid=%v, pid=%v, softwareVersion=%v, certificationType=%v is already in provisional state on the ledger",
		vid, pid, sv, certificationType,
	)
}

func NewErrComplianceInfoAlreadyExist(vid interface{}, pid interface{}, sv interface{}, certificationType interface{}) error {
	return errors.Wrapf(
		ErrComplianceInfoAlreadyExist,
		"Model with vid=%v, pid=%v, softwareVersion=%v, certificationType=%v already has compliance info on the ledger",
		vid, pid, sv, certificationType,
	)
}

func NewErrComplianceInfoDoesNotExist(vid interface{}, pid interface{}, sv interface{}, certificationType interface{}) error {
	return errors.Wrapf(
		ErrComplianceInfoDoesNotExist,
		"Model with vid=%v, pid=%v, softwareVersion=%v, certificationType=%v has no compliance info on the ledger",
		vid, pid, sv, certificationType,
	)
}

func NewErrInvalidUint32ForCdVersionNumber(vid interface{}, pid interface{}, sv interface{}, certificationType interface{}, cdVersionNumber interface{}) error {
	return errors.Wrapf(
		ErrInvalidUint32ForCdVersionNumber,
		"Compliance info with vid=%v, pid=%v, softwareVersion=%v, certificationType=%v cannot be updated with an invalid uint32 cd version number %v",
		vid, pid, sv, certificationType, cdVersionNumber,
	)
}

func NewErrModelVersionStringDoesNotMatch(vid interface{}, pid interface{},
	softwareVersion interface{}, softwareVersionString interface{},
) error {
	return errors.Wrapf(
		ErrModelVersionStringDoesNotMatch,
		"Model with vid=%v, pid=%v, softwareVersion=%v present on the ledger does not have"+
			" matching softwareVersionString=%v",
		vid, pid, softwareVersion, softwareVersionString,
	)
}

func NewErrModelVersionCDVersionNumberDoesNotMatch(vid interface{}, pid interface{},
	softwareVersion interface{}, cDVersionNumber interface{},
) error {
	return errors.Wrapf(
		ErrModelVersionStringDoesNotMatch,
		"Model with vid=%v, pid=%v, softwareVersion=%v present on the ledger does not have"+
			" matching CDVersionNumber=%v",
		vid, pid, softwareVersion, cDVersionNumber,
	)
}

func NewErrInvalidTestDateFormat(testDate interface{}) error {
	return errors.Wrapf(ErrInvalidTestDateFormat,
		"Invalid TestDate \"%v\": it must be RFC3339 encoded date, for example 2019-10-12T07:20:50.52Z",
		testDate,
	)
}

func NewErrInvalidCertificationType(certType interface{}, certList interface{}) error {
	return errors.Wrapf(ErrInvalidCertificationType,
		"Invalid CertificationType: \"%s\". Supported types: [%s]",
		certType, certList,
	)
}

func NewErrInvalidPFCCertificationRoute(certRoute interface{}, certList interface{}) error {
	return errors.Wrapf(ErrInvalidPFCCertificationRoute,
		"Invalid PFCCertificationRoute: \"%s\". Supported types: [%s]",
		certRoute, certList,
	)
}

func NewErrInvalidCertificationRoute(certRoute interface{}, certList interface{}) error {
	return errors.Wrapf(ErrInvalidCertificationRoute,
		"Invalid CertificationRoute: \"%s\". Supported routes: [%s]",
		certRoute, certList,
	)
}

func NewErrInvalidFamilyID(familyID interface{}) error {
	return errors.Wrapf(
		ErrInvalidFamilyID,
		"Invalid FamilyID: \"%v\", It should start with the 'FAM' prefix, followed by alphanumeric characters",
		familyID,
	)
}

func NewErrInvalidTransport(transport interface{}, transportList interface{}) error {
	return errors.Wrapf(
		ErrInvalidTransport,
		"Invalid Transport: \"%v\". Supported transports: [%s]. When multiple transports are supported, "+
			"they must be comma-separated without spaces or duplicates (e.g. \"wi-fi,ethernet,bluetooth\")",
		transport, transportList,
	)
}

func NewErrInvalidSupportedClusters(supportedClusters interface{}) error {
	return errors.Wrapf(
		ErrInvalidSupportedClusters,
		"Invalid SupportedClusters: \"%v\". It must be a comma-separated list of hexadecimal cluster IDs, "+
			"each formatted as 0x followed by 1-4 hex digits, without spaces or duplicates "+
			"(e.g. \"0x0003,0x0004,0x0006,0x0008,0x0062,0x0300\")",
		supportedClusters,
	)
}

func NewErrInvalidProgramType(programType interface{}, programTypeList interface{}) error {
	return errors.Wrapf(ErrInvalidProgramType,
		"Invalid ProgramType: \"%s\". Supported types: [%s]",
		programType, programTypeList,
	)
}
