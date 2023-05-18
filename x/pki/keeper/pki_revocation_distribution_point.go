package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

// SetPkiRevocationDistributionPoint set a specific pKIRevocationDistributionPoint in the store from its index
func (k Keeper) SetPkiRevocationDistributionPoint(ctx sdk.Context, pKIRevocationDistributionPoint types.PkiRevocationDistributionPoint) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.PkiRevocationDistributionPointKeyPrefix))
	b := k.cdc.MustMarshal(&pKIRevocationDistributionPoint)
	store.Set(types.PkiRevocationDistributionPointKey(
		pKIRevocationDistributionPoint.Vid,
		pKIRevocationDistributionPoint.Label,
		pKIRevocationDistributionPoint.IssuerSubjectKeyID,
	), b)
}

// GetPkiRevocationDistributionPoint returns a pKIRevocationDistributionPoint from its index
func (k Keeper) GetPkiRevocationDistributionPoint(
	ctx sdk.Context,
	vid int32,
	label string,
	issuerSubjectKeyID string,

) (val types.PkiRevocationDistributionPoint, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.PkiRevocationDistributionPointKeyPrefix))

	b := store.Get(types.PkiRevocationDistributionPointKey(
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

// RemovePkiRevocationDistributionPoint removes a pKIRevocationDistributionPoint from the store
func (k Keeper) RemovePkiRevocationDistributionPoint(
	ctx sdk.Context,
	vid int32,
	label string,
	issuerSubjectKeyID string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.PkiRevocationDistributionPointKeyPrefix))
	store.Delete(types.PkiRevocationDistributionPointKey(
		vid,
		label,
		issuerSubjectKeyID,
	))
}

// GetAllPkiRevocationDistributionPoint returns all pKIRevocationDistributionPoint
func (k Keeper) GetAllPkiRevocationDistributionPoint(ctx sdk.Context) (list []types.PkiRevocationDistributionPoint) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.PkiRevocationDistributionPointKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.PkiRevocationDistributionPoint
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
