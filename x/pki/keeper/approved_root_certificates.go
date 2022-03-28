package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

// SetApprovedRootCertificates set approvedRootCertificates in the store.
func (k Keeper) SetApprovedRootCertificates(ctx sdk.Context, approvedRootCertificates types.ApprovedRootCertificates) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ApprovedRootCertificatesKeyPrefix))
	b := k.cdc.MustMarshal(&approvedRootCertificates)
	store.Set(types.ApprovedRootCertificatesKey, b)
}

// GetApprovedRootCertificates returns approvedRootCertificates.
func (k Keeper) GetApprovedRootCertificates(ctx sdk.Context) (val types.ApprovedRootCertificates, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ApprovedRootCertificatesKeyPrefix))

	b := store.Get(types.ApprovedRootCertificatesKey)
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)

	return val, true
}

// RemoveApprovedRootCertificates removes approvedRootCertificates from the store.
func (k Keeper) RemoveApprovedRootCertificates(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ApprovedRootCertificatesKeyPrefix))
	store.Delete(types.ApprovedRootCertificatesKey)
}

// Add root certificate to the list.
func (k Keeper) AddApprovedRootCertificate(ctx sdk.Context, certID types.CertificateIdentifier) {
	rootCertificates, _ := k.GetApprovedRootCertificates(ctx)

	// Check if the root cert is already there
	for _, existingCertID := range rootCertificates.Certs {
		if *existingCertID == certID {
			return
		}
	}

	rootCertificates.Certs = append(rootCertificates.Certs, &certID)

	k.SetApprovedRootCertificates(ctx, rootCertificates)
}

// Remove root certificate from the list.
func (k Keeper) RemoveApprovedRootCertificate(ctx sdk.Context, certID types.CertificateIdentifier) {
	rootCertificates, _ := k.GetApprovedRootCertificates(ctx)

	certIDIndex := -1

	for i, existingIdentifier := range rootCertificates.Certs {
		if *existingIdentifier == certID {
			certIDIndex = i

			break
		}
	}
	if certIDIndex == -1 {
		return
	}

	rootCertificates.Certs = append(rootCertificates.Certs[:certIDIndex], rootCertificates.Certs[certIDIndex+1:]...)

	k.SetApprovedRootCertificates(ctx, rootCertificates)
}
