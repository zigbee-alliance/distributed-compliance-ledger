package keeper

import (
	"fmt"
	"math"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/libs/log"
	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	authTypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

type (
	Keeper struct {
		cdc      codec.BinaryCodec
		storeKey sdk.StoreKey
		memKey   sdk.StoreKey

		dclauthKeeper types.DclauthKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey sdk.StoreKey,

	dclauthKeeper types.DclauthKeeper,
) *Keeper {
	return &Keeper{
		cdc:      cdc,
		storeKey: storeKey,
		memKey:   memKey,

		dclauthKeeper: dclauthKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", pkitypes.ModuleName))
}

func (k Keeper) CertificateApprovalsCount(ctx sdk.Context, authKeeper types.DclauthKeeper) int {
	return int(math.Ceil(types.RootCertificateApprovalsPercent *
		float64(authKeeper.CountAccountsWithRole(ctx, authTypes.Trustee))))
}

func (k Keeper) CertificateRejectApprovalsCount(ctx sdk.Context, authKeeper types.DclauthKeeper) int {
	return authKeeper.CountAccountsWithRole(ctx, authTypes.Trustee) - k.CertificateApprovalsCount(ctx, authKeeper) + 1
}

func (k Keeper) EnsureVidMatches(ctx sdk.Context, owner string, signer string) error {
	// get signer VID
	signerAddr, err := sdk.AccAddressFromBech32(signer)
	if err != nil {
		return pkitypes.NewErrInvalidAddress(err)
	}

	signerAccount, _ := k.dclauthKeeper.GetAccountO(ctx, signerAddr)
	signerVid := signerAccount.VendorID

	// get owner VID
	ownerAddr, err := sdk.AccAddressFromBech32(owner)
	if err != nil {
		return pkitypes.NewErrInvalidAddress(err)
	}

	ownerAccount, _ := k.dclauthKeeper.GetAccountO(ctx, ownerAddr)
	ownerVid := ownerAccount.VendorID

	if signerVid != ownerVid {
		return pkitypes.NewErrUnauthorizedCertVendor(ownerVid)
	}

	return nil
}

func removeCertFromList(issuer string, serialNumber string, certs *[]*types.Certificate) {
	certIndex := -1

	for i, cert := range *certs {
		if cert.SerialNumber == serialNumber && cert.Issuer == issuer {
			certIndex = i

			break
		}
	}
	if certIndex == -1 {
		return
	}
	*certs = append((*certs)[:certIndex], (*certs)[certIndex+1:]...)
}

func findCertificate(serialNumber string, certificates *[]*types.Certificate) (*types.Certificate, bool) {
	for _, cert := range *certificates {
		if cert.SerialNumber == serialNumber {
			return cert, true
		}
	}

	return nil, false
}
