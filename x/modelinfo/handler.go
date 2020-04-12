package modelinfo

import (
	"fmt"

	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/authz"

	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/modelinfo/internal/keeper"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/modelinfo/internal/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewHandler(keeper keeper.Keeper, authzKeeper authz.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case types.MsgAddModelInfo:
			return handleMsgAddModelInfo(ctx, keeper, authzKeeper, msg)
		case types.MsgUpdateModelInfo:
			return handleMsgUpdateModelInfo(ctx, keeper, authzKeeper, msg)
			/*		case types.MsgDeleteModelInfo:
					return handleMsgDeleteModelInfo(ctx, keeper, authzKeeper, msg)*/
		default:
			errMsg := fmt.Sprintf("unrecognized nameservice Msg type: %v", msg.Type())
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

func handleMsgAddModelInfo(ctx sdk.Context, keeper keeper.Keeper, authzKeeper authz.Keeper,
	msg types.MsgAddModelInfo) sdk.Result {
	if keeper.IsModelInfoPresent(ctx, msg.VID, msg.PID) {
		return types.ErrModelInfoAlreadyExists().Result()
	}

	if err := checkAddModelRights(ctx, authzKeeper, msg.Signer); err != nil {
		return err.Result()
	}

	modelInfo := types.NewModelInfo(
		msg.VID,
		msg.PID,
		msg.CID,
		msg.Name,
		msg.Signer,
		msg.Description,
		msg.SKU,
		msg.FirmwareVersion,
		msg.HardwareVersion,
		msg.Custom,
		msg.CertificateID,
		msg.CertifiedDate,
		msg.TisOrTrpTestingCompleted,
	)

	keeper.SetModelInfo(ctx, modelInfo)

	return sdk.Result{}
}

func handleMsgUpdateModelInfo(ctx sdk.Context, keeper keeper.Keeper, authzKeeper authz.Keeper,
	msg types.MsgUpdateModelInfo) sdk.Result {
	if !keeper.IsModelInfoPresent(ctx, msg.VID, msg.PID) {
		return types.ErrModelInfoDoesNotExist().Result()
	}

	modelInfo := keeper.GetModelInfo(ctx, msg.VID, msg.PID)

	if err := checkUpdateModelRights(modelInfo.Owner, msg.Signer); err != nil {
		return err.Result()
	}

	modelInfo.CID = msg.NewCID
	modelInfo.Description = msg.NewDescription
	modelInfo.Custom = msg.NewCustom
	modelInfo.CertificateID = msg.NewCertificateID
	modelInfo.CertifiedDate = msg.NewCertifiedDate
	modelInfo.TisOrTrpTestingCompleted = msg.NewTisOrTrpTestingCompleted

	keeper.SetModelInfo(ctx, modelInfo)

	return sdk.Result{}
}

func handleMsgDeleteModelInfo(ctx sdk.Context, keeper keeper.Keeper, authzKeeper authz.Keeper,
	msg types.MsgDeleteModelInfo) sdk.Result {
	if !keeper.IsModelInfoPresent(ctx, msg.VID, msg.PID) {
		return types.ErrModelInfoDoesNotExist().Result()
	}

	modelInfo := keeper.GetModelInfo(ctx, msg.VID, msg.PID)

	if err := checkUpdateModelRights(modelInfo.Owner, msg.Signer); err != nil {
		return err.Result()
	}

	keeper.DeleteModelInfo(ctx, msg.VID, msg.PID)

	return sdk.Result{}
}

func checkAddModelRights(ctx sdk.Context, authzKeeper authz.Keeper, signer sdk.AccAddress) sdk.Error {
	if !authzKeeper.HasRole(ctx, signer, authz.Vendor) {
		return sdk.ErrUnauthorized("MsgUpdateModelInfo tx should be signed either by vendor")
	}

	return nil
}

func checkUpdateModelRights(owner sdk.AccAddress, signer sdk.AccAddress) sdk.Error {
	if !signer.Equals(owner) {
		return sdk.ErrUnauthorized("MsgUpdateModelInfo tx should be signed by owner")
	}

	return nil
}
