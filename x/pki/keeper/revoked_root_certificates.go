package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

// SetRevokedRootCertificates set revokedRootCertificates in the store.
func (k Keeper) SetRevokedRootCertificates(ctx sdk.Context, revokedRootCertificates types.RevokedRootCertificates) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.RevokedRootCertificatesKeyPrefix))
	b := k.cdc.MustMarshal(&revokedRootCertificates)
	store.Set(pkitypes.RevokedRootCertificatesKey, b)
}

// GetRevokedRootCertificates returns revokedRootCertificates.
func (k Keeper) GetRevokedRootCertificates(ctx sdk.Context) (val types.RevokedRootCertificates, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.RevokedRootCertificatesKeyPrefix))

	b := store.Get(pkitypes.RevokedRootCertificatesKey)
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)

	return val, true
}

// RemoveRevokedRootCertificates removes revokedRootCertificates from the store.
func (k Keeper) RemoveRevokedRootCertificates(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.RevokedRootCertificatesKeyPrefix))
	store.Delete(pkitypes.RevokedRootCertificatesKey)
}

// Add revoked root certificate to the list.
func (k Keeper) AddRevokedRootCertificate(ctx sdk.Context, certID types.CertificateIdentifier) {
	rootCertificates, _ := k.GetRevokedRootCertificates(ctx)

	// Check if the root cert is already there
	for _, existingCertID := range rootCertificates.Certs {
		if *existingCertID == certID {
			return
		}
	}

	rootCertificates.Certs = append(rootCertificates.Certs, &certID)

	k.SetRevokedRootCertificates(ctx, rootCertificates)
}

// Remove revoked root certificate from the list.
func (k Keeper) RemoveRevokedRootCertificate(ctx sdk.Context, certID types.CertificateIdentifier) {
	rootCertificates, _ := k.GetRevokedRootCertificates(ctx)

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

	k.SetRevokedRootCertificates(ctx, rootCertificates)
}
