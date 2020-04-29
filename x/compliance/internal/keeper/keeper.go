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
	complianceInfoPrefix = "1"
)

func NewKeeper(storeKey sdk.StoreKey, cdc *codec.Codec) Keeper {
	return Keeper{storeKey: storeKey, cdc: cdc}
}

// Gets the entire ComplianceInfo struct for a ComplianceInfoID
func (k Keeper) GetComplianceInfo(ctx sdk.Context, certificationType types.CertificationType, vid uint16, pid uint16) types.ComplianceInfo {
	if !k.IsComplianceInfoPresent(ctx, certificationType, vid, pid) {
		panic("ComplianceInfo does not exist")
	}

	store := ctx.KVStore(k.storeKey)
	bz := store.Get([]byte(ComplianceInfoId(certificationType, vid, pid)))

	var device types.ComplianceInfo

	k.cdc.MustUnmarshalBinaryBare(bz, &device)

	return device
}

// Sets the entire ComplianceInfo metadata struct for a ComplianceInfoID
func (k Keeper) SetComplianceInfo(ctx sdk.Context, model types.ComplianceInfo) {
	store := ctx.KVStore(k.storeKey)
	store.Set([]byte(ComplianceInfoId(model.CertificationType, model.VID, model.PID)), k.cdc.MustMarshalBinaryBare(model))
}

// Iterate over all ComplianceInfos
func (k Keeper) IterateComplianceInfos(ctx sdk.Context, certificationType types.CertificationType, process func(info types.ComplianceInfo) (stop bool)) {
	store := ctx.KVStore(k.storeKey)

	iter := sdk.KVStorePrefixIterator(store, []byte(prefix(certificationType)))
	defer iter.Close()

	for {
		if !iter.Valid() {
			return
		}

		val := iter.Value()

		var certifiedModel types.ComplianceInfo

		k.cdc.MustUnmarshalBinaryBare(val, &certifiedModel)

		if process(certifiedModel) {
			return
		}

		iter.Next()
	}
}

func (k Keeper) CountTotalComplianceInfo(ctx sdk.Context, certificationType types.CertificationType) int {
	return k.countTotal(ctx, prefix(certificationType))
}

// Check if the ComplianceInfo is present in the store or not
func (k Keeper) IsComplianceInfoPresent(ctx sdk.Context, certificationType types.CertificationType, vid uint16, pid uint16) bool {
	return k.isRecordPresent(ctx, ComplianceInfoId(certificationType, vid, pid))
}

// Id builder for ComplianceInfo
func ComplianceInfoId(certificationType types.CertificationType, vid interface{}, pid interface{}) string {
	return fmt.Sprintf("%s:%v:%v", prefix(certificationType), vid, pid)
}

func prefix(certificationType types.CertificationType) string {
	return fmt.Sprintf("%s:%v", complianceInfoPrefix, certificationType)
}

// Check if the record is present in the store or not
func (k Keeper) isRecordPresent(ctx sdk.Context, id string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(id))
}

//  TODO: Is iteration the only way to calculate the total number of elements?
//  It looks like that in a non-pagination case we iterate twice: to get the total number of elements and to get the real content.
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
