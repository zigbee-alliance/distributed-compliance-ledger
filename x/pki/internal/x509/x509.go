package x509

// nolint:goimports
import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"strings"

	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/pki/internal/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type X509Certificate struct {
	Issuer         string
	SerialNumber   string
	Subject        string
	SubjectKeyID   string
	AuthorityKeyID string
	Certificate    *x509.Certificate
}

func DecodeX509Certificate(pemCertificate string) (*X509Certificate, sdk.Error) {
	block, _ := pem.Decode([]byte(pemCertificate))
	if block == nil {
		return nil, types.ErrCodeInvalidCertificate("Could not decode pem certificate")
	}

	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, types.ErrCodeInvalidCertificate(fmt.Sprintf("Could not parse certificate: %v", err.Error()))
	}

	certificate := X509Certificate{
		Issuer:         cert.Issuer.String(),
		SerialNumber:   cert.SerialNumber.String(),
		Subject:        cert.Subject.String(),
		SubjectKeyID:   BytesToHex(cert.SubjectKeyId),
		AuthorityKeyID: BytesToHex(cert.AuthorityKeyId),
		Certificate:    cert,
	}

	return &certificate, nil
}

func BytesToHex(bytes []byte) string {
	if bytes == nil {
		return ""
	}

	bytesHex := make([]string, len(bytes))
	for i, byte_ := range bytes {
		bytesHex[i] = fmt.Sprintf("%X", byte_)
	}

	return strings.Join(bytesHex, ":")
}

func (c X509Certificate) Verify(parent *X509Certificate) sdk.Error {
	roots := x509.NewCertPool()
	roots.AddCert(parent.Certificate)

	opts := x509.VerifyOptions{Roots: roots}

	if _, err := c.Certificate.Verify(opts); err != nil {
		return types.ErrCodeInvalidCertificate(fmt.Sprintf("Certificate verification failed. Error: %v", err))
	}

	return nil
}

func (c X509Certificate) IsSelfSigned() bool {
	if len(c.AuthorityKeyID) > 0 {
		return c.Issuer == c.Subject && c.AuthorityKeyID == c.SubjectKeyID
	} else {
		return c.Issuer == c.Subject
	}
}
