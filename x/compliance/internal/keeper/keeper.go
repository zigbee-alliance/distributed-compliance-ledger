package keeper

import (
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliance/internal/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Keeper struct {
	// Unexposed key to access store from sdk.Context
	storeKey sdk.StoreKey

	// The wire codec for binary encoding/decoding
	cdc *codec.Codec
}

func NewKeeper(storeKey sdk.StoreKey, cdc *codec.Codec) Keeper {
	return Keeper{storeKey: storeKey, cdc: cdc}
}

// Gets the entire ModelInfo metadata struct for a id
func (k Keeper) GetModelInfo(ctx sdk.Context, id string) types.ModelInfo {
	if !k.IsModelInfoPresent(ctx, id) {
		panic("ModelInfo does not exist")
	}

	store := ctx.KVStore(k.storeKey)
	bz := store.Get([]byte(id))

	var device types.ModelInfo

	k.cdc.MustUnmarshalBinaryBare(bz, &device)

	return device
}

// Sets the entire ModelInfo metadata struct for a id
func (k Keeper) SetModelInfo(ctx sdk.Context, id string, device types.ModelInfo) {
	store := ctx.KVStore(k.storeKey)
	store.Set([]byte(id), k.cdc.MustMarshalBinaryBare(device))
}

// Deletes the ModelInfo from the store
func (k Keeper) DeleteModelInfo(ctx sdk.Context, id string) {
	if !k.IsModelInfoPresent(ctx, id) {
		panic("ModelInfo does not exist")
	}

	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(id))
}

// Get an iterator over all names in which the keys are the names and the values are the whois
func (k Keeper) GetModelInfoIDIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, nil)
}

// Check if the ModelInfo is present in the store or not
func (k Keeper) IsModelInfoPresent(ctx sdk.Context, id string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(id))
}
