// Copyright 2022 DSR Corporation
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
	// ApprovedCertificatesBySubjectKeyPrefix is the prefix to retrieve all ApprovedCertificatesBySubject.
	ApprovedCertificatesBySubjectKeyPrefix = "ApprovedCertificatesBySubject/value/"
)

// ApprovedCertificatesBySubjectKey returns the store key to retrieve a ApprovedCertificatesBySubject from the index fields.
func ApprovedCertificatesBySubjectKey(
	subject string,
) []byte {
	var key []byte

	subjectBytes := []byte(subject)
	key = append(key, subjectBytes...)
	key = append(key, []byte("/")...)

	return key
}
