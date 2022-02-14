package keeper

import (
	"context"
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliance/types"
	dclauthtypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
	modeltypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/model/types"
)

func (k msgServer) CertifyModel(goCtx context.Context, msg *types.MsgCertifyModel) (*types.MsgCertifyModelResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	signerAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid Address: (%s)", err)
	}

	// check if sender has enough rights to certify model
	// sender must have CertificationCenter role to certify/revoke model
	if !k.dclauthKeeper.HasRole(ctx, signerAddr, dclauthtypes.CertificationCenter) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized,
			"MsgAddTestingResult transaction should be signed by an account with the %s role",
			dclauthtypes.CertificationCenter,
		)
	}

	// The corresponding Model Version must be present on ledger.
	modelVersion, found := k.modelKeeper.GetModelVersion(ctx, msg.Vid, msg.Pid, msg.SoftwareVersion)
	if !found {
		return nil, modeltypes.NewErrModelVersionDoesNotExist(msg.Vid, msg.Pid, msg.SoftwareVersion)
	}

	// check if softwareVersionString matches with what is stored for the given version
	if modelVersion.SoftwareVersionString != msg.SoftwareVersionString {
		return nil, types.NewErrModelVersionStringDoesNotMatch(msg.Vid, msg.Pid, msg.SoftwareVersion, msg.SoftwareVersionString)
	}

	complianceInfo, found := k.GetComplianceInfo(ctx, msg.Vid, msg.Pid, msg.SoftwareVersion, msg.CertificationType)
	// nolint:nestif
	if found {
		// Compliance record already exist. Cases:
		// 1) We want to re-certify compliance which is already certified now -> Error.
		// 2) We want to certify provisioned compliance. So no revocations were done earlier and thus certification
		// will be tracked on ledger. So we have to ensure that there are tesing results at first.
		// 3) We want to certify revoked compliance. Either earlier certification was done before revocation
		// and thus the corresponding test results are already on the ledger, or only revocation is tracked on the ledger
		// and so the test results are not required to be present.

		// check if compliance is already in certified state
		if complianceInfo.SoftwareVersionCertificationStatus == types.CodeCertified {
			return nil, types.NewErrAlreadyCertified(msg.Vid, msg.Pid, msg.SoftwareVersion, msg.CertificationType)
		}

		// if state changes on `certified` check that certification date is after provisional/revocation date
		newDate, err := time.Parse(time.RFC3339, msg.CertificationDate)
		if err != nil {
			return nil, types.NewErrInvalidTestDateFormat(msg.CertificationDate)
		}
		oldDate, err := time.Parse(time.RFC3339, complianceInfo.Date)
		if err != nil {
			return nil, types.NewErrInvalidTestDateFormat(complianceInfo.Date)
		}
		if newDate.Before(oldDate) {
			return nil, types.NewErrInconsistentDates(
				fmt.Sprintf("The `certification_date`:%v must be after the current `date`:%v to "+
					"certify model", msg.CertificationDate, complianceInfo.Date),
			)
		}

		complianceInfo.SetCertifiedStatus(msg.CertificationDate, msg.Reason)
	} else {
		// There is no compliance record yet. So certification will be tracked on ledger.

		complianceInfo = types.ComplianceInfo{
			Vid:                                msg.Vid,
			Pid:                                msg.Pid,
			SoftwareVersion:                    msg.SoftwareVersion,
			SoftwareVersionString:              msg.SoftwareVersionString,
			CertificationType:                  msg.CertificationType,
			Date:                               msg.CertificationDate,
			Reason:                             msg.Reason,
			Owner:                              msg.Signer,
			SoftwareVersionCertificationStatus: types.CodeCertified,
			History:                            []*types.ComplianceHistoryItem{},
			CDVersionNumber:                    msg.CDVersionNumber,
		}
	}

	// store compliance info
	k.SetComplianceInfo(ctx, complianceInfo)

	// update certified, revoked and provisional index
	certifiedModel := types.CertifiedModel{
		Vid:               msg.Vid,
		Pid:               msg.Pid,
		SoftwareVersion:   msg.SoftwareVersion,
		CertificationType: msg.CertificationType,
		Value:             true,
	}
	k.SetCertifiedModel(ctx, certifiedModel)
	revokedModel := types.RevokedModel{
		Vid:               msg.Vid,
		Pid:               msg.Pid,
		SoftwareVersion:   msg.SoftwareVersion,
		CertificationType: msg.CertificationType,
		Value:             false,
	}
	k.SetRevokedModel(ctx, revokedModel)
	provisionalModel := types.ProvisionalModel{
		Vid:               msg.Vid,
		Pid:               msg.Pid,
		SoftwareVersion:   msg.SoftwareVersion,
		CertificationType: msg.CertificationType,
		Value:             false,
	}
	k.SetProvisionalModel(ctx, provisionalModel)

	return &types.MsgCertifyModelResponse{}, nil
}
