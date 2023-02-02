package keeper

// DONTCOVER

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	v4 "github.com/gridiron-zone/huddle/x/posts/legacy/v4"

	v3 "github.com/gridiron-zone/huddle/x/posts/legacy/v3"

	v2 "github.com/gridiron-zone/huddle/x/posts/legacy/v2"
	"github.com/gridiron-zone/huddle/x/posts/types"
)

// Migrator is a struct for handling in-place store migrations.
type Migrator struct {
	k  Keeper
	sk types.SubspacesKeeper
}

// NewMigrator returns a new Migrator
func NewMigrator(keeper Keeper, sk types.SubspacesKeeper) Migrator {
	return Migrator{
		k:  keeper,
		sk: sk,
	}
}

// Migrate1to2 migrates from version 1 to 2.
func (m Migrator) Migrate1to2(ctx sdk.Context) error {
	return v2.MigrateStore(ctx, m.k.storeKey, m.k.paramsSubspace, m.sk)
}

// Migrate2to3 migrates from version 2 to 3.
func (m Migrator) Migrate2to3(ctx sdk.Context) error {
	return v3.MigrateStore(ctx, m.k.storeKey, m.k.cdc)
}

// Migrate3to4 migrates from version 3 to 4.
func (m Migrator) Migrate3to4(ctx sdk.Context) error {
	return v4.MigrateStore(ctx, m.k.storeKey, m.k.cdc)
}
