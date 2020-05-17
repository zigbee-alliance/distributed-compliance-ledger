//nolint:testpackage
package types

// nolint:goimports
import (
	"git.dsr-corporation.com/zb-ledger/zb-ledger/integration_tests/constants"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"testing"
)

/*
	MsgProposeAddX509RootCert
*/

func TestNewMsgProposeAddX509RootCert(t *testing.T) {
	var msg = NewMsgProposeAddX509RootCert(testconstants.RootCertPem, testconstants.Signer)

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
	var msg = NewMsgProposeAddX509RootCert(testconstants.StubCert, testconstants.Signer)
	res := msg.GetSignBytes()

	expected := `{"type":"pki/ProposeAddX509RootCert","value":{"cert":` +
		`"pem certificate","signer":"cosmos1p72j8mgkf39qjzcmr283w8l8y9qv30qpj056uz"}}`
	require.Equal(t, expected, string(res))
}

/*
	MsgApproveAddX509RootCert(
*/

func TestNewMsgApproveAddX509RootCert(t *testing.T) {
	var msg = NewMsgApproveAddX509RootCert(testconstants.LeafSubject,
		testconstants.LeafSubjectKeyID, testconstants.Signer)

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
			testconstants.LeafSubject, testconstants.LeafSubjectKeyID, testconstants.Signer)},
		{false, NewMsgApproveAddX509RootCert(
			"", testconstants.LeafSubjectKeyID, testconstants.Signer)},
		{false, NewMsgApproveAddX509RootCert(
			testconstants.LeafSubject, "", testconstants.Signer)},
		{false, NewMsgApproveAddX509RootCert(
			testconstants.LeafSubject, testconstants.LeafSubjectKeyID, nil)},
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
	var msg = NewMsgApproveAddX509RootCert(testconstants.LeafSubject,
		testconstants.LeafSubjectKeyID, testconstants.Signer)

	expected := `{"type":"pki/ApproveAddX509RootCert","value":{"signer":` +
		`"cosmos1p72j8mgkf39qjzcmr283w8l8y9qv30qpj056uz",` +
		`"subject":"CN=dsr-corporation.com","subject_key_id":` +
		`"8A:34:B:5C:D8:42:18:F2:C1:2A:AC:7A:B3:8F:6E:90:66:F4:4E:5C"}}`
	require.Equal(t, expected, string(msg.GetSignBytes()))
}

/*
	MsgAddX509Cert
*/

func TestNewMsgAddX509Cert(t *testing.T) {
	var msg = NewMsgAddX509Cert(testconstants.LeafCertPem, testconstants.Signer)

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

func TestMsgMsgAddX509Cert(t *testing.T) {
	var msg = NewMsgAddX509Cert(testconstants.StubCert, testconstants.Signer)
	res := msg.GetSignBytes()

	expected := `{"type":"pki/AddX509Cert","value":{"cert":"pem certificate",` +
		`"signer":"cosmos1p72j8mgkf39qjzcmr283w8l8y9qv30qpj056uz"}}`
	require.Equal(t, expected, string(res))
}
