package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/cli"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/common"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/model/types"
)

func CmdCreateModel() *cobra.Command {
	var (
		vid                                        int32
		pid                                        int32
		deviceTypeID                               int32
		productName                                string
		productLabel                               string
		partNumber                                 string
		discoveryCapabilitiesBitmask               uint32
		commissioningCustomFlow                    int32
		commissioningCustomFlowURL                 string
		commissioningModeInitialStepsHint          uint32
		commissioningModeInitialStepsInstruction   string
		commissioningModeSecondaryStepsHint        uint32
		commissioningModeSecondaryStepsInstruction string
		icdUserActiveModeTriggerHint               uint32
		icdUserActiveModeTriggerInstruction        string
		factoryResetStepsHint                      uint32
		factoryResetStepsInstruction               string
		userManualURL                              string
		supportURL                                 string
		productURL                                 string
		lsfURL                                     string
		enhancedSetupFlowOptions                   int32
		enhancedSetupFlowTCURL                     string
		enhancedSetupFlowTCRevision                int32
		enhancedSetupFlowTCDigest                  string
		enhancedSetupFlowTCFileSize                uint32
		maintenanceURL                             string
		commissioningFallbackURL                   string
		schemaVersion                              uint32
	)

	cmd := &cobra.Command{
		Use:   "add-model",
		Short: "Add new Model",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			isCommissioningFallbackURLSpecified := cmd.Flags().Changed(FlagCommissioningFallbackURL)
			isDiscoveryCapabilitiesBitmaskSpecified := cmd.Flags().Changed(FlagDiscoveryCapabilitiesBitmask)

			if isCommissioningFallbackURLSpecified && !isDiscoveryCapabilitiesBitmaskSpecified {
				return types.ErrFallbackURLRequiresBitmask
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			productLabel, err := utils.ReadFromFile(productLabel)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateModel(
				clientCtx.GetFromAddress().String(),
				vid,
				pid,
				deviceTypeID,
				productName,
				productLabel,
				partNumber,
				discoveryCapabilitiesBitmask,
				commissioningCustomFlow,
				commissioningCustomFlowURL,
				commissioningModeInitialStepsHint,
				commissioningModeInitialStepsInstruction,
				commissioningModeSecondaryStepsHint,
				commissioningModeSecondaryStepsInstruction,
				icdUserActiveModeTriggerHint,
				icdUserActiveModeTriggerInstruction,
				factoryResetStepsHint,
				factoryResetStepsInstruction,
				userManualURL,
				supportURL,
				productURL,
				lsfURL,
				schemaVersion,
				enhancedSetupFlowOptions,
				enhancedSetupFlowTCURL,
				enhancedSetupFlowTCRevision,
				enhancedSetupFlowTCDigest,
				enhancedSetupFlowTCFileSize,
				maintenanceURL,
				commissioningFallbackURL,
			)

			// validate basic will be called in GenerateOrBroadcastTxCLI
			err = tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
			if cli.IsWriteInsteadReadRPCError(err) {
				return clientCtx.PrintString(cli.LightClientProxyForWriteRequests)
			}

			return err
		},
	}

	cmd.Flags().Int32Var(&vid, FlagVid, 0,
		"Model vendor ID (positive non-zero uint16)")
	cmd.Flags().Int32Var(&pid, FlagPid, 0,
		"Model product ID (positive non-zero uint16)")
	cmd.Flags().Int32Var(&deviceTypeID, FlagDeviceTypeID, 0,
		"Model category ID")
	cmd.Flags().StringVarP(&productName, FlagProductName, FlagProductNameShortcut, "",
		"Model name")
	cmd.Flags().StringVarP(&productLabel, FlagProductLabel, FlagProductLabelShortcut, "",
		"Model description (string or path to file containing data)")
	cmd.Flags().StringVar(&partNumber, FlagPartNumber, "",
		"Model Part Number (or sku)")
	cmd.Flags().Uint32Var(&discoveryCapabilitiesBitmask, FlagDiscoveryCapabilitiesBitmask, 0,
		`This field identifies the device's available technologies for device discovery. 
This field SHALL be populated if CommissioningFallbackUrl is populated`)
	cmd.Flags().Int32Var(&commissioningCustomFlow, FlagCommissioningCustomFlow, 0,
		`A value of 1 indicates that user interaction with the device (pressing a button, for example) is 
required before commissioning can take place. When CommissioningCustomflow is set to a value of 2, 
the commissioner SHOULD attempt to obtain a URL which MAY be used to provide an end-user with 
the necessary details for how to configure the product for initial commissioning.`)
	cmd.Flags().StringVar(&commissioningCustomFlowURL, FlagCommissioningCustomFlowURL, "",
		`commissioningCustomFlowURL SHALL identify a vendor specific commissioning URL for the 
device model when the commissioningCustomFlow field is set to '2'`)
	cmd.Flags().Uint32Var(&commissioningModeInitialStepsHint, FlagCommissioningModeInitialStepsHint, 1,
		`commissioningModeInitialStepsHint SHALL 
identify a hint for the steps that can be used to put into commissioning mode a device that 
has not yet been commissioned. This field is a bitmap with values defined in the Pairing Hint Table. 
For example, a value of 1 (bit 0 is set) indicates 
that a device that has not yet been commissioned will enter Commissioning Mode upon a power cycle (default 1).`)
	cmd.Flags().StringVar(&commissioningModeInitialStepsInstruction, FlagCommissioningModeInitialStepsInstruction, "",
		`commissioningModeInitialStepsInstruction SHALL contain text which relates to specific 
values of commissioningModeInitialStepsHint. Certain values of CommissioningModeInitialStepsHint, 
as defined in the Pairing Hint Table, indicate a Pairing Instruction (PI) dependency, and for these 
values the commissioningModeInitialStepsInstruction SHALL be set`)
	cmd.Flags().Uint32Var(&commissioningModeSecondaryStepsHint, FlagCommissioningModeSecondaryStepsHint, 0,
		`commissioningModeSecondaryStepsHint SHALL identify a hint for steps that can 
be used to put into commissioning mode a device that has already been commissioned. 
This field is a bitmap with values defined in the Pairing Hint Table. For example, a value of 4 (bit 2 is set) 
indicates that a device that has already been commissioned will require the user to visit a 
current CHIP Administrator to put the device into commissioning mode.`)
	cmd.Flags().StringVar(&commissioningModeSecondaryStepsInstruction, FlagCommissioningModeSecondaryStepsInstruction, "",
		`commissioningModeSecondaryStepInstruction SHALL contain text which relates to specific values 
of commissioningModeSecondaryStepsHint. Certain values of commissioningModeSecondaryStepsHint, 
as defined in the Pairing Hint Table, indicate a Pairing Instruction (PI) dependency, 
and for these values the commissioningModeSecondaryStepInstruction SHALL be set`)
	cmd.Flags().Uint32Var(&icdUserActiveModeTriggerHint, FlagIcdUserActiveModeTriggerHint, 0,
		`IcdUserActiveModeTriggerHint (when provided) is applicable to 
an ICD that supports the UserActiveModeTriggerFeature feature. This field SHALL indicate which user action(s) 
will trigger the ICD to switch to Active mode. This field SHALL follow the requirements specified in UserActiveModeTriggerHint.`)
	cmd.Flags().StringVar(&icdUserActiveModeTriggerInstruction, FlagIcdUserActiveModeTriggerInstruction, "",
		`IcdUserActiveModeTriggerInstruction (when provided) is applicable to an ICD that supports the UserActiveModeTriggerFeature 
feature. The meaning of this field is dependent upon the UserActiveModeTriggerHint field value, and the conformance is indicated 
in the "dependency" column in UserActiveModeTriggerHintTable. This field SHALL follow the requirements specified in 
UserActiveModeTriggerInstruction.`)
	cmd.Flags().Uint32Var(&factoryResetStepsHint, FlagFactoryResetStepsHint, 0,
		`FactoryResetStepsHint SHALL identify a hint for the steps that MAY be used to factory 
reset a device. This field is a bitmap with values defined in the Pairing/Reset Hint Table. For example, 
a value of 64 (bit 6 is set) indicates that a device will be factory reset when the Reset Button is pressed.`)
	cmd.Flags().StringVar(&factoryResetStepsInstruction, FlagFactoryResetStepsInstruction, "",
		`FactoryResetStepsInstruction SHALL be populated with the appropriate 
factory reset instruction for those values of FactoryResetStepsHint, for which the Pairing/Reset Hint Table 
indicates a dependency in the Instruction Dependency column.`)
	cmd.Flags().StringVar(&userManualURL, FlagUserManualURL, "",
		"URL that contains product specific web page that contains user manual for the device model.")
	cmd.Flags().StringVar(&supportURL, FlagSupportURL, "",
		"URL that contains product specific web page that contains support details for the device model.")
	cmd.Flags().StringVar(&productURL, FlagProductURL, "",
		"URL that contains product specific web page that contains details for the device model.")
	cmd.Flags().StringVar(&lsfURL, FlagLsfURL, "", "URL to the Localized String File of this product")
	cli.AddTxFlagsToCmd(cmd)
	cmd.Flags().Uint32Var(&schemaVersion, common.FlagSchemaVersion, 0, "Schema version - default is 0, the value should be equal to 0")
	cmd.Flags().Int32Var(&enhancedSetupFlowOptions, FlagEnhancedSetupFlowOptions, 0,
		"enhancedSetupFlowOptions SHALL identify the configuration options for the Enhanced Setup Flow.")
	cmd.Flags().StringVar(&enhancedSetupFlowTCURL, FlagEnhancedSetupFlowTCURL, "",
		"enhancedSetupFlowTCURL SHALL identify a link to the Enhanced Setup Flow Terms and Condition File for this product. This field SHALL be present if and only if the EnhancedSetupFlowOptions field has bit 0 set.")
	cmd.Flags().Int32Var(&enhancedSetupFlowTCRevision, FlagEnhancedSetupFlowTCRevision, 0,
		"enhancedSetupFlowTCRevision is an increasing positive integer indicating the latest available version of the Enhanced Setup Flow Terms and Conditions file. This field SHALL be present if and only if the EnhancedSetupFlowOptions field has bit 0 set.")
	cmd.Flags().StringVar(&enhancedSetupFlowTCDigest, FlagEnhancedSetupFlowTCDigest, "",
		"enhancedSetupFlowTCDigest SHALL contain the digest of the entire contents of the associated file downloaded from the EnhancedSetupFlowTCUrl field, encoded in base64 string representation and SHALL be used to ensure the contents of the downloaded file are authentic. This field SHALL be present if and only if the EnhancedSetupFlowOptions field has bit 0 set.")
	cmd.Flags().Uint32Var(&enhancedSetupFlowTCFileSize, FlagEnhancedSetupFlowTCFileSize, 0,
		"enhancedSetupFlowTCFileSize SHALL indicate the total size of the Enhanced Setup Flow Terms and Conditions file in bytes, and SHALL be used to ensure the downloaded file size is within the bounds of EnhancedSetupFlowTCFileSize. This field SHALL be present if and only if the EnhancedSetupFlowOptions field has bit 0 set.")
	cmd.Flags().StringVar(&maintenanceURL, FlagMaintenanceURL, "",
		"maintenanceURL SHALL identify a link to a vendor-specific URL which SHALL provide a manufacturer specific means to resolve any functionality limitations indicated by the TERMS_AND_CONDITIONS_CHANGED status code. This field SHALL be present if and only if the EnhancedSetupFlowOptions field has bit 0 set.")
	cmd.Flags().StringVar(&commissioningFallbackURL, FlagCommissioningFallbackURL, "",
		"This field SHALL identify a vendor-specific commissioning-fallback URL for the device model, which can be used by a Commissioner in case commissioning fails to direct the user to a manufacturer-provided mechanism to provide resolution to commissioning issues.")

	_ = cmd.MarkFlagRequired(FlagVid)
	_ = cmd.MarkFlagRequired(FlagPid)
	_ = cmd.MarkFlagRequired(FlagDeviceTypeID)
	_ = cmd.MarkFlagRequired(FlagProductName)
	_ = cmd.MarkFlagRequired(flags.FlagFrom)

	return cmd
}

func CmdUpdateModel() *cobra.Command {
	var (
		vid                                        int32
		pid                                        int32
		productName                                string
		productLabel                               string
		partNumber                                 string
		commissioningCustomFlowURL                 string
		commissioningModeInitialStepsInstruction   string
		commissioningModeSecondaryStepsInstruction string
		icdUserActiveModeTriggerHint               uint32
		icdUserActiveModeTriggerInstruction        string
		factoryResetStepsHint                      uint32
		factoryResetStepsInstruction               string
		userManualURL                              string
		supportURL                                 string
		productURL                                 string
		lsfURL                                     string
		lsfRevision                                int32
		schemaVersion                              uint32
		commissioningModeInitialStepsHint          uint32
		enhancedSetupFlowOptions                   int32
		enhancedSetupFlowTCURL                     string
		enhancedSetupFlowTCRevision                int32
		enhancedSetupFlowTCDigest                  string
		enhancedSetupFlowTCFileSize                uint32
		maintenanceURL                             string
		commissioningFallbackURL                   string
		commissioningModeSecondaryStepsHint        uint32
	)

	cmd := &cobra.Command{
		Use:   "update-model",
		Short: "Update existing Model",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			productLabel, err := utils.ReadFromFile(productLabel)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateModel(
				clientCtx.GetFromAddress().String(),
				vid,
				pid,
				productName,
				productLabel,
				partNumber,
				commissioningCustomFlowURL,
				commissioningModeInitialStepsInstruction,
				commissioningModeSecondaryStepsInstruction,
				userManualURL,
				supportURL,
				productURL,
				lsfURL,
				lsfRevision,
				schemaVersion,
				commissioningModeInitialStepsHint,
				enhancedSetupFlowOptions,
				enhancedSetupFlowTCURL,
				enhancedSetupFlowTCRevision,
				enhancedSetupFlowTCDigest,
				enhancedSetupFlowTCFileSize,
				maintenanceURL,
				commissioningFallbackURL,
				commissioningModeSecondaryStepsHint,
				icdUserActiveModeTriggerHint,
				icdUserActiveModeTriggerInstruction,
				factoryResetStepsHint,
				factoryResetStepsInstruction,
			)

			// validate basic will be called in GenerateOrBroadcastTxCLI
			err = tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
			if cli.IsWriteInsteadReadRPCError(err) {
				return clientCtx.PrintString(cli.LightClientProxyForWriteRequests)
			}

			return err
		},
	}

	cmd.Flags().Int32Var(&vid, FlagVid, 0,
		"Model vendor ID (positive non-zero uint16)")
	cmd.Flags().Int32Var(&pid, FlagPid, 0,
		"Model product ID (positive non-zero uint16)")
	cmd.Flags().StringVarP(&productName, FlagProductName, FlagProductNameShortcut, "",
		"Model name")
	cmd.Flags().StringVarP(&productLabel, FlagProductLabel, FlagProductLabelShortcut, "",
		"Model description (string or path to file containing data)")
	cmd.Flags().StringVar(&partNumber, FlagPartNumber, "",
		"Model Part Number (or sku)")
	cmd.Flags().StringVar(&commissioningCustomFlowURL, FlagCommissioningCustomFlowURL, "",
		`commissioningCustomFlowURL SHALL identify a vendor specific commissioning URL for the 
device model when the commissioningCustomFlow field is set to '2'`)
	cmd.Flags().StringVar(&commissioningModeInitialStepsInstruction, FlagCommissioningModeInitialStepsInstruction, "",
		`commissioningModeInitialStepsInstruction SHALL contain text which relates to specific 
values of commissioningModeInitialStepsHint. Certain values of CommissioningModeInitialStepsHint, 
as defined in the Pairing Hint Table, indicate a Pairing Instruction (PI) dependency, and for these 
values the commissioningModeInitialStepsInstruction SHALL be set`)
	cmd.Flags().StringVar(&commissioningModeSecondaryStepsInstruction, FlagCommissioningModeSecondaryStepsInstruction, "",
		`commissioningModeSecondaryStepInstruction SHALL contain text which relates to specific values 
of commissioningModeSecondaryStepsHint. Certain values of commissioningModeSecondaryStepsHint, 
as defined in the Pairing Hint Table, indicate a Pairing Instruction (PI) dependency, 
and for these values the commissioningModeSecondaryStepInstruction SHALL be set`)
	cmd.Flags().Uint32Var(&icdUserActiveModeTriggerHint, FlagIcdUserActiveModeTriggerHint, 0,
		`IcdUserActiveModeTriggerHint (when provided) is applicable to 
an ICD that supports the UserActiveModeTriggerFeature feature. This field SHALL indicate which user action(s) 
will trigger the ICD to switch to Active mode. This field SHALL follow the requirements specified in UserActiveModeTriggerHint.`)
	cmd.Flags().StringVar(&icdUserActiveModeTriggerInstruction, FlagIcdUserActiveModeTriggerInstruction, "",
		`IcdUserActiveModeTriggerInstruction (when provided) is applicable to an ICD that supports the UserActiveModeTriggerFeature 
feature. The meaning of this field is dependent upon the UserActiveModeTriggerHint field value, and the conformance is indicated 
in the "dependency" column in UserActiveModeTriggerHintTable. This field SHALL follow the requirements specified in 
UserActiveModeTriggerInstruction.`)
	cmd.Flags().StringVar(&factoryResetStepsInstruction, FlagFactoryResetStepsInstruction, "",
		`FactoryResetStepsInstruction SHALL be populated with the appropriate 
factory reset instruction for those values of FactoryResetStepsHint, for which the Pairing/Reset Hint Table 
indicates a dependency in the Instruction Dependency column.`)
	cmd.Flags().StringVar(&userManualURL, FlagUserManualURL, "",
		"URL that contains product specific web page that contains user manual for the device model.")
	cmd.Flags().StringVar(&supportURL, FlagSupportURL, "",
		"URL that contains product specific web page that contains support details for the device model.")
	cmd.Flags().StringVar(&productURL, FlagProductURL, "",
		"URL that contains product specific web page that contains details for the device model.")
	cmd.Flags().StringVar(&lsfURL, FlagLsfURL, "", "URL to the Localized String File of this product")
	cmd.Flags().Int32Var(&lsfRevision, FlagLsfRevision, 0,
		"LsfRevision is a monotonically increasing positive integer indicating the latest available version of Localized String File")
	cmd.Flags().Uint32Var(&schemaVersion, common.FlagSchemaVersion, 0, "Schema version")
	cmd.Flags().Uint32Var(&commissioningModeInitialStepsHint, FlagCommissioningModeInitialStepsHint, 0,
		`commissioningModeInitialStepsHint SHALL 
identify a hint for the steps that can be used to put into commissioning mode a device that 
has not yet been commissioned. This field is a bitmap with values defined in the Pairing Hint Table. 
For example, a value of 1 (bit 0 is set) indicates that a device that has not yet been commissioned 
will enter Commissioning Mode upon a power cycle. Note that this value cannot be updated to 0. (default 1).`)
	cmd.Flags().Uint32Var(&commissioningModeSecondaryStepsHint, FlagCommissioningModeSecondaryStepsHint, 0,
		`commissioningModeSecondaryStepsHint SHALL 
identify a hint for the steps that can be used to put into commissioning mode a device that 
has not yet been commissioned. This field is a bitmap with values defined in the Pairing Hint Table. 
For example, a value of 1 (bit 0 is set) indicates that a device that has not yet been commissioned 
will enter Commissioning Mode upon a power cycle. Note that this value cannot be updated to 0. (default 1).`)
	cmd.Flags().Uint32Var(&factoryResetStepsHint, FlagFactoryResetStepsHint, 0,
		`FactoryResetStepsHint SHALL identify a hint for the steps that MAY be used to factory 
reset a device. This field is a bitmap with values defined in the Pairing/Reset Hint Table. For example, 
a value of 64 (bit 6 is set) indicates that a device will be factory reset when the Reset Button is pressed.`)
	cmd.Flags().Int32Var(&enhancedSetupFlowOptions, FlagEnhancedSetupFlowOptions, 0,
		"enhancedSetupFlowOptions SHALL identify the configuration options for the Enhanced Setup Flow.")
	cmd.Flags().StringVar(&enhancedSetupFlowTCURL, FlagEnhancedSetupFlowTCURL, "",
		"enhancedSetupFlowTCURL SHALL identify a link to the Enhanced Setup Flow Terms and Condition File for this product. This field SHALL be present if and only if the EnhancedSetupFlowOptions field has bit 0 set.")
	cmd.Flags().Int32Var(&enhancedSetupFlowTCRevision, FlagEnhancedSetupFlowTCRevision, 0,
		"enhancedSetupFlowTCRevision is an increasing positive integer indicating the latest available version of the Enhanced Setup Flow Terms and Conditions file. This field SHALL be present if and only if the EnhancedSetupFlowOptions field has bit 0 set.")
	cmd.Flags().StringVar(&enhancedSetupFlowTCDigest, FlagEnhancedSetupFlowTCDigest, "",
		"enhancedSetupFlowTCDigest SHALL contain the digest of the entire contents of the associated file downloaded from the EnhancedSetupFlowTCUrl field, encoded in base64 string representation and SHALL be used to ensure the contents of the downloaded file are authentic. This field SHALL be present if and only if the EnhancedSetupFlowOptions field has bit 0 set.")
	cmd.Flags().Uint32Var(&enhancedSetupFlowTCFileSize, FlagEnhancedSetupFlowTCFileSize, 0,
		"enhancedSetupFlowTCFileSize SHALL indicate the total size of the Enhanced Setup Flow Terms and Conditions file in bytes, and SHALL be used to ensure the downloaded file size is within the bounds of EnhancedSetupFlowTCFileSize. This field SHALL be present if and only if the EnhancedSetupFlowOptions field has bit 0 set.")
	cmd.Flags().StringVar(&maintenanceURL, FlagMaintenanceURL, "",
		"maintenanceURL SHALL identify a link to a vendor-specific URL which SHALL provide a manufacturer specific means to resolve any functionality limitations indicated by the TERMS_AND_CONDITIONS_CHANGED status code. This field SHALL be present if and only if the EnhancedSetupFlowOptions field has bit 0 set.")
	cmd.Flags().StringVar(&commissioningFallbackURL, FlagCommissioningFallbackURL, "",
		"This field SHALL identify a vendor-specific commissioning-fallback URL for the device model, which can be used by a Commissioner in case commissioning fails to direct the user to a manufacturer-provided mechanism to provide resolution to commissioning issues.")

	cli.AddTxFlagsToCmd(cmd)

	_ = cmd.MarkFlagRequired(FlagVid)
	_ = cmd.MarkFlagRequired(FlagPid)
	_ = cmd.MarkFlagRequired(flags.FlagFrom)

	return cmd
}

func CmdDeleteModel() *cobra.Command {
	var (
		vid int32
		pid int32
	)

	cmd := &cobra.Command{
		Use:   "delete-model",
		Short: "Delete existing Model",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			msg := types.NewMsgDeleteModel(
				clientCtx.GetFromAddress().String(),
				vid,
				pid,
			)

			// validate basic will be called in GenerateOrBroadcastTxCLI
			err = tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
			if cli.IsWriteInsteadReadRPCError(err) {
				return clientCtx.PrintString(cli.LightClientProxyForWriteRequests)
			}

			return err
		},
	}

	cmd.Flags().Int32Var(&vid, FlagVid, 0, "Model vendor ID (positive non-zero uint16)")
	cmd.Flags().Int32Var(&pid, FlagPid, 0, "Model product ID (positive non-zero uint16)")

	cli.AddTxFlagsToCmd(cmd)

	_ = cmd.MarkFlagRequired(FlagVid)
	_ = cmd.MarkFlagRequired(FlagPid)
	_ = cmd.MarkFlagRequired(flags.FlagFrom)

	return cmd
}
