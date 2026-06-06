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
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/cli"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

var _ = strconv.Itoa(0)

func CmdRevokeNocX509IcaCert() *cobra.Command {
	cmd := &cobra.Command{
		Use: "revoke-noc-x509-ica-cert",
		Short: "Revokes the given NOC intermediate or leaf certificate (ICAC). " +
			"If revoke-child flag is set to true then all the certificates in the subtree signed by the revoked " +
			"certificate will be revoked as well.",
		Args: cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			subject := viper.GetString(FlagSubject)
			subjectKeyID := viper.GetString(FlagSubjectKeyID)
			serialNumber := viper.GetString(FlagSerialNumber)
			revokeChild := viper.GetBool(FlagRevokeChild)
			infoArg := viper.GetString(FlagInfo)

			msg := types.NewMsgRevokeNocX509IcaCert(
				clientCtx.GetFromAddress().String(),
				subject,
				subjectKeyID,
				serialNumber,
				infoArg,
				revokeChild,
			)
			// validate basic will be called in GenerateOrBroadcastTxCLI
			err = tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
			if cli.IsWriteInsteadReadRPCError(err) {
				return clientCtx.PrintString(cli.LightClientProxyForWriteRequests)
			}

			return err
		},
	}

	cmd.Flags().StringP(FlagSubject, FlagSubjectShortcut, "", "Certificate's subject")
	cmd.Flags().StringP(FlagSubjectKeyID, FlagSubjectKeyIDShortcut, "", "Certificate's subject key id (hex)")
	cmd.Flags().StringP(FlagSerialNumber, FlagSerialNumberShortcut, "", "Certificate's serial number")
	cmd.Flags().StringP(FlagRevokeChild, FlagRevokeChildShortcut, "", "If flag is true then all the certificates in the subtree will be revoked as well - default is false")
	cmd.Flags().String(FlagInfo, "", FlagInfoUsage)

	cli.AddTxFlagsToCmd(cmd)

	_ = cmd.MarkFlagRequired(FlagSubject)
	_ = cmd.MarkFlagRequired(FlagSubjectKeyID)
	_ = cmd.MarkFlagRequired(flags.FlagFrom)

	return cmd
}
