// Copyright 2022 DSR Corporation
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

const (
	// ModuleName defines the module name.
	ModuleName = "dclauth"

	// ModuleName defines the module name to use in user interactions.
	ModuleNameUser = "auth"

	// command name for the module.
	CmdName = ModuleNameUser

	// StoreKey defines the primary module store key.
	StoreKey = ModuleName

	// RouterKey is the message route for slashing.
	RouterKey = ModuleNameUser

	// QuerierRoute defines the module's query routing key.
	QuerierRoute = ModuleNameUser

	// MemStoreKey defines the in-memory store key.
	MemStoreKey = "mem_" + ModuleName
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	AccountStatKey = "AccountStat-value-"
)
