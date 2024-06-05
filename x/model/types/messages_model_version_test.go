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

func TestMsgCreateModelVersion_ValidateBasic(t *testing.T) {
	negativeTests := []struct {
		name string
		msg  *MsgCreateModelVersion
		err  error
	}{
		{
			name: "Creator is omitted",
			msg: func(msg *MsgCreateModelVersion) *MsgCreateModelVersion {
				msg.Creator = ""

				return msg
			}(validMsgCreateModelVersion()),
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "Creator is not valid address",
			msg: func(msg *MsgCreateModelVersion) *MsgCreateModelVersion {
				msg.Creator = "not valid address"

				return msg
			}(validMsgCreateModelVersion()),
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "Vid < 0",
			msg: func(msg *MsgCreateModelVersion) *MsgCreateModelVersion {
				msg.Vid = -1

				return msg
			}(validMsgCreateModelVersion()),
			err: validator.ErrFieldLowerBoundViolated,
		},
		{
			name: "Vid == 0",
			msg: func(msg *MsgCreateModelVersion) *MsgCreateModelVersion {
				msg.Vid = 0

				return msg
			}(validMsgCreateModelVersion()),
			err: validator.ErrFieldLowerBoundViolated,
		},
		{
			name: "Vid > 65535",
			msg: func(msg *MsgCreateModelVersion) *MsgCreateModelVersion {
				msg.Vid = 65536

				return msg
			}(validMsgCreateModelVersion()),
			err: validator.ErrFieldUpperBoundViolated,
		},
		{
			name: "Pid < 0",
			msg: func(msg *MsgCreateModelVersion) *MsgCreateModelVersion {
				msg.Pid = -1

				return msg
			}(validMsgCreateModelVersion()),
			err: validator.ErrFieldLowerBoundViolated,
		},
		{
			name: "Pid == 0",
			msg: func(msg *MsgCreateModelVersion) *MsgCreateModelVersion {
				msg.Pid = 0

				return msg
			}(validMsgCreateModelVersion()),
			err: validator.ErrFieldLowerBoundViolated,
		},
		{
			name: "Pid > 65535",
			msg: func(msg *MsgCreateModelVersion) *MsgCreateModelVersion {
				msg.Pid = 65536

				return msg
			}(validMsgCreateModelVersion()),
			err: validator.ErrFieldUpperBoundViolated,
		},
		{
			name: "SoftwareVersionString is omitted",
			msg: func(msg *MsgCreateModelVersion) *MsgCreateModelVersion {
				msg.SoftwareVersionString = ""

				return msg
			}(validMsgCreateModelVersion()),
			err: validator.ErrRequiredFieldMissing,
		},
		{
			name: "SoftwareVersionString length > 64",
			msg: func(msg *MsgCreateModelVersion) *MsgCreateModelVersion {
				msg.SoftwareVersionString = tmrand.Str(65)

				return msg
			}(validMsgCreateModelVersion()),
			err: validator.ErrFieldMaxLengthExceeded,
		},
		{
			name: "CdVersionNumber < 0",
			msg: func(msg *MsgCreateModelVersion) *MsgCreateModelVersion {
				msg.CdVersionNumber = -1

				return msg
			}(validMsgCreateModelVersion()),
			err: validator.ErrFieldLowerBoundViolated,
		},
		{
			name: "CdVersionNumber > 65535",
			msg: func(msg *MsgCreateModelVersion) *MsgCreateModelVersion {
				msg.CdVersionNumber = 65536

				return msg
			}(validMsgCreateModelVersion()),
			err: validator.ErrFieldUpperBoundViolated,
		},
		{
			name: "FirmwareInformation length > 512",
			msg: func(msg *MsgCreateModelVersion) *MsgCreateModelVersion {
				msg.FirmwareInformation = tmrand.Str(513)

				return msg
			}(validMsgCreateModelVersion()),
			err: validator.ErrFieldMaxLengthExceeded,
		},
		{
			name: "OtaUrl is not valid URL",
			msg: func(msg *MsgCreateModelVersion) *MsgCreateModelVersion {
				msg.OtaUrl = "not valid URL"

				return msg
			}(validMsgCreateModelVersion()),
			err: validator.ErrFieldNotValid,
		},
		{
			name: "OtaUrl starts with http:",
			msg: func(msg *MsgCreateModelVersion) *MsgCreateModelVersion {
				msg.OtaUrl = "http://sampleflowurl.dclmodel"

				return msg
			}(validMsgCreateModelVersion()),
			err: validator.ErrFieldNotValid,
		},
		{
			name: "OtaUrl length > 256",
			msg: func(msg *MsgCreateModelVersion) *MsgCreateModelVersion {
				msg.OtaUrl = "https://sampleflowurl.dclauth/" + tmrand.Str(257-30) // length = 257

				return msg
			}(validMsgCreateModelVersion()),
			err: validator.ErrFieldMaxLengthExceeded,
		},
		{
			name: "OtaFileSize == 0 when OtaUrl is set",
			msg: func(msg *MsgCreateModelVersion) *MsgCreateModelVersion {
				msg.OtaUrl = "https://sampleflowurl.dclmodel"
				msg.OtaFileSize = 0

				return msg
			}(validMsgCreateModelVersion()),
			err: validator.ErrRequiredFieldMissing,
		},
		{
			name: "OtaChecksum is omitted when OtaUrl is set",
			msg: func(msg *MsgCreateModelVersion) *MsgCreateModelVersion {
				msg.OtaUrl = "https://sampleflowurl.dclmodel"
				msg.OtaChecksum = ""

				return msg
			}(validMsgCreateModelVersion()),
			err: validator.ErrRequiredFieldMissing,
		},
		{
			name: "OtaChecksum length > 64",
			msg: func(msg *MsgCreateModelVersion) *MsgCreateModelVersion {
				msg.OtaChecksum = "SGVsbG8gd29ybGQhSGVsbG8gd29ybGQhSGVsbG8gd29ybGQhSGVsbG8gd29ybGQhSGVsbG8gd29ybGQhSGVsbG8gd29ybGQhSGVsbG8gd29ybGQh"

				return msg
			}(validMsgCreateModelVersion()),
			err: validator.ErrFieldMaxLengthExceeded,
		},
		{
			name: "OtaChecksum is not base64 encoded",
			msg: func(msg *MsgCreateModelVersion) *MsgCreateModelVersion {
				msg.OtaUrl = "https://sampleflowurl.dclmodel"
				msg.OtaChecksum = "not_base64_encoded"

				return msg
			}(validMsgCreateModelVersion()),
			err: ErrOtaChecksumIsNotValid,
		},
		{
			name: "OtaChecksumType == 0 when OtaUrl is set",
			msg: func(msg *MsgCreateModelVersion) *MsgCreateModelVersion {
				msg.OtaUrl = "https://sampleflowurl.dclmodel"
				msg.OtaChecksumType = 0

				return msg
			}(validMsgCreateModelVersion()),
			err: validator.ErrRequiredFieldMissing,
		},
		{
			name: "OtaChecksumType < 0",
			msg: func(msg *MsgCreateModelVersion) *MsgCreateModelVersion {
				msg.OtaChecksumType = -1

				return msg
			}(validMsgCreateModelVersion()),
			err: validator.ErrFieldLowerBoundViolated,
		},
		{
			name: "OtaChecksumType > 65535",
			msg: func(msg *MsgCreateModelVersion) *MsgCreateModelVersion {
				msg.OtaChecksumType = 65536

				return msg
			}(validMsgCreateModelVersion()),
			err: validator.ErrFieldUpperBoundViolated,
		},
		{
			name: "MinApplicableSoftwareVersion > MaxApplicableSoftwareVersion " +
				"and MaxApplicableSoftwareVersion == 0",
			msg: func(msg *MsgCreateModelVersion) *MsgCreateModelVersion {
				msg.MinApplicableSoftwareVersion = 1
				msg.MaxApplicableSoftwareVersion = 0

				return msg
			}(validMsgCreateModelVersion()),
			err: validator.ErrFieldNotValid,
		},
		{
			name: "MinApplicableSoftwareVersion > MaxApplicableSoftwareVersion " +
				"and MaxApplicableSoftwareVersion > 0",
			msg: func(msg *MsgCreateModelVersion) *MsgCreateModelVersion {
				msg.MinApplicableSoftwareVersion = 8
				msg.MaxApplicableSoftwareVersion = 7

				return msg
			}(validMsgCreateModelVersion()),
			err: validator.ErrFieldNotValid,
		},
		{
			name: "ReleaseNotesUrl is not valid URL",
			msg: func(msg *MsgCreateModelVersion) *MsgCreateModelVersion {
				msg.ReleaseNotesUrl = "not valid URL"

				return msg
			}(validMsgCreateModelVersion()),
			err: validator.ErrFieldNotValid,
		},
		{
			name: "ReleaseNotesUrl starts with http:",
			msg: func(msg *MsgCreateModelVersion) *MsgCreateModelVersion {
				msg.ReleaseNotesUrl = "http://sampleflowurl.dclmodel"

				return msg
			}(validMsgCreateModelVersion()),
			err: validator.ErrFieldNotValid,
		},
		{
			name: "ReleaseNotesUrl length > 256",
			msg: func(msg *MsgCreateModelVersion) *MsgCreateModelVersion {
				msg.ReleaseNotesUrl = "https://sampleflowurl.dclauth/" + tmrand.Str(257-30) // length = 257

				return msg
			}(validMsgCreateModelVersion()),
			err: validator.ErrFieldMaxLengthExceeded,
		},
		{
			name: "schemaVersion > 65535",
			msg: func(msg *MsgCreateModelVersion) *MsgCreateModelVersion {
				msg.Creator = sample.AccAddress()
				msg.SchemaVersion = 65536

				return msg
			}(validMsgCreateModelVersion()),
			err: validator.ErrFieldUpperBoundViolated,
		},
	}

	positiveTests := []struct {
		name string
		msg  *MsgCreateModelVersion
	}{
		{
			name: "valid message",
			msg:  validMsgCreateModelVersion(),
		},
		{
			name: "Creator is valid address",
			msg: func(msg *MsgCreateModelVersion) *MsgCreateModelVersion {
				msg.Creator = sample.AccAddress()

				return msg
			}(validMsgCreateModelVersion()),
		},
		{
			name: "Vid == 1",
			msg: func(msg *MsgCreateModelVersion) *MsgCreateModelVersion {
				msg.Vid = 1

				return msg
			}(validMsgCreateModelVersion()),
		},
		{
			name: "Vid == 65535",
			msg: func(msg *MsgCreateModelVersion) *MsgCreateModelVersion {
				msg.Vid = 65535

				return msg
			}(validMsgCreateModelVersion()),
		},
		{
			name: "Pid == 1",
			msg: func(msg *MsgCreateModelVersion) *MsgCreateModelVersion {
				msg.Pid = 1

				return msg
			}(validMsgCreateModelVersion()),
		},
		{
			name: "Pid == 65535",
			msg: func(msg *MsgCreateModelVersion) *MsgCreateModelVersion {
				msg.Pid = 65535

				return msg
			}(validMsgCreateModelVersion()),
		},
		{
			name: "SoftwareVersion == 0",
			msg: func(msg *MsgCreateModelVersion) *MsgCreateModelVersion {
				msg.SoftwareVersion = 0

				return msg
			}(validMsgCreateModelVersion()),
		},
		{
			name: "SoftwareVersion > 0",
			msg: func(msg *MsgCreateModelVersion) *MsgCreateModelVersion {
				msg.SoftwareVersion = 1

				return msg
			}(validMsgCreateModelVersion()),
		},
		{
			name: "SoftwareVersionString length == 64",
			msg: func(msg *MsgCreateModelVersion) *MsgCreateModelVersion {
				msg.SoftwareVersionString = tmrand.Str(64)

				return msg
			}(validMsgCreateModelVersion()),
		},
		{
			name: "CdVersionNumber == 0",
			msg: func(msg *MsgCreateModelVersion) *MsgCreateModelVersion {
				msg.CdVersionNumber = 0

				return msg
			}(validMsgCreateModelVersion()),
		},
		{
			name: "CdVersionNumber == 65535",
			msg: func(msg *MsgCreateModelVersion) *MsgCreateModelVersion {
				msg.CdVersionNumber = 65535

				return msg
			}(validMsgCreateModelVersion()),
		},
		{
			name: "FirmwareInformation is omitted",
			msg: func(msg *MsgCreateModelVersion) *MsgCreateModelVersion {
				msg.FirmwareInformation = ""

				return msg
			}(validMsgCreateModelVersion()),
		},
		{
			name: "FirmwareInformation length == 512",
			msg: func(msg *MsgCreateModelVersion) *MsgCreateModelVersion {
				msg.FirmwareInformation = tmrand.Str(512)

				return msg
			}(validMsgCreateModelVersion()),
		},
		{
			name: "OtaUrl is omitted",
			msg: func(msg *MsgCreateModelVersion) *MsgCreateModelVersion {
				msg.OtaUrl = ""

				return msg
			}(validMsgCreateModelVersion()),
		},
		{
			name: "OtaUrl length == 256",
			msg: func(msg *MsgCreateModelVersion) *MsgCreateModelVersion {
				msg.OtaUrl = "https://sampleflowurl.dclauth/" + tmrand.Str(256-30) // length = 256

				return msg
			}(validMsgCreateModelVersion()),
		},
		{
			name: "OtaFileSize == 0 when OtaUrl is omitted",
			msg: func(msg *MsgCreateModelVersion) *MsgCreateModelVersion {
				msg.OtaUrl = ""
				msg.OtaFileSize = 0

				return msg
			}(validMsgCreateModelVersion()),
		},
		{
			name: "OtaFileSize > 0 when OtaUrl is set",
			msg: func(msg *MsgCreateModelVersion) *MsgCreateModelVersion {
				msg.OtaUrl = "https://sampleflowurl.dclmodel"
				msg.OtaFileSize = 1

				return msg
			}(validMsgCreateModelVersion()),
		},
		{
			name: "OtaChecksum is omitted when OtaUrl is omitted",
			msg: func(msg *MsgCreateModelVersion) *MsgCreateModelVersion {
				msg.OtaUrl = ""
				msg.OtaChecksum = ""

				return msg
			}(validMsgCreateModelVersion()),
		},
		{
			name: "OtaChecksum is set when OtaUrl is set",
			msg: func(msg *MsgCreateModelVersion) *MsgCreateModelVersion {
				msg.OtaUrl = "https://sampleflowurl.dclmodel"
				msg.OtaChecksum = "SGVsbG8gd29ybGQh"

				return msg
			}(validMsgCreateModelVersion()),
		},
		{
			name: "OtaChecksum length == 64",
			msg: func(msg *MsgCreateModelVersion) *MsgCreateModelVersion {
				msg.OtaChecksum = tmrand.Str(64)

				return msg
			}(validMsgCreateModelVersion()),
		},
		{
			name: "OtaChecksumType == 0 when OtaUrl is omitted",
			msg: func(msg *MsgCreateModelVersion) *MsgCreateModelVersion {
				msg.OtaUrl = ""
				msg.OtaChecksumType = 0

				return msg
			}(validMsgCreateModelVersion()),
		},
		{
			name: "OtaChecksumType != 0 when OtaUrl is set",
			msg: func(msg *MsgCreateModelVersion) *MsgCreateModelVersion {
				msg.OtaUrl = "https://sampleflowurl.dclmodel"
				msg.OtaChecksumType = 1

				return msg
			}(validMsgCreateModelVersion()),
		},
		{
			name: "OtaChecksumType == 1",
			msg: func(msg *MsgCreateModelVersion) *MsgCreateModelVersion {
				msg.OtaChecksumType = 1

				return msg
			}(validMsgCreateModelVersion()),
		},
		{
			name: "OtaChecksumType == 65535",
			msg: func(msg *MsgCreateModelVersion) *MsgCreateModelVersion {
				msg.OtaChecksumType = 65535

				return msg
			}(validMsgCreateModelVersion()),
		},
		{
			name: "MinApplicableSoftwareVersion == 0 and MaxApplicableSoftwareVersion == 0",
			msg: func(msg *MsgCreateModelVersion) *MsgCreateModelVersion {
				msg.MinApplicableSoftwareVersion = 0
				msg.MaxApplicableSoftwareVersion = 0

				return msg
			}(validMsgCreateModelVersion()),
		},
		{
			name: "MinApplicableSoftwareVersion == 0 and MaxApplicableSoftwareVersion > 0",
			msg: func(msg *MsgCreateModelVersion) *MsgCreateModelVersion {
				msg.MinApplicableSoftwareVersion = 0
				msg.MaxApplicableSoftwareVersion = 1

				return msg
			}(validMsgCreateModelVersion()),
		},
		{
			name: "MinApplicableSoftwareVersion > 0, MaxApplicableSoftwareVersion > 0 " +
				"and MinApplicableSoftwareVersion < MaxApplicableSoftwareVersion",
			msg: func(msg *MsgCreateModelVersion) *MsgCreateModelVersion {
				msg.MinApplicableSoftwareVersion = 5
				msg.MaxApplicableSoftwareVersion = 10

				return msg
			}(validMsgCreateModelVersion()),
		},
		{
			name: "MinApplicableSoftwareVersion > 0, MaxApplicableSoftwareVersion > 0 " +
				"and MinApplicableSoftwareVersion == MaxApplicableSoftwareVersion",
			msg: func(msg *MsgCreateModelVersion) *MsgCreateModelVersion {
				msg.MinApplicableSoftwareVersion = 7
				msg.MaxApplicableSoftwareVersion = 7

				return msg
			}(validMsgCreateModelVersion()),
		},
		{
			name: "ReleaseNotesUrl is omitted",
			msg: func(msg *MsgCreateModelVersion) *MsgCreateModelVersion {
				msg.ReleaseNotesUrl = ""

				return msg
			}(validMsgCreateModelVersion()),
		},
		{
			name: "ReleaseNotesUrl length == 256",
			msg: func(msg *MsgCreateModelVersion) *MsgCreateModelVersion {
				msg.ReleaseNotesUrl = "https://sampleflowurl.dclauth/" + tmrand.Str(256-30) // length = 256

				return msg
			}(validMsgCreateModelVersion()),
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

func TestMsgUpdateModelVersion_ValidateBasic(t *testing.T) {
	negativeTests := []struct {
		name string
		msg  *MsgUpdateModelVersion
		err  error
	}{
		{
			name: "Creator is omitted",
			msg: func(msg *MsgUpdateModelVersion) *MsgUpdateModelVersion {
				msg.Creator = ""

				return msg
			}(validMsgUpdateModelVersion()),
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "Creator is not valid address",
			msg: func(msg *MsgUpdateModelVersion) *MsgUpdateModelVersion {
				msg.Creator = "not valid address"

				return msg
			}(validMsgUpdateModelVersion()),
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "Vid < 0",
			msg: func(msg *MsgUpdateModelVersion) *MsgUpdateModelVersion {
				msg.Vid = -1

				return msg
			}(validMsgUpdateModelVersion()),
			err: validator.ErrFieldLowerBoundViolated,
		},
		{
			name: "Vid == 0",
			msg: func(msg *MsgUpdateModelVersion) *MsgUpdateModelVersion {
				msg.Vid = 0

				return msg
			}(validMsgUpdateModelVersion()),
			err: validator.ErrFieldLowerBoundViolated,
		},
		{
			name: "Vid > 65535",
			msg: func(msg *MsgUpdateModelVersion) *MsgUpdateModelVersion {
				msg.Vid = 65536

				return msg
			}(validMsgUpdateModelVersion()),
			err: validator.ErrFieldUpperBoundViolated,
		},
		{
			name: "Pid < 0",
			msg: func(msg *MsgUpdateModelVersion) *MsgUpdateModelVersion {
				msg.Pid = -1

				return msg
			}(validMsgUpdateModelVersion()),
			err: validator.ErrFieldLowerBoundViolated,
		},
		{
			name: "Pid == 0",
			msg: func(msg *MsgUpdateModelVersion) *MsgUpdateModelVersion {
				msg.Pid = 0

				return msg
			}(validMsgUpdateModelVersion()),
			err: validator.ErrFieldLowerBoundViolated,
		},
		{
			name: "Pid > 65535",
			msg: func(msg *MsgUpdateModelVersion) *MsgUpdateModelVersion {
				msg.Pid = 65536

				return msg
			}(validMsgUpdateModelVersion()),
			err: validator.ErrFieldUpperBoundViolated,
		},
		{
			name: "OtaUrl is not valid URL",
			msg: func(msg *MsgUpdateModelVersion) *MsgUpdateModelVersion {
				msg.OtaUrl = "not valid URL"

				return msg
			}(validMsgUpdateModelVersion()),
			err: validator.ErrFieldNotValid,
		},
		{
			name: "OtaUrl starts with http:",
			msg: func(msg *MsgUpdateModelVersion) *MsgUpdateModelVersion {
				msg.OtaUrl = "http://sampleflowurl.dclmodel"

				return msg
			}(validMsgUpdateModelVersion()),
			err: validator.ErrFieldNotValid,
		},
		{
			name: "OtaUrl length > 256",
			msg: func(msg *MsgUpdateModelVersion) *MsgUpdateModelVersion {
				msg.OtaUrl = "https://sampleflowurl.dclauth/" + tmrand.Str(257-30) // length = 257

				return msg
			}(validMsgUpdateModelVersion()),
			err: validator.ErrFieldMaxLengthExceeded,
		},
		{
			name: "OtaChecksum is not base64 encoded",
			msg: func(msg *MsgUpdateModelVersion) *MsgUpdateModelVersion {
				msg.OtaUrl = "https://sampleflowurl.dclmodel"
				msg.OtaChecksum = "not_base64_encoded"

				return msg
			}(validMsgUpdateModelVersion()),
			err: ErrOtaChecksumIsNotValid,
		},
		{
			name: "MinApplicableSoftwareVersion and MaxApplicableSoftwareVersion are set " +
				"and MinApplicableSoftwareVersion > MaxApplicableSoftwareVersion",
			msg: func(msg *MsgUpdateModelVersion) *MsgUpdateModelVersion {
				msg.MinApplicableSoftwareVersion = 8
				msg.MaxApplicableSoftwareVersion = 7

				return msg
			}(validMsgUpdateModelVersion()),
			err: validator.ErrFieldNotValid,
		},
		{
			name: "ReleaseNotesUrl is not valid URL",
			msg: func(msg *MsgUpdateModelVersion) *MsgUpdateModelVersion {
				msg.ReleaseNotesUrl = "not valid URL"

				return msg
			}(validMsgUpdateModelVersion()),
			err: validator.ErrFieldNotValid,
		},
		{
			name: "ReleaseNotesUrl starts with http:",
			msg: func(msg *MsgUpdateModelVersion) *MsgUpdateModelVersion {
				msg.ReleaseNotesUrl = "http://sampleflowurl.dclmodel"

				return msg
			}(validMsgUpdateModelVersion()),
			err: validator.ErrFieldNotValid,
		},
		{
			name: "ReleaseNotesUrl length > 256",
			msg: func(msg *MsgUpdateModelVersion) *MsgUpdateModelVersion {
				msg.ReleaseNotesUrl = "https://sampleflowurl.dclauth/" + tmrand.Str(257-30) // length = 257

				return msg
			}(validMsgUpdateModelVersion()),
			err: validator.ErrFieldMaxLengthExceeded,
		},
		{
			name: "schemaVersion > 65535",
			msg: func(msg *MsgUpdateModelVersion) *MsgUpdateModelVersion {
				msg.Creator = sample.AccAddress()
				msg.SchemaVersion = 65536

				return msg
			}(validMsgUpdateModelVersion()),
			err: validator.ErrFieldUpperBoundViolated,
		},
	}

	positiveTests := []struct {
		name string
		msg  *MsgUpdateModelVersion
	}{
		{
			name: "valid message",
			msg:  validMsgUpdateModelVersion(),
		},
		{
			name: "Creator is valid address",
			msg: func(msg *MsgUpdateModelVersion) *MsgUpdateModelVersion {
				msg.Creator = sample.AccAddress()

				return msg
			}(validMsgUpdateModelVersion()),
		},
		{
			name: "Vid == 1",
			msg: func(msg *MsgUpdateModelVersion) *MsgUpdateModelVersion {
				msg.Vid = 1

				return msg
			}(validMsgUpdateModelVersion()),
		},
		{
			name: "Vid == 65535",
			msg: func(msg *MsgUpdateModelVersion) *MsgUpdateModelVersion {
				msg.Vid = 65535

				return msg
			}(validMsgUpdateModelVersion()),
		},
		{
			name: "Pid == 1",
			msg: func(msg *MsgUpdateModelVersion) *MsgUpdateModelVersion {
				msg.Pid = 1

				return msg
			}(validMsgUpdateModelVersion()),
		},
		{
			name: "Pid == 65535",
			msg: func(msg *MsgUpdateModelVersion) *MsgUpdateModelVersion {
				msg.Pid = 65535

				return msg
			}(validMsgUpdateModelVersion()),
		},
		{
			name: "SoftwareVersion == 0",
			msg: func(msg *MsgUpdateModelVersion) *MsgUpdateModelVersion {
				msg.SoftwareVersion = 0

				return msg
			}(validMsgUpdateModelVersion()),
		},
		{
			name: "SoftwareVersion > 0",
			msg: func(msg *MsgUpdateModelVersion) *MsgUpdateModelVersion {
				msg.SoftwareVersion = 1

				return msg
			}(validMsgUpdateModelVersion()),
		},
		{
			name: "OtaUrl is omitted",
			msg: func(msg *MsgUpdateModelVersion) *MsgUpdateModelVersion {
				msg.OtaUrl = ""

				return msg
			}(validMsgUpdateModelVersion()),
		},
		{
			name: "OtaChecksum is base64 encoded",
			msg: func(msg *MsgUpdateModelVersion) *MsgUpdateModelVersion {
				msg.OtaUrl = "https://sampleflowurl.dclmodel"
				msg.OtaChecksum = "SGVsbG8gd29ybGQh"

				return msg
			}(validMsgUpdateModelVersion()),
		},
		{
			name: "OtaUrl length == 256",
			msg: func(msg *MsgUpdateModelVersion) *MsgUpdateModelVersion {
				msg.OtaUrl = "https://sampleflowurl.dclauth/" + tmrand.Str(256-30) // length = 256

				return msg
			}(validMsgUpdateModelVersion()),
		},
		{
			name: "MinApplicableSoftwareVersion and MaxApplicableSoftwareVersion are not set",
			msg: func(msg *MsgUpdateModelVersion) *MsgUpdateModelVersion {
				msg.MinApplicableSoftwareVersion = 0
				msg.MaxApplicableSoftwareVersion = 0

				return msg
			}(validMsgUpdateModelVersion()),
		},
		{
			name: "MinApplicableSoftwareVersion is set and MaxApplicableSoftwareVersion is not set",
			msg: func(msg *MsgUpdateModelVersion) *MsgUpdateModelVersion {
				msg.MinApplicableSoftwareVersion = 1
				msg.MaxApplicableSoftwareVersion = 0

				return msg
			}(validMsgUpdateModelVersion()),
		},
		{
			name: "MinApplicableSoftwareVersion is not set and MaxApplicableSoftwareVersion is set",
			msg: func(msg *MsgUpdateModelVersion) *MsgUpdateModelVersion {
				msg.MinApplicableSoftwareVersion = 0
				msg.MaxApplicableSoftwareVersion = 1

				return msg
			}(validMsgUpdateModelVersion()),
		},
		{
			name: "MinApplicableSoftwareVersion and MaxApplicableSoftwareVersion are set " +
				"and MinApplicableSoftwareVersion < MaxApplicableSoftwareVersion",
			msg: func(msg *MsgUpdateModelVersion) *MsgUpdateModelVersion {
				msg.MinApplicableSoftwareVersion = 5
				msg.MaxApplicableSoftwareVersion = 10

				return msg
			}(validMsgUpdateModelVersion()),
		},
		{
			name: "MinApplicableSoftwareVersion and MaxApplicableSoftwareVersion are set " +
				"and MinApplicableSoftwareVersion == MaxApplicableSoftwareVersion",
			msg: func(msg *MsgUpdateModelVersion) *MsgUpdateModelVersion {
				msg.MinApplicableSoftwareVersion = 7
				msg.MaxApplicableSoftwareVersion = 7

				return msg
			}(validMsgUpdateModelVersion()),
		},
		{
			name: "ReleaseNotesUrl is omitted",
			msg: func(msg *MsgUpdateModelVersion) *MsgUpdateModelVersion {
				msg.ReleaseNotesUrl = ""

				return msg
			}(validMsgUpdateModelVersion()),
		},
		{
			name: "ReleaseNotesUrl length == 256",
			msg: func(msg *MsgUpdateModelVersion) *MsgUpdateModelVersion {
				msg.ReleaseNotesUrl = "https://sampleflowurl.dclauth/" + tmrand.Str(256-30) // length = 256

				return msg
			}(validMsgUpdateModelVersion()),
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

func validMsgCreateModelVersion() *MsgCreateModelVersion {
	return &MsgCreateModelVersion{
		Creator:                      sample.AccAddress(),
		Vid:                          testconstants.Vid,
		Pid:                          testconstants.Pid,
		SoftwareVersion:              testconstants.SoftwareVersion,
		SoftwareVersionString:        testconstants.SoftwareVersionString,
		CdVersionNumber:              testconstants.CdVersionNumber,
		FirmwareInformation:          testconstants.FirmwareInformation,
		SoftwareVersionValid:         testconstants.SoftwareVersionValid,
		OtaUrl:                       testconstants.OtaURL,
		OtaFileSize:                  testconstants.OtaFileSize,
		OtaChecksum:                  testconstants.OtaChecksum,
		OtaChecksumType:              testconstants.OtaChecksumType,
		MinApplicableSoftwareVersion: testconstants.MinApplicableSoftwareVersion,
		MaxApplicableSoftwareVersion: testconstants.MaxApplicableSoftwareVersion,
		ReleaseNotesUrl:              testconstants.ReleaseNotesURL,
	}
}

func validMsgUpdateModelVersion() *MsgUpdateModelVersion {
	return &MsgUpdateModelVersion{
		Creator:                      sample.AccAddress(),
		Vid:                          testconstants.Vid,
		Pid:                          testconstants.Pid,
		SoftwareVersion:              testconstants.SoftwareVersion,
		SoftwareVersionValid:         !testconstants.SoftwareVersionValid,
		OtaUrl:                       testconstants.OtaURL + "/updated",
		MinApplicableSoftwareVersion: testconstants.MinApplicableSoftwareVersion + 1,
		MaxApplicableSoftwareVersion: testconstants.MaxApplicableSoftwareVersion + 1,
		ReleaseNotesUrl:              testconstants.ReleaseNotesURL + "/updated",
	}
}

func TestMsgDeleteModelVersion_ValidateBasic(t *testing.T) {
	negativeTests := []struct {
		name string
		msg  *MsgDeleteModelVersion
		err  error
	}{
		{
			name: "Creator is omitted",
			msg: func(msg *MsgDeleteModelVersion) *MsgDeleteModelVersion {
				msg.Creator = ""

				return msg
			}(validMsgDeleteModelVersion()),
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "Creator is not valid address",
			msg: func(msg *MsgDeleteModelVersion) *MsgDeleteModelVersion {
				msg.Creator = "not valid address"

				return msg
			}(validMsgDeleteModelVersion()),
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "Vid < 0",
			msg: func(msg *MsgDeleteModelVersion) *MsgDeleteModelVersion {
				msg.Vid = -1

				return msg
			}(validMsgDeleteModelVersion()),
			err: validator.ErrFieldLowerBoundViolated,
		},
		{
			name: "Vid == 0",
			msg: func(msg *MsgDeleteModelVersion) *MsgDeleteModelVersion {
				msg.Vid = 0

				return msg
			}(validMsgDeleteModelVersion()),
			err: validator.ErrFieldLowerBoundViolated,
		},
		{
			name: "Vid > 65535",
			msg: func(msg *MsgDeleteModelVersion) *MsgDeleteModelVersion {
				msg.Vid = 65536

				return msg
			}(validMsgDeleteModelVersion()),
			err: validator.ErrFieldUpperBoundViolated,
		},
		{
			name: "Pid < 0",
			msg: func(msg *MsgDeleteModelVersion) *MsgDeleteModelVersion {
				msg.Pid = -1

				return msg
			}(validMsgDeleteModelVersion()),
			err: validator.ErrFieldLowerBoundViolated,
		},
		{
			name: "Pid == 0",
			msg: func(msg *MsgDeleteModelVersion) *MsgDeleteModelVersion {
				msg.Pid = 0

				return msg
			}(validMsgDeleteModelVersion()),
			err: validator.ErrFieldLowerBoundViolated,
		},
		{
			name: "Pid > 65535",
			msg: func(msg *MsgDeleteModelVersion) *MsgDeleteModelVersion {
				msg.Pid = 65536

				return msg
			}(validMsgDeleteModelVersion()),
			err: validator.ErrFieldUpperBoundViolated,
		},
	}

	positiveTests := []struct {
		name string
		msg  *MsgDeleteModelVersion
	}{
		{
			name: "valid message",
			msg:  validMsgDeleteModelVersion(),
		},
		{
			name: "Creator is valid address",
			msg: func(msg *MsgDeleteModelVersion) *MsgDeleteModelVersion {
				msg.Creator = sample.AccAddress()

				return msg
			}(validMsgDeleteModelVersion()),
		},
		{
			name: "Vid == 1",
			msg: func(msg *MsgDeleteModelVersion) *MsgDeleteModelVersion {
				msg.Vid = 1

				return msg
			}(validMsgDeleteModelVersion()),
		},
		{
			name: "Vid == 65535",
			msg: func(msg *MsgDeleteModelVersion) *MsgDeleteModelVersion {
				msg.Vid = 65535

				return msg
			}(validMsgDeleteModelVersion()),
		},
		{
			name: "Pid == 1",
			msg: func(msg *MsgDeleteModelVersion) *MsgDeleteModelVersion {
				msg.Pid = 1

				return msg
			}(validMsgDeleteModelVersion()),
		},
		{
			name: "Pid == 65535",
			msg: func(msg *MsgDeleteModelVersion) *MsgDeleteModelVersion {
				msg.Pid = 65535

				return msg
			}(validMsgDeleteModelVersion()),
		},
		{
			name: "SoftwareVersion == 0",
			msg: func(msg *MsgDeleteModelVersion) *MsgDeleteModelVersion {
				msg.SoftwareVersion = 0

				return msg
			}(validMsgDeleteModelVersion()),
		},
		{
			name: "SoftwareVersion > 0",
			msg: func(msg *MsgDeleteModelVersion) *MsgDeleteModelVersion {
				msg.SoftwareVersion = 1

				return msg
			}(validMsgDeleteModelVersion()),
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

func validMsgDeleteModelVersion() *MsgDeleteModelVersion {
	return &MsgDeleteModelVersion{
		Creator:         sample.AccAddress(),
		Vid:             testconstants.Vid,
		Pid:             testconstants.Pid,
		SoftwareVersion: testconstants.SoftwareVersion,
	}
}
