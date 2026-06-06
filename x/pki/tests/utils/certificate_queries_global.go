// Copyright 2020 DSR Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
