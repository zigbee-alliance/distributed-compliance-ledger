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
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Response Payload for a list query with pagination.
type ListModelInfoItems struct {
	Total int             `json:"total"`
	Items []ModelInfoItem `json:"items"`
}

// Implement fmt.Stringer.
func (n ListModelInfoItems) String() string {
	res, err := json.Marshal(n)
	if err != nil {
		panic(err)
	}

	return string(res)
}

type ModelInfoItem struct {
	VID   uint16         `json:"vid"`
	PID   uint16         `json:"pid"`
	Name  string         `json:"name"`
	SKU   string         `json:"sku"`
	Owner sdk.AccAddress `json:"owner"`
}

// Response Payload for a list query with pagination.
type ListVendorItems struct {
	Total int          `json:"total"`
	Items []VendorItem `json:"items"`
}

// Implement fmt.Stringer.
func (n ListVendorItems) String() string {
	res, err := json.Marshal(n)
	if err != nil {
		panic(err)
	}

	return string(res)
}

type VendorItem struct {
	VID uint16 `json:"vid"`
}
