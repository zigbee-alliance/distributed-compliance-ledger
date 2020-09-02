package pagination

import (
	"fmt"
	"net/http"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/viper"
)

const (
	FlagSkip      = "skip"
	FlagSkipUsage = "amount of records to skip"
	FlagTake      = "take"
	FlagTakeUsage = "take of records to take"
)

// request Payload for a list query with pagination.
type PaginationParams struct {
	Skip int
	Take int
}

func NewPaginationParams(skip int, take int) PaginationParams {
	return PaginationParams{Skip: skip, Take: take}
}

func ParsePaginationParamsFromFlags() PaginationParams {
	return NewPaginationParams(
		viper.GetInt(FlagSkip),
		viper.GetInt(FlagTake),
	)
}

func ParsePaginationParamsFromRequest(r *http.Request) (PaginationParams, error) {
	skip := 0

	if str := r.FormValue("skip"); len(str) > 0 {
		val_, err := strconv.Atoi(str)
		if err != nil {
			return PaginationParams{}, error(sdk.ErrUnknownRequest(
				fmt.Sprintf("Invalid query parameter `skip`: Parsing Error: %v must be number", str)))
		}

		skip = val_
	}

	take := 0

	if str := r.FormValue("take"); len(str) > 0 {
		val_, err := strconv.Atoi(str)
		if err != nil {
			return PaginationParams{}, error(sdk.ErrUnknownRequest(
				fmt.Sprintf("Invalid query parameter `take`: Parsing Error: %v must be number", str)))
		}

		take = val_
	}

	return NewPaginationParams(skip, take), nil
}
