package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

// SetPkiRevocationDistributionPointsByIssuerSubjectKeyID set a specific pkiRevocationDistributionPointsByIssuerSubjectKeyID in the store from its index.
func (k Keeper) SetPkiRevocationDistributionPointsByIssuerSubjectKeyID(ctx sdk.Context, pkiRevocationDistributionPointsByIssuerSubjectKeyID types.PkiRevocationDistributionPointsByIssuerSubjectKeyID) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.PkiRevocationDistributionPointsByIssuerSubjectKeyIDKeyPrefix))
	b := k.cdc.MustMarshal(&pkiRevocationDistributionPointsByIssuerSubjectKeyID)
	store.Set(types.PkiRevocationDistributionPointsByIssuerSubjectKeyIDKey(
		pkiRevocationDistributionPointsByIssuerSubjectKeyID.IssuerSubjectKeyID,
	), b)
}

// GetPkiRevocationDistributionPointsByIssuerSubjectKeyID returns a pkiRevocationDistributionPointsByIssuerSubjectKeyID from its index.
func (k Keeper) GetPkiRevocationDistributionPointsByIssuerSubjectKeyID(
	ctx sdk.Context,
	issuerSubjectKeyID string,

) (val types.PkiRevocationDistributionPointsByIssuerSubjectKeyID, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.PkiRevocationDistributionPointsByIssuerSubjectKeyIDKeyPrefix))

	b := store.Get(types.PkiRevocationDistributionPointsByIssuerSubjectKeyIDKey(
		issuerSubjectKeyID,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)

	return val, true
}

// RemovePkiRevocationDistributionPointsByIssuerSubjectKeyID removes a pkiRevocationDistributionPointsByIssuerSubjectKeyID from the store.
func (k Keeper) RemovePkiRevocationDistributionPointsByIssuerSubjectKeyID(
	ctx sdk.Context,
	issuerSubjectKeyID string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.PkiRevocationDistributionPointsByIssuerSubjectKeyIDKeyPrefix))
	store.Delete(types.PkiRevocationDistributionPointsByIssuerSubjectKeyIDKey(
		issuerSubjectKeyID,
	))
}

// GetAllPkiRevocationDistributionPointsByIssuerSubjectKeyID returns all pkiRevocationDistributionPointsByIssuerSubjectKeyID.
func (k Keeper) GetAllPkiRevocationDistributionPointsByIssuerSubjectKeyID(ctx sdk.Context) (list []types.PkiRevocationDistributionPointsByIssuerSubjectKeyID) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.PkiRevocationDistributionPointsByIssuerSubjectKeyIDKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.PkiRevocationDistributionPointsByIssuerSubjectKeyID
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// Add PkiRevocationDistributionPoint to a subjectKeyID->revocationPoint index.
func (k Keeper) AddPkiRevocationDistributionPointBySubjectKeyID(ctx sdk.Context, pkiRevocationDistributionPoint types.PkiRevocationDistributionPoint) {
	revocationPoints, isFound := k.GetPkiRevocationDistributionPointsByIssuerSubjectKeyID(ctx, pkiRevocationDistributionPoint.IssuerSubjectKeyID)

	if !isFound {
		revocationPoints.IssuerSubjectKeyID = pkiRevocationDistributionPoint.IssuerSubjectKeyID
	}

	// Check if revocation point is already there
	for _, revocationPoint := range revocationPoints.Points {
		if revocationPoint.Vid == pkiRevocationDistributionPoint.Vid &&
			revocationPoint.Label == pkiRevocationDistributionPoint.Label &&
			revocationPoint.IssuerSubjectKeyID == pkiRevocationDistributionPoint.IssuerSubjectKeyID {
			return
		}
	}
	revocationPoints.Points = append(revocationPoints.Points, &pkiRevocationDistributionPoint)

	k.SetPkiRevocationDistributionPointsByIssuerSubjectKeyID(ctx, revocationPoints)
}

// Remove PkiRevocationDistributionPoint from a subjectKeyID->revocationPoint index.
func (k Keeper) RemovePkiRevocationDistributionPointBySubjectKeyID(ctx sdk.Context, pkiRevocationDistributionPoint types.PkiRevocationDistributionPoint) {
	revocationPoints, _ := k.GetPkiRevocationDistributionPointsByIssuerSubjectKeyID(ctx, pkiRevocationDistributionPoint.IssuerSubjectKeyID)

	pointIndex := -1
	for i, revocationPoint := range revocationPoints.Points {
		if revocationPoint.Vid == pkiRevocationDistributionPoint.Vid &&
			revocationPoint.Label == pkiRevocationDistributionPoint.Label &&
			revocationPoint.IssuerSubjectKeyID == pkiRevocationDistributionPoint.IssuerSubjectKeyID {
			pointIndex = i

			break
		}
	}
	if pointIndex == -1 {
		return
	}

	revocationPoints.Points = append(revocationPoints.Points[:pointIndex], revocationPoints.Points[pointIndex+1:]...)

	if len(revocationPoints.Points) > 0 {
		k.SetPkiRevocationDistributionPointsByIssuerSubjectKeyID(ctx, revocationPoints)
	} else {
		k.RemovePkiRevocationDistributionPointsByIssuerSubjectKeyID(ctx, pkiRevocationDistributionPoint.IssuerSubjectKeyID)
	}
}

// Update PkiRevocationDistributionPoint in a subjectKeyID->revocationPoint index.
func (k Keeper) UpdatePkiRevocationDistributionPointBySubjectKeyID(ctx sdk.Context, pkiRevocationDistributionPoint types.PkiRevocationDistributionPoint) {
	revocationPoints, _ := k.GetPkiRevocationDistributionPointsByIssuerSubjectKeyID(ctx, pkiRevocationDistributionPoint.IssuerSubjectKeyID)

	// Check if revocation point is there
	for i, revocationPoint := range revocationPoints.Points {
		if revocationPoint.Vid == pkiRevocationDistributionPoint.Vid &&
			revocationPoint.Label == pkiRevocationDistributionPoint.Label &&
			revocationPoint.IssuerSubjectKeyID == pkiRevocationDistributionPoint.IssuerSubjectKeyID {
			revocationPoints.Points[i] = &pkiRevocationDistributionPoint
		}
	}

	k.SetPkiRevocationDistributionPointsByIssuerSubjectKeyID(ctx, revocationPoints)
}
