package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/x509"
)

func (k msgServer) AddPkiRevocationDistributionPoint(goCtx context.Context, msg *types.MsgAddPkiRevocationDistributionPoint) (*types.MsgAddPkiRevocationDistributionPointResponse, error) {
	_ = sdk.UnwrapSDKContext(goCtx)

	x509Certificate, err := x509.DecodeX509Certificate(msg.CrlSignerCertificate)
	if err != nil {
		return nil, pkitypes.NewErrInvalidCertificate(err)
	}

	if x509Certificate.IsSelfSigned() {
		return nil, pkitypes.NewErrInappropriateCertificateType(
			"Inappropriate Certificate Type: Passed certificate is self-signed, " +
				"so it cannot be added to the system as a non-root certificate. " +
				"To propose adding a root certificate please use `PROPOSE_ADD_X509_ROOT_CERT` transaction.")
	}

	return &types.MsgAddPkiRevocationDistributionPointResponse{}, nil
}
