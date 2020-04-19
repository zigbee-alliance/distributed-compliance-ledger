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
		ZbCertificationType, test_constants.Reason, test_constants.Signer)

	require.Equal(t, msg.Route(), RouterKey)
	require.Equal(t, msg.Type(), "certify_model")
	require.Equal(t, msg.GetSigners(), []sdk.AccAddress{test_constants.Signer})
}

func TestMsgCertifyModelValidation(t *testing.T) {
	cases := []struct {
		valid bool
		msg   MsgCertifyModel
	}{
		{true, NewMsgCertifyModel(
			test_constants.VID, test_constants.PID, test_constants.CertificationDate,
			CertificationType(test_constants.CertificationType), test_constants.Reason, test_constants.Signer)},
		{false, NewMsgCertifyModel(
			test_constants.VID, 0, test_constants.CertificationDate,
			CertificationType(test_constants.CertificationType), test_constants.Reason, test_constants.Signer)},
		{false, NewMsgCertifyModel(
			test_constants.VID, 0, test_constants.CertificationDate,
			CertificationType(test_constants.CertificationType), test_constants.Reason, test_constants.Signer)},
		{false, NewMsgCertifyModel(
			test_constants.VID, test_constants.PID, time.Time{},
			CertificationType(test_constants.CertificationType), test_constants.Reason, test_constants.Signer)},
		{true, NewMsgCertifyModel(
			test_constants.VID, test_constants.PID, test_constants.CertificationDate,
			"", test_constants.Reason, test_constants.Signer)},
		{false, NewMsgCertifyModel(
			test_constants.VID, test_constants.PID, test_constants.CertificationDate,
			"Other Type", test_constants.Reason, test_constants.Signer)},
		{true, NewMsgCertifyModel(
			test_constants.VID, test_constants.PID, test_constants.CertificationDate,
			CertificationType(test_constants.CertificationType), "", test_constants.Signer)},
		{false, NewMsgCertifyModel(
			test_constants.VID, test_constants.PID, test_constants.CertificationDate,
			CertificationType(test_constants.CertificationType), test_constants.Reason, nil)},
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

func TestMsgCertifyModelGetSignBytes(t *testing.T) {
	var msg = NewMsgCertifyModel(test_constants.VID, test_constants.PID, test_constants.CertificationDate,
		CertificationType(test_constants.CertificationType), test_constants.EmptyString, test_constants.Signer)
	res := msg.GetSignBytes()

	expected := `{"type":"compliance/CertifyModel","value":{"certification_date":"2020-01-01T00:00:00Z",` +
		`"certification_type":"zb","pid":22,"signer":"cosmos1p72j8mgkf39qjzcmr283w8l8y9qv30qpj056uz","vid":1}}`

	require.Equal(t, expected, string(res))
}

func TestNewMsgRevokeModel(t *testing.T) {
	var msg = NewMsgRevokeModel(test_constants.VID, test_constants.PID, test_constants.RevocationDate,
		CertificationType(test_constants.CertificationType), test_constants.RevocationReason, test_constants.Signer)

	require.Equal(t, msg.Route(), RouterKey)
	require.Equal(t, msg.Type(), "revoke_model")
	require.Equal(t, msg.GetSigners(), []sdk.AccAddress{test_constants.Signer})
}

func TestMsgRevokeModelValidation(t *testing.T) {
	cases := []struct {
		valid bool
		msg   MsgRevokeModel
	}{
		{true, NewMsgRevokeModel(
			test_constants.VID, test_constants.PID, test_constants.RevocationDate,
			CertificationType(test_constants.CertificationType), test_constants.RevocationReason, test_constants.Signer)},
		{false, NewMsgRevokeModel(
			0, test_constants.PID, test_constants.RevocationDate,
			CertificationType(test_constants.CertificationType), test_constants.RevocationReason, test_constants.Signer)},
		{false, NewMsgRevokeModel(
			test_constants.VID, 0, test_constants.RevocationDate,
			CertificationType(test_constants.CertificationType), test_constants.RevocationReason, test_constants.Signer)},
		{false, NewMsgRevokeModel(
			test_constants.VID, test_constants.PID, time.Time{},
			CertificationType(test_constants.CertificationType), test_constants.RevocationReason, test_constants.Signer)},
		{true, NewMsgRevokeModel(
			test_constants.VID, test_constants.PID, test_constants.RevocationDate,
			CertificationType(test_constants.CertificationType), "", test_constants.Signer)},
		{true, NewMsgRevokeModel(
			test_constants.VID, test_constants.PID, test_constants.RevocationDate,
			"", test_constants.RevocationReason, test_constants.Signer)},
		{false, NewMsgRevokeModel(
			test_constants.VID, test_constants.PID, test_constants.RevocationDate,
			CertificationType(test_constants.CertificationType), test_constants.RevocationReason, nil)},
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

func TestMsRevokeModelGetSignBytes(t *testing.T) {
	var msg = NewMsgRevokeModel(test_constants.VID, test_constants.PID, test_constants.RevocationDate,
		CertificationType(test_constants.CertificationType), test_constants.RevocationReason, test_constants.Signer)
	res := msg.GetSignBytes()

	expected := `{"type":"compliance/RevokeModel","value":{"certification_type":"zb","pid":22,"reason":"Some Reason",` +
		`"revocation_date":"2020-03-03T03:30:00Z","signer":"cosmos1p72j8mgkf39qjzcmr283w8l8y9qv30qpj056uz","vid":1}}`
	require.Equal(t, expected, string(res))
}
