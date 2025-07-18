# Model and Model Version

<!-- markdownlint-disable MD036 -->

### ADD_MODEL

**Status: Implemented**

Adds a new Model identified by a unique combination of `vid` (vendor ID) and `pid` (product ID). If `account` was created with product ID ranges then the `pid` must fall within that specified range.

Not all fields can be edited (see `EDIT_MODEL`).

- Parameters:
  - vid: `uint16` -  model vendor ID (positive non-zero)
  - pid: `uint16` -  model product ID (positive non-zero)
  - deviceTypeID: `uint16` -  DeviceTypeID is the device type identifier. For example, DeviceTypeID 10 (0x000a), is the device type identifier for a Door Lock.
  - productName: `string` -  model name
  - productLabel: `optional(string)` -  model description (string or path to file containing data)
  - partNumber: `optional(string)` -  stock keeping unit
  - commissioningCustomFlow: `optional(uint8)` - A value of 1 indicates that user interaction with the device (pressing a button, for example) is required before commissioning can take place. When CommissioningCustomflow is set to a value of 2, the commissioner SHOULD attempt to obtain a URL which MAY be used to provide an end user with the necessary details for how to configure the product for initial commissioning
  - commissioningCustomFlowURL: `optional(string)` - commissioningCustomFlowURL SHALL identify a vendor specific commissioning URL for the device model when the commissioningCustomFlow field is set to '2'
  - commissioningModeInitialStepsHint: `optional(uint32)` - commissioningModeInitialStepsHint SHALL identify a hint for the steps that can be used to put into commissioning mode a device that has not yet been commissioned. This field is a bitmap with values defined in the Pairing Hint Table. For example, a value of 1 (bit 0 is set) indicates that a device that has not yet been commissioned will enter Commissioning Mode upon a power cycle (default 1).
  - commissioningModeInitialStepsInstruction: `optional(string)` - commissioningModeInitialStepsInstruction SHALL contain text which relates to specific values of CommissioningModeInitialStepsHint. Certain values of CommissioningModeInitialStepsHint, as defined in the Pairing Hint Table, indicate a Pairing Instruction (PI) dependency, and for these values the commissioningModeInitialStepsInstruction SHALL be set
  - commissioningModeSecondaryStepsHint: `optional(uint32)` - commissioningModeSecondaryStepsHint SHALL identify a hint for steps that can be used to put into commissioning mode a device that has already been commissioned. This field is a bitmap with values defined in the Pairing Hint Table. For example, a value of 4 (bit 2 is set) indicates that a device that has already been commissioned will require the user to visit a current CHIP Administrator to put the device into commissioning mode (default 1).
  - commissioningModeSecondaryStepInstruction: `optional(string)` - commissioningModeSecondaryStepInstruction SHALL contain text which relates to specific values of commissioningModeSecondaryStepsHint. Certain values of commissioningModeSecondaryStepsHint, as defined in the Pairing Hint Table, indicate a Pairing Instruction (PI) dependency, and for these values the commissioningModeSecondaryStepInstruction SHALL be set
  - userManualURL: `optional(string)` - URL that contains product specific web page that contains user manual for the device model.
  - supportURL: `optional(string)` - URL that contains product specific web page that contains support details for the device model.
  - productURL: `optional(string)` - URL that contains product specific web page that contains details for the device model.
  - lsfURL: `optional(string)` - URL to the Localized String File of this product.
  - enhancedSetupFlowOptions: `optional(uint16)` - enhancedSetupFlowOptions SHALL identify the configuration options for the Enhanced Setup Flow.
  - enhancedSetupFlowTCUrl: `optional(string)` - enhancedSetupFlowTCUrl SHALL identify a link to the Enhanced Setup Flow Terms and Condition File for this product. This field SHALL be present if and only if the EnhancedSetupFlowOptions field has bit 0 set.
  - enhancedSetupFlowTCRevision: `optional(uint16)` - enhancedSetupFlowTCRevision is an increasing positive integer indicating the latest available version of the Enhanced Setup Flow Terms and Conditions file. This field SHALL be present if and only if the EnhancedSetupFlowOptions field has bit 0 set.
  - enhancedSetupFlowTCDigest: `optional(string)` - enhancedSetupFlowTCDigest SHALL contain the digest of the entire contents of the associated file downloaded from the EnhancedSetupFlowTCUrl field, encoded in base64 string representation and SHALL be used to ensure the contents of the downloaded file are authentic. This field SHALL be present if and only if the EnhancedSetupFlowOptions field has bit 0 set.
  - enhancedSetupFlowTCFileSize: `optional(uint32)` - enhancedSetupFlowTCFileSize SHALL indicate the total size of the Enhanced Setup Flow Terms and Conditions file in bytes, and SHALL be used to ensure the downloaded file size is within the bounds of EnhancedSetupFlowTCFileSize. This field SHALL be present if and only if the EnhancedSetupFlowOptions field has bit 0 set.
  - maintenanceUrl: `optional(string)` - maintenanceUrl SHALL identify a link to a vendor-specific URL which SHALL provide a manufacturer specific means to resolve any functionality limitations indicated by the TERMS_AND_CONDITIONS_CHANGED status code. This field SHALL be present if and only if the EnhancedSetupFlowOptions field has bit 0 set.
  - schemaVersion: `optional(uint16)` - Schema version to support backward/forward compatability. Should be equal to 0 (default 0)
  - discoveryCapabilitiesBitmask: `optional(uint16)` - Identifies the device's available technologies for device discovery (default 0). This field SHALL be populated if CommissioningFallbackUrl is populated
  - commissioningFallbackURL: `optional(string)` - This field SHALL identify a vendor-specific commissioning-fallback URL for the device model, which can be used by a Commissioner in case commissioning fails to direct the user to a manufacturer-provided mechanism to provide resolution to commissioning issues.
- In State:
  - `model/Model/value/<vid>/<pid>`
  - `model/VendorProducts/value/<vid>`
- Who can send:
  - Vendor account who is associated with the given vid
- CLI command minimal:

```bash
dcld tx model add-model --vid=<uint16> --pid=<uint16> --deviceTypeID=<uint16> --productName=<string> --productLabel=<string or path> --partNumber=<string> --from=<account>
```

- CLI command full:

```bash
dcld tx model add-model --vid=<uint16> --pid=<uint16> --deviceTypeID=<uint16> --productName=<string> --productLabel=<string or path> --partNumber=<string> 
    --commissioningCustomFlow=<uint8> --commissioningCustomFlowUrl=<string> --commissioningModeInitialStepsHint=<uint32> --commissioningModeInitialStepsInstruction=<string>
    --commissioningModeSecondaryStepsHint=<uint32> --commissioningModeSecondaryStepsInstruction=<string> --userManualURL=<string> --supportURL=<string> --productURL=<string> --lsfURL=<string> --discoveryCapabilitiesBitmask=<uint16> --commissioningFallbackURL<string>
    --from=<account>
```

### EDIT_MODEL

**Status: Implemented**

Edits an existing Model identified by a unique combination of `vid` (vendor ID) and `pid` (product ID)
by the vendor account. If `account` was created with product ID ranges then the `pid` must fall within that specified range.

Only the fields listed below (except `vid` and `pid`) can be edited. If other fields need to be edited -
a new model info with a new `vid` or `pid` can be created.

All non-edited fields remain the same.

If one of EnhancedSetupFlow or MaintenanceUrl fields needs to be updated, ALL EnhancedSetupFlow fields MUST be specified, and EnhancedSetupFlowOptions field must have bit 0 set.

- Parameters:
  - vid: `uint16` -  model vendor ID (positive non-zero)
  - pid: `uint16` -  model product ID (positive non-zero)
  - productName: `optional(string)` -  model name
  - productLabel: `optional(string)` -  model description (string or path to file containing data)
  - partNumber: `optional(string)` -  stock keeping unit
  - commissioningCustomFlowURL: `optional(string)` - commissioningCustomFlowURL SHALL identify a vendor specific commissioning URL for the device model when the commissioningCustomFlow field is set to '2'
  - commissioningModeInitialStepsInstruction: `optional(string)` - commissioningModeInitialStepsInstruction SHALL contain text which relates to specific values of CommissioningModeInitialStepsHint. Certain values of CommissioningModeInitialStepsHint, as defined in the Pairing Hint Table, indicate a Pairing Instruction (PI) dependency, and for these values the commissioningModeInitialStepsInstruction SHALL be set
  - commissioningModeSecondaryStepInstruction: `optional(string)` - commissioningModeSecondaryStepInstruction SHALL contain text which relates to specific values of commissioningModeSecondaryStepsHint. Certain values of commissioningModeSecondaryStepsHint, as defined in the Pairing Hint Table, indicate a Pairing Instruction (PI) dependency, and for these values the commissioningModeSecondaryStepInstruction SHALL be set
  - userManualURL: `optional(string)` - URL that contains product specific web page that contains user manual for the device model.
  - supportURL: `optional(string)` - URL that contains product specific web page that contains support details for the device model.
  - productURL: `optional(string)` - URL that contains product specific web page that contains details for the device model.  
  - lsfURL: `optional(string)` - URL to the Localized String File of this product.
  - lsfRevision: `optional(uint32)` - LsfRevision is a monotonically increasing positive integer indicating the latest available version of Localized String File.
  - commissioningModeInitialStepsHint: `optional(uint32)` - commissioningModeInitialStepsHint SHALL identify a hint for the steps that can be used to put into commissioning mode a device that has not yet been commissioned. This field is a bitmap with values defined in the Pairing Hint Table. For example, a value of 1 (bit 0 is set) indicates that a device that has not yet been commissioned will enter Commissioning Mode upon a power cycle. Note that this value cannot be updated to 0.
  - commissioningModeSecondaryStepsHint: `optional(uint32)` - commissioningModeSecondaryStepsHint SHALL identify a hint for steps that can be used to put into commissioning mode a device that has already been commissioned. This field is a bitmap with values defined in the Pairing Hint Table. For example, a value of 4 (bit 2 is set) indicates that a device that has already been commissioned will require the user to visit a current CHIP Administrator to put the device into commissioning mode. Note that this value cannot be updated to 0.
  - enhancedSetupFlowOptions: `optional(uint16)` - enhancedSetupFlowOptions SHALL identify the configuration options for the Enhanced Setup Flow.
  - enhancedSetupFlowTCUrl: `optional(string)` - enhancedSetupFlowTCUrl SHALL identify a link to the Enhanced Setup Flow Terms and Condition File for this product. This field SHALL be present if and only if the EnhancedSetupFlowOptions field has bit 0 set.
  - enhancedSetupFlowTCRevision: `optional(uint16)` - enhancedSetupFlowTCRevision is an increasing positive integer indicating the latest available version of the Enhanced Setup Flow Terms and Conditions file. This field SHALL be present if and only if the EnhancedSetupFlowOptions field has bit 0 set.
  - enhancedSetupFlowTCDigest: `optional(string)` - enhancedSetupFlowTCDigest SHALL contain the digest of the entire contents of the associated file downloaded from the EnhancedSetupFlowTCUrl field, encoded in base64 string representation and SHALL be used to ensure the contents of the downloaded file are authentic. This field SHALL be present if and only if the EnhancedSetupFlowOptions field has bit 0 set.
  - enhancedSetupFlowTCFileSize: `optional(uint32)` - enhancedSetupFlowTCFileSize SHALL indicate the total size of the Enhanced Setup Flow Terms and Conditions file in bytes, and SHALL be used to ensure the downloaded file size is within the bounds of EnhancedSetupFlowTCFileSize. This field SHALL be present if and only if the EnhancedSetupFlowOptions field has bit 0 set.
  - maintenanceUrl: `optional(string)` - maintenanceUrl SHALL identify a link to a vendor-specific URL which SHALL provide a manufacturer specific means to resolve any functionality limitations indicated by the TERMS_AND_CONDITIONS_CHANGED status code. This field SHALL be present if and only if the EnhancedSetupFlowOptions field has bit 0 set.
  - schemaVersion: `optional(uint16)` - Schema version to support backward/forward compatability. Should be equal to 0 (default 0)
  - commissioningFallbackURL: `optional(string)` - This field SHALL identify a vendor-specific commissioning-fallback URL for the device model, which can be used by a Commissioner in case commissioning fails to direct the user to a manufacturer-provided mechanism to provide resolution to commissioning issues.
- In State: `model/Model/value/<vid>/<pid>`
- Who can send:
  - Vendor account associated with the same vid who has created the model
- CLI command:
  - `dcld tx model update-model --vid=<uint16> --pid=<uint16> ... --from=<account>`

### DELETE_MODEL

**Status: Implemented**

Deletes an existing Model identified by a unique combination of `vid` (vendor ID) and `pid` (product ID)
by the vendor account. If `account` was created with product ID ranges then the `pid` must fall within that specified range.

If one of Model Versions associated with the Model is certified then Model can not be deleted. When Model is deleted, all associated Model Versions will be deleted as well.

- Parameters:
  - vid: `uint16` -  model vendor ID (positive non-zero)
  - pid: `uint16` -  model product ID (positive non-zero)
- In State: `model/Model/value/<vid>/<pid>`
- Who can send:
  - Vendor account associated with the same vid who has created the model
- CLI command:
  - `dcld tx model delete-model --vid=<uint16> --pid=<uint16> --from=<account>`

### ADD_MODEL_VERSION

**Status: Implemented**

Adds a new Model Software Version identified by a unique combination of `vid` (vendor ID), `pid` (product ID) and `softwareVersion`.
If `account` was created with product ID ranges then the `pid` must fall within that specified range

Not all Model Software Version fields can be edited (see `EDIT_MODEL_VERSION`).

If one of `OTA_URl`, `OTA_checksum` or `OTA_checksum_type` fields is set, then the other two must also be set.

- Parameters:
  - vid: `uint16` -  model vendor ID (positive non-zero)
  - pid: `uint16` -  model product ID (positive non-zero)
  - softwareVersion: `uint32` - model software version
  - softwareVersionSting: `string` - model software version string
  - cdVersionNumber `uint16` - CD Version Number of the certification
  - minApplicableSoftwareVersion `uint32` - MinApplicableSoftwareVersion should specify the lowest SoftwareVersion for which this image can be applied
  - maxApplicableSoftwareVersion `uint32` - MaxApplicableSoftwareVersion should specify the highest SoftwareVersion for which this image can be applied
  - firmwareInformation `optional(string)` - FirmwareInformation field included in the Device Attestation response when this Software Image boots on the device
  - softwareVersionValid `optional(bool)` - Flag to indicate whether the software version is valid or not (default true)
  - otaURL `optional(string)` - URL where to obtain the OTA image
  - otaFileSize `optional(string)`  - OtaFileSize is the total size of the OTA software image in bytes
  - otaChecksum `optional(string)` - Digest of the entire contents of the associated OTA Software Update Image under the OtaUrl attribute, encoded in base64 string representation. The digest SHALL have been computed using the algorithm specified in OtaChecksumType
  - otaChecksumType `optional(string)` - Numeric identifier as defined in IANA Named Information Hash Algorithm Registry for the type of otaChecksum. For example, a value of 1 would match the sha-256 identifier, which maps to the SHA-256 digest algorithm
  - releaseNotesURL `optional(string)` - URL that contains product specific web page that contains release notes for the device model.
  - schemaVersion: `optional(uint16)` - Schema version to support backward/forward compatability. Should be equal to 0 (default 0)
- In State:
  - `model/ModelVersion/value/<vid>/<pid>/<softwareVersion>`
  - `model/ModelVersions/value/<vid>/<pid>`
- Who can send:
  - Vendor with same vid who created the Model
- CLI command minimal:

```bash
dcld tx model add-model-version --vid=<uint16> --pid=<uint16> --softwareVersion=<uint32> --softwareVersionString=<string> --cdVersionNumber=<uint32>
--minApplicableSoftwareVersion=<uint32> --maxApplicableSoftwareVersion=<uint32> --from=<account>
```

- CLI command full:

```bash
dcld tx model add-model-version --vid=<uint16> --pid=<uint16> --softwareVersion=<uint32> --softwareVersionString=<string> --cdVersionNumber=<uint32>
--minApplicableSoftwareVersion=<uint32> --maxApplicableSoftwareVersion=<uint32>
--firmwareInformation=<string> --softwareVersionValid=<bool> --otaURL=<string> --otaFileSize=<string> --otaChecksum=<string> --otaChecksumType=<string> --releaseNotesURL=<string> 
--from=<account>
```

### EDIT_MODEL_VERSION

**Status: Implemented**

Edits an existing Model Software Version identified by a unique combination of `vid` (vendor ID) `pid` (product ID) and `softwareVersion`
by the vendor. If `account` was created with product ID ranges then the `pid` must fall within that specified range.

Only the fields listed below (except `vid` `pid` and `softwareVersion`)  can be edited.

All non-edited fields remain the same.

`otaURL` can be edited only if  `otaFileSize`, `otaChecksum` and `otaChecksumType` are already set.

- Parameters:
  - vid: `uint16` -  model vendor ID (positive non-zero)
  - pid: `uint16` -  model product ID (positive non-zero)
  - softwareVersion: `uint32` - model software version (positive non-zero)
  - softwareVersionValid `optional(bool)` - Flag to indicate whether the software version is valid or not (default true)
  - otaURL `optional(string)` - URL where to obtain the OTA image
  - maxApplicableSoftwareVersion `optional(uint32)` - MaxApplicableSoftwareVersion should specify the highest SoftwareVersion for which this image can be applied
  - minApplicableSoftwareVersion `optional(uint32)` - MinApplicableSoftwareVersion should specify the lowest SoftwareVersion for which this image can be applied
  - releaseNotesURL `optional(string)` - URL that contains product specific web page that contains release notes for the device model.
  - otaURL `optional(string)` - URL where to obtain the OTA image
  - otaFileSize `optional(string)`  - OtaFileSize is the total size of the OTA software image in bytes
  - otaChecksum `optional(string)` - Digest of the entire contents of the associated OTA Software Update Image under the OtaUrl attribute, encoded in base64 string representation. The digest SHALL have been computed using the algorithm specified in OtaChecksumType
  - schemaVersion: `optional(uint16)` - Schema version to support backward/forward compatability. Should be equal to 0 (default 0)

- In State: `model/ModelVersion/value/<vid>/<pid>/<softwareVersion>`
- Who can send:
  - Vendor associated with the same vid who created the Model
- CLI command:
  - `dcld tx model update-model-version --vid=<uint16> --pid=<uint16> --softwareVersion=<uint32> ... --from=<account>`

### DELETE_MODEL_VERSION

**Status: Implemented**

Deletes an existing Model Version identified by a unique combination of `vid` (vendor ID), `pid` (product ID) and `softwareVersion`
by the vendor account. If `account` was created with product ID ranges then the `pid` must fall within that specified range.

Model Version can be deleted only before it is certified.

- Parameters:
  - vid: `uint16` -  model version vendor ID (positive non-zero)
  - pid: `uint16` -  model version product ID (positive non-zero)
  - softwareVersion: `uint32` - model version software version (positive non-zero)
- In State: `model/ModelVersion/value/<vid>/<pid>/<softwareVersion>`
- Who can send:
  - Vendor account associated with the same vid who has created the model version
- CLI command:
  - `dcld tx model delete-model-version --vid=< uint16 > --pid=< uint16 > --softwareVersion=<uint32> --from=<account>`

### GET_MODEL

**Status: Implemented**

Gets a Model Info with the given `vid` (vendor ID) and `pid` (product ID).

- Parameters:
  - vid: `uint16` -  model vendor ID (positive non-zero)
  - pid: `uint16` -  model product ID (positive non-zero)
- CLI command:
  - `dcld query model get-model --vid=<uint16> --pid=<uint16>`
- REST API:
  - GET `/dcl/model/models/{vid}/{pid}`

### GET_MODEL_VERSION

**Status: Implemented**

Gets a Model Software Versions for the given `vid`, `pid` and `softwareVersion`.

- Parameters
  - vid: `uint16` -  model vendor ID (positive non-zero)
  - pid: `uint16` -  model product ID (positive non-zero)
  - softwareVersion: `uint32` - model software version (positive non-zero)
- CLI command:
  - `dcld query model model-version --vid=<uint16> --pid=<uint16> --softwareVersion=<uint32>`
- REST API:
  - GET `/dcl/model/versions/{vid}/{pid}/{softwareVersion}`

### GET_ALL_MODELS

**Status: Implemented**

Gets all Model Infos for all vendors.

Should be sent to trusted nodes only.

- Parameters:
  - Common pagination parameters (see [pagination-params](#common-pagination-parameters))
- CLI command:
  - `dcld query model all-models`
- REST API:
  - GET `/dcl/model/models`

### GET_ALL_VENDOR_MODELS

**Status: Implemented**

Gets all Model Infos by the given Vendor (`vid`).

- Parameters:
  - vid: `uint16` -  model vendor ID (positive non-zero)
- CLI command:
  - `dcld query model vendor-models --vid=<uint16>`
- REST API:
  - GET `/dcl/model/models/{vid}`

### GET_ALL_MODEL_VERSIONS

**Status: Implemented**

Gets all Model Software Versions for the given `vid` and `pid` combination.

- Parameters:
  - vid: `uint16` -  model vendor ID (positive non-zero)
  - pid: `uint16` -  model product ID (positive non-zero)
- CLI command:
  - `dcld query model all-model-versions --vid=<uint16> --pid=<uint16>`
- REST API:
  - GET `/dcl/model/versions/{vid}/{pid}`

<!-- markdownlint-enable MD036 -->
