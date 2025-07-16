package types

import (
	context "context"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/validator"
)

const TypeMsgProposeUpgrade = "propose_upgrade"
const GitReleaseAPIURL = "https://api.github.com/repos/zigbee-alliance/distributed-compliance-ledger/releases/tags"

var ExistingUpgradesMap = map[string]string{
	// plan name <--> git tag
	"v0.10.0": "v0.10.0",
	"v0.11.0": "v0.11.0",
	"v0.12.0": "v0.12.0",
	"v1.2":    "v1.2.2",
	"v1.4":    "v1.4.3",
	"v1.4.4":  "v1.4.4",
}

var _ sdk.Msg = &MsgProposeUpgrade{}

func NewMsgProposeUpgrade(creator string, plan Plan, info string) *MsgProposeUpgrade {
	return &MsgProposeUpgrade{
		Creator: creator,
		Plan:    plan,
		Info:    info,
		Time:    time.Now().Unix(),
	}
}

func (msg *MsgProposeUpgrade) Route() string {
	return RouterKey
}

func (msg *MsgProposeUpgrade) Type() string {
	return TypeMsgProposeUpgrade
}

func (msg *MsgProposeUpgrade) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{creator}
}

func (msg *MsgProposeUpgrade) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)

	return sdk.MustSortJSON(bz)
}

func ValidateBinaries(msg *MsgProposeUpgrade, gitBaseURL string) error {
	if len(msg.Plan.Info) == 0 {
		return nil
	}

	var planInfoJSON map[string]map[string]string

	err := json.Unmarshal([]byte(msg.Plan.Info), &planInfoJSON)
	if err != nil {
		return sdkerrors.ErrJSONUnmarshal
	}

	binariesLen := len(planInfoJSON["binaries"])

	if binariesLen > 1 {
		return errors.Wrapf(sdkerrors.ErrJSONUnmarshal, "invalid parsing, supports only one binary file")
	}

	if binariesLen == 0 {
		return errors.Wrapf(sdkerrors.ErrJSONUnmarshal, "invalid parsing, binary files not found")
	}

	for _, urlWithSum := range planInfoJSON["binaries"] {
		fileURL, sha256Sum, foundSep := strings.Cut(urlWithSum, "?")
		if !foundSep || !strings.HasPrefix(sha256Sum, "checksum=") {
			return errors.Wrapf(sdkerrors.ErrJSONUnmarshal, "invalid parsing upgrade plan url")
		}

		sha256Sum = strings.TrimPrefix(sha256Sum, "checksum=")

		partsURL := strings.Split(fileURL, "/")
		urlGitTag := partsURL[7]

		// support previous updates where there is no direct matching of plan name and git tag
		existingGitTag, upgradeExist := ExistingUpgradesMap[msg.Plan.Name]
		if (!upgradeExist || urlGitTag != existingGitTag) && msg.Plan.Name != urlGitTag {
			return errors.Wrapf(sdkerrors.ErrInvalidRequest, "planName is not equal to the binary file version")
		}

		ctx := context.Background()
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, gitBaseURL+"/"+urlGitTag, http.NoBody)
		if err != nil {
			return errors.Wrapf(sdkerrors.ErrInvalidRequest, "binary file info create request failed")
		}

		gitToken := os.Getenv("GH_TOKEN")
		if len(gitToken) > 0 {
			req.Header.Add("Authorization", "token "+gitToken)
		}

		client := &http.Client{}
		resp, err := client.Do(req)

		if err != nil {
			return errors.Wrapf(sdkerrors.ErrInvalidRequest, "binary file info do request failed")
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return errors.Wrapf(sdkerrors.ErrInvalidRequest, "binary file info request failed")
		}

		var parsedBody map[string]any

		err = json.Unmarshal(body, &parsedBody)
		if err != nil {
			return errors.Wrapf(sdkerrors.ErrJSONUnmarshal, "invalid parsing binary file info")
		}

		assets, assetsExist := parsedBody["assets"]

		if !assetsExist {
			return errors.Wrapf(sdkerrors.ErrJSONUnmarshal, "invalid assets in json response")
		}

		valid := false

		for _, asset := range assets.([]any) {
			assetMap := asset.(map[string]any)

			if assetMap["name"] == "dcld" &&
				assetMap["state"] == "uploaded" &&
				assetMap["browser_download_url"] == fileURL &&
				(assetMap["digest"] == nil || assetMap["digest"] == sha256Sum) {
				valid = true

				break
			}
		}

		if !valid {
			return errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid binary file")
		}
	}

	return nil
}

func (msg *MsgProposeUpgrade) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	err = validator.Validate(msg)
	if err != nil {
		return err
	}

	err = msg.Plan.ValidateBasic()
	if err != nil {
		return err
	}

	return nil
}

func (msg *MsgProposeUpgrade) ValidateBasicCLI() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	err = validator.Validate(msg)
	if err != nil {
		return err
	}

	err = msg.Plan.ValidateBasic()
	if err != nil {
		return err
	}

	err = ValidateBinaries(msg, GitReleaseAPIURL)
	if err != nil {
		return err
	}

	return nil
}
