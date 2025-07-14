package types

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"time"

	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/validator"
)

const TypeMsgProposeUpgrade = "propose_upgrade"
const GitReleaseApiUrl = "https://api.github.com/repos/zigbee-alliance/distributed-compliance-ledger/releases/tags"

// plan name <--> git tag
var ExistingUpgradesMap = map[string]string{
	"v0.10.0":     "v0.10.0",
	"v0.11.0":     "v0.11.0",
	"v0.12.0":     "v0.12.0",
	"v0.13.0-pre": "v0.13.0-pre",
	"v1.2":        "v1.2.2",
	"v1.4":        "v1.4.3",
	"v1.4.4":      "v1.4.4",
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

func ValidateBinaries(msg *MsgProposeUpgrade, gitBaseUrl string) error {

	if len(msg.Plan.Info) == 0 {
		return nil
	}

	var planInfoJson map[string]map[string]string

	err := json.Unmarshal([]byte(msg.Plan.Info), &planInfoJson)
	if err != nil {
		return sdkerrors.ErrJSONUnmarshal
	}

	binariesLen := len(planInfoJson["binaries"])

	if binariesLen > 1 {
		return errors.Wrapf(sdkerrors.ErrJSONUnmarshal, "invalid parsing, supports only one binary file")
	}

	if binariesLen == 0 {
		return errors.Wrapf(sdkerrors.ErrJSONUnmarshal, "invalid parsing, binary files not found")
	}

	for _, urlWithSum := range planInfoJson["binaries"] {
		fileUrl, sha256Sum, foundSep := strings.Cut(urlWithSum, "?")
		if !foundSep || !strings.HasPrefix(sha256Sum, "checksum=") {
			return errors.Wrapf(sdkerrors.ErrJSONUnmarshal, "invalid parsing upgrade plan url")
		}

		sha256Sum = strings.TrimPrefix(sha256Sum, "checksum=")

		partsUrl := strings.Split(fileUrl, "/")
		urlGitTag := partsUrl[7]

		// support previous updates where there is no direct matching of plan name and git tag
		existingGitTag, isUpgradeExist := ExistingUpgradesMap[msg.Plan.Name]
		if (!isUpgradeExist || urlGitTag != existingGitTag) && msg.Plan.Name != urlGitTag {
			return errors.Wrapf(sdkerrors.ErrInvalidRequest, "planName is not equal to the binary file version")
		}

		resp, err := http.Get(gitBaseUrl + "/" + urlGitTag)
		if err != nil {
			return errors.Wrapf(sdkerrors.ErrInvalidRequest, "binary file info request failed")
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return errors.Wrapf(sdkerrors.ErrInvalidRequest, "binary file info request failed")
		}

		var parsedBody map[string]any

		err = json.Unmarshal([]byte(body), &parsedBody)
		if err != nil {
			return errors.Wrapf(sdkerrors.ErrJSONUnmarshal, "invalid parsing binary file info")
		}

		var valid bool = false

		for _, asset := range parsedBody["assets"].([]any) {

			assetMap := asset.(map[string]any)

			if assetMap["name"] == "dcld" &&
				assetMap["state"] == "uploaded" &&
				assetMap["browser_download_url"] == fileUrl &&
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

	err = ValidateBinaries(msg, GitReleaseApiUrl)
	if err != nil {
		return err
	}

	return nil
}
