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

import (
	"encoding/json"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/bytes"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/pagination"
)

type RangeResult struct {
	Total int               `json:"total"`
	Items []json.RawMessage `json:"items"`
}

func QueryRangeWithTotal(ctx Context, storeKey string, prefix []byte, params pagination.RangeParams,
	totalKey []byte, valueUnmarshaler func([]byte) json.RawMessage) (RangeResult, int64, error) {
	// Is all range queried
	isAllRange := false

	startKey := append(prefix, params.StartKey...)

	if !bytes.IsAllZeroes(params.StartKey) {
		isAllRange = false
	}

	var endKey []byte
	if len(params.EndKey) == 0 {
		endKey = bytes.CpIncr(prefix)
	} else {
		// nolint: gocritic
		endKey = append(prefix, params.EndKey...)
		isAllRange = false
	}

	// Query range
	rangeRes, rangeHeight, err := ctx.QueryRange(startKey, endKey, params.Limit, storeKey)
	if err != nil {
		return RangeResult{}, 0, sdk.ErrInternal(fmt.Sprintf("Could not get data: %s\n", err))
	}

	// Query total number of items at the same height
	totalRes, totalHeight, err := ctx.QueryStoreAtHeight(rangeHeight, totalKey, storeKey)
	if err != nil {
		return RangeResult{}, 0, sdk.ErrInternal(fmt.Sprintf("Could not get data: %s\n", err))
	}

	if rangeHeight != totalHeight {
		panic("should not happen")
	}

	var total int
	if totalRes == nil {
		total = 0
	} else {
		codec.Cdc.MustUnmarshalBinaryLengthPrefixed(totalRes, &total)
	}

	// Compare length of response with the PROVED total number if all range is requested
	if isAllRange && len(rangeRes.Values) != total {
		return RangeResult{}, 0, sdk.ErrInternal(
			fmt.Sprintf("Response length doesn't match value stored in totoal key: %s\n", err))
	}

	// Convert result to json
	result := RangeResult{
		total,
		make([]json.RawMessage, 0, len(rangeRes.Values)),
	}

	for _, valueBytes := range rangeRes.Values {
		result.Items = append(result.Items, valueUnmarshaler(valueBytes))
	}

	return result, rangeHeight, nil
}
