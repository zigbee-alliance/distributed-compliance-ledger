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

func TestDclauthMsgMethods(t *testing.T) {
	signer := sample.AccAddress()
	msgs := []legacyMsg{
		&MsgApproveAddAccount{Signer: signer},
		&MsgApproveRevokeAccount{Signer: signer},
		&MsgProposeAddAccount{Signer: signer},
		&MsgProposeRevokeAccount{Signer: signer},
		&MsgRejectAddAccount{Signer: signer},
	}
	for _, msg := range msgs {
		t.Run(msg.Type(), func(t *testing.T) {
			requireMsgMethods(t, msg)
		})
	}
}

func TestDclauthMsgGetSignersPanicsOnInvalidSigner(t *testing.T) {
	msgs := []legacyMsg{
		&MsgApproveAddAccount{Signer: "invalid"},
		&MsgApproveRevokeAccount{Signer: "invalid"},
		&MsgProposeAddAccount{Signer: "invalid"},
		&MsgProposeRevokeAccount{Signer: "invalid"},
		&MsgRejectAddAccount{Signer: "invalid"},
	}
	for _, msg := range msgs {
		t.Run(msg.Type(), func(t *testing.T) {
			require.Panics(t, func() { _ = msg.GetSigners() })
		})
	}
}
