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
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/cli"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/conversions"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/auth/internal/types"
)

func GetTxCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	authTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Authorization subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	authTxCmd.AddCommand(cli.SignedCommands(client.PostCommands(
		GetCmdProposeAddAccount(cdc),
		GetCmdApproveAddAccount(cdc),
		GetCmdProposeRevokeAccount(cdc),
		GetCmdApproveRevokeAccount(cdc),
	)...)...)

	return authTxCmd
}

func GetCmdProposeAddAccount(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "propose-add-account",
		Short: "Propose a new account with the given address, public key and roles",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := cli.NewCLIContext().WithCodec(cdc)

			address, err := sdk.AccAddressFromBech32(viper.GetString(FlagAddress))
			if err != nil {
				return err
			}

			pubkey := viper.GetString(FlagPubKey)
			_, err = sdk.GetAccPubKeyBech32(pubkey)
			if err != nil {
				return err
			}

			var roles types.AccountRoles
			if rolesStr := viper.GetString(FlagRoles); len(rolesStr) > 0 {
				for _, role := range strings.Split(rolesStr, ",") {
					roles = append(roles, types.AccountRole(role))
				}
			}

			var vendorId uint16
			if viper.GetString(FlagVID) != "" {
				var err_ sdk.Error
				vendorId, err_ = conversions.ParseVID(viper.GetString(FlagVID))
				if err_ != nil {
					return err_
				}
			}
			fmt.Println(vendorId)

			msg := types.NewMsgProposeAddAccount(address, pubkey, roles, vendorId, cliCtx.FromAddress())

			return cliCtx.HandleWriteMessage(msg)
		},
	}

	cmd.Flags().String(FlagAddress, "", "Bench32 encoded account address")
	cmd.Flags().String(FlagPubKey, "", "Bench32 encoded account public key")
	cmd.Flags().String(FlagRoles, "",
		fmt.Sprintf("The list of roles, comma-separated, assigning to the account (supported roles: %v)",
			types.Roles))
	cmd.Flags().String(FlagVID, "", "Vendor ID associated with this account. Required only for Vendor Roles")

	_ = cmd.MarkFlagRequired(FlagAddress)
	_ = cmd.MarkFlagRequired(FlagPubKey)

	return cmd
}

func GetCmdApproveAddAccount(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "approve-add-account",
		Short: "Approve the proposed account with the given address",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := cli.NewCLIContext().WithCodec(cdc)

			address, err := sdk.AccAddressFromBech32(viper.GetString(FlagAddress))
			if err != nil {
				return err
			}

			msg := types.NewMsgApproveAddAccount(address, cliCtx.FromAddress())

			return cliCtx.HandleWriteMessage(msg)
		},
	}

	cmd.Flags().String(FlagAddress, "", "Bench32 encoded account address")

	_ = cmd.MarkFlagRequired(FlagAddress)

	return cmd
}

func GetCmdProposeRevokeAccount(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "propose-revoke-account",
		Short: "Propose revocation of the account with the given address",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := cli.NewCLIContext().WithCodec(cdc)

			address, err := sdk.AccAddressFromBech32(viper.GetString(FlagAddress))
			if err != nil {
				return err
			}

			msg := types.NewMsgProposeRevokeAccount(address, cliCtx.FromAddress())

			return cliCtx.HandleWriteMessage(msg)
		},
	}

	cmd.Flags().String(FlagAddress, "", "Bench32 encoded account address")

	_ = cmd.MarkFlagRequired(FlagAddress)

	return cmd
}

func GetCmdApproveRevokeAccount(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "approve-revoke-account",
		Short: "Approve the proposed revocation of the account with the given address",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := cli.NewCLIContext().WithCodec(cdc)

			address, err := sdk.AccAddressFromBech32(viper.GetString(FlagAddress))
			if err != nil {
				return err
			}

			msg := types.NewMsgApproveRevokeAccount(address, cliCtx.FromAddress())

			return cliCtx.HandleWriteMessage(msg)
		},
	}

	cmd.Flags().String(FlagAddress, "", "Bench32 encoded account address")

	_ = cmd.MarkFlagRequired(FlagAddress)

	return cmd
}
