// Copyright 2022 DSR Corporation
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

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
)

func (k msgServer) ProposeRevokeAccount(goCtx context.Context, msg *types.MsgProposeRevokeAccount) (*types.MsgProposeRevokeAccountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	signerAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid Signer: (%s)", err)
	}

	// check that sender has enough rights to propose account revocation
	if !k.HasRole(ctx, signerAddr, types.Trustee) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized,
			"MsgProposeRevokeAccount transaction should be signed by an account with the %s role",
			types.Trustee,
		)
	}

	accAddr, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid Address: (%s)", err)
	}

	// check that account exists
	if !k.IsAccountPresent(ctx, accAddr) {
		return nil, types.ErrAccountDoesNotExist(msg.Address)
	}

	// check that pending account revocation does not exist yet
	if k.IsPendingAccountRevocationPresent(ctx, accAddr) {
		return nil, types.ErrPendingAccountRevocationAlreadyExists(msg.Address)
	}

	// if more than 1 trustee's approval is needed, create pending account revocation else delete the account.
	if AccountApprovalsCount(ctx, k.Keeper) > 1 {
		// create and store pending account revocation record
		revoc := types.NewPendingAccountRevocation(accAddr, signerAddr)
		k.SetPendingAccountRevocation(ctx, revoc)
	} else {
		// delete account record
		k.RemoveAccount(ctx, accAddr)
	}

	return &types.MsgProposeRevokeAccountResponse{}, nil
}
