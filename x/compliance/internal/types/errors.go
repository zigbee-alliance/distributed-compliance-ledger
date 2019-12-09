package types

import sdk "github.com/cosmos/cosmos-sdk/types"

const (
	DefaultCodespace sdk.CodespaceType = ModuleName

	CodeDeviceAlreadyExists sdk.CodeType = 101
	CodeDeviceDoesNotExist  sdk.CodeType = 102
)

func ErrDeviceAlreadyExists(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeDeviceAlreadyExists, "ModelInfo already exists")
}

func ErrDeviceDoesNotExist(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeDeviceDoesNotExist, "ModelInfo does not exist")
}
