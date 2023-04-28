package cli

import (
	"math"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
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
			)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().Int32Var(&vid, FlagVID, 0,
		"Model vendor ID (positive non-zero uint16)")
	cmd.Flags().Int32Var(&pid, FlagPID, 0,
		"Model product ID (positive non-zero uint16)")
	cmd.Flags().Uint32VarP(&softwareVersion, FlagSoftwareVersion, FlagSoftwareVersionShortcut, math.MaxUint32,
		"Software Version of model (uint32)")
	cmd.Flags().StringVar(&cdVersionNumber, FlagCDVersionNumber, "",
		"CD Version Number of the certification")
	cmd.Flags().StringVarP(&certificationType, FlagCertificationType, FlagCertificationTypeShortcut, "", TextCertificationType)
	cmd.Flags().StringVarP(&certificationDate, FlagCertificationDate, FlagDateShortcut, "",
		"The date of model certification (rfc3339 encoded), for example 2019-10-12T07:20:50.52Z")
	cmd.Flags().StringVar(&reason, FlagReason, "",
		"Optional comment describing the reason of certification")
	cmd.Flags().StringVar(&owner, FlagOwner, "", "Signer of certification")
	cmd.Flags().StringVar(&programTypeVersion, FlagProgramTypeVersion, "",
		"Program Type Version of the certification")
	cmd.Flags().StringVar(&CDCertificateID, FlagCDCertificateID, "",
		"CD Certification ID of the certification")
	cmd.Flags().StringVar(&familyID, FlagFamilyID, "",
		"Family ID of the certification")
	cmd.Flags().StringVar(&supportedClusters, FlagSupportedClusters, "",
		"Supported Clusters of the certification")
	cmd.Flags().StringVar(&compliantPlatformUsed, FlagCompliantPlatformUsed, "",
		"Compliant Platform Used of the certification")
	cmd.Flags().StringVar(&compliantPlatformVersion, FlagCompliantPlatformVersion, "",
		"Compliant Platform Version of the certification")
	cmd.Flags().StringVar(&OSVersion, FlagOSVersion, "",
		"OS Version of the certification")
	cmd.Flags().StringVar(&certificationRoute, FlagCertificationRoute, "",
		"Certification Route of the certification")
	cmd.Flags().StringVar(&programType, FlagProgramType, "",
		"Program Type of the certification")
	cmd.Flags().StringVar(&transport, FlagTransport, "",
		"Transport of the certification")
	cmd.Flags().StringVar(&parentChild, FlagParentChild, "",
		"Parent or Child  of the PFC certification route")
	cmd.Flags().StringVar(&certificationIDOfSoftwareComponent, FlagCertificationIDOfSoftwareComponent, "",
		"certification ID of software component")

	_ = cmd.MarkFlagRequired(FlagVID)
	_ = cmd.MarkFlagRequired(FlagPID)
	_ = cmd.MarkFlagRequired(FlagSoftwareVersion)
	_ = cmd.MarkFlagRequired(FlagCertificationType)
	_ = cmd.MarkFlagRequired(flags.FlagFrom)

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
