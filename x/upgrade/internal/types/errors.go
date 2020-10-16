package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SDK error codes
const (
	// Upgrade error codes
	ErrInvalidRequest sdk.CodeType = 0
	ErrUnknownRequest sdk.CodeType = 1
	ErrInvalidProposalContent sdk.CodeType = 2
	ErrJSONMarshal sdk.CodeType = 3
	ErrJSONUnmarshal sdk.CodeType = 4

	// CodespaceRoot is a codespace for error codes in the upgrade module only.
	CodespaceUpgrade sdk.CodespaceType = ModuleName
)

// NewError - create an error with the module's codespace.
func NewError(code sdk.CodeType, format string, args ...interface{}) sdk.Error {
	return sdk.NewError(CodespaceUpgrade, code, format, args...)
}
