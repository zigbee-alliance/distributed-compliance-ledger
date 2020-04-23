package pki

import (
	constants "git.dsr-corporation.com/zb-ledger/zb-ledger/integration_tests/constants"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/authz"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/pki/internal/keeper"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/pki/internal/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	"testing"
)

func TestHandler_ProposeAddX509RootCertByNotTrustee(t *testing.T) {
	setup := Setup()

	for _, role := range []authz.AccountRole{authz.Administrator, authz.TestHouse, authz.ZBCertificationCenter, authz.Vendor} {
		// assign role
		setup.AuthzKeeper.AssignRole(setup.Ctx, constants.Address1, role)

		// propose x509 root certificate
		proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(constants.RootCertPem, constants.Address1)
		result := setup.Handler(setup.Ctx, proposeAddX509RootCert)
		require.Equal(t, sdk.CodeOK, result.Code)

		// query proposed certificate
		proposedCertificate, _ := queryProposedCertificate(setup, constants.RootSubject, constants.RootSubjectKeyId)

		// check
		require.Equal(t, proposeAddX509RootCert.Cert, proposedCertificate.PemCert)
		require.Equal(t, proposeAddX509RootCert.Signer, proposedCertificate.Owner)
		require.Equal(t, constants.RootSubject, proposedCertificate.Subject)
		require.Equal(t, constants.RootSubjectKeyId, proposedCertificate.SubjectKeyId)
		require.Nil(t, proposedCertificate.Approvals)

		// try to query approved certificate
		_, err := querySingleCertificate(setup, constants.RootSubject, constants.RootSubjectKeyId)
		require.Equal(t, types.CodeCertificateDoesNotExist, err.Code())

		// delete proposed certificate for next iteration
		setup.PkiKeeper.DeleteProposedCertificate(setup.Ctx, constants.RootSubject, constants.RootSubjectKeyId)
	}
}

func TestHandler_ProposeAddX509RootCertByTrustee(t *testing.T) {
	setup := Setup()

	// propose x509 root certificate
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(constants.RootCertPem, setup.Trustee)
	result := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.Equal(t, sdk.CodeOK, result.Code)

	// query proposed certificate
	proposedCertificate, _ := queryProposedCertificate(setup, constants.RootSubject, constants.RootSubjectKeyId)

	// check
	require.Equal(t, proposeAddX509RootCert.Cert, proposedCertificate.PemCert)
	require.Equal(t, proposeAddX509RootCert.Signer, proposedCertificate.Owner)
	require.Equal(t, constants.RootSubject, proposedCertificate.Subject)
	require.Equal(t, constants.RootSubjectKeyId, proposedCertificate.SubjectKeyId)
	require.Equal(t, []sdk.AccAddress{proposeAddX509RootCert.Signer}, proposedCertificate.Approvals)

	// query approved certificate
	_, err := querySingleCertificate(setup, constants.RootSubject, constants.RootSubjectKeyId)
	require.Equal(t, types.CodeCertificateDoesNotExist, err.Code())
}

func TestHandler_ProposeAddX509RootCert_ForInvalidCertificate(t *testing.T) {
	setup := Setup()

	// propose x509 root certificate
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(constants.StubCert, setup.Trustee)
	result := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.Equal(t, sdk.CodeInternal, result.Code)
}

func TestHandler_ProposeAddX509RootCert_ForNotRootCertificate(t *testing.T) {
	setup := Setup()

	// propose x509 leaf certificate as root
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(constants.LeafCertPem, setup.Trustee)
	result := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.Equal(t, types.CodeInappropriateCertificateType, result.Code)
}

func TestHandler_ProposeAddX509RootCert_Twice(t *testing.T) {
	setup := Setup()

	// propose x509 root certificate
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(constants.RootCertPem, setup.Trustee)
	result := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.Equal(t, sdk.CodeOK, result.Code)

	// again propose add x509 root certificate
	result = setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.Equal(t, types.CodeCertificateAlreadyExists, result.Code)
}

func TestHandler_ProposeAddX509RootCert_CertificateAlreadyExists(t *testing.T) {
	setup := Setup()

	// store root certificate
	certificate := rootCertificate(setup.Trustee)
	setup.PkiKeeper.SetCertificate(setup.Ctx, certificate)

	for _, role := range []authz.AccountRole{authz.Trustee, authz.Administrator, authz.TestHouse, authz.ZBCertificationCenter, authz.Vendor} {
		// assign role
		setup.AuthzKeeper.AssignRole(setup.Ctx, constants.Address1, role)

		// again propose adding of x509 root certificate
		proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(certificate.PemCert, constants.Address1)
		result := setup.Handler(setup.Ctx, proposeAddX509RootCert)
		require.Equal(t, types.CodeCertificateAlreadyExists, result.Code)
	}
}

func TestHandler_ApproveAddX509RootCert_ForNotEnoughApprovals(t *testing.T) {
	setup := Setup()

	// propose x509 root certificate
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(constants.RootCertPem, constants.Address1)
	result := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.Equal(t, sdk.CodeOK, result.Code)

	// approve
	approveAddX509RootCert := types.NewMsgApproveAddX509RootCert(constants.RootSubject, constants.RootSubjectKeyId, setup.Trustee)
	result = setup.Handler(setup.Ctx, approveAddX509RootCert)
	require.Equal(t, sdk.CodeOK, result.Code)

	// query certificate
	proposedCertificate, _ := queryProposedCertificate(setup, constants.RootSubject, constants.RootSubjectKeyId)
	require.Equal(t, proposeAddX509RootCert.Cert, proposedCertificate.PemCert)
	require.Equal(t, []sdk.AccAddress{setup.Trustee}, proposedCertificate.Approvals)

	// query approved certificate
	_, err := querySingleCertificate(setup, constants.RootSubject, constants.RootSubjectKeyId)
	require.Equal(t, types.CodeCertificateDoesNotExist, err.Code())
}

func TestHandler_ApproveAddX509RootCert_ForEnoughApprovals(t *testing.T) {
	setup := Setup()

	// propose add x509 root certificate as Trustee
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(constants.RootCertPem, setup.Trustee)
	result := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.Equal(t, sdk.CodeOK, result.Code)

	// set second Trustee
	setup.AuthzKeeper.AssignRole(setup.Ctx, constants.Address1, authz.Trustee)

	// second Trustee approve
	approveAddX509RootCert := types.NewMsgApproveAddX509RootCert(constants.RootSubject, constants.RootSubjectKeyId, constants.Address1)
	result = setup.Handler(setup.Ctx, approveAddX509RootCert)
	require.Equal(t, sdk.CodeOK, result.Code)

	// query proposed certificate
	_, err := queryProposedCertificate(setup, constants.RootSubject, constants.RootSubjectKeyId)
	require.Equal(t, types.CodePendingCertificateDoesNotExist, err.Code())

	// query approved certificate
	approvedCertificate, _ := querySingleCertificate(setup, constants.RootSubject, constants.RootSubjectKeyId)
	require.Equal(t, proposeAddX509RootCert.Cert, approvedCertificate.PemCert)
	require.Equal(t, proposeAddX509RootCert.Signer, approvedCertificate.Owner)
	require.Equal(t, constants.RootSubject, approvedCertificate.Subject)
	require.Equal(t, constants.RootSubjectKeyId, approvedCertificate.SubjectKeyId)
	require.Equal(t, types.RootCertificate, approvedCertificate.Type)
}

func TestHandler_ApproveAddX509RootCert_ForUnknownProposedCertificate(t *testing.T) {
	setup := Setup()

	// approve
	approveAddX509RootCert := types.NewMsgApproveAddX509RootCert(constants.RootSubject, constants.RootSubjectKeyId, setup.Trustee)
	result := setup.Handler(setup.Ctx, approveAddX509RootCert)
	require.Equal(t, types.CodePendingCertificateDoesNotExist, result.Code)
}

func TestHandler_ApproveAddX509RootCert_ForNotTrustee(t *testing.T) {
	setup := Setup()

	// propose add x509 root certificate
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(constants.RootCertPem, setup.Trustee)
	result := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.Equal(t, sdk.CodeOK, result.Code)

	for _, role := range []authz.AccountRole{authz.Administrator, authz.TestHouse, authz.ZBCertificationCenter, authz.Vendor} {
		// assign role
		setup.AuthzKeeper.AssignRole(setup.Ctx, constants.Address1, role)

		// approve
		approveAddX509RootCert := types.NewMsgApproveAddX509RootCert(constants.RootSubject, constants.RootSubjectKeyId, constants.Address1)
		result = setup.Handler(setup.Ctx, approveAddX509RootCert)
		require.Equal(t, sdk.CodeUnauthorized, result.Code)
	}
}

func TestHandler_ApproveAddX509RootCert_Twice(t *testing.T) {
	setup := Setup()

	// propose add x509 root certificate
	proposeAddX509RootCert := types.NewMsgProposeAddX509RootCert(constants.RootCertPem, constants.Address1)
	result := setup.Handler(setup.Ctx, proposeAddX509RootCert)
	require.Equal(t, sdk.CodeOK, result.Code)

	// approve
	approveAddX509RootCert := types.NewMsgApproveAddX509RootCert(constants.RootSubject, constants.RootSubjectKeyId, setup.Trustee)
	result = setup.Handler(setup.Ctx, approveAddX509RootCert)
	require.Equal(t, sdk.CodeOK, result.Code)

	// approve second time
	result = setup.Handler(setup.Ctx, approveAddX509RootCert)
	require.Equal(t, sdk.CodeInternal, result.Code)
}

func TestHandler_AddX509Cert(t *testing.T) {
	setup := Setup()

	// store root certificate
	rootCertificate := rootCertificate(setup.Trustee)
	setup.PkiKeeper.SetCertificate(setup.Ctx, rootCertificate)

	for _, role := range []authz.AccountRole{authz.Trustee, authz.Administrator, authz.TestHouse, authz.ZBCertificationCenter, authz.Vendor} {
		// assign role
		setup.AuthzKeeper.AssignRole(setup.Ctx, constants.Address1, role)

		// add x509 certificate
		addX509Cert := types.NewMsgAddX509Cert(constants.IntermediateCertPem, constants.Address1)
		result := setup.Handler(setup.Ctx, addX509Cert)
		require.Equal(t, sdk.CodeOK, result.Code)

		// query certificate
		certificate, _ := querySingleCertificate(setup, constants.IntermediateSubject, constants.IntermediateSubjectKeyId)

		// check
		require.Equal(t, addX509Cert.Cert, certificate.PemCert)
		require.Equal(t, addX509Cert.Signer, certificate.Owner)
		require.Equal(t, constants.IntermediateSubject, certificate.Subject)
		require.Equal(t, constants.IntermediateSubjectKeyId, certificate.SubjectKeyId)
		require.Equal(t, types.IntermediateCertificate, certificate.Type)
		require.Equal(t, constants.RootSubjectKeyId, certificate.RootSubjectKeyId)

		// query proposed certificate
		_, err := queryProposedCertificate(setup, constants.RootSubject, constants.RootSubjectKeyId)
		require.Equal(t, types.CodePendingCertificateDoesNotExist, err.Code())

		// delete for next iteration
		setup.PkiKeeper.DeleteCertificates(setup.Ctx, certificate.Subject, certificate.SubjectKeyId)
	}
}

func TestHandler_AddX509Cert_ForInvalidCertificate(t *testing.T) {
	setup := Setup()

	// add x509 certificate
	addX509Cert := types.NewMsgAddX509Cert(constants.StubCert, setup.Trustee)
	result := setup.Handler(setup.Ctx, addX509Cert)
	require.Equal(t, sdk.CodeInternal, result.Code)
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
	setup.PkiKeeper.SetCertificate(setup.Ctx, rootCertificate)

	// store root intermediate
	intermediateCertificate := intermediateCertificate(setup.Trustee)
	setup.PkiKeeper.SetCertificate(setup.Ctx, intermediateCertificate)

	// store root intermediate second time
	addX509Cert := types.NewMsgAddX509Cert(constants.IntermediateCertPem, constants.Address1)
	result := setup.Handler(setup.Ctx, addX509Cert)
	require.Equal(t, types.CodeCertificateAlreadyExists, result.Code)
}

func TestHandler_AddX509Cert_ForDifferentSerialNumber(t *testing.T) {
	setup := Setup()

	// store root certificate
	rootCertificate := rootCertificate(setup.Trustee)
	setup.PkiKeeper.SetCertificate(setup.Ctx, rootCertificate)

	// store root intermediate with different serial number
	intermediateCertificate := intermediateCertificate(setup.Trustee)
	intermediateCertificate.SerialNumber = "12345678"
	setup.PkiKeeper.SetCertificate(setup.Ctx, intermediateCertificate)

	// store root intermediate second time
	addX509Cert := types.NewMsgAddX509Cert(constants.IntermediateCertPem, setup.Trustee)
	result := setup.Handler(setup.Ctx, addX509Cert)
	require.Equal(t, sdk.CodeOK, result.Code)

	// query certificate
	certificates, _ := queryCertificates(setup, constants.IntermediateSubject, constants.IntermediateSubjectKeyId)

	// check
	require.Equal(t, 2, len(certificates.Items))
	require.NotEqual(t, certificates.Items[0].SerialNumber, certificates.Items[1].SerialNumber)

	for _, certificate := range certificates.Items {
		require.Equal(t, addX509Cert.Cert, certificate.PemCert)
		require.Equal(t, addX509Cert.Signer, certificate.Owner)
		require.Equal(t, constants.IntermediateSubject, certificate.Subject)
		require.Equal(t, constants.IntermediateSubjectKeyId, certificate.SubjectKeyId)
		require.Equal(t, types.IntermediateCertificate, certificate.Type)
		require.Equal(t, constants.RootSubjectKeyId, certificate.RootSubjectKeyId)
	}
}

func TestHandler_AddX509Cert_ForDifferentSerialNumberDifferentSigner(t *testing.T) {
	setup := Setup()

	// store root certificate
	rootCertificate := rootCertificate(setup.Trustee)
	setup.PkiKeeper.SetCertificate(setup.Ctx, rootCertificate)

	// store root intermediate with different serial number
	intermediateCertificate := intermediateCertificate(setup.Trustee)
	intermediateCertificate.SerialNumber = "12345678"
	setup.PkiKeeper.SetCertificate(setup.Ctx, intermediateCertificate)

	// store root intermediate second time
	addX509Cert := types.NewMsgAddX509Cert(constants.IntermediateCertPem, constants.Address1)
	result := setup.Handler(setup.Ctx, addX509Cert)
	require.Equal(t, sdk.CodeUnauthorized, result.Code)
}

func TestHandler_AddX509Cert_ForAbsentDirectParentCert(t *testing.T) {
	setup := Setup()

	// add intermediate x509 certificate
	addX509Cert := types.NewMsgAddX509Cert(constants.IntermediateCertPem, constants.Address1)
	result := setup.Handler(setup.Ctx, addX509Cert)
	require.Equal(t, types.CodeCertificateDoesNotExist, result.Code)
}

func TestHandler_AddX509Cert_ForNoRootCert(t *testing.T) {
	setup := Setup()

	// add intermediate
	intermediateCertificate := intermediateCertificate(setup.Trustee)
	setup.PkiKeeper.SetCertificate(setup.Ctx, intermediateCertificate)

	// add leaf x509 certificate
	addX509Cert := types.NewMsgAddX509Cert(constants.LeafCertPem, setup.Trustee)
	result := setup.Handler(setup.Ctx, addX509Cert)
	require.Equal(t, sdk.CodeInternal, result.Code)
}

func TestHandler_AddX509Cert_ForFailedCertificateVerification(t *testing.T) {
	setup := Setup()

	// add invalid root
	intermediateCertificate := types.NewRootCertificate(constants.StubCert, constants.RootSubject, constants.RootSubjectKeyId, constants.RootSerialNumber, setup.Trustee)
	setup.PkiKeeper.SetCertificate(setup.Ctx, intermediateCertificate)

	// add intermediate x509 certificate
	addX509Cert := types.NewMsgAddX509Cert(constants.IntermediateCertPem, setup.Trustee)
	result := setup.Handler(setup.Ctx, addX509Cert)
	require.Equal(t, sdk.CodeInternal, result.Code)
}

func TestHandler_AddX509Cert_ChildRefersToTwoParents(t *testing.T) {
	setup := Setup()

	// store root certificate
	rootCert := rootCertificate(setup.Trustee)
	setup.PkiKeeper.SetCertificate(setup.Ctx, rootCert)

	// store second root certificate
	rootCert = rootCertificate(setup.Trustee)
	rootCert.SerialNumber = "12345678"
	setup.PkiKeeper.SetCertificate(setup.Ctx, rootCert)

	// store root intermediate with different serial number
	intermediateCertificate := intermediateCertificate(setup.Trustee)
	intermediateCertificate.SerialNumber = "12345678"
	setup.PkiKeeper.SetCertificate(setup.Ctx, intermediateCertificate)

	// store intermediate second time
	addX509Cert := types.NewMsgAddX509Cert(constants.IntermediateCertPem, setup.Trustee)
	result := setup.Handler(setup.Ctx, addX509Cert)
	require.Equal(t, sdk.CodeOK, result.Code)

	// store leaf which matches to two intermediate
	addX509Cert = types.NewMsgAddX509Cert(constants.LeafCertPem, setup.Trustee)
	result = setup.Handler(setup.Ctx, addX509Cert)
	require.Equal(t, sdk.CodeOK, result.Code)

	certificate, _ := querySingleCertificate(setup, constants.LeafSubject, constants.LeafSubjectKeyId)
	require.Equal(t, addX509Cert.Cert, certificate.PemCert)
	require.Equal(t, addX509Cert.Signer, certificate.Owner)
	require.Equal(t, constants.LeafSubject, certificate.Subject)
	require.Equal(t, constants.LeafSubjectKeyId, certificate.SubjectKeyId)
	require.Equal(t, types.IntermediateCertificate, certificate.Type)
	require.Equal(t, constants.RootSubjectKeyId, certificate.RootSubjectKeyId)
}

func TestHandler_AddX509Cert_ForChain(t *testing.T) {
	setup := Setup()

	// add root x509 certificate
	rootCertificate := rootCertificate(setup.Trustee)
	setup.PkiKeeper.SetCertificate(setup.Ctx, rootCertificate)

	// add intermediate x509 certificate
	addIntermediateX509Cert := types.NewMsgAddX509Cert(constants.IntermediateCertPem, constants.Address1)
	result := setup.Handler(setup.Ctx, addIntermediateX509Cert)
	require.Equal(t, sdk.CodeOK, result.Code)

	// add leaf x509 certificate
	addLeafX509Cert := types.NewMsgAddX509Cert(constants.LeafCertPem, constants.Address1)
	result = setup.Handler(setup.Ctx, addLeafX509Cert)
	require.Equal(t, sdk.CodeOK, result.Code)

	// query intermediate certificate
	intermediateCertificate, _ := querySingleCertificate(setup, constants.IntermediateSubject, constants.IntermediateSubjectKeyId)
	require.Equal(t, addIntermediateX509Cert.Cert, intermediateCertificate.PemCert)
	require.Equal(t, constants.RootSubjectKeyId, intermediateCertificate.RootSubjectKeyId)

	// query leaf certificate
	leafCertificate, _ := querySingleCertificate(setup, constants.LeafSubject, constants.LeafSubjectKeyId)
	require.Equal(t, addLeafX509Cert.Cert, leafCertificate.PemCert)
	require.Equal(t, constants.RootSubjectKeyId, leafCertificate.RootSubjectKeyId)

	// check child certificates for leaf certificate
	leafCertChild := setup.PkiKeeper.GetChildCertificates(setup.Ctx, leafCertificate.Subject, leafCertificate.SubjectKeyId)

	require.Equal(t, 0, len(leafCertChild.ChildCertificates))

	// check child certificates for intermediate certificate
	intermediateCertChild := setup.PkiKeeper.GetChildCertificates(setup.Ctx, intermediateCertificate.Subject, intermediateCertificate.SubjectKeyId)

	require.Equal(t, 1, len(intermediateCertChild.ChildCertificates))
	require.Equal(t, types.NewCertificateIdentifier(constants.LeafSubject, constants.LeafSubjectKeyId), intermediateCertChild.ChildCertificates[0])

	// check child certificates for root certificate
	rootCertChild := setup.PkiKeeper.GetChildCertificates(setup.Ctx, rootCertificate.Subject, rootCertificate.SubjectKeyId)

	require.Equal(t, 1, len(rootCertChild.ChildCertificates))
	require.Equal(t, types.NewCertificateIdentifier(constants.IntermediateSubject, constants.IntermediateSubjectKeyId), rootCertChild.ChildCertificates[0])
}

func queryProposedCertificate(setup TestSetup, subject string, subjectKeyId string) (*types.ProposedCertificate, sdk.Error) {
	// query proposed certificate
	result, err := setup.Querier(
		setup.Ctx,
		[]string{keeper.QueryProposedX509RootCert, subject, subjectKeyId},
		abci.RequestQuery{},
	)

	if err != nil {
		return nil, err
	}

	var pendingCertificate types.ProposedCertificate
	_ = setup.Cdc.UnmarshalJSON(result, &pendingCertificate)

	return &pendingCertificate, nil
}

func querySingleCertificate(setup TestSetup, subject string, subjectKeyId string) (*types.Certificate, sdk.Error) {
	certificates, err := queryCertificates(setup, subject, subjectKeyId)
	if err != nil {
		return nil, err
	}

	if len(certificates.Items) > 1 {
		return nil, sdk.ErrInternal("More then 1 certificate returned")
	}

	return &certificates.Items[0], nil
}

func queryCertificates(setup TestSetup, subject string, subjectKeyId string) (types.Certificates, sdk.Error) {
	// query certificate
	result, err := setup.Querier(
		setup.Ctx,
		[]string{keeper.QueryX509Cert, subject, subjectKeyId},
		abci.RequestQuery{},
	)

	if err != nil {
		return types.Certificates{}, err
	}

	var certificate types.Certificates
	_ = setup.Cdc.UnmarshalJSON(result, &certificate)

	return certificate, nil
}

func rootCertificate(address sdk.AccAddress) types.Certificate {
	return types.NewRootCertificate(
		constants.RootCertPem,
		constants.RootSubject,
		constants.RootSubjectKeyId,
		constants.RootSerialNumber,
		address,
	)
}

func intermediateCertificate(address sdk.AccAddress) types.Certificate {
	return types.NewIntermediateCertificate(
		constants.IntermediateCertPem,
		constants.IntermediateSubject,
		constants.IntermediateSubjectKeyId,
		constants.IntermediateSerialNumber,
		constants.RootSubject,
		constants.RootSubjectKeyId,
		address,
	)
}