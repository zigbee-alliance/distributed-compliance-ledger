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
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

// Default parameter values.
const (
	DclMaxMemoCharacters          uint64  = authtypes.DefaultMaxMemoCharacters
	DclTxSigLimit                 uint64  = authtypes.DefaultTxSigLimit
	DclTxSizeCostPerByte          uint64  = 0 // gas is not needed in DCL
	DclSigVerifyCostED25519       uint64  = 0 // gas is not needed in DCL
	DclSigVerifyCostSecp256k1     uint64  = 0 // gas is not needed in DCL
	AccountApprovalsPercent       float64 = 0.66
	VendorAccountApprovalsPercent float64 = 0.33
)
