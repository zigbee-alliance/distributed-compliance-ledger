package types

import (
	"testing"

	tmrand "github.com/cometbft/cometbft/libs/rand"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	"github.com/zigbee-alliance/distributed-compliance-ledger/testutil/sample"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/validator"
)

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
			name: "ProductLabel length > 256",
			msg: func(msg *MsgCreateModel) *MsgCreateModel {
				msg.ProductLabel = tmrand.Str(257)

				return msg
			}(validMsgCreateModel()),
			err: validator.ErrFieldMaxLengthExceeded,
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
			name: "DiscoveryCapabilitiesBitmask > 14",
			msg: func(msg *MsgCreateModel) *MsgCreateModel {
				msg.DiscoveryCapabilitiesBitmask = 15

				return msg
			}(validMsgCreateModel()),
			err: validator.ErrFieldUpperBoundViolated,
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
			name: "CommissioningCustomFlowUrl starts not with https:",
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
			name: "IcdUserActiveModeTriggerInstruction length > 1024",
			msg: func(msg *MsgCreateModel) *MsgCreateModel {
				msg.IcdUserActiveModeTriggerInstruction = tmrand.Str(1025)

				return msg
			}(validMsgCreateModel()),
			err: validator.ErrFieldMaxLengthExceeded,
		},
		{
			name: "FactoryResetStepsInstruction length > 1024",
			msg: func(msg *MsgCreateModel) *MsgCreateModel {
				msg.FactoryResetStepsInstruction = tmrand.Str(1025)

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
			name: "UserManualUrl starts not with https:",
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
			name: "SupportUrl starts not with https:",
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
			name: "ProductUrl starts not with https:",
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
			name: "LsfUrl starts not with https:",
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
			name: "schemaVersion != 0",
			msg: func(msg *MsgCreateModel) *MsgCreateModel {
				msg.Creator = sample.AccAddress()
				msg.SchemaVersion = 5

				return msg
			}(validMsgCreateModel()),
			err: validator.ErrFieldEqualBoundViolated,
		},
		{
			name: "EnhancedSetupFlowOptions > 65535",
			msg: func(msg *MsgCreateModel) *MsgCreateModel {
				msg.Creator = sample.AccAddress()
				msg.EnhancedSetupFlowOptions = 65536

				return msg
			}(validMsgCreateModel()),
			err: validator.ErrFieldUpperBoundViolated,
		},
		{
			name: "EnhancedSetupFlowTCUrl starts not with https:",
			msg: func(msg *MsgCreateModel) *MsgCreateModel {
				msg.EnhancedSetupFlowOptions = 1
				msg.EnhancedSetupFlowTCUrl = "http://sampleflowurl.dclmodel"

				return msg
			}(validMsgCreateModel()),
			err: validator.ErrFieldNotValid,
		},
		{
			name: "EnhancedSetupFlowTCUrl starts with non-http:",
			msg: func(msg *MsgCreateModel) *MsgCreateModel {
				msg.EnhancedSetupFlowOptions = 1
				msg.EnhancedSetupFlowTCUrl = "ftp://sampleflowurl.dclmodel"

				return msg
			}(validMsgCreateModel()),
			err: validator.ErrFieldNotValid,
		},
		{
			name: "EnhancedSetupFlowTCUrl length > 256",
			msg: func(msg *MsgCreateModel) *MsgCreateModel {
				msg.EnhancedSetupFlowOptions = 1
				msg.EnhancedSetupFlowTCUrl = "https://sampleflowurl.dclauth/" + tmrand.Str(257-30) // length = 257

				return msg
			}(validMsgCreateModel()),
			err: validator.ErrFieldMaxLengthExceeded,
		},
		{
			name: "EnhancedSetupFlowTCUrl, EnhancedSetupFlowTCRevision, EnhancedSetupFlowTCDigest, EnhancedSetupFlowTCFileSize and MaintenanceUrl are omitted when EnhancedSetupFlowOptions&1 == 1",
			msg: func(msg *MsgCreateModel) *MsgCreateModel {
				msg.EnhancedSetupFlowOptions = 1
				msg.EnhancedSetupFlowTCUrl = ""
				msg.EnhancedSetupFlowTCRevision = 0
				msg.EnhancedSetupFlowTCDigest = ""
				msg.EnhancedSetupFlowTCFileSize = 0
				msg.MaintenanceUrl = ""

				return msg
			}(validMsgCreateModel()),
			err: validator.ErrRequiredFieldMissing,
		},
		{
			name: "EnhancedSetupFlowTCUrl is omitted when EnhancedSetupFlowOptions&1 == 1",
			msg: func(msg *MsgCreateModel) *MsgCreateModel {
				msg.EnhancedSetupFlowOptions = 1
				msg.EnhancedSetupFlowTCUrl = ""

				return msg
			}(validMsgCreateModel()),
			err: validator.ErrRequiredFieldMissing,
		},
		{
			name: "EnhancedSetupFlowTCRevision is omitted when EnhancedSetupFlowOptions&1 == 1",
			msg: func(msg *MsgCreateModel) *MsgCreateModel {
				msg.EnhancedSetupFlowOptions = 1
				msg.EnhancedSetupFlowTCRevision = 0

				return msg
			}(validMsgCreateModel()),
			err: validator.ErrRequiredFieldMissing,
		},
		{
			name: "EnhancedSetupFlowTCDigest is omitted when EnhancedSetupFlowOptions&1 == 1",
			msg: func(msg *MsgCreateModel) *MsgCreateModel {
				msg.EnhancedSetupFlowOptions = 1
				msg.EnhancedSetupFlowTCDigest = ""

				return msg
			}(validMsgCreateModel()),
			err: validator.ErrRequiredFieldMissing,
		},
		{
			name: "EnhancedSetupFlowTCFileSize is omitted when EnhancedSetupFlowOptions&1 == 1",
			msg: func(msg *MsgCreateModel) *MsgCreateModel {
				msg.EnhancedSetupFlowOptions = 1
				msg.EnhancedSetupFlowTCFileSize = 0

				return msg
			}(validMsgCreateModel()),
			err: validator.ErrRequiredFieldMissing,
		},
		{
			name: "MaintenanceUrl is omitted when EnhancedSetupFlowOptions&1 == 1",
			msg: func(msg *MsgCreateModel) *MsgCreateModel {
				msg.EnhancedSetupFlowOptions = 1
				msg.MaintenanceUrl = ""

				return msg
			}(validMsgCreateModel()),
			err: validator.ErrRequiredFieldMissing,
		},
		{
			name: "EnhancedSetupFlowTCUrl, EnhancedSetupFlowTCRevision, EnhancedSetupFlowTCDigest, EnhancedSetupFlowTCFileSize and MaintenanceUrl are specified when EnhancedSetupFlowOptions&1 == 0",
			msg: func(msg *MsgCreateModel) *MsgCreateModel {
				msg.EnhancedSetupFlowOptions = 0

				return msg
			}(validMsgCreateModel()),
			err: validator.ErrRequiredFieldMissing,
		},
		{
			name: "EnhancedSetupFlowTCDigest is not base64 encoded string",
			msg: func(msg *MsgCreateModel) *MsgCreateModel {
				msg.EnhancedSetupFlowOptions = 1
				msg.EnhancedSetupFlowTCFileSize = 1
				msg.EnhancedSetupFlowTCRevision = 1
				msg.EnhancedSetupFlowTCUrl = "https://sampleflowurl.dclauth/"
				msg.MaintenanceUrl = "https://sampleflowurl.dclmodel"
				msg.EnhancedSetupFlowTCDigest = "--"

				return msg
			}(validMsgCreateModel()),
			err: ErrFieldIsNotBase64Encoded,
		},
		{
			name: "MaintenanceUrl starts not with https:",
			msg: func(msg *MsgCreateModel) *MsgCreateModel {
				msg.EnhancedSetupFlowOptions = 1
				msg.EnhancedSetupFlowTCFileSize = 1
				msg.EnhancedSetupFlowTCRevision = 1
				msg.EnhancedSetupFlowTCUrl = "https://sampleflowurl.dclauth/"
				msg.EnhancedSetupFlowTCDigest = "=="
				msg.MaintenanceUrl = "http://sampleflowurl.dclmodel"

				return msg
			}(validMsgCreateModel()),
			err: validator.ErrFieldNotValid,
		},
		{
			name: "MaintenanceUrl starts with non-http",
			msg: func(msg *MsgCreateModel) *MsgCreateModel {
				msg.EnhancedSetupFlowOptions = 1
				msg.EnhancedSetupFlowTCFileSize = 1
				msg.EnhancedSetupFlowTCRevision = 1
				msg.EnhancedSetupFlowTCUrl = "https://sampleflowurl.dclauth/"
				msg.EnhancedSetupFlowTCDigest = "=="
				msg.MaintenanceUrl = "ftp://sampleflowurl.dclmodel"

				return msg
			}(validMsgCreateModel()),
			err: validator.ErrFieldNotValid,
		},
		{
			name: "MaintenanceUrl length > 256",
			msg: func(msg *MsgCreateModel) *MsgCreateModel {
				msg.EnhancedSetupFlowOptions = 1
				msg.EnhancedSetupFlowTCFileSize = 1
				msg.EnhancedSetupFlowTCRevision = 1
				msg.EnhancedSetupFlowTCUrl = "https://sampleflowurl.dclauth/"
				msg.EnhancedSetupFlowTCDigest = "=="
				msg.MaintenanceUrl = "https://sampleflowurl.dclauth/" + tmrand.Str(257-30) // length = 257

				return msg
			}(validMsgCreateModel()),
			err: validator.ErrFieldMaxLengthExceeded,
		},
		{
			name: "CommissioningCustomFlowUrl can't start with non-https URLs",
			msg: func(msg *MsgCreateModel) *MsgCreateModel {
				msg.CommissioningCustomFlowUrl = "ftp://sampleflowurl.dclmodel"

				return msg
			}(validMsgCreateModel()),
			err: validator.ErrFieldNotValid,
		},
		{
			name: "UserManualUrl can't start with non-https URLs",
			msg: func(msg *MsgCreateModel) *MsgCreateModel {
				msg.UserManualUrl = "ftp://sampleflowurl.dclmodel"

				return msg
			}(validMsgCreateModel()),
			err: validator.ErrFieldNotValid,
		},
		{
			name: "SupportUrl can't start with non-https URLs",
			msg: func(msg *MsgCreateModel) *MsgCreateModel {
				msg.SupportUrl = "ftp://sampleflowurl.dclmodel"

				return msg
			}(validMsgCreateModel()),
			err: validator.ErrFieldNotValid,
		},
		{
			name: "ProductUrl can't start with non-https URLs",
			msg: func(msg *MsgCreateModel) *MsgCreateModel {
				msg.ProductUrl = "ftp://sampleflowurl.dclmodel"

				return msg
			}(validMsgCreateModel()),
			err: validator.ErrFieldNotValid,
		},
		{
			name: "LsfUrl can't start with non-https URLs",
			msg: func(msg *MsgCreateModel) *MsgCreateModel {
				msg.LsfUrl = "ftp://sampleflowurl.dclmodel"

				return msg
			}(validMsgCreateModel()),
			err: validator.ErrFieldNotValid,
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
			name: "DiscoveryCapabilitiesBitmask == 0",
			msg: func(msg *MsgCreateModel) *MsgCreateModel {
				msg.DiscoveryCapabilitiesBitmask = 0

				return msg
			}(validMsgCreateModel()),
		},
		{
			name: "DiscoveryCapabilitiesBitmask == 14",
			msg: func(msg *MsgCreateModel) *MsgCreateModel {
				msg.DiscoveryCapabilitiesBitmask = 14

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
			name: "IcdUserActiveModeTriggerHint == 0",
			msg: func(msg *MsgCreateModel) *MsgCreateModel {
				msg.IcdUserActiveModeTriggerHint = 0

				return msg
			}(validMsgCreateModel()),
		},
		{
			name: "FactoryResetStepsHint == 0",
			msg: func(msg *MsgCreateModel) *MsgCreateModel {
				msg.FactoryResetStepsHint = 0

				return msg
			}(validMsgCreateModel()),
		},
		{
			name: "IcdUserActiveModeTriggerHint > 0",
			msg: func(msg *MsgCreateModel) *MsgCreateModel {
				msg.IcdUserActiveModeTriggerHint = 1

				return msg
			}(validMsgCreateModel()),
		},
		{
			name: "FactoryResetStepsHint > 0",
			msg: func(msg *MsgCreateModel) *MsgCreateModel {
				msg.FactoryResetStepsHint = 1

				return msg
			}(validMsgCreateModel()),
		},
		{
			name: "IcdUserActiveModeTriggerInstruction is omitted",
			msg: func(msg *MsgCreateModel) *MsgCreateModel {
				msg.IcdUserActiveModeTriggerInstruction = ""

				return msg
			}(validMsgCreateModel()),
		},
		{
			name: "FactoryResetStepsInstruction is omitted",
			msg: func(msg *MsgCreateModel) *MsgCreateModel {
				msg.FactoryResetStepsInstruction = ""

				return msg
			}(validMsgCreateModel()),
		},
		{
			name: "IcdUserActiveModeTriggerInstruction length == 1024",
			msg: func(msg *MsgCreateModel) *MsgCreateModel {
				msg.IcdUserActiveModeTriggerInstruction = tmrand.Str(1024)

				return msg
			}(validMsgCreateModel()),
		},
		{
			name: "FactoryResetStepsInstruction length == 1024",
			msg: func(msg *MsgCreateModel) *MsgCreateModel {
				msg.FactoryResetStepsInstruction = tmrand.Str(1024)

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
		{
			name: "EnhancedSetupFlowOptions&1 == 0 and EnhancedSetupFlowTCUrl, EnhancedSetupFlowTCRevision, EnhancedSetupFlowTCDigest, EnhancedSetupFlowTCFileSize and MaintenanceUrl are omitted",
			msg: func(msg *MsgCreateModel) *MsgCreateModel {
				msg.EnhancedSetupFlowOptions = 0
				msg.EnhancedSetupFlowTCUrl = ""
				msg.EnhancedSetupFlowTCRevision = 0
				msg.EnhancedSetupFlowTCDigest = ""
				msg.EnhancedSetupFlowTCFileSize = 0
				msg.MaintenanceUrl = ""

				return msg
			}(validMsgCreateModel()),
		},
		{
			name: "EnhancedSetupFlowOptions&1 == 1 and EnhancedSetupFlowTCUrl, EnhancedSetupFlowTCRevision, EnhancedSetupFlowTCDigest, EnhancedSetupFlowTCFileSize and MaintenanceUrl are valid",
			msg: func(msg *MsgCreateModel) *MsgCreateModel {
				msg.EnhancedSetupFlowOptions = 1
				msg.EnhancedSetupFlowTCUrl = "https://sampleflowurl.dclauth/" + tmrand.Str(256-30) // length = 256
				msg.EnhancedSetupFlowTCRevision = 1
				msg.EnhancedSetupFlowTCDigest = "MWRjNGE0NDA0MWRjYWYxMTU0NWI3NTQzZGZlOTQyZjQ3NDJmNTY4YmU2OGZlZTI3NTQ0MWIwOTJiYjYwZGVlZA=="
				msg.EnhancedSetupFlowTCFileSize = 1024
				msg.MaintenanceUrl = "https://sampleflowurl.dclauth/" + tmrand.Str(256-30) // length = 256

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
			name: "CommissioningCustomFlowUrl starts not with https:",
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
			name: "IcdUserActiveModeTriggerInstruction length > 1024",
			msg: func(msg *MsgUpdateModel) *MsgUpdateModel {
				msg.IcdUserActiveModeTriggerInstruction = tmrand.Str(1025)

				return msg
			}(validMsgUpdateModel()),
			err: validator.ErrFieldMaxLengthExceeded,
		},
		{
			name: "FactoryResetStepsInstruction length > 1024",
			msg: func(msg *MsgUpdateModel) *MsgUpdateModel {
				msg.FactoryResetStepsInstruction = tmrand.Str(1025)

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
			name: "UserManualUrl starts not with https:",
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
			name: "SupportUrl starts not with https:",
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
			name: "ProductUrl starts not with https:",
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
		{
			name: "LsfUrl is not valid URL",
			msg: func(msg *MsgUpdateModel) *MsgUpdateModel {
				msg.LsfUrl = "not valid URL"

				return msg
			}(validMsgUpdateModel()),
			err: validator.ErrFieldNotValid,
		},
		{
			name: "LsfUrl starts not with https:",
			msg: func(msg *MsgUpdateModel) *MsgUpdateModel {
				msg.LsfUrl = "http://sampleflowurl.dclmodel"

				return msg
			}(validMsgUpdateModel()),
			err: validator.ErrFieldNotValid,
		},
		{
			name: "LsfUrl length > 256",
			msg: func(msg *MsgUpdateModel) *MsgUpdateModel {
				msg.LsfUrl = "https://sampleflowurl.dclauth/" + tmrand.Str(257-30) // length = 257

				return msg
			}(validMsgUpdateModel()),
			err: validator.ErrFieldMaxLengthExceeded,
		},
		{
			name: "LsfRevision is negative",
			msg: func(msg *MsgUpdateModel) *MsgUpdateModel {
				msg.LsfRevision = -1

				return msg
			}(validMsgUpdateModel()),
			err: validator.ErrFieldLowerBoundViolated,
		},
		{
			name: "LsfRevision is greater then max uint16",
			msg: func(msg *MsgUpdateModel) *MsgUpdateModel {
				msg.LsfRevision = 65536

				return msg
			}(validMsgUpdateModel()),
			err: validator.ErrFieldUpperBoundViolated,
		},
		{
			name: "schemaVersion != 0",
			msg: func(msg *MsgUpdateModel) *MsgUpdateModel {
				msg.Creator = sample.AccAddress()
				msg.SchemaVersion = 5

				return msg
			}(validMsgUpdateModel()),
			err: validator.ErrFieldEqualBoundViolated,
		},
		{
			name: "EnhancedSetupFlowOptions > 65535",
			msg: func(msg *MsgUpdateModel) *MsgUpdateModel {
				msg.EnhancedSetupFlowOptions = 65536

				return msg
			}(validMsgUpdateModel()),
			err: validator.ErrFieldUpperBoundViolated,
		},
		{
			name: "EnhancedSetupFlowTCUrl starts not with https:",
			msg: func(msg *MsgUpdateModel) *MsgUpdateModel {
				msg.EnhancedSetupFlowOptions = 1
				msg.EnhancedSetupFlowTCUrl = "http://sampleflowurl.dclmodel"

				return msg
			}(validMsgUpdateModel()),
			err: validator.ErrFieldNotValid,
		},
		{
			name: "EnhancedSetupFlowTCUrl starts with non-http",
			msg: func(msg *MsgUpdateModel) *MsgUpdateModel {
				msg.EnhancedSetupFlowOptions = 1
				msg.EnhancedSetupFlowTCUrl = "ftp://sampleflowurl.dclmodel"

				return msg
			}(validMsgUpdateModel()),
			err: validator.ErrFieldNotValid,
		},
		{
			name: "EnhancedSetupFlowTCUrl length > 256",
			msg: func(msg *MsgUpdateModel) *MsgUpdateModel {
				msg.EnhancedSetupFlowOptions = 1
				msg.EnhancedSetupFlowTCUrl = "https://sampleflowurl.dclauth/" + tmrand.Str(257-30) // length = 257

				return msg
			}(validMsgUpdateModel()),
			err: validator.ErrFieldMaxLengthExceeded,
		},
		{
			name: "EnhancedSetupFlowTCUrl, EnhancedSetupFlowTCRevision, EnhancedSetupFlowTCDigest, EnhancedSetupFlowTCFileSize and MaintenanceUrl are omitted when EnhancedSetupFlowOptions&1 == 1",
			msg: func(msg *MsgUpdateModel) *MsgUpdateModel {
				msg.EnhancedSetupFlowOptions = 1
				msg.EnhancedSetupFlowTCUrl = ""
				msg.EnhancedSetupFlowTCRevision = 0
				msg.EnhancedSetupFlowTCDigest = ""
				msg.EnhancedSetupFlowTCFileSize = 0
				msg.MaintenanceUrl = ""

				return msg
			}(validMsgUpdateModel()),
			err: validator.ErrRequiredFieldMissing,
		},
		{
			name: "EnhancedSetupFlowTCUrl is omitted when EnhancedSetupFlowOptions&1 == 1",
			msg: func(msg *MsgUpdateModel) *MsgUpdateModel {
				msg.EnhancedSetupFlowOptions = 1
				msg.EnhancedSetupFlowTCUrl = ""

				return msg
			}(validMsgUpdateModel()),
			err: validator.ErrRequiredFieldMissing,
		},
		{
			name: "EnhancedSetupFlowTCRevision is omitted when EnhancedSetupFlowOptions&1 == 1",
			msg: func(msg *MsgUpdateModel) *MsgUpdateModel {
				msg.EnhancedSetupFlowOptions = 1
				msg.EnhancedSetupFlowTCRevision = 0

				return msg
			}(validMsgUpdateModel()),
			err: validator.ErrRequiredFieldMissing,
		},
		{
			name: "EnhancedSetupFlowTCDigest is omitted when EnhancedSetupFlowOptions&1 == 1",
			msg: func(msg *MsgUpdateModel) *MsgUpdateModel {
				msg.EnhancedSetupFlowOptions = 1
				msg.EnhancedSetupFlowTCDigest = ""

				return msg
			}(validMsgUpdateModel()),
			err: validator.ErrRequiredFieldMissing,
		},
		{
			name: "EnhancedSetupFlowTCFileSize is omitted when EnhancedSetupFlowOptions&1 == 1",
			msg: func(msg *MsgUpdateModel) *MsgUpdateModel {
				msg.EnhancedSetupFlowOptions = 1
				msg.EnhancedSetupFlowTCFileSize = 0

				return msg
			}(validMsgUpdateModel()),
			err: validator.ErrRequiredFieldMissing,
		},
		{
			name: "MaintenanceUrl is omitted when EnhancedSetupFlowOptions&1 == 1",
			msg: func(msg *MsgUpdateModel) *MsgUpdateModel {
				msg.EnhancedSetupFlowOptions = 1
				msg.MaintenanceUrl = ""

				return msg
			}(validMsgUpdateModel()),
			err: validator.ErrRequiredFieldMissing,
		},
		{
			name: "EnhancedSetupFlowTCUrl, EnhancedSetupFlowTCRevision, EnhancedSetupFlowTCDigest, EnhancedSetupFlowTCFileSize and MaintenanceUrl are specified when EnhancedSetupFlowOptions&1 == 0",
			msg: func(msg *MsgUpdateModel) *MsgUpdateModel {
				msg.EnhancedSetupFlowOptions = 0

				return msg
			}(validMsgUpdateModel()),
			err: validator.ErrRequiredFieldMissing,
		},
		{
			name: "EnhancedSetupFlowOptions = 0, but EnhancedSetupFlowTCUrl is set",
			msg: func(msg *MsgUpdateModel) *MsgUpdateModel {
				msg.EnhancedSetupFlowOptions = 0
				msg.EnhancedSetupFlowTCUrl = "https://sampleflowurl.dclauth/"

				return msg
			}(validMsgUpdateModel()),
			err: validator.ErrRequiredFieldMissing,
		},
		{
			name: "EnhancedSetupFlowOptions = 0, but EnhancedSetupFlowTCFileSize is set",
			msg: func(msg *MsgUpdateModel) *MsgUpdateModel {
				msg.EnhancedSetupFlowOptions = 0
				msg.EnhancedSetupFlowTCFileSize = 1

				return msg
			}(validMsgUpdateModel()),
			err: validator.ErrRequiredFieldMissing,
		},
		{
			name: "EnhancedSetupFlowOptions = 0, but EnhancedSetupFlowTCRevision is set",
			msg: func(msg *MsgUpdateModel) *MsgUpdateModel {
				msg.EnhancedSetupFlowOptions = 0
				msg.EnhancedSetupFlowTCRevision = 1

				return msg
			}(validMsgUpdateModel()),
			err: validator.ErrRequiredFieldMissing,
		},
		{
			name: "EnhancedSetupFlowOptions = 0, but MaintenanceUrl is set",
			msg: func(msg *MsgUpdateModel) *MsgUpdateModel {
				msg.EnhancedSetupFlowOptions = 0
				msg.MaintenanceUrl = "https://sampleflowurl.dclmodel"

				return msg
			}(validMsgUpdateModel()),
			err: validator.ErrRequiredFieldMissing,
		},
		{
			name: "EnhancedSetupFlowOptions = 0, but EnhancedSetupFlowTCDigest is set",
			msg: func(msg *MsgUpdateModel) *MsgUpdateModel {
				msg.EnhancedSetupFlowOptions = 0
				msg.EnhancedSetupFlowTCDigest = "--"

				return msg
			}(validMsgUpdateModel()),
			err: validator.ErrRequiredFieldMissing,
		},
		{
			name: "EnhancedSetupFlowTCDigest is not base64 encoded string",
			msg: func(msg *MsgUpdateModel) *MsgUpdateModel {
				msg.EnhancedSetupFlowOptions = 1
				msg.EnhancedSetupFlowTCFileSize = 1
				msg.EnhancedSetupFlowTCRevision = 1
				msg.EnhancedSetupFlowTCUrl = "https://sampleflowurl.dclauth/"
				msg.MaintenanceUrl = "https://sampleflowurl.dclmodel"
				msg.EnhancedSetupFlowTCDigest = "--"

				return msg
			}(validMsgUpdateModel()),
			err: ErrFieldIsNotBase64Encoded,
		},
		{
			name: "MaintenanceUrl starts not with https:",
			msg: func(msg *MsgUpdateModel) *MsgUpdateModel {
				msg.EnhancedSetupFlowOptions = 1
				msg.EnhancedSetupFlowTCFileSize = 1
				msg.EnhancedSetupFlowTCRevision = 1
				msg.EnhancedSetupFlowTCUrl = "https://sampleflowurl.dclauth/"
				msg.EnhancedSetupFlowTCDigest = "=="
				msg.MaintenanceUrl = "http://sampleflowurl.dclmodel"

				return msg
			}(validMsgUpdateModel()),
			err: validator.ErrFieldNotValid,
		},
		{
			name: "MaintenanceUrl starts with non-http:",
			msg: func(msg *MsgUpdateModel) *MsgUpdateModel {
				msg.EnhancedSetupFlowOptions = 1
				msg.EnhancedSetupFlowTCFileSize = 1
				msg.EnhancedSetupFlowTCRevision = 1
				msg.EnhancedSetupFlowTCUrl = "https://sampleflowurl.dclauth/"
				msg.EnhancedSetupFlowTCDigest = "=="
				msg.MaintenanceUrl = "ftp://sampleflowurl.dclmodel"

				return msg
			}(validMsgUpdateModel()),
			err: validator.ErrFieldNotValid,
		},
		{
			name: "MaintenanceUrl length > 256",
			msg: func(msg *MsgUpdateModel) *MsgUpdateModel {
				msg.EnhancedSetupFlowOptions = 1
				msg.EnhancedSetupFlowTCFileSize = 1
				msg.EnhancedSetupFlowTCRevision = 1
				msg.EnhancedSetupFlowTCUrl = "https://sampleflowurl.dclauth/"
				msg.EnhancedSetupFlowTCDigest = "=="
				msg.MaintenanceUrl = "https://sampleflowurl.dclauth/" + tmrand.Str(257-30) // length = 257

				return msg
			}(validMsgUpdateModel()),
			err: validator.ErrFieldMaxLengthExceeded,
		},
		{
			name: "CommissioningCustomFlowUrl can't start with non-https URLs",
			msg: func(msg *MsgUpdateModel) *MsgUpdateModel {
				msg.CommissioningCustomFlowUrl = "ftp://sampleflowurl.dclmodel"

				return msg
			}(validMsgUpdateModel()),
			err: validator.ErrFieldNotValid,
		},
		{
			name: "UserManualUrl can't start with non-https URLs",
			msg: func(msg *MsgUpdateModel) *MsgUpdateModel {
				msg.UserManualUrl = "ftp://sampleflowurl.dclmodel"

				return msg
			}(validMsgUpdateModel()),
			err: validator.ErrFieldNotValid,
		},
		{
			name: "SupportUrl can't start with non-https URLs",
			msg: func(msg *MsgUpdateModel) *MsgUpdateModel {
				msg.SupportUrl = "ftp://sampleflowurl.dclmodel"

				return msg
			}(validMsgUpdateModel()),
			err: validator.ErrFieldNotValid,
		},
		{
			name: "ProductUrl can't start with non-https URLs",
			msg: func(msg *MsgUpdateModel) *MsgUpdateModel {
				msg.ProductUrl = "ftp://sampleflowurl.dclmodel"

				return msg
			}(validMsgUpdateModel()),
			err: validator.ErrFieldNotValid,
		},
		{
			name: "LsfUrl can't start with non-https URLs",
			msg: func(msg *MsgUpdateModel) *MsgUpdateModel {
				msg.LsfUrl = "ftp://sampleflowurl.dclmodel"

				return msg
			}(validMsgUpdateModel()),
			err: validator.ErrFieldNotValid,
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
			name: "IcdUserActiveModeTriggerInstruction is omitted",
			msg: func(msg *MsgUpdateModel) *MsgUpdateModel {
				msg.IcdUserActiveModeTriggerInstruction = ""

				return msg
			}(validMsgUpdateModel()),
		},
		{
			name: "FactoryResetStepsInstruction is omitted",
			msg: func(msg *MsgUpdateModel) *MsgUpdateModel {
				msg.FactoryResetStepsInstruction = ""

				return msg
			}(validMsgUpdateModel()),
		},
		{
			name: "IcdUserActiveModeTriggerInstruction length == 1024",
			msg: func(msg *MsgUpdateModel) *MsgUpdateModel {
				msg.IcdUserActiveModeTriggerInstruction = tmrand.Str(1024)

				return msg
			}(validMsgUpdateModel()),
		},
		{
			name: "FactoryResetStepsInstruction length == 1024",
			msg: func(msg *MsgUpdateModel) *MsgUpdateModel {
				msg.FactoryResetStepsInstruction = tmrand.Str(1024)

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
		{
			name: "LsfUrl is omitted and LsfRevision is omitted 1",
			msg: func(msg *MsgUpdateModel) *MsgUpdateModel {
				msg.LsfUrl = ""
				msg.LsfRevision = 0

				return msg
			}(validMsgUpdateModel()),
		},

		{
			name: "LsfUrl length == 256",
			msg: func(msg *MsgUpdateModel) *MsgUpdateModel {
				msg.LsfUrl = "https://sampleflowurl.dclauth/" + tmrand.Str(256-30) // length = 256

				return msg
			}(validMsgUpdateModel()),
		},
		{
			name: "LsfRevision is set to 65535",
			msg: func(msg *MsgUpdateModel) *MsgUpdateModel {
				msg.LsfRevision = 65535

				return msg
			}(validMsgUpdateModel()),
		},
		{
			name: "EnhancedSetupFlowOptions is valid and EnhancedSetupFlowTCUrl, EnhancedSetupFlowTCRevision, EnhancedSetupFlowTCDigest, EnhancedSetupFlowTCFileSize and MaintenanceUrl are omitted",
			msg: func(msg *MsgUpdateModel) *MsgUpdateModel {
				return msg
			}(validMsgUpdateModel()),
		},
		{
			name: "EnhancedSetupFlowOptions&1 == 1 and EnhancedSetupFlowTCUrl, EnhancedSetupFlowTCRevision, EnhancedSetupFlowTCDigest, EnhancedSetupFlowTCFileSize and MaintenanceUrl are valid",
			msg: func(msg *MsgUpdateModel) *MsgUpdateModel {
				msg.EnhancedSetupFlowOptions = 1
				msg.EnhancedSetupFlowTCUrl = "https://sampleflowurl.dclauth/" + tmrand.Str(256-30) // length = 256
				msg.EnhancedSetupFlowTCRevision = 1
				msg.EnhancedSetupFlowTCDigest = "MWRjNGE0NDA0MWRjYWYxMTU0NWI3NTQzZGZlOTQyZjQ3NDJmNTY4YmU2OGZlZTI3NTQ0MWIwOTJiYjYwZGVlZA=="
				msg.EnhancedSetupFlowTCFileSize = 1024
				msg.MaintenanceUrl = "https://sampleflowurl.dclauth/" + tmrand.Str(256-30) // length = 256

				return msg
			}(validMsgUpdateModel()),
		},
		{
			name: "EnhancedSetupFlowOptions&1 == 0 and EnhancedSetupFlowTCUrl, EnhancedSetupFlowTCRevision, EnhancedSetupFlowTCDigest, EnhancedSetupFlowTCFileSize and MaintenanceUrl are not set",
			msg: func(msg *MsgUpdateModel) *MsgUpdateModel {
				msg.EnhancedSetupFlowOptions = 0
				msg.EnhancedSetupFlowTCUrl = ""
				msg.EnhancedSetupFlowTCRevision = 0
				msg.EnhancedSetupFlowTCDigest = ""
				msg.EnhancedSetupFlowTCFileSize = 0
				msg.MaintenanceUrl = ""

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
		DeviceTypeId:                             testconstants.DeviceTypeID,
		ProductName:                              testconstants.ProductName,
		ProductLabel:                             testconstants.ProductLabel,
		PartNumber:                               testconstants.PartNumber,
		CommissioningCustomFlow:                  testconstants.CommissioningCustomFlow,
		CommissioningCustomFlowUrl:               testconstants.CommissioningCustomFlowURL,
		CommissioningModeInitialStepsHint:        testconstants.CommissioningModeInitialStepsHint,
		CommissioningModeInitialStepsInstruction: testconstants.CommissioningModeInitialStepsInstruction,
		CommissioningModeSecondaryStepsHint:      testconstants.CommissioningModeSecondaryStepsHint,
		CommissioningModeSecondaryStepsInstruction: testconstants.CommissioningModeSecondaryStepsInstruction,
		IcdUserActiveModeTriggerHint:               testconstants.IcdUserActiveModeTriggerHint,
		IcdUserActiveModeTriggerInstruction:        testconstants.IcdUserActiveModeTriggerInstruction,
		FactoryResetStepsHint:                      testconstants.FactoryResetStepsHint,
		FactoryResetStepsInstruction:               testconstants.FactoryResetStepsInstruction,
		UserManualUrl:                              testconstants.UserManualURL,
		SupportUrl:                                 testconstants.SupportURL,
		ProductUrl:                                 testconstants.ProductURL,
		LsfUrl:                                     testconstants.LsfURL,
		EnhancedSetupFlowOptions:                   testconstants.EnhancedSetupFlowOptions,
		EnhancedSetupFlowTCUrl:                     testconstants.EnhancedSetupFlowTCURL,
		EnhancedSetupFlowTCRevision:                int32(testconstants.EnhancedSetupFlowTCRevision),
		EnhancedSetupFlowTCDigest:                  testconstants.EnhancedSetupFlowTCDigest,
		EnhancedSetupFlowTCFileSize:                uint32(testconstants.EnhancedSetupFlowTCFileSize),
		MaintenanceUrl:                             testconstants.MaintenanceURL,
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
		CommissioningCustomFlowUrl:               testconstants.CommissioningCustomFlowURL + "/updated",
		CommissioningModeInitialStepsInstruction: testconstants.CommissioningModeInitialStepsInstruction + "-updated",
		CommissioningModeSecondaryStepsInstruction: testconstants.CommissioningModeSecondaryStepsInstruction + "-updated",
		IcdUserActiveModeTriggerInstruction:        testconstants.IcdUserActiveModeTriggerInstruction + "-updated",
		FactoryResetStepsInstruction:               testconstants.FactoryResetStepsInstruction + "-updated",
		UserManualUrl:                              testconstants.UserManualURL + "/updated",
		SupportUrl:                                 testconstants.SupportURL + "/updated",
		ProductUrl:                                 testconstants.ProductURL + "/updated",
		LsfUrl:                                     testconstants.LsfURL + "/updated",
		LsfRevision:                                testconstants.LsfRevision + 1,
		EnhancedSetupFlowOptions:                   testconstants.EnhancedSetupFlowOptions + 2,
		EnhancedSetupFlowTCUrl:                     testconstants.EnhancedSetupFlowTCURL + "/updated",
		EnhancedSetupFlowTCRevision:                int32(testconstants.EnhancedSetupFlowTCRevision + 1),
		EnhancedSetupFlowTCDigest:                  testconstants.EnhancedSetupFlowTCDigest,
		EnhancedSetupFlowTCFileSize:                uint32(testconstants.EnhancedSetupFlowTCFileSize + 1),
		MaintenanceUrl:                             testconstants.MaintenanceURL + "/updated",
	}
}
