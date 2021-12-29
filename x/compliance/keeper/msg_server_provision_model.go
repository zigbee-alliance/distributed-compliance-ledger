package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliance/types"
	dclauthtypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
)

func (k msgServer) ProvisionModel(goCtx context.Context, msg *types.MsgProvisionModel) (*types.MsgProvisionModelResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	signerAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid Address: (%s)", err)
	}

	// check if sender has enough rights to provision model
	// sender must have CertificationCenter role to certify/revoke model
	if !k.dclauthKeeper.HasRole(ctx, signerAddr, dclauthtypes.CertificationCenter) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized,
			"MsgAddTestingResult transaction should be signed by an account with the %s role",
			dclauthtypes.CertificationCenter,
		)
	}

	// can set provisional state only if there is no compliance info on the ledger
	complianceInfo, found := k.GetComplianceInfo(ctx, msg.Vid, msg.Pid, msg.SoftwareVersion, msg.CertificationType)
	if found {
		switch status := complianceInfo.SoftwareVersionCertificationStatus; status {
		case types.CodeProvisional:
			return nil, types.NewErrAlreadyProvisional(msg.Vid, msg.Pid, msg.SoftwareVersion, msg.CertificationType)
		case types.CodeCertified:
			return nil, types.NewErrAlreadyCertified(msg.Vid, msg.Pid, msg.SoftwareVersion, msg.CertificationType)
		case types.CodeRevoked:
			return nil, types.NewErrAlreadyRevoked(msg.Vid, msg.Pid, msg.SoftwareVersion, msg.CertificationType)
		default:
			return nil, types.NewErrComplianceInfoAlreadyExist(msg.Vid, msg.Pid, msg.SoftwareVersion, msg.CertificationType)
		}
	}

	complianceInfo = types.ComplianceInfo{
		Vid:                                msg.Vid,
		Pid:                                msg.Pid,
		SoftwareVersion:                    msg.SoftwareVersion,
		SoftwareVersionString:              msg.SoftwareVersionString,
		CertificationType:                  msg.CertificationType,
		Date:                               msg.ProvisionalDate,
		Reason:                             msg.Reason,
		Owner:                              msg.Signer,
		SoftwareVersionCertificationStatus: types.CodeProvisional,
		History:                            []*types.ComplianceHistoryItem{},
		CDVersionNumber:                    msg.CDVersionNumber,
	}

	// store compliance info
	k.SetComplianceInfo(ctx, complianceInfo)

	// update provisional index
	provisionalModel := types.ProvisionalModel{
		Vid:               msg.Vid,
		Pid:               msg.Pid,
		SoftwareVersion:   msg.SoftwareVersion,
		CertificationType: msg.CertificationType,
		Value:             true,
	}
	k.SetProvisionalModel(ctx, provisionalModel)

	return &types.MsgProvisionModelResponse{}, nil
}
