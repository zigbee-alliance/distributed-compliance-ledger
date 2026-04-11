package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Migrator is a struct for handling in-place store migrations.
type Migrator struct {
	keeper Keeper
}

// NewMigrator returns a new Migrator.
func NewMigrator(keeper Keeper) Migrator {
	return Migrator{keeper: keeper}
}

// Migrate3to4 migrates from version 3 to 4.
func (m Migrator) Migrate3to4(ctx sdk.Context) error {
	models := m.keeper.GetAllModel(ctx)
	for _, model := range models {
		if model.GetCommissioningModeSecondaryStepsHint() != 1 {
			continue
		}
		model.CommissioningModeSecondaryStepsHint = 4
		m.keeper.SetModel(ctx, model)
	}

	return nil
}

// Migrate4to5 migrates from version 4 to 5.
func (m Migrator) Migrate4to5(ctx sdk.Context) error {
	modelVersions := m.keeper.GetAllModelVersions(ctx)
	for _, modelVersion := range modelVersions {
		for _, softwareVersion := range modelVersion.SoftwareVersions {
			_, found := m.keeper.GetModelVersion(ctx, modelVersion.Vid, modelVersion.Pid, softwareVersion)
			if !found {
				m.keeper.RemoveSoftwareVersion(ctx, modelVersion.Vid, modelVersion.Pid, softwareVersion)
			}
		}
	}

	return nil
}
