package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/x509"
)

func (k msgServer) ProposeAddX509RootCert(goCtx context.Context, msg *types.MsgProposeAddX509RootCert) (*types.MsgProposeAddX509RootCertResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// decode pem certificate
	x509Certificate, err := x509.DecodeX509Certificate(msg.Cert)
	if err != nil {
		return nil, types.NewErrInvalidCertificate(err)
	}

	// fail if certificate is not self-signed
	if !x509Certificate.IsSelfSigned() {
		return nil, types.NewErrInappropriateCertificateType(
			"Inappropriate Certificate Type: Passed certificate is not self-signed, " +
				"so it cannot be used as a root certificate.")
	}

	// check if `Proposed` certificate with the same Subject/SubjectKeyId combination already exists
	if k.IsProposedCertificatePresent(ctx, x509Certificate.Subject, x509Certificate.SubjectKeyID) {
		return nil, types.NewErrProposedCertificateAlreadyExists(x509Certificate.Subject, x509Certificate.SubjectKeyID)
	}

	// check if certificate with Issuer/Serial Number combination already exists
	if k.IsUniqueCertificatePresent(ctx, x509Certificate.Issuer, x509Certificate.SerialNumber) {
		return nil, types.NewErrCertificateAlreadyExists(x509Certificate.Issuer, x509Certificate.SerialNumber)
	}

	// verify certificate
	_, _, err = k.verifyCertificate(ctx, x509Certificate)
	if err != nil {
		return nil, err
	}

	// Get list of certificates for Subject / Subject Key Id combination
	existingCertificates, found := k.GetApprovedCertificates(ctx, x509Certificate.Subject, x509Certificate.SubjectKeyID)
	if found {
		// Issuer and authorityKeyID must be the same as ones of exisiting certificates with the same subject and
		// subjectKeyID. Since new certificate is self-signed, we have to ensure that the exisiting certificates are
		// self-signed too, consequently are root certificates.
		if !existingCertificates.Certs[0].IsRoot {
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized,
				"Issuer and authorityKeyID of new certificate with subject=%v and subjectKeyID=%v "+
					"must be the same as ones of existing certificates with the same subject and subjectKeyID",
				x509Certificate.Subject, x509Certificate.SubjectKeyID)
		}

		// signer must be same as owner of existing certificates
		if msg.Signer != existingCertificates.Certs[0].Owner {
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized,
				"Only owner of existing certificates with subject=%v and subjectKeyID=%v "+
					"can add new certificate with the same subject and subjectKeyID",
				x509Certificate.Subject, x509Certificate.SubjectKeyID)
		}
	}

	// create a new proposed certificate with empty approvals list
	proposedCertificate := types.ProposedCertificate{
		Subject:      x509Certificate.Subject,
		SubjectKeyId: x509Certificate.SubjectKeyID,
		PemCert:      msg.Cert,
		SerialNumber: x509Certificate.SerialNumber,
		Owner:        msg.Signer,
		Approvals:    []types.Grant{},
	}

	// if signer has `RootCertificateApprovalRole` append approval
	signerAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid Address: (%s)", err)
	}
	if k.dclauthKeeper.HasRole(ctx, signerAddr, types.RootCertificateApprovalRole) {
		grant := types.Grant{
			Address: signerAddr.String(),
			Time:    msg.Time,
			Info:    msg.Info,
		}
		proposedCertificate.Approvals = append(proposedCertificate.Approvals, grant)
	}

	// store proposed certificate
	k.SetProposedCertificate(ctx, proposedCertificate)

	// register the unique certificate key
	uniqueCertificate := types.UniqueCertificate{
		Issuer:       x509Certificate.Issuer,
		SerialNumber: x509Certificate.SerialNumber,
		Present:      true,
	}
	k.SetUniqueCertificate(ctx, uniqueCertificate)

	return &types.MsgProposeAddX509RootCertResponse{}, nil
}
