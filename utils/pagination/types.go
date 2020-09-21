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

package pagination

import (
	"fmt"
	"net/http"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
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

func AddPaginationParams(cmd *cobra.Command) {
	cmd.Flags().Int(FlagSkip, 0, FlagSkipUsage)
	cmd.Flags().Int(FlagTake, 0, FlagTakeUsage)
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

const (
	FlagStartKey      = "start-key"
	FlagStartKeyUsage = "utf-8 str to start iteration from"
	FlagEndKey        = "end-key"
	FlagEndKeyUsage   = "utf-8 str to stop iteration at"
	FlagLimit         = "limit"
	FlagLimitUsage    = "maximum number of records"
)

// request Payload for a range query.
type RangeParams struct {
	StartKey []byte
	EndKey   []byte
	Limit    int
}

func NewRangeParams(startKey []byte, endKey []byte, limit int) RangeParams {
	return RangeParams{StartKey: startKey, EndKey: endKey, Limit: limit}
}

func AddRangeParams(cmd *cobra.Command) {
	cmd.Flags().String(FlagStartKey, "", FlagStartKeyUsage)
	cmd.Flags().String(FlagEndKey, "", FlagEndKeyUsage)
	cmd.Flags().Int(FlagLimit, 0, FlagLimitUsage)
}

func ParseRangeParamsFromFlags() (RangeParams, error) {
	return NewRangeParams(
		[]byte(viper.GetString(FlagStartKey)),
		[]byte(viper.GetString(FlagEndKey)),
		viper.GetInt(FlagLimit),
	), nil
}

// nolint: godox
// TODO
//func ParsePaginationParamsFromRequest(r *http.Request) (PaginationParams, error) {
//	skip := 0
//
//	if str := r.FormValue("skip"); len(str) > 0 {
//		val_, err := strconv.Atoi(str)
//		if err != nil {
//			return PaginationParams{}, error(sdk.ErrUnknownRequest(
//				fmt.Sprintf("Invalid query parameter `skip`: Parsing Error: %v must be number", str)))
//		}
//
//		skip = val_
//	}
//
//	take := 0
//
//	if str := r.FormValue("take"); len(str) > 0 {
//		val_, err := strconv.Atoi(str)
//		if err != nil {
//			return PaginationParams{}, error(sdk.ErrUnknownRequest(
//				fmt.Sprintf("Invalid query parameter `take`: Parsing Error: %v must be number", str)))
//		}
//
//		take = val_
//	}
//
//	return NewPaginationParams(skip, take), nil
//}
