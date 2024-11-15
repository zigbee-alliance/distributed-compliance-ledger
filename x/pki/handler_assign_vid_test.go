package pki

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	dclauthtypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

// Main

func TestHandler_AssignVid_certificateWithoutSubjectVid(t *testing.T) {
	setup := Setup(t)

	vendorAcc := GenerateAccAddress()
	setup.AddAccount(vendorAcc, []dclauthtypes.AccountRole{dclauthtypes.VendorAdmin}, 0)

	// propose and approve x509 root certificate
	rootCertOptions := createTestRootCertOptions()
	rootCertOptions.vid = 0
	proposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertOptions)

	assignVid := types.MsgAssignVid{
		Signer:       vendorAcc.String(),
		Subject:      rootCertOptions.subject,
		SubjectKeyId: rootCertOptions.subjectKeyID,
		Vid:          testconstants.Vid,
	}

	_, err := setup.Handler(setup.Ctx, &assignVid)
	require.NoError(t, err)

	// DA certificates indexes checks

	// DaCertificates: Subject and SKID
	approvedCertificate, _ := querySingleApprovedCertificate(setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.Equal(t, testconstants.Vid, approvedCertificate.Vid)

	// DaCertificates: SKID
	certificateBySubjectKeyID, _ := queryAllApprovedCertificatesBySubjectKeyID(setup, testconstants.RootSubjectKeyID)
	require.Equal(t, 1, len(certificateBySubjectKeyID))
	require.Equal(t, 1, len(certificateBySubjectKeyID[0].Certs))
	require.Equal(t, testconstants.Vid, certificateBySubjectKeyID[0].Certs[0].Vid)

	// All certificates indexes checks

	// AllCertificate: Subject and SKID
	allCertificate, err := querySingleCertificateFromAllCertificatesIndex(setup, testconstants.RootSubject, testconstants.RootSubjectKeyID)
	require.NoError(t, err)
	require.Equal(t, testconstants.Vid, allCertificate.Vid)
}

func TestHandler_AssignVid_certificateWithSubjectVid(t *testing.T) {
	setup := Setup(t)

	vendorAcc := GenerateAccAddress()
	setup.AddAccount(vendorAcc, []dclauthtypes.AccountRole{dclauthtypes.VendorAdmin}, 0)

	// propose and approve x509 root certificate
	rootCertOptions := createPAACertWithNumericVidOptions()
	rootCertOptions.vid = 0
	proposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertOptions)

	assignVid := types.MsgAssignVid{
		Signer:       vendorAcc.String(),
		Subject:      rootCertOptions.subject,
		SubjectKeyId: rootCertOptions.subjectKeyID,
		Vid:          testconstants.PAACertWithNumericVidVid,
	}

	_, err := setup.Handler(setup.Ctx, &assignVid)
	require.NoError(t, err)

	// DA certificates indexes checks

	// DaCertificates: Subject and SKID
	approvedCertificate, _ := querySingleApprovedCertificate(setup, rootCertOptions.subject, rootCertOptions.subjectKeyID)
	require.Equal(t, testconstants.PAACertWithNumericVidVid, approvedCertificate.Vid)

	// DaCertificates: SKID
	certificateBySubjectKeyID, _ := queryAllApprovedCertificatesBySubjectKeyID(setup, rootCertOptions.subjectKeyID)
	require.Equal(t, 1, len(certificateBySubjectKeyID))
	require.Equal(t, 1, len(certificateBySubjectKeyID[0].Certs))
	require.Equal(t, testconstants.PAACertWithNumericVidVid, certificateBySubjectKeyID[0].Certs[0].Vid)

	// All certificates indexes checks

	// AllCertificate: Subject and SKID
	allCertificate, err := querySingleCertificateFromAllCertificatesIndex(setup, rootCertOptions.subject, rootCertOptions.subjectKeyID)
	require.NoError(t, err)
	require.Equal(t, testconstants.PAACertWithNumericVidVid, allCertificate.Vid)
}

// Extra cases

// Error cases

func TestHandler_AssignVid_SenderNotVendorAdmin(t *testing.T) {
	setup := Setup(t)

	assignVid := types.MsgAssignVid{
		Signer:       setup.Trustee1.String(),
		Subject:      testconstants.TestSubject,
		SubjectKeyId: testconstants.TestSubjectKeyID,
		Vid:          testconstants.TestCertPemVid,
	}

	_, err := setup.Handler(setup.Ctx, &assignVid)
	require.ErrorIs(t, err, sdkerrors.ErrUnauthorized)
}

func TestHandler_AssignVid_CertificateDoesNotExist(t *testing.T) {
	setup := Setup(t)

	vendorAcc := GenerateAccAddress()
	setup.AddAccount(vendorAcc, []dclauthtypes.AccountRole{dclauthtypes.VendorAdmin}, 0)

	assignVid := types.MsgAssignVid{
		Signer:       vendorAcc.String(),
		Subject:      testconstants.TestSubject,
		SubjectKeyId: testconstants.TestSubjectKeyID,
		Vid:          testconstants.TestCertPemVid,
	}

	_, err := setup.Handler(setup.Ctx, &assignVid)
	require.ErrorIs(t, err, pkitypes.ErrCertificateDoesNotExist)
}

func TestHandler_AssignVid_ForNonRootCertificate(t *testing.T) {
	setup := Setup(t)

	vendorAcc := GenerateAccAddress()
	setup.AddAccount(vendorAcc, []dclauthtypes.AccountRole{dclauthtypes.VendorAdmin}, 0)

	// propose and approve x509 root certificate
	rootCertOptions := createTestRootCertOptions()
	proposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertOptions)

	// Add vendor account
	vendorAccAddress := GenerateAccAddress()
	setup.AddAccount(vendorAccAddress, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testconstants.Vid)

	// add x509 intermediate certificate
	addX509Cert := types.NewMsgAddX509Cert(vendorAccAddress.String(), testconstants.IntermediateCertPem, testconstants.CertSchemaVersion)
	_, err := setup.Handler(setup.Ctx, addX509Cert)
	require.NoError(t, err)

	assignVid := types.MsgAssignVid{
		Signer:       vendorAcc.String(),
		Subject:      testconstants.IntermediateSubject,
		SubjectKeyId: testconstants.IntermediateSubjectKeyID,
		Vid:          testconstants.PAACertWithNumericVidVid,
	}

	_, err = setup.Handler(setup.Ctx, &assignVid)
	require.ErrorIs(t, err, pkitypes.ErrInappropriateCertificateType)
}

func TestHandler_AssignVid_CertificateAlreadyHasVid(t *testing.T) {
	setup := Setup(t)

	vendorAcc := GenerateAccAddress()
	setup.AddAccount(vendorAcc, []dclauthtypes.AccountRole{dclauthtypes.VendorAdmin}, 0)

	// propose and approve x509 root certificate
	rootCertOptions := createPAACertWithNumericVidOptions()
	proposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertOptions)

	assignVid := types.MsgAssignVid{
		Signer:       vendorAcc.String(),
		Subject:      rootCertOptions.subject,
		SubjectKeyId: rootCertOptions.subjectKeyID,
		Vid:          testconstants.PAACertWithNumericVidVid,
	}

	_, err := setup.Handler(setup.Ctx, &assignVid)
	require.ErrorIs(t, err, pkitypes.ErrNotEmptyVid)
}

func TestHandler_AssignVid_MessageVidAndCertificateVidNotEqual(t *testing.T) {
	setup := Setup(t)

	vendorAcc := GenerateAccAddress()
	setup.AddAccount(vendorAcc, []dclauthtypes.AccountRole{dclauthtypes.VendorAdmin}, 0)

	// propose and approve x509 root certificate
	rootCertOptions := createPAACertWithNumericVidOptions()
	rootCertOptions.vid = 0
	proposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertOptions)

	assignVid := types.MsgAssignVid{
		Signer:       vendorAcc.String(),
		Subject:      rootCertOptions.subject,
		SubjectKeyId: rootCertOptions.subjectKeyID,
		Vid:          1,
	}

	_, err := setup.Handler(setup.Ctx, &assignVid)
	require.ErrorIs(t, err, pkitypes.ErrCertificateVidNotEqualMsgVid)
}
