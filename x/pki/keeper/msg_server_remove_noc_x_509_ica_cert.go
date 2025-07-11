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
	if cert.CertificateType != types.CertificateType_OperationalPKI && cert.CertificateType != types.CertificateType_VIDSignerPKI {
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

	if msg.SerialNumber != "" { //nolint:nestif
		certBySerialNumber, found := FindCertificateInList(msg.SerialNumber, &certificates)
		if !found {
			return nil, pkitypes.NewErrCertificateBySerialNumberDoesNotExist(msg.Subject, msg.SubjectKeyId, msg.SerialNumber)
		}

		// remove from subject with serialNumber map
		k.RemoveUniqueCertificate(ctx, certBySerialNumber.Issuer, certBySerialNumber.SerialNumber)

		if foundActive {
			// Remove from certificates lists
			k.RemoveNocCertBySerialNumber(
				ctx,
				certBySerialNumber.Subject,
				certBySerialNumber.SubjectKeyId,
				&icaCerts,
				accountVid,
				certBySerialNumber.SerialNumber,
				certBySerialNumber.Issuer,
				false,
			)
			if len(icaCerts.Certs) == 0 {
				k.RemoveChildCertificate(ctx, certBySerialNumber.Issuer, certBySerialNumber.AuthorityKeyId, types.CertificateIdentifier{
					Subject:      icaCerts.Subject,
					SubjectKeyId: icaCerts.SubjectKeyId,
				})
			}
		}

		if foundRevoked {
			RemoveCertFromList(certBySerialNumber.Issuer, certBySerialNumber.SerialNumber, &revCerts.Certs)
			k.removeRevokedNocX509IcaCert(ctx, certID, &revCerts)
		}
	} else {
		// remove from revoked list
		k.RemoveRevokedNocIcaCertificates(ctx, certID.Subject, certID.SubjectKeyId)
		// remove from noc certificates map
		k.RemoveNocCertificate(ctx, cert.Subject, cert.SubjectKeyId, accountVid, false)
		// Remove certificate identifier from issuer's ChildCertificates record
		k.RemoveChildCertificate(ctx, certificates[0].Issuer, certificates[0].AuthorityKeyId, certID)
		// remove from subject with serialNumber map
		for _, cert := range certificates {
			k.RemoveUniqueCertificate(ctx, cert.Issuer, cert.SerialNumber)
		}
	}

	return &types.MsgRemoveNocX509IcaCertResponse{}, nil
}

func (k msgServer) removeRevokedNocX509IcaCert(ctx sdk.Context, certID types.CertificateIdentifier, certificates *types.RevokedNocIcaCertificates) {
	if len(certificates.Certs) == 0 {
		k.RemoveRevokedNocIcaCertificates(ctx, certID.Subject, certID.SubjectKeyId)
	} else {
		k.SetRevokedNocIcaCertificates(ctx, *certificates)
	}
}
