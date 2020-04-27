package types

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
	var msg = NewMsgProposeAddX509RootCert(test_constants.RootCertPem, test_constants.Signer)

	require.Equal(t, msg.Route(), RouterKey)
	require.Equal(t, msg.Type(), "propose_add_x509_root_cert")
	require.Equal(t, msg.GetSigners(), []sdk.AccAddress{test_constants.Signer})
}

func TestValidateMsgProposeAddX509RootCert(t *testing.T) {
	cases := []struct {
		valid bool
		msg   MsgProposeAddX509RootCert
	}{
		{true, NewMsgProposeAddX509RootCert(
			test_constants.RootCertPem, test_constants.Signer)},
		{false, NewMsgProposeAddX509RootCert(
			"", test_constants.Signer)},
		{false, NewMsgProposeAddX509RootCert(
			test_constants.RootCertPem, nil)},
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
	var msg = NewMsgProposeAddX509RootCert(test_constants.StubCert, test_constants.Signer)
	res := msg.GetSignBytes()

	expected := `{"type":"pki/ProposeAddX509RootCert","value":{"cert":"pem certificate","signer":"cosmos1p72j8mgkf39qjzcmr283w8l8y9qv30qpj056uz"}}`
	require.Equal(t, expected, string(res))
}

/*
	MsgApproveAddX509RootCert(
*/

func TestNewMsgApproveAddX509RootCert(t *testing.T) {
	var msg = NewMsgApproveAddX509RootCert(test_constants.LeafSubject, test_constants.LeafSubjectKeyId, test_constants.Signer)

	require.Equal(t, msg.Route(), RouterKey)
	require.Equal(t, msg.Type(), "approve_add_x509_root_cert")
	require.Equal(t, msg.GetSigners(), []sdk.AccAddress{test_constants.Signer})
}

func TestValidateMsgApproveAddX509RootCert(t *testing.T) {
	cases := []struct {
		valid bool
		msg   MsgApproveAddX509RootCert
	}{
		{true, NewMsgApproveAddX509RootCert(
			test_constants.LeafSubject, test_constants.LeafSubjectKeyId, test_constants.Signer)},
		{false, NewMsgApproveAddX509RootCert(
			"", test_constants.LeafSubjectKeyId, test_constants.Signer)},
		{false, NewMsgApproveAddX509RootCert(
			test_constants.LeafSubject, "", test_constants.Signer)},
		{false, NewMsgApproveAddX509RootCert(
			test_constants.LeafSubject, test_constants.LeafSubjectKeyId, nil)},
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
	var msg = NewMsgApproveAddX509RootCert(test_constants.LeafSubject, test_constants.LeafSubjectKeyId, test_constants.Signer)
	res := msg.GetSignBytes()

	expected := `{"type":"pki/ApproveAddX509RootCert","value":{"signer":"cosmos1p72j8mgkf39qjzcmr283w8l8y9qv30qpj056uz",` +
		`"subject":"CN=dsr-corporation.com","subject_key_id":"8A:34:B:5C:D8:42:18:F2:C1:2A:AC:7A:B3:8F:6E:90:66:F4:4E:5C"}}`
	require.Equal(t, expected, string(res))
}

/*
	MsgAddX509Cert
*/

func TestNewMsgAddX509Cert(t *testing.T) {
	var msg = NewMsgAddX509Cert(test_constants.LeafCertPem, test_constants.Signer)

	require.Equal(t, msg.Route(), RouterKey)
	require.Equal(t, msg.Type(), "add_x509_cert")
	require.Equal(t, msg.GetSigners(), []sdk.AccAddress{test_constants.Signer})
}

func TestValidateMsgAddX509Cert(t *testing.T) {
	cases := []struct {
		valid bool
		msg   MsgAddX509Cert
	}{
		{true, NewMsgAddX509Cert(
			test_constants.LeafCertPem, test_constants.Signer)},
		{false, NewMsgAddX509Cert(
			"", test_constants.Signer)},
		{false, NewMsgAddX509Cert(
			test_constants.LeafCertPem, nil)},
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
	var msg = NewMsgAddX509Cert(test_constants.StubCert, test_constants.Signer)
	res := msg.GetSignBytes()

	expected := `{"type":"pki/AddX509Cert","value":{"cert":"pem certificate","signer":"cosmos1p72j8mgkf39qjzcmr283w8l8y9qv30qpj056uz"}}`
	require.Equal(t, expected, string(res))
}
