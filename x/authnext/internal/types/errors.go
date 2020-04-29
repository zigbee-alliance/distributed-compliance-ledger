package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	Codespace sdk.CodespaceType = ModuleName

	CodeAccountDoesNotExist sdk.CodeType = 501
)

func ErrAccountDoesNotExist(address interface{}) sdk.Error {
	return sdk.NewError(Codespace, CodeAccountDoesNotExist,
		fmt.Sprintf("No account associated with the address=%v on the ledger", address))
}
