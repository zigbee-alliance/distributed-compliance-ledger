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

package tests

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
		SchemaVersion:        0,
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
		SchemaVersion:        0,
	}
}

func createAddRevocationMessageWithPAICertNoVid(signer string) *types.MsgAddPkiRevocationDistributionPoint {
	return &types.MsgAddPkiRevocationDistributionPoint{
		Signer:               signer,
		Vid:                  testconstants.RootCertWithVidVid,
		IsPAA:                false,
		Pid:                  0,
		CrlSignerCertificate: testconstants.IntermediateCertPem,
		Label:                label,
		DataURL:              testconstants.DataURL,
		IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
		RevocationType:       types.CRLRevocationType,
		SchemaVersion:        0,
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
		SchemaVersion:        0,
	}
}

func createAddRevocationMessageWithPAICertWithVid(signer string) *types.MsgAddPkiRevocationDistributionPoint {
	return &types.MsgAddPkiRevocationDistributionPoint{
		Signer:               signer,
		Vid:                  testconstants.PAICertWithVidVid,
		IsPAA:                false,
		CrlSignerCertificate: testconstants.PAICertWithVid,
		Label:                label,
		DataURL:              testconstants.DataURL,
		IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
		RevocationType:       types.CRLRevocationType,
		SchemaVersion:        0,
	}
}

func createAddRevocationMessageWithPAACertNoVid(signer string, vid int32) *types.MsgAddPkiRevocationDistributionPoint {
	return &types.MsgAddPkiRevocationDistributionPoint{
		Signer:               signer,
		Vid:                  vid,
		IsPAA:                true,
		Pid:                  0,
		CrlSignerCertificate: testconstants.PAACertNoVid,
		Label:                label,
		DataURL:              testconstants.DataURL,
		IssuerSubjectKeyID:   testconstants.SubjectKeyIDWithoutColons,
		RevocationType:       types.CRLRevocationType,
		SchemaVersion:        0,
	}
}

func createAddRevocationMessageWithLeafCertWithVid(signer string) *types.MsgAddPkiRevocationDistributionPoint {
	return &types.MsgAddPkiRevocationDistributionPoint{
		Signer:               signer,
		Vid:                  testconstants.LeafCertWithVidVid,
		IsPAA:                false,
		CrlSignerCertificate: testconstants.LeafCertWithVid,
		CrlSignerDelegator:   testconstants.IntermediateCertWithVid1,
		Label:                label,
		DataURL:              testconstants.DataURL,
		IssuerSubjectKeyID:   testconstants.IntermediateCertWithVid1SubjectKeyIDWithoutColumns,
		RevocationType:       types.CRLRevocationType,
		SchemaVersion:        0,
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
