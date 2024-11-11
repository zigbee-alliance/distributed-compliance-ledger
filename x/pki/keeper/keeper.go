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

func (k msgServer) AddCertificateToAllCertificateIndexes(ctx sdk.Context, certificate types.Certificate) {
	// Add to the global list of certificates
	k.AddAllCertificate(ctx, certificate)

	// append to global list of certificates indexed by subject
	k.AddAllCertificateBySubject(ctx, certificate.Subject, certificate.SubjectKeyId)
}

func (k msgServer) AddCertificateToDaCertificateIndexes(
	ctx sdk.Context,
	certificate types.Certificate,
	isRoot bool) {
	// append new certificate to list of certificates with the same Subject/SubjectKeyID combination and store updated list
	k.AddApprovedCertificate(ctx, certificate)

	// add to subject -> subject key ID map
	k.AddApprovedCertificateBySubject(ctx, certificate.Subject, certificate.SubjectKeyId)

	// add to subject key ID -> certificates map
	k.AddApprovedCertificateBySubjectKeyID(ctx, certificate)

	if isRoot {
		// add to root certificates index
		k.AddApprovedRootCertificate(ctx, certificate)
	} else {
		// add the certificate identifier to the issuer's Child Certificates record
		k.AddChildCertificate(ctx, certificate)
	}
}

func (k msgServer) AddCertificateToNocCertificateIndexes(
	ctx sdk.Context,
	certificate types.Certificate,
	isRoot bool) {
	// Add to the list of all NOC certificates
	k.AddNocCertificate(ctx, certificate)

	// add to certificates map indexed by { vid, subject key id }
	k.AddNocCertificateByVidAndSkid(ctx, certificate)

	// add to certificates map indexed by { subject }
	k.AddNocCertificateBySubject(ctx, certificate)

	// add to certificates map indexed by { subject key id }
	k.AddNocCertificateBySubjectKeyID(ctx, certificate)

	if isRoot {
		// Add to the list of NOC root certificates with the same VID
		k.AddNocRootCertificate(ctx, certificate)
	} else {
		// Add to the list of NOC ica certificates with the same VID
		k.AddNocIcaCertificate(ctx, certificate)
		// add the certificate identifier to the issuer's Child Certificates record
		k.AddChildCertificate(ctx, certificate)
	}
}

func (k msgServer) RemoveCertificateFromAllCertificateIndexes(ctx sdk.Context, certID types.CertificateIdentifier) {
	// remove from global certificates map
	k.RemoveAllCertificates(ctx, certID.Subject, certID.SubjectKeyId)
	// remove from global subject -> subject key ID map
	k.RemoveAllCertificateBySubject(ctx, certID.Subject, certID.SubjectKeyId)
}

func (k msgServer) RemoveCertificateFromDaCertificateIndexes(
	ctx sdk.Context,
	certID types.CertificateIdentifier,
	isRoot bool) {
	// remove from approved certificates map
	k.RemoveApprovedCertificates(ctx, certID.Subject, certID.SubjectKeyId)
	// remove from subject -> subject key ID map
	k.RemoveApprovedCertificateBySubject(ctx, certID.Subject, certID.SubjectKeyId)
	// remove from subject key ID -> certificates map
	k.RemoveApprovedCertificatesBySubjectKeyID(ctx, certID.Subject, certID.SubjectKeyId)
	if isRoot {
		k.RemoveApprovedRootCertificate(ctx, certID)
	}
}

func (k msgServer) RemoveCertificateFromNocCertificateIndexes(
	ctx sdk.Context,
	certID types.CertificateIdentifier,
	accountVid int32,
	isRoot bool) {
	// remove from noc certificates map
	k.RemoveNocCertificates(ctx, certID.Subject, certID.SubjectKeyId)
	// remove from vid, subject key id map
	k.RemoveNocCertificatesByVidAndSkid(ctx, accountVid, certID.SubjectKeyId)
	// remove from subject -> subject key ID map
	k.RemoveNocCertificateBySubject(ctx, certID.Subject, certID.SubjectKeyId)
	// remove from subject key ID -> certificates map
	k.RemoveNocCertificatesBySubjectAndSubjectKeyID(ctx, certID.Subject, certID.SubjectKeyId)
	if isRoot {
		// remove from noc root certificates map
		k.RemoveNocRootCertificate(ctx, certID.Subject, certID.SubjectKeyId, accountVid)
	} else {
		// remove from noc ica certificates map
		k.RemoveNocIcaCertificate(ctx, certID.Subject, certID.SubjectKeyId, accountVid)
	}
}

func (k msgServer) removeDaX509Cert(
	ctx sdk.Context,
	certID types.CertificateIdentifier,
	certificates *types.ApprovedCertificates,
	serialNumber string) {
	if len(certificates.Certs) == 0 {
		// remove from global certificates map
		k.RemoveCertificateFromAllCertificateIndexes(ctx, certID)
		// remove from noc certificates map
		k.RemoveCertificateFromDaCertificateIndexes(ctx, certID, false)
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
		// remove from global certificates map
		k.RemoveCertificateFromAllCertificateIndexes(ctx, certID)
		// remove from noc certificates map
		k.RemoveCertificateFromNocCertificateIndexes(ctx, certID, accountVid, isRoot)
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
