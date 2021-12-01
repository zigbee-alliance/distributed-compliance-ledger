package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
)

// SetAccountStat set accountStat in the store
func (k Keeper) SetAccountStat(ctx sdk.Context, accountStat types.AccountStat) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AccountStatKey))
	b := k.cdc.MustMarshal(&accountStat)
	store.Set([]byte{0}, b)
}

// GetAccountStat returns accountStat
func (k Keeper) GetAccountStat(ctx sdk.Context) (val types.AccountStat, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AccountStatKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveAccountStat removes accountStat from the store
func (k Keeper) RemoveAccountStat(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AccountStatKey))
	store.Delete([]byte{0})
}
