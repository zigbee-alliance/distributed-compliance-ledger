package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

func (k msgServer) RemoveX509Cert(goCtx context.Context, msg *types.MsgRemoveX509Cert) (*types.MsgRemoveX509CertResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	certificates, found := k.GetApprovedCertificates(ctx, msg.Subject, msg.SubjectKeyId)
	if !found {
		return nil, pkitypes.NewErrCertificateDoesNotExist(msg.Subject, msg.SubjectKeyId)
	}

	if certificates.Certs[0].IsRoot {
		return nil, pkitypes.NewErrInappropriateCertificateType(
			fmt.Sprintf("Inappropriate Certificate Type: Certificate with subject=%v and subjectKeyID=%v "+
				"is a root certificate. To propose revocation of a root certificate please use "+
				"`PROPOSE_REVOKE_X509_ROOT_CERT` transaction.", msg.Subject, msg.SubjectKeyId),
		)
	}

	if msg.Signer != certificates.Certs[0].Owner {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized,
			"Only owner can revoke certificate using `REVOKE_X509_CERT`",
		)
	}

	certID := types.CertificateIdentifier{
		Subject:      msg.Subject,
		SubjectKeyId: msg.SubjectKeyId,
	}

	if msg.SerialNumber != "" {
		certBySerialNumber, found := findCertificate(msg.SerialNumber, &certificates.Certs)
		if !found {
			return nil, pkitypes.NewErrCertificateBySerialNumberDoesNotExist(msg.Subject, msg.SubjectKeyId, msg.SerialNumber)
		}
		// remove from subject with serialNumber map
		k.RemoveUniqueCertificate(ctx, certBySerialNumber.Issuer, certBySerialNumber.SerialNumber)

		k.removeCertFromList(certBySerialNumber.Issuer, certBySerialNumber.SerialNumber, &certificates)
		k._removeX509Cert(ctx, certID, certificates)
	} else {
		k.RemoveApprovedCertificates(ctx, certID.Subject, certID.SubjectKeyId)
		// remove from subject -> subject key ID map
		k.RemoveApprovedCertificateBySubject(ctx, certID.Subject, certID.SubjectKeyId)
		// remove from subject key ID -> certificates map
		k.RemoveApprovedCertificatesBySubjectKeyID(ctx, certID.Subject, certID.SubjectKeyId)
		// remove from subject with serialNumber map
		for _, cert := range certificates.Certs {
			k.RemoveUniqueCertificate(ctx, cert.Issuer, cert.SerialNumber)
		}
	}

	return &types.MsgRemoveX509CertResponse{}, nil
}

func (k msgServer) _removeX509Cert(ctx sdk.Context, certID types.CertificateIdentifier, certificates types.ApprovedCertificates) {
	if len(certificates.Certs) == 0 {
		k.RemoveApprovedCertificates(ctx, certID.Subject, certID.SubjectKeyId)
		k.RemoveApprovedCertificateBySubject(ctx, certID.Subject, certID.SubjectKeyId)
		k.RemoveApprovedCertificatesBySubjectKeyID(ctx, certID.Subject, certID.SubjectKeyId)
	} else {
		k.SetApprovedCertificates(ctx, certificates)
		k.SetApprovedCertificatesBySubjectKeyID(
			ctx,
			types.ApprovedCertificatesBySubjectKeyId{SubjectKeyId: certID.SubjectKeyId, Certs: certificates.Certs},
		)
	}
}
