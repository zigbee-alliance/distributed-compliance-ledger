package types

import sdk "github.com/cosmos/cosmos-sdk/types"

const (
	Codespace sdk.CodespaceType = ModuleName

	CodeDeviceComplianceAlreadyExists sdk.CodeType = 301
	CodeDeviceComplianceDoesNotExist  sdk.CodeType = 302
)

func ErrDeviceComplianceAlreadyExists() sdk.Error {
	return sdk.NewError(Codespace, CodeDeviceComplianceAlreadyExists, "CertifiedModel already exists")
}

func ErrDeviceComplianceoDoesNotExist() sdk.Error {
	return sdk.NewError(Codespace, CodeDeviceComplianceDoesNotExist, "CertifiedModel does not exist")
}
