package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	dclauthtypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

func (k msgServer) RemoveX509Cert(goCtx context.Context, msg *types.MsgRemoveX509Cert) (*types.MsgRemoveX509CertResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	signerAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return nil, pkitypes.NewErrInvalidAddress(err)
	}

	// check if signer has vendor role
	if !k.dclauthKeeper.HasRole(ctx, signerAddr, dclauthtypes.Vendor) {
		return nil, pkitypes.NewErrUnauthorizedRole("MsgRemoveX509Cert", dclauthtypes.Vendor)
	}

	aprCerts, foundApproved := k.GetApprovedCertificates(ctx, msg.Subject, msg.SubjectKeyId)
	revCerts, foundRevoked := k.GetRevokedCertificates(ctx, msg.Subject, msg.SubjectKeyId)
	certificates := aprCerts.Certs
	certificates = append(certificates, revCerts.Certs...)
	if len(certificates) == 0 {
		return nil, pkitypes.NewErrCertificateDoesNotExist(msg.Subject, msg.SubjectKeyId)
	}

	if certificates[0].IsRoot {
		return nil, pkitypes.NewErrMessageExpectedNonRoot(msg.Subject, msg.SubjectKeyId)
	}

	// Existing certificate must not be NOC certificate
	if certificates[0].IsNoc {
		return nil, pkitypes.NewErrProvidedNotNocCertButExistingNoc(msg.Subject, msg.SubjectKeyId)
	}

	if err := k.EnsureVidMatches(ctx, certificates[0].Owner, msg.Signer); err != nil {
		return nil, err
	}

	certID := types.CertificateIdentifier{
		Subject:      msg.Subject,
		SubjectKeyId: msg.SubjectKeyId,
	}

	if msg.SerialNumber != "" {
		certBySerialNumber, found := findCertificate(msg.SerialNumber, &certificates)
		if !found {
			return nil, pkitypes.NewErrCertificateBySerialNumberDoesNotExist(msg.Subject, msg.SubjectKeyId, msg.SerialNumber)
		}

		// remove from subject with serialNumber map
		k.RemoveUniqueCertificate(ctx, certBySerialNumber.Issuer, certBySerialNumber.SerialNumber)

		if foundApproved {
			removeCertFromList(certBySerialNumber.Issuer, certBySerialNumber.SerialNumber, &aprCerts.Certs)
			k.removeApprovedX509Cert(ctx, certID, &aprCerts, msg.SerialNumber)
		}
		if foundRevoked {
			removeCertFromList(certBySerialNumber.Issuer, certBySerialNumber.SerialNumber, &revCerts.Certs)
			k.removeOrUpdateRevokedX509Cert(ctx, certID, &revCerts)
		}
	} else {
		k.RemoveApprovedCertificates(ctx, certID.Subject, certID.SubjectKeyId)
		// remove from subject -> subject key ID map
		k.RemoveApprovedCertificateBySubject(ctx, certID.Subject, certID.SubjectKeyId)
		// remove from subject key ID -> certificates map
		k.RemoveApprovedCertificatesBySubjectKeyID(ctx, certID.Subject, certID.SubjectKeyId)
		// remove from revoked list
		k.RemoveRevokedCertificates(ctx, certID.Subject, certID.SubjectKeyId)
		// remove from subject with serialNumber map
		for _, cert := range certificates {
			k.RemoveUniqueCertificate(ctx, cert.Issuer, cert.SerialNumber)
		}
	}

	return &types.MsgRemoveX509CertResponse{}, nil
}
