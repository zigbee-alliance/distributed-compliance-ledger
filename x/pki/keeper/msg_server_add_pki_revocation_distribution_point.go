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

	// decode CrlSignerCertificate
	crlSignerCertificate, err := x509.DecodeX509Certificate(msg.CrlSignerCertificate)
	if err != nil {
		return nil, pkitypes.NewErrInvalidCertificate(err)
	}

	// check if signer has vendor role
	signerAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid Address: (%s)", err)
	}
	signerAccount, _ := k.dclauthKeeper.GetAccountO(ctx, signerAddr)
	if !k.dclauthKeeper.HasRole(ctx, signerAddr, dclauthtypes.Vendor) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized,
			"MsgAddPkiRevocationDistributionPoint transaction should be signed by an account with the \"%s\" role",
			dclauthtypes.Vendor,
		)
	}

	// compare VID in message and Vendor acount
	if msg.Vid != signerAccount.VendorID {
		return nil, sdkerrors.Wrap(pkitypes.ErrCRLSignerCertificateVidNotEqualAccountVid,
			"MsgAddPkiRevocationDistributionPoint signer must have the same vid as provided in message",
		)
	}

	// check that distribution point doesn't exist yet
	_, isFound := k.GetPkiRevocationDistributionPoint(ctx, msg.Vid, msg.Label, msg.IssuerSubjectKeyID)
	if isFound {
		return nil, pkitypes.NewErrPkiRevocationDistributionPointAlreadyExists("PKI revocation distribution point already exist")
	}

	if crlSignerCertificate.IsSelfSigned() {
		// check that crlSignerCertificate cert is present on the ledger and has the same VID
		err = k.checkRootCert(ctx, crlSignerCertificate, msg)
	} else {
		// check that crlSignerCertificate is chained back to a certificate on the ledger
		err = k.checkNonRootCert(ctx, crlSignerCertificate)
	}
	if err != nil {
		return nil, err
	}

	// add to state
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

func (k msgServer) checkRootCert(ctx sdk.Context, crlSignerCertificate *x509.Certificate, msg *types.MsgAddPkiRevocationDistributionPoint) error {
	// find the cert on the ledger
	approvedCertificates, isFound := k.GetApprovedCertificates(ctx, crlSignerCertificate.Subject, crlSignerCertificate.SubjectKeyID)
	if !isFound {
		return sdkerrors.Wrap(pkitypes.NewErrCertificateDoesNotExist(crlSignerCertificate.Subject, crlSignerCertificate.SubjectKeyID), "CRL signer Certificate must be a root certificate present on the ledger if isPAA = True")
	}

	// check that it has the same PEM value
	var foundRootCert *types.Certificate
	for _, approvedCertificate := range approvedCertificates.Certs {
		if approvedCertificate.PemCert == msg.CrlSignerCertificate {
			foundRootCert = approvedCertificate

			break
		}
	}
	if foundRootCert == nil {
		return pkitypes.NewErrPemValuesNotEqual("PEM values of the CRL signer certificate and a certificate found by its Subject and SubjectKeyID are not equal")
	}

	// check that root cert has the same VID as in the message if it's non-VID scoped
	// (vid-scoped has been already checked as patr of static validation + equality of PEM values
	ledgerRootVid, err := x509.GetVidFromSubject(foundRootCert.SubjectAsText)
	if err != nil {
		return sdkerrors.Wrapf(pkitypes.ErrInvalidVidFormat, "Could not parse vid: %s", err)
	}
	if ledgerRootVid == 0 && msg.Vid != foundRootCert.Vid {
		return pkitypes.NewErrMessageVidNotEqualRootCertVid(msg.Vid, foundRootCert.Vid)
	}

	return nil
}

func (k msgServer) checkNonRootCert(ctx sdk.Context, crlSignerCertificate *x509.Certificate) error {
	// check that it's chained back to a cert on DCL
	if _, _, err := k.verifyCertificate(ctx, crlSignerCertificate); err != nil {
		return pkitypes.NewErrCertNotChainedBack()
	}

	return nil
}
