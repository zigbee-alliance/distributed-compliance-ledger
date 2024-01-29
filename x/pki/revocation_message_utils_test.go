package pki

import (
	"testing"

	"github.com/stretchr/testify/require"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

const label = "label"

func createAddRevocationMessageWithPAACertWithNumericVid(signer string) *types.MsgAddPkiRevocationDistributionPoint {
	return &types.MsgAddPkiRevocationDistributionPoint{
		Signer:               signer,
		Vid:                  testconstants.PAACertWithNumericVidVid,
		IsPAA:                true,
		Pid:                  0,
		CrlSignerCertificate: testconstants.PAACertWithNumericVid,
		Label:                label,
		DataURL:              testconstants.DataURL,
		IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
		RevocationType:       types.CRLRevocationType,
	}
}

func createAddRevocationMessageWithPAICertWithNumericVidPid(signer string) *types.MsgAddPkiRevocationDistributionPoint {
	return &types.MsgAddPkiRevocationDistributionPoint{
		Signer:               signer,
		Vid:                  testconstants.PAICertWithNumericPidVidVid,
		IsPAA:                false,
		Pid:                  testconstants.PAICertWithNumericPidVidPid,
		CrlSignerCertificate: testconstants.PAICertWithNumericPidVid,
		Label:                label,
		DataURL:              testconstants.DataURL,
		IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
		RevocationType:       types.CRLRevocationType,
	}
}

func createAddRevocationMessageWithPAICertWithVidPid(signer string) *types.MsgAddPkiRevocationDistributionPoint {
	return &types.MsgAddPkiRevocationDistributionPoint{
		Signer:               signer,
		Vid:                  testconstants.PAICertWithPidVidVid,
		IsPAA:                false,
		Pid:                  testconstants.PAICertWithPidVidPid,
		CrlSignerCertificate: testconstants.PAICertWithPidVid,
		Label:                label,
		DataURL:              testconstants.DataURL,
		IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
		RevocationType:       types.CRLRevocationType,
	}
}

func createAddRevocationMessageWithPAACertNoVid(signer string, vid int32) *types.MsgAddPkiRevocationDistributionPoint {
	return &types.MsgAddPkiRevocationDistributionPoint{
		Signer:               signer,
		Vid:                  vid,
		IsPAA:                false,
		Pid:                  0,
		CrlSignerCertificate: testconstants.PAACertNoVid,
		Label:                label,
		DataURL:              testconstants.DataURL,
		IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
		RevocationType:       types.CRLRevocationType,
	}
}

func assertRevocationPointEqual(t *testing.T, expected *types.MsgAddPkiRevocationDistributionPoint, actual *types.PkiRevocationDistributionPoint) {
	t.Helper()
	require.Equal(t, expected.CrlSignerCertificate, actual.CrlSignerCertificate)
	require.Equal(t, expected.CrlSignerCertificate, actual.CrlSignerCertificate)
	require.Equal(t, expected.DataDigest, actual.DataDigest)
	require.Equal(t, expected.DataDigestType, actual.DataDigestType)
	require.Equal(t, expected.DataFileSize, actual.DataFileSize)
	require.Equal(t, expected.DataURL, actual.DataURL)
	require.Equal(t, expected.IsPAA, actual.IsPAA)
	require.Equal(t, expected.IssuerSubjectKeyID, actual.IssuerSubjectKeyID)
	require.Equal(t, expected.Label, actual.Label)
	require.Equal(t, expected.Pid, actual.Pid)
	require.Equal(t, expected.RevocationType, actual.RevocationType)
	require.Equal(t, expected.Vid, actual.Vid)
}
