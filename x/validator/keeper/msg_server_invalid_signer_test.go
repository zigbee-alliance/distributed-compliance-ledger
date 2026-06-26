package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/validator/types"
)

func TestMsgServerInvalidSigner(t *testing.T) {
	ms, ctx := setupMsgServer(t)
	const bad = "invalid-signer-address"

	cases := []struct {
		name string
		call func() error
	}{
		{"CreateValidator", func() error {
			_, err := ms.CreateValidator(ctx, &types.MsgCreateValidator{Signer: bad})
			return err
		}},
		{"DisableValidator", func() error {
			_, err := ms.DisableValidator(ctx, &types.MsgDisableValidator{Creator: bad})
			return err
		}},
		{"EnableValidator", func() error {
			_, err := ms.EnableValidator(ctx, &types.MsgEnableValidator{Creator: bad})
			return err
		}},
		{"ApproveDisableValidator", func() error {
			_, err := ms.ApproveDisableValidator(ctx, &types.MsgApproveDisableValidator{Creator: bad})
			return err
		}},
		{"ProposeDisableValidator", func() error {
			_, err := ms.ProposeDisableValidator(ctx, &types.MsgProposeDisableValidator{Creator: bad})
			return err
		}},
		{"RejectDisableValidator", func() error {
			_, err := ms.RejectDisableValidator(ctx, &types.MsgRejectDisableValidator{Creator: bad})
			return err
		}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			require.Error(t, tc.call())
		})
	}
}
