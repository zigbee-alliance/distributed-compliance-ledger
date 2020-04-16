package keeper

import (
	"fmt"
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

const (
	certifiedModelPrefix  = "cm"
	vendorCertifiedModels = "vcm"
)

func NewKeeper(storeKey sdk.StoreKey, cdc *codec.Codec) Keeper {
	return Keeper{storeKey: storeKey, cdc: cdc}
}

// Gets the entire CertifiedModel struct for a CertifiedModelID
func (k Keeper) GetCertifiedModel(ctx sdk.Context, vid int16, pid int16) types.CertifiedModel {
	if !k.IsCertifiedModelPresent(ctx, vid, pid) {
		panic("CertifiedModel does not exist")
	}

	store := ctx.KVStore(k.storeKey)
	bz := store.Get([]byte(CertifiedModelId(vid, pid)))

	var device types.CertifiedModel

	k.cdc.MustUnmarshalBinaryBare(bz, &device)

	return device
}

// Sets the entire CertifiedModel metadata struct for a CertifiedModelID
func (k Keeper) SetCertifiedModel(ctx sdk.Context, model types.CertifiedModel) {
	store := ctx.KVStore(k.storeKey)
	store.Set([]byte(CertifiedModelId(model.VID, model.PID)), k.cdc.MustMarshalBinaryBare(model))
}

// Deletes the CertifiedModel from the store
func (k Keeper) DeleteCertifiedModel(ctx sdk.Context, vid int16, pid int16) {
	if !k.IsCertifiedModelPresent(ctx, vid, pid) {
		panic("CertifiedModel does not exist")
	}

	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(CertifiedModelId(vid, pid)))
}

// Iterate over all CertifiedModels
func (k Keeper) IterateCertifiedModels(ctx sdk.Context, process func(info types.CertifiedModel) (stop bool)) {
	store := ctx.KVStore(k.storeKey)

	iter := sdk.KVStorePrefixIterator(store, []byte(certifiedModelPrefix))
	defer iter.Close()

	for {
		if !iter.Valid() {
			return
		}

		val := iter.Value()

		var certifiedModel types.CertifiedModel

		k.cdc.MustUnmarshalBinaryBare(val, &certifiedModel)

		if process(certifiedModel) {
			return
		}

		iter.Next()
	}
}

func (k Keeper) CountTotalCertifiedModel(ctx sdk.Context) int {
	return k.countTotal(ctx, certifiedModelPrefix)
}

// Check if the CertifiedModel is present in the store or not
func (k Keeper) IsCertifiedModelPresent(ctx sdk.Context, vid int16, pid int16) bool {
	return k.isRecordPresent(ctx, CertifiedModelId(vid, pid))
}

// Id builder for CertifiedModel
func CertifiedModelId(vid interface{}, pid interface{}) string {
	return fmt.Sprintf("%s:%v:%v", certifiedModelPrefix, vid, pid)
}

// Check if the record is present in the store or not
func (k Keeper) isRecordPresent(ctx sdk.Context, id string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(id))
}

func (k Keeper) countTotal(ctx sdk.Context, prefix string) int {
	store := ctx.KVStore(k.storeKey)
	res := 0

	iter := sdk.KVStorePrefixIterator(store, []byte(prefix))
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		res++
	}

	return res
}
