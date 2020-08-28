package cli

import (
	"fmt"

	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/auth/internal/types"
)

const (
	FlagAddress      = "address"
	FlagAddressUsage = "Bench32 encoded account address"
	FlagPubKey       = "pubkey"
	FlagPubKeyUsage  = "Bench32 encoded account public key"
	FlagRoles        = "roles"
)

var FlagRolesUsage = fmt.Sprintf("The list of roles, comma-separated, assigning to the account "+
	"(supported roles: %v)", types.Roles)
