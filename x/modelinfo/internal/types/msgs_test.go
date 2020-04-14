package types

import (
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/modelinfo/test_constants"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewMsgAddModelInfo(t *testing.T) {
	var msg = NewMsgAddModelInfo(test_constants.VID, test_constants.PID, test_constants.CID, test_constants.Name,
		test_constants.Description, test_constants.Sku, test_constants.FirmwareVersion,
		test_constants.HardwareVersion, test_constants.Custom, test_constants.TisOrTrpTestingCompleted, test_constants.Signer)

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
			test_constants.VID, test_constants.PID, test_constants.CID, test_constants.Name, test_constants.Description,
			test_constants.Sku, test_constants.FirmwareVersion, test_constants.HardwareVersion,
			test_constants.Custom, test_constants.TisOrTrpTestingCompleted, test_constants.Signer)},
		{false, NewMsgAddModelInfo(
			test_constants.VID, 0, test_constants.CID, test_constants.Name, test_constants.Description,
			test_constants.Sku, test_constants.FirmwareVersion, test_constants.HardwareVersion,
			test_constants.Custom, test_constants.TisOrTrpTestingCompleted, test_constants.Signer)},
		{false, NewMsgAddModelInfo(
			test_constants.VID, 0, test_constants.CID, test_constants.Name, test_constants.Description,
			test_constants.Sku, test_constants.FirmwareVersion, test_constants.HardwareVersion,
			test_constants.Custom, test_constants.TisOrTrpTestingCompleted, test_constants.Signer)},
		{true, NewMsgAddModelInfo(
			test_constants.VID, test_constants.PID, 0, test_constants.Name, test_constants.Description,
			test_constants.Sku, test_constants.FirmwareVersion, test_constants.HardwareVersion,
			test_constants.Custom, test_constants.TisOrTrpTestingCompleted, test_constants.Signer)},
		{false, NewMsgAddModelInfo(
			test_constants.VID, test_constants.PID, test_constants.CID, "", test_constants.Description,
			test_constants.Sku, test_constants.FirmwareVersion, test_constants.HardwareVersion,
			test_constants.Custom, test_constants.TisOrTrpTestingCompleted, test_constants.Signer)},
		{false, NewMsgAddModelInfo(
			test_constants.VID, test_constants.PID, test_constants.CID, test_constants.Name, "",
			test_constants.Sku, test_constants.FirmwareVersion, test_constants.HardwareVersion,
			test_constants.Custom, test_constants.TisOrTrpTestingCompleted, test_constants.Signer)},
		{false, NewMsgAddModelInfo(
			test_constants.VID, test_constants.PID, test_constants.CID, test_constants.Name, test_constants.Description,
			"", test_constants.FirmwareVersion, test_constants.HardwareVersion,
			test_constants.Custom, test_constants.TisOrTrpTestingCompleted, test_constants.Signer)},
		{false, NewMsgAddModelInfo(
			test_constants.VID, test_constants.PID, test_constants.CID, test_constants.Name, test_constants.Description,
			test_constants.Sku, "", test_constants.HardwareVersion,
			test_constants.Custom, test_constants.TisOrTrpTestingCompleted, test_constants.Signer)},
		{false, NewMsgAddModelInfo(
			test_constants.VID, test_constants.PID, test_constants.CID, test_constants.Name, test_constants.Description,
			test_constants.Sku, test_constants.FirmwareVersion, "",
			test_constants.Custom, test_constants.TisOrTrpTestingCompleted, test_constants.Signer)},
		{true, NewMsgAddModelInfo(
			test_constants.VID, test_constants.PID, test_constants.CID, test_constants.Name, test_constants.Description,
			test_constants.Sku, test_constants.FirmwareVersion, test_constants.HardwareVersion,
			"", test_constants.TisOrTrpTestingCompleted, test_constants.Signer)},
		{false, NewMsgAddModelInfo(
			test_constants.VID, test_constants.PID, test_constants.CID, test_constants.Name, test_constants.Description,
			test_constants.Sku, test_constants.FirmwareVersion, test_constants.HardwareVersion,
			test_constants.Custom, test_constants.TisOrTrpTestingCompleted, nil)},
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
	var msg = NewMsgAddModelInfo(test_constants.VID, test_constants.PID, test_constants.CID, test_constants.Name,
		test_constants.Description, test_constants.Sku, test_constants.FirmwareVersion, test_constants.HardwareVersion,
		test_constants.Custom, test_constants.TisOrTrpTestingCompleted, test_constants.Signer)
	res := msg.GetSignBytes()

	expected := `{"type":"modelinfo/AddModelInfo","value":{"cid":12345,"custom":"Custom data","description":"Device Description",` +
		`"firmware_version":"1.0","hardware_version":"2.0","name":"Device Name","pid":22,"signer":"cosmos1p72j8mgkf39qjzcmr283w8l8y9qv30qpj056uz",` +
		`"sku":"RCU2205A","tis_or_trp_testing_completed":false,"vid":1}}`

	require.Equal(t, expected, string(res))
}
