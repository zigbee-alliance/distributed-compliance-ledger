package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/validator"
)

const (
	TypeMsgCreateModelVersion = "create_model_version"
	TypeMsgUpdateModelVersion = "update_model_version"
	TypeMsgDeleteModelVersion = "delete_model_version"
)

var _ sdk.Msg = &MsgCreateModelVersion{}

func NewMsgCreateModelVersion(
	creator string,
	vid int32,
	pid int32,
	softwareVersion uint32,
	softwareVersionString string,
	cdVersionNumber int32,
	firmwareInformation string,
	softwareVersionValid bool,
	otaURL string,
	otaFileSize uint64,
	otaChecksum string,
	otaChecksumType int32,
	minApplicableSoftwareVersion uint32,
	maxApplicableSoftwareVersion uint32,
	releaseNotesURL string,
) *MsgCreateModelVersion {
	return &MsgCreateModelVersion{
		Creator:                      creator,
		Vid:                          vid,
		Pid:                          pid,
		SoftwareVersion:              softwareVersion,
		SoftwareVersionString:        softwareVersionString,
		CdVersionNumber:              cdVersionNumber,
		FirmwareInformation:          firmwareInformation,
		SoftwareVersionValid:         softwareVersionValid,
		OtaUrl:                       otaURL,
		OtaFileSize:                  otaFileSize,
		OtaChecksum:                  otaChecksum,
		OtaChecksumType:              otaChecksumType,
		MinApplicableSoftwareVersion: minApplicableSoftwareVersion,
		MaxApplicableSoftwareVersion: maxApplicableSoftwareVersion,
		ReleaseNotesUrl:              releaseNotesURL,
	}
}

func (msg *MsgCreateModelVersion) Route() string {
	return RouterKey
}

func (msg *MsgCreateModelVersion) Type() string {
	return TypeMsgCreateModelVersion
}

func (msg *MsgCreateModelVersion) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateModelVersion) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)

	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateModelVersion) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	err = validator.Validate(msg)
	if err != nil {
		return err
	}

	return nil
}

var _ sdk.Msg = &MsgUpdateModelVersion{}

func NewMsgUpdateModelVersion(
	creator string,
	vid int32,
	pid int32,
	softwareVersion uint32,
	softwareVersionValid bool,
	otaURL string,
	minApplicableSoftwareVersion uint32,
	maxApplicableSoftwareVersion uint32,
	releaseNotesURL string,
) *MsgUpdateModelVersion {
	return &MsgUpdateModelVersion{
		Creator:                      creator,
		Vid:                          vid,
		Pid:                          pid,
		SoftwareVersion:              softwareVersion,
		SoftwareVersionValid:         softwareVersionValid,
		OtaUrl:                       otaURL,
		MinApplicableSoftwareVersion: minApplicableSoftwareVersion,
		MaxApplicableSoftwareVersion: maxApplicableSoftwareVersion,
		ReleaseNotesUrl:              releaseNotesURL,
	}
}

func (msg *MsgUpdateModelVersion) Route() string {
	return RouterKey
}

func (msg *MsgUpdateModelVersion) Type() string {
	return TypeMsgUpdateModelVersion
}

func (msg *MsgUpdateModelVersion) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateModelVersion) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)

	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateModelVersion) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	err = validator.Validate(msg)
	if err != nil {
		return err
	}

	return nil
}

var _ sdk.Msg = &MsgDeleteModelVersion{}

func NewMsgDeleteModelVersion(creator string, vid int32, pid int32, softwareVersion uint32) *MsgDeleteModelVersion {
	return &MsgDeleteModelVersion{
		Creator:         creator,
		Vid:             vid,
		Pid:             pid,
		SoftwareVersion: softwareVersion,
	}
}

func (msg *MsgDeleteModelVersion) Route() string {
	return RouterKey
}

func (msg *MsgDeleteModelVersion) Type() string {
	return TypeMsgDeleteModelVersion
}

func (msg *MsgDeleteModelVersion) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{signer}
}

func (msg *MsgDeleteModelVersion) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)

	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteModelVersion) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signer address (%s)", err)
	}

	return nil
}
