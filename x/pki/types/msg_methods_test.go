package types

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/zigbee-alliance/distributed-compliance-ledger/testutil/sample"
)

// legacyMsg captures the standard cosmos sdk.LegacyMsg boilerplate implemented
// by every transaction message in this module.
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

func TestPkiMsgMethods(t *testing.T) {
	signer := sample.AccAddress()
	msgs := []legacyMsg{
		&MsgAddNocX509IcaCert{Signer: signer},
		&MsgAddNocX509RootCert{Signer: signer},
		&MsgAddPkiRevocationDistributionPoint{Signer: signer},
		&MsgAddX509Cert{Signer: signer},
		&MsgApproveAddX509RootCert{Signer: signer},
		&MsgApproveRevokeX509RootCert{Signer: signer},
		&MsgAssignVid{Signer: signer},
		&MsgDeletePkiRevocationDistributionPoint{Signer: signer},
		&MsgProposeAddX509RootCert{Signer: signer},
		&MsgProposeRevokeX509RootCert{Signer: signer},
		&MsgRejectAddX509RootCert{Signer: signer},
		&MsgRemoveNocX509IcaCert{Signer: signer},
		&MsgRemoveNocX509RootCert{Signer: signer},
		&MsgRemoveX509Cert{Signer: signer},
		&MsgRevokeNocX509IcaCert{Signer: signer},
		&MsgRevokeNocX509RootCert{Signer: signer},
		&MsgRevokeX509Cert{Signer: signer},
		&MsgUpdatePkiRevocationDistributionPoint{Signer: signer},
	}
	for _, msg := range msgs {
		t.Run(msg.Type(), func(t *testing.T) {
			requireMsgMethods(t, msg)
		})
	}
}

func TestPkiMsgGetSignersPanicsOnInvalidSigner(t *testing.T) {
	msgs := []legacyMsg{
		&MsgAddNocX509IcaCert{Signer: "invalid"},
		&MsgAddNocX509RootCert{Signer: "invalid"},
		&MsgAddPkiRevocationDistributionPoint{Signer: "invalid"},
		&MsgAddX509Cert{Signer: "invalid"},
		&MsgApproveAddX509RootCert{Signer: "invalid"},
		&MsgApproveRevokeX509RootCert{Signer: "invalid"},
		&MsgAssignVid{Signer: "invalid"},
		&MsgDeletePkiRevocationDistributionPoint{Signer: "invalid"},
		&MsgProposeAddX509RootCert{Signer: "invalid"},
		&MsgProposeRevokeX509RootCert{Signer: "invalid"},
		&MsgRejectAddX509RootCert{Signer: "invalid"},
		&MsgRemoveNocX509IcaCert{Signer: "invalid"},
		&MsgRemoveNocX509RootCert{Signer: "invalid"},
		&MsgRemoveX509Cert{Signer: "invalid"},
		&MsgRevokeNocX509IcaCert{Signer: "invalid"},
		&MsgRevokeNocX509RootCert{Signer: "invalid"},
		&MsgRevokeX509Cert{Signer: "invalid"},
		&MsgUpdatePkiRevocationDistributionPoint{Signer: "invalid"},
	}
	for _, msg := range msgs {
		t.Run(msg.Type(), func(t *testing.T) {
			require.Panics(t, func() { _ = msg.GetSigners() })
		})
	}
}
