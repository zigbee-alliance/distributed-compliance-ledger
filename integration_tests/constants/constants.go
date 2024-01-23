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
	"github.com/zigbee-alliance/distributed-compliance-ledger/app"
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
	defEncConfig = app.MakeEncodingConfig()

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
	LsfURL                                            = "https://url.lsfurl.dclmodel"
	DataURL                                           = "https://url.data.dclmodel"
	DataURL2                                          = "https://url.data.dclmodel2"
	URLWithoutProtocol                                = "url.dclmodel"
	LsfRevision                                int32  = 1
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
	PAACertWithNumericVidSubject      = "MDAxGDAWBgNVBAMMD01hdHRlciBUZXN0IFBBQTEUMBIGCisGAQQBgqJ8AgEMBEZGRjE="
	PAACertWithNumericVidSubjectKeyID = "6A:FD:22:77:1F:51:1F:EC:BF:16:41:97:67:10:DC:DC:31:A1:71:7E"
	PAACertWithNumericVidVid          = 65521

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
	PAICertWithNumericPidVidVid = 65521
	PAICertWithNumericPidVidPid = 32768

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

	RootIssuer        = "MDQxCzAJBgNVBAYTAkFVMRMwEQYDVQQIDApzb21lLXN0YXRlMRAwDgYDVQQKDAdyb290LWNh"
	RootSubject       = "MDQxCzAJBgNVBAYTAkFVMRMwEQYDVQQIDApzb21lLXN0YXRlMRAwDgYDVQQKDAdyb290LWNh"
	RootSubjectAsText = "O=root-ca,ST=some-state,C=AU"
	RootSubjectKeyID  = "5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB"
	RootSerialNumber  = "442314047376310867378175982234956458728610743315"

	IntermediateIssuer         = "MDQxCzAJBgNVBAYTAkFVMRMwEQYDVQQIDApzb21lLXN0YXRlMRAwDgYDVQQKDAdyb290LWNh"
	IntermediateAuthorityKeyID = "5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB"
	IntermediateSubject        = "MDwxCzAJBgNVBAYTAkFVMRMwEQYDVQQIDApzb21lLXN0YXRlMRgwFgYDVQQKDA9pbnRlcm1lZGlhdGUtY2E="
	IntermediateSubjectAsText  = "O=intermediate-ca,ST=some-state,C=AU"
	IntermediateSubjectKeyID   = "4E:3B:73:F4:70:4D:C2:98:0D:DB:C8:5A:5F:02:3B:BF:86:25:56:2B"
	IntermediateSerialNumber   = "169917617234879872371588777545667947720450185023"

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
)
