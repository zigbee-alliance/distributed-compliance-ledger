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
	"time"
)

const (
	// Default power every validator created with.
	Power int32 = 10

	// Zero power is used to demote validator.
	ZeroPower int32 = 0

	// Maximum number of nodes.
	MaxNodes = 100

	// Maximum time to accept double-sign evidence.
	MaxEvidenceAge = 60 * 2 * time.Second

	// Size (number of blocks) of the sliding window used to track validator liveness.
	SignedBlocksWindow int64 = 100

	// Minimal number of blocks must have been signed per window.
	MinSignedPerWindow int64 = 50
)
