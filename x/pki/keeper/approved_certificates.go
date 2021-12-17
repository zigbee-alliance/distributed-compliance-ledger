package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

// SetApprovedCertificates set a specific approvedCertificates in the store from its index
func (k Keeper) SetApprovedCertificates(ctx sdk.Context, approvedCertificates types.ApprovedCertificates) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ApprovedCertificatesKeyPrefix))
	b := k.cdc.MustMarshal(&approvedCertificates)
	store.Set(types.ApprovedCertificatesKey(
		approvedCertificates.Subject,
		approvedCertificates.SubjectKeyId,
	), b)
}

// GetApprovedCertificates returns a approvedCertificates from its index
func (k Keeper) GetApprovedCertificates(
	ctx sdk.Context,
	subject string,
	subjectKeyId string,

) (val types.ApprovedCertificates, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ApprovedCertificatesKeyPrefix))

	b := store.Get(types.ApprovedCertificatesKey(
		subject,
		subjectKeyId,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveApprovedCertificates removes a approvedCertificates from the store
func (k Keeper) RemoveApprovedCertificates(
	ctx sdk.Context,
	subject string,
	subjectKeyId string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ApprovedCertificatesKeyPrefix))
	store.Delete(types.ApprovedCertificatesKey(
		subject,
		subjectKeyId,
	))
}

// GetAllApprovedCertificates returns all approvedCertificates
func (k Keeper) GetAllApprovedCertificates(ctx sdk.Context) (list []types.ApprovedCertificates) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ApprovedCertificatesKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ApprovedCertificates
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
