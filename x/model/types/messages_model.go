package types

import (
	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/validator"
)

const (
	TypeMsgCreateModel = "create_model"
	TypeMsgUpdateModel = "update_model"
	TypeMsgDeleteModel = "delete_model"
)

var _ sdk.Msg = &MsgCreateModel{}

func NewMsgCreateModel(
	creator string,
	vid int32,
	pid int32,
	deviceTypeID int32,
	productName string,
	productLabel string,
	partNumber string,
	commissioningCustomFlow int32,
	commissioningCustomFlowURL string,
	commissioningModeInitialStepsHint uint32,
	commissioningModeInitialStepsInstruction string,
	commissioningModeSecondaryStepsHint uint32,
	commissioningModeSecondaryStepsInstruction string,
	commissionerRemoteUIFlowURL string,
	userManualURL string,
	supportURL string,
	productURL string,
	lsfURL string,
	schemaVersion uint32,
) *MsgCreateModel {
	return &MsgCreateModel{
		Creator:                                  creator,
		Vid:                                      vid,
		Pid:                                      pid,
		DeviceTypeId:                             deviceTypeID,
		ProductName:                              productName,
		ProductLabel:                             productLabel,
		PartNumber:                               partNumber,
		CommissioningCustomFlow:                  commissioningCustomFlow,
		CommissioningCustomFlowUrl:               commissioningCustomFlowURL,
		CommissioningModeInitialStepsHint:        commissioningModeInitialStepsHint,
		CommissioningModeInitialStepsInstruction: commissioningModeInitialStepsInstruction,
		CommissioningModeSecondaryStepsHint:      commissioningModeSecondaryStepsHint,
		CommissioningModeSecondaryStepsInstruction: commissioningModeSecondaryStepsInstruction,
		CommissionerRemoteUiFlowUrl:                commissionerRemoteUIFlowURL,
		UserManualUrl:                              userManualURL,
		SupportUrl:                                 supportURL,
		ProductUrl:                                 productURL,
		LsfUrl:                                     lsfURL,
		SchemaVersion:                              schemaVersion,
	}
}

func (msg *MsgCreateModel) Route() string {
	return RouterKey
}

func (msg *MsgCreateModel) Type() string {
	return TypeMsgCreateModel
}

func (msg *MsgCreateModel) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateModel) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)

	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateModel) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	err = validator.Validate(msg)
	if err != nil {
		return err
	}

	return nil
}

var _ sdk.Msg = &MsgUpdateModel{}

func NewMsgUpdateModel(
	creator string,
	vid int32,
	pid int32,
	productName string,
	productLabel string,
	partNumber string,
	commissioningCustomFlowURL string,
	commissioningModeInitialStepsInstruction string,
	commissioningModeSecondaryStepsInstruction string,
	commissionerRemoteUIFlowURL string,
	userManualURL string,
	supportURL string,
	productURL string,
	lsfURL string,
	lsfRevision int32,
	schemaVersion uint32,
	commissioningModeInitialStepsHint uint32,
) *MsgUpdateModel {
	return &MsgUpdateModel{
		Creator:                                  creator,
		Vid:                                      vid,
		Pid:                                      pid,
		ProductName:                              productName,
		ProductLabel:                             productLabel,
		PartNumber:                               partNumber,
		CommissioningCustomFlowUrl:               commissioningCustomFlowURL,
		CommissioningModeInitialStepsInstruction: commissioningModeInitialStepsInstruction,
		CommissioningModeSecondaryStepsInstruction: commissioningModeSecondaryStepsInstruction,
		CommissionerRemoteUiFlowUrl:                commissionerRemoteUIFlowURL,
		UserManualUrl:                              userManualURL,
		SupportUrl:                                 supportURL,
		ProductUrl:                                 productURL,
		LsfUrl:                                     lsfURL,
		LsfRevision:                                lsfRevision,
		SchemaVersion:                              schemaVersion,
		CommissioningModeInitialStepsHint:          commissioningModeInitialStepsHint,
	}
}

func (msg *MsgUpdateModel) Route() string {
	return RouterKey
}

func (msg *MsgUpdateModel) Type() string {
	return TypeMsgUpdateModel
}

func (msg *MsgUpdateModel) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateModel) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)

	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateModel) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	err = validator.Validate(msg)
	if err != nil {
		return err
	}

	return nil
}

var _ sdk.Msg = &MsgDeleteModel{}

func NewMsgDeleteModel(
	creator string,
	vid int32,
	pid int32,
) *MsgDeleteModel {
	return &MsgDeleteModel{
		Creator: creator,
		Vid:     vid,
		Pid:     pid,
	}
}

func (msg *MsgDeleteModel) Route() string {
	return RouterKey
}

func (msg *MsgDeleteModel) Type() string {
	return TypeMsgDeleteModel
}

func (msg *MsgDeleteModel) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteModel) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)

	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteModel) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	err = validator.Validate(msg)
	if err != nil {
		return err
	}

	return nil
}
