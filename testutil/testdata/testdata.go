package testdata

import (
	sdktestdata "github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func GenerateAccAddress() sdk.AccAddress {
	_, _, accAddress := sdktestdata.KeyTestPubAddr()
	return accAddress
}
