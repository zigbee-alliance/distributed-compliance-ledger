package cli

import (
	"flag"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/validator/types"
)

var _ = strconv.Itoa(0)

func CmdCreateValidator() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-validator",
		Short: "create new validator",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			txf := tx.NewFactoryCLI(clientCtx, cmd.Flags()).
				WithTxConfig(clientCtx.TxConfig).WithAccountRetriever(clientCtx.AccountRetriever)
			txf, msg, err := newBuildCreateValidatorMsg(clientCtx, txf, cmd.Flags())
			if err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	fsCreateValidator := InitValidatorFlags()
	cmd.Flags().AddFlagSet(fsCreateValidator)

	cmd.Flags().String(FlagIP, "", fmt.Sprintf("The node's public IP. It takes effect only when used in combination with --%s", flags.FlagGenerateOnly))
	cmd.Flags().String(FlagNodeID, "", "The node's ID")
	flags.AddTxFlagsToCmd(cmd)

	_ = cmd.MarkFlagRequired(flags.FlagFrom)
	_ = cmd.MarkFlagRequired(FlagPubKey)
	_ = cmd.MarkFlagRequired(FlagName)

	return cmd
}

func InitValidatorFlags() (fs *flag.FlagSet) {
	fsCreateValidator := flag.NewFlagSet("", flag.ContinueOnError)
	fsCreateValidator.String(FlagPubKey, "", "The validator's Protobuf JSON encoded public key")
	fsCreateValidator.String(FlagName, "", "The validator's name")
	fsCreateValidator.String(FlagWebsite, "", "The validator's (optional) website")
	fsCreateValidator.String(FlagDetails, "", "The validator's (optional) details")
	fsCreateValidator.String(FlagIdentity, "", "The (optional) identity signature (ex. UPort or Keybase)")

	return fsCreateValidator
}

func newBuildCreateValidatorMsg(clientCtx client.Context, txf tx.Factory, fs *flag.FlagSet) (tx.Factory, *types.MsgCreateValidator, error) {
	valAddr := clientCtx.GetFromAddress()
	pkStr, err := fs.GetString(FlagPubKey)
	if err != nil {
		return txf, nil, err
	}

	var pk cryptotypes.PubKey
	if err := clientCtx.Codec.UnmarshalInterfaceJSON([]byte(pkStr), &pk); err != nil {
		return txf, nil, err
	}

	name, _ := fs.GetString(FlagName)
	identity, _ := fs.GetString(FlagIdentity)
	website, _ := fs.GetString(FlagWebsite)
	details, _ := fs.GetString(FlagDetails)
	description := types.NewDescription(
		name,
		identity,
		website,
		security,
		details,
	)

	msg, err := types.NewMsgCreateValidator(valAddr, pk, description)
	if err != nil {
		return txf, nil, err
	}
	if err := msg.ValidateBasic(); err != nil {
		return txf, nil, err
	}

	genOnly, _ := fs.GetBool(flags.FlagGenerateOnly)
	if genOnly {
		ip, _ := fs.GetString(FlagIP)
		nodeID, _ := fs.GetString(FlagNodeID)

		if nodeID != "" && ip != "" {
			txf = txf.WithMemo(fmt.Sprintf("%s@%s:26656", nodeID, ip))
		}
	}

	return txf, msg, nil
}

/* FIXME issue 99, from old DCL
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
*/

/* FIXME issue 99, from cosmos
func CreateValidatorMsgFlagSet(ipDefault string) (fs *flag.FlagSet, defaultsDesc string) {
	fsCreateValidator := flag.NewFlagSet("", flag.ContinueOnError)
	fsCreateValidator.String(FlagIP, ipDefault, "The node's public IP")
	fsCreateValidator.String(FlagNodeID, "", "The node's NodeID")
	fsCreateValidator.String(FlagName, "", "The validator's (optional) name")
	fsCreateValidator.String(FlagWebsite, "", "The validator's (optional) website")
	fsCreateValidator.String(FlagDetails, "", "The validator's (optional) details")
	fsCreateValidator.String(FlagIdentity, "", "The (optional) identity signature (ex. UPort or Keybase)")
	fsCreateValidator.AddFlagSet(FlagSetPublicKey())

	return fsCreateValidator, defaultsDesc
}
*/
