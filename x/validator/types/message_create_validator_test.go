package types_test

import (
	"strings"
	"testing"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	"github.com/stretchr/testify/require"
	"github.com/zigbee-alliance/distributed-compliance-ledger/testutil/sample"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/validator/types"
)

func TestMsgCreateValidator_ValidateBasicBranches(t *testing.T) {
	goodPk, err := codectypes.NewAnyWithValue(ed25519.GenPrivKey().PubKey())
	require.NoError(t, err)
	badPk, err := codectypes.NewAnyWithValue(&types.Validator{Owner: "x"})
	require.NoError(t, err)

	signer := sample.ValAddress()
	goodDesc := types.NewDescription("moniker", "identity", "website", "details")

	for _, tc := range []struct {
		name string
		msg  types.MsgCreateValidator
	}{
		{"invalid signer", types.MsgCreateValidator{Signer: "invalid"}},
		{"nil pubkey", types.MsgCreateValidator{Signer: signer, PubKey: nil, Description: goodDesc}},
		{"non-pubkey any", types.MsgCreateValidator{Signer: signer, PubKey: badPk, Description: goodDesc}},
		{"empty description", types.MsgCreateValidator{Signer: signer, PubKey: goodPk}},
		{"invalid description", types.MsgCreateValidator{
			Signer:      signer,
			PubKey:      goodPk,
			Description: types.NewDescription(strings.Repeat("a", types.MaxNameLength+1), "", "", ""),
		}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			msg := tc.msg
			require.Error(t, msg.ValidateBasic())
		})
	}
}
