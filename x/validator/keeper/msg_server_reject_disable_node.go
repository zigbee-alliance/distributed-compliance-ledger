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
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/validator/types"
)

func (k msgServer) RejectDisableValidator(goCtx context.Context, msg *types.MsgRejectDisableValidator) (*types.MsgRejectDisableValidatorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	creatorAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid creator address: (%s)", err)
	}

	validatorAddr, err := sdk.ValAddressFromBech32(msg.Address)
	if err != nil {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid validator address: (%s)", err)
	}

	// check if message creator has enough rights to reject disable validator
	if !k.dclauthKeeper.HasRole(ctx, creatorAddr, types.VoteForDisableValidatorRole) {
		return nil, errors.Wrapf(sdkerrors.ErrUnauthorized,
			"MsgRejectDisableValidator transaction should be signed by an account with the %s role",
			types.VoteForDisableValidatorRole,
		)
	}

	// check if proposed disable validator exists
	proposedDisableValidator, isFound := k.GetProposedDisableValidator(ctx, validatorAddr.String())
	if !isFound {
		return nil, types.NewErrProposedDisableValidatorDoesNotExist(msg.Address)
	}

	// check if disable validator already has reject from message creator
	if proposedDisableValidator.HasRejectDisableFrom(creatorAddr) {
		return nil, errors.Wrapf(sdkerrors.ErrUnauthorized,
			"Disabled validator with address=%v already has reject from=%v",
			msg.Address,
			msg.Creator,
		)
	}

	// append approval
	grant := types.Grant{
		Address: creatorAddr.String(),
		Time:    msg.Time,
		Info:    msg.Info,
	}

	// check if disable validator has approval from message creator
	if proposedDisableValidator.HasApprovalFrom(creatorAddr) {
		// Remove propose disable validator if there are no rejects and other approvals
		if len(proposedDisableValidator.Approvals) == 1 && len(proposedDisableValidator.Rejects) == 0 {
			k.RemoveProposedDisableValidator(ctx, proposedDisableValidator.Address)

			return &types.MsgRejectDisableValidatorResponse{}, nil
		}

		// Remove approval from the list of approvals
		for i, other := range proposedDisableValidator.Approvals {
			if other.Address == grant.Address {
				proposedDisableValidator.Approvals = append(proposedDisableValidator.Approvals[:i], proposedDisableValidator.Approvals[i+1:]...)

				break
			}
		}
	}
	proposedDisableValidator.Rejects = append(proposedDisableValidator.Rejects, &grant)

	// check if proposed disable validator has enough reject approvals
	if len(proposedDisableValidator.Rejects) >= k.DisableValidatorRejectApprovalsCount(ctx) {
		k.RemoveProposedDisableValidator(ctx, proposedDisableValidator.Address)
		rejectedDisableValidator := types.RejectedDisableValidator(proposedDisableValidator)
		k.SetRejectedNode(ctx, rejectedDisableValidator)
	} else {
		// update proposed disable validator
		k.SetProposedDisableValidator(ctx, proposedDisableValidator)
	}

	return &types.MsgRejectDisableValidatorResponse{}, nil
}
