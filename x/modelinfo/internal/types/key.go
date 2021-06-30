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
	ModuleName = "modelinfo"

	// StoreKey to be used when creating the KVStore.
	StoreKey = ModuleName
)

var (
	ModelInfoPrefix      = []byte{0x01} // prefix for each key to a model info
	ModelVersionPrefix   = []byte{0x02} // prefix for each key to a model with sv and hv
	VendorProductsPrefix = []byte{0x03} // prefix for each key to a vendor products
)

// Key builder for Model Info (with versions).
func GetModelInfoKey(vid uint16, pid uint16, softwareVersion uint32, hardwareVersion uint32) []byte {
	var key []byte

	key = append(key, ModelInfoPrefix...)

	v := make([]byte, 2)
	binary.LittleEndian.PutUint16(v, vid)
	key = append(key, v...)

	p := make([]byte, 2)
	binary.LittleEndian.PutUint16(p, pid)
	key = append(key, p...)

	sv := make([]byte, 4)
	binary.LittleEndian.PutUint32(sv, softwareVersion)
	key = append(key, sv...)

	hv := make([]byte, 4)
	binary.LittleEndian.PutUint32(hv, hardwareVersion)
	key = append(key, hv...)

	return key
}

// Key builder for Model Info (without versions).
func GetProductKey(vid uint16, pid uint16) []byte {
	var key []byte

	key = append(key, ModelInfoPrefix...)

	v := make([]byte, 2)
	binary.LittleEndian.PutUint16(v, vid)
	key = append(key, v...)

	p := make([]byte, 2)
	binary.LittleEndian.PutUint16(p, pid)
	key = append(key, p...)

	return key
}

// Key builder for Vendor Products.
func GetVendorProductsKey(vid uint16) []byte {
	b := make([]byte, 2)
	binary.LittleEndian.PutUint16(b, vid)

	return append(VendorProductsPrefix, b...)
}
