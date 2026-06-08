package keeper

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	x509std "crypto/x509"
	"encoding/asn1"

	sdk "github.com/cosmos/cosmos-sdk/types"

	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	dclauthtypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/x509"
)

func (k msgServer) AddPkiRevocationDistributionPoint(goCtx context.Context, msg *types.MsgAddPkiRevocationDistributionPoint) (*types.MsgAddPkiRevocationDistributionPointResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// check if signer has vendor role
	signerAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return nil, pkitypes.NewErrInvalidAddress(err)
	}
	signerAccount, _ := k.dclauthKeeper.GetAccountO(ctx, signerAddr)
	if !k.dclauthKeeper.HasRole(ctx, signerAddr, dclauthtypes.Vendor) {
		return nil, pkitypes.NewErrUnauthorizedRole("MsgAddPkiRevocationDistributionPoint", dclauthtypes.Vendor)
	}

	// decode CrlSignerCertificate
	crlSignerCertificate, err := x509.ParseAndValidateCertificate(msg.CrlSignerCertificate)
	if err != nil {
		return nil, pkitypes.NewErrInvalidCertificate(err)
	}

	// verify CrlSignerCertificate
	err = VerifyCrlSignerCertificate(crlSignerCertificate, msg)
	if err != nil {
		return nil, err
	}

	// compare VID in message and Vendor account
	if msg.Vid != signerAccount.VendorID {
		return nil, pkitypes.NewErrMessageVidNotEqualAccountVid(msg.Vid, signerAccount.VendorID)
	}

	// check that distribution point doesn't exist yet
	_, isFound := k.GetPkiRevocationDistributionPoint(ctx, msg.Vid, msg.Label, msg.IssuerSubjectKeyID)
	if isFound {
		return nil, pkitypes.NewErrPkiRevocationDistributionPointWithVidAndLabelAlreadyExists(msg.Vid, msg.Label, msg.IssuerSubjectKeyID)
	}

	if crlSignerCertificate.IsSelfSigned() {
		// check that crlSignerCertificate cert is present on the ledger and has the same VID
		err = k.checkRootCert(ctx, crlSignerCertificate, msg)
	} else {
		// check that crlSignerCertificate is chained back to a certificate on the ledger
		err = k.checkCRLSignerNonRootCert(ctx, crlSignerCertificate, msg.CrlSignerDelegator, msg.IsPAA)
	}
	if err != nil {
		return nil, err
	}

	revocationList, isFound := k.GetPkiRevocationDistributionPointsByIssuerSubjectKeyID(ctx, msg.IssuerSubjectKeyID)
	if isFound {
		for _, revocationPoint := range revocationList.Points {
			if revocationPoint.DataURL == msg.DataURL && revocationPoint.Vid == msg.Vid {
				return nil, pkitypes.NewErrPkiRevocationDistributionPointWithDataURLAlreadyExists(msg.DataURL, msg.IssuerSubjectKeyID)
			}
		}
	}

	// add to state
	pkiRevocationDistributionPoint := types.PkiRevocationDistributionPoint{
		Vid:                  msg.Vid,
		Label:                msg.Label,
		IssuerSubjectKeyID:   msg.IssuerSubjectKeyID,
		Pid:                  msg.Pid,
		IsPAA:                msg.IsPAA,
		CrlSignerCertificate: msg.CrlSignerCertificate,
		CrlSignerDelegator:   msg.CrlSignerDelegator,
		DataURL:              msg.DataURL,
		DataFileSize:         msg.DataFileSize,
		DataDigest:           msg.DataDigest,
		DataDigestType:       msg.DataDigestType,
		RevocationType:       msg.RevocationType,
		SchemaVersion:        msg.SchemaVersion,
	}

	k.SetPkiRevocationDistributionPoint(ctx, pkiRevocationDistributionPoint)
	k.AddPkiRevocationDistributionPointBySubjectKeyID(ctx, pkiRevocationDistributionPoint)

	return &types.MsgAddPkiRevocationDistributionPointResponse{}, nil
}

func (k msgServer) checkRootCert(ctx sdk.Context, crlSignerCertificate *x509.Certificate, msg *types.MsgAddPkiRevocationDistributionPoint) error {
	// find the cert on the ledger
	approvedCertificates, isFound := k.GetApprovedCertificates(ctx, crlSignerCertificate.Subject, crlSignerCertificate.SubjectKeyID)
	if !isFound {
		return pkitypes.NewErrRootCertificateDoesNotExist(crlSignerCertificate.Subject, crlSignerCertificate.SubjectKeyID)
	}

	// check that it has the same PEM value
	var foundRootCert *types.Certificate
	for _, approvedCertificate := range approvedCertificates.Certs {
		if x509.RemoveWhitespaces(approvedCertificate.PemCert) == x509.RemoveWhitespaces(msg.CrlSignerCertificate) {
			foundRootCert = approvedCertificate

			break
		}
	}
	if foundRootCert == nil {
		return pkitypes.NewErrPemValuesNotEqual(crlSignerCertificate.Subject, crlSignerCertificate.SubjectKeyID)
	}

	// check that root cert has the same VID as in the message if it's non-VID scoped
	// (vid-scoped has been already checked as patr of static validation + equality of PEM values
	ledgerRootVid, err := x509.GetVidFromSubject(foundRootCert.SubjectAsText)
	if err != nil {
		return pkitypes.NewErrInvalidVidFormat(err)
	}
	if ledgerRootVid == 0 && msg.Vid != foundRootCert.Vid {
		return pkitypes.NewErrMessageVidNotEqualRootCertVid(msg.Vid, foundRootCert.Vid)
	}

	return nil
}

func (k msgServer) checkCRLSignerNonRootCert(ctx sdk.Context, crlSignerCertificate *x509.Certificate, crlSignerDelegator string, isPAA bool) error {
	if crlSignerDelegator != "" && !isPAA {
		crlSignerDelegatorCert, err := x509.ParseAndValidateCertificate(crlSignerDelegator)
		if err != nil {
			return pkitypes.NewErrInvalidCertificate(err)
		}

		// verify CRL Signer certificate against Delegated PAI certificate
		if err = crlSignerCertificate.Verify(crlSignerDelegatorCert, ctx.BlockTime()); err != nil {
			return pkitypes.NewErrCRLSignerCertNotChainedBackToDelegator()
		}

		if _, err = k.verifyCertificate(ctx, crlSignerDelegatorCert); err != nil {
			return pkitypes.NewErrCRLSignerCertDelegatorNotChainedBack()
		}

		return nil
	}

	// check that it's chained back to a cert on DCL
	_, err := k.verifyCertificate(ctx, crlSignerCertificate)
	if err != nil {
		return pkitypes.NewErrCertNotChainedBack()
	}

	return nil
}

func VerifyCrlSignerCertificate(cert *x509.Certificate, msg *types.MsgAddPkiRevocationDistributionPoint) error {
	if msg.IsPAA {
		return verifyPAA(cert, msg)
	}

	return verifyPAI(cert, msg)
}

func verifyPAA(cert *x509.Certificate, msg *types.MsgAddPkiRevocationDistributionPoint) error {
	if msg.Pid != 0 {
		return pkitypes.NewErrNotEmptyPidForRootCertificate()
	}

	pid, _ := x509.GetPidFromSubject(cert.SubjectAsText)
	if pid != 0 {
		return pkitypes.NewErrNotEmptyPidForRootCertificate()
	}

	// verify VID
	vid, err := x509.GetVidFromSubject(cert.SubjectAsText)
	if err != nil {
		return pkitypes.NewErrInvalidVidFormat(err)
	}

	if vid > 0 && vid != msg.Vid {
		return pkitypes.NewErrCRLSignerCertificateVidNotEqualMsgVid(vid, msg.Vid)
	}

	if !cert.IsSelfSigned() {
		err = VerifyCRLSignerCertFormat(cert)
		if err != nil {
			return err
		}
	}

	return nil
}

var oidKeyUsage = asn1.ObjectIdentifier{2, 5, 29, 15}

func verifyPAI(cert *x509.Certificate, msg *types.MsgAddPkiRevocationDistributionPoint) error {
	if cert.IsSelfSigned() {
		return pkitypes.NewErrNonRootCertificateSelfSigned()
	}

	// verify VID
	vid, err := x509.GetVidFromSubject(cert.SubjectAsText)
	if err != nil {
		return pkitypes.NewErrInvalidVidFormat(err)
	}

	if vid > 0 && vid != msg.Vid {
		return pkitypes.NewErrCRLSignerCertificateVidNotEqualMsgVid(vid, msg.Vid)
	}

	// verify PID
	pid, err := x509.GetPidFromSubject(cert.SubjectAsText)
	if err != nil {
		return pkitypes.NewErrInvalidPidFormat(err)
	}
	if pid == 0 && msg.Pid != 0 {
		return pkitypes.NewErrNotEmptyPidForNonRootCertificate()
	}
	if pid != 0 && msg.Pid == 0 {
		return pkitypes.NewErrPidNotFoundInMessage(pid)
	}
	if pid > 0 && pid != msg.Pid {
		return pkitypes.NewErrCRLSignerCertificatePidNotEqualMsgPid(pid, msg.Pid)
	}

	if msg.CrlSignerDelegator != "" {
		err = VerifyCRLSignerCertFormat(cert)
		if err != nil {
			return err
		}
	}

	return nil
}

func VerifyCRLSignerCertFormat(certificate *x509.Certificate) error {
	if certificate.SubjectKeyID == "" {
		return pkitypes.NewErrWrongSubjectKeyIDFormat()
	}

	cert := certificate.Certificate
	if cert.Version != 3 {
		return pkitypes.NewErrCRLSignerCertificateInvalidVersion(
			"The version field SHALL be set to 2 to indicate v3 certificate",
		)
	}

	if cert.SignatureAlgorithm != x509std.ECDSAWithSHA256 {
		return pkitypes.NewErrCRLSignerCertificateInvalidFormat(
			"The signature field SHALL contain the identifier for signatureAlgorithm ecdsa-with-SHA256",
		)
	}

	// Type assert to get the ECDSA public key
	ecdsaPubKey, ok := cert.PublicKey.(*ecdsa.PublicKey)
	if !ok {
		return pkitypes.NewErrCRLSignerCertificateInvalidFormat(
			"Public key is not of type ECDSA",
		)
	}

	// Check if the curve parameters match prime256v1 (secp256r1 / P-256)
	if ecdsaPubKey.Curve != elliptic.P256() {
		return pkitypes.NewErrCRLSignerCertificateInvalidFormat(
			"The public key must use prime256v1 curve",
		)
	}

	// Basic Constraint extension should be marked critical and have the cA field set to false
	if !cert.BasicConstraintsValid || cert.IsCA {
		return pkitypes.NewErrCRLSignerCertificateInvalidFormat(
			"Basic Constraint extension's cA field SHALL be set to FALSE",
		)
	}

	// Basic Constraint extension should be marked critical
	isCritical := false
	for _, ext := range cert.Extensions {
		if ext.Id.Equal(oidKeyUsage) {
			isCritical = ext.Critical

			break
		}
	}

	if !isCritical {
		return pkitypes.NewErrCRLSignerCertificateInvalidFormat("Basic Constraint extension SHALL be marked critical")
	}

	if cert.KeyUsage&x509std.KeyUsageCRLSign == 0 {
		return pkitypes.NewErrCRLSignerCertificateInvalidFormat("The cRLSign bits SHALL be set in the KeyUsage bitstring")
	}

	if cert.KeyUsage&^(x509std.KeyUsageCRLSign|x509std.KeyUsageDigitalSignature) != 0 {
		return pkitypes.NewErrCRLSignerCertificateInvalidFormat("The KeyUsage bitstring can only include the cRLSign and digitalSignature bits")
	}

	return nil
}
