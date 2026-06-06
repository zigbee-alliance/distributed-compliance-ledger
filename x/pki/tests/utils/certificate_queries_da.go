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
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func QueryProposedCertificate(
	setup *TestSetup,
	subject string,
	subjectKeyID string,
) (*types.ProposedCertificate, error) {
	req := &types.QueryGetProposedCertificateRequest{
		Subject:      subject,
		SubjectKeyId: subjectKeyID,
	}

	resp, err := setup.Keeper.ProposedCertificate(setup.Wctx, req)
	if err != nil {
		require.Nil(setup.T, resp)

		return nil, err
	}

	require.NotNil(setup.T, resp)

	return &resp.ProposedCertificate, nil
}

func QueryApprovedCertificates(
	setup *TestSetup,
	subject string,
	subjectKeyID string,
) (*types.ApprovedCertificates, error) {
	req := &types.QueryGetApprovedCertificatesRequest{
		Subject:      subject,
		SubjectKeyId: subjectKeyID,
	}

	resp, err := setup.Keeper.ApprovedCertificates(setup.Wctx, req)
	if err != nil {
		require.Nil(setup.T, resp)

		return nil, err
	}

	require.NotNil(setup.T, resp)

	return &resp.ApprovedCertificates, nil
}

func QueryApprovedCertificatesBySubject(
	setup *TestSetup,
	subject string,
) (*types.ApprovedCertificatesBySubject, error) {
	req := &types.QueryGetApprovedCertificatesBySubjectRequest{
		Subject: subject,
	}

	resp, err := setup.Keeper.ApprovedCertificatesBySubject(setup.Wctx, req)
	if err != nil {
		require.Nil(setup.T, resp)

		return nil, err
	}

	require.NotNil(setup.T, resp)

	return &resp.ApprovedCertificatesBySubject, nil
}

func QueryApprovedCertificatesBySubjectKeyID(
	setup *TestSetup,
	subjectKeyID string,
) ([]types.ApprovedCertificates, error) {
	req := &types.QueryAllApprovedCertificatesRequest{
		SubjectKeyId: subjectKeyID,
	}

	resp, err := setup.Keeper.ApprovedCertificatesAll(setup.Wctx, req)
	if err != nil {
		require.Nil(setup.T, resp)

		return nil, err
	}

	require.NotNil(setup.T, resp)

	return resp.ApprovedCertificates, nil
}

func QueryApprovedRootCertificates(
	setup *TestSetup,
	subject string,
	subjectKeyID string,
) (*types.CertificateIdentifier, error) {
	req := &types.QueryGetApprovedRootCertificatesRequest{}

	resp, err := setup.Keeper.ApprovedRootCertificates(setup.Wctx, req)
	if err != nil {
		require.Nil(setup.T, resp)

		return nil, err
	}

	require.NotNil(setup.T, resp)

	for _, cert := range resp.ApprovedRootCertificates.Certs {
		if cert.Subject == subject && cert.SubjectKeyId == subjectKeyID {
			return cert, nil
		}
	}

	return nil, status.Error(codes.NotFound, "not found")
}

func QueryProposedCertificateRevocation(
	setup *TestSetup,
	subject string,
	subjectKeyID string,
	serialNumber string,
) (*types.ProposedCertificateRevocation, error) {
	// query proposed certificate revocation
	req := &types.QueryGetProposedCertificateRevocationRequest{
		Subject:      subject,
		SubjectKeyId: subjectKeyID,
		SerialNumber: serialNumber,
	}

	resp, err := setup.Keeper.ProposedCertificateRevocation(setup.Wctx, req)
	if err != nil {
		require.Nil(setup.T, resp)

		return nil, err
	}

	require.NotNil(setup.T, resp)

	return &resp.ProposedCertificateRevocation, nil
}

func QueryRevokedCertificates(
	setup *TestSetup,
	subject string,
	subjectKeyID string,
) (*types.RevokedCertificates, error) {
	req := &types.QueryGetRevokedCertificatesRequest{
		Subject:      subject,
		SubjectKeyId: subjectKeyID,
	}

	resp, err := setup.Keeper.RevokedCertificates(setup.Wctx, req)
	if err != nil {
		require.Nil(setup.T, resp)

		return nil, err
	}

	require.NotNil(setup.T, resp)

	return &resp.RevokedCertificates, nil
}

func QueryRevokedRootCertificates(setup *TestSetup) (*types.RevokedRootCertificates, error) {
	req := &types.QueryGetRevokedRootCertificatesRequest{}

	resp, err := setup.Keeper.RevokedRootCertificates(setup.Wctx, req)
	if err != nil {
		require.Nil(setup.T, resp)

		return nil, err
	}

	require.NotNil(setup.T, resp)

	return &resp.RevokedRootCertificates, nil
}

func QueryRejectedCertificates(
	setup *TestSetup,
	subject string,
	subjectKeyID string,
) (*types.RejectedCertificate, error) {
	req := &types.QueryGetRejectedCertificatesRequest{
		Subject:      subject,
		SubjectKeyId: subjectKeyID,
	}

	resp, err := setup.Keeper.RejectedCertificate(setup.Wctx, req)
	if err != nil {
		require.Nil(setup.T, resp)

		return nil, err
	}

	require.NotNil(setup.T, resp)

	return &resp.RejectedCertificate, nil
}

func QueryAllProposedCertificates(
	setup *TestSetup,
) ([]types.ProposedCertificate, error) {
	req := &types.QueryAllProposedCertificateRequest{}

	resp, err := setup.Keeper.ProposedCertificateAll(setup.Wctx, req)
	if err != nil {
		require.Nil(setup.T, resp)

		return nil, err
	}

	require.NotNil(setup.T, resp)

	return resp.ProposedCertificate, nil
}

func QueryAllApprovedCertificates(
	setup *TestSetup,
) ([]types.ApprovedCertificates, error) {
	req := &types.QueryAllApprovedCertificatesRequest{}

	resp, err := setup.Keeper.ApprovedCertificatesAll(setup.Wctx, req)
	if err != nil {
		require.Nil(setup.T, resp)

		return nil, err
	}

	require.NotNil(setup.T, resp)

	return resp.ApprovedCertificates, nil
}

func QueryAllRevokedCertificates(
	setup *TestSetup,
) ([]types.RevokedCertificates, error) {
	req := &types.QueryAllRevokedCertificatesRequest{}

	resp, err := setup.Keeper.RevokedCertificatesAll(setup.Wctx, req)
	if err != nil {
		require.Nil(setup.T, resp)

		return nil, err
	}

	require.NotNil(setup.T, resp)

	return resp.RevokedCertificates, nil
}

func QueryAllProposedCertificateRevocations(
	setup *TestSetup,
) ([]types.ProposedCertificateRevocation, error) {
	req := &types.QueryAllProposedCertificateRevocationRequest{}

	resp, err := setup.Keeper.ProposedCertificateRevocationAll(setup.Wctx, req)
	if err != nil {
		require.Nil(setup.T, resp)

		return nil, err
	}

	require.NotNil(setup.T, resp)

	return resp.ProposedCertificateRevocation, nil
}

func IsRevokedRootCertificatePresent(
	setup *TestSetup,
	subject string,
	subjectKeyID string,
) bool {
	req := &types.QueryGetRevokedRootCertificatesRequest{}

	resp, err := setup.Keeper.RevokedRootCertificates(setup.Wctx, req)
	if err != nil {
		require.Nil(setup.T, resp)

		return false
	}

	require.NotNil(setup.T, resp)

	for _, cert := range resp.RevokedRootCertificates.Certs {
		if cert.Subject == subject && cert.SubjectKeyId == subjectKeyID {
			return true
		}
	}

	return false
}
