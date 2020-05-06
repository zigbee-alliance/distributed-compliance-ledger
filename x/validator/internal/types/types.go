package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Last validator power, needed for validator set update logic
type LastValidatorPower struct {
	Address sdk.ValAddress
	Power   int64
}
