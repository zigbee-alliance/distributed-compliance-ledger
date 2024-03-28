package keeper

import (
	"context"

	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliance/types"
	dclauthtypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
)

func (k msgServer) DeleteComplianceInfo(goCtx context.Context, msg *types.MsgDeleteComplianceInfo) (*types.MsgDeleteComplianceInfoResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	signerAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid Address: (%s)", err)
	}

	// check if sender has enough rights to delete model
	// sender must have CertificationCenter role to certify/revoke model
	if !k.dclauthKeeper.HasRole(ctx, signerAddr, dclauthtypes.CertificationCenter) {
		return nil, errors.Wrapf(sdkerrors.ErrUnauthorized,
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
	if len(deviceSoftwareCompliance.ComplianceInfo) != 0 {
		k.SetDeviceSoftwareCompliance(ctx, deviceSoftwareCompliance)
	}
	// If we don't have compliance info in Device Software Compliance - we should delete this Device Software Compliance
	if len(deviceSoftwareCompliance.ComplianceInfo) == 0 {
		k.RemoveDeviceSoftwareCompliance(ctx, deviceSoftwareCompliance.CDCertificateId)
	}

	k.RemoveRevokedModel(ctx, msg.Vid, msg.Pid, msg.SoftwareVersion, msg.CertificationType)
	k.RemoveProvisionalModel(ctx, msg.Vid, msg.Pid, msg.SoftwareVersion, msg.CertificationType)
	k.RemoveCertifiedModel(ctx, msg.Vid, msg.Pid, msg.SoftwareVersion, msg.CertificationType)

	// store compliance info
	k.RemoveComplianceInfo(ctx, msg.Vid, msg.Pid, msg.SoftwareVersion, msg.CertificationType)

	return &types.MsgDeleteComplianceInfoResponse{}, nil
}
