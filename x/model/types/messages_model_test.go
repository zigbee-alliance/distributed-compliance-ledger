package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	tmrand "github.com/tendermint/tendermint/libs/rand"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	"github.com/zigbee-alliance/distributed-compliance-ledger/testutil/sample"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/validator"
)

//nolint:goconst
func TestMsgCreateModel_ValidateBasic(t *testing.T) {
	negativeTests := []struct {
		name string
		msg  *MsgCreateModel
		err  error
	}{
		{
			name: "Creator is omitted",
			msg: func(msg *MsgCreateModel) *MsgCreateModel {
				msg.Creator = ""
				return msg
			}(validMsgCreateModel()),
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "Creator is not valid address",
			msg: func(msg *MsgCreateModel) *MsgCreateModel {
				msg.Creator = "not valid address"
				return msg
			}(validMsgCreateModel()),
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "Vid < 0",
			msg: func(msg *MsgCreateModel) *MsgCreateModel {
				msg.Vid = -1
				return msg
			}(validMsgCreateModel()),
			err: validator.ErrFieldLowerBoundViolated,
		},
		{
			name: "Vid == 0",
			msg: func(msg *MsgCreateModel) *MsgCreateModel {
				msg.Vid = 0
				return msg
			}(validMsgCreateModel()),
			err: validator.ErrFieldLowerBoundViolated,
		},
		{
			name: "Vid > 65535",
			msg: func(msg *MsgCreateModel) *MsgCreateModel {
				msg.Vid = 65536
				return msg
			}(validMsgCreateModel()),
			err: validator.ErrFieldUpperBoundViolated,
		},
		{
			name: "Pid < 0",
			msg: func(msg *MsgCreateModel) *MsgCreateModel {
				msg.Pid = -1
				return msg
			}(validMsgCreateModel()),
			err: validator.ErrFieldLowerBoundViolated,
		},
		{
			name: "Pid == 0",
			msg: func(msg *MsgCreateModel) *MsgCreateModel {
				msg.Pid = 0
				return msg
			}(validMsgCreateModel()),
			err: validator.ErrFieldLowerBoundViolated,
		},
		{
			name: "Pid > 65535",
			msg: func(msg *MsgCreateModel) *MsgCreateModel {
				msg.Pid = 65536
				return msg
			}(validMsgCreateModel()),
			err: validator.ErrFieldUpperBoundViolated,
		},
		{
			name: "DeviceTypeId < 0",
			msg: func(msg *MsgCreateModel) *MsgCreateModel {
				msg.DeviceTypeId = -1
				return msg
			}(validMsgCreateModel()),
			err: validator.ErrFieldLowerBoundViolated,
		},
		{
			name: "DeviceTypeId > 65535",
			msg: func(msg *MsgCreateModel) *MsgCreateModel {
				msg.DeviceTypeId = 65536
				return msg
			}(validMsgCreateModel()),
			err: validator.ErrFieldUpperBoundViolated,
		},
		{
			name: "ProductName is omitted",
			msg: func(msg *MsgCreateModel) *MsgCreateModel {
				msg.ProductName = ""
				return msg
			}(validMsgCreateModel()),
			err: validator.ErrRequiredFieldMissing,
		},
		{
			name: "ProductName length > 128",
			msg: func(msg *MsgCreateModel) *MsgCreateModel {
				msg.ProductName = tmrand.Str(129)
				return msg
			}(validMsgCreateModel()),
			err: validator.ErrFieldMaxLengthExceeded,
		},
		{
			name: "ProductLabel is omitted",
			msg: func(msg *MsgCreateModel) *MsgCreateModel {
				msg.ProductLabel = ""
				return msg
			}(validMsgCreateModel()),
			err: validator.ErrRequiredFieldMissing,
		},
		{
			name: "ProductLabel length > 256",
			msg: func(msg *MsgCreateModel) *MsgCreateModel {
				msg.ProductLabel = tmrand.Str(257)
				return msg
			}(validMsgCreateModel()),
			err: validator.ErrFieldMaxLengthExceeded,
		},
		{
			name: "PartNumber is omitted",
			msg: func(msg *MsgCreateModel) *MsgCreateModel {
				msg.PartNumber = ""
				return msg
			}(validMsgCreateModel()),
			err: validator.ErrRequiredFieldMissing,
		},
		{
			name: "PartNumber length > 32",
			msg: func(msg *MsgCreateModel) *MsgCreateModel {
				msg.PartNumber = tmrand.Str(33)
				return msg
			}(validMsgCreateModel()),
			err: validator.ErrFieldMaxLengthExceeded,
		},
		{
			name: "CommissioningCustomFlow < 0",
			msg: func(msg *MsgCreateModel) *MsgCreateModel {
				msg.CommissioningCustomFlow = -1
				return msg
			}(validMsgCreateModel()),
			err: validator.ErrFieldLowerBoundViolated,
		},
		{
			name: "CommissioningCustomFlow > 2",
			msg: func(msg *MsgCreateModel) *MsgCreateModel {
				msg.CommissioningCustomFlow = 3
				return msg
			}(validMsgCreateModel()),
			err: validator.ErrFieldUpperBoundViolated,
		},
		{
			name: "CommissioningCustomFlowUrl is omitted when CommissioningCustomFlow == 2",
			msg: func(msg *MsgCreateModel) *MsgCreateModel {
				msg.CommissioningCustomFlow = 2
				msg.CommissioningCustomFlowUrl = ""
				return msg
			}(validMsgCreateModel()),
			err: validator.ErrRequiredFieldMissing,
		},
		{
			name: "CommissioningCustomFlowUrl is not valid URL",
			msg: func(msg *MsgCreateModel) *MsgCreateModel {
				msg.CommissioningCustomFlowUrl = "not valid URL"
				return msg
			}(validMsgCreateModel()),
			err: validator.ErrFieldNotValid,
		},
		{
			name: "CommissioningCustomFlowUrl starts with http:",
			msg: func(msg *MsgCreateModel) *MsgCreateModel {
				msg.CommissioningCustomFlowUrl = "http://sampleflowurl.dclmodel"
				return msg
			}(validMsgCreateModel()),
			err: validator.ErrFieldNotValid,
		},
		{
			name: "CommissioningCustomFlowUrl length > 256",
			msg: func(msg *MsgCreateModel) *MsgCreateModel {
				msg.CommissioningCustomFlowUrl = "https://sampleflowurl.dclauth/" + tmrand.Str(257-30) // length = 257
				return msg
			}(validMsgCreateModel()),
			err: validator.ErrFieldMaxLengthExceeded,
		},
		{
			name: "CommissioningModeInitialStepsInstruction length > 1024",
			msg: func(msg *MsgCreateModel) *MsgCreateModel {
				msg.CommissioningModeInitialStepsInstruction = tmrand.Str(1025)
				return msg
			}(validMsgCreateModel()),
			err: validator.ErrFieldMaxLengthExceeded,
		},
		{
			name: "CommissioningModeSecondaryStepsInstruction length > 1024",
			msg: func(msg *MsgCreateModel) *MsgCreateModel {
				msg.CommissioningModeSecondaryStepsInstruction = tmrand.Str(1025)
				return msg
			}(validMsgCreateModel()),
			err: validator.ErrFieldMaxLengthExceeded,
		},
		{
			name: "UserManualUrl is not valid URL",
			msg: func(msg *MsgCreateModel) *MsgCreateModel {
				msg.UserManualUrl = "not valid URL"
				return msg
			}(validMsgCreateModel()),
			err: validator.ErrFieldNotValid,
		},
		{
			name: "UserManualUrl starts with http:",
			msg: func(msg *MsgCreateModel) *MsgCreateModel {
				msg.UserManualUrl = "http://sampleflowurl.dclmodel"
				return msg
			}(validMsgCreateModel()),
			err: validator.ErrFieldNotValid,
		},
		{
			name: "UserManualUrl length > 256",
			msg: func(msg *MsgCreateModel) *MsgCreateModel {
				msg.UserManualUrl = "https://sampleflowurl.dclauth/" + tmrand.Str(257-30) // length = 257
				return msg
			}(validMsgCreateModel()),
			err: validator.ErrFieldMaxLengthExceeded,
		},
		{
			name: "SupportUrl is not valid URL",
			msg: func(msg *MsgCreateModel) *MsgCreateModel {
				msg.SupportUrl = "not valid URL"
				return msg
			}(validMsgCreateModel()),
			err: validator.ErrFieldNotValid,
		},
		{
			name: "SupportUrl starts with http:",
			msg: func(msg *MsgCreateModel) *MsgCreateModel {
				msg.SupportUrl = "http://sampleflowurl.dclmodel"
				return msg
			}(validMsgCreateModel()),
			err: validator.ErrFieldNotValid,
		},
		{
			name: "SupportUrl length > 256",
			msg: func(msg *MsgCreateModel) *MsgCreateModel {
				msg.SupportUrl = "https://sampleflowurl.dclauth/" + tmrand.Str(257-30) // length = 257
				return msg
			}(validMsgCreateModel()),
			err: validator.ErrFieldMaxLengthExceeded,
		},
		{
			name: "ProductUrl is not valid URL",
			msg: func(msg *MsgCreateModel) *MsgCreateModel {
				msg.ProductUrl = "not valid URL"
				return msg
			}(validMsgCreateModel()),
			err: validator.ErrFieldNotValid,
		},
		{
			name: "ProductUrl starts with http:",
			msg: func(msg *MsgCreateModel) *MsgCreateModel {
				msg.ProductUrl = "http://sampleflowurl.dclmodel"
				return msg
			}(validMsgCreateModel()),
			err: validator.ErrFieldNotValid,
		},
		{
			name: "ProductUrl length > 256",
			msg: func(msg *MsgCreateModel) *MsgCreateModel {
				msg.ProductUrl = "https://sampleflowurl.dclauth/" + tmrand.Str(257-30) // length = 257
				return msg
			}(validMsgCreateModel()),
			err: validator.ErrFieldMaxLengthExceeded,
		},
		{
			name: "LsfUrl is not valid URL",
			msg: func(msg *MsgCreateModel) *MsgCreateModel {
				msg.LsfUrl = "not valid URL"
				return msg
			}(validMsgCreateModel()),
			err: validator.ErrFieldNotValid,
		},
		{
			name: "LsfUrl starts with http:",
			msg: func(msg *MsgCreateModel) *MsgCreateModel {
				msg.LsfUrl = "http://sampleflowurl.dclmodel"
				return msg
			}(validMsgCreateModel()),
			err: validator.ErrFieldNotValid,
		},
		{
			name: "LsfUrl length > 256",
			msg: func(msg *MsgCreateModel) *MsgCreateModel {
				msg.LsfUrl = "https://sampleflowurl.dclauth/" + tmrand.Str(257-30) // length = 257
				return msg
			}(validMsgCreateModel()),
			err: validator.ErrFieldMaxLengthExceeded,
		},
		{
			name: "LsfRevision is missing",
			msg: func(msg *MsgCreateModel) *MsgCreateModel {
				msg.LsfRevision = 0
				return msg
			}(validMsgCreateModel()),
			err: validator.ErrRequiredFieldMissing,
		},
	}

	positiveTests := []struct {
		name string
		msg  *MsgCreateModel
	}{
		{
			name: "valid message",
			msg:  validMsgCreateModel(),
		},
		{
			name: "Creator is valid address",
			msg: func(msg *MsgCreateModel) *MsgCreateModel {
				msg.Creator = sample.AccAddress()
				return msg
			}(validMsgCreateModel()),
		},
		{
			name: "Vid == 1",
			msg: func(msg *MsgCreateModel) *MsgCreateModel {
				msg.Vid = 1
				return msg
			}(validMsgCreateModel()),
		},
		{
			name: "Vid == 65535",
			msg: func(msg *MsgCreateModel) *MsgCreateModel {
				msg.Vid = 65535
				return msg
			}(validMsgCreateModel()),
		},
		{
			name: "Pid == 1",
			msg: func(msg *MsgCreateModel) *MsgCreateModel {
				msg.Pid = 1
				return msg
			}(validMsgCreateModel()),
		},
		{
			name: "Pid == 65535",
			msg: func(msg *MsgCreateModel) *MsgCreateModel {
				msg.Pid = 65535
				return msg
			}(validMsgCreateModel()),
		},
		{
			name: "DeviceTypeId == 0",
			msg: func(msg *MsgCreateModel) *MsgCreateModel {
				msg.DeviceTypeId = 0
				return msg
			}(validMsgCreateModel()),
		},
		{
			name: "DeviceTypeId == 65535",
			msg: func(msg *MsgCreateModel) *MsgCreateModel {
				msg.DeviceTypeId = 65535
				return msg
			}(validMsgCreateModel()),
		},
		{
			name: "ProductName length == 128",
			msg: func(msg *MsgCreateModel) *MsgCreateModel {
				msg.ProductName = tmrand.Str(128)
				return msg
			}(validMsgCreateModel()),
		},
		{
			name: "ProductLabel length == 256",
			msg: func(msg *MsgCreateModel) *MsgCreateModel {
				msg.ProductLabel = tmrand.Str(256)
				return msg
			}(validMsgCreateModel()),
		},
		{
			name: "PartNumber length == 32",
			msg: func(msg *MsgCreateModel) *MsgCreateModel {
				msg.PartNumber = tmrand.Str(32)
				return msg
			}(validMsgCreateModel()),
		},
		{
			name: "CommissioningCustomFlow == 0",
			msg: func(msg *MsgCreateModel) *MsgCreateModel {
				msg.CommissioningCustomFlow = 0
				return msg
			}(validMsgCreateModel()),
		},
		{
			name: "CommissioningCustomFlow == 2",
			msg: func(msg *MsgCreateModel) *MsgCreateModel {
				msg.CommissioningCustomFlow = 2
				return msg
			}(validMsgCreateModel()),
		},
		{
			name: "CommissioningCustomFlowUrl is omitted when CommissioningCustomFlow != 2",
			msg: func(msg *MsgCreateModel) *MsgCreateModel {
				msg.CommissioningCustomFlow = 1
				msg.CommissioningCustomFlowUrl = ""
				return msg
			}(validMsgCreateModel()),
		},
		{
			name: "CommissioningCustomFlowUrl is set when CommissioningCustomFlow == 2",
			msg: func(msg *MsgCreateModel) *MsgCreateModel {
				msg.CommissioningCustomFlow = 2
				msg.CommissioningCustomFlowUrl = "https://sampleflowurl.dclmodel"
				return msg
			}(validMsgCreateModel()),
		},
		{
			name: "CommissioningCustomFlowUrl length == 256",
			msg: func(msg *MsgCreateModel) *MsgCreateModel {
				msg.CommissioningCustomFlowUrl = "https://sampleflowurl.dclauth/" + tmrand.Str(256-30) // length = 256
				return msg
			}(validMsgCreateModel()),
		},
		{
			name: "CommissioningModeInitialStepsHint == 0",
			msg: func(msg *MsgCreateModel) *MsgCreateModel {
				msg.CommissioningModeInitialStepsHint = 0
				return msg
			}(validMsgCreateModel()),
		},
		{
			name: "CommissioningModeInitialStepsHint > 0",
			msg: func(msg *MsgCreateModel) *MsgCreateModel {
				msg.CommissioningModeInitialStepsHint = 1
				return msg
			}(validMsgCreateModel()),
		},
		{
			name: "CommissioningModeInitialStepsInstruction is omitted",
			msg: func(msg *MsgCreateModel) *MsgCreateModel {
				msg.CommissioningModeInitialStepsInstruction = ""
				return msg
			}(validMsgCreateModel()),
		},
		{
			name: "CommissioningModeInitialStepsInstruction length == 1024",
			msg: func(msg *MsgCreateModel) *MsgCreateModel {
				msg.CommissioningModeInitialStepsInstruction = tmrand.Str(1024)
				return msg
			}(validMsgCreateModel()),
		},
		{
			name: "CommissioningModeSecondaryStepsHint == 0",
			msg: func(msg *MsgCreateModel) *MsgCreateModel {
				msg.CommissioningModeSecondaryStepsHint = 0
				return msg
			}(validMsgCreateModel()),
		},
		{
			name: "CommissioningModeSecondaryStepsHint > 0",
			msg: func(msg *MsgCreateModel) *MsgCreateModel {
				msg.CommissioningModeSecondaryStepsHint = 1
				return msg
			}(validMsgCreateModel()),
		},
		{
			name: "CommissioningModeSecondaryStepsInstruction is omitted",
			msg: func(msg *MsgCreateModel) *MsgCreateModel {
				msg.CommissioningModeSecondaryStepsInstruction = ""
				return msg
			}(validMsgCreateModel()),
		},
		{
			name: "CommissioningModeSecondaryStepsInstruction length == 1024",
			msg: func(msg *MsgCreateModel) *MsgCreateModel {
				msg.CommissioningModeSecondaryStepsInstruction = tmrand.Str(1024)
				return msg
			}(validMsgCreateModel()),
		},
		{
			name: "UserManualUrl is omitted",
			msg: func(msg *MsgCreateModel) *MsgCreateModel {
				msg.UserManualUrl = ""
				return msg
			}(validMsgCreateModel()),
		},
		{
			name: "UserManualUrl length == 256",
			msg: func(msg *MsgCreateModel) *MsgCreateModel {
				msg.UserManualUrl = "https://sampleflowurl.dclauth/" + tmrand.Str(256-30) // length = 256
				return msg
			}(validMsgCreateModel()),
		},
		{
			name: "SupportUrl is omitted",
			msg: func(msg *MsgCreateModel) *MsgCreateModel {
				msg.SupportUrl = ""
				return msg
			}(validMsgCreateModel()),
		},
		{
			name: "SupportUrl length == 256",
			msg: func(msg *MsgCreateModel) *MsgCreateModel {
				msg.SupportUrl = "https://sampleflowurl.dclauth/" + tmrand.Str(256-30) // length = 256
				return msg
			}(validMsgCreateModel()),
		},
		{
			name: "ProductUrl is omitted",
			msg: func(msg *MsgCreateModel) *MsgCreateModel {
				msg.ProductUrl = ""
				return msg
			}(validMsgCreateModel()),
		},
		{
			name: "ProductUrl length == 256",
			msg: func(msg *MsgCreateModel) *MsgCreateModel {
				msg.ProductUrl = "https://sampleflowurl.dclauth/" + tmrand.Str(256-30) // length = 256
				return msg
			}(validMsgCreateModel()),
		},
		{
			name: "LsfUrl is omitted",
			msg: func(msg *MsgCreateModel) *MsgCreateModel {
				msg.LsfUrl = ""
				return msg
			}(validMsgCreateModel()),
		},
		{
			name: "LsfUrl length == 256",
			msg: func(msg *MsgCreateModel) *MsgCreateModel {
				msg.LsfUrl = "https://sampleflowurl.dclauth/" + tmrand.Str(256-30) // length = 256
				return msg
			}(validMsgCreateModel()),
		},
	}

	for _, tt := range negativeTests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			require.Error(t, err)
			require.ErrorIs(t, err, tt.err)
		})
	}

	for _, tt := range positiveTests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			require.NoError(t, err)
		})
	}
}

func TestMsgUpdateModel_ValidateBasic(t *testing.T) {
	negativeTests := []struct {
		name string
		msg  *MsgUpdateModel
		err  error
	}{
		{
			name: "Creator is omitted",
			msg: func(msg *MsgUpdateModel) *MsgUpdateModel {
				msg.Creator = ""
				return msg
			}(validMsgUpdateModel()),
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "Creator is not valid address",
			msg: func(msg *MsgUpdateModel) *MsgUpdateModel {
				msg.Creator = "not valid address"
				return msg
			}(validMsgUpdateModel()),
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "Vid < 0",
			msg: func(msg *MsgUpdateModel) *MsgUpdateModel {
				msg.Vid = -1
				return msg
			}(validMsgUpdateModel()),
			err: validator.ErrFieldLowerBoundViolated,
		},
		{
			name: "Vid == 0",
			msg: func(msg *MsgUpdateModel) *MsgUpdateModel {
				msg.Vid = 0
				return msg
			}(validMsgUpdateModel()),
			err: validator.ErrFieldLowerBoundViolated,
		},
		{
			name: "Vid > 65535",
			msg: func(msg *MsgUpdateModel) *MsgUpdateModel {
				msg.Vid = 65536
				return msg
			}(validMsgUpdateModel()),
			err: validator.ErrFieldUpperBoundViolated,
		},
		{
			name: "Pid < 0",
			msg: func(msg *MsgUpdateModel) *MsgUpdateModel {
				msg.Pid = -1
				return msg
			}(validMsgUpdateModel()),
			err: validator.ErrFieldLowerBoundViolated,
		},
		{
			name: "Pid == 0",
			msg: func(msg *MsgUpdateModel) *MsgUpdateModel {
				msg.Pid = 0
				return msg
			}(validMsgUpdateModel()),
			err: validator.ErrFieldLowerBoundViolated,
		},
		{
			name: "Pid > 65535",
			msg: func(msg *MsgUpdateModel) *MsgUpdateModel {
				msg.Pid = 65536
				return msg
			}(validMsgUpdateModel()),
			err: validator.ErrFieldUpperBoundViolated,
		},
		{
			name: "ProductName length > 128",
			msg: func(msg *MsgUpdateModel) *MsgUpdateModel {
				msg.ProductName = tmrand.Str(129)
				return msg
			}(validMsgUpdateModel()),
			err: validator.ErrFieldMaxLengthExceeded,
		},
		{
			name: "ProductLabel length > 256",
			msg: func(msg *MsgUpdateModel) *MsgUpdateModel {
				msg.ProductLabel = tmrand.Str(257)
				return msg
			}(validMsgUpdateModel()),
			err: validator.ErrFieldMaxLengthExceeded,
		},
		{
			name: "PartNumber length > 32",
			msg: func(msg *MsgUpdateModel) *MsgUpdateModel {
				msg.PartNumber = tmrand.Str(33)
				return msg
			}(validMsgUpdateModel()),
			err: validator.ErrFieldMaxLengthExceeded,
		},
		{
			name: "CommissioningCustomFlowUrl is not valid URL",
			msg: func(msg *MsgUpdateModel) *MsgUpdateModel {
				msg.CommissioningCustomFlowUrl = "not valid URL"
				return msg
			}(validMsgUpdateModel()),
			err: validator.ErrFieldNotValid,
		},
		{
			name: "CommissioningCustomFlowUrl starts with http:",
			msg: func(msg *MsgUpdateModel) *MsgUpdateModel {
				msg.CommissioningCustomFlowUrl = "http://sampleflowurl.dclmodel"
				return msg
			}(validMsgUpdateModel()),
			err: validator.ErrFieldNotValid,
		},
		{
			name: "CommissioningCustomFlowUrl length > 256",
			msg: func(msg *MsgUpdateModel) *MsgUpdateModel {
				msg.CommissioningCustomFlowUrl = "https://sampleflowurl.dclauth/" + tmrand.Str(257-30) // length = 257
				return msg
			}(validMsgUpdateModel()),
			err: validator.ErrFieldMaxLengthExceeded,
		},
		{
			name: "CommissioningModeInitialStepsInstruction length > 1024",
			msg: func(msg *MsgUpdateModel) *MsgUpdateModel {
				msg.CommissioningModeInitialStepsInstruction = tmrand.Str(1025)
				return msg
			}(validMsgUpdateModel()),
			err: validator.ErrFieldMaxLengthExceeded,
		},
		{
			name: "CommissioningModeSecondaryStepsInstruction length > 1024",
			msg: func(msg *MsgUpdateModel) *MsgUpdateModel {
				msg.CommissioningModeSecondaryStepsInstruction = tmrand.Str(1025)
				return msg
			}(validMsgUpdateModel()),
			err: validator.ErrFieldMaxLengthExceeded,
		},
		{
			name: "UserManualUrl is not valid URL",
			msg: func(msg *MsgUpdateModel) *MsgUpdateModel {
				msg.UserManualUrl = "not valid URL"
				return msg
			}(validMsgUpdateModel()),
			err: validator.ErrFieldNotValid,
		},
		{
			name: "UserManualUrl starts with http:",
			msg: func(msg *MsgUpdateModel) *MsgUpdateModel {
				msg.UserManualUrl = "http://sampleflowurl.dclmodel"
				return msg
			}(validMsgUpdateModel()),
			err: validator.ErrFieldNotValid,
		},
		{
			name: "UserManualUrl length > 256",
			msg: func(msg *MsgUpdateModel) *MsgUpdateModel {
				msg.UserManualUrl = "https://sampleflowurl.dclauth/" + tmrand.Str(257-30) // length = 257
				return msg
			}(validMsgUpdateModel()),
			err: validator.ErrFieldMaxLengthExceeded,
		},
		{
			name: "SupportUrl is not valid URL",
			msg: func(msg *MsgUpdateModel) *MsgUpdateModel {
				msg.SupportUrl = "not valid URL"
				return msg
			}(validMsgUpdateModel()),
			err: validator.ErrFieldNotValid,
		},
		{
			name: "SupportUrl starts with http:",
			msg: func(msg *MsgUpdateModel) *MsgUpdateModel {
				msg.SupportUrl = "http://sampleflowurl.dclmodel"
				return msg
			}(validMsgUpdateModel()),
			err: validator.ErrFieldNotValid,
		},
		{
			name: "SupportUrl length > 256",
			msg: func(msg *MsgUpdateModel) *MsgUpdateModel {
				msg.SupportUrl = "https://sampleflowurl.dclauth/" + tmrand.Str(257-30) // length = 257
				return msg
			}(validMsgUpdateModel()),
			err: validator.ErrFieldMaxLengthExceeded,
		},
		{
			name: "ProductUrl is not valid URL",
			msg: func(msg *MsgUpdateModel) *MsgUpdateModel {
				msg.ProductUrl = "not valid URL"
				return msg
			}(validMsgUpdateModel()),
			err: validator.ErrFieldNotValid,
		},
		{
			name: "ProductUrl starts with http:",
			msg: func(msg *MsgUpdateModel) *MsgUpdateModel {
				msg.ProductUrl = "http://sampleflowurl.dclmodel"
				return msg
			}(validMsgUpdateModel()),
			err: validator.ErrFieldNotValid,
		},
		{
			name: "ProductUrl length > 256",
			msg: func(msg *MsgUpdateModel) *MsgUpdateModel {
				msg.ProductUrl = "https://sampleflowurl.dclauth/" + tmrand.Str(257-30) // length = 257
				return msg
			}(validMsgUpdateModel()),
			err: validator.ErrFieldMaxLengthExceeded,
		},
	}

	positiveTests := []struct {
		name string
		msg  *MsgUpdateModel
	}{
		{
			name: "valid message",
			msg:  validMsgUpdateModel(),
		},
		{
			name: "Creator is valid address",
			msg: func(msg *MsgUpdateModel) *MsgUpdateModel {
				msg.Creator = sample.AccAddress()
				return msg
			}(validMsgUpdateModel()),
		},
		{
			name: "Vid == 1",
			msg: func(msg *MsgUpdateModel) *MsgUpdateModel {
				msg.Vid = 1
				return msg
			}(validMsgUpdateModel()),
		},
		{
			name: "Vid == 65535",
			msg: func(msg *MsgUpdateModel) *MsgUpdateModel {
				msg.Vid = 65535
				return msg
			}(validMsgUpdateModel()),
		},
		{
			name: "Pid == 1",
			msg: func(msg *MsgUpdateModel) *MsgUpdateModel {
				msg.Pid = 1
				return msg
			}(validMsgUpdateModel()),
		},
		{
			name: "Pid == 65535",
			msg: func(msg *MsgUpdateModel) *MsgUpdateModel {
				msg.Pid = 65535
				return msg
			}(validMsgUpdateModel()),
		},
		{
			name: "ProductName is omitted",
			msg: func(msg *MsgUpdateModel) *MsgUpdateModel {
				msg.ProductName = ""
				return msg
			}(validMsgUpdateModel()),
		},
		{
			name: "ProductName length == 128",
			msg: func(msg *MsgUpdateModel) *MsgUpdateModel {
				msg.ProductName = tmrand.Str(128)
				return msg
			}(validMsgUpdateModel()),
		},
		{
			name: "ProductLabel is omitted",
			msg: func(msg *MsgUpdateModel) *MsgUpdateModel {
				msg.ProductLabel = ""
				return msg
			}(validMsgUpdateModel()),
		},
		{
			name: "ProductLabel length == 256",
			msg: func(msg *MsgUpdateModel) *MsgUpdateModel {
				msg.ProductLabel = tmrand.Str(256)
				return msg
			}(validMsgUpdateModel()),
		},
		{
			name: "PartNumber is omitted",
			msg: func(msg *MsgUpdateModel) *MsgUpdateModel {
				msg.PartNumber = ""
				return msg
			}(validMsgUpdateModel()),
		},
		{
			name: "PartNumber length == 32",
			msg: func(msg *MsgUpdateModel) *MsgUpdateModel {
				msg.PartNumber = tmrand.Str(32)
				return msg
			}(validMsgUpdateModel()),
		},
		{
			name: "CommissioningCustomFlowUrl is omitted",
			msg: func(msg *MsgUpdateModel) *MsgUpdateModel {
				msg.CommissioningCustomFlowUrl = ""
				return msg
			}(validMsgUpdateModel()),
		},
		{
			name: "CommissioningCustomFlowUrl length == 256",
			msg: func(msg *MsgUpdateModel) *MsgUpdateModel {
				msg.CommissioningCustomFlowUrl = "https://sampleflowurl.dclauth/" + tmrand.Str(256-30) // length = 256
				return msg
			}(validMsgUpdateModel()),
		},
		{
			name: "CommissioningModeInitialStepsInstruction is omitted",
			msg: func(msg *MsgUpdateModel) *MsgUpdateModel {
				msg.CommissioningModeInitialStepsInstruction = ""
				return msg
			}(validMsgUpdateModel()),
		},
		{
			name: "CommissioningModeInitialStepsInstruction length == 1024",
			msg: func(msg *MsgUpdateModel) *MsgUpdateModel {
				msg.CommissioningModeInitialStepsInstruction = tmrand.Str(1024)
				return msg
			}(validMsgUpdateModel()),
		},
		{
			name: "CommissioningModeSecondaryStepsInstruction is omitted",
			msg: func(msg *MsgUpdateModel) *MsgUpdateModel {
				msg.CommissioningModeSecondaryStepsInstruction = ""
				return msg
			}(validMsgUpdateModel()),
		},
		{
			name: "CommissioningModeSecondaryStepsInstruction length == 1024",
			msg: func(msg *MsgUpdateModel) *MsgUpdateModel {
				msg.CommissioningModeSecondaryStepsInstruction = tmrand.Str(1024)
				return msg
			}(validMsgUpdateModel()),
		},
		{
			name: "UserManualUrl is omitted",
			msg: func(msg *MsgUpdateModel) *MsgUpdateModel {
				msg.UserManualUrl = ""
				return msg
			}(validMsgUpdateModel()),
		},
		{
			name: "UserManualUrl length == 256",
			msg: func(msg *MsgUpdateModel) *MsgUpdateModel {
				msg.UserManualUrl = "https://sampleflowurl.dclauth/" + tmrand.Str(256-30) // length = 256
				return msg
			}(validMsgUpdateModel()),
		},
		{
			name: "SupportUrl is omitted",
			msg: func(msg *MsgUpdateModel) *MsgUpdateModel {
				msg.SupportUrl = ""
				return msg
			}(validMsgUpdateModel()),
		},
		{
			name: "SupportUrl length == 256",
			msg: func(msg *MsgUpdateModel) *MsgUpdateModel {
				msg.SupportUrl = "https://sampleflowurl.dclauth/" + tmrand.Str(256-30) // length = 256
				return msg
			}(validMsgUpdateModel()),
		},
		{
			name: "ProductUrl is omitted",
			msg: func(msg *MsgUpdateModel) *MsgUpdateModel {
				msg.ProductUrl = ""
				return msg
			}(validMsgUpdateModel()),
		},
		{
			name: "ProductUrl length == 256",
			msg: func(msg *MsgUpdateModel) *MsgUpdateModel {
				msg.ProductUrl = "https://sampleflowurl.dclauth/" + tmrand.Str(256-30) // length = 256
				return msg
			}(validMsgUpdateModel()),
		},
	}

	for _, tt := range negativeTests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			require.Error(t, err)
			require.ErrorIs(t, err, tt.err)
		})
	}

	for _, tt := range positiveTests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			require.NoError(t, err)
		})
	}
}

func TestMsgDeleteModel_ValidateBasic(t *testing.T) {
	negativeTests := []struct {
		name string
		msg  *MsgDeleteModel
		err  error
	}{
		{
			name: "Creator is omitted",
			msg: &MsgDeleteModel{
				Creator: "",
				Vid:     testconstants.Vid,
				Pid:     testconstants.Pid,
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "Creator is not valid address",
			msg: &MsgDeleteModel{
				Creator: "not valid address",
				Vid:     testconstants.Vid,
				Pid:     testconstants.Pid,
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "Vid < 0",
			msg: &MsgDeleteModel{
				Creator: sample.AccAddress(),
				Vid:     -1,
				Pid:     testconstants.Pid,
			},
			err: validator.ErrFieldLowerBoundViolated,
		},
		{
			name: "Vid == 0",
			msg: &MsgDeleteModel{
				Creator: sample.AccAddress(),
				Vid:     0,
				Pid:     testconstants.Pid,
			},
			err: validator.ErrFieldLowerBoundViolated,
		},
		{
			name: "Vid > 65535",
			msg: &MsgDeleteModel{
				Creator: sample.AccAddress(),
				Vid:     65536,
				Pid:     testconstants.Pid,
			},
			err: validator.ErrFieldUpperBoundViolated,
		},
		{
			name: "Pid < 0",
			msg: &MsgDeleteModel{
				Creator: sample.AccAddress(),
				Vid:     testconstants.Vid,
				Pid:     -1,
			},
			err: validator.ErrFieldLowerBoundViolated,
		},
		{
			name: "Pid == 0",
			msg: &MsgDeleteModel{
				Creator: sample.AccAddress(),
				Vid:     testconstants.Vid,
				Pid:     0,
			},
			err: validator.ErrFieldLowerBoundViolated,
		},
		{
			name: "Pid > 65535",
			msg: &MsgDeleteModel{
				Creator: sample.AccAddress(),
				Vid:     testconstants.Vid,
				Pid:     65536,
			},
			err: validator.ErrFieldUpperBoundViolated,
		},
	}

	positiveTests := []struct {
		name string
		msg  *MsgDeleteModel
	}{
		{
			name: "Creator is valid address",
			msg: &MsgDeleteModel{
				Creator: sample.AccAddress(),
				Vid:     testconstants.Vid,
				Pid:     testconstants.Pid,
			},
		},
		{
			name: "Vid == 1",
			msg: &MsgDeleteModel{
				Creator: sample.AccAddress(),
				Vid:     1,
				Pid:     testconstants.Pid,
			},
		},
		{
			name: "Vid == 65535",
			msg: &MsgDeleteModel{
				Creator: sample.AccAddress(),
				Vid:     65535,
				Pid:     testconstants.Pid,
			},
		},
		{
			name: "Pid == 1",
			msg: &MsgDeleteModel{
				Creator: sample.AccAddress(),
				Vid:     testconstants.Vid,
				Pid:     1,
			},
		},
		{
			name: "Pid == 65535",
			msg: &MsgDeleteModel{
				Creator: sample.AccAddress(),
				Vid:     testconstants.Vid,
				Pid:     65535,
			},
		},
	}

	for _, tt := range negativeTests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			require.Error(t, err)
			require.ErrorIs(t, err, tt.err)
		})
	}

	for _, tt := range positiveTests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			require.NoError(t, err)
		})
	}
}

func validMsgCreateModel() *MsgCreateModel {
	return &MsgCreateModel{
		Creator:                                  sample.AccAddress(),
		Vid:                                      testconstants.Vid,
		Pid:                                      testconstants.Pid,
		DeviceTypeId:                             testconstants.DeviceTypeId,
		ProductName:                              testconstants.ProductName,
		ProductLabel:                             testconstants.ProductLabel,
		PartNumber:                               testconstants.PartNumber,
		CommissioningCustomFlow:                  testconstants.CommissioningCustomFlow,
		CommissioningCustomFlowUrl:               testconstants.CommissioningCustomFlowUrl,
		CommissioningModeInitialStepsHint:        testconstants.CommissioningModeInitialStepsHint,
		CommissioningModeInitialStepsInstruction: testconstants.CommissioningModeInitialStepsInstruction,
		CommissioningModeSecondaryStepsHint:      testconstants.CommissioningModeSecondaryStepsHint,
		CommissioningModeSecondaryStepsInstruction: testconstants.CommissioningModeSecondaryStepsInstruction,
		UserManualUrl: testconstants.UserManualUrl,
		SupportUrl:    testconstants.SupportUrl,
		ProductUrl:    testconstants.ProductUrl,
		LsfUrl:        testconstants.LsfUrl,
		LsfRevision:   testconstants.LsfRevision,
	}
}

func validMsgUpdateModel() *MsgUpdateModel {
	return &MsgUpdateModel{
		Creator:                                  sample.AccAddress(),
		Vid:                                      testconstants.Vid,
		Pid:                                      testconstants.Pid,
		ProductName:                              testconstants.ProductName + "-updated",
		ProductLabel:                             testconstants.ProductLabel + "-updated",
		PartNumber:                               testconstants.PartNumber + "-updated",
		CommissioningCustomFlowUrl:               testconstants.CommissioningCustomFlowUrl + "/updated",
		CommissioningModeInitialStepsInstruction: testconstants.CommissioningModeInitialStepsInstruction + "-updated",
		CommissioningModeSecondaryStepsInstruction: testconstants.CommissioningModeSecondaryStepsInstruction + "-updated",
		UserManualUrl: testconstants.UserManualUrl + "/updated",
		SupportUrl:    testconstants.SupportUrl + "/updated",
		ProductUrl:    testconstants.ProductUrl + "/updated",
		LsfUrl:        testconstants.LsfUrl + "/updated",
		LsfRevision:   testconstants.LsfRevision + 1,
	}
}
