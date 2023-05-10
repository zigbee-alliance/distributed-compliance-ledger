package keeper

import (
	"context"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliance/types"
	dclauthtypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
	modeltypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/model/types"
)

func (k msgServer) UpdateComplianceInfo(goCtx context.Context, msg *types.MsgUpdateComplianceInfo) (*types.MsgUpdateComplianceInfoResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// check if signer has enough rights to update model
	signerAddr, err := sdk.AccAddressFromBech32(msg.Creator)

	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid Address: (%s)", err)
	}

	if !k.dclauthKeeper.HasRole(ctx, signerAddr, dclauthtypes.CertificationCenter) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "%s transaction should be "+
			"signed by an account with the %s role", "MsgUpdateComplianceInfo", dclauthtypes.CertificationCenter)
	}

	complianceInfo, isFound := k.GetComplianceInfo(ctx, msg.Vid, msg.Pid, msg.SoftwareVersion, msg.CertificationType)

	if !isFound {
		return nil, types.NewErrComplianceInfoDoesNotExist(msg.Vid, msg.Pid, msg.SoftwareVersion, msg.CertificationType)
	}

	if msg.CDVersionNumber != "" {
		cdVersionNumber, err := strconv.ParseUint(msg.CDVersionNumber, 10, 32)

		if err != nil {
			return nil, types.NewErrInvalidUint32ForCdVersionNumber(msg.Vid, msg.Pid, msg.SoftwareVersion, msg.CertificationType, msg.CDVersionNumber)
		}

		modelVersion, isFound := k.modelKeeper.GetModelVersion(ctx, msg.Vid, msg.Pid, msg.SoftwareVersion)

		if !isFound {
			return nil, modeltypes.NewErrModelVersionDoesNotExist(msg.Vid, msg.Pid, msg.SoftwareVersion)
		}

		if modelVersion.CdVersionNumber != int32(cdVersionNumber) {
			return nil, types.NewErrModelVersionCDVersionNumberDoesNotMatch(msg.Vid, msg.Pid, msg.SoftwareVersion, msg.CDVersionNumber)
		}

		complianceInfo.CDVersionNumber = uint32(cdVersionNumber)
	}

	if msg.CertificationIdOfSoftwareComponent != "" {
		complianceInfo.CertificationIdOfSoftwareComponent = msg.CertificationIdOfSoftwareComponent
	}

	if msg.CertificationRoute != "" {
		complianceInfo.CertificationRoute = msg.CertificationRoute
	}

	if msg.CompliantPlatformUsed != "" {
		complianceInfo.CompliantPlatformUsed = msg.CompliantPlatformUsed
	}

	if msg.CompliantPlatformVersion != "" {
		complianceInfo.CompliantPlatformVersion = msg.CompliantPlatformVersion
	}

	if msg.Date != "" {
		complianceInfo.Date = msg.Date
	}

	if msg.FamilyId != "" {
		complianceInfo.FamilyId = msg.FamilyId
	}

	if msg.OSVersion != "" {
		complianceInfo.OSVersion = msg.OSVersion
	}

	if msg.ParentChild != "" {
		complianceInfo.ParentChild = msg.ParentChild
	}

	if msg.ProgramType != "" {
		complianceInfo.ProgramType = msg.ProgramType
	}

	if msg.ProgramTypeVersion != "" {
		complianceInfo.ProgramTypeVersion = msg.ProgramTypeVersion
	}

	if msg.Reason != "" {
		complianceInfo.Reason = msg.Reason
	}

	if msg.SupportedClusters != "" {
		complianceInfo.SupportedClusters = msg.SupportedClusters
	}

	if msg.Transport != "" {
		complianceInfo.Transport = msg.Transport
	}

	//nolint:nestif
	if msg.CDCertificateId != "" && msg.CDCertificateId != complianceInfo.CDCertificateId {
		// remove the compliance info from a device software compliance with its old cd certificate id
		currentDeviceSoftwareCompliance, _ := k.GetDeviceSoftwareCompliance(ctx, complianceInfo.CDCertificateId)

		index, found := currentDeviceSoftwareCompliance.IsComplianceInfoExist(msg.Vid, msg.Pid, msg.SoftwareVersion)
		if found {
			currentDeviceSoftwareCompliance.RemoveComplianceInfo(index)
		}

		if len(currentDeviceSoftwareCompliance.ComplianceInfo) == 0 {
			k.RemoveDeviceSoftwareCompliance(ctx, currentDeviceSoftwareCompliance.CDCertificateId)
		} else {
			k.SetDeviceSoftwareCompliance(ctx, currentDeviceSoftwareCompliance)
		}

		// update the compliance info cd certificate id field
		complianceInfo.CDCertificateId = msg.CDCertificateId

		// insert the compliance info to a device software compliance with its new cd certificate id if exists, else create one.
		targetDeviceSoftwareCompliance, isFound := k.GetDeviceSoftwareCompliance(ctx, complianceInfo.CDCertificateId)

		if !isFound {
			targetDeviceSoftwareCompliance.CDCertificateId = msg.CDCertificateId
		}
		targetDeviceSoftwareCompliance.ComplianceInfo = append(targetDeviceSoftwareCompliance.ComplianceInfo, &complianceInfo)
		k.SetDeviceSoftwareCompliance(ctx, targetDeviceSoftwareCompliance)
	} else { // update the corresponding device software compliance to sync with compliance info.
		deviceSoftwareCompliance, _ := k.GetDeviceSoftwareCompliance(ctx, complianceInfo.CDCertificateId)

		index, found := deviceSoftwareCompliance.IsComplianceInfoExist(msg.Vid, msg.Pid, msg.SoftwareVersion)
		if found {
			deviceSoftwareCompliance.ComplianceInfo[index] = &complianceInfo
		}

		k.SetDeviceSoftwareCompliance(ctx, deviceSoftwareCompliance)
	}

	k.SetComplianceInfo(ctx, complianceInfo)

	return &types.MsgUpdateComplianceInfoResponse{}, nil
}
