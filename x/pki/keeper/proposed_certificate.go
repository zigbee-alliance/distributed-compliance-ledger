package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

// SetProposedCertificate set a specific proposedCertificate in the store from its index
func (k Keeper) SetProposedCertificate(ctx sdk.Context, proposedCertificate types.ProposedCertificate) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProposedCertificateKeyPrefix))
	b := k.cdc.MustMarshal(&proposedCertificate)
	store.Set(types.ProposedCertificateKey(
		proposedCertificate.Subject,
		proposedCertificate.SubjectKeyId,
	), b)
}

// GetProposedCertificate returns a proposedCertificate from its index
func (k Keeper) GetProposedCertificate(
	ctx sdk.Context,
	subject string,
	subjectKeyId string,

) (val types.ProposedCertificate, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProposedCertificateKeyPrefix))

	b := store.Get(types.ProposedCertificateKey(
		subject,
		subjectKeyId,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveProposedCertificate removes a proposedCertificate from the store
func (k Keeper) RemoveProposedCertificate(
	ctx sdk.Context,
	subject string,
	subjectKeyId string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProposedCertificateKeyPrefix))
	store.Delete(types.ProposedCertificateKey(
		subject,
		subjectKeyId,
	))
}

// GetAllProposedCertificate returns all proposedCertificate
func (k Keeper) GetAllProposedCertificate(ctx sdk.Context) (list []types.ProposedCertificate) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProposedCertificateKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ProposedCertificate
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
