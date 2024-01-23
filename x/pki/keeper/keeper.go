package keeper

import (
	"fmt"
	"math"

	"github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	authTypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

type (
	Keeper struct {
		cdc      codec.BinaryCodec
		storeKey storetypes.StoreKey
		memKey   storetypes.StoreKey

		dclauthKeeper types.DclauthKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,

	dclauthKeeper types.DclauthKeeper,
) *Keeper {
	return &Keeper{
		cdc:      cdc,
		storeKey: storeKey,
		memKey:   memKey,

		dclauthKeeper: dclauthKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", pkitypes.ModuleName))
}

func (k Keeper) CertificateApprovalsCount(ctx sdk.Context, authKeeper types.DclauthKeeper) int {
	return int(math.Ceil(types.RootCertificateApprovalsPercent *
		float64(authKeeper.CountAccountsWithRole(ctx, authTypes.Trustee))))
}

func (k Keeper) CertificateRejectApprovalsCount(ctx sdk.Context, authKeeper types.DclauthKeeper) int {
	return authKeeper.CountAccountsWithRole(ctx, authTypes.Trustee) - k.CertificateApprovalsCount(ctx, authKeeper) + 1
}
