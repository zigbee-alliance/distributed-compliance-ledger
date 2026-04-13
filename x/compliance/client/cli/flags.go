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
	FlagVID                                = "vid"
	FlagPID                                = "pid"
	FlagSoftwareVersion                    = "softwareVersion"
	FlagSoftwareVersionShortcut            = "v"
	FlagSoftwareVersionString              = "softwareVersionString"
	FlagCertificationType                  = "certificationType"
	FlagCertificationTypeShortcut          = "t"
	FlagCertificationDate                  = "certificationDate"
	FlagDateShortcut                       = "d"
	FlagIsProvisional                      = "isProvisional"
	FlagRevocationDate                     = "revocationDate"
	FlagReason                             = "reason"
	FlagReasonShortcut                     = "r"
	FlagCDVersionNumber                    = "cdVersionNumber"
	FlagProvisionalDate                    = "provisionalDate"
	FlagCertificationTypeVersion           = "certificationTypeVersion"
	FlagCDCertificateID                    = "cdCertificateId"
	FlagFamilyID                           = "familyId"
	FlagSupportedClusters                  = "supportedClusters"
	FlagCompliantPlatformUsed              = "compliantPlatformUsed"
	FlagCompliantPlatformVersion           = "compliantPlatformVersion"
	FlagOSNameAndVersion                   = "OSNameAndVersion"
	FlagCertificationRoute                 = "certificationRoute"
	FlagProductType                        = "productType"
	FlagTransport                          = "transport"
	FlagParentChild                        = "parentChild"
	FlagOwner                              = "owner"
	FlagSoftwareVersionCertificationStatus = "softwareVersionCertificationStatus"

	TextVID                      = "Model vendor ID (positive non-zero uint16)"
	TextPID                      = "Model product ID (positive non-zero uint16)"
	TextSoftwareVersion          = "Software Version of model (uint32)"
	TextSoftwareVersionString    = "Software Version String of model"
	TextCDVersionNumber          = "CD Version Number of the certification"
	TextCertificationType        = "Certification type - Currently 'zigbee' and 'matter' types are supported"
	TextProvisionalDate          = "The date of model provisional certification (rfc3339 encoded), for example 2019-10-12T07:20:50.52Z"
	TextCertificationDate        = "The date of model certification (rfc3339 encoded), for example 2019-10-12T07:20:50.52Z"
	TextRevocationDate           = "The date of model revocation (rfc3339 encoded), for example 2019-10-12T07:20:50.52Z"
	TextProvisionalReason        = "Optional comment describing the reason of provisioning"
	TextCertificationReason      = "Optional comment describing the reason of certification"
	TextRevocationReason         = "Optional comment describing the reason of revocation"
	TextCertificationTypeVersion = "Version of the certification program (see `certificationType` for supported programs)"
	TextCDCertificateID          = "Connectivity Standards Alliance certification's certificate ID applied to the model certification"
	TextFamilyID                 = "Product family to which the certified model belongs. The possible value should start with the prefix `FAM` and be followed by an alphanumeric character (e.g. `FAM123456`)"
	TextSupportedClusters        = "Cluster IDs supported by the application. Supported cluster IDs are `0x0003`, `0x0004`, `0x0006`, `0x0062`, `0x0008`, and `0x0406`"
	TextCompliantPlatformUsed    = "Certification ID of the compliant platform used with the product"
	TextCompliantPlatformVersion = "Version of the used compliant platform (see `FlagCompliantPlatformUsed` for compliant platform)"
	TextOSNameAndVersion         = "Operating system name and version running on the device at the time of certification"
	TextCertificationRoute       = "Certification Route of the certification. Supported values: `zigbee`, `matter`"
	TextProductType              = "Product type. Supported values are `endProduct`, `softwareComponent` or `compliantPlatform`"
	TextTransport                = "Underlying communication technology the device uses to connect and exchange data. Supported values are `thread`, `wi-fi`, `ethernet`, and `bluetooth`"
	TextParentChild              = "Parent vs. child characteristic when using the Product Family Certification or Portfolio Certification Program. Supported values are `parent` and `child`"
	TextOwner                    = "Key to sign the transaction"
	TextSchemaVersion            = "Schema version"
)
