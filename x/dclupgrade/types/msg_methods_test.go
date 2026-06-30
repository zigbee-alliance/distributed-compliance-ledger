package types

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/zigbee-alliance/distributed-compliance-ledger/testutil/sample"
)

type legacyMsg interface {
	Route() string
	Type() string
	GetSignBytes() []byte
	GetSigners() []sdk.AccAddress
}

func requireMsgMethods(t *testing.T, msg legacyMsg) {
	t.Helper()
	require.NotEmpty(t, msg.Route())
	require.NotEmpty(t, msg.Type())
	require.NotPanics(t, func() { _ = msg.GetSignBytes() })
	require.Len(t, msg.GetSigners(), 1)
}

func TestDclupgradeMsgMethods(t *testing.T) {
	creator := sample.AccAddress()
	msgs := []legacyMsg{
		&MsgApproveUpgrade{Creator: creator},
		&MsgProposeUpgrade{Creator: creator},
		&MsgRejectUpgrade{Creator: creator},
	}
	for _, msg := range msgs {
		t.Run(msg.Type(), func(t *testing.T) {
			requireMsgMethods(t, msg)
		})
	}
}

func TestDclupgradeMsgGetSignersPanicsOnInvalidSigner(t *testing.T) {
	msgs := []legacyMsg{
		&MsgApproveUpgrade{Creator: "invalid"},
		&MsgProposeUpgrade{Creator: "invalid"},
		&MsgRejectUpgrade{Creator: "invalid"},
	}
	for _, msg := range msgs {
		t.Run(msg.Type(), func(t *testing.T) {
			require.Panics(t, func() { _ = msg.GetSigners() })
		})
	}
}
