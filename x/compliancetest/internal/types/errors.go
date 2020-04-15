package types

import sdk "github.com/cosmos/cosmos-sdk/types"

const (
	Codespace sdk.CodespaceType = ModuleName

	CodeTestingResultDoesNotExist sdk.CodeType = 201
)

func ErrTestingResultDoesNotExist() sdk.Error {
	return sdk.NewError(Codespace, CodeTestingResultDoesNotExist, "TestingResults do not exist")
}
