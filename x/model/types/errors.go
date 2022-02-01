package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	// Model Error Codes.
	ErrModelAlreadyExists       = sdkerrors.Register(ModuleName, 501, "model already exists")
	ErrModelDoesNotExist        = sdkerrors.Register(ModuleName, 502, "model does not exist")
	ErrVendorProductsDoNotExist = sdkerrors.Register(ModuleName, 504, "vendor products do not exist")

	// Model Version Error Codes.
	ErrSoftwareVersionStringInvalid = sdkerrors.Register(ModuleName, 511, "software version string invalid")
	ErrFirmwareDigestsInvalid       = sdkerrors.Register(ModuleName, 512, "firmware digests invalid")
	ErrCDVersionNumberInvalid       = sdkerrors.Register(ModuleName, 513, "CD version number invalid")
	ErrOtaURLInvalid                = sdkerrors.Register(ModuleName, 514, "OTA URL invalid")
	ErrOtaMissingInformation        = sdkerrors.Register(ModuleName, 515, "OTA missing information")
	ErrReleaseNotesURLInvalid       = sdkerrors.Register(ModuleName, 516, "release notes URL invalid")
	ErrModelVersionDoesNotExist     = sdkerrors.Register(ModuleName, 517, "model version does not exist")
	ErrNoModelVersionsExist         = sdkerrors.Register(ModuleName, 518, "no model versions exist")
	ErrModelVersionAlreadyExists    = sdkerrors.Register(ModuleName, 519, "model version already exists")
	ErrOtaURLCannotBeSet            = sdkerrors.Register(ModuleName, 520, "OTA URL cannot be set")
	ErrMaxSVLessThanMinSV           = sdkerrors.Register(ModuleName, 521, "max software version less than min software version")
	ErrLsfRevisionIsNotHigher       = sdkerrors.Register(ModuleName, 522, "LsfRevision should be greater then existing revision")
)

func NewErrModelAlreadyExists(vid interface{}, pid interface{}) error {
	return sdkerrors.Wrapf(ErrModelAlreadyExists,
		"Model associated with vid=%v and pid=%v already exists on the ledger",
		vid, pid)
}

func NewErrModelDoesNotExist(vid interface{}, pid interface{}) error {
	return sdkerrors.Wrapf(ErrModelDoesNotExist,
		"No model associated with vid=%v and pid=%v exist on the ledger",
		vid, pid)
}

func NewErrVendorProductsDoNotExist(vid interface{}) error {
	return sdkerrors.Wrapf(ErrVendorProductsDoNotExist,
		"No vendor products associated with vid=%v exist on the ledger",
		vid)
}

func NewErrSoftwareVersionStringInvalid(softwareVersion interface{}) error {
	return sdkerrors.Wrapf(ErrSoftwareVersionStringInvalid,
		"SoftwareVersionString %v is invalid. It should be greater then 1 and less then 64 character long",
		softwareVersion)
}

func NewErrFirmwareDigestsInvalid(firmwareDigests interface{}) error {
	return sdkerrors.Wrapf(ErrFirmwareDigestsInvalid,
		"firmwareDigests %v is of invalid length. Maximum length should be less then 512",
		firmwareDigests)
}

func NewErrCDVersionNumberInvalid(cdVersionNumber interface{}) error {
	return sdkerrors.Wrapf(ErrCDVersionNumberInvalid,
		"CDVersionNumber %v is invalid. It should be a 16 bit unsigned integer",
		cdVersionNumber)
}

func NewErrOtaURLInvalid(otaURL interface{}) error {
	return sdkerrors.Wrapf(ErrOtaURLInvalid,
		"OtaURL %v is invalid. Maximum length should be less then 256",
		otaURL)
}

func NewErrOtaMissingInformation() error {
	return sdkerrors.Wrap(ErrOtaMissingInformation,
		"OtaFileSize, OtaChecksum and OtaChecksumType are required if OtaUrl is provided")
}

func NewErrReleaseNotesURLInvalid(releaseNotesURL interface{}) error {
	return sdkerrors.Wrapf(ErrReleaseNotesURLInvalid,
		"ReleaseNotesURLInvalid %v is invalid. Maximum length should be less then 256",
		releaseNotesURL)
}

func NewErrModelVersionDoesNotExist(vid interface{}, pid interface{}, softwareVersion interface{}) error {
	return sdkerrors.Wrapf(ErrModelVersionDoesNotExist,
		"No model version associated with vid=%v, pid=%v and softwareVersion=%v exist on the ledger",
		vid, pid, softwareVersion)
}

func NewErrNoModelVersionsExist(vid interface{}, pid interface{}) error {
	return sdkerrors.Wrapf(ErrNoModelVersionsExist,
		"No versions associated with vid=%v and pid=%v exist on the ledger",
		vid, pid)
}

func NewErrModelVersionAlreadyExists(vid interface{}, pid interface{}, softwareVersion interface{}) error {
	return sdkerrors.Wrapf(ErrModelVersionAlreadyExists,
		"Model Version already exists on ledger with vid=%v pid=%v and softwareVersion=%v exist on the ledger",
		vid, pid, softwareVersion)
}

func NewErrOtaURLCannotBeSet(vid interface{}, pid interface{}, softwareVersion interface{}) error {
	return sdkerrors.Wrapf(ErrOtaURLCannotBeSet,
		"OTA URL cannot be set for model version associated with vid=%v, pid=%v "+
			"and softwareVersion=%v because OTA was not set for this model initially",
		vid, pid, softwareVersion)
}

func NewErrMaxSVLessThanMinSV(minApplicableSoftwareVersion interface{},
	maxApplicableSoftwareVersion interface{}) error {
	return sdkerrors.Wrapf(ErrMaxSVLessThanMinSV,
		"MaxApplicableSoftwareVersion %v is less than MinApplicableSoftwareVersion %v",
		maxApplicableSoftwareVersion, minApplicableSoftwareVersion)
}

func NewErrLsfRevisionIsNotHigher(previousLsfVersion interface{},
	currentLsfVersion interface{}) error {
	return sdkerrors.Wrapf(ErrLsfRevisionIsNotHigher,
		"LsfRevision %v should be greater then existing lsfRevision %v",
		currentLsfVersion, previousLsfVersion)
}
