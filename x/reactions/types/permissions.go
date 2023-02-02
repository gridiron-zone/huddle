package types

// DONTCOVER

import (
	subspacestypes "github.com/gridiron-zone/huddle/x/subspaces/types"
)

var (
	PermissionsReact                    = subspacestypes.RegisterPermission("add reaction")
	PermissionManageRegisteredReactions = subspacestypes.RegisterPermission("manage registered reactions")
	PermissionManageReactionParams      = subspacestypes.RegisterPermission("manage reaction params")
)
