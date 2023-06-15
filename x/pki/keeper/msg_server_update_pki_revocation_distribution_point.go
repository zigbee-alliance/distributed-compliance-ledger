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

func verifyUpdatedPAA(updatedCrlSignerCertificate *x509.Certificate, revocationPointVid int32) error {
	if !updatedCrlSignerCertificate.IsSelfSigned() {
		return pkitypes.NewErrRootCertificateIsNotSelfSigned("Updated CRL signer certificate must be self-signed since old one was self-signed")
	}

	vid, err := x509.GetVidFromSubject(updatedCrlSignerCertificate.SubjectAsText)

	if err != nil {
		return sdkerrors.Wrapf(pkitypes.ErrInvalidVidFormat, "Could not parse vid: %s", err)
	}

	if vid == 0 {
		return pkitypes.NewErrUnsupportedOperation("publishing a revocation point for non-VID scoped root certificates is currently not supported")
	}

	if vid != revocationPointVid {
		return pkitypes.NewErrCRLSignerCertificateVidNotEqualMsgVid("CRL Signer Certificate's vid must be equal to the provided vid in the message")
	}

	return nil
}

func verifyUpdatedPAI(updatedCrlSignerCertificate *x509.Certificate, msgVid int32, revocationPointPid int32) error {
	if updatedCrlSignerCertificate.IsSelfSigned() {
		return pkitypes.NewErrNonRootCertificateSelfSigned("Updated CRL signer certificate must not be self-signed since old one was not self-signed")
	}

	vid, err := x509.GetVidFromSubject(updatedCrlSignerCertificate.SubjectAsText)

	if err != nil {
		return sdkerrors.Wrapf(pkitypes.ErrInvalidVidFormat, "Could not parse vid: %s", err)
	}

	if vid == 0 {
		return pkitypes.NewErrVidNotFound("vid must be present in updated non-root CRL signer certificate")
	}

	if vid != msgVid {
		return pkitypes.NewErrCRLSignerCertificateVidNotEqualMsgVid("CRL Signer Certificate's vid must be equal to the provided vid in the message")
	}

	pid, err := x509.GetPidFromSubject(updatedCrlSignerCertificate.SubjectAsText)

	if err != nil {
		return sdkerrors.Wrapf(pkitypes.ErrInvalidPidFormat, "Could not parse pid: %s", err)
	}

	if pid != 0 && pid != revocationPointPid {
		return pkitypes.NewErrCRLSignerCertificatePidNotEqualMsgPid("pid in updated CRL Signer Certificate must be equal to pid in revocation point")
	}

	if pid == 0 && pid != revocationPointPid {
		return pkitypes.NewErrPidNotFound("pid not found in updated CRL Signer Certificate when it is provided in revocation point")
	}

	return nil
}

func verifyUpdatedCertificate(updatedCertificate string, revocationPoint types.PkiRevocationDistributionPoint, msgVid int32, isPrevCertPAA bool) error {
	updatedCrlSignerCertificate, err := x509.DecodeX509Certificate(updatedCertificate)
	if err != nil {
		return pkitypes.NewErrInvalidCertificate(err)
	}

	if isPrevCertPAA {
		err := verifyUpdatedPAA(updatedCrlSignerCertificate, revocationPoint.Vid)
		if err != nil {
			return err
		}
	} else {
		err := verifyUpdatedPAI(updatedCrlSignerCertificate, msgVid, revocationPoint.Pid)
		if err != nil {
			return err
		}
	}

	return nil
}

func (k msgServer) UpdatePkiRevocationDistributionPoint(goCtx context.Context, msg *types.MsgUpdatePkiRevocationDistributionPoint) (*types.MsgUpdatePkiRevocationDistributionPointResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	signerAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid Address: (%s)", err)
	}

	signerAccount, _ := k.dclauthKeeper.GetAccountO(ctx, signerAddr)

	// check if signer has vendor role
	if !k.dclauthKeeper.HasRole(ctx, signerAddr, dclauthtypes.Vendor) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized,
			"MsgUpdatePkiRevocationDistributionPoint transaction should be signed by an account with the \"%s\" role",
			dclauthtypes.Vendor,
		)
	}

	pkiRevocationDistributionPoint, isFound := k.GetPkiRevocationDistributionPoint(ctx, msg.Vid, msg.Label, msg.IssuerSubjectKeyID)
	if !isFound {
		return nil, pkitypes.NewErrPkiRevocationDistributionPointDoesNotExists("PKI revocation distribution point does not exist")
	}

	crlSignerCertificate, err := x509.DecodeX509Certificate(pkiRevocationDistributionPoint.CrlSignerCertificate)
	if err != nil {
		return nil, pkitypes.NewErrInvalidCertificate(err)
	}

	if pkiRevocationDistributionPoint.Vid != signerAccount.VendorID {
		return nil, sdkerrors.Wrap(pkitypes.ErrCRLSignerCertificateVidNotEqualAccountVid,
			"MsgUpdatePkiRevocationDistributionPoint signer must have the same vid as provided in an existing certificate from the revocation point",
		)
	}

	if msg.CrlSignerCertificate != "" {
		err = verifyUpdatedCertificate(msg.CrlSignerCertificate, pkiRevocationDistributionPoint, msg.Vid, crlSignerCertificate.IsSelfSigned())

		if err != nil {
			return nil, err
		}

		pkiRevocationDistributionPoint.CrlSignerCertificate = msg.CrlSignerCertificate
	}

	if msg.DataURL != "" {
		pkiRevocationDistributionPoint.DataURL = msg.DataURL
	}

	if msg.DataFileSize != 0 {
		pkiRevocationDistributionPoint.DataFileSize = msg.DataFileSize
	}

	if msg.DataDigest != "" {
		pkiRevocationDistributionPoint.DataDigest = msg.DataDigest
	}

	if msg.DataDigestType != 0 {
		pkiRevocationDistributionPoint.DataDigestType = msg.DataDigestType
	}

	k.SetPkiRevocationDistributionPoint(ctx, pkiRevocationDistributionPoint)
	k.UpdatePkiRevocationDistributionPointBySubjectKeyID(ctx, pkiRevocationDistributionPoint)

	return &types.MsgUpdatePkiRevocationDistributionPointResponse{}, nil
}
