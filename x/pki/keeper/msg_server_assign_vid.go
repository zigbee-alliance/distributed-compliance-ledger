package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	dclauthtypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/x509"
)

func (k msgServer) AssignVid(goCtx context.Context, msg *types.MsgAssignVid) (*types.MsgAssignVidResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	signerAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid Address: (%s)", err)
	}

	// check if signer has vendor admin role
	if !k.dclauthKeeper.HasRole(ctx, signerAddr, dclauthtypes.VendorAdmin) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized,
			"AssignVid transaction should be signed by an account with the \"%s\" role", dclauthtypes.VendorAdmin,
		)
	}

	// get corresponding approved certificates
	certificates, found := k.GetApprovedCertificates(ctx, msg.Subject, msg.SubjectKeyId)
	if !found {
		return nil, pkitypes.NewErrCertificateDoesNotExist(msg.Subject, msg.SubjectKeyId)
	}

	certificate := certificates.Certs[0]

	// fail if certificates are not root
	if !certificate.IsRoot {
		return nil, pkitypes.NewErrInappropriateCertificateType(
			fmt.Sprintf("Inappropriate Certificate Type: Certificate with subject=%v and subjectKeyID=%v "+
				"is not a root certificate.", msg.Subject, msg.SubjectKeyId),
		)
	}

	// check that the certificate VID and Message VID are equal
	subjectVid, err := x509.GetVidFromSubject(certificate.SubjectAsText)
	if err != nil {
		return nil, pkitypes.NewErrInvalidCertificate(err)
	}
	if subjectVid != 0 && subjectVid != msg.Vid {
		return nil, pkitypes.NewErrCertificateVidNotEqualMsgVid(fmt.Sprintf("Certificate VID=%d is not equal to the msg VID=%d", subjectVid, msg.Vid))
	}

	hasCertificateWithoutVid := false

	// assign the VID to certificates that don't have it
	for _, certificate := range certificates.Certs {
		if certificate.Vid != 0 {
			continue
		}

		hasCertificateWithoutVid = true

		certificate.Vid = msg.Vid
	}

	// check that the VID has been set for at least one certificate
	if !hasCertificateWithoutVid {
		return nil, pkitypes.NewErrNotEmptyVid("Vendor ID (VID) already present in certificates")
	}

	k.SetApprovedCertificates(ctx, certificates)

	return &types.MsgAssignVidResponse{}, nil
}
