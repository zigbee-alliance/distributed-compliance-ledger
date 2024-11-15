package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	dclauthtypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

func (k msgServer) RemoveNocX509IcaCert(goCtx context.Context, msg *types.MsgRemoveNocX509IcaCert) (*types.MsgRemoveNocX509IcaCertResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	signerAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return nil, pkitypes.NewErrInvalidAddress(err)
	}

	// check if signer has vendor role
	if !k.dclauthKeeper.HasRole(ctx, signerAddr, dclauthtypes.Vendor) {
		return nil, pkitypes.NewErrUnauthorizedRole("MsgRemoveNocX509IcaCert", dclauthtypes.Vendor)
	}

	signerAccount, _ := k.dclauthKeeper.GetAccountO(ctx, signerAddr)
	accountVid := signerAccount.VendorID

	icaCerts, foundActive := k.GetNocCertificates(ctx, msg.Subject, msg.SubjectKeyId)
	revCerts, foundRevoked := k.GetRevokedNocIcaCertificates(ctx, msg.Subject, msg.SubjectKeyId)
	certificates := icaCerts.Certs
	certificates = append(certificates, revCerts.Certs...)
	if len(certificates) == 0 {
		return nil, pkitypes.NewErrCertificateDoesNotExist(msg.Subject, msg.SubjectKeyId)
	}

	cert := certificates[0]
	// Existing certificate must be Root certificate
	if cert.IsRoot {
		return nil, pkitypes.NewErrMessageExpectedNonRoot(cert.Subject, cert.SubjectKeyId)
	}

	// Existing certificate must be NOC certificate
	if cert.CertificateType != types.CertificateType_OperationalPKI {
		return nil, pkitypes.NewErrProvidedNocCertButExistingNotNoc(msg.Subject, msg.SubjectKeyId)
	}

	// account VID must be same as VID of existing certificates
	if accountVid != cert.Vid {
		return nil, pkitypes.NewErrRevokeCertVidNotEqualToAccountVid(cert.Vid, accountVid)
	}

	if err = k.EnsureVidMatches(ctx, certificates[0].Owner, msg.Signer); err != nil {
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

		if foundActive {
			// Remove from certificates lists
			removeCertFromList(certBySerialNumber.Issuer, certBySerialNumber.SerialNumber, &icaCerts.Certs)
			k.removeNocX509Cert(ctx, certID, &icaCerts, accountVid, msg.SerialNumber, false)
		}

		if foundRevoked {
			removeCertFromList(certBySerialNumber.Issuer, certBySerialNumber.SerialNumber, &revCerts.Certs)
			k._removeRevokedNocX509IcaCert(ctx, certID, &revCerts)
		}
	} else {
		// remove from global certificates map
		k.RemoveAllCertificates(ctx, certID.Subject, certID.SubjectKeyId)
		// remove from global subject -> subject key ID map
		k.RemoveAllCertificateBySubject(ctx, certID.Subject, certID.SubjectKeyId)
		// remove from noc certificates map
		k.RemoveNocCertificates(ctx, certID.Subject, certID.SubjectKeyId)
		// remove from noc ica certificates map
		k.RemoveNocIcaCertificate(ctx, certID.Subject, certID.SubjectKeyId, accountVid)
		// remove from vid, subject key id map
		k.RemoveNocCertificatesByVidAndSkid(ctx, accountVid, certID.SubjectKeyId)
		// remove from subject -> subject key ID map
		k.RemoveNocCertificateBySubject(ctx, certID.Subject, certID.SubjectKeyId)
		// remove from subject key ID -> certificates map
		k.RemoveNocCertificatesBySubjectAndSubjectKeyID(ctx, certID.Subject, certID.SubjectKeyId)
		// remove from revoked list
		k.RemoveRevokedNocIcaCertificates(ctx, certID.Subject, certID.SubjectKeyId)
		// remove from subject with serialNumber map
		for _, cert := range certificates {
			k.RemoveUniqueCertificate(ctx, cert.Issuer, cert.SerialNumber)
		}
	}

	return &types.MsgRemoveNocX509IcaCertResponse{}, nil
}

func (k msgServer) _removeRevokedNocX509IcaCert(ctx sdk.Context, certID types.CertificateIdentifier, certificates *types.RevokedNocIcaCertificates) {
	if len(certificates.Certs) == 0 {
		k.RemoveRevokedNocIcaCertificates(ctx, certID.Subject, certID.SubjectKeyId)
	} else {
		k.SetRevokedNocIcaCertificates(ctx, *certificates)
	}
}
