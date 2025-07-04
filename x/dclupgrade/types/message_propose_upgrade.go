package types

import (
	"encoding/json"
	"fmt"
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
const TmpFileForValidateBinaries = "/tmp/testfile"
const GitReleaseApiUrl = "https://api.github.com/repos/zigbee-alliance/distributed-compliance-ledger/releases/tags/"

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

func ValidateBinaries(planInfo *string) error {
	fmt.Fprintln(os.Stderr, "Start validate binaries")
	var planInfoJson map[string]map[string]string

	err := json.Unmarshal([]byte(*planInfo), &planInfoJson)
	if err != nil {
		return sdkerrors.ErrJSONUnmarshal
	}

	for _, urlWithSum := range planInfoJson["binaries"] {
		fileUrl, sha256Sum, foundSep := strings.Cut(urlWithSum, "?")
		if !foundSep {
			return errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid parsing upgrade plan url")
		}

		partsUrl := strings.Split(fileUrl, "/")

		gitTag := partsUrl[7]

		resp, err := http.Get(GitReleaseApiUrl + gitTag)
		if err != nil {
			return errors.Wrapf(sdkerrors.ErrInvalidRequest, "request error")
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return errors.Wrapf(sdkerrors.ErrInvalidRequest, "read request error")
		}

		var parsedBody map[string]any

		err = json.Unmarshal([]byte(body), &parsedBody)

		if err != nil {
			return sdkerrors.ErrJSONUnmarshal
		}

		var valid bool = false

		for _, asset := range parsedBody["assets"].([]any) {

			assetMap := asset.(map[string]any)

			if assetMap["name"] == "dcld" && assetMap["browser_download_url"] == fileUrl &&
				(assetMap["digest"] == nil || assetMap["digest"] == sha256Sum) {

				valid = true
				fmt.Printf("Valid link\n")
				// for key, value := range assetMap {
				// 	fmt.Printf("key: %s, value: %s\n", key, value)
				// }
			}

		}

		if !valid {
			return errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid binary file")
		}
	}

	println("Binary files validation completed successfully")

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

	if len(msg.Plan.Info) > 0 {
		err = ValidateBinaries(&msg.Plan.Info)
		if err != nil {
			return err
		}
	}

	return nil
}
