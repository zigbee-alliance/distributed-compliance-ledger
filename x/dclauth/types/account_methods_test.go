package types_test

import (
	"testing"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/stretchr/testify/require"
	"github.com/zigbee-alliance/distributed-compliance-ledger/testutil/sample"
	commontypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/common/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
)

func TestAccountRole_Validate(t *testing.T) {
	require.NoError(t, types.Vendor.Validate())
	require.NoError(t, types.Trustee.Validate())
	require.Error(t, types.AccountRole("NotARole").Validate())
}

func baseAccount(t *testing.T) *authtypes.BaseAccount {
	t.Helper()
	addr := sample.AccAddress()

	return &authtypes.BaseAccount{Address: addr}
}

func TestAccount_Validate(t *testing.T) {
	t.Run("valid non-vendor", func(t *testing.T) {
		acc := types.NewAccount(baseAccount(t), types.AccountRoles{types.Trustee}, nil, nil, 0, nil)
		require.NoError(t, acc.Validate())
	})
	t.Run("valid vendor with vendorID", func(t *testing.T) {
		acc := types.NewAccount(baseAccount(t), types.AccountRoles{types.Vendor}, nil, nil, 1, nil)
		require.NoError(t, acc.Validate())
	})
	t.Run("invalid role", func(t *testing.T) {
		acc := types.NewAccount(baseAccount(t), types.AccountRoles{types.AccountRole("Bogus")}, nil, nil, 0, nil)
		require.Error(t, acc.Validate())
	})
	t.Run("vendor without vendorID", func(t *testing.T) {
		acc := types.NewAccount(baseAccount(t), types.AccountRoles{types.Vendor}, nil, nil, 0, nil)
		require.Error(t, acc.Validate())
	})
}

func TestAccount_RoleHelpers(t *testing.T) {
	acc := types.NewAccount(baseAccount(t), types.AccountRoles{types.Vendor}, nil, nil, 1, nil)
	require.True(t, acc.HasRole(types.Vendor))
	require.False(t, acc.HasRole(types.Trustee))
	require.True(t, acc.HasOnlyVendorRole(types.Vendor))

	multi := types.NewAccount(baseAccount(t), types.AccountRoles{types.Vendor, types.Trustee}, nil, nil, 1, nil)
	require.False(t, multi.HasOnlyVendorRole(types.Vendor))
}

func TestAccount_HasRightsToChange(t *testing.T) {
	// No products -> rights to change anything.
	noProducts := types.NewAccount(baseAccount(t), types.AccountRoles{types.Vendor}, nil, nil, 1, nil)
	require.True(t, noProducts.HasRightsToChange(123))

	scoped := types.NewAccount(baseAccount(t), types.AccountRoles{types.Vendor}, nil, nil, 1,
		[]*commontypes.Uint16Range{{Min: 10, Max: 20}})
	require.True(t, scoped.HasRightsToChange(10))  // min boundary
	require.True(t, scoped.HasRightsToChange(20))  // max boundary
	require.True(t, scoped.HasRightsToChange(15))  // inside
	require.False(t, scoped.HasRightsToChange(9))  // below
	require.False(t, scoped.HasRightsToChange(21)) // above
}
