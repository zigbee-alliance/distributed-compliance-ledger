package keeper

import (
	"context"
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	dclcompltypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/compliance"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliance/types"
	dclauthtypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
	modeltypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/model/types"
)

func (k msgServer) RevokeModel(goCtx context.Context, msg *types.MsgRevokeModel) (*types.MsgRevokeModelResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	signerAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid Address: (%s)", err)
	}

	// check if sender has enough rights to revoke model
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
	//nolint:nestif
	if found {
		// Compliance record already exist. Cases:
		// 1) We want to re-revoke compliance which is already in revoked state now -> Error.
		// 2) We want to revoke certified or provisioned compliance.

		// check if compliance is already in revoked state
		if complianceInfo.SoftwareVersionCertificationStatus == dclcompltypes.CodeRevoked {
			return nil, types.NewErrAlreadyRevoked(msg.Vid, msg.Pid, msg.SoftwareVersion, msg.CertificationType)
		}
		// if state changes on `revoked` check that revocation date is after certification/provisional date
		newDate, err := time.Parse(time.RFC3339, msg.RevocationDate)
		if err != nil {
			return nil, types.NewErrInvalidTestDateFormat(msg.RevocationDate)
		}
		oldDate, err := time.Parse(time.RFC3339, complianceInfo.Date)
		if err != nil {
			return nil, types.NewErrInvalidTestDateFormat(complianceInfo.Date)
		}

		if newDate.Before(oldDate) {
			return nil, types.NewErrInconsistentDates(
				fmt.Sprintf("The `revocation_date`:%v must be after the `certification_date`:%v to "+
					"revoke model", msg.RevocationDate, complianceInfo.Date),
			)
		}

		complianceInfo.SetRevokedStatus(msg.RevocationDate, msg.Reason)

		// remove compliance info from the entity Device Compliance Info
		deviceSoftwareCompliance, found := k.GetDeviceSoftwareCompliance(ctx, complianceInfo.CDCertificateId)
		if found {
			index, found := deviceSoftwareCompliance.IsComplianceInfoExist(msg.Vid, msg.Pid, msg.SoftwareVersion)
			if found {
				deviceSoftwareCompliance.RemoveComplianceInfo(index)
			}
		}
		k.SetDeviceSoftwareCompliance(ctx, deviceSoftwareCompliance)

		// If we don't have compliance info in Device Software Compliance - we should delete this Device Software Compliance
		if len(deviceSoftwareCompliance.ComplianceInfo) == 0 {
			k.RemoveDeviceSoftwareCompliance(ctx, deviceSoftwareCompliance.CDCertificateId)
		}

		// update certified and provisional index
		certifiedModel := types.CertifiedModel{
			Vid:               msg.Vid,
			Pid:               msg.Pid,
			SoftwareVersion:   msg.SoftwareVersion,
			CertificationType: msg.CertificationType,
			Value:             false,
		}
		k.SetCertifiedModel(ctx, certifiedModel)
		provisionalModel := types.ProvisionalModel{
			Vid:               msg.Vid,
			Pid:               msg.Pid,
			SoftwareVersion:   msg.SoftwareVersion,
			CertificationType: msg.CertificationType,
			Value:             false,
		}
		k.SetProvisionalModel(ctx, provisionalModel)
	} else {
		// There is no compliance record yet. So only revocation will be tracked on ledger.

		complianceInfo = dclcompltypes.ComplianceInfo{
			Vid:                                msg.Vid,
			Pid:                                msg.Pid,
			SoftwareVersion:                    msg.SoftwareVersion,
			SoftwareVersionString:              msg.SoftwareVersionString,
			CertificationType:                  msg.CertificationType,
			Date:                               msg.RevocationDate,
			Reason:                             msg.Reason,
			Owner:                              msg.Signer,
			SoftwareVersionCertificationStatus: dclcompltypes.CodeRevoked,
			History:                            []*dclcompltypes.ComplianceHistoryItem{},
			CDVersionNumber:                    msg.CDVersionNumber,
		}
	}

	// store compliance info
	k.SetComplianceInfo(ctx, complianceInfo)

	// update revoked index
	revokedModel := types.RevokedModel{
		Vid:               msg.Vid,
		Pid:               msg.Pid,
		SoftwareVersion:   msg.SoftwareVersion,
		CertificationType: msg.CertificationType,
		Value:             true,
	}
	k.SetRevokedModel(ctx, revokedModel)

	return &types.MsgRevokeModelResponse{}, nil
}
