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

	vendorAcc := utils.GenerateAccAddress()
	setup.AddAccount(vendorAcc, []dclauthtypes.AccountRole{dclauthtypes.VendorAdmin}, 0)

	// propose and approve x509 root certificate
	rootCertificate := utils.CreateTestRootCert()
	rootCertOptions := utils.CreateTestRootCertOptions()
	rootCertOptions.Vid = 0
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertOptions)

	assignVid := types.MsgAssignVid{
		Signer:       vendorAcc.String(),
		Subject:      rootCertOptions.Subject,
		SubjectKeyId: rootCertOptions.SubjectKeyID,
		Vid:          testconstants.Vid,
	}

	_, err := setup.Handler(setup.Ctx, &assignVid)
	require.NoError(t, err)

	// DA certificates indexes checks
	// Check indexes
	indexes := []utils.TestIndex{
		{Key: types.ProposedCertificateKeyPrefix, Exist: false},
		{Key: types.UniqueCertificateKeyPrefix, Exist: true},
		{Key: types.AllCertificatesKeyPrefix, Exist: true},
		{Key: types.AllCertificatesBySubjectKeyPrefix, Exist: true},
		{Key: types.AllCertificatesBySubjectKeyIDKeyPrefix, Exist: true},
		{Key: types.ApprovedCertificatesKeyPrefix, Exist: true},
		{Key: types.ApprovedCertificatesBySubjectKeyPrefix, Exist: true},
		{Key: types.ApprovedCertificatesBySubjectKeyIDKeyPrefix, Exist: true},
		{Key: types.ApprovedRootCertificatesKeyPrefix, Exist: true},
	}
	resolvedCertificates := utils.CheckCertificateStateIndexes(t, setup, rootCertificate, indexes)

	// Check VID is assigned
	require.Equal(t, testconstants.Vid, resolvedCertificates.ApprovedCertificates.Certs[0].Vid)
	require.Equal(t, testconstants.Vid, resolvedCertificates.ApprovedCertificatesBySubjectKeyId[0].Certs[0].Vid)
	require.Equal(t, testconstants.Vid, resolvedCertificates.AllCertificates.Certs[0].Vid)
	require.Equal(t, testconstants.Vid, resolvedCertificates.AllCertificatesBySubjectKeyId[0].Certs[0].Vid)
}

func TestHandler_AssignVid_certificateWithSubjectVid(t *testing.T) {
	setup := utils.Setup(t)

	vendorAcc := utils.GenerateAccAddress()
	setup.AddAccount(vendorAcc, []dclauthtypes.AccountRole{dclauthtypes.VendorAdmin}, 0)

	// propose and approve x509 root certificate
	rootCertOptions := utils.CreatePAACertWithNumericVidOptions()
	rootCertOptions.Vid = 0
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertOptions)

	assignVid := types.MsgAssignVid{
		Signer:       vendorAcc.String(),
		Subject:      rootCertOptions.Subject,
		SubjectKeyId: rootCertOptions.SubjectKeyID,
		Vid:          testconstants.PAACertWithNumericVidVid,
	}

	_, err := setup.Handler(setup.Ctx, &assignVid)
	require.NoError(t, err)

	// DA certificates indexes checks

	// DaCertificates: Subject and SKID
	approvedCertificate, _ := utils.QueryApprovedCertificates(setup, rootCertOptions.Subject, rootCertOptions.SubjectKeyID)
	require.Equal(t, testconstants.PAACertWithNumericVidVid, approvedCertificate.Certs[0].Vid)

	// DaCertificates: SKID
	certificateBySubjectKeyID, _ := utils.QueryApprovedCertificatesBySubjectKeyID(setup, rootCertOptions.SubjectKeyID)
	require.Equal(t, 1, len(certificateBySubjectKeyID))
	require.Equal(t, 1, len(certificateBySubjectKeyID[0].Certs))
	require.Equal(t, testconstants.PAACertWithNumericVidVid, certificateBySubjectKeyID[0].Certs[0].Vid)

	// All certificates indexes checks

	// AllCertificates: Subject and SKID
	allCertificate, err := utils.QueryAllCertificates(setup, rootCertOptions.Subject, rootCertOptions.SubjectKeyID)
	require.NoError(t, err)
	require.Equal(t, testconstants.PAACertWithNumericVidVid, allCertificate.Certs[0].Vid)

	// AllCertificates: SKID
	allCertificateBySkid, err := utils.QueryAllCertificatesBySubjectKeyID(setup, rootCertOptions.SubjectKeyID)
	require.NoError(t, err)
	require.Equal(t, testconstants.PAACertWithNumericVidVid, allCertificateBySkid[0].Certs[0].Vid)
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

	vendorAcc := utils.GenerateAccAddress()
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
	setup := utils.Setup(t)

	vendorAcc := utils.GenerateAccAddress()
	setup.AddAccount(vendorAcc, []dclauthtypes.AccountRole{dclauthtypes.VendorAdmin}, 0)

	// propose and approve x509 root certificate
	rootCertOptions := utils.CreateTestRootCertOptions()
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertOptions)

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

	vendorAcc := utils.GenerateAccAddress()
	setup.AddAccount(vendorAcc, []dclauthtypes.AccountRole{dclauthtypes.VendorAdmin}, 0)

	// propose and approve x509 root certificate
	rootCertOptions := utils.CreatePAACertWithNumericVidOptions()
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertOptions)

	assignVid := types.MsgAssignVid{
		Signer:       vendorAcc.String(),
		Subject:      rootCertOptions.Subject,
		SubjectKeyId: rootCertOptions.SubjectKeyID,
		Vid:          testconstants.PAACertWithNumericVidVid,
	}

	_, err := setup.Handler(setup.Ctx, &assignVid)
	require.ErrorIs(t, err, pkitypes.ErrNotEmptyVid)
}

func TestHandler_AssignVid_MessageVidAndCertificateVidNotEqual(t *testing.T) {
	setup := utils.Setup(t)

	vendorAcc := utils.GenerateAccAddress()
	setup.AddAccount(vendorAcc, []dclauthtypes.AccountRole{dclauthtypes.VendorAdmin}, 0)

	// propose and approve x509 root certificate
	rootCertOptions := utils.CreatePAACertWithNumericVidOptions()
	rootCertOptions.Vid = 0
	utils.ProposeAndApproveRootCertificate(setup, setup.Trustee1, rootCertOptions)

	assignVid := types.MsgAssignVid{
		Signer:       vendorAcc.String(),
		Subject:      rootCertOptions.Subject,
		SubjectKeyId: rootCertOptions.SubjectKeyID,
		Vid:          1,
	}

	_, err := setup.Handler(setup.Ctx, &assignVid)
	require.ErrorIs(t, err, pkitypes.ErrCertificateVidNotEqualMsgVid)
}
