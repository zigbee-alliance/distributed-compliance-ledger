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

//nolint:testpackage
package keeper

import (
	"testing"

	"github.com/stretchr/testify/require"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/internal/types"
)

func TestKeeper_ApprovedCertificateGetSet(t *testing.T) {
	setup := Setup()

	// check if certificate present
	require.False(t, setup.PkiKeeper.IsApprovedCertificatesPresent(
		setup.Ctx, testconstants.LeafSubject, testconstants.LeafSubjectKeyID))

	// no certificate before its created
	certificates := setup.PkiKeeper.GetApprovedCertificates(
		setup.Ctx, testconstants.LeafSubject, testconstants.LeafSubjectKeyID)
	require.Equal(t, 0, len(certificates.Items))

	// store certificate
	certificate := types.NewNonRootCertificate(
		testconstants.LeafCertPem,
		testconstants.LeafSubject,
		testconstants.LeafSubjectKeyID,
		testconstants.LeafSerialNumber,
		testconstants.LeafIssuer,
		testconstants.LeafAuthorityKeyID,
		testconstants.RootSubject,
		testconstants.RootSubjectKeyID,
		testconstants.Address1,
	)

	setup.PkiKeeper.AddApprovedCertificate(setup.Ctx, certificate)

	// check if certificate present
	require.True(t, setup.PkiKeeper.IsApprovedCertificatesPresent(
		setup.Ctx, testconstants.LeafSubject, testconstants.LeafSubjectKeyID))

	// get certificate
	receivedCertificates := setup.PkiKeeper.GetApprovedCertificates(
		setup.Ctx, testconstants.LeafSubject, certificate.SubjectKeyID)
	require.Equal(t, 1, len(receivedCertificates.Items))

	receivedCertificate := receivedCertificates.Items[0]
	require.Equal(t, certificate.PemCert, receivedCertificate.PemCert)
	require.Equal(t, certificate.Subject, receivedCertificate.Subject)
	require.Equal(t, certificate.SubjectKeyID, receivedCertificate.SubjectKeyID)
	require.Equal(t, certificate.SerialNumber, receivedCertificate.SerialNumber)
	require.Equal(t, certificate.Issuer, receivedCertificate.Issuer)
	require.Equal(t, certificate.AuthorityKeyID, receivedCertificate.AuthorityKeyID)
	require.Equal(t, certificate.RootSubject, receivedCertificate.RootSubject)
	require.Equal(t, certificate.RootSubjectKeyID, receivedCertificate.RootSubjectKeyID)
	require.Equal(t, certificate.IsRoot, receivedCertificate.IsRoot)
	require.Equal(t, certificate.Owner, receivedCertificate.Owner)
}

// nolint:wsl
func TestKeeper_ProposedCertificateGetSet(t *testing.T) {
	setup := Setup()

	// check if proposed certificate present
	require.False(t, setup.PkiKeeper.IsProposedCertificatePresent(
		setup.Ctx, testconstants.RootSubject, testconstants.RootSubjectKeyID))

	// no certificate before its created
	require.Panics(t, func() {
		setup.PkiKeeper.GetProposedCertificate(setup.Ctx, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	})

	// store certificate
	certificate := types.NewProposedCertificate(
		testconstants.RootCertPem,
		testconstants.RootSubject,
		testconstants.RootSubjectKeyID,
		testconstants.RootSerialNumber,
		testconstants.Address1,
	)

	setup.PkiKeeper.SetProposedCertificate(setup.Ctx, certificate)

	// check if certificate present
	require.True(t, setup.PkiKeeper.IsProposedCertificatePresent(
		setup.Ctx, testconstants.RootSubject, testconstants.RootSubjectKeyID))

	// get certificate
	receivedCertificate := setup.PkiKeeper.GetProposedCertificate(
		setup.Ctx, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.Equal(t, certificate.PemCert, receivedCertificate.PemCert)
	require.Equal(t, certificate.Subject, receivedCertificate.Subject)
	require.Equal(t, certificate.SubjectKeyID, receivedCertificate.SubjectKeyID)
	require.Equal(t, certificate.SerialNumber, receivedCertificate.SerialNumber)
	require.Equal(t, certificate.Owner, receivedCertificate.Owner)
	// Amino marshals empty slices as nulls: https://github.com/tendermint/go-amino/issues/275
	// require.Equal(t, certificate.Approvals, receivedCertificate.Approvals)
}

func TestKeeper_RevokedCertificateGetSet(t *testing.T) {
	setup := Setup()

	// check if revoked certificate present
	require.False(t, setup.PkiKeeper.IsRevokedCertificatesPresent(
		setup.Ctx, testconstants.LeafSubject, testconstants.LeafSubjectKeyID))

	// no revoked certificate before its created
	certificates := setup.PkiKeeper.GetRevokedCertificates(
		setup.Ctx, testconstants.LeafSubject, testconstants.LeafSubjectKeyID)
	require.Equal(t, 0, len(certificates.Items))

	// store revoked certificate
	certificate := types.NewNonRootCertificate(
		testconstants.LeafCertPem,
		testconstants.LeafSubject,
		testconstants.LeafSubjectKeyID,
		testconstants.LeafSerialNumber,
		testconstants.LeafIssuer,
		testconstants.LeafAuthorityKeyID,
		testconstants.RootSubject,
		testconstants.RootSubjectKeyID,
		testconstants.Address1,
	)

	setup.PkiKeeper.AddRevokedCertificates(setup.Ctx, certificate.Subject, certificate.SubjectKeyID,
		types.NewCertificates([]types.Certificate{certificate}))

	// check if revoked certificate present
	require.True(t, setup.PkiKeeper.IsRevokedCertificatesPresent(
		setup.Ctx, testconstants.LeafSubject, testconstants.LeafSubjectKeyID))

	// get certificate
	receivedCertificates := setup.PkiKeeper.GetRevokedCertificates(
		setup.Ctx, testconstants.LeafSubject, certificate.SubjectKeyID)
	require.Equal(t, 1, len(receivedCertificates.Items))

	receivedCertificate := receivedCertificates.Items[0]
	require.Equal(t, certificate.PemCert, receivedCertificate.PemCert)
	require.Equal(t, certificate.Subject, receivedCertificate.Subject)
	require.Equal(t, certificate.SubjectKeyID, receivedCertificate.SubjectKeyID)
	require.Equal(t, certificate.SerialNumber, receivedCertificate.SerialNumber)
	require.Equal(t, certificate.Issuer, receivedCertificate.Issuer)
	require.Equal(t, certificate.AuthorityKeyID, receivedCertificate.AuthorityKeyID)
	require.Equal(t, certificate.RootSubject, receivedCertificate.RootSubject)
	require.Equal(t, certificate.RootSubjectKeyID, receivedCertificate.RootSubjectKeyID)
	require.Equal(t, certificate.IsRoot, receivedCertificate.IsRoot)
	require.Equal(t, certificate.Owner, receivedCertificate.Owner)
}

func TestKeeper_ProposedCertificateRevocationGetSet(t *testing.T) {
	setup := Setup()

	// check if proposed certificate present
	require.False(t, setup.PkiKeeper.IsProposedCertificateRevocationPresent(
		setup.Ctx, testconstants.RootSubject, testconstants.RootSubjectKeyID))

	// no certificate before its created
	require.Panics(t, func() {
		setup.PkiKeeper.GetProposedCertificateRevocation(
			setup.Ctx, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	})

	// store certificate
	revocation := types.NewProposedCertificateRevocation(
		testconstants.RootSubject,
		testconstants.RootSubjectKeyID,
		testconstants.Address1,
	)

	setup.PkiKeeper.SetProposedCertificateRevocation(setup.Ctx, revocation)

	// check if certificate present
	require.True(t, setup.PkiKeeper.IsProposedCertificateRevocationPresent(
		setup.Ctx, testconstants.RootSubject, testconstants.RootSubjectKeyID))

	// get certificate
	receivedRevocation := setup.PkiKeeper.GetProposedCertificateRevocation(
		setup.Ctx, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.Equal(t, revocation.Subject, receivedRevocation.Subject)
	require.Equal(t, revocation.SubjectKeyID, receivedRevocation.SubjectKeyID)
	require.Equal(t, revocation.Approvals, revocation.Approvals)
}

func TestKeeper_ChildCertificatesGetSet(t *testing.T) {
	setup := Setup()

	// check if child certificates list present
	require.False(t, setup.PkiKeeper.IsChildCertificatesPresent(setup.Ctx,
		testconstants.RootSubject, testconstants.RootSubjectKeyID))

	// no child certificates before its created
	receivedChildCertificates :=
		setup.PkiKeeper.GetChildCertificates(setup.Ctx, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.Equal(t, 0, len(receivedChildCertificates.CertIdentifiers))

	// store child certificates
	childCertificates := types.NewChildCertificates(testconstants.RootSubject, testconstants.RootSubjectKeyID)
	childCertificates.CertIdentifiers = append(childCertificates.CertIdentifiers,
		types.NewCertificateIdentifier(testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID))

	setup.PkiKeeper.SetChildCertificates(setup.Ctx, childCertificates)

	// check if child certificates present
	require.True(t, setup.PkiKeeper.IsChildCertificatesPresent(setup.Ctx,
		testconstants.RootSubject, testconstants.RootSubjectKeyID))

	// get child certificates
	receivedChildCertificates =
		setup.PkiKeeper.GetChildCertificates(setup.Ctx, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.Equal(t, childCertificates.Issuer, receivedChildCertificates.Issuer)
	require.Equal(t, childCertificates.AuthorityKeyID, receivedChildCertificates.AuthorityKeyID)
	require.Equal(t, childCertificates.CertIdentifiers, receivedChildCertificates.CertIdentifiers)

	// store second child
	childCertificates.CertIdentifiers = append(childCertificates.CertIdentifiers,
		types.NewCertificateIdentifier(testconstants.LeafSubject, testconstants.LeafSubjectKeyID))

	setup.PkiKeeper.SetChildCertificates(setup.Ctx, childCertificates)

	// get child certificates
	receivedChildCertificates =
		setup.PkiKeeper.GetChildCertificates(setup.Ctx, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.Equal(t, childCertificates.Issuer, receivedChildCertificates.Issuer)
	require.Equal(t, childCertificates.AuthorityKeyID, receivedChildCertificates.AuthorityKeyID)
	require.Equal(t, childCertificates.CertIdentifiers, receivedChildCertificates.CertIdentifiers)
}

func TestKeeper_UniqueCertificateKeyGetSet(t *testing.T) {
	setup := Setup()

	// check if unique certificate key is busy
	require.False(t, setup.PkiKeeper.IsUniqueCertificateKeyPresent(setup.Ctx,
		testconstants.IntermediateIssuer, testconstants.IntermediateSerialNumber))

	// register unique certificate key
	setup.PkiKeeper.SetUniqueCertificateKey(setup.Ctx,
		testconstants.IntermediateIssuer, testconstants.IntermediateSerialNumber)

	// check if unique certificate key is busy
	require.True(t, setup.PkiKeeper.IsUniqueCertificateKeyPresent(setup.Ctx,
		testconstants.IntermediateIssuer, testconstants.IntermediateSerialNumber))
}

// nolint:dupl
func TestKeeper_ApprovedCertificatesIterator(t *testing.T) {
	setup := Setup()

	genCerts := setup.PopulateStoreWithMixedCertificates()

	// get iterator
	var iteratedApprovedCerts []types.Certificate

	setup.PkiKeeper.IterateApprovedCertificatesRecords(setup.Ctx, "", func(certificates types.Certificates) (stop bool) {
		iteratedApprovedCerts = append(iteratedApprovedCerts, certificates.Items...)

		return false
	})

	allApproved := CombineCertLists(genCerts.ApprovedRoots, genCerts.ApprovedNonRoots)

	require.Equal(t, len(allApproved), len(iteratedApprovedCerts))

	for i := 0; i < len(allApproved); i++ {
		require.Equal(t, allApproved[i].PemCert, iteratedApprovedCerts[i].PemCert)
		require.Equal(t, allApproved[i].Subject, iteratedApprovedCerts[i].Subject)
		require.Equal(t, allApproved[i].SubjectKeyID, iteratedApprovedCerts[i].SubjectKeyID)
		require.Equal(t, allApproved[i].SerialNumber, iteratedApprovedCerts[i].SerialNumber)
		require.Equal(t, allApproved[i].Issuer, iteratedApprovedCerts[i].Issuer)
		require.Equal(t, allApproved[i].AuthorityKeyID, iteratedApprovedCerts[i].AuthorityKeyID)
		require.Equal(t, allApproved[i].RootSubject, iteratedApprovedCerts[i].RootSubject)
		require.Equal(t, allApproved[i].RootSubjectKeyID, iteratedApprovedCerts[i].RootSubjectKeyID)
		require.Equal(t, allApproved[i].IsRoot, iteratedApprovedCerts[i].IsRoot)
		require.Equal(t, allApproved[i].Owner, iteratedApprovedCerts[i].Owner)
	}
}

func TestKeeper_ProposedCertificateIterator(t *testing.T) {
	setup := Setup()

	genCerts := setup.PopulateStoreWithMixedCertificates()

	// get iterator
	var iteratedProposedCerts []types.ProposedCertificate

	setup.PkiKeeper.IterateProposedCertificates(setup.Ctx,
		func(proposedCertificate types.ProposedCertificate) (stop bool) {
			iteratedProposedCerts = append(iteratedProposedCerts, proposedCertificate)

			return false
		})

	require.Equal(t, len(genCerts.ProposedRoots), len(iteratedProposedCerts))

	// nolint:wsl
	for i := 0; i < len(genCerts.ProposedRoots); i++ {
		require.Equal(t, genCerts.ProposedRoots[i].PemCert, iteratedProposedCerts[i].PemCert)
		require.Equal(t, genCerts.ProposedRoots[i].Subject, iteratedProposedCerts[i].Subject)
		require.Equal(t, genCerts.ProposedRoots[i].SubjectKeyID, iteratedProposedCerts[i].SubjectKeyID)
		require.Equal(t, genCerts.ProposedRoots[i].SerialNumber, iteratedProposedCerts[i].SerialNumber)
		require.Equal(t, genCerts.ProposedRoots[i].Owner, iteratedProposedCerts[i].Owner)
		// Amino marshals empty slices as nulls: https://github.com/tendermint/go-amino/issues/275
		// require.Equal(t, genCerts.ProposedRoots[i].Approvals, iteratedProposedCerts[i].Approvals)
	}
}

// nolint:dupl
func TestKeeper_RevokedCertificatesIterator(t *testing.T) {
	setup := Setup()

	genCerts := setup.PopulateStoreWithMixedCertificates()

	// get iterator
	var iteratedRevokedCerts []types.Certificate

	setup.PkiKeeper.IterateRevokedCertificatesRecords(setup.Ctx, "", func(certificates types.Certificates) (stop bool) {
		iteratedRevokedCerts = append(iteratedRevokedCerts, certificates.Items...)

		return false
	})

	allRevoked := CombineCertLists(genCerts.RevokedRoots, genCerts.RevokedNonRoots)

	require.Equal(t, len(allRevoked), len(iteratedRevokedCerts))

	for i := 0; i < len(allRevoked); i++ {
		require.Equal(t, allRevoked[i].PemCert, iteratedRevokedCerts[i].PemCert)
		require.Equal(t, allRevoked[i].Subject, iteratedRevokedCerts[i].Subject)
		require.Equal(t, allRevoked[i].SubjectKeyID, iteratedRevokedCerts[i].SubjectKeyID)
		require.Equal(t, allRevoked[i].SerialNumber, iteratedRevokedCerts[i].SerialNumber)
		require.Equal(t, allRevoked[i].Issuer, iteratedRevokedCerts[i].Issuer)
		require.Equal(t, allRevoked[i].AuthorityKeyID, iteratedRevokedCerts[i].AuthorityKeyID)
		require.Equal(t, allRevoked[i].RootSubject, iteratedRevokedCerts[i].RootSubject)
		require.Equal(t, allRevoked[i].RootSubjectKeyID, iteratedRevokedCerts[i].RootSubjectKeyID)
		require.Equal(t, allRevoked[i].IsRoot, iteratedRevokedCerts[i].IsRoot)
		require.Equal(t, allRevoked[i].Owner, iteratedRevokedCerts[i].Owner)
	}
}

func TestKeeper_ProposedCertificateRevocationIterator(t *testing.T) {
	setup := Setup()

	genCerts := setup.PopulateStoreWithMixedCertificates()

	// get iterator
	var iteratedProposedRevocations []types.ProposedCertificateRevocation

	setup.PkiKeeper.IterateProposedCertificateRevocations(
		setup.Ctx,
		func(revocation types.ProposedCertificateRevocation) (stop bool) {
			iteratedProposedRevocations = append(iteratedProposedRevocations, revocation)

			return false
		},
	)

	require.Equal(t, len(genCerts.ProposedRootRevocations), len(iteratedProposedRevocations))

	for i := 0; i < len(genCerts.ProposedRootRevocations); i++ {
		require.Equal(t, genCerts.ProposedRootRevocations[i].Subject, iteratedProposedRevocations[i].Subject)
		require.Equal(t, genCerts.ProposedRootRevocations[i].SubjectKeyID, iteratedProposedRevocations[i].SubjectKeyID)
		require.Equal(t, genCerts.ProposedRootRevocations[i].Approvals, iteratedProposedRevocations[i].Approvals)
	}
}

func TestKeeper_ChildCertificatesIterator(t *testing.T) {
	setup := Setup()

	// store child certificates
	childCertificates1 := types.NewChildCertificates(
		"CN=DST Root CA X3,O=Digital Signature Trust Co.",
		testconstants.IntermediateAuthorityKeyID)
	childCertificates1.CertIdentifiers = append(childCertificates1.CertIdentifiers,
		types.NewCertificateIdentifier(testconstants.IntermediateSubject, testconstants.IntermediateSubjectKeyID))

	setup.PkiKeeper.SetChildCertificates(setup.Ctx, childCertificates1)

	childCertificates2 := types.NewChildCertificates(
		"CN=Let's Encrypt Authority X3,O=Let's Encrypt,C=US",
		testconstants.LeafAuthorityKeyID)
	childCertificates2.CertIdentifiers = append(childCertificates2.CertIdentifiers,
		types.NewCertificateIdentifier(testconstants.LeafSubject, testconstants.LeafSubjectKeyID))

	setup.PkiKeeper.SetChildCertificates(setup.Ctx, childCertificates2)

	// get iterator
	var iteratedChildCertificatesRecords []types.ChildCertificates

	setup.PkiKeeper.IterateChildCertificatesRecords(
		setup.Ctx,
		func(childCertificates types.ChildCertificates) (stop bool) {
			iteratedChildCertificatesRecords = append(iteratedChildCertificatesRecords, childCertificates)

			return false
		},
	)

	require.Equal(t, 2, len(iteratedChildCertificatesRecords))
	require.Equal(t, childCertificates1, iteratedChildCertificatesRecords[0])
	require.Equal(t, childCertificates2, iteratedChildCertificatesRecords[1])
}
