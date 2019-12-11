package types

import sdk "github.com/cosmos/cosmos-sdk/types"

const (
	DefaultCodespace sdk.CodespaceType = ModuleName

	CodeModelInfoAlreadyExists sdk.CodeType = 101
	CodeModelInfoDoesNotExist  sdk.CodeType = 102
)

func ErrModelInfoAlreadyExists(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeModelInfoAlreadyExists, "ModelInfo already exists")
}

func ErrModelInfoDoesNotExist(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeModelInfoDoesNotExist, "ModelInfo does not exist")
}
