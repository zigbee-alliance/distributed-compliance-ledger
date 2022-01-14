package keeper

import (
	"context"
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliance/types"
	compliancetesttypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/compliancetest/types"
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
		// 1) Only revocation is tracked on the ledger. We want to certify revoked compliance.
		// The corresponding Model Info and test results are not required to be on the ledger.
		// 2) Compliance is tracked on ledger. We want to certify revoked compliance.
		// `Else` branch was passed on first certification. So Model Info and test results exist on the ledger.

		// check if certification is already done
		if complianceInfo.SoftwareVersionCertificationStatus == types.CodeCertified {
			return nil, types.NewErrAlreadyCertified(msg.Vid, msg.Pid, msg.SoftwareVersion, msg.CertificationType)
		} else {
			// if state changes on `certified` check that certification_date is after revocation_date
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
		}
	} else {
		// Compliance is tracked on ledger. There is no compliance record yet.

		// The corresponding test results must be present on ledger.
		_, found = k.compliancetestKeeper.GetTestingResults(ctx, msg.Vid, msg.Pid, msg.SoftwareVersion)
		if !found {
			return nil, compliancetesttypes.NewErrTestingResultsDoNotExist(msg.Vid, msg.Pid, msg.SoftwareVersion)
		}

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

	// update certified/revoked index
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
