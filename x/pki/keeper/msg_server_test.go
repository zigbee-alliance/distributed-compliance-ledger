package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/zigbee-alliance/distributed-compliance-ledger/testutil/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

//nolint:unused
func setupMsgServer(tb testing.TB) (types.MsgServer, context.Context) {
	tb.Helper()
	k, ctx := keepertest.PkiKeeper(tb, nil)

	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}

func TestMsgServer_AllHandlers(t *testing.T) {
	msgServer, ctx := setupMsgServer(t)
	
	tests := []struct {
		name        string
		msg         interface{}
		expectError bool
		errorMsg    string
	}{
		{
			name: "ProposeAddX509RootCert",
			msg: &types.MsgProposeAddX509RootCert{
				Cert:   "test-cert",
				Signer: "cosmos1test",
			},
			expectError: true,
		},
		{
			name: "ApproveAddX509RootCert",
			msg: &types.MsgApproveAddX509RootCert{
				Subject:      "test-subject",
				SubjectKeyId: "test-key-id",
				Signer:       "cosmos1test",
			},
			expectError: true,
		},
		{
			name: "AddX509Cert",
			msg: &types.MsgAddX509Cert{
				Cert:   "test-cert",
				Signer: "cosmos1test",
			},
			expectError: true,
		},
		{
			name: "ProposeRevokeX509RootCert",
			msg: &types.MsgProposeRevokeX509RootCert{
				Subject:      "test-subject",
				SubjectKeyId: "test-key-id",
				Signer:       "cosmos1test",
			},
			expectError: true,
		},
		{
			name: "ApproveRevokeX509RootCert",
			msg: &types.MsgApproveRevokeX509RootCert{
				Subject:      "test-subject",
				SubjectKeyId: "test-key-id",
				Signer:       "cosmos1test",
			},
			expectError: true,
		},
		{
			name: "RevokeX509Cert",
			msg: &types.MsgRevokeX509Cert{
				Subject:      "test-subject",
				SubjectKeyId: "test-key-id",
				Signer:       "cosmos1test",
			},
			expectError: true,
		},
		{
			name: "RejectAddX509RootCert",
			msg: &types.MsgRejectAddX509RootCert{
				Subject:      "test-subject",
				SubjectKeyId: "test-key-id",
				Signer:       "cosmos1test",
			},
			expectError: true,
		},
		{
			name: "AddPkiRevocationDistributionPoint",
			msg: &types.MsgAddPkiRevocationDistributionPoint{
				Vid:                  1,
				Label:                "test-label",
				IssuerSubjectKeyID:   "test-issuer",
				DataURL:              "https://test.com",
				DataFileSize:         100,
				DataDigest:           "test-digest",
				DataDigestType:       1,
				CrlSignerCertificate: "test-cert",
				Signer:               "cosmos1test",
			},
			expectError: true,
		},
		{
			name: "UpdatePkiRevocationDistributionPoint",
			msg: &types.MsgUpdatePkiRevocationDistributionPoint{
				Vid:                  1,
				Label:                "test-label",
				IssuerSubjectKeyID:   "test-issuer",
				DataURL:              "https://test.com",
				DataFileSize:         100,
				DataDigest:           "test-digest",
				DataDigestType:       1,
				CrlSignerCertificate: "test-cert",
				Signer:               "cosmos1test",
			},
			expectError: true,
		},
		{
			name: "DeletePkiRevocationDistributionPoint",
			msg: &types.MsgDeletePkiRevocationDistributionPoint{
				Vid:                1,
				Label:              "test-label",
				IssuerSubjectKeyID: "test-issuer",
				Signer:             "cosmos1test",
			},
			expectError: true,
		},
		{
			name: "AssignVid",
			msg: &types.MsgAssignVid{
				Vid:    1,
				Signer: "cosmos1test",
			},
			expectError: true,
		},
		{
			name: "AddNocX509RootCert",
			msg: &types.MsgAddNocX509RootCert{
				Cert:   "test-cert",
				Signer: "cosmos1test",
			},
			expectError: true,
		},
		{
			name: "RemoveX509Cert",
			msg: &types.MsgRemoveX509Cert{
				Subject:      "test-subject",
				SubjectKeyId: "test-key-id",
				Signer:       "cosmos1test",
			},
			expectError: true,
		},
		{
			name: "AddNocX509IcaCert",
			msg: &types.MsgAddNocX509IcaCert{
				Cert:   "test-cert",
				Signer: "cosmos1test",
			},
			expectError: true,
		},
		{
			name: "RevokeNocX509RootCert",
			msg: &types.MsgRevokeNocX509RootCert{
				Subject:      "test-subject",
				SubjectKeyId: "test-key-id",
				Signer:       "cosmos1test",
			},
			expectError: true,
		},
		{
			name: "RevokeNocX509IcaCert",
			msg: &types.MsgRevokeNocX509IcaCert{
				Subject:      "test-subject",
				SubjectKeyId: "test-key-id",
				Signer:       "cosmos1test",
			},
			expectError: true,
		},
		{
			name: "RemoveNocX509IcaCert",
			msg: &types.MsgRemoveNocX509IcaCert{
				Subject:      "test-subject",
				SubjectKeyId: "test-key-id",
				Signer:       "cosmos1test",
			},
			expectError: true,
		},
		{
			name: "RemoveNocX509RootCert",
			msg: &types.MsgRemoveNocX509RootCert{
				Subject:      "test-subject",
				SubjectKeyId: "test-key-id",
				Signer:       "cosmos1test",
			},
			expectError: true,
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var err error
			switch msg := tt.msg.(type) {
			case *types.MsgProposeAddX509RootCert:
				_, err = msgServer.ProposeAddX509RootCert(ctx, msg)
			case *types.MsgApproveAddX509RootCert:
				_, err = msgServer.ApproveAddX509RootCert(ctx, msg)
			case *types.MsgAddX509Cert:
				_, err = msgServer.AddX509Cert(ctx, msg)
			case *types.MsgProposeRevokeX509RootCert:
				_, err = msgServer.ProposeRevokeX509RootCert(ctx, msg)
			case *types.MsgApproveRevokeX509RootCert:
				_, err = msgServer.ApproveRevokeX509RootCert(ctx, msg)
			case *types.MsgRevokeX509Cert:
				_, err = msgServer.RevokeX509Cert(ctx, msg)
			case *types.MsgRejectAddX509RootCert:
				_, err = msgServer.RejectAddX509RootCert(ctx, msg)
			case *types.MsgAddPkiRevocationDistributionPoint:
				_, err = msgServer.AddPkiRevocationDistributionPoint(ctx, msg)
			case *types.MsgUpdatePkiRevocationDistributionPoint:
				_, err = msgServer.UpdatePkiRevocationDistributionPoint(ctx, msg)
			case *types.MsgDeletePkiRevocationDistributionPoint:
				_, err = msgServer.DeletePkiRevocationDistributionPoint(ctx, msg)
			case *types.MsgAssignVid:
				_, err = msgServer.AssignVid(ctx, msg)
			case *types.MsgAddNocX509RootCert:
				_, err = msgServer.AddNocX509RootCert(ctx, msg)
			case *types.MsgRemoveX509Cert:
				_, err = msgServer.RemoveX509Cert(ctx, msg)
			case *types.MsgAddNocX509IcaCert:
				_, err = msgServer.AddNocX509IcaCert(ctx, msg)
			case *types.MsgRevokeNocX509RootCert:
				_, err = msgServer.RevokeNocX509RootCert(ctx, msg)
			case *types.MsgRevokeNocX509IcaCert:
				_, err = msgServer.RevokeNocX509IcaCert(ctx, msg)
			case *types.MsgRemoveNocX509IcaCert:
				_, err = msgServer.RemoveNocX509IcaCert(ctx, msg)
			case *types.MsgRemoveNocX509RootCert:
				_, err = msgServer.RemoveNocX509RootCert(ctx, msg)
			}
			
			if tt.expectError {
				require.Error(t, err)
				if tt.errorMsg != "" {
					require.Contains(t, err.Error(), tt.errorMsg)
				}
			} else {
				require.NoError(t, err)
			}
		})
	}
}
