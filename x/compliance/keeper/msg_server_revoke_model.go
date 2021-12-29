package keeper

import (
	"context"
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliance/types"
	dclauthtypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
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

	complianceInfo, found := k.GetComplianceInfo(ctx, msg.Vid, msg.Pid, msg.SoftwareVersion, msg.CertificationType)
	// nolint:nestif
	if found {
		// Compliance record already exist. Cases:
		// 1) Only revocation is tracked on the ledger. We want to re-revoke.
		// The corresponding Model Info and test results are not required to be on the ledger.
		// 2) Compliance is tracked on ledger. We want to revoke certified compliance.
		// `Else` branch in CertifyModel was passed on first certification. So Model Info and test results exist on the ledger.

		// check if certification is already revoked
		if complianceInfo.SoftwareVersionCertificationStatus == types.CodeRevoked {
			// TODO: do we allow re-revocation (date update) by the same signer?
			// if complianceInfo.Owner != msg.Signer {
			// 	return nil, types.NewErrAlreadyRevoked(msg.Vid, msg.Pid)
			// }
			return nil, types.NewErrAlreadyRevoked(msg.Vid, msg.Pid, msg.SoftwareVersion, msg.CertificationType)
		} else {
			// if state changes on `revoked` check that revocation_date is after certification_date
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
		}

		// update certified/revoked index
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
		// Only revocation is tracked on the ledger. There is no compliance record yet.
		// The corresponding Model Info and test results are not required to be on the ledger.

		complianceInfo = types.ComplianceInfo{
			Vid:                                msg.Vid,
			Pid:                                msg.Pid,
			SoftwareVersion:                    msg.SoftwareVersion,
			SoftwareVersionString:              msg.SoftwareVersionString,
			CertificationType:                  msg.CertificationType,
			Date:                               msg.RevocationDate,
			Reason:                             msg.Reason,
			Owner:                              msg.Signer,
			SoftwareVersionCertificationStatus: types.CodeRevoked,
			History:                            []*types.ComplianceHistoryItem{},
			CDVersionNumber:                    msg.CDVersionNumber,
		}
	}

	// store compliance info
	k.SetComplianceInfo(ctx, complianceInfo)

	// update certified/revoked index
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
