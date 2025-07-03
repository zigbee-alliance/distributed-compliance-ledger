package types

import (
	"crypto/sha256"
	"encoding/hex"
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
		fileUrl, sh256Sum, found1 := strings.Cut(urlWithSum, "?")
		if !found1 {
			return errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid parsing raw url")
		}

		_, expectedSum, found2 := strings.Cut(sh256Sum, ":")
		if !found2 {
			return errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid parsing raw url")
		}

		// println("Trying download file from ", fileUrl)
		fmt.Fprintln(os.Stderr, "Trying download file from ", fileUrl)

		// Create the file
		out, err := os.Create(TmpFileForValidateBinaries)
		if err != nil {
			return errors.Wrapf(sdkerrors.ErrInvalidRequest, "error creatinging temp binary file")
		}
		defer out.Close()
		defer os.Remove(TmpFileForValidateBinaries)

		resp, err := http.Get(fileUrl)
		if err != nil {
			return errors.Wrapf(sdkerrors.ErrInvalidRequest, "error downloading a binary file")
		}
		defer resp.Body.Close()

		hash256 := sha256.New()
		_, err = io.Copy(hash256, resp.Body)
		if err != nil {
			return err
		}

		realSum := hex.EncodeToString(hash256.Sum(nil))
		if expectedSum != realSum {
			return errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid file checksum")
		}
	}

	println("Download and checksum verification completed successfully")

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
