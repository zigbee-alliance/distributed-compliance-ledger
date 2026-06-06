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

package cli

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	commontypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/common/types"
)

func TestAccount_getPidRanges(t *testing.T) {
	tests := []struct {
		name      string
		pidRanges string
		want      []*commontypes.Uint16Range
		err       error
	}{
		{
			name:      "get ProductID ranges from \"1-100\"",
			pidRanges: "1-100",
			want:      []*commontypes.Uint16Range{{Min: 1, Max: 100}},
			err:       nil,
		},
		{
			name:      "get ProductID ranges from \"1-100,200-300\"",
			pidRanges: "1-100,200-300",
			want:      []*commontypes.Uint16Range{{Min: 1, Max: 100}, {Min: 200, Max: 300}},
			err:       nil,
		},
		{
			name:      "get ProductID ranges from \"100-100\"",
			pidRanges: "100-100",
			want:      []*commontypes.Uint16Range{{Min: 100, Max: 100}},
			err:       nil,
		},
		{
			name:      "get ProductID ranges from \"1-100-200\"",
			pidRanges: "1-100-200",
			want:      nil,
			err:       fmt.Errorf("failed to parse PID Range"),
		},
		{
			name:      "get ProductID ranges from \"1-10a,200-300\"",
			pidRanges: "1-10a,100-200",
			want:      nil,
			err:       fmt.Errorf("failed to parse PID Range"),
		},
		{
			name:      "get ProductID ranges from \"1O-100\"",
			pidRanges: "1O-100",
			want:      nil,
			err:       fmt.Errorf("failed to parse PID Range"),
		},
		{
			name:      "get ProductID ranges from \"1-10O,101-200\"",
			pidRanges: "1-10O,101-200",
			want:      nil,
			err:       fmt.Errorf("failed to parse PID Range"),
		},
		{
			name:      "get ProductID ranges from \"0-100,200-300\"",
			pidRanges: "0-100,200-300",
			want:      nil,
			err:       fmt.Errorf("invalid PID Range is provided"),
		},
		{
			name:      "get ProductID ranges from \"1-100,100-200\"",
			pidRanges: "1-100,100-200",
			want:      nil,
			err:       fmt.Errorf("invalid PID Range is provided"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, err := getPidRanges(tt.pidRanges); err != nil && tt.err == nil {
				t.Errorf("getPidRanges(%s) = %v, want %v", tt.pidRanges, got, tt.want)
			} else {
				require.Equal(t, tt.want, got)
			}
		})
	}
}
