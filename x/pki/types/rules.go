package types

import "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"

// nolint:godox
// TODO: 1. Move it to separate module  	2. Make it configurable		3. Save into store.
var (
	RootCertificateApprovals    = 2
	RootCertificateApprovalRole = types.Trustee
)
