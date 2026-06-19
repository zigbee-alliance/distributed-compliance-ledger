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

package testconstants

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module/testutil"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/common/types"
)

func strToPubKey(pkStr string, cdc codec.Codec) cryptotypes.PubKey {
	var pk cryptotypes.PubKey
	if err := cdc.UnmarshalInterfaceJSON([]byte(pkStr), &pk); err != nil {
		panic(err)
	}

	return pk
}

var (
	// default context
	// TODO issue 99: design test context better.
	defEncConfig = testutil.MakeTestEncodingConfig()

	// Base constants.
	JackAccount  = "jack"
	AliceAccount = "alice"
	BobAccount   = "bob"
	AnnaAccount  = "anna"
	ChainID      = "dclchain"
	AccountName  = JackAccount
	Passphrase   = "test1234"
	EmptyString  = ""

	// Model Info.
	Vid                                        int32  = 1
	VendorName                                        = "Vendor Name"
	CompanyLegalName                                  = "Legal Company Name"
	CompanyPreferredName                              = "Company Preferred Name"
	VendorLandingPageURL                              = "https://www.example.com"
	Pid                                        int32  = 22
	DeviceTypeID                               int32  = 12345
	Version                                           = "1.0"
	ProductName                                       = "Device Name"
	ProductLabel                                      = "Product Label and/or Product Description"
	PartNumber                                        = "RCU2205A"
	SoftwareVersion                            uint32 = 1
	SoftwareVersionString                             = "1.0"
	HardwareVersion                            uint32 = 21
	HardwareVersionString                             = "2.1"
	CdVersionNumber                            int32  = 312
	FirmwareInformation                               = "Firmware Information String"
	Revoked                                           = false
	SoftwareVersionValid                              = true
	OtaURL                                            = "https://ota.firmware.com"
	OtaFileSize                                uint64 = 12345678
	OtaChecksum                                       = "MjFiZmYxN2YyMTRlMGJiMGMwNzhlNzIzOGIxZWE1ODk1Mzg4MjA3ZmFhNmM2NTg2YTBmNDU0MDk3YTU0ZWIzMw==" //nolint:lll
	OtaChecksumType                            int32  = 1
	OtaBlob                                           = "OTABlob Text"
	CommissioningCustomFlow                    int32  = 1
	CommissioningCustomFlowURL                        = "https://sampleflowurl.dclmodel"
	CommissioningModeInitialStepsHint          uint32 = 2
	CommissioningModeInitialStepsInstruction          = "commissioningModeInitialStepsInstruction details"
	CommissioningModeSecondaryStepsHint        uint32 = 3
	CommissioningModeSecondaryStepsInstruction        = "commissioningModeSecondaryStepsInstruction steps"
	IcdUserActiveModeTriggerHint               uint32 = 5
	IcdUserActiveModeTriggerInstruction               = "icdUserActiveModeTriggerInstruction steps"
	FactoryResetStepsHint                      uint32 = 4
	FactoryResetStepsInstruction                      = "factoryResetStepsInstruction steps"
	ReleaseNotesURL                                   = "https://url.releasenotes.dclmodel"
	UserManualURL                                     = "https://url.usermanual.dclmodel"
	SupportURL                                        = "https://url.supporturl.dclmodel"
	ProductURL                                        = "https://url.producturl.dclmodel"
	EnhancedSetupFlowTCURL                            = "https://url.enhansedsetupflowurl.dclmodel"
	EnhancedSetupFlowTCRevision                       = 1
	EnhancedSetupFlowTCDigest                         = "MmNmMjRkYmE1ZmIwYTMwZTI2ZTgzYjJhYzViOWUyOWUxYjE2MWU1YzFmYTc0MjVlNzMwNDMzNjI5MzhiOTgyNA=="
	EnhancedSetupFlowTCFileSize                       = 1
	MaintenanceURL                                    = "https://url.maintenanceurl.dclmodel"
	CommissioningFallbackURL                          = "https://url.commissioningfallbackurl.dclmodel"
	DiscoveryCapabilitiesBitmask               uint32 = 0
	LsfURL                                            = "https://url.lsfurl.dclmodel"
	DataURL                                           = "https://url.data.dclmodel"
	DataURL2                                          = "https://url.data.dclmodel2"
	URLWithoutProtocol                                = "url.dclmodel"
	URLStartsWithW3                                   = "www.example.org/path/to/file.txt"
	LsfRevision                                int32  = 1
	EnhancedSetupFlowOptions                   int32  = 1
	EmptyLsfRevision                           int32
	ChipBlob                                          = "Chip Blob Text"
	VendorBlob                                        = "Vendor Blob Text"
	MinApplicableSoftwareVersion               uint32 = 1
	MaxApplicableSoftwareVersion               uint32 = 1000
	Owner                                             = Address1
	Info                                              = "Information for Proposal/Approval/Revoke"
	Info2                                             = "Alternate Text 2 for Information for Proposal/Approval"
	Info3                                             = "Alternate Text 3 for Information for Proposal/Approval"
	Time                                       int64  = 1645809254
	Time2                                      int64  = 1645809261
	Time3                                      int64  = 1645809278

	// Compliance.
	ProvisionalDate                    = "2019-12-12T00:00:00Z"
	CertificationDate                  = "2020-01-01T00:00:00Z"
	RevocationDate                     = "2020-03-03T03:30:00Z"
	Reason                             = "Some Reason"
	RevocationReason                   = "Some Reason"
	CertificationType                  = "zigbee"
	CDCertificateID                    = "12345678910abcdefgh"
	FamilyID                           = "FAM123456abc"
	SupportedClusters                  = "0x0003,0x0004,0x0006,0x0008,0x0062,0x0300"
	CompliantPlatformUsed              = "Some Compliance Platform Used"
	CompliantPlatformVersion           = "Some Compliance Platform Version"
	OSVersion                          = "Some OS Version"
	CertificationRoute                 = "fullTested"
	ProgramType                        = "endProduct"
	ProgramTypeVersion                 = "1.4.2"
	Transport                          = "wi-fi,ethernet"
	SoftwareVersionCertificationStatus = uint32(
		3,
	)
	ParentChild1                       = "parent"
	ParentChild2                       = "child"
	CertificationIDOfSoftwareComponent = "some certification ID of software component"
	FirstJanuary                       = "2020-01-01T00:00:01Z"

	// Testing Result.
	TestResult = "http://test.result.com"
	TestDate   = "2020-02-02T02:00:00Z"

	// Upgrade.
	UpgradePlanNameV1_2_0 = "v1.2.0"
	UpgradePlanInfoV1_2_0 = "{\"binaries\":{\"linux/amd64\":\"https://github.com/zigbee-alliance/distributed-compliance-ledger/releases/download/v1.2.0/dcld?checksum=sha256:e4031c6a77aa8e58add391be671a334613271bcf6e7f11d23b04a0881ece6958\"}}"
	UpgradePlanNameV1_2_1 = "v1.2.1"
	UpgradePlanInfoV1_2_1 = "{\"binaries\":{\"linux/amd64\":\"https://github.com/zigbee-alliance/distributed-compliance-ledger/releases/download/v1.2.1/dcld?checksum=sha256:e4031c6a77aa8e58add391be671a334613271bcf6e7f11d23b04a0881ece6958\"}}"
	UpgradePlanNameV1_2_2 = "v1.2.2"
	UpgradePlanInfoV1_2_2 = "{\"binaries\":{\"linux/amd64\":\"https://github.com/zigbee-alliance/distributed-compliance-ledger/releases/download/v1.2.2/dcld?checksum=sha256:e4031c6a77aa8e58add391be671a334613271bcf6e7f11d23b04a0881ece6958\"}}"
	UpgradePlanNameV1_4_0 = "v1.4.0"
	UpgradePlanInfoV1_4_0 = "{\"binaries\":{\"linux/amd64\":\"https://github.com/zigbee-alliance/distributed-compliance-ledger/releases/download/v1.4.0/dcld?checksum=sha256:e4031c6a77aa8e58add391be671a334613271bcf6e7f11d23b04a0881ece6958\"}}"
	UpgradePlanNameV1_4_1 = "v1.4.1"
	UpgradePlanInfoV1_4_1 = "{\"binaries\":{\"linux/amd64\":\"https://github.com/zigbee-alliance/distributed-compliance-ledger/releases/download/v1.4.1/dcld?checksum=sha256:e4031c6a77aa8e58add391be671a334613271bcf6e7f11d23b04a0881ece6958\"}}"
	UpgradePlanNameV1_4_2 = "v1.4.2"
	UpgradePlanInfoV1_4_2 = "{\"binaries\":{\"linux/amd64\":\"https://github.com/zigbee-alliance/distributed-compliance-ledger/releases/download/v1.4.2/dcld?checksum=sha256:e4031c6a77aa8e58add391be671a334613271bcf6e7f11d23b04a0881ece6958\"}}"

	UpgradePlanName                 = "v1.4.4"
	UpgradePlanHeight         int64 = 1337
	UpgradePlanInfo                 = "{\"binaries\":{\"linux/amd64\":\"https://github.com/zigbee-alliance/distributed-compliance-ledger/releases/download/v1.4.4/dcld?checksum=sha256:e4031c6a77aa8e58add391be671a334613271bcf6e7f11d23b04a0881ece6958\"}}"
	UpgradeBrowserDownloadURL       = "https://github.com/zigbee-alliance/distributed-compliance-ledger/releases/download/v1.4.4/dcld"
	UpgradeGitAPIJSONResponse       = "{\"assets\":[{\"name\": \"dcld\", \"state\": \"uploaded\", \"digest\": \"sha256:e4031c6a77aa8e58add391be671a334613271bcf6e7f11d23b04a0881ece6958\", \"browser_download_url\":\"" + UpgradeBrowserDownloadURL + "\"}]}"
	//
	Address1, _           = sdk.AccAddressFromBech32("cosmos1s5xf3aanx7w84hgplk9z3l90qfpantg6nsmhpf")
	Address2, _           = sdk.AccAddressFromBech32("cosmos1nl4uaesk9gtu7su3n89lne6xpa6lq8gljn79rq")
	Address3, _           = sdk.AccAddressFromBech32("cosmos12r9vsus5js32pvnayt33zhcd4y9wcqcly45gr9")
	Address4, _           = sdk.AccAddressFromBech32("cosmos1vvwldfef3yuggm7ge9p34d6dvpz5s74nus6n7g")
	VendorID1       int32 = 1000
	VendorID2       int32 = 2000
	VendorID3       int32 = 3000
	VendorID4       int32 = 4000
	ProductIDsEmpty []*types.Uint16Range
	ProductIDsFull  = append([]*types.Uint16Range{}, &types.Uint16Range{Min: 1, Max: 65535})
	ProductIDs100   = append([]*types.Uint16Range{}, &types.Uint16Range{Min: 1, Max: 100})
	ProductIDs200   = append([]*types.Uint16Range{}, &types.Uint16Range{Min: 1, Max: 100}, &types.Uint16Range{Min: 101, Max: 200})
	PubKey1         = strToPubKey(
		`{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"Aw1XXHQ8i6JVNKsFQ9eQArJVt2GXEO0EBFsQL6XJ5BxY"}`,
		defEncConfig.Codec,
	)
	PubKey2 = strToPubKey(
		`{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"A2wJ7uOEE5Zm04K52czFTXfDj1qF2mholzi1zOJVlKlr"}`,
		defEncConfig.Codec,
	)
	PubKey3 = strToPubKey(
		`{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"A0GnKr6hItYE1A7dzoxNSMwMZuu1zauOLWAqJWen1RzF"}`,
		defEncConfig.Codec,
	)
	PubKey4 = strToPubKey(
		`{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"AnQC2MkMN1TOQyAJ0zjakPDcak+5FLtEoL4yBsgFO8Xe"}`,
		defEncConfig.Codec,
	)
	Signer           = Address1
	ValidatorPubKey1 = strToPubKey(
		`{"@type":"/cosmos.crypto.ed25519.PubKey","key":"1e+1/jHGaJi0b2zgCN46eelKCYpKiuTgPN18mL3fzx8="}`,
		defEncConfig.Codec,
	)
	ValidatorPubKey2 = strToPubKey(
		`{"@type":"/cosmos.crypto.ed25519.PubKey","key":"NB8hcdxKYDCaPWR67OiUXUSltZfYYOWYryPDUdbWRlA="}`,
		defEncConfig.Codec,
	)
	ValidatorAddress1 = "cosmosvaloper156dzj776tf3lmsahgmtnrphflaqf7n58kug5qe"
	ValidatorAddress2 = "cosmosvaloper12tg2p3rjsaczddufmsjjrw9nvhg8wkc4hcz3zw"
	ValidHTTPSURL     = "https://valid.url.com"
	ValidHTTPURL      = "http://valid.url.com"
	NotAValidURL      = "not a valid url"

	// CertSchemaVersion schema version of certificate.
	CertSchemaVersion uint32

	// SchemaVersion initial default value.
	SchemaVersion        uint32
	SpecificationVersion uint32 = 0x01040200
)

/*
	Certificates are taken from dsr-corporation.com
*/

const (
	StubCertPem = `pem certificate`

	RootCertPem = `
-----BEGIN CERTIFICATE-----
MIIB0DCCAXWgAwIBAgIUDj4iUMmpY5DIU3yCgcK6E9B+4fUwCgYIKoZIzj0EAwIw
NDELMAkGA1UEBhMCQVUxEzARBgNVBAgTCnNvbWUtc3RhdGUxEDAOBgNVBAoTB3Jv
b3QtY2EwIBcNMjAwOTExMDk0MDM4WhgPNDc1ODA4MDgwOTQwMzhaMDQxCzAJBgNV
BAYTAkFVMRMwEQYDVQQIEwpzb21lLXN0YXRlMRAwDgYDVQQKEwdyb290LWNhMFkw
EwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEEs2oo/gqooonqJeGl58I4Du4FTxL7+cN
8XH4Yhn5ryOn2pCs+DqCoWNZehXc7vV84Y2YyT0H4+Eiv0eHgMWQD6NjMGEwDgYD
VR0PAQH/BAQDAgGGMA8GA1UdEwEB/wQFMAMBAf8wHQYDVR0OBBYEFN9Or7CMnDd4
GudTEsrkeGtIHq+wMB8GA1UdIwQYMBaAFN9Or7CMnDd4GudTEsrkeGtIHq+wMAoG
CCqGSM49BAMCA0kAMEYCIQCQlgzDlYMx5YoNdnZnWPGYznT4Nm9+DCgcpDyeexPm
dgIhAOlessgLGBB1DQyBabefuJJM1ToVf4ZrY4LE2+h7Or/g
-----END CERTIFICATE-----`

	IntermediateCertPem = `
-----BEGIN CERTIFICATE-----
MIIB2jCCAYCgAwIBAgIUVUIAvKkPpBwMWSV6Qmf6m2XCV8swCgYIKoZIzj0EAwIw
NDELMAkGA1UEBhMCQVUxEzARBgNVBAgTCnNvbWUtc3RhdGUxEDAOBgNVBAoTB3Jv
b3QtY2EwIBcNMjAwOTExMDk0MDM4WhgPNDc1ODA4MDgwOTQwMzhaMDwxCzAJBgNV
BAYTAkFVMRMwEQYDVQQIEwpzb21lLXN0YXRlMRgwFgYDVQQKEw9pbnRlcm1lZGlh
dGUtY2EwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASJO5c8FeVaA3z6jK6/2gIo
HTPBCBpm8SETnMpFxBxLlOP/gJ5T6J/T80wgU9XLYCePZNWKeXugzf4owXQ/xCCY
o2YwZDAOBgNVHQ8BAf8EBAMCAYYwEgYDVR0TAQH/BAgwBgEB/wIBADAdBgNVHQ4E
FgQUG3MqkTRGipAqhxmR5L2PaTr5BHcwHwYDVR0jBBgwFoAU306vsIycN3ga51MS
yuR4a0ger7AwCgYIKoZIzj0EAwIDSAAwRQIgbJbBd0lVhXGz+ovgC60UQ9+A3hOC
Al5Cw8Y5bfMnzP4CIQCcztnBZWCpZWfzP5CSXXraj4Kn/wGLN0RpjDcLX3Dbog==
-----END CERTIFICATE-----`

	LeafCertPem = `
-----BEGIN CERTIFICATE-----
MIIB0DCCAXegAwIBAgIUR8MvYSChpIkgG5sw9x91igOzA2gwCgYIKoZIzj0EAwIw
PDELMAkGA1UEBhMCQVUxEzARBgNVBAgTCnNvbWUtc3RhdGUxGDAWBgNVBAoTD2lu
dGVybWVkaWF0ZS1jYTAgFw0yMDA5MTEwOTQwMzhaGA80NzU4MDgwODA5NDAzOFow
MTELMAkGA1UEBhMCQVUxEzARBgNVBAgTCnNvbWUtc3RhdGUxDTALBgNVBAoTBGxl
YWYwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATSDwF+M/6HSJy+xZlTveqaJmGF
1zg425OeU7f6GAd0lcaxB7v12W/Kb9RDRxGGy0LT5/+Pr7qt55IpJGndlM5Ao2Aw
XjAOBgNVHQ8BAf8EBAMCB4AwDAYDVR0TAQH/BAIwADAdBgNVHQ4EFgQUKjGNOW5Q
2pbflcWYg2jwWLIVszowHwYDVR0jBBgwFoAUG3MqkTRGipAqhxmR5L2PaTr5BHcw
CgYIKoZIzj0EAwIDRwAwRAIgSeacsa1ZrO9JAurbiKGBtpZ39aV26ewb8oCDyy1+
VCcCICd1T061CSnglx4ZPNrL4JLQ/H0CuQxh0mDVcV1sjHlY
-----END CERTIFICATE-----`

	GoogleCertPem = `
-----BEGIN CERTIFICATE-----
MIIB7TCCAZOgAwIBAgIBATAKBggqhkjOPQQDAjBLMQswCQYDVQQGEwJVUzEPMA0G
A1UECgwGR29vZ2xlMRUwEwYDVQQDDAxNYXR0ZXIgUEFBIDExFDASBgorBgEEAYKi
fAIBDAQ2MDA2MCAXDTIxMTIwODIwMjYwM1oYDzIxMjExMjA4MjAyNjAzWjBLMQsw
CQYDVQQGEwJVUzEPMA0GA1UECgwGR29vZ2xlMRUwEwYDVQQDDAxNYXR0ZXIgUEFB
IDExFDASBgorBgEEAYKifAIBDAQ2MDA2MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcD
QgAE8iZX+exx8NDV7jYKorx3EcsD1gessexUTSimIfvFI2PySlReMjJDVCGIzXor
hTYFOzwMAx4b6ogNMIUmcW7uT6NmMGQwEgYDVR0TAQH/BAgwBgEB/wIBATAOBgNV
HQ8BAf8EBAMCAQYwHQYDVR0OBBYEFLAAVoG4iGKJYoDhIRihqL4J3pMhMB8GA1Ud
IwQYMBaAFLAAVoG4iGKJYoDhIRihqL4J3pMhMAoGCCqGSM49BAMCA0gAMEUCIQCV
c26cVlyqjhQfcgN3udpne6zZQdyVMNLRWZn3EENBkAIgasUeFU8zaUt8bKNWd0k+
4RQp5Cp5wYzrE8AxJ9BiA/E=
-----END CERTIFICATE-----`

	TestCertPem = `
-----BEGIN CERTIFICATE-----
MIIBvDCCAWKgAwIBAgIGAX+LduKHMAoGCCqGSM49BAMCMDAxGDAWBgNVBAMMD01h
dHRlciBUZXN0IFBBQTEUMBIGCisGAQQBgqJ8AgEMBDEyNUQwIBcNMjIwMzE1MDI0
NDU4WhgPMjEyMjAzMTUwMjQ0NThaMDAxGDAWBgNVBAMMD01hdHRlciBUZXN0IFBB
QTEUMBIGCisGAQQBgqJ8AgEMBDEyNUQwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNC
AAQJ5S9EoWLHKeQc8mfaoVHf0+xgv6kfVxVPm5kStvm1lToFvAGWeq2QqhXWZfcN
x2++l2fDriD0oXKnssJJ0hx5o2YwZDASBgNVHRMBAf8ECDAGAQH/AgEBMB8GA1Ud
IwQYMBaAFOKQjTacPKPBE7sJ4k3BzMWmZpHUMB0GA1UdDgQWBBTikI02nDyjwRO7
CeJNwczFpmaR1DAOBgNVHQ8BAf8EBAMCAQYwCgYIKoZIzj0EAwIDSAAwRQIhAPZJ
skxY48EcSnatPseu6GcuFZw/bE/7uvp/PknnofJVAiAFXbU9SkxGi+Lqqa4YQRx9
tpcQ/mhg7DECwutZLCxKyA==
-----END CERTIFICATE-----`

	PAACertWithNumericVid = `
-----BEGIN CERTIFICATE-----
MIIBvTCCAWSgAwIBAgIITqjoMYLUHBwwCgYIKoZIzj0EAwIwMDEYMBYGA1UEAwwP
TWF0dGVyIFRlc3QgUEFBMRQwEgYKKwYBBAGConwCAQwERkZGMTAgFw0yMTA2Mjgx
NDIzNDNaGA85OTk5MTIzMTIzNTk1OVowMDEYMBYGA1UEAwwPTWF0dGVyIFRlc3Qg
UEFBMRQwEgYKKwYBBAGConwCAQwERkZGMTBZMBMGByqGSM49AgEGCCqGSM49AwEH
A0IABLbLY3KIfyko9brIGqnZOuJDHK2p154kL2UXfvnO2TKijs0Duq9qj8oYShpQ
NUKWDUU/MD8fGUIddR6Pjxqam3WjZjBkMBIGA1UdEwEB/wQIMAYBAf8CAQEwDgYD
VR0PAQH/BAQDAgEGMB0GA1UdDgQWBBRq/SJ3H1Ef7L8WQZdnENzcMaFxfjAfBgNV
HSMEGDAWgBRq/SJ3H1Ef7L8WQZdnENzcMaFxfjAKBggqhkjOPQQDAgNHADBEAiBQ
qoAC9NkyqaAFOPZTaK0P/8jvu8m+t9pWmDXPmqdRDgIgI7rI/g8j51RFtlM5CBpH
mUkpxyqvChVI1A0DTVFLJd4=
-----END CERTIFICATE-----`
	PAACertWithNumericVidSubject                    = "MDAxGDAWBgNVBAMMD01hdHRlciBUZXN0IFBBQTEUMBIGCisGAQQBgqJ8AgEMBEZGRjE="
	PAACertWithNumericVidSubjectAsText              = "CN=Matter Test PAA,1.3.6.1.4.1.37244.2.1=FFF1"
	PAACertWithNumericVidSubjectKeyID               = "6A:FD:22:77:1F:51:1F:EC:BF:16:41:97:67:10:DC:DC:31:A1:71:7E"
	PAACertWithNumericVidSerialNumber               = "4ea8e83182d41c1c"
	PAACertWithNumericVidVid                  int32 = 65521
	PAACertWithNumericVidDifferentWhitespaces       = `
-----BEGIN CERTIFICATE-----
MIIBvTCCAWSgAwIBAgIITqjoMY
LUHBwwCgYIKoZIzj0EAwIwMDEYMBYGA1UEAwwP
TWF0dGVyIFRlc3QgUEFBMRQ
wEgYKKwYBBAGConwCAQwERkZGMTAgFw0yMTA2Mjgx
ND  IzNDNaGA85OTk5MTI
zMTIzNTk1OVowMDEYMBYGA1UEAwwPTWF0dGVyIFRlc3Qg
UEFBMRQwEgYKKwYBBAGConwCAQwERkZGMTBZMBMGByqGSM49AgEGCCqGSM49AwEH
A0IABLbLY3KIfyko9brIGqnZOuJDHK2p154kL2UXfvnO2TKijs0Duq9qj8oYShpQ
NUKWDUU/	MD8fGUIddR6Pjxqam3WjZjBkMBIGA1UdEwEB/wQIMAYBAf8CAQEwDgYD
VR0PAQH/BAQDAgEGMB0GA1Ud     DgQWBBRq/SJ3H1Ef7L8WQZdnENzcMaFxfjAfBgNV
HSMEGDAWgBRq/SJ3H1Ef7L8WQZdnENzcMaFxfjAKBggqhkjOPQQDAgNHADBEAiBQ
qoAC9NkyqaAFOPZTaK0P/8jvu8m+t9pWmDXPmqdRDgIgI7rI/g8j51RFtlM5CBpH
mUkpxyqvChVI1A0DTVFLJd4=
-----END CERTIFICATE-----`

	PAACertNoVid = `
-----BEGIN CERTIFICATE-----
MIIBkTCCATegAwIBAgIHC4+6qN2G7jAKBggqhkjOPQQDAjAaMRgwFgYDVQQDDA9N
YXR0ZXIgVGVzdCBQQUEwIBcNMjEwNjI4MTQyMzQzWhgPOTk5OTEyMzEyMzU5NTla
MBoxGDAWBgNVBAMMD01hdHRlciBUZXN0IFBBQTBZMBMGByqGSM49AgEGCCqGSM49
AwEHA0IABBDvAqgah7aBIfuo0xl4+AejF+UKqKgoRGgokUuTPejt1KXDnJ/3Gkzj
ZH/X9iZTt9JJX8ukwPR/h2iAA54HIEqjZjBkMBIGA1UdEwEB/wQIMAYBAf8CAQEw
DgYDVR0PAQH/BAQDAgEGMB0GA1UdDgQWBBR4XOcFuGuPTm/Hk6pgy0PqaWiC1TAf
BgNVHSMEGDAWgBR4XOcFuGuPTm/Hk6pgy0PqaWiC1TAKBggqhkjOPQQDAgNIADBF
AiEAue/bPqBqUuwL8B5h2u0sLRVt22zwFBAdq3mPrAX6R+UCIGAGHT411g2dSw1E
ja12EvfoXFguP8MS3Bh5TdNzcV5d
-----END CERTIFICATE-----`
	PAACertNoVidSubject      = "MBoxGDAWBgNVBAMMD01hdHRlciBUZXN0IFBBQQ=="
	PAACertNoVidSubjectKeyID = "78:5C:E7:05:B8:6B:8F:4E:6F:C7:93:AA:60:CB:43:EA:69:68:82:D5"

	PAACertWithNumericVid1 = `
-----BEGIN CERTIFICATE-----
MIIBvDCCAWKgAwIBAgIIUU31T4F/bycwCgYIKoZIzj0EAwIwMDEYMBYGA1UEAwwP
TWF0dGVyIFRlc3QgUEFBMRQwEgYKKwYBBAGConwCAQwERkZGMjAeFw0zMTA2Mjgx
NDIzNDNaFw0zMjA2MjcxNDIzNDJaMDAxGDAWBgNVBAMMD01hdHRlciBUZXN0IFBB
QTEUMBIGCisGAQQBgqJ8AgEMBEZGRjIwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNC
AAQjYtkTW7E7w26mNn1LTLN/93IZp/xgOrAGP9ye/8bPC166EyCmbvTjSBvKu+HM
UlUmt2r79bkb9Swzhg3GXxA5o2YwZDASBgNVHRMBAf8ECDAGAQH/AgEBMA4GA1Ud
DwEB/wQEAwIBBjAdBgNVHQ4EFgQUfx2q8kSYuYZoDqCPwYkh6EhInRcwHwYDVR0j
BBgwFoAUfx2q8kSYuYZoDqCPwYkh6EhInRcwCgYIKoZIzj0EAwIDSAAwRQIgbBOM
5XDzrA1cWOTHigR4go96aos4+W1pA8Irlj2LxRcCIQDje3+2Gvz7UW9rRkf8p/SG
NbKsuLiNm8I5idctQg3eaw==
-----END CERTIFICATE-----`
	PAACertWithNumericVid1Subject      = "MDAxGDAWBgNVBAMMD01hdHRlciBUZXN0IFBBQTEUMBIGCisGAQQBgqJ8AgEMBEZGRjI="
	PAACertWithNumericVid1SubjectKeyID = "7F:1D:AA:F2:44:98:B9:86:68:0E:A0:8F:C1:89:21:E8:48:48:9D:17"
	PAACertWithNumericVid1Vid          = 65522

	PAICertWithNumericPidVid = `
-----BEGIN CERTIFICATE-----
MIIB1DCCAXqgAwIBAgIIPmzmUJrYQM0wCgYIKoZIzj0EAwIwMDEYMBYGA1UEAwwP
TWF0dGVyIFRlc3QgUEFBMRQwEgYKKwYBBAGConwCAQwERkZGMTAgFw0yMTA2Mjgx
NDIzNDNaGA85OTk5MTIzMTIzNTk1OVowRjEYMBYGA1UEAwwPTWF0dGVyIFRlc3Qg
UEFJMRQwEgYKKwYBBAGConwCAQwERkZGMTEUMBIGCisGAQQBgqJ8AgIMBDgwMDAw
WTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASA3fEbIo8+MfY7z1eY2hRiOuu96C7z
eO6tv7GP4avOMdCO1LIGBLbMxtm1+rZOfeEMt0vgF8nsFRYFbXDyzQsio2YwZDAS
BgNVHRMBAf8ECDAGAQH/AgEAMA4GA1UdDwEB/wQEAwIBBjAdBgNVHQ4EFgQUr0K3
CU3r1RXsbs8zuBEVIl8yUogwHwYDVR0jBBgwFoAUav0idx9RH+y/FkGXZxDc3DGh
cX4wCgYIKoZIzj0EAwIDSAAwRQIhAJbJyM8uAYhgBdj1vHLAe3X9mldpWsSRETET
i+oDPOUDAiAlVJQ75X1T1sR199I+v8/CA2zSm6Y5PsfvrYcUq3GCGQ==
-----END CERTIFICATE-----`

	PAICertWithNumericPidVidSubject       = "MEYxGDAWBgNVBAMMD01hdHRlciBUZXN0IFBBSTEUMBIGCisGAQQBgqJ8AgEMBEZGRjExFDASBgorBgEEAYKifAICDAQ4MDAw"
	PAICertWithNumericPidVidSubjectAsText = "CN=Matter Test PAI,1.3.6.1.4.1.37244.2.1=FFF1,1.3.6.1.4.1.37244.2.2=8000"
	PAICertWithNumericPidVidSubjectKeyID  = "AF:42:B7:09:4D:EB:D5:15:EC:6E:CF:33:B8:11:15:22:5F:32:52:88"
	PAICertWithNumericPidVidVid           = 65521
	PAICertWithNumericPidVidPid           = 32768
	PAICertWithNumericPidVidSerialNumber  = "4498223361705918669"

	PAICertWithPidVid = `
-----BEGIN CERTIFICATE-----
MIIBpjCCAUygAwIBAgIIVUDpotyYkzswCgYIKoZIzj0EAwIwGjEYMBYGA1UEAwwP
TWF0dGVyIFRlc3QgUEFBMCAXDTIxMDYyODE0MjM0M1oYDzk5OTkxMjMxMjM1OTU5
WjAuMSwwKgYDVQQDDCNNYXR0ZXIgVGVzdCBQQUkgTXZpZDpGRkYyIE1waWQ6ODAw
NDBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABNPzm8LjNudsiNIxOAibyWOQ6CsH
cUlCQlOb9XzaPM4rmR1jjepSu9xoMIO23HzEOIfByUP18+r0L1QvEn6mNwyjZjBk
MBIGA1UdEwEB/wQIMAYBAf8CAQAwDgYDVR0PAQH/BAQDAgEGMB0GA1UdDgQWBBQV
bs2MFL6AtBCu5AKj8jMX5zQGdDAfBgNVHSMEGDAWgBR4XOcFuGuPTm/Hk6pgy0Pq
aWiC1TAKBggqhkjOPQQDAgNIADBFAiEA7+WO/UkVZ4DGULOTLIItVhG7rC+mnqJI
fAuwib9kCRACIFaMCdDo/n+E+hOBXDXVemlbz0znMaLn/KcquoxDIfb7
-----END CERTIFICATE-----`
	PAICertWithPidVidVid = 65522
	PAICertWithPidVidPid = 32772

	PAICertWithNumericVid = `
-----BEGIN CERTIFICATE-----
MIIBqDCCAU6gAwIBAgIIPXS7VllxEBwwCgYIKoZIzj0EAwIwGjEYMBYGA1UEAwwP
TWF0dGVyIFRlc3QgUEFBMCAXDTIxMDYyODE0MjM0M1oYDzk5OTkxMjMxMjM1OTU5
WjAwMRgwFgYDVQQDDA9NYXR0ZXIgVGVzdCBQQUkxFDASBgorBgEEAYKifAIBDARG
RkYyMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE2L+TR5LPjq7awk/8lmyRdiD7
ly+6uY7G1RMUoHrpjhoD+0GR0m4tEny5UnYhw26XOhhsVtDK2ZmwQcJwqbHLP6Nm
MGQwEgYDVR0TAQH/BAgwBgEB/wIBADAOBgNVHQ8BAf8EBAMCAQYwHQYDVR0OBBYE
FGE90Ic1XvCLrgHkxpqPxz2sjH39MB8GA1UdIwQYMBaAFHhc5wW4a49Ob8eTqmDL
Q+ppaILVMAoGCCqGSM49BAMCA0gAMEUCIQDfwJ3oS/qVbWDW/vTirREL3iIqMogw
pn4/F7keUYUaeAIgce2XGOSIsrjPlUQ1zj/zLqUFVhQ8TyycBaIK8z7Uytk=
-----END CERTIFICATE-----`
	PAICertWithNumericVidVid          = 65522
	PAICertWithNumericVidSubject      = "MDAxGDAWBgNVBAMMD01hdHRlciBUZXN0IFBBSTEUMBIGCisGAQQBgqJ8AgEMBEZGRjI="
	PAICertWithNumericVidSubjectKeyID = "61:3D:D0:87:35:5E:F0:8B:AE:01:E4:C6:9A:8F:C7:3D:AC:8C:7D:FD"

	PAICertWithVid = `-----BEGIN CERTIFICATE-----
MIIBmzCCAUKgAwIBAgIIIt8JcSeGaqMwCgYIKoZIzj0EAwIwGjEYMBYGA1UEAwwP
TWF0dGVyIFRlc3QgUEFBMCAXDTIxMDYyODE0MjM0M1oYDzk5OTkxMjMxMjM1OTU5
WjAkMSIwIAYDVQQDDBlNYXR0ZXIgVGVzdCBQQUkgTXZpZDpGRkYyMFkwEwYHKoZI
zj0CAQYIKoZIzj0DAQcDQgAEanvx+yP6s3+8o1jmPZIEtwNXBr6tftPgSwsb/ipZ
VnUf6hVvlNNGzihh9Ouv2CKTwudkXMx6Wqy/sX5B5eJEgaNmMGQwEgYDVR0TAQH/
BAgwBgEB/wIBADAOBgNVHQ8BAf8EBAMCAQYwHQYDVR0OBBYEFBn6lHSTKdp0D1Vq
Mz+1E3PXCBdWMB8GA1UdIwQYMBaAFHhc5wW4a49Ob8eTqmDLQ+ppaILVMAoGCCqG
SM49BAMCA0cAMEQCIHZQ4Yv8BJhq6w3Gjhu8AZlvRLSwNLDYDI2UpothBjIDAiB4
/ryct/QEzO8ZXM8eywlUQ4vlpZ10iumuMTkNmxJb/g==
-----END CERTIFICATE-----`
	PAICertWithVidVid = 65522

	PAACertExpired = `-----BEGIN CERTIFICATE-----
MIIBvDCCAWKgAwIBAgIIAwFMXQkbuSQwCgYIKoZIzj0EAwIwMDEYMBYGA1UEAwwP
TWF0dGVyIFRlc3QgUEFBMRQwEgYKKwYBBAGConwCAQwERkZGMjAeFw0yMTA2Mjgx
NDIzNDNaFw0yMjA2MjgxNDIzNDJaMDAxGDAWBgNVBAMMD01hdHRlciBUZXN0IFBB
QTEUMBIGCisGAQQBgqJ8AgEMBEZGRjIwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNC
AAS3ChaWbeRtv7g4oICrIP92kJfHqKAh37RKtoeyBOARLTsQffxqcPdTKIHbdlbU
kCTiECbFVfuSLYkJzG8MBbHKo2YwZDASBgNVHRMBAf8ECDAGAQH/AgEBMA4GA1Ud
DwEB/wQEAwIBBjAdBgNVHQ4EFgQUm7KZu6ZJfbWgTh4Yx8IOrsVMQukwHwYDVR0j
BBgwFoAUm7KZu6ZJfbWgTh4Yx8IOrsVMQukwCgYIKoZIzj0EAwIDSAAwRQIgED1f
neH8IHXIiGiNf/knk09MXh43Cqmj8tAWJGNqtGECIQCjqZb5p6tl5zLGr4d70sVO
2GL5RcHKohx+qL0o57dAuQ==
-----END CERTIFICATE-----`

	PAACertWithSameSubjectID1 = `-----BEGIN CERTIFICATE-----
MIIB5jCCAYygAwIBAgICA+kwCgYIKoZIzj0EAwIwWjELMAkGA1UEBhMCVVoxDDAK
BgNVBAgTA1RTSDERMA8GA1UEBxMIVEFTSEtFTlQxDDAKBgNVBAoTA0RTUjELMAkG
A1UECxMCREMxDzANBgNVBAMTBk1BVFRFUjAeFw0yNTA0MjkxODMxMjVaFw0zNTA0
MjgxODMxMjVaMFoxCzAJBgNVBAYTAlVaMQwwCgYDVQQIEwNUU0gxETAPBgNVBAcT
CFRBU0hLRU5UMQwwCgYDVQQKEwNEU1IxCzAJBgNVBAsTAkRDMQ8wDQYDVQQDEwZN
QVRURVIwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATqwLpFXWRXqvrNHwiThBaQ
ItJnyqCJg4T9wVQvR1NsLpFKMAFQUvmu3yMwE2KR4QRP0WXbw36jk4jgSfISOgyD
o0IwQDAOBgNVHQ8BAf8EBAMCAYYwDwYDVR0TAQH/BAUwAwEB/zAdBgNVHQ4EFgQU
gbf4QxVsD2cwgNmSuTBsVVnT1mUwCgYIKoZIzj0EAwIDSAAwRQIhAKvskVz5ZaGP
tXbPxIh4g3iR2dpRZnB6fcJwZjirv69OAiBUPEZWcMnlgOoQnbXIiS1cP01Awqrs
YZuybMMjutqvcQ==
-----END CERTIFICATE-----`

	PAACertWithSameSubjectID2 = `-----BEGIN CERTIFICATE-----
MIIB8jCCAZigAwIBAgICA+owCgYIKoZIzj0EAwIwYDELMAkGA1UEBhMCVVoxDDAK
BgNVBAgTA1RTSDERMA8GA1UEBxMIVEFTSEtFTlQxDDAKBgNVBAoTA0RTUjEQMA4G
A1UECxMHTUFUVEVSMjEQMA4GA1UEAxMHTUFUVEVSMjAeFw0yNTA0MjkxODMxMjVa
Fw0zNTA0MjgxODMxMjVaMGAxCzAJBgNVBAYTAlVaMQwwCgYDVQQIEwNUU0gxETAP
BgNVBAcTCFRBU0hLRU5UMQwwCgYDVQQKEwNEU1IxEDAOBgNVBAsTB01BVFRFUjIx
EDAOBgNVBAMTB01BVFRFUjIwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATqwLpF
XWRXqvrNHwiThBaQItJnyqCJg4T9wVQvR1NsLpFKMAFQUvmu3yMwE2KR4QRP0WXb
w36jk4jgSfISOgyDo0IwQDAOBgNVHQ8BAf8EBAMCAYYwDwYDVR0TAQH/BAUwAwEB
/zAdBgNVHQ4EFgQUgbf4QxVsD2cwgNmSuTBsVVnT1mUwCgYIKoZIzj0EAwIDSAAw
RQIhAOWD4HmJJGKv1+sd02m39b2IM9HRthXHKAFvP3iw4ZI1AiBsrab/TxnXryYY
i9/bPMUvFdoff+dwG25rLD0AEurzxg==
-----END CERTIFICATE-----`

	RootCertWithSameSubjectAndSKID1 = `-----BEGIN CERTIFICATE-----
MIICODCCAd+gAwIBAgIBATAKBggqhkjOPQQDAjCBgjELMAkGA1UEBhMCVVMxETAP
BgNVBAgTCE5ldyBZb3JrMREwDwYDVQQHEwhOZXcgWW9yazEYMBYGA1UEChMPRXhh
bXBsZSBDb21wYW55MRkwFwYDVQQLExBUZXN0aW5nIERpdmlzaW9uMRgwFgYDVQQD
Ew93d3cuZXhhbXBsZS5jb20wIBcNMjQwMjE2MDY1OTAyWhgPMzAyMzA2MTkwNjU5
MDJaMIGCMQswCQYDVQQGEwJVUzERMA8GA1UECBMITmV3IFlvcmsxETAPBgNVBAcT
CE5ldyBZb3JrMRgwFgYDVQQKEw9FeGFtcGxlIENvbXBhbnkxGTAXBgNVBAsTEFRl
c3RpbmcgRGl2aXNpb24xGDAWBgNVBAMTD3d3dy5leGFtcGxlLmNvbTBZMBMGByqG
SM49AgEGCCqGSM49AwEHA0IABK1FfAnV0tE18vOeFWhjTQ5Ty7bQiHq1hW9yZI/2
JQdCeWa+rhwYS8yGLBc2jBRbbk+1Y3lJT68feOXU1nEPJL6jQjBAMA4GA1UdDwEB
/wQEAwIBhjAPBgNVHRMBAf8EBTADAQH/MB0GA1UdDgQWBBTBSGbtbyPYKBrZN3xY
rD/aBMFB6DAKBggqhkjOPQQDAgNHADBEAiANi2hxBoGgpxD7Y9Gr6273K5tODyOg
uIE5t5m3wUmziQIgLzRawTfn2ceA3s+H179zwWBU/di3RRU3enqPBzSiVGg=
-----END CERTIFICATE-----`

	RootCertWithSameSubjectAndSKID2 = `-----BEGIN CERTIFICATE-----
MIICOjCCAd+gAwIBAgIBAjAKBggqhkjOPQQDAjCBgjELMAkGA1UEBhMCVVMxETAP
BgNVBAgTCE5ldyBZb3JrMREwDwYDVQQHEwhOZXcgWW9yazEYMBYGA1UEChMPRXhh
bXBsZSBDb21wYW55MRkwFwYDVQQLExBUZXN0aW5nIERpdmlzaW9uMRgwFgYDVQQD
Ew93d3cuZXhhbXBsZS5jb20wIBcNMjQwMjE2MDY1OTAyWhgPMzAyMzA2MTkwNjU5
MDJaMIGCMQswCQYDVQQGEwJVUzERMA8GA1UECBMITmV3IFlvcmsxETAPBgNVBAcT
CE5ldyBZb3JrMRgwFgYDVQQKEw9FeGFtcGxlIENvbXBhbnkxGTAXBgNVBAsTEFRl
c3RpbmcgRGl2aXNpb24xGDAWBgNVBAMTD3d3dy5leGFtcGxlLmNvbTBZMBMGByqG
SM49AgEGCCqGSM49AwEHA0IABK1FfAnV0tE18vOeFWhjTQ5Ty7bQiHq1hW9yZI/2
JQdCeWa+rhwYS8yGLBc2jBRbbk+1Y3lJT68feOXU1nEPJL6jQjBAMA4GA1UdDwEB
/wQEAwIBhjAPBgNVHRMBAf8EBTADAQH/MB0GA1UdDgQWBBTBSGbtbyPYKBrZN3xY
rD/aBMFB6DAKBggqhkjOPQQDAgNJADBGAiEAyXs4797Z3LCZVdks82xGnljJttA2
J0Bd0tu1G7yEay4CIQC5iE5/pPI62yKKReLftNTO9mpFlhXEbQFwXIcdmZQZ5g==
-----END CERTIFICATE-----`

	IntermediateWithSameSubjectAndSKID1 = `-----BEGIN CERTIFICATE-----
MIICIDCCAcWgAwIBAgIBAzAKBggqhkjOPQQDAjCBgjELMAkGA1UEBhMCVVMxETAP
BgNVBAgTCE5ldyBZb3JrMREwDwYDVQQHEwhOZXcgWW9yazEYMBYGA1UEChMPRXhh
bXBsZSBDb21wYW55MRkwFwYDVQQLExBUZXN0aW5nIERpdmlzaW9uMRgwFgYDVQQD
Ew93d3cuZXhhbXBsZS5jb20wIBcNMjQwMjE2MDY1OTAyWhgPMzAyMzA2MTkwNjU5
MDJaMEUxCzAJBgNVBAYTAkFVMRMwEQYDVQQIEwpTb21lLVN0YXRlMSEwHwYDVQQK
ExhJbnRlcm5ldCBXaWRnaXRzIFB0eSBMdGQwWTATBgcqhkjOPQIBBggqhkjOPQMB
BwNCAASVi2ZPDYX8Tu5EJTYQgukEUngK5PHLQbgBBZ5YR5z6BvtPTjsvrA5qSGI3
4mYaGOzf6nl6cz/1MllMRnW8+mRho2YwZDAOBgNVHQ8BAf8EBAMCAYYwEgYDVR0T
AQH/BAgwBgEB/wIBADAdBgNVHQ4EFgQUoeCSifoYghIUnbiuGUO+RDFr8fUwHwYD
VR0jBBgwFoAUwUhm7W8j2Cga2Td8WKw/2gTBQegwCgYIKoZIzj0EAwIDSQAwRgIh
AJLVYTBo2Xoe3QK0Za83ont/6e35mBJk7MZE1/4KJkoiAiEAzMb7lkVaFsg75+d6
HKHWeuUXSca9LIQAhJ+dowZVT1A=
-----END CERTIFICATE-----`

	IntermediateWithSameSubjectAndSKID2 = `-----BEGIN CERTIFICATE-----
MIICHzCCAcWgAwIBAgIBBDAKBggqhkjOPQQDAjCBgjELMAkGA1UEBhMCVVMxETAP
BgNVBAgTCE5ldyBZb3JrMREwDwYDVQQHEwhOZXcgWW9yazEYMBYGA1UEChMPRXhh
bXBsZSBDb21wYW55MRkwFwYDVQQLExBUZXN0aW5nIERpdmlzaW9uMRgwFgYDVQQD
Ew93d3cuZXhhbXBsZS5jb20wIBcNMjQwMjE2MDY1OTAyWhgPMzAyMzA2MTkwNjU5
MDJaMEUxCzAJBgNVBAYTAkFVMRMwEQYDVQQIEwpTb21lLVN0YXRlMSEwHwYDVQQK
ExhJbnRlcm5ldCBXaWRnaXRzIFB0eSBMdGQwWTATBgcqhkjOPQIBBggqhkjOPQMB
BwNCAASVi2ZPDYX8Tu5EJTYQgukEUngK5PHLQbgBBZ5YR5z6BvtPTjsvrA5qSGI3
4mYaGOzf6nl6cz/1MllMRnW8+mRho2YwZDAOBgNVHQ8BAf8EBAMCAYYwEgYDVR0T
AQH/BAgwBgEB/wIBADAdBgNVHQ4EFgQUoeCSifoYghIUnbiuGUO+RDFr8fUwHwYD
VR0jBBgwFoAUwUhm7W8j2Cga2Td8WKw/2gTBQegwCgYIKoZIzj0EAwIDSAAwRQIg
b0dWC2ZE246iU72OqNBij3QfO37jxBdkIDg0lds2Y18CIQDA3t1dg/cUDPjAB+qJ
nmBj6GRxpq3EPYFWPsM7unmENw==
-----END CERTIFICATE-----`

	LeafCertWithSameSubjectAndSKID = `-----BEGIN CERTIFICATE-----
MIIB2zCCAYGgAwIBAgIBBTAKBggqhkjOPQQDAjBFMQswCQYDVQQGEwJBVTETMBEG
A1UECBMKU29tZS1TdGF0ZTEhMB8GA1UEChMYSW50ZXJuZXQgV2lkZ2l0cyBQdHkg
THRkMCAXDTI0MDIxNjA2NTkwMloYDzMwMjMwNjE5MDY1OTAyWjBFMQswCQYDVQQG
EwJBVTETMBEGA1UECBMKU29tZS1TdGF0ZTEhMB8GA1UEChMYSW50ZXJuZXQgV2lk
Z2l0cyBQdHkgTHRkMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE1vWE7i2oGi4E
O2C5yeH6y/oOf/X2PPmvHUZWjnpgXk6PjCUV0dZP1dfyvcpA56IQxn2vtFxwvd1c
2qlu8ln6baNgMF4wDgYDVR0PAQH/BAQDAgeAMAwGA1UdEwEB/wQCMAAwHQYDVR0O
BBYEFJCBhMfsuIEUZmEvgrvpUWfyTZmjMB8GA1UdIwQYMBaAFKHgkon6GIISFJ24
rhlDvkQxa/H1MAoGCCqGSM49BAMCA0gAMEUCIQDGTb+/1toLx9c7DsgtP435NmXx
wtFjk2980dRKbPQHEQIgWLO4bbDSe0ss9C3D8W5n9HUgBYQZ+AdenfYutUKqXJI=
-----END CERTIFICATE-----`

	LeafCertWithoutBasicConstraints = `-----BEGIN CERTIFICATE-----
MIIBkTCCATigAwIBAgIBCzAKBggqhkjOPQQDAjA5MQswCQYDVQQGEwJBVTEUMBIG
A1UEChMLTm9CQy1QYXJlbnQxFDASBgNVBAMTC05vQkMgcGFyZW50MCAXDTI0MDIx
NjA2NTkwMloYDzMwMjMwNjE5MDY1OTAyWjA1MQswCQYDVQQGEwJBVTESMBAGA1UE
ChMJTm9CQy1MZWFmMRIwEAYDVQQDEwlOb0JDIGxlYWYwWTATBgcqhkjOPQIBBggq
hkjOPQMBBwNCAATV8ZN2gIlQ9JnphS/Ir1B7HjSvOFebo8ZluWTlv2AZ2AAgMs05
Q6n/qj2SrafMTAHmNIN3xsFfcq5+RAhQBQO4ozMwMTAOBgNVHQ8BAf8EBAMCB4Aw
HwYDVR0jBBgwFoAU9l4vLGgmuT+d9AtN+AmSTyCQYpAwCgYIKoZIzj0EAwIDRwAw
RAIgMdB+xhNDkNjgX+uwi+5bWPTrI2pFoNeJMCTKDDyIC34CIH6hpZc3hANMnaOG
nT1SO4vp46SzmvdrGRUhDb7gB91i
-----END CERTIFICATE-----`

	// MatterDACShaped is a synthetic ECDSA P-256 leaf that satisfies the
	// Matter R1.6 §6.2.2.3 DAC structural profile: BasicConstraints critical
	// with cA=FALSE, KeyUsage critical with exactly digitalSignature, plus SKI
	// and AKI.
	MatterDACShaped = `-----BEGIN CERTIFICATE-----
MIIBpDCCAUqgAwIBAgIBCjAKBggqhkjOPQQDAjAnMQswCQYDVQQGEwJVUzEYMBYG
A1UEAxMPTWF0dGVyIFRlc3QgUEFJMCAXDTI0MDEwMTAwMDAwMFoYDzk5OTkxMjMx
MjM1OTU5WjAsMQswCQYDVQQGEwJVUzEdMBsGA1UEAxMUTWF0dGVyIFRlc3QgREFD
IDAwMDEwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASLC9D0DeE7L3FYX3s3KONn
XGPk+QpyWsTWJbXMkNKtag0LVrflO+WehVb1r3EHmiCGr4xBkVwU6YAbVTL9NMXp
o2AwXjAOBgNVHQ8BAf8EBAMCB4AwDAYDVR0TAQH/BAIwADAdBgNVHQ4EFgQUJTdb
TFxSXx874rark+MG2pyv4bAwHwYDVR0jBBgwFoAUj/enZ2zBUoVDPaG6NzY9FKne
ix8wCgYIKoZIzj0EAwIDSAAwRQIhALzjiaDd8HnkH0lOofEjfB/sAUE1SYcTYYi7
v/w6GteWAiAFUPvSjKa+YudRxdFhm8QmZJik8cS3gkUiCFO55F1DYg==
-----END CERTIFICATE-----`

	// MatterNOCShaped is a synthetic ECDSA P-256 leaf that satisfies the Matter
	// R1.6 §6.5.12 NOC extension profile: BasicConstraints critical with
	// is-ca=FALSE, KeyUsage critical with exactly digitalSignature,
	// ExtendedKeyUsage critical with exactly {serverAuth, clientAuth}, plus
	// SKI and AKI.
	MatterNOCShaped = `-----BEGIN CERTIFICATE-----
MIIBwzCCAWqgAwIBAgIBFTAKBggqhkjOPQQDAjAoMQswCQYDVQQGEwJVUzEZMBcG
A1UEAxMQTWF0dGVyIFRlc3QgUkNBQzAgFw0yNDAxMDEwMDAwMDBaGA85OTk5MTIz
MTIzNTk1OVowJzELMAkGA1UEBhMCVVMxGDAWBgNVBAMTD01hdHRlciBUZXN0IE5P
QzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABJnf+/VvcVaa9lIsssmAAqrZ47j4
2jD2OQ+1ovOk/IqOFZO3U+MhzmhvV/MR/gJ+XL0r9LP32vy0dhqmDJ2ozzGjgYMw
gYAwDgYDVR0PAQH/BAQDAgeAMAwGA1UdEwEB/wQCMAAwHQYDVR0OBBYEFEbBgVot
PJ6J16NBeYjlVTbOTx0TMB8GA1UdIwQYMBaAFOBfA24KwNTKMiqUD5Ey6J7IW6iI
MCAGA1UdJQEB/wQWMBQGCCsGAQUFBwMBBggrBgEFBQcDAjAKBggqhkjOPQQDAgNH
ADBEAiAuZaOUCjHwV3YN91cr1+8GDfweapLiJjFImKRzvToNZgIgBYXoYLHgLCSm
zifykV/VENlmobferpJekJEZ0//p7FY=
-----END CERTIFICATE-----`

	RootCertWithVid = `
-----BEGIN CERTIFICATE-----
MIIChjCCAiygAwIBAgIBATAKBggqhkjOPQQDAjCBmDELMAkGA1UEBhMCVVMxETAP
BgNVBAgTCE5ldyBZb3JrMREwDwYDVQQHEwhOZXcgWW9yazEYMBYGA1UEChMPRXhh
bXBsZSBDb21wYW55MRkwFwYDVQQLExBUZXN0aW5nIERpdmlzaW9uMRgwFgYDVQQD
Ew93d3cuZXhhbXBsZS5jb20xFDASBgorBgEEAYKifAIBEwRGRkYxMCAXDTI0MDIy
NjExNTQzMVoYDzMwMjMwNjI5MTE1NDMxWjCBmDELMAkGA1UEBhMCVVMxETAPBgNV
BAgTCE5ldyBZb3JrMREwDwYDVQQHEwhOZXcgWW9yazEYMBYGA1UEChMPRXhhbXBs
ZSBDb21wYW55MRkwFwYDVQQLExBUZXN0aW5nIERpdmlzaW9uMRgwFgYDVQQDEw93
d3cuZXhhbXBsZS5jb20xFDASBgorBgEEAYKifAIBEwRGRkYxMFkwEwYHKoZIzj0C
AQYIKoZIzj0DAQcDQgAE4BNRvFdOZD7xT0uJZZlox1DC5tDjIlM/0J37EbcJTbcZ
7kMqLiSlVVWpl/qTjw48QohUH6neKzhaLBkWz7UJE6NjMGEwDgYDVR0PAQH/BAQD
AgGGMA8GA1UdEwEB/wQFMAMBAf8wHQYDVR0OBBYEFGuMdx6ty6g8M5wvECdfQgMd
CvSOMB8GA1UdIwQYMBaAFGuMdx6ty6g8M5wvECdfQgMdCvSOMAoGCCqGSM49BAMC
A0gAMEUCIHERQ+c7EsIaLLfoZaDkFgFr04kh/BHl+jxNeBL0iNKyAiEAxE5RhBPW
qHETvDFmP+6SIbqM0FacO30f+mn2ux9aTSE=
-----END CERTIFICATE-----
`

	IntermediateCertWithVid1 = `
-----BEGIN CERTIFICATE-----
MIICnzCCAkWgAwIBAgIBAzAKBggqhkjOPQQDAjCBmDELMAkGA1UEBhMCVVMxETAP
BgNVBAgTCE5ldyBZb3JrMREwDwYDVQQHEwhOZXcgWW9yazEYMBYGA1UEChMPRXhh
bXBsZSBDb21wYW55MRkwFwYDVQQLExBUZXN0aW5nIERpdmlzaW9uMRgwFgYDVQQD
Ew93d3cuZXhhbXBsZS5jb20xFDASBgorBgEEAYKifAIBEwRGRkYxMCAXDTI0MDIy
NjExNTQzMVoYDzMwMjMwNjI5MTE1NDMxWjCBrjELMAkGA1UEBhMCVVMxETAPBgNV
BAgTCE5ldyBZb3JrMREwDwYDVQQHEwhOZXcgWW9yazEYMBYGA1UEChMPRXhhbXBs
ZSBDb21wYW55MRkwFwYDVQQLExBUZXN0aW5nIERpdmlzaW9uMRgwFgYDVQQDEw93
d3cuZXhhbXBsZS5jb20xFDASBgorBgEEAYKifAIBEwRGRkYxMRQwEgYKKwYBBAGC
onwCAhMERkZGMTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABNXrgyGcS03b065a
2WVLfYUxdBdRjVBbBYF4ExILOyCuT6xtMqU9hh6wU+4jqNwB75V5fV83mSBLMpTJ
yR/s4R2jZjBkMA4GA1UdDwEB/wQEAwIBhjASBgNVHRMBAf8ECDAGAQH/AgEAMB0G
A1UdDgQWBBSwez/xRQGRj8H67suaAQbHR5td7DAfBgNVHSMEGDAWgBRrjHcercuo
PDOcLxAnX0IDHQr0jjAKBggqhkjOPQQDAgNIADBFAiEAnZHGqtMcgCkNkeHtcBry
jAL/LxZacm4zgr967c0FGcACIBmtfLqHsddfy30XV/zpYgp9oUp1YJaPsDlQvwF2
grbH
-----END CERTIFICATE-----
`

	IntermediateCertWithVid2 = `
-----BEGIN CERTIFICATE-----
MIICnzCCAkWgAwIBAgIBBjAKBggqhkjOPQQDAjCBmDELMAkGA1UEBhMCVVMxETAP
BgNVBAgTCE5ldyBZb3JrMREwDwYDVQQHEwhOZXcgWW9yazEYMBYGA1UEChMPRXhh
bXBsZSBDb21wYW55MRkwFwYDVQQLExBUZXN0aW5nIERpdmlzaW9uMRgwFgYDVQQD
Ew93d3cuZXhhbXBsZS5jb20xFDASBgorBgEEAYKifAIBEwRGRkYxMCAXDTI0MDIy
NjExNTQzMVoYDzMwMjMwNjI5MTE1NDMxWjCBrjELMAkGA1UEBhMCVVMxETAPBgNV
BAgTCE5ldyBZb3JrMREwDwYDVQQHEwhOZXcgWW9yazEYMBYGA1UEChMPRXhhbXBs
ZSBDb21wYW55MRkwFwYDVQQLExBUZXN0aW5nIERpdmlzaW9uMRgwFgYDVQQDEw93
d3cuZXhhbXBsZS5jb20xFDASBgorBgEEAYKifAIBEwRGRkYyMRQwEgYKKwYBBAGC
onwCAhMERkZGMjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABGWTKZpQBQa55HiR
DlgRdbofXa6CxlDPjsflAaXNaxKhOFaTAv6cDyQtziF4xm85FJ/RWTi5Dh0n2s50
bkMzReOjZjBkMA4GA1UdDwEB/wQEAwIBhjASBgNVHRMBAf8ECDAGAQH/AgEAMB0G
A1UdDgQWBBT8XnQdu5V8lQ821tb3DpUFGFt+AzAfBgNVHSMEGDAWgBRrjHcercuo
PDOcLxAnX0IDHQr0jjAKBggqhkjOPQQDAgNIADBFAiBNxNjaKuZQONgD4ZuZJSs3
whGarynKSrofLSJYf9gydwIhAOvaYGXDmqPLYuGv/zHbYt7CRt5uV7DIWFPXapDG
9PKH
-----END CERTIFICATE-----
`

	IntermediateCertWithoutVidPid = `
-----BEGIN CERTIFICATE-----
MIICczCCAhmgAwIBAgIBBzAKBggqhkjOPQQDAjCBmDELMAkGA1UEBhMCVVMxETAP
BgNVBAgTCE5ldyBZb3JrMREwDwYDVQQHEwhOZXcgWW9yazEYMBYGA1UEChMPRXhh
bXBsZSBDb21wYW55MRkwFwYDVQQLExBUZXN0aW5nIERpdmlzaW9uMRgwFgYDVQQD
Ew93d3cuZXhhbXBsZS5jb20xFDASBgorBgEEAYKifAIBEwRGRkYxMCAXDTI0MDIy
NjExNTQzMVoYDzMwMjMwNjI5MTE1NDMxWjCBgjELMAkGA1UEBhMCVVMxETAPBgNV
BAgTCE5ldyBZb3JrMREwDwYDVQQHEwhOZXcgWW9yazEYMBYGA1UEChMPRXhhbXBs
ZSBDb21wYW55MRkwFwYDVQQLExBUZXN0aW5nIERpdmlzaW9uMRgwFgYDVQQDEw93
d3cuZXhhbXBsZS5jb20wWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAT5fnAAV4+O
bLTyeiyq7LU4WUYHlK+YNnOovoS0ejToaiQ1acjBBTtVpoPf6ltAi+Nh8BRRpJ2+
dPenX51zWyw2o2YwZDAOBgNVHQ8BAf8EBAMCAYYwEgYDVR0TAQH/BAgwBgEB/wIB
ADAdBgNVHQ4EFgQUoe4uCAtmVPiEgIqPlNZp/HVanQgwHwYDVR0jBBgwFoAUa4x3
Hq3LqDwznC8QJ19CAx0K9I4wCgYIKoZIzj0EAwIDSAAwRQIhALlwvvjlvIBZtGsw
sMxUebWm/JTGxBcx5fr2TWgmzpGrAiAgS/w8FjDVq/KvbcjXex3HeDAKN9mYB/ug
aZ1gNyAECQ==
-----END CERTIFICATE-----
`

	LeafCertWithVid = `
-----BEGIN CERTIFICATE-----
MIICrzCCAlSgAwIBAgIUHl25sAyYOxtPeqKkHeRXYwAKOiYwCgYIKoZIzj0EAwIw
ga4xCzAJBgNVBAYTAlVTMREwDwYDVQQIEwhOZXcgWW9yazERMA8GA1UEBxMITmV3
IFlvcmsxGDAWBgNVBAoTD0V4YW1wbGUgQ29tcGFueTEZMBcGA1UECxMQVGVzdGlu
ZyBEaXZpc2lvbjEYMBYGA1UEAxMPd3d3LmV4YW1wbGUuY29tMRQwEgYKKwYBBAGC
onwCARMERkZGMTEUMBIGCisGAQQBgqJ8AgITBEZGRjEwIBcNMjQwMjI2MTE1NDMx
WhgPMzAyMzA2MjkxMTU0MzFaMIGaMQswCQYDVQQGEwJVUzERMA8GA1UECBMITmV3
IFlvcmsxETAPBgNVBAcTCE5ldyBZb3JrMRowGAYDVQQKExFDUkwtbGVhZiB3aXRo
IFZJRDEZMBcGA1UECxMQVGVzdGluZyBEaXZpc2lvbjEYMBYGA1UEAxMPd3d3LmV4
YW1wbGUuY29tMRQwEgYKKwYBBAGConwCARMERkZGMTBZMBMGByqGSM49AgEGCCqG
SM49AwEHA0IABPEBoluRBbOuTZEzRKgNMCA63unDmG3YAA3o+7wYrvC+6pmelaGX
cnAo8PeFbMtRQ2aij2DiJjMgnMMAdvl2VlGjYDBeMA4GA1UdDwEB/wQEAwIBgjAM
BgNVHRMBAf8EAjAAMB0GA1UdDgQWBBRr9RQ9uxoRjmkgM4K5fsrGQ9GeYjAfBgNV
HSMEGDAWgBSwez/xRQGRj8H67suaAQbHR5td7DAKBggqhkjOPQQDAgNJADBGAiEA
1Jtl+jN+Q3dmCREsNI77OOT6j8Zg+YDSQMWKqHATuU4CIQCUoCPkaqadr0Brm+Ue
QiHQtlXK8M8VwJ50OEM+FqlXcg==
-----END CERTIFICATE-----
`

	LeafCertWithVidPid = `
-----BEGIN CERTIFICATE-----
MIICyzCCAnKgAwIBAgIUWX6tKvaztLmIhjGTVvMe8hojy14wCgYIKoZIzj0EAwIw
ga4xCzAJBgNVBAYTAlVTMREwDwYDVQQIEwhOZXcgWW9yazERMA8GA1UEBxMITmV3
IFlvcmsxGDAWBgNVBAoTD0V4YW1wbGUgQ29tcGFueTEZMBcGA1UECxMQVGVzdGlu
ZyBEaXZpc2lvbjEYMBYGA1UEAxMPd3d3LmV4YW1wbGUuY29tMRQwEgYKKwYBBAGC
onwCARMERkZGMTEUMBIGCisGAQQBgqJ8AgITBEZGRjEwIBcNMjQwMjI2MTE1NDMx
WhgPMzAyMzA2MjkxMTU0MzFaMIG4MQswCQYDVQQGEwJVUzERMA8GA1UECBMITmV3
IFlvcmsxETAPBgNVBAcTCE5ldyBZb3JrMSIwIAYDVQQKExlDUkwtbGVhZiB3aXRo
IFZJRCBhbmQgUElEMRkwFwYDVQQLExBUZXN0aW5nIERpdmlzaW9uMRgwFgYDVQQD
Ew93d3cuZXhhbXBsZS5jb20xFDASBgorBgEEAYKifAIBEwRGRkYxMRQwEgYKKwYB
BAGConwCAhMERkZGMTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABJhwAIcHMCUM
7gHfLGhXEkgCiyIzkOUcq9bOg3dmOyMTZVTlLNl6squxgilkZf25jE2ELHXbVhaL
QO09cmnMib+jYDBeMA4GA1UdDwEB/wQEAwIBgjAMBgNVHRMBAf8EAjAAMB0GA1Ud
DgQWBBRACqNJkv74DJN+WRLuPdwWEhX3pzAfBgNVHSMEGDAWgBSwez/xRQGRj8H6
7suaAQbHR5td7DAKBggqhkjOPQQDAgNHADBEAiBqYQYHn0whzLr8CxTEaFck8di4
iYtDkq5gKXncsq+TTgIgULBJ82BNoKrPjymJcTz8m3G/x8r6NY91jKiZqaXOKas=
-----END CERTIFICATE-----
`

	LeafCertWithoutVidPid = `
-----BEGIN CERTIFICATE-----
MIICozCCAkmgAwIBAgIUC6BOWakTVKMR38KH6xSz3CJ7rpIwCgYIKoZIzj0EAwIw
ga4xCzAJBgNVBAYTAlVTMREwDwYDVQQIEwhOZXcgWW9yazERMA8GA1UEBxMITmV3
IFlvcmsxGDAWBgNVBAoTD0V4YW1wbGUgQ29tcGFueTEZMBcGA1UECxMQVGVzdGlu
ZyBEaXZpc2lvbjEYMBYGA1UEAxMPd3d3LmV4YW1wbGUuY29tMRQwEgYKKwYBBAGC
onwCARMERkZGMTEUMBIGCisGAQQBgqJ8AgITBEZGRjEwIBcNMjQwMjI2MTE1NDMx
WhgPMzAyMzA2MjkxMTU0MzFaMIGPMQswCQYDVQQGEwJVUzERMA8GA1UECBMITmV3
IFlvcmsxETAPBgNVBAcTCE5ldyBZb3JrMSUwIwYDVQQKExxDUkwtbGVhZiB3aXRo
b3V0IFZJRCBhbmQgUElEMRkwFwYDVQQLExBUZXN0aW5nIERpdmlzaW9uMRgwFgYD
VQQDEw93d3cuZXhhbXBsZS5jb20wWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATh
2x7cYw5BHTYNBz67WzF3c+R1EvKedlvk80K1BNSNgfO8HFVWGR4xvC3mJE5cBVB8
S1MzXrLH1pY/3ds4b75Lo2AwXjAOBgNVHQ8BAf8EBAMCAYIwDAYDVR0TAQH/BAIw
ADAdBgNVHQ4EFgQUJgN7taAIhLB8QIzeB3Iux398WFQwHwYDVR0jBBgwFoAUsHs/
8UUBkY/B+u7LmgEGx0ebXewwCgYIKoZIzj0EAwIDSAAwRQIgFOenqyPaLe++AicB
F3GvTS0sJnz80nDXDtoftuHpTe0CIQDS2N1WBAENrj2poT/+rHWtEGoyuEaL7UXy
eZwFu4fZgg==
-----END CERTIFICATE-----
`

	CertWithSerialNumber21octets = `-----BEGIN CERTIFICATE-----
MIIBXzCCAQSgAwIBAgIVASNFZ4mrze8BI0VniavN7wEjRWeJMAoGCCqGSM49BAMC
MBwxGjAYBgNVBAMMEU92ZXJzaXplZC1FQy1UZXN0MCAXDTI2MDIxNDA5NTAzM1oY
DzIxMjYwMTIxMDk1MDMzWjAcMRowGAYDVQQDDBFPdmVyc2l6ZWQtRUMtVGVzdDBZ
MBMGByqGSM49AgEGCCqGSM49AwEHA0IABDLr1YtJFdyyiOcN5hrqzDK6Y2q5vBCx
pw+ngG7+0GV3X4oEhKl8FoX1tFWZ8VhQffZuvv+w+OgCYJ2knfbCbLujITAfMB0G
A1UdDgQWBBQdANeX8HpQXhv/J+xX/HY+eW1W+TAKBggqhkjOPQQDAgNJADBGAiEA
iOBJB5ipLIANiD9Vvp8j5kE/wTlu2Pk0vhpxCwVus9wCIQCC6b+TeViQceE2l3Da
OoWTfXtini5G54diYLC/2uHrpw==
-----END CERTIFICATE-----`

	CertWithInvalidSerialNumber = `-----BEGIN CERTIFICATE-----
MIIBSTCB8KADAgECAgEAMAoGCCqGSM49BAMCMBwxGjAYBgNVBAMMEU92ZXJzaXpl
ZC1FQy1UZXN0MCAXDTI2MDIxNDA5NTYyM1oYDzIxMjYwMTIxMDk1NjIzWjAcMRow
GAYDVQQDDBFPdmVyc2l6ZWQtRUMtVGVzdDBZMBMGByqGSM49AgEGCCqGSM49AwEH
A0IABDLr1YtJFdyyiOcN5hrqzDK6Y2q5vBCxpw+ngG7+0GV3X4oEhKl8FoX1tFWZ
8VhQffZuvv+w+OgCYJ2knfbCbLujITAfMB0GA1UdDgQWBBQdANeX8HpQXhv/J+xX
/HY+eW1W+TAKBggqhkjOPQQDAgNIADBFAiBNw0mLQcOI9k7ZEtg+dgrniJqc0yYw
DJAQ0+r/O9wutAIhANQxSyFpLrpDKbmAQvnuzq8LlVoOxYhA2wzcURyCH747
-----END CERTIFICATE-----`

	CertWithSizeGreater2KB = `-----BEGIN CERTIFICATE-----
MIIHFTCCBrugAwIBAgIUP4MM/GLQPb8mLfmZOvlUpr/n2uYwCgYIKoZIzj0EAwIw
HDEaMBgGA1UEAwwRT3ZlcnNpemVkLUVDLVRlc3QwIBcNMjYwMjE4MDk0NDM5WhgP
MjEyNjAxMjUwOTQ0MzlaMBwxGjAYBgNVBAMMEU92ZXJzaXplZC1FQy1UZXN0MFkw
EwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEMuvVi0kV3LKI5w3mGurMMrpjarm8ELGn
D6eAbv7QZXdfigSEqXwWhfW0VZnxWFB99m6+/7D46AJgnaSd9sJsu6OCBdcwggXT
MBYGA1UdEQQPMA2CC2V4YW1wbGUuY29tMIIFmAYJYIZIAYb4QgENBIIFiRaCBYVM
b3JlbSBpcHN1bSBkb2xvciBzaXQgYW1ldCwgY29uc2VjdGV0dXIgYWRpcGlzY2lu
ZyBlbGl0LiBEdWlzIGJsYW5kaXQgaXBzdW0gYXVndWUsIGEgcnV0cnVtIGxpYmVy
byBlZ2VzdGFzIGEuIFZlc3RpYnVsdW0gcXVpcyBsaWd1bGEgZnJpbmdpbGxhLCBp
YWN1bGlzIHJpc3VzIGEsIHZhcml1cyBtaS4gUGVsbGVudGVzcXVlIHZvbHV0cGF0
IGlwc3VtIHNpdCBhbWV0IGR1aSBzYWdpdHRpcyBzZW1wZXIuIEV0aWFtIGVyb3Mg
bmVxdWUsIGZpbmlidXMgcXVpcyBtb2xsaXMgc2l0IGFtZXQsIHN1c2NpcGl0IG5v
biB2ZWxpdC4gTnVuYyBlZ2V0IGVyYXQgaW4gZW5pbSBlZmZpY2l0dXIgbG9ib3J0
aXMgc2VtcGVyIG5vbiBlcm9zLiBQZWxsZW50ZXNxdWUgaGFiaXRhbnQgbW9yYmkg
dHJpc3RpcXVlIHNlbmVjdHVzIGV0IG5ldHVzIGV0IG1hbGVzdWFkYSBmYW1lcyBh
YyB0dXJwaXMgZWdlc3Rhcy4gUGVsbGVudGVzcXVlIGVnZXQgaW1wZXJkaWV0IG1l
dHVzLiBDdXJhYml0dXIgc3VzY2lwaXQgbnVsbGEgaW4gYWxpcXVhbSB1bGxhbWNv
cnBlci4gUXVpc3F1ZSBvcmNpIG51bGxhLCBhY2N1bXNhbiB2ZWwgdmVoaWN1bGEg
bmVjLCBsdWN0dXMgdXQgbG9yZW0uRXRpYW0gdGVtcG9yIHF1YW0gZXUgZW5pbSBz
ZW1wZXIgb3JuYXJlLiBOdW5jIHZlbmVuYXRpcyBuaXNpIHF1aXMgbGVvIHJ1dHJ1
bSwgZXQgbWF4aW11cyBlbGl0IGJpYmVuZHVtLiBQZWxsZW50ZXNxdWUgaGFiaXRh
bnQgbW9yYmkgdHJpc3RpcXVlIHNlbmVjdHVzIGV0IG5ldHVzIGV0IG1hbGVzdWFk
YSBmYW1lcyBhYyB0dXJwaXMgZWdlc3Rhcy4gQWVuZWFuIHByZXRpdW0sIGxpYmVy
byBub24gcHJldGl1bSBwb3N1ZXJlLCBuaWJoIGxlbyBjdXJzdXMgZHVpLCBldSBp
YWN1bGlzIG1hZ25hIGxlY3R1cyBpZCB0ZWxsdXMuIFV0IGxpYmVybyBhbnRlLCBw
ZWxsZW50ZXNxdWUgbmVjIGlwc3VtIGluLCB2ZXN0aWJ1bHVtIGNvbmd1ZSBzYXBp
ZW4uIENyYXMgdmVsaXQgZXJvcywgdWxsYW1jb3JwZXIgdmVsIGZhdWNpYnVzIGV0
LCBibGFuZGl0IHV0IGVsaXQuIE1hdXJpcyBmaW5pYnVzIG9yY2kgdml0YWUgdmVs
aXQgdWx0cmljaWVzLCB2aXRhZSB2ZWhpY3VsYSB0dXJwaXMgY29uZ3VlLiBRdWlz
cXVlIHJob25jdXMgaW50ZXJkdW0gdWxsYW1jb3JwZXIuIERvbmVjIHZlaGljdWxh
IGxlY3R1cyBuZWMgc2FwaWVuIG9ybmFyZSBhbGlxdWFtLiBEb25lYyBjb25ndWUg
c2VtIGV0IG9yY2kgdml2ZXJyYSBldWlzbW9kIGNvbnNlY3RldHVyIG5vbiBhdWd1
ZS4gUHJvaW4gc2l0IGFtZXQgZXVpc21vZCBsZW8sIG5vbiBibGFuZGl0IG51bGxh
LiBQcm9pbiBwdWx2aW5hciBqdXN0byB2ZWwgZG9sb3IgYmxhbmRpdCwgcXVpcyBv
cm5hcmUgYXVndWUgcmhvbmN1cy4gTmFtIGlhY3VsaXMgcmlzdXMgZXUgdHVycGlz
IHN1c2NpcGl0IGNvbnNlcXVhdC4wHQYDVR0OBBYEFB0A15fwelBeG/8n7Ff8dj55
bVb5MAoGCCqGSM49BAMCA0gAMEUCIEK+KrlapPMVX8009oalJb21le7l1cjI2Z1P
HsesWPZNAiEAg1rO3UB1UCteAplSzGRLIAPSrdL+IBANEeSqcCoQI5U=
-----END CERTIFICATE-----`

	RootIssuer                     = "MDQxCzAJBgNVBAYTAkFVMRMwEQYDVQQIEwpzb21lLXN0YXRlMRAwDgYDVQQKEwdyb290LWNh"
	RootSubject                    = "MDQxCzAJBgNVBAYTAkFVMRMwEQYDVQQIEwpzb21lLXN0YXRlMRAwDgYDVQQKEwdyb290LWNh"
	RootSubjectAsText              = "O=root-ca,ST=some-state,C=AU"
	RootSubjectKeyID               = "DF:4E:AF:B0:8C:9C:37:78:1A:E7:53:12:CA:E4:78:6B:48:1E:AF:B0"
	RootSubjectKeyIDWithoutColumns = "DF4EAFB08C9C37781AE75312CAE4786B481EAFB0"
	RootSerialNumber               = "81311506302208030248766861785118937702312370677"

	RootCertWithSameSubjectAndSKIDSubject               = "MIGCMQswCQYDVQQGEwJVUzERMA8GA1UECBMITmV3IFlvcmsxETAPBgNVBAcTCE5ldyBZb3JrMRgwFgYDVQQKEw9FeGFtcGxlIENvbXBhbnkxGTAXBgNVBAsTEFRlc3RpbmcgRGl2aXNpb24xGDAWBgNVBAMTD3d3dy5leGFtcGxlLmNvbQ=="
	RootCertWithSameSubjectAndSKIDSubjectAsText         = "C=US,ST=New York,L=New York,O=Example Company,OU=Testing Division,CN=www.example.com"
	RootCertWithSameSubjectAndSKIDSubjectKeyID          = "C1:48:66:ED:6F:23:D8:28:1A:D9:37:7C:58:AC:3F:DA:04:C1:41:E8"
	RootCertWithSameSubjectAndSKID1SerialNumber         = "1"
	RootCertWithSameSubjectAndSKID2SerialNumber         = "2"
	RootCertWithSameSubjectAndSKID1Issuer               = "MIGCMQswCQYDVQQGEwJVUzERMA8GA1UECBMITmV3IFlvcmsxETAPBgNVBAcTCE5ldyBZb3JrMRgwFgYDVQQKEw9FeGFtcGxlIENvbXBhbnkxGTAXBgNVBAsTEFRlc3RpbmcgRGl2aXNpb24xGDAWBgNVBAMTD3d3dy5leGFtcGxlLmNvbQ=="
	RootCertWithSameSubjectAndSKID2Issuer               = "MIGCMQswCQYDVQQGEwJVUzERMA8GA1UECBMITmV3IFlvcmsxETAPBgNVBAcTCE5ldyBZb3JrMRgwFgYDVQQKEw9FeGFtcGxlIENvbXBhbnkxGTAXBgNVBAsTEFRlc3RpbmcgRGl2aXNpb24xGDAWBgNVBAMTD3d3dy5leGFtcGxlLmNvbQ=="
	IntermediateCertWithSameSubjectAndSKIDSubject       = "MEUxCzAJBgNVBAYTAkFVMRMwEQYDVQQIEwpTb21lLVN0YXRlMSEwHwYDVQQKExhJbnRlcm5ldCBXaWRnaXRzIFB0eSBMdGQ="
	IntermediateCertWithSameSubjectAndSKIDSubjectAsText = "C=AU,ST=Some-State,O=Internet Widgits Pty Ltd"

	IntermediateCertWithSameSubjectIssuer               = RootCertWithSameSubjectAndSKIDSubject
	IntermediateCertWithSameSubjectAuthorityKeyID       = RootCertWithSameSubjectAndSKIDSubjectKeyID
	IntermediateCertWithSameSubjectAndSKIDSubjectKeyID  = "A1:E0:92:89:FA:18:82:12:14:9D:B8:AE:19:43:BE:44:31:6B:F1:F5"
	IntermediateCertWithSameSubjectAndSKIDIssuer        = RootCertWithSameSubjectAndSKIDSubject
	IntermediateCertWithSameSubjectAndSKID1SerialNumber = "3"
	IntermediateCertWithSameSubjectAndSKID2SerialNumber = "4"
	LeafCertWithSameSubjectAndSKIDSubject               = "MEUxCzAJBgNVBAYTAkFVMRMwEQYDVQQIEwpTb21lLVN0YXRlMSEwHwYDVQQKExhJbnRlcm5ldCBXaWRnaXRzIFB0eSBMdGQ="
	LeafCertWithSameSubjectAndSKIDSubjectAsText         = "C=AU,ST=Some-State,O=Internet Widgits Pty Ltd"
	LeafCertWithSameSubjectAndSKIDSubjectKeyID          = "90:81:84:C7:EC:B8:81:14:66:61:2F:82:BB:E9:51:67:F2:4D:99:A3"
	LeafCertWithSameSubjectAndSKIDSerialNumber          = "5"
	LeafCertWithSameSubjectIssuer                       = IntermediateCertWithSameSubjectAndSKIDSubject
	LeafCertWithSameSubjectAuthorityKeyID               = IntermediateCertWithSameSubjectAndSKIDSubjectKeyID

	IntermediateIssuer                     = "MDQxCzAJBgNVBAYTAkFVMRMwEQYDVQQIEwpzb21lLXN0YXRlMRAwDgYDVQQKEwdyb290LWNh"
	IntermediateAuthorityKeyID             = "DF:4E:AF:B0:8C:9C:37:78:1A:E7:53:12:CA:E4:78:6B:48:1E:AF:B0"
	IntermediateSubject                    = "MDwxCzAJBgNVBAYTAkFVMRMwEQYDVQQIEwpzb21lLXN0YXRlMRgwFgYDVQQKEw9pbnRlcm1lZGlhdGUtY2E="
	IntermediateSubjectAsText              = "O=intermediate-ca,ST=some-state,C=AU"
	IntermediateSubjectKeyID               = "1B:73:2A:91:34:46:8A:90:2A:87:19:91:E4:BD:8F:69:3A:F9:04:77"
	IntermediateSubjectKeyIDWithoutColumns = "1B732A9134468A902A871991E4BD8F693AF90477"
	IntermediateSerialNumber               = "486736128900935106101503663840421220667833341899"

	LeafIssuer         = IntermediateSubject
	LeafAuthorityKeyID = IntermediateSubjectKeyID
	LeafSubject        = "MDExCzAJBgNVBAYTAkFVMRMwEQYDVQQIEwpzb21lLXN0YXRlMQ0wCwYDVQQKEwRsZWFm"
	LeafSubjectAsText  = "O=leaf,ST=some-state,C=AU"
	LeafSubjectKeyID   = "2A:31:8D:39:6E:50:DA:96:DF:95:C5:98:83:68:F0:58:B2:15:B3:3A"
	LeafSerialNumber   = "409691117370409054634487600348183880852961428328"

	GoogleIssuer         = "MEsxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZHb29nbGUxFTATBgNVBAMMDE1hdHRlciBQQUEgMTEUMBIGCisGAQQBgqJ8AgEMBDYwMDY="
	GoogleAuthorityKeyID = ""
	GoogleSubject        = "MEsxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZHb29nbGUxFTATBgNVBAMMDE1hdHRlciBQQUEgMTEUMBIGCisGAQQBgqJ8AgEMBDYwMDY="
	GoogleSubjectAsText  = "CN=Matter PAA 1,O=Google,C=US,vid=0x6006"
	GoogleSubjectKeyID   = "B0:00:56:81:B8:88:62:89:62:80:E1:21:18:A1:A8:BE:09:DE:93:21"
	GoogleSerialNumber   = "1"
	GoogleVid            = 65521

	TestIssuer         = "MDAxGDAWBgNVBAMMD01hdHRlciBUZXN0IFBBQTEUMBIGCisGAQQBgqJ8AgEMBDEyNUQ="
	TestAuthorityKeyID = "E2:90:8D:36:9C:3C:A3:C1:13:BB:09:E2:4D:C1:CC:C5:A6:66:91:D4"
	TestSubject        = "MDAxGDAWBgNVBAMMD01hdHRlciBUZXN0IFBBQTEUMBIGCisGAQQBgqJ8AgEMBDEyNUQ="
	TestSubjectAsText  = "CN=Matter Test PAA,vid=0x125D"
	TestSubjectKeyID   = "E2:90:8D:36:9C:3C:A3:C1:13:BB:09:E2:4D:C1:CC:C5:A6:66:91:D4"
	TestSerialNumber   = "1647312298631"

	PAACertWithSameSubjectID1Subject       = "MFoxCzAJBgNVBAYTAlVaMQwwCgYDVQQIEwNUU0gxETAPBgNVBAcTCFRBU0hLRU5UMQwwCgYDVQQKEwNEU1IxCzAJBgNVBAsTAkRDMQ8wDQYDVQQDEwZNQVRURVI="
	PAACertWithSameSubjectID1SubjectAsText = "C=UZ,ST=TSH,L=TASHKENT,O=DSR,OU=DC,CN=MATTER"
	PAACertWithSameSubjectID2Subject       = "MGAxCzAJBgNVBAYTAlVaMQwwCgYDVQQIEwNUU0gxETAPBgNVBAcTCFRBU0hLRU5UMQwwCgYDVQQKEwNEU1IxEDAOBgNVBAsTB01BVFRFUjIxEDAOBgNVBAMTB01BVFRFUjI="
	PAACertWithSameSubjectIDSubjectKeyID   = "81:B7:F8:43:15:6C:0F:67:30:80:D9:92:B9:30:6C:55:59:D3:D6:65"
	PAACertWithSameSubjectIssuer           = "MFoxCzAJBgNVBAYTAlVaMQwwCgYDVQQIEwNUU0gxETAPBgNVBAcTCFRBU0hLRU5UMQwwCgYDVQQKEwNEU1IxCzAJBgNVBAsTAkRDMQ8wDQYDVQQDEwZNQVRURVI="
	PAACertWithSameSubjectSerialNumber     = "1001"
	PAACertWithSameSubject2Issuer          = "MGAxCzAJBgNVBAYTAlVaMQwwCgYDVQQIEwNUU0gxETAPBgNVBAcTCFRBU0hLRU5UMQwwCgYDVQQKEwNEU1IxEDAOBgNVBAsTB01BVFRFUjIxEDAOBgNVBAMTB01BVFRFUjI="
	PAACertWithSameSubject2SerialNumber    = "1002"

	TestVID1String            = "0xA13"
	TestPID1String            = "0xA11"
	TestVID2String            = "0xA14"
	TestPID2String            = "0xA15"
	TestVID3String            = "0xA16"
	TestPID3String            = "0xA17"
	SubjectKeyIDWithoutColons = "5A880E6C3653D07FB08971A3F473790930E62BDB"
	DataDigest                = "9a5d2c1f4b3e6f8d7b1a0c9e2f5d8b7"

	TestCertPemVid = 4701

	RootCertWithVidSubject                    = "MIGYMQswCQYDVQQGEwJVUzERMA8GA1UECBMITmV3IFlvcmsxETAPBgNVBAcTCE5ldyBZb3JrMRgwFgYDVQQKEw9FeGFtcGxlIENvbXBhbnkxGTAXBgNVBAsTEFRlc3RpbmcgRGl2aXNpb24xGDAWBgNVBAMTD3d3dy5leGFtcGxlLmNvbTEUMBIGCisGAQQBgqJ8AgETBEZGRjE="
	RootCertWithVidSubjectSubjectAsText       = "C=US,ST=New York,L=New York,O=Example Company,OU=Testing Division,CN=www.example.com,vid=0xFFF1"
	RootCertWithVidSubjectKeyID               = "6B:8C:77:1E:AD:CB:A8:3C:33:9C:2F:10:27:5F:42:03:1D:0A:F4:8E"
	RootCertWithVidSubjectKeyIDWithoutColumns = "6B8C771EADCBA83C339C2F10275F42031D0AF48E"
	RootCertWithVidVid                        = 65521
	RootCertWithVidSerialNumber               = "1"

	IntermediateCertWithVidIssuer                      = RootCertWithVidSubject
	IntermediateCertWithVid1Subject                    = "MIGuMQswCQYDVQQGEwJVUzERMA8GA1UECBMITmV3IFlvcmsxETAPBgNVBAcTCE5ldyBZb3JrMRgwFgYDVQQKEw9FeGFtcGxlIENvbXBhbnkxGTAXBgNVBAsTEFRlc3RpbmcgRGl2aXNpb24xGDAWBgNVBAMTD3d3dy5leGFtcGxlLmNvbTEUMBIGCisGAQQBgqJ8AgETBEZGRjExFDASBgorBgEEAYKifAICEwRGRkYx"
	IntermediateCertWithVid1SubjectKeyID               = "B0:7B:3F:F1:45:01:91:8F:C1:FA:EE:CB:9A:01:06:C7:47:9B:5D:EC"
	IntermediateCertWithVid1SubjectKeyIDWithoutColumns = "B07B3FF14501918FC1FAEECB9A0106C7479B5DEC"
	IntermediateCertWithVid1SerialNumber               = "3"
	IntermediateCertWithVid1Vid                        = 65521

	IntermediateCertWithVid2Subject      = "MIGuMQswCQYDVQQGEwJVUzERMA8GA1UECBMITmV3IFlvcmsxETAPBgNVBAcTCE5ldyBZb3JrMRgwFgYDVQQKEw9FeGFtcGxlIENvbXBhbnkxGTAXBgNVBAsTEFRlc3RpbmcgRGl2aXNpb24xGDAWBgNVBAMTD3d3dy5leGFtcGxlLmNvbTEUMBIGCisGAQQBgqJ8AgETBEZGRjIxFDASBgorBgEEAYKifAICEwRGRkYy"
	IntermediateCertWithVid2SubjectKeyID = "FC:5E:74:1D:BB:95:7C:95:0F:36:D6:D6:F7:0E:95:05:18:5B:7E:03"
	IntermediateCertWithVid2SerialNumber = "6"
	IntermediateCertWithVid2Vid          = 65522

	IntermediateCertWithoutVidPidSubject      = "MIGCMQswCQYDVQQGEwJVUzERMA8GA1UECBMITmV3IFlvcmsxETAPBgNVBAcTCE5ldyBZb3JrMRgwFgYDVQQKEw9FeGFtcGxlIENvbXBhbnkxGTAXBgNVBAsTEFRlc3RpbmcgRGl2aXNpb24xGDAWBgNVBAMTD3d3dy5leGFtcGxlLmNvbQ=="
	IntermediateCertWithoutVidPidSubjectKeyID = "A1:EE:2E:08:0B:66:54:F8:84:80:8A:8F:94:D6:69:FC:75:5A:9D:08"
	IntermediateCertWithoutVidPidSerialNumber = "7"

	LeafCertWithVidIssuer         = IntermediateCertWithVid1Subject
	LeafCertWithVidSubject        = "MIGaMQswCQYDVQQGEwJVUzERMA8GA1UECBMITmV3IFlvcmsxETAPBgNVBAcTCE5ldyBZb3JrMRowGAYDVQQKExFDUkwtbGVhZiB3aXRoIFZJRDEZMBcGA1UECxMQVGVzdGluZyBEaXZpc2lvbjEYMBYGA1UEAxMPd3d3LmV4YW1wbGUuY29tMRQwEgYKKwYBBAGConwCARMERkZGMQ=="
	LeafCertWithVidSubjectAsText  = "CN=www.example.com,OU=Testing Division,O=CRL-leaf with VID,L=New York,ST=New York,C=US,vid=0xFFF1"
	LeafCertWithVidSubjectKeyID   = "6B:F5:14:3D:BB:1A:11:8E:69:20:33:82:B9:7E:CA:C6:43:D1:9E:62"
	LeafCertWithVidAuthorityKeyID = IntermediateCertWithVid1SubjectKeyID
	LeafCertWithVidSerialNumber   = "173359868107513651307797193493378479362570598950"
	LeafCertWithVidVid            = 65521

	LeafCertWithVidPidSubject        = "MIG4MQswCQYDVQQGEwJVUzERMA8GA1UECBMITmV3IFlvcmsxETAPBgNVBAcTCE5ldyBZb3JrMSIwIAYDVQQKExlDUkwtbGVhZiB3aXRoIFZJRCBhbmQgUElEMRkwFwYDVQQLExBUZXN0aW5nIERpdmlzaW9uMRgwFgYDVQQDEw93d3cuZXhhbXBsZS5jb20xFDASBgorBgEEAYKifAIBEwRGRkYxMRQwEgYKKwYBBAGConwCAhMERkZGMQ=="
	LeafCertWithVidPidSubjectAsText  = "CN=www.example.com,OU=Testing Division,O=CRL-leaf with VID and PID,L=New York,ST=New York,C=US,pid=0xFFF1,vid=0xFFF1"
	LeafCertWithVidPidSubjectKeyID   = "40:0A:A3:49:92:FE:F8:0C:93:7E:59:12:EE:3D:DC:16:12:15:F7:A7"
	LeafCertWithVidPidAuthorityKeyID = IntermediateCertWithVid1SubjectKeyID
	LeafCertWithVidPidSerialNumber   = "510925157543585355008626589851569388005255007070"
	LeafCertWithVidPidVid            = 65521
	LeafCertWithVidPidPid            = 65521

	LeafCertWithoutVidPidSubject        = "MIGPMQswCQYDVQQGEwJVUzERMA8GA1UECBMITmV3IFlvcmsxETAPBgNVBAcTCE5ldyBZb3JrMSUwIwYDVQQKExxDUkwtbGVhZiB3aXRob3V0IFZJRCBhbmQgUElEMRkwFwYDVQQLExBUZXN0aW5nIERpdmlzaW9uMRgwFgYDVQQDEw93d3cuZXhhbXBsZS5jb20="
	LeafCertWithoutVidPidSubjectAsText  = "CN=www.example.com,OU=Testing Division,O=CRL-leaf without VID and PID,L=New York,ST=New York,C=US"
	LeafCertWithoutVidPidSubjectKeyID   = "26:03:7B:B5:A0:08:84:B0:7C:40:8C:DE:07:72:2E:C7:7F:7C:58:54"
	LeafCertWithoutVidPidAuthorityKeyID = IntermediateCertWithoutVidPidSubjectKeyID
	LeafCertWithoutVidPidSerialNumber   = "66373842979000369302063877566075311743961902738"
)
