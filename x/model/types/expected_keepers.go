package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliance/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
)

type DclauthKeeper interface {
	// Methods imported from dclauth should be defined here
	HasRole(ctx sdk.Context, addr sdk.AccAddress, roleToCheck types.AccountRole) bool
	HasVendorID(ctx sdk.Context, addr sdk.AccAddress, vid int32) bool
}

type ComplianceKeeper interface {
	// Methods imported from compliance should be defined here
	GetComplianceInfo(
		ctx sdk.Context,
		vid int32,
		pid int32,
		softwareVersion uint32,
		certificationType string,
	) (val types.ComplianceInfo, found bool)

	GetProvisionalModel(
		ctx sdk.Context,
		vid int32,
		pid int32,
		softwareVersion uint32,
		certificationType string,
	) (val types.ProvisionalModel, found bool)

	GetRevokedModel(
		ctx sdk.Context,
		vid int32,
		pid int32,
		softwareVersion uint32,
		certificationType string,
	) (val types.RevokedModel, found bool)
}
