package cli

import (
	"fmt"
	// "strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	// sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/zigbee-alliance/distributed-compliance-ledger/x/validator/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group validator queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdListValidator())
	cmd.AddCommand(CmdShowValidator())
	cmd.AddCommand(CmdListLastValidatorPower())
	cmd.AddCommand(CmdShowLastValidatorPower())
	cmd.AddCommand(CmdListValidatorSigningInfo())
	cmd.AddCommand(CmdShowValidatorSigningInfo())
	cmd.AddCommand(CmdListValidatorMissedBlockBitArray())
	cmd.AddCommand(CmdShowValidatorMissedBlockBitArray())
	cmd.AddCommand(CmdListValidatorOwner())
	cmd.AddCommand(CmdShowValidatorOwner())
	// this line is used by starport scaffolding # 1

	return cmd
}
