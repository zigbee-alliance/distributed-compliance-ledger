package keeper

import (
	"context"

	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/x509"
)

func (k msgServer) ProposeAddX509RootCert(goCtx context.Context, msg *types.MsgProposeAddX509RootCert) (*types.MsgProposeAddX509RootCertResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	signerAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid Address: (%s)", err)
	}

	// check if sender has enough rights to propose a x509 root cert
	if !k.dclauthKeeper.HasRole(ctx, signerAddr, types.RootCertificateApprovalRole) {
		return nil, errors.Wrapf(sdkerrors.ErrUnauthorized,
			"MsgProposeAddX509RootCert transaction should be signed by an account with the %s role",
			types.RootCertificateApprovalRole,
		)
	}

	// decode pem certificate
	x509Certificate, err := x509.DecodeX509Certificate(msg.Cert)
	if err != nil {
		return nil, pkitypes.NewErrInvalidCertificate(err)
	}

	// fail if certificate is not self-signed
	if !x509Certificate.IsSelfSigned() {
		return nil, pkitypes.NewErrInappropriateCertificateType(
			"Inappropriate Certificate Type: Passed certificate is not self-signed, " +
				"so it cannot be used as a root certificate.")
	}

	// check if `Proposed` certificate with the same Subject/SubjectKeyID combination already exists
	if k.IsProposedCertificatePresent(ctx, x509Certificate.Subject, x509Certificate.SubjectKeyID) {
		return nil, pkitypes.NewErrProposedCertificateAlreadyExists(x509Certificate.Subject, x509Certificate.SubjectKeyID)
	}

	// check if certificate with Issuer/Serial Number combination already exists
	if k.IsUniqueCertificatePresent(ctx, x509Certificate.Issuer, x509Certificate.SerialNumber) {
		return nil, pkitypes.NewErrCertificateAlreadyExists(x509Certificate.Issuer, x509Certificate.SerialNumber)
	}

	// verify certificate
	_, err = k.verifyCertificate(ctx, x509Certificate)
	if err != nil {
		return nil, err
	}

	// Get list of certificates for Subject / Subject Key Id combination
	existingCertificates, found := k.GetApprovedCertificates(ctx, x509Certificate.Subject, x509Certificate.SubjectKeyID)
	if found && len(existingCertificates.Certs) > 0 {
		existingCertificate := existingCertificates.Certs[0]

		// Issuer and authorityKeyID must be the same as ones of exisiting certificates with the same subject and
		// subjectKeyID. Since new certificate is self-signed, we have to ensure that the exisiting certificates are
		// self-signed too, consequently are root certificates.
		if !existingCertificate.IsRoot {
			return nil, pkitypes.NewErrUnauthorizedCertIssuer(x509Certificate.Subject, x509Certificate.SubjectKeyID)
		}

		// Existing certificate must not be NOC certificate
		if existingCertificate.CertificateType == types.CertificateType_OperationalPKI {
			return nil, pkitypes.NewErrProvidedNotNocCertButExistingNoc(x509Certificate.Subject, x509Certificate.SubjectKeyID)
		}

		// signer must be same as owner of existing certificates
		if msg.Signer != existingCertificate.Owner {
			return nil, pkitypes.NewErrUnauthorizedCertOwner(x509Certificate.Subject, x509Certificate.SubjectKeyID)
		}
	}

	grant := types.Grant{
		Address: signerAddr.String(),
		Time:    msg.Time,
		Info:    msg.Info,
	}

	// create a new proposed certificate with empty approvals list
	proposedCertificate := types.ProposedCertificate{
		Subject:           x509Certificate.Subject,
		SubjectAsText:     x509Certificate.SubjectAsText,
		SubjectKeyId:      x509Certificate.SubjectKeyID,
		PemCert:           msg.Cert,
		SerialNumber:      x509Certificate.SerialNumber,
		Owner:             msg.Signer,
		Approvals:         []*types.Grant{},
		Vid:               msg.Vid,
		CertSchemaVersion: msg.CertSchemaVersion,
	}

	proposedCertificate.Approvals = append(proposedCertificate.Approvals, &grant)

	// store proposed certificate
	k.SetProposedCertificate(ctx, proposedCertificate)

	_, isFound := k.GetRejectedCertificate(ctx, proposedCertificate.Subject, proposedCertificate.SubjectKeyId)
	if isFound {
		k.RemoveRejectedCertificate(ctx, proposedCertificate.Subject, proposedCertificate.SubjectKeyId)
	}

	// register the unique certificate key
	k.SetUniqueX509Certificate(ctx, x509Certificate)

	return &types.MsgProposeAddX509RootCertResponse{}, nil
}
