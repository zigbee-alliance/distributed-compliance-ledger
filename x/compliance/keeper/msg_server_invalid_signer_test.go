package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "github.com/zigbee-alliance/distributed-compliance-ledger/testutil/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliance/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliance/types"
)

func TestMsgServerInvalidSigner(t *testing.T) {
	k, ctx := keepertest.ComplianceKeeper(t, nil, nil)
	ms := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)
	const bad = "invalid-signer-address"

	cases := []struct {
		name string
		call func() error
	}{
		{"CertifyModel", func() error {
			_, err := ms.CertifyModel(wctx, &types.MsgCertifyModel{Signer: bad})

			return err
		}},
		{"RevokeModel", func() error {
			_, err := ms.RevokeModel(wctx, &types.MsgRevokeModel{Signer: bad})

			return err
		}},
		{"ProvisionModel", func() error {
			_, err := ms.ProvisionModel(wctx, &types.MsgProvisionModel{Signer: bad})

			return err
		}},
		{"UpdateComplianceInfo", func() error {
			_, err := ms.UpdateComplianceInfo(wctx, &types.MsgUpdateComplianceInfo{Creator: bad})

			return err
		}},
		{"DeleteComplianceInfo", func() error {
			_, err := ms.DeleteComplianceInfo(wctx, &types.MsgDeleteComplianceInfo{Creator: bad})

			return err
		}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			require.Error(t, tc.call())
		})
	}
}
