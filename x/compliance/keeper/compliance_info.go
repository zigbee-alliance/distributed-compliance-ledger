package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	dclcompltypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/compliance"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliance/types"
)

// SetComplianceInfo set a specific complianceInfo in the store from its index.
func (k Keeper) SetComplianceInfo(ctx sdk.Context, complianceInfo dclcompltypes.ComplianceInfo) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ComplianceInfoKeyPrefix))
	b := k.cdc.MustMarshal(&complianceInfo)
	store.Set(types.ComplianceInfoKey(
		complianceInfo.Vid,
		complianceInfo.Pid,
		complianceInfo.SoftwareVersion,
		complianceInfo.CertificationType,
	), b)
}

// GetComplianceInfo returns a complianceInfo from its index.
func (k Keeper) GetComplianceInfo(
	ctx sdk.Context,
	vid int32,
	pid int32,
	softwareVersion uint32,
	certificationType string,
) (val dclcompltypes.ComplianceInfo, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ComplianceInfoKeyPrefix))

	b := store.Get(types.ComplianceInfoKey(
		vid,
		pid,
		softwareVersion,
		certificationType,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)

	return val, true
}

// RemoveComplianceInfo removes a complianceInfo from the store.
func (k Keeper) RemoveComplianceInfo(
	ctx sdk.Context,
	vid int32,
	pid int32,
	softwareVersion uint32,
	certificationType string,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ComplianceInfoKeyPrefix))
	store.Delete(types.ComplianceInfoKey(
		vid,
		pid,
		softwareVersion,
		certificationType,
	))
}

// GetAllComplianceInfo returns all complianceInfo.
func (k Keeper) GetAllComplianceInfo(ctx sdk.Context) (list []dclcompltypes.ComplianceInfo) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ComplianceInfoKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val dclcompltypes.ComplianceInfo
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
