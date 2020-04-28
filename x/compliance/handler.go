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
	// check if sender has enough rights to certify model
	if err := checkZbCertificationRights(ctx, authzKeeper, msg.Signer, msg.CertificationType); err != nil {
		return err.Result()
	}

	var complianceInfo types.ComplianceInfo

	if keeper.IsComplianceInfoPresent(ctx, msg.CertificationType, msg.VID, msg.PID) {
		// Compliance record already exist. Cases:
		// 1) Only revocation is tracked on the ledger. We want to certify revoked compliance. The corresponding Model Info and test results are not required to be on the ledger.
		// 2) Compliance is tracked on ledger. We want to certify revoked compliance. `Else` branch was passed on first certification. So Model Info and test results are exists on the ledger.
		complianceInfo = keeper.GetComplianceInfo(ctx, msg.CertificationType, msg.VID, msg.PID)

		// if state changes on `certified` check that certification_date is after revocation_date
		if complianceInfo.State == types.Revoked {
			if msg.CertificationDate.Before(complianceInfo.Date) {
				return sdk.ErrInternal(
					fmt.Sprintf("The `certification_date`:%v must be after the current `date`:%v to certify model", msg.CertificationDate, complianceInfo.Date)).Result()
			}

			complianceInfo.UpdateComplianceInfo(msg.CertificationDate, msg.Reason)
		}
		// TODO: else allow setting different certification date?
	} else {
		// Compliance is tracked on ledger. There is no compliance record yet. The corresponding Model Info and test results must be present on ledger.
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

	// store compliance info
	keeper.SetComplianceInfo(ctx, complianceInfo)

	return sdk.Result{}
}

func handleMsgRevokeModel(ctx sdk.Context, keeper keeper.Keeper, authzKeeper authz.Keeper,
	msg types.MsgRevokeModel) sdk.Result {
	// check if sender has enough rights to revoke model
	if err := checkZbCertificationRights(ctx, authzKeeper, msg.Signer, msg.CertificationType); err != nil {
		return err.Result()
	}

	var complianceInfo types.ComplianceInfo

	if keeper.IsComplianceInfoPresent(ctx, msg.CertificationType, msg.VID, msg.PID) {
		// Compliance record already exist.
		complianceInfo = keeper.GetComplianceInfo(ctx, msg.CertificationType, msg.VID, msg.PID)

		// if state changes on `revoked` check that revocation_date is after certification_date
		if complianceInfo.State == types.Certified {
			if msg.RevocationDate.Before(complianceInfo.Date) {
				return sdk.ErrInternal(
					fmt.Sprintf("The `revocation_date`:%v must be after the `certification_date`:%v to revoke model", msg.RevocationDate, complianceInfo.Date)).Result()
			}

			complianceInfo.UpdateComplianceInfo(msg.RevocationDate, msg.Reason)
		}
		// TODO: else allow setting different revocation date?
	} else {
		// Only revocation is tracked on the ledger. There is no compliance record yet. The corresponding Model Info and test results are not required to be on the ledger.
		complianceInfo = types.NewRevokedComplianceInfo(
			msg.VID,
			msg.PID,
			msg.CertificationType,
			msg.RevocationDate,
			msg.Reason,
			msg.Signer,
		)
	}

	// store compliance info
	keeper.SetComplianceInfo(ctx, complianceInfo)

	return sdk.Result{}
}

func checkZbCertificationRights(ctx sdk.Context, authzKeeper authz.Keeper, signer sdk.AccAddress, certificationType types.CertificationType) sdk.Error {
	// rights are depend on certification type
	if certificationType == types.ZbCertificationType {
		// sender must have ZBCertificationCenter role to certify/revoke model
		if !authzKeeper.HasRole(ctx, signer, authz.ZBCertificationCenter) {
			return sdk.ErrUnauthorized(fmt.Sprintf("MsgCertifyModel/MsgRevokeMode transaction should be signed by an account with the %s role", authz.ZBCertificationCenter))
		}
	} else {
		return sdk.ErrUnknownRequest(fmt.Sprintf("Unexpected CertificationType: \"%s\". Supported types: [%s]", certificationType, ZbCertificationType))
	}
	return nil
}
