package types

import "git.dsr-corporation.com/zb-ledger/zb-ledger/x/authz"

// TODO:
// 	1. Move it to separate module  	2. Make it configurable		3. Save into store
var (
	RootCertificateApprovals    = 2
	RootCertificateApprovalRole = authz.Trustee
)
