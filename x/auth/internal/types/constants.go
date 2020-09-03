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

// Default parameter values.
const (
	MaxMemoCharacters             uint64  = 256
	TxSizeCostPerByte             uint64  = 10
	DefaultSigVerifyCostED25519   uint64  = 590
	DefaultSigVerifyCostSecp256k1 uint64  = 1000
	AccountApprovalPercent        float64 = 0.66
)
