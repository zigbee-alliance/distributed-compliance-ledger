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

	"git.dsr-corporation.com/zb-ledger/zb-ledger/utils/cli"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/validator/internal/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/spf13/cobra"
	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
	cfg "github.com/tendermint/tendermint/config"
	"github.com/tendermint/tendermint/crypto"
)

// GetTxCmd returns the transaction commands for this module.
func GetTxCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	validatorTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "validator transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	validatorTxCmd.AddCommand(flags.PostCommands(
		GetCmdCreateValidator(cdc),
	)...)

	return validatorTxCmd
}

// GetCmdCreateValidator implements the create validator command handler.
func GetCmdCreateValidator(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-node",
		Short: "Adds a new validator node",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := cli.NewCLIContext().WithCodec(cdc)

			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			txBldr, msg, err := BuildCreateValidatorMsg(cliCtx.Context(), txBldr, false)
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx.Context(), txBldr, []sdk.Msg{msg})
		},
	}

	fsCreateValidator := InitValidatorFlags()
	cmd.Flags().AddFlagSet(fsCreateValidator)

	_ = cmd.MarkFlagRequired(flags.FlagFrom)
	_ = cmd.MarkFlagRequired(FlagAddress)
	_ = cmd.MarkFlagRequired(FlagPubKey)
	_ = cmd.MarkFlagRequired(FlagName)

	return cmd
}

// Return the flagset for create validator command.
func InitValidatorFlags() (fs *flag.FlagSet) {
	fsCreateValidator := flag.NewFlagSet("", flag.ContinueOnError)
	fsCreateValidator.String(FlagAddress, "", "The Bech32 encoded Address of the validator")
	fsCreateValidator.String(FlagPubKey, "", "The Bech32 encoded ConsensusPubkey of the validator")
	fsCreateValidator.String(FlagName, "", "The validator's name")
	fsCreateValidator.String(FlagWebsite, "", "The validator's (optional) website")
	fsCreateValidator.String(FlagDetails, "", "The validator's (optional) details")
	fsCreateValidator.String(FlagIdentity, "", "The (optional) identity signature (ex. UPort or Keybase)")

	return fsCreateValidator
}

// prepare flags in config.
func PrepareFlagsForTxCreateValidator(
	config *cfg.Config, nodeID, chainID string, valPubKey crypto.PubKey,
) {
	viper.Set(flags.FlagChainID, chainID)
	viper.Set(FlagNodeID, nodeID)
	viper.Set(FlagAddress, sdk.ConsAddress(valPubKey.Address()).String())
	viper.Set(FlagPubKey, sdk.MustBech32ifyConsPub(valPubKey))
	viper.Set(FlagName, config.Moniker)

	if config.Moniker == "" {
		viper.Set(FlagName, viper.GetString(flags.FlagName))
	}
}

// BuildCreateValidatorMsg makes a new MsgCreateValidator.
func BuildCreateValidatorMsg(cliCtx context.CLIContext, txBldr auth.TxBuilder,
	isGenesis bool) (auth.TxBuilder, sdk.Msg, error) {
	signer := cliCtx.GetFromAddress()

	address, err := sdk.ConsAddressFromBech32(viper.GetString(FlagAddress))
	if err != nil {
		return txBldr, nil, err
	}

	pubkey := viper.GetString(FlagPubKey)

	description := types.NewDescription(
		viper.GetString(FlagName),
		viper.GetString(FlagIdentity),
		viper.GetString(FlagWebsite),
		viper.GetString(FlagDetails),
	)

	msg := types.NewMsgCreateValidator(address, pubkey, description, signer)

	if isGenesis {
		ip := viper.GetString(FlagIP)
		nodeID := viper.GetString(FlagNodeID)

		if nodeID != "" && ip != "" {
			txBldr = txBldr.WithMemo(fmt.Sprintf("%s@%s:26656", nodeID, ip))
		}
	}

	err = msg.ValidateBasic()
	if err != nil {
		return txBldr, nil, err
	}

	return txBldr, msg, nil
}
