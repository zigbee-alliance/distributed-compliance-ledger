package keeper

import (
	"math"

	sdk "github.com/cosmos/cosmos-sdk/types"
	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	authTypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

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

func RemoveCertFromList(issuer string, serialNumber string, certs *[]*types.Certificate) {
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

func FindCertificateInList(serialNumber string, certificates *[]*types.Certificate) (*types.Certificate, bool) {
	for _, cert := range *certificates {
		if cert.SerialNumber == serialNumber {
			return cert, true
		}
	}

	return nil, false
}

func FilterCertificateList(certificates *[]*types.Certificate, predicate CertificatePredicate) []*types.Certificate {
	var result []*types.Certificate

	for _, s := range *certificates {
		if predicate(s) {
			result = append(result, s)
		}
	}

	return result
}

func (k msgServer) AddCertificateToGlobalCertificateIndexes(
	ctx sdk.Context,
	certificate types.Certificate,
) {
	// add to the global list of certificates
	k.AddAllCertificate(ctx, certificate)
	// add to the global list of certificates indexed by subject key id
	k.AddAllCertificateBySubjectKeyID(ctx, certificate)
	// add to the global list of certificates indexed by subject
	k.AddAllCertificateBySubject(ctx, certificate.Subject, certificate.SubjectKeyId)
}

func (k msgServer) RemoveCertificateFromGlobalCertificateIndexes(
	ctx sdk.Context,
	subject string,
	subjectKeyID string,
) {
	// remove from the global list of certificates
	k.RemoveAllCertificates(ctx, subject, subjectKeyID)
	// remove from the global list of certificates indexed by subject key id
	k.RemoveAllCertificatesBySubjectKeyID(ctx, subject, subjectKeyID)
	// remove from the global list of certificates indexed by subject
	k.RemoveAllCertificateBySubject(ctx, subject, subjectKeyID)
}

func (k msgServer) StoreDaCertificate(
	ctx sdk.Context,
	certificate types.Certificate,
	isRoot bool,
) {
	// add to Global certificates indexes
	k.AddCertificateToGlobalCertificateIndexes(ctx, certificate)

	// add to list of certificates with the same Subject/SubjectKeyID combination and store updated list
	k.AddApprovedCertificate(ctx, certificate)

	// add to list of certificates indexed by subject
	k.AddApprovedCertificateBySubject(ctx, certificate.Subject, certificate.SubjectKeyId)

	// add to list of certificates indexed by subject key id
	k.AddApprovedCertificateBySubjectKeyID(ctx, certificate)

	if isRoot {
		// add to root certificates index
		k.AddApprovedRootCertificate(ctx, certificate)
	} else {
		// add the certificate identifier to the issuer's Child Certificates record
		k.AddChildCertificate(ctx, certificate)
	}
}

func (k msgServer) RemoveDaCertificate(
	ctx sdk.Context,
	subject string,
	subjectKeyID string,
	isRoot bool,
) {
	// remove from global list
	k.RemoveCertificateFromGlobalCertificateIndexes(ctx, subject, subjectKeyID)
	// remove from approved certificates map
	k.RemoveApprovedCertificates(ctx, subject, subjectKeyID)
	// remove from subject -> subject key ID map
	k.RemoveApprovedCertificateBySubject(ctx, subject, subjectKeyID)
	// remove from subject key ID -> certificates map
	k.RemoveApprovedCertificatesBySubjectKeyID(ctx, subject, subjectKeyID)
	if isRoot {
		k.RemoveApprovedRootCertificate(ctx, subject, subjectKeyID)
	}
}

func (k msgServer) RemoveDaCertificateBySerialNumber(
	ctx sdk.Context,
	subject string,
	subjectKeyID string,
	certificates *types.ApprovedCertificates,
	serialNumber string,
	issuer string,
	isRoot bool,
) {
	RemoveCertFromList(issuer, serialNumber, &certificates.Certs)

	if len(certificates.Certs) == 0 {
		k.RemoveDaCertificate(ctx, subject, subjectKeyID, isRoot)
	} else {
		k.RemoveAllCertificatesBySerialNumber(ctx, subject, subjectKeyID, serialNumber)
		k.RemoveAllCertificatesBySubjectKeyIDBySerialNumber(ctx, subject, subjectKeyID, serialNumber)
		k.RemoveApprovedCertificatesBySerialNumber(ctx, subject, subjectKeyID, serialNumber)
		k.RemoveApprovedCertificatesBySubjectKeyIDBySerialNumber(ctx, subject, subjectKeyID, serialNumber)
	}
}

func (k msgServer) StoreNocCertificate(
	ctx sdk.Context,
	certificate types.Certificate,
	isRoot bool) {
	// add to Global certificates indexes
	k.AddCertificateToGlobalCertificateIndexes(ctx, certificate)

	// add to the list of all NOC certificates
	k.AddNocCertificate(ctx, certificate)

	// add to certificates map indexed by vid/skid
	k.AddNocCertificateByVidAndSkid(ctx, certificate)

	// add to certificates map indexed by subject
	k.AddNocCertificateBySubject(ctx, certificate)

	// add to certificates map indexed by subject key id
	k.AddNocCertificateBySubjectKeyID(ctx, certificate)

	if isRoot {
		// add to the list of NOC root certificates with the same VID
		k.AddNocRootCertificate(ctx, certificate)
	} else {
		// add to the list of NOC ica certificates with the same VID
		k.AddNocIcaCertificate(ctx, certificate)
		// add the certificate identifier to the issuer's Child Certificates record
		k.AddChildCertificate(ctx, certificate)
	}
}

func (k msgServer) RemoveNocCertificate(
	ctx sdk.Context,
	subject string,
	subjectKeyID string,
	accountVid int32,
	isRoot bool,
) {
	// remove from global list
	k.RemoveCertificateFromGlobalCertificateIndexes(ctx, subject, subjectKeyID)
	// remove from noc certificates map
	k.RemoveNocCertificates(ctx, subject, subjectKeyID)
	// remove from vid, subject key id map
	k.RemoveNocCertificatesByVidAndSkid(ctx, accountVid, subjectKeyID)
	// remove from subject -> subject key ID map
	k.RemoveNocCertificateBySubject(ctx, subject, subjectKeyID)
	// remove from subject key ID -> certificates map
	k.RemoveNocCertificatesBySubjectAndSubjectKeyID(ctx, subject, subjectKeyID)
	if isRoot {
		// remove from noc root certificates map
		k.RemoveNocRootCertificate(ctx, subject, subjectKeyID, accountVid)
	} else {
		// remove from noc ica certificates map
		k.RemoveNocIcaCertificate(ctx, subject, subjectKeyID, accountVid)
	}
}

func (k msgServer) RemoveNocCertBySerialNumber(
	ctx sdk.Context,
	subject string,
	subjectKeyID string,
	certificates *types.NocCertificates,
	accountVid int32,
	serialNumber string,
	issuer string,
	isRoot bool,
) {
	RemoveCertFromList(issuer, serialNumber, &certificates.Certs)

	if len(certificates.Certs) == 0 {
		k.RemoveNocCertificate(ctx, subject, subjectKeyID, accountVid, isRoot)
	} else {
		k.RemoveAllCertificatesBySerialNumber(ctx, subject, subjectKeyID, serialNumber)
		k.RemoveAllCertificatesBySubjectKeyIDBySerialNumber(ctx, subject, subjectKeyID, serialNumber)
		k.RemoveNocCertificatesBySerialNumber(ctx, subject, subjectKeyID, serialNumber)
		k.RemoveNocCertificatesBySubjectKeyIDBySerialNumber(ctx, subject, subjectKeyID, serialNumber)
		k.RemoveNocCertificatesByVidAndSkidBySerialNumber(ctx, accountVid, subject, subjectKeyID, serialNumber)
		if isRoot {
			k.RemoveNocRootCertificateBySerialNumber(ctx, subject, subjectKeyID, accountVid, serialNumber)
		} else {
			k.RemoveNocIcaCertificateBySerialNumber(ctx, subject, subjectKeyID, accountVid, serialNumber)
		}
	}
}
