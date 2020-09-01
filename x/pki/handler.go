package pki

import (
	"fmt"

	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/auth"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/pki/internal/keeper"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/pki/internal/types"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/pki/internal/x509"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewHandler(keeper keeper.Keeper, authKeeper auth.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case types.MsgProposeAddX509RootCert:
			return handleMsgProposeAddX509RootCert(ctx, keeper, authKeeper, msg)
		case types.MsgApproveAddX509RootCert:
			return handleMsgApproveAddX509RootCert(ctx, keeper, authKeeper, msg)
		case types.MsgAddX509Cert:
			return handleMsgAddX509Cert(ctx, keeper, msg)
		case types.MsgRevokeX509Cert:
			return handleMsgRevokeX509Cert(ctx, keeper, msg)
		default:
			errMsg := fmt.Sprintf("unrecognized pki Msg type: %v", msg.Type())
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

// nolint:funlen
func handleMsgProposeAddX509RootCert(ctx sdk.Context, keeper keeper.Keeper, authKeeper auth.Keeper,
	msg types.MsgProposeAddX509RootCert) sdk.Result {
	// decode pem certificate
	x509Certificate, err := x509.DecodeX509Certificate(msg.Cert)
	if err != nil {
		return types.ErrCodeInvalidCertificate(err.Data()).Result()
	}

	// fail if certificate is not self-signed
	if !x509Certificate.IsSelfSigned() {
		return types.ErrInappropriateCertificateType(
			"Inappropriate Certificate Type: Passed certificate is not self-signed, " +
				"so it cannot be used as a root certificate.").Result()
	}

	// check if `Proposed` certificate with the same Subject/SubjectKeyId combination already exists
	if keeper.IsProposedCertificatePresent(ctx, x509Certificate.Subject, x509Certificate.SubjectKeyID) {
		return types.ErrProposedCertificateAlreadyExists(x509Certificate.Subject, x509Certificate.SubjectKeyID).Result()
	}

	// check if certificate with Issuer/Serial Number combination already exists
	if keeper.IsUniqueCertificateKeyPresent(ctx, x509Certificate.Issuer, x509Certificate.SerialNumber) {
		return types.ErrCertificateAlreadyExists(x509Certificate.Issuer, x509Certificate.SerialNumber).Result()
	}

	// verify certificate
	_, _, err = verifyCertificate(ctx, keeper, x509Certificate)
	if err != nil {
		return err.Result()
	}

	// Get list of certificates for Subject / Subject Key Id combination
	existingCertificates := keeper.GetApprovedCertificates(ctx, x509Certificate.Subject, x509Certificate.SubjectKeyID)

	if len(existingCertificates.Items) > 0 {
		// Issuer and authorityKeyID must be the same as ones of exisiting certificates with the same subject and
		// subjectKeyID. Since new certificate is self-signed, we have to ensure that the exisiting certificates are
		// self-signed too, consequently are root certificates.
		if !existingCertificates.Items[0].IsRoot {
			return sdk.ErrUnauthorized(
				fmt.Sprintf("Issuer and authorityKeyID of new certificate with subject=%v and subjectKeyID=%v "+
					"must be the same as ones of existing certificates with the same subject and subjectKeyID",
					x509Certificate.Subject, x509Certificate.SubjectKeyID)).Result()
		}

		// signer must be same as owner of existing certificates
		if !msg.Signer.Equals(existingCertificates.Items[0].Owner) {
			return sdk.ErrUnauthorized(
				fmt.Sprintf("Only owner of existing certificates with subject=%v and subjectKeyID=%v "+
					"can add new certificate with the same subject and subjectKeyID",
					x509Certificate.Subject, x509Certificate.SubjectKeyID)).Result()
		}
	}

	// create new proposed certificate with empty approvals list
	proposedCertificate := types.NewProposedCertificate(
		msg.Cert,
		x509Certificate.Subject,
		x509Certificate.SubjectKeyID,
		x509Certificate.SerialNumber,
		msg.Signer,
	)

	// if signer has `RootCertificateApprovalRole` append approval
	if authKeeper.HasRole(ctx, msg.Signer, types.RootCertificateApprovalRole) {
		proposedCertificate.Approvals = append(proposedCertificate.Approvals, msg.Signer)
	}

	// store proposed certificate
	keeper.SetProposedCertificate(ctx, proposedCertificate)

	// register the unique certificate key
	keeper.SetUniqueCertificateKey(ctx, x509Certificate.Issuer, x509Certificate.SerialNumber)

	return sdk.Result{}
}

func handleMsgApproveAddX509RootCert(ctx sdk.Context, keeper keeper.Keeper, authKeeper auth.Keeper,
	msg types.MsgApproveAddX509RootCert) sdk.Result {
	// check if signer has root certificate approval role
	if !authKeeper.HasRole(ctx, msg.Signer, types.RootCertificateApprovalRole) {
		return sdk.ErrUnauthorized(
			fmt.Sprintf("MsgApproveAddX509RootCert transaction should be signed by "+
				"an account with the \"%s\" role", types.RootCertificateApprovalRole)).Result()
	}

	// check if corresponding proposed certificate exists
	if !keeper.IsProposedCertificatePresent(ctx, msg.Subject, msg.SubjectKeyID) {
		return types.ErrProposedCertificateDoesNotExist(msg.Subject, msg.SubjectKeyID).Result()
	}

	// get proposed certificate
	proposedCertificate := keeper.GetProposedCertificate(ctx, msg.Subject, msg.SubjectKeyID)

	// check if proposed certificate already has approval form signer
	if proposedCertificate.HasApprovalFrom(msg.Signer) {
		return sdk.ErrUnauthorized(
			fmt.Sprintf("Certificate associated with subject=%v and subjectKeyID=%v combination "+
				"already has approval from=%v", msg.Subject, msg.SubjectKeyID, msg.Signer)).Result()
	}

	// append approval
	proposedCertificate.Approvals = append(proposedCertificate.Approvals, msg.Signer)

	// check if proposed certificate has enough approvals
	if len(proposedCertificate.Approvals) == types.RootCertificateApprovals {
		// create approved certificate
		rootCertificate := types.NewRootCertificate(
			proposedCertificate.PemCert,
			proposedCertificate.Subject,
			proposedCertificate.SubjectKeyID,
			proposedCertificate.SerialNumber,
			proposedCertificate.Owner,
		)

		// add approved certificate to stored list of certificates with the same Subject/SubjectKeyId combination
		keeper.AddApprovedCertificate(ctx, rootCertificate)

		// delete proposed certificate
		keeper.DeleteProposedCertificate(ctx, msg.Subject, msg.SubjectKeyID)
	} else {
		// update proposed certificate
		keeper.SetProposedCertificate(ctx, proposedCertificate)
	}

	return sdk.Result{}
}

// nolint:funlen
func handleMsgAddX509Cert(ctx sdk.Context, keeper keeper.Keeper, msg types.MsgAddX509Cert) sdk.Result {
	// decode pem certificate
	x509Certificate, err := x509.DecodeX509Certificate(msg.Cert)
	if err != nil {
		return err.Result()
	}

	// fail if certificate is self-signed
	if x509Certificate.IsSelfSigned() {
		return types.ErrInappropriateCertificateType(
			"Inappropriate Certificate Type: Passed certificate is self-signed, " +
				"so it cannot be added to the system as a non-root certificate. " +
				"To propose adding a root certificate please use `PROPOSE_ADD_X509_ROOT_CERT` transaction.").Result()
	}

	// check if certificate with Issuer/Serial Number combination already exists
	if keeper.IsUniqueCertificateKeyPresent(ctx, x509Certificate.Issuer, x509Certificate.SerialNumber) {
		return types.ErrCertificateAlreadyExists(x509Certificate.Issuer, x509Certificate.SerialNumber).Result()
	}

	// Get list of certificates for Subject / Subject Key Id combination
	certificates := keeper.GetApprovedCertificates(ctx, x509Certificate.Subject, x509Certificate.SubjectKeyID)

	if len(certificates.Items) > 0 {
		// Issuer and authorityKeyID must be the same as ones of exisiting certificates with the same subject and
		// subjectKeyID. Since new certificate is not self-signed, we have to ensure that the exisiting certificates
		// are not self-signed too, consequently are non-root certificates, before to match issuer and authorityKeyID.
		if certificates.Items[0].IsRoot || x509Certificate.Issuer != certificates.Items[0].Issuer ||
			x509Certificate.AuthorityKeyID != certificates.Items[0].AuthorityKeyID {
			return sdk.ErrUnauthorized(
				fmt.Sprintf("Issuer and authorityKeyID of new certificate with subject=%v and subjectKeyID=%v "+
					"must be the same as ones of existing certificates with the same subject and subjectKeyID",
					x509Certificate.Subject, x509Certificate.SubjectKeyID)).Result()
		}

		// signer must be same as owner of existing certificates
		if !msg.Signer.Equals(certificates.Items[0].Owner) {
			return sdk.ErrUnauthorized(
				fmt.Sprintf("Only owner of existing certificates with subject=%v and subjectKeyID=%v "+
					"can add new certificate with the same subject and subjectKeyID",
					x509Certificate.Subject, x509Certificate.SubjectKeyID)).Result()
		}
	}

	// Valid certificate chain must be built for new certificate
	rootCertificateSubject, rootCertificateSubjectKeyID, err := verifyCertificate(ctx, keeper, x509Certificate)
	if err != nil {
		return types.ErrCodeInvalidCertificate(
			fmt.Sprintf("Cannot build valid certificate chain for certificate with subject=%v and subjectKeyID=%v",
				x509Certificate.Subject, x509Certificate.SubjectKeyID)).Result()
	}

	// create new certificate
	certificate := types.NewNonRootCertificate(
		msg.Cert,
		x509Certificate.Subject,
		x509Certificate.SubjectKeyID,
		x509Certificate.SerialNumber,
		x509Certificate.Issuer,
		x509Certificate.AuthorityKeyID,
		rootCertificateSubject,
		rootCertificateSubjectKeyID,
		msg.Signer,
	)

	// append new certificate to list of certificates with the same Subject/SubjectKeyId combination and store updated list
	certificates.Items = append(certificates.Items, certificate)
	keeper.SetApprovedCertificates(ctx, certificate.Subject, certificate.SubjectKeyID, certificates)

	// add the certificate identifier to the issuer's Child Certificates record
	certificateIdentifier := types.NewCertificateIdentifier(certificate.Subject, certificate.SubjectKeyID)
	addChildCertificateEntry(ctx, keeper, certificate.Issuer, certificate.AuthorityKeyID, certificateIdentifier)

	// register the unique certificate key
	keeper.SetUniqueCertificateKey(ctx, x509Certificate.Issuer, x509Certificate.SerialNumber)

	return sdk.Result{}
}

func handleMsgProposeRevokeX509RootCert(ctx sdk.Context, keeper keeper.Keeper, authKeeper auth.Keeper,
	msg types.MsgProposeRevokeX509RootCert) sdk.Result {
	// check if signer has root certificate approval role
	if !authKeeper.HasRole(ctx, msg.Signer, types.RootCertificateApprovalRole) {
		return sdk.ErrUnauthorized(
			fmt.Sprintf("MsgProposeRevokeX509RootCert transaction should be signed by "+
				"an account with the \"%s\" role", types.RootCertificateApprovalRole)).Result()
	}

	// check if corresponding approved certificates exist
	if !keeper.IsApprovedCertificatesPresent(ctx, msg.Subject, msg.SubjectKeyID) {
		return types.ErrCertificateDoesNotExist(msg.Subject, msg.SubjectKeyID).Result()
	}

	// check that proposed certificate revocation does not exist yet
	if keeper.IsProposedCertificateRevocationPresent(ctx, msg.Subject, msg.SubjectKeyID) {
		return types.ErrProposedCertificateRevocationAlreadyExists(msg.Subject, msg.SubjectKeyID).Result()
	}

	// get corresponding approved certificates
	certificates := keeper.GetApprovedCertificates(ctx, msg.Subject, msg.SubjectKeyID)

	// fail if certificates are not root
	if !certificates.Items[0].IsRoot {
		return types.ErrInappropriateCertificateType(
			fmt.Sprintf("Inappropriate Certificate Type: Certificate with subject=%v and subjectKeyID=%v "+
				"is not a root certificate.", msg.Subject, msg.SubjectKeyID)).Result()
	}

	// create new proposed certificate revocation with approval from signer
	revocation := types.NewProposedCertificateRevocation(
		msg.Subject,
		msg.SubjectKeyID,
		msg.Signer,
	)

	// store proposed certificate revocation
	keeper.SetProposedCertificateRevocation(ctx, revocation)

	return sdk.Result{}
}

func handleMsgApproveRevokeX509RootCert(ctx sdk.Context, keeper keeper.Keeper, authKeeper auth.Keeper,
	msg types.MsgApproveRevokeX509RootCert) sdk.Result {
	// check if signer has root certificate approval role
	if !authKeeper.HasRole(ctx, msg.Signer, types.RootCertificateApprovalRole) {
		return sdk.ErrUnauthorized(
			fmt.Sprintf("MsgApproveRevokeX509RootCert transaction should be signed by "+
				"an account with the \"%s\" role", types.RootCertificateApprovalRole)).Result()
	}

	// check if corresponding proposed certificate revocation exists
	if !keeper.IsProposedCertificateRevocationPresent(ctx, msg.Subject, msg.SubjectKeyID) {
		return types.ErrProposedCertificateRevocationDoesNotExist(msg.Subject, msg.SubjectKeyID).Result()
	}

	// get proposed certificate revocation
	revocation := keeper.GetProposedCertificateRevocation(ctx, msg.Subject, msg.SubjectKeyID)

	// check if proposed certificate revocation already has approval form signer
	if revocation.HasApprovalFrom(msg.Signer) {
		return sdk.ErrUnauthorized(
			fmt.Sprintf("Certificate revocation associated with subject=%v and subjectKeyID=%v combination "+
				"already has approval from=%v", msg.Subject, msg.SubjectKeyID, msg.Signer)).Result()
	}

	// append approval
	revocation.Approvals = append(revocation.Approvals, msg.Signer)

	// check if proposed certificate revocation has enough approvals
	if len(revocation.Approvals) == types.RootCertificateApprovals {
		certificates := keeper.GetApprovedCertificates(ctx, msg.Subject, msg.SubjectKeyID)

		keeper.AddRevokedCertificates(ctx, msg.Subject, msg.SubjectKeyID, certificates)
		keeper.DeleteApprovedCertificates(ctx, msg.Subject, msg.SubjectKeyID)

		revokeChildCertificates(ctx, keeper, msg.Subject, msg.SubjectKeyID)
	} else {
		keeper.SetProposedCertificateRevocation(ctx, revocation)
	}

	return sdk.Result{}
}

func handleMsgRevokeX509Cert(ctx sdk.Context, keeper keeper.Keeper, msg types.MsgRevokeX509Cert) sdk.Result {
	if !keeper.IsApprovedCertificatesPresent(ctx, msg.Subject, msg.SubjectKeyID) {
		return types.ErrProposedCertificateDoesNotExist(msg.Subject, msg.SubjectKeyID).Result()
	}

	certificates := keeper.GetApprovedCertificates(ctx, msg.Subject, msg.SubjectKeyID)

	if certificates.Items[0].IsRoot {
		return types.ErrInappropriateCertificateType(
			fmt.Sprintf("Inappropriate Certificate Type: Certificate with subject=%v and subjectKeyID=%v "+
				"is a root certificate. To propose revocation of a root certificate please use "+
				"`PROPOSE_REVOKE_X509_ROOT_CERT` transaction.", msg.Subject, msg.SubjectKeyID)).Result()
	}

	if !msg.Signer.Equals(certificates.Items[0].Owner) {
		return sdk.ErrUnauthorized("Only owner can revoke certificate using `REVOKE_X509_CERT`").Result()
	}

	issuer := certificates.Items[0].Issuer
	authorityKeyID := certificates.Items[0].AuthorityKeyID

	// Revoke certificates with given subject/subjectKeyID
	keeper.AddRevokedCertificates(ctx, msg.Subject, msg.SubjectKeyID, certificates)
	keeper.DeleteApprovedCertificates(ctx, msg.Subject, msg.SubjectKeyID)

	// Remove certificate identifier from issuer's ChildCertificates record
	certIdentifier := types.NewCertificateIdentifier(msg.Subject, msg.SubjectKeyID)
	removeChildCertificateEntry(ctx, keeper, issuer, authorityKeyID, certIdentifier)

	revokeChildCertificates(ctx, keeper, msg.Subject, msg.SubjectKeyID)

	return sdk.Result{}
}

func revokeChildCertificates(ctx sdk.Context, keeper keeper.Keeper, issuer string, authorityKeyID string) {
	// Get issuer's ChildCertificates record
	childCertificates := keeper.GetChildCertificates(ctx, issuer, authorityKeyID)

	// For each child certificate subject/subjectKeyID combination
	for _, certIdentifier := range childCertificates.CertIdentifiers {
		// Revoke certificates with this subject/subjectKeyID combination
		certificates := keeper.GetApprovedCertificates(ctx, certIdentifier.Subject, certIdentifier.SubjectKeyID)
		keeper.AddRevokedCertificates(ctx, certIdentifier.Subject, certIdentifier.SubjectKeyID, certificates)
		keeper.DeleteApprovedCertificates(ctx, certIdentifier.Subject, certIdentifier.SubjectKeyID)

		// Process child certificates recursively
		revokeChildCertificates(ctx, keeper, certIdentifier.Subject, certIdentifier.SubjectKeyID)
	}

	// Delete entire ChildCertificates record of issuer
	keeper.DeleteChildCertificates(ctx, issuer, authorityKeyID)
}

func addChildCertificateEntry(ctx sdk.Context, keeper keeper.Keeper, issuer string, authorityKeyID string,
	certIdentifier types.CertificateIdentifier) {
	childCertificates := keeper.GetChildCertificates(ctx, issuer, authorityKeyID)

	for _, existingIdentifier := range childCertificates.CertIdentifiers {
		if existingIdentifier == certIdentifier {
			return
		}
	}

	childCertificates.CertIdentifiers = append(childCertificates.CertIdentifiers, certIdentifier)
	keeper.SetChildCertificates(ctx, childCertificates)
}

func removeChildCertificateEntry(ctx sdk.Context, keeper keeper.Keeper, issuer string, authorityKeyID string,
	certIdentifier types.CertificateIdentifier) {
	childCertificates := keeper.GetChildCertificates(ctx, issuer, authorityKeyID)

	certIDIndex := -1

	for i, existingIdentifier := range childCertificates.CertIdentifiers {
		if existingIdentifier == certIdentifier {
			certIDIndex = i

			break
		}
	}

	if certIDIndex == -1 {
		return
	}

	childCertificates.CertIdentifiers =
		append(childCertificates.CertIdentifiers[:certIDIndex], childCertificates.CertIdentifiers[certIDIndex+1:]...)

	if len(childCertificates.CertIdentifiers) > 0 {
		keeper.SetChildCertificates(ctx, childCertificates)
	} else {
		keeper.DeleteChildCertificates(ctx, issuer, authorityKeyID)
	}
}

// Tries to build a valid certificate chain for the given certificate.
// Returns the RootSubject/RootSubjectKeyID combination or an error in case no valid certificate chain can be built.
func verifyCertificate(ctx sdk.Context, keeper keeper.Keeper,
	x509Certificate *x509.X509Certificate) (string, string, sdk.Error) {
	if x509Certificate.IsSelfSigned() {
		// in this system a certificate is self-signed if and only if it is a root certificate
		if err := x509Certificate.Verify(x509Certificate); err == nil {
			return x509Certificate.Subject, x509Certificate.SubjectKeyID, nil
		}
	} else {
		parentCertificates := keeper.GetApprovedCertificates(ctx, x509Certificate.Issuer, x509Certificate.AuthorityKeyID)

		for _, cert := range parentCertificates.Items {
			parentX509Certificate, err := x509.DecodeX509Certificate(cert.PemCert)
			if err != nil {
				continue
			}

			// verify certificate against parent
			if err := x509Certificate.Verify(parentX509Certificate); err != nil {
				continue
			}

			// verify parent certificate
			if subject, subjectKeyID, err := verifyCertificate(ctx, keeper, parentX509Certificate); err == nil {
				return subject, subjectKeyID, nil
			}
		}
	}

	return "", "", types.ErrCodeInvalidCertificate(
		fmt.Sprintf("Certificate verification failed for certificate with subject=%v and subjectKeyID=%v",
			x509Certificate.Subject, x509Certificate.SubjectKeyID))
}
