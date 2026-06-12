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
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/tests/utils"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
	dclx509 "github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/x509"
)

// These tests submit the Matter R1.6 specification's example certificates
// (§6.2.2 DAC / PAI / PAA examples and the §6.5 RCAC / ICAC / NOC examples)
// through the production add-handlers, exercising the full chain: structural
// profile checks, signature verification, ledger storage, and VID scoping.

func parseSpecCert(t *testing.T, pem string) *dclx509.Certificate {
	t.Helper()
	cert, err := dclx509.ParseAndValidateCertificate(pem)
	require.NoError(t, err)

	return cert
}

func TestHandler_ProposeAddDaRootCert_MatterSpecExample_PAA(t *testing.T) {
	setup := utils.Setup(t)

	msg := types.NewMsgProposeAddX509RootCert(
		setup.Trustee1.String(),
		testconstants.MatterSpecPAA,
		testconstants.Info,
		testconstants.MatterSpecPAAVid,
		testconstants.CertSchemaVersion,
	)
	_, err := setup.Handler(setup.Ctx, msg)
	require.NoError(t, err)
}

func TestHandler_AddDaIntermediateCert_MatterSpecExample_PAI(t *testing.T) {
	setup := utils.Setup(t)

	// Pre-seed the spec PAA so the spec PAI has a parent to chain against.
	specPAA := parseSpecCert(t, testconstants.MatterSpecPAA)
	proposeMsg := types.NewMsgProposeAddX509RootCert(
		setup.Trustee1.String(),
		testconstants.MatterSpecPAA,
		testconstants.Info,
		testconstants.MatterSpecPAAVid,
		testconstants.CertSchemaVersion,
	)
	_, err := setup.Handler(setup.Ctx, proposeMsg)
	require.NoError(t, err)

	approveMsg := types.NewMsgApproveAddX509RootCert(
		setup.Trustee2.String(),
		specPAA.Subject,
		specPAA.SubjectKeyID,
		testconstants.Info,
	)
	_, err = setup.Handler(setup.Ctx, approveMsg)
	require.NoError(t, err)

	// Vendor account scoped to the spec VID (0xFFF1) submits the PAI.
	vendor := setup.CreateVendorAccount(testconstants.MatterSpecPAAVid)
	addPAI := types.NewMsgAddX509Cert(vendor.String(), testconstants.MatterSpecPAI, testconstants.CertSchemaVersion)
	_, err = setup.Handler(setup.Ctx, addPAI)
	require.NoError(t, err)
}

func TestHandler_AddDaIntermediateCert_MatterSpecExample_DAC(t *testing.T) {
	setup := utils.Setup(t)

	// Pre-seed the spec PAA + PAI so the spec DAC has a chain to follow.
	specPAA := parseSpecCert(t, testconstants.MatterSpecPAA)
	proposeMsg := types.NewMsgProposeAddX509RootCert(
		setup.Trustee1.String(),
		testconstants.MatterSpecPAA,
		testconstants.Info,
		testconstants.MatterSpecPAAVid,
		testconstants.CertSchemaVersion,
	)
	_, err := setup.Handler(setup.Ctx, proposeMsg)
	require.NoError(t, err)

	approveMsg := types.NewMsgApproveAddX509RootCert(
		setup.Trustee2.String(),
		specPAA.Subject,
		specPAA.SubjectKeyID,
		testconstants.Info,
	)
	_, err = setup.Handler(setup.Ctx, approveMsg)
	require.NoError(t, err)

	vendor := setup.CreateVendorAccount(testconstants.MatterSpecPAAVid)
	addPAI := types.NewMsgAddX509Cert(vendor.String(), testconstants.MatterSpecPAI, testconstants.CertSchemaVersion)
	_, err = setup.Handler(setup.Ctx, addPAI)
	require.NoError(t, err)

	// Same VID-scoped vendor submits the DAC (cA=FALSE, dispatched onto the
	// DAC profile branch of VerifyDAChainNonRoot).
	addDAC := types.NewMsgAddX509Cert(vendor.String(), testconstants.MatterSpecDAC, testconstants.CertSchemaVersion)
	_, err = setup.Handler(setup.Ctx, addDAC)
	require.NoError(t, err)
}

func TestHandler_AddNocRootCert_MatterSpecExample_RCAC(t *testing.T) {
	setup := utils.Setup(t)

	msg := types.NewMsgAddNocX509RootCert(
		setup.Vendor1.String(),
		testconstants.MatterSpecRCAC,
		testconstants.CertSchemaVersion,
		false,
	)
	_, err := setup.Handler(setup.Ctx, msg)
	require.NoError(t, err)
}

func TestHandler_AddNocIcaCert_MatterSpecExample_ICAC(t *testing.T) {
	setup := utils.Setup(t)

	// Pre-seed the spec RCAC under Vendor1 so the spec ICAC chains under it.
	addRCAC := types.NewMsgAddNocX509RootCert(
		setup.Vendor1.String(),
		testconstants.MatterSpecRCAC,
		testconstants.CertSchemaVersion,
		false,
	)
	_, err := setup.Handler(setup.Ctx, addRCAC)
	require.NoError(t, err)

	addICAC := types.NewMsgAddNocX509IcaCert(
		setup.Vendor1.String(),
		testconstants.MatterSpecICAC,
		testconstants.CertSchemaVersion,
		false,
	)
	_, err = setup.Handler(setup.Ctx, addICAC)
	require.NoError(t, err)
}

func TestHandler_AddNocIcaCert_MatterSpecExample_NOC(t *testing.T) {
	setup := utils.Setup(t)

	addRCAC := types.NewMsgAddNocX509RootCert(
		setup.Vendor1.String(),
		testconstants.MatterSpecRCAC,
		testconstants.CertSchemaVersion,
		false,
	)
	_, err := setup.Handler(setup.Ctx, addRCAC)
	require.NoError(t, err)

	addICAC := types.NewMsgAddNocX509IcaCert(
		setup.Vendor1.String(),
		testconstants.MatterSpecICAC,
		testconstants.CertSchemaVersion,
		false,
	)
	_, err = setup.Handler(setup.Ctx, addICAC)
	require.NoError(t, err)

	// NOC end-entity (cA=FALSE) is dispatched to the §6.5.12 NOC profile branch
	// of VerifyNOCChainNonRoot.
	addNOC := types.NewMsgAddNocX509IcaCert(
		setup.Vendor1.String(),
		testconstants.MatterSpecNOC,
		testconstants.CertSchemaVersion,
		false,
	)
	_, err = setup.Handler(setup.Ctx, addNOC)
	require.NoError(t, err)
}
