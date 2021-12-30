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
