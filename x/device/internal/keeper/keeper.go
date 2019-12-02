package keeper

import (
	"github.com/askolesov/zb-ledger/x/device/internal/types"
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

// Gets the entire Device metadata struct for a id
func (k Keeper) GetDevice(ctx sdk.Context, id string) types.Device {
	store := ctx.KVStore(k.storeKey)

	if !k.IsIDPresent(ctx, id) {
		return types.NewDevice()
	}

	bz := store.Get([]byte(id))

	var device types.Device

	k.cdc.MustUnmarshalBinaryBare(bz, &device)

	return device
}

// Sets the entire Device metadata struct for a id
func (k Keeper) SetDevice(ctx sdk.Context, id string, device types.Device) {
	if device.Owner.Empty() {
		return
	}

	store := ctx.KVStore(k.storeKey)

	store.Set([]byte(id), k.cdc.MustMarshalBinaryBare(device))
}

// Check if the id is present in the store or not
func (k Keeper) IsIDPresent(ctx sdk.Context, id string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(id))
}
