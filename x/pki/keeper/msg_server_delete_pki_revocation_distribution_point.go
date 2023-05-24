package keeper

import (
	"context"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	dclauthtypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/x509"
)

func (k msgServer) DeletePkiRevocationDistributionPoint(goCtx context.Context, msg *types.MsgDeletePkiRevocationDistributionPoint) (*types.MsgDeletePkiRevocationDistributionPointResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	pkiRevocationDistributionPoint, isFound := k.GetPkiRevocationDistributionPoint(ctx, msg.Vid, msg.Label, msg.IssuerSubjectKeyID)
	if !isFound {
		return nil, pkitypes.NewErrPkiRevocationDistributionPointDoesNotExists("PKI revocation distribution point does not exist")
	}

	crlSignerCertificate, err := x509.DecodeX509Certificate(pkiRevocationDistributionPoint.CrlSignerCertificate)
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

		strVid, found := subjectAsMap["vid"]
		if found {
			_, err := strconv.ParseInt(strVid, 10, 32)
			if err != nil {
				return nil, err
			}

			// check if signer has vendor role
			if !k.dclauthKeeper.HasRole(ctx, signerAddr, dclauthtypes.Vendor) {
				return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized,
					"MsgDeletePkiRevocationDistributionPoint transaction should be signed by an account with the \"%s\" role",
					dclauthtypes.Vendor,
				)
			}

			if signerAccount.VendorID != msg.Vid {
				return nil, pkitypes.NewErrCRLSignerCertificateVidNotEqualAccountVid("CRL signer Certificate vid must equal to signer account vid")
			}
		} else {
			// check if signer has vendor admin role
			if !k.dclauthKeeper.HasRole(ctx, signerAddr, dclauthtypes.VendorAdmin) {
				return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized,
					"MsgDeletePkiRevocationDistributionPoint transaction should be signed by an account with the \"%s\" role",
					dclauthtypes.VendorAdmin,
				)
			}
		}
	} else {
		// check if signer has vendor role
		if !k.dclauthKeeper.HasRole(ctx, signerAddr, dclauthtypes.Vendor) {
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized,
				"MsgDeletePkiRevocationDistributionPoint transaction should be signed by an account with the \"%s\" role",
				dclauthtypes.Vendor,
			)
		}

		if signerAccount.VendorID != msg.Vid {
			return nil, pkitypes.NewErrCRLSignerCertificateVidNotEqualAccountVid("CRL signer Certificate vid must equal to signer account vid")
		}

		subjectAsMap := x509.SubjectAsTextToMap(crlSignerCertificate.SubjectAsText)

		strVid, found := subjectAsMap["vid"]
		if found {
			vid, err := strconv.ParseInt(strVid, 10, 32)
			if err != nil {
				return nil, err
			}

			if int32(vid) != msg.Vid {
				return nil, pkitypes.NewErrCRLSignerCertificateVidNotEqualMsgVid("CRL signer Certificate vid must equal to message vid")
			}
		}
	}

	k.RemovePkiRevocationDistributionPoint(ctx, msg.Vid, msg.Label, msg.IssuerSubjectKeyID)

	return &types.MsgDeletePkiRevocationDistributionPointResponse{}, nil
}
