package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	Codespace sdk.CodespaceType = ModuleName

	CodeDeviceComplianceAlreadyExists sdk.CodeType = 301
	CodeDeviceComplianceDoesNotExist  sdk.CodeType = 302
)

func ErrDeviceComplianceAlreadyExists(vid interface{}, pid interface{}) sdk.Error {
	return sdk.NewError(Codespace, CodeDeviceComplianceAlreadyExists,
		fmt.Sprintf("The model with vid=%v and pid=%v is already certified.", vid, pid))
}

func ErrDeviceComplianceoDoesNotExist(vid interface{}, pid interface{}) sdk.Error {
	return sdk.NewError(Codespace, CodeDeviceComplianceDoesNotExist,
		fmt.Sprintf("No certification information about the model with vid=%v and pid=%v on the ledger. "+
			"This means that the model is either not certified yet or certified by default (off-ledger).", vid, pid))
}
