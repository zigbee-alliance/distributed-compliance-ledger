package types_test

// import (
// 	"testing"

// 	"github.com/stretchr/testify/require"
// 	"github.com/zigbee-alliance/distributed-compliance-ledger/x/vendorinfo/types"
// )

// func TestGenesisState_Validate(t *testing.T) {
// 	for _, tc := range []struct {
// 		desc     string
// 		genState *types.GenesisState
// 		valid    bool
// 	}{
// 		{
// 			desc:     "default is valid",
// 			genState: types.DefaultGenesis(),
// 			valid:    true,
// 		},
// 		{
// 			desc: "valid genesis state",
// 			genState: &types.GenesisState{
// 				VendorInfoList: []types.VendorInfo{
// 					{
// 						VendorID: 0,
// 					},
// 					{
// 						VendorID: 1,
// 					},
// 				},
// 				// this line is used by starport scaffolding # types/genesis/validField
// 			},
// 			valid: true,
// 		},
// 		{
// 			desc: "duplicated vendorInfo",
// 			genState: &types.GenesisState{
// 				VendorInfoList: []types.VendorInfo{
// 					{
// 						VendorID: 0,
// 					},
// 					{
// 						VendorID: 0,
// 					},
// 				},
// 			},
// 			valid: false,
// 		},
// 		// this line is used by starport scaffolding # types/genesis/testcase
// 	} {
// 		t.Run(tc.desc, func(t *testing.T) {
// 			err := tc.genState.Validate()
// 			if tc.valid {
// 				require.NoError(t, err)
// 			} else {
// 				require.Error(t, err)
// 			}
// 		})
// 	}
// }
