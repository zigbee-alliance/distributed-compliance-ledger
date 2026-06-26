package types_test

import (
	"testing"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/zigbee-alliance/distributed-compliance-ledger/testutil/sample"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
)

func TestMsgProposeAddAccount_ValidateBasicBranches(t *testing.T) {
	pk := ed25519.GenPrivKey().PubKey()
	accAddr := sdk.AccAddress(pk.Address())
	pkAny, err := codectypes.NewAnyWithValue(pk)
	require.NoError(t, err)
	badAny, err := codectypes.NewAnyWithValue(&types.MsgProposeAddAccount{})
	require.NoError(t, err)
	signer := sample.AccAddress()

	base := func() types.MsgProposeAddAccount {
		return types.MsgProposeAddAccount{
			Address:  accAddr.String(),
			Signer:   signer,
			PubKey:   pkAny,
			Roles:    types.AccountRoles{types.Vendor},
			VendorID: 1,
		}
	}

	mutate := func(f func(*types.MsgProposeAddAccount)) types.MsgProposeAddAccount {
		m := base()
		f(&m)

		return m
	}

	for _, tc := range []struct {
		name string
		msg  types.MsgProposeAddAccount
	}{
		{"empty address", mutate(func(m *types.MsgProposeAddAccount) { m.Address = "" })},
		{"invalid address", mutate(func(m *types.MsgProposeAddAccount) { m.Address = "invalid" })},
		{"empty signer", mutate(func(m *types.MsgProposeAddAccount) { m.Signer = "" })},
		{"invalid signer", mutate(func(m *types.MsgProposeAddAccount) { m.Signer = "invalid" })},
		{"nil pubkey", mutate(func(m *types.MsgProposeAddAccount) { m.PubKey = nil })},
		{"non-pubkey any", mutate(func(m *types.MsgProposeAddAccount) { m.PubKey = badAny })},
		{"pubkey/address mismatch", mutate(func(m *types.MsgProposeAddAccount) { m.Address = sample.AccAddress() })},
		{"no roles", mutate(func(m *types.MsgProposeAddAccount) { m.Roles = nil })},
		{"invalid role", mutate(func(m *types.MsgProposeAddAccount) { m.Roles = types.AccountRoles{types.AccountRole("bogus")} })},
		{"vendor without vendorID", mutate(func(m *types.MsgProposeAddAccount) { m.VendorID = 0 })},
	} {
		t.Run(tc.name, func(t *testing.T) {
			msg := tc.msg
			require.Error(t, msg.ValidateBasic())
		})
	}
}
