package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	Codespace sdk.CodespaceType = ModuleName

	CodeValidatorOperatorAddressExist sdk.CodeType = 601
	CodeValidatorDoesNotExist         sdk.CodeType = 602
)

func ErrValidatorExists(address interface{}) sdk.Error {
	return sdk.NewError(Codespace, CodeValidatorOperatorAddressExist,
		fmt.Sprintf("Validator associated with the validator_address=%v already exists on the ledger", address))
}

func ErrValidatorDoesNotExist(address interface{}) sdk.Error {
	return sdk.NewError(Codespace, CodeValidatorDoesNotExist,
		fmt.Sprintf("No validator associated with the operator_address=%v on the ledger", address))
}
