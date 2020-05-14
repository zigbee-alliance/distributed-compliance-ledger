package compliancetest

//nolint:goimports
import (
	"fmt"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/auth"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/modelinfo"

	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliancetest/internal/keeper"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliancetest/internal/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewHandler(keeper keeper.Keeper, modelinfoKeeper modelinfo.Keeper, authzKeeper auth.Keeper) sdk.Handler {
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

func handleMsgAddTestingResult(ctx sdk.Context, keeper keeper.Keeper, modelinfoKeeper modelinfo.Keeper,
	authzKeeper authz.Keeper, msg types.MsgAddTestingResult) sdk.Result {
	// check if sender has enough rights to add testing results
	if err := checkAddTestingResultRights(ctx, authzKeeper, msg.Signer); err != nil {
		return err.Result()
	}

	// check that corresponding model exists on the ledger
	if !modelinfoKeeper.IsModelInfoPresent(ctx, msg.VID, msg.PID) {
		return modelinfo.ErrModelInfoDoesNotExist(msg.VID, msg.PID).Result()
	}

	testingResult := types.NewTestingResult(
		msg.VID,
		msg.PID,
		msg.Signer,
		msg.TestResult,
		msg.TestDate,
	)

	// store testing results. it extends existing value if testing results already exists
	keeper.AddTestingResult(ctx, testingResult)

	return sdk.Result{}
}

func checkAddTestingResultRights(ctx sdk.Context, authzKeeper auth.Keeper, signer sdk.AccAddress) sdk.Error {
	// sender must have TestHouse role to add new model
	if !authzKeeper.HasRole(ctx, signer, authz.TestHouse) {
		return sdk.ErrUnauthorized(fmt.Sprintf(
			"MsgAddTestingResult transaction should be signed by an account with the %s role", authz.TestHouse))
	}

	return nil
}
