package pagination

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/viper"
	"net/http"
	"strconv"
)

const (
	FlagSkip = "skip"
	FlagTake = "take"
)

// Request Payload for a list query with pagination
type PaginationParams struct {
	Skip int
	Take int
}

func NewPaginationParams(skip int, take int) PaginationParams {
	return PaginationParams{Skip: skip, Take: take}
}

func ParsePaginationParamsFromFlags(cdc *codec.Codec) []byte {
	params := NewPaginationParams(
		viper.GetInt(FlagSkip),
		viper.GetInt(FlagTake),
	)
	return cdc.MustMarshalJSON(params)
}

func ParsePaginationParamsFromRequest(cdc *codec.Codec, r *http.Request) ([]byte, error) {
	skip := 0
	if str := r.FormValue("skip"); len(str) > 0 {
		val_, err := strconv.Atoi(str)
		if err != nil {
			return nil, error(sdk.ErrInternal(fmt.Sprintf("Invalid query parameter `skip`: Parsing Error: %v must be number", str)))
		}
		skip = val_
	}

	take := 0
	if str := r.FormValue("take"); len(str) > 0 {
		val_, err := strconv.Atoi(str)
		if err != nil {
			return nil, error(sdk.ErrInternal(fmt.Sprintf("Invalid query parameter `take`: Parsing Error: %v must be number", str)))
		}
		take = val_
	}

	params := NewPaginationParams(skip, take)
	return cdc.MustMarshalJSON(params), nil
}
