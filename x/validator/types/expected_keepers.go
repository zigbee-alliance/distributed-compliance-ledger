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

package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	dclauthtypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
)

type DclauthKeeper interface {
	// Methods imported from dclauth should be defined here
	HasRole(ctx sdk.Context, addr sdk.AccAddress, roleToCheck dclauthtypes.AccountRole) bool
	CountAccountsWithRole(ctx sdk.Context, roleToCount dclauthtypes.AccountRole) int
	GetAccountO(ctx sdk.Context, address sdk.AccAddress) (val dclauthtypes.Account, found bool)
	SetRevokedAccount(ctx sdk.Context, revokedAccount dclauthtypes.RevokedAccount)
	RemoveAccount(ctx sdk.Context, address sdk.AccAddress)
	AddAccountToRevokedAccount(
		ctx sdk.Context, accAddr sdk.AccAddress, approvals []*dclauthtypes.Grant, reason dclauthtypes.RevokedAccount_Reason, //nolint:nosnakecase
	) (*dclauthtypes.RevokedAccount, error)
}
