package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

// TestMsgServerInvalidSigner exercises the invalid-signer-address error branch
// that fronts every pki transaction handler. The address is parsed before any
// keeper dependency is touched, so a nil dclauth keeper is fine here.
func TestMsgServerInvalidSigner(t *testing.T) {
	ms, ctx := setupMsgServer(t)
	const bad = "invalid-signer-address"

	cases := []struct {
		name string
		call func() error
	}{
		{"AddX509Cert", func() error {
			_, err := ms.AddX509Cert(ctx, &types.MsgAddX509Cert{Signer: bad})

			return err
		}},
		{"AddNocX509RootCert", func() error {
			_, err := ms.AddNocX509RootCert(ctx, &types.MsgAddNocX509RootCert{Signer: bad})

			return err
		}},
		{"AddNocX509IcaCert", func() error {
			_, err := ms.AddNocX509IcaCert(ctx, &types.MsgAddNocX509IcaCert{Signer: bad})

			return err
		}},
		{"AddPkiRevocationDistributionPoint", func() error {
			_, err := ms.AddPkiRevocationDistributionPoint(ctx, &types.MsgAddPkiRevocationDistributionPoint{Signer: bad})

			return err
		}},
		{"UpdatePkiRevocationDistributionPoint", func() error {
			_, err := ms.UpdatePkiRevocationDistributionPoint(ctx, &types.MsgUpdatePkiRevocationDistributionPoint{Signer: bad})

			return err
		}},
		{"DeletePkiRevocationDistributionPoint", func() error {
			_, err := ms.DeletePkiRevocationDistributionPoint(ctx, &types.MsgDeletePkiRevocationDistributionPoint{Signer: bad})

			return err
		}},
		{"ApproveAddX509RootCert", func() error {
			_, err := ms.ApproveAddX509RootCert(ctx, &types.MsgApproveAddX509RootCert{Signer: bad})

			return err
		}},
		{"ApproveRevokeX509RootCert", func() error {
			_, err := ms.ApproveRevokeX509RootCert(ctx, &types.MsgApproveRevokeX509RootCert{Signer: bad})

			return err
		}},
		{"AssignVid", func() error {
			_, err := ms.AssignVid(ctx, &types.MsgAssignVid{Signer: bad})

			return err
		}},
		{"ProposeAddX509RootCert", func() error {
			_, err := ms.ProposeAddX509RootCert(ctx, &types.MsgProposeAddX509RootCert{Signer: bad})

			return err
		}},
		{"ProposeRevokeX509RootCert", func() error {
			_, err := ms.ProposeRevokeX509RootCert(ctx, &types.MsgProposeRevokeX509RootCert{Signer: bad})

			return err
		}},
		{"RejectAddX509RootCert", func() error {
			_, err := ms.RejectAddX509RootCert(ctx, &types.MsgRejectAddX509RootCert{Signer: bad})

			return err
		}},
		{"RemoveX509Cert", func() error {
			_, err := ms.RemoveX509Cert(ctx, &types.MsgRemoveX509Cert{Signer: bad})

			return err
		}},
		{"RemoveNocX509RootCert", func() error {
			_, err := ms.RemoveNocX509RootCert(ctx, &types.MsgRemoveNocX509RootCert{Signer: bad})

			return err
		}},
		{"RemoveNocX509IcaCert", func() error {
			_, err := ms.RemoveNocX509IcaCert(ctx, &types.MsgRemoveNocX509IcaCert{Signer: bad})

			return err
		}},
		{"RevokeX509Cert", func() error {
			_, err := ms.RevokeX509Cert(ctx, &types.MsgRevokeX509Cert{Signer: bad})

			return err
		}},
		{"RevokeNocX509RootCert", func() error {
			_, err := ms.RevokeNocX509RootCert(ctx, &types.MsgRevokeNocX509RootCert{Signer: bad})

			return err
		}},
		{"RevokeNocX509IcaCert", func() error {
			_, err := ms.RevokeNocX509IcaCert(ctx, &types.MsgRevokeNocX509IcaCert{Signer: bad})

			return err
		}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			require.Error(t, tc.call())
		})
	}
}
