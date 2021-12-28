package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliancetest/types"
	dclauthtypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
	modeltypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/model/types"
)

func (k msgServer) AddTestingResult(goCtx context.Context, msg *types.MsgAddTestingResult) (*types.MsgAddTestingResultResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	signerAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid Address: (%s)", err)
	}

	// sender must have TestHouse role to add new model
	if !k.dclauthKeeper.HasRole(ctx, signerAddr, dclauthtypes.TestHouse) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized,
			"MsgAddTestingResult transaction should be signed by an account with the %s role",
			dclauthtypes.TestHouse,
		)
	}

	// check that corresponding model exists on the ledger
	modelVersion, found := k.modelKeeper.GetModelVersion(ctx, msg.Vid, msg.Pid, msg.SoftwareVersion)
	if !found {
		return nil, modeltypes.NewErrModelVersionDoesNotExist(msg.Vid, msg.Pid, msg.SoftwareVersion)
	}

	// check if softwareVersionString matches with what is stored for the given version
	if modelVersion.SoftwareVersionString != msg.SoftwareVersionString {
		return nil, types.NewErrModelVersionStringDoesNotMatch(msg.Vid, msg.Pid, msg.SoftwareVersion, msg.SoftwareVersionString)
	}

	testingResult := types.TestingResult{
		Vid:                   msg.Vid,
		Pid:                   msg.Pid,
		SoftwareVersion:       msg.SoftwareVersion,
		SoftwareVersionString: msg.SoftwareVersionString,
		Owner:                 msg.Signer,
		TestResult:            msg.TestResult,
		TestDate:              msg.TestDate,
	}

	// store testing results. it extends existing value if testing results already exists
	k.AppendTestingResult(ctx, testingResult)

	return &types.MsgAddTestingResultResponse{}, nil
}
