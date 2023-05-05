package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
)

// SetPKIRevocationDistributionPoint set a specific pKIRevocationDistributionPoint in the store from its index
func (k Keeper) SetPKIRevocationDistributionPoint(ctx sdk.Context, pKIRevocationDistributionPoint types.PKIRevocationDistributionPoint) {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PKIRevocationDistributionPointKeyPrefix))
	b := k.cdc.MustMarshal(&pKIRevocationDistributionPoint)
	store.Set(types.PKIRevocationDistributionPointKey(
        pKIRevocationDistributionPoint.Vid,
    pKIRevocationDistributionPoint.Label,
    pKIRevocationDistributionPoint.IssuerSubjectKeyID,
    ), b)
}

// GetPKIRevocationDistributionPoint returns a pKIRevocationDistributionPoint from its index
func (k Keeper) GetPKIRevocationDistributionPoint(
    ctx sdk.Context,
    vid uint64,
    label string,
    issuerSubjectKeyID string,
    
) (val types.PKIRevocationDistributionPoint, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PKIRevocationDistributionPointKeyPrefix))

	b := store.Get(types.PKIRevocationDistributionPointKey(
        vid,
    label,
    issuerSubjectKeyID,
    ))
    if b == nil {
        return val, false
    }

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemovePKIRevocationDistributionPoint removes a pKIRevocationDistributionPoint from the store
func (k Keeper) RemovePKIRevocationDistributionPoint(
    ctx sdk.Context,
    vid uint64,
    label string,
    issuerSubjectKeyID string,
    
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PKIRevocationDistributionPointKeyPrefix))
	store.Delete(types.PKIRevocationDistributionPointKey(
	    vid,
    label,
    issuerSubjectKeyID,
    ))
}

// GetAllPKIRevocationDistributionPoint returns all pKIRevocationDistributionPoint
func (k Keeper) GetAllPKIRevocationDistributionPoint(ctx sdk.Context) (list []types.PKIRevocationDistributionPoint) {
    store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PKIRevocationDistributionPointKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.PKIRevocationDistributionPoint
		k.cdc.MustUnmarshal(iterator.Value(), &val)
        list = append(list, val)
	}

    return
}
