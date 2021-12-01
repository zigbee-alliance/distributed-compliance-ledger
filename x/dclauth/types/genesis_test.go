package types_test

/*

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
)

func TestGenesisState_Validate(t *testing.T) {
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
					{
						Address: "0",
					},
					{
						Address: "1",
					},
				},
				PendingAccountList: []types.PendingAccount{
					{
						Address: "0",
					},
					{
						Address: "1",
					},
				},
				PendingAccountRevocationList: []types.PendingAccountRevocation{
					{
						Address: "0",
					},
					{
						Address: "1",
					},
				},
				AccountStat: &types.AccountStat{
					Number: 94,
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated account",
			genState: &types.GenesisState{
				AccountList: []types.Account{
					{
						Address: "0",
					},
					{
						Address: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated pendingAccount",
			genState: &types.GenesisState{
				PendingAccountList: []types.PendingAccount{
					{
						Address: "0",
					},
					{
						Address: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated pendingAccountRevocation",
			genState: &types.GenesisState{
				PendingAccountRevocationList: []types.PendingAccountRevocation{
					{
						Address: "0",
					},
					{
						Address: "0",
					},
				},
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
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
*/
