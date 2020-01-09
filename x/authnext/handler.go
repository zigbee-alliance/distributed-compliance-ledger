package authnext

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewHandler(accKeeper AccountKeeper, cdc *codec.Codec) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		errMsg := fmt.Sprintf("unrecognized authnext Msg type: %v", msg.Type())
		return sdk.ErrUnknownRequest(errMsg).Result()
	}
}
