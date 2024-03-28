package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

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
		return nil, pkitypes.NewErrInvalidAddress(err)
	}
	signerAccount, _ := k.dclauthKeeper.GetAccountO(ctx, signerAddr)
	if !k.dclauthKeeper.HasRole(ctx, signerAddr, dclauthtypes.Vendor) {
		return nil, pkitypes.NewErrUnauthorizedRole("MsgAddPkiRevocationDistributionPoint", dclauthtypes.Vendor)
	}

	// compare VID in message and Vendor acount
	if msg.Vid != signerAccount.VendorID {
		return nil, pkitypes.NewErrMessageVidNotEqualAccountVid(msg.Vid, signerAccount.VendorID)
	}

	// check that distribution point doesn't exist yet
	_, isFound := k.GetPkiRevocationDistributionPoint(ctx, msg.Vid, msg.Label, msg.IssuerSubjectKeyID)
	if isFound {
		return nil, pkitypes.NewErrPkiRevocationDistributionPointWithVidAndLabelAlreadyExists(msg.Vid, msg.Label, msg.IssuerSubjectKeyID)
	}

	if crlSignerCertificate.IsSelfSigned() {
		// check that crlSignerCertificate cert is present on the ledger and has the same VID
		err = k.checkRootCert(ctx, crlSignerCertificate, msg)
	} else {
		// check that crlSignerCertificate is chained back to a certificate on the ledger
		err = k.checkCRLSignerNonRootCert(ctx, crlSignerCertificate, msg.CrlSignerDelegator, msg.IsPAA)
	}
	if err != nil {
		return nil, err
	}

	revocationList, isFound := k.GetPkiRevocationDistributionPointsByIssuerSubjectKeyID(ctx, msg.IssuerSubjectKeyID)
	if isFound {
		for _, revocationPoint := range revocationList.Points {
			if revocationPoint.DataURL == msg.DataURL && revocationPoint.Vid == msg.Vid {
				return nil, pkitypes.NewErrPkiRevocationDistributionPointWithDataURLAlreadyExists(msg.DataURL, msg.IssuerSubjectKeyID)
			}
		}
	}

	// add to state
	pkiRevocationDistributionPoint := types.PkiRevocationDistributionPoint{
		Vid:                  msg.Vid,
		Label:                msg.Label,
		IssuerSubjectKeyID:   msg.IssuerSubjectKeyID,
		Pid:                  msg.Pid,
		IsPAA:                msg.IsPAA,
		CrlSignerCertificate: msg.CrlSignerCertificate,
		CrlSignerDelegator:   msg.CrlSignerDelegator,
		DataURL:              msg.DataURL,
		DataFileSize:         msg.DataFileSize,
		DataDigest:           msg.DataDigest,
		DataDigestType:       msg.DataDigestType,
		RevocationType:       msg.RevocationType,
		SchemaVersion:        msg.SchemaVersion,
	}

	k.SetPkiRevocationDistributionPoint(ctx, pkiRevocationDistributionPoint)
	k.AddPkiRevocationDistributionPointBySubjectKeyID(ctx, pkiRevocationDistributionPoint)

	return &types.MsgAddPkiRevocationDistributionPointResponse{}, nil
}

func (k msgServer) checkRootCert(ctx sdk.Context, crlSignerCertificate *x509.Certificate, msg *types.MsgAddPkiRevocationDistributionPoint) error {
	// find the cert on the ledger
	approvedCertificates, isFound := k.GetApprovedCertificates(ctx, crlSignerCertificate.Subject, crlSignerCertificate.SubjectKeyID)
	if !isFound {
		return pkitypes.NewErrRootCertificateDoesNotExist(crlSignerCertificate.Subject, crlSignerCertificate.SubjectKeyID)
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
		return pkitypes.NewErrPemValuesNotEqual(crlSignerCertificate.Subject, crlSignerCertificate.SubjectKeyID)
	}

	// check that root cert has the same VID as in the message if it's non-VID scoped
	// (vid-scoped has been already checked as patr of static validation + equality of PEM values
	ledgerRootVid, err := x509.GetVidFromSubject(foundRootCert.SubjectAsText)
	if err != nil {
		return pkitypes.NewErrInvalidVidFormat(err)
	}
	if ledgerRootVid == 0 && msg.Vid != foundRootCert.Vid {
		return pkitypes.NewErrMessageVidNotEqualRootCertVid(msg.Vid, foundRootCert.Vid)
	}

	return nil
}

func (k msgServer) checkCRLSignerNonRootCert(ctx sdk.Context, crlSignerCertificate *x509.Certificate, crlSignerDelegator string, isPAA bool) error {
	if crlSignerDelegator != "" && !isPAA {
		crlSignerDelegatorCert, err := x509.DecodeX509Certificate(crlSignerDelegator)
		if err != nil {
			return pkitypes.NewErrInvalidCertificate(err)
		}

		// verify CRL Signer certificate against Delegated PAI certificate
		if err = crlSignerCertificate.Verify(crlSignerDelegatorCert, ctx.BlockTime()); err != nil {
			return pkitypes.NewErrCRLSignerCertNotChainedBackToDelegator()
		}

		if _, err = k.verifyCertificate(ctx, crlSignerDelegatorCert); err != nil {
			return pkitypes.NewErrCRLSignerCertDelegatorNotChainedBack()
		}

		return nil
	}

	// check that it's chained back to a cert on DCL
	_, err := k.verifyCertificate(ctx, crlSignerCertificate)
	if err != nil {
		return pkitypes.NewErrCertNotChainedBack()
	}

	return nil
}
