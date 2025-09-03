package types

import (
	"testing"

	"fmt"
	"net/http"
	"net/http/httptest"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	"github.com/zigbee-alliance/distributed-compliance-ledger/testutil/sample"
)

func TestMsgProposeUpgrade_ValidateBasic(t *testing.T) {
	negativeTests := []struct {
		name string
		msg  MsgProposeUpgrade
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgProposeUpgrade{
				Creator: "invalid_address",
				Plan: Plan{
					Name:   testconstants.UpgradePlanName,
					Height: testconstants.UpgradePlanHeight,
					Info:   testconstants.UpgradePlanInfo,
				},
				Info: testconstants.Info,
				Time: testconstants.Time,
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "omitted address",
			msg: MsgProposeUpgrade{
				Creator: "",
				Plan: Plan{
					Name:   testconstants.UpgradePlanName,
					Height: testconstants.UpgradePlanHeight,
					Info:   testconstants.UpgradePlanInfo,
				},
				Info: testconstants.Info,
				Time: testconstants.Time,
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "plan name is not set",
			msg: MsgProposeUpgrade{
				Creator: sample.AccAddress(),
				Plan: Plan{
					Name:   "",
					Height: testconstants.UpgradePlanHeight,
					Info:   testconstants.UpgradePlanInfo,
				},
				Info: testconstants.Info,
				Time: testconstants.Time,
			},
			err: sdkerrors.ErrInvalidRequest,
		},
		{
			name: "plan height is 0",
			msg: MsgProposeUpgrade{
				Creator: sample.AccAddress(),
				Plan: Plan{
					Name:   testconstants.UpgradePlanName,
					Height: 0,
					Info:   testconstants.UpgradePlanInfo,
				},
				Info: testconstants.Info,
				Time: testconstants.Time,
			},
			err: sdkerrors.ErrInvalidRequest,
		},
		{
			name: "plan height is less than 0",
			msg: MsgProposeUpgrade{
				Creator: sample.AccAddress(),
				Plan: Plan{
					Name:   testconstants.UpgradePlanName,
					Height: -1,
					Info:   testconstants.UpgradePlanInfo,
				},
				Info: testconstants.Info,
				Time: testconstants.Time,
			},
			err: sdkerrors.ErrInvalidRequest,
		},
	}

	positiveTests := []struct {
		name string
		msg  MsgProposeUpgrade
	}{
		{
			name: "valid MsgProposeUpgrade message",
			msg: MsgProposeUpgrade{
				Creator: sample.AccAddress(),
				Plan: Plan{
					Name:   testconstants.UpgradePlanName,
					Height: testconstants.UpgradePlanHeight,
					Info:   testconstants.UpgradePlanInfo,
				},
				Info: testconstants.Info,
				Time: testconstants.Time,
			},
		},
		{
			name: "info is not set",
			msg: MsgProposeUpgrade{
				Creator: sample.AccAddress(),
				Plan: Plan{
					Name:   testconstants.UpgradePlanName,
					Height: testconstants.UpgradePlanHeight,
					Info:   testconstants.UpgradePlanInfo,
				},
				Info: "",
				Time: testconstants.Time,
			},
		},
		{
			name: "plan height is greater than 0",
			msg: MsgProposeUpgrade{
				Creator: sample.AccAddress(),
				Plan: Plan{
					Name:   testconstants.UpgradePlanName,
					Height: 1,
					Info:   testconstants.UpgradePlanInfo,
				},
				Info: testconstants.Info,
				Time: testconstants.Time,
			},
		},
		{
			name: "plan info is not set",
			msg: MsgProposeUpgrade{
				Creator: sample.AccAddress(),
				Plan: Plan{
					Name:   testconstants.UpgradePlanName,
					Height: testconstants.UpgradePlanHeight,
					Info:   "",
				},
				Info: testconstants.Info,
				Time: testconstants.Time,
			},
		},
	}
	for _, tt := range positiveTests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			require.NoError(t, err)
		})
	}

	for _, tt := range negativeTests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			require.Error(t, err)
			require.ErrorIs(t, err, tt.err)
		})
	}
}

func TestMsgProposeUpgrade_ValidateBinaries(t *testing.T) {
	var expectedResponse string

	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, expectedResponse)
	}))
	defer svr.Close()

	negativeTests := []struct {
		name     string
		expected string
		msg      MsgProposeUpgrade
		err      error
	}{
		{
			name:     "invalid binary format 1",
			expected: testconstants.UpgradeGitAPIJSONResponse,
			msg: MsgProposeUpgrade{
				Creator: sample.AccAddress(),
				Plan: Plan{
					Name:   testconstants.UpgradePlanName,
					Height: testconstants.UpgradePlanHeight,
					Info:   "{\"binaries\":{\"linux/amd64\":\"https://github.com/zigbee-alliance/distributed-compliance-ledger/releases/download/v1.4.4/dcld-checksum=sha256:e4031c6a77aa8e58add391be671a334613271bcf6e7f11d23b04a0881ece6958\"}}",
				},
				Info: testconstants.Info,
				Time: testconstants.Time,
			},
			err: sdkerrors.ErrJSONUnmarshal,
		},
		{
			name:     "invalid binary format 2",
			expected: testconstants.UpgradeGitAPIJSONResponse,
			msg: MsgProposeUpgrade{
				Creator: sample.AccAddress(),
				Plan: Plan{
					Name:   testconstants.UpgradePlanName,
					Height: testconstants.UpgradePlanHeight,
					Info:   "{\"binaries\":{\"linux/amd64\":\"https://github.com/zigbee-alliance/distributed-compliance-ledger/releases/download/v1.4.4/dcld?checksum-sha256:e4031c6a77aa8e58add391be671a334613271bcf6e7f11d23b04a0881ece6958\"}}",
				},
				Info: testconstants.Info,
				Time: testconstants.Time,
			},
			err: sdkerrors.ErrJSONUnmarshal,
		},
		{
			name:     "invalid binary format 3",
			expected: testconstants.UpgradeGitAPIJSONResponse,
			msg: MsgProposeUpgrade{
				Creator: sample.AccAddress(),
				Plan: Plan{
					Name:   testconstants.UpgradePlanName,
					Height: testconstants.UpgradePlanHeight,
					Info:   "{\"binaries\":[\"linux/amd64\":\"https://github.com/zigbee-alliance/distributed-compliance-ledger/releases/download/v1.4.4/dcld?checksum-sha256:e4031c6a77aa8e58add391be671a334613271bcf6e7f11d23b04a0881ece6958\"]}",
				},
				Info: testconstants.Info,
				Time: testconstants.Time,
			},
			err: sdkerrors.ErrJSONUnmarshal,
		},
		{
			name:     "lots of binary files",
			expected: testconstants.UpgradeGitAPIJSONResponse,
			msg: MsgProposeUpgrade{
				Creator: sample.AccAddress(),
				Plan: Plan{
					Name:   testconstants.UpgradePlanName,
					Height: testconstants.UpgradePlanHeight,
					Info:   "{\"binaries\":{\"linux/amd64\":\"URL1\", \"mac\":\"URL2\"}}",
				},
				Info: testconstants.Info,
				Time: testconstants.Time,
			},
			err: sdkerrors.ErrJSONUnmarshal,
		},
		{
			name:     "unsupported os platform",
			expected: testconstants.UpgradeGitAPIJSONResponse,
			msg: MsgProposeUpgrade{
				Creator: sample.AccAddress(),
				Plan: Plan{
					Name:   testconstants.UpgradePlanName,
					Height: testconstants.UpgradePlanHeight,
					Info:   "{\"binaries\":[\"mac\":\"https://github.com/zigbee-alliance/distributed-compliance-ledger/releases/download/v1.4.4/dcld?checksum-sha256:e4031c6a77aa8e58add391be671a334613271bcf6e7f11d23b04a0881ece6958\"]}",
				},
				Info: testconstants.Info,
				Time: testconstants.Time,
			},
			err: sdkerrors.ErrJSONUnmarshal,
		},
		{
			name:     "no binary files",
			expected: testconstants.UpgradeGitAPIJSONResponse,
			msg: MsgProposeUpgrade{
				Creator: sample.AccAddress(),
				Plan: Plan{
					Name:   testconstants.UpgradePlanName,
					Height: testconstants.UpgradePlanHeight,
					Info:   "{\"binaries\":{}}",
				},
				Info: testconstants.Info,
				Time: testconstants.Time,
			},
			err: sdkerrors.ErrJSONUnmarshal,
		},
	}

	positiveTests := []struct {
		name     string
		expected string
		msg      MsgProposeUpgrade
	}{
		{
			name:     "valid binary file without checksum",
			expected: "{\"assets\":[{\"name\": \"dcld\", \"state\": \"uploaded\", \"digest\": null, \"browser_download_url\":\"" + testconstants.UpgradeBrowserDownloadURL + "\"}]}",
			msg: MsgProposeUpgrade{
				Creator: sample.AccAddress(),
				Plan: Plan{
					Name:   testconstants.UpgradePlanName,
					Height: testconstants.UpgradePlanHeight,
					Info:   testconstants.UpgradePlanInfo,
				},
				Info: testconstants.Info,
				Time: testconstants.Time,
			},
		},
		{
			name:     "valid binary file with checksum",
			expected: testconstants.UpgradeGitAPIJSONResponse,
			msg: MsgProposeUpgrade{
				Creator: sample.AccAddress(),
				Plan: Plan{
					Name:   testconstants.UpgradePlanName,
					Height: testconstants.UpgradePlanHeight,
					Info:   testconstants.UpgradePlanInfo,
				},
				Info: testconstants.Info,
				Time: testconstants.Time,
			},
		},
	}
	for _, tt := range positiveTests {
		t.Run(tt.name, func(t *testing.T) {
			expectedResponse = tt.expected
			err := ValidateBinaries(&tt.msg, svr.URL)
			require.NoError(t, err)
		})
	}

	for _, tt := range negativeTests {
		t.Run(tt.name, func(t *testing.T) {
			expectedResponse = tt.expected
			err := ValidateBinaries(&tt.msg, svr.URL)
			require.Error(t, err)
			require.ErrorIs(t, err, tt.err)
		})
	}
}
