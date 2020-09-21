package cli

import (
	"encoding/json"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/bytes"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/pagination"
)

type RangeResult struct {
	Total int               `json:"total"`
	Items []json.RawMessage `json:"items"`
}

// Takes into account cli flags, validates total, handles output.
func (ctx CliContext) QueryRangeWithToatalAndHandleCLIIO(storeKey string,
	prefix []byte, totalKey []byte, valueUnmarshaler func([]byte) json.RawMessage) error {
	params, err := pagination.ParseRangeParamsFromFlags()
	if err != nil {
		return err
	}

	result, height, err := ctx.QueryRangeWithTotal(storeKey, prefix, params, totalKey, valueUnmarshaler)
	if err != nil {
		return err
	}

	return ctx.PrintWithHeight(ctx.Codec().MustMarshalJSON(result), height)
}

func (ctx CliContext) QueryRangeWithTotal(storeKey string, prefix []byte, params pagination.RangeParams,
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
	totalRes, totalHeight, err := ctx.context.WithHeight(rangeHeight).QueryStore(totalKey, storeKey)
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
		ctx.Codec().MustUnmarshalBinaryLengthPrefixed(totalRes, &total)
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
