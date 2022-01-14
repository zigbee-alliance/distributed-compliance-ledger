package ante

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Sets Infinite Gas Meter instead of default one in SetUpContextDecorator.
type InfiniteGasSetUpContextDecorator struct{}

func NewInfiniteGasSetUpContextDecorator() InfiniteGasSetUpContextDecorator {
	return InfiniteGasSetUpContextDecorator{}
}

func (sud InfiniteGasSetUpContextDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	newCtx = ctx.WithGasMeter(sdk.NewInfiniteGasMeter())
	return next(newCtx, tx, simulate)
}
