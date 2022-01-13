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

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	compliancetesttypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/compliancetest/types"
	dclauthtypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
	modeltypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/model/types"
)

type DclauthKeeper interface {
	// Methods imported from dclauth should be defined here

	HasRole(ctx sdk.Context, addr sdk.AccAddress, roleToCheck dclauthtypes.AccountRole) bool
}

type ModelKeeper interface {
	// Methods imported from model should be defined here

	GetModelVersion(ctx sdk.Context, vid int32, pid int32, softwareVersion uint32) (val modeltypes.ModelVersion, found bool)
}

type CompliancetestKeeper interface {
	// Methods imported from compliancetest should be defined here

	GetTestingResults(ctx sdk.Context, vid int32, pid int32, softwareVersion uint32) (val compliancetesttypes.TestingResults, found bool)
}
