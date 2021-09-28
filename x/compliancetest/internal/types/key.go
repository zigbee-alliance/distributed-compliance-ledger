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
	ModuleName = "compliancetest"

	// StoreKey to be used when creating the KVStore.
	StoreKey = ModuleName
)

var TestingResultsPrefix = []byte{0x01} // prefix for each key to a testing results

// Key builder for Testing Results.
func GetTestingResultsKey(vid uint16, pid uint16, softwareVersion uint32) []byte {

	var key []byte

	key = append(key, TestingResultsPrefix...)
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
