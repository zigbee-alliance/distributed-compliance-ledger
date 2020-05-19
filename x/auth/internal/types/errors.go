package types

//nolint:goimports
import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	DefaultCodespace sdk.CodespaceType = ModuleName

	CodeAccountAlreadyExist        sdk.CodeType = 101
	CodeAccountDoesNotExist        sdk.CodeType = 102
	CodePendingAccountDoesNotExist sdk.CodeType = 103
)

func ErrAccountAlreadyExistExist(address interface{}) sdk.Error {
	return sdk.NewError(DefaultCodespace, CodeAccountAlreadyExist,
		fmt.Sprintf("Account associated with the address=%v already exists on the ledger", address))
}

func ErrAccountDoesNotExist(address interface{}) sdk.Error {
	return sdk.NewError(DefaultCodespace, CodeAccountDoesNotExist,
		fmt.Sprintf("No account associated with the address=%v on the ledger", address))
}

func ErrPendingAccountDoesNotExist(address interface{}) sdk.Error {
	return sdk.NewError(DefaultCodespace, CodePendingAccountDoesNotExist,
		fmt.Sprintf("No pending account associated with the address=%v on the ledger", address))
}
