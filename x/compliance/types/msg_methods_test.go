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

func TestComplianceMsgMethods(t *testing.T) {
	addr := sample.AccAddress()
	msgs := []legacyMsg{
		&MsgCertifyModel{Signer: addr},
		&MsgRevokeModel{Signer: addr},
		&MsgProvisionModel{Signer: addr},
		&MsgUpdateComplianceInfo{Creator: addr},
		&MsgDeleteComplianceInfo{Creator: addr},
	}
	for _, msg := range msgs {
		t.Run(msg.Type(), func(t *testing.T) {
			requireMsgMethods(t, msg)
		})
	}
}

func TestComplianceMsgGetSignersPanicsOnInvalidSigner(t *testing.T) {
	msgs := []legacyMsg{
		&MsgCertifyModel{Signer: "invalid"},
		&MsgRevokeModel{Signer: "invalid"},
		&MsgProvisionModel{Signer: "invalid"},
		&MsgUpdateComplianceInfo{Creator: "invalid"},
		&MsgDeleteComplianceInfo{Creator: "invalid"},
	}
	for _, msg := range msgs {
		t.Run(msg.Type(), func(t *testing.T) {
			require.Panics(t, func() { _ = msg.GetSigners() })
		})
	}
}
