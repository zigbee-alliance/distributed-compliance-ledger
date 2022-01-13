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

import (
	"encoding/binary"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ binary.ByteOrder

const (
	// ValidatorKeyPrefix is the prefix to retrieve all Validator.
	ValidatorKeyPrefix = "Validator/value/"
	// ValidatorByConsAddrKeyPrefix is the prefix to retrieve all Validator by consensus address.
	ValidatorByConsAddrKeyPrefix = "ValidatorByConsAddr/value/"
)

// ValidatorKey returns the store key to retrieve a Validator from the index fields.
func ValidatorKey(
	owner sdk.ValAddress,
) []byte {
	var key []byte

	ownerBytes := []byte(owner)
	key = append(key, ownerBytes...)
	key = append(key, []byte("/")...)

	return key
}

func ValidatorByConsAddrKey(
	addr sdk.ConsAddress,
) []byte {
	var key []byte

	ownerBytes := []byte(addr)
	key = append(key, ownerBytes...)
	key = append(key, []byte("/")...)

	return key
}
