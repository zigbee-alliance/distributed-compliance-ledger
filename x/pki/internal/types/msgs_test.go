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
package types

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
)

/*
	MsgProposeAddX509RootCert
*/

func TestNewMsgProposeAddX509RootCert(t *testing.T) {
	msg := NewMsgProposeAddX509RootCert(testconstants.RootCertPem, testconstants.Signer)

	require.Equal(t, msg.Route(), RouterKey)
	require.Equal(t, msg.Type(), "propose_add_x509_root_cert")
	require.Equal(t, msg.GetSigners(), []sdk.AccAddress{testconstants.Signer})
}

func TestValidateMsgProposeAddX509RootCert(t *testing.T) {
	cases := []struct {
		valid bool
		msg   MsgProposeAddX509RootCert
	}{
		{true, NewMsgProposeAddX509RootCert(
			testconstants.RootCertPem, testconstants.Signer)},
		{false, NewMsgProposeAddX509RootCert(
			"", testconstants.Signer)},
		{false, NewMsgProposeAddX509RootCert(
			testconstants.RootCertPem, nil)},
	}

	for _, tc := range cases {
		err := tc.msg.ValidateBasic()

		if tc.valid {
			require.Nil(t, err)
		} else {
			require.NotNil(t, err)
		}
	}
}

func TestMsgProposeAddX509RootCertGetSignBytes(t *testing.T) {
	msg := NewMsgProposeAddX509RootCert(testconstants.StubCertPem, testconstants.Signer)
	res := msg.GetSignBytes()

	expected := `{"type":"pki/ProposeAddX509RootCert","value":{` +
		`"cert":"pem certificate",` +
		`"signer":"cosmos1p72j8mgkf39qjzcmr283w8l8y9qv30qpj056uz"}}`
	require.Equal(t, expected, string(res))
}

/*
	MsgApproveAddX509RootCert
*/

func TestNewMsgApproveAddX509RootCert(t *testing.T) {
	msg := NewMsgApproveAddX509RootCert(testconstants.RootSubject,
		testconstants.RootSubjectKeyID, testconstants.Signer)

	require.Equal(t, msg.Route(), RouterKey)
	require.Equal(t, msg.Type(), "approve_add_x509_root_cert")
	require.Equal(t, msg.GetSigners(), []sdk.AccAddress{testconstants.Signer})
}

func TestValidateMsgApproveAddX509RootCert(t *testing.T) {
	cases := []struct {
		valid bool
		msg   MsgApproveAddX509RootCert
	}{
		{true, NewMsgApproveAddX509RootCert(
			testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Signer)},
		{false, NewMsgApproveAddX509RootCert(
			"", testconstants.RootSubjectKeyID, testconstants.Signer)},
		{false, NewMsgApproveAddX509RootCert(
			testconstants.RootSubject, "", testconstants.Signer)},
		{false, NewMsgApproveAddX509RootCert(
			testconstants.RootSubject, testconstants.RootSubjectKeyID, nil)},
	}

	for _, tc := range cases {
		err := tc.msg.ValidateBasic()

		if tc.valid {
			require.Nil(t, err)
		} else {
			require.NotNil(t, err)
		}
	}
}

func TestMsgApproveAddX509RootCertGetSignBytes(t *testing.T) {
	msg := NewMsgApproveAddX509RootCert(testconstants.RootSubject,
		testconstants.RootSubjectKeyID, testconstants.Signer)

	expected := `{"type":"pki/ApproveAddX509RootCert","value":{` +
		`"signer":"cosmos1p72j8mgkf39qjzcmr283w8l8y9qv30qpj056uz",` +
		`"subject":"CN=DST Root CA X3,O=Digital Signature Trust Co.",` +
		`"subject_key_id":"C4:A7:B1:A4:7B:2C:71:FA:DB:E1:4B:90:75:FF:C4:15:60:85:89:10"}}`
	require.Equal(t, expected, string(msg.GetSignBytes()))
}

/*
	MsgAddX509Cert
*/

func TestNewMsgAddX509Cert(t *testing.T) {
	msg := NewMsgAddX509Cert(testconstants.LeafCertPem, testconstants.Signer)

	require.Equal(t, msg.Route(), RouterKey)
	require.Equal(t, msg.Type(), "add_x509_cert")
	require.Equal(t, msg.GetSigners(), []sdk.AccAddress{testconstants.Signer})
}

func TestValidateMsgAddX509Cert(t *testing.T) {
	cases := []struct {
		valid bool
		msg   MsgAddX509Cert
	}{
		{true, NewMsgAddX509Cert(
			testconstants.LeafCertPem, testconstants.Signer)},
		{false, NewMsgAddX509Cert(
			"", testconstants.Signer)},
		{false, NewMsgAddX509Cert(
			testconstants.LeafCertPem, nil)},
	}

	for _, tc := range cases {
		err := tc.msg.ValidateBasic()

		if tc.valid {
			require.Nil(t, err)
		} else {
			require.NotNil(t, err)
		}
	}
}

func TestMsgAddX509CertGetSignBytes(t *testing.T) {
	msg := NewMsgAddX509Cert(testconstants.StubCertPem, testconstants.Signer)
	res := msg.GetSignBytes()

	expected := `{"type":"pki/AddX509Cert","value":{` +
		`"cert":"pem certificate",` +
		`"signer":"cosmos1p72j8mgkf39qjzcmr283w8l8y9qv30qpj056uz"}}`
	require.Equal(t, expected, string(res))
}

/*
	MsgProposeRevokeX509RootCert
*/

func TestNewMsgProposeRevokeX509RootCert(t *testing.T) {
	msg := NewMsgProposeRevokeX509RootCert(testconstants.RootSubject,
		testconstants.RootSubjectKeyID, testconstants.Signer)

	require.Equal(t, msg.Route(), RouterKey)
	require.Equal(t, msg.Type(), "propose_revoke_x509_root_cert")
	require.Equal(t, msg.GetSigners(), []sdk.AccAddress{testconstants.Signer})
}

func TestValidateMsgProposeRevokeX509RootCert(t *testing.T) {
	cases := []struct {
		valid bool
		msg   MsgProposeRevokeX509RootCert
	}{
		{true, NewMsgProposeRevokeX509RootCert(
			testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Signer)},
		{false, NewMsgProposeRevokeX509RootCert(
			"", testconstants.RootSubjectKeyID, testconstants.Signer)},
		{false, NewMsgProposeRevokeX509RootCert(
			testconstants.RootSubject, "", testconstants.Signer)},
		{false, NewMsgProposeRevokeX509RootCert(
			testconstants.RootSubject, testconstants.RootSubjectKeyID, nil)},
	}

	for _, tc := range cases {
		err := tc.msg.ValidateBasic()

		if tc.valid {
			require.Nil(t, err)
		} else {
			require.NotNil(t, err)
		}
	}
}

func TestMsgProposeRevokeX509RootCertGetSignBytes(t *testing.T) {
	msg := NewMsgProposeRevokeX509RootCert(testconstants.RootSubject,
		testconstants.RootSubjectKeyID, testconstants.Signer)

	expected := `{"type":"pki/ProposeRevokeX509RootCert","value":{` +
		`"signer":"cosmos1p72j8mgkf39qjzcmr283w8l8y9qv30qpj056uz",` +
		`"subject":"CN=DST Root CA X3,O=Digital Signature Trust Co.",` +
		`"subject_key_id":"C4:A7:B1:A4:7B:2C:71:FA:DB:E1:4B:90:75:FF:C4:15:60:85:89:10"}}`
	require.Equal(t, expected, string(msg.GetSignBytes()))
}

/*
	MsgApproveRevokeX509RootCert
*/

func TestNewMsgApproveRevokeX509RootCert(t *testing.T) {
	msg := NewMsgApproveRevokeX509RootCert(testconstants.RootSubject,
		testconstants.RootSubjectKeyID, testconstants.Signer)

	require.Equal(t, msg.Route(), RouterKey)
	require.Equal(t, msg.Type(), "approve_revoke_x509_root_cert")
	require.Equal(t, msg.GetSigners(), []sdk.AccAddress{testconstants.Signer})
}

func TestValidateMsgApproveRevokeX509RootCert(t *testing.T) {
	cases := []struct {
		valid bool
		msg   MsgApproveRevokeX509RootCert
	}{
		{true, NewMsgApproveRevokeX509RootCert(
			testconstants.RootSubject, testconstants.RootSubjectKeyID, testconstants.Signer)},
		{false, NewMsgApproveRevokeX509RootCert(
			"", testconstants.RootSubjectKeyID, testconstants.Signer)},
		{false, NewMsgApproveRevokeX509RootCert(
			testconstants.RootSubject, "", testconstants.Signer)},
		{false, NewMsgApproveRevokeX509RootCert(
			testconstants.RootSubject, testconstants.RootSubjectKeyID, nil)},
	}

	for _, tc := range cases {
		err := tc.msg.ValidateBasic()

		if tc.valid {
			require.Nil(t, err)
		} else {
			require.NotNil(t, err)
		}
	}
}

func TestMsgApproveRevokeX509RootCertGetSignBytes(t *testing.T) {
	msg := NewMsgApproveRevokeX509RootCert(testconstants.RootSubject,
		testconstants.RootSubjectKeyID, testconstants.Signer)

	expected := `{"type":"pki/ApproveRevokeX509RootCert","value":{` +
		`"signer":"cosmos1p72j8mgkf39qjzcmr283w8l8y9qv30qpj056uz",` +
		`"subject":"CN=DST Root CA X3,O=Digital Signature Trust Co.",` +
		`"subject_key_id":"C4:A7:B1:A4:7B:2C:71:FA:DB:E1:4B:90:75:FF:C4:15:60:85:89:10"}}`
	require.Equal(t, expected, string(msg.GetSignBytes()))
}

/*
	MsgRevokeX509Cert
*/

func TestNewMsgRevokeX509Cert(t *testing.T) {
	msg := NewMsgRevokeX509Cert(testconstants.LeafSubject, testconstants.LeafSubjectKeyID, testconstants.Signer)

	require.Equal(t, RouterKey, msg.Route())
	require.Equal(t, "revoke_x509_cert", msg.Type())
	require.Equal(t, []sdk.AccAddress{testconstants.Signer}, msg.GetSigners())
}

func TestValidateMsgRevokeX509Cert(t *testing.T) {
	cases := []struct {
		valid bool
		msg   MsgRevokeX509Cert
	}{
		{true, NewMsgRevokeX509Cert(testconstants.LeafSubject, testconstants.LeafSubjectKeyID, testconstants.Signer)},
		{false, NewMsgRevokeX509Cert("", testconstants.LeafSubjectKeyID, testconstants.Signer)},
		{false, NewMsgRevokeX509Cert(testconstants.LeafSubject, "", testconstants.Signer)},
		{false, NewMsgRevokeX509Cert(testconstants.LeafSubject, testconstants.LeafSubjectKeyID, nil)},
	}

	for _, tc := range cases {
		err := tc.msg.ValidateBasic()

		if tc.valid {
			require.Nil(t, err)
		} else {
			require.NotNil(t, err)
		}
	}
}

func TestMsgRevokeX509CertGetSignBytes(t *testing.T) {
	msg := NewMsgRevokeX509Cert(testconstants.LeafSubject, testconstants.LeafSubjectKeyID, testconstants.Signer)
	res := msg.GetSignBytes()

	expected := `{"type":"pki/RevokeX509Cert","value":{` +
		`"signer":"cosmos1p72j8mgkf39qjzcmr283w8l8y9qv30qpj056uz",` +
		`"subject":"CN=dsr-corporation.com",` +
		`"subject_key_id":"8A:E9:AC:D4:16:81:2F:87:66:8E:61:BE:A9:C5:1C:0:1B:F7:BB:AE"}}`
	require.Equal(t, expected, string(res))
}
