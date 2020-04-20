package compliance

import (
	"fmt"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/authz"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliance/internal/keeper"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliance/internal/types"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliancetest"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/modelinfo"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewHandler(keeper keeper.Keeper, modelinfoKeeper modelinfo.Keeper,
	compliancetestKeeper compliancetest.Keeper, authzKeeper authz.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case types.MsgCertifyModel:
			return handleMsgCertifyModel(ctx, keeper, modelinfoKeeper, compliancetestKeeper, authzKeeper, msg)
		case types.MsgRevokeModel:
			return handleMsgRevokeModel(ctx, keeper, authzKeeper, msg)
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

	if err := checkZbCertificationRights(ctx, authzKeeper, msg.Signer, msg.CertificationType); err != nil {
		return err.Result()
	}

	var complianceInfo types.ComplianceInfo

	if keeper.IsComplianceInfoPresent(ctx, msg.CertificationType, msg.VID, msg.PID) {
		complianceInfo = keeper.GetComplianceInfo(ctx, msg.CertificationType, msg.VID, msg.PID)

		if !complianceInfo.Owner.Equals(msg.Signer) {
			return sdk.ErrUnauthorized("MsgCertifyModel transaction should be signed by the owner for editing of the existing record").Result()
		}

		if complianceInfo.State == types.Revoked {
			if msg.CertificationDate.Before(complianceInfo.Date) {
				return sdk.ErrInternal(
					fmt.Sprintf("The `certification_date`:%v must be after the current `date`:%v to certify model", msg.CertificationDate, complianceInfo.Date)).Result()
			}

			complianceInfo.UpdateComplianceInfo(msg.CertificationDate, msg.Reason)
		}
		// TODO: else allow setting different certification date?
	} else {
		if !modelinfoKeeper.IsModelInfoPresent(ctx, msg.VID, msg.PID) {
			return modelinfo.ErrModelInfoDoesNotExist(msg.VID, msg.PID).Result()
		}

		if !compliancetestKeeper.IsTestingResultsPresents(ctx, msg.VID, msg.PID) {
			return compliancetest.ErrTestingResultDoesNotExist(msg.VID, msg.PID).Result()
		}

		complianceInfo = types.NewCertifiedComplianceInfo(
			msg.VID,
			msg.PID,
			msg.CertificationType,
			msg.CertificationDate,
			msg.Reason,
			msg.Signer,
		)
	}

	keeper.SetComplianceInfo(ctx, complianceInfo)

	return sdk.Result{}
}

func handleMsgRevokeModel(ctx sdk.Context, keeper keeper.Keeper, authzKeeper authz.Keeper,
	msg types.MsgRevokeModel) sdk.Result {

	if err := msg.ValidateBasic(); err != nil {
		return err.Result()
	}

	if err := checkZbCertificationRights(ctx, authzKeeper, msg.Signer, msg.CertificationType); err != nil {
		return err.Result()
	}

	var complianceInfo types.ComplianceInfo

	if keeper.IsComplianceInfoPresent(ctx, msg.CertificationType, msg.VID, msg.PID) {
		complianceInfo = keeper.GetComplianceInfo(ctx, msg.CertificationType, msg.VID, msg.PID)

		if !complianceInfo.Owner.Equals(msg.Signer) {
			return sdk.ErrUnauthorized("MsgRevokeModel transaction should be signed by the owner for editing of the existing record").Result()
		}

		if complianceInfo.State == types.Certified {

			if msg.RevocationDate.Before(complianceInfo.Date) {
				return sdk.ErrInternal(
					fmt.Sprintf("The `revocation_date`:%v must be after the `certification_date`:%v to revoke model", msg.RevocationDate, complianceInfo.Date)).Result()
			}

			complianceInfo.UpdateComplianceInfo(msg.RevocationDate, msg.Reason)
		}
		// TODO: else allow setting different revocation date?
	} else {
		complianceInfo = types.NewRevokedComplianceInfo(
			msg.VID,
			msg.PID,
			msg.CertificationType,
			msg.RevocationDate,
			msg.Reason,
			msg.Signer,
		)
	}

	keeper.SetComplianceInfo(ctx, complianceInfo)

	return sdk.Result{}
}

func checkZbCertificationRights(ctx sdk.Context, authzKeeper authz.Keeper, signer sdk.AccAddress, certificationType types.CertificationType) sdk.Error {
	if certificationType == types.EmptyCertificationType || certificationType == types.ZbCertificationType { // certification type is empty or ZbCertificationType
		if !authzKeeper.HasRole(ctx, signer, authz.ZBCertificationCenter) {
			return sdk.ErrUnauthorized("MsgCertifyModel/MsgRevokeModel transaction should be signed by an account with the ZBCertificationCenter role")
		}
	}

	return nil
}
