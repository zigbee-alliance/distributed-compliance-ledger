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

package auth

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type GenesisState struct {
	Accounts                  []Account                  `json:"accounts"`
	PendingAccounts           []PendingAccount           `json:"pending_accounts"`
	PendingAccountRevocations []PendingAccountRevocation `json:"pending_account_revocations"`
}

func NewGenesisState() GenesisState {
	return GenesisState{
		Accounts:                  []Account{},
		PendingAccounts:           []PendingAccount{},
		PendingAccountRevocations: []PendingAccountRevocation{},
	}
}

func ValidateGenesis(data GenesisState) error {
	for _, record := range data.Accounts {
		if err := record.Validate(); err != nil {
			return err
		}
	}

	for _, record := range data.PendingAccounts {
		if err := record.Validate(); err != nil {
			return err
		}
	}

	for _, record := range data.PendingAccountRevocations {
		if err := record.Validate(); err != nil {
			return err
		}
	}

	return nil
}

func DefaultGenesisState() GenesisState {
	return NewGenesisState()
}

func InitGenesis(ctx sdk.Context, keeper Keeper, data GenesisState) {
	for _, record := range data.Accounts {
		keeper.SetAccount(ctx, record)
	}

	for _, record := range data.PendingAccounts {
		keeper.SetPendingAccount(ctx, record)
	}

	for _, record := range data.PendingAccountRevocations {
		keeper.SetPendingAccountRevocation(ctx, record)
	}
}

func ExportGenesis(ctx sdk.Context, k Keeper) GenesisState {
	var (
		accounts                  []Account
		pendingAccounts           []PendingAccount
		pendingAccountRevocations []PendingAccountRevocation
	)

	k.IterateAccounts(ctx, func(account Account) (stop bool) {
		accounts = append(accounts, account)

		return false
	})

	k.IteratePendingAccounts(ctx, func(pendingAccount PendingAccount) (stop bool) {
		pendingAccounts = append(pendingAccounts, pendingAccount)

		return false
	})

	k.IteratePendingAccountRevocations(ctx, func(pendingAccountRevocation PendingAccountRevocation) (stop bool) {
		pendingAccountRevocations = append(pendingAccountRevocations, pendingAccountRevocation)

		return false
	})

	return GenesisState{
		Accounts:                  accounts,
		PendingAccounts:           pendingAccounts,
		PendingAccountRevocations: pendingAccountRevocations,
	}
}
