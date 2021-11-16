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
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/crypto/secp256k1"
)

var (
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
	VID                                        uint16 = 1
	VendorName                                        = "Vendor Name"
	CompanyLegalName                                  = "Legal Company Name"
	CompanyPreferredName                              = "Company Preferred Name"
	VendorLandingPageURL                              = "https://www.example.com"
	PID                                        uint16 = 22
	DeviceTypeID                               uint16 = 12345
	Version                                           = "1.0"
	ProductName                                       = "Device Name"
	ProductLabel                                      = "Product Label and/or Product Description"
	PartNumber                                        = "RCU2205A"
	SoftwareVersion                            uint32 = 1
	SoftwareVersionString                             = "1.0"
	HardwareVersion                            uint32 = 21
	HardwareVersionString                             = "2.1"
	CDVersionNumber                            uint16 = 312
	FirmwareDigests                                   = "Firmware Digest String"
	Revoked                                           = false
	SoftwareVersionValid                              = true
	OtaURL                                            = "https://ota.firmware.com"
	OtaFileSize                                uint64 = 12345678
	OtaChecksum                                       = "2fd4e1c67a2d28fced849ee1bb76e7391b93eb12" //nolint:lll
	OtaChecksumType                            uint16 = 1
	OtaBlob                                           = "OTABlob Text"
	CommissioningCustomFlow                    uint8  = 1
	CommissioningCustomFlowURL                        = "https://sampleflowurl.dclmodel"
	CommissioningModeInitialStepsHint          uint32 = 2
	CommissioningModeInitialStepsInstruction          = "commissioningModeInitialStepsInstruction details"
	CommissioningModeSecondaryStepsHint        uint32 = 3
	CommissioningModeSecondaryStepsInstruction        = "commissioningModeSecondaryStepsInstruction steps"
	ReleaseNotesURL                                   = "https://url.releasenotes.dclmodel"
	UserManualURL                                     = "https://url.usermanual.dclmodel"
	SupportURL                                        = "https://url.supporturl.dclmodel"
	ProductURL                                        = "https://url.producturl.dclmodel"
	ChipBlob                                          = "Chip Blob Text"
	VendorBlob                                        = "Vendor Blob Text"
	MinApplicableSoftwareVersion               uint32 = 1
	MaxApplicableSoftwareVersion               uint32 = 1000
	Owner                                             = Address1

	// Compliance.
	CertificationDate = time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	RevocationDate    = time.Date(2020, 3, 3, 3, 30, 0, 0, time.UTC)
	Reason            = "Some Reason"
	RevocationReason  = "Some Reason"
	CertificationType = "zigbee"

	// Testing Result.
	TestResult = "http://test.result.com"
	TestDate   = time.Date(2020, 2, 2, 2, 0, 0, 0, time.UTC)

	//
	Address1, _              = sdk.AccAddressFromBech32("cosmos1p72j8mgkf39qjzcmr283w8l8y9qv30qpj056uz")
	Address2, _              = sdk.AccAddressFromBech32("cosmos1j8x9urmqs7p44va5p4cu29z6fc3g0cx2c2vxx2")
	Address3, _              = sdk.AccAddressFromBech32("cosmos1j7tc5f4f54fd8hns42nsavzhadr0gchddz6vfl")
	VendorID1         uint16 = 1000
	VendorID2         uint16 = 2000
	VendorID3         uint16 = 3000
	Pubkey1Str               = "cosmospub1addwnpepq28rlfval9n8khmgqz55mlfwn4rlh0jk80k9n7fvtu4g4u37qtvry76ww9h"
	PubKey1, _               = sdk.GetAccPubKeyBech32(Pubkey1Str)
	PubKey2Str               = "cosmospub1addwnpepq086aynq08ey3nyhdvd3nma5fqyh00yuqtwzz06g6juqaqclcpqvcft9yng"
	PubKey2, _               = sdk.GetAccPubKeyBech32(PubKey2Str)
	PubKey3Str               = "cosmospub1addwnpepqwsq3gh4k5xat4n6s0e3murz4xgmwu9jv9wl0zwhp709f2eyn5ljv8z60zn"
	PubKey3, _               = sdk.GetAccPubKeyBech32(PubKey3Str)
	Signer                   = Address1
	ValidatorPubKey1         = "cosmosvalconspub1zcjduepqdmmjdfyvh2mrwl8p8wkwp23kh8lvjrd9u45snxqz6te6y6lwk6gqts45r3"
	ValidatorPubKey2         = "cosmosvalconspub1zcjduepqdtar5ynhrhc78mymwg5sqksdnfafqyqu6sar3gg745u6dsw32krscaqv8u"
	ValidatorAddress1        = sdk.ConsAddress(sdk.MustGetConsPubKeyBech32(ValidatorPubKey1).Address())
	ValidatorAddress2        = sdk.ConsAddress(sdk.MustGetConsPubKeyBech32(ValidatorPubKey2).Address())
	ValidHTTPSURL            = "https://valid.url.com"
	ValidHTTPURL             = "http://valid.url.com"
	NotAValidURL             = "not a valid url"
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

func TestAddress() (sdk.AccAddress, crypto.PubKey, string) {
	key := secp256k1.GenPrivKey()
	pub := key.PubKey()
	addr := sdk.AccAddress(pub.Address())
	pubStr := sdk.MustBech32ifyAccPub(pub)

	return addr, pub, pubStr
}
