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

func TestValidatorMsgMethods(t *testing.T) {
	acc := sample.AccAddress()
	val := sample.ValAddress()
	msgs := []legacyMsg{
		// signer is a validator operator (valoper) address
		&MsgCreateValidator{Signer: val},
		&MsgDisableValidator{Creator: val},
		&MsgEnableValidator{Creator: val},
		// signer is a regular account address
		&MsgApproveDisableValidator{Creator: acc},
		&MsgProposeDisableValidator{Creator: acc},
		&MsgRejectDisableValidator{Creator: acc},
	}
	for _, msg := range msgs {
		t.Run(msg.Type(), func(t *testing.T) {
			requireMsgMethods(t, msg)
		})
	}
}

func TestValidatorMsgGetSignersPanicsOnInvalidSigner(t *testing.T) {
	msgs := []legacyMsg{
		&MsgCreateValidator{Signer: "invalid"},
		&MsgDisableValidator{Creator: "invalid"},
		&MsgEnableValidator{Creator: "invalid"},
		&MsgApproveDisableValidator{Creator: "invalid"},
		&MsgProposeDisableValidator{Creator: "invalid"},
		&MsgRejectDisableValidator{Creator: "invalid"},
	}
	for _, msg := range msgs {
		t.Run(msg.Type(), func(t *testing.T) {
			require.Panics(t, func() { _ = msg.GetSigners() })
		})
	}
}
