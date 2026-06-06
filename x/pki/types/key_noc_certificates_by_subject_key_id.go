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

package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// NocCertificatesBySubjectKeyIDKeyPrefix is the prefix to retrieve all NocCertificatesBySubjectKeyID
	NocCertificatesBySubjectKeyIDKeyPrefix = "NocCertificatesBySubjectKeyID/value/"
)

// NocCertificatesBySubjectKeyIDKey returns the store key to retrieve a NocCertificatesBySubjectKeyID from the index fields
func NocCertificatesBySubjectKeyIDKey(
	subjectKeyID string,
) []byte {
	var key []byte

	subjectKeyIDBytes := []byte(subjectKeyID)
	key = append(key, subjectKeyIDBytes...)
	key = append(key, []byte("/")...)

	return key
}
