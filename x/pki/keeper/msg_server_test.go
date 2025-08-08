package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "github.com/zigbee-alliance/distributed-compliance-ledger/testutil/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

const (
	unauthorizedError = "unauthorized"
)

//nolint:unused
func setupMsgServer(tb testing.TB) (types.MsgServer, context.Context) {
	tb.Helper()
	k, ctx := keepertest.PkiKeeper(tb, nil)

	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}

func TestMsgServer_ProposeAddX509RootCert(t *testing.T) {
	msgServer, ctx := setupMsgServer(t)
	
	msg := &types.MsgProposeAddX509RootCert{
		Cert:   "test-cert",
		Signer: "cosmos1test",
	}
	
	_, err := msgServer.ProposeAddX509RootCert(ctx, msg)
	require.Error(t, err)
	require.Contains(t, err.Error(), unauthorizedError)
}

func TestMsgServer_ApproveAddX509RootCert(t *testing.T) {
	msgServer, ctx := setupMsgServer(t)
	
	msg := &types.MsgApproveAddX509RootCert{
		Subject:      "test-subject",
		SubjectKeyId: "test-key-id",
		Signer:       "cosmos1test",
	}
	
	_, err := msgServer.ApproveAddX509RootCert(ctx, msg)
	require.Error(t, err)
	require.Contains(t, err.Error(), unauthorizedError)
}

func TestMsgServer_AddX509Cert(t *testing.T) {
	msgServer, ctx := setupMsgServer(t)
	
	msg := &types.MsgAddX509Cert{
		Cert:   "test-cert",
		Signer: "cosmos1test",
	}
	
	_, err := msgServer.AddX509Cert(ctx, msg)
	require.Error(t, err)
	require.Contains(t, err.Error(), unauthorizedError)
}

func TestMsgServer_ProposeRevokeX509RootCert(t *testing.T) {
	msgServer, ctx := setupMsgServer(t)
	
	msg := &types.MsgProposeRevokeX509RootCert{
		Subject:      "test-subject",
		SubjectKeyId: "test-key-id",
		Signer:       "cosmos1test",
	}
	
	_, err := msgServer.ProposeRevokeX509RootCert(ctx, msg)
	require.Error(t, err)
	require.Contains(t, err.Error(), unauthorizedError)
}

func TestMsgServer_ApproveRevokeX509RootCert(t *testing.T) {
	msgServer, ctx := setupMsgServer(t)
	
	msg := &types.MsgApproveRevokeX509RootCert{
		Subject:      "test-subject",
		SubjectKeyId: "test-key-id",
		Signer:       "cosmos1test",
	}
	
	_, err := msgServer.ApproveRevokeX509RootCert(ctx, msg)
	require.Error(t, err)
	require.Contains(t, err.Error(), unauthorizedError)
}

func TestMsgServer_RevokeX509Cert(t *testing.T) {
	msgServer, ctx := setupMsgServer(t)
	
	msg := &types.MsgRevokeX509Cert{
		Subject:      "test-subject",
		SubjectKeyId: "test-key-id",
		Signer:       "cosmos1test",
	}
	
	_, err := msgServer.RevokeX509Cert(ctx, msg)
	require.Error(t, err)
	require.Contains(t, err.Error(), unauthorizedError)
}

func TestMsgServer_RejectAddX509RootCert(t *testing.T) {
	msgServer, ctx := setupMsgServer(t)
	
	msg := &types.MsgRejectAddX509RootCert{
		Subject:      "test-subject",
		SubjectKeyId: "test-key-id",
		Signer:       "cosmos1test",
	}
	
	_, err := msgServer.RejectAddX509RootCert(ctx, msg)
	require.Error(t, err)
	require.Contains(t, err.Error(), unauthorizedError)
}

func TestMsgServer_AddPkiRevocationDistributionPoint(t *testing.T) {
	msgServer, ctx := setupMsgServer(t)
	
	msg := &types.MsgAddPkiRevocationDistributionPoint{
		Vid:                  1,
		Label:                "test-label",
		IssuerSubjectKeyID:   "test-issuer",
		DataURL:              "https://test.com",
		DataFileSize:         100,
		DataDigest:           "test-digest",
		DataDigestType:       1,
		CrlSignerCertificate: "test-cert",
		Signer:               "cosmos1test",
	}
	
	_, err := msgServer.AddPkiRevocationDistributionPoint(ctx, msg)
	require.Error(t, err)
	require.Contains(t, err.Error(), unauthorizedError)
}

func TestMsgServer_UpdatePkiRevocationDistributionPoint(t *testing.T) {
	msgServer, ctx := setupMsgServer(t)
	
	msg := &types.MsgUpdatePkiRevocationDistributionPoint{
		Vid:                  1,
		Label:                "test-label",
		IssuerSubjectKeyID:   "test-issuer",
		DataURL:              "https://test.com",
		DataFileSize:         100,
		DataDigest:           "test-digest",
		DataDigestType:       1,
		CrlSignerCertificate: "test-cert",
		Signer:               "cosmos1test",
	}
	
	_, err := msgServer.UpdatePkiRevocationDistributionPoint(ctx, msg)
	require.Error(t, err)
	require.Contains(t, err.Error(), unauthorizedError)
}

func TestMsgServer_DeletePkiRevocationDistributionPoint(t *testing.T) {
	msgServer, ctx := setupMsgServer(t)
	
	msg := &types.MsgDeletePkiRevocationDistributionPoint{
		Vid:                1,
		Label:              "test-label",
		IssuerSubjectKeyID: "test-issuer",
		Signer:             "cosmos1test",
	}
	
	_, err := msgServer.DeletePkiRevocationDistributionPoint(ctx, msg)
	require.Error(t, err)
	require.Contains(t, err.Error(), unauthorizedError)
}

func TestMsgServer_AssignVid(t *testing.T) {
	msgServer, ctx := setupMsgServer(t)
	
	msg := &types.MsgAssignVid{
		Vid:    1,
		Signer: "cosmos1test",
	}
	
	_, err := msgServer.AssignVid(ctx, msg)
	require.Error(t, err)
	require.Contains(t, err.Error(), unauthorizedError)
}

func TestMsgServer_AddNocX509RootCert(t *testing.T) {
	msgServer, ctx := setupMsgServer(t)
	
	msg := &types.MsgAddNocX509RootCert{
		Cert:   "test-cert",
		Signer: "cosmos1test",
	}
	
	_, err := msgServer.AddNocX509RootCert(ctx, msg)
	require.Error(t, err)
	require.Contains(t, err.Error(), unauthorizedError)
}

func TestMsgServer_RemoveX509Cert(t *testing.T) {
	msgServer, ctx := setupMsgServer(t)
	
	msg := &types.MsgRemoveX509Cert{
		Subject:      "test-subject",
		SubjectKeyId: "test-key-id",
		Signer:       "cosmos1test",
	}
	
	_, err := msgServer.RemoveX509Cert(ctx, msg)
	require.Error(t, err)
	require.Contains(t, err.Error(), unauthorizedError)
}

func TestMsgServer_AddNocX509IcaCert(t *testing.T) {
	msgServer, ctx := setupMsgServer(t)
	
	msg := &types.MsgAddNocX509IcaCert{
		Cert:   "test-cert",
		Signer: "cosmos1test",
	}
	
	_, err := msgServer.AddNocX509IcaCert(ctx, msg)
	require.Error(t, err)
	require.Contains(t, err.Error(), unauthorizedError)
}

func TestMsgServer_RevokeNocX509RootCert(t *testing.T) {
	msgServer, ctx := setupMsgServer(t)
	
	msg := &types.MsgRevokeNocX509RootCert{
		Subject:      "test-subject",
		SubjectKeyId: "test-key-id",
		Signer:       "cosmos1test",
	}
	
	_, err := msgServer.RevokeNocX509RootCert(ctx, msg)
	require.Error(t, err)
	require.Contains(t, err.Error(), unauthorizedError)
}

func TestMsgServer_RevokeNocX509IcaCert(t *testing.T) {
	msgServer, ctx := setupMsgServer(t)
	
	msg := &types.MsgRevokeNocX509IcaCert{
		Subject:      "test-subject",
		SubjectKeyId: "test-key-id",
		Signer:       "cosmos1test",
	}
	
	_, err := msgServer.RevokeNocX509IcaCert(ctx, msg)
	require.Error(t, err)
	require.Contains(t, err.Error(), unauthorizedError)
}

func TestMsgServer_RemoveNocX509IcaCert(t *testing.T) {
	msgServer, ctx := setupMsgServer(t)
	
	msg := &types.MsgRemoveNocX509IcaCert{
		Subject:      "test-subject",
		SubjectKeyId: "test-key-id",
		Signer:       "cosmos1test",
	}
	
	_, err := msgServer.RemoveNocX509IcaCert(ctx, msg)
	require.Error(t, err)
	require.Contains(t, err.Error(), unauthorizedError)
}

func TestMsgServer_RemoveNocX509RootCert(t *testing.T) {
	msgServer, ctx := setupMsgServer(t)
	
	msg := &types.MsgRemoveNocX509RootCert{
		Subject:      "test-subject",
		SubjectKeyId: "test-key-id",
		Signer:       "cosmos1test",
	}
	
	_, err := msgServer.RemoveNocX509RootCert(ctx, msg)
	require.Error(t, err)
	require.Contains(t, err.Error(), unauthorizedError)
}
