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
	"reflect"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

func TestAccountRole_Validate(t *testing.T) {
	tests := []struct {
		name    string
		role    AccountRole
		wantErr bool
	}{
		{
			name:    "invalid role",
			wantErr: true,
		},
		{
			name:    "valid  vendor role",
			role:    Vendor,
			wantErr: false,
		},
		{
			name:    "valid  certification center role",
			role:    CertificationCenter,
			wantErr: false,
		},
		{
			name:    "valid  trustee role",
			role:    Trustee,
			wantErr: false,
		},
		{
			name:    "valid  node admin role",
			role:    NodeAdmin,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.role.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("AccountRole.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewAccount(t *testing.T) {
	type args struct {
		ba        *authtypes.BaseAccount
		roles     AccountRoles
		approvals []*Grant
		vendorID  int32
	}
	tests := []struct {
		name string
		args args
		want *Account
	}{
		{
			name: "valid account all roles",
			args: args{
				ba:        &authtypes.BaseAccount{},
				roles:     []AccountRole{Vendor, CertificationCenter, Trustee, NodeAdmin},
				approvals: []*Grant{},
				vendorID:  1,
			},
			want: &Account{
				BaseAccount: &authtypes.BaseAccount{},
				Roles:       []AccountRole{Vendor, CertificationCenter, Trustee, NodeAdmin},
				Approvals:   []*Grant{},
				VendorID:    1,
			},
		},
		{
			name: "invalid account vendor role",
			args: args{
				ba:        &authtypes.BaseAccount{},
				roles:     []AccountRole{Vendor},
				approvals: []*Grant{},
				vendorID:  2,
			},
			want: &Account{
				BaseAccount: &authtypes.BaseAccount{},
				Roles:       []AccountRole{Vendor},
				Approvals:   []*Grant{},
				VendorID:    2,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAccount(tt.args.ba, tt.args.roles, tt.args.approvals, tt.args.vendorID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAccount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccount_Validate(t *testing.T) {
	type fields struct {
		BaseAccount *types.BaseAccount
		Roles       []AccountRole
		Approvals   []*Grant
		VendorID    int32
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "valid account with vendor ID",
			fields: fields{
				BaseAccount: &types.BaseAccount{},
				Roles:       []AccountRole{Vendor},
				Approvals:   []*Grant{},
				VendorID:    1,
			},
			wantErr: false,
		},
		{
			name: "valid account with certification center role",
			fields: fields{
				BaseAccount: &types.BaseAccount{},
				Roles:       []AccountRole{CertificationCenter},
				Approvals:   []*Grant{},
				VendorID:    0,
			},
			wantErr: false,
		},
		{
			name: "invalid vendor account with missing vendor ID",
			fields: fields{
				BaseAccount: &types.BaseAccount{},
				Roles:       []AccountRole{Vendor},
				Approvals:   []*Grant{},
				VendorID:    0,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			acc := Account{
				BaseAccount: tt.fields.BaseAccount,
				Roles:       tt.fields.Roles,
				Approvals:   tt.fields.Approvals,
				VendorID:    tt.fields.VendorID,
			}
			if err := acc.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Account.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAccount_GetRoles(t *testing.T) {
	type fields struct {
		BaseAccount *types.BaseAccount
		Roles       []AccountRole
		Approvals   []*Grant
		VendorID    int32
	}
	tests := []struct {
		name   string
		fields fields
		want   []AccountRole
	}{
		{
			name: "account with Vendor and Trustee roles",
			fields: fields{
				BaseAccount: &types.BaseAccount{},
				Roles:       []AccountRole{Vendor, Trustee},
				Approvals:   []*Grant{},
				VendorID:    1,
			},
			want: []AccountRole{Vendor, Trustee},
		},
		{
			name: "account with Vendor, Trustee and NodeAdmin roles",
			fields: fields{
				BaseAccount: &types.BaseAccount{},
				Roles:       []AccountRole{Vendor, Trustee, NodeAdmin},
				Approvals:   []*Grant{},
				VendorID:    1,
			},
			want: []AccountRole{Vendor, Trustee, NodeAdmin},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			acc := Account{
				BaseAccount: tt.fields.BaseAccount,
				Roles:       tt.fields.Roles,
				Approvals:   tt.fields.Approvals,
				VendorID:    tt.fields.VendorID,
			}
			if got := acc.GetRoles(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Account.GetRoles() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccount_GetVendorID(t *testing.T) {
	type fields struct {
		BaseAccount *types.BaseAccount
		Roles       []AccountRole
		Approvals   []*Grant
		VendorID    int32
	}
	tests := []struct {
		name   string
		fields fields
		want   int32
	}{
		{
			name: "account with vendor ID 45",
			fields: fields{
				BaseAccount: &types.BaseAccount{},
				Roles:       []AccountRole{Vendor},
				Approvals:   []*Grant{},
				VendorID:    45,
			},
			want: 45,
		},
		{
			name: "account with vendor ID 0",
			fields: fields{
				BaseAccount: &types.BaseAccount{},
				Roles:       []AccountRole{Trustee},
				Approvals:   []*Grant{},
				VendorID:    0,
			},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			acc := Account{
				BaseAccount: tt.fields.BaseAccount,
				Roles:       tt.fields.Roles,
				Approvals:   tt.fields.Approvals,
				VendorID:    tt.fields.VendorID,
			}
			if got := acc.GetVendorID(); got != tt.want {
				t.Errorf("Account.GetVendorID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccount_HasRole(t *testing.T) {
	type fields struct {
		BaseAccount *types.BaseAccount
		Roles       []AccountRole
		Approvals   []*Grant
		VendorID    int32
	}
	type args struct {
		targetRole AccountRole
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "account with Vendor and Trustee roles",
			fields: fields{
				BaseAccount: &types.BaseAccount{},
				Roles:       []AccountRole{Vendor, Trustee},
				Approvals:   []*Grant{},
				VendorID:    1,
			},
			args: args{
				targetRole: Vendor,
			},
			want: true,
		},
		{
			name: "account with Vendor, Trustee and NodeAdmin roles",
			fields: fields{
				BaseAccount: &types.BaseAccount{},
				Roles:       []AccountRole{Vendor, Trustee, NodeAdmin},
				Approvals:   []*Grant{},
				VendorID:    1,
			},
			args: args{
				targetRole: NodeAdmin,
			},
			want: true,
		},
		{
			name: "account with Vendor, Trustee and NodeAdmin roles",
			fields: fields{
				BaseAccount: &types.BaseAccount{},
				Roles:       []AccountRole{Vendor, Trustee, NodeAdmin},
				Approvals:   []*Grant{},
				VendorID:    1,
			},
			args: args{
				targetRole: CertificationCenter,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			acc := Account{
				BaseAccount: tt.fields.BaseAccount,
				Roles:       tt.fields.Roles,
				Approvals:   tt.fields.Approvals,
				VendorID:    tt.fields.VendorID,
			}
			if got := acc.HasRole(tt.args.targetRole); got != tt.want {
				t.Errorf("Account.HasRole() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPendingAccount_HasApprovalFrom(t *testing.T) {
	type fields struct {
		Account *Account
	}
	type args struct {
		address sdk.AccAddress
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			acc := PendingAccount{
				Account: tt.fields.Account,
			}
			if got := acc.HasApprovalFrom(tt.args.address); got != tt.want {
				t.Errorf("PendingAccount.HasApprovalFrom() = %v, want %v", got, tt.want)
			}
		})
	}
}
