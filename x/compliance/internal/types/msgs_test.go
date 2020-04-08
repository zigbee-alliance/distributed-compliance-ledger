package types

import (
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliance/test_constants"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestNewMsgAddModelInfo(t *testing.T) {
	var msg = NewMsgAddModelInfo(test_constants.Id, test_constants.Name, test_constants.Owner, test_constants.Description, test_constants.Sku, test_constants.FirmwareVersion,
		test_constants.HardwareVersion, test_constants.CertificateID, test_constants.CertifiedDate, test_constants.TisOrTrpTestingCompleted, test_constants.Signer)

	require.Equal(t, msg.Route(), RouterKey)
	require.Equal(t, msg.Type(), "add_model_info")
	require.Equal(t, msg.GetSigners(), []sdk.AccAddress{test_constants.Signer})
}

func TestMsgAddModelInfoValidation(t *testing.T) {
	cases := []struct {
		valid bool
		msg   MsgAddModelInfo
	}{
		{true, NewMsgAddModelInfo(
			test_constants.Id, test_constants.Name, test_constants.Owner, test_constants.Description, test_constants.Sku, test_constants.FirmwareVersion,
			test_constants.HardwareVersion, test_constants.CertificateID, test_constants.CertifiedDate, test_constants.TisOrTrpTestingCompleted, test_constants.Signer)},
		{false, NewMsgAddModelInfo(
			"", test_constants.Name, test_constants.Owner, test_constants.Description, test_constants.Sku, test_constants.FirmwareVersion,
			test_constants.HardwareVersion, test_constants.CertificateID, test_constants.CertifiedDate, test_constants.TisOrTrpTestingCompleted, test_constants.Signer)},
		{false, NewMsgAddModelInfo(
			test_constants.Id, "", test_constants.Owner, test_constants.Description, test_constants.Sku, test_constants.FirmwareVersion,
			test_constants.HardwareVersion, test_constants.CertificateID, test_constants.CertifiedDate, test_constants.TisOrTrpTestingCompleted, test_constants.Signer)},
		{false, NewMsgAddModelInfo(
			test_constants.Id, test_constants.Name, nil, test_constants.Description, test_constants.Sku, test_constants.FirmwareVersion,
			test_constants.HardwareVersion, test_constants.CertificateID, test_constants.CertifiedDate, test_constants.TisOrTrpTestingCompleted, test_constants.Signer)},
		{false, NewMsgAddModelInfo(
			test_constants.Id, test_constants.Name, test_constants.Owner, "", test_constants.Sku, test_constants.FirmwareVersion,
			test_constants.HardwareVersion, test_constants.CertificateID, test_constants.CertifiedDate, test_constants.TisOrTrpTestingCompleted, test_constants.Signer)},
		{false, NewMsgAddModelInfo(
			test_constants.Id, test_constants.Name, test_constants.Owner, test_constants.Description, "", test_constants.FirmwareVersion,
			test_constants.HardwareVersion, test_constants.CertificateID, test_constants.CertifiedDate, test_constants.TisOrTrpTestingCompleted, test_constants.Signer)},
		{false, NewMsgAddModelInfo(
			test_constants.Id, test_constants.Name, test_constants.Owner, test_constants.Description, test_constants.Sku, "", test_constants.HardwareVersion,
			test_constants.CertificateID, test_constants.CertifiedDate, test_constants.TisOrTrpTestingCompleted, test_constants.Signer)},
		{false, NewMsgAddModelInfo(
			test_constants.Id, test_constants.Name, test_constants.Owner, test_constants.Description, test_constants.Sku, test_constants.FirmwareVersion, "",
			test_constants.CertificateID, test_constants.CertifiedDate, test_constants.TisOrTrpTestingCompleted, test_constants.Signer)},
		{true, NewMsgAddModelInfo(
			test_constants.Id, test_constants.Name, test_constants.Owner, test_constants.Description, test_constants.Sku, test_constants.FirmwareVersion, test_constants.HardwareVersion,
			"", test_constants.CertifiedDate, test_constants.TisOrTrpTestingCompleted, test_constants.Signer)},
		{true, NewMsgAddModelInfo(
			test_constants.Id, test_constants.Name, test_constants.Owner, test_constants.Description, test_constants.Sku, test_constants.FirmwareVersion, test_constants.HardwareVersion,
			test_constants.CertificateID, time.Time{}, test_constants.TisOrTrpTestingCompleted, test_constants.Signer)},
		{false, NewMsgAddModelInfo(
			test_constants.Id, test_constants.Name, test_constants.Owner, test_constants.Description, test_constants.Sku, test_constants.FirmwareVersion, test_constants.HardwareVersion,
			test_constants.CertificateID, test_constants.CertifiedDate, test_constants.TisOrTrpTestingCompleted, nil)},
		{true, NewMsgAddModelInfo(
			test_constants.Id, test_constants.Name, test_constants.Owner, test_constants.Description, test_constants.Sku, test_constants.FirmwareVersion, test_constants.HardwareVersion,
			test_constants.CertificateID, test_constants.CertifiedDate, true, test_constants.Owner)},
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

func TestMsgAddModelInfoGetSignBytes(t *testing.T) {
	var msg = NewMsgAddModelInfo(test_constants.Id, test_constants.Name, test_constants.Owner, test_constants.Description, test_constants.Sku, test_constants.FirmwareVersion, test_constants.HardwareVersion,
		test_constants.CertificateID, test_constants.CertifiedDate, test_constants.TisOrTrpTestingCompleted, test_constants.Signer)
	res := msg.GetSignBytes()

	expected := `{"type":"compliance/AddModelInfo","value":{"certificate_id":"ZIG12345678",` +
		`"certified_date":"2020-01-01T00:00:00Z","description":"Device Description","firmware_version":"1.0",` +
		`"hardware_version":"2.0","id":"Device Id","name":"Device Name","owner":"cosmos1damkuetjzyud4a",` +
		`"signer":"cosmos1wd5kwmn9wgr5dmap","sku":"RCU2205A","tis_or_trp_testing_completed":false}}`

	require.Equal(t, expected, string(res))
}
