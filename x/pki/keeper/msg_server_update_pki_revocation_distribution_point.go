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

func (k msgServer) UpdatePkiRevocationDistributionPoint(goCtx context.Context, msg *types.MsgUpdatePkiRevocationDistributionPoint) (*types.MsgUpdatePkiRevocationDistributionPointResponse, error) {
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
		if msg.CrlSignerCertificate != "" {
			updatedCrlSignerCertificate, err := x509.DecodeX509Certificate(msg.CrlSignerCertificate)
			if err != nil {
				return nil, pkitypes.NewErrInvalidCertificate(err)
			}

			if !updatedCrlSignerCertificate.IsSelfSigned() {
				return nil, pkitypes.NewErrRootCertificateIsNotSelfSigned("Updated CRL signer certificate must be self-signed since old one was self-signed")
			}
		}

		subjectAsMap := x509.SubjectAsTextToMap(crlSignerCertificate.SubjectAsText)

		strVid, found := subjectAsMap["Mvid"]

		if !found {
			return nil, pkitypes.NewErrVidNotFound("vid must be encoded in Revocation Distribution Point's Signer Certificate")
		}

		_, err := strconv.ParseInt(strings.Trim(strVid, "0x"), 16, 32)
		if err != nil {
			return nil, err
		}

		// check if signer has vendor role
		if !k.dclauthKeeper.HasRole(ctx, signerAddr, dclauthtypes.Vendor) {
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized,
				"MsgUpdatePkiRevocationDistributionPoint transaction should be signed by an account with the \"%s\" role",
				dclauthtypes.Vendor,
			)
		}

		if signerAccount.VendorID != msg.Vid {
			return nil, pkitypes.NewErrCRLSignerCertificateVidNotEqualAccountVid("CRL signer Certificate vid must equal to signer account vid")
		}

		if pkiRevocationDistributionPoint.CrlSignerCertificate != "" {
			subjectAsMap := x509.SubjectAsTextToMap(crlSignerCertificate.SubjectAsText)

			strVid, found := subjectAsMap["Mvid"]
			if found {
				vid, err := strconv.ParseInt(strings.Trim(strVid, "0x"), 16, 32)
				if err != nil {
					return nil, err
				}

				if int32(vid) != msg.Vid {
					return nil, pkitypes.NewErrCRLSignerCertificateVidNotEqualMsgVid("CRL signer Certificate vid must equal to message vid")
				}
			}
		}
	} else {
		// check if signer has vendor role
		if !k.dclauthKeeper.HasRole(ctx, signerAddr, dclauthtypes.Vendor) {
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized,
				"MsgUpdatePkiRevocationDistributionPoint transaction should be signed by an account with the \"%s\" role",
				dclauthtypes.Vendor,
			)
		}

		if msg.Vid != signerAccount.VendorID {
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized,
				"MsgUpdatePkiRevocationDistributionPoint signer must have the same vid as provided in message",
			)
		}

		if msg.CrlSignerCertificate != "" {
			subjectAsMap := x509.SubjectAsTextToMap(msg.CrlSignerCertificate)

			strVid, found := subjectAsMap["Mvid"]
			if found {
				vid, err := strconv.ParseInt(strings.Trim(strVid, "0x"), 16, 32)
				if err != nil {
					return nil, err
				}

				if int32(vid) != msg.Vid {
					return nil, pkitypes.NewErrCRLSignerCertificateVidNotEqualMsgVid("CRL signer Certificate vid must equal to message vid")
				}
			}

			strPid, found := subjectAsMap["Mpid"]
			if found {
				pid, err := strconv.ParseInt(strings.Trim(strPid, "0x"), 16, 32)
				if err != nil {
					return nil, err
				}

				if int32(pid) != pkiRevocationDistributionPoint.Pid {
					return nil, pkitypes.NewErrCRLSignerCertificateVidNotEqualMsgVid("CRL signer Certificate pid must equal to message pid")
				}
			}

			if pkiRevocationDistributionPoint.Pid != 0 {
				strPid, found := subjectAsMap["Mpid"]
				if found {
					pid, err := strconv.ParseInt(strings.Trim(strPid, "0x"), 16, 32)
					if err != nil {
						return nil, err
					}

					if int32(pid) != pkiRevocationDistributionPoint.Pid {
						return nil, err
					}
				} else {
					return nil, pkitypes.NewErrPidNotFound("pid must be encoded in updated CRL signer certificate since pid fields presented in PKI revocation distribution point")
				}
			}
		}
	}

	if msg.CrlSignerCertificate != "" {
		pkiRevocationDistributionPoint.CrlSignerCertificate = msg.CrlSignerCertificate
	}

	if msg.DataUrl != "" {
		pkiRevocationDistributionPoint.DataUrl = msg.DataUrl
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

	return &types.MsgUpdatePkiRevocationDistributionPointResponse{}, nil
}
