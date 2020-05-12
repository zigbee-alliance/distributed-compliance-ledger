package cli

import (
	"fmt"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/utils/cli"
	"os"

	"github.com/spf13/cobra"
	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"

	cfg "github.com/tendermint/tendermint/config"
	"github.com/tendermint/tendermint/crypto"

	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/validator/internal/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
)

// GetTxCmd returns the transaction commands for this module
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

			txBldr, msg, err := BuildCreateValidatorMsg(cliCtx.Context(), txBldr)
			if err != nil {
				return err
			}

			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx.Context(), txBldr, []sdk.Msg{msg})
		},
	}

	cmd.Flags().String(FlagName, "", "The validator's name")

	fsCreateValidator := InitValidatorFlags("")
	cmd.Flags().AddFlagSet(fsCreateValidator)

	cmd.MarkFlagRequired(flags.FlagFrom)
	cmd.MarkFlagRequired(FlagAddress)
	cmd.MarkFlagRequired(FlagPubKey)
	cmd.MarkFlagRequired(FlagName)

	return cmd
}

// Return the flagset for create validator command
func InitValidatorFlags(ipDefault string) (fs *flag.FlagSet) {
	fsCreateValidator := flag.NewFlagSet("", flag.ContinueOnError)
	fsCreateValidator.String(FlagAddress, "", "The Bech32 encoded Address of the validator")
	fsCreateValidator.String(FlagPubKey, "", "The Bech32 encoded ConsensusPubkey of the validator")
	fsCreateValidator.String(FlagIP, ipDefault, "The node's public IP")
	fsCreateValidator.String(FlagNodeID, "", "The node's NodeID")
	fsCreateValidator.String(FlagWebsite, "", "The validator's (optional) website")
	fsCreateValidator.String(FlagDetails, "", "The validator's (optional) details")
	fsCreateValidator.String(FlagIdentity, "", "The (optional) identity signature (ex. UPort or Keybase)")
	return fsCreateValidator
}

// prepare flags in config
func PrepareFlagsForTxCreateValidator(
	config *cfg.Config, nodeID, chainID string, valPubKey crypto.PubKey,
) {

	ip := viper.GetString(FlagIP)
	if ip == "" {
		fmt.Fprintf(os.Stderr, "couldn't retrieve an external IP; "+
			"the tx's memo field will be unset")
	}

	website := viper.GetString(FlagWebsite)
	details := viper.GetString(FlagDetails)
	identity := viper.GetString(FlagIdentity)

	viper.Set(flags.FlagChainID, chainID)
	viper.Set(flags.FlagFrom, viper.GetString(flags.FlagName))
	viper.Set(FlagNodeID, nodeID)
	viper.Set(FlagIP, ip)
	viper.Set(FlagAddress, sdk.ConsAddress(valPubKey.Address()).String())
	viper.Set(FlagPubKey, sdk.MustBech32ifyConsPub(valPubKey))
	viper.Set(FlagName, config.Moniker)
	viper.Set(FlagWebsite, website)
	viper.Set(FlagDetails, details)
	viper.Set(FlagIdentity, identity)

	if config.Moniker == "" {
		viper.Set(FlagName, viper.GetString(flags.FlagName))
	}
}

// BuildCreateValidatorMsg makes a new MsgCreateValidator.
func BuildCreateValidatorMsg(cliCtx context.CLIContext, txBldr auth.TxBuilder) (auth.TxBuilder, sdk.Msg, error) {

	signer := cliCtx.GetFromAddress()

	valConsAddr, err := sdk.ConsAddressFromBech32(viper.GetString(FlagAddress))
	if err != nil {
		return txBldr, nil, err
	}

	valConsPubkeyStr := viper.GetString(FlagPubKey)
	valConsPubkey, err := sdk.GetConsPubKeyBech32(valConsPubkeyStr)
	if err != nil {
		return txBldr, nil, err
	}

	if !valConsAddr.Equals(sdk.ConsAddress(valConsPubkey.Address())) {
		return txBldr, nil, sdk.ErrUnknownRequest("Validator Consensus Pubkey does not match to Validator Address.")
	}

	description := types.NewDescription(
		viper.GetString(FlagName),
		viper.GetString(FlagIdentity),
		viper.GetString(FlagWebsite),
		viper.GetString(FlagDetails),
	)

	msg := types.NewMsgCreateValidator(valConsAddr, valConsPubkeyStr, description, signer)

	if viper.GetBool(client.FlagGenerateOnly) {
		ip := viper.GetString(FlagIP)
		nodeID := viper.GetString(FlagNodeID)
		if nodeID != "" && ip != "" {
			txBldr = txBldr.WithMemo(fmt.Sprintf("%s@%s:26656", nodeID, ip))
		}
	}

	return txBldr, msg, nil
}
