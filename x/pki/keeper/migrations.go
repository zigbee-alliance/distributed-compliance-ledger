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

// Migrate1to2 migrates from version 1 to 2.
func (m Migrator) Migrate1to2(_ sdk.Context) error { return nil }

// Migrate2to3 migrates from version 2 to 3.
func (m Migrator) Migrate2to3(ctx sdk.Context) error {
	approvedCertificates := m.keeper.GetAllApprovedCertificates(ctx)
	for _, cert := range approvedCertificates {
		m.keeper.AddApprovedCertificatesBySubjectKeyID(ctx, cert)
	}

	return nil
}

// Migrate3to4 migrates from version 3 to 4.
func (m Migrator) Migrate3to4(ctx sdk.Context) error {
	approvedCertificates := m.keeper.GetAllApprovedCertificates(ctx)
	for _, cert := range approvedCertificates {
		m.keeper.AddAllCertificates(ctx, cert.Subject, cert.SubjectKeyId, cert.SchemaVersion, cert.Certs)
	}

	return nil
}

// Migrate2to4 migrates from version 2 to 4.
func (m Migrator) Migrate2to4(ctx sdk.Context) error {
	err := m.Migrate2to3(ctx)
	if err != nil {
		return err
	}
	err = m.Migrate3to4(ctx)
	if err != nil {
		return err
	}
	return nil
}
