package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	dclcompltypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/compliance"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliance/types"
	dclauthtypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
)

func (k msgServer) DeleteComplianceInfo(goCtx context.Context, msg *types.MsgDeleteComplianceInfo) (*types.MsgDeleteComplianceInfoResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	signerAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid Address: (%s)", err)
	}

	// check if sender has enough rights to delete model
	// sender must have CertificationCenter role to certify/revoke model
	if !k.dclauthKeeper.HasRole(ctx, signerAddr, dclauthtypes.CertificationCenter) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized,
			"MsgDeleteComplianceInfo transaction should be signed by an account with the %s role",
			dclauthtypes.CertificationCenter,
		)
	}

	complianceInfo, found := k.GetComplianceInfo(ctx, msg.Vid, msg.Pid, msg.SoftwareVersion, msg.CertificationType)
	if !found {
		return nil, types.NewErrComplianceInfoDoesNotExist(msg.Vid, msg.Pid, msg.SoftwareVersion, msg.CertificationType)
	}

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

	switch complianceInfo.SoftwareVersionCertificationStatus {
	case dclcompltypes.CodeRevoked:
		k.RemoveRevokedModel(ctx, msg.Vid, msg.Pid, msg.SoftwareVersion, msg.CertificationType)
	case dclcompltypes.CodeProvisional:
		k.RemoveProvisionalModel(ctx, msg.Vid, msg.Pid, msg.SoftwareVersion, msg.CertificationType)
	case dclcompltypes.CodeCertified:
		k.RemoveCertifiedModel(ctx, msg.Vid, msg.Pid, msg.SoftwareVersion, msg.CertificationType)
	default:
		break
	}

	// store compliance info
	k.RemoveComplianceInfo(ctx, msg.Vid, msg.Pid, msg.SoftwareVersion, msg.CertificationType)

	return &types.MsgDeleteComplianceInfoResponse{}, nil
}
