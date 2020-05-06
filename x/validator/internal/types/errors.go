package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"strings"
)

const (
	Codespace sdk.CodespaceType = ModuleName

	CodeValidatorOperatorAddressExist sdk.CodeType = 601
	CodeValidatorPubKeyExist          sdk.CodeType = 602
	CodeValidatorDoesNotExist         sdk.CodeType = 603
)

func ErrValidatorOperatorAddressExists(address interface{}) sdk.Error {
	return sdk.NewError(Codespace, CodeValidatorOperatorAddressExist,
		fmt.Sprintf("Validator associated with the operator_address=%v already exists on the ledger", address))
}

func ErrValidatorPubKeyExists(pubkey interface{}) sdk.Error {
	return sdk.NewError(Codespace, CodeValidatorPubKeyExist,
		fmt.Sprintf("Validator associated with the public_key=%v already exists on the ledger", pubkey))
}

func ErrValidatorDoesNotExist(address interface{}) sdk.Error {
	return sdk.NewError(Codespace, CodeValidatorDoesNotExist,
		fmt.Sprintf("No validator associated with the operator_address=%v on the ledger", address))
}

func ErrValidatorPubKeyTypeNotSupported(keyType string, supportedTypes []string) sdk.Error {
	return sdk.ErrUnknownRequest(
		fmt.Sprintf("Validator pubkey type \"%s\" is not supported. Supported types: [%s]", keyType, strings.Join(supportedTypes, ",")))
}
