package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/cli"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/common"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliance/types"
)

var _ = strconv.Itoa(0)

func CmdCertifyModel() *cobra.Command {
	var (
		vid                                int32
		pid                                int32
		softwareVersion                    uint32
		softwareVersionString              string
		certificationDate                  string
		certificationType                  string
		reason                             string
		cdVersionNumber                    uint32
		programTypeVersion                 string
		CDCertificateID                    string
		familyID                           string
		supportedClusters                  string
		compliantPlatformUsed              string
		compliantPlatformVersion           string
		OSVersion                          string
		certificationRoute                 string
		programType                        string
		transport                          string
		parentChild                        string
		certificationIDOfSoftwareComponent string
		specificationVersion               uint32
		schemaVersion                      uint32
	)

	cmd := &cobra.Command{
		Use:   "certify-model",
		Short: "Certify an existing model",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCertifyModel(
				clientCtx.GetFromAddress().String(),
				vid,
				pid,
				softwareVersion,
				softwareVersionString,
				cdVersionNumber,
				certificationDate,
				certificationType,
				reason,
				programTypeVersion,
				CDCertificateID,
				familyID,
				supportedClusters,
				compliantPlatformUsed,
				compliantPlatformVersion,
				OSVersion,
				certificationRoute,
				programType,
				transport,
				parentChild,
				certificationIDOfSoftwareComponent,
				specificationVersion,
				schemaVersion,
			)

			// validate basic will be called in GenerateOrBroadcastTxCLI
			err = tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
			if cli.IsWriteInsteadReadRPCError(err) {
				return clientCtx.PrintString(cli.LightClientProxyForWriteRequests)
			}

			return err
		},
	}
	cmd.Flags().Int32Var(&vid, FlagVID, 0, TextVID)
	cmd.Flags().Int32Var(&pid, FlagPID, 0, TextPID)
	cmd.Flags().Uint32VarP(&softwareVersion, FlagSoftwareVersion, FlagSoftwareVersionShortcut, 0, TextSoftwareVersion)
	cmd.Flags().StringVar(&softwareVersionString, FlagSoftwareVersionString, "", TextSoftwareVersionString)
	cmd.Flags().Uint32Var(&cdVersionNumber, FlagCDVersionNumber, 0, TextCDVersionNumber)
	cmd.Flags().StringVarP(&certificationType, FlagCertificationType, FlagCertificationTypeShortcut, "", TextCertificationType)
	cmd.Flags().Uint32Var(&specificationVersion, FlagSpecificationVersion, 0, TextSpecificationVersion)
	cmd.Flags().StringVarP(&certificationDate, FlagCertificationDate, FlagDateShortcut, "", TextCertificationDate)
	cmd.Flags().StringVar(&reason, FlagReason, "", TextCertificationReason)
	cmd.Flags().StringVar(&CDCertificateID, FlagCDCertificateID, "", TextCDCertificateID)
	cmd.Flags().StringVar(&familyID, FlagFamilyID, "", TextFamilyID)
	cmd.Flags().StringVar(&supportedClusters, FlagSupportedClusters, "", TextSupportedClusters)
	cmd.Flags().StringVar(&certificationRoute, FlagCertificationRoute, "", TextCertificationRoute)
	cmd.Flags().StringVar(&transport, FlagTransport, "", TextTransport)
	cmd.Flags().StringVar(&parentChild, FlagParentChild, "", TextParentChild)
	cmd.Flags().Uint32Var(&schemaVersion, common.FlagSchemaVersion, 0, TextSchemaVersion)
	cmd.Flags().StringVar(&programType, FlagProgramType, "", TextProgramType)
	cmd.Flags().StringVar(&programTypeVersion, FlagProgramTypeVersion, "", TextProgramTypeVersion)
	// Deprecated fields
	cmd.Flags().StringVar(&compliantPlatformUsed, FlagCompliantPlatformUsed, "", TextCompliantPlatformUsed)
	_ = cmd.Flags().MarkDeprecated(FlagCompliantPlatformUsed, DeprecatedTextCompliantPlatformUsed)
	cmd.Flags().StringVar(&compliantPlatformVersion, FlagCompliantPlatformVersion, "", TextCompliantPlatformVersion)
	_ = cmd.Flags().MarkDeprecated(FlagCompliantPlatformVersion, DeprecatedTextCompliantPlatformVersion)
	cmd.Flags().StringVar(&certificationIDOfSoftwareComponent, FlagCertificationIDOfSoftwareComponent, "", TextCertificationIDOfSoftwareComponent)
	_ = cmd.Flags().MarkDeprecated(FlagCertificationIDOfSoftwareComponent, DeprecatedTextCertificationIDOfSoftwareComponent)
	cmd.Flags().StringVar(&OSVersion, FlagOSVersion, "", TextOSVersion)
	_ = cmd.Flags().MarkDeprecated(FlagOSVersion, DeprecatedTextOSVersion)

	// Required fields
	_ = cmd.MarkFlagRequired(FlagVID)
	_ = cmd.MarkFlagRequired(FlagPID)
	_ = cmd.MarkFlagRequired(FlagSoftwareVersion)
	_ = cmd.MarkFlagRequired(FlagSoftwareVersionString)
	_ = cmd.MarkFlagRequired(FlagCertificationType)
	_ = cmd.MarkFlagRequired(FlagSpecificationVersion)
	_ = cmd.MarkFlagRequired(FlagCertificationDate)
	_ = cmd.MarkFlagRequired(FlagCDCertificateID)
	_ = cmd.MarkFlagRequired(flags.FlagFrom)

	cli.AddTxFlagsToCmd(cmd)

	return cmd
}
