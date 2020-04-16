package types

import (
	"git.dsr-corporation.com/zb-ledger/zb-ledger/integration_tests/constants"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestNewMsgCertifyModel(t *testing.T) {
	var msg = NewMsgCertifyModel(test_constants.VID, test_constants.PID, test_constants.CertificationDate,
		test_constants.CertificationType, test_constants.Signer)

	require.Equal(t, msg.Route(), RouterKey)
	require.Equal(t, msg.Type(), "certify_model")
	require.Equal(t, msg.GetSigners(), []sdk.AccAddress{test_constants.Signer})
}

func TestNewMsgCertifyModelValidation(t *testing.T) {
	cases := []struct {
		valid bool
		msg   MsgCertifyModel
	}{
		{true, NewMsgCertifyModel(
			test_constants.VID, test_constants.PID, test_constants.CertificationDate,
			test_constants.CertificationType, test_constants.Signer)},
		{false, NewMsgCertifyModel(
			test_constants.VID, 0, test_constants.CertificationDate,
			test_constants.CertificationType, test_constants.Signer)},
		{false, NewMsgCertifyModel(
			test_constants.VID, 0, test_constants.CertificationDate,
			test_constants.CertificationType, test_constants.Signer)},
		{false, NewMsgCertifyModel(
			test_constants.VID, test_constants.PID, time.Time{},
			test_constants.CertificationType, test_constants.Signer)},
		{true, NewMsgCertifyModel(
			test_constants.VID, test_constants.PID, test_constants.CertificationDate,
			"", test_constants.Signer)},
		{false, NewMsgCertifyModel(
			test_constants.VID, test_constants.PID, test_constants.CertificationDate,
			"Other Type", test_constants.Signer)},
		{false, NewMsgCertifyModel(
			test_constants.VID, test_constants.PID, test_constants.CertificationDate,
			test_constants.CertificationType, nil)},
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

func TestMsgMsgCertifyModelGetSignBytes(t *testing.T) {
	var msg = NewMsgCertifyModel(test_constants.VID, test_constants.PID, test_constants.CertificationDate,
		test_constants.CertificationType, test_constants.Signer)
	res := msg.GetSignBytes()

	expected := `{"type":"compliance/CertifyModel","value":{"certification_date":"2020-01-01T00:00:00Z",` +
		`"certification_type":"zb","pid":22,"signer":"cosmos1p72j8mgkf39qjzcmr283w8l8y9qv30qpj056uz","vid":1}}`

	require.Equal(t, expected, string(res))
}
