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

func TestModelMsgMethods(t *testing.T) {
	creator := sample.AccAddress()
	msgs := []legacyMsg{
		&MsgCreateModel{Creator: creator},
		&MsgUpdateModel{Creator: creator},
		&MsgDeleteModel{Creator: creator},
		&MsgCreateModelVersion{Creator: creator},
		&MsgUpdateModelVersion{Creator: creator},
		&MsgDeleteModelVersion{Creator: creator},
	}
	for _, msg := range msgs {
		t.Run(msg.Type(), func(t *testing.T) {
			requireMsgMethods(t, msg)
		})
	}
}

func TestModelMsgGetSignersPanicsOnInvalidSigner(t *testing.T) {
	msgs := []legacyMsg{
		&MsgCreateModel{Creator: "invalid"},
		&MsgUpdateModel{Creator: "invalid"},
		&MsgDeleteModel{Creator: "invalid"},
		&MsgCreateModelVersion{Creator: "invalid"},
		&MsgUpdateModelVersion{Creator: "invalid"},
		&MsgDeleteModelVersion{Creator: "invalid"},
	}
	for _, msg := range msgs {
		t.Run(msg.Type(), func(t *testing.T) {
			require.Panics(t, func() { _ = msg.GetSigners() })
		})
	}
}
