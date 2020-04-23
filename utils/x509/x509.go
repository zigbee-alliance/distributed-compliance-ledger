package x509

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"strings"
)

type X509Certificate struct {
	Issuer         string
	SerialNumber   string
	Subject        string
	SubjectKeyId   string
	AuthorityKeyId string
	Certificate    *x509.Certificate
}

func DecodeX509Certificate(pemCertificate string) (*X509Certificate, sdk.Error) {
	block, _ := pem.Decode([]byte(pemCertificate))
	if block == nil {
		return nil, sdk.ErrInternal("Could not decode pemCertificate")
	}

	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, sdk.ErrInternal(fmt.Sprintf("Could not parse pemCertificate" + err.Error()))
	}

	certificate := X509Certificate{
		Issuer:         cert.Issuer.String(),
		SerialNumber:   cert.SerialNumber.String(),
		Subject:        cert.Subject.String(),
		SubjectKeyId:   BytesToHex(cert.SubjectKeyId),
		AuthorityKeyId: BytesToHex(cert.AuthorityKeyId),
		Certificate:    cert,
	}

	return &certificate, nil
}

func BytesToHex(bytes []byte) string {
	if bytes == nil {
		return ""
	}

	var bytesHex []string
	for _, byte_ := range bytes {
		bytesHex = append(bytesHex, fmt.Sprintf("%X", byte_))
	}

	return strings.Join(bytesHex, ":")
}

func (c X509Certificate) VerifyX509Certificate(parent *x509.Certificate) sdk.Error {
	roots := x509.NewCertPool()
	roots.AddCert(parent)

	opts := x509.VerifyOptions{Roots: roots}

	if _, err := c.Certificate.Verify(opts); err != nil {
		return sdk.ErrInternal(fmt.Sprintf("Certificate verification failed. Error: %v", err))
	}

	return nil
}

func (c X509Certificate) IsRootCertificate() bool {
	return c.Subject == c.Issuer || (len(c.AuthorityKeyId) > 0 && c.AuthorityKeyId == c.SubjectKeyId)
}
