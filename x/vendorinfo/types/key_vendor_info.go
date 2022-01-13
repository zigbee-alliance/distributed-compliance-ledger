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
	// VendorInfoKeyPrefix is the prefix to retrieve all VendorInfo.
	VendorInfoKeyPrefix = "VendorInfo/value/"
)

// VendorInfoKey returns the store key to retrieve a VendorInfo from the index fields.
func VendorInfoKey(
	vendorID int32,
) []byte {
	var key []byte

	vendorIDBytes := make([]byte, 8)
	binary.BigEndian.PutUint32(vendorIDBytes, uint32(vendorID))
	key = append(key, vendorIDBytes...)
	key = append(key, []byte("/")...)

	return key
}
