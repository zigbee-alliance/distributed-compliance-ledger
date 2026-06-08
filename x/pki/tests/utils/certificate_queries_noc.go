package utils

import (
	"github.com/stretchr/testify/require"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

func QueryNocCertificates(
	setup *TestSetup,
	subject string,
	subjectKeyID string,
) (*types.NocCertificates, error) {
	req := &types.QueryGetNocCertificatesRequest{
		Subject:      subject,
		SubjectKeyId: subjectKeyID,
	}

	resp, err := setup.Keeper.NocCertificates(setup.Wctx, req)
	if err != nil {
		require.Nil(setup.T, resp)

		return nil, err
	}

	require.NotNil(setup.T, resp)

	return &resp.NocCertificates, nil
}

func QueryNocCertificatesByVidAndSkid(
	setup *TestSetup,
	vid int32,
	subjectKeyID string,
) (*types.NocCertificatesByVidAndSkid, error) {
	req := &types.QueryGetNocCertificatesByVidAndSkidRequest{
		Vid:          vid,
		SubjectKeyId: subjectKeyID,
	}

	resp, err := setup.Keeper.NocCertificatesByVidAndSkid(setup.Wctx, req)
	if err != nil {
		require.Nil(setup.T, resp)

		return nil, err
	}

	require.NotNil(setup.T, resp)

	return &resp.NocCertificatesByVidAndSkid, nil
}

func QueryNocCertificatesBySubject(
	setup *TestSetup,
	subject string,
) (*types.NocCertificatesBySubject, error) {
	req := &types.QueryGetNocCertificatesBySubjectRequest{
		Subject: subject,
	}

	resp, err := setup.Keeper.NocCertificatesBySubject(setup.Wctx, req)
	if err != nil {
		require.Nil(setup.T, resp)

		return nil, err
	}

	require.NotNil(setup.T, resp)

	return &resp.NocCertificatesBySubject, nil
}

func QueryNocCertificatesBySubjectKeyID(
	setup *TestSetup,
	subjectKeyID string,
) ([]types.NocCertificates, error) {
	req := &types.QueryNocCertificatesRequest{
		SubjectKeyId: subjectKeyID,
	}

	resp, err := setup.Keeper.NocCertificatesAll(setup.Wctx, req)
	if err != nil {
		require.Nil(setup.T, resp)

		return nil, err
	}

	require.NotNil(setup.T, resp)

	return resp.NocCertificates, nil
}

func QueryNocRootCertificatesByVid(
	setup *TestSetup,
	vid int32,
) (*types.NocRootCertificates, error) {
	req := &types.QueryGetNocRootCertificatesRequest{
		Vid: vid,
	}

	resp, err := setup.Keeper.NocRootCertificates(setup.Wctx, req)
	if err != nil {
		require.Nil(setup.T, resp)

		return nil, err
	}

	require.NotNil(setup.T, resp)

	return &resp.NocRootCertificates, nil
}

func QueryNocIcaCertificatesByVid(
	setup *TestSetup,
	vid int32,
) (*types.NocIcaCertificates, error) {
	req := &types.QueryGetNocIcaCertificatesRequest{
		Vid: vid,
	}

	resp, err := setup.Keeper.NocIcaCertificates(setup.Wctx, req)
	if err != nil {
		require.Nil(setup.T, resp)

		return nil, err
	}

	require.NotNil(setup.T, resp)

	return &resp.NocIcaCertificates, nil
}

func QueryAllNocCertificates(
	setup *TestSetup,
) ([]types.NocCertificates, error) {
	req := &types.QueryNocCertificatesRequest{}

	resp, err := setup.Keeper.NocCertificatesAll(setup.Wctx, req)
	if err != nil {
		require.Nil(setup.T, resp)

		return nil, err
	}

	require.NotNil(setup.T, resp)

	return resp.NocCertificates, nil
}

func QueryAllNocRootCertificates(
	setup *TestSetup,
) ([]types.NocRootCertificates, error) {
	req := &types.QueryAllNocRootCertificatesRequest{}

	resp, err := setup.Keeper.NocRootCertificatesAll(setup.Wctx, req)
	if err != nil {
		require.Nil(setup.T, resp)

		return nil, err
	}

	require.NotNil(setup.T, resp)

	return resp.NocRootCertificates, nil
}

func QueryAllNocIcaCertificates(
	setup *TestSetup,
) ([]types.NocIcaCertificates, error) {
	req := &types.QueryAllNocIcaCertificatesRequest{}

	resp, err := setup.Keeper.NocIcaCertificatesAll(setup.Wctx, req)
	if err != nil {
		require.Nil(setup.T, resp)

		return nil, err
	}

	require.NotNil(setup.T, resp)

	return resp.NocIcaCertificates, nil
}

func QueryNocRevokedRootCertificates(
	setup *TestSetup,
	subject string,
	subjectKeyID string,
) (*types.RevokedNocRootCertificates, error) {
	req := &types.QueryGetRevokedNocRootCertificatesRequest{
		Subject:      subject,
		SubjectKeyId: subjectKeyID,
	}

	resp, err := setup.Keeper.RevokedNocRootCertificates(setup.Wctx, req)
	if err != nil {
		require.Nil(setup.T, resp)

		return nil, err
	}

	require.NotNil(setup.T, resp)

	return &resp.RevokedNocRootCertificates, nil
}

func QueryNocRevokedIcaCertificates(
	setup *TestSetup,
	subject string,
	subjectKeyID string,
) (*types.RevokedNocIcaCertificates, error) {
	req := &types.QueryGetRevokedNocIcaCertificatesRequest{
		Subject:      subject,
		SubjectKeyId: subjectKeyID,
	}

	resp, err := setup.Keeper.RevokedNocIcaCertificates(setup.Wctx, req)
	if err != nil {
		require.Nil(setup.T, resp)

		return nil, err
	}

	require.NotNil(setup.T, resp)

	return &resp.RevokedNocIcaCertificates, nil
}

func QueryAllNocRevokedIcaCertificates(
	setup *TestSetup,
) ([]types.RevokedNocIcaCertificates, error) {
	req := &types.QueryAllRevokedNocIcaCertificatesRequest{}

	resp, err := setup.Keeper.RevokedNocIcaCertificatesAll(setup.Wctx, req)
	if err != nil {
		require.Nil(setup.T, resp)

		return nil, err
	}

	require.NotNil(setup.T, resp)

	return resp.RevokedNocIcaCertificates, nil
}
