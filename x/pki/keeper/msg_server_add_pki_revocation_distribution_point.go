package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	dclauthtypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"

	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/x509"
)

func (k msgServer) AddPkiRevocationDistributionPoint(goCtx context.Context, msg *types.MsgAddPkiRevocationDistributionPoint) (*types.MsgAddPkiRevocationDistributionPointResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	crlSignerCertificate, err := x509.DecodeX509Certificate(msg.CrlSignerCertificate)
	if err != nil {
		return nil, pkitypes.NewErrInvalidCertificate(err)
	}

	signerAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid Address: (%s)", err)
	}

	signerAccount, _ := k.dclauthKeeper.GetAccountO(ctx, signerAddr)

	// check if signer has vendor role
	if !k.dclauthKeeper.HasRole(ctx, signerAddr, dclauthtypes.Vendor) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized,
			"MsgAddPkiRevocationDistributionPoint transaction should be signed by an account with the \"%s\" role",
			dclauthtypes.Vendor,
		)
	}

	if msg.Vid != signerAccount.VendorID {
		return nil, sdkerrors.Wrap(pkitypes.ErrCRLSignerCertificateVidNotEqualAccountVid,
			"MsgAddPkiRevocationDistributionPoint signer must have the same vid as provided in message",
		)
	}

	if crlSignerCertificate.IsSelfSigned() {
		approvedCertificates, isFound := k.GetApprovedCertificates(ctx, crlSignerCertificate.Subject, crlSignerCertificate.SubjectKeyID)
		if !isFound {
			return nil, sdkerrors.Wrap(pkitypes.NewErrCertificateDoesNotExist(crlSignerCertificate.Subject, crlSignerCertificate.SubjectKeyID), "CRL signer Certificate must be a root certificate present on the ledger if isPAA = True")
		}

		if approvedCertificates.Certs[0].PemCert != msg.CrlSignerCertificate {
			return nil, pkitypes.NewErrPemValuesNotEqual("PEM values of the CRL signer certificate and a certificate found by its Subject and SubjectKeyID are not equal")
		}
	} else {
		_, _, err = k.verifyCertificate(ctx, crlSignerCertificate)

		if err != nil {
			return nil, err
		}
	}

	_, isFound := k.GetPkiRevocationDistributionPoint(ctx, msg.Vid, msg.Label, msg.IssuerSubjectKeyID)
	if isFound {
		return nil, pkitypes.NewErrPkiRevocationDistributionPointAlreadyExists("PKI revocation distribution point already exist")
	}

	pkiRevocationDistributionPoint := types.PkiRevocationDistributionPoint{
		Vid:                  msg.Vid,
		Label:                msg.Label,
		IssuerSubjectKeyID:   msg.IssuerSubjectKeyID,
		Pid:                  msg.Pid,
		IsPAA:                msg.IsPAA,
		CrlSignerCertificate: msg.CrlSignerCertificate,
		DataURL:              msg.DataURL,
		DataFileSize:         msg.DataFileSize,
		DataDigest:           msg.DataDigest,
		DataDigestType:       msg.DataDigestType,
		RevocationType:       msg.RevocationType,
	}

	k.SetPkiRevocationDistributionPoint(ctx, pkiRevocationDistributionPoint)
	k.AddPkiRevocationDistributionPointBySubjectKeyID(ctx, pkiRevocationDistributionPoint)

	return &types.MsgAddPkiRevocationDistributionPointResponse{}, nil
}
