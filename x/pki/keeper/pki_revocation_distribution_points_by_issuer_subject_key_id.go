package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

// SetPkiRevocationDistributionPointsByIssuerSubjectKeyId set a specific pkiRevocationDistributionPointsByIssuerSubjectKeyId in the store from its index
func (k Keeper) SetPkiRevocationDistributionPointsByIssuerSubjectKeyId(ctx sdk.Context, pkiRevocationDistributionPointsByIssuerSubjectKeyId types.PkiRevocationDistributionPointsByIssuerSubjectKeyId) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.PkiRevocationDistributionPointsByIssuerSubjectKeyIdKeyPrefix))
	b := k.cdc.MustMarshal(&pkiRevocationDistributionPointsByIssuerSubjectKeyId)
	store.Set(types.PkiRevocationDistributionPointsByIssuerSubjectKeyIdKey(
		pkiRevocationDistributionPointsByIssuerSubjectKeyId.IssuerSubjectKeyId,
	), b)
}

// GetPkiRevocationDistributionPointsByIssuerSubjectKeyId returns a pkiRevocationDistributionPointsByIssuerSubjectKeyId from its index
func (k Keeper) GetPkiRevocationDistributionPointsByIssuerSubjectKeyId(
	ctx sdk.Context,
	issuerSubjectKeyId string,

) (val types.PkiRevocationDistributionPointsByIssuerSubjectKeyId, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.PkiRevocationDistributionPointsByIssuerSubjectKeyIdKeyPrefix))

	b := store.Get(types.PkiRevocationDistributionPointsByIssuerSubjectKeyIdKey(
		issuerSubjectKeyId,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)

	return val, true
}

// RemovePkiRevocationDistributionPointsByIssuerSubjectKeyId removes a pkiRevocationDistributionPointsByIssuerSubjectKeyId from the store
func (k Keeper) RemovePkiRevocationDistributionPointsByIssuerSubjectKeyId(
	ctx sdk.Context,
	issuerSubjectKeyId string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.PkiRevocationDistributionPointsByIssuerSubjectKeyIdKeyPrefix))
	store.Delete(types.PkiRevocationDistributionPointsByIssuerSubjectKeyIdKey(
		issuerSubjectKeyId,
	))
}

// GetAllPkiRevocationDistributionPointsByIssuerSubjectKeyId returns all pkiRevocationDistributionPointsByIssuerSubjectKeyId
func (k Keeper) GetAllPkiRevocationDistributionPointsByIssuerSubjectKeyId(ctx sdk.Context) (list []types.PkiRevocationDistributionPointsByIssuerSubjectKeyId) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.PkiRevocationDistributionPointsByIssuerSubjectKeyIdKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.PkiRevocationDistributionPointsByIssuerSubjectKeyId
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
