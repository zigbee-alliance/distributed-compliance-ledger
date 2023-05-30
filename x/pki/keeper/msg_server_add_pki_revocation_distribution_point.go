package keeper

import (
	"context"
	"strconv"
	"strings"

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

	if crlSignerCertificate.IsSelfSigned() {
		subjectAsMap := x509.SubjectAsTextToMap(crlSignerCertificate.SubjectAsText)

		strVid, found := subjectAsMap["Mvid"]
		if found {
			vid, err := strconv.ParseInt(strings.Trim(strVid, "0x"), 16, 32)
			if err != nil {
				return nil, err
			}

			// check if signer has vendor role
			if !k.dclauthKeeper.HasRole(ctx, signerAddr, dclauthtypes.Vendor) {
				return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized,
					"MsgAddPkiRevocationDistributionPoint transaction should be signed by an account with the \"%s\" role",
					dclauthtypes.Vendor,
				)
			}

			if int32(vid) != signerAccount.VendorID {
				return nil, pkitypes.NewErrCRLSignerCertificateVidNotEqualAccountVid("CRL signer Certificate's vid must be equal to signer account's vid")
			}
		}

		approvedCertificates, isFound := k.GetApprovedCertificates(ctx, crlSignerCertificate.Subject, crlSignerCertificate.SubjectKeyID)
		if !isFound {
			return nil, pkitypes.NewErrCertificateDoesNotExist(crlSignerCertificate.Subject, crlSignerCertificate.SubjectKeyID)
		}

		if approvedCertificates.Certs[0].PemCert != msg.CrlSignerCertificate {
			return nil, pkitypes.NewErrPemValuesNotEqual("PEM value of CRL signer certificate and certificate found by its Subject and SubjectKeyID on the ledger is not equal")
		}
	} else {
		// check if signer has vendor role
		if !k.dclauthKeeper.HasRole(ctx, signerAddr, dclauthtypes.Vendor) {
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized,
				"MsgAddPkiRevocationDistributionPoint transaction should be signed by an account with the \"%s\" role",
				dclauthtypes.Vendor,
			)
		}

		if msg.Vid != signerAccount.VendorID {
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized,
				"MsgAddPkiRevocationDistributionPoint signer must have the same vid as provided in message",
			)
		}

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
		IsPAA:                *msg.IsPAA,
		CrlSignerCertificate: msg.CrlSignerCertificate,
		DataUrl:              msg.DataUrl,
		DataFileSize:         msg.DataFileSize,
		DataDigest:           msg.DataDigest,
		DataDigestType:       msg.DataDigestType,
		RevocationType:       msg.RevocationType,
	}

	k.SetPkiRevocationDistributionPoint(ctx, pkiRevocationDistributionPoint)

	return &types.MsgAddPkiRevocationDistributionPointResponse{}, nil
}
