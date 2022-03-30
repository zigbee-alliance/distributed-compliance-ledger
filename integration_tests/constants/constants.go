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
	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
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
	defEncConfig = simapp.MakeTestEncodingConfig()

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
	FirmwareDigests                                   = "Firmware Digest String"
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
	ProvisionalDate   = "2019-12-12T00:00:00Z"
	CertificationDate = "2020-01-01T00:00:00Z"
	RevocationDate    = "2020-03-03T03:30:00Z"
	Reason            = "Some Reason"
	RevocationReason  = "Some Reason"
	CertificationType = "zigbee"

	// Testing Result.
	TestResult = "http://test.result.com"
	TestDate   = "2020-02-02T02:00:00Z"

	// Upgrade.
	UpgradePlanName         = "TestUpgrade"
	UpgradePlanHeight int64 = 1337
	UpgradePlanInfo         = "Some upgrade info"

	//
	Address1, _       = sdk.AccAddressFromBech32("cosmos1s5xf3aanx7w84hgplk9z3l90qfpantg6nsmhpf")
	Address2, _       = sdk.AccAddressFromBech32("cosmos1nl4uaesk9gtu7su3n89lne6xpa6lq8gljn79rq")
	Address3, _       = sdk.AccAddressFromBech32("cosmos12r9vsus5js32pvnayt33zhcd4y9wcqcly45gr9")
	VendorID1   int32 = 1000
	VendorID2   int32 = 2000
	VendorID3   int32 = 3000
	PubKey1           = strToPubKey(
		`{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"Aw1XXHQ8i6JVNKsFQ9eQArJVt2GXEO0EBFsQL6XJ5BxY"}`,
		defEncConfig.Marshaler,
	)
	PubKey2 = strToPubKey(
		`{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"A2wJ7uOEE5Zm04K52czFTXfDj1qF2mholzi1zOJVlKlr"}`,
		defEncConfig.Marshaler,
	)
	PubKey3 = strToPubKey(
		`{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"A0GnKr6hItYE1A7dzoxNSMwMZuu1zauOLWAqJWen1RzF"}`,
		defEncConfig.Marshaler,
	)
	Signer           = Address1
	ValidatorPubKey1 = strToPubKey(
		`{"@type":"/cosmos.crypto.ed25519.PubKey","key":"1e+1/jHGaJi0b2zgCN46eelKCYpKiuTgPN18mL3fzx8="}`,
		defEncConfig.Marshaler,
	)
	ValidatorPubKey2 = strToPubKey(
		`{"@type":"/cosmos.crypto.ed25519.PubKey","key":"NB8hcdxKYDCaPWR67OiUXUSltZfYYOWYryPDUdbWRlA="}`,
		defEncConfig.Marshaler,
	)
	ValidatorAddress1 = "cosmosvalcons1uks7yvlwqsfyp730w6da64g5fw20d9ynh00k53"
	ValidatorAddress2 = "cosmosvalcons12tg2p3rjsaczddufmsjjrw9nvhg8wkc4hcz3zw"
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

	RootIssuer        = "Tz1yb290LWNhLFNUPXNvbWUtc3RhdGUsQz1BVQ=="
	RootSubject       = "Tz1yb290LWNhLFNUPXNvbWUtc3RhdGUsQz1BVQ=="
	RootSubjectAsText = "O=root-ca,ST=some-state,C=AU"
	RootSubjectKeyID  = "5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB"
	RootSerialNumber  = "442314047376310867378175982234956458728610743315"

	IntermediateIssuer         = "Tz1yb290LWNhLFNUPXNvbWUtc3RhdGUsQz1BVQ=="
	IntermediateAuthorityKeyID = "5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB"
	IntermediateSubject        = "Tz1pbnRlcm1lZGlhdGUtY2EsU1Q9c29tZS1zdGF0ZSxDPUFV"
	IntermediateSubjectAsText  = "O=intermediate-ca,ST=some-state,C=AU"
	IntermediateSubjectKeyID   = "4E:3B:73:F4:70:4D:C2:98:0D:DB:C8:5A:5F:02:3B:BF:86:25:56:2B"
	IntermediateSerialNumber   = "169917617234879872371588777545667947720450185023"

	LeafIssuer         = "Tz1pbnRlcm1lZGlhdGUtY2EsU1Q9c29tZS1zdGF0ZSxDPUFV"
	LeafAuthorityKeyID = "4E:3B:73:F4:70:4D:C2:98:D:DB:C8:5A:5F:02:3B:BF:86:25:56:2B"
	LeafSubject        = "Tz1sZWFmLFNUPXNvbWUtc3RhdGUsQz1BVQ=="
	LeafSubjectAsText  = "O=leaf,ST=some-state,C=AU"
	LeafSubjectKeyID   = "30:F4:65:75:14:20:B2:AF:3D:14:71:17:AC:49:90:93:3E:24:A0:1F"
	LeafSerialNumber   = "143290473708569835418599774898811724528308722063"

	GoogleIssuer         = "Q049TWF0dGVyIFBBQSAxLE89R29vZ2xlLEM9VVMsMS4zLjYuMS40LjEuMzcyNDQuMi4xPSMxMzA0MzYzMDMwMzY="
	GoogleAuthorityKeyID = ""
	GoogleSubject        = "Q049TWF0dGVyIFBBQSAxLE89R29vZ2xlLEM9VVMsMS4zLjYuMS40LjEuMzcyNDQuMi4xPSMxMzA0MzYzMDMwMzY="
	GoogleSubjectAsText  = "CN=Matter PAA 1,O=Google,C=US,vid=0x6006"
	GoogleSubjectKeyID   = "B0:00:56:81:B8:88:62:89:62:80:E1:21:18:A1:A8:BE:09:DE:93:21"
	GoogleSerialNumber   = "1"

	TestIssuer         = "Q049TWF0dGVyIFRlc3QgUEFBLDEuMy42LjEuNC4xLjM3MjQ0LjIuMT0jMTMwNDMxMzIzNTQ0"
	TestAuthorityKeyID = "E2:90:8D:36:9C:3C:A3:C1:13:BB:09:E2:4D:C1:CC:C5:A6:66:91:D4"
	TestSubject        = "Q049TWF0dGVyIFRlc3QgUEFBLDEuMy42LjEuNC4xLjM3MjQ0LjIuMT0jMTMwNDMxMzIzNTQ0"
	TestSubjectAsText  = "CN=Matter Test PAA,vid=0x125D"
	TestSubjectKeyID   = "E2:90:8D:36:9C:3C:A3:C1:13:BB:09:E2:4D:C1:CC:C5:A6:66:91:D4"
	TestSerialNumber   = "1647312298631"
)
