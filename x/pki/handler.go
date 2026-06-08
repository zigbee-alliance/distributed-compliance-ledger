package pki

import (
	"fmt"

	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

// NewHandler ...
func NewHandler(k keeper.Keeper) sdk.Handler {
	msgServer := keeper.NewMsgServerImpl(k)

	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		case *types.MsgProposeAddX509RootCert:
			res, err := msgServer.ProposeAddX509RootCert(sdk.WrapSDKContext(ctx), msg)

			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgApproveAddX509RootCert:
			res, err := msgServer.ApproveAddX509RootCert(sdk.WrapSDKContext(ctx), msg)

			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgAddX509Cert:
			res, err := msgServer.AddX509Cert(sdk.WrapSDKContext(ctx), msg)

			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgProposeRevokeX509RootCert:
			res, err := msgServer.ProposeRevokeX509RootCert(sdk.WrapSDKContext(ctx), msg)

			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgApproveRevokeX509RootCert:
			res, err := msgServer.ApproveRevokeX509RootCert(sdk.WrapSDKContext(ctx), msg)

			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgRevokeX509Cert:
			res, err := msgServer.RevokeX509Cert(sdk.WrapSDKContext(ctx), msg)

			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgRejectAddX509RootCert:
			res, err := msgServer.RejectAddX509RootCert(sdk.WrapSDKContext(ctx), msg)

			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgAddPkiRevocationDistributionPoint:
			res, err := msgServer.AddPkiRevocationDistributionPoint(sdk.WrapSDKContext(ctx), msg)

			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgUpdatePkiRevocationDistributionPoint:
			res, err := msgServer.UpdatePkiRevocationDistributionPoint(sdk.WrapSDKContext(ctx), msg)

			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgDeletePkiRevocationDistributionPoint:
			res, err := msgServer.DeletePkiRevocationDistributionPoint(sdk.WrapSDKContext(ctx), msg)

			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgAssignVid:
			res, err := msgServer.AssignVid(sdk.WrapSDKContext(ctx), msg)

			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgAddNocX509RootCert:
			res, err := msgServer.AddNocX509RootCert(sdk.WrapSDKContext(ctx), msg)

			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgRemoveX509Cert:
			res, err := msgServer.RemoveX509Cert(sdk.WrapSDKContext(ctx), msg)

			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgAddNocX509IcaCert:
			res, err := msgServer.AddNocX509IcaCert(sdk.WrapSDKContext(ctx), msg)

			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgRevokeNocX509RootCert:
			res, err := msgServer.RevokeNocX509RootCert(sdk.WrapSDKContext(ctx), msg)

			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgRevokeNocX509IcaCert:
			res, err := msgServer.RevokeNocX509IcaCert(sdk.WrapSDKContext(ctx), msg)

			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgRemoveNocX509IcaCert:
			res, err := msgServer.RemoveNocX509IcaCert(sdk.WrapSDKContext(ctx), msg)

			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgRemoveNocX509RootCert:
			res, err := msgServer.RemoveNocX509RootCert(sdk.WrapSDKContext(ctx), msg)

			return sdk.WrapServiceResult(ctx, res, err)
			// this line is used by starport scaffolding # 1
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", pkitypes.ModuleName, msg)

			return nil, errors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}
