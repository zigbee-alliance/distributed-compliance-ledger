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
package pki

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	constants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/pagination"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/auth"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/internal/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/internal/types"
)

const SerialNumber = "12345678"

func TestHandler_ProposeAddX509RootCert_ByNotTrustee(t *testing.T) {
	setup := Setup()

	for _, role := range []auth.AccountRole{auth.TestHouse, auth.ZBCertificationCenter, auth.Vendor} {
		// assign role
		account := auth.NewAccount(constants.Address1, constants.PubKey1, auth.AccountRoles{role}, constants.VendorId1)
		setup.AuthKeeper.SetAccount(setup.Ctx, account)

		// propose x509 root certificate
		proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(constants.RootCertPem, constants.Address1)
		result := setup.Handler(setup.Ctx, proposeAddX509RootCert)
		require.Equal(t, sdk.CodeOK, result.Code)

		// query proposed certificate
		proposedCertificate, _ := queryProposedCertificate(&setup, constants.RootSubject, constants.RootSubjectKeyID)

		// check proposed certificate
		require.Equal(t, proposeAddX509RootCert.Cert, proposedCertificate.PemCert)
		require.Equal(t, proposeAddX509RootCert.Signer, proposedCertificate.Owner)
		require.Equal(t, constants.RootSubject, proposedCertificate.Subject)
		require.Equal(t, constants.RootSubjectKeyID, proposedCertificate.SubjectKeyID)
		require.Equal(t, constants.RootSerialNumber, proposedCertificate.SerialNumber)
		require.Nil(t, proposedCertificate.Approvals)

		// check that unique certificate key is registered
		require.True(t, setup.PkiKeeper.IsUniqueCertificateKeyPresent(setup.Ctx,
			constants.RootIssuer, constants.RootSerialNumber))

		// try to query approved certificate
		_, err := querySingleApprovedCertificate(&setup, constants.RootSubject, constants.RootSubjectKeyID)
		require.Equal(t, types.CodeCertificateDoesNotExist, err.Code())

		// cleanup for next iteration
		setup.PkiKeeper.DeleteProposedCertificate(setup.Ctx, constants.RootSubject, constants.RootSubjectKeyID)
		setup.PkiKeeper.DeleteUniqueCertificateKey(setup.Ctx, constants.RootIssuer, constants.RootSerialNumber)
	}
}

func TestHandler_ProposeAddX509RootCert_ByTrustee(t *testing.T) {
	setup := Setup()

	// propose x509 root certificate
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(constants.RootCertPem, setup.Trustee)
	result := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.Equal(t, sdk.CodeOK, result.Code)

	// query proposed certificate
	proposedCertificate, _ := queryProposedCertificate(&setup, constants.RootSubject, constants.RootSubjectKeyID)

	// check proposed certificate
	require.Equal(t, proposeAddX509RootCert.Cert, proposedCertificate.PemCert)
	require.Equal(t, proposeAddX509RootCert.Signer, proposedCertificate.Owner)
	require.Equal(t, constants.RootSubject, proposedCertificate.Subject)
	require.Equal(t, constants.RootSubjectKeyID, proposedCertificate.SubjectKeyID)
	require.Equal(t, constants.RootSerialNumber, proposedCertificate.SerialNumber)
	require.Equal(t, []sdk.AccAddress{proposeAddX509RootCert.Signer}, proposedCertificate.Approvals)

	// check that unique certificate key is registered
	require.True(t, setup.PkiKeeper.IsUniqueCertificateKeyPresent(setup.Ctx,
		constants.RootIssuer, constants.RootSerialNumber))

	// query approved certificate
	_, err := querySingleApprovedCertificate(&setup, constants.RootSubject, constants.RootSubjectKeyID)
	require.Equal(t, types.CodeCertificateDoesNotExist, err.Code())
}

func TestHandler_ProposeAddX509RootCert_ForInvalidCertificate(t *testing.T) {
	setup := Setup()

	// propose x509 root certificate
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(constants.StubCertPem, setup.Trustee)
	result := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.Equal(t, types.CodeInvalidCertificate, result.Code)
}

func TestHandler_ProposeAddX509RootCert_ForNonRootCertificate(t *testing.T) {
	setup := Setup()

	// propose x509 leaf certificate as root
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(constants.LeafCertPem, setup.Trustee)
	result := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.Equal(t, types.CodeInappropriateCertificateType, result.Code)
}

func TestHandler_ProposeAddX509RootCert_ProposedCertificateAlreadyExists(t *testing.T) {
	setup := Setup()

	// propose adding of x509 root certificate
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(constants.RootCertPem, setup.Trustee)
	result := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.Equal(t, sdk.CodeOK, result.Code)

	// store another account
	account := auth.NewAccount(constants.Address1, constants.PubKey1, auth.AccountRoles{}, 0)
	setup.AuthKeeper.SetAccount(setup.Ctx, account)

	// propose adding of the same x509 root certificate again
	proposeAddX509RootCert = types.NewMsgProposeAddX509RootCert(constants.RootCertPem, constants.Address1)
	result = setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.Equal(t, types.CodeProposedCertificateAlreadyExists, result.Code)
}

func TestHandler_ProposeAddX509RootCert_CertificateAlreadyExists(t *testing.T) {
	setup := Setup()

	// store x509 root certificate
	rootCertificate := rootCertificate(constants.Address1)
	setup.PkiKeeper.SetUniqueCertificateKey(setup.Ctx, rootCertificate.Subject, rootCertificate.SerialNumber)
	setup.PkiKeeper.AddApprovedCertificate(setup.Ctx, rootCertificate)

	// propose adding of the same x509 root certificate
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(constants.RootCertPem, setup.Trustee)
	result := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.Equal(t, types.CodeCertificateAlreadyExists, result.Code)
}

func TestHandler_ProposeAddX509RootCert_ForDifferentSerialNumber(t *testing.T) {
	setup := Setup()

	// store root certificate with different serial number
	rootCertificate := rootCertificate(setup.Trustee)
	rootCertificate.SerialNumber = SerialNumber
	setup.PkiKeeper.SetUniqueCertificateKey(setup.Ctx, rootCertificate.Subject, rootCertificate.SerialNumber)
	setup.PkiKeeper.AddApprovedCertificate(setup.Ctx, rootCertificate)

	// propose second root certificate
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(constants.RootCertPem, setup.Trustee)
	result := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.Equal(t, sdk.CodeOK, result.Code)

	// check
	certificate, _ := querySingleApprovedCertificate(&setup, constants.RootSubject, constants.RootSubjectKeyID)
	require.True(t, certificate.IsRoot)
	require.Equal(t, constants.RootIssuer, certificate.Subject)
	require.Equal(t, SerialNumber, certificate.SerialNumber)

	proposedCertificate, _ := queryProposedCertificate(&setup, constants.RootSubject, constants.RootSubjectKeyID)
	require.Equal(t, constants.RootIssuer, proposedCertificate.Subject)
	require.Equal(t, constants.RootSerialNumber, proposedCertificate.SerialNumber)

	require.NotEqual(t, certificate.SerialNumber, proposedCertificate.SerialNumber)
}

func TestHandler_ProposeAddX509RootCert_ForDifferentSerialNumberDifferentSigner(t *testing.T) {
	setup := Setup()

	// store root certificate with different serial number
	rootCertificate := rootCertificate(constants.Address1)
	rootCertificate.SerialNumber = SerialNumber
	setup.PkiKeeper.SetUniqueCertificateKey(setup.Ctx, rootCertificate.Subject, rootCertificate.SerialNumber)
	setup.PkiKeeper.AddApprovedCertificate(setup.Ctx, rootCertificate)

	// propose second root certificate
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(constants.RootCertPem, setup.Trustee)
	result := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.Equal(t, sdk.CodeUnauthorized, result.Code)
}

func TestHandler_ApproveAddX509RootCert_ForNotEnoughApprovals(t *testing.T) {
	setup := Setup()

	// store account
	account := auth.NewAccount(constants.Address1, constants.PubKey1, auth.AccountRoles{}, 0)
	setup.AuthKeeper.SetAccount(setup.Ctx, account)

	// propose x509 root certificate by account without trustee role
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(constants.RootCertPem, constants.Address1)
	result := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.Equal(t, sdk.CodeOK, result.Code)

	// approve
	approveAddX509RootCert := types.NewMsgApproveAddX509RootCert(
		constants.RootSubject, constants.RootSubjectKeyID, setup.Trustee)
	result = setup.Handler(setup.Ctx, approveAddX509RootCert)
	require.Equal(t, sdk.CodeOK, result.Code)

	// query certificate
	proposedCertificate, _ := queryProposedCertificate(&setup, constants.RootSubject, constants.RootSubjectKeyID)
	require.Equal(t, proposeAddX509RootCert.Cert, proposedCertificate.PemCert)
	require.Equal(t, []sdk.AccAddress{setup.Trustee}, proposedCertificate.Approvals)

	// query approved certificate
	_, err := querySingleApprovedCertificate(&setup, constants.RootSubject, constants.RootSubjectKeyID)
	require.Equal(t, types.CodeCertificateDoesNotExist, err.Code())
}

func TestHandler_ApproveAddX509RootCert_ForEnoughApprovals(t *testing.T) {
	setup := Setup()

	// propose add x509 root certificate by trustee
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(constants.RootCertPem, setup.Trustee)
	result := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.Equal(t, sdk.CodeOK, result.Code)

	// store second trustee
	account := auth.NewAccount(constants.Address1, constants.PubKey1, auth.AccountRoles{auth.Trustee}, 0)
	setup.AuthKeeper.SetAccount(setup.Ctx, account)

	// approve by second trustee
	approveAddX509RootCert := types.NewMsgApproveAddX509RootCert(
		constants.RootSubject, constants.RootSubjectKeyID, constants.Address1)
	result = setup.Handler(setup.Ctx, approveAddX509RootCert)
	require.Equal(t, sdk.CodeOK, result.Code)

	// query proposed certificate
	_, err := queryProposedCertificate(&setup, constants.RootSubject, constants.RootSubjectKeyID)
	require.Equal(t, types.CodeProposedCertificateDoesNotExist, err.Code())

	// query approved certificate
	approvedCertificate, _ := querySingleApprovedCertificate(&setup, constants.RootSubject, constants.RootSubjectKeyID)
	require.Equal(t, proposeAddX509RootCert.Cert, approvedCertificate.PemCert)
	require.Equal(t, proposeAddX509RootCert.Signer, approvedCertificate.Owner)
	require.Equal(t, constants.RootSubject, approvedCertificate.Subject)
	require.Equal(t, constants.RootSubjectKeyID, approvedCertificate.SubjectKeyID)
	require.Equal(t, constants.RootSerialNumber, approvedCertificate.SerialNumber)
	require.True(t, approvedCertificate.IsRoot)
	require.Empty(t, approvedCertificate.RootSubject)
	require.Empty(t, approvedCertificate.RootSubjectKeyID)
	require.Empty(t, approvedCertificate.Issuer)
	require.Empty(t, approvedCertificate.AuthorityKeyID)

	// check that unique certificate key is registered
	require.True(t, setup.PkiKeeper.IsUniqueCertificateKeyPresent(setup.Ctx,
		constants.RootIssuer, constants.RootSerialNumber))
}

func TestHandler_ApproveAddX509RootCert_ForUnknownProposedCertificate(t *testing.T) {
	setup := Setup()

	// approve
	approveAddX509RootCert := types.NewMsgApproveAddX509RootCert(
		constants.RootSubject, constants.RootSubjectKeyID, setup.Trustee)
	result := setup.Handler(setup.Ctx, approveAddX509RootCert)
	require.Equal(t, types.CodeProposedCertificateDoesNotExist, result.Code)
}

func TestHandler_ApproveAddX509RootCert_ByNotTrustee(t *testing.T) {
	setup := Setup()

	// propose add x509 root certificate
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(constants.RootCertPem, setup.Trustee)
	result := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.Equal(t, sdk.CodeOK, result.Code)

	for _, role := range []auth.AccountRole{auth.TestHouse, auth.ZBCertificationCenter, auth.Vendor} {
		// assign role
		account := auth.NewAccount(constants.Address1, constants.PubKey1, auth.AccountRoles{role}, constants.VendorId1)
		setup.AuthKeeper.SetAccount(setup.Ctx, account)

		// approve
		approveAddX509RootCert := types.NewMsgApproveAddX509RootCert(
			constants.RootSubject, constants.RootSubjectKeyID, constants.Address1)
		result = setup.Handler(setup.Ctx, approveAddX509RootCert)
		require.Equal(t, sdk.CodeUnauthorized, result.Code)
	}
}

func TestHandler_ApproveAddX509RootCert_Twice(t *testing.T) {
	setup := Setup()

	// store account
	account := auth.NewAccount(constants.Address1, constants.PubKey1, auth.AccountRoles{}, constants.VendorId1)
	setup.AuthKeeper.SetAccount(setup.Ctx, account)

	// propose add x509 root certificate
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(constants.RootCertPem, constants.Address1)
	result := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.Equal(t, sdk.CodeOK, result.Code)

	// approve
	approveAddX509RootCert := types.NewMsgApproveAddX509RootCert(
		constants.RootSubject, constants.RootSubjectKeyID, setup.Trustee)
	result = setup.Handler(setup.Ctx, approveAddX509RootCert)
	require.Equal(t, sdk.CodeOK, result.Code)

	// approve second time
	result = setup.Handler(setup.Ctx, approveAddX509RootCert)
	require.Equal(t, sdk.CodeUnauthorized, result.Code)
}

func TestHandler_AddX509Cert(t *testing.T) {
	setup := Setup()

	// store root certificate
	rootCertificate := rootCertificate(setup.Trustee)
	setup.PkiKeeper.AddApprovedCertificate(setup.Ctx, rootCertificate)

	for _, role := range []auth.AccountRole{auth.Trustee, auth.TestHouse, auth.ZBCertificationCenter, auth.Vendor} {
		// assign role
		account := auth.NewAccount(constants.Address1, constants.PubKey1, auth.AccountRoles{role}, constants.VendorId1)
		setup.AuthKeeper.SetAccount(setup.Ctx, account)

		// add x509 certificate
		addX509Cert := types.NewMsgAddX509Cert(constants.IntermediateCertPem, constants.Address1)
		result := setup.Handler(setup.Ctx, addX509Cert)
		require.Equal(t, sdk.CodeOK, result.Code)

		// query certificate
		certificate, _ := querySingleApprovedCertificate(&setup,
			constants.IntermediateSubject, constants.IntermediateSubjectKeyID)

		// check
		require.Equal(t, addX509Cert.Cert, certificate.PemCert)
		require.Equal(t, addX509Cert.Signer, certificate.Owner)
		require.Equal(t, constants.IntermediateSubject, certificate.Subject)
		require.Equal(t, constants.IntermediateSubjectKeyID, certificate.SubjectKeyID)
		require.Equal(t, constants.IntermediateSerialNumber, certificate.SerialNumber)
		require.False(t, certificate.IsRoot)
		require.Equal(t, constants.IntermediateIssuer, certificate.Issuer)
		require.Equal(t, constants.IntermediateAuthorityKeyID, certificate.AuthorityKeyID)
		require.Equal(t, constants.RootSubject, certificate.RootSubject)
		require.Equal(t, constants.RootSubjectKeyID, certificate.RootSubjectKeyID)

		// check that unique certificate key is registered
		require.True(t, setup.PkiKeeper.IsUniqueCertificateKeyPresent(setup.Ctx,
			constants.IntermediateIssuer, constants.IntermediateSerialNumber))

		// check that child certificates of issuer contains certificate identifier
		issuerChildren := setup.PkiKeeper.GetChildCertificates(setup.Ctx,
			constants.IntermediateIssuer, constants.IntermediateAuthorityKeyID)

		require.Equal(t, 1, len(issuerChildren.CertIdentifiers))
		require.Equal(t,
			types.NewCertificateIdentifier(constants.IntermediateSubject, constants.IntermediateSubjectKeyID),
			issuerChildren.CertIdentifiers[0])

		// check that no proposed certificate has been created
		_, err := queryProposedCertificate(&setup, constants.IntermediateSubject, constants.IntermediateSubjectKeyID)
		require.Equal(t, types.CodeProposedCertificateDoesNotExist, err.Code())

		// cleanup for next iteration
		setup.PkiKeeper.DeleteApprovedCertificates(setup.Ctx,
			constants.IntermediateSubject, constants.IntermediateSubjectKeyID)
		setup.PkiKeeper.DeleteUniqueCertificateKey(setup.Ctx,
			constants.IntermediateIssuer, constants.IntermediateSerialNumber)
		setup.PkiKeeper.DeleteChildCertificates(setup.Ctx,
			constants.IntermediateIssuer, constants.IntermediateAuthorityKeyID)
	}
}

func TestHandler_AddX509Cert_ForInvalidCertificate(t *testing.T) {
	setup := Setup()

	// add x509 certificate
	addX509Cert := types.NewMsgAddX509Cert(constants.StubCertPem, setup.Trustee)
	result := setup.Handler(setup.Ctx, addX509Cert)
	require.Equal(t, types.CodeInvalidCertificate, result.Code)
}

func TestHandler_AddX509Cert_ForRootCertificate(t *testing.T) {
	setup := Setup()

	// add root certificate as leaf x509 certificate
	addX509Cert := types.NewMsgAddX509Cert(constants.RootCertPem, setup.Trustee)
	result := setup.Handler(setup.Ctx, addX509Cert)
	require.Equal(t, types.CodeInappropriateCertificateType, result.Code)
}

func TestHandler_AddX509Cert_ForDuplicate(t *testing.T) {
	setup := Setup()

	// store root certificate
	rootCertificate := rootCertificate(setup.Trustee)
	setup.PkiKeeper.AddApprovedCertificate(setup.Ctx, rootCertificate)

	// store intermediate certificate
	addX509Cert := types.NewMsgAddX509Cert(constants.IntermediateCertPem, setup.Trustee)
	result := setup.Handler(setup.Ctx, addX509Cert)
	require.Equal(t, sdk.CodeOK, result.Code)

	// store intermediate certificate second time
	result = setup.Handler(setup.Ctx, addX509Cert)
	require.Equal(t, types.CodeCertificateAlreadyExists, result.Code)
}

func TestHandler_AddX509Cert_ForDifferentSerialNumber(t *testing.T) {
	setup := Setup()

	// store root certificate
	rootCertificate := rootCertificate(setup.Trustee)
	setup.PkiKeeper.AddApprovedCertificate(setup.Ctx, rootCertificate)

	// store intermediate certificate with different serial number
	intermediateCertificate := intermediateCertificate(setup.Trustee)
	intermediateCertificate.SerialNumber = SerialNumber
	setup.PkiKeeper.SetUniqueCertificateKey(setup.Ctx,
		intermediateCertificate.Issuer, intermediateCertificate.SerialNumber)
	setup.PkiKeeper.AddApprovedCertificate(setup.Ctx, intermediateCertificate)

	// store intermediate certificate second time
	addX509Cert := types.NewMsgAddX509Cert(constants.IntermediateCertPem, setup.Trustee)
	result := setup.Handler(setup.Ctx, addX509Cert)
	require.Equal(t, sdk.CodeOK, result.Code)

	// query certificate
	certificates, _ := queryApprovedCertificates(&setup, constants.IntermediateSubject, constants.IntermediateSubjectKeyID)

	// check
	require.Equal(t, 2, len(certificates.Items))
	require.NotEqual(t, certificates.Items[0].SerialNumber, certificates.Items[1].SerialNumber)

	for _, certificate := range certificates.Items {
		require.Equal(t, addX509Cert.Cert, certificate.PemCert)
		require.Equal(t, addX509Cert.Signer, certificate.Owner)
		require.Equal(t, constants.IntermediateSubject, certificate.Subject)
		require.Equal(t, constants.IntermediateSubjectKeyID, certificate.SubjectKeyID)
		require.False(t, certificate.IsRoot)
		require.Equal(t, constants.RootSubject, certificate.RootSubject)
		require.Equal(t, constants.RootSubjectKeyID, certificate.RootSubjectKeyID)
		require.Equal(t, constants.IntermediateIssuer, certificate.Issuer)
		require.Equal(t, constants.IntermediateAuthorityKeyID, certificate.AuthorityKeyID)
	}
}

func TestHandler_AddX509Cert_ForDifferentSerialNumberDifferentSigner(t *testing.T) {
	setup := Setup()

	// store root certificate
	rootCertificate := rootCertificate(constants.Address1)
	setup.PkiKeeper.AddApprovedCertificate(setup.Ctx, rootCertificate)

	// store intermediate certificate with different serial number
	intermediateCertificate := intermediateCertificate(constants.Address1)
	intermediateCertificate.SerialNumber = SerialNumber
	setup.PkiKeeper.SetUniqueCertificateKey(setup.Ctx,
		intermediateCertificate.Issuer, intermediateCertificate.SerialNumber)
	setup.PkiKeeper.AddApprovedCertificate(setup.Ctx, intermediateCertificate)

	// store intermediate certificate second time
	addX509Cert := types.NewMsgAddX509Cert(constants.IntermediateCertPem, setup.Trustee)
	result := setup.Handler(setup.Ctx, addX509Cert)
	require.Equal(t, sdk.CodeUnauthorized, result.Code)
}

func TestHandler_AddX509Cert_ForAbsentDirectParentCert(t *testing.T) {
	setup := Setup()

	// add intermediate x509 certificate
	addX509Cert := types.NewMsgAddX509Cert(constants.IntermediateCertPem, setup.Trustee)
	result := setup.Handler(setup.Ctx, addX509Cert)
	require.Equal(t, types.CodeInvalidCertificate, result.Code)
}

func TestHandler_AddX509Cert_ForNoRootCert(t *testing.T) {
	setup := Setup()

	// add intermediate certificate
	intermediateCertificate := intermediateCertificate(setup.Trustee)
	setup.PkiKeeper.AddApprovedCertificate(setup.Ctx, intermediateCertificate)

	// add leaf x509 certificate
	addX509Cert := types.NewMsgAddX509Cert(constants.LeafCertPem, setup.Trustee)
	result := setup.Handler(setup.Ctx, addX509Cert)
	require.Equal(t, types.CodeInvalidCertificate, result.Code)
}

func TestHandler_AddX509Cert_ForFailedCertificateVerification(t *testing.T) {
	setup := Setup()

	// add invalid root
	invalidRootCertificate := types.NewRootCertificate(constants.StubCertPem,
		constants.RootSubject, constants.RootSubjectKeyID, constants.RootSerialNumber, setup.Trustee)
	setup.PkiKeeper.AddApprovedCertificate(setup.Ctx, invalidRootCertificate)

	// add intermediate x509 certificate
	addX509Cert := types.NewMsgAddX509Cert(constants.IntermediateCertPem, setup.Trustee)
	result := setup.Handler(setup.Ctx, addX509Cert)
	require.Equal(t, types.CodeInvalidCertificate, result.Code)
}

func TestHandler_AddX509Cert_ForTree(t *testing.T) {
	setup := Setup()

	// add root x509 certificate
	proposeAndApproveRootCertificate(t, &setup, setup.Trustee)

	// add intermediate x509 certificate
	addIntermediateX509Cert := types.NewMsgAddX509Cert(constants.IntermediateCertPem, setup.Trustee)
	result := setup.Handler(setup.Ctx, addIntermediateX509Cert)
	require.Equal(t, sdk.CodeOK, result.Code)

	// add leaf x509 certificate
	addLeafX509Cert := types.NewMsgAddX509Cert(constants.LeafCertPem, setup.Trustee)
	result = setup.Handler(setup.Ctx, addLeafX509Cert)
	require.Equal(t, sdk.CodeOK, result.Code)

	// query root certificate
	rootCertificate, _ := querySingleApprovedCertificate(&setup, constants.RootSubject, constants.RootSubjectKeyID)
	require.Equal(t, constants.RootCertPem, rootCertificate.PemCert)

	// check child certificate identifiers of root certificate
	rootCertChildren := setup.PkiKeeper.GetChildCertificates(setup.Ctx,
		constants.RootSubject, constants.RootSubjectKeyID)

	require.Equal(t, 1, len(rootCertChildren.CertIdentifiers))
	require.Equal(t,
		types.NewCertificateIdentifier(constants.IntermediateSubject, constants.IntermediateSubjectKeyID),
		rootCertChildren.CertIdentifiers[0])

	// query intermediate certificate
	intermediateCertificate, _ :=
		querySingleApprovedCertificate(&setup, constants.IntermediateSubject, constants.IntermediateSubjectKeyID)
	require.Equal(t, constants.IntermediateCertPem, intermediateCertificate.PemCert)

	// check child certificate identifiers of intermediate certificate
	intermediateCertChildren := setup.PkiKeeper.GetChildCertificates(setup.Ctx,
		constants.IntermediateSubject, constants.IntermediateSubjectKeyID)

	require.Equal(t, 1, len(intermediateCertChildren.CertIdentifiers))
	require.Equal(t,
		types.NewCertificateIdentifier(constants.LeafSubject, constants.LeafSubjectKeyID),
		intermediateCertChildren.CertIdentifiers[0])

	// query leaf certificate
	leafCertificate, _ := querySingleApprovedCertificate(&setup, constants.LeafSubject, constants.LeafSubjectKeyID)
	require.Equal(t, constants.LeafCertPem, leafCertificate.PemCert)

	// check child certificate identifiers of leaf certificate
	leafCertChildren :=
		setup.PkiKeeper.GetChildCertificates(setup.Ctx, constants.LeafSubject, constants.LeafSubjectKeyID)

	require.Equal(t, 0, len(leafCertChildren.CertIdentifiers))
}

//nolint:funlen
func TestHandler_AddX509Cert_EachChildCertRefersToTwoParentCerts(t *testing.T) {
	setup := Setup()

	// store root certificate
	rootCert := rootCertificate(setup.Trustee)

	setup.PkiKeeper.AddApprovedCertificate(setup.Ctx, rootCert)
	setup.PkiKeeper.SetUniqueCertificateKey(setup.Ctx, rootCert.Subject, rootCert.SerialNumber)

	// store second root certificate
	rootCert = rootCertificate(setup.Trustee)
	rootCert.SerialNumber = SerialNumber

	setup.PkiKeeper.AddApprovedCertificate(setup.Ctx, rootCert)
	setup.PkiKeeper.SetUniqueCertificateKey(setup.Ctx, rootCert.Subject, rootCert.SerialNumber)

	// store intermediate certificate (it refers to two parent certificates)
	intermediateCertificate := intermediateCertificate(setup.Trustee)
	intermediateCertificate.SerialNumber = SerialNumber

	setup.PkiKeeper.AddApprovedCertificate(setup.Ctx, intermediateCertificate)
	setup.PkiKeeper.SetUniqueCertificateKey(setup.Ctx,
		intermediateCertificate.Issuer, intermediateCertificate.SerialNumber)

	rootChildCertificates := setup.PkiKeeper.GetChildCertificates(setup.Ctx,
		intermediateCertificate.Issuer, intermediateCertificate.AuthorityKeyID)
	rootChildCertificates.CertIdentifiers = append(rootChildCertificates.CertIdentifiers,
		types.NewCertificateIdentifier(intermediateCertificate.Subject, intermediateCertificate.SubjectKeyID))
	setup.PkiKeeper.SetChildCertificates(setup.Ctx, rootChildCertificates)

	// store second intermediate certificate (it refers to two parent certificates)
	addX509Cert := types.NewMsgAddX509Cert(constants.IntermediateCertPem, setup.Trustee)
	result := setup.Handler(setup.Ctx, addX509Cert)
	require.Equal(t, sdk.CodeOK, result.Code)

	// store leaf certificate (it refers to two parent certificates)
	addX509Cert = types.NewMsgAddX509Cert(constants.LeafCertPem, setup.Trustee)
	result = setup.Handler(setup.Ctx, addX509Cert)
	require.Equal(t, sdk.CodeOK, result.Code)

	// query root certificate
	rootCertificates, _ := queryApprovedCertificates(&setup, constants.RootSubject, constants.RootSubjectKeyID)
	require.Equal(t, 2, len(rootCertificates.Items))

	// check child certificate identifiers of root certificate
	rootCertChildren := setup.PkiKeeper.GetChildCertificates(setup.Ctx,
		constants.RootSubject, constants.RootSubjectKeyID)

	require.Equal(t, 1, len(rootCertChildren.CertIdentifiers))
	require.Equal(t,
		types.NewCertificateIdentifier(constants.IntermediateSubject, constants.IntermediateSubjectKeyID),
		rootCertChildren.CertIdentifiers[0])

	// query intermediate certificate
	intermediateCertificates, _ := queryApprovedCertificates(&setup,
		constants.IntermediateSubject, constants.IntermediateSubjectKeyID)
	require.Equal(t, 2, len(intermediateCertificates.Items))

	// check child certificate identifiers of intermediate certificate
	intermediateCertChildren := setup.PkiKeeper.GetChildCertificates(setup.Ctx,
		constants.IntermediateSubject, constants.IntermediateSubjectKeyID)

	require.Equal(t, 1, len(intermediateCertChildren.CertIdentifiers))
	require.Equal(t,
		types.NewCertificateIdentifier(constants.LeafSubject, constants.LeafSubjectKeyID),
		intermediateCertChildren.CertIdentifiers[0])

	// query leaf certificate
	leafCertificates, _ := queryApprovedCertificates(&setup, constants.LeafSubject, constants.LeafSubjectKeyID)
	require.Equal(t, 1, len(leafCertificates.Items))

	// check child certificate identifiers of intermediate certificate
	leafCertChildren := setup.PkiKeeper.GetChildCertificates(setup.Ctx,
		constants.LeafSubject, constants.LeafSubjectKeyID)

	require.Equal(t, 0, len(leafCertChildren.CertIdentifiers))
}

func TestHandler_ProposeRevokeX509RootCert_ByTrusteeOwner(t *testing.T) {
	setup := Setup()

	// propose x509 root certificate by `setup.Trustee` and approve by another trustee
	proposeAndApproveRootCertificate(t, &setup, setup.Trustee)

	// propose revocation of x509 root certificate by `setup.Trustee`
	proposeRevokeX509RootCert := types.NewMsgProposeRevokeX509RootCert(
		constants.RootSubject, constants.RootSubjectKeyID, setup.Trustee)
	result := setup.Handler(setup.Ctx, proposeRevokeX509RootCert)
	require.Equal(t, sdk.CodeOK, result.Code)

	// query and check proposed certificate revocation
	proposedRevocation, _ := queryProposedCertificateRevocation(&setup, constants.RootSubject, constants.RootSubjectKeyID)
	require.Equal(t, constants.RootSubject, proposedRevocation.Subject)
	require.Equal(t, constants.RootSubjectKeyID, proposedRevocation.SubjectKeyID)
	require.Equal(t, []sdk.AccAddress{setup.Trustee}, proposedRevocation.Approvals)

	// check that approved certificate still exists
	certificate, _ := querySingleApprovedCertificate(&setup, constants.RootSubject, constants.RootSubjectKeyID)
	require.NotNil(t, certificate)

	// check that revoked certificate does not exist
	_, err := queryRevokedCertificates(&setup, constants.RootSubject, constants.RootSubjectKeyID)
	require.Equal(t, types.CodeRevokedCertificateDoesNotExist, err.Code())

	// check that unique certificate key stays registered
	require.True(t,
		setup.PkiKeeper.IsUniqueCertificateKeyPresent(setup.Ctx, constants.RootIssuer, constants.RootSerialNumber))
}

func TestHandler_ProposeRevokeX509RootCert_ByTrusteeNotOwner(t *testing.T) {
	setup := Setup()

	// propose x509 root certificate by `setup.Trustee` and approve by another trustee
	proposeAndApproveRootCertificate(t, &setup, setup.Trustee)

	// store new trustee
	account := auth.NewAccount(constants.Address1, constants.PubKey1, auth.AccountRoles{auth.Trustee}, constants.VendorId1)
	setup.AuthKeeper.SetAccount(setup.Ctx, account)

	// propose revocation of x509 root certificate by new trustee
	proposeRevokeX509RootCert := types.NewMsgProposeRevokeX509RootCert(
		constants.RootSubject, constants.RootSubjectKeyID, constants.Address1)
	result := setup.Handler(setup.Ctx, proposeRevokeX509RootCert)
	require.Equal(t, sdk.CodeOK, result.Code)

	// query and check proposed certificate revocation
	proposedRevocation, _ := queryProposedCertificateRevocation(&setup, constants.RootSubject, constants.RootSubjectKeyID)
	require.Equal(t, constants.RootSubject, proposedRevocation.Subject)
	require.Equal(t, constants.RootSubjectKeyID, proposedRevocation.SubjectKeyID)
	require.Equal(t, []sdk.AccAddress{constants.Address1}, proposedRevocation.Approvals)

	// check that approved certificate still exists
	certificate, _ := querySingleApprovedCertificate(&setup, constants.RootSubject, constants.RootSubjectKeyID)
	require.NotNil(t, certificate)

	// check that revoked certificate does not exist
	_, err := queryRevokedCertificates(&setup, constants.RootSubject, constants.RootSubjectKeyID)
	require.Equal(t, types.CodeRevokedCertificateDoesNotExist, err.Code())

	// check that unique certificate key stays registered
	require.True(t,
		setup.PkiKeeper.IsUniqueCertificateKeyPresent(setup.Ctx, constants.RootIssuer, constants.RootSerialNumber))
}

func TestHandler_ProposeRevokeX509RootCert_ByNotTrustee(t *testing.T) {
	setup := Setup()

	// propose and approve x509 root certificate
	proposeAndApproveRootCertificate(t, &setup, setup.Trustee)

	for _, role := range []auth.AccountRole{auth.TestHouse, auth.ZBCertificationCenter, auth.Vendor} {
		// assign role
		account := auth.NewAccount(constants.Address1, constants.PubKey1, auth.AccountRoles{role}, constants.VendorId1)
		setup.AuthKeeper.SetAccount(setup.Ctx, account)

		// propose revocation of x509 root certificate
		proposeRevokeX509RootCert := types.NewMsgProposeRevokeX509RootCert(
			constants.RootSubject, constants.RootSubjectKeyID, constants.Address1)
		result := setup.Handler(setup.Ctx, proposeRevokeX509RootCert)
		require.Equal(t, sdk.CodeUnauthorized, result.Code)
	}
}

func TestHandler_ProposeRevokeX509RootCert_CertificateDoesNotExist(t *testing.T) {
	setup := Setup()

	// propose revocation of not existing certificate
	proposeRevokeX509RootCert := types.NewMsgProposeRevokeX509RootCert(
		constants.RootSubject, constants.RootSubjectKeyID, setup.Trustee)
	result := setup.Handler(setup.Ctx, proposeRevokeX509RootCert)
	require.Equal(t, types.CodeCertificateDoesNotExist, result.Code)
}

func TestHandler_ProposeRevokeX509RootCert_ForProposedCertificate(t *testing.T) {
	setup := Setup()

	// propose x509 root certificate
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(constants.RootCertPem, setup.Trustee)
	result := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.Equal(t, sdk.CodeOK, result.Code)

	// check that proposed certificate is present
	proposedCertificate, _ := queryProposedCertificate(&setup, constants.RootSubject, constants.RootSubjectKeyID)
	require.NotNil(t, proposedCertificate)

	// propose revocation of proposed root certificate
	proposeRevokeX509RootCert := types.NewMsgProposeRevokeX509RootCert(
		constants.RootSubject, constants.RootSubjectKeyID, setup.Trustee)
	result = setup.Handler(setup.Ctx, proposeRevokeX509RootCert)
	require.Equal(t, types.CodeCertificateDoesNotExist, result.Code)
}

func TestHandler_ProposeRevokeX509RootCert_ProposedRevocationAlreadyExists(t *testing.T) {
	setup := Setup()

	// propose and approve x509 root certificate
	proposeAndApproveRootCertificate(t, &setup, setup.Trustee)

	// propose revocation of x509 root certificate
	proposeRevokeX509RootCert := types.NewMsgProposeRevokeX509RootCert(
		constants.RootSubject, constants.RootSubjectKeyID, setup.Trustee)
	result := setup.Handler(setup.Ctx, proposeRevokeX509RootCert)
	require.Equal(t, sdk.CodeOK, result.Code)

	// store another trustee account
	account := auth.NewAccount(constants.Address1, constants.PubKey1, auth.AccountRoles{auth.Trustee}, constants.VendorId1)
	setup.AuthKeeper.SetAccount(setup.Ctx, account)

	// propose revocation of the same x509 root certificate again
	proposeRevokeX509RootCert = types.NewMsgProposeRevokeX509RootCert(
		constants.RootSubject, constants.RootSubjectKeyID, setup.Trustee)
	result = setup.Handler(setup.Ctx, proposeRevokeX509RootCert)
	require.Equal(t, types.CodeProposedCertificateRevocationAlreadyExists, result.Code)
}

func TestHandler_ProposeRevokeX509RootCert_ForNonRootCertificate(t *testing.T) {
	setup := Setup()

	// store x509 root certificate
	rootCertificate := rootCertificate(setup.Trustee)
	setup.PkiKeeper.AddApprovedCertificate(setup.Ctx, rootCertificate)

	// store x509 intermediate certificate
	addX509Cert := types.NewMsgAddX509Cert(constants.IntermediateCertPem, setup.Trustee)
	result := setup.Handler(setup.Ctx, addX509Cert)
	require.Equal(t, sdk.CodeOK, result.Code)

	// propose revocation of x509 intermediate certificate
	proposeRevokeX509RootCert := types.NewMsgProposeRevokeX509RootCert(
		constants.IntermediateSubject, constants.IntermediateSubjectKeyID, setup.Trustee)
	result = setup.Handler(setup.Ctx, proposeRevokeX509RootCert)
	require.Equal(t, types.CodeInappropriateCertificateType, result.Code)
}

func TestHandler_ApproveRevokeX509RootCert_ForNotEnoughApprovals(t *testing.T) {
	setup := Setup()

	// propose and approve x509 root certificate
	proposeAndApproveRootCertificate(t, &setup, setup.Trustee)

	// increase the number of approvals required for root certificates control to three
	oldRootCertificateApprovals := types.RootCertificateApprovals

	defer func() {
		types.RootCertificateApprovals = oldRootCertificateApprovals
	}()

	types.RootCertificateApprovals = 3

	// propose revocation of x509 root certificate
	proposeRevokeX509RootCert := types.NewMsgProposeRevokeX509RootCert(
		constants.RootSubject, constants.RootSubjectKeyID, setup.Trustee)
	result := setup.Handler(setup.Ctx, proposeRevokeX509RootCert)
	require.Equal(t, sdk.CodeOK, result.Code)

	// store second trustee
	account := auth.NewAccount(constants.Address1, constants.PubKey1, auth.AccountRoles{auth.Trustee}, 0)
	setup.AuthKeeper.SetAccount(setup.Ctx, account)

	// approve
	approveRevokeX509RootCert := types.NewMsgApproveRevokeX509RootCert(
		constants.RootSubject, constants.RootSubjectKeyID, constants.Address1)
	result = setup.Handler(setup.Ctx, approveRevokeX509RootCert)
	require.Equal(t, sdk.CodeOK, result.Code)

	// query and check proposed certificate revocation
	proposedRevocation, _ := queryProposedCertificateRevocation(&setup, constants.RootSubject, constants.RootSubjectKeyID)
	require.Equal(t, constants.RootSubject, proposedRevocation.Subject)
	require.Equal(t, constants.RootSubjectKeyID, proposedRevocation.SubjectKeyID)
	require.Equal(t, []sdk.AccAddress{setup.Trustee, constants.Address1}, proposedRevocation.Approvals)

	// check that approved certificate still exists
	certificate, _ := querySingleApprovedCertificate(&setup, constants.RootSubject, constants.RootSubjectKeyID)
	require.NotNil(t, certificate)

	// check that revoked certificate does not exist
	_, err := queryRevokedCertificates(&setup, constants.RootSubject, constants.RootSubjectKeyID)
	require.Equal(t, types.CodeRevokedCertificateDoesNotExist, err.Code())

	// check that unique certificate key stays registered
	require.True(t,
		setup.PkiKeeper.IsUniqueCertificateKeyPresent(setup.Ctx, constants.RootIssuer, constants.RootSerialNumber))
}

func TestHandler_ApproveRevokeX509RootCert_ForEnoughApprovals(t *testing.T) {
	setup := Setup()

	// propose and approve x509 root certificate
	proposeAndApproveRootCertificate(t, &setup, setup.Trustee)

	// propose revocation of x509 root certificate
	proposeRevokeX509RootCert := types.NewMsgProposeRevokeX509RootCert(
		constants.RootSubject, constants.RootSubjectKeyID, setup.Trustee)
	result := setup.Handler(setup.Ctx, proposeRevokeX509RootCert)
	require.Equal(t, sdk.CodeOK, result.Code)

	// get certificate for further comparison
	certificateBeforeRevocation, _ :=
		querySingleApprovedCertificate(&setup, constants.RootSubject, constants.RootSubjectKeyID)
	require.NotNil(t, certificateBeforeRevocation)

	// store second trustee
	account := auth.NewAccount(constants.Address1, constants.PubKey1, auth.AccountRoles{auth.Trustee}, 0)
	setup.AuthKeeper.SetAccount(setup.Ctx, account)

	// approve
	approveRevokeX509RootCert := types.NewMsgApproveRevokeX509RootCert(
		constants.RootSubject, constants.RootSubjectKeyID, constants.Address1)
	result = setup.Handler(setup.Ctx, approveRevokeX509RootCert)
	require.Equal(t, sdk.CodeOK, result.Code)

	// check that proposed certificate revocation does not exist anymore
	_, err := queryProposedCertificateRevocation(&setup, constants.RootSubject, constants.RootSubjectKeyID)
	require.Equal(t, types.CodeProposedCertificateRevocationDoesNotExist, err.Code())

	// check that approved certificate does not exist anymore
	_, err = queryApprovedCertificates(&setup, constants.RootSubject, constants.RootSubjectKeyID)
	require.Equal(t, types.CodeCertificateDoesNotExist, err.Code())

	// query and check revoked certificate
	revokedCertificate, _ := querySingleRevokedCertificate(&setup, constants.RootSubject, constants.RootSubjectKeyID)
	require.Equal(t, certificateBeforeRevocation, revokedCertificate)

	// check that unique certificate key stays registered
	require.True(t,
		setup.PkiKeeper.IsUniqueCertificateKeyPresent(setup.Ctx, constants.RootIssuer, constants.RootSerialNumber))
}

func TestHandler_ApproveRevokeX509RootCert_ByNotTrustee(t *testing.T) {
	setup := Setup()

	// propose and approve x509 root certificate
	proposeAndApproveRootCertificate(t, &setup, setup.Trustee)

	// propose revocation of x509 root certificate
	proposeRevokeX509RootCert := types.NewMsgProposeRevokeX509RootCert(
		constants.RootSubject, constants.RootSubjectKeyID, setup.Trustee)
	result := setup.Handler(setup.Ctx, proposeRevokeX509RootCert)
	require.Equal(t, sdk.CodeOK, result.Code)

	for _, role := range []auth.AccountRole{auth.TestHouse, auth.ZBCertificationCenter, auth.Vendor} {
		// assign role
		account := auth.NewAccount(constants.Address1, constants.PubKey1, auth.AccountRoles{role}, constants.VendorId1)
		setup.AuthKeeper.SetAccount(setup.Ctx, account)

		// approve
		approveRevokeX509RootCert := types.NewMsgApproveRevokeX509RootCert(
			constants.RootSubject, constants.RootSubjectKeyID, constants.Address1)
		result = setup.Handler(setup.Ctx, approveRevokeX509RootCert)
		require.Equal(t, sdk.CodeUnauthorized, result.Code)
	}
}

func TestHandler_ApproveRevokeX509RootCert_ProposedRevocationDoesNotExist(t *testing.T) {
	setup := Setup()

	// propose and approve x509 root certificate
	proposeAndApproveRootCertificate(t, &setup, setup.Trustee)

	// approve revocation of x509 root certificate
	approveRevokeX509RootCert := types.NewMsgApproveRevokeX509RootCert(
		constants.RootSubject, constants.RootSubjectKeyID, setup.Trustee)
	result := setup.Handler(setup.Ctx, approveRevokeX509RootCert)
	require.Equal(t, types.CodeProposedCertificateRevocationDoesNotExist, result.Code)
}

func TestHandler_ApproveRevokeX509RootCert_Twice(t *testing.T) {
	setup := Setup()

	// propose and approve x509 root certificate
	proposeAndApproveRootCertificate(t, &setup, setup.Trustee)

	// propose revocation of x509 root certificate
	proposeRevokeX509RootCert := types.NewMsgProposeRevokeX509RootCert(
		constants.RootSubject, constants.RootSubjectKeyID, setup.Trustee)
	result := setup.Handler(setup.Ctx, proposeRevokeX509RootCert)
	require.Equal(t, sdk.CodeOK, result.Code)

	// approve revocation by the same trustee
	approveRevokeX509RootCert := types.NewMsgApproveRevokeX509RootCert(
		constants.RootSubject, constants.RootSubjectKeyID, setup.Trustee)
	result = setup.Handler(setup.Ctx, approveRevokeX509RootCert)
	require.Equal(t, sdk.CodeUnauthorized, result.Code)
}

//nolint:funlen
func TestHandler_ApproveRevokeX509RootCert_ForTree(t *testing.T) {
	setup := Setup()

	// add root x509 certificate
	proposeAndApproveRootCertificate(t, &setup, setup.Trustee)

	// add intermediate x509 certificate
	addIntermediateX509Cert := types.NewMsgAddX509Cert(constants.IntermediateCertPem, setup.Trustee)
	result := setup.Handler(setup.Ctx, addIntermediateX509Cert)
	require.Equal(t, sdk.CodeOK, result.Code)

	// add leaf x509 certificate
	addLeafX509Cert := types.NewMsgAddX509Cert(constants.LeafCertPem, setup.Trustee)
	result = setup.Handler(setup.Ctx, addLeafX509Cert)
	require.Equal(t, sdk.CodeOK, result.Code)

	// propose revocation of x509 root certificate
	proposeRevokeX509RootCert := types.NewMsgProposeRevokeX509RootCert(
		constants.RootSubject, constants.RootSubjectKeyID, setup.Trustee)
	result = setup.Handler(setup.Ctx, proposeRevokeX509RootCert)
	require.Equal(t, sdk.CodeOK, result.Code)

	// store second trustee
	account := auth.NewAccount(constants.Address1, constants.PubKey1, auth.AccountRoles{auth.Trustee}, constants.VendorId1)
	setup.AuthKeeper.SetAccount(setup.Ctx, account)

	// approve
	approveRevokeX509RootCert := types.NewMsgApproveRevokeX509RootCert(
		constants.RootSubject, constants.RootSubjectKeyID, constants.Address1)
	result = setup.Handler(setup.Ctx, approveRevokeX509RootCert)
	require.Equal(t, sdk.CodeOK, result.Code)

	// check that root, intermediate and leaf certificates have been revoked
	allRevokedCertificates, _ := queryAllRevokedCertificates(&setup)
	require.Equal(t, 3, len(allRevokedCertificates.Items))
	require.Equal(t, constants.IntermediateSubject, allRevokedCertificates.Items[0].Subject)
	require.Equal(t, constants.IntermediateSubjectKeyID, allRevokedCertificates.Items[0].SubjectKeyID)
	require.Equal(t, constants.IntermediateCertPem, allRevokedCertificates.Items[0].PemCert)
	require.Equal(t, constants.LeafSubject, allRevokedCertificates.Items[1].Subject)
	require.Equal(t, constants.LeafSubjectKeyID, allRevokedCertificates.Items[1].SubjectKeyID)
	require.Equal(t, constants.LeafCertPem, allRevokedCertificates.Items[1].PemCert)
	require.Equal(t, constants.RootSubject, allRevokedCertificates.Items[2].Subject)
	require.Equal(t, constants.RootSubjectKeyID, allRevokedCertificates.Items[2].SubjectKeyID)
	require.Equal(t, constants.RootCertPem, allRevokedCertificates.Items[2].PemCert)

	// check that no certificates stays approved
	allApprovedCertificates, _ := queryAllApprovedCertificates(&setup)
	require.Equal(t, 0, len(allApprovedCertificates.Items))

	// check that no proposed certificate revocations exist
	allProposedCertificateRevocations, _ := queryAllProposedCertificateRevocations(&setup)
	require.Equal(t, 0, len(allProposedCertificateRevocations.Items))

	// check that no child certificate identifiers are registered for revoked root certificate
	rootCertChildren := setup.PkiKeeper.GetChildCertificates(setup.Ctx,
		constants.RootSubject, constants.RootSubjectKeyID)
	require.Equal(t, 0, len(rootCertChildren.CertIdentifiers))

	// check that no child certificate identifiers are registered for revoked intermediate certificate
	intermediateCertChildren := setup.PkiKeeper.GetChildCertificates(setup.Ctx,
		constants.IntermediateSubject, constants.IntermediateSubjectKeyID)
	require.Equal(t, 0, len(intermediateCertChildren.CertIdentifiers))

	// check that no child certificate identifiers are registered for revoked leaf certificate
	leafCertChildren := setup.PkiKeeper.GetChildCertificates(setup.Ctx,
		constants.LeafSubject, constants.LeafSubjectKeyID)
	require.Equal(t, 0, len(leafCertChildren.CertIdentifiers))
}

func TestHandler_RevokeX509Cert(t *testing.T) {
	setup := Setup()

	// store root certificate
	rootCertificate := rootCertificate(setup.Trustee)
	setup.PkiKeeper.AddApprovedCertificate(setup.Ctx, rootCertificate)

	for _, role := range []auth.AccountRole{auth.Trustee, auth.TestHouse, auth.ZBCertificationCenter, auth.Vendor} {
		// assign role
		account := auth.NewAccount(constants.Address1, constants.PubKey1, auth.AccountRoles{role}, constants.VendorId1)
		setup.AuthKeeper.SetAccount(setup.Ctx, account)

		// add x509 certificate
		addX509Cert := types.NewMsgAddX509Cert(constants.IntermediateCertPem, constants.Address1)
		result := setup.Handler(setup.Ctx, addX509Cert)
		require.Equal(t, sdk.CodeOK, result.Code)

		// get certificate for further comparison
		certificateBeforeRevocation, _ := querySingleApprovedCertificate(&setup,
			constants.IntermediateSubject, constants.IntermediateSubjectKeyID)
		require.NotNil(t, certificateBeforeRevocation)

		// revoke x509 certificate
		revokeX509Cert := types.NewMsgRevokeX509Cert(
			constants.IntermediateSubject, constants.IntermediateSubjectKeyID, constants.Address1)
		result = setup.Handler(setup.Ctx, revokeX509Cert)
		require.Equal(t, sdk.CodeOK, result.Code)

		// check that intermediate certificate has been revoked
		allRevokedCertificates, _ := queryAllRevokedCertificates(&setup)
		require.Equal(t, 1, len(allRevokedCertificates.Items))
		require.Equal(t, constants.IntermediateSubject, allRevokedCertificates.Items[0].Subject)
		require.Equal(t, constants.IntermediateSubjectKeyID, allRevokedCertificates.Items[0].SubjectKeyID)
		require.Equal(t, certificateBeforeRevocation, &allRevokedCertificates.Items[0])

		// check that root certificate stays approved
		allApprovedCertificates, _ := queryAllApprovedCertificates(&setup)
		require.Equal(t, 1, len(allApprovedCertificates.Items))
		require.Equal(t, constants.IntermediateSubject, allRevokedCertificates.Items[0].Subject)
		require.Equal(t, constants.IntermediateSubjectKeyID, allRevokedCertificates.Items[0].SubjectKeyID)

		// check that no proposed certificate revocations have been created
		allProposedCertificateRevocations, _ := queryAllProposedCertificateRevocations(&setup)
		require.Equal(t, 0, len(allProposedCertificateRevocations.Items))

		// check that child certificate identifiers list of issuer do not exist anymore
		require.False(t, setup.PkiKeeper.IsChildCertificatesPresent(setup.Ctx,
			constants.IntermediateIssuer, constants.IntermediateAuthorityKeyID))

		// check that unique certificate key stays registered
		require.True(t, setup.PkiKeeper.IsUniqueCertificateKeyPresent(setup.Ctx,
			constants.IntermediateIssuer, constants.IntermediateSerialNumber))

		// cleanup for next iteration
		setup.PkiKeeper.DeleteRevokedCertificates(setup.Ctx,
			constants.IntermediateSubject, constants.IntermediateSubjectKeyID)
		setup.PkiKeeper.DeleteUniqueCertificateKey(setup.Ctx,
			constants.IntermediateIssuer, constants.IntermediateSerialNumber)
	}
}

func TestHandler_RevokeX509Cert_CertificateDoesNotExist(t *testing.T) {
	setup := Setup()

	// revoke x509 certificate
	revokeX509Cert := types.NewMsgRevokeX509Cert(
		constants.IntermediateSubject, constants.IntermediateSubjectKeyID, setup.Trustee)
	result := setup.Handler(setup.Ctx, revokeX509Cert)
	require.Equal(t, types.CodeCertificateDoesNotExist, result.Code)
}

func TestHandler_RevokeX509Cert_ForRootCertificate(t *testing.T) {
	setup := Setup()

	// propose and approve x509 root certificate
	proposeAndApproveRootCertificate(t, &setup, setup.Trustee)

	// revoke x509 root certificate
	revokeX509Cert := types.NewMsgRevokeX509Cert(
		constants.RootSubject, constants.RootSubjectKeyID, setup.Trustee)
	result := setup.Handler(setup.Ctx, revokeX509Cert)
	require.Equal(t, types.CodeInappropriateCertificateType, result.Code)
}

func TestHandler_RevokeX509Cert_ByNotOwner(t *testing.T) {
	setup := Setup()

	// store root certificate
	rootCertificate := rootCertificate(setup.Trustee)
	setup.PkiKeeper.AddApprovedCertificate(setup.Ctx, rootCertificate)

	// add x509 certificate by `setup.Trustee`
	addX509Cert := types.NewMsgAddX509Cert(constants.IntermediateCertPem, setup.Trustee)
	result := setup.Handler(setup.Ctx, addX509Cert)
	require.Equal(t, sdk.CodeOK, result.Code)

	// store another account
	account := auth.NewAccount(constants.Address1, constants.PubKey1, auth.AccountRoles{auth.Trustee}, constants.VendorId1)
	setup.AuthKeeper.SetAccount(setup.Ctx, account)

	// revoke x509 certificate by another account
	revokeX509Cert := types.NewMsgRevokeX509Cert(
		constants.IntermediateSubject, constants.IntermediateSubjectKeyID, constants.Address1)
	result = setup.Handler(setup.Ctx, revokeX509Cert)
	require.Equal(t, sdk.CodeUnauthorized, result.Code)
}

func TestHandler_RevokeX509Cert_ForTree(t *testing.T) {
	setup := Setup()

	// add root x509 certificate
	proposeAndApproveRootCertificate(t, &setup, setup.Trustee)

	// add intermediate x509 certificate
	addIntermediateX509Cert := types.NewMsgAddX509Cert(constants.IntermediateCertPem, setup.Trustee)
	result := setup.Handler(setup.Ctx, addIntermediateX509Cert)
	require.Equal(t, sdk.CodeOK, result.Code)

	// add leaf x509 certificate
	addLeafX509Cert := types.NewMsgAddX509Cert(constants.LeafCertPem, setup.Trustee)
	result = setup.Handler(setup.Ctx, addLeafX509Cert)
	require.Equal(t, sdk.CodeOK, result.Code)

	// revoke x509 certificate
	revokeX509Cert := types.NewMsgRevokeX509Cert(
		constants.IntermediateSubject, constants.IntermediateSubjectKeyID, setup.Trustee)
	result = setup.Handler(setup.Ctx, revokeX509Cert)
	require.Equal(t, sdk.CodeOK, result.Code)

	// check that intermediate and leaf certificates have been revoked
	allRevokedCertificates, _ := queryAllRevokedCertificates(&setup)
	require.Equal(t, 2, len(allRevokedCertificates.Items))
	require.Equal(t, constants.IntermediateSubject, allRevokedCertificates.Items[0].Subject)
	require.Equal(t, constants.IntermediateSubjectKeyID, allRevokedCertificates.Items[0].SubjectKeyID)
	require.Equal(t, constants.IntermediateCertPem, allRevokedCertificates.Items[0].PemCert)
	require.Equal(t, constants.LeafSubject, allRevokedCertificates.Items[1].Subject)
	require.Equal(t, constants.LeafSubjectKeyID, allRevokedCertificates.Items[1].SubjectKeyID)
	require.Equal(t, constants.LeafCertPem, allRevokedCertificates.Items[1].PemCert)

	// check that root certificate stays approved
	allApprovedCertificates, _ := queryAllApprovedCertificates(&setup)
	require.Equal(t, 1, len(allApprovedCertificates.Items))
	require.Equal(t, constants.RootSubject, allApprovedCertificates.Items[0].Subject)
	require.Equal(t, constants.RootSubjectKeyID, allApprovedCertificates.Items[0].SubjectKeyID)

	// check that no proposed certificate revocations have been created
	allProposedCertificateRevocations, _ := queryAllProposedCertificateRevocations(&setup)
	require.Equal(t, 0, len(allProposedCertificateRevocations.Items))

	// check that no child certificate identifiers are now registered for root certificate
	rootCertChildren := setup.PkiKeeper.GetChildCertificates(setup.Ctx,
		constants.RootSubject, constants.RootSubjectKeyID)
	require.Equal(t, 0, len(rootCertChildren.CertIdentifiers))

	// check that no child certificate identifiers are registered for revoked intermediate certificate
	intermediateCertChildren := setup.PkiKeeper.GetChildCertificates(setup.Ctx,
		constants.IntermediateSubject, constants.IntermediateSubjectKeyID)
	require.Equal(t, 0, len(intermediateCertChildren.CertIdentifiers))

	// check that no child certificate identifiers are registered for revoked leaf certificate
	leafCertChildren := setup.PkiKeeper.GetChildCertificates(setup.Ctx,
		constants.LeafSubject, constants.LeafSubjectKeyID)
	require.Equal(t, 0, len(leafCertChildren.CertIdentifiers))
}

func proposeAndApproveRootCertificate(t *testing.T, setup *TestSetup, ownerTrustee sdk.AccAddress) {
	// ensure that `ownerTrustee` is trustee to eventually have enough approvals
	require.True(t, setup.AuthKeeper.HasRole(setup.Ctx, ownerTrustee, types.RootCertificateApprovalRole))

	// propose x509 root certificate by `ownerTrustee`
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(constants.RootCertPem, ownerTrustee)
	result := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.Equal(t, sdk.CodeOK, result.Code)

	// store another trustee account
	account := auth.NewAccount(constants.Address3, constants.PubKey3, auth.AccountRoles{auth.Trustee}, constants.VendorId3)
	setup.AuthKeeper.SetAccount(setup.Ctx, account)

	// approve x509 root certificate by another trustee
	approveAddX509RootCert := types.NewMsgApproveAddX509RootCert(
		constants.RootSubject, constants.RootSubjectKeyID, constants.Address3)
	result = setup.Handler(setup.Ctx, approveAddX509RootCert)
	require.Equal(t, sdk.CodeOK, result.Code)

	// check that root certificate has been approved
	approvedCertificate, _ := querySingleApprovedCertificate(setup, constants.RootSubject, constants.RootSubjectKeyID)
	require.NotNil(t, approvedCertificate)
}

func queryProposedCertificate(setup *TestSetup, subject string,
	subjectKeyID string) (*types.ProposedCertificate, sdk.Error) {
	// query proposed certificate
	result, err := setup.Querier(
		setup.Ctx,
		[]string{keeper.QueryProposedX509RootCert, subject, subjectKeyID},
		abci.RequestQuery{},
	)
	if err != nil {
		return nil, err
	}

	var proposedCertificate types.ProposedCertificate
	_ = setup.Cdc.UnmarshalJSON(result, &proposedCertificate)

	return &proposedCertificate, nil
}

func queryAllApprovedCertificates(setup *TestSetup) (*types.ListCertificates, sdk.Error) {
	// query all certificates
	result, err := setup.Querier(
		setup.Ctx,
		[]string{keeper.QueryAllX509Certs},
		abci.RequestQuery{Data: emptyParams(setup)},
	)
	if err != nil {
		return nil, err
	}

	var certificates types.ListCertificates
	_ = setup.Cdc.UnmarshalJSON(result, &certificates)

	return &certificates, nil
}

func querySingleApprovedCertificate(setup *TestSetup,
	subject string, subjectKeyID string) (*types.Certificate, sdk.Error) {
	certificates, err := queryApprovedCertificates(setup, subject, subjectKeyID)
	if err != nil {
		return nil, err
	}

	if len(certificates.Items) > 1 {
		return nil, sdk.ErrInternal("More than 1 certificate returned")
	}

	return &certificates.Items[0], nil
}

func queryApprovedCertificates(setup *TestSetup, subject string, subjectKeyID string) (*types.Certificates, sdk.Error) {
	// query certificate
	result, err := setup.Querier(
		setup.Ctx,
		[]string{keeper.QueryX509Cert, subject, subjectKeyID},
		abci.RequestQuery{},
	)
	if err != nil {
		return nil, err
	}

	var certificates types.Certificates
	_ = setup.Cdc.UnmarshalJSON(result, &certificates)

	return &certificates, nil
}

func queryAllProposedCertificateRevocations(setup *TestSetup) (*types.ListProposedCertificateRevocations, sdk.Error) {
	// query all proposed certificate revocations
	result, err := setup.Querier(
		setup.Ctx,
		[]string{keeper.QueryAllProposedX509RootCertRevocations},
		abci.RequestQuery{Data: emptyParams(setup)},
	)
	if err != nil {
		return nil, err
	}

	var listProposedCertificateRevocations types.ListProposedCertificateRevocations
	_ = setup.Cdc.UnmarshalJSON(result, &listProposedCertificateRevocations)

	return &listProposedCertificateRevocations, nil
}

func queryProposedCertificateRevocation(setup *TestSetup, subject string,
	subjectKeyID string) (*types.ProposedCertificateRevocation, sdk.Error) {
	// query proposed certificate revocation
	result, err := setup.Querier(
		setup.Ctx,
		[]string{keeper.QueryProposedX509RootCertRevocation, subject, subjectKeyID},
		abci.RequestQuery{},
	)
	if err != nil {
		return nil, err
	}

	var proposedCertificateRevocation types.ProposedCertificateRevocation
	_ = setup.Cdc.UnmarshalJSON(result, &proposedCertificateRevocation)

	return &proposedCertificateRevocation, nil
}

func queryAllRevokedCertificates(setup *TestSetup) (*types.ListCertificates, sdk.Error) {
	// query all revoked certificates
	result, err := setup.Querier(
		setup.Ctx,
		[]string{keeper.QueryAllRevokedX509Certs},
		abci.RequestQuery{Data: emptyParams(setup)},
	)
	if err != nil {
		return nil, err
	}

	var certificates types.ListCertificates
	_ = setup.Cdc.UnmarshalJSON(result, &certificates)

	return &certificates, nil
}

func querySingleRevokedCertificate(setup *TestSetup,
	subject string, subjectKeyID string) (*types.Certificate, sdk.Error) {
	certificates, err := queryRevokedCertificates(setup, subject, subjectKeyID)
	if err != nil {
		return nil, err
	}

	if len(certificates.Items) > 1 {
		return nil, sdk.ErrInternal("More than 1 revoked certificate returned")
	}

	return &certificates.Items[0], nil
}

func queryRevokedCertificates(setup *TestSetup, subject string, subjectKeyID string) (*types.Certificates, sdk.Error) {
	// query revoked certificate
	result, err := setup.Querier(
		setup.Ctx,
		[]string{keeper.QueryRevokedX509Cert, subject, subjectKeyID},
		abci.RequestQuery{},
	)
	if err != nil {
		return nil, err
	}

	var certificates types.Certificates
	_ = setup.Cdc.UnmarshalJSON(result, &certificates)

	return &certificates, nil
}

func rootCertificate(address sdk.AccAddress) types.Certificate {
	return types.NewRootCertificate(
		constants.RootCertPem,
		constants.RootSubject,
		constants.RootSubjectKeyID,
		constants.RootSerialNumber,
		address,
	)
}

func intermediateCertificate(address sdk.AccAddress) types.Certificate {
	return types.NewNonRootCertificate(
		constants.IntermediateCertPem,
		constants.IntermediateSubject,
		constants.IntermediateSubjectKeyID,
		constants.IntermediateSerialNumber,
		constants.IntermediateIssuer,
		constants.IntermediateAuthorityKeyID,
		constants.RootSubject,
		constants.RootSubjectKeyID,
		address,
	)
}

func emptyParams(setup *TestSetup) []byte {
	paginationParams := pagination.NewPaginationParams(0, 0)

	return setup.Cdc.MustMarshalJSON(types.NewPkiQueryParams(paginationParams, "", ""))
}
