package types

import sdk "github.com/cosmos/cosmos-sdk/types"

const (
	Codespace sdk.CodespaceType = ModuleName

	CodeModelInfoAlreadyExists sdk.CodeType = 101
	CodeModelInfoDoesNotExist  sdk.CodeType = 102
)

func ErrModelInfoAlreadyExists() sdk.Error {
	return sdk.NewError(Codespace, CodeModelInfoAlreadyExists, "ModelInfo already exists")
}

func ErrModelInfoDoesNotExist() sdk.Error {
	return sdk.NewError(Codespace, CodeModelInfoDoesNotExist, "ModelInfo does not exist")
}
