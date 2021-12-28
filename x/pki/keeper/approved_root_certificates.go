package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

// SetApprovedRootCertificates set approvedRootCertificates in the store.
func (k Keeper) SetApprovedRootCertificates(ctx sdk.Context, approvedRootCertificates types.ApprovedRootCertificates) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ApprovedRootCertificatesKey))
	b := k.cdc.MustMarshal(&approvedRootCertificates)
	store.Set([]byte{0}, b)
}

// GetApprovedRootCertificates returns approvedRootCertificates.
func (k Keeper) GetApprovedRootCertificates(ctx sdk.Context) (val types.ApprovedRootCertificates, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ApprovedRootCertificatesKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveApprovedRootCertificates removes approvedRootCertificates from the store.
func (k Keeper) RemoveApprovedRootCertificates(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ApprovedRootCertificatesKey))
	store.Delete([]byte{0})
}

// Add root certificate to the list.
func (k Keeper) AddApprovedRootCertificate(ctx sdk.Context, certId types.CertificateIdentifier) {
	rootCertificates, _ := k.GetApprovedRootCertificates(ctx)

	// Check if the root cert is already there
	for _, existingСertId := range rootCertificates.Certs {
		if *existingСertId == certId {
			return
		}
	}

	rootCertificates.Certs = append(rootCertificates.Certs, &certId)

	k.SetApprovedRootCertificates(ctx, rootCertificates)
}

// Remove root certificate from the list.
func (k Keeper) RemoveApprovedRootCertificate(ctx sdk.Context, certId types.CertificateIdentifier) {
	rootCertificates, _ := k.GetApprovedRootCertificates(ctx)

	certIDIndex := -1
	for i, existingIdentifier := range rootCertificates.Certs {
		if *existingIdentifier == certId {
			certIDIndex = i
			break
		}
	}
	if certIDIndex == -1 {
		return
	}

	rootCertificates.Certs =
		append(rootCertificates.Certs[:certIDIndex], rootCertificates.Certs[certIDIndex+1:]...)

	k.SetApprovedRootCertificates(ctx, rootCertificates)
}
