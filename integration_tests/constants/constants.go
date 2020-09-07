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
	VID                      uint16 = 1
	PID                      uint16 = 22
	CID                      uint16 = 12345
	Name                            = "Device Name"
	Owner                           = Address1
	Description                     = "Device Description"
	Sku                             = "RCU2205A"
	FirmwareVersion                 = "1.0"
	HardwareVersion                 = "2.0"
	Custom                          = "Custom data"
	TisOrTrpTestingCompleted        = false

	// Compliance.
	CertificationDate = time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	RevocationDate    = time.Date(2020, 3, 3, 3, 30, 0, 0, time.UTC)
	Reason            = "Some Reason"
	RevocationReason  = "Some Reason"
	CertificationType = "zb"

	// Testing Result.
	TestResult = "http://test.result.com"
	TestDate   = time.Date(2020, 2, 2, 2, 0, 0, 0, time.UTC)

	//
	Address1, _       = sdk.AccAddressFromBech32("cosmos1p72j8mgkf39qjzcmr283w8l8y9qv30qpj056uz")
	Address2, _       = sdk.AccAddressFromBech32("cosmos1j8x9urmqs7p44va5p4cu29z6fc3g0cx2c2vxx2")
	Address3, _       = sdk.AccAddressFromBech32("cosmos1j7tc5f4f54fd8hns42nsavzhadr0gchddz6vfl")
	Pubkey1Str        = "cosmospub1addwnpepq28rlfval9n8khmgqz55mlfwn4rlh0jk80k9n7fvtu4g4u37qtvry76ww9h"
	PubKey1, _        = sdk.GetAccPubKeyBech32(Pubkey1Str)
	PubKey2Str        = "cosmospub1addwnpepq086aynq08ey3nyhdvd3nma5fqyh00yuqtwzz06g6juqaqclcpqvcft9yng"
	PubKey2, _        = sdk.GetAccPubKeyBech32(PubKey2Str)
	PubKey3Str        = "cosmospub1addwnpepqwsq3gh4k5xat4n6s0e3murz4xgmwu9jv9wl0zwhp709f2eyn5ljv8z60zn"
	PubKey3, _        = sdk.GetAccPubKeyBech32(PubKey3Str)
	Signer            = Address1
	ValidatorPubKey1  = "cosmosvalconspub1zcjduepqdmmjdfyvh2mrwl8p8wkwp23kh8lvjrd9u45snxqz6te6y6lwk6gqts45r3"
	ValidatorPubKey2  = "cosmosvalconspub1zcjduepqdtar5ynhrhc78mymwg5sqksdnfafqyqu6sar3gg745u6dsw32krscaqv8u"
	ValidatorAddress1 = sdk.ConsAddress(sdk.MustGetConsPubKeyBech32(ValidatorPubKey1).Address())
	ValidatorAddress2 = sdk.ConsAddress(sdk.MustGetConsPubKeyBech32(ValidatorPubKey2).Address())
)

/*
	Certificates are taken from dsr-corporation.com
*/

const (
	StubCertPem = `pem certificate`

	RootCertPem = `
-----BEGIN CERTIFICATE-----
MIIDSjCCAjKgAwIBAgIQRK+wgNajJ7qJMDmGLvhAazANBgkqhkiG9w0BAQUFADA/
MSQwIgYDVQQKExtEaWdpdGFsIFNpZ25hdHVyZSBUcnVzdCBDby4xFzAVBgNVBAMT
DkRTVCBSb290IENBIFgzMB4XDTAwMDkzMDIxMTIxOVoXDTIxMDkzMDE0MDExNVow
PzEkMCIGA1UEChMbRGlnaXRhbCBTaWduYXR1cmUgVHJ1c3QgQ28uMRcwFQYDVQQD
Ew5EU1QgUm9vdCBDQSBYMzCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEB
AN+v6ZdQCINXtMxiZfaQguzH0yxrMMpb7NnDfcdAwRgUi+DoM3ZJKuM/IUmTrE4O
rz5Iy2Xu/NMhD2XSKtkyj4zl93ewEnu1lcCJo6m67XMuegwGMoOifooUMM0RoOEq
OLl5CjH9UL2AZd+3UWODyOKIYepLYYHsUmu5ouJLGiifSKOeDNoJjj4XLh7dIN9b
xiqKqy69cK3FCxolkHRyxXtqqzTWMIn/5WgTe1QLyNau7Fqckh49ZLOMxt+/yUFw
7BZy1SbsOFU5Q9D8/RhcQPGX69Wam40dutolucbY38EVAjqr2m7xPi71XAicPNaD
aeQQmxkqtilX4+U9m5/wAl0CAwEAAaNCMEAwDwYDVR0TAQH/BAUwAwEB/zAOBgNV
HQ8BAf8EBAMCAQYwHQYDVR0OBBYEFMSnsaR7LHH62+FLkHX/xBVghYkQMA0GCSqG
SIb3DQEBBQUAA4IBAQCjGiybFwBcqR7uKGY3Or+Dxz9LwwmglSBd49lZRNI+DT69
ikugdB/OEIKcdBodfpga3csTS7MgROSR6cz8faXbauX+5v3gTt23ADq1cEmv8uXr
AvHRAosZy5Q6XkjEGB5YGV8eAlrwDPGxrancWYaLbumR9YbK+rlmM6pZW87ipxZz
R8srzJmwN0jP41ZL9c8PDHIyh8bwRLtTcm1D9SZImlJnt1ir/md2cXjbDaJWFBM5
JDGFoqgCWjBH4d1QB7wCCZAA62RjYJsWvIjJEubSfZGL+T0yjWW06XyxV3bqxbYo
Ob8VZRzI9neWagqNdwvYkQsEjgfbKbYK7p2CNTUQ
-----END CERTIFICATE-----`

	IntermediateCertPem = `
-----BEGIN CERTIFICATE-----
MIIEkjCCA3qgAwIBAgIQCgFBQgAAAVOFc2oLheynCDANBgkqhkiG9w0BAQsFADA/
MSQwIgYDVQQKExtEaWdpdGFsIFNpZ25hdHVyZSBUcnVzdCBDby4xFzAVBgNVBAMT
DkRTVCBSb290IENBIFgzMB4XDTE2MDMxNzE2NDA0NloXDTIxMDMxNzE2NDA0Nlow
SjELMAkGA1UEBhMCVVMxFjAUBgNVBAoTDUxldCdzIEVuY3J5cHQxIzAhBgNVBAMT
GkxldCdzIEVuY3J5cHQgQXV0aG9yaXR5IFgzMIIBIjANBgkqhkiG9w0BAQEFAAOC
AQ8AMIIBCgKCAQEAnNMM8FrlLke3cl03g7NoYzDq1zUmGSXhvb418XCSL7e4S0EF
q6meNQhY7LEqxGiHC6PjdeTm86dicbp5gWAf15Gan/PQeGdxyGkOlZHP/uaZ6WA8
SMx+yk13EiSdRxta67nsHjcAHJyse6cF6s5K671B5TaYucv9bTyWaN8jKkKQDIZ0
Z8h/pZq4UmEUEz9l6YKHy9v6Dlb2honzhT+Xhq+w3Brvaw2VFn3EK6BlspkENnWA
a6xK8xuQSXgvopZPKiAlKQTGdMDQMc2PMTiVFrqoM7hD8bEfwzB/onkxEz0tNvjj
/PIzark5McWvxI0NHWQWM6r6hCm21AvA2H3DkwIDAQABo4IBfTCCAXkwEgYDVR0T
AQH/BAgwBgEB/wIBADAOBgNVHQ8BAf8EBAMCAYYwfwYIKwYBBQUHAQEEczBxMDIG
CCsGAQUFBzABhiZodHRwOi8vaXNyZy50cnVzdGlkLm9jc3AuaWRlbnRydXN0LmNv
bTA7BggrBgEFBQcwAoYvaHR0cDovL2FwcHMuaWRlbnRydXN0LmNvbS9yb290cy9k
c3Ryb290Y2F4My5wN2MwHwYDVR0jBBgwFoAUxKexpHsscfrb4UuQdf/EFWCFiRAw
VAYDVR0gBE0wSzAIBgZngQwBAgEwPwYLKwYBBAGC3xMBAQEwMDAuBggrBgEFBQcC
ARYiaHR0cDovL2Nwcy5yb290LXgxLmxldHNlbmNyeXB0Lm9yZzA8BgNVHR8ENTAz
MDGgL6AthitodHRwOi8vY3JsLmlkZW50cnVzdC5jb20vRFNUUk9PVENBWDNDUkwu
Y3JsMB0GA1UdDgQWBBSoSmpjBH3duubRObemRWXv86jsoTANBgkqhkiG9w0BAQsF
AAOCAQEA3TPXEfNjWDjdGBX7CVW+dla5cEilaUcne8IkCJLxWh9KEik3JHRRHGJo
uM2VcGfl96S8TihRzZvoroed6ti6WqEBmtzw3Wodatg+VyOeph4EYpr/1wXKtx8/
wApIvJSwtmVi4MFU5aMqrSDE6ea73Mj2tcMyo5jMd6jmeWUHK8so/joWUoHOUgwu
X4Po1QYz+3dszkDqMp4fklxBwXRsW10KXzPMTZ+sOPAveyxindmjkW8lGy+QsRlG
PfZ+G6Z6h7mjem0Y+iWlkYcV4PIWL1iwBi8saCbGS5jN2p8M+X+Q7UNKEkROb3N6
KOqkqm57TH2H3eDJAkSnh6/DNFu0Qg==
-----END CERTIFICATE-----`

	LeafCertPem = `
-----BEGIN CERTIFICATE-----
MIIGBjCCBO6gAwIBAgISBIWVFpvXXfNwjp+/5C8v//29MA0GCSqGSIb3DQEBCwUA
MEoxCzAJBgNVBAYTAlVTMRYwFAYDVQQKEw1MZXQncyBFbmNyeXB0MSMwIQYDVQQD
ExpMZXQncyBFbmNyeXB0IEF1dGhvcml0eSBYMzAeFw0yMDA2MTIxMDA2NDNaFw0y
MDA5MTAxMDA2NDNaMB4xHDAaBgNVBAMTE2Rzci1jb3Jwb3JhdGlvbi5jb20wggEi
MA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQDsdmARA5TRQNk4uavZA43ILVPg
TubMuMy6lXy6ur5yOdN/sFArVmpKhmQGTeeCxtSQx2qdVtlukBq2tUuwP0V0KuGz
HMAB8muJ+rH9r8UCssU5Sa+zVy6gl0JMyKdwut1S6vkc58OooHoTjCUfhTEGq6WP
VDQvpdfU1BjHSuu2mQgox2+tZLK3RkjiPa72apivCZ5b1qYsfedEhO3XJUgP229l
zE0ivgwUbiieVO3X9QMKL2iTyfhz2eVR1dFCiRyFnyVmJUh2cAb1pzEm3BOXBTKe
g5cBKocdxJaQiIiMQiWn3ejdOQqorC/L9cjEFjlB6ytimWRe2WQIxANTWkXVAgMB
AAGjggMQMIIDDDAOBgNVHQ8BAf8EBAMCBaAwHQYDVR0lBBYwFAYIKwYBBQUHAwEG
CCsGAQUFBwMCMAwGA1UdEwEB/wQCMAAwHQYDVR0OBBYEFIrprNQWgS+HZo5hvqnF
HAAb97uuMB8GA1UdIwQYMBaAFKhKamMEfd265tE5t6ZFZe/zqOyhMG8GCCsGAQUF
BwEBBGMwYTAuBggrBgEFBQcwAYYiaHR0cDovL29jc3AuaW50LXgzLmxldHNlbmNy
eXB0Lm9yZzAvBggrBgEFBQcwAoYjaHR0cDovL2NlcnQuaW50LXgzLmxldHNlbmNy
eXB0Lm9yZy8wgccGA1UdEQSBvzCBvIIZKi5kZW4uZHNyLWNvcnBvcmF0aW9uLmNv
bYIZKi5kZXYuZHNyLWNvcnBvcmF0aW9uLmNvbYIRKi5kc3ItY29tcGFueS5jb22C
FSouZHNyLWNvcnBvcmF0aW9uLmNvbYIZKi5vcG8uZHNyLWNvcnBvcmF0aW9uLmNv
bYIZKi52b3ouZHNyLWNvcnBvcmF0aW9uLmNvbYIPZHNyLWNvbXBhbnkuY29tghNk
c3ItY29ycG9yYXRpb24uY29tMEwGA1UdIARFMEMwCAYGZ4EMAQIBMDcGCysGAQQB
gt8TAQEBMCgwJgYIKwYBBQUHAgEWGmh0dHA6Ly9jcHMubGV0c2VuY3J5cHQub3Jn
MIIBAgYKKwYBBAHWeQIEAgSB8wSB8ADuAHUA8JWkWfIA0YJAEC0vk4iOrUv+HUfj
meHQNKawqKqOsnMAAAFyqDXiowAABAMARjBEAiAXMUP/VFsOZGT1ej2I0/BBnZmh
zsAthlar5TMMivGJsAIgWCGEH4cYUyuV2KyYiOaaCw3kxAsoBfT9lNqCVBSQzj0A
dQCyHgXMi6LNiiBOh2b5K7mKJSBna9r6cOeySVMt74uQXgAAAXKoNeKUAAAEAwBG
MEQCIEJHPLV0jdjc+wZYCRTrf5GCwG6VNWOIz6+Y2mkFYVK5AiBhhXIpr5eThbkj
izndaDzH138IRrL+X2njajXlupmWITANBgkqhkiG9w0BAQsFAAOCAQEAcb53VyNY
jp98EthlIRYSQ6n3iyZEj4txRDzPUDeyW6/2WRl5sqlL2N7iFaAhYHeiHJa+vtuk
IcyZPNl0/xHlLugtMtSuZe5bQLh4c8F/DK6Jvl3FJjoCNJi/Vaa15HSZOQclnZsK
FQ9Xh3mhjLGVIP562CcHNvv11fzTqOR2vXvXmpRX7gbU2AjoauO5wNu3+YF5gvnx
lY1XYUJe8+ByFwIiUGA6MIYKKAvbld9vmIaU/zmiBQ974aDREkl3xtXIJ3ZDh+lw
5jocGegjiQFa26Q6d77mr7Sv4O4grJvZEjY50zx7kjZl5ABScxiuKsq+T1OTqIr6
eTXgtecouKOE6g==
-----END CERTIFICATE-----`

	RootIssuer       = "CN=DST Root CA X3,O=Digital Signature Trust Co."
	RootSubject      = "CN=DST Root CA X3,O=Digital Signature Trust Co."
	RootSubjectKeyID = "C4:A7:B1:A4:7B:2C:71:FA:DB:E1:4B:90:75:FF:C4:15:60:85:89:10"
	RootSerialNumber = "91299735575339953335919266965803778155"

	IntermediateIssuer         = "CN=DST Root CA X3,O=Digital Signature Trust Co."
	IntermediateAuthorityKeyID = "C4:A7:B1:A4:7B:2C:71:FA:DB:E1:4B:90:75:FF:C4:15:60:85:89:10"
	IntermediateSubject        = "CN=Let's Encrypt Authority X3,O=Let's Encrypt,C=US"
	IntermediateSubjectKeyID   = "A8:4A:6A:63:4:7D:DD:BA:E6:D1:39:B7:A6:45:65:EF:F3:A8:EC:A1"
	IntermediateSerialNumber   = "13298795840390663119752826058995181320"

	LeafIssuer         = "CN=Let's Encrypt Authority X3,O=Let's Encrypt,C=US"
	LeafAuthorityKeyID = "A8:4A:6A:63:4:7D:DD:BA:E6:D1:39:B7:A6:45:65:EF:F3:A8:EC:A1"
	LeafSubject        = "CN=dsr-corporation.com"
	LeafSubjectKeyID   = "8A:E9:AC:D4:16:81:2F:87:66:8E:61:BE:A9:C5:1C:0:1B:F7:BB:AE"
	LeafSerialNumber   = "393904870890265262371394210372104514174397"
)

func TestAddress() (sdk.AccAddress, crypto.PubKey, string) {
	key := secp256k1.GenPrivKey()
	pub := key.PubKey()
	addr := sdk.AccAddress(pub.Address())
	pubStr := sdk.MustBech32ifyAccPub(pub)

	return addr, pub, pubStr
}
