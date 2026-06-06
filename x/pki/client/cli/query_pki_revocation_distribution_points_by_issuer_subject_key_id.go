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
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/cli"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

func CmdShowPkiRevocationDistributionPointsByIssuerSubjectKeyID() *cobra.Command {
	var issuerSubjectKeyID string

	cmd := &cobra.Command{
		Use:   "revocation-points",
		Short: "Gets all revocation points associated with issuer's subject key id",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			var res types.PkiRevocationDistributionPointsByIssuerSubjectKeyID

			return cli.QueryWithProof(
				clientCtx,
				pkitypes.StoreKey,
				types.PkiRevocationDistributionPointsByIssuerSubjectKeyIDKeyPrefix,
				types.PkiRevocationDistributionPointsByIssuerSubjectKeyIDKey(issuerSubjectKeyID),
				&res,
			)
		},
	}

	cmd.Flags().StringVar(&issuerSubjectKeyID, FlagIssuerSubjectKeyID, "", "Issuer's subject key id")
	flags.AddQueryFlagsToCmd(cmd)

	_ = cmd.MarkFlagRequired(FlagIssuerSubjectKeyID)

	return cmd
}
