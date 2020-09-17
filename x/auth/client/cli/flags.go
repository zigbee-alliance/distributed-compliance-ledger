// Copyright 2020 DSR Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cli

import (
	"fmt"

	"github.com/zigbee-alliance/distributed-compliance-ledger/x/auth/internal/types"
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
