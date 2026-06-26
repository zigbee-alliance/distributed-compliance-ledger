package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/model/types"
)

func TestMsgServerInvalidSigner(t *testing.T) {
	ms, ctx := setupMsgServer(t)
	const bad = "invalid-signer-address"

	cases := []struct {
		name string
		call func() error
	}{
		{"CreateModel", func() error {
			_, err := ms.CreateModel(ctx, &types.MsgCreateModel{Creator: bad})
			return err
		}},
		{"UpdateModel", func() error {
			_, err := ms.UpdateModel(ctx, &types.MsgUpdateModel{Creator: bad})
			return err
		}},
		{"DeleteModel", func() error {
			_, err := ms.DeleteModel(ctx, &types.MsgDeleteModel{Creator: bad})
			return err
		}},
		{"CreateModelVersion", func() error {
			_, err := ms.CreateModelVersion(ctx, &types.MsgCreateModelVersion{Creator: bad})
			return err
		}},
		{"UpdateModelVersion", func() error {
			_, err := ms.UpdateModelVersion(ctx, &types.MsgUpdateModelVersion{Creator: bad})
			return err
		}},
		{"DeleteModelVersion", func() error {
			_, err := ms.DeleteModelVersion(ctx, &types.MsgDeleteModelVersion{Creator: bad})
			return err
		}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			require.Error(t, tc.call())
		})
	}
}
