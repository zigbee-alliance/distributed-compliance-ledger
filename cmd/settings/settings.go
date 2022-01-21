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

package settings

import (
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/store/types"
)

const (
	// Default broadcast mode used for write transactions.
	DefaultBroadcastMode = flags.BroadcastBlock
)

// PruningStrategy of the application: Store every 100th state. Keep last 100 states. Do pruning at every 10th height.
var PruningStrategy = types.NewPruningOptions(100, 100, 10)
