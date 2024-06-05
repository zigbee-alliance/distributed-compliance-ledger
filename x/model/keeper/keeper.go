package keeper

import (
	"fmt"

	"github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	compliancetypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/compliance/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/model/types"
)

type (
	Keeper struct {
		cdc      codec.BinaryCodec
		storeKey storetypes.StoreKey
		memKey   storetypes.StoreKey

		dclauthKeeper    types.DclauthKeeper
		complianceKeeper ComplianceKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,

	dclauthKeeper types.DclauthKeeper,
	complianceKeeper ComplianceKeeper,
) *Keeper {
	return &Keeper{
		cdc:      cdc,
		storeKey: storeKey,
		memKey:   memKey,

		dclauthKeeper:    dclauthKeeper,
		complianceKeeper: complianceKeeper,
	}
}

func (k *Keeper) SetComplianceKeeper(complianceKeeper ComplianceKeeper) {
	k.complianceKeeper = complianceKeeper
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

type ComplianceKeeper interface {
	// Methods imported from compliance should be defined here
	GetComplianceInfo(
		ctx sdk.Context,
		vid int32,
		pid int32,
		softwareVersion uint32,
		certificationType string,
	) (val compliancetypes.ComplianceInfo, found bool)
}
