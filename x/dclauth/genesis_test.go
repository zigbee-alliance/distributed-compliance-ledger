package dclauth_test

/*

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/zigbee-alliance/distributed-compliance-ledger/testutil/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
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
			Number: 26,
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.DclauthKeeper(t)
	dclauth.InitGenesis(ctx, *k, genesisState)
	got := dclauth.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	require.ElementsMatch(t, genesisState.AccountList, got.AccountList)
	require.ElementsMatch(t, genesisState.PendingAccountList, got.PendingAccountList)
	require.ElementsMatch(t, genesisState.PendingAccountRevocationList, got.PendingAccountRevocationList)
	require.Equal(t, genesisState.AccountStat, got.AccountStat)
	// this line is used by starport scaffolding # genesis/test/assert
}
*/
