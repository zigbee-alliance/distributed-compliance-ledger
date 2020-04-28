package test_constants

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"time"
)

var (
	ChainId                        = "zblchain"
	AccountName                    = "jack"
	Passphrase                     = "test1234"
	VID                      int16 = 1
	PID                      int16 = 22
	CID                      int16 = 12345
	Name                           = "Device Name"
	Owner                          = Address1
	Description                    = "Device Description"
	Sku                            = "RCU2205A"
	FirmwareVersion                = "1.0"
	HardwareVersion                = "2.0"
	Custom                         = "Custom data"
	CertificateID                  = "ZIG12345678"
	CertificationDate              = time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	RevocationDate                 = time.Date(2020, 3, 3, 3, 30, 0, 0, time.UTC)
	Reason                         = "Some Reason"
	RevocationReason               = "Some Reason"
	TisOrTrpTestingCompleted       = false
	TestResult                     = "http://test.result.com"
	TestDate                       = time.Date(2020, 2, 2, 2, 0, 0, 0, time.UTC)
	Address1, _                    = sdk.AccAddressFromBech32("cosmos1p72j8mgkf39qjzcmr283w8l8y9qv30qpj056uz")
	Address2, _                    = sdk.AccAddressFromBech32("cosmos1j8x9urmqs7p44va5p4cu29z6fc3g0cx2c2vxx2")
	Address3, _                    = sdk.AccAddressFromBech32("cosmos1j7tc5f4f54fd8hns42nsavzhadr0gchddz6vfl")
	PubKey                         = "cosmospub1addwnpepqtrnrp93hswlsrzvltc3n8z7hjg9dxuh3n4rkp2w2verwfr8yg27c95l4k3"
	Signer                         = Address1
	CertificationType              = "zb"
	EmptyString                    = ""
)

/*
	Certificates are taken from dsr-corporation.com
*/

const (
	StubCert = `pem certificate`

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
MIIGCDCCBPCgAwIBAgISA5VDSEqrx1nKdEk5EZSCCZ7eMA0GCSqGSIb3DQEBCwUA
MEoxCzAJBgNVBAYTAlVTMRYwFAYDVQQKEw1MZXQncyBFbmNyeXB0MSMwIQYDVQQD
ExpMZXQncyBFbmNyeXB0IEF1dGhvcml0eSBYMzAeFw0yMDA0MTAxMDA0NTFaFw0y
MDA3MDkxMDA0NTFaMB4xHDAaBgNVBAMTE2Rzci1jb3Jwb3JhdGlvbi5jb20wggEi
MA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQChMNeaBTYkFXYWqFODO4q9+lQP
5FrdfdC6xDNOP2TUqRl70r8EIpElH6AwnNejej/2JIQq0Lu1y+KML0esHOM66DP+
YVHLq3aZYL2JorL9vufpPUzgzdl+bhYFu6w0pbg8Y7ua/OYFIUK9exbR2ZCIkG/x
xr44MD2kg0DDPbjQso5CcRubOyOlJQcb++64Duz+gKJt7g3nPcx++LuGR7WVoVoZ
3rSbXiXbLAcrZkqMt0Y7QCifFJjq7l2EMcYpxYolS6/QFmcKkgMiP0koMb0tAD2p
MqeY06DRw5n62NeQKz0bj9Ko0LemAB3l89pxMvaDKf0p0IniYY4sXdiw7iutAgMB
AAGjggMSMIIDDjAOBgNVHQ8BAf8EBAMCBaAwHQYDVR0lBBYwFAYIKwYBBQUHAwEG
CCsGAQUFBwMCMAwGA1UdEwEB/wQCMAAwHQYDVR0OBBYEFIo0C1zYQhjywSqserOP
bpBm9E5cMB8GA1UdIwQYMBaAFKhKamMEfd265tE5t6ZFZe/zqOyhMG8GCCsGAQUF
BwEBBGMwYTAuBggrBgEFBQcwAYYiaHR0cDovL29jc3AuaW50LXgzLmxldHNlbmNy
eXB0Lm9yZzAvBggrBgEFBQcwAoYjaHR0cDovL2NlcnQuaW50LXgzLmxldHNlbmNy
eXB0Lm9yZy8wgccGA1UdEQSBvzCBvIIZKi5kZW4uZHNyLWNvcnBvcmF0aW9uLmNv
bYIZKi5kZXYuZHNyLWNvcnBvcmF0aW9uLmNvbYIRKi5kc3ItY29tcGFueS5jb22C
FSouZHNyLWNvcnBvcmF0aW9uLmNvbYIZKi5vcG8uZHNyLWNvcnBvcmF0aW9uLmNv
bYIZKi52b3ouZHNyLWNvcnBvcmF0aW9uLmNvbYIPZHNyLWNvbXBhbnkuY29tghNk
c3ItY29ycG9yYXRpb24uY29tMEwGA1UdIARFMEMwCAYGZ4EMAQIBMDcGCysGAQQB
gt8TAQEBMCgwJgYIKwYBBQUHAgEWGmh0dHA6Ly9jcHMubGV0c2VuY3J5cHQub3Jn
MIIBBAYKKwYBBAHWeQIEAgSB9QSB8gDwAHYA8JWkWfIA0YJAEC0vk4iOrUv+HUfj
meHQNKawqKqOsnMAAAFxY8OJqgAABAMARzBFAiEA3+XxvlbpW4iOsH+uJ9cLUNkI
7/NYb9Dj/+I/uxn8wGECIETeVnfiECiFnL2LNZMvE01qK+ORFup4yAra0pQGKUWp
AHYAB7dcG+V9aP/xsMYdIxXHuuZXfFeUt2ruvGE6GmnTohwAAAFxY8OJ3QAABAMA
RzBFAiAspZ8d8yAOXvKfm7x6cW/w42J8ixDVpx8u4Mae9QqE0QIhAIDgzDTIHwx3
sE/+Va6XXeez5kciFMeFHykTOv4tQsN7MA0GCSqGSIb3DQEBCwUAA4IBAQAtvQNQ
SvyOChOmP7Fsa8LeuGRyth2XQFsniwdQSD0oEkq8RF7h6GMmTapUHQiihYZbjsjZ
Cs3NW4/4Yfs/2do7JUD94mFpFr8IlpAkTUtfWhHFSi2hJbIwRx9bv8bMikyqcZxV
SEQS1fu1qgX4hbh5QUdtdBD/fQNIiFwgzgik9EURk/mQXWzt7d6OLLptTp0OclfS
psS5wa0aHSO2RsRzpOG8+t30EJm5fHL33VVrNJKLSNqPQy7VNDC/E71GO0RqG0HF
jjnAZrhowZT5bTsGAIMjIw90Pqz/77Xjam50VkMo7R/Ik8D25/Wy0TboZPK75NaX
nKzrbFlTU4d9Cmib
-----END CERTIFICATE-----`

	RootIssuer       = "CN=DST Root CA X3,O=Digital Signature Trust Co."
	RootSubject      = "CN=DST Root CA X3,O=Digital Signature Trust Co."
	RootSubjectKeyId = "C4:A7:B1:A4:7B:2C:71:FA:DB:E1:4B:90:75:FF:C4:15:60:85:89:10"
	RootSerialNumber = "91299735575339953335919266965803778155"

	IntermediateIssuer       = "CN=DST Root CA X3,O=Digital Signature Trust Co."
	IntermediateSubject      = "CN=Let's Encrypt Authority X3,O=Let's Encrypt,C=US"
	IntermediateSubjectKeyId = "A8:4A:6A:63:4:7D:DD:BA:E6:D1:39:B7:A6:45:65:EF:F3:A8:EC:A1"
	IntermediateSerialNumber = "13298795840390663119752826058995181320"

	LeafIssuer       = "CN=Let's Encrypt Authority X3,O=Let's Encrypt,C=US"
	LeafSubject      = "CN=dsr-corporation.com"
	LeafSubjectKeyId = "8A:34:B:5C:D8:42:18:F2:C1:2A:AC:7A:B3:8F:6E:90:66:F4:4E:5C"
	LeafSerialNumber = "312128364102099997394566658874957944692446"
)
