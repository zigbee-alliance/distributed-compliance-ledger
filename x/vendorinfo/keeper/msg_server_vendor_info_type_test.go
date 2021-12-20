package keeper_test

import (
    "strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

    keepertest "github.com/zigbee-alliance/distributed-compliance-ledger/testutil/keeper"
    "github.com/zigbee-alliance/distributed-compliance-ledger/x/vendorinfo/keeper"
    "github.com/zigbee-alliance/distributed-compliance-ledger/x/vendorinfo/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestVendorInfoTypeMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.VendorinfoKeeper(t)
	srv := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateVendorInfoType{Creator: creator,
		    Index: strconv.Itoa(i),
            
		}
		_, err := srv.CreateVendorInfoType(wctx, expected)
		require.NoError(t, err)
		rst, found := k.GetVendorInfoType(ctx,
		    expected.Index,
            
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestVendorInfoTypeMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateVendorInfoType
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgUpdateVendorInfoType{Creator: creator,
			    Index: strconv.Itoa(0),
                
			},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateVendorInfoType{Creator: "B",
			    Index: strconv.Itoa(0),
                
			},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "KeyNotFound",
			request: &types.MsgUpdateVendorInfoType{Creator: creator,
			    Index: strconv.Itoa(100000),
                
			},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.VendorinfoKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)
			expected := &types.MsgCreateVendorInfoType{Creator: creator,
			    Index: strconv.Itoa(0),
                
			}
			_, err := srv.CreateVendorInfoType(wctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateVendorInfoType(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetVendorInfoType(ctx,
				    expected.Index,
                    
				)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestVendorInfoTypeMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteVendorInfoType
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgDeleteVendorInfoType{Creator: creator,
			    Index: strconv.Itoa(0),
                
			},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgDeleteVendorInfoType{Creator: "B",
			    Index: strconv.Itoa(0),
                
			},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "KeyNotFound",
			request: &types.MsgDeleteVendorInfoType{Creator: creator,
			    Index: strconv.Itoa(100000),
                
			},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.VendorinfoKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)

			_, err := srv.CreateVendorInfoType(wctx, &types.MsgCreateVendorInfoType{Creator: creator,
			    Index: strconv.Itoa(0),
                
			})
			require.NoError(t, err)
			_, err = srv.DeleteVendorInfoType(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetVendorInfoType(ctx,
				    tc.request.Index,
                    
				)
				require.False(t, found)
			}
		})
	}
}
