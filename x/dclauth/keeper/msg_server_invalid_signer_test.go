package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
)

func TestMsgServerInvalidSigner(t *testing.T) {
	ms, ctx := setupMsgServer(t)
	const bad = "invalid-signer-address"

	cases := []struct {
		name string
		call func() error
	}{
		{"ProposeAddAccount", func() error {
			_, err := ms.ProposeAddAccount(ctx, &types.MsgProposeAddAccount{Signer: bad})
			return err
		}},
		{"ApproveAddAccount", func() error {
			_, err := ms.ApproveAddAccount(ctx, &types.MsgApproveAddAccount{Signer: bad})
			return err
		}},
		{"RejectAddAccount", func() error {
			_, err := ms.RejectAddAccount(ctx, &types.MsgRejectAddAccount{Signer: bad})
			return err
		}},
		{"ProposeRevokeAccount", func() error {
			_, err := ms.ProposeRevokeAccount(ctx, &types.MsgProposeRevokeAccount{Signer: bad})
			return err
		}},
		{"ApproveRevokeAccount", func() error {
			_, err := ms.ApproveRevokeAccount(ctx, &types.MsgApproveRevokeAccount{Signer: bad})
			return err
		}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			require.Error(t, tc.call())
		})
	}
}
