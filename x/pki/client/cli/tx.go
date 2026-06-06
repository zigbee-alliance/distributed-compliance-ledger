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
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"
	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	// "github.com/cosmos/cosmos-sdk/client/flags".
)

var DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())

// GetTxCmd returns the transaction commands for this module.
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        pkitypes.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", pkitypes.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdProposeAddX509RootCert())
	cmd.AddCommand(CmdApproveAddX509RootCert())
	cmd.AddCommand(CmdAddX509Cert())
	cmd.AddCommand(CmdProposeRevokeX509RootCert())
	cmd.AddCommand(CmdApproveRevokeX509RootCert())
	cmd.AddCommand(CmdRevokeX509Cert())
	cmd.AddCommand(CmdRejectAddX509RootCert())
	cmd.AddCommand(CmdAddPkiRevocationDistributionPoint())
	cmd.AddCommand(CmdUpdatePkiRevocationDistributionPoint())
	cmd.AddCommand(CmdDeletePkiRevocationDistributionPoint())
	cmd.AddCommand(CmdAssignVid())
	cmd.AddCommand(CmdAddNocX509RootCert())
	cmd.AddCommand(CmdRemoveX509Cert())
	cmd.AddCommand(CmdAddNocX509IcaCert())
	cmd.AddCommand(CmdRevokeNocX509RootCert())
	cmd.AddCommand(CmdRevokeNocX509IcaCert())
	cmd.AddCommand(CmdRemoveNocX509IcaCert())
	cmd.AddCommand(CmdRemoveNocX509RootCert())
	// this line is used by starport scaffolding # 1

	return cmd
}
