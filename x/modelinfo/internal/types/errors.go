package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	Codespace sdk.CodespaceType = ModuleName

	CodeModelInfoAlreadyExists   sdk.CodeType = 501
	CodeModelInfoDoesNotExist    sdk.CodeType = 502
	CodeVendorProductsDoNotExist sdk.CodeType = 503
)

func ErrModelInfoAlreadyExists(vid interface{}, pid interface{}) sdk.Error {
	return sdk.NewError(Codespace, CodeModelInfoAlreadyExists,
		fmt.Sprintf("Model info associated with the vid=%v and pid=%v already exists on the ledger", vid, pid))
}

func ErrModelInfoDoesNotExist(vid interface{}, pid interface{}) sdk.Error {
	return sdk.NewError(Codespace, CodeModelInfoDoesNotExist,
		fmt.Sprintf("No model info associated with the vid=%v and pid=%v on the ledger", vid, pid))
}

func ErrVendorProductsDoNotExist(vid interface{}) sdk.Error {
	return sdk.NewError(Codespace, CodeVendorProductsDoNotExist,
		fmt.Sprintf("No models associated with the vid=%v on the ledger", vid))
}
