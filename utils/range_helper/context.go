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

// nolint: stylecheck
package range_helper

import "github.com/cosmos/cosmos-sdk/store/iavl"

type Context interface {
	QueryRange(startKey []byte, endKey []byte, limit int, storeName string) (iavl.RangeRes, int64, error)
	QueryStoreAtHeight(height int64, key []byte, storeName string) ([]byte, int64, error)
}
