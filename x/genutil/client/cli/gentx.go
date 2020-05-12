package cli

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/tendermint/tendermint/libs/common"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/pkg/errors"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	tmtypes "github.com/tendermint/tendermint/types"

	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/genutil"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/genutil/types"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	kbkeys "github.com/cosmos/cosmos-sdk/crypto/keys"
	"github.com/cosmos/cosmos-sdk/server"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"

	validator "git.dsr-corporation.com/zb-ledger/zb-ledger/x/validator/client/cli"
)

// GenTxCmd builds the application's gentx command.
// nolint: errcheck
func GenTxCmd(ctx *server.Context, cdc *codec.Codec, mbm module.BasicManager,
	genAccIterator types.GenesisAccountsIterator, defaultNodeHome, defaultCLIHome string) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "gentx",
		Short: "Generate a genesis transaction to create a validator",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {

			config := ctx.Config
			config.SetRoot(viper.GetString(client.FlagHome))

			nodeID, valPubKey, err := genutil.InitializeNodeValidatorFiles(ctx.Config)
			if err != nil {
				return err
			}

			// Read --nodeID, if empty take it from priv_validator.json
			if nodeIDString := viper.GetString(validator.FlagNodeID); nodeIDString != "" {
				nodeID = nodeIDString
			}
			// Read --pubkey, if empty take it from priv_validator.json
			if valPubKeyString := viper.GetString(validator.FlagPubKey); valPubKeyString != "" {
				valPubKey, err = sdk.GetConsPubKeyBech32(valPubKeyString)
				if err != nil {
					return err
				}
			}

			genDoc, err := tmtypes.GenesisDocFromFile(config.GenesisFile())
			if err != nil {
				return err
			}

			var genesisState map[string]json.RawMessage
			if err = cdc.UnmarshalJSON(genDoc.AppState, &genesisState); err != nil {
				return err
			}

			if err = mbm.ValidateGenesis(genesisState); err != nil {
				return err
			}

			kb, err := client.NewKeyBaseFromDir(viper.GetString(flagClientHome))
			if err != nil {
				return err
			}

			from := viper.GetString(client.FlagFrom)
			key, err := kb.Get(from)
			if err != nil {
				return err
			}

			// Set flags for creating gentx
			viper.Set(client.FlagHome, viper.GetString(flagClientHome))
			validator.PrepareFlagsForTxCreateValidator(config, nodeID, genDoc.ChainID, valPubKey)

			err = genutil.ValidateAccountInGenesis(genesisState, genAccIterator, key.GetAddress(), cdc)
			if err != nil {
				return err
			}

			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := client.NewCLIContext().WithCodec(cdc)

			// Set the generate-only flag here after the CLI context has
			// been created. This allows the from from/key to be correctly populated.
			viper.Set(client.FlagGenerateOnly, true)

			// create a 'create-validator' message
			txBldr, msg, err := validator.BuildCreateValidatorMsg(cliCtx, txBldr)
			if err != nil {
				return errors.Wrap(err, "failed to build create-validator message")
			}

			info, err := txBldr.Keybase().Get(from)
			if err != nil {
				return errors.Wrap(err, "failed to read from tx builder keybase")
			}

			if info.GetType() == kbkeys.TypeOffline || info.GetType() == kbkeys.TypeMulti {
				fmt.Println("Offline key passed in. Use `tx sign` command to sign:")
				return utils.PrintUnsignedStdTx(txBldr, cliCtx, []sdk.Msg{msg})
			}

			// write the unsigned transaction to the buffer
			w := bytes.NewBuffer([]byte{})
			cliCtx = cliCtx.WithOutput(w)

			if err = utils.PrintUnsignedStdTx(txBldr, cliCtx, []sdk.Msg{msg}); err != nil {
				return errors.Wrap(err, "failed to print unsigned std tx")
			}

			// read the transaction
			stdTx, err := readUnsignedGenTxFile(cdc, w)
			if err != nil {
				return errors.Wrap(err, "failed to read unsigned gen tx file")
			}

			// sign the transaction and write it to the output file
			signedTx, err := utils.SignStdTx(txBldr, cliCtx, from, stdTx, false, true)
			if err != nil {
				return errors.Wrap(err, "failed to sign std tx")
			}

			// Fetch output file from
			outputDocument := viper.GetString(flags.FlagOutputDocument)
			if outputDocument == "" {
				outputDocument, err = makeOutputFilepath(config.RootDir, nodeID)
				if err != nil {
					return errors.Wrap(err, "failed to create output file path")
				}
			}

			if err := writeSignedGenTx(cdc, outputDocument, signedTx); err != nil {
				return errors.Wrap(err, "failed to write signed gen tx")
			}

			fmt.Fprintf(os.Stderr, "Genesis transaction written to %q\n", outputDocument)
			return nil

		},
	}

	cmd.Flags().String(flags.FlagHome, defaultNodeHome, "node's home directory")
	cmd.Flags().String(flagClientHome, defaultCLIHome, "client's home directory")
	cmd.Flags().String(flags.FlagFrom, "", "Name or address of private key with which to sign the gentx")
	cmd.Flags().String(flags.FlagOutputDocument, "",
		"write the genesis transaction JSON document to the given file instead of the default location")

	ipDefault, _ := server.ExternalIP()
	cmd.Flags().AddFlagSet(validator.InitValidatorFlags(ipDefault))

	cmd.MarkFlagRequired(flags.FlagFrom)
	return cmd
}

func makeOutputFilepath(rootDir, nodeID string) (string, error) {
	writePath := filepath.Join(rootDir, "config", "gentx")
	if err := common.EnsureDir(writePath, 0700); err != nil {
		return "", err
	}
	return filepath.Join(writePath, fmt.Sprintf("gentx-%v.json", nodeID)), nil
}

func readUnsignedGenTxFile(cdc *codec.Codec, r io.Reader) (auth.StdTx, error) {
	var stdTx auth.StdTx
	bytes, err := ioutil.ReadAll(r)
	if err != nil {
		return stdTx, err
	}
	err = cdc.UnmarshalJSON(bytes, &stdTx)
	return stdTx, err
}

func writeSignedGenTx(cdc *codec.Codec, outputDocument string, tx auth.StdTx) error {
	outputFile, err := os.OpenFile(outputDocument, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer outputFile.Close()
	json, err := cdc.MarshalJSON(tx)
	if err != nil {
		return err
	}
	_, err = fmt.Fprintf(outputFile, "%s\n", json)
	return err
}
