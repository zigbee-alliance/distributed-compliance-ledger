package types_test

import (
	"strings"
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/zigbee-alliance/distributed-compliance-ledger/testutil/sample"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/validator/types"
)

func newTestValidator(t *testing.T) (types.Validator, sdk.ValAddress) {
	t.Helper()
	pk := ed25519.GenPrivKey().PubKey()
	owner := sdk.ValAddress(pk.Address())
	v, err := types.NewValidator(owner, pk, types.NewDescription("moniker", "identity", "website", "details"))
	require.NoError(t, err)

	return v, owner
}

func TestValidator_Accessors(t *testing.T) {
	v, owner := newTestValidator(t)

	require.Equal(t, types.Power, v.GetPower())
	require.Equal(t, "moniker", v.GetMoniker())
	require.False(t, v.IsJailed())
	require.Equal(t, owner, v.GetOwner())

	consAddr, err := v.GetConsAddress()
	require.NoError(t, err)
	require.NotEmpty(t, consAddr)

	consPubKey, err := v.GetConsPubKey()
	require.NoError(t, err)
	require.NotNil(t, consPubKey)

	consAddr2, err := v.GetConsAddr()
	require.NoError(t, err)
	require.Equal(t, consAddr, consAddr2)

	_, err = v.TmConsPublicKey()
	require.NoError(t, err)

	require.NotPanics(t, func() { _ = v.ABCIValidatorUpdate() })
	require.NotPanics(t, func() { _ = v.ABCIValidatorUpdateZero() })
	require.NotEmpty(t, v.String())

	registry := codectypes.NewInterfaceRegistry()
	cryptocodec.RegisterInterfaces(registry)
	require.NoError(t, v.UnpackInterfaces(registry))
}

func TestValidator_GetOwner(t *testing.T) {
	require.Nil(t, types.Validator{Owner: ""}.GetOwner())
	require.Panics(t, func() { _ = types.Validator{Owner: "invalid-address"}.GetOwner() })
}

func TestValidator_ConsKeyErrors(t *testing.T) {
	// PubKey Any whose cached value is not a cryptotypes.PubKey.
	bad := types.Validator{PubKey: mustNonPubKeyAny(t)}
	_, err := bad.GetConsAddress()
	require.Error(t, err)
	_, err = bad.GetConsPubKey()
	require.Error(t, err)
	_, err = bad.GetConsAddr()
	require.Error(t, err)
	_, err = bad.TmConsPublicKey()
	require.Error(t, err)
}

func mustNonPubKeyAny(t *testing.T) *codectypes.Any {
	t.Helper()
	// Use a proto message that is NOT a cryptotypes.PubKey.
	a, err := codectypes.NewAnyWithValue(&types.Validator{Owner: "x"})
	require.NoError(t, err)

	return a
}

func TestValidator_MarshalRoundtrip(t *testing.T) {
	v, _ := newTestValidator(t)
	registry := codectypes.NewInterfaceRegistry()
	cryptocodec.RegisterInterfaces(registry)
	cdc := codec.NewProtoCodec(registry)

	bz := types.MustMarshalValidator(cdc, &v)
	require.NotEmpty(t, bz)

	got := types.MustUnmarshalValidator(cdc, bz)
	require.Equal(t, v.Owner, got.Owner)

	got2, err := types.UnmarshalValidator(cdc, bz)
	require.NoError(t, err)
	require.Equal(t, v.Owner, got2.Owner)
}

func TestDescription_Validate(t *testing.T) {
	for _, tc := range []struct {
		name  string
		desc  types.Description
		valid bool
	}{
		{"valid", types.NewDescription("m", "i", "w", "d"), true},
		{"empty moniker", types.NewDescription("", "i", "w", "d"), false},
		{"moniker too long", types.NewDescription(strings.Repeat("a", types.MaxNameLength+1), "", "", ""), false},
		{"identity too long", types.NewDescription("m", strings.Repeat("a", types.MaxIdentityLength+1), "", ""), false},
		{"website too long", types.NewDescription("m", "", strings.Repeat("a", types.MaxWebsiteLength+1), ""), false},
		{"details too long", types.NewDescription("m", "", "", strings.Repeat("a", types.MaxDetailsLength+1)), false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.desc.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
	require.NotEmpty(t, types.NewDescription("m", "", "", "").String())
}

func TestLastValidatorPower_Methods(t *testing.T) {
	addr, err := sdk.ValAddressFromBech32(sample.ValAddress())
	require.NoError(t, err)

	lvp := types.NewLastValidatorPower(addr)
	require.Equal(t, addr, lvp.GetOwner())
	require.Equal(t, types.Power, lvp.GetPower())
	require.NotEmpty(t, lvp.String())

	require.Nil(t, types.LastValidatorPower{Owner: ""}.GetOwner())
	require.Panics(t, func() { _ = types.LastValidatorPower{Owner: "invalid"}.GetOwner() })
}
