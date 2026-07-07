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

const (
	FlagVID                       = "vid"
	FlagPID                       = "pid"
	FlagSoftwareVersion           = "softwareVersion"
	FlagSoftwareVersionShortcut   = "v"
	FlagSoftwareVersionString     = "softwareVersionString"
	FlagCertificationType         = "certificationType"
	FlagCertificationTypeShortcut = "t"
	FlagCertificationDate         = "certificationDate"
	FlagDateShortcut              = "d"
	FlagIsProvisional             = "isProvisional"
	FlagRevocationDate            = "revocationDate"
	FlagReason                    = "reason"
	FlagReasonShortcut            = "r"
	FlagCDVersionNumber           = "cdVersionNumber"
	FlagProvisionalDate           = "provisionalDate"
	FlagProgramTypeVersion        = "programTypeVersion"
	FlagCDCertificateID           = "cdCertificateId"
	FlagFamilyID                  = "familyId"
	FlagSupportedClusters         = "supportedClusters"
	// Deprecated: The flag is kept for backward compatibility and will be removed in a future release.
	FlagCompliantPlatformUsed = "compliantPlatformUsed"
	// Deprecated: The flag is kept for backward compatibility and will be removed in a future release.
	FlagCompliantPlatformVersion = "compliantPlatformVersion"
	FlagCertificationRoute       = "certificationRoute"
	FlagProgramType              = "programType"
	FlagTransport                = "transport"
	FlagParentChild              = "parentChild"
	// Deprecated: The flag is kept for backward compatibility and will be removed in a future release.
	FlagCertificationIDOfSoftwareComponent = "certificationIDOfSoftwareComponent"
	FlagSpecificationVersion               = "specificationVersion"
	FlagOwner                              = "owner"
	// Deprecated: OSVersion is superseded by specificationVersion;
	// the flag is kept for backward compatibility and will be removed in a future release.
	FlagOSVersion = "OSVersion"

	TextVID                                          = "Model vendor ID (positive non-zero uint16)"
	TextPID                                          = "Model product ID (positive non-zero uint16)"
	TextSoftwareVersion                              = "Software Version of model (uint32)"
	TextSoftwareVersionString                        = "Software Version String of model. If the corresponding Model Version is present on the ledger, the value must match it."
	TextCDVersionNumber                              = "CD Version Number of the certification. If the corresponding Model Version is present on the ledger, the value must match it."
	TextCertificationType                            = "Certification program applied to the model. Supported values are 'zigbee', 'matter' or 'aliro'."
	TextProvisionalDate                              = "The date of model provisional certification (rfc3339 encoded), for example 2019-10-12T07:20:50.52Z"
	TextCertificationDate                            = "The date of model certification (rfc3339 encoded), for example 2019-10-12T07:20:50.52Z"
	TextRevocationDate                               = "The date of model revocation (rfc3339 encoded), for example 2019-10-12T07:20:50.52Z"
	TextProvisionalReason                            = "Optional comment describing the reason of provisioning"
	TextCertificationReason                          = "Optional comment describing the reason of certification"
	TextRevocationReason                             = "Optional comment describing the reason of revocation"
	TextCDCertificateID                              = "Connectivity Standards Alliance certification’s certificate ID for the Certification that applies to this record. The value of this field is used in the Certification Declaration's certificate_id field for products using the VendorID, ProductID and SoftwareVersion in this schema entry. Must be exactly 19 characters."
	TextFamilyID                                     = "Product family to which the certified model belongs. Typical family IDs have the prefix FAM followed by a sequence of alphanumeric characters (e.g. FAM123456). Only alphanumeric characters are allowed."
	TextSupportedClusters                            = "Application cluster IDs supported by the device, as hexadecimal numbers in a comma-separated list. For example, for an Extended Color Light (implementing Matter 1.5) this field would contain (at least) 0x0003,0x0004,0x0006,0x0008,0x0062,0x0300. Each cluster ID must be 0x followed by 1-4 hexadecimal digits; comma-separated, no duplicates."
	TextCertificationRoute                           = "Various certification paths, such as Fully Tested, Certification by Similarity, Family/Portfolio Certification, Certification Transfer etc. Supported values are fullTested, similarity, rapid-recert, fastTrack, ctp, family, and portfolio. Note that some values could be added or removed in the future."
	TextProgramType                                  = "Product type. Supported values are endProduct, softwareComponent or compliantPlatform"
	TextProgramTypeVersion                           = "Version of programType (see `programType` flag for supported programs). Can only be used when programType is provided."
	TextTransport                                    = "Underlying communication technology the device uses to connect and exchange data. Supported transports are thread, wi-fi, ethernet, bluetooth and nfc. When multiple transports supported - should be used with comma-separator (e.g. wi-fi,ethernet,bluetooth). Duplicate values are not allowed."
	TextParentChild                                  = "Parent vs. child characteristic when using the Product Family Certification or Portfolio Certification Program. Supported values are `parent` and `child`"
	TextOwner                                        = "Key to sign the transaction"
	TextSpecificationVersion                         = "Specification version applicable to the device model, and it matches the SpecificationVersion attribute in the Basic Information Cluster of a device running the software certified by this DeviceModel record. For example, for version 1.4.2.0 this field should contain 0x01040200."
	TextCertificationIDOfSoftwareComponent           = "Certification ID of software component."
	DeprecatedTextCertificationIDOfSoftwareComponent = "Flag 'certificationIDOfSoftwareComponent' will be removed in a future release."
	TextSchemaVersion                                = "Schema version - default is 1, the value should be equal to 1"
	TextOSVersion                                    = "Name and version of operating system separated by whitespace. For example, Android 16 or iOS 26.4."
	DeprecatedTextOSVersion                          = "Flag 'OSVersion' will be removed in a future release."
	TextCompliantPlatformUsed                        = "Certification ID of the compliant platform used with the product."
	DeprecatedTextCompliantPlatformUsed              = "Flag 'compliantPlatformUsed' will be removed in a future release."
	TextCompliantPlatformVersion                     = "Certified firmware version of Compliant Platform."
	DeprecatedTextCompliantPlatformVersion           = "Flag 'compliantPlatformVersion' will be removed in a future release."
)
