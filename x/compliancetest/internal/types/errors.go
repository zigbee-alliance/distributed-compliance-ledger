package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	Codespace sdk.CodespaceType = ModuleName

	CodeTestingResultsDoNotExist sdk.CodeType = 201
)

func ErrTestingResultDoesNotExist(vid interface{}, pid interface{}) sdk.Error {
	return sdk.NewError(Codespace, CodeTestingResultsDoNotExist,
		fmt.Sprintf("No testing results about the model with vid=%v and pid=%v on the ledger", vid, pid))
}
