package cli

import (
	"math"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"

	"github.com/zigbee-alliance/distributed-compliance-ledger/x/common"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliance/types"
)

var _ = strconv.Itoa(0)

func CmdUpdateComplianceInfo() *cobra.Command {
	var (
		vid                                int32
		pid                                int32
		softwareVersion                    uint32
		certificationType                  string
		cdVersionNumber                    string
		certificationDate                  string
		reason                             string
		owner                              string
		CDCertificateID                    string
		certificationRoute                 string
		programType                        string
		programTypeVersion                 string
		compliantPlatformUsed              string
		compliantPlatformVersion           string
		transport                          string
		familyID                           string
		supportedClusters                  string
		OSVersion                          string
		parentChild                        string
		certificationIDOfSoftwareComponent string
		specificationVersion               uint32
		schemaVersion                      uint32
	)

	cmd := &cobra.Command{
		Use:   "update-compliance-info",
		Short: "Update compliance info which exists on the ledger",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateComplianceInfo(
				clientCtx.GetFromAddress().String(),
				vid,
				pid,
				softwareVersion,
				certificationType,
				cdVersionNumber,
				certificationDate,
				reason,
				owner,
				CDCertificateID,
				certificationRoute,
				programType,
				programTypeVersion,
				compliantPlatformUsed,
				compliantPlatformVersion,
				transport,
				familyID,
				supportedClusters,
				OSVersion,
				parentChild,
				certificationIDOfSoftwareComponent,
				specificationVersion,
				schemaVersion,
			)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().Int32Var(&vid, FlagVID, 0, TextVID)
	cmd.Flags().Int32Var(&pid, FlagPID, 0, TextPID)
	cmd.Flags().Uint32VarP(&softwareVersion, FlagSoftwareVersion, FlagSoftwareVersionShortcut, math.MaxUint32, TextSoftwareVersion)
	cmd.Flags().StringVar(&cdVersionNumber, FlagCDVersionNumber, "", TextCDVersionNumber)
	cmd.Flags().StringVarP(&certificationType, FlagCertificationType, FlagCertificationTypeShortcut, "", TextCertificationType)
	cmd.Flags().Uint32Var(&specificationVersion, FlagSpecificationVersion, 0, TextSpecificationVersion)
	cmd.Flags().StringVarP(&certificationDate, FlagCertificationDate, FlagDateShortcut, "", TextCertificationDate)
	cmd.Flags().StringVar(&reason, FlagReason, "", TextCertificationReason)
	cmd.Flags().StringVar(&owner, FlagOwner, "", TextOwner)
	cmd.Flags().StringVar(&programTypeVersion, FlagProgramTypeVersion, "", TextProgramTypeVersion)
	cmd.Flags().StringVar(&CDCertificateID, FlagCDCertificateID, "", TextCDCertificateID)
	cmd.Flags().StringVar(&familyID, FlagFamilyID, "", TextFamilyID)
	cmd.Flags().StringVar(&supportedClusters, FlagSupportedClusters, "", TextSupportedClusters)
	cmd.Flags().StringVar(&compliantPlatformUsed, FlagCompliantPlatformUsed, "", TextCompliantPlatformUsed)
	cmd.Flags().StringVar(&compliantPlatformVersion, FlagCompliantPlatformVersion, "", TextCompliantPlatformVersion)
	cmd.Flags().StringVar(&OSVersion, FlagOSVersion, "", TextOSVersion)
	cmd.Flags().StringVar(&certificationRoute, FlagCertificationRoute, "", TextCertificationRoute)
	cmd.Flags().StringVar(&programType, FlagProgramType, "", TextProgramType)
	cmd.Flags().StringVar(&transport, FlagTransport, "", TextTransport)
	cmd.Flags().StringVar(&parentChild, FlagParentChild, "", TextParentChild)
	cmd.Flags().Uint32Var(&schemaVersion, common.FlagSchemaVersion, 0, TextSchemaVersion)

	_ = cmd.MarkFlagRequired(FlagVID)
	_ = cmd.MarkFlagRequired(FlagPID)
	_ = cmd.MarkFlagRequired(FlagSoftwareVersion)
	_ = cmd.MarkFlagRequired(FlagCertificationType)
	_ = cmd.MarkFlagRequired(flags.FlagFrom)

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
