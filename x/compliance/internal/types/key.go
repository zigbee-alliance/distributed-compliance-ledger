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

import (
	"encoding/binary"
)

const (
	// ModuleName is the name of the module.
	ModuleName = "compliance"

	// StoreKey to be used when creating the KVStore.
	StoreKey = ModuleName
)

var ComplianceInfoPrefix = []byte{0x01} // prefix for each key to a compliance info

// Key builder for Compliance Info.
func GetComplianceInfoKey(certificationType CertificationType, vid uint16, pid uint16, softwareVersion uint32) []byte {

	var key []byte

	key = append(key, ComplianceInfoPrefix...)
	key = append(key, []byte(certificationType)...)

	v := make([]byte, 2)
	binary.LittleEndian.PutUint16(v, vid)
	key = append(key, v...)

	p := make([]byte, 2)
	binary.LittleEndian.PutUint16(p, pid)
	key = append(key, p...)

	sv := make([]byte, 4)
	binary.LittleEndian.PutUint32(sv, softwareVersion)
	key = append(key, sv...)

	return key
}

// Key builder for Compliance Info.
func GetCertificationPrefix(certificationType CertificationType) []byte {
	return append(ComplianceInfoPrefix, []byte(certificationType)...)
}
