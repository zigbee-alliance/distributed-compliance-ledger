// Copyright 2020 DSR Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package compliance

import (
	"bytes"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/auth"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliance/internal/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliance/internal/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliancetest"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/modelversion"
)

func NewHandler(keeper keeper.Keeper, modelversionKeeper modelversion.Keeper,
	compliancetestKeeper compliancetest.Keeper, authKeeper auth.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case types.MsgCertifyModel:
			return handleMsgCertifyModel(ctx, keeper, modelversionKeeper, compliancetestKeeper, authKeeper, msg)
		case types.MsgRevokeModel:
			return handleMsgRevokeModel(ctx, keeper, modelversionKeeper, authKeeper, msg)
		default:
			errMsg := fmt.Sprintf("unrecognized nameservice Msg type: %v", msg.Type())

			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

func handleMsgCertifyModel(ctx sdk.Context, keeper keeper.Keeper, modelversionKeeper modelversion.Keeper,
	compliancetestKeeper compliancetest.Keeper, authKeeper auth.Keeper,
	msg types.MsgCertifyModel) sdk.Result {
	// check if sender has enough rights to certify model
	if err := checkZbCertificationRights(ctx, authKeeper, msg.Signer, msg.CertificationType); err != nil {
		return err.Result()
	}

	if err := checkZbCertificationDone(ctx, keeper, authKeeper, msg.Signer, msg); err != nil {
		return err.Result()
	}

	var complianceInfo types.ComplianceInfo

	// nolint:nestif
	if keeper.IsComplianceInfoPresent(ctx, msg.CertificationType, msg.VID, msg.PID, msg.SoftwareVersion) {
		// Compliance record already exist. Cases:
		// 1) Only revocation is tracked on the ledger. We want to certify revoked compliance.
		//The corresponding Model Info and test results are not required to be on the ledger.
		// 2) Compliance is tracked on ledger. We want to certify revoked compliance.
		//`Else` branch was passed on first certification. So Model Info and test results are exists on the ledger.
		complianceInfo = keeper.GetComplianceInfo(ctx, msg.CertificationType, msg.VID, msg.PID, msg.SoftwareVersion)

		// if state changes on `certified` check that certification_date is after revocation_date
		if complianceInfo.SoftwareVersionCertificationStatus == types.CodeRevoked {
			if msg.CertificationDate.Before(complianceInfo.Date) {
				return types.ErrInconsistentDates(
					fmt.Sprintf("The `certification_date`:%v must be after the current `date`:%v to "+
						"certify model", msg.CertificationDate, complianceInfo.Date)).Result()
			}

			complianceInfo.UpdateComplianceInfo(msg.CertificationDate, msg.Reason)
		}
	} else {
		// Compliance is tracked on ledger. There is no compliance record yet.
		// The corresponding Model Info and test results must be present on ledger.
		if !modelversionKeeper.IsModelPresent(ctx, msg.VID, msg.PID) {
			return modelversion.ErrModelVersionDoesNotExist(msg.VID, msg.PID, msg.SoftwareVersion).Result()
		}

		if !compliancetestKeeper.IsTestingResultsPresents(ctx, msg.VID, msg.PID, msg.SoftwareVersion) {
			return compliancetest.ErrTestingResultDoesNotExist(msg.VID, msg.PID, msg.SoftwareVersion).Result()
		}

		complianceInfo = types.NewCertifiedComplianceInfo(
			msg.VID,
			msg.PID,
			msg.SoftwareVersion,
			msg.SoftwareVersionString,
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

func handleMsgRevokeModel(ctx sdk.Context, keeper keeper.Keeper, modelversionKeeper modelversion.Keeper,
	authKeeper auth.Keeper, msg types.MsgRevokeModel) sdk.Result {
	// check if sender has enough rights to revoke model
	if err := checkZbCertificationRights(ctx, authKeeper, msg.Signer, msg.CertificationType); err != nil {
		return err.Result()
	}

	var complianceInfo types.ComplianceInfo

	// nolint: gocritic, nestif
	if keeper.IsComplianceInfoPresent(ctx, msg.CertificationType, msg.VID, msg.PID, msg.SoftwareVersion) {
		// Compliance record already exist.
		complianceInfo = keeper.GetComplianceInfo(ctx, msg.CertificationType, msg.VID, msg.PID, msg.SoftwareVersion)

		// if state changes on `revoked` check that revocation_date is after certification_date
		if complianceInfo.SoftwareVersionCertificationStatus == types.CodeCertified {
			if msg.RevocationDate.Before(complianceInfo.Date) {
				return types.ErrInconsistentDates(
					fmt.Sprintf("The `revocation_date`:%v must be after the `certification_date`:%v to "+
						"revoke model", msg.RevocationDate, complianceInfo.Date)).Result()
			}

			complianceInfo.UpdateComplianceInfo(msg.RevocationDate, msg.Reason)
		}
	} else if modelversionKeeper.IsModelVersionPresent(ctx, msg.VID, msg.PID, msg.SoftwareVersion) {
		// Only revocation is tracked on the ledger. There is no compliance record yet.
		// The corresponding Model Info and test results are not required to be on the ledger.
		complianceInfo = types.NewRevokedComplianceInfo(
			msg.VID,
			msg.PID,
			msg.SoftwareVersion,
			msg.SoftwareVersionString,
			msg.CertificationType,
			msg.RevocationDate,
			msg.Reason,
			msg.Signer,
		)
	} else {
		return types.ErrModelDoesNotExist(msg.VID, msg.PID).Result()
	}

	// store compliance info
	keeper.SetComplianceInfo(ctx, complianceInfo)

	return sdk.Result{}
}

func checkZbCertificationRights(ctx sdk.Context, authKeeper auth.Keeper, signer sdk.AccAddress,
	certificationType types.CertificationType) sdk.Error {
	// rights are depend on certification type
	if certificationType == types.ZbCertificationType {
		// sender must have ZBCertificationCenter role to certify/revoke model
		if !authKeeper.HasRole(ctx, signer, auth.ZBCertificationCenter) {
			return sdk.ErrUnauthorized(fmt.Sprintf("MsgCertifyModel/MsgRevokeMode transaction should be "+
				"signed by an account with the %s role", auth.ZBCertificationCenter))
		}
	} else {
		return sdk.ErrUnknownRequest(fmt.Sprintf("Unexpected CertificationType: \"%s\". Supported types: [%s]",
			certificationType, ZbCertificationType))
	}

	return nil
}

func checkZbCertificationDone(
	ctx sdk.Context,
	keeper keeper.Keeper,
	authKeeper auth.Keeper,
	signer sdk.AccAddress,
	msg types.MsgCertifyModel) sdk.Error {
	if !keeper.IsComplianceInfoPresent(ctx, msg.CertificationType, msg.VID, msg.PID, msg.SoftwareVersion) {
		return nil
	}

	complianceInfo := keeper.GetComplianceInfo(ctx, msg.CertificationType, msg.VID, msg.PID, msg.SoftwareVersion)

	if complianceInfo.SoftwareVersionCertificationStatus != types.CodeCertified {
		return nil
	}

	if bytes.Equal(complianceInfo.Owner, signer) {
		return nil
	}

	return types.ErrAlreadyCertifyed(msg.VID, msg.PID)
}
