package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/dclupgrade/types"
)

func TestMsgServerInvalidSigner(t *testing.T) {
	ms, ctx := setupMsgServer(t)
	const bad = "invalid-signer-address"

	cases := []struct {
		name string
		call func() error
	}{
		{"ProposeUpgrade", func() error {
			_, err := ms.ProposeUpgrade(ctx, &types.MsgProposeUpgrade{Creator: bad})
			return err
		}},
		{"ApproveUpgrade", func() error {
			_, err := ms.ApproveUpgrade(ctx, &types.MsgApproveUpgrade{Creator: bad})
			return err
		}},
		{"RejectUpgrade", func() error {
			_, err := ms.RejectUpgrade(ctx, &types.MsgRejectUpgrade{Creator: bad})
			return err
		}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			require.Error(t, tc.call())
		})
	}
}
