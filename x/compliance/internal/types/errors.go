package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	Codespace sdk.CodespaceType = ModuleName

	CodeComplianceInfoDoesNotExist sdk.CodeType = 301
)

func ErrComplianceInfoDoesNotExist(vid interface{}, pid interface{}, certificationType interface{}) sdk.Error {
	return sdk.NewError(Codespace, CodeComplianceInfoDoesNotExist,
		fmt.Sprintf("No certification information about the model with vid=%v, pid=%v and certification_type=%v on the ledger. "+
			"This means that the model is either not certified yet or certified by default (off-ledger).", vid, pid, certificationType))
}
