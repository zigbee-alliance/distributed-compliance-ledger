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

package x509

import (
	"bytes"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"strings"

	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

type X509Certificate struct {
	Issuer         string
	SerialNumber   string
	Subject        string
	SubjectKeyID   string
	AuthorityKeyID string
	Certificate    *x509.Certificate
}

func DecodeX509Certificate(pemCertificate string) (*X509Certificate, error) {
	block, _ := pem.Decode([]byte(pemCertificate))
	if block == nil {
		return nil, types.NewErrInvalidCertificate("Could not decode pem certificate")
	}

	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, types.NewErrInvalidCertificate(fmt.Sprintf("Could not parse certificate: %v", err.Error()))
	}

	certificate := X509Certificate{
		Issuer:         TranslateToVidOrPid([]byte(cert.Issuer.String())),
		SerialNumber:   cert.SerialNumber.String(),
		Subject:        TranslateToVidOrPid([]byte(cert.Subject.String())),
		SubjectKeyID:   BytesToHex(cert.SubjectKeyId),
		AuthorityKeyID: BytesToHex(cert.AuthorityKeyId),
		Certificate:    cert,
	}

	return &certificate, nil
}

func TranslateToVidOrPid(subject []byte) string {
	oldVID := []byte("1.3.6.1.4.1.37244.2.1")
	oldPID := []byte("1.3.6.1.4.1.37244.2.2")

	var vid, vidValue, pid, pidValue []byte

	if i := bytes.Index(subject, oldVID); i >= 0 {
		// get vid value from subject
		vidValue = subject[i+len(oldVID)+2 : len(string(subject))]
		// get last 8 numbers from vidValue
		vidValue = vidValue[len(vidValue)-8:]

		decoded, _ := hex.DecodeString(string(vidValue))

		vid = []byte(("vid=0x") + string(decoded))

		subject = subject[0:i]
		subject = []byte(string(subject) + string(vid))
	} else if i := bytes.Index(subject, oldPID); i >= 0 {
		// get pid value from subject
		pidValue = subject[i+len(oldPID)+2 : len(string(subject))]
		// get last 8 number from pidValue
		pidValue = pidValue[len(pidValue)-8:]

		decoded, _ := hex.DecodeString(string(pidValue))

		pid = []byte(("pid=0x") + string(decoded))

		subject = subject[0:i]
		subject = []byte(string(subject) + string(pid))
	}

	return string(subject)
}

func BytesToHex(bytes []byte) string {
	if bytes == nil {
		return ""
	}

	bytesHex := make([]string, len(bytes))
	for i, b := range bytes {
		bytesHex[i] = fmt.Sprintf("%X", b)
	}

	return strings.Join(bytesHex, ":")
}

func (c X509Certificate) Verify(parent *X509Certificate) error {
	roots := x509.NewCertPool()
	roots.AddCert(parent.Certificate)

	opts := x509.VerifyOptions{Roots: roots}

	if _, err := c.Certificate.Verify(opts); err != nil {
		return types.NewErrInvalidCertificate(fmt.Sprintf("Certificate verification failed. Error: %v", err))
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
