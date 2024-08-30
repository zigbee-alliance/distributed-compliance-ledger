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
	OtaChecksum                                       = "2fd4e1c67a2d28fced849ee1bb76e7391b93eb12" //nolint:lll
	OtaChecksumType                            int32  = 1
	OtaBlob                                           = "OTABlob Text"
	CommissioningCustomFlow                    int32  = 1
	CommissioningCustomFlowURL                        = "https://sampleflowurl.dclmodel"
	CommissioningModeInitialStepsHint          uint32 = 2
	CommissioningModeInitialStepsInstruction          = "commissioningModeInitialStepsInstruction details"
	CommissioningModeSecondaryStepsHint        uint32 = 3
	CommissioningModeSecondaryStepsInstruction        = "commissioningModeSecondaryStepsInstruction steps"
	ReleaseNotesURL                                   = "https://url.releasenotes.dclmodel"
	UserManualURL                                     = "https://url.usermanual.dclmodel"
	SupportURL                                        = "https://url.supporturl.dclmodel"
	ProductURL                                        = "https://url.producturl.dclmodel"
	EnhancedSetupFlowTCURL                            = "https://url.enhansedsetupflowurl.dclmodel"
	EnhancedSetupFlowTCRevision                       = 1
	EnhancedSetupFlowTCDigest                         = "MmNmMjRkYmE1ZmIwYTMwZTI2ZTgzYjJhYzViOWUyOWUxYjE2MWU1YzFmYTc0MjVlNzMwNDMzNjI5MzhiOTgyNA=="
	EnhancedSetupFlowTCFileSize                       = 1
	MaintenanceURL                                    = "https://url.maintenanceurl.dclmodel"
	LsfURL                                            = "https://url.lsfurl.dclmodel"
	DataURL                                           = "https://url.data.dclmodel"
	DataURL2                                          = "https://url.data.dclmodel2"
	URLWithoutProtocol                                = "url.dclmodel"
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
	CDCertificateID                    = "15DEXF"
	FamilyID                           = "Some Family ID"
	SupportedClusters                  = "Some Supported Clusters"
	CompliantPlatformUsed              = "Some Compliance Platform Used"
	CompliantPlatformVersion           = "Some Compliance Platform Version"
	OSVersion                          = "Some OS Version"
	CertificationRoute                 = "Some Certification Route"
	ProgramType                        = "Some Program Type"
	ProgramTypeVersion                 = "Some Program Type Version"
	Transport                          = "Some Transport"
	SoftwareVersionCertificationStatus = uint32(3)
	ParentChild1                       = "parent"
	ParentChild2                       = "child"
	CertificationIDOfSoftwareComponent = "some certification ID of software component"
	FirstJanuary                       = "2020-01-01T00:00:01Z"

	// Testing Result.
	TestResult = "http://test.result.com"
	TestDate   = "2020-02-02T02:00:00Z"

	// Upgrade.
	UpgradePlanName         = "TestUpgrade"
	UpgradePlanHeight int64 = 1337
	UpgradePlanInfo         = "Some upgrade info"

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
	SchemaVersion uint32
)

/*
	Certificates are taken from dsr-corporation.com
*/

const (
	StubCertPem = `pem certificate`

	RootCertPem = `
-----BEGIN CERTIFICATE-----
MIIBvjCCAWWgAwIBAgIUTXoMP/NTKMkiXcBqcmVzMsNSRBMwCgYIKoZIzj0EAwIw
NDELMAkGA1UEBhMCQVUxEzARBgNVBAgMCnNvbWUtc3RhdGUxEDAOBgNVBAoMB3Jv
b3QtY2EwIBcNMjAwOTExMDk0MDM4WhgPNDc1ODA4MDgwOTQwMzhaMDQxCzAJBgNV
BAYTAkFVMRMwEQYDVQQIDApzb21lLXN0YXRlMRAwDgYDVQQKDAdyb290LWNhMFkw
EwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEN/59wQ7D+o9NXwK0kXkklzP5FR9kwduu
T4LggVdNTHM4DCUN22OaL37xkjFMZFd7avGigQaaZXb9FHkuSVRxHKNTMFEwHQYD
VR0OBBYEFFqIDmw2U9B/sIlxo/RzeQkw5ivbMB8GA1UdIwQYMBaAFFqIDmw2U9B/
sIlxo/RzeQkw5ivbMA8GA1UdEwEB/wQFMAMBAf8wCgYIKoZIzj0EAwIDRwAwRAIg
OA++8ir7a/b8hxgMG9RQiKM/Dvxg3+MJCXF6v+IV404CIDfe7CYKy3sOgcam2bFY
VtvkclFxeP9KADpcPOXXQLzG
-----END CERTIFICATE-----`

	IntermediateCertPem = `
-----BEGIN CERTIFICATE-----
MIIB0TCCAXegAwIBAgIUHcNelel2YdCqtFAAQdsjY8F7Fz8wCgYIKoZIzj0EAwIw
NDELMAkGA1UEBhMCQVUxEzARBgNVBAgMCnNvbWUtc3RhdGUxEDAOBgNVBAoMB3Jv
b3QtY2EwIBcNMjAwOTExMDk0MDM4WhgPNDc1ODA4MDgwOTQwMzhaMDwxCzAJBgNV
BAYTAkFVMRMwEQYDVQQIDApzb21lLXN0YXRlMRgwFgYDVQQKDA9pbnRlcm1lZGlh
dGUtY2EwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQu/epaxQzpTAPoLOm0Su5Y
zhg/5Qi+VPvP/qBT/r4k9oLfmnM1tkX78H/Gb2KIYpZWOIIeTw79siSNMYhHqQ9Y
o10wWzAMBgNVHRMEBTADAQH/MB0GA1UdDgQWBBROO3P0cE3CmA3byFpfAju/hiVW
KzAfBgNVHSMEGDAWgBRaiA5sNlPQf7CJcaP0c3kJMOYr2zALBgNVHQ8EBAMCAqQw
CgYIKoZIzj0EAwIDSAAwRQIhAPy+V4Z/NE1XepIcroa30B+gYIpTIHbYiE1lprwn
NvP9AiATc6OC9FU4WE193XVy9xOn7TZ11BJjNDfBNUt3qrgp5w==
-----END CERTIFICATE-----`

	LeafCertPem = `
-----BEGIN CERTIFICATE-----
MIIBzDCCAXGgAwIBAgIUGRld9+18qI6qmbtyRk68F/DYGY8wCgYIKoZIzj0EAwIw
PDELMAkGA1UEBhMCQVUxEzARBgNVBAgMCnNvbWUtc3RhdGUxGDAWBgNVBAoMD2lu
dGVybWVkaWF0ZS1jYTAgFw0yMDA5MTEwOTQwMzhaGA80NzU4MDgwODA5NDAzOFow
MTELMAkGA1UEBhMCQVUxEzARBgNVBAgMCnNvbWUtc3RhdGUxDTALBgNVBAoMBGxl
YWYwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATPnm0zlXXZKPbOb/2R9tmuBW/A
J9zB2ZUssik8FtC7vxxQn12KgeItR1GZf/L9kSfMTcnV0qH/+dbL+lsIkGgDo1ow
WDAJBgNVHRMEAjAAMB0GA1UdDgQWBBQw9GV1FCCyrz0UcResSZCTPiSgHzAfBgNV
HSMEGDAWgBROO3P0cE3CmA3byFpfAju/hiVWKzALBgNVHQ8EBAMCBaAwCgYIKoZI
zj0EAwIDSQAwRgIhAPq8sXrMDueq9XplZBcbS/3VlTakULzdOlo7PzquUdDnAiEA
0JQR8xr2SnNKb+eeCuxsgZFe7RkxkWGdQwXzF2chq34=
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
	PAACertWithNumericVidSubject              = "MDAxGDAWBgNVBAMMD01hdHRlciBUZXN0IFBBQTEUMBIGCisGAQQBgqJ8AgEMBEZGRjE="
	PAACertWithNumericVidSubjectKeyID         = "6A:FD:22:77:1F:51:1F:EC:BF:16:41:97:67:10:DC:DC:31:A1:71:7E"
	PAACertWithNumericVidVid                  = 65521
	PAACertWithNumericVidDifferentWhitespaces = `
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

	PAICertWithNumericPidVidSubject      = "MEYxGDAWBgNVBAMMD01hdHRlciBUZXN0IFBBSTEUMBIGCisGAQQBgqJ8AgEMBEZGRjExFDASBgorBgEEAYKifAICDAQ4MDAw"
	PAICertWithNumericPidVidSubjectKeyID = "AF:42:B7:09:4D:EB:D5:15:EC:6E:CF:33:B8:11:15:22:5F:32:52:88"
	PAICertWithNumericPidVidVid          = 65521
	PAICertWithNumericPidVidPid          = 32768

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
MIIDYzCCAkugAwIBAgIUCS2EDr4/3u40BiqmXDwCa4DQAbMwDQYJKoZIhvcNAQEL
BQAwWjELMAkGA1UEBhMCVVoxDDAKBgNVBAgMA1RTSDERMA8GA1UEBwwIVEFTSEtF
TlQxDDAKBgNVBAoMA0RTUjELMAkGA1UECwwCREMxDzANBgNVBAMMBk1BVFRFUjAe
Fw0yNDAxMTgwOTAxMThaFw0yNTAxMTcwOTAxMThaMFoxCzAJBgNVBAYTAlVaMQww
CgYDVQQIDANUU0gxETAPBgNVBAcMCFRBU0hLRU5UMQwwCgYDVQQKDANEU1IxCzAJ
BgNVBAsMAkRDMQ8wDQYDVQQDDAZNQVRURVIwggEiMA0GCSqGSIb3DQEBAQUAA4IB
DwAwggEKAoIBAQCdmh5EgxuSYGwxjikI+ABz1e1dj42RbLjSzYrpSijF5LXo83AB
4LODFlegOupAIB6Wnepsu+VmJrSPDyQKY2nvu0asZ6LDSr3TtGoOejxqFWY8RyHE
RAYkWTaeq5W9cNzNY5iGPsQFxJd8V6Dry/rNegPpSsSVJH8R0R1GxBLkfO5lDmPc
PUmqazC7fwj82liSt0G66uhiGhbBVkUXItO1LHlc29dv88W9OkJsNvB9E4vaEEUG
R0IN0ixeJn9MYktYVmCHG9pcVqE6/nYR/ZsyShfYSh+qsTlIlM67WIeOItyvjIP6
F0Kw/9nQjKVuP6fSxn5hI8ZYytZmCDlOOGFpAgMBAAGjITAfMB0GA1UdDgQWBBR/
xUxhpypAAtqzc/uooKxCLER3BTANBgkqhkiG9w0BAQsFAAOCAQEAb0fGOrQQU3EX
PVs73DKEqe1bGSPmsHIxl2HYUSFlh/VwCG5fvL8++jTHYCragPFv79EzrE9BUh87
+Hye31hXxR2TBRdYb4VutQ8o+4/U9y9GiYbpyc6hfjK15Xqs/pjxxa9HQ3NwzzIG
E3puly+3ppAZosbx28S+uHOVDsalrkeHjXEbZG1TkgVMHOC+ueS06hdBr49sf+42
IoM39VX5mrGfG93j5ALoo7el+e3gNW8V3BmXxqab1WbJkuWV0IEzoSTpeew1ibDJ
W1jf5CtfmklzKeKcYzetjdw9fW5riaM7085OLn/ov8EvmF4P0zy43EmZ1iUjznR5
4m4o3AHotw==
-----END CERTIFICATE-----`

	PAACertWithSameSubjectID2 = `-----BEGIN CERTIFICATE-----
MIIDbzCCAlegAwIBAgIUbItNCEkX2sqrtOnCtIPnAoDsuNQwDQYJKoZIhvcNAQEL
BQAwYDELMAkGA1UEBhMCVVoxDDAKBgNVBAgMA1RTSDERMA8GA1UEBwwIVEFTSEtF
TlQxDDAKBgNVBAoMA0RTUjEQMA4GA1UECwwHTUFUVEVSMjEQMA4GA1UEAwwHTUFU
VEVSMjAeFw0yNDAxMTgwOTA0MThaFw0yNTAxMTcwOTA0MThaMGAxCzAJBgNVBAYT
AlVaMQwwCgYDVQQIDANUU0gxETAPBgNVBAcMCFRBU0hLRU5UMQwwCgYDVQQKDANE
U1IxEDAOBgNVBAsMB01BVFRFUjIxEDAOBgNVBAMMB01BVFRFUjIwggEiMA0GCSqG
SIb3DQEBAQUAA4IBDwAwggEKAoIBAQDU0wWxxthYuUU0twy/MokJeXDKZqbrTGgW
cxBYSgDySqsxRoxD97SXQ/RyqmCHGlu7qBHvs6CGRX76IQfX2TpSGaNffYEwKfz5
0n9e6UHugJlPxSXIXiy54M6Z3/1+ySbwiDcOkNQklrlwWsSlL+g5ynDtfUhOu1/n
Xl68+1lU1EmO72CP/KnFLoL27WfuekvvWmC9ZO3OtzgIgvTLiquEzDez/l1UH9po
hhXzGx6Bukqr07OJFOu1JYzWwmeBw9X7F7Gd1onmlaGNRdsVJBdack49E10HVVEQ
TeW+vSfH3BiIiWLF1qySsRUwmFMZc4QD/3TJ75XoZsVbgSlQgESrAgMBAAGjITAf
MB0GA1UdDgQWBBR/xUxhpypAAtqzc/uooKxCLER3BTANBgkqhkiG9w0BAQsFAAOC
AQEAJscFyfnRDcsvHI7SLixSFFkcw2Us6j/9T73ccDBuJFbqGd+YYx7c1JVYU/qI
5MrsDg6D229qW+AFDuoLrX+u+rrSZbEADR5GVR66W6CgL4tQNkLWAAn5bJFV3Ugv
IBHlx8AWzAuLUwHE7hqnYz3ROq+u3idJMgUJtIQrXWqq9oqOl6VP1WC6e/RLmN36
a0u7AgG5ftl0Qx8nqht2Qdr+fIPivdR9JFweH/irVgTjgAo4Tj/0okgWdjS/qfIx
41rK5Rk4u875N8SOyJIU+DPbh7+bGV6QwX35aoPckfz12nC1uwDJ1SuLI55qy1cx
KER9kcWhvxgD0H/rtNnFVgbvQw==
-----END CERTIFICATE-----`

	RootCertWithSameSubjectAndSKID1 = `-----BEGIN CERTIFICATE-----
MIID0zCCArugAwIBAgIBATANBgkqhkiG9w0BAQsFADCBgjELMAkGA1UEBhMCVVMx
ETAPBgNVBAgMCE5ldyBZb3JrMREwDwYDVQQHDAhOZXcgWW9yazEYMBYGA1UECgwP
RXhhbXBsZSBDb21wYW55MRkwFwYDVQQLDBBUZXN0aW5nIERpdmlzaW9uMRgwFgYD
VQQDDA93d3cuZXhhbXBsZS5jb20wIBcNMjQwMjE1MTE1MDUwWhgPMzAyMzA2MTgx
MTUwNTBaMIGCMQswCQYDVQQGEwJVUzERMA8GA1UECAwITmV3IFlvcmsxETAPBgNV
BAcMCE5ldyBZb3JrMRgwFgYDVQQKDA9FeGFtcGxlIENvbXBhbnkxGTAXBgNVBAsM
EFRlc3RpbmcgRGl2aXNpb24xGDAWBgNVBAMMD3d3dy5leGFtcGxlLmNvbTCCASIw
DQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAMv3aL1eJ3UJsELx8Y6k/bCPnFHH
IMRR2Apxdov3bSb20qdTCb/tLbGNxeywlAXmZnY7CjRbEPPdTuZzNy9qm20MRX+o
CaYo9ysyWZGVuvXe2+Ma/cD4EiqUetB9ItYDfn9HQ04wAMIuntsNutOSGAmeCAtk
CICRsoSECWv0OsVFFstwhNOSgIq4B80YgA2YX4p7AC2LIxTjZ4x8QwmOFnaDMXIt
G/Ad1k+bEYBt/HwPnq5JR0Hpdx/aK/Cvabz+FD9y0Ohwy8Dgav9BjpWqv8ihZHnr
CJ2f3ccU0DjxVwOiL3KDrlQ7Iy5k2bxYjdvA8Fmvq4RgYm99j4LCir0M95cCAwEA
AaNQME4wHQYDVR0OBBYEFDNeDAdE+LWczVUBm21xI4Nv0NS+MB8GA1UdIwQYMBaA
FDNeDAdE+LWczVUBm21xI4Nv0NS+MAwGA1UdEwQFMAMBAf8wDQYJKoZIhvcNAQEL
BQADggEBADQxNXj4UBiIBG5wN6bOtMnVy9xw8OXUfihzfjGxLGjzvdUrRqb2tpDR
uIJEScHsUXAKsE3x0cuXMSf5aMiydNgQTpsmOBckL1huGqMg3LfloXQ6/pE+ATaB
rc292tYLmmCVHLB5kpFmMn8DsF6eTAPKuPtFeKRQ5c8IvW4E6CksK1JLTY73fWd/
p6lc3hTEQsQZsUwVzH74wu+whWZdKHHrEY7rONc/QiLmwZl+w2nGs+S62z20GueU
XSNIRw5NvAwLCvnog8A47MIqpuF211kdKvu2QFM/ekMvduL8BpkIFVKULSOY1t9d
XPz2ZlXABob+/ovGOyGPDw/3tUmlBXU=
-----END CERTIFICATE-----`

	RootCertWithSameSubjectAndSKID2 = `-----BEGIN CERTIFICATE-----
MIID0zCCArugAwIBAgIBAjANBgkqhkiG9w0BAQsFADCBgjELMAkGA1UEBhMCVVMx
ETAPBgNVBAgMCE5ldyBZb3JrMREwDwYDVQQHDAhOZXcgWW9yazEYMBYGA1UECgwP
RXhhbXBsZSBDb21wYW55MRkwFwYDVQQLDBBUZXN0aW5nIERpdmlzaW9uMRgwFgYD
VQQDDA93d3cuZXhhbXBsZS5jb20wIBcNMjQwMjE1MTE1MDU1WhgPMzAyMzA2MTgx
MTUwNTVaMIGCMQswCQYDVQQGEwJVUzERMA8GA1UECAwITmV3IFlvcmsxETAPBgNV
BAcMCE5ldyBZb3JrMRgwFgYDVQQKDA9FeGFtcGxlIENvbXBhbnkxGTAXBgNVBAsM
EFRlc3RpbmcgRGl2aXNpb24xGDAWBgNVBAMMD3d3dy5leGFtcGxlLmNvbTCCASIw
DQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAMv3aL1eJ3UJsELx8Y6k/bCPnFHH
IMRR2Apxdov3bSb20qdTCb/tLbGNxeywlAXmZnY7CjRbEPPdTuZzNy9qm20MRX+o
CaYo9ysyWZGVuvXe2+Ma/cD4EiqUetB9ItYDfn9HQ04wAMIuntsNutOSGAmeCAtk
CICRsoSECWv0OsVFFstwhNOSgIq4B80YgA2YX4p7AC2LIxTjZ4x8QwmOFnaDMXIt
G/Ad1k+bEYBt/HwPnq5JR0Hpdx/aK/Cvabz+FD9y0Ohwy8Dgav9BjpWqv8ihZHnr
CJ2f3ccU0DjxVwOiL3KDrlQ7Iy5k2bxYjdvA8Fmvq4RgYm99j4LCir0M95cCAwEA
AaNQME4wHQYDVR0OBBYEFDNeDAdE+LWczVUBm21xI4Nv0NS+MB8GA1UdIwQYMBaA
FDNeDAdE+LWczVUBm21xI4Nv0NS+MAwGA1UdEwQFMAMBAf8wDQYJKoZIhvcNAQEL
BQADggEBAENwaLKvyNz0IW1BNH2eTYNvgFl4f/I1pVYPOlA0O0ZB4BupqtgfKdTF
0DyPydGsZVMU/4MqEXe7qtvgQUbKk3kNTdi4s/y9rYKNyBKDfoTyDUtuDFHJNaOr
3F1Kjx1TWfu39UxSulutDTxnsdjVIxPxCT13vkekDYoILv746xWJaqkYR3L9u40w
583R+dA2BCVhPREnURneVcFuDcYuWNm3b6W5Vo0CBRhoIS1w+4Y+btRNL+0BvxBO
XWQLy4RZmLIkrm7vj3uFWRpi7lOBkAPOgCm04RTYqJJwnI0UeJmvaxoFd42J+k3D
xEsSrRoqMgkOX01+kkNn8Ugv3bEfeJ4=
-----END CERTIFICATE-----`

	IntermediateWithSameSubjectAndSKID1 = `-----BEGIN CERTIFICATE-----
MIIDlTCCAn2gAwIBAgIBAzANBgkqhkiG9w0BAQsFADCBgjELMAkGA1UEBhMCVVMx
ETAPBgNVBAgMCE5ldyBZb3JrMREwDwYDVQQHDAhOZXcgWW9yazEYMBYGA1UECgwP
RXhhbXBsZSBDb21wYW55MRkwFwYDVQQLDBBUZXN0aW5nIERpdmlzaW9uMRgwFgYD
VQQDDA93d3cuZXhhbXBsZS5jb20wIBcNMjQwMjE2MDY1NTA0WhgPMzAyMzA2MTkw
NjU1MDRaMEUxCzAJBgNVBAYTAkFVMRMwEQYDVQQIDApTb21lLVN0YXRlMSEwHwYD
VQQKDBhJbnRlcm5ldCBXaWRnaXRzIFB0eSBMdGQwggEiMA0GCSqGSIb3DQEBAQUA
A4IBDwAwggEKAoIBAQDmdvkss9fh7NRVNAKbrt9zReTNUwMsUxf6ryvepNUxEc6o
bGDS8yunS9QVpqf9BVQfM0tCzrewif1EDetdPYIMnC9o34nF095/4E5v+aoKabzG
wuyibKjtKVSl+dy0p42UJtScwzZRqLXIZxhmefh2CZT4q9Fs4y2qnCBtFqaWCToT
rcNWAQNBZ0E6S2ZulxXsdMoOGJ4iYPjAhqSbejcrN0McBudYq97pvEInG3HbyX2o
IGtZznGtwG326l9SV4OvmofxrvLjhx/nOauSBbbJcPWy3L10FWDZZ2h8ddpvx6I6
oCfqYRNMqsDPfd7eagbDObMApguhQ1Hl60NJW2KBAgMBAAGjUDBOMB0GA1UdDgQW
BBQuEztEUiww6ez7Rfpd5QQKwcbmuTAfBgNVHSMEGDAWgBQzXgwHRPi1nM1VAZtt
cSODb9DUvjAMBgNVHRMEBTADAQH/MA0GCSqGSIb3DQEBCwUAA4IBAQBm1hCN9XFd
2LnrxZkNr0RGMuSsFKIT8B7/gXmgDtkFKb84GcT/rFYGch4Nc6sVXQLFQWGsYCO9
OKZdiJGx4TOkGZuo1waa3/JzwDxLHh/2d7CLrEuGQnEo2evw+yj8AwkjWml+5zjn
L3bG9iZ1JQSUsGiVIHtUVYTGetyUy+SuJxuVvx359hc5foRuNiKDhQW7mwXWv+ua
xfBPUaspjiGALO8hBKlbVxt0RWv5MGyg2JJbSt9Ijexa6aoLzynq5gpSoEfQABUp
wbfDZe4Cbio4ndASlsbtpo/5ZOuQKn9Wp54meOotFDrFntnD7XFohxMJc5YY0F1q
Yk3FHd02VN0M
-----END CERTIFICATE-----`

	IntermediateWithSameSubjectAndSKID2 = `-----BEGIN CERTIFICATE-----
MIIDlTCCAn2gAwIBAgIBBDANBgkqhkiG9w0BAQsFADCBgjELMAkGA1UEBhMCVVMx
ETAPBgNVBAgMCE5ldyBZb3JrMREwDwYDVQQHDAhOZXcgWW9yazEYMBYGA1UECgwP
RXhhbXBsZSBDb21wYW55MRkwFwYDVQQLDBBUZXN0aW5nIERpdmlzaW9uMRgwFgYD
VQQDDA93d3cuZXhhbXBsZS5jb20wIBcNMjQwMjE2MDY1NzQ4WhgPMzAyMzA2MTkw
NjU3NDhaMEUxCzAJBgNVBAYTAkFVMRMwEQYDVQQIDApTb21lLVN0YXRlMSEwHwYD
VQQKDBhJbnRlcm5ldCBXaWRnaXRzIFB0eSBMdGQwggEiMA0GCSqGSIb3DQEBAQUA
A4IBDwAwggEKAoIBAQDmdvkss9fh7NRVNAKbrt9zReTNUwMsUxf6ryvepNUxEc6o
bGDS8yunS9QVpqf9BVQfM0tCzrewif1EDetdPYIMnC9o34nF095/4E5v+aoKabzG
wuyibKjtKVSl+dy0p42UJtScwzZRqLXIZxhmefh2CZT4q9Fs4y2qnCBtFqaWCToT
rcNWAQNBZ0E6S2ZulxXsdMoOGJ4iYPjAhqSbejcrN0McBudYq97pvEInG3HbyX2o
IGtZznGtwG326l9SV4OvmofxrvLjhx/nOauSBbbJcPWy3L10FWDZZ2h8ddpvx6I6
oCfqYRNMqsDPfd7eagbDObMApguhQ1Hl60NJW2KBAgMBAAGjUDBOMB0GA1UdDgQW
BBQuEztEUiww6ez7Rfpd5QQKwcbmuTAfBgNVHSMEGDAWgBQzXgwHRPi1nM1VAZtt
cSODb9DUvjAMBgNVHRMEBTADAQH/MA0GCSqGSIb3DQEBCwUAA4IBAQClVhy+Gtd7
5lOwDaT4m+J5FXyxMI6Lh7OydDcpsbUeIKtw0ODgCBFTrwtmYfqlJ35msKOwCOvp
L4LVxnUJGCV6dRuQgGFYnhJDwKdAI9aH2b51ZJoVuHxpoqpwkwFjYEzRzP1otkW9
b5VmtGeujfP12ptzqhmwmQ/z9yFDkKjwTcq02n1NGzc4CcygkQKZbr64HW17nsq/
AWNUMVCMnWrkYaFjBH21+RZ+zkdOVQXjtp7EZhQzLPhqrg4phLe1L+SztCJN/VrN
/sptlM3aW8XXI2pLbYAMCNpqQftnkdOOYfnJmf1A7x3eeDQHWiRj3FBHHGqZydmV
OtwplWOnxNQw
-----END CERTIFICATE-----`

	LeafCertWithSameSubjectAndSKID = `-----BEGIN CERTIFICATE-----
MIIDSTCCAjGgAwIBAgIBBTANBgkqhkiG9w0BAQsFADBFMQswCQYDVQQGEwJBVTET
MBEGA1UECAwKU29tZS1TdGF0ZTEhMB8GA1UECgwYSW50ZXJuZXQgV2lkZ2l0cyBQ
dHkgTHRkMCAXDTI0MDIxNjA2NTkwMloYDzMwMjMwNjE5MDY1OTAyWjBFMQswCQYD
VQQGEwJBVTETMBEGA1UECAwKU29tZS1TdGF0ZTEhMB8GA1UECgwYSW50ZXJuZXQg
V2lkZ2l0cyBQdHkgTHRkMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA
oCiuqhgIgU89pQ3/nT7ccFbqrmSbyRc/5jfYw83VDuSE4SkPwEml268k8XRoUjdh
sLkDROlR3hab0Ez/eI9LBLXXFWB3qlAJ2OfDkZ+RMLVIGqMxckn2bfnYtUITSVzD
FWWvanb7+wACW6MDH3f/FLLblII21XAcljluzPQbVZoCimh33CS5VOI68y/Esjlz
8HDhXDuxyswIjGMjCnKaFS0vQsySC+M4nqLKdJ0HDPNVMjeUtVIUCKke0rOFDTUZ
nDmBHfdOoHtBiCeZzNH7s2ER4NCwXZmSNXev8CHwmZqaIUSRMsmOxbw3I7KpRBFf
hObeLm55aVS2FEMH68H0FQIDAQABo0IwQDAdBgNVHQ4EFgQUEhZVjl4q3wTX5v7R
U2lhmO8XLwMwHwYDVR0jBBgwFoAULhM7RFIsMOns+0X6XeUECsHG5rkwDQYJKoZI
hvcNAQELBQADggEBAHK9fmY6C9FyzVNh6RTKNT6FL3ozr+WvmKJmE7WcxqAEW6JZ
rtihObu2y1B7e74umOwa1QJd7EFyMm4qnXYT2PepnanxTnz0EST9ZuhM3GpM1FP6
fjlqLDHoQ1UhBmEnocFTqd7QEZtUbRWPnlJw0ZK2uFK7IYmlnBKkewPCLVGI3ihx
al/8sTx3xx7fWpS+rJ3jviCpHgP+cGV/ANg8hOlyr68u0FE+x6pye00TmxcFzDuo
5/OA9jGQln82Z8inmc05wZPQPpjZxdCQteqJkNl7PrklgO5EevG9JlUArIets2Py
2Vciq5eYOIi+PlP+HI5QzlZYxSqFjJrFcfzYCJ4=
-----END CERTIFICATE-----`

	RootCertWithVid = `-----BEGIN CERTIFICATE-----
MIICdDCCAhmgAwIBAgIBATAKBggqhkjOPQQDAjCBmDELMAkGA1UEBhMCVVMxETAP
BgNVBAgMCE5ldyBZb3JrMREwDwYDVQQHDAhOZXcgWW9yazEYMBYGA1UECgwPRXhh
bXBsZSBDb21wYW55MRkwFwYDVQQLDBBUZXN0aW5nIERpdmlzaW9uMRgwFgYDVQQD
DA93d3cuZXhhbXBsZS5jb20xFDASBgorBgEEAYKifAIBDARGRkYxMCAXDTI0MDIy
NjExNTQzMVoYDzMwMjMwNjI5MTE1NDMxWjCBmDELMAkGA1UEBhMCVVMxETAPBgNV
BAgMCE5ldyBZb3JrMREwDwYDVQQHDAhOZXcgWW9yazEYMBYGA1UECgwPRXhhbXBs
ZSBDb21wYW55MRkwFwYDVQQLDBBUZXN0aW5nIERpdmlzaW9uMRgwFgYDVQQDDA93
d3cuZXhhbXBsZS5jb20xFDASBgorBgEEAYKifAIBDARGRkYxMFkwEwYHKoZIzj0C
AQYIKoZIzj0DAQcDQgAEDcguargOjH5nh4SCsflFk1ACqNBOR6Wua8huVYPBfse6
uwfkgmyTJrCBCUAq9ayPD83jPVor1NN9YAx/V0zbsKNQME4wHQYDVR0OBBYEFM6o
kmbq4IC9K7Vo5AsHxPosNG0xMB8GA1UdIwQYMBaAFM6okmbq4IC9K7Vo5AsHxPos
NG0xMAwGA1UdEwQFMAMBAf8wCgYIKoZIzj0EAwIDSQAwRgIhAOdYHo1krgzyV+CT
G+RKcYoxHr6YS9ddNOJibjBx/I63AiEAxNl6kcOH0Rovwi2wySHvTD26kfUYJAmi
HGBcCo5whZU=
-----END CERTIFICATE-----`

	IntermediateCertWithVid1 = `-----BEGIN CERTIFICATE-----
MIICiTCCAi+gAwIBAgIBAzAKBggqhkjOPQQDAjCBmDELMAkGA1UEBhMCVVMxETAP
BgNVBAgMCE5ldyBZb3JrMREwDwYDVQQHDAhOZXcgWW9yazEYMBYGA1UECgwPRXhh
bXBsZSBDb21wYW55MRkwFwYDVQQLDBBUZXN0aW5nIERpdmlzaW9uMRgwFgYDVQQD
DA93d3cuZXhhbXBsZS5jb20xFDASBgorBgEEAYKifAIBDARGRkYxMCAXDTI0MDMy
NzA2MDcxMloYDzMwMjMwNzI5MDYwNzEyWjCBrjELMAkGA1UEBhMCVVMxETAPBgNV
BAgMCE5ldyBZb3JrMREwDwYDVQQHDAhOZXcgWW9yazEYMBYGA1UECgwPRXhhbXBs
ZSBDb21wYW55MRkwFwYDVQQLDBBUZXN0aW5nIERpdmlzaW9uMRgwFgYDVQQDDA93
d3cuZXhhbXBsZS5jb20xFDASBgorBgEEAYKifAIBDARGRkYxMRQwEgYKKwYBBAGC
onwCAgwERkZGMTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABOTNh8u27CnRGdj0
G0/z0oo9rsKcpgUogQ8fYYEg/QClYFHJuhFbf1M+VdeMScbllpt4kGH2ih7aU1b7
1jRkVsyjUDBOMB0GA1UdDgQWBBQOjOjIuKpQvCWFVrmxnMLH2cUvFzAfBgNVHSME
GDAWgBTOqJJm6uCAvSu1aOQLB8T6LDRtMTAMBgNVHRMEBTADAQH/MAoGCCqGSM49
BAMCA0gAMEUCIQCy8SeF6UXIGM+0X6fc5tqSrgAQ1nCN5cvsWyfZvH0y9wIgQ45S
TXQomsOa4eHQpJzsY/JQqprA0FapY1nsvL+PQFg=
-----END CERTIFICATE-----`

	IntermediateCertWithVid2 = `-----BEGIN CERTIFICATE-----
MIICiDCCAi+gAwIBAgIBBDAKBggqhkjOPQQDAjCBmDELMAkGA1UEBhMCVVMxETAP
BgNVBAgMCE5ldyBZb3JrMREwDwYDVQQHDAhOZXcgWW9yazEYMBYGA1UECgwPRXhh
bXBsZSBDb21wYW55MRkwFwYDVQQLDBBUZXN0aW5nIERpdmlzaW9uMRgwFgYDVQQD
DA93d3cuZXhhbXBsZS5jb20xFDASBgorBgEEAYKifAIBDARGRkYxMCAXDTI0MDMy
NzE1MzQxMVoYDzMwMjMwNzI5MTUzNDExWjCBrjELMAkGA1UEBhMCVVMxETAPBgNV
BAgMCE5ldyBZb3JrMREwDwYDVQQHDAhOZXcgWW9yazEYMBYGA1UECgwPRXhhbXBs
ZSBDb21wYW55MRkwFwYDVQQLDBBUZXN0aW5nIERpdmlzaW9uMRgwFgYDVQQDDA93
d3cuZXhhbXBsZS5jb20xFDASBgorBgEEAYKifAIBDARGRkYyMRQwEgYKKwYBBAGC
onwCAgwERkZGMjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABOTNh8u27CnRGdj0
G0/z0oo9rsKcpgUogQ8fYYEg/QClYFHJuhFbf1M+VdeMScbllpt4kGH2ih7aU1b7
1jRkVsyjUDBOMB0GA1UdDgQWBBQOjOjIuKpQvCWFVrmxnMLH2cUvFzAfBgNVHSME
GDAWgBTOqJJm6uCAvSu1aOQLB8T6LDRtMTAMBgNVHRMEBTADAQH/MAoGCCqGSM49
BAMCA0cAMEQCIHkhL7r/xEi16827IYysHe0w8X0rsbU5zcHcbK1wt0ALAiASEZMI
NN1ZIQJHBjCm+vWh3Jsjt2wUHKIM5i64Wd9kPA==
-----END CERTIFICATE-----`

	IntermediateCertWithoutVidPid = `-----BEGIN CERTIFICATE-----
MIICfjCCAiOgAwIBAgIUApsGBeXsNPxNq4brOXLNfbYysakwCgYIKoZIzj0EAwIw
gZgxCzAJBgNVBAYTAlVTMREwDwYDVQQIDAhOZXcgWW9yazERMA8GA1UEBwwITmV3
IFlvcmsxGDAWBgNVBAoMD0V4YW1wbGUgQ29tcGFueTEZMBcGA1UECwwQVGVzdGlu
ZyBEaXZpc2lvbjEYMBYGA1UEAwwPd3d3LmV4YW1wbGUuY29tMRQwEgYKKwYBBAGC
onwCAQwERkZGMTAgFw0yNDAzMjgxMzEzMjVaGA8zMDIzMDczMDEzMTMyNVowgYIx
CzAJBgNVBAYTAlVTMREwDwYDVQQIDAhOZXcgWW9yazERMA8GA1UEBwwITmV3IFlv
cmsxGDAWBgNVBAoMD0V4YW1wbGUgQ29tcGFueTEZMBcGA1UECwwQVGVzdGluZyBE
aXZpc2lvbjEYMBYGA1UEAwwPd3d3LmV4YW1wbGUuY29tMFkwEwYHKoZIzj0CAQYI
KoZIzj0DAQcDQgAE5M2Hy7bsKdEZ2PQbT/PSij2uwpymBSiBDx9hgSD9AKVgUcm6
EVt/Uz5V14xJxuWWm3iQYfaKHtpTVvvWNGRWzKNdMFswHQYDVR0OBBYEFA6M6Mi4
qlC8JYVWubGcwsfZxS8XMB8GA1UdIwQYMBaAFM6okmbq4IC9K7Vo5AsHxPosNG0x
MAkGA1UdEwQCMAAwDgYDVR0PAQH/BAQDAgGCMAoGCCqGSM49BAMCA0kAMEYCIQDm
jhpYAW9UseDLyoF2bmvy36jV7Hwvst+R3wJi0jh4xAIhAPXCfe8DUCoRV32q97C0
IYJElzT/KwBY6c2Xyu4gsjqh
-----END CERTIFICATE-----`

	LeafCertWithVid = `-----BEGIN CERTIFICATE-----
MIICrjCCAlSgAwIBAgIUBCg+BsyaPLK2sNxttFUIbDF/FPAwCgYIKoZIzj0EAwIw
ga4xCzAJBgNVBAYTAlVTMREwDwYDVQQIDAhOZXcgWW9yazERMA8GA1UEBwwITmV3
IFlvcmsxGDAWBgNVBAoMD0V4YW1wbGUgQ29tcGFueTEZMBcGA1UECwwQVGVzdGlu
ZyBEaXZpc2lvbjEYMBYGA1UEAwwPd3d3LmV4YW1wbGUuY29tMRQwEgYKKwYBBAGC
onwCAQwERkZGMTEUMBIGCisGAQQBgqJ8AgIMBEZGRjEwIBcNMjQwMzI2MTAyNDI1
WhgPMzAyMzA3MjgxMDI0MjVaMIGaMQswCQYDVQQGEwJVUzERMA8GA1UECAwITmV3
IFlvcmsxETAPBgNVBAcMCE5ldyBZb3JrMRowGAYDVQQKDBFDUkwtbGVhZiB3aXRo
IFZJRDEZMBcGA1UECwwQVGVzdGluZyBEaXZpc2lvbjEYMBYGA1UEAwwPd3d3LmV4
YW1wbGUuY29tMRQwEgYKKwYBBAGConwCAQwERkZGMTBZMBMGByqGSM49AgEGCCqG
SM49AwEHA0IABNk/8AZJsYEd7kBVDv5c+Mm4kNsuyMF1d+UTOTlptsCzx4YwLlCX
SSr2SwDHbkRvMbp5cfFt9uyNc0Tx3bVVyPWjYDBeMB0GA1UdDgQWBBTWmCYQvqwj
dAkKQAvNOWVT8Xaw9TAfBgNVHSMEGDAWgBQOjOjIuKpQvCWFVrmxnMLH2cUvFzAM
BgNVHRMBAf8EAjAAMA4GA1UdDwEB/wQEAwIBgjAKBggqhkjOPQQDAgNIADBFAiEA
nAoa731+XkR5/0XaESqHG40IZysduxN8sJo2sJpPvvwCICGn7oAwDmQh0umEJ6dK
Vtv3RJ9iuKtC/fkzUzhv9c0z
-----END CERTIFICATE-----`

	LeafCertWithVidPid = `-----BEGIN CERTIFICATE-----
MIICzDCCAnKgAwIBAgIUG6W5A5QhAdUKiVAG9yo5VrndE2IwCgYIKoZIzj0EAwIw
ga4xCzAJBgNVBAYTAlVTMREwDwYDVQQIDAhOZXcgWW9yazERMA8GA1UEBwwITmV3
IFlvcmsxGDAWBgNVBAoMD0V4YW1wbGUgQ29tcGFueTEZMBcGA1UECwwQVGVzdGlu
ZyBEaXZpc2lvbjEYMBYGA1UEAwwPd3d3LmV4YW1wbGUuY29tMRQwEgYKKwYBBAGC
onwCAQwERkZGMTEUMBIGCisGAQQBgqJ8AgIMBEZGRjEwIBcNMjQwMzI2MTAzNTI4
WhgPMzAyMzA3MjgxMDM1MjhaMIG4MQswCQYDVQQGEwJVUzERMA8GA1UECAwITmV3
IFlvcmsxETAPBgNVBAcMCE5ldyBZb3JrMSIwIAYDVQQKDBlDUkwtbGVhZiB3aXRo
IFZJRCBhbmQgUElEMRkwFwYDVQQLDBBUZXN0aW5nIERpdmlzaW9uMRgwFgYDVQQD
DA93d3cuZXhhbXBsZS5jb20xFDASBgorBgEEAYKifAIBDARGRkYxMRQwEgYKKwYB
BAGConwCAgwERkZGMTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABNk/8AZJsYEd
7kBVDv5c+Mm4kNsuyMF1d+UTOTlptsCzx4YwLlCXSSr2SwDHbkRvMbp5cfFt9uyN
c0Tx3bVVyPWjYDBeMB0GA1UdDgQWBBTWmCYQvqwjdAkKQAvNOWVT8Xaw9TAfBgNV
HSMEGDAWgBQOjOjIuKpQvCWFVrmxnMLH2cUvFzAMBgNVHRMBAf8EAjAAMA4GA1Ud
DwEB/wQEAwIBgjAKBggqhkjOPQQDAgNIADBFAiEAhs/qxSBUSsRdqXfC8tQlPIPU
CNbAI81hYOHbiOx6fD0CIFz63D+Ug7xurPSqAPHoTAY6MhseK4IrbAjKRPA0sQl5
-----END CERTIFICATE-----`

	LeafCertWithoutVidPid = `-----BEGIN CERTIFICATE-----
MIICozCCAkmgAwIBAgIUDXi3VEZsSRTrSqZuIqDWX0Ar4egwCgYIKoZIzj0EAwIw
ga4xCzAJBgNVBAYTAlVTMREwDwYDVQQIDAhOZXcgWW9yazERMA8GA1UEBwwITmV3
IFlvcmsxGDAWBgNVBAoMD0V4YW1wbGUgQ29tcGFueTEZMBcGA1UECwwQVGVzdGlu
ZyBEaXZpc2lvbjEYMBYGA1UEAwwPd3d3LmV4YW1wbGUuY29tMRQwEgYKKwYBBAGC
onwCAQwERkZGMTEUMBIGCisGAQQBgqJ8AgIMBEZGRjEwIBcNMjQwMzI2MTEwNjIz
WhgPMzAyMzA3MjgxMTA2MjNaMIGPMQswCQYDVQQGEwJVUzERMA8GA1UECAwITmV3
IFlvcmsxETAPBgNVBAcMCE5ldyBZb3JrMSUwIwYDVQQKDBxDUkwtbGVhZiB3aXRo
b3V0IFZJRCBhbmQgUElEMRkwFwYDVQQLDBBUZXN0aW5nIERpdmlzaW9uMRgwFgYD
VQQDDA93d3cuZXhhbXBsZS5jb20wWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATZ
P/AGSbGBHe5AVQ7+XPjJuJDbLsjBdXflEzk5abbAs8eGMC5Ql0kq9ksAx25EbzG6
eXHxbfbsjXNE8d21Vcj1o2AwXjAdBgNVHQ4EFgQU1pgmEL6sI3QJCkALzTllU/F2
sPUwHwYDVR0jBBgwFoAUDozoyLiqULwlhVa5sZzCx9nFLxcwDAYDVR0TAQH/BAIw
ADAOBgNVHQ8BAf8EBAMCAYIwCgYIKoZIzj0EAwIDSAAwRQIhAPIzS2Tlov+9/R6U
fJhEWAA8mOgN9OVCdPWAegWuN3b2AiApXciu/dT4B5db3puPWrAsMjAUYF2Owc/D
eujhLsD51w==
-----END CERTIFICATE-----`

	RootIssuer                     = "MDQxCzAJBgNVBAYTAkFVMRMwEQYDVQQIDApzb21lLXN0YXRlMRAwDgYDVQQKDAdyb290LWNh"
	RootSubject                    = "MDQxCzAJBgNVBAYTAkFVMRMwEQYDVQQIDApzb21lLXN0YXRlMRAwDgYDVQQKDAdyb290LWNh"
	RootSubjectAsText              = "O=root-ca,ST=some-state,C=AU"
	RootSubjectKeyID               = "5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB"
	RootSubjectKeyIDWithoutColumns = "5A880E6C3653D07FB08971A3F473790930E62BDB"
	RootSerialNumber               = "442314047376310867378175982234956458728610743315"

	RootCertWithSameSubjectAndSKIDSubject         = "MIGCMQswCQYDVQQGEwJVUzERMA8GA1UECAwITmV3IFlvcmsxETAPBgNVBAcMCE5ldyBZb3JrMRgwFgYDVQQKDA9FeGFtcGxlIENvbXBhbnkxGTAXBgNVBAsMEFRlc3RpbmcgRGl2aXNpb24xGDAWBgNVBAMMD3d3dy5leGFtcGxlLmNvbQ=="
	RootCertWithSameSubjectAndSKIDSubjectKeyID    = "33:5E:0C:07:44:F8:B5:9C:CD:55:01:9B:6D:71:23:83:6F:D0:D4:BE"
	RootCertWithSameSubjectAndSKID1SerialNumber   = "1"
	RootCertWithSameSubjectAndSKID2SerialNumber   = "2"
	IntermediateCertWithSameSubjectAndSKIDSubject = "MEUxCzAJBgNVBAYTAkFVMRMwEQYDVQQIDApTb21lLVN0YXRlMSEwHwYDVQQKDBhJbnRlcm5ldCBXaWRnaXRzIFB0eSBMdGQ="

	IntermediateCertWithSameSubjectAndSKIDSubjectKeyID  = "2E:13:3B:44:52:2C:30:E9:EC:FB:45:FA:5D:E5:04:0A:C1:C6:E6:B9"
	IntermediateCertWithSameSubjectAndSKID1SerialNumber = "3"
	IntermediateCertWithSameSubjectAndSKID2SerialNumber = "4"
	LeafCertWithSameSubjectAndSKIDSubject               = "MEUxCzAJBgNVBAYTAkFVMRMwEQYDVQQIDApTb21lLVN0YXRlMSEwHwYDVQQKDBhJbnRlcm5ldCBXaWRnaXRzIFB0eSBMdGQ="
	LeafCertWithSameSubjectAndSKIDSubjectKeyID          = "12:16:55:8E:5E:2A:DF:04:D7:E6:FE:D1:53:69:61:98:EF:17:2F:03"
	LeafCertWithSameSubjectAndSKIDSerialNumber          = "5"

	IntermediateIssuer                     = "MDQxCzAJBgNVBAYTAkFVMRMwEQYDVQQIDApzb21lLXN0YXRlMRAwDgYDVQQKDAdyb290LWNh"
	IntermediateAuthorityKeyID             = "5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB"
	IntermediateSubject                    = "MDwxCzAJBgNVBAYTAkFVMRMwEQYDVQQIDApzb21lLXN0YXRlMRgwFgYDVQQKDA9pbnRlcm1lZGlhdGUtY2E="
	IntermediateSubjectAsText              = "O=intermediate-ca,ST=some-state,C=AU"
	IntermediateSubjectKeyID               = "4E:3B:73:F4:70:4D:C2:98:0D:DB:C8:5A:5F:02:3B:BF:86:25:56:2B"
	IntermediateSubjectKeyIDWithoutColumns = "4E3B73F4704DC2980DDBC85A5F023BBF8625562B"
	IntermediateSerialNumber               = "169917617234879872371588777545667947720450185023"

	LeafIssuer         = "MDwxCzAJBgNVBAYTAkFVMRMwEQYDVQQIDApzb21lLXN0YXRlMRgwFgYDVQQKDA9pbnRlcm1lZGlhdGUtY2E="
	LeafAuthorityKeyID = "4E:3B:73:F4:70:4D:C2:98:D:DB:C8:5A:5F:02:3B:BF:86:25:56:2B"
	LeafSubject        = "MDExCzAJBgNVBAYTAkFVMRMwEQYDVQQIDApzb21lLXN0YXRlMQ0wCwYDVQQKDARsZWFm"
	LeafSubjectAsText  = "O=leaf,ST=some-state,C=AU"
	LeafSubjectKeyID   = "30:F4:65:75:14:20:B2:AF:3D:14:71:17:AC:49:90:93:3E:24:A0:1F"
	LeafSerialNumber   = "143290473708569835418599774898811724528308722063"

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

	PAACertWithSameSubjectID1Subject  = "MFoxCzAJBgNVBAYTAlVaMQwwCgYDVQQIDANUU0gxETAPBgNVBAcMCFRBU0hLRU5UMQwwCgYDVQQKDANEU1IxCzAJBgNVBAsMAkRDMQ8wDQYDVQQDDAZNQVRURVI="
	PAACertWithSameSubjectID2Subject  = "MGAxCzAJBgNVBAYTAlVaMQwwCgYDVQQIDANUU0gxETAPBgNVBAcMCFRBU0hLRU5UMQwwCgYDVQQKDANEU1IxEDAOBgNVBAsMB01BVFRFUjIxEDAOBgNVBAMMB01BVFRFUjI="
	PAACertWithSameSubjectIDSubjectID = "7F:C5:4C:61:A7:2A:40:02:DA:B3:73:FB:A8:A0:AC:42:2C:44:77:05"

	TestVID1String            = "0xA13"
	TestPID1String            = "0xA11"
	TestVID2String            = "0xA14"
	TestPID2String            = "0xA15"
	TestVID3String            = "0xA16"
	TestPID3String            = "0xA17"
	SubjectKeyIDWithoutColons = "5A880E6C3653D07FB08971A3F473790930E62BDB"
	DataDigest                = "9a5d2c1f4b3e6f8d7b1a0c9e2f5d8b7"

	TestCertPemVid = 4701

	RootCertWithVidSubject                    = "MIGYMQswCQYDVQQGEwJVUzERMA8GA1UECAwITmV3IFlvcmsxETAPBgNVBAcMCE5ldyBZb3JrMRgwFgYDVQQKDA9FeGFtcGxlIENvbXBhbnkxGTAXBgNVBAsMEFRlc3RpbmcgRGl2aXNpb24xGDAWBgNVBAMMD3d3dy5leGFtcGxlLmNvbTEUMBIGCisGAQQBgqJ8AgEMBEZGRjE="
	RootCertWithVidSubjectKeyID               = "CE:A8:92:66:EA:E0:80:BD:2B:B5:68:E4:0B:07:C4:FA:2C:34:6D:31"
	RootCertWithVidSubjectKeyIDWithoutColumns = "CEA89266EAE080BD2BB568E40B07C4FA2C346D31"
	RootCertWithVidVid                        = 65521

	IntermediateCertWithVid1Subject                    = "MIGuMQswCQYDVQQGEwJVUzERMA8GA1UECAwITmV3IFlvcmsxETAPBgNVBAcMCE5ldyBZb3JrMRgwFgYDVQQKDA9FeGFtcGxlIENvbXBhbnkxGTAXBgNVBAsMEFRlc3RpbmcgRGl2aXNpb24xGDAWBgNVBAMMD3d3dy5leGFtcGxlLmNvbTEUMBIGCisGAQQBgqJ8AgEMBEZGRjExFDASBgorBgEEAYKifAICDARGRkYx"
	IntermediateCertWithVid1SubjectKeyID               = "0E:8C:E8:C8:B8:AA:50:BC:25:85:56:B9:B1:9C:C2:C7:D9:C5:2F:17"
	IntermediateCertWithVid1SubjectKeyIDWithoutColumns = "0E8CE8C8B8AA50BC258556B9B19CC2C7D9C52F17"
	IntermediateCertWithVid1SerialNumber               = "3"
	IntermediateCertWithVid1Vid                        = 65521

	IntermediateCertWithVid2SubjectKeyID = "0E:8C:E8:C8:B8:AA:50:BC:25:85:56:B9:B1:9C:C2:C7:D9:C5:2F:17"
	IntermediateCertWithVid2SerialNumber = "4"
	IntermediateCertWithVid2Vid          = 65522

	IntermediateCertWithoutVidPidSubject      = "MIGCMQswCQYDVQQGEwJVUzERMA8GA1UECAwITmV3IFlvcmsxETAPBgNVBAcMCE5ldyBZb3JrMRgwFgYDVQQKDA9FeGFtcGxlIENvbXBhbnkxGTAXBgNVBAsMEFRlc3RpbmcgRGl2aXNpb24xGDAWBgNVBAMMD3d3dy5leGFtcGxlLmNvbQ=="
	IntermediateCertWithoutVidPidSubjectKeyID = "0E:8C:E8:C8:B8:AA:50:BC:25:85:56:B9:B1:9C:C2:C7:D9:C5:2F:17"
	IntermediateCertWithoutVidPidSerialNumber = "14875121728167018569770528052537472929544450473"

	LeafCertWithVidSubject        = "MIGaMQswCQYDVQQGEwJVUzERMA8GA1UECAwITmV3IFlvcmsxETAPBgNVBAcMCE5ldyBZb3JrMRowGAYDVQQKDBFDUkwtbGVhZiB3aXRoIFZJRDEZMBcGA1UECwwQVGVzdGluZyBEaXZpc2lvbjEYMBYGA1UEAwwPd3d3LmV4YW1wbGUuY29tMRQwEgYKKwYBBAGConwCAQwERkZGMQ=="
	LeafCertWithVidSubjectAsText  = "CN=www.example.com,OU=Testing Division,O=CRL-leaf with VID,L=New York,ST=New York,C=US,vid=0xFFF1"
	LeafCertWithVidSubjectKeyID   = "D6:98:26:10:BE:AC:23:74:09:0A:40:0B:CD:39:65:53:F1:76:B0:F5"
	LeafCertWithVidAuthorityKeyID = IntermediateCertWithVid1SubjectKeyID
	LeafCertWithVidSerialNumber   = "23733396166621909643583307546615137635389084912"
	LeafCertWithVidVid            = 65521

	LeafCertWithVidPidSubject        = "MIG4MQswCQYDVQQGEwJVUzERMA8GA1UECAwITmV3IFlvcmsxETAPBgNVBAcMCE5ldyBZb3JrMSIwIAYDVQQKDBlDUkwtbGVhZiB3aXRoIFZJRCBhbmQgUElEMRkwFwYDVQQLDBBUZXN0aW5nIERpdmlzaW9uMRgwFgYDVQQDDA93d3cuZXhhbXBsZS5jb20xFDASBgorBgEEAYKifAIBDARGRkYxMRQwEgYKKwYBBAGConwCAgwERkZGMQ=="
	LeafCertWithVidPidSubjectAsText  = "CN=www.example.com,OU=Testing Division,O=CRL-leaf with VID and PID,L=New York,ST=New York,C=US,pid=0xFFF1,vid=0xFFF1"
	LeafCertWithVidPidSubjectKeyID   = "D6:98:26:10:BE:AC:23:74:09:0A:40:0B:CD:39:65:53:F1:76:B0:F5"
	LeafCertWithVidPidAuthorityKeyID = IntermediateCertWithVid1SubjectKeyID
	LeafCertWithVidPidSerialNumber   = "157838490760642822714861562571853387507185816418"
	LeafCertWithVidPidVid            = 65521
	LeafCertWithVidPidPid            = 65521

	LeafCertWithoutVidPidSubject        = "MIGPMQswCQYDVQQGEwJVUzERMA8GA1UECAwITmV3IFlvcmsxETAPBgNVBAcMCE5ldyBZb3JrMSUwIwYDVQQKDBxDUkwtbGVhZiB3aXRob3V0IFZJRCBhbmQgUElEMRkwFwYDVQQLDBBUZXN0aW5nIERpdmlzaW9uMRgwFgYDVQQDDA93d3cuZXhhbXBsZS5jb20="
	LeafCertWithoutVidPidSubjectAsText  = "CN=www.example.com,OU=Testing Division,O=CRL-leaf without VID and PID,L=New York,ST=New York,C=US"
	LeafCertWithoutVidPidSubjectKeyID   = "D6:98:26:10:BE:AC:23:74:09:0A:40:0B:CD:39:65:53:F1:76:B0:F5"
	LeafCertWithoutVidPidAuthorityKeyID = IntermediateCertWithVid1SubjectKeyID
	LeafCertWithoutVidPidSerialNumber   = "76908939670186132114931832808683834138281370088"
)
