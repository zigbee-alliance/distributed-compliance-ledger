package utils

import (
	"github.com/stretchr/testify/require"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

func QueryAllCertificatesAll(
	setup *TestSetup,
) ([]types.AllCertificates, error) {
	req := &types.QueryAllCertificatesRequest{}

	resp, err := setup.Keeper.CertificatesAll(setup.Wctx, req)
	if err != nil {
		require.Nil(setup.T, resp)

		return nil, err
	}

	require.NotNil(setup.T, resp)

	return resp.Certificates, nil
}

func QueryAllCertificates(
	setup *TestSetup,
	subject string,
	subjectKeyID string,
) (*types.AllCertificates, error) {
	req := &types.QueryGetCertificatesRequest{
		Subject:      subject,
		SubjectKeyId: subjectKeyID,
	}

	resp, err := setup.Keeper.Certificates(setup.Wctx, req)
	if err != nil {
		require.Nil(setup.T, resp)

		return nil, err
	}

	require.NotNil(setup.T, resp)

	return &resp.Certificates, nil
}

func QueryAllCertificatesBySubject(
	setup *TestSetup,
	subject string,
) (*types.AllCertificatesBySubject, error) {
	req := &types.QueryGetAllCertificatesBySubjectRequest{
		Subject: subject,
	}

	resp, err := setup.Keeper.AllCertificatesBySubject(setup.Wctx, req)
	if err != nil {
		require.Nil(setup.T, resp)

		return nil, err
	}

	require.NotNil(setup.T, resp)

	return &resp.AllCertificatesBySubject, nil
}

func QueryAllCertificatesBySubjectKeyID(
	setup *TestSetup,
	subjectKeyID string,
) ([]types.AllCertificates, error) {
	req := &types.QueryAllCertificatesRequest{
		SubjectKeyId: subjectKeyID,
	}

	resp, err := setup.Keeper.CertificatesAll(setup.Wctx, req)
	if err != nil {
		require.Nil(setup.T, resp)

		return nil, err
	}

	require.NotNil(setup.T, resp)

	return resp.Certificates, nil
}

func QueryChildCertificates(
	setup *TestSetup,
	issuer string,
	authorityKeyID string,
) (*types.ChildCertificates, error) {
	req := &types.QueryGetChildCertificatesRequest{
		Issuer:         issuer,
		AuthorityKeyId: authorityKeyID,
	}

	resp, err := setup.Keeper.ChildCertificates(setup.Wctx, req)
	if err != nil {
		require.Nil(setup.T, resp)

		return nil, err
	}

	require.NotNil(setup.T, resp)

	return &resp.ChildCertificates, nil
}
