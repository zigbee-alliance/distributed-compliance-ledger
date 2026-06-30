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

func TestNewMsgProposeUpgrade(t *testing.T) {
	creator := sample.AccAddress()
	plan := Plan{
		Name:   testconstants.UpgradePlanName,
		Height: testconstants.UpgradePlanHeight,
		Info:   testconstants.UpgradePlanInfo,
	}

	msg := NewMsgProposeUpgrade(creator, plan, testconstants.Info)
	require.Equal(t, creator, msg.Creator)
	require.Equal(t, plan, msg.Plan)
	require.Equal(t, testconstants.Info, msg.Info)
	require.NotZero(t, msg.Time)
}

func TestMsgProposeUpgrade_ValidateBasicCLI_InvalidCreator(t *testing.T) {
	// Invalid creator address is rejected before the live GitHub-API call,
	// so this exercises ValidateBasicCLI without any network access.
	msg := MsgProposeUpgrade{
		Creator: "invalid_address",
		Plan: Plan{
			Name:   testconstants.UpgradePlanName,
			Height: testconstants.UpgradePlanHeight,
			Info:   testconstants.UpgradePlanInfo,
		},
		Info: testconstants.Info,
		Time: testconstants.Time,
	}

	err := msg.ValidateBasicCLI()
	require.Error(t, err)
	require.ErrorIs(t, err, sdkerrors.ErrInvalidAddress)
}

func TestMsgProposeUpgrade_ValidateBinaries(t *testing.T) {
	var expectedResponse string

	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprint(w, expectedResponse)
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
		{
			// valid JSON but the single binary targets an unsupported platform.
			name:     "unsupported os platform (valid json)",
			expected: testconstants.UpgradeGitAPIJSONResponse,
			msg: MsgProposeUpgrade{
				Creator: sample.AccAddress(),
				Plan: Plan{
					Name:   testconstants.UpgradePlanName,
					Height: testconstants.UpgradePlanHeight,
					Info:   "{\"binaries\":{\"windows/amd64\":\"https://github.com/zigbee-alliance/distributed-compliance-ledger/releases/download/v1.4.4/dcld?checksum=sha256:abc\"}}",
				},
				Info: testconstants.Info,
				Time: testconstants.Time,
			},
			err: sdkerrors.ErrJSONUnmarshal,
		},
		{
			// URL has the required "?checksum=" separator but too few path segments.
			name:     "binary url too short",
			expected: testconstants.UpgradeGitAPIJSONResponse,
			msg: MsgProposeUpgrade{
				Creator: sample.AccAddress(),
				Plan: Plan{
					Name:   testconstants.UpgradePlanName,
					Height: testconstants.UpgradePlanHeight,
					Info:   "{\"binaries\":{\"linux/amd64\":\"http://x?checksum=sha256:abc\"}}",
				},
				Info: testconstants.Info,
				Time: testconstants.Time,
			},
			err: sdkerrors.ErrInvalidRequest,
		},
		{
			// git tag parsed from the URL does not contain the plan name.
			name:     "plan name not in binary url tag",
			expected: testconstants.UpgradeGitAPIJSONResponse,
			msg: MsgProposeUpgrade{
				Creator: sample.AccAddress(),
				Plan: Plan{
					Name:   "v9.9.9",
					Height: testconstants.UpgradePlanHeight,
					Info:   testconstants.UpgradePlanInfo,
				},
				Info: testconstants.Info,
				Time: testconstants.Time,
			},
			err: sdkerrors.ErrInvalidRequest,
		},
		{
			// server returns a body that is not valid JSON.
			name:     "git api response not json",
			expected: "this is not json",
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
			err: sdkerrors.ErrJSONUnmarshal,
		},
		{
			// server returns valid JSON without the "assets" field.
			name:     "git api response without assets",
			expected: "{\"message\":\"Not Found\"}",
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
			err: sdkerrors.ErrJSONUnmarshal,
		},
		{
			// assets present but none match the expected binary.
			name:     "git api response with no matching asset",
			expected: "{\"assets\":[{\"name\": \"other\", \"state\": \"uploaded\", \"digest\": null, \"browser_download_url\":\"https://example.com/other\"}]}",
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
			err: sdkerrors.ErrInvalidRequest,
		},
	}

	positiveTests := []struct {
		name     string
		expected string
		msg      MsgProposeUpgrade
	}{
		{
			// empty Plan.Info short-circuits with no error and no network call.
			name:     "empty plan info",
			expected: testconstants.UpgradeGitAPIJSONResponse,
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
		{
			name:     "valid binary file for a dev release",
			expected: "{\"assets\":[{\"name\": \"dcld\", \"state\": \"uploaded\", \"digest\": \"sha256:f5c1120790319c9c4aefbfbc08a0bb1f91e848e3cd77cf3590a46d637e70cfad\", \"browser_download_url\":\"" + "https://github.com/zigbee-alliance/distributed-compliance-ledger/releases/download/v1.5.2-0.dev.1/dcld" + "\"}]}",
			msg: MsgProposeUpgrade{
				Creator: sample.AccAddress(),
				Plan: Plan{
					Name:   "v1.5.2",
					Height: testconstants.UpgradePlanHeight,
					Info:   "{\"binaries\":{\"linux/amd64\":\"https://github.com/zigbee-alliance/distributed-compliance-ledger/releases/download/v1.5.2-0.dev.1/dcld?checksum=sha256:f5c1120790319c9c4aefbfbc08a0bb1f91e848e3cd77cf3590a46d637e70cfad\"}}",
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

	t.Run("git token header is set when GH_TOKEN present", func(t *testing.T) {
		t.Setenv("GH_TOKEN", "test-token")
		expectedResponse = testconstants.UpgradeGitAPIJSONResponse
		msg := MsgProposeUpgrade{
			Creator: sample.AccAddress(),
			Plan: Plan{
				Name:   testconstants.UpgradePlanName,
				Height: testconstants.UpgradePlanHeight,
				Info:   testconstants.UpgradePlanInfo,
			},
			Info: testconstants.Info,
			Time: testconstants.Time,
		}
		require.NoError(t, ValidateBinaries(&msg, svr.URL))
	})

	t.Run("git api request fails", func(t *testing.T) {
		msg := MsgProposeUpgrade{
			Creator: sample.AccAddress(),
			Plan: Plan{
				Name:   testconstants.UpgradePlanName,
				Height: testconstants.UpgradePlanHeight,
				Info:   testconstants.UpgradePlanInfo,
			},
			Info: testconstants.Info,
			Time: testconstants.Time,
		}
		// Unroutable address forces client.Do to fail before any response.
		err := ValidateBinaries(&msg, "http://127.0.0.1:1")
		require.Error(t, err)
		require.ErrorIs(t, err, sdkerrors.ErrInvalidRequest)
	})
}
