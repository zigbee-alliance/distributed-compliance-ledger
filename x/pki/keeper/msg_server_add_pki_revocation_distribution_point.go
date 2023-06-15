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

func verifyVid(subjectAsText string, signerVid int32) error {
	vid, err := x509.GetVidFromSubject(subjectAsText)

	if err != nil {
		return sdkerrors.Wrapf(pkitypes.ErrInvalidVidFormat, "Could not parse vid: %s", err)
	}

	if vid == 0 {
		return pkitypes.NewErrUnsupportedOperation("publishing a revocation point for non-VID scoped root certificates is currently not supported")
	}

	if vid != signerVid {
		return pkitypes.NewErrCRLSignerCertificateVidNotEqualAccountVid("CRL signer Certificate's vid must be equal to signer account's vid")
	}

	return nil
}

func (k msgServer) verifyPAA(crlSignerCertificate *x509.Certificate, certPemValue string, signerVid int32, approvedCertificates types.ApprovedCertificates) error {
	err := verifyVid(crlSignerCertificate.SubjectAsText, signerVid)

	if err != nil {
		return err
	}

	if approvedCertificates.Certs[0].PemCert != certPemValue {
		return pkitypes.NewErrPemValuesNotEqual("Pem values of CRL signer certificate and certificate found by its Subject and SubjectKeyID are is not equal ")
	}

	return nil
}

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

		err = k.verifyPAA(crlSignerCertificate, msg.CrlSignerCertificate, signerAccount.VendorID, approvedCertificates)

		if err != nil {
			return nil, err
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
