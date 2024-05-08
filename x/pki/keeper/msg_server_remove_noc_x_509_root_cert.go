package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	dclauthtypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

func (k msgServer) RemoveNocX509RootCert(goCtx context.Context, msg *types.MsgRemoveNocX509RootCert) (*types.MsgRemoveNocX509RootCertResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	signerAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return nil, pkitypes.NewErrInvalidAddress(err)
	}

	// check if signer has vendor role
	if !k.dclauthKeeper.HasRole(ctx, signerAddr, dclauthtypes.Vendor) {
		return nil, pkitypes.NewErrUnauthorizedRole("MsgRemoveNocX509RootCert", dclauthtypes.Vendor)
	}

	signerAccount, _ := k.dclauthKeeper.GetAccountO(ctx, signerAddr)
	accountVid := signerAccount.VendorID

	nocCerts, foundActive := k.GetApprovedCertificates(ctx, msg.Subject, msg.SubjectKeyId)
	revCerts, foundRevoked := k.GetRevokedNocRootCertificates(ctx, msg.Subject, msg.SubjectKeyId)
	certificates := nocCerts.Certs
	certificates = append(certificates, revCerts.Certs...)
	if len(certificates) == 0 {
		return nil, pkitypes.NewErrCertificateDoesNotExist(msg.Subject, msg.SubjectKeyId)
	}

	cert := certificates[0]
	// Existing certificate must be Root certificate
	if !cert.IsRoot {
		return nil, pkitypes.NewErrMessageExistingCertIsNotRoot(cert.Subject, cert.SubjectKeyId)
	}

	// Existing certificate must be NOC certificate
	if !cert.IsNoc {
		return nil, pkitypes.NewErrProvidedNocCertButExistingNotNoc(msg.Subject, msg.SubjectKeyId)
	}

	// account VID must be same as VID of existing certificates
	if accountVid != cert.Vid {
		return nil, pkitypes.NewErrRevokeCertVidNotEqualToAccountVid(cert.Vid, accountVid)
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
		k.RemoveUniqueCertificate(ctx, certBySerialNumber.Subject, certBySerialNumber.SerialNumber)

		if foundActive {
			// Remove from Approved lists
			removeCertFromList(certBySerialNumber.Issuer, certBySerialNumber.SerialNumber, &nocCerts.Certs)
			k.removeApprovedX509Cert(ctx, certID, &nocCerts, msg.SerialNumber)

			// Remove from NOC lists
			k.RemoveNocRootCertificateBySerialNumber(ctx, accountVid, certID.Subject, certID.SubjectKeyId, msg.SerialNumber)
			k.RemoveNocRootCertificateByVidSubjectSkidAndSerialNumber(ctx, accountVid, certID.Subject, certID.SubjectKeyId, msg.SerialNumber)
		}

		if foundRevoked {
			removeCertFromList(certBySerialNumber.Issuer, certBySerialNumber.SerialNumber, &revCerts.Certs)
			k._removeRevokedNocX509RootCert(ctx, certID, &revCerts)
		}
	} else {
		k.RemoveNocRootCertificate(ctx, accountVid, certID.Subject, certID.SubjectKeyId)
		// remove from vid, subject key id map
		k.RemoveNocRootCertificatesByVidAndSkid(ctx, accountVid, certID.SubjectKeyId)
		// remove from revoked noc root certs
		k.RemoveRevokedNocRootCertificates(ctx, certID.Subject, certID.SubjectKeyId)
		// remove from revoked list
		k.RemoveRevokedCertificates(ctx, certID.Subject, certID.SubjectKeyId)
		// remove from approved list
		k.RemoveApprovedCertificates(ctx, certID.Subject, certID.SubjectKeyId)
		// remove from subject -> subject key ID map
		k.RemoveApprovedCertificateBySubject(ctx, certID.Subject, certID.SubjectKeyId)
		// remove from subject key ID -> certificates map
		k.RemoveApprovedCertificatesBySubjectKeyID(ctx, certID.Subject, certID.SubjectKeyId)
		// remove from subject with serialNumber map
		for _, cert := range certificates {
			k.RemoveUniqueCertificate(ctx, cert.Subject, cert.SerialNumber)
		}
	}

	return &types.MsgRemoveNocX509RootCertResponse{}, nil
}

func (k msgServer) _removeRevokedNocX509RootCert(ctx sdk.Context, certID types.CertificateIdentifier, certificates *types.RevokedNocRootCertificates) {
	if len(certificates.Certs) == 0 {
		k.RemoveRevokedNocRootCertificates(ctx, certID.Subject, certID.SubjectKeyId)
		k.RemoveRevokedCertificates(ctx, certID.Subject, certID.SubjectKeyId)
	} else {
		k.SetRevokedNocRootCertificates(ctx, *certificates)
		k.SetRevokedCertificates(
			ctx,
			types.RevokedCertificates{
				Subject:      certificates.Subject,
				SubjectKeyId: certificates.SubjectKeyId,
				Certs:        certificates.Certs,
			},
		)
	}
}
