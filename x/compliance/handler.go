package compliance

import (
	"fmt"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliancetest"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/modelinfo"

	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/authz"

	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliance/internal/keeper"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliance/internal/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewHandler(keeper keeper.Keeper, modelinfoKeeper modelinfo.Keeper,
	compliancetestKeeper compliancetest.Keeper, authzKeeper authz.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case types.MsgCertifyModel:
			return handleMsgCertifyModel(ctx, keeper, modelinfoKeeper, compliancetestKeeper, authzKeeper, msg)
		default:
			errMsg := fmt.Sprintf("unrecognized nameservice Msg type: %v", msg.Type())
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

func handleMsgCertifyModel(ctx sdk.Context, keeper keeper.Keeper, modelinfoKeeper modelinfo.Keeper,
	compliancetestKeeper compliancetest.Keeper, authzKeeper authz.Keeper,
	msg types.MsgCertifyModel) sdk.Result {

	if err := msg.ValidateBasic(); err != nil {
		return err.Result()
	}

	//if keeper.IsCertifiedModelPresent(ctx, msg.VID, msg.PID) {
	//	return types.ErrDeviceComplianceAlreadyExists(msg.VID, msg.PID).Result()
	//}

	if !modelinfoKeeper.IsModelInfoPresent(ctx, msg.VID, msg.PID) {
		return modelinfo.ErrModelInfoDoesNotExist(msg.VID, msg.PID).Result()
	}

	if !compliancetestKeeper.IsTestingResultsPresents(ctx, msg.VID, msg.PID) {
		return compliancetest.ErrTestingResultDoesNotExist(msg.VID, msg.PID).Result()
	}

	if err := checkCertifyModelsRights(ctx, authzKeeper, msg.Signer, msg.CertificationType); err != nil {
		return err.Result()
	}

	certifiedModel := types.NewCertifiedModel(
		msg.VID,
		msg.PID,
		msg.CertificationDate,
		msg.CertificationType,
		msg.Signer,
	)

	keeper.SetCertifiedModel(ctx, certifiedModel)

	return sdk.Result{}
}

func checkCertifyModelsRights(ctx sdk.Context, authzKeeper authz.Keeper, signer sdk.AccAddress, certificationType string) sdk.Error {
	if len(certificationType) == 0 || certificationType == types.ZbCertificationType { // certification type is empty or ZbCertificationType
		if !authzKeeper.HasRole(ctx, signer, authz.ZBCertificationCenter) {
			return sdk.ErrUnauthorized("MsgCertifyModel transaction should be signed by an account with the ZBCertificationCenter role")
		}
	}

	return nil
}
