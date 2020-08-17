package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	DefaultCodespace sdk.CodespaceType = ModuleName

	CodeAccountAlreadyExists                  sdk.CodeType = 101
	CodeAccountDoesNotExist                   sdk.CodeType = 102
	CodePendingAccountAlreadyExists           sdk.CodeType = 103
	CodePendingAccountDoesNotExist            sdk.CodeType = 104
	CodePendingAccountRevocationAlreadyExists sdk.CodeType = 105
	CodePendingAccountRevocationDoesNotExist  sdk.CodeType = 106
)

func ErrAccountAlreadyExist(address interface{}) sdk.Error {
	return sdk.NewError(DefaultCodespace, CodeAccountAlreadyExists,
		fmt.Sprintf("Account associated with the address=%v already exists on the ledger", address))
}

func ErrAccountDoesNotExist(address interface{}) sdk.Error {
	return sdk.NewError(DefaultCodespace, CodeAccountDoesNotExist,
		fmt.Sprintf("No account associated with the address=%v on the ledger", address))
}

func ErrPendingAccountAlreadyExist(address interface{}) sdk.Error {
	return sdk.NewError(DefaultCodespace, CodePendingAccountAlreadyExists,
		fmt.Sprintf("Pending account associated with the address=%v already exists on the ledger", address))
}

func ErrPendingAccountDoesNotExist(address interface{}) sdk.Error {
	return sdk.NewError(DefaultCodespace, CodePendingAccountDoesNotExist,
		fmt.Sprintf("No pending account associated with the address=%v on the ledger", address))
}

func ErrPendingAccountRevocationAlreadyExist(address interface{}) sdk.Error {
	return sdk.NewError(DefaultCodespace, CodePendingAccountRevocationAlreadyExists,
		fmt.Sprintf("Pending account revocation associated with the address=%v already exists on the ledger", address))
}

func ErrPendingAccountRevocationDoesNotExist(address interface{}) sdk.Error {
	return sdk.NewError(DefaultCodespace, CodePendingAccountRevocationDoesNotExist,
		fmt.Sprintf("No pending account revocation associated with the address=%v on the ledger", address))
}
