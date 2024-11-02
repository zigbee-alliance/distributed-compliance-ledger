package keeper

import (
	"fmt"
	"math"

	"github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	authTypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

type (
	Keeper struct {
		cdc      codec.BinaryCodec
		storeKey storetypes.StoreKey
		memKey   storetypes.StoreKey

		dclauthKeeper types.DclauthKeeper
	}
	CertificatePredicate func(*types.Certificate) bool
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,

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

func filterCertificates(certificates *[]*types.Certificate, predicate CertificatePredicate) []*types.Certificate {
	var result []*types.Certificate

	for _, s := range *certificates {
		if predicate(s) {
			result = append(result, s)
		}
	}

	return result
}

func (k msgServer) removeApprovedX509Cert(ctx sdk.Context, certID types.CertificateIdentifier, certificates *types.ApprovedCertificates, serialNumber string) {
	if len(certificates.Certs) == 0 {
		k.RemoveAllCertificates(ctx, certID.Subject, certID.SubjectKeyId)
		k.RemoveApprovedCertificates(ctx, certID.Subject, certID.SubjectKeyId)
		k.RemoveApprovedCertificateBySubject(ctx, certID.Subject, certID.SubjectKeyId)
		k.RemoveApprovedCertificatesBySubjectKeyID(ctx, certID.Subject, certID.SubjectKeyId)
	} else {
		k.RemoveAllCertificatesBySerialNumber(ctx, certID.Subject, certID.SubjectKeyId, serialNumber)
		k.RemoveApprovedCertificatesBySerialNumber(ctx, certID.Subject, certID.SubjectKeyId, serialNumber)
		k.RemoveApprovedCertificatesBySubjectKeyIDBySerialNumber(ctx, certID.Subject, certID.SubjectKeyId, serialNumber)
	}
}

func (k msgServer) removeNocX509Cert(
	ctx sdk.Context,
	certID types.CertificateIdentifier,
	certificates *types.NocCertificates,
	accountVid int32,
	serialNumber string,
	isRoot bool,
) {
	if len(certificates.Certs) == 0 { //nolint:nestif
		k.RemoveAllCertificates(ctx, certID.Subject, certID.SubjectKeyId)
		k.RemoveNocCertificates(ctx, certID.Subject, certID.SubjectKeyId)
		k.RemoveNocCertificateBySubject(ctx, certID.Subject, certID.SubjectKeyId)
		k.RemoveNocCertificatesBySubjectAndSubjectKeyID(ctx, certID.Subject, certID.SubjectKeyId)
		k.RemoveNocCertificatesByVidAndSkid(ctx, accountVid, certID.SubjectKeyId)

		if isRoot {
			k.RemoveNocRootCertificate(ctx, certID.Subject, certID.SubjectKeyId, accountVid)
		} else {
			k.RemoveNocIcaCertificate(ctx, certID.Subject, certID.SubjectKeyId, accountVid)
		}
	} else {
		k.RemoveAllCertificatesBySerialNumber(ctx, certID.Subject, certID.SubjectKeyId, serialNumber)
		k.RemoveNocCertificatesBySerialNumber(ctx, certID.Subject, certID.SubjectKeyId, serialNumber)
		k.RemoveNocCertificatesBySubjectKeyIDBySerialNumber(ctx, certID.Subject, certID.SubjectKeyId, serialNumber)
		k.RemoveNocCertificatesByVidAndSkidBySerialNumber(ctx, accountVid, certID.Subject, certID.SubjectKeyId, serialNumber)

		if isRoot {
			k.RemoveNocRootCertificateBySerialNumber(ctx, certID.Subject, certID.SubjectKeyId, accountVid, serialNumber)
		} else {
			k.RemoveNocIcaCertificateBySerialNumber(ctx, certID.Subject, certID.SubjectKeyId, accountVid, serialNumber)
		}
	}
}
