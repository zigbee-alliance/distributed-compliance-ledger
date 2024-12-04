package tests

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	dclauthtypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/tests/utils"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

// Main

func TestHandler_AssignVid_certificateWithoutSubjectVid(t *testing.T) {
	setup := utils.Setup(t)

	vendorAcc := setup.CreateVendorAdminAccount(0)

	// propose and approve x509 root certificate
	rootCertificate := utils.CreateTestRootCert()
	rootCertificate.VID = 0
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, &rootCertificate)

	// assign Vid
	utils.AssignCertificateVid(setup, vendorAcc, rootCertificate.Subject, rootCertificate.SubjectKeyID, testconstants.Vid)

	// Check state indexes
	indexes := utils.TestIndexes{
		Present: []utils.TestIndex{
			{Key: types.UniqueCertificateKeyPrefix},
			{Key: types.AllCertificatesKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.ApprovedCertificatesKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.ApprovedRootCertificatesKeyPrefix},
		},
		Missing: []utils.TestIndex{
			{Key: types.ProposedCertificateKeyPrefix},
		},
	}
	resolvedCertificates := utils.CheckCertificateStateIndexes(t, setup, rootCertificate, indexes)

	// Check VID is assigned
	require.Equal(t, testconstants.Vid, resolvedCertificates.ApprovedCertificates.Certs[0].Vid)
	require.Equal(t, testconstants.Vid, resolvedCertificates.ApprovedCertificatesBySubjectKeyID[0].Certs[0].Vid)
	require.Equal(t, testconstants.Vid, resolvedCertificates.AllCertificates.Certs[0].Vid)
	require.Equal(t, testconstants.Vid, resolvedCertificates.AllCertificatesBySubjectKeyID[0].Certs[0].Vid)
}

func TestHandler_AssignVid_certificateWithSubjectVid(t *testing.T) {
	setup := utils.Setup(t)

	vendorAcc := setup.CreateVendorAdminAccount(0)

	// propose and approve x509 root certificate
	rootCertificate := utils.CreateTestPAACertWithNumericVid()
	rootCertificate.VID = 0
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, &rootCertificate)

	// assign Vid
	utils.AssignCertificateVid(setup, vendorAcc, rootCertificate.Subject, rootCertificate.SubjectKeyID, testconstants.PAACertWithNumericVidVid)

	// Check state indexes
	indexes := utils.TestIndexes{
		Present: []utils.TestIndex{
			{Key: types.AllCertificatesKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyPrefix},
			{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.ApprovedCertificatesKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyPrefix},
			{Key: types.ApprovedCertificatesBySubjectKeyIDKeyPrefix},
			{Key: types.ApprovedRootCertificatesKeyPrefix},
		},
		Missing: []utils.TestIndex{
			{Key: types.ProposedCertificateKeyPrefix},
		},
	}
	resolvedCertificates := utils.CheckCertificateStateIndexes(t, setup, rootCertificate, indexes)

	// Check VID is assigned
	require.Equal(t, testconstants.PAACertWithNumericVidVid, resolvedCertificates.ApprovedCertificates.Certs[0].Vid)
	require.Equal(t, testconstants.PAACertWithNumericVidVid, resolvedCertificates.ApprovedCertificatesBySubjectKeyID[0].Certs[0].Vid)
	require.Equal(t, testconstants.PAACertWithNumericVidVid, resolvedCertificates.AllCertificates.Certs[0].Vid)
	require.Equal(t, testconstants.PAACertWithNumericVidVid, resolvedCertificates.AllCertificatesBySubjectKeyID[0].Certs[0].Vid)
}

// Extra cases

// Error cases

func TestHandler_AssignVid_SenderNotVendorAdmin(t *testing.T) {
	setup := utils.Setup(t)

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
	setup := utils.Setup(t)

	vendorAcc := setup.CreateVendorAdminAccount(0)

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
	setup := utils.Setup(t)

	vendorAcc := setup.CreateVendorAdminAccount(0)

	// propose and approve x509 root certificate
	rootCert := utils.CreateTestRootCert()
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, &rootCert)

	// Add vendor account
	vendorAccAddress := utils.GenerateAccAddress()
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
	setup := utils.Setup(t)

	vendorAcc := setup.CreateVendorAdminAccount(0)

	// propose and approve x509 root certificate
	rootCert := utils.CreateTestPAACertWithNumericVid()
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, &rootCert)

	assignVid := types.MsgAssignVid{
		Signer:       vendorAcc.String(),
		Subject:      rootCert.Subject,
		SubjectKeyId: rootCert.SubjectKeyID,
		Vid:          testconstants.PAACertWithNumericVidVid,
	}

	_, err := setup.Handler(setup.Ctx, &assignVid)
	require.ErrorIs(t, err, pkitypes.ErrNotEmptyVid)
}

func TestHandler_AssignVid_MessageVidAndCertificateVidNotEqual(t *testing.T) {
	setup := utils.Setup(t)

	vendorAcc := setup.CreateVendorAdminAccount(0)

	// propose and approve x509 root certificate
	rootCert := utils.CreateTestPAACertWithNumericVid()
	rootCert.VID = 0
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, &rootCert)

	assignVid := types.MsgAssignVid{
		Signer:       vendorAcc.String(),
		Subject:      rootCert.Subject,
		SubjectKeyId: rootCert.SubjectKeyID,
		Vid:          1,
	}

	_, err := setup.Handler(setup.Ctx, &assignVid)
	require.ErrorIs(t, err, pkitypes.ErrCertificateVidNotEqualMsgVid)
}
