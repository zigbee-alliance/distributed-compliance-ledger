package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
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
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid Address: (%s)", err)
	}
	signerAccount, _ := k.dclauthKeeper.GetAccountO(ctx, signerAddr)
	if !k.dclauthKeeper.HasRole(ctx, signerAddr, dclauthtypes.Vendor) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized,
			"MsgUpdatePkiRevocationDistributionPoint transaction should be signed by an account with the \"%s\" role",
			dclauthtypes.Vendor,
		)
	}

	// check that Revocation Point exists
	pkiRevocationDistributionPoint, isFound := k.GetPkiRevocationDistributionPoint(ctx, msg.Vid, msg.Label, msg.IssuerSubjectKeyID)
	if !isFound {
		return nil, pkitypes.NewErrPkiRevocationDistributionPointDoesNotExists("PKI revocation distribution point does not exist")
	}

	// check that Vendor has the same VID as the Revocation Point
	if pkiRevocationDistributionPoint.Vid != signerAccount.VendorID {
		return nil, sdkerrors.Wrap(pkitypes.ErrCRLSignerCertificateVidNotEqualAccountVid,
			"MsgUpdatePkiRevocationDistributionPoint signer must have the same vid as provided in an existing certificate from the revocation point",
		)
	}

	// validate and update new values
	if msg.CrlSignerCertificate != "" {
		if err := k.verifyUpdatedCertificate(ctx, msg.CrlSignerCertificate, &pkiRevocationDistributionPoint); err != nil {
			return nil, err
		}
		pkiRevocationDistributionPoint.CrlSignerCertificate = msg.CrlSignerCertificate
	}

	if pkiRevocationDistributionPoint.RevocationType == types.CRLRevocationType && (msg.DataFileSize != 0 || msg.DataDigest != "" || msg.DataDigestType != 0) {
		return nil, pkitypes.NewErrDataFieldPresented(fmt.Sprintf("Data Digest, Data File Size and Data Digest Type must be omitted for Revocation Type %d", types.CRLRevocationType))
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
		return pkitypes.NewErrRootCertificateIsNotSelfSigned("Updated CRL signer certificate must be self-signed since old one was self-signed")
	}

	// check that VID is the same
	newVid, err := x509.GetVidFromSubject(newCertificate.SubjectAsText)
	if err != nil {
		return sdkerrors.Wrapf(pkitypes.ErrInvalidVidFormat, "Could not parse vid: %s", err)
	}
	if newVid != 0 && newVid != revocationPoint.Vid {
		return pkitypes.NewErrCRLSignerCertificateVidNotEqualMsgVid("CRL Signer Certificate's vid must be equal to the provided vid in the message")
	}

	// find the cert on the ledger
	approvedCertificates, isFound := k.GetApprovedCertificates(ctx, newCertificate.Subject, newCertificate.SubjectKeyID)
	if !isFound {
		return sdkerrors.Wrap(pkitypes.NewErrCertificateDoesNotExist(newCertificate.Subject, newCertificate.SubjectKeyID), "CRL signer Certificate must be a root certificate present on the ledger if isPAA = True")
	}

	// check that it has the same PEM value
	var foundRootCert *types.Certificate = nil
	for _, approvedCertificate := range approvedCertificates.Certs {
		if approvedCertificate.PemCert == newCertificatePem {
			foundRootCert = approvedCertificate
			break
		}
	}
	if foundRootCert == nil {
		return pkitypes.NewErrPemValuesNotEqual("PEM values of the CRL signer certificate and a certificate found by its Subject and SubjectKeyID are not equal")
	}

	// check that new cert has the same VID as in the message if it's non-VID scoped
	// (vid-scoped has been already checked as patr of static validation + equality of PEM values
	ledgerRootVid, err := x509.GetVidFromSubject(foundRootCert.SubjectAsText)
	if err != nil {
		return sdkerrors.Wrapf(pkitypes.ErrInvalidVidFormat, "Could not parse vid: %s", err)
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
		return pkitypes.NewErrNonRootCertificateSelfSigned("Updated CRL signer certificate must not be self-signed since old one was not self-signed")
	}

	// check that VID is the same
	newVid, err := x509.GetVidFromSubject(newCertificate.SubjectAsText)
	if err != nil {
		return sdkerrors.Wrapf(pkitypes.ErrInvalidVidFormat, "Could not parse vid: %s", err)
	}
	if newVid != revocationPoint.Vid {
		return pkitypes.NewErrCRLSignerCertificateVidNotEqualRevocationPointVid(revocationPoint.Vid, newVid)
	}

	// check PID
	newPid, err := x509.GetPidFromSubject(newCertificate.SubjectAsText)
	if err != nil {
		return sdkerrors.Wrapf(pkitypes.ErrInvalidPidFormat, "Could not parse pid: %s", err)
	}
	if newPid != 0 && newPid != revocationPoint.Pid {
		return pkitypes.NewErrCRLSignerCertificatePidNotEqualMsgPid("pid in updated CRL Signer Certificate must be equal to pid in revocation point")
	}
	if newPid == 0 && newPid != revocationPoint.Pid {
		return pkitypes.NewErrPidNotFound("pid not found in updated CRL Signer Certificate when it is provided in revocation point")
	}

	// check that it's chained back to a cert on DCL
	if _, _, err := k.verifyCertificate(ctx, newCertificate); err != nil {
		return pkitypes.NewErrCertNotChainedBack()
	}

	return nil
}
