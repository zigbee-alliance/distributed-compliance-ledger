package keeper

import (
	"context"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/validator"
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
		cdVersionNumber, err := ParseCDVersionNumber(msg.CDVersionNumber)

		if err != nil {
			return nil, err
		}

		modelVersion, isFound := k.modelKeeper.GetModelVersion(ctx, msg.Vid, msg.Pid, msg.SoftwareVersion)

		if !isFound {
			return nil, modeltypes.NewErrModelVersionDoesNotExist(msg.Vid, msg.Pid, msg.SoftwareVersion)
		}

		if modelVersion.CdVersionNumber != int32(cdVersionNumber) {
			return nil, types.NewErrModelVersionCDVersionNumberDoesNotMatch(msg.Vid, msg.Pid, msg.SoftwareVersion, msg.CDVersionNumber)
		}

		complianceInfo.CDVersionNumber = cdVersionNumber
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

	// if cdCertificateId is present, update all related indices as well.
	if msg.CDCertificateId != "" {
		deviceSoftwareCompliance, isFound := k.GetDeviceSoftwareCompliance(ctx, complianceInfo.CDCertificateId)

		if !isFound {
			deviceSoftwareCompliance.CDCertificateId = msg.CDCertificateId
			deviceSoftwareCompliance.ComplianceInfo = append(deviceSoftwareCompliance.ComplianceInfo, &complianceInfo)
		}

		for _, info := range deviceSoftwareCompliance.ComplianceInfo {
			info.CDCertificateId = msg.CDCertificateId
			k.SetComplianceInfo(ctx, *info)
		}

		complianceInfo.CDCertificateId = msg.CDCertificateId
		deviceSoftwareCompliance.CDCertificateId = msg.CDCertificateId

		k.SetDeviceSoftwareCompliance(ctx, deviceSoftwareCompliance)
	}

	k.SetComplianceInfo(ctx, complianceInfo)

	return &types.MsgUpdateComplianceInfoResponse{}, nil
}

func ParseCDVersionNumber(cdVersionNumberString string) (uint32, error) {
	cdVersionNumber, err := strconv.ParseUint(cdVersionNumberString, 10, 32)

	if err != nil {
		return 65536, err
	}

	if cdVersionNumber > 65535 {
		return 65536, sdkerrors.Wrap(validator.ErrFieldUpperBoundViolated, "CDVersionNumber must not be greater than 65535: field upper bound violatedError")
	}

	return uint32(cdVersionNumber), nil
}
