package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

// SetProposedCertificateRevocation set a specific proposedCertificateRevocation in the store from its index
func (k Keeper) SetProposedCertificateRevocation(ctx sdk.Context, proposedCertificateRevocation types.ProposedCertificateRevocation) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProposedCertificateRevocationKeyPrefix))
	b := k.cdc.MustMarshal(&proposedCertificateRevocation)
	store.Set(types.ProposedCertificateRevocationKey(
		proposedCertificateRevocation.Subject,
		proposedCertificateRevocation.SubjectKeyId,
	), b)
}

// GetProposedCertificateRevocation returns a proposedCertificateRevocation from its index
func (k Keeper) GetProposedCertificateRevocation(
	ctx sdk.Context,
	subject string,
	subjectKeyId string,

) (val types.ProposedCertificateRevocation, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProposedCertificateRevocationKeyPrefix))

	b := store.Get(types.ProposedCertificateRevocationKey(
		subject,
		subjectKeyId,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveProposedCertificateRevocation removes a proposedCertificateRevocation from the store
func (k Keeper) RemoveProposedCertificateRevocation(
	ctx sdk.Context,
	subject string,
	subjectKeyId string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProposedCertificateRevocationKeyPrefix))
	store.Delete(types.ProposedCertificateRevocationKey(
		subject,
		subjectKeyId,
	))
}

// GetAllProposedCertificateRevocation returns all proposedCertificateRevocation
func (k Keeper) GetAllProposedCertificateRevocation(ctx sdk.Context) (list []types.ProposedCertificateRevocation) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProposedCertificateRevocationKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ProposedCertificateRevocation
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
