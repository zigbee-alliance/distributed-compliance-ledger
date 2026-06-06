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

package keeper

import (
	"context"

	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/dclupgrade/types"
)

func (k msgServer) ApproveUpgrade(goCtx context.Context, msg *types.MsgApproveUpgrade) (*types.MsgApproveUpgradeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	creatorAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid Address: (%s)", err)
	}

	// check if message creator has enough rights to approve upgrade
	if !k.dclauthKeeper.HasRole(ctx, creatorAddr, types.UpgradeApprovalRole) {
		return nil, errors.Wrapf(sdkerrors.ErrUnauthorized,
			"MsgApproveUpgrade transaction should be signed by an account with the %s role",
			types.UpgradeApprovalRole,
		)
	}

	// check if proposed upgrade exists
	proposedUpgrade, isFound := k.GetProposedUpgrade(ctx, msg.Name)
	if !isFound {
		return nil, types.NewErrProposedUpgradeDoesNotExist(msg.Name)
	}

	// check if proposed upgrade already has approval form message creator
	if proposedUpgrade.HasApprovalFrom(creatorAddr) {
		return nil, errors.Wrapf(sdkerrors.ErrUnauthorized,
			"Proposed upgrade with name=%v already has approval from=%v",
			msg.Name, msg.Creator,
		)
	}

	// append approval
	grant := types.Grant{
		Address: creatorAddr.String(),
		Time:    msg.Time,
		Info:    msg.Info,
	}

	// check if proposed upgrade already has reject from message creator
	if proposedUpgrade.HasRejectFrom(creatorAddr) {
		for i, other := range proposedUpgrade.Rejects {
			if other.Address == grant.Address {
				proposedUpgrade.Rejects = append(proposedUpgrade.Rejects[:i], proposedUpgrade.Rejects[i+1:]...)

				break
			}
		}
	}
	proposedUpgrade.Approvals = append(proposedUpgrade.Approvals, &grant)

	// check if proposed upgrade has enough approvals
	if len(proposedUpgrade.Approvals) >= k.UpgradeApprovalsCount(ctx) {
		// schedule upgrade
		err = k.upgradeKeeper.ScheduleUpgrade(ctx, proposedUpgrade.Plan)
		if err != nil {
			return nil, err
		}

		// remove proposed upgrade
		k.RemoveProposedUpgrade(ctx, proposedUpgrade.Plan.Name)

		approvedUpgrage := types.ApprovedUpgrade(proposedUpgrade)
		k.SetApprovedUpgrade(ctx, approvedUpgrage)
	} else {
		// Execute scheduling upgrade in a new context branch (with branched store)
		// to validate msg.Plan before the proposal proceeds through the approval process.
		// State is not persisted.
		cacheCtx, _ := ctx.CacheContext()
		err = k.upgradeKeeper.ScheduleUpgrade(cacheCtx, proposedUpgrade.Plan)
		if err != nil {
			return nil, err
		}
		// update proposed upgrade
		k.SetProposedUpgrade(ctx, proposedUpgrade)
	}

	return &types.MsgApproveUpgradeResponse{}, nil
}
