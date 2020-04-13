package compliancetest

import (
	"fmt"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/authz"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/modelinfo"

	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliancetest/internal/keeper"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliancetest/internal/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewHandler(keeper keeper.Keeper, modelinfoKeeper modelinfo.Keeper, authzKeeper authz.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case types.MsgAddTestingResult:
			return handleMsgAddTestingResult(ctx, keeper, modelinfoKeeper, authzKeeper, msg)
		default:
			errMsg := fmt.Sprintf("unrecognized compliancetest Msg type: %v", msg.Type())
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

func handleMsgAddTestingResult(ctx sdk.Context, keeper keeper.Keeper, modelinfoKeeper modelinfo.Keeper, authzKeeper authz.Keeper,
	msg types.MsgAddTestingResult) sdk.Result {
	if err := checkAddTestingResultRights(ctx, authzKeeper, msg.Signer); err != nil {
		return err.Result()
	}

	if !modelinfoKeeper.IsModelInfoPresent(ctx, msg.VID, msg.PID) {
		return modelinfo.ErrModelInfoDoesNotExist().Result()
	}

	if keeper.IsTestingResultPresents(ctx, msg.VID, msg.PID, msg.Signer) {
		return types.ErrTestingResultAlreadyExists().Result()
	}

	testingResult := types.NewTestingResult(
		msg.VID,
		msg.PID,
		msg.TestResult,
		msg.Signer,
	)

	if err := keeper.AddTestingResult(ctx, testingResult); err != nil {
		return err.Result()
	}

	return sdk.Result{}
}

func checkAddTestingResultRights(ctx sdk.Context, authzKeeper authz.Keeper, signer sdk.AccAddress) sdk.Error {
	if !authzKeeper.HasRole(ctx, signer, authz.TestHouse) {
		return sdk.ErrUnauthorized("MsgAddTestingResult transaction should be signed by an account with the TestHouse role")
	}

	return nil
}
