package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliance/types"
)

// SetDeviceSoftwareCompliance set a specific deviceSoftwareCompliance in the store from its index
func (k Keeper) SetDeviceSoftwareCompliance(ctx sdk.Context, deviceSoftwareCompliance types.DeviceSoftwareCompliance) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DeviceSoftwareComplianceKeyPrefix))
	b := k.cdc.MustMarshal(&deviceSoftwareCompliance)
	store.Set(types.DeviceSoftwareComplianceKey(
		deviceSoftwareCompliance.CDCertificateId,
	), b)
}

// GetDeviceSoftwareCompliance returns a deviceSoftwareCompliance from its index
func (k Keeper) GetDeviceSoftwareCompliance(
	ctx sdk.Context,
	cdCertificateId string,

) (val types.DeviceSoftwareCompliance, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DeviceSoftwareComplianceKeyPrefix))

	b := store.Get(types.DeviceSoftwareComplianceKey(
		cdCertificateId,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveDeviceSoftwareCompliance removes a deviceSoftwareCompliance from the store
func (k Keeper) RemoveDeviceSoftwareCompliance(
	ctx sdk.Context,
	cdCertificateId string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DeviceSoftwareComplianceKeyPrefix))
	store.Delete(types.DeviceSoftwareComplianceKey(
		cdCertificateId,
	))
}

// GetAllDeviceSoftwareCompliance returns all deviceSoftwareCompliance
func (k Keeper) GetAllDeviceSoftwareCompliance(ctx sdk.Context) (list []types.DeviceSoftwareCompliance) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DeviceSoftwareComplianceKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.DeviceSoftwareCompliance
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
