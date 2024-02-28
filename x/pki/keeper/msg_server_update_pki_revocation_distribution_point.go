package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	dclauthtypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/x509"
)

func (k msgServer) UpdatePkiRevocationDistributionPoint(goCtx context.Context, msg *types.MsgUpdatePkiRevocationDistributionPoint) (*types.MsgUpdatePkiRevocationDistributionPointResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// check if signer has vendor role
	signerAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return nil, pkitypes.NewErrInvalidAddress(err)
	}
	if !k.dclauthKeeper.HasRole(ctx, signerAddr, dclauthtypes.Vendor) {
		return nil, pkitypes.NewErrUnauthorizedRole("MsgUpdatePkiRevocationDistributionPoint", dclauthtypes.Vendor)
	}

	// compare VID in message and Vendor acount
	signerAccount, _ := k.dclauthKeeper.GetAccountO(ctx, signerAddr)
	if msg.Vid != signerAccount.VendorID {
		return nil, pkitypes.NewErrMessageVidNotEqualAccountVid(msg.Vid, signerAccount.VendorID)
	}

	// check that Revocation Point exists
	pkiRevocationDistributionPoint, isFound := k.GetPkiRevocationDistributionPoint(ctx, msg.Vid, msg.Label, msg.IssuerSubjectKeyID)
	if !isFound {
		return nil, pkitypes.NewErrPkiRevocationDistributionPointDoesNotExists(msg.Vid, msg.Label, msg.IssuerSubjectKeyID)
	}

	// validate and update new values
	if msg.CrlSignerCertificate != "" {
		if err := k.verifyUpdatedCertificate(ctx, msg.CrlSignerCertificate, &pkiRevocationDistributionPoint); err != nil {
			return nil, err
		}
		pkiRevocationDistributionPoint.CrlSignerCertificate = msg.CrlSignerCertificate
	}

	if pkiRevocationDistributionPoint.RevocationType == types.CRLRevocationType && (msg.DataFileSize != 0 || msg.DataDigest != "" || msg.DataDigestType != 0) {
		return nil, pkitypes.NewErrDataFieldPresented(types.CRLRevocationType)
	}

	if msg.DataURL != "" {
		pkiRevocationDistributionPoint.DataURL = msg.DataURL
	}

	if msg.DataFileSize != 0 {
		pkiRevocationDistributionPoint.DataFileSize = msg.DataFileSize
	}

	if msg.DataDigest != "" {
		pkiRevocationDistributionPoint.DataDigest = msg.DataDigest
	}

	if msg.DataDigestType != 0 {
		pkiRevocationDistributionPoint.DataDigestType = msg.DataDigestType
	}

	revocationList, isFound := k.GetPkiRevocationDistributionPointsByIssuerSubjectKeyID(ctx, msg.IssuerSubjectKeyID)
	if isFound {
		for _, revocationPoint := range revocationList.Points {
			if revocationPoint.DataURL == msg.DataURL && revocationPoint.Vid == msg.Vid && revocationPoint.Label != msg.Label {
				return nil, pkitypes.NewErrPkiRevocationDistributionPointWithDataURLAlreadyExists(msg.DataURL, msg.IssuerSubjectKeyID)
			}
		}
	}

	k.SetPkiRevocationDistributionPoint(ctx, pkiRevocationDistributionPoint)
	k.UpdatePkiRevocationDistributionPointBySubjectKeyID(ctx, pkiRevocationDistributionPoint)

	return &types.MsgUpdatePkiRevocationDistributionPointResponse{}, nil
}

func (k msgServer) verifyUpdatedCertificate(ctx sdk.Context, newCertificatePem string, revocationPoint *types.PkiRevocationDistributionPoint) error {
	oldCertificate, err := x509.DecodeX509Certificate(revocationPoint.CrlSignerCertificate)
	if err != nil {
		return pkitypes.NewErrInvalidCertificate(err)
	}

	if oldCertificate.IsSelfSigned() {
		err = k.verifyUpdatedPAA(ctx, newCertificatePem, revocationPoint)
	} else {
		err = k.verifyUpdatedPAI(ctx, newCertificatePem, revocationPoint)
	}

	if err != nil {
		return err
	}

	return nil
}

func (k msgServer) verifyUpdatedPAA(ctx sdk.Context, newCertificatePem string, revocationPoint *types.PkiRevocationDistributionPoint) error {
	// decode new cert
	newCertificate, err := x509.DecodeX509Certificate(newCertificatePem)
	if err != nil {
		return pkitypes.NewErrInvalidCertificate(err)
	}

	// check that it's self-signed
	if !newCertificate.IsSelfSigned() {
		return pkitypes.NewErrRootCertificateIsNotSelfSigned()
	}

	// check that VID is the same
	newCertificateVid, err := x509.GetVidFromSubject(newCertificate.SubjectAsText)
	if err != nil {
		return pkitypes.NewErrInvalidVidFormat(err)
	}
	if newCertificateVid != 0 && newCertificateVid != revocationPoint.Vid {
		return pkitypes.NewErrCRLSignerCertificateVidNotEqualRevocationPointVid(newCertificateVid, revocationPoint.Vid)
	}

	// find the cert on the ledger
	approvedCertificates, isFound := k.GetApprovedCertificates(ctx, newCertificate.Subject, newCertificate.SubjectKeyID)
	if !isFound {
		return pkitypes.NewErrRootCertificateDoesNotExist(newCertificate.Subject, newCertificate.SubjectKeyID)
	}

	// check that it has the same PEM value
	var foundRootCert *types.Certificate
	for _, approvedCertificate := range approvedCertificates.Certs {
		if approvedCertificate.PemCert == newCertificatePem {
			foundRootCert = approvedCertificate

			break
		}
	}
	if foundRootCert == nil {
		return pkitypes.NewErrPemValuesNotEqual(newCertificate.Subject, newCertificate.SubjectKeyID)
	}

	// check that new cert has the same VID as in the message if it's non-VID scoped
	// (vid-scoped has been already checked as part of static validation + equality of PEM values)
	ledgerRootVid, err := x509.GetVidFromSubject(foundRootCert.SubjectAsText)
	if err != nil {
		return pkitypes.NewErrInvalidVidFormat(err)
	}
	if ledgerRootVid == 0 && revocationPoint.Vid != foundRootCert.Vid {
		return pkitypes.NewErrMessageVidNotEqualRootCertVid(revocationPoint.Vid, foundRootCert.Vid)
	}
	if ledgerRootVid != 0 && revocationPoint.Vid != ledgerRootVid {
		return pkitypes.NewErrMessageVidNotEqualRootCertVid(revocationPoint.Vid, foundRootCert.Vid)
	}

	return nil
}

func (k msgServer) verifyUpdatedPAI(ctx sdk.Context, newCertificatePem string, revocationPoint *types.PkiRevocationDistributionPoint) error {
	// decode new cert
	newCertificate, err := x509.DecodeX509Certificate(newCertificatePem)
	if err != nil {
		return pkitypes.NewErrInvalidCertificate(err)
	}

	// check that it's not self-signed
	if newCertificate.IsSelfSigned() {
		return pkitypes.NewErrNonRootCertificateSelfSigned()
	}

	// check that VID is the same
	newCertificateVid, err := x509.GetVidFromSubject(newCertificate.SubjectAsText)
	if err != nil {
		return pkitypes.NewErrInvalidVidFormat(err)
	}
	if newCertificateVid != revocationPoint.Vid {
		return pkitypes.NewErrCRLSignerCertificateVidNotEqualRevocationPointVid(revocationPoint.Vid, newCertificateVid)
	}

	// check PID
	newCertificatePid, err := x509.GetPidFromSubject(newCertificate.SubjectAsText)
	if err != nil {
		return pkitypes.NewErrInvalidPidFormat(err)
	}
	if newCertificatePid != 0 && newCertificatePid != revocationPoint.Pid {
		return pkitypes.NewErrCRLSignerCertificatePidNotEqualRevocationPointPid(newCertificatePid, revocationPoint.Pid)
	}
	if newCertificatePid == 0 && newCertificatePid != revocationPoint.Pid {
		return pkitypes.NewErrPidNotFoundInCertificateButProvidedInRevocationPoint()
	}

	// check that it's chained back to a cert on DCL
	if _, err = k.verifyCertificate(ctx, newCertificate); err != nil {
		return pkitypes.NewErrCertNotChainedBack()
	}

	return nil
}
