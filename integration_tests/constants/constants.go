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
	VendorLandingPageUrl                              = "https://www.example.com"
	Pid                                        int32  = 22
	DeviceTypeId                               int32  = 12345
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
	OtaUrl                                            = "https://ota.firmware.com"
	OtaFileSize                                uint64 = 12345678
	OtaChecksum                                       = "2fd4e1c67a2d28fced849ee1bb76e7391b93eb12" //nolint:lll
	OtaChecksumType                            int32  = 1
	OtaBlob                                           = "OTABlob Text"
	CommissioningCustomFlow                    int32  = 1
	CommissioningCustomFlowUrl                        = "https://sampleflowurl.dclmodel"
	CommissioningModeInitialStepsHint          uint32 = 2
	CommissioningModeInitialStepsInstruction          = "commissioningModeInitialStepsInstruction details"
	CommissioningModeSecondaryStepsHint        uint32 = 3
	CommissioningModeSecondaryStepsInstruction        = "commissioningModeSecondaryStepsInstruction steps"
	ReleaseNotesUrl                                   = "https://url.releasenotes.dclmodel"
	UserManualUrl                                     = "https://url.usermanual.dclmodel"
	SupportUrl                                        = "https://url.supporturl.dclmodel"
	ProductUrl                                        = "https://url.producturl.dclmodel"
	LsfUrl                                            = "https://url.lsfurl.dclmodel"
	LsfRevision                                int32  = 1
	ChipBlob                                          = "Chip Blob Text"
	VendorBlob                                        = "Vendor Blob Text"
	MinApplicableSoftwareVersion               uint32 = 1
	MaxApplicableSoftwareVersion               uint32 = 1000
	Owner                                             = Address1

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

	RootIssuer       = "O=root-ca,ST=some-state,C=AU"
	RootSubject      = "O=root-ca,ST=some-state,C=AU"
	RootSubjectKeyID = "5A:88:E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:9:30:E6:2B:DB"
	RootSerialNumber = "442314047376310867378175982234956458728610743315"

	IntermediateIssuer         = "O=root-ca,ST=some-state,C=AU"
	IntermediateAuthorityKeyID = "5A:88:E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:9:30:E6:2B:DB"
	IntermediateSubject        = "O=intermediate-ca,ST=some-state,C=AU"
	IntermediateSubjectKeyID   = "4E:3B:73:F4:70:4D:C2:98:D:DB:C8:5A:5F:2:3B:BF:86:25:56:2B"
	IntermediateSerialNumber   = "169917617234879872371588777545667947720450185023"

	LeafIssuer         = "O=intermediate-ca,ST=some-state,C=AU"
	LeafAuthorityKeyID = "4E:3B:73:F4:70:4D:C2:98:D:DB:C8:5A:5F:2:3B:BF:86:25:56:2B"
	LeafSubject        = "O=leaf,ST=some-state,C=AU"
	LeafSubjectKeyID   = "30:F4:65:75:14:20:B2:AF:3D:14:71:17:AC:49:90:93:3E:24:A0:1F"
	LeafSerialNumber   = "143290473708569835418599774898811724528308722063"
)
