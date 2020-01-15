package types

import sdk "github.com/cosmos/cosmos-sdk/types"

const (
	DefaultCodespace sdk.CodespaceType = ModuleName

	CodeAccountRolesDoesNotExist sdk.CodeType = 101
)

func ErrAccountRolesDoesNotExist(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeAccountRolesDoesNotExist, "AccountRoles does not exist")
}
