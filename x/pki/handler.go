package pki

// nolint:goimports
import (
	"fmt"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/auth"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/pki/internal/keeper"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/pki/internal/types"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/pki/internal/x509"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewHandler(keeper keeper.Keeper, authzKeeper auth.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case types.MsgProposeAddX509RootCert:
			return handleMsgProposeAddX509RootCert(ctx, keeper, authzKeeper, msg)
		case types.MsgApproveAddX509RootCert:
			return handleMsgApproveAddX509RootCert(ctx, keeper, authzKeeper, msg)
		case types.MsgAddX509Cert:
			return handleMsgAddX509Cert(ctx, keeper, msg)
		default:
			errMsg := fmt.Sprintf("unrecognized compliancetest Msg type: %v", msg.Type())
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

// nolint:funlen
func handleMsgProposeAddX509RootCert(ctx sdk.Context, keeper keeper.Keeper, authzKeeper auth.Keeper,
	msg types.MsgProposeAddX509RootCert) sdk.Result {
	// decode pem certificate
	certificate, err := x509.DecodeX509Certificate(msg.Cert)
	if err != nil {
		return types.ErrCodeInvalidCertificate(err.Data()).Result()
	}

	// verify certificate against self
	if err := certificate.VerifyX509Certificate(certificate.Certificate); err != nil {
		return types.ErrCodeInvalidCertificate(err.Data()).Result()
	}

	// check if certificate is `root`
	if !certificate.IsRootCertificate() {
		return types.ErrInappropriateCertificateType(
			"Inappropriate Certificate Type: Passed certificate is not a root certificate so " +
				"it cannot be used for root certificates proposing.").Result()
	}

	// check if `Proposed` certificate with the same Subject/SubjectKeyId combination already exists
	if keeper.IsProposedCertificatePresent(ctx, certificate.Subject, certificate.SubjectKeyID) {
		return types.ErrCertificateAlreadyExists(certificate.Subject,
			certificate.SubjectKeyID, certificate.SerialNumber).Result()
	}

	// check if certificate with Issuer/Serial Number combination already exists
	if keeper.IsCertificateExists(ctx, certificate.Issuer, certificate.SerialNumber) {
		return sdk.NewError(types.Codespace, types.CodeCertificateAlreadyExists,
			fmt.Sprintf("X509 certificate with the combination of issuer=%v, "+
				"serialNumber=%v already exists on the ledger", certificate.Issuer, certificate.SerialNumber)).Result()
	}

	// Get list of certificates for Subject / Subject Key Id combination
	existingCertificates := keeper.GetCertificates(ctx, certificate.Subject, certificate.SubjectKeyID)

	if len(existingCertificates.Items) > 0 {
		// signer must be same as owner of existing certificates
		if !msg.Signer.Equals(existingCertificates.Items[0].Owner) {
			return sdk.ErrUnauthorized(
				fmt.Sprintf("Only owner can append next certificate corresponding to the same subject=%v and "+
					"subjectKeyId=%v combination", certificate.Subject, certificate.SubjectKeyID)).Result()
		}
	}

	// create new proposed certificate with empty approvals list
	pendingCertificate := types.NewProposedCertificate(
		msg.Cert,
		certificate.Subject,
		certificate.SubjectKeyID,
		certificate.SerialNumber,
		msg.Signer,
	)

	// if signer has `RootCertificateApprovalRole` append approval
	if authzKeeper.HasRole(ctx, msg.Signer, types.RootCertificateApprovalRole) {
		pendingCertificate.Approvals = append(pendingCertificate.Approvals, msg.Signer)
	}

	// store proposed certificate
	keeper.SetProposedCertificate(ctx, pendingCertificate)

	// set certificate existence flag
	keeper.AddCertificateExistenceFlag(ctx, certificate.Issuer, certificate.SerialNumber)

	return sdk.Result{}
}

func handleMsgApproveAddX509RootCert(ctx sdk.Context, keeper keeper.Keeper, authzKeeper auth.Keeper,
	msg types.MsgApproveAddX509RootCert) sdk.Result {
	// check if corresponding proposed certificate exists
	if !keeper.IsProposedCertificatePresent(ctx, msg.Subject, msg.SubjectKeyID) {
		return types.ErrProposedCertificateDoesNotExist(msg.Subject, msg.SubjectKeyID).Result()
	}

	// check if signer has approval role
	if !authzKeeper.HasRole(ctx, msg.Signer, types.RootCertificateApprovalRole) {
		return sdk.ErrUnauthorized(
			fmt.Sprintf("MsgApproveAddX509RootCert transaction should be signed by "+
				"an account with the \"%s\" role", types.RootCertificateApprovalRole)).Result()
	}

	// get proposed certificate
	pendingCertificate := keeper.GetProposedCertificate(ctx, msg.Subject, msg.SubjectKeyID)

	// check if certificate already has approval form signer
	if pendingCertificate.HasApprovalFrom(msg.Signer) {
		return sdk.ErrUnauthorized(
			fmt.Sprintf("Certificate associated with the subject=%v and subjectKeyId=%v "+
				"combination already has approval from=%v", msg.Subject, msg.SubjectKeyID, msg.Signer)).Result()
	}

	// append approval
	pendingCertificate.Approvals = append(pendingCertificate.Approvals, msg.Signer)

	// check if certificate has enough approvals
	if len(pendingCertificate.Approvals) == types.RootCertificateApprovals {
		// create approved certificate
		rootCertificate := types.NewRootCertificate(
			pendingCertificate.PemCert,
			pendingCertificate.Subject,
			pendingCertificate.SubjectKeyID,
			pendingCertificate.SerialNumber,
			pendingCertificate.Owner,
		)

		// get certificate referring to the same Subject / Subject Key Id combination
		certificates := keeper.GetCertificates(ctx, rootCertificate.Subject, rootCertificate.SubjectKeyID)

		// append new
		certificates.Items = append(certificates.Items, rootCertificate)

		// store updated certificates and delete proposed
		keeper.SetCertificates(ctx, rootCertificate.Subject, rootCertificate.SubjectKeyID, certificates)
		keeper.DeleteProposedCertificate(ctx, msg.Subject, msg.SubjectKeyID)
	} else {
		// update proposed certificate
		keeper.SetProposedCertificate(ctx, pendingCertificate)
	}

	return sdk.Result{}
}

// nolint:funlen
func handleMsgAddX509Cert(ctx sdk.Context, keeper keeper.Keeper, msg types.MsgAddX509Cert) sdk.Result {
	// decode pem certificate
	certificate, err := x509.DecodeX509Certificate(msg.Cert)
	if err != nil {
		return err.Result()
	}

	// check if certificate is NOT root
	if certificate.IsRootCertificate() {
		return types.ErrInappropriateCertificateType(
			"Inappropriate Certificate Type: Passed certificate is a root certificate. " +
				"Please use `PROPOSE_ADD_X509_ROOT_CERT` to propose a root certificate").Result()
	}

	// check if certificate with Issuer/Serial Number combination already exists
	if keeper.IsCertificateExists(ctx, certificate.Issuer, certificate.SerialNumber) {
		return sdk.NewError(types.Codespace, types.CodeCertificateAlreadyExists,
			fmt.Sprintf("X509 certificate with the combination of issuer=%v, "+
				"serialNumber=%v already exists on the ledger", certificate.Issuer, certificate.SerialNumber)).Result()
	}

	// Get list of certificates for Subject / Subject Key Id combination
	existingCertificates := keeper.GetCertificates(ctx, certificate.Subject, certificate.SubjectKeyID)

	if len(existingCertificates.Items) > 0 {
		// signer must be same as owner of existing certificates
		if !msg.Signer.Equals(existingCertificates.Items[0].Owner) {
			return sdk.ErrUnauthorized(
				fmt.Sprintf("Only owner can append next certificate corresponding to the same "+
					"subject=%v and subjectKeyId=%v combination", certificate.Subject, certificate.SubjectKeyID)).Result()
		}
	}

	// valid certificate chain can be build for new certificate
	rootCertificateSubject, rootCertificateSubjectKeyID, err := VerifyCertificate(ctx, keeper, certificate)
	if err != nil {
		return types.ErrCodeInvalidCertificate(
			fmt.Sprintf("Cannot build valid chain to root for certificate with subject=%v and SubjectKeyID=%v. "+
				"Error: %v", certificate.Subject, certificate.SubjectKeyID, err.Data())).Result()
	}

	// create new certificate
	leafCertificate := types.NewIntermediateCertificate(
		msg.Cert,
		certificate.Subject,
		certificate.SubjectKeyID,
		certificate.SerialNumber,
		rootCertificateSubject,
		rootCertificateSubjectKeyID,
		msg.Signer,
	)

	// append new certificate to existing
	existingCertificates.Items = append(existingCertificates.Items, leafCertificate)
	keeper.SetCertificates(ctx, leafCertificate.Subject, leafCertificate.SubjectKeyID, existingCertificates)

	// append to parent certificate reference on child
	keeper.AddChildCertificate(ctx, certificate.Issuer, certificate.AuthorityKeyID, leafCertificate)

	// set certificate existence flag
	keeper.AddCertificateExistenceFlag(ctx, certificate.Issuer, certificate.SerialNumber)

	return sdk.Result{}
}

func VerifyCertificate(ctx sdk.Context, keeper keeper.Keeper,
	certificate *x509.X509Certificate) (string, string, sdk.Error) {
	// exit if root found
	if certificate.IsRootCertificate() {
		return certificate.Subject, certificate.SubjectKeyID, nil
	}

	// check parent certificates exists
	if !keeper.IsCertificatePresent(ctx, certificate.Issuer, certificate.AuthorityKeyID) {
		return "", "", sdk.NewError(types.Codespace, types.CodeCertificateDoesNotExist,
			fmt.Sprintf("No parent X509 certificate associated with the issuer=%v and "+
				"authorityKeyId=%v on the ledger", certificate.Issuer, certificate.AuthorityKeyID))
	}

	parentCertificates := keeper.GetCertificates(ctx, certificate.Issuer, certificate.AuthorityKeyID)

	for _, cert := range parentCertificates.Items {
		parentX509Certificate, err := x509.DecodeX509Certificate(cert.PemCert)
		if err != nil {
			continue
		}

		// verify certificate against parent
		if err := certificate.VerifyX509Certificate(parentX509Certificate.Certificate); err != nil {
			continue
		}

		// verify parent certificate. exit if root is find
		if subject, subjectKeyID, err := VerifyCertificate(ctx, keeper, parentX509Certificate); err == nil {
			return subject, subjectKeyID, nil
		}
	}

	return "", "", types.ErrCodeInvalidCertificate(
		fmt.Sprintf("Certificate verification failed for certificate "+
			"with sibject=%v and subjectKeyId=%v", certificate.SubjectKeyID, certificate.SubjectKeyID))
}
