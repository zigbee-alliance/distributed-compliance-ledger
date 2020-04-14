package types

import sdk "github.com/cosmos/cosmos-sdk/types"

const (
	Codespace sdk.CodespaceType = ModuleName

	CodeTestingResultAlreadyExists sdk.CodeType = 201
	CodeTestingResultDoesNotExist  sdk.CodeType = 202
)

func ErrTestingResultAlreadyExists() sdk.Error {
	return sdk.NewError(Codespace, CodeTestingResultAlreadyExists, "TestingResult already exists")
}

func ErrTestingResultDoesNotExist() sdk.Error {
	return sdk.NewError(Codespace, CodeTestingResultDoesNotExist, "TestingResults do not exist")
}
