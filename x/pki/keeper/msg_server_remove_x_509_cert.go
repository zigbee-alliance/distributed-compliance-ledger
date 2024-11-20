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
	if certificates[0].CertificateType == types.CertificateType_OperationalPKI {
		return nil, pkitypes.NewErrProvidedNotNocCertButExistingNoc(msg.Subject, msg.SubjectKeyId)
	}

	if err := k.EnsureVidMatches(ctx, certificates[0].Owner, msg.Signer); err != nil {
		return nil, err
	}

	if msg.SerialNumber != "" {
		certBySerialNumber, found := FindCertificateInList(msg.SerialNumber, &certificates)
		if !found {
			return nil, pkitypes.NewErrCertificateBySerialNumberDoesNotExist(msg.Subject, msg.SubjectKeyId, msg.SerialNumber)
		}

		// remove from subject with serialNumber map
		k.RemoveUniqueCertificate(ctx, certBySerialNumber.Issuer, certBySerialNumber.SerialNumber)

		if foundApproved {
			k.RemoveDaCertificateBySerialNumber(
				ctx,
				certBySerialNumber.Subject,
				certBySerialNumber.SubjectKeyId,
				&aprCerts,
				certBySerialNumber.SerialNumber,
				certBySerialNumber.Issuer,
				false,
			)
		}
		if foundRevoked {
			RemoveCertFromList(certBySerialNumber.Issuer, certBySerialNumber.SerialNumber, &revCerts.Certs)
			k.removeOrUpdateRevokedX509Cert(ctx, msg.Subject, msg.SubjectKeyId, &revCerts)
		}
	} else {
		// remove from noc certificates map
		k.RemoveDaCertificate(ctx, msg.Subject, msg.SubjectKeyId, false)
		// remove from revoked list
		k.RemoveRevokedCertificates(ctx, msg.Subject, msg.SubjectKeyId)
		// remove from subject with serialNumber map
		for _, cert := range certificates {
			k.RemoveUniqueCertificate(ctx, cert.Issuer, cert.SerialNumber)
		}
	}

	return &types.MsgRemoveX509CertResponse{}, nil
}
