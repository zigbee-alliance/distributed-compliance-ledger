package types_test

import (
	"testing"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/stretchr/testify/require"
	"github.com/zigbee-alliance/distributed-compliance-ledger/testutil/sample"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
)

func testAccount(address string) *types.Account {
	return &types.Account{BaseAccount: &authtypes.BaseAccount{Address: address}}
}

func TestGenesisState_Validate(t *testing.T) {
	addr1 := sample.AccAddress()
	addr2 := sample.AccAddress()

	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{
				AccountList: []types.Account{
					*testAccount(addr1),
					*testAccount(addr2),
				},
				PendingAccountList: []types.PendingAccount{
					{Account: testAccount(addr1)},
					{Account: testAccount(addr2)},
				},
				PendingAccountRevocationList: []types.PendingAccountRevocation{
					{Address: addr1},
					{Address: addr2},
				},
				RevokedAccountList: []types.RevokedAccount{
					{Account: testAccount(addr1)},
					{Account: testAccount(addr2)},
				},
				RejectedAccountList: []types.RejectedAccount{
					{Account: testAccount(addr1)},
					{Account: testAccount(addr2)},
				},
			},
			valid: true,
		},
		{
			desc: "invalid account address",
			genState: &types.GenesisState{
				AccountList: []types.Account{*testAccount("invalid-address")},
			},
			valid: false,
		},
		{
			desc: "duplicated account",
			genState: &types.GenesisState{
				AccountList: []types.Account{
					*testAccount(addr1),
					*testAccount(addr1),
				},
			},
			valid: false,
		},
		{
			desc: "invalid pendingAccount address",
			genState: &types.GenesisState{
				PendingAccountList: []types.PendingAccount{
					{Account: testAccount("invalid-address")},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated pendingAccount",
			genState: &types.GenesisState{
				PendingAccountList: []types.PendingAccount{
					{Account: testAccount(addr1)},
					{Account: testAccount(addr1)},
				},
			},
			valid: false,
		},
		{
			desc: "invalid pendingAccountRevocation address",
			genState: &types.GenesisState{
				PendingAccountRevocationList: []types.PendingAccountRevocation{
					{Address: "invalid-address"},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated pendingAccountRevocation",
			genState: &types.GenesisState{
				PendingAccountRevocationList: []types.PendingAccountRevocation{
					{Address: addr1},
					{Address: addr1},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated revokedAccount",
			genState: &types.GenesisState{
				RevokedAccountList: []types.RevokedAccount{
					{Account: testAccount(addr1)},
					{Account: testAccount(addr1)},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated rejectedAccount",
			genState: &types.GenesisState{
				RejectedAccountList: []types.RejectedAccount{
					{Account: testAccount(addr1)},
					{Account: testAccount(addr1)},
				},
			},
			valid: false,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
